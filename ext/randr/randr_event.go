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
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	v.Rotation, v.Sequence = r.ReadEventHeader()

	v.Timestamp = x.Timestamp(r.Read4b())

	v.ConfigTimestamp = x.Timestamp(r.Read4b())

	v.Root = x.Window(r.Read4b())

	v.RequestWindow = x.Window(r.Read4b()) // 5

	v.SizeID = r.Read2b()
	v.SubpixelOrder = r.Read2b()

	v.Width = r.Read2b()
	v.Height = r.Read2b()

	v.MmWidth = r.Read2b()
	v.MmHeight = r.Read2b() // 8

	return nil
}

type NotifyEvent struct {
	SubCode  uint8
	Sequence uint16
	Data     []byte
}

func readNotifyEvent(r *x.Reader, v *NotifyEvent) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	v.SubCode, v.Sequence = r.ReadEventHeader()

	v.Data = r.MustReadBytes(28) // 8

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

func (e *NotifyEvent) NewResourceChangeNotifyEvent() (*ResourceChangeNotifyEvent,
	error) {
	var ev ResourceChangeNotifyEvent
	r := x.NewReaderFromData(e.Data)
	err := readResourceChangeNotifyEvent(r, &ev)
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
	if !r.RemainAtLeast4b(7) {
		return x.ErrDataLenShort
	}
	v.Timestamp = x.Timestamp(r.Read4b())

	v.Window = x.Window(r.Read4b())

	v.Crtc = Crtc(r.Read4b())

	v.Mode = Mode(r.Read4b())

	v.Rotation = r.Read2b()
	r.ReadPad(2) // 5

	v.X = int16(r.Read2b())
	v.Y = int16(r.Read2b())

	v.Width = r.Read2b()
	v.Height = r.Read2b() // 7

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
	if !r.RemainAtLeast4b(7) {
		return x.ErrDataLenShort
	}
	v.Timestamp = x.Timestamp(r.Read4b())

	v.ConfigTimestamp = x.Timestamp(r.Read4b())

	v.Window = x.Window(r.Read4b())

	v.Output = Output(r.Read4b())

	v.Crtc = Crtc(r.Read4b()) // 5

	v.Mode = Mode(r.Read4b())

	v.Rotation = r.Read2b()
	v.Connection = r.Read1b()
	v.SubPixelOrder = r.Read1b() // 7

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
	if !r.RemainAtLeast4b(5) {
		return x.ErrDataLenShort
	}
	v.Window = x.Window(r.Read4b())

	v.Output = Output(r.Read4b())

	v.Atom = x.Atom(r.Read4b())

	v.Timestamp = x.Timestamp(r.Read4b())

	v.Status = r.Read1b() // 5

	return nil
}

type ResourceChangeNotifyEvent struct {
	Timestamp x.Timestamp
	Window    x.Window
}

func readResourceChangeNotifyEvent(r *x.Reader, v *ResourceChangeNotifyEvent) error {
	if !r.RemainAtLeast4b(2) {
		return x.ErrDataLenShort
	}
	v.Timestamp = x.Timestamp(r.Read4b())
	v.Window = x.Window(r.Read4b())
	return nil
}
