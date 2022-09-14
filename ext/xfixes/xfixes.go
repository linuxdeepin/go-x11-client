// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package xfixes

import (
	"math"

	"github.com/linuxdeepin/go-x11-client"
)

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
func encodeChangeSaveSet(mode, target, map0 uint8, window x.Window) (b x.RequestBody) {
	b.AddBlock(2).
		Write1b(mode).
		Write1b(target).
		Write1b(map0).
		WritePad(1).
		Write4b(uint32(window)).
		End()
	return
}

type SelectionNotifyEvent struct {
	Subtype            uint8
	Sequence           uint16
	Window             x.Window
	Owner              x.Window
	Selection          x.Atom
	Timestamp          x.Timestamp
	SelectionTimestamp x.Timestamp
}

func readSelectionNotifyEvent(r *x.Reader, v *SelectionNotifyEvent) error {
	if !r.RemainAtLeast4b(6) {
		return x.ErrDataLenShort
	}
	v.Subtype, v.Sequence = r.ReadEventHeader()

	v.Window = x.Window(r.Read4b())

	v.Owner = x.Window(r.Read4b())

	v.Selection = x.Atom(r.Read4b())

	v.Timestamp = x.Timestamp(r.Read4b())

	v.SelectionTimestamp = x.Timestamp(r.Read4b()) // 6

	return nil
}

// #WREQ
func encodeSelectSelectionInput(window x.Window, selection x.Atom,
	eventMask uint32) (b x.RequestBody) {
	b.AddBlock(3).
		Write4b(uint32(window)).
		Write4b(uint32(selection)).
		Write4b(eventMask).
		End()
	return
}

type CursorNotifyEvent struct {
	Subtype      uint8
	Sequence     uint16
	Window       x.Window
	CursorSerial uint32
	Timestamp    x.Timestamp
	Name         x.Atom
}

func readCursorNotifyEvent(r *x.Reader, v *CursorNotifyEvent) error {
	if !r.RemainAtLeast4b(5) {
		return x.ErrDataLenShort
	}
	v.Subtype, v.Sequence = r.ReadEventHeader()

	v.Window = x.Window(r.Read4b())

	v.CursorSerial = r.Read4b()

	v.Timestamp = x.Timestamp(r.Read4b())

	v.Name = x.Atom(r.Read4b()) // 5

	return nil
}

// #WREQ
func encodeSelectCursorInput(window x.Window, eventMask uint32) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(window)).
		Write4b(eventMask).
		End()
	return
}

// #WREQ
func encodeGetCursorImage() (b x.RequestBody) {
	return
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
	if !r.RemainAtLeast4b(6) {
		return x.ErrDataLenShort
	}
	r.ReadPad(8)

	v.X = int16(r.Read2b())
	v.Y = int16(r.Read2b())

	v.Width = r.Read2b()
	v.Height = r.Read2b()

	v.XHot = r.Read2b()
	v.YHot = r.Read2b()

	v.CursorSerial = r.Read4b() // 6

	dataLen := int(v.Width) * int(v.Height)
	if dataLen > 0 {
		if !r.RemainAtLeast4b(dataLen) {
			return x.ErrDataLenShort
		}
		v.CursorImage = make([]uint32, dataLen)
		for i := 0; i < dataLen; i++ {
			v.CursorImage[i] = r.Read4b()
		}
	}

	return nil
}

// #WREQ
func encodeCreateRegion(region Region, rects []x.Rectangle) (b x.RequestBody) {
	b0 := b.AddBlock(1 + 2*len(rects)).
		Write4b(uint32(region))
	for _, rect := range rects {
		x.WriteRectangle(b0, rect)
	}
	b0.End()
	return
}

// #WREQ
func encodeCreateRegionFromBitmap(region Region, bitmap x.Pixmap) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(region)).
		Write4b(uint32(bitmap)).
		End()
	return
}

// #WREQ
//func encodeCreateRegionFromWindow

// #WREQ
func encodeCreateRegionFromGC(region Region, gc x.GContext) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(region)).
		Write4b(uint32(gc)).
		End()
	return
}

// #WREQ
//func encodeCreateRegionFromPicture

// #WREQ
func encodeDestroyRegion(region Region) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(region)).
		End()
	return
}

// #WREQ
func encodeSetRegion(region Region, rects []x.Rectangle) (b x.RequestBody) {
	b0 := b.AddBlock(1 + 2*len(rects)).
		Write4b(uint32(region))
	for _, rect := range rects {
		x.WriteRectangle(b0, rect)
	}
	b0.End()
	return
}

// #WREQ
func encodeCopyRegion(source, destination Region) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(source)).
		Write4b(uint32(destination)).
		End()
	return
}

// #WREQ
func encodeUnionRegion(source1, source2, destination Region) (b x.RequestBody) {
	b.AddBlock(3).
		Write4b(uint32(source1)).
		Write4b(uint32(source2)).
		Write4b(uint32(destination)).
		End()
	return
}

// #WREQ
func encodeIntersectRegion(source1, source2, destination Region) (b x.RequestBody) {
	b.AddBlock(3).
		Write4b(uint32(source1)).
		Write4b(uint32(source2)).
		Write4b(uint32(destination)).
		End()
	return
}

// #WREQ
func encodeSubtractRegion(source1, source2, destination Region) (b x.RequestBody) {
	b.AddBlock(3).
		Write4b(uint32(source1)).
		Write4b(uint32(source2)).
		Write4b(uint32(destination)).
		End()
	return
}

