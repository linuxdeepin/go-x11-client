package client

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

	eventCh chan GenericEvent
	errorCh chan *GenericError
}

const (
	eventChanBufferSize = 5000
	errorChanBufferSize = 1000
)

func NewIn(mu *sync.Mutex) *In {
	in := &In{}
	in.replies = make(map[uint64]*list.List)
	in.pendingReplies = list.New()
	in.readers = list.New()

	in.eventCh = make(chan GenericEvent, eventChanBufferSize)
	in.errorCh = make(chan *GenericError, errorChanBufferSize)
	return in
}

func (in *In) addEvent(e GenericEvent) {
	Logger.Println("add event", e)
	in.eventCh <- e
}

func (in *In) addError(e *GenericError) bool {
	select {
	case in.errorCh <- e:
		Logger.Println("add error", e.Error())
		return true

	default:
		Logger.Println("error chan buffer full")
		return false
	}
}

func (in *In) close() {
	close(in.errorCh)
	close(in.eventCh)
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

type InExported struct {
	RequestExpected  uint64
	RequestRead      uint64
	RequestCompleted uint64

	CurrentReply *list.List
	Replies      map[uint64]*list.List

	Readers *list.List
}

func (in *In) Export() *InExported {
	return &InExported{
		RequestExpected:  in.requestExpected,
		RequestRead:      in.requestRead,
		RequestCompleted: in.requestCompleted,

		CurrentReply: in.currentReply,
		Replies:      in.replies,
		Readers:      in.readers,
	}
}
