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
	if !r.RemainAtLeast4b(3) {
		return x.ErrDataLenShort
	}

	r.ReadPad(8)

	v.MajorVersion = r.Read2b()
	v.MinorVersion = r.Read2b() // 3

	return nil
}
