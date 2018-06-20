package x

import (
	"bufio"
	"log"
	"net"
	"os"
	"sync"
)

var Logger *log.Logger
var debugEnabled bool

func init() {
	if os.Getenv("DEBUG_X11_CLIENT") == "1" {
		debugEnabled = true
		Logger = log.New(os.Stderr, "[x] ", log.Lshortfile)
	}
}

func logPrintln(v ...interface{}) {
	if debugEnabled {
		Logger.Println(v...)
	}
}
func logPrintf(format string, v ...interface{}) {
	if debugEnabled {
		Logger.Printf(format, v...)
	}
}

type Conn struct {
	conn      net.Conn
	bufReader *bufio.Reader

	host          string
	display       string
	DisplayNumber int
	ScreenNumber  int
	setup         *Setup

	ioMu sync.Mutex
	in   *in
	out  *out

	ext    ext
	xidGen *xidGenerator
}

func (c *Conn) GetSetup() *Setup {
	return c.setup
}

func (c *Conn) GetDefaultScreen() *Screen {
	return &c.setup.Roots[c.ScreenNumber]
}

func (c *Conn) Close() {
	c.conn.Close()
	c.in.close()
}

func (c *Conn) AddEventChan(eventChan chan<- GenericEvent) {
	c.in.addEventChan(eventChan)
}

func (c *Conn) RemoveEventChan(eventChan chan<- GenericEvent) {
	c.in.removeEventChan(eventChan)
}

func (c *Conn) AddErrorChan(errorChan chan<- Error) {
	c.in.addErrorChan(errorChan)
}

func (c *Conn) RemoveErrorChan(errorChan chan<- Error) {
	c.in.removeErrorChan(errorChan)
}
