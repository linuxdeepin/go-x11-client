// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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

// size: 3 * 4b
type HierarchyInfo struct {
	DeviceId   DeviceId
	Attachment DeviceId
	Type       uint8
	Enabled    bool
	Flags      uint32
}

func readHierarchyInfo(r *x.Reader, v *HierarchyInfo) {
	v.DeviceId = DeviceId(r.Read2b())
	v.Attachment = DeviceId(r.Read2b())

	v.Type = r.Read1b()
	v.Enabled = r.ReadBool()
	r.ReadPad(2)

	v.Flags = r.Read4b() // 3
}

// EventHeader size: 6b
func readEventHeader(r *x.Reader) EventHeader {
	var v EventHeader
	v.DeviceId = DeviceId(r.Read2b())
	v.Time = x.Timestamp(r.Read4b())
	return v
}

func readHierarchyEvent(r *x.Reader, v *HierarchyEvent) error {
	if !r.RemainAtLeast(22) {
		return x.ErrDataLenShort
	}
	v.EventHeader = readEventHeader(r)

	v.Flags = r.Read4b()

	infosLen := int(r.Read2b())

	r.ReadPad(10) // 22b

	if infosLen > 0 {
		v.Infos = make([]HierarchyInfo, infosLen)
		for i := 0; i < infosLen; i++ {
			readHierarchyInfo(r, &v.Infos[i])
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
	if !r.RemainAtLeast(22) {
		return x.ErrDataLenShort
	}
	v.EventHeader = readEventHeader(r)

	classesLen := int(r.Read2b())
	v.SourceId = DeviceId(r.Read2b())

	v.Reason = r.Read1b()

	r.ReadPad(11) // 22b

	if classesLen > 0 {
		v.Classes = make([]DeviceClass, classesLen)
		for i := 0; i < classesLen; i++ {
			var err error
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
	if !r.RemainAtLeast(70) {
		return x.ErrDataLenShort
	}
	v.EventHeader = readEventHeader(r)

	v.Detail = r.Read4b()

	v.Root = x.Window(r.Read4b())

	v.Event = x.Window(r.Read4b())

	v.Child = x.Window(r.Read4b()) // 22b

	v.RootX = readFP1616(r)

	v.RootY = readFP1616(r)

	v.EventX = readFP1616(r)

	v.EventY = readFP1616(r) // 38b

	buttonsLen := int(r.Read2b())
	valuatorsLen := int(r.Read2b())

	v.SourceId = DeviceId(r.Read2b())
	r.ReadPad(2) // 46b

	v.Flags = r.Read4b()

	readModifierInfo(r, &v.Mods)

	readGroupInfo(r, &v.Group) // 70b

	if buttonsLen > 0 {
		if !r.RemainAtLeast4b(buttonsLen) {
			return x.ErrDataLenShort
		}
		v.ButtonMask = make([]uint32, buttonsLen)
		for i := 0; i < buttonsLen; i++ {
			v.ButtonMask[i] = r.Read4b()
		}
	}

	if valuatorsLen > 0 {
		if !r.RemainAtLeast4b(valuatorsLen) {
			return x.ErrDataLenShort
		}
		v.ValuatorMask = make([]uint32, valuatorsLen)
		for i := 0; i < valuatorsLen; i++ {
			v.ValuatorMask[i] = r.Read4b()
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
	if !r.RemainAtLeast(62) {
		return x.ErrDataLenShort
	}
	v.EventHeader = readEventHeader(r)

	v.SourceId = DeviceId(r.Read2b())
	v.Mode = r.Read1b()
	v.Detail = r.Read1b() // 10b

	v.Root = x.Window(r.Read4b())

	v.Event = x.Window(r.Read4b())

	v.Child = x.Window(r.Read4b()) // 22b

	v.RootX = readFP1616(r)

	v.RootY = readFP1616(r)

	v.EventX = readFP1616(r)

	v.EventY = readFP1616(r) // 38b

	v.SameScreen = r.ReadBool()
	v.Focus = r.ReadBool()
	buttonsLen := int(r.Read2b()) // 42b

	readModifierInfo(r, &v.Mods)

	readGroupInfo(r, &v.Group) // 62b

	if buttonsLen > 0 {
		if !r.RemainAtLeast4b(buttonsLen) {
			return x.ErrDataLenShort
		}
		v.Buttons = make([]uint32, buttonsLen)
		for i := 0; i < buttonsLen; i++ {
			v.Buttons[i] = r.Read4b()
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
	if !r.RemainAtLeast(11) {
		return x.ErrDataLenShort
	}
	v.EventHeader = readEventHeader(r)

	v.Property = x.Atom(r.Read4b())

	v.What = r.Read1b() // 11b

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
	if !r.RemainAtLeast(22) {
		return x.ErrDataLenShort
	}
	v.EventHeader = readEventHeader(r)

	v.Detail = r.Read4b()

	v.SourceId = DeviceId(r.Read2b())
	valuatorsLen := int(r.Read2b()) // 14b

	v.Flags = r.Read4b()

	r.ReadPad(4) // 22b

	if valuatorsLen > 0 {
		if !r.RemainAtLeast4b(valuatorsLen) {
			return x.ErrDataLenShort
		}
		v.ValuatorMask = make([]uint32, valuatorsLen)
		for i := 0; i < valuatorsLen; i++ {
			v.ValuatorMask[i] = r.Read4b()
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
	if !r.RemainAtLeast(30) {
		return x.ErrDataLenShort
	}
	v.EventHeader = readEventHeader(r)

	v.TouchId = r.Read4b() // 10b

	v.Root = x.Window(r.Read4b())

	v.Event = x.Window(r.Read4b())

	v.Child = x.Window(r.Read4b())

	v.SourceId = DeviceId(r.Read2b())
	r.ReadPad(2)

	v.Flags = r.Read4b() // 30b

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
	if !r.RemainAtLeast(58) {
		return x.ErrDataLenShort
	}
	v.EventHeader = readEventHeader(r)

	v.EventId = r.Read4b() // 10b

	v.Root = x.Window(r.Read4b())

	v.Event = x.Window(r.Read4b())

	v.Barrier = xfixes.Barrier(r.Read4b())

	v.DTime = r.Read4b()

	v.Flags = r.Read4b()

	v.SourceId = DeviceId(r.Read2b())
	r.ReadPad(2) // 34b

	v.RootX = readFP1616(r)

	v.RootY = readFP1616(r)

	v.DX = readFP3232(r)

	v.DY = readFP3232(r) //58b

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
	if !r.RemainAtLeast4b(2) {
		return x.ErrDataLenShort
	}
	v.DeviceId, v.Sequence = r.ReadEventHeader()

	v.DeviceState = r.Read2b()
	valuatorsLen := int(r.Read1b())
	v.FirstValuator = r.Read1b() // 2

	if valuatorsLen > 0 {
		if !r.RemainAtLeast4b(valuatorsLen) {
			return x.ErrDataLenShort
		}
		v.Valuators = make([]int32, valuatorsLen)
		for i := 0; i < valuatorsLen; i++ {
			v.Valuators[i] = int32(r.Read4b())
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
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	v.Detail, v.Sequence = r.ReadEventHeader()

	v.Time = x.Timestamp(r.Read4b())

	v.Root = x.Window(r.Read4b())

	v.Event = x.Window(r.Read4b())

	v.Child = x.Window(r.Read4b()) // 5

	v.RootX = int16(r.Read2b())
	v.RootY = int16(r.Read2b())

	v.EventX = int16(r.Read2b())
	v.EventY = int16(r.Read2b())

	v.State = r.Read2b()
	v.SameScreen = r.ReadBool()
	v.DeviceId = r.Read1b() // 8

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
	if !r.RemainAtLeast4b(4) {
		return x.ErrDataLenShort
	}
	v.Detail, v.Sequence = r.ReadEventHeader()

	v.Time = x.Timestamp(r.Read4b())

	v.Window = x.Window(r.Read4b())

	v.Mode = r.Read1b()
	v.DeviceId = r.Read1b() // 4

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
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	v.DeviceId, v.Sequence = r.ReadEventHeader()

	v.Time = x.Timestamp(r.Read4b())

	v.NumKeys = r.Read1b()
	v.NumButtons = r.Read1b()
	v.NumValuators = r.Read1b()
	v.ClassesReported = r.Read1b() // 3

	v.Buttons = r.MustReadBytes(4)

	v.Keys = r.MustReadBytes(4)

	v.Valuators = make([]uint32, 3) // 8
	for i := 0; i < 3; i++ {
		v.Valuators[i] = r.Read4b()
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
	if !r.RemainAtLeast4b(3) {
		return x.ErrDataLenShort
	}
	v.DeviceId, v.Sequence = r.ReadEventHeader()

	v.Request = r.Read1b()
	v.FirstKeycode = x.Keycode(r.Read1b())
	v.Count = r.Read1b()
	r.ReadPad(1)

	v.Time = x.Timestamp(r.Read4b()) // 3

	return nil
}

type ChangeDeviceNotifyEvent struct {
	DeviceId uint8
	Sequence uint16
	Time     x.Timestamp
	Request  uint8
}

func readChangeDeviceNotifyEvent(r *x.Reader, v *ChangeDeviceNotifyEvent) error {
	if !r.RemainAtLeast4b(3) {
		return x.ErrDataLenShort
	}
	v.DeviceId, v.Sequence = r.ReadEventHeader()

	v.Time = x.Timestamp(r.Read4b())

	v.Request = r.Read1b() // 3

	return nil
}

type DeviceKeyStateNotifyEvent struct {
	DeviceId uint8
	Sequence uint16
	Keys     []uint8 // len: 28
}

func readDeviceKeyStateNotifyEvent(r *x.Reader, v *DeviceKeyStateNotifyEvent) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	v.DeviceId, v.Sequence = r.ReadEventHeader()

	v.Keys = r.MustReadBytes(28) // 8

	return nil
}

type DeviceButtonStateNotifyEvent struct {
	DeviceId uint8
	Sequence uint16
	Buttons  []uint8 // len: 28
}

func readDeviceButtonStateNotifyEvent(r *x.Reader, v *DeviceButtonStateNotifyEvent) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	v.DeviceId, v.Sequence = r.ReadEventHeader()

	v.Buttons = r.MustReadBytes(28) // 8

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
	if !r.RemainAtLeast4b(3) {
		return x.ErrDataLenShort
	}
	_, v.Sequence = r.ReadEventHeader()

	v.Time = x.Timestamp(r.Read4b())

	v.DevChange = r.Read1b()
	v.DeviceId = r.Read1b()
	v.Control = r.Read2b() // 3

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
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	v.State, v.Sequence = r.ReadEventHeader()

	v.Time = x.Timestamp(r.Read4b())

	v.Property = x.Atom(r.Read4b())

	r.ReadPad(19)

	v.DeviceId = r.Read1b() // 8

	return nil
}
