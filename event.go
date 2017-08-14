package client

import (
	"fmt"
)

type GenericEvent []byte

func (ge GenericEvent) GetEventCode() uint8 {
	return ge[0] &^ 0x80
}

// TODO
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

//type Event interface {
//GetBytes() []byte
//GetNumber() uint8
//}

//// Event is an interface that can contain any of the events returned by the
//// server. Use a type assertion switch to extract the Event structs.
//type Event interface {
//Bytes() []byte
//String() string
//}

//// NewEventFun is the type of function use to construct events from raw bytes.
//// It should not be used. It is exported for use in the extension sub-packages.
//type NewEventFun func(buf []byte) Event

//// NewEventFuncs is a map from event numbers to functions that create
//// the corresponding event. It should not be used. It is exported for use
//// in the extension sub-packages.
//var NewEventFuncs = make(map[int]NewEventFun)

//// NewExtEventFuncs is a temporary map that stores event constructor functions
//// for each extension. When an extension is initialized, each event for that
//// extension is added to the 'NewEventFuncs' map. It should not be used. It is
//// exported for use in the extension sub-packages.
//var NewExtEventFuncs = make(map[string]map[int]NewEventFun)
