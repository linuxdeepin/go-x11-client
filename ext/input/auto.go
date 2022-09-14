// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package input

import x "github.com/linuxdeepin/go-x11-client"

// _ns.ext_name: Input
const MajorVersion = 2
const MinorVersion = 3

var _ext *x.Extension

func Ext() *x.Extension {
	return _ext
}

// simple ('xcb', 'Input', 'EventClass')
type EventClass uint32

// simple ('xcb', 'Input', 'KeyCode')
type KeyCode uint8

// simple ('xcb', 'Input', 'DeviceId')
type DeviceId uint16

// simple ('xcb', 'Input', 'FP1616')
type Fp1616 int32

const GetExtensionVersionOpcode = 1

type GetExtensionVersionCookie x.SeqNum

// enum DeviceUse
const (
	DeviceUseIsXPointer           = 0
	DeviceUseIsXKeyboard          = 1
	DeviceUseIsXExtensionDevice   = 2
	DeviceUseIsXExtensionKeyboard = 3
	DeviceUseIsXExtensionPointer  = 4
)

// enum InputClass
const (
	InputClassKey       = 0
	InputClassButton    = 1
	InputClassValuator  = 2
	InputClassFeedback  = 3
	InputClassProximity = 4
	InputClassFocus     = 5
	InputClassOther     = 6
)

// enum ValuatorMode
const (
	ValuatorModeRelative = 0
	ValuatorModeAbsolute = 1
)

const ListInputDevicesOpcode = 2

type ListInputDevicesCookie x.SeqNum

// simple ('xcb', 'Input', 'EventTypeBase')
type EventTypeBase uint8

const OpenDeviceOpcode = 3

type OpenDeviceCookie x.SeqNum

const CloseDeviceOpcode = 4
const SetDeviceModeOpcode = 5

type SetDeviceModeCookie x.SeqNum

const SelectExtensionEventOpcode = 6
const GetSelectedExtensionEventsOpcode = 7

type GetSelectedExtensionEventsCookie x.SeqNum

// enum PropagateMode
const (
	PropagateModeAddToList      = 0
	PropagateModeDeleteFromList = 1
)

const ChangeDeviceDontPropagateListOpcode = 8
const GetDeviceDontPropagateListOpcode = 9

type GetDeviceDontPropagateListCookie x.SeqNum

const GetDeviceMotionEventsOpcode = 10

type GetDeviceMotionEventsCookie x.SeqNum

const ChangeKeyboardDeviceOpcode = 11

type ChangeKeyboardDeviceCookie x.SeqNum

const ChangePointerDeviceOpcode = 12

type ChangePointerDeviceCookie x.SeqNum

const GrabDeviceOpcode = 13

type GrabDeviceCookie x.SeqNum

const UngrabDeviceOpcode = 14

// enum ModifierDevice
const (
	ModifierDeviceUseXKeyboard = 255
)

const GrabDeviceKeyOpcode = 15
const UngrabDeviceKeyOpcode = 16
const GrabDeviceButtonOpcode = 17
const UngrabDeviceButtonOpcode = 18

// enum DeviceInputMode
const (
	DeviceInputModeAsyncThisDevice   = 0
	DeviceInputModeSyncThisDevice    = 1
	DeviceInputModeReplayThisDevice  = 2
	DeviceInputModeAsyncOtherDevices = 3
	DeviceInputModeAsyncAll          = 4
	DeviceInputModeSyncAll           = 5
)

const AllowDeviceEventsOpcode = 19
const GetDeviceFocusOpcode = 20

type GetDeviceFocusCookie x.SeqNum

const SetDeviceFocusOpcode = 21

// enum FeedbackClass
const (
	FeedbackClassKeyboard = 0
	FeedbackClassPointer  = 1
	FeedbackClassString   = 2
	FeedbackClassInteger  = 3
	FeedbackClassLed      = 4
	FeedbackClassBell     = 5
)

