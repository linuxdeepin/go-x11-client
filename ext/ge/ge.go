package ge

import "github.com/linuxdeepin/go-x11-client"

// #WREQ
func encodeQueryVersion(majorVersion, minorVersion uint16) (b x.RequestBody) {
	b.AddBlock(1).
		Write2b(majorVersion).
		Write2b(minorVersion).
		End()
	return
}

type QueryVersionReply struct {
	MajorVersion uint16
	MinorVersion uint16
}

func readQueryVersionReply(r *x.Reader, v *QueryVersionReply) error {
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

	v.MajorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MinorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(20)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}
