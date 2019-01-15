package xkb

import x "github.com/linuxdeepin/go-x11-client"

// _ns.ext_name: xkb
const MajorVersion = 1
const MinorVersion = 0

var _ext *x.Extension

func Ext() *x.Extension {
	return _ext
}

// enum Const
const (
	ConstMaxLegalKeyCode    = 255
	ConstPerKeyBitArraySize = 32
	ConstKeyNameLength      = 4
)

// enum EventType
const (
	EventTypeNewKeyboardNotify     = 1
	EventTypeMapNotify             = 2
	EventTypeStateNotify           = 4
	EventTypeControlsNotify        = 8
	EventTypeIndicatorStateNotify  = 16
	EventTypeIndicatorMapNotify    = 32
	EventTypeNamesNotify           = 64
	EventTypeCompatMapNotify       = 128
	EventTypeBellNotify            = 256
	EventTypeActionMessage         = 512
	EventTypeAccessXNotify         = 1024
	EventTypeExtensionDeviceNotify = 2048
)

// enum NKNDetail
const (
	NKNDetailKeycodes = 1
	NKNDetailGeometry = 2
	NKNDetailDeviceId = 4
)

// enum AXNDetail
const (
	AXNDetailSkPress    = 1
	AXNDetailSkAccept   = 2
	AXNDetailSkReject   = 4
	AXNDetailSkRelease  = 8
	AXNDetailBkAccept   = 16
	AXNDetailBkReject   = 32
	AXNDetailAxkWarning = 64
)

// enum MapPart
const (
	MapPartKeyTypes           = 1
	MapPartKeySyms            = 2
	MapPartModifierMap        = 4
	MapPartExplicitComponents = 8
	MapPartKeyActions         = 16
	MapPartKeyBehaviors       = 32
	MapPartVirtualMods        = 64
	MapPartVirtualModMap      = 128
)

// enum SetMapFlags
const (
	SetMapFlagsResizeTypes      = 1
	SetMapFlagsRecomputeActions = 2
)

// enum StatePart
const (
	StatePartModifierState    = 1
	StatePartModifierBase     = 2
	StatePartModifierLatch    = 4
	StatePartModifierLock     = 8
	StatePartGroupState       = 16
	StatePartGroupBase        = 32
	StatePartGroupLatch       = 64
	StatePartGroupLock        = 128
	StatePartCompatState      = 256
	StatePartGrabMods         = 512
	StatePartCompatGrabMods   = 1024
	StatePartLookupMods       = 2048
	StatePartCompatLookupMods = 4096
	StatePartPointerButtons   = 8192
)

// enum BoolCtrl
const (
	BoolCtrlRepeatKeys          = 1
	BoolCtrlSlowKeys            = 2
	BoolCtrlBounceKeys          = 4
	BoolCtrlStickyKeys          = 8
	BoolCtrlMouseKeys           = 16
	BoolCtrlMouseKeysAccel      = 32
	BoolCtrlAccessXKeys         = 64
	BoolCtrlAccessXTimeoutMask  = 128
	BoolCtrlAccessXFeedbackMask = 256
	BoolCtrlAudibleBellMask     = 512
	BoolCtrlOverlay1Mask        = 1024
	BoolCtrlOverlay2Mask        = 2048
	BoolCtrlIgnoreGroupLockMask = 4096
)

// enum Control
const (
	ControlGroupsWrap      = 134217728
	ControlInternalMods    = 268435456
	ControlIgnoreLockMods  = 536870912
	ControlPerKeyRepeat    = 1073741824
	ControlControlsEnabled = 2147483648
)

// enum AXOption
const (
	AXOptionSkPressFb    = 1
	AXOptionSkAcceptFb   = 2
	AXOptionFeatureFb    = 4
	AXOptionSlowWarnFb   = 8
	AXOptionIndicatorFb  = 16
	AXOptionStickyKeysFb = 32
	AXOptionTwoKeys      = 64
	AXOptionLatchToLock  = 128
	AXOptionSkReleaseFb  = 256
	AXOptionSkRejectFb   = 512
	AXOptionBkRejectFb   = 1024
	AXOptionDumbBell     = 2048
)

// simple ('xcb', 'xkb', 'DeviceSpec')
type DeviceSpec uint16

