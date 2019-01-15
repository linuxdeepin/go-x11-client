package composite

import "github.com/linuxdeepin/go-x11-client"

// #WREQ
func encodeQueryVersion(majorVersion, minorVersion uint32) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(majorVersion).
		Write4b(minorVersion).
		End()
	return
}

type QueryVersionReply struct {
	MajorVersion uint32
	MinorVersion uint32
}

func readQueryVersionReply(r *x.Reader, v *QueryVersionReply) error {
	if !r.RemainAtLeast4b(4) {
		return x.ErrDataLenShort
	}

	r.ReadPad(8)

	v.MajorVersion = r.Read4b()

	v.MinorVersion = r.Read4b() // 4

	return nil
}

// #WREQ
func encodeRedirectWindow(window x.Window, update uint8) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(window)).
		Write1b(update).
		WritePad(3).
		End()
	return
}

// #WREQ
func encodeRedirectSubwindows(window x.Window, update uint8) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(window)).
		Write1b(update).
		WritePad(3).
		End()
	return
}

// #WREQ
func encodeUnredirectWindow(window x.Window) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

// #WREQ
func encodeUnredirectSubwindows(window x.Window) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

// #WREQ
func encodeNameWindowPixmap(window x.Window, pixmap x.Pixmap) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(window)).
		Write4b(uint32(pixmap)).
		End()
	return
}
