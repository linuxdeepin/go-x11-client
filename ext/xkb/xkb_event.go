package xkb

import "github.com/linuxdeepin/go-x11-client"

type eventHeader struct {
	XkbType  uint8
	Sequence uint16
	Time     x.Timestamp
	DeviceID uint8
}

func readEventHeader(r *x.Reader, v *eventHeader) {
	v.XkbType, v.Sequence = r.ReadEventHeader()
	if r.Err() != nil {
		return
	}

	v.Time = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return
	}

	v.DeviceID = r.Read1b()
	return
}

type NewKeyboardNotifyEvent struct {
	eventHeader
	OldDeviceID   uint8
	MinKeyCode    x.Keycode
	MaxKeyCode    x.Keycode
	OldMinKeyCode x.Keycode
	OldMaxKeyCode x.Keycode
	RequestMajor  uint8
	RequestMinor  uint8
	Changed       uint16
}

func readNewKeyboardNotifyEvent(r *x.Reader, v *NewKeyboardNotifyEvent) error {
	readEventHeader(r, &v.eventHeader)
	if r.Err() != nil {
		return r.Err()
	}

	v.OldDeviceID = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MinKeyCode = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.MaxKeyCode = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.OldMinKeyCode = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.OldMaxKeyCode = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.RequestMajor = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.RequestMinor = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Changed = r.Read2b()
	return r.Err()
}

type MapNotifyEvent struct {
	eventHeader
	PtrBtnActions    uint8
	Changed          uint16
	MinKeyCode       x.Keycode
	MaxKeyCode       x.Keycode
	FirstType        uint8
	NTypes           uint8
	FirstKeySym      x.Keycode
	NKeySyms         uint8
	FirstKeyAct      x.Keycode
	NKeyActs         uint8
	FirstKeyBehavior x.Keycode
	NKeyBehavior     uint8
	FirstKeyExplicit x.Keycode
	NKeyExplicit     uint8
	FirstModMapKey   x.Keycode
	NModMapKeys      uint8
	FirstVModMapKey  x.Keycode
	NVModMapKeys     uint8
	VirtualMods      uint16
}

func readMapNotifyEvent(r *x.Reader, v *MapNotifyEvent) error {
	readEventHeader(r, &v.eventHeader)
	if r.Err() != nil {
		return r.Err()
	}

	v.PtrBtnActions = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Changed = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MinKeyCode = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.MaxKeyCode = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstType = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NTypes = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstKeySym = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.NKeySyms = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstKeyAct = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.NKeyActs = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstKeyBehavior = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.NKeyBehavior = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstKeyExplicit = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.NKeyExplicit = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstModMapKey = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.NModMapKeys = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstVModMapKey = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.NVModMapKeys = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.VirtualMods = r.Read2b()
	return r.Err()
}

type StateNotifyEvent struct {
	eventHeader
	Mods             uint8
	BaseMods         uint8
	LatchedMods      uint8
	LockedMods       uint8
	Group            uint8
	BaseGroup        int16
	LatchedGroup     int16
	LockedGroup      uint8
	CompatState      uint8
	GrabMods         uint8
	CompatGrabMods   uint8
	LookupMods       uint8
	CompatLookupMods uint8
	PtrBtnState      uint16
	Changed          uint16
	Keycode          x.Keycode
	EventType        uint8
	RequestMajor     uint8
	RequestMinor     uint8
}

func readStateNotifyEvent(r *x.Reader, v *StateNotifyEvent) error {
	readEventHeader(r, &v.eventHeader)
	if r.Err() != nil {
		return r.Err()
	}

	v.Mods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.BaseMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.LatchedMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.LockedMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Group = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.BaseGroup = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.LatchedGroup = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.LockedGroup = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.CompatState = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.GrabMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.CompatGrabMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.LookupMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.CompatLookupMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.PtrBtnState = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Changed = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Keycode = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.EventType = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.RequestMajor = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.RequestMinor = r.Read1b()
	return r.Err()
}

type ControlsNotifyEvent struct {
	eventHeader
	NumGroups             uint8
	ChangedControls       uint32
	EnabledControls       uint32
	EnabledControlChanges uint32
	Keycode               x.Keycode
	EventType             uint8
	RequestMajor          uint8
	RequestMinor          uint8
}

func readControlsNotifyEvent(r *x.Reader, v *ControlsNotifyEvent) error {
	readEventHeader(r, &v.eventHeader)
	if r.Err() != nil {
		return r.Err()
	}

	v.NumGroups = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(2)
	if r.Err() != nil {
		return r.Err()
	}

	v.ChangedControls = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.EnabledControls = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.EnabledControlChanges = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Keycode = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.EventType = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.RequestMajor = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.RequestMinor = r.Read1b()
	return r.Err()
}

