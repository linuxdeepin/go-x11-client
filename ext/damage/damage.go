package damage

import (
	x "github.com/linuxdeepin/go-x11-client"
)

type NotifyEvent struct {
	Level     uint8
	Sequence  uint16
	Drawable  x.Drawable
	Damage    Damage
	Timestamp x.Timestamp
	Area      x.Rectangle
	Geometry  x.Rectangle
}

func readNotifyEvent(r *x.Reader, v *NotifyEvent) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	v.Level, v.Sequence = r.ReadEventHeader()

	v.Drawable = x.Drawable(r.Read4b())

	v.Damage = Damage(r.Read4b())

	v.Timestamp = x.Timestamp(r.Read4b()) // 4

	v.Area = x.ReadRectangle(r) // 6

	v.Geometry = x.ReadRectangle(r) // 8

	return nil
}

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
func encodeCreate(damage Damage, drawable x.Drawable, level uint8) (b x.RequestBody) {
	b.AddBlock(3).
		Write4b(uint32(damage)).
		Write4b(uint32(drawable)).
		Write1b(level).
		WritePad(3).
		End()
	return
}

// #WREQ
func encodeDestroy(damage Damage) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(damage)).
		End()
	return
}

// #WREQ
//func Subtract() {
//
//}
