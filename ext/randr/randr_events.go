package randr

import "github.com/linuxdeepin/go-x11-client"

type ScreenChangeNotifyEvent struct {
	Rotation        uint8
	Sequence        uint16
	Timestamp       x.Timestamp
	ConfigTimestamp x.Timestamp
	Root            x.Window
	RequestWindow   x.Window
	SizeID          uint16
	SubpixelOrder   uint16
	Width           uint16
	Height          uint16
	MmWidth         uint16
	MmHeight        uint16
}

func readScreenChangeNotifyEvent(r *x.Reader, v *ScreenChangeNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.Rotation = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Timestamp = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.ConfigTimestamp = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Root = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.RequestWindow = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.SizeID = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SubpixelOrder = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Width = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Height = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MmWidth = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MmHeight = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type NotifyEvent struct {
	SubCode  uint8
	Sequence uint16
}

func readNotifyEvent(r *x.Reader, v *NotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.SubCode = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// TODO
	return nil
}
