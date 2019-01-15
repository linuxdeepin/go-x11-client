package xkb

import "github.com/linuxdeepin/go-x11-client"

// size: 9b
type eventHeader struct {
	XkbType  uint8
	Sequence uint16
	Time     x.Timestamp
	DeviceID uint8
}

func readEventHeader(r *x.Reader, v *eventHeader) {
	v.XkbType, v.Sequence = r.ReadEventHeader()
	v.Time = x.Timestamp(r.Read4b())
	v.DeviceID = r.Read1b()
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
	if !r.RemainAtLeast4b(5) {
		return x.ErrDataLenShort
	}
	readEventHeader(r, &v.eventHeader)
	v.OldDeviceID = r.Read1b()
	v.MinKeyCode = x.Keycode(r.Read1b())
	v.MaxKeyCode = x.Keycode(r.Read1b()) // 3

	v.OldMinKeyCode = x.Keycode(r.Read1b())
	v.OldMaxKeyCode = x.Keycode(r.Read1b())
	v.RequestMajor = r.Read1b()
	v.RequestMinor = r.Read1b()

	v.Changed = r.Read2b() // 5
	return nil
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
	if !r.RemainAtLeast4b(7) {
		return x.ErrDataLenShort
	}
	readEventHeader(r, &v.eventHeader)
	v.PtrBtnActions = r.Read1b()
	v.Changed = r.Read2b()
	v.MinKeyCode = x.Keycode(r.Read1b()) // 3

	v.MaxKeyCode = x.Keycode(r.Read1b())
	v.FirstType = r.Read1b()
	v.NTypes = r.Read1b()
	v.FirstKeySym = x.Keycode(r.Read1b())

	v.NKeySyms = r.Read1b()
	v.FirstKeyAct = x.Keycode(r.Read1b())
	v.NKeyActs = r.Read1b()
	v.FirstKeyBehavior = x.Keycode(r.Read1b())

	v.NKeyBehavior = r.Read1b()
	v.FirstKeyExplicit = x.Keycode(r.Read1b())
	v.NKeyExplicit = r.Read1b()
	v.FirstModMapKey = x.Keycode(r.Read1b())

	v.NModMapKeys = r.Read1b()
	v.FirstVModMapKey = x.Keycode(r.Read1b())
	v.NVModMapKeys = r.Read1b()
	v.VirtualMods = r.Read2b() // 7
	return nil
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
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	readEventHeader(r, &v.eventHeader)
	v.Mods = r.Read1b()
	v.BaseMods = r.Read1b()
	v.LatchedMods = r.Read1b() // 3

	v.LockedMods = r.Read1b()
	v.Group = r.Read1b()
	v.BaseGroup = int16(r.Read2b())

	v.LatchedGroup = int16(r.Read2b())
	v.LockedGroup = r.Read1b()
	v.CompatState = r.Read1b()

	v.GrabMods = r.Read1b()
	v.CompatGrabMods = r.Read1b()
	v.LookupMods = r.Read1b()
	v.CompatLookupMods = r.Read1b()

	v.PtrBtnState = r.Read2b()
	v.Changed = r.Read2b()

	v.Keycode = x.Keycode(r.Read1b())
	v.EventType = r.Read1b()
	v.RequestMajor = r.Read1b()
	v.RequestMinor = r.Read1b() // 8
	return nil
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
	if !r.RemainAtLeast4b(7) {
		return x.ErrDataLenShort
	}
	readEventHeader(r, &v.eventHeader)
	v.NumGroups = r.Read1b()
	r.ReadPad(2) // 3

	v.ChangedControls = r.Read4b()

	v.EnabledControls = r.Read4b()

	v.EnabledControlChanges = r.Read4b()

	v.Keycode = x.Keycode(r.Read1b())
	v.EventType = r.Read1b()
	v.RequestMajor = r.Read1b()
	v.RequestMinor = r.Read1b() // 7
	return nil
}

type IndicatorStateNotifyEvent struct {
	eventHeader
	State        uint32
	StateChanged uint32
}

