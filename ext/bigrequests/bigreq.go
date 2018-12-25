package bigrequests

import "github.com/linuxdeepin/go-x11-client"

// #WREQ
func encodeEnable() (b x.RequestBody) {
	return
}

type EnableReply struct {
	MaximumRequestLength uint32
}

func readEnableReply(r *x.Reader, v *EnableReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// sequence
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MaximumRequestLength = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}