const GetFeedbackControlOpcode = 22

type GetFeedbackControlCookie x.SeqNum

// enum ChangeFeedbackControlMask
const (
	ChangeFeedbackControlMaskKeyClickPercent = 1
	ChangeFeedbackControlMaskPercent         = 2
	ChangeFeedbackControlMaskPitch           = 4
	ChangeFeedbackControlMaskDuration        = 8
	ChangeFeedbackControlMaskLed             = 16
	ChangeFeedbackControlMaskLedMode         = 32
	ChangeFeedbackControlMaskKey             = 64
	ChangeFeedbackControlMaskAutoRepeatMode  = 128
	ChangeFeedbackControlMaskString          = 1
	ChangeFeedbackControlMaskInteger         = 1
	ChangeFeedbackControlMaskAccelNum        = 1
	ChangeFeedbackControlMaskAccelDenom      = 2
	ChangeFeedbackControlMaskThreshold       = 4
)

const ChangeFeedbackControlOpcode = 23
const GetDeviceKeyMappingOpcode = 24

type GetDeviceKeyMappingCookie x.SeqNum

const ChangeDeviceKeyMappingOpcode = 25
const GetDeviceModifierMappingOpcode = 26

type GetDeviceModifierMappingCookie x.SeqNum

const SetDeviceModifierMappingOpcode = 27

type SetDeviceModifierMappingCookie x.SeqNum

const GetDeviceButtonMappingOpcode = 28

type GetDeviceButtonMappingCookie x.SeqNum

const SetDeviceButtonMappingOpcode = 29

type SetDeviceButtonMappingCookie x.SeqNum

// enum ValuatorStateModeMask
const (
	ValuatorStateModeMaskDeviceModeAbsolute = 1
	ValuatorStateModeMaskOutOfProximity     = 2
)

const QueryDeviceStateOpcode = 30

type QueryDeviceStateCookie x.SeqNum

const DeviceBellOpcode = 32
const SetDeviceValuatorsOpcode = 33

type SetDeviceValuatorsCookie x.SeqNum

// enum DeviceControl
const (
	DeviceControlResolution = 1
	DeviceControlAbsCalib   = 2
	DeviceControlCore       = 3
	DeviceControlEnable     = 4
	DeviceControlAbsArea    = 5
)

const GetDeviceControlOpcode = 34

type GetDeviceControlCookie x.SeqNum

const ChangeDeviceControlOpcode = 35

type ChangeDeviceControlCookie x.SeqNum

const ListDevicePropertiesOpcode = 36

type ListDevicePropertiesCookie x.SeqNum

// enum PropertyFormat
const (
	PropertyFormat8Bits  = 8
	PropertyFormat16Bits = 16
	PropertyFormat32Bits = 32
)

const ChangeDevicePropertyOpcode = 37
const DeleteDevicePropertyOpcode = 38
const GetDevicePropertyOpcode = 39

type GetDevicePropertyCookie x.SeqNum

// enum Device
const (
	DeviceAll       = 0
	DeviceAllMaster = 1
)

const XIQueryPointerOpcode = 40

type XIQueryPointerCookie x.SeqNum

const XIWarpPointerOpcode = 41
const XIChangeCursorOpcode = 42

// enum HierarchyChangeType
const (
	HierarchyChangeTypeAddMaster    = 1
	HierarchyChangeTypeRemoveMaster = 2
	HierarchyChangeTypeAttachSlave  = 3
	HierarchyChangeTypeDetachSlave  = 4
)

// enum ChangeMode
const (
	ChangeModeAttach = 1
	ChangeModeFloat  = 2
)

const XIChangeHierarchyOpcode = 43
const XISetClientPointerOpcode = 44
const XIGetClientPointerOpcode = 45

type XIGetClientPointerCookie x.SeqNum