func readIndicatorStateNotifyEvent(r *x.Reader, v *IndicatorStateNotifyEvent) error {
	if !r.RemainAtLeast4b(5) {
		return x.ErrDataLenShort
	}
	readEventHeader(r, &v.eventHeader)
	r.ReadPad(3) // 3

	v.State = r.Read4b()

	v.StateChanged = r.Read4b() // 5
	return nil
}

type IndicatorMapNotifyEvent struct {
	eventHeader
	State      uint32
	MapChanged uint32
}

func readIndicatorMapNotifyEvent(r *x.Reader, v *IndicatorMapNotifyEvent) error {
	if !r.RemainAtLeast4b(5) {
		return x.ErrDataLenShort
	}
	readEventHeader(r, &v.eventHeader)
	r.ReadPad(3) // 3

	v.State = r.Read4b()

	v.MapChanged = r.Read4b() // 5
	return nil
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
	if !r.RemainAtLeast4b(7) {
		return x.ErrDataLenShort
	}
	readEventHeader(r, &v.eventHeader)
	r.ReadPad(1)
	v.Changed = r.Read2b() // 3

	v.FirstType = r.Read1b()
	v.NTypes = r.Read1b()
	v.FirstLevelName = r.Read1b()
	v.NLevelNames = r.Read1b()

	r.ReadPad(1)
	v.NRadioGroups = r.Read1b()
	v.NKeyAliases = r.Read1b()
	v.ChangedGroupNames = r.Read1b()

	v.ChangedVirtualMods = r.Read2b()
	v.FirstKey = x.Keycode(r.Read1b())
	v.NKeys = r.Read1b()

	v.ChangedIndicators = r.Read4b() // 7
	return nil
}

type CompatMapNotifyEvent struct {
	eventHeader
	ChangedGroup uint8
	FirstSI      uint16
	NSI          uint16
	NTotalSI     uint16
}

func readCompatMapNotifyEvent(r *x.Reader, v *CompatMapNotifyEvent) error {
	if !r.RemainAtLeast4b(4) {
		return x.ErrDataLenShort
	}
	readEventHeader(r, &v.eventHeader)
	v.ChangedGroup = r.Read1b()
	v.FirstSI = r.Read2b() // 3

	v.NSI = r.Read2b()
	v.NTotalSI = r.Read2b()
	return nil
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
	if !r.RemainAtLeast4b(7) {
		return x.ErrDataLenShort
	}
	readEventHeader(r, &v.eventHeader)
	v.BellClass = r.Read1b()
	v.BellID = r.Read1b()
	v.Percent = r.Read1b() // 3

	v.Pitch = r.Read2b()
	v.Duration = r.Read2b()

	v.Name = x.Atom(r.Read4b())

	v.Window = x.Window(r.Read4b())

	v.EventOnly = r.ReadBool() // 7
	return nil
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
	if !r.RemainAtLeast(22) {
		return x.ErrDataLenShort
	}
	readEventHeader(r, &v.eventHeader)
	v.Keycode = x.Keycode(r.Read1b())
	v.Press = x.Uint8ToBool(r.Read1b())
	v.KeyEventFollows = x.Uint8ToBool(r.Read1b()) // 3

	v.Mods = r.Read1b()
	v.Group = r.Read1b()

	msg := r.MustReadBytes(8)
	// TODO: 0 in msg
	v.Message = string(msg) // 22b

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
	if !r.RemainAtLeast4b(4) {
		return x.ErrDataLenShort
	}
	readEventHeader(r, &v.eventHeader)
	v.Keycode = x.Keycode(r.Read1b())
	v.Detail = r.Read2b() // 3

	v.SlowKeysDelay = r.Read2b()
	v.DebounceDelay = r.Read2b() // 4
	return nil
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
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	readEventHeader(r, &v.eventHeader)
	r.ReadPad(1)
	v.Reason = r.Read2b() // 3

	v.LedClass = r.Read2b()
	v.LedID = r.Read2b()

	v.LedsDefined = r.Read4b()

	v.LedState = r.Read4b()

	v.FirstButton = r.Read1b()
	v.NButtons = r.Read1b()
	v.Supported = r.Read2b()

	v.Unsupported = r.Read2b() // 8
	return nil
}
