package test

import (
	x "github.com/linuxdeepin/go-x11-client"
)

// #WREQ
func writeGetVersion(w *x.Writer, majorVersion uint8, minorVersion uint16) {
	w.WritePad(4)
	w.Write1b(majorVersion)
	w.WritePad(1)
	w.Write2b(minorVersion)
}

type GetVersionReply struct {
	MajorVersion uint8
	MinorVersion uint16
}

func readGetVersionReply(r *x.Reader, v *GetVersionReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MajorVersion = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// len
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MinorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeCompareCursor(w *x.Writer, window x.Window, cursor x.Cursor) {
	w.WritePad(4)
	w.Write4b(uint32(window))
	w.Write4b(uint32(cursor))
}

type CompareCursorReply struct {
	Same bool
}

func readCompareCursorReply(r *x.Reader, v *CompareCursorReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Same = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// len
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}
	return nil
}

// #WREQ
func writeFakeInput(w *x.Writer, evType uint8, detail uint8, time x.Timestamp, root x.Window,
	rootX, rootY int16, deviceId uint8) {
	w.WritePad(4)

	w.Write1b(evType)
	w.Write1b(detail)
	w.WritePad(2)

	w.Write4b(uint32(time))
	w.Write4b(uint32(root))
	w.WritePad(8)

	w.Write2b(uint16(rootX))
	w.Write2b(uint16(rootY))

	w.WritePad(7)
	w.Write1b(deviceId)
}

// #WREQ
func writeGrabControl(w *x.Writer, impervious bool) {
	w.WritePad(4)

	w.Write1b(x.BoolToUint8(impervious))
	w.WritePad(3)
}