// enum XIEventMask
const (
	XIEventMaskDeviceChanged    = 2
	XIEventMaskKeyPress         = 4
	XIEventMaskKeyRelease       = 8
	XIEventMaskButtonPress      = 16
	XIEventMaskButtonRelease    = 32
	XIEventMaskMotion           = 64
	XIEventMaskEnter            = 128
	XIEventMaskLeave            = 256
	XIEventMaskFocusIn          = 512
	XIEventMaskFocusOut         = 1024
	XIEventMaskHierarchy        = 2048
	XIEventMaskProperty         = 4096
	XIEventMaskRawKeyPress      = 8192
	XIEventMaskRawKeyRelease    = 16384
	XIEventMaskRawButtonPress   = 32768
	XIEventMaskRawButtonRelease = 65536
	XIEventMaskRawMotion        = 131072
	XIEventMaskTouchBegin       = 262144
	XIEventMaskTouchUpdate      = 524288
	XIEventMaskTouchEnd         = 1048576
	XIEventMaskTouchOwnership   = 2097152
	XIEventMaskRawTouchBegin    = 4194304
	XIEventMaskRawTouchUpdate   = 8388608
	XIEventMaskRawTouchEnd      = 16777216
	XIEventMaskBarrierHit       = 33554432
	XIEventMaskBarrierLeave     = 67108864
)

const XISelectEventsOpcode = 46
const XIQueryVersionOpcode = 47

type XIQueryVersionCookie x.SeqNum

// enum DeviceClassType
const (
	DeviceClassTypeKey      = 0
	DeviceClassTypeButton   = 1
	DeviceClassTypeValuator = 2
	DeviceClassTypeScroll   = 3
	DeviceClassTypeTouch    = 8
)

// enum DeviceType
const (
	DeviceTypeMasterPointer  = 1
	DeviceTypeMasterKeyboard = 2
	DeviceTypeSlavePointer   = 3
	DeviceTypeSlaveKeyboard  = 4
	DeviceTypeFloatingSlave  = 5
)

// enum ScrollFlags
const (
	ScrollFlagsNoEmulation = 1
	ScrollFlagsPreferred   = 2
)

// enum ScrollType
const (
	ScrollTypeVertical   = 1
	ScrollTypeHorizontal = 2
)

// enum TouchMode
const (
	TouchModeDirect    = 1
	TouchModeDependent = 2
)

const XIQueryDeviceOpcode = 48

type XIQueryDeviceCookie x.SeqNum

const XISetFocusOpcode = 49
const XIGetFocusOpcode = 50

type XIGetFocusCookie x.SeqNum

// enum GrabOwner
const (
	GrabOwnerNoOwner = 0
	GrabOwnerOwner   = 1
)

const XIGrabDeviceOpcode = 51

type XIGrabDeviceCookie x.SeqNum

const XIUngrabDeviceOpcode = 52

// enum EventMode
const (
	EventModeAsyncDevice       = 0
	EventModeSyncDevice        = 1
	EventModeReplayDevice      = 2
	EventModeAsyncPairedDevice = 3
	EventModeAsyncPair         = 4
	EventModeSyncPair          = 5
	EventModeAcceptTouch       = 6
	EventModeRejectTouch       = 7
)

const XIAllowEventsOpcode = 53

// enum GrabMode22
const (
	GrabMode22Sync  = 0
	GrabMode22Async = 1
	GrabMode22Touch = 2
)

// enum GrabType
const (
	GrabTypeButton     = 0
	GrabTypeKeycode    = 1
	GrabTypeEnter      = 2
	GrabTypeFocusIn    = 3
	GrabTypeTouchBegin = 4
)

// enum ModifierMask
const (
	ModifierMaskAny = 2147483648
)

const XIPassiveGrabDeviceOpcode = 54

type XIPassiveGrabDeviceCookie x.SeqNum

const XIPassiveUngrabDeviceOpcode = 55
const XIListPropertiesOpcode = 56

