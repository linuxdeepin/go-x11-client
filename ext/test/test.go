package test

import (
	x "github.com/linuxdeepin/go-x11-client"
)

// #WREQ
func encodeGetVersion(majorVersion uint8, minorVersion uint16) (b x.RequestBody) {
	b.AddBlock(1).
		Write1b(majorVersion).
		WritePad(1).
		Write2b(minorVersion).
		End()
	return
}

type GetVersionReply struct {
	MajorVersion uint8
	MinorVersion uint16
}

func readGetVersionReply(r *x.Reader, v *GetVersionReply) error {
	if !r.RemainAtLeast4b(3) {
		return x.ErrDataLenShort
	}
	v.MajorVersion, _ = r.ReadReplyHeader()

	v.MinorVersion = r.Read2b() // 3

	return nil
}

// #WREQ
func encodeCompareCursor(window x.Window, cursor x.Cursor) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(window)).
		Write4b(uint32(cursor)).
		End()
	return
}

type CompareCursorReply struct {
	Same bool
}

func readCompareCursorReply(r *x.Reader, v *CompareCursorReply) error {
	if !r.RemainAtLeast4b(2) {
		return x.ErrDataLenShort
	}

	same, _ := r.ReadReplyHeader() // 2
	v.Same = x.Uint8ToBool(same)
	return nil
}

// #WREQ
func encodeFakeInput(evType uint8, detail uint8, time x.Timestamp, root x.Window,
	rootX, rootY int16, deviceId uint8) (b x.RequestBody) {

	b.AddBlock(8).
		Write1b(evType).
		Write1b(detail).
		WritePad(2).
		Write4b(uint32(time)).
		Write4b(uint32(root)).
		WritePad(8).
		Write2b(uint16(rootX)).
		Write2b(uint16(rootY)).
		WritePad(7).
		Write1b(deviceId).
		End()
	return
}

// #WREQ
func encodeGrabControl(impervious bool) (b x.RequestBody) {
	b.AddBlock(1).
		Write1b(x.BoolToUint8(impervious)).
		WritePad(3).
		End()
	return
}
