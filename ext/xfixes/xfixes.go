package xfixes

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

	r.Read1b()
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

	v.MajorVersion = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MinorVersion = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(16)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeChangeSaveSet(w *x.Writer, mode, target, map0 uint8, window x.Window) {
	w.WritePad(4)

	w.Write1b(mode)
	w.Write1b(target)
	w.Write1b(map0)
	w.WritePad(1)

	w.Write4b(uint32(window))
}

type SelectionNotifyEvent struct {
	Subtype            uint8
	Window             x.Window
	Owner              x.Window
	Selection          x.Atom
	Timestamp          x.Timestamp
	SelectionTimestamp x.Timestamp
}

func readSelectionNotifyEvent(r *x.Reader, v *SelectionNotifyEvent) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Subtype = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Owner = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Selection = x.Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Timestamp = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.SelectionTimestamp = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(8)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeSelectSelectionInput(w *x.Writer, window x.Window, selection x.Atom, eventMask uint32) {
	w.WritePad(4)
	w.Write4b(uint32(window))
	w.Write4b(uint32(selection))
	w.Write4b(eventMask)
}

type CursorNotifyEvent struct {
	Subtype      uint8
	Window       x.Window
	CursorSerial uint32
	Timestamp    x.Timestamp
	Name         x.Atom
}

func readCursorNotifyEvent(r *x.Reader, v *CursorNotifyEvent) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Subtype = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.CursorSerial = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Timestamp = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Name = x.Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(12)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeSelectCursorInput(w *x.Writer, window x.Window, eventMask uint32) {
	w.WritePad(4)
	w.Write4b(uint32(window))
	w.Write4b(eventMask)
}

// #WREQ
func writeGetCursorImage(w *x.Writer) {
	w.WritePad(4)
}

type GetCursorImageReply struct {
	X            int16
	Y            int16
	Width        uint16
	Height       uint16
	XHot         uint16
	YHot         uint16
	CursorSerial uint32
	CursorImage  []uint32
}

