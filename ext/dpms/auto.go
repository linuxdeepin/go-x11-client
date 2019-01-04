package dpms

import x "github.com/linuxdeepin/go-x11-client"

// _ns.ext_name: DPMS
const MajorVersion = 0
const MinorVersion = 0

var _ext *x.Extension

func Ext() *x.Extension {
	return _ext
}

const GetVersionOpcode = 0

type GetVersionCookie x.SeqNum

const CapableOpcode = 1

type CapableCookie x.SeqNum

const GetTimeoutsOpcode = 2

type GetTimeoutsCookie x.SeqNum

const SetTimeoutsOpcode = 3
const EnableOpcode = 4
const DisableOpcode = 5

// enum DPMSMode
const (
	DPMSModeOn      = 0
	DPMSModeStandby = 1
	DPMSModeSuspend = 2
	DPMSModeOff     = 3
)

const ForceLevelOpcode = 6
const InfoOpcode = 7

type InfoCookie x.SeqNum

var requestOpcodeNameMap = map[uint]string{
	GetVersionOpcode:  "GetVersion",
	CapableOpcode:     "Capable",
	GetTimeoutsOpcode: "GetTimeouts",
	SetTimeoutsOpcode: "SetTimeouts",
	EnableOpcode:      "Enable",
	DisableOpcode:     "Disable",
	ForceLevelOpcode:  "ForceLevel",
	InfoOpcode:        "Info",
}

func init() {
	_ext = x.NewExtension("DPMS", 0, nil, requestOpcodeNameMap)
}
