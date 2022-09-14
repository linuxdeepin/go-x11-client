// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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
	if !r.RemainAtLeast4b(5) {
		return x.ErrDataLenShort
	}
	v.State, v.Sequence = r.ReadEventHeader()

	v.Time = x.Timestamp(r.Read4b())

	v.Root = x.Window(r.Read4b())

	v.Window = x.Window(r.Read4b())

	v.Kind = r.Read1b()
	v.Forced = r.ReadBool() // 5

	return nil
}

// #WREQ
func encodeQueryVersion(clientMajorVersion, clientMinorVersion uint8) (b x.RequestBody) {
	b.AddBlock(1).
		Write1b(clientMajorVersion).
		Write1b(clientMinorVersion).
		WritePad(2).
		End()
	return
}

type QueryVersionReply struct {
	ServerMajorVersion uint16
	ServerMinorVersion uint16
}

func readQueryVersionReply(r *x.Reader, v *QueryVersionReply) error {
	if !r.RemainAtLeast4b(3) {
		return x.ErrDataLenShort
	}
	r.ReadPad(8)

	v.ServerMajorVersion = r.Read2b()
	v.ServerMinorVersion = r.Read2b() // 3

	return nil
}

// #WREQ
func encodeQueryInfo(drawable x.Drawable) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(drawable)).
		End()
	return
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
	if !r.RemainAtLeast4b(7) {
		return x.ErrDataLenShort
	}
	v.State, _ = r.ReadReplyHeader()

	v.SaverWindow = x.Window(r.Read4b())

	v.MsUntilServer = r.Read4b()

	v.MsSinceUserInput = r.Read4b()

	v.EventMask = r.Read4b() // 6

	v.Kind = r.Read1b() // 7

	return nil
}

// #WREQ
func encodeSelectInput(drawable x.Drawable, eventMask uint32) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(drawable)).
		Write4b(eventMask).
		End()
	return
}

// #WREQ
func encodeSetAttributes(drawable x.Drawable, X, y int16, width, height,
	boardWidth uint16, class, depth uint8, visual x.VisualID, valueMask uint32,
	valueList []uint32) (b x.RequestBody) {

	b0 := b.AddBlock(6 + len(valueList)).
		Write4b(uint32(drawable)).
		Write2b(uint16(X)).
		Write2b(uint16(y)).
		Write2b(width).
		Write2b(height).
		Write2b(boardWidth).
		Write1b(class).
		Write1b(depth).
		Write4b(uint32(visual)).
		Write4b(valueMask)

	for _, value := range valueList {
		b0.Write4b(value)
	}
	b0.End()
	return
}

// #WREQ
func encodeUnsetAttributes(drawable x.Drawable) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(drawable)).
		End()
	return
}

// #WREQ
func encodeSuspend(suspend bool) (b x.RequestBody) {
	b.AddBlock(1).
		WriteBool(suspend).
		WritePad(3).
		End()
	return
}
