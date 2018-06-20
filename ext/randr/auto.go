package randr

import x "github.com/linuxdeepin/go-x11-client"

// _ns.ext_name: RandR
const MajorVersion = 1
const MinorVersion = 5

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

type QueryVersionCookie uint64

// enum SetConfig
const (
	SetConfigSuccess           = 0
	SetConfigInvalidConfigTime = 1
	SetConfigInvalidTime       = 2
	SetConfigFailed            = 3
)

const SetScreenConfigOpcode = 2

type SetScreenConfigCookie uint64

// enum NotifyMask
const (
	NotifyMaskScreenChange     = 1
	NotifyMaskCrtcChange       = 2
	NotifyMaskOutputChange     = 4
	NotifyMaskOutputProperty   = 8
	NotifyMaskProviderChange   = 16
	NotifyMaskProviderProperty = 32
	NotifyMaskResourceChange   = 64
)

const SelectInputOpcode = 4
const GetScreenInfoOpcode = 5

type GetScreenInfoCookie uint64

const GetScreenSizeRangeOpcode = 6

type GetScreenSizeRangeCookie uint64

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

type GetScreenResourcesCookie uint64

// enum Connection
const (
	ConnectionConnected    = 0
	ConnectionDisconnected = 1
	ConnectionUnknown      = 2
)

const GetOutputInfoOpcode = 9

type GetOutputInfoCookie uint64

const ListOutputPropertiesOpcode = 10

type ListOutputPropertiesCookie uint64

const QueryOutputPropertyOpcode = 11

type QueryOutputPropertyCookie uint64

const ConfigureOutputPropertyOpcode = 12
const ChangeOutputPropertyOpcode = 13
const DeleteOutputPropertyOpcode = 14
const GetOutputPropertyOpcode = 15

type GetOutputPropertyCookie uint64

const CreateModeOpcode = 16

type CreateModeCookie uint64

const DestroyModeOpcode = 17
const AddOutputModeOpcode = 18
const DeleteOutputModeOpcode = 19
const GetCrtcInfoOpcode = 20

type GetCrtcInfoCookie uint64

const SetCrtcConfigOpcode = 21

type SetCrtcConfigCookie uint64

const GetCrtcGammaSizeOpcode = 22

type GetCrtcGammaSizeCookie uint64

const GetCrtcGammaOpcode = 23

type GetCrtcGammaCookie uint64

const SetCrtcGammaOpcode = 24
const GetScreenResourcesCurrentOpcode = 25

type GetScreenResourcesCurrentCookie uint64

// enum Transform
const (
	TransformUnit       = 1
	TransformScaleUp    = 2
	TransformScaleDown  = 4
	TransformProjective = 8
)

const SetCrtcTransformOpcode = 26
const GetCrtcTransformOpcode = 27

type GetCrtcTransformCookie uint64

const GetPanningOpcode = 28

type GetPanningCookie uint64

const SetPanningOpcode = 29

type SetPanningCookie uint64

const SetOutputPrimaryOpcode = 30
const GetOutputPrimaryOpcode = 31

type GetOutputPrimaryCookie uint64

const GetProvidersOpcode = 32

type GetProvidersCookie uint64

// enum ProviderCapability
const (
	ProviderCapabilitySourceOutput  = 1
	ProviderCapabilitySinkOutput    = 2
	ProviderCapabilitySourceOffload = 4
	ProviderCapabilitySinkOffload   = 8
)

const GetProviderInfoOpcode = 33

type GetProviderInfoCookie uint64

const SetProviderOffloadSinkOpcode = 34
const SetProviderOutputSourceOpcode = 35
const ListProviderPropertiesOpcode = 36

type ListProviderPropertiesCookie uint64

const QueryProviderPropertyOpcode = 37

type QueryProviderPropertyCookie uint64

const ConfigureProviderPropertyOpcode = 38
const ChangeProviderPropertyOpcode = 39
const DeleteProviderPropertyOpcode = 40
const GetProviderPropertyOpcode = 41

type GetProviderPropertyCookie uint64

const ScreenChangeNotifyEventCode = 0

func NewScreenChangeNotifyEvent(data []byte) (*ScreenChangeNotifyEvent, error) {
	var ev ScreenChangeNotifyEvent
	r := x.NewReaderFromData(data)
	readScreenChangeNotifyEvent(r, &ev)
	if err := r.Err(); err != nil {
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
)

const NotifyEventCode = 1

func NewNotifyEvent(data []byte) (*NotifyEvent, error) {
	var ev NotifyEvent
	r := x.NewReaderFromData(data)
	readNotifyEvent(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

const GetMonitorsOpcode = 42

type GetMonitorsCookie uint64

const SetMonitorOpcode = 43
const DeleteMonitorOpcode = 44

var readErrorFuncMap = make(map[uint8]x.ReadErrorFunc, 4)

func init() {
	readErrorFuncMap[BadOutputErrorCode] = readBadOutputError
	readErrorFuncMap[BadCrtcErrorCode] = readBadCrtcError
	readErrorFuncMap[BadModeErrorCode] = readBadModeError
	readErrorFuncMap[BadProviderErrorCode] = readBadProviderError
	_ext = x.NewExtension("RANDR", 3, readErrorFuncMap)
}
