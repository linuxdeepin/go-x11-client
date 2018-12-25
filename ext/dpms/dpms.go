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
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ServerMajorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ServerMinorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

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
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Capable = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(23)
	if r.Err() != nil {
		return r.Err()
	}

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
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.StandbyTimeout = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SuspendTimeout = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.OffTimeout = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(18)
	if r.Err() != nil {
		return r.Err()
	}

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
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.PowerLevel = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.State = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(21)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}
