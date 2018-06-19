package x

import (
	"sync"
)

type lazyReplyTag uint

const (
	lazyNone lazyReplyTag = iota
	lazyCookie
	lazyForced
)

type Extension struct {
	name     string // extension-xname
	globalID int    // get from query extension request
}

func (ext *Extension) Name() string {
	return ext.name
}

func NewExtension(name string) *Extension {
	return &Extension{
		name: name,
	}
}

type lazyReply struct {
	tag    lazyReplyTag
	reply  *QueryExtensionReply
	cookie QueryExtensionCookie
}

type exts struct {
	mu      sync.Mutex
	replies []lazyReply
}

func newExts() *exts {
	return &exts{
		replies: make([]lazyReply, 3),
	}
}

func (exts *exts) grow(n int) {
	if n <= len(exts.replies) {
		return
	}

	//Logger.Println("exts grow", len(exts.replies), " -> ", n)
	bigger := make([]lazyReply, n)
	copy(bigger, exts.replies)
	exts.replies = bigger
}

func (exts *exts) getByIndex(idx int) *lazyReply {
	// idx is extension globalID
	if idx > len(exts.replies) {
		exts.grow(idx * 2)
	}

	return &exts.replies[idx-1]
}

var nextGlobalID int
var globalIDMu sync.Mutex

func (exts *exts) getLazyReply(conn *Conn, ext *Extension) (lzr *lazyReply) {
	globalIDMu.Lock()
	if ext.globalID == 0 {
		nextGlobalID++
		ext.globalID = nextGlobalID
		//Logger.Println("ext", ext.name, "gloabl ID", ext.globalID)
	}
	globalIDMu.Unlock()

	lzr = exts.getByIndex(ext.globalID)
	// lazyNone -> lazyCookie
	if lzr.tag == lazyNone {
		lzr.tag = lazyCookie
		lzr.cookie = QueryExtension(conn, ext.name)
	}
	return
}

func (c *Conn) GetExtensionData(ext *Extension) (data *QueryExtensionReply) {
	c.exts.mu.Lock()
	lzr := c.exts.getLazyReply(c, ext)

	// lazyCookie -> lazyForced
	if lzr.tag == lazyCookie {
		lzr.tag = lazyForced
		lzr.reply, _ = lzr.cookie.Reply(c)
	}
	data = lzr.reply
	c.exts.mu.Unlock()
	return
}

func (c *Conn) PrefetchExtensionData(ext *Extension) {
	c.exts.mu.Lock()
	c.exts.getLazyReply(c, ext)
	c.exts.mu.Unlock()
}
