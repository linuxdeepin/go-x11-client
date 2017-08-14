package client

import (
	"container/list"
	"io"
	"sync"
)

const (
	ResponseTypeError = 0
	ResponseTypeReply = 1
)

//enum xcb_send_request_flags_t {
//XCB_REQUEST_CHECKED = 1 << 0,
//XCB_REQUEST_RAW = 1 << 1,
//XCB_REQUEST_DISCARD_REPLY = 1 << 2,
//XCB_REQUEST_REPLY_FDS = 1 << 3
//};
const (
	RequestChecked      = 1 << 0
	RequestRaw          = 1 << 1 // ?
	RequestDiscardReply = 1 << 2
	RequestReplyFds     = 1 << 3 // ?
)

//typedef struct {
//uint8_t   response_type;  [>*< Type of the response <]
//uint8_t  pad0;           [>*< Padding <]
//uint16_t sequence;       [>*< Sequence number <]
//uint32_t length;         [>*< Length of the response <]
//} xcb_generic_reply_t;

type GenericReply struct {
	responseType uint8
	sequence     uint16
	length       uint32
}

func NewGenericReply(buf []byte) *GenericReply {
	r := &GenericReply{}
	r.responseType = buf[0]
	// skip pad buf[1]
	r.sequence = Get16(buf[2:])
	r.length = Get32(buf[4:])
	return r
}

func (c *Conn) DebugIn() *InExported {
	return c.in.Export()
}

// go c.readLoop
func (c *Conn) readLoop() {
	for {
		err := c.readPacket()
		if err != nil {
			return
		}
	}
}

func (c *Conn) readPacket() error {
	length := 32
	buf := make([]byte, length)
	_, err := io.ReadFull(c.conn, buf)
	if err != nil {
		return err
	}
	genrep := NewGenericReply(buf)
	Logger.Printf("genrep: %#v\n", genrep)

	if genrep.responseType == ResponseTypeReply && genrep.length > 0 {
		length += int(genrep.length) * 4
		// grow buf
		biggerBuf := make([]byte, length)
		copy(biggerBuf[:32], buf)
		_, err = io.ReadFull(c.conn, biggerBuf[32:])
		if err != nil {
			return err
		}
		buf = biggerBuf
	}

	c.iolock.Lock()
	defer c.iolock.Unlock()
	// update c.in.request*
	if genrep.responseType != 11 /*XCB_KEYMAP_NOTIFY*/ {
		lastRead := c.in.requestRead
		c.in.requestRead = (lastRead & 0xffffffffffff0000) | uint64(genrep.sequence)
		if c.in.requestRead < lastRead {
			c.in.requestRead += 0x10000
		}

		if c.in.requestRead > c.in.requestExpected {
			c.in.requestExpected = c.in.requestRead
		}

		if c.in.requestRead != lastRead {
			curReply := c.in.currentReply
			if curReply != nil && curReply.Len() != 0 {
				c.in.replies[lastRead] = curReply
				c.in.currentReply = nil
			}
			c.in.requestCompleted = c.in.requestRead - 1
		}

		c.in.removeFinishedPendingReplies()

		if genrep.responseType == ResponseTypeError {
			c.in.requestCompleted = c.in.requestRead
		}
		c.in.removeFinishedReaders()
	}

	var pend *PendingReply
	if genrep.responseType == ResponseTypeReply ||
		genrep.responseType == ResponseTypeError {

		if prFront := c.in.pendingReplies.Front(); prFront != nil {
			pend = prFront.Value.(*PendingReply)

			if !(pend.firstRequest <= c.in.requestRead &&
				c.in.requestRead <= pend.lastRequest) {
				pend = nil
			}
		}
	}

	Logger.Printf("pend %#v\n", pend)
	if pend != nil && pend.flags&RequestDiscardReply != 0 {
		// discard reply
		Logger.Println("discard reply", c.in.requestRead)
		return nil
	}

	/* reply, or checked error */
	if genrep.responseType == ResponseTypeReply ||
		(genrep.responseType == ResponseTypeError &&
			pend != nil && pend.flags&RequestChecked != 0) {

		if c.in.currentReply == nil {
			c.in.currentReply = list.New()
		}
		c.in.currentReply.PushBack(buf)
		Logger.Printf("pushBack buf %d len=%d\n", c.in.requestRead, len(buf))

		front := c.in.readers.Front()
		if front != nil {
			reader := front.Value.(*ReplyReader)
			if reader.request == c.in.requestRead {
				Logger.Printf("readPacket reader %d signal\n", reader.request)
				reader.cond.Signal()
			}
		}
		return nil
	}

	/* event, or unchecked error */
	// not special event
	if genrep.responseType == ResponseTypeError {
		// is unchecked error
		c.in.addError(NewGenericError(buf))
	} else {
		// is event
		c.in.addEvent(GenericEvent(buf))
	}

	return nil
}

