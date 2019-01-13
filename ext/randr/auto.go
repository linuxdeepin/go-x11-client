package randr

import x "github.com/linuxdeepin/go-x11-client"

// _ns.ext_name: RandR
const MajorVersion = 1
const MinorVersion = 6

var _ext *x.Extension

func Ext() *x.Extension {
	return _ext
}

// simple ('xcb', 'RandR', 'MODE')
type Mode uint32

// simple ('xcb', 'RandR', 'CRTC')
type Crtc uint32

// simple ('xcb', 'RandR', 'OUTPUT')
type Output uint32

// simple ('xcb', 'RandR', 'PROVIDER')
type Provider uint32

// simple ('xcb', 'RandR', 'LEASE')
type Lease uint32

const BadOutputErrorCode = 0
const BadCrtcErrorCode = 1
const BadModeErrorCode = 2
const BadProviderErrorCode = 3

// enum Rotation
const (
	RotationRotate0   = 1
	RotationRotate90  = 2
	RotationRotate180 = 4
	RotationRotate270 = 8
	RotationReflectX  = 16
	RotationReflectY  = 32
)

const QueryVersionOpcode = 0

type QueryVersionCookie x.SeqNum

// enum SetConfig
const (
	SetConfigSuccess           = 0
	SetConfigInvalidConfigTime = 1
	SetConfigInvalidTime       = 2
	SetConfigFailed            = 3
)

const SetScreenConfigOpcode = 2

type SetScreenConfigCookie x.SeqNum

// enum NotifyMask
const (
	NotifyMaskScreenChange     = 1
	NotifyMaskCrtcChange       = 2
	NotifyMaskOutputChange     = 4
	NotifyMaskOutputProperty   = 8
	NotifyMaskProviderChange   = 16
	NotifyMaskProviderProperty = 32
	NotifyMaskResourceChange   = 64
	NotifyMaskLease            = 128
)

const SelectInputOpcode = 4
const GetScreenInfoOpcode = 5

type GetScreenInfoCookie x.SeqNum

const GetScreenSizeRangeOpcode = 6

type GetScreenSizeRangeCookie x.SeqNum

const SetScreenSizeOpcode = 7

// enum ModeFlag
const (
	ModeFlagHsyncPositive  = 1
	ModeFlagHsyncNegative  = 2
	ModeFlagVsyncPositive  = 4
	ModeFlagVsyncNegative  = 8
	ModeFlagInterlace      = 16
	ModeFlagDoubleScan     = 32
	ModeFlagCsync          = 64
	ModeFlagCsyncPositive  = 128
	ModeFlagCsyncNegative  = 256
	ModeFlagHskewPresent   = 512
	ModeFlagBcast          = 1024
	ModeFlagPixelMultiplex = 2048
	ModeFlagDoubleClock    = 4096
	ModeFlagHalveClock     = 8192
)

const GetScreenResourcesOpcode = 8

type GetScreenResourcesCookie x.SeqNum

// enum Connection
const (
	ConnectionConnected    = 0
	ConnectionDisconnected = 1
	ConnectionUnknown      = 2
)

const GetOutputInfoOpcode = 9

type GetOutputInfoCookie x.SeqNum

const ListOutputPropertiesOpcode = 10

type ListOutputPropertiesCookie x.SeqNum

const QueryOutputPropertyOpcode = 11

type QueryOutputPropertyCookie x.SeqNum

const ConfigureOutputPropertyOpcode = 12
const ChangeOutputPropertyOpcode = 13
const DeleteOutputPropertyOpcode = 14
const GetOutputPropertyOpcode = 15

type GetOutputPropertyCookie x.SeqNum

const CreateModeOpcode = 16

type CreateModeCookie x.SeqNum

const DestroyModeOpcode = 17
const AddOutputModeOpcode = 18
const DeleteOutputModeOpcode = 19
const GetCrtcInfoOpcode = 20

type GetCrtcInfoCookie x.SeqNum

const SetCrtcConfigOpcode = 21

type SetCrtcConfigCookie x.SeqNum

const GetCrtcGammaSizeOpcode = 22

type GetCrtcGammaSizeCookie x.SeqNum

const GetCrtcGammaOpcode = 23

type GetCrtcGammaCookie x.SeqNum

const SetCrtcGammaOpcode = 24
const GetScreenResourcesCurrentOpcode = 25

type GetScreenResourcesCurrentCookie x.SeqNum

// enum Transform
const (
	TransformUnit       = 1
	TransformScaleUp    = 2
	TransformScaleDown  = 4
	TransformProjective = 8
)

const SetCrtcTransformOpcode = 26
const GetCrtcTransformOpcode = 27

type GetCrtcTransformCookie x.SeqNum

const GetPanningOpcode = 28

type GetPanningCookie x.SeqNum

const SetPanningOpcode = 29

type SetPanningCookie x.SeqNum

const SetOutputPrimaryOpcode = 30
const GetOutputPrimaryOpcode = 31

type GetOutputPrimaryCookie x.SeqNum

const GetProvidersOpcode = 32

type GetProvidersCookie x.SeqNum

// enum ProviderCapability
const (
	ProviderCapabilitySourceOutput  = 1
	ProviderCapabilitySinkOutput    = 2
	ProviderCapabilitySourceOffload = 4
	ProviderCapabilitySinkOffload   = 8
)

const GetProviderInfoOpcode = 33

