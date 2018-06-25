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
	Data     []byte
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

	v.Data = r.ReadBytes(28)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

func (e *NotifyEvent) NewCrtcChangeNotifyEvent() (*CrtcChangeNotifyEvent, error) {
	var ev CrtcChangeNotifyEvent
	r := x.NewReaderFromData(e.Data)
	err := readCrtcChangeNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

func (e *NotifyEvent) NewOutputChangeNotifyEvent() (*OutputChangeNotifyEvent, error) {
	var ev OutputChangeNotifyEvent
	r := x.NewReaderFromData(e.Data)
	err := readOutputChangeNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

func (e *NotifyEvent) NewOutputPropertyNotifyEvent() (*OutputPropertyNotifyEvent,
	error) {
	var ev OutputPropertyNotifyEvent
	r := x.NewReaderFromData(e.Data)
	err := readOutputPropertyNotifyEvent(r, &ev)
	if err != nil {
		return nil, err
	}
	return &ev, nil
}

type CrtcChangeNotifyEvent struct {
	Timestamp     x.Timestamp
	Window        x.Window
	Crtc          Crtc
	Mode          Mode
	Rotation      uint16
	X, Y          int16
	Width, Height uint16
}

func readCrtcChangeNotifyEvent(r *x.Reader, v *CrtcChangeNotifyEvent) error {
	v.Timestamp = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Crtc = Crtc(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Mode = Mode(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Rotation = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(2)
	if r.Err() != nil {
		return r.Err()
	}

	v.X = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Y = int16(r.Read2b())
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

	return nil
}

type OutputChangeNotifyEvent struct {
	Timestamp       x.Timestamp
	ConfigTimestamp x.Timestamp
	Window          x.Window
	Output          Output
	Crtc            Crtc
	Mode            Mode
	Rotation        uint16
	Connection      uint8
	SubPixelOrder   uint8
}

func readOutputChangeNotifyEvent(r *x.Reader, v *OutputChangeNotifyEvent) error {
	v.Timestamp = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.ConfigTimestamp = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Output = Output(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Crtc = Crtc(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Mode = Mode(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Rotation = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Connection = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SubPixelOrder = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type OutputPropertyNotifyEvent struct {
	Window    x.Window
	Output    Output
	Atom      x.Atom
	Timestamp x.Timestamp
	Status    uint8
}

func readOutputPropertyNotifyEvent(r *x.Reader, v *OutputPropertyNotifyEvent) error {
	v.Window = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Output = Output(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Atom = x.Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Timestamp = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Status = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}
