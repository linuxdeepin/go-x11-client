package x

import (
	"container/list"
	"sync"
)

type In struct {
	requestExpected  uint64
	requestRead      uint64
	requestCompleted uint64

	currentReply   *list.List
	replies        map[uint64]*list.List
	pendingReplies *list.List

	readers *list.List

	eventChans []chan<- GenericEvent
	errorChans []chan<- Error
	chansMu    sync.Mutex
}

func newIn() *In {
	in := &In{}
	in.replies = make(map[uint64]*list.List)
	in.pendingReplies = list.New()
	in.readers = list.New()
	return in
}

func (in *In) addEventChan(eventChan chan<- GenericEvent) {
	in.chansMu.Lock()

	for _, ch := range in.eventChans {
		if ch == eventChan {
			// exist
			in.chansMu.Unlock()
			return
		}
	}

	in.eventChans = append(in.eventChans, eventChan)
	in.chansMu.Unlock()
}

func (in *In) removeEventChan(eventChan chan<- GenericEvent) {
	in.chansMu.Lock()

	chans := in.eventChans
	idx := -1
	for i, ch := range chans {
		if ch == eventChan {
			idx = i
			break
		}
	}

	if idx == -1 {
		// not found
		in.chansMu.Unlock()
		return
	}

	chans[idx] = chans[len(chans)-1]
	chans[len(chans)-1] = nil
	chans = chans[:len(chans)-1]

	in.chansMu.Unlock()
}

func (in *In) addErrorChan(errorChan chan<- Error) {
	in.chansMu.Lock()

	for _, ch := range in.errorChans {
		if ch == errorChan {
			// exist
			in.chansMu.Unlock()
			return
		}
	}

	in.errorChans = append(in.errorChans, errorChan)
	in.chansMu.Unlock()
}

func (in *In) removeErrorChan(errorChan chan<- Error) {
	in.chansMu.Lock()

	chans := in.errorChans
	idx := -1
	for i, ch := range chans {
		if ch == errorChan {
			idx = i
			break
		}
	}

	if idx == -1 {
		// not found
		in.chansMu.Unlock()
		return
	}

	chans[idx] = chans[len(chans)-1]
	chans[len(chans)-1] = nil
	chans = chans[:len(chans)-1]

	in.chansMu.Unlock()
}

func (in *In) addEvent(e GenericEvent) {
	Logger.Println("add event", e)
	in.chansMu.Lock()

	for _, ch := range in.eventChans {
		ch <- e
	}

	in.chansMu.Unlock()
}

func (in *In) addError(e Error) {
	in.chansMu.Lock()

	for _, ch := range in.errorChans {
		select {
		case ch <- e:
		default:
		}
	}

	in.chansMu.Unlock()
}

func (in *In) close() {
	in.chansMu.Lock()

	for _, ch := range in.eventChans {
		close(ch)
	}

	for _, ch := range in.errorChans {
		close(ch)
	}

	in.chansMu.Unlock()
}

type ReplyReader struct {
	request uint64
	cond    *sync.Cond
}

func (in *In) insertNewReader(request uint64, cond *sync.Cond) *ReplyReader {
	r := &ReplyReader{
		request: request,
		cond:    cond,
	}

	var mark *list.Element
	l := in.readers
	for e := l.Front(); e != nil; e = e.Next() {
		reader := e.Value.(*ReplyReader)
		if reader.request >= request {
			mark = e
			break
		}
	}

	if mark != nil {
		Logger.Printf("insertNewReader %d before %d\n", request, mark.Value.(*ReplyReader).request)
		l.InsertBefore(r, mark)
	} else {
		Logger.Printf("insertNewReader %d at end\n", request)
		l.PushBack(r)
	}
	return r
}

func (in *In) removeReader(r *ReplyReader) {
	l := in.readers
	for e := l.Front(); e != nil; e = e.Next() {
		reader := e.Value.(*ReplyReader)
		if reader.request == r.request {
			Logger.Println("remove reader", reader.request)
			l.Remove(e)
			break
		}
	}
}

func (in *In) removeFinishedReaders() {
	l := in.readers
	e := l.Front()
	for e != nil {
		reader := e.Value.(*ReplyReader)
		if reader.request <= in.requestCompleted {
			reader.cond.Signal()
			Logger.Println("remove finshed reader", reader.request)
			tmp := e
			e = e.Next()
			l.Remove(tmp)
		} else {
			break
		}
	}
}

func (in *In) wakeUpNextReader() {
	if in.readers.Front() != nil {
		reader := in.readers.Front().Value.(*ReplyReader)
		Logger.Println("wake up next reader", reader.request)
		reader.cond.Signal()
	}
}

type PendingReply struct {
	firstRequest uint64
	lastRequest  uint64
	workaround   uint
	flags        uint
}

func (in *In) expectReply(request uint64, workaround uint, flags uint) {
	pend := &PendingReply{
		firstRequest: request,
		lastRequest:  request,
		workaround:   workaround,
		flags:        flags,
	}
	in.pendingReplies.PushBack(pend)
}

func (in *In) removeFinishedPendingReplies() {
	l := in.pendingReplies
	e := l.Front()
	for e != nil {
		pend := e.Value.(*PendingReply)
		if pend.lastRequest <= in.requestCompleted {
			// remove pend from list
			tmp := e
			e = e.Next()
			l.Remove(tmp)
		} else {
			break
		}
	}
}
