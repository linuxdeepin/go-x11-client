package client

import (
	"fmt"
)

type GenericEvent []byte

func (ge GenericEvent) GetEventCode() uint8 {
	return ge[0] &^ 0x80
}

func (ge GenericEvent) Real() bool {
	return ge[0]&0x80 == 0
}

func (ge GenericEvent) String() string {
	return fmt.Sprintf("GenericEvent{ EventCode: %d, Real: %v }", ge.GetEventCode(), ge.Real())
}

type EventHandler interface {
	Run(ge GenericEvent)
	Detach(xid uint32)
}

// Detach removes all callbacks associated with a particular x resource.
func (conn *Conn) Detach(xid uint32) {
	conn.EventHandlersMu.RLock()
	for _, eh := range conn.EventHandlers {
		eh.Detach(xid)
	}
	conn.EventHandlersMu.RUnlock()
}

func (conn *Conn) EventLoop() {
	Logger.Println("start event loop")
	for {
		ge := conn.WaitForEvent()
		Logger.Println("eventLoop: get ge", ge)
		if ge == nil {
			return
		}

		conn.EventHandlersMu.RLock()
		eh := conn.EventHandlers[ge.GetEventCode()]
		conn.EventHandlersMu.RUnlock()

		Logger.Println("get eh", eh)

		if eh == nil {
			Logger.Println("ignore event", ge)
			continue
		}
		Logger.Println("handle event", ge)
		eh.Run(ge)
	}
}