// enum LedClassResult
const (
	LedClassResultKbdFeedbackClass = 0
	LedClassResultLedFeedbackClass = 4
)

// enum LedClass
const (
	LedClassKbdFeedbackClass = 0
	LedClassLedFeedbackClass = 4
	LedClassDfltXiClass      = 768
	LedClassAllXiClasses     = 1280
)

// simple ('xcb', 'xkb', 'LedClassSpec')
type LedClassSpec uint16

// enum BellClassResult
const (
	BellClassResultKbdFeedbackClass  = 0
	BellClassResultBellFeedbackClass = 5
)

// enum BellClass
const (
	BellClassKbdFeedbackClass  = 0
	BellClassBellFeedbackClass = 5
	BellClassDfltXiClass       = 768
)

// simple ('xcb', 'xkb', 'BellClassSpec')
type BellClassSpec uint16

// enum ID
const (
	IDUseCoreKbd  = 256
	IDUseCorePtr  = 512
	IDDfltXiClass = 768
	IDDfltXiId    = 1024
	IDAllXiClass  = 1280
	IDAllXiId     = 1536
	IDXiNone      = 65280
)

// simple ('xcb', 'xkb', 'IDSpec')
type IdSpec uint16

// enum Group
const (
	Group1 = 0
	Group2 = 1
	Group3 = 2
	Group4 = 3
)

// enum Groups
const (
	GroupsAny = 254
	GroupsAll = 255
)

// enum SetOfGroup
const (
	SetOfGroupGroup1 = 1
	SetOfGroupGroup2 = 2
	SetOfGroupGroup3 = 4
	SetOfGroupGroup4 = 8
)

// enum SetOfGroups
const (
	SetOfGroupsAny = 128
)

// enum GroupsWrap
const (
	GroupsWrapWrapIntoRange     = 0
	GroupsWrapClampIntoRange    = 64
	GroupsWrapRedirectIntoRange = 128
)

// enum VModsHigh
const (
	VModsHigh15 = 128
	VModsHigh14 = 64
	VModsHigh13 = 32
	VModsHigh12 = 16
	VModsHigh11 = 8
	VModsHigh10 = 4
	VModsHigh9  = 2
	VModsHigh8  = 1
)

// enum VModsLow
const (
	VModsLow7 = 128
	VModsLow6 = 64
	VModsLow5 = 32
	VModsLow4 = 16
	VModsLow3 = 8
	VModsLow2 = 4
	VModsLow1 = 2
	VModsLow0 = 1
)

// enum VMod
const (
	VMod15 = 32768
	VMod14 = 16384
	VMod13 = 8192
	VMod12 = 4096
	VMod11 = 2048
	VMod10 = 1024
	VMod9  = 512
	VMod8  = 256
	VMod7  = 128
	VMod6  = 64
	VMod5  = 32
	VMod4  = 16
	VMod3  = 8
	VMod2  = 4
	VMod1  = 2
	VMod0  = 1
)

// enum Explicit
const (
	ExplicitVModMap    = 128
	ExplicitBehavior   = 64
	ExplicitAutoRepeat = 32
	ExplicitInterpret  = 16
	ExplicitKeyType4   = 8
	ExplicitKeyType3   = 4
	ExplicitKeyType2   = 2
	ExplicitKeyType1   = 1
)

// enum SymInterpretMatch
const (
	SymInterpretMatchNoneOf      = 0
	SymInterpretMatchAnyOfOrNone = 1
	SymInterpretMatchAnyOf       = 2
	SymInterpretMatchAllOf       = 3
	SymInterpretMatchExactly     = 4
)

// enum SymInterpMatch
const (
	SymInterpMatchLevelOneOnly = 128
	SymInterpMatchOpMask       = 127
)

// enum IMFlag
const (
	IMFlagNoExplicit  = 128
	IMFlagNoAutomatic = 64
	IMFlagLedDrivesKb = 32
)

// enum IMModsWhich
const (
	IMModsWhichUseCompat    = 16
	IMModsWhichUseEffective = 8
	IMModsWhichUseLocked    = 4
	IMModsWhichUseLatched   = 2
	IMModsWhichUseBase      = 1
)

// enum IMGroupsWhich
const (
	IMGroupsWhichUseCompat    = 16
	IMGroupsWhichUseEffective = 8
	IMGroupsWhichUseLocked    = 4
	IMGroupsWhichUseLatched   = 2
	IMGroupsWhichUseBase      = 1
)

