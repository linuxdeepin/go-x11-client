package x

type KeyEvent struct {
	Detail     Keycode
	Sequence   uint16
	Time       Timestamp
	Root       Window
	Event      Window
	Child      Window
	RootX      int16
	RootY      int16
	EventX     int16
	EventY     int16
	State      uint16
	SameScreen bool
}

func readKeyEvent(r *Reader, v *KeyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.Detail = Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Time = Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Root = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Event = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Child = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.RootX = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.RootY = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.EventX = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.EventY = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.State = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SameScreen = Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type KeyPressEvent struct {
	KeyEvent
}

func readKeyPressEvent(r *Reader, v *KeyPressEvent) error {
	return readKeyEvent(r, &v.KeyEvent)
}

type KeyReleaseEvent struct {
	KeyEvent
}

func readKeyReleaseEvent(r *Reader, v *KeyReleaseEvent) error {
	return readKeyEvent(r, &v.KeyEvent)
}

type ButtonEvent struct {
	Detail     Button
	Sequence   uint16
	Time       Timestamp
	Root       Window
	Event      Window
	Child      Window
	RootX      int16
	RootY      int16
	EventX     int16
	EventY     int16
	State      uint16
	SameScreen bool
}

func readButtonEvent(r *Reader, v *ButtonEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.Detail = Button(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Time = Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Root = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Event = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Child = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.RootX = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.RootY = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.EventX = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.EventY = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.State = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SameScreen = Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type ButtonPressEvent struct {
	ButtonEvent
}

func readButtonPressEvent(r *Reader, v *ButtonPressEvent) error {
	return readButtonEvent(r, &v.ButtonEvent)
}

type ButtonReleaseEvent struct {
	ButtonEvent
}

func readButtonReleaseEvent(r *Reader, v *ButtonReleaseEvent) error {
	return readButtonEvent(r, &v.ButtonEvent)
}

type MotionNotifyEvent struct {
	Detail     uint8
	Sequence   uint16
	Time       Timestamp
	Root       Window
	Event      Window
	Child      Window
	RootX      int16
	RootY      int16
	EventX     int16
	EventY     int16
	State      uint16
	SameScreen bool
}

func readMotionNotifyEvent(r *Reader, v *MotionNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.Detail = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Time = Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Root = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Event = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Child = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.RootX = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.RootY = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.EventX = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.EventY = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.State = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SameScreen = Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type PointerWindowEvent struct {
	Detail          uint8
	Sequence        uint16
	Time            Timestamp
	Root            Window
	Event           Window
	Child           Window
	RootX           int16
	RootY           int16
	EventX          int16
	EventY          int16
	State           uint16
	Mode            uint8
	SameScreenFocus uint8
}

func readPointerWindowEvent(r *Reader, v *PointerWindowEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.Detail = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Time = Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Root = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Event = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Child = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.RootX = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.RootY = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.EventX = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.EventY = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.State = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Mode = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SameScreenFocus = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type EnterNotifyEvent struct {
	PointerWindowEvent
}

func readEnterNotifyEvent(r *Reader, v *EnterNotifyEvent) error {
	return readPointerWindowEvent(r, &v.PointerWindowEvent)
}

type LeaveNotifyEvent struct {
	PointerWindowEvent
}

func readLeaveNotifyEvent(r *Reader, v *LeaveNotifyEvent) error {
	return readPointerWindowEvent(r, &v.PointerWindowEvent)
}

type FocusEvent struct {
	Detail   uint8
	Sequence uint16
	Event    Window
	Mode     uint8
}

func readFocusEvent(r *Reader, v *FocusEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.Detail = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Mode = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(23)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type FocusInEvent struct {
	FocusEvent
}

func readFocusInEvent(r *Reader, v *FocusInEvent) error {
	return readFocusEvent(r, &v.FocusEvent)
}

type FocusOutEvent struct {
	FocusEvent
}

func readFocusOutEvent(r *Reader, v *FocusOutEvent) error {
	return readFocusEvent(r, &v.FocusEvent)
}

type KeymapNotifyEvent struct {
	Keys []byte
}

func readKeymapNotifyEvent(r *Reader, v *KeymapNotifyEvent) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Keys = r.ReadBytes(31)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type ExposeEvent struct {
	Sequence uint16
	Window   Window
	X        uint16
	Y        uint16
	Width    uint16
	Height   uint16
	Count    uint16
}

func readExposeEvent(r *Reader, v *ExposeEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.X = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Y = r.Read2b()
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

	v.Count = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(14)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type GraphicsExposureEvent struct {
	Sequence    uint16
	Drawable    Drawable
	X           uint16
	Y           uint16
	Width       uint16
	Height      uint16
	MinorOpcode uint16
	Count       uint16
	MajorOpcode uint8
}

func readGraphicsExposureEvent(r *Reader, v *GraphicsExposureEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Drawable = Drawable(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.X = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Y = r.Read2b()
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

	v.MinorOpcode = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Count = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MajorOpcode = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(11)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type NoExposureEvent struct {
	Sequence    uint16
	Drawable    Drawable
	MinorOpcode uint16
	MajorOpcode uint8
}

func readNoExposureEvent(r *Reader, v *NoExposureEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Drawable = Drawable(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.MinorOpcode = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MajorOpcode = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(21)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type VisibilityNotifyEvent struct {
	Sequence uint16
	Window   Window
	State    uint8
}

func readVisibilityNotifyEvent(r *Reader, v *VisibilityNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.State = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(23)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type CreateNotifyEvent struct {
	Sequence         uint16
	Parent           Window
	Window           Window
	X                int16
	Y                int16
	Width            uint16
	Height           uint16
	BorderWidth      uint16
	OverrideRedirect bool
}

func readCreateNotifyEvent(r *Reader, v *CreateNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = Window(r.Read4b())
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

	v.BorderWidth = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.OverrideRedirect = Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(9)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type DestroyNotifyEvent struct {
	Sequence uint16
	Event    Window
	Window   Window
}

func readDestroyNotifyEvent(r *Reader, v *DestroyNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Event = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(20)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type UnmapNotifyEvent struct {
	Sequence      uint16
	Event         Window
	Window        Window
	FromConfigure bool
}

func readUnmapNotifyEvent(r *Reader, v *UnmapNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Event = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.FromConfigure = Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(19)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type MapNotifyEvent struct {
	Sequence         uint16
	Event            Window
	Window           Window
	OverrideRedirect bool
}

func readMapNotifyEvent(r *Reader, v *MapNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Event = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.OverrideRedirect = Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(19)
	if r.Err() != nil {
		return r.Err()
	}
	return nil
}

type MapRequestEvent struct {
	Sequence uint16
	Parent   Window
	Window   Window
}

func readMapRequestEvent(r *Reader, v *MapRequestEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Parent = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(20)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type ReparentNotifyEvent struct {
	Sequence         uint16
	Event            Window
	Window           Window
	Parent           Window
	X                int16
	Y                int16
	OverrideRedirect bool
}

func readReparentNotifyEvent(r *Reader, v *ReparentNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Event = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Parent = Window(r.Read4b())
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

	v.OverrideRedirect = Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(11)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type ConfigureNotifyEvent struct {
	Sequence         uint16
	Event            Window
	Window           Window
	AboveSibling     Window
	X                int16
	Y                int16
	Width            uint16
	Height           uint16
	BorderWidth      uint16
	OverrideRedirect bool
}

func readConfigureNotifyEvent(r *Reader, v *ConfigureNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Event = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.AboveSibling = Window(r.Read4b())
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

	v.BorderWidth = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.OverrideRedirect = Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(5)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type ConfigureRequestEvent struct {
	StackMode   uint8
	Sequence    uint16
	Parent      Window
	Window      Window
	Sibling     Window
	X           int16
	Y           int16
	Width       uint16
	Height      uint16
	BorderWidth uint16
	ValueMask   uint16
}

func readConfigureRequestEvent(r *Reader, v *ConfigureRequestEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.StackMode = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Parent = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Sibling = Window(r.Read4b())
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

	v.BorderWidth = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ValueMask = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(4)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type GravityNotifyEvent struct {
	Sequence uint16
	Event    Window
	Window   Window
	X        int16
	Y        int16
}

func readGravityNotifyEvent(r *Reader, v *GravityNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Event = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = Window(r.Read4b())
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

	r.ReadPad(16)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type ResizeRequestEvent struct {
	Sequence uint16
	Window   Window
	Width    uint16
	Height   uint16
}

func readResizeRequestEvent(r *Reader, v *ResizeRequestEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = Window(r.Read4b())
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

	r.ReadPad(20)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type CirculateEvent struct {
	Sequence uint16
	Event    Window
	Window   Window
	Place    uint8
}

func readCirculateEvent(r *Reader, v *CirculateEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Event = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Place = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(15)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type CirculateNotifyEvent struct {
	CirculateEvent
}

func readCirculateNotifyEvent(r *Reader, v *CirculateNotifyEvent) error {
	return readCirculateEvent(r, &v.CirculateEvent)
}

type CirculateRequestEvent struct {
	CirculateEvent
}

func readCirculateRequestEvent(r *Reader, v *CirculateRequestEvent) error {
	return readCirculateEvent(r, &v.CirculateEvent)
}

type PropertyNotifyEvent struct {
	Sequence uint16
	Window   Window
	Atom     Atom
	Time     Timestamp
	State    uint8
}

func readPropertyNotifyEvent(r *Reader, v *PropertyNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Atom = Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Time = Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.State = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(15)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type SelectionClearEvent struct {
	Sequence  uint16
	Time      Timestamp
	Owner     Window
	Selection Atom
}

func readSelectionClearEvent(r *Reader, v *SelectionClearEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Time = Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Owner = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Selection = Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(16)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type SelectionRequestEvent struct {
	Sequence  uint16
	Time      Timestamp
	Owner     Window
	Requestor Window
	Selection Atom
	Target    Atom
	Property  Atom
}

func readSelectionRequestEvent(r *Reader, v *SelectionRequestEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Time = Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Owner = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Requestor = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Selection = Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Target = Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Property = Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(4)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type SelectionNotifyEvent struct {
	Sequence  uint16
	Time      Timestamp
	Requestor Window
	Selection Atom
	Target    Atom
	Property  Atom
}

func readSelectionNotifyEvent(r *Reader, v *SelectionNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Time = Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Requestor = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Selection = Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Target = Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Property = Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(8)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type ColormapNotifyEvent struct {
	Sequence uint16
	Window   Window
	Colormap Colormap
	New      bool
	State    uint8
}

func readColormapNotifyEvent(r *Reader, v *ColormapNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Colormap = Colormap(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.New = Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.State = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(18)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type ClientMessageEvent struct {
	Format   uint8
	Sequence uint16
	Window   Window
	Type     Atom
	Data     ClientMessageData
}

func readClientMessageEvent(r *Reader, v *ClientMessageEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.Format = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Type = Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	data := r.ReadBytes(20)
	if r.Err() != nil {
		return r.Err()
	}
	v.Data = ClientMessageData{
		data: data,
	}

	return nil
}

func WriteClientMessageEvent(w *Writer, v *ClientMessageEvent) {
	w.Write1b(ClientMessageEventCode)
	w.Write1b(v.Format)
	w.Write2b(v.Sequence)
	w.Write4b(uint32(v.Window))
	w.Write4b(uint32(v.Type))
	writeClientMessageData(w, &v.Data)
}

type MappingNotifyEvent struct {
	Sequence     uint16
	Request      uint8
	FirstKeycode Keycode
	Count        uint8
}

func readMappingNotifyEvent(r *Reader, v *MappingNotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Request = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstKeycode = Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Count = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(25)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type GeGenericEvent struct {
	Extension uint8
	Sequence  uint16
	Length    uint32
	EventType uint16
}

func readGeGenericEvent(r *Reader, v *GeGenericEvent) error {
	// TODO
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	v.Extension = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Length = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.EventType = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}
