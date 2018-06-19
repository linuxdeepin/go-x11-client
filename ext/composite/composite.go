package composite

import "github.com/linuxdeepin/go-x11-client"

// #WREQ
func writeQueryVersion(w *x.Writer, majorVersion, minorVersion uint32) {
	w.WritePad(4)
	w.Write4b(majorVersion)
	w.Write4b(minorVersion)
}

type QueryVersionReply struct {
	MajorVersion uint32
	MinorVersion uint32
}

func readQueryVersionReply(r *x.Reader, v *QueryVersionReply) error {
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

	v.MajorVersion = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MinorVersion = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeRedirectWindow(w *x.Writer, window x.Window, update uint8) {
	w.WritePad(4)
	w.Write4b(uint32(window))
	w.Write1b(update)
	w.WritePad(3)
}

// #WREQ
func writeRedirectSubwindows(w *x.Writer, window x.Window, update uint8) {
	w.WritePad(4)
	w.Write4b(uint32(window))
	w.Write1b(update)
	w.WritePad(3)
}

// #WREQ
func writeUnredirectWindow(w *x.Writer, window x.Window) {
	w.WritePad(4)
	w.Write4b(uint32(window))
}

// #WREQ
func writeUnredirectSubwindows(w *x.Writer, window x.Window) {
	w.WritePad(4)
	w.Write4b(uint32(window))
}

// #WREQ
func writeNameWindowPixmap(w *x.Writer, window x.Window, pixmap x.Pixmap) {
	w.WritePad(4)
	w.Write4b(uint32(window))
	w.Write4b(uint32(pixmap))
}