// enum CMDetail
const (
	CMDetailSymInterp   = 1
	CMDetailGroupCompat = 2
)

// enum NameDetail
const (
	NameDetailKeycodes        = 1
	NameDetailGeometry        = 2
	NameDetailSymbols         = 4
	NameDetailPhysSymbols     = 8
	NameDetailTypes           = 16
	NameDetailCompat          = 32
	NameDetailKeyTypeNames    = 64
	NameDetailKtLevelNames    = 128
	NameDetailIndicatorNames  = 256
	NameDetailKeyNames        = 512
	NameDetailKeyAliases      = 1024
	NameDetailVirtualModNames = 2048
	NameDetailGroupNames      = 4096
	NameDetailRgNames         = 8192
)

// enum GBNDetail
const (
	GBNDetailTypes         = 1
	GBNDetailCompatMap     = 2
	GBNDetailClientSymbols = 4
	GBNDetailServerSymbols = 8
	GBNDetailIndicatorMaps = 16
	GBNDetailKeyNames      = 32
	GBNDetailGeometry      = 64
	GBNDetailOtherNames    = 128
)

// enum XIFeature
const (
	XIFeatureKeyboards      = 1
	XIFeatureButtonActions  = 2
	XIFeatureIndicatorNames = 4
	XIFeatureIndicatorMaps  = 8
	XIFeatureIndicatorState = 16
)

// enum PerClientFlag
const (
	PerClientFlagDetectableAutoRepeat   = 1
	PerClientFlagGrabsUseXkbState       = 2
	PerClientFlagAutoResetControls      = 4
	PerClientFlagLookupStateWhenGrabbed = 8
	PerClientFlagSendEventUsesXkbState  = 16
)

// enum BehaviorType
const (
	BehaviorTypeDefault             = 0
	BehaviorTypeLock                = 1
	BehaviorTypeRadioGroup          = 2
	BehaviorTypeOverlay1            = 3
	BehaviorTypeOverlay2            = 4
	BehaviorTypePermamentLock       = 129
	BehaviorTypePermamentRadioGroup = 130
	BehaviorTypePermamentOverlay1   = 131
	BehaviorTypePermamentOverlay2   = 132
)

// simple ('xcb', 'xkb', 'STRING8')
type String8 byte

// enum DoodadType
const (
	DoodadTypeOutline   = 1
	DoodadTypeSolid     = 2
	DoodadTypeText      = 3
	DoodadTypeIndicator = 4
	DoodadTypeLogo      = 5
)

// enum Error
const (
	ErrorBadDevice = 255
	ErrorBadClass  = 254
	ErrorBadId     = 253
)

const KeyboardErrorCode = 0

// enum SA
const (
	SAClearLocks    = 1
	SALatchToLock   = 2
	SAUseModMapMods = 4
	SAGroupAbsolute = 4
)

// enum SAType
const (
	SATypeNoAction       = 0
	SATypeSetMods        = 1
	SATypeLatchMods      = 2
	SATypeLockMods       = 3
	SATypeSetGroup       = 4
	SATypeLatchGroup     = 5
	SATypeLockGroup      = 6
	SATypeMovePtr        = 7
	SATypePtrBtn         = 8
	SATypeLockPtrBtn     = 9
	SATypeSetPtrDflt     = 10
	SATypeIsoLock        = 11
	SATypeTerminate      = 12
	SATypeSwitchScreen   = 13
	SATypeSetControls    = 14
	SATypeLockControls   = 15
	SATypeActionMessage  = 16
	SATypeRedirectKey    = 17
	SATypeDeviceBtn      = 18
	SATypeLockDeviceBtn  = 19
	SATypeDeviceValuator = 20
)

// enum SAMovePtrFlag
const (
	SAMovePtrFlagNoAcceleration = 1
	SAMovePtrFlagMoveAbsoluteX  = 2
	SAMovePtrFlagMoveAbsoluteY  = 4
)

// enum SASetPtrDfltFlag
const (
	SASetPtrDfltFlagDfltBtnAbsolute  = 4
	SASetPtrDfltFlagAffectDfltButton = 1
)

