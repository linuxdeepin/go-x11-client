package damage

import (
	"fmt"

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

type BadDamageError struct {
	Code        uint8
	Sequence    uint16
	Bad         Damage
	MinorOpcode uint16
	MajorOpcode uint8
}

func (e BadDamageError) GetCode() uint8 {
	return e.Code
}

func (e BadDamageError) GetSequenceNumber() uint16 {
	return e.Sequence
}

func (e BadDamageError) GetMinorOpcode() uint16 {
	return e.MinorOpcode
}

func (e BadDamageError) GetMajorOpcode() uint8 {
	return e.MajorOpcode
}

func (e BadDamageError) Error() string {
	return fmt.Sprintf("damage.BadDamageError{code=%d, seq=%d, minor=%d, major=%d, bad=%d}", e.Code, e.Sequence, e.MinorOpcode, e.MajorOpcode, e.Bad)
}

func readBadDamageError(r *x.Reader) x.Error {
	var e BadDamageError
	// Error
	r.Read1b()
	e.Code = r.Read1b()
	e.Sequence = r.Read2b()

	e.Bad = Damage(r.Read4b())

	e.MinorOpcode = r.Read2b()
	e.MajorOpcode = r.Read1b()

	r.ReadPad(21)
	return e
}

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
func writeCreate(w *x.Writer, damage Damage, drawable x.Drawable, level uint8) {
	w.WritePad(4)
	w.Write4b(uint32(damage))
	w.Write4b(uint32(drawable))
	w.Write1b(level)
	w.WritePad(3)
}

// #WREQ
func writeDestroy(w *x.Writer, damage Damage) {
	w.WritePad(4)
	w.Write4b(uint32(damage))
}

// #WREQ
//func Subtract() {
//
//}