type IndicatorStateNotifyEvent struct {
	eventHeader
	State        uint32
	StateChanged uint32
}

func readIndicatorStateNotifyEvent(r *x.Reader, v *IndicatorStateNotifyEvent) error {
	readEventHeader(r, &v.eventHeader)
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(3)
	if r.Err() != nil {
		return r.Err()
	}

	v.State = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.StateChanged = r.Read4b()
	return r.Err()
}

type IndicatorMapNotifyEvent struct {
	eventHeader
	State      uint32
	MapChanged uint32
}

func readIndicatorMapNotifyEvent(r *x.Reader, v *IndicatorMapNotifyEvent) error {
	readEventHeader(r, &v.eventHeader)
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(3)
	if r.Err() != nil {
		return r.Err()
	}

	v.State = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MapChanged = r.Read4b()
	return r.Err()
}

type NamesNotifyEvent struct {
	eventHeader
	Changed            uint16
	FirstType          uint8
	NTypes             uint8
	FirstLevelName     uint8
	NLevelNames        uint8
	NRadioGroups       uint8
	NKeyAliases        uint8
	ChangedGroupNames  uint8
	ChangedVirtualMods uint16
	FirstKey           x.Keycode
	NKeys              uint8
	ChangedIndicators  uint32
}

func readNamesNotifyEvent(r *x.Reader, v *NamesNotifyEvent) error {
	readEventHeader(r, &v.eventHeader)
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.Changed = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstType = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NTypes = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstLevelName = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NLevelNames = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.NRadioGroups = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NKeyAliases = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ChangedGroupNames = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ChangedVirtualMods = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstKey = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.NKeys = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ChangedIndicators = r.Read4b()
	return r.Err()
}

type CompatMapNotifyEvent struct {
	eventHeader
	ChangedGroup uint8
	FirstSI      uint16
	NSI          uint16
	NTotalSI     uint16
}

func readCompatMapNotifyEvent(r *x.Reader, v *CompatMapNotifyEvent) error {
	readEventHeader(r, &v.eventHeader)
	if r.Err() != nil {
		return r.Err()
	}

	v.ChangedGroup = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstSI = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NSI = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NTotalSI = r.Read2b()
	return r.Err()
}

type BellNotifyEvent struct {
	eventHeader
	BellClass uint8
	BellID    uint8
	Percent   uint8
	Pitch     uint16
	Duration  uint16
	Name      x.Atom
	Window    x.Window
	EventOnly bool
}

func readBellNotifyEvent(r *x.Reader, v *BellNotifyEvent) error {
	readEventHeader(r, &v.eventHeader)
	if r.Err() != nil {
		return r.Err()
	}

	v.BellClass = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.BellID = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Percent = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Pitch = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Duration = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Name = x.Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.EventOnly = x.Uint8ToBool(r.Read1b())
	return r.Err()
}

type ActionMessageEvent struct {
	eventHeader
	Keycode         x.Keycode
	Press           bool
	KeyEventFollows bool
	Mods            uint8
	Group           uint8
	Message         string
}

func readActionMessageEvent(r *x.Reader, v *ActionMessageEvent) error {
	readEventHeader(r, &v.eventHeader)
	if r.Err() != nil {
		return r.Err()
	}

	v.Keycode = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Press = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.KeyEventFollows = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Mods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Group = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	msg := r.ReadBytes(8)
	if r.Err() != nil {
		return r.Err()
	}
	v.Message = string(msg)

	return nil
}

type AccessXNotifyEvent struct {
	eventHeader
	Keycode       x.Keycode
	Detail        uint16
	SlowKeysDelay uint16
	DebounceDelay uint16
}

func readAccessXNotifyEvent(r *x.Reader, v *AccessXNotifyEvent) error {
	readEventHeader(r, &v.eventHeader)
	if r.Err() != nil {
		return r.Err()
	}

	v.Keycode = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Detail = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SlowKeysDelay = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.DebounceDelay = r.Read2b()
	return r.Err()
}

type ExtensionDeviceNotifyEvent struct {
	eventHeader
	Reason      uint16
	LedClass    uint16
	LedID       uint16
	LedsDefined uint32
	LedState    uint32
	FirstButton uint8
	NButtons    uint8
	Supported   uint16
	Unsupported uint16
}

func readExtensionDeviceNotifyEvent(r *x.Reader, v *ExtensionDeviceNotifyEvent) error {
	readEventHeader(r, &v.eventHeader)
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.Reason = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.LedClass = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.LedID = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.LedsDefined = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.LedState = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstButton = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NButtons = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Supported = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Unsupported = r.Read2b()
	return r.Err()
}