func (c *Conn) pollForReply(request uint64) (replyBuf []byte, isErr, stop bool) {
	Logger.Println("pollForReply", request)
	var front *list.Element
	if request == 0 {
		//replyBuf = nil
		//isErr = false
		stop = true
		return
	} else if request < c.in.requestRead {
		/* We've read requests past the one we want, so if it has replies we have
		 * them all and they're in the replies map. */
		l := c.in.replies[request]
		Logger.Println("reply in replies map")
		if l != nil {
			// pop front
			front = l.Front()
			if front != nil {
				l.Remove(front)
			}
			if l.Len() == 0 {
				delete(c.in.replies, request)
			}
		}
	} else if request == c.in.requestRead && c.in.currentReply != nil && c.in.currentReply.Front() != nil {
		/* We're currently processing the responses to the request we want, and we
		 * have a reply ready to return. So just return it without blocking. */
		Logger.Println("reply in currentReply")
		front = c.in.currentReply.Front()
		c.in.currentReply.Remove(front)

	} else if request == c.in.requestCompleted {
		/* We know this request can't have any more replies, and we've already
		 * established it doesn't have a reply now. Don't bother blocking. */
		Logger.Println("request completed")
		stop = true
		return
	} else {
		/* We may have more replies on the way for this request: block until we're sure. */
		// stop = false
		return
	}

	if front != nil {
		replyBuf = front.Value.([]byte)
		respType := replyBuf[0]
		if respType == ResponseTypeError {
			isErr = true
		}
	}
	stop = true
	return
}

func (c *Conn) waitForReply(request uint64) (replyBuf []byte, isErr bool) {
	var stop bool
	cond := sync.NewCond(&c.iolock)
	r := c.in.insertNewReader(request, cond)
	for {
		replyBuf, isErr, stop = c.pollForReply(request)
		if stop {
			break
		}

		Logger.Printf("waitForReply reader %d wait\n", request)
		r.cond.Wait()
	}
	c.in.removeReader(r)
	c.in.wakeUpNextReader()
	return
}

func (c *Conn) WaitForReply(request uint64) (replyBuf []byte, isErr bool) {
	c.iolock.Lock()
	replyBuf, isErr = c.waitForReply(request)
	c.iolock.Unlock()
	return
}

func (c *Conn) PollForReply(request uint64) (replyBuf []byte, isErr, stop bool) {
	c.iolock.Lock()
	replyBuf, isErr, stop = c.pollForReply(request)
	c.iolock.Unlock()
	return
}

type VoidCookie uint64

func (cookie VoidCookie) Check(c *Conn) error {
	return c.requestCheck(uint64(cookie))
}

func (c *Conn) requestCheck(request uint64) error {
	c.iolock.Lock()
	if request >= c.in.requestExpected &&
		request > c.in.requestCompleted {

		// send sync request
		c.sendSync()
	}
	replyBuf, isErr := c.waitForReply(request)
	c.iolock.Unlock()

	if isErr {
		return NewGenericError(replyBuf)
	} else {
		// if not err, replyBuf must be nil
		if replyBuf != nil {
			panic("replyBuf is not nil")
		}
	}
	return nil
}

//func (c *Conn) getEvent() []byte {
//l := c.in.events

//if l != nil {
//front := l.Front()
//if front != nil {
//buf := front.Value.([]byte)
//l.Remove(front)
//// buf is event or unchecked error
//return buf
//}
//}
//return nil
//}

//func (c *Conn) getEvent() *GenericEvent {
//select {
//case e := <-c.in.eventCh:
//return e
//default:
//return nil
//}
//}

//func (c *Conn) pollForNextEvent(queued bool) *GenericEvent {
//ret := c.getEvent()
//if ret == nil && !queued {
//// wait event come
//ret = c.getEvent()
//}
//return ret
//}

//func (c *Conn) pollForNextEvent(queued bool) []byte {
//c.iolock.Lock()
//ret := c.getEvent()
//if ret == nil && !queued {
//// wait event come
//c.in.eventCond.Wait()
//ret = c.getEvent()
//}
//c.iolock.Unlock()
//return ret
//}

func (c *Conn) PollForEvent() GenericEvent {
	select {
	case e := <-c.in.eventCh:
		return e
	default:
		return nil
	}
}

func (c *Conn) WaitForEvent() GenericEvent {
	return <-c.in.eventCh
}

//func (c *Conn) PollForQueuedEvent() []byte {
//return c.pollForNextEvent(true)
//}
