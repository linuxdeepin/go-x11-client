package render

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

func QueryPictFormats(conn *x.Conn) QueryPictFormatsCookie {
	w := x.NewWriter()
	writeQueryPictFormats(w)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: QueryPictFormatsOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return QueryPictFormatsCookie(seq)
}

func (cookie QueryPictFormatsCookie) Reply(conn *x.Conn) (*QueryPictFormatsReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply QueryPictFormatsReply
	err = readQueryPictFormatsReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryPictIndexValues(conn *x.Conn, format PictFormat) QueryPictIndexValuesCookie {
	w := x.NewWriter()
	writeQueryPictIndexValues(w, format)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: QueryPictIndexValuesOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return QueryPictIndexValuesCookie(seq)
}

func (cookie QueryPictIndexValuesCookie) Reply(conn *x.Conn) (*QueryPictIndexValuesReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply QueryPictIndexValuesReply
	err = readQueryPictIndexValuesReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryFilters(conn *x.Conn, drawable x.Drawable) QueryFiltersCookie {
	w := x.NewWriter()
	writeQueryFilters(w, drawable)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: QueryFiltersOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return QueryFiltersCookie(seq)
}

func (cookie QueryFiltersCookie) Reply(conn *x.Conn) (*QueryFiltersReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply QueryFiltersReply
	err = readQueryFiltersReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func CreatePicture(conn *x.Conn, pid Picture, drawable x.Drawable, format PictFormat, valueMask uint32, valueList []uint32) {
	w := x.NewWriter()
	writeCreatePicture(w, pid, drawable, format, valueMask, valueList)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreatePictureOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CreatePictureChecked(conn *x.Conn, pid Picture, drawable x.Drawable, format PictFormat, valueMask uint32, valueList []uint32) x.VoidCookie {
	w := x.NewWriter()
	writeCreatePicture(w, pid, drawable, format, valueMask, valueList)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreatePictureOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func ChangePicture(conn *x.Conn, picture Picture, valueMask uint32, valueList []uint32) {
	w := x.NewWriter()
	writeChangePicture(w, picture, valueMask, valueList)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  ChangePictureOpcode,
	}
	conn.SendRequest(0, d, req)
}

func ChangePictureChecked(conn *x.Conn, picture Picture, valueMask uint32, valueList []uint32) x.VoidCookie {
	w := x.NewWriter()
	writeChangePicture(w, picture, valueMask, valueList)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  ChangePictureOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func SetPictureClipRectangles(conn *x.Conn, picture Picture, clipXOrigin, clipYOrigin int16, rectangles []x.Rectangle) {
	w := x.NewWriter()
	writeSetPictureClipRectangles(w, picture, clipXOrigin, clipYOrigin, rectangles)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SetPictureClipRectanglesOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SetPictureClipRectanglesChecked(conn *x.Conn, picture Picture, clipXOrigin, clipYOrigin int16, rectangles []x.Rectangle) x.VoidCookie {
	w := x.NewWriter()
	writeSetPictureClipRectangles(w, picture, clipXOrigin, clipYOrigin, rectangles)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SetPictureClipRectanglesOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func SetPictureTransform(conn *x.Conn, picture Picture, transform *Transform) {
	w := x.NewWriter()
	writeSetPictureTransform(w, picture, transform)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SetPictureTransformOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SetPictureTransformChecked(conn *x.Conn, picture Picture, transform *Transform) x.VoidCookie {
	w := x.NewWriter()
	writeSetPictureTransform(w, picture, transform)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SetPictureTransformOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func SetPictureFilter(conn *x.Conn, picture Picture, filter string, values []Fixed) {
	w := x.NewWriter()
	writeSetPictureFilter(w, picture, filter, values)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SetPictureFilterOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SetPictureFilterChecked(conn *x.Conn, picture Picture, filter string, values []Fixed) x.VoidCookie {
	w := x.NewWriter()
	writeSetPictureFilter(w, picture, filter, values)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SetPictureFilterOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func FreePicture(conn *x.Conn, picture Picture) {
	w := x.NewWriter()
	writeFreePicture(w, picture)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  FreePictureOpcode,
	}
	conn.SendRequest(0, d, req)
}

func FreePictureChecked(conn *x.Conn, picture Picture) x.VoidCookie {
	w := x.NewWriter()
	writeFreePicture(w, picture)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  FreePictureOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func Composite(conn *x.Conn, op uint8, src, mask, dst Picture, srcX, srcY, maskX, maskY, dstX, dstY int16, width, height uint16) {
	w := x.NewWriter()
	writeComposite(w, op, src, mask, dst, srcX, srcY, maskX, maskY, dstX, dstY, width, height)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CompositeOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CompositeChecked(conn *x.Conn, op uint8, src, mask, dst Picture, srcX, srcY, maskX, maskY, dstX, dstY int16, width, height uint16) x.VoidCookie {
	w := x.NewWriter()
	writeComposite(w, op, src, mask, dst, srcX, srcY, maskX, maskY, dstX, dstY, width, height)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CompositeOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func FillRectangles(conn *x.Conn, op uint8, dst Picture, color Color, rects []x.Rectangle) {
	w := x.NewWriter()
	writeFillRectangles(w, op, dst, color, rects)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  FillRectanglesOpcode,
	}
	conn.SendRequest(0, d, req)
}

func FillRectanglesChecked(conn *x.Conn, op uint8, dst Picture, color Color, rects []x.Rectangle) x.VoidCookie {
	w := x.NewWriter()
	writeFillRectangles(w, op, dst, color, rects)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  FillRectanglesOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func Trapezoids(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, srcX, srcY int16, traps []Trapezoid) {
	w := x.NewWriter()
	writeTrapezoids(w, op, src, dst, maskFormat, srcX, srcY, traps)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  TrapezoidsOpcode,
	}
	conn.SendRequest(0, d, req)
}

func TrapezoidsChecked(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, srcX, srcY int16, traps []Trapezoid) x.VoidCookie {
	w := x.NewWriter()
	writeTrapezoids(w, op, src, dst, maskFormat, srcX, srcY, traps)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  TrapezoidsOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func Triangles(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, srcX, srcY int16, triangles []Triangle) {
	w := x.NewWriter()
	writeTriangles(w, op, src, dst, maskFormat, srcX, srcY, triangles)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  TrianglesOpcode,
	}
	conn.SendRequest(0, d, req)
}

func TrianglesChecked(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, srcX, srcY int16, triangles []Triangle) x.VoidCookie {
	w := x.NewWriter()
	writeTriangles(w, op, src, dst, maskFormat, srcX, srcY, triangles)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  TrianglesOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func TriStrip(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, srcX, srcY int16, points []PointFixed) {
	w := x.NewWriter()
	writeTriStrip(w, op, src, dst, maskFormat, srcX, srcY, points)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  TriStripOpcode,
	}
	conn.SendRequest(0, d, req)
}

func TriStripChecked(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, srcX, srcY int16, points []PointFixed) x.VoidCookie {
	w := x.NewWriter()
	writeTriStrip(w, op, src, dst, maskFormat, srcX, srcY, points)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  TriStripOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func TriFan(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, srcX, srcY int16, points []PointFixed) {
	w := x.NewWriter()
	writeTriFan(w, op, src, dst, maskFormat, srcX, srcY, points)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  TriFanOpcode,
	}
	conn.SendRequest(0, d, req)
}

func TriFanChecked(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, srcX, srcY int16, points []PointFixed) x.VoidCookie {
	w := x.NewWriter()
	writeTriFan(w, op, src, dst, maskFormat, srcX, srcY, points)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  TriFanOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func CreateGlyphSet(conn *x.Conn, gsId GlyphSet, format PictFormat) {
	w := x.NewWriter()
	writeCreateGlyphSet(w, gsId, format)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateGlyphSetOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CreateGlyphSetChecked(conn *x.Conn, gsId GlyphSet, format PictFormat) x.VoidCookie {
	w := x.NewWriter()
	writeCreateGlyphSet(w, gsId, format)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateGlyphSetOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func ReferenceGlyphSet(conn *x.Conn, gsId, existing GlyphSet) {
	w := x.NewWriter()
	writeReferenceGlyphSet(w, gsId, existing)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  ReferenceGlyphSetOpcode,
	}
	conn.SendRequest(0, d, req)
}

func ReferenceGlyphSetChecked(conn *x.Conn, gsId, existing GlyphSet) x.VoidCookie {
	w := x.NewWriter()
	writeReferenceGlyphSet(w, gsId, existing)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  ReferenceGlyphSetOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func FreeGlyphSet(conn *x.Conn, glyphSet GlyphSet) {
	w := x.NewWriter()
	writeFreeGlyphSet(w, glyphSet)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  FreeGlyphSetOpcode,
	}
	conn.SendRequest(0, d, req)
}

func FreeGlyphSetChecked(conn *x.Conn, glyphSet GlyphSet) x.VoidCookie {
	w := x.NewWriter()
	writeFreeGlyphSet(w, glyphSet)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  FreeGlyphSetOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func AddGlyphs(conn *x.Conn, glyphSet GlyphSet, glyphIds []uint32, glyphs []GlyphInfo, data []byte) {
	w := x.NewWriter()
	writeAddGlyphs(w, glyphSet, glyphIds, glyphs, data)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  AddGlyphsOpcode,
	}
	conn.SendRequest(0, d, req)
}

func AddGlyphsChecked(conn *x.Conn, glyphSet GlyphSet, glyphIds []uint32, glyphs []GlyphInfo, data []byte) x.VoidCookie {
	w := x.NewWriter()
	writeAddGlyphs(w, glyphSet, glyphIds, glyphs, data)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  AddGlyphsOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func FreeGlyphs(conn *x.Conn, glyphSet GlyphSet, glyphs []Glyph) {
	w := x.NewWriter()
	writeFreeGlyphs(w, glyphSet, glyphs)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  FreeGlyphsOpcode,
	}
	conn.SendRequest(0, d, req)
}

func FreeGlyphsChecked(conn *x.Conn, glyphSet GlyphSet, glyphs []Glyph) x.VoidCookie {
	w := x.NewWriter()
	writeFreeGlyphs(w, glyphSet, glyphs)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  FreeGlyphsOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func CompositeGlyphs8(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) {
	w := x.NewWriter()
	writeCompositeGlyphs8(w, op, src, dst, maskFormat, glyphSet, srcX, srcY, glyphCmds)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CompositeGlyphs8Opcode,
	}
	conn.SendRequest(0, d, req)
}

func CompositeGlyphs8Checked(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) x.VoidCookie {
	w := x.NewWriter()
	writeCompositeGlyphs8(w, op, src, dst, maskFormat, glyphSet, srcX, srcY, glyphCmds)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CompositeGlyphs8Opcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func CompositeGlyphs16(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) {
	w := x.NewWriter()
	writeCompositeGlyphs16(w, op, src, dst, maskFormat, glyphSet, srcX, srcY, glyphCmds)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CompositeGlyphs16Opcode,
	}
	conn.SendRequest(0, d, req)
}

func CompositeGlyphs16Checked(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) x.VoidCookie {
	w := x.NewWriter()
	writeCompositeGlyphs16(w, op, src, dst, maskFormat, glyphSet, srcX, srcY, glyphCmds)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CompositeGlyphs16Opcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func CompositeGlyphs32(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) {
	w := x.NewWriter()
	writeCompositeGlyphs32(w, op, src, dst, maskFormat, glyphSet, srcX, srcY, glyphCmds)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CompositeGlyphs32Opcode,
	}
	conn.SendRequest(0, d, req)
}

func CompositeGlyphs32Checked(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) x.VoidCookie {
	w := x.NewWriter()
	writeCompositeGlyphs32(w, op, src, dst, maskFormat, glyphSet, srcX, srcY, glyphCmds)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CompositeGlyphs32Opcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func CreateCursor(conn *x.Conn, cid x.Cursor, source Picture, X, y uint16) {
	w := x.NewWriter()
	writeCreateCursor(w, cid, source, X, y)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateCursorOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CreateCursorChecked(conn *x.Conn, cid x.Cursor, source Picture, X, y uint16) x.VoidCookie {
	w := x.NewWriter()
	writeCreateCursor(w, cid, source, X, y)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateCursorOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func CreateAnimCursor(conn *x.Conn, cid x.Cursor, cursors []AnimCursorElt) {
	w := x.NewWriter()
	writeCreateAnimCursor(w, cid, cursors)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateAnimCursorOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CreateAnimCursorChecked(conn *x.Conn, cid x.Cursor, cursors []AnimCursorElt) x.VoidCookie {
	w := x.NewWriter()
	writeCreateAnimCursor(w, cid, cursors)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateAnimCursorOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func AddTraps(conn *x.Conn, picture Picture, xOff, yOff int16, traps []Trap) {
	w := x.NewWriter()
	writeAddTraps(w, picture, xOff, yOff, traps)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  AddTrapsOpcode,
	}
	conn.SendRequest(0, d, req)
}

func AddTrapsChecked(conn *x.Conn, picture Picture, xOff, yOff int16, traps []Trap) x.VoidCookie {
	w := x.NewWriter()
	writeAddTraps(w, picture, xOff, yOff, traps)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  AddTrapsOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func CreateSolidFill(conn *x.Conn, pid Picture, color Color) {
	w := x.NewWriter()
	writeCreateSolidFill(w, pid, color)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateSolidFillOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CreateSolidFillChecked(conn *x.Conn, pid Picture, color Color) x.VoidCookie {
	w := x.NewWriter()
	writeCreateSolidFill(w, pid, color)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateSolidFillOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func CreateLinearGradient(conn *x.Conn, picture Picture, p1, p2 PointFixed, stops []Fixed, colors []Color) {
	w := x.NewWriter()
	writeCreateLinearGradient(w, picture, p1, p2, stops, colors)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateLinearGradientOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CreateLinearGradientChecked(conn *x.Conn, picture Picture, p1, p2 PointFixed, stops []Fixed, colors []Color) x.VoidCookie {
	w := x.NewWriter()
	writeCreateLinearGradient(w, picture, p1, p2, stops, colors)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateLinearGradientOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func CreateRadialGradient(conn *x.Conn, picture Picture, inner, outer PointFixed, innerRadius, outerRadius Fixed, stops []Fixed, colors []Color) {
	w := x.NewWriter()
	writeCreateRadialGradient(w, picture, inner, outer, innerRadius, outerRadius, stops, colors)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateRadialGradientOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CreateRadialGradientChecked(conn *x.Conn, picture Picture, inner, outer PointFixed, innerRadius, outerRadius Fixed, stops []Fixed, colors []Color) x.VoidCookie {
	w := x.NewWriter()
	writeCreateRadialGradient(w, picture, inner, outer, innerRadius, outerRadius, stops, colors)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateRadialGradientOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func CreateConicalGradient(conn *x.Conn, picture Picture, center PointFixed, angle Fixed, stops []Fixed, colors []Color) {
	w := x.NewWriter()
	writeCreateConicalGradient(w, picture, center, angle, stops, colors)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateConicalGradientOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CreateConicalGradientChecked(conn *x.Conn, picture Picture, center PointFixed, angle Fixed, stops []Fixed, colors []Color) x.VoidCookie {
	w := x.NewWriter()
	writeCreateConicalGradient(w, picture, center, angle, stops, colors)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateConicalGradientOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}
