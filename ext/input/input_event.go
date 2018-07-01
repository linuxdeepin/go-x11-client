package input

import (
	"github.com/linuxdeepin/go-x11-client"
	"github.com/linuxdeepin/go-x11-client/ext/xfixes"
)

type EventHeader struct {
	DeviceId DeviceId
	Time     x.Timestamp
}

// ge event
type HierarchyEvent struct {
	EventHeader
	Flags uint32
	Infos []HierarchyInfo
}

type HierarchyInfo struct {
	DeviceId   DeviceId
	Attachment DeviceId
	Type       uint8
	Enabled    bool
	Flags      uint32
}

func readHierarchyInfo(r *x.Reader, v *HierarchyInfo) error {
	v.DeviceId = DeviceId(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Attachment = DeviceId(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Type = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Enabled = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(2)
	if r.Err() != nil {
		return r.Err()
	}

	v.Flags = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}
	return nil
}

func readEventHeader(r *x.Reader) (EventHeader, error) {
	var v EventHeader
	v.DeviceId = DeviceId(r.Read2b())
	if r.Err() != nil {
		return EventHeader{}, r.Err()
	}

	v.Time = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return EventHeader{}, r.Err()
	}

	return v, nil
}

func readHierarchyEvent(r *x.Reader, v *HierarchyEvent) error {
	var err error
	v.EventHeader, err = readEventHeader(r)
	if err != nil {
		return err
	}

	v.Flags = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	infosLen := int(r.Read2b())

	r.ReadPad(10)
	if r.Err() != nil {
		return r.Err()
	}

	if infosLen > 0 {
		v.Infos = make([]HierarchyInfo, infosLen)
		for i := 0; i < infosLen; i++ {
			err = readHierarchyInfo(r, &v.Infos[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type DeviceChangedEvent struct {
	EventHeader
	SourceId DeviceId
	Reason   uint8
	Classes  []DeviceClass
}

func readDeviceChangedEvent(r *x.Reader, v *DeviceChangedEvent) error {
	var err error
	v.EventHeader, err = readEventHeader(r)
	if err != nil {
		return err
	}

	classesLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.SourceId = DeviceId(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Reason = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(11)
	if r.Err() != nil {
		return r.Err()
	}

	if classesLen > 0 {
		v.Classes = make([]DeviceClass, classesLen)
		for i := 0; i < classesLen; i++ {
			v.Classes[i], err = readDeviceClass(r)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// ge event
type DeviceEvent struct {
	EventHeader
	Detail       uint32
	Root         x.Window
	Event        x.Window
	Child        x.Window
	RootX        float32
	RootY        float32
	EventX       float32
	EventY       float32
	SourceId     DeviceId
	Flags        uint32
	Mods         ModifierInfo
	Group        GroupInfo
	ButtonMask   []uint32
	ValuatorMask []uint32
	AxisValues   []float32
}

func readDeviceEvent(r *x.Reader, v *DeviceEvent) error {
	var err error
	v.EventHeader, err = readEventHeader(r)
	if err != nil {
		return err
	}

	v.Detail = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Root = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Event = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Child = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.RootX, err = readFP1616(r)
	if err != nil {
		return err
	}

	v.RootY, err = readFP1616(r)
	if err != nil {
		return err
	}

	v.EventX, err = readFP1616(r)
	if err != nil {
		return err
	}

	v.EventY, err = readFP1616(r)
	if err != nil {
		return err
	}

	buttonsLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	valuatorsLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.SourceId = DeviceId(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(2)
	if r.Err() != nil {
		return r.Err()
	}

	v.Flags = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	err = readModifierInfo(r, &v.Mods)
	if err != nil {
		return err
	}

	err = readGroupInfo(r, &v.Group)
	if err != nil {
		return err
	}

	if buttonsLen > 0 {
		v.ButtonMask = make([]uint32, buttonsLen)
		for i := 0; i < buttonsLen; i++ {
			v.ButtonMask[i] = r.Read4b()
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	if valuatorsLen > 0 {
		v.ValuatorMask = make([]uint32, valuatorsLen)
		for i := 0; i < valuatorsLen; i++ {
			v.ValuatorMask[i] = r.Read4b()
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	// TODO AxisValue
	return nil
}

type KeyPressEvent struct {
	DeviceEvent
}

func readKeyPressEvent(r *x.Reader, v *KeyPressEvent) error {
	return readDeviceEvent(r, &v.DeviceEvent)
}

type KeyReleaseEvent struct {
	DeviceEvent
}

func readKeyReleaseEvent(r *x.Reader, v *KeyReleaseEvent) error {
	return readDeviceEvent(r, &v.DeviceEvent)
}

type ButtonPressEvent struct {
	DeviceEvent
}

func readButtonPressEvent(r *x.Reader, v *ButtonPressEvent) error {
	return readDeviceEvent(r, &v.DeviceEvent)
}

type ButtonReleaseEvent struct {
	DeviceEvent
}

func readButtonReleaseEvent(r *x.Reader, v *ButtonReleaseEvent) error {
	return readDeviceEvent(r, &v.DeviceEvent)
}

type MotionEvent struct {
	DeviceEvent
}

func readMotionEvent(r *x.Reader, v *MotionEvent) error {
	return readDeviceEvent(r, &v.DeviceEvent)
}

// ge event
type EnterLeaveEvent struct {
	EventHeader
	SourceId   DeviceId
	Mode       uint8
	Detail     uint8
	Root       x.Window
	Event      x.Window
	Child      x.Window
	RootX      float32
	RootY      float32
	EventX     float32
	EventY     float32
	SameScreen bool
	Focus      bool
	Mods       ModifierInfo
	Group      GroupInfo
	Buttons    []uint32
}

func readEnterLeaveEvent(r *x.Reader, v *EnterLeaveEvent) error {
	var err error
	v.EventHeader, err = readEventHeader(r)
	if err != nil {
		return err
	}

	v.SourceId = DeviceId(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Mode = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Detail = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Root = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Event = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Child = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.RootX, err = readFP1616(r)
	if err != nil {
		return err
	}

	v.RootY, err = readFP1616(r)
	if err != nil {
		return err
	}

	v.EventX, err = readFP1616(r)
	if err != nil {
		return err
	}

	v.EventY, err = readFP1616(r)
	if err != nil {
		return err
	}

	v.SameScreen = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Focus = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	buttonsLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	err = readModifierInfo(r, &v.Mods)
	if err != nil {
		return err
	}

	err = readGroupInfo(r, &v.Group)
	if err != nil {
		return err
	}

	if buttonsLen > 0 {
		v.Buttons = make([]uint32, buttonsLen)
		for i := 0; i < buttonsLen; i++ {
			v.Buttons[i] = r.Read4b()
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	return nil
}

type EnterEvent struct {
	EnterLeaveEvent
}

func readEnterEvent(r *x.Reader, v *EnterEvent) error {
	return readEnterLeaveEvent(r, &v.EnterLeaveEvent)
}

type LeaveEvent struct {
	EnterLeaveEvent
}

func readLeaveEvent(r *x.Reader, v *LeaveEvent) error {
	return readEnterLeaveEvent(r, &v.EnterLeaveEvent)
}

type FocusInEvent struct {
	EnterLeaveEvent
}

func readFocusInEvent(r *x.Reader, v *FocusInEvent) error {
	return readEnterLeaveEvent(r, &v.EnterLeaveEvent)
}

type FocusOutEvent struct {
	EnterLeaveEvent
}

func readFocusOutEvent(r *x.Reader, v *FocusOutEvent) error {
	return readEnterLeaveEvent(r, &v.EnterLeaveEvent)
}

// ge event
type PropertyEvent struct {
	EventHeader
	Property x.Atom
	What     uint8
}

func readPropertyEvent(r *x.Reader, v *PropertyEvent) error {
	var err error
	v.EventHeader, err = readEventHeader(r)
	if err != nil {
		return err
	}

	v.Property = x.Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.What = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(11)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// ge event
type RawEvent struct {
	EventHeader
	Detail        uint32
	SourceId      DeviceId
	Flags         uint32
	ValuatorMask  []uint32
	AxisValues    []float32
	AxisValuesRaw []float32
}

func readRawEvent(r *x.Reader, v *RawEvent) error {
	var err error
	v.EventHeader, err = readEventHeader(r)
	if err != nil {
		return err
	}

	v.Detail = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SourceId = DeviceId(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	valuatorsLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Flags = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(4)
	if r.Err() != nil {
		return r.Err()
	}

	if valuatorsLen > 0 {
		v.ValuatorMask = make([]uint32, valuatorsLen)
		for i := 0; i < valuatorsLen; i++ {
			v.ValuatorMask[i] = r.Read4b()
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	// TODO: axisValues
	// TODO: axisValuesRaw

	return nil
}

type RawKeyPressEvent struct {
	RawEvent
}

func readRawKeyPressEvent(r *x.Reader, v *RawKeyPressEvent) error {
	return readRawEvent(r, &v.RawEvent)
}

type RawKeyReleaseEvent struct {
	RawEvent
}

func readRawKeyReleaseEvent(r *x.Reader, v *RawKeyReleaseEvent) error {
	return readRawEvent(r, &v.RawEvent)
}

type RawButtonPressEvent struct {
	RawEvent
}

func readRawButtonPressEvent(r *x.Reader, v *RawButtonPressEvent) error {
	return readRawEvent(r, &v.RawEvent)
}

type RawButtonReleaseEvent struct {
	RawEvent
}

func readRawButtonReleaseEvent(r *x.Reader, v *RawButtonReleaseEvent) error {
	return readRawEvent(r, &v.RawEvent)
}

type RawMotionEvent struct {
	RawEvent
}

func readRawMotionEvent(r *x.Reader, v *RawMotionEvent) error {
	return readRawEvent(r, &v.RawEvent)
}

type TouchBeginEvent struct {
	DeviceEvent
}

func readTouchBeginEvent(r *x.Reader, v *TouchBeginEvent) error {
	return readDeviceEvent(r, &v.DeviceEvent)
}

type TouchUpdateEvent struct {
	DeviceEvent
}

func readTouchUpdateEvent(r *x.Reader, v *TouchUpdateEvent) error {
	return readDeviceEvent(r, &v.DeviceEvent)
}

type TouchEndEvent struct {
	DeviceEvent
}

func readTouchEndEvent(r *x.Reader, v *TouchEndEvent) error {
	return readDeviceEvent(r, &v.DeviceEvent)
}

type TouchOwnershipEvent struct {
	EventHeader
	TouchId  uint32
	Root     x.Window
	Event    x.Window
	Child    x.Window
	SourceId DeviceId
	Flags    uint32
}

func readTouchOwnershipEvent(r *x.Reader, v *TouchOwnershipEvent) error {
	var err error
	v.EventHeader, err = readEventHeader(r)
	if err != nil {
		return err
	}

	v.TouchId = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Root = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Event = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Child = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.SourceId = DeviceId(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(2)
	if r.Err() != nil {
		return r.Err()
	}

	v.Flags = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(8)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type RawTouchBeginEvent struct {
	RawEvent
}

func readRawTouchBeginEvent(r *x.Reader, v *RawTouchBeginEvent) error {
	return readRawEvent(r, &v.RawEvent)
}

type RawTouchUpdateEvent struct {
	RawEvent
}

func readRawTouchUpdateEvent(r *x.Reader, v *RawTouchUpdateEvent) error {
	return readRawEvent(r, &v.RawEvent)
}

type RawTouchEndEvent struct {
	RawEvent
}

func readRawTouchEndEvent(r *x.Reader, v *RawTouchEndEvent) error {
	return readRawEvent(r, &v.RawEvent)
}

// ge event
type BarrierEvent struct {
	EventHeader
	EventId  uint32
	Root     x.Window
	Event    x.Window
	Barrier  xfixes.Barrier
	DTime    uint32
	Flags    uint32
	SourceId DeviceId
	RootX    float32
	RootY    float32
	DX       float64
	DY       float64
}

func readBarrierEvent(r *x.Reader, v *BarrierEvent) error {
	var err error
	v.EventHeader, err = readEventHeader(r)
	if err != nil {
		return err
	}

	v.EventId = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Root = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Event = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Barrier = xfixes.Barrier(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.DTime = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Flags = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SourceId = DeviceId(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(2)
	if r.Err() != nil {
		return r.Err()
	}

	v.RootX, err = readFP1616(r)
	if err != nil {
		return err
	}

	v.RootY, err = readFP1616(r)
	if err != nil {
		return err
	}

	v.DX, err = readFP3232(r)
	if err != nil {
		return err
	}

	v.DY, err = readFP3232(r)
	if err != nil {
		return err
	}

	return nil
}

type BarrierHitEvent struct {
	BarrierEvent
}

func readBarrierHitEvent(r *x.Reader, v *BarrierHitEvent) error {
	return readBarrierEvent(r, &v.BarrierEvent)
}

type BarrierLeaveEvent struct {
	BarrierHitEvent
}

func readBarrierLeaveEvent(r *x.Reader, v *BarrierLeaveEvent) error {
	return readBarrierEvent(r, &v.BarrierEvent)
}

type DeviceValuatorEvent struct {
	DeviceId      uint8
	Sequence      uint16
	DeviceState   uint16
	FirstValuator uint8
	Valuators     []int32
}

func readDeviceValuatorEvent(r *x.Reader, v *DeviceValuatorEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.DeviceId = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	//seq
	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.DeviceState = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	valuatorsLen := int(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstValuator = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	if valuatorsLen > 0 {
		v.Valuators = make([]int32, valuatorsLen)
		for i := 0; i < valuatorsLen; i++ {
			v.Valuators[i] = int32(r.Read4b())
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	return nil
}

type DeviceKeyEvent struct {
	Detail     uint8
	Sequence   uint16
	Time       x.Timestamp
	Root       x.Window
	Event      x.Window
	Child      x.Window
	RootX      int16
	RootY      int16
	EventX     int16
	EventY     int16
	State      uint16
	SameScreen bool
	DeviceId   uint8
}

func readDeviceKeyEvent(r *x.Reader, v *DeviceKeyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.Detail = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	//seq
	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Time = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Root = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Event = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Child = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.RootX = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.RootY = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.EventX = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.EventY = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.State = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SameScreen = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.DeviceId = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type DeviceKeyPressEvent struct {
	DeviceKeyEvent
}

func readDeviceKeyPressEvent(r *x.Reader, v *DeviceKeyPressEvent) error {
	return readDeviceKeyEvent(r, &v.DeviceKeyEvent)
}

type DeviceKeyReleaseEvent struct {
	DeviceKeyEvent
}

func readDeviceKeyReleaseEvent(r *x.Reader, v *DeviceKeyReleaseEvent) error {
	return readDeviceKeyEvent(r, &v.DeviceKeyEvent)
}

type DeviceButtonPressEvent struct {
	DeviceKeyEvent
}

func readDeviceButtonPressEvent(r *x.Reader, v *DeviceButtonPressEvent) error {
	return readDeviceKeyEvent(r, &v.DeviceKeyEvent)
}

type DeviceButtonReleaseEvent struct {
	DeviceKeyEvent
}

func readDeviceButtonReleaseEvent(r *x.Reader, v *DeviceButtonReleaseEvent) error {
	return readDeviceKeyEvent(r, &v.DeviceKeyEvent)
}

type DeviceMotionNotifyEvent struct {
	DeviceKeyEvent
}

func readDeviceMotionNotifyEvent(r *x.Reader, v *DeviceMotionNotifyEvent) error {
	return readDeviceKeyEvent(r, &v.DeviceKeyEvent)
}

type DeviceFocusEvent struct {
	Detail   uint8
	Sequence uint16
	Time     x.Timestamp
	Window   x.Window
	Mode     uint8
	DeviceId uint8
}

func readDeviceFocusEvent(r *x.Reader, v *DeviceFocusEvent) error {
	//code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.Detail = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Time = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Mode = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.DeviceId = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type DeviceFocusInEvent struct {
	DeviceFocusEvent
}

func readDeviceFocusInEvent(r *x.Reader, v *DeviceFocusInEvent) error {
	return readDeviceFocusEvent(r, &v.DeviceFocusEvent)
}

type DeviceFocusOutEvent struct {
	DeviceFocusEvent
}

func readDeviceFocusOutEvent(r *x.Reader, v *DeviceFocusOutEvent) error {
	return readDeviceFocusEvent(r, &v.DeviceFocusEvent)
}

type ProximityInEvent struct {
	DeviceKeyEvent
}

func readProximityInEvent(r *x.Reader, v *ProximityInEvent) error {
	return readDeviceKeyEvent(r, &v.DeviceKeyEvent)
}

type ProximityOutEvent struct {
	DeviceKeyEvent
}

func readProximityOutEvent(r *x.Reader, v *ProximityOutEvent) error {
	return readDeviceKeyEvent(r, &v.DeviceKeyEvent)
}

type DeviceStateNotifyEvent struct {
	DeviceId        uint8
	Sequence        uint16
	Time            x.Timestamp
	NumKeys         uint8
	NumButtons      uint8
	NumValuators    uint8
	ClassesReported uint8
	Buttons         []uint8  // len 4
	Keys            []uint8  // len 4
	Valuators       []uint32 // len 3
}

func readDeviceStateNotifyEvent(r *x.Reader, v *DeviceStateNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.DeviceId = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Time = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.NumKeys = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NumButtons = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NumValuators = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ClassesReported = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Buttons = r.ReadBytes(4)
	if r.Err() != nil {
		return r.Err()
	}

	v.Keys = r.ReadBytes(4)
	if r.Err() != nil {
		return r.Err()
	}

	v.Valuators = make([]uint32, 3)
	for i := 0; i < 3; i++ {
		v.Valuators[i] = r.Read4b()
		if r.Err() != nil {
			return r.Err()
		}
	}

	return nil
}

type DeviceMappingNotifyEvent struct {
	DeviceId     uint8
	Sequence     uint16
	Request      uint8
	FirstKeycode x.Keycode
	Count        uint8
	Time         x.Timestamp
}

func readDeviceMappingNotifyEvent(r *x.Reader, v *DeviceMappingNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.DeviceId = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Request = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstKeycode = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Count = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.Time = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(20)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type ChangeDeviceNotifyEvent struct {
	DeviceId uint8
	Sequence uint16
	Time     x.Timestamp
	Request  uint8
}

func readChangeDeviceNotifyEvent(r *x.Reader, v *ChangeDeviceNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.DeviceId = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Time = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Request = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(23)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type DeviceKeyStateNotifyEvent struct {
	DeviceId uint8
	Sequence uint16
	Keys     []uint8 // len: 28
}

func readDeviceKeyStateNotifyEvent(r *x.Reader, v *DeviceKeyStateNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.DeviceId = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Keys = r.ReadBytes(28)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type DeviceButtonStateNotifyEvent struct {
	DeviceId uint8
	Sequence uint16
	Buttons  []uint8 // len: 28
}

func readDeviceButtonStateNotifyEvent(r *x.Reader, v *DeviceButtonStateNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.DeviceId = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Buttons = r.ReadBytes(28)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type DevicePresenceNotifyEvent struct {
	Sequence  uint16
	Time      x.Timestamp
	DevChange uint8
	DeviceId  uint8
	Control   uint16
}

func readDevicePresenceNotifyEvent(r *x.Reader, v *DevicePresenceNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Time = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.DevChange = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.DeviceId = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Control = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(20)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type DevicePropertyNotifyEvent struct {
	State    uint8
	Sequence uint16
	Time     x.Timestamp
	Property x.Atom
	DeviceId uint8
}

func readDevicePropertyNotifyEvent(r *x.Reader, v *DevicePropertyNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.State = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Time = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Property = x.Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(19)
	if r.Err() != nil {
		return r.Err()
	}

	v.DeviceId = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}
