package xfixes

import x "github.com/linuxdeepin/go-x11-client"

func QueryVersion(conn *x.Conn, majorVersion, minorVersion uint32) QueryVersionCookie {
	w := x.NewWriter()
	writeQueryVersion(w, majorVersion, minorVersion)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: QueryVersionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return QueryVersionCookie(seq)
}

func (cookie QueryVersionCookie) Reply(conn *x.Conn) (*QueryVersionReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply QueryVersionReply
	err = readQueryVersionReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func ChangeSaveSet(conn *x.Conn, mode, target, map0 uint8, window x.Window) {
	w := x.NewWriter()
	writeChangeSaveSet(w, mode, target, map0, window)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  ChangeSaveSetOpcode,
	}
	conn.SendRequest(0, d, req)
}

func ChangeSaveSetChecked(conn *x.Conn, mode, target, map0 uint8, window x.Window) x.VoidCookie {
	w := x.NewWriter()
	writeChangeSaveSet(w, mode, target, map0, window)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  ChangeSaveSetOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func SelectSelectionInput(conn *x.Conn, window x.Window, selection x.Atom, eventMask uint32) {
	w := x.NewWriter()
	writeSelectSelectionInput(w, window, selection, eventMask)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SelectSelectionInputOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SelectSelectionInputChecked(conn *x.Conn, window x.Window, selection x.Atom, eventMask uint32) x.VoidCookie {
	w := x.NewWriter()
	writeSelectSelectionInput(w, window, selection, eventMask)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SelectSelectionInputOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func SelectCursorInput(conn *x.Conn, window x.Window, eventMask uint32) {
	w := x.NewWriter()
	writeSelectCursorInput(w, window, eventMask)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SelectCursorInputOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SelectCursorInputChecked(conn *x.Conn, window x.Window, eventMask uint32) x.VoidCookie {
	w := x.NewWriter()
	writeSelectCursorInput(w, window, eventMask)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SelectCursorInputOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func GetCursorImage(conn *x.Conn) GetCursorImageCookie {
	w := x.NewWriter()
	writeGetCursorImage(w)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: GetCursorImageOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return GetCursorImageCookie(seq)
}

func (cookie GetCursorImageCookie) Reply(conn *x.Conn) (*GetCursorImageReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetCursorImageReply
	err = readGetCursorImageReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func CreateRegion(conn *x.Conn, region Region, rects []x.Rectangle) {
	w := x.NewWriter()
	writeCreateRegion(w, region, rects)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateRegionOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CreateRegionChecked(conn *x.Conn, region Region, rects []x.Rectangle) x.VoidCookie {
	w := x.NewWriter()
	writeCreateRegion(w, region, rects)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateRegionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func CreateRegionFromBitmap(conn *x.Conn, region Region, bitmap x.Pixmap) {
	w := x.NewWriter()
	writeCreateRegionFromBitmap(w, region, bitmap)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateRegionFromBitmapOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CreateRegionFromBitmapChecked(conn *x.Conn, region Region, bitmap x.Pixmap) x.VoidCookie {
	w := x.NewWriter()
	writeCreateRegionFromBitmap(w, region, bitmap)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateRegionFromBitmapOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func CreateRegionFromGC(conn *x.Conn, region Region, gc x.GContext) {
	w := x.NewWriter()
	writeCreateRegionFromGC(w, region, gc)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateRegionFromGCOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CreateRegionFromGCChecked(conn *x.Conn, region Region, gc x.GContext) x.VoidCookie {
	w := x.NewWriter()
	writeCreateRegionFromGC(w, region, gc)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateRegionFromGCOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func DestroyRegion(conn *x.Conn, region Region) {
	w := x.NewWriter()
	writeDestroyRegion(w, region)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  DestroyRegionOpcode,
	}
	conn.SendRequest(0, d, req)
}

func DestroyRegionChecked(conn *x.Conn, region Region) x.VoidCookie {
	w := x.NewWriter()
	writeDestroyRegion(w, region)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  DestroyRegionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func SetRegion(conn *x.Conn, region Region, rects []x.Rectangle) {
	w := x.NewWriter()
	writeSetRegion(w, region, rects)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SetRegionOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SetRegionChecked(conn *x.Conn, region Region, rects []x.Rectangle) x.VoidCookie {
	w := x.NewWriter()
	writeSetRegion(w, region, rects)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SetRegionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func CopyRegion(conn *x.Conn, source, destination Region) {
	w := x.NewWriter()
	writeCopyRegion(w, source, destination)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CopyRegionOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CopyRegionChecked(conn *x.Conn, source, destination Region) x.VoidCookie {
	w := x.NewWriter()
	writeCopyRegion(w, source, destination)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CopyRegionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func UnionRegion(conn *x.Conn, source1, source2, destination Region) {
	w := x.NewWriter()
	writeUnionRegion(w, source1, source2, destination)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  UnionRegionOpcode,
	}
	conn.SendRequest(0, d, req)
}

func UnionRegionChecked(conn *x.Conn, source1, source2, destination Region) x.VoidCookie {
	w := x.NewWriter()
	writeUnionRegion(w, source1, source2, destination)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  UnionRegionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func IntersectRegion(conn *x.Conn, source1, source2, destination Region) {
	w := x.NewWriter()
	writeIntersectRegion(w, source1, source2, destination)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  IntersectRegionOpcode,
	}
	conn.SendRequest(0, d, req)
}

func IntersectRegionChecked(conn *x.Conn, source1, source2, destination Region) x.VoidCookie {
	w := x.NewWriter()
	writeIntersectRegion(w, source1, source2, destination)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  IntersectRegionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func SubtractRegion(conn *x.Conn, source1, source2, destination Region) {
	w := x.NewWriter()
	writeSubtractRegion(w, source1, source2, destination)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SubtractRegionOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SubtractRegionChecked(conn *x.Conn, source1, source2, destination Region) x.VoidCookie {
	w := x.NewWriter()
	writeSubtractRegion(w, source1, source2, destination)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SubtractRegionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func InvertRegion(conn *x.Conn, source Region, bounds x.Rectangle, destination Region) {
	w := x.NewWriter()
	writeInvertRegion(w, source, bounds, destination)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  InvertRegionOpcode,
	}
	conn.SendRequest(0, d, req)
}

func InvertRegionChecked(conn *x.Conn, source Region, bounds x.Rectangle, destination Region) x.VoidCookie {
	w := x.NewWriter()
	writeInvertRegion(w, source, bounds, destination)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  InvertRegionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func TranslateRegion(conn *x.Conn, region Region, dx, dy int16) {
	w := x.NewWriter()
	writeTranslateRegion(w, region, dx, dy)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  TranslateRegionOpcode,
	}
	conn.SendRequest(0, d, req)
}

func TranslateRegionChecked(conn *x.Conn, region Region, dx, dy int16) x.VoidCookie {
	w := x.NewWriter()
	writeTranslateRegion(w, region, dx, dy)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  TranslateRegionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func RegionExtents(conn *x.Conn, source, destination Region) {
	w := x.NewWriter()
	writeRegionExtents(w, source, destination)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  RegionExtentsOpcode,
	}
	conn.SendRequest(0, d, req)
}

func RegionExtentsChecked(conn *x.Conn, source, destination Region) x.VoidCookie {
	w := x.NewWriter()
	writeRegionExtents(w, source, destination)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  RegionExtentsOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func FetchRegion(conn *x.Conn, region Region) FetchRegionCookie {
	w := x.NewWriter()
	writeFetchRegion(w, region)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: FetchRegionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return FetchRegionCookie(seq)
}

func (cookie FetchRegionCookie) Reply(conn *x.Conn) (*FetchRegionReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply FetchRegionReply
	err = readFetchRegionReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetGCClipRegion(conn *x.Conn, gc x.GContext, clipXOrigin, clipYOrigin int16, region Region) {
	w := x.NewWriter()
	writeSetGCClipRegion(w, gc, clipXOrigin, clipYOrigin, region)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SetGCClipRegionOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SetGCClipRegionChecked(conn *x.Conn, gc x.GContext, clipXOrigin, clipYOrigin int16, region Region) x.VoidCookie {
	w := x.NewWriter()
	writeSetGCClipRegion(w, gc, clipXOrigin, clipYOrigin, region)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SetGCClipRegionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func SetCursorName(conn *x.Conn, cursor x.Cursor, name string) {
	w := x.NewWriter()
	writeSetCursorName(w, cursor, name)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SetCursorNameOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SetCursorNameChecked(conn *x.Conn, cursor x.Cursor, name string) x.VoidCookie {
	w := x.NewWriter()
	writeSetCursorName(w, cursor, name)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SetCursorNameOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func GetCursorImageAndName(conn *x.Conn) GetCursorImageAndNameCookie {
	w := x.NewWriter()
	writeGetCursorImageAndName(w)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: GetCursorImageAndNameOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return GetCursorImageAndNameCookie(seq)
}

func (cookie GetCursorImageAndNameCookie) Reply(conn *x.Conn) (*GetCursorImageAndNameReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetCursorImageAndNameReply
	err = readGetCursorImageAndNameReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func ChangeCursor(conn *x.Conn, source, destination x.Cursor) {
	w := x.NewWriter()
	writeChangeCursor(w, source, destination)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  ChangeCursorOpcode,
	}
	conn.SendRequest(0, d, req)
}

func ChangeCursorChecked(conn *x.Conn, source, destination x.Cursor) x.VoidCookie {
	w := x.NewWriter()
	writeChangeCursor(w, source, destination)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  ChangeCursorOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func ChangeCursorByName(conn *x.Conn, src x.Cursor, name string) {
	w := x.NewWriter()
	writeChangeCursorByName(w, src, name)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  ChangeCursorByNameOpcode,
	}
	conn.SendRequest(0, d, req)
}

func ChangeCursorByNameChecked(conn *x.Conn, src x.Cursor, name string) x.VoidCookie {
	w := x.NewWriter()
	writeChangeCursorByName(w, src, name)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  ChangeCursorByNameOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func ExpandRegion(conn *x.Conn, source, destination Region, left, right, top, bottom uint16) {
	w := x.NewWriter()
	writeExpandRegion(w, source, destination, left, right, top, bottom)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  ExpandRegionOpcode,
	}
	conn.SendRequest(0, d, req)
}

func ExpandRegionChecked(conn *x.Conn, source, destination Region, left, right, top, bottom uint16) x.VoidCookie {
	w := x.NewWriter()
	writeExpandRegion(w, source, destination, left, right, top, bottom)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  ExpandRegionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func HideCursor(conn *x.Conn, window x.Window) {
	w := x.NewWriter()
	writeHideCursor(w, window)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  HideCursorOpcode,
	}
	conn.SendRequest(0, d, req)
}

func HideCursorChecked(conn *x.Conn, window x.Window) x.VoidCookie {
	w := x.NewWriter()
	writeHideCursor(w, window)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  HideCursorOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func ShowCursor(conn *x.Conn, window x.Window) {
	w := x.NewWriter()
	writeShowCursor(w, window)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  ShowCursorOpcode,
	}
	conn.SendRequest(0, d, req)
}

func ShowCursorChecked(conn *x.Conn, window x.Window) x.VoidCookie {
	w := x.NewWriter()
	writeShowCursor(w, window)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  ShowCursorOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func CreatePointerBarrier(conn *x.Conn, barrier Barrier, drawable x.Drawable, x1, y1, x2, y2 int16, directions uint32, devices []uint16) {
	w := x.NewWriter()
	writeCreatePointerBarrier(w, barrier, drawable, x1, y1, x2, y2, directions, devices)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreatePointerBarrierOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CreatePointerBarrierChecked(conn *x.Conn, barrier Barrier, drawable x.Drawable, x1, y1, x2, y2 int16, directions uint32, devices []uint16) x.VoidCookie {
	w := x.NewWriter()
	writeCreatePointerBarrier(w, barrier, drawable, x1, y1, x2, y2, directions, devices)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreatePointerBarrierOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func DeletePointerBarrier(conn *x.Conn, barrier Barrier) {
	w := x.NewWriter()
	writeDeletePointerBarrier(w, barrier)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  DeletePointerBarrierOpcode,
	}
	conn.SendRequest(0, d, req)
}

func DeletePointerBarrierChecked(conn *x.Conn, barrier Barrier) x.VoidCookie {
	w := x.NewWriter()
	writeDeletePointerBarrier(w, barrier)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  DeletePointerBarrierOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}
