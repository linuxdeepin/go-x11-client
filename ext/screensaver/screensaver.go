package screensaver

import "github.com/linuxdeepin/go-x11-client"

type NotifyEvent struct {
	State    uint8
	Sequence uint16
	Time     x.Timestamp
	Root     x.Window
	Window   x.Window
	Kind     uint8
	Forced   bool
}

func readNotifyEvent(r *x.Reader, v *NotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.State = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Time = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Root = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Kind = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Forced = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(14)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeQueryVersion(w *x.Writer, clientMajorVersion, clientMinorVersion uint8) {
	w.WritePad(4)
	w.Write1b(clientMajorVersion)
	w.Write1b(clientMinorVersion)
	w.WritePad(2)
}

type QueryVersionReply struct {
	ServerMajorVersion uint16
	ServerMinorVersion uint16
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

	v.ServerMajorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ServerMinorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(20)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeQueryInfo(w *x.Writer, drawable x.Drawable) {
	w.WritePad(4)
	w.Write4b(uint32(drawable))
}

type QueryInfoReply struct {
	State            uint8
	SaverWindow      x.Window
	MsUntilServer    uint32
	MsSinceUserInput uint32
	EventMask        uint32
	Kind             uint8
}

func readQueryInfoReply(r *x.Reader, v *QueryInfoReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.State = r.Read1b()
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

	v.SaverWindow = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.MsUntilServer = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MsSinceUserInput = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.EventMask = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Kind = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// read field Pad0
	r.ReadPad(7)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeSelectInput(w *x.Writer, drawable x.Drawable, eventMask uint32) {
	w.WritePad(4)
	w.Write4b(uint32(drawable))
	w.Write4b(eventMask)
}

// #WREQ
func writeSetAttributes(w *x.Writer, drawable x.Drawable, X, y int16, width, height,
	boardWidth uint16, class, depth uint8, visual x.VisualID, valueMask uint32, valueList []uint32) {
	w.WritePad(4)
	w.Write4b(uint32(drawable))

	w.Write2b(uint16(X))
	w.Write2b(uint16(y))

	w.Write2b(width)
	w.Write2b(height)

	w.Write2b(boardWidth)
	w.Write1b(class)
	w.Write1b(depth)

	w.Write4b(uint32(visual))
	w.Write4b(valueMask)

	for _, value := range valueList {
		w.Write4b(value)
	}
}

// #WREQ
func writeUnsetAttributes(w *x.Writer, drawable x.Drawable) {
	w.WritePad(4)
	w.Write4b(uint32(drawable))
}

// #WREQ
func writeSuspend(w *x.Writer, suspend bool) {
	w.WritePad(4)
	w.Write1b(x.BoolToUint8(suspend))
	w.WritePad(3)
}