// enum SAIsoLockFlag
const (
	SAIsoLockFlagNoLock         = 1
	SAIsoLockFlagNoUnlock       = 2
	SAIsoLockFlagUseModMapMods  = 4
	SAIsoLockFlagGroupAbsolute  = 4
	SAIsoLockFlagIsoDfltIsGroup = 8
)

// enum SAIsoLockNoAffect
const (
	SAIsoLockNoAffectCtrls = 8
	SAIsoLockNoAffectPtr   = 16
	SAIsoLockNoAffectGroup = 32
	SAIsoLockNoAffectMods  = 64
)

// enum SwitchScreenFlag
const (
	SwitchScreenFlagApplication = 1
	SwitchScreenFlagAbsolute    = 4
)

// enum BoolCtrlsHigh
const (
	BoolCtrlsHighAccessXFeedback = 1
	BoolCtrlsHighAudibleBell     = 2
	BoolCtrlsHighOverlay1        = 4
	BoolCtrlsHighOverlay2        = 8
	BoolCtrlsHighIgnoreGroupLock = 16
)

// enum BoolCtrlsLow
const (
	BoolCtrlsLowRepeatKeys     = 1
	BoolCtrlsLowSlowKeys       = 2
	BoolCtrlsLowBounceKeys     = 4
	BoolCtrlsLowStickyKeys     = 8
	BoolCtrlsLowMouseKeys      = 16
	BoolCtrlsLowMouseKeysAccel = 32
	BoolCtrlsLowAccessXKeys    = 64
	BoolCtrlsLowAccessXTimeout = 128
)

// enum ActionMessageFlag
const (
	ActionMessageFlagOnPress     = 1
	ActionMessageFlagOnRelease   = 2
	ActionMessageFlagGenKeyEvent = 4
)

// enum LockDeviceFlags
const (
	LockDeviceFlagsNoLock   = 1
	LockDeviceFlagsNoUnlock = 2
)

// enum SAValWhat
const (
	SAValWhatIgnoreVal      = 0
	SAValWhatSetValMin      = 1
	SAValWhatSetValCenter   = 2
	SAValWhatSetValMax      = 3
	SAValWhatSetValRelative = 4
	SAValWhatSetValAbsolute = 5
)

const UseExtensionOpcode = 0

type UseExtensionCookie x.SeqNum

const SelectEventsOpcode = 1
const BellOpcode = 3
const GetStateOpcode = 4

type GetStateCookie x.SeqNum

const LatchLockStateOpcode = 5
const GetControlsOpcode = 6

type GetControlsCookie x.SeqNum

const SetControlsOpcode = 7
const GetMapOpcode = 8

type GetMapCookie x.SeqNum

const SetMapOpcode = 9
const GetCompatMapOpcode = 10

type GetCompatMapCookie x.SeqNum

const SetCompatMapOpcode = 11
const GetIndicatorStateOpcode = 12

type GetIndicatorStateCookie x.SeqNum

const GetIndicatorMapOpcode = 13

type GetIndicatorMapCookie x.SeqNum

const SetIndicatorMapOpcode = 14
const GetNamedIndicatorOpcode = 15

type GetNamedIndicatorCookie x.SeqNum

const SetNamedIndicatorOpcode = 16
const GetNamesOpcode = 17

type GetNamesCookie x.SeqNum

const SetNamesOpcode = 18
const PerClientFlagsOpcode = 21

type PerClientFlagsCookie x.SeqNum

const ListComponentsOpcode = 22

type ListComponentsCookie x.SeqNum

const GetKbdByNameOpcode = 23

type GetKbdByNameCookie x.SeqNum

const GetDeviceInfoOpcode = 24

type GetDeviceInfoCookie x.SeqNum

const SetDeviceInfoOpcode = 25
const SetDebuggingFlagsOpcode = 101

type SetDebuggingFlagsCookie x.SeqNum

const NewKeyboardNotifyEventCode = 0

