package screensaver

import x "github.com/linuxdeepin/go-x11-client"

// _ns.ext_name: ScreenSaver
const MajorVersion = 1
const MinorVersion = 1

var _ext *x.Extension

func Ext() *x.Extension {
	return _ext
}

// enum Kind
const (
	KindBlanked  = 0
	KindInternal = 1
	KindExternal = 2
)

// enum Event
const (
	EventNotifyMask = 1
	EventCycleMask  = 2
)

// enum State
const (
	StateOff      = 0
	StateOn       = 1
	StateCycle    = 2
	StateDisabled = 3
)

const QueryVersionOpcode = 0

type QueryVersionCookie x.SeqNum

const QueryInfoOpcode = 1

type QueryInfoCookie x.SeqNum

const SelectInputOpcode = 2
const SetAttributesOpcode = 3
const UnsetAttributesOpcode = 4
const SuspendOpcode = 5
const NotifyEventCode = 0

func NewNotifyEvent(data []byte) (*NotifyEvent, error) {
	var ev NotifyEvent
	r := x.NewReaderFromData(data)
	readNotifyEvent(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

var requestOpcodeNameMap = map[uint]string{
	QueryVersionOpcode:    "QueryVersion",
	QueryInfoOpcode:       "QueryInfo",
	SelectInputOpcode:     "SelectInput",
	SetAttributesOpcode:   "SetAttributes",
	UnsetAttributesOpcode: "UnsetAttributes",
	SuspendOpcode:         "Suspend",
}

func init() {
	_ext = x.NewExtension("MIT-SCREEN-SAVER", 0, nil, requestOpcodeNameMap)
}
