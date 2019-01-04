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
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.Level = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Drawable = x.Drawable(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Damage = Damage(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Timestamp = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	var err error
	v.Area, err = x.ReadRectangle(r)
	if err != nil {
		return err
	}

	v.Geometry, err = x.ReadRectangle(r)
	if err != nil {
		return err
	}

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
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// sequence
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
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