type GetProviderInfoCookie x.SeqNum

const SetProviderOffloadSinkOpcode = 34
const SetProviderOutputSourceOpcode = 35
const ListProviderPropertiesOpcode = 36

type ListProviderPropertiesCookie x.SeqNum

const QueryProviderPropertyOpcode = 37

type QueryProviderPropertyCookie x.SeqNum

const ConfigureProviderPropertyOpcode = 38
const ChangeProviderPropertyOpcode = 39
const DeleteProviderPropertyOpcode = 40
const GetProviderPropertyOpcode = 41

type GetProviderPropertyCookie x.SeqNum

const ScreenChangeNotifyEventCode = 0

func NewScreenChangeNotifyEvent(data []byte) (*ScreenChangeNotifyEvent, error) {
	var ev ScreenChangeNotifyEvent
	r := x.NewReaderFromData(data)
	if !r.RemainAtLeast4b(8) {
		return nil, x.ErrDataLenShort
	}
	err := readScreenChangeNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

// enum Notify
const (
	NotifyCrtcChange       = 0
	NotifyOutputChange     = 1
	NotifyOutputProperty   = 2
	NotifyProviderChange   = 3
	NotifyProviderProperty = 4
	NotifyResourceChange   = 5
	NotifyLease            = 6
)

const GetMonitorsOpcode = 42

type GetMonitorsCookie x.SeqNum

const SetMonitorOpcode = 43
const DeleteMonitorOpcode = 44
const CreateLeaseOpcode = 45

type CreateLeaseCookie x.SeqNum

const FreeLeaseOpcode = 46
const NotifyEventCode = 1

func NewNotifyEvent(data []byte) (*NotifyEvent, error) {
	var ev NotifyEvent
	r := x.NewReaderFromData(data)
	if !r.RemainAtLeast4b(8) {
		return nil, x.ErrDataLenShort
	}
	err := readNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

var errorCodeNameMap = map[uint8]string{
	BadOutputErrorCode:   "BadOutput",
	BadCrtcErrorCode:     "BadCrtc",
	BadModeErrorCode:     "BadMode",
	BadProviderErrorCode: "BadProvider",
}
var requestOpcodeNameMap = map[uint]string{
	QueryVersionOpcode:              "QueryVersion",
	SetScreenConfigOpcode:           "SetScreenConfig",
	SelectInputOpcode:               "SelectInput",
	GetScreenInfoOpcode:             "GetScreenInfo",
	GetScreenSizeRangeOpcode:        "GetScreenSizeRange",
	SetScreenSizeOpcode:             "SetScreenSize",
	GetScreenResourcesOpcode:        "GetScreenResources",
	GetOutputInfoOpcode:             "GetOutputInfo",
	ListOutputPropertiesOpcode:      "ListOutputProperties",
	QueryOutputPropertyOpcode:       "QueryOutputProperty",
	ConfigureOutputPropertyOpcode:   "ConfigureOutputProperty",
	ChangeOutputPropertyOpcode:      "ChangeOutputProperty",
	DeleteOutputPropertyOpcode:      "DeleteOutputProperty",
	GetOutputPropertyOpcode:         "GetOutputProperty",
	CreateModeOpcode:                "CreateMode",
	DestroyModeOpcode:               "DestroyMode",
	AddOutputModeOpcode:             "AddOutputMode",
	DeleteOutputModeOpcode:          "DeleteOutputMode",
	GetCrtcInfoOpcode:               "GetCrtcInfo",
	SetCrtcConfigOpcode:             "SetCrtcConfig",
	GetCrtcGammaSizeOpcode:          "GetCrtcGammaSize",
	GetCrtcGammaOpcode:              "GetCrtcGamma",
	SetCrtcGammaOpcode:              "SetCrtcGamma",
	GetScreenResourcesCurrentOpcode: "GetScreenResourcesCurrent",
	SetCrtcTransformOpcode:          "SetCrtcTransform",
	GetCrtcTransformOpcode:          "GetCrtcTransform",
	GetPanningOpcode:                "GetPanning",
	SetPanningOpcode:                "SetPanning",
	SetOutputPrimaryOpcode:          "SetOutputPrimary",
	GetOutputPrimaryOpcode:          "GetOutputPrimary",
	GetProvidersOpcode:              "GetProviders",
	GetProviderInfoOpcode:           "GetProviderInfo",
	SetProviderOffloadSinkOpcode:    "SetProviderOffloadSink",
	SetProviderOutputSourceOpcode:   "SetProviderOutputSource",
	ListProviderPropertiesOpcode:    "ListProviderProperties",
	QueryProviderPropertyOpcode:     "QueryProviderProperty",
	ConfigureProviderPropertyOpcode: "ConfigureProviderProperty",
	ChangeProviderPropertyOpcode:    "ChangeProviderProperty",
	DeleteProviderPropertyOpcode:    "DeleteProviderProperty",
	GetProviderPropertyOpcode:       "GetProviderProperty",
	GetMonitorsOpcode:               "GetMonitors",
	SetMonitorOpcode:                "SetMonitor",
	DeleteMonitorOpcode:             "DeleteMonitor",
	CreateLeaseOpcode:               "CreateLease",
	FreeLeaseOpcode:                 "FreeLease",
}

func init() {
	_ext = x.NewExtension("RANDR", 3, errorCodeNameMap, requestOpcodeNameMap)
}