type XIListPropertiesCookie x.SeqNum

const XIChangePropertyOpcode = 57
const XIDeletePropertyOpcode = 58
const XIGetPropertyOpcode = 59

type XIGetPropertyCookie x.SeqNum

const XIGetSelectedEventsOpcode = 60

type XIGetSelectedEventsCookie x.SeqNum

const XIBarrierReleasePointerOpcode = 61
const DeviceValuatorEventCode = 0

func NewDeviceValuatorEvent(data []byte) (*DeviceValuatorEvent, error) {
	var ev DeviceValuatorEvent
	r := x.NewReaderFromData(data)
	err := readDeviceValuatorEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

// enum MoreEventsMask
const (
	MoreEventsMaskMoreEvents = 128
)

const DeviceKeyPressEventCode = 1

func NewDeviceKeyPressEvent(data []byte) (*DeviceKeyPressEvent, error) {
	var ev DeviceKeyPressEvent
	r := x.NewReaderFromData(data)
	err := readDeviceKeyPressEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const DeviceKeyReleaseEventCode = 2

func NewDeviceKeyReleaseEvent(data []byte) (*DeviceKeyReleaseEvent, error) {
	var ev DeviceKeyReleaseEvent
	r := x.NewReaderFromData(data)
	err := readDeviceKeyReleaseEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const DeviceButtonPressEventCode = 3

func NewDeviceButtonPressEvent(data []byte) (*DeviceButtonPressEvent, error) {
	var ev DeviceButtonPressEvent
	r := x.NewReaderFromData(data)
	err := readDeviceButtonPressEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const DeviceButtonReleaseEventCode = 4

func NewDeviceButtonReleaseEvent(data []byte) (*DeviceButtonReleaseEvent, error) {
	var ev DeviceButtonReleaseEvent
	r := x.NewReaderFromData(data)
	err := readDeviceButtonReleaseEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const DeviceMotionNotifyEventCode = 5

func NewDeviceMotionNotifyEvent(data []byte) (*DeviceMotionNotifyEvent, error) {
	var ev DeviceMotionNotifyEvent
	r := x.NewReaderFromData(data)
	err := readDeviceMotionNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const DeviceFocusInEventCode = 6

func NewDeviceFocusInEvent(data []byte) (*DeviceFocusInEvent, error) {
	var ev DeviceFocusInEvent
	r := x.NewReaderFromData(data)
	err := readDeviceFocusInEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const DeviceFocusOutEventCode = 7

func NewDeviceFocusOutEvent(data []byte) (*DeviceFocusOutEvent, error) {
	var ev DeviceFocusOutEvent
	r := x.NewReaderFromData(data)
	err := readDeviceFocusOutEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const ProximityInEventCode = 8

func NewProximityInEvent(data []byte) (*ProximityInEvent, error) {
	var ev ProximityInEvent
	r := x.NewReaderFromData(data)
	err := readProximityInEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const ProximityOutEventCode = 9

func NewProximityOutEvent(data []byte) (*ProximityOutEvent, error) {
	var ev ProximityOutEvent
	r := x.NewReaderFromData(data)
	err := readProximityOutEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

// enum ClassesReportedMask
const (
	ClassesReportedMaskOutOfProximity     = 128
	ClassesReportedMaskDeviceModeAbsolute = 64
	ClassesReportedMaskReportingValuators = 4
	ClassesReportedMaskReportingButtons   = 2
	ClassesReportedMaskReportingKeys      = 1
)

const DeviceStateNotifyEventCode = 10

func NewDeviceStateNotifyEvent(data []byte) (*DeviceStateNotifyEvent, error) {
	var ev DeviceStateNotifyEvent
	r := x.NewReaderFromData(data)
	err := readDeviceStateNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const DeviceMappingNotifyEventCode = 11

func NewDeviceMappingNotifyEvent(data []byte) (*DeviceMappingNotifyEvent, error) {
	var ev DeviceMappingNotifyEvent
	r := x.NewReaderFromData(data)
	err := readDeviceMappingNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

// enum ChangeDevice
const (
	ChangeDeviceNewPointer  = 0
	ChangeDeviceNewKeyboard = 1
)

const ChangeDeviceNotifyEventCode = 12

func NewChangeDeviceNotifyEvent(data []byte) (*ChangeDeviceNotifyEvent, error) {
	var ev ChangeDeviceNotifyEvent
	r := x.NewReaderFromData(data)
	err := readChangeDeviceNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const DeviceKeyStateNotifyEventCode = 13

func NewDeviceKeyStateNotifyEvent(data []byte) (*DeviceKeyStateNotifyEvent, error) {
	var ev DeviceKeyStateNotifyEvent
	r := x.NewReaderFromData(data)
	err := readDeviceKeyStateNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const DeviceButtonStateNotifyEventCode = 14

func NewDeviceButtonStateNotifyEvent(data []byte) (*DeviceButtonStateNotifyEvent, error) {
	var ev DeviceButtonStateNotifyEvent
	r := x.NewReaderFromData(data)
	err := readDeviceButtonStateNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

// enum DeviceChange
const (
	DeviceChangeAdded          = 0
	DeviceChangeRemoved        = 1
	DeviceChangeEnabled        = 2
	DeviceChangeDisabled       = 3
	DeviceChangeUnrecoverable  = 4
	DeviceChangeControlChanged = 5
)

const DevicePresenceNotifyEventCode = 15

func NewDevicePresenceNotifyEvent(data []byte) (*DevicePresenceNotifyEvent, error) {
	var ev DevicePresenceNotifyEvent
	r := x.NewReaderFromData(data)
	err := readDevicePresenceNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const DevicePropertyNotifyEventCode = 16

func NewDevicePropertyNotifyEvent(data []byte) (*DevicePropertyNotifyEvent, error) {
	var ev DevicePropertyNotifyEvent
	r := x.NewReaderFromData(data)
	err := readDevicePropertyNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

// enum ChangeReason
const (
	ChangeReasonSlaveSwitch  = 1
	ChangeReasonDeviceChange = 2
)

const DeviceChangedEventCode = 1

func NewDeviceChangedEvent(data []byte) (*DeviceChangedEvent, error) {
	var ev DeviceChangedEvent
	r := x.NewReaderFromData(data)
	err := readDeviceChangedEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

// enum KeyEventFlags
const (
	KeyEventFlagsKeyRepeat = 65536
)

const KeyPressEventCode = 2

func NewKeyPressEvent(data []byte) (*KeyPressEvent, error) {
	var ev KeyPressEvent
	r := x.NewReaderFromData(data)
	err := readKeyPressEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const KeyReleaseEventCode = 3

func NewKeyReleaseEvent(data []byte) (*KeyReleaseEvent, error) {
	var ev KeyReleaseEvent
	r := x.NewReaderFromData(data)
	err := readKeyReleaseEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

// enum PointerEventFlags
const (
	PointerEventFlagsPointerEmulated = 65536
)

const ButtonPressEventCode = 4

func NewButtonPressEvent(data []byte) (*ButtonPressEvent, error) {
	var ev ButtonPressEvent
	r := x.NewReaderFromData(data)
	err := readButtonPressEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const ButtonReleaseEventCode = 5

func NewButtonReleaseEvent(data []byte) (*ButtonReleaseEvent, error) {
	var ev ButtonReleaseEvent
	r := x.NewReaderFromData(data)
	err := readButtonReleaseEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const MotionEventCode = 6

func NewMotionEvent(data []byte) (*MotionEvent, error) {
	var ev MotionEvent
	r := x.NewReaderFromData(data)
	err := readMotionEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

// enum NotifyMode
const (
	NotifyModeNormal        = 0
	NotifyModeGrab          = 1
	NotifyModeUngrab        = 2
	NotifyModeWhileGrabbed  = 3
	NotifyModePassiveGrab   = 4
	NotifyModePassiveUngrab = 5
)

// enum NotifyDetail
const (
	NotifyDetailAncestor         = 0
	NotifyDetailVirtual          = 1
	NotifyDetailInferior         = 2
	NotifyDetailNonlinear        = 3
	NotifyDetailNonlinearVirtual = 4
	NotifyDetailPointer          = 5
	NotifyDetailPointerRoot      = 6
	NotifyDetailNone             = 7
)

const EnterEventCode = 7

func NewEnterEvent(data []byte) (*EnterEvent, error) {
	var ev EnterEvent
	r := x.NewReaderFromData(data)
	err := readEnterEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const LeaveEventCode = 8

func NewLeaveEvent(data []byte) (*LeaveEvent, error) {
	var ev LeaveEvent
	r := x.NewReaderFromData(data)
	err := readLeaveEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const FocusInEventCode = 9

func NewFocusInEvent(data []byte) (*FocusInEvent, error) {
	var ev FocusInEvent
	r := x.NewReaderFromData(data)
	err := readFocusInEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const FocusOutEventCode = 10

func NewFocusOutEvent(data []byte) (*FocusOutEvent, error) {
	var ev FocusOutEvent
	r := x.NewReaderFromData(data)
	err := readFocusOutEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

// enum HierarchyMask
const (
	HierarchyMaskMasterAdded    = 1
	HierarchyMaskMasterRemoved  = 2
	HierarchyMaskSlaveAdded     = 4
	HierarchyMaskSlaveRemoved   = 8
	HierarchyMaskSlaveAttached  = 16
	HierarchyMaskSlaveDetached  = 32
	HierarchyMaskDeviceEnabled  = 64
	HierarchyMaskDeviceDisabled = 128
)

const HierarchyEventCode = 11

func NewHierarchyEvent(data []byte) (*HierarchyEvent, error) {
	var ev HierarchyEvent
	r := x.NewReaderFromData(data)
	err := readHierarchyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

// enum PropertyFlag
const (
	PropertyFlagDeleted  = 0
	PropertyFlagCreated  = 1
	PropertyFlagModified = 2
)

const PropertyEventCode = 12

func NewPropertyEvent(data []byte) (*PropertyEvent, error) {
	var ev PropertyEvent
	r := x.NewReaderFromData(data)
	err := readPropertyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const RawKeyPressEventCode = 13

func NewRawKeyPressEvent(data []byte) (*RawKeyPressEvent, error) {
	var ev RawKeyPressEvent
	r := x.NewReaderFromData(data)
	err := readRawKeyPressEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const RawKeyReleaseEventCode = 14

func NewRawKeyReleaseEvent(data []byte) (*RawKeyReleaseEvent, error) {
	var ev RawKeyReleaseEvent
	r := x.NewReaderFromData(data)
	err := readRawKeyReleaseEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const RawButtonPressEventCode = 15

func NewRawButtonPressEvent(data []byte) (*RawButtonPressEvent, error) {
	var ev RawButtonPressEvent
	r := x.NewReaderFromData(data)
	err := readRawButtonPressEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const RawButtonReleaseEventCode = 16

func NewRawButtonReleaseEvent(data []byte) (*RawButtonReleaseEvent, error) {
	var ev RawButtonReleaseEvent
	r := x.NewReaderFromData(data)
	err := readRawButtonReleaseEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const RawMotionEventCode = 17

func NewRawMotionEvent(data []byte) (*RawMotionEvent, error) {
	var ev RawMotionEvent
	r := x.NewReaderFromData(data)
	err := readRawMotionEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

// enum TouchEventFlags
const (
	TouchEventFlagsTouchPendingEnd       = 65536
	TouchEventFlagsTouchEmulatingPointer = 131072
)

const TouchBeginEventCode = 18

func NewTouchBeginEvent(data []byte) (*TouchBeginEvent, error) {
	var ev TouchBeginEvent
	r := x.NewReaderFromData(data)
	err := readTouchBeginEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const TouchUpdateEventCode = 19

func NewTouchUpdateEvent(data []byte) (*TouchUpdateEvent, error) {
	var ev TouchUpdateEvent
	r := x.NewReaderFromData(data)
	err := readTouchUpdateEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const TouchEndEventCode = 20

func NewTouchEndEvent(data []byte) (*TouchEndEvent, error) {
	var ev TouchEndEvent
	r := x.NewReaderFromData(data)
	err := readTouchEndEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

// enum TouchOwnershipFlags
const (
	TouchOwnershipFlagsNone = 0
)

const TouchOwnershipEventCode = 21

func NewTouchOwnershipEvent(data []byte) (*TouchOwnershipEvent, error) {
	var ev TouchOwnershipEvent
	r := x.NewReaderFromData(data)
	err := readTouchOwnershipEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const RawTouchBeginEventCode = 22

func NewRawTouchBeginEvent(data []byte) (*RawTouchBeginEvent, error) {
	var ev RawTouchBeginEvent
	r := x.NewReaderFromData(data)
	err := readRawTouchBeginEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const RawTouchUpdateEventCode = 23

func NewRawTouchUpdateEvent(data []byte) (*RawTouchUpdateEvent, error) {
	var ev RawTouchUpdateEvent
	r := x.NewReaderFromData(data)
	err := readRawTouchUpdateEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const RawTouchEndEventCode = 24

func NewRawTouchEndEvent(data []byte) (*RawTouchEndEvent, error) {
	var ev RawTouchEndEvent
	r := x.NewReaderFromData(data)
	err := readRawTouchEndEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

// enum BarrierFlags
const (
	BarrierFlagsPointerReleased = 1
	BarrierFlagsDeviceIsGrabbed = 2
)

const BarrierHitEventCode = 25

func NewBarrierHitEvent(data []byte) (*BarrierHitEvent, error) {
	var ev BarrierHitEvent
	r := x.NewReaderFromData(data)
	err := readBarrierHitEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const BarrierLeaveEventCode = 26

func NewBarrierLeaveEvent(data []byte) (*BarrierLeaveEvent, error) {
	var ev BarrierLeaveEvent
	r := x.NewReaderFromData(data)
	err := readBarrierLeaveEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const SendExtensionEventOpcode = 31
const DeviceErrorCode = 0
const EventErrorCode = 1
const ModeErrorCode = 2
const DeviceBusyErrorCode = 3
const ClassErrorCode = 4

var errorCodeNameMap = map[uint8]string{
	DeviceErrorCode:     "BadDevice",
	EventErrorCode:      "BadEvent",
	ModeErrorCode:       "BadMode",
	DeviceBusyErrorCode: "DeviceBusy",
	ClassErrorCode:      "BadClass",
}
var requestOpcodeNameMap = map[uint]string{
	GetExtensionVersionOpcode:           "GetExtensionVersion",
	ListInputDevicesOpcode:              "ListInputDevices",
	OpenDeviceOpcode:                    "OpenDevice",
	CloseDeviceOpcode:                   "CloseDevice",
	SetDeviceModeOpcode:                 "SetDeviceMode",
	SelectExtensionEventOpcode:          "SelectExtensionEvent",
	GetSelectedExtensionEventsOpcode:    "GetSelectedExtensionEvents",
	ChangeDeviceDontPropagateListOpcode: "ChangeDeviceDontPropagateList",
	GetDeviceDontPropagateListOpcode:    "GetDeviceDontPropagateList",
	GetDeviceMotionEventsOpcode:         "GetDeviceMotionEvents",
	ChangeKeyboardDeviceOpcode:          "ChangeKeyboardDevice",
	ChangePointerDeviceOpcode:           "ChangePointerDevice",
	GrabDeviceOpcode:                    "GrabDevice",
	UngrabDeviceOpcode:                  "UngrabDevice",
	GrabDeviceKeyOpcode:                 "GrabDeviceKey",
	UngrabDeviceKeyOpcode:               "UngrabDeviceKey",
	GrabDeviceButtonOpcode:              "GrabDeviceButton",
	UngrabDeviceButtonOpcode:            "UngrabDeviceButton",
	AllowDeviceEventsOpcode:             "AllowDeviceEvents",
	GetDeviceFocusOpcode:                "GetDeviceFocus",
	SetDeviceFocusOpcode:                "SetDeviceFocus",
	GetFeedbackControlOpcode:            "GetFeedbackControl",
	ChangeFeedbackControlOpcode:         "ChangeFeedbackControl",
	GetDeviceKeyMappingOpcode:           "GetDeviceKeyMapping",
	ChangeDeviceKeyMappingOpcode:        "ChangeDeviceKeyMapping",
	GetDeviceModifierMappingOpcode:      "GetDeviceModifierMapping",
	SetDeviceModifierMappingOpcode:      "SetDeviceModifierMapping",
	GetDeviceButtonMappingOpcode:        "GetDeviceButtonMapping",
	SetDeviceButtonMappingOpcode:        "SetDeviceButtonMapping",
	QueryDeviceStateOpcode:              "QueryDeviceState",
	DeviceBellOpcode:                    "DeviceBell",
	SetDeviceValuatorsOpcode:            "SetDeviceValuators",
	GetDeviceControlOpcode:              "GetDeviceControl",
	ChangeDeviceControlOpcode:           "ChangeDeviceControl",
	ListDevicePropertiesOpcode:          "ListDeviceProperties",
	ChangeDevicePropertyOpcode:          "ChangeDeviceProperty",
	DeleteDevicePropertyOpcode:          "DeleteDeviceProperty",
	GetDevicePropertyOpcode:             "GetDeviceProperty",
	XIQueryPointerOpcode:                "XIQueryPointer",
	XIWarpPointerOpcode:                 "XIWarpPointer",
	XIChangeCursorOpcode:                "XIChangeCursor",
	XIChangeHierarchyOpcode:             "XIChangeHierarchy",
	XISetClientPointerOpcode:            "XISetClientPointer",
	XIGetClientPointerOpcode:            "XIGetClientPointer",
	XISelectEventsOpcode:                "XISelectEvents",
	XIQueryVersionOpcode:                "XIQueryVersion",
	XIQueryDeviceOpcode:                 "XIQueryDevice",
	XISetFocusOpcode:                    "XISetFocus",
	XIGetFocusOpcode:                    "XIGetFocus",
	XIGrabDeviceOpcode:                  "XIGrabDevice",
	XIUngrabDeviceOpcode:                "XIUngrabDevice",
	XIAllowEventsOpcode:                 "XIAllowEvents",
	XIPassiveGrabDeviceOpcode:           "XIPassiveGrabDevice",
	XIPassiveUngrabDeviceOpcode:         "XIPassiveUngrabDevice",
	XIListPropertiesOpcode:              "XIListProperties",
	XIChangePropertyOpcode:              "XIChangeProperty",
	XIDeletePropertyOpcode:              "XIDeleteProperty",
	XIGetPropertyOpcode:                 "XIGetProperty",
	XIGetSelectedEventsOpcode:           "XIGetSelectedEvents",
	XIBarrierReleasePointerOpcode:       "XIBarrierReleasePointer",
	SendExtensionEventOpcode:            "SendExtensionEvent",
}

func init() {
	_ext = x.NewExtension("XInputExtension", 4, errorCodeNameMap, requestOpcodeNameMap)
}