func readGetCursorImageReply(r *x.Reader, v *GetCursorImageReply) error {
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

	// len
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.X = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Y = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Width = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Height = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.XHot = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.YHot = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.CursorSerial = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	dataLen := int(v.Width) * int(v.Height)
	if dataLen > 0 {
		v.CursorImage = make([]uint32, dataLen)
		for i := 0; i < dataLen; i++ {
			v.CursorImage[i] = r.Read4b()
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	return nil
}

// #WREQ
func writeCreateRegion(w *x.Writer, region Region, rects []x.Rectangle) {
	w.WritePad(4)
	w.Write4b(uint32(region))
	for _, rect := range rects {
		x.WriteRectangle(w, rect)
	}
}

// #WREQ
func writeCreateRegionFromBitmap(w *x.Writer, region Region, bitmap x.Pixmap) {
	w.WritePad(4)
	w.Write4b(uint32(region))
	w.Write4b(uint32(bitmap))
}

// #WREQ
//func writeCreateRegionFromWindow

// #WREQ
func writeCreateRegionFromGC(w *x.Writer, region Region, gc x.GContext) {
	w.WritePad(4)
	w.Write4b(uint32(region))
	w.Write4b(uint32(gc))
}

// #WREQ
//func writeCreateRegionFromPicture

// #WREQ
func writeDestroyRegion(w *x.Writer, region Region) {
	w.WritePad(4)
	w.Write4b(uint32(region))
}

// #WREQ
func writeSetRegion(w *x.Writer, region Region, rects []x.Rectangle) {
	w.WritePad(4)
	w.Write4b(uint32(region))
	for _, rect := range rects {
		x.WriteRectangle(w, rect)
	}
}

// #WREQ
func writeCopyRegion(w *x.Writer, source, destination Region) {
	w.WritePad(4)
	w.Write4b(uint32(source))
	w.Write4b(uint32(destination))
}

// #WREQ
func writeUnionRegion(w *x.Writer, source1, source2, destination Region) {
	w.WritePad(4)
	w.Write4b(uint32(source1))
	w.Write4b(uint32(source2))
	w.Write4b(uint32(destination))
}

// #WREQ
func writeIntersectRegion(w *x.Writer, source1, source2, destination Region) {
	w.WritePad(4)
	w.Write4b(uint32(source1))
	w.Write4b(uint32(source2))
	w.Write4b(uint32(destination))
}

// #WREQ
func writeSubtractRegion(w *x.Writer, source1, source2, destination Region) {
	w.WritePad(4)
	w.Write4b(uint32(source1))
	w.Write4b(uint32(source2))
	w.Write4b(uint32(destination))
}

// #WREQ
func writeInvertRegion(w *x.Writer, source Region, bounds x.Rectangle, destination Region) {
	w.WritePad(4)
	w.Write4b(uint32(source))
	x.WriteRectangle(w, bounds)
	w.Write4b(uint32(destination))
}

// #WREQ
func writeTranslateRegion(w *x.Writer, region Region, dx, dy int16) {
	w.WritePad(4)
	w.Write4b(uint32(region))
	w.Write2b(uint16(dx))
	w.Write2b(uint16(dy))
}

// #WREQ
func writeRegionExtents(w *x.Writer, source, destination Region) {
	w.WritePad(4)
	w.Write4b(uint32(source))
	w.Write4b(uint32(destination))
}

// #WREQ
func writeFetchRegion(w *x.Writer, region Region) {
	w.WritePad(4)
	w.Write4b(uint32(region))
}

type FetchRegionReply struct {
	Extents    x.Rectangle
	Rectangles []x.Rectangle
}

func readFetchRegionReply(r *x.Reader, v *FetchRegionReply) error {
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

	// len
	replyLen := int(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	var err error
	v.Extents, err = x.ReadRectangle(r)
	if err != nil {
		return err
	}

	r.ReadPad(16)
	if err != nil {
		return err
	}

	rectanglesLen := replyLen / 2
	if rectanglesLen > 0 {
		v.Rectangles = make([]x.Rectangle, rectanglesLen)
		for i := 0; i < rectanglesLen; i++ {
			v.Rectangles[i], err = x.ReadRectangle(r)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// #WREQ
func writeSetGCClipRegion(w *x.Writer, gc x.GContext, clipXOrigin, clipYOrigin int16,
	region Region) {

	w.WritePad(4)
	w.Write4b(uint32(gc))
	w.Write4b(uint32(region))
	w.Write2b(uint16(clipXOrigin))
	w.Write2b(uint16(clipYOrigin))
}

// #WREQ
func writeSetCursorName(w *x.Writer, cursor x.Cursor, name string) {
	w.WritePad(4)
	w.Write4b(uint32(cursor))
	nameLen := len(name)
	if nameLen > 0xffff {
		name = name[:0xffff]
		nameLen = 0xffff
	}
	w.Write2b(uint16(nameLen))
	w.WriteString(name)
	w.WritePad(x.Pad(nameLen))
}

// #WREQ
func writeGetCursorImageAndName(w *x.Writer) {
	w.WritePad(4)
}

type GetCursorImageAndNameReply struct {
	X, Y          int16
	Width, Height uint16
	XHot, YHot    uint16
	CursorSerial  uint32
	CursorAtom    x.Atom
	CursorName    string
	CursorImage   []uint32
}

func readGetCursorImageAndNameReply(r *x.Reader, v *GetCursorImageAndNameReply) error {
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

	// len
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.X = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Y = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Width = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Height = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.XHot = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.YHot = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.CursorSerial = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.CursorAtom = x.Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	cursorNameLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(2)
	if r.Err() != nil {
		return r.Err()
	}

	cursorImageLen := int(v.Width) * int(v.Height)
	if cursorImageLen > 0 {
		v.CursorImage = make([]uint32, cursorImageLen)
		for i := 0; i < cursorImageLen; i++ {
			v.CursorImage[i] = r.Read4b()
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	v.CursorName = r.ReadString(cursorNameLen)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeChangeCursor(w *x.Writer, source, destination x.Cursor) {
	w.WritePad(4)
	w.Write4b(uint32(source))
	w.Write4b(uint32(destination))
}

// #WREQ
func writeChangeCursorByName(w *x.Writer, src x.Cursor, name string) {
	w.WritePad(4)
	w.Write4b(uint32(src))
	nameLen := len(name)
	if nameLen > 0xffff {
		name = name[:0xffff]
		nameLen = 0xffff
	}
	w.Write2b(uint16(nameLen))
	w.WritePad(2)

	w.WriteString(name)
	w.WritePad(x.Pad(nameLen))
}

// #WREQ
func writeExpandRegion(w *x.Writer, source, destination Region, left, right,
	top, bottom uint16) {

	w.WritePad(4)
	w.Write4b(uint32(source))
	w.Write4b(uint32(destination))
	w.Write2b(left)
	w.Write2b(right)
	w.Write2b(top)
	w.Write2b(bottom)
}

// #WREQ
func writeHideCursor(w *x.Writer, window x.Window) {
	w.WritePad(4)
	w.Write4b(uint32(window))
}

// #WREQ
func writeShowCursor(w *x.Writer, window x.Window) {
	w.WritePad(4)
	w.Write4b(uint32(window))
}

// #WREQ
func writeCreatePointerBarrier(w *x.Writer, barrier Barrier, drawable x.Drawable,
	x1, y1, x2, y2 int16, directions uint32, devices []uint16) {
	w.WritePad(4)
	w.Write4b(uint32(barrier))
	w.Write4b(uint32(drawable))
	w.Write2b(uint16(x1))
	w.Write2b(uint16(y1))
	w.Write2b(uint16(x2))
	w.Write2b(uint16(y2))
	w.Write4b(directions)
	w.WritePad(2)
	w.Write2b(uint16(len(devices)))
	for _, device := range devices {
		w.Write2b(device)
	}
}

// #WREQ
func writeDeletePointerBarrier(w *x.Writer, barrier Barrier) {
	w.WritePad(4)
	w.Write4b(uint32(barrier))
}

type BadRegionError struct {
	x.ResourceErrorBase
}

func readBadRegionError(r *x.Reader) x.Error {
	return BadRegionError{x.ReadResourceErrorBase(r)}
}

func (e BadRegionError) Error() string {
	return "BadRegionError" + e.Msg()
}
