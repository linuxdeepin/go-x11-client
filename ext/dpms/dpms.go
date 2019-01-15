package dpms

import "github.com/linuxdeepin/go-x11-client"

// #WREQ
func encodeGetVersion(clientMajorVersion, clientMinorVersion uint16) (b x.RequestBody) {
	b.AddBlock(1).
		Write2b(clientMajorVersion).
		Write2b(clientMinorVersion).
		End()
	return
}

type GetVersionReply struct {
	ServerMajorVersion uint16
	ServerMinorVersion uint16
}

func readGetVersionReply(r *x.Reader, v *GetVersionReply) error {
	if !r.RemainAtLeast4b(3) {
		return x.ErrDataLenShort
	}
	r.ReadPad(8)

	v.ServerMajorVersion = r.Read2b()
	v.ServerMinorVersion = r.Read2b() // 3

	return nil
}

// #WREQ
func encodeCapable() (b x.RequestBody) {
	return
}

type CapableReply struct {
	Capable bool
}

func readCapableReply(r *x.Reader, v *CapableReply) error {
	if !r.RemainAtLeast4b(3) {
		return x.ErrDataLenShort
	}

	r.ReadPad(8)

	v.Capable = r.ReadBool() // 3

	return nil
}

// #WREQ
func encodeGetTimeouts() (b x.RequestBody) {
	return
}

type GetTimeoutsReply struct {
	StandbyTimeout uint16
	SuspendTimeout uint16
	OffTimeout     uint16
}

func readGetTimeoutsReply(r *x.Reader, v *GetTimeoutsReply) error {
	if !r.RemainAtLeast4b(4) {
		return x.ErrDataLenShort
	}
	r.ReadPad(8)

	v.StandbyTimeout = r.Read2b()
	v.SuspendTimeout = r.Read2b()

	v.OffTimeout = r.Read2b() // 4

	return nil
}

// #WREQ
func encodeSetTimeouts(standbyTimeout, suspendTimeout, offTimeout uint16) (b x.RequestBody) {
	b.AddBlock(2).
		Write2b(standbyTimeout).
		Write2b(suspendTimeout).
		Write2b(offTimeout).
		WritePad(2).
		End()
	return
}

// #WREQ
func encodeEnable() (b x.RequestBody) {
	return
}

// #WREQ
func encodeDisable() (b x.RequestBody) {
	return
}

// #WREQ
func encodeForceLevel(powerLevel uint16) (b x.RequestBody) {
	b.AddBlock(1).
		Write2b(powerLevel).
		WritePad(2).
		End()
	return
}

// #WREQ
func encodeInfo() (b x.RequestBody) {
	return
}

type InfoReply struct {
	PowerLevel uint16
	State      bool
}

func readInfoReply(r *x.Reader, v *InfoReply) error {
	if !r.RemainAtLeast4b(3) {
		return x.ErrDataLenShort
	}
	r.ReadPad(8)

	v.PowerLevel = r.Read2b()
	v.State = r.ReadBool() // 3

	return nil
}
