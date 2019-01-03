package render

import x "github.com/linuxdeepin/go-x11-client"

func QueryVersion(conn *x.Conn, majorVersion, minorVersion uint32) QueryVersionCookie {
	body := encodeQueryVersion(majorVersion, minorVersion)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: QueryVersionOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return QueryVersionCookie(seq)
}

func (cookie QueryVersionCookie) Reply(conn *x.Conn) (*QueryVersionReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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
	body := encodeQueryPictFormats()
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: QueryPictFormatsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return QueryPictFormatsCookie(seq)
}

func (cookie QueryPictFormatsCookie) Reply(conn *x.Conn) (*QueryPictFormatsReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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
	body := encodeQueryPictIndexValues(format)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: QueryPictIndexValuesOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return QueryPictIndexValuesCookie(seq)
}

func (cookie QueryPictIndexValuesCookie) Reply(conn *x.Conn) (*QueryPictIndexValuesReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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
	body := encodeQueryFilters(drawable)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: QueryFiltersOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return QueryFiltersCookie(seq)
}

func (cookie QueryFiltersCookie) Reply(conn *x.Conn) (*QueryFiltersReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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
	body := encodeCreatePicture(pid, drawable, format, valueMask, valueList)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreatePictureOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CreatePictureChecked(conn *x.Conn, pid Picture, drawable x.Drawable, format PictFormat, valueMask uint32, valueList []uint32) x.VoidCookie {
	body := encodeCreatePicture(pid, drawable, format, valueMask, valueList)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreatePictureOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func ChangePicture(conn *x.Conn, picture Picture, valueMask uint32, valueList []uint32) {
	body := encodeChangePicture(picture, valueMask, valueList)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ChangePictureOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func ChangePictureChecked(conn *x.Conn, picture Picture, valueMask uint32, valueList []uint32) x.VoidCookie {
	body := encodeChangePicture(picture, valueMask, valueList)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ChangePictureOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func SetPictureClipRectangles(conn *x.Conn, picture Picture, clipXOrigin, clipYOrigin int16, rectangles []x.Rectangle) {
	body := encodeSetPictureClipRectangles(picture, clipXOrigin, clipYOrigin, rectangles)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetPictureClipRectanglesOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetPictureClipRectanglesChecked(conn *x.Conn, picture Picture, clipXOrigin, clipYOrigin int16, rectangles []x.Rectangle) x.VoidCookie {
	body := encodeSetPictureClipRectangles(picture, clipXOrigin, clipYOrigin, rectangles)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetPictureClipRectanglesOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func SetPictureTransform(conn *x.Conn, picture Picture, transform *Transform) {
	body := encodeSetPictureTransform(picture, transform)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetPictureTransformOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetPictureTransformChecked(conn *x.Conn, picture Picture, transform *Transform) x.VoidCookie {
	body := encodeSetPictureTransform(picture, transform)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetPictureTransformOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func SetPictureFilter(conn *x.Conn, picture Picture, filter string, values []Fixed) {
	body := encodeSetPictureFilter(picture, filter, values)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetPictureFilterOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetPictureFilterChecked(conn *x.Conn, picture Picture, filter string, values []Fixed) x.VoidCookie {
	body := encodeSetPictureFilter(picture, filter, values)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetPictureFilterOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func FreePicture(conn *x.Conn, picture Picture) {
	body := encodeFreePicture(picture)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: FreePictureOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func FreePictureChecked(conn *x.Conn, picture Picture) x.VoidCookie {
	body := encodeFreePicture(picture)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: FreePictureOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func Composite(conn *x.Conn, op uint8, src, mask, dst Picture, srcX, srcY, maskX, maskY, dstX, dstY int16, width, height uint16) {
	body := encodeComposite(op, src, mask, dst, srcX, srcY, maskX, maskY, dstX, dstY, width, height)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CompositeOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CompositeChecked(conn *x.Conn, op uint8, src, mask, dst Picture, srcX, srcY, maskX, maskY, dstX, dstY int16, width, height uint16) x.VoidCookie {
	body := encodeComposite(op, src, mask, dst, srcX, srcY, maskX, maskY, dstX, dstY, width, height)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CompositeOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func FillRectangles(conn *x.Conn, op uint8, dst Picture, color Color, rects []x.Rectangle) {
	body := encodeFillRectangles(op, dst, color, rects)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: FillRectanglesOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func FillRectanglesChecked(conn *x.Conn, op uint8, dst Picture, color Color, rects []x.Rectangle) x.VoidCookie {
	body := encodeFillRectangles(op, dst, color, rects)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: FillRectanglesOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func Trapezoids(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, srcX, srcY int16, traps []Trapezoid) {
	body := encodeTrapezoids(op, src, dst, maskFormat, srcX, srcY, traps)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: TrapezoidsOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func TrapezoidsChecked(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, srcX, srcY int16, traps []Trapezoid) x.VoidCookie {
	body := encodeTrapezoids(op, src, dst, maskFormat, srcX, srcY, traps)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: TrapezoidsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func Triangles(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, srcX, srcY int16, triangles []Triangle) {
	body := encodeTriangles(op, src, dst, maskFormat, srcX, srcY, triangles)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: TrianglesOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func TrianglesChecked(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, srcX, srcY int16, triangles []Triangle) x.VoidCookie {
	body := encodeTriangles(op, src, dst, maskFormat, srcX, srcY, triangles)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: TrianglesOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func TriStrip(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, srcX, srcY int16, points []PointFixed) {
	body := encodeTriStrip(op, src, dst, maskFormat, srcX, srcY, points)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: TriStripOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func TriStripChecked(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, srcX, srcY int16, points []PointFixed) x.VoidCookie {
	body := encodeTriStrip(op, src, dst, maskFormat, srcX, srcY, points)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: TriStripOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func TriFan(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, srcX, srcY int16, points []PointFixed) {
	body := encodeTriFan(op, src, dst, maskFormat, srcX, srcY, points)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: TriFanOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func TriFanChecked(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, srcX, srcY int16, points []PointFixed) x.VoidCookie {
	body := encodeTriFan(op, src, dst, maskFormat, srcX, srcY, points)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: TriFanOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func CreateGlyphSet(conn *x.Conn, gsId GlyphSet, format PictFormat) {
	body := encodeCreateGlyphSet(gsId, format)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateGlyphSetOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CreateGlyphSetChecked(conn *x.Conn, gsId GlyphSet, format PictFormat) x.VoidCookie {
	body := encodeCreateGlyphSet(gsId, format)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateGlyphSetOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func ReferenceGlyphSet(conn *x.Conn, gsId, existing GlyphSet) {
	body := encodeReferenceGlyphSet(gsId, existing)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ReferenceGlyphSetOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func ReferenceGlyphSetChecked(conn *x.Conn, gsId, existing GlyphSet) x.VoidCookie {
	body := encodeReferenceGlyphSet(gsId, existing)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ReferenceGlyphSetOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func FreeGlyphSet(conn *x.Conn, glyphSet GlyphSet) {
	body := encodeFreeGlyphSet(glyphSet)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: FreeGlyphSetOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func FreeGlyphSetChecked(conn *x.Conn, glyphSet GlyphSet) x.VoidCookie {
	body := encodeFreeGlyphSet(glyphSet)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: FreeGlyphSetOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func AddGlyphs(conn *x.Conn, glyphSet GlyphSet, glyphIds []uint32, glyphs []GlyphInfo, data []byte) {
	body := encodeAddGlyphs(glyphSet, glyphIds, glyphs, data)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: AddGlyphsOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func AddGlyphsChecked(conn *x.Conn, glyphSet GlyphSet, glyphIds []uint32, glyphs []GlyphInfo, data []byte) x.VoidCookie {
	body := encodeAddGlyphs(glyphSet, glyphIds, glyphs, data)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: AddGlyphsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func FreeGlyphs(conn *x.Conn, glyphSet GlyphSet, glyphs []Glyph) {
	body := encodeFreeGlyphs(glyphSet, glyphs)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: FreeGlyphsOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func FreeGlyphsChecked(conn *x.Conn, glyphSet GlyphSet, glyphs []Glyph) x.VoidCookie {
	body := encodeFreeGlyphs(glyphSet, glyphs)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: FreeGlyphsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func CompositeGlyphs8(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) {
	body := encodeCompositeGlyphs8(op, src, dst, maskFormat, glyphSet, srcX, srcY, glyphCmds)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CompositeGlyphs8Opcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CompositeGlyphs8Checked(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) x.VoidCookie {
	body := encodeCompositeGlyphs8(op, src, dst, maskFormat, glyphSet, srcX, srcY, glyphCmds)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CompositeGlyphs8Opcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func CompositeGlyphs16(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) {
	body := encodeCompositeGlyphs16(op, src, dst, maskFormat, glyphSet, srcX, srcY, glyphCmds)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CompositeGlyphs16Opcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CompositeGlyphs16Checked(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) x.VoidCookie {
	body := encodeCompositeGlyphs16(op, src, dst, maskFormat, glyphSet, srcX, srcY, glyphCmds)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CompositeGlyphs16Opcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func CompositeGlyphs32(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) {
	body := encodeCompositeGlyphs32(op, src, dst, maskFormat, glyphSet, srcX, srcY, glyphCmds)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CompositeGlyphs32Opcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CompositeGlyphs32Checked(conn *x.Conn, op uint8, src, dst Picture, maskFormat PictFormat, glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) x.VoidCookie {
	body := encodeCompositeGlyphs32(op, src, dst, maskFormat, glyphSet, srcX, srcY, glyphCmds)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CompositeGlyphs32Opcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func CreateCursor(conn *x.Conn, cid x.Cursor, source Picture, X, y uint16) {
	body := encodeCreateCursor(cid, source, X, y)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateCursorOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CreateCursorChecked(conn *x.Conn, cid x.Cursor, source Picture, X, y uint16) x.VoidCookie {
	body := encodeCreateCursor(cid, source, X, y)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateCursorOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func CreateAnimCursor(conn *x.Conn, cid x.Cursor, cursors []AnimCursorElt) {
	body := encodeCreateAnimCursor(cid, cursors)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateAnimCursorOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CreateAnimCursorChecked(conn *x.Conn, cid x.Cursor, cursors []AnimCursorElt) x.VoidCookie {
	body := encodeCreateAnimCursor(cid, cursors)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateAnimCursorOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func AddTraps(conn *x.Conn, picture Picture, xOff, yOff int16, traps []Trap) {
	body := encodeAddTraps(picture, xOff, yOff, traps)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: AddTrapsOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func AddTrapsChecked(conn *x.Conn, picture Picture, xOff, yOff int16, traps []Trap) x.VoidCookie {
	body := encodeAddTraps(picture, xOff, yOff, traps)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: AddTrapsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func CreateSolidFill(conn *x.Conn, pid Picture, color Color) {
	body := encodeCreateSolidFill(pid, color)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateSolidFillOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CreateSolidFillChecked(conn *x.Conn, pid Picture, color Color) x.VoidCookie {
	body := encodeCreateSolidFill(pid, color)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateSolidFillOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func CreateLinearGradient(conn *x.Conn, picture Picture, p1, p2 PointFixed, stops []Fixed, colors []Color) {
	body := encodeCreateLinearGradient(picture, p1, p2, stops, colors)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateLinearGradientOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CreateLinearGradientChecked(conn *x.Conn, picture Picture, p1, p2 PointFixed, stops []Fixed, colors []Color) x.VoidCookie {
	body := encodeCreateLinearGradient(picture, p1, p2, stops, colors)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateLinearGradientOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func CreateRadialGradient(conn *x.Conn, picture Picture, inner, outer PointFixed, innerRadius, outerRadius Fixed, stops []Fixed, colors []Color) {
	body := encodeCreateRadialGradient(picture, inner, outer, innerRadius, outerRadius, stops, colors)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateRadialGradientOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CreateRadialGradientChecked(conn *x.Conn, picture Picture, inner, outer PointFixed, innerRadius, outerRadius Fixed, stops []Fixed, colors []Color) x.VoidCookie {
	body := encodeCreateRadialGradient(picture, inner, outer, innerRadius, outerRadius, stops, colors)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateRadialGradientOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func CreateConicalGradient(conn *x.Conn, picture Picture, center PointFixed, angle Fixed, stops []Fixed, colors []Color) {
	body := encodeCreateConicalGradient(picture, center, angle, stops, colors)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateConicalGradientOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CreateConicalGradientChecked(conn *x.Conn, picture Picture, center PointFixed, angle Fixed, stops []Fixed, colors []Color) x.VoidCookie {
	body := encodeCreateConicalGradient(picture, center, angle, stops, colors)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateConicalGradientOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}