func NewNewKeyboardNotifyEvent(data []byte) (*NewKeyboardNotifyEvent, error) {
	var ev NewKeyboardNotifyEvent
	r := x.NewReaderFromData(data)
	err := readNewKeyboardNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const MapNotifyEventCode = 1

func NewMapNotifyEvent(data []byte) (*MapNotifyEvent, error) {
	var ev MapNotifyEvent
	r := x.NewReaderFromData(data)
	err := readMapNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const StateNotifyEventCode = 2

func NewStateNotifyEvent(data []byte) (*StateNotifyEvent, error) {
	var ev StateNotifyEvent
	r := x.NewReaderFromData(data)
	err := readStateNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const ControlsNotifyEventCode = 3

func NewControlsNotifyEvent(data []byte) (*ControlsNotifyEvent, error) {
	var ev ControlsNotifyEvent
	r := x.NewReaderFromData(data)
	err := readControlsNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const IndicatorStateNotifyEventCode = 4

func NewIndicatorStateNotifyEvent(data []byte) (*IndicatorStateNotifyEvent, error) {
	var ev IndicatorStateNotifyEvent
	r := x.NewReaderFromData(data)
	err := readIndicatorStateNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const IndicatorMapNotifyEventCode = 5

func NewIndicatorMapNotifyEvent(data []byte) (*IndicatorMapNotifyEvent, error) {
	var ev IndicatorMapNotifyEvent
	r := x.NewReaderFromData(data)
	err := readIndicatorMapNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const NamesNotifyEventCode = 6

func NewNamesNotifyEvent(data []byte) (*NamesNotifyEvent, error) {
	var ev NamesNotifyEvent
	r := x.NewReaderFromData(data)
	err := readNamesNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const CompatMapNotifyEventCode = 7

func NewCompatMapNotifyEvent(data []byte) (*CompatMapNotifyEvent, error) {
	var ev CompatMapNotifyEvent
	r := x.NewReaderFromData(data)
	err := readCompatMapNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const BellNotifyEventCode = 8

func NewBellNotifyEvent(data []byte) (*BellNotifyEvent, error) {
	var ev BellNotifyEvent
	r := x.NewReaderFromData(data)
	err := readBellNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const ActionMessageEventCode = 9

func NewActionMessageEvent(data []byte) (*ActionMessageEvent, error) {
	var ev ActionMessageEvent
	r := x.NewReaderFromData(data)
	err := readActionMessageEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const AccessXNotifyEventCode = 10

func NewAccessXNotifyEvent(data []byte) (*AccessXNotifyEvent, error) {
	var ev AccessXNotifyEvent
	r := x.NewReaderFromData(data)
	err := readAccessXNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

const ExtensionDeviceNotifyEventCode = 11

func NewExtensionDeviceNotifyEvent(data []byte) (*ExtensionDeviceNotifyEvent, error) {
	var ev ExtensionDeviceNotifyEvent
	r := x.NewReaderFromData(data)
	err := readExtensionDeviceNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

var errorCodeNameMap = map[uint8]string{
	KeyboardErrorCode: "BadKeyboard",
}
var requestOpcodeNameMap = map[uint]string{
	UseExtensionOpcode:      "UseExtension",
	SelectEventsOpcode:      "SelectEvents",
	BellOpcode:              "Bell",
	GetStateOpcode:          "GetState",
	LatchLockStateOpcode:    "LatchLockState",
	GetControlsOpcode:       "GetControls",
	SetControlsOpcode:       "SetControls",
	GetMapOpcode:            "GetMap",
	SetMapOpcode:            "SetMap",
	GetCompatMapOpcode:      "GetCompatMap",
	SetCompatMapOpcode:      "SetCompatMap",
	GetIndicatorStateOpcode: "GetIndicatorState",
	GetIndicatorMapOpcode:   "GetIndicatorMap",
	SetIndicatorMapOpcode:   "SetIndicatorMap",
	GetNamedIndicatorOpcode: "GetNamedIndicator",
	SetNamedIndicatorOpcode: "SetNamedIndicator",
	GetNamesOpcode:          "GetNames",
	SetNamesOpcode:          "SetNames",
	PerClientFlagsOpcode:    "PerClientFlags",
	ListComponentsOpcode:    "ListComponents",
	GetKbdByNameOpcode:      "GetKbdByName",
	GetDeviceInfoOpcode:     "GetDeviceInfo",
	SetDeviceInfoOpcode:     "SetDeviceInfo",
	SetDebuggingFlagsOpcode: "SetDebuggingFlags",
}

func init() {
	_ext = x.NewExtension("XKEYBOARD", 0, errorCodeNameMap, requestOpcodeNameMap)
}