// #WREQ
func encodeInvertRegion(source Region, bounds x.Rectangle, destination Region) (b x.RequestBody) {
	b0 := b.AddBlock(4).
		Write4b(uint32(source))
	x.WriteRectangle(b0, bounds)
	b0.Write4b(uint32(destination)).
		End()
	return
}

// #WREQ
func encodeTranslateRegion(region Region, dx, dy int16) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(region)).
		Write2b(uint16(dx)).
		Write2b(uint16(dy)).
		End()
	return
}

// #WREQ
func encodeRegionExtents(source, destination Region) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(source)).
		Write4b(uint32(destination)).
		End()
	return
}

// #WREQ
func encodeFetchRegion(region Region) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(region)).
		End()
	return
}

type FetchRegionReply struct {
	Extents    x.Rectangle
	Rectangles []x.Rectangle
}

func readFetchRegionReply(r *x.Reader, v *FetchRegionReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	_, replyLen := r.ReadReplyHeader()

	v.Extents = x.ReadRectangle(r) // 4

	r.ReadPad(16) // 8

	rectanglesLen := int(replyLen / 2)
	if rectanglesLen > 0 {
		if !r.RemainAtLeast4b(2 * rectanglesLen) {
			return x.ErrDataLenShort
		}
		v.Rectangles = make([]x.Rectangle, rectanglesLen)
		for i := 0; i < rectanglesLen; i++ {
			v.Rectangles[i] = x.ReadRectangle(r)
		}
	}

	return nil
}

// #WREQ
func encodeSetGCClipRegion(gc x.GContext, clipXOrigin, clipYOrigin int16,
	region Region) (b x.RequestBody) {

	b.AddBlock(3).
		Write4b(uint32(gc)).
		Write4b(uint32(region)).
		Write2b(uint16(clipXOrigin)).
		Write2b(uint16(clipYOrigin)).
		End()
	return
}

// #WREQ
func encodeSetCursorName(cursor x.Cursor, name string) (b x.RequestBody) {
	name = x.TruncateStr(name, math.MaxUint16)
	nameLen := len(name)
	b.AddBlock(2 + x.SizeIn4bWithPad(nameLen)).
		Write4b(uint32(cursor)).
		Write2b(uint16(nameLen)).
		WritePad(2).
		WriteString(name).
		WritePad(x.Pad(nameLen)).
		End()
	return
}

// #WREQ
func encodeGetCursorImageAndName() (b x.RequestBody) {
	return
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
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	r.ReadPad(8)

	v.X = int16(r.Read2b())
	v.Y = int16(r.Read2b())

	v.Width = r.Read2b()
	v.Height = r.Read2b()

	v.XHot = r.Read2b()
	v.YHot = r.Read2b() // 5

	v.CursorSerial = r.Read4b()

	v.CursorAtom = x.Atom(r.Read4b())

	cursorNameLen := int(r.Read2b())
	r.ReadPad(2) // 8

	dataLen := int(v.Width) * int(v.Height)
	if dataLen > 0 {
		if !r.RemainAtLeast4b(dataLen) {
			return x.ErrDataLenShort
		}
		v.CursorImage = make([]uint32, dataLen)
		for i := 0; i < dataLen; i++ {
			v.CursorImage[i] = r.Read4b()
		}
	}

	var err error
	v.CursorName, err = r.ReadString(cursorNameLen)
	return err
}

// #WREQ
func encodeChangeCursor(source, destination x.Cursor) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(source)).
		Write4b(uint32(destination)).
		End()
	return
}

// #WREQ
func encodeChangeCursorByName(src x.Cursor, name string) (b x.RequestBody) {
	name = x.TruncateStr(name, math.MaxUint16)
	nameLen := len(name)
	b.AddBlock(2 + x.SizeIn4bWithPad(nameLen)).
		Write4b(uint32(src)).
		Write2b(uint16(nameLen)).
		WritePad(2).
		WriteString(name).
		WritePad(x.Pad(nameLen)).
		End()
	return
}

// #WREQ
func encodeExpandRegion(source, destination Region, left, right,
	top, bottom uint16) (b x.RequestBody) {

	b.AddBlock(4).
		Write4b(uint32(source)).
		Write4b(uint32(destination)).
		Write2b(left).
		Write2b(right).
		Write2b(top).
		Write2b(bottom).
		End()
	return
}

// #WREQ
func encodeHideCursor(window x.Window) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

// #WREQ
func encodeShowCursor(window x.Window) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

// #WREQ
func encodeCreatePointerBarrier(barrier Barrier, drawable x.Drawable,
	x1, y1, x2, y2 int16, directions uint32, devices []uint16) (b x.RequestBody) {
	devicesLen := len(devices)
	b0 := b.AddBlock(6 + x.SizeIn4bWithPad(devicesLen*2)).
		Write4b(uint32(barrier)).
		Write4b(uint32(drawable)).
		Write2b(uint16(x1)).
		Write2b(uint16(y1)).
		Write2b(uint16(x2)).
		Write2b(uint16(y2)).
		Write4b(directions).
		WritePad(2).
		Write2b(uint16(len(devices)))
	for _, device := range devices {
		b0.Write2b(device)
	}
	b0.WritePad(x.Pad(devicesLen * 2))
	b0.End()
	return
}

// #WREQ
func encodeDeletePointerBarrier(barrier Barrier) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(barrier)).
		End()
	return
}
