package dpms

import x "github.com/linuxdeepin/go-x11-client"

// _ns.ext_name: DPMS
const MajorVersion = 0
const MinorVersion = 0

var _ext = x.NewExtension("DPMS")

func Ext() *x.Extension {
	return _ext
}

const GetVersionOpcode = 0

type GetVersionCookie uint64

const CapableOpcode = 1

type CapableCookie uint64

const GetTimeoutsOpcode = 2

type GetTimeoutsCookie uint64

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

type InfoCookie uint64
