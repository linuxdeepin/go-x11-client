// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package xfixes

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

func ChangeSaveSet(conn *x.Conn, mode, target, map0 uint8, window x.Window) {
	body := encodeChangeSaveSet(mode, target, map0, window)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ChangeSaveSetOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func ChangeSaveSetChecked(conn *x.Conn, mode, target, map0 uint8, window x.Window) x.VoidCookie {
	body := encodeChangeSaveSet(mode, target, map0, window)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ChangeSaveSetOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func SelectSelectionInput(conn *x.Conn, window x.Window, selection x.Atom, eventMask uint32) {
	body := encodeSelectSelectionInput(window, selection, eventMask)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SelectSelectionInputOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SelectSelectionInputChecked(conn *x.Conn, window x.Window, selection x.Atom, eventMask uint32) x.VoidCookie {
	body := encodeSelectSelectionInput(window, selection, eventMask)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SelectSelectionInputOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func SelectCursorInput(conn *x.Conn, window x.Window, eventMask uint32) {
	body := encodeSelectCursorInput(window, eventMask)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SelectCursorInputOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SelectCursorInputChecked(conn *x.Conn, window x.Window, eventMask uint32) x.VoidCookie {
	body := encodeSelectCursorInput(window, eventMask)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SelectCursorInputOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func GetCursorImage(conn *x.Conn) GetCursorImageCookie {
	body := encodeGetCursorImage()
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetCursorImageOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetCursorImageCookie(seq)
}

func (cookie GetCursorImageCookie) Reply(conn *x.Conn) (*GetCursorImageReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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
	body := encodeCreateRegion(region, rects)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateRegionOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CreateRegionChecked(conn *x.Conn, region Region, rects []x.Rectangle) x.VoidCookie {
	body := encodeCreateRegion(region, rects)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateRegionOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func CreateRegionFromBitmap(conn *x.Conn, region Region, bitmap x.Pixmap) {
	body := encodeCreateRegionFromBitmap(region, bitmap)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateRegionFromBitmapOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CreateRegionFromBitmapChecked(conn *x.Conn, region Region, bitmap x.Pixmap) x.VoidCookie {
	body := encodeCreateRegionFromBitmap(region, bitmap)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateRegionFromBitmapOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func CreateRegionFromGC(conn *x.Conn, region Region, gc x.GContext) {
	body := encodeCreateRegionFromGC(region, gc)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateRegionFromGCOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CreateRegionFromGCChecked(conn *x.Conn, region Region, gc x.GContext) x.VoidCookie {
	body := encodeCreateRegionFromGC(region, gc)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateRegionFromGCOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func DestroyRegion(conn *x.Conn, region Region) {
	body := encodeDestroyRegion(region)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: DestroyRegionOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func DestroyRegionChecked(conn *x.Conn, region Region) x.VoidCookie {
	body := encodeDestroyRegion(region)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: DestroyRegionOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func SetRegion(conn *x.Conn, region Region, rects []x.Rectangle) {
	body := encodeSetRegion(region, rects)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetRegionOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetRegionChecked(conn *x.Conn, region Region, rects []x.Rectangle) x.VoidCookie {
	body := encodeSetRegion(region, rects)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetRegionOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func CopyRegion(conn *x.Conn, source, destination Region) {
	body := encodeCopyRegion(source, destination)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CopyRegionOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CopyRegionChecked(conn *x.Conn, source, destination Region) x.VoidCookie {
	body := encodeCopyRegion(source, destination)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CopyRegionOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func UnionRegion(conn *x.Conn, source1, source2, destination Region) {
	body := encodeUnionRegion(source1, source2, destination)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: UnionRegionOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func UnionRegionChecked(conn *x.Conn, source1, source2, destination Region) x.VoidCookie {
	body := encodeUnionRegion(source1, source2, destination)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: UnionRegionOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func IntersectRegion(conn *x.Conn, source1, source2, destination Region) {
	body := encodeIntersectRegion(source1, source2, destination)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: IntersectRegionOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func IntersectRegionChecked(conn *x.Conn, source1, source2, destination Region) x.VoidCookie {
	body := encodeIntersectRegion(source1, source2, destination)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: IntersectRegionOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func SubtractRegion(conn *x.Conn, source1, source2, destination Region) {
	body := encodeSubtractRegion(source1, source2, destination)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SubtractRegionOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SubtractRegionChecked(conn *x.Conn, source1, source2, destination Region) x.VoidCookie {
	body := encodeSubtractRegion(source1, source2, destination)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SubtractRegionOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func InvertRegion(conn *x.Conn, source Region, bounds x.Rectangle, destination Region) {
	body := encodeInvertRegion(source, bounds, destination)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: InvertRegionOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func InvertRegionChecked(conn *x.Conn, source Region, bounds x.Rectangle, destination Region) x.VoidCookie {
	body := encodeInvertRegion(source, bounds, destination)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: InvertRegionOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func TranslateRegion(conn *x.Conn, region Region, dx, dy int16) {
	body := encodeTranslateRegion(region, dx, dy)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: TranslateRegionOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func TranslateRegionChecked(conn *x.Conn, region Region, dx, dy int16) x.VoidCookie {
	body := encodeTranslateRegion(region, dx, dy)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: TranslateRegionOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func RegionExtents(conn *x.Conn, source, destination Region) {
	body := encodeRegionExtents(source, destination)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: RegionExtentsOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func RegionExtentsChecked(conn *x.Conn, source, destination Region) x.VoidCookie {
	body := encodeRegionExtents(source, destination)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: RegionExtentsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func FetchRegion(conn *x.Conn, region Region) FetchRegionCookie {
	body := encodeFetchRegion(region)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: FetchRegionOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return FetchRegionCookie(seq)
}

func (cookie FetchRegionCookie) Reply(conn *x.Conn) (*FetchRegionReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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
	body := encodeSetGCClipRegion(gc, clipXOrigin, clipYOrigin, region)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetGCClipRegionOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetGCClipRegionChecked(conn *x.Conn, gc x.GContext, clipXOrigin, clipYOrigin int16, region Region) x.VoidCookie {
	body := encodeSetGCClipRegion(gc, clipXOrigin, clipYOrigin, region)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetGCClipRegionOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func SetCursorName(conn *x.Conn, cursor x.Cursor, name string) {
	body := encodeSetCursorName(cursor, name)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetCursorNameOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetCursorNameChecked(conn *x.Conn, cursor x.Cursor, name string) x.VoidCookie {
	body := encodeSetCursorName(cursor, name)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetCursorNameOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func GetCursorImageAndName(conn *x.Conn) GetCursorImageAndNameCookie {
	body := encodeGetCursorImageAndName()
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetCursorImageAndNameOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetCursorImageAndNameCookie(seq)
}

func (cookie GetCursorImageAndNameCookie) Reply(conn *x.Conn) (*GetCursorImageAndNameReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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
	body := encodeChangeCursor(source, destination)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ChangeCursorOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func ChangeCursorChecked(conn *x.Conn, source, destination x.Cursor) x.VoidCookie {
	body := encodeChangeCursor(source, destination)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ChangeCursorOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func ChangeCursorByName(conn *x.Conn, src x.Cursor, name string) {
	body := encodeChangeCursorByName(src, name)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ChangeCursorByNameOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func ChangeCursorByNameChecked(conn *x.Conn, src x.Cursor, name string) x.VoidCookie {
	body := encodeChangeCursorByName(src, name)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ChangeCursorByNameOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func ExpandRegion(conn *x.Conn, source, destination Region, left, right, top, bottom uint16) {
	body := encodeExpandRegion(source, destination, left, right, top, bottom)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ExpandRegionOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func ExpandRegionChecked(conn *x.Conn, source, destination Region, left, right, top, bottom uint16) x.VoidCookie {
	body := encodeExpandRegion(source, destination, left, right, top, bottom)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ExpandRegionOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func HideCursor(conn *x.Conn, window x.Window) {
	body := encodeHideCursor(window)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: HideCursorOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func HideCursorChecked(conn *x.Conn, window x.Window) x.VoidCookie {
	body := encodeHideCursor(window)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: HideCursorOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func ShowCursor(conn *x.Conn, window x.Window) {
	body := encodeShowCursor(window)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ShowCursorOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func ShowCursorChecked(conn *x.Conn, window x.Window) x.VoidCookie {
	body := encodeShowCursor(window)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ShowCursorOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func CreatePointerBarrier(conn *x.Conn, barrier Barrier, drawable x.Drawable, x1, y1, x2, y2 int16, directions uint32, devices []uint16) {
	body := encodeCreatePointerBarrier(barrier, drawable, x1, y1, x2, y2, directions, devices)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreatePointerBarrierOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CreatePointerBarrierChecked(conn *x.Conn, barrier Barrier, drawable x.Drawable, x1, y1, x2, y2 int16, directions uint32, devices []uint16) x.VoidCookie {
	body := encodeCreatePointerBarrier(barrier, drawable, x1, y1, x2, y2, directions, devices)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreatePointerBarrierOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func DeletePointerBarrier(conn *x.Conn, barrier Barrier) {
	body := encodeDeletePointerBarrier(barrier)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: DeletePointerBarrierOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func DeletePointerBarrierChecked(conn *x.Conn, barrier Barrier) x.VoidCookie {
	body := encodeDeletePointerBarrier(barrier)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: DeletePointerBarrierOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}
