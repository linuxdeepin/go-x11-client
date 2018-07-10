package x

import (
	"bytes"
	"errors"
	"math"
)

type Setup struct {
	ProtocolMajorVersion     uint16
	ProtocolMinorVersion     uint16
	ReleaseNumber            uint32
	ResourceIdBase           uint32
	ResourceIdMask           uint32
	MonitorBufferSize        uint32
	MaximumRequestLength     uint16
	ImageByteOrder           uint8
	BitmapFormatBitOrder     uint8
	BitmapFormatScanlineUint uint8
	BitmapFormatScanlinePad  uint8
	MinKeycode               Keycode
	MaxKeycode               Keycode
	Vendor                   string
	PixmapFormats            []Format
	Roots                    []Screen
}

func readSetup(r *Reader, v *Setup) error {
	// status
	status := r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}
	if status != 1 {
		return errors.New("status != 1")
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ProtocolMajorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ProtocolMinorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length in 4-byte units of additional data
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ReleaseNumber = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ResourceIdBase = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ResourceIdMask = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MonitorBufferSize = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	vendorLen := r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MaximumRequestLength = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	screensLen := r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	formatsLen := r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ImageByteOrder = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.BitmapFormatBitOrder = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.BitmapFormatScanlineUint = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.BitmapFormatScanlinePad = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MinKeycode = Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.MaxKeycode = Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Vendor = r.ReadString(int(vendorLen))
	if r.Err() != nil {
		return r.Err()
	}

	// formats
	if formatsLen > 0 {
		v.PixmapFormats = make([]Format, int(formatsLen))
		for i := 0; i < int(formatsLen); i++ {
			err := readFormat(r, &v.PixmapFormats[i])
			if err != nil {
				return err
			}
		}
	}

	// screens
	if screensLen > 0 {
		v.Roots = make([]Screen, int(screensLen))
		for i := 0; i < int(screensLen); i++ {
			err := readScreen(r, &v.Roots[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type Format struct {
	Depth        uint8
	BitsPerPixel uint8
	ScanlinePad  uint8
}

func readFormat(r *Reader, v *Format) error {
	v.Depth = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.BitsPerPixel = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ScanlinePad = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(5)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type Screen struct {
	Root                Window
	DefaultColorMap     Colormap
	WhitePixel          uint32
	BlackPixel          uint32
	CurrentInputMask    uint32
	WidthInPixels       uint16
	HeightInPixels      uint16
	WidthInMillimeters  uint16
	HeightInMillimeters uint16
	MinInstalledMaps    uint16
	MaxInstalledMaps    uint16
	RootVisual          VisualID
	BackingStores       uint8
	SaveUnders          bool
	RootDepth           uint8
	AllowedDepths       []Depth
}

func readScreen(r *Reader, v *Screen) error {
	v.Root = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.DefaultColorMap = Colormap(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.WhitePixel = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.BlackPixel = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.CurrentInputMask = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.WidthInPixels = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.HeightInPixels = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.WidthInMillimeters = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.HeightInMillimeters = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MinInstalledMaps = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MaxInstalledMaps = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.RootVisual = VisualID(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.BackingStores = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	saveUnders := r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	if saveUnders == 0 {
		v.SaveUnders = false
	} else {
		v.SaveUnders = true
	}

	v.RootDepth = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	depthsLen := r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// depths
	if depthsLen > 0 {
		v.AllowedDepths = make([]Depth, int(depthsLen))
		for i := 0; i < int(depthsLen); i++ {
			err := readDepth(r, &v.AllowedDepths[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type Depth struct {
	Depth   uint8
	Visuals []VisualType
}

func readDepth(r *Reader, v *Depth) error {
	v.Depth = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	visualsLen := r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	// visuals
	if visualsLen > 0 {
		v.Visuals = make([]VisualType, int(visualsLen))
		for i := 0; i < int(visualsLen); i++ {
			err := readVisualType(r, &v.Visuals[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type VisualType struct {
	Id              VisualID
	Class           uint8
	BitsPerRGBValue uint8
	ColorMapEntries uint16
	RedMask         uint32
	GreenMask       uint32
	BlueMask        uint32
}

func readVisualType(r *Reader, v *VisualType) error {
	v.Id = VisualID(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Class = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.BitsPerRGBValue = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ColorMapEntries = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.RedMask = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.GreenMask = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.BlueMask = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type Rectangle struct {
	X, Y          int16
	Width, Height uint16
}

func ReadRectangle(r *Reader) (Rectangle, error) {
	var v Rectangle
	v.X = int16(r.Read2b())
	if r.Err() != nil {
		return Rectangle{}, r.Err()
	}

	v.Y = int16(r.Read2b())
	if r.Err() != nil {
		return Rectangle{}, r.Err()
	}

	v.Width = r.Read2b()
	if r.Err() != nil {
		return Rectangle{}, r.Err()
	}

	v.Height = r.Read2b()
	if r.Err() != nil {
		return Rectangle{}, r.Err()
	}

	return v, nil
}

func WriteRectangle(w *Writer, v Rectangle) {
	w.Write2b(uint16(v.X))
	w.Write2b(uint16(v.Y))
	w.Write2b(v.Width)
	w.Write2b(v.Height)
}

// #WREQ
func writeCreateWindow(w *Writer, depth uint8, wid, parent Window, x, y int16,
	width, height, borderWidth uint16, class uint16, visual VisualID,
	valueMask uint32, valueList []uint32) {
	// opcode
	w.WritePad(1)

	w.Write1b(depth)

	// length
	w.WritePad(2)

	w.Write4b(uint32(wid))

	w.Write4b(uint32(parent))

	w.Write2b(uint16(x))
	w.Write2b(uint16(y))

	w.Write2b(width)
	w.Write2b(height)
	w.Write2b(borderWidth)
	w.Write2b(class)
	w.Write4b(uint32(visual))
	w.Write4b(valueMask)

	for _, value := range valueList {
		w.Write4b(value)
	}
}

// #WREQ
func writeChangeWindowAttributes(w *Writer, window Window, valueMask uint32,
	valueList []uint32) {

	w.WritePad(4)
	w.Write4b(uint32(window))
	w.Write4b(valueMask)
	for _, value := range valueList {
		w.Write4b(value)
	}
}

// #WREQ
func writeGetWindowAttributes(w *Writer, window Window) {
	w.WritePad(4)
	w.Write4b(uint32(window))
}

type GetWindowAttributesReply struct {
	BackingStore       uint8
	Visual             VisualID
	Class              uint16
	BitGravity         uint8
	WinGravity         uint8
	BackingPlanes      uint32
	BackingPixel       uint32
	SaveUnder          bool
	MapIsInstalled     bool
	MapState           uint8
	OverrideRedirect   bool
	Colormap           Colormap
	AllEventMasks      uint32
	YourEventMask      uint32
	DoNotPropagateMask uint16
}

func readGetWindowAttributesReply(r *Reader, v *GetWindowAttributesReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.BackingStore = r.Read1b()
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

	v.Visual = VisualID(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Class = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.BitGravity = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.WinGravity = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.BackingPlanes = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.BackingPixel = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	saveUnder := r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	if saveUnder == 0 {
		v.SaveUnder = false
	} else {
		v.SaveUnder = true
	}

	mapIsInstalled := r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	if mapIsInstalled == 0 {
		v.MapIsInstalled = false
	} else {
		v.MapIsInstalled = true
	}

	v.MapState = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	overrideRedirect := r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	if overrideRedirect == 0 {
		v.OverrideRedirect = false
	} else {
		v.OverrideRedirect = true
	}

	v.Colormap = Colormap(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.AllEventMasks = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.YourEventMask = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.DoNotPropagateMask = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}
	return nil
}

// #WREQ
func writeDestroyWindow(w *Writer, window Window) {
	w.WritePad(4)
	w.Write4b(uint32(window))
}

// #WREQ
func writeDestroySubwindows(w *Writer, window Window) {
	w.WritePad(4)
	w.Write4b(uint32(window))
}

// #WREQ
func writeChangeSaveSet(w *Writer, mode uint8, window Window) {
	w.WritePad(1)
	w.Write1b(mode)
	w.WritePad(2)
	w.Write4b(uint32(window))
}

// #WREQ
func writeReparentWindow(w *Writer, window, parent Window, x, y int16) {
	w.WritePad(4)
	w.Write4b(uint32(window))
	w.Write4b(uint32(parent))
	w.Write2b(uint16(x))
	w.Write2b(uint16(y))
}

// #WREQ
func writeMapWindow(w *Writer, window Window) {
	w.WritePad(4)
	w.Write4b(uint32(window))
}

// #WREQ
func writeMapSubwindows(w *Writer, window Window) {
	w.WritePad(4)
	w.Write4b(uint32(window))
}

// #WREQ
func writeUnmapWindow(w *Writer, window Window) {
	w.WritePad(4)
	w.Write4b(uint32(window))
}

// #WREQ
func writeUnmapSubwindows(w *Writer, window Window) {
	w.WritePad(4)
	w.Write4b(uint32(window))
}

// #WREQ
func writeConfigureWindow(w *Writer, window Window, valueMask uint16, valueList []uint32) {
	w.WritePad(4)
	w.Write4b(uint32(window))
	w.Write2b(valueMask)
	w.WritePad(2)
	for _, value := range valueList {
		w.Write4b(value)
	}
}

// #WREQ
func writeCirculateWindow(w *Writer, direction uint8, window Window) {
	w.WritePad(1)
	w.Write1b(direction)
	w.WritePad(2)
	w.Write4b(uint32(window))
}

// #WREQ
func writeGetGeometry(w *Writer, drawable Drawable) {
	w.WritePad(4)
	w.Write4b(uint32(drawable))
}

type GetGeometryReply struct {
	Depth       uint8
	Root        Window
	X           int16
	Y           int16
	Width       uint16
	Height      uint16
	BorderWidth uint16
}

func readGetGeometryReply(r *Reader, v *GetGeometryReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Depth = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Root = Window(r.Read4b())
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

	// unused
	r.ReadPad(10)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeQueryTree(w *Writer, window Window) {
	w.WritePad(4)
	w.Write4b(uint32(window))
}

type QueryTreeReply struct {
	Root, Parent Window
	Children     []Window
}

func readQueryTreeReply(r *Reader, v *QueryTreeReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Root = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Parent = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	childrenLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(14)
	if r.Err() != nil {
		return r.Err()
	}

	if childrenLen > 0 {
		v.Children = make([]Window, childrenLen)
		for i := 0; i < childrenLen; i++ {
			v.Children[i] = Window(r.Read4b())
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	return nil
}

// #WREQ
func writeInternAtom(w *Writer, onlyIfExists bool, name string) {
	w.WritePad(1)
	w.Write1b(BoolToUint8(onlyIfExists))
	nameLen := len(name)
	if nameLen > math.MaxInt16 {
		nameLen = math.MaxInt16
		name = name[:math.MaxInt16]
	}
	w.WritePad(2)

	w.Write2b(uint16(nameLen))
	w.WritePad(2)

	// name
	w.WriteString(name)
	w.WritePad(Pad(nameLen))
}

type InternAtomReply struct {
	Atom Atom
}

func readInternAtomReply(r *Reader, v *InternAtomReply) error {
	// reply
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Atom = Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(20)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeGetAtomName(w *Writer, atom Atom) {
	w.WritePad(4)
	w.Write4b(uint32(atom))
}

type GetAtomNameReply struct {
	Name string
}

func readGetAtomNameReply(r *Reader, v *GetAtomNameReply) error {
	// reply
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	// name len
	nameLen := r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(22)
	if r.Err() != nil {
		return r.Err()
	}

	v.Name = r.ReadString(int(nameLen))
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(Pad(int(nameLen)))
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeChangeProperty(w *Writer, mode uint8, window Window, property, type0 Atom,
	format uint8, data []byte) {
	w.WritePad(1)
	w.Write1b(mode)
	w.WritePad(2)

	w.Write4b(uint32(window))
	w.Write4b(uint32(property))
	w.Write4b(uint32(type0))
	w.Write1b(format)
	w.WritePad(3)

	dataLen := len(data)
	w.Write4b(uint32(dataLen / (int(format) / 8)))
	w.WriteBytes(data)
	w.WritePad(Pad(dataLen))
}

// #WREQ
func writeDeleteProperty(w *Writer, window Window, property Atom) {
	w.WritePad(4)
	w.Write4b(uint32(window))
	w.Write4b(uint32(property))
}

// #WREQ
func writeGetProperty(w *Writer, delete bool, window Window, property, type0 Atom,
	longOffset, longLength uint32) {
	w.WritePad(1)
	w.Write1b(BoolToUint8(delete))
	w.WritePad(2)
	w.Write4b(uint32(window))
	w.Write4b(uint32(property))
	w.Write4b(uint32(type0))
	w.Write4b(longOffset)
	w.Write4b(longLength)
}

type GetPropertyReply struct {
	Format     uint8
	Type       Atom
	BytesAfter uint32
	ValueLen   uint32
	Value      []byte
}

func readGetPropertyReply(r *Reader, v *GetPropertyReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Format = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Type = Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.BytesAfter = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ValueLen = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(12)
	if r.Err() != nil {
		return r.Err()
	}

	n := int(v.ValueLen) * int(v.Format/8)
	v.Value = r.ReadBytes(n)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(Pad(n))
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeListProperties(w *Writer, window Window) {
	w.WritePad(4)
	w.Write4b(uint32(window))
}

type ListPropertiesReply struct {
	Atoms []Atom
}

func readListPropertiesReply(r *Reader, v *ListPropertiesReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	atomsLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(22)
	if r.Err() != nil {
		return r.Err()
	}

	if atomsLen > 0 {
		v.Atoms = make([]Atom, atomsLen)
		for i := 0; i < atomsLen; i++ {
			v.Atoms[i] = Atom(r.Read4b())
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	return nil
}

// #WREQ
func writeSetSelectionOwner(w *Writer, owner Window, selection Atom, time Timestamp) {
	w.WritePad(4)
	w.Write4b(uint32(owner))
	w.Write4b(uint32(selection))
	w.Write4b(uint32(time))
}

// #WREQ
func writeGetSelectionOwner(w *Writer, selection Atom) {
	w.WritePad(4)
	w.Write4b(uint32(selection))
}

type GetSelectionOwnerReply struct {
	Owner Window
}

func readGetSelectionOwnerReply(r *Reader, v *GetSelectionOwnerReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Owner = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(20)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeConvertSelection(w *Writer, requestor Window, selection, target, property Atom,
	time Timestamp) {

	w.WritePad(4)
	w.Write4b(uint32(requestor))
	w.Write4b(uint32(selection))
	w.Write4b(uint32(target))
	w.Write4b(uint32(property))
	w.Write4b(uint32(time))
}

// #WREQ
func writeSendEvent(w *Writer, propagate bool, destination Window, eventMask uint32,
	event []byte) {

	w.WritePad(1)
	w.Write1b(BoolToUint8(propagate))
	w.WritePad(2)
	w.Write4b(uint32(destination))
	w.Write4b(eventMask)
	w.WriteBytes(event)
}

// #WREQ
func writeGrabPointer(w *Writer, ownerEvents bool, grabWindow Window, eventMask uint16,
	pointerMode, keyboardMode uint8, confineTo Window, cursor Cursor, time Timestamp) {
	w.WritePad(1)
	w.Write1b(BoolToUint8(ownerEvents))
	w.WritePad(2)
	w.Write4b(uint32(grabWindow))
	w.Write2b(eventMask)
	w.Write1b(pointerMode)
	w.Write1b(keyboardMode)
	w.Write4b(uint32(confineTo))
	w.Write4b(uint32(cursor))
	w.Write4b(uint32(time))
}

type GrabPointerReply struct {
	Status uint8
}

func readGrabPointerReply(r *Reader, v *GrabPointerReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Status = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(24)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeUngrabPointer(w *Writer, time Timestamp) {
	w.WritePad(4)
	w.Write4b(uint32(time))
}

// #WREQ
func writeGrabButton(w *Writer, ownerEvents bool, grabWindow Window, eventMask uint16,
	pointerMode, keyboardMode uint8, confineTo Window, cursor Cursor, button uint8,
	modifiers uint16) {

	w.WritePad(1)
	w.Write1b(BoolToUint8(ownerEvents))
	w.WritePad(2)

	w.Write4b(uint32(grabWindow))

	w.Write2b(eventMask)
	w.Write1b(pointerMode)
	w.Write1b(keyboardMode)

	w.Write4b(uint32(confineTo))
	w.Write4b(uint32(cursor))

	w.Write1b(button)
	w.WritePad(1)
	w.Write2b(modifiers)
}

// #WREQ
func writeUngrabButton(w *Writer, button uint8, grabWindow Window, modifiers uint16) {
	w.WritePad(1)
	w.Write1b(button)
	w.WritePad(2)

	w.Write4b(uint32(grabWindow))

	w.Write2b(modifiers)
	w.WritePad(2)
}

// #WREQ
func writeChangeActivePointerGrab(w *Writer, cursor Cursor, time Timestamp, eventMask uint16) {
	w.WritePad(4)
	w.Write4b(uint32(cursor))
	w.Write4b(uint32(time))

	w.Write2b(eventMask)
	w.WritePad(2)
}

// #WREQ
func writeGrabKeyboard(w *Writer, ownerEvents bool, grabWindow Window, time Timestamp,
	pointerMode, keyboardMode uint8) {
	w.WritePad(1)
	w.Write1b(BoolToUint8(ownerEvents))
	w.WritePad(2)
	w.Write4b(uint32(grabWindow))
	w.Write4b(uint32(time))
	w.Write1b(pointerMode)
	w.Write1b(keyboardMode)
	w.WritePad(2)
}

type GrabKeyboardReply struct {
	Status uint8
}

func readGrabKeyboardReply(r *Reader, v *GrabKeyboardReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Status = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(24)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeUngrabKeyboard(w *Writer, time Timestamp) {
	w.WritePad(4)
	w.Write4b(uint32(time))
}

// #WREQ
func writeGrabKey(w *Writer, ownerEvents bool, grabWindow Window, modifiers uint16,
	key Keycode, pointerMode, keyboardMode uint8) {
	w.WritePad(1)
	w.Write1b(BoolToUint8(ownerEvents))
	w.WritePad(2)
	w.Write4b(uint32(grabWindow))
	w.Write2b(modifiers)
	w.Write1b(uint8(key))
	w.Write1b(pointerMode)
	w.Write1b(keyboardMode)
	w.WritePad(3)
}

// #WREQ
func writeUngrabKey(w *Writer, key Keycode, grabWindow Window, modifiers uint16) {
	w.WritePad(1)
	w.Write1b(uint8(key))
	w.WritePad(2)
	w.Write4b(uint32(grabWindow))
	w.Write2b(modifiers)
	w.WritePad(2)
}

// #WREQ
func writeAllowEvents(w *Writer, mode uint8, time Timestamp) {
	w.WritePad(1)
	w.Write1b(mode)
	w.WritePad(2)

	w.Write4b(uint32(time))
}

// #WREQ
func writeGrabServer(w *Writer) {
	w.WritePad(4)
}

// #WREQ
func writeUngrabServer(w *Writer) {
	w.WritePad(4)
}

// #WREQ
func writeQueryPointer(w *Writer, window Window) {
	w.WritePad(4)
	w.Write4b(uint32(window))
}

type QueryPointerReply struct {
	SameScreen bool
	Root       Window
	Child      Window
	RootX      int16
	RootY      int16
	WinX       int16
	WinY       int16
	Mask       uint16
}

func readQueryPointerReply(r *Reader, v *QueryPointerReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SameScreen = Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Root = Window(r.Read4b())
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

	v.WinX = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.WinY = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Mask = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(6)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeGetMotionEvents(w *Writer, window Window, start, stop Timestamp) {
	w.WritePad(4)
	w.Write4b(uint32(window))
	w.Write4b(uint32(start))
	w.Write4b(uint32(stop))
}

type GetMotionEventsReply struct {
	Events []TimeCoord
}

type TimeCoord struct {
	Time Timestamp
	X, Y int16
}

func readTimeCoord(r *Reader, v *TimeCoord) error {
	v.Time = Timestamp(r.Read4b())
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

	return nil
}

func readGetMotionEventsReply(r *Reader, v *GetMotionEventsReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	eventsLen := int(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(20)
	if r.Err() != nil {
		return r.Err()
	}

	if eventsLen > 0 {
		v.Events = make([]TimeCoord, eventsLen)
		for i := 0; i < eventsLen; i++ {
			err := readTimeCoord(r, &v.Events[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// #WREQ
func writeTranslateCoordinates(w *Writer, srcWindow, dstWindow Window, srcX, srcY int16) {
	w.WritePad(4)
	w.Write4b(uint32(srcWindow))
	w.Write4b(uint32(dstWindow))

	w.Write2b(uint16(srcX))
	w.Write2b(uint16(srcY))
}

type TranslateCoordinatesReply struct {
	SameScreen bool
	Child      Window
	DstX       int16
	DstY       int16
}

func readTranslateCoordinatesReply(r *Reader, v *TranslateCoordinatesReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SameScreen = Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Child = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.DstX = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.DstY = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(16)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeWarpPointer(w *Writer, srcWindow, dstWindow Window, srcX, srcY int16,
	srcWidth, srcHeight uint16, dstX, dstY int16) {
	w.WritePad(4)
	w.Write4b(uint32(srcWindow))
	w.Write4b(uint32(dstWindow))
	w.Write2b(uint16(srcX))
	w.Write2b(uint16(srcY))
	w.Write2b(srcWidth)
	w.Write2b(srcHeight)
	w.Write2b(uint16(dstX))
	w.Write2b(uint16(dstY))
}

// #WREQ
func writeSetInputFocus(w *Writer, revertTo uint8, focus Window, time Timestamp) {
	w.WritePad(1)
	w.Write1b(revertTo)
	w.WritePad(2)

	w.Write4b(uint32(focus))
	w.Write4b(uint32(time))
}

// #WREQ
func writeGetInputFocus(w *Writer) {
	w.WritePad(4)
}

type GetInputFocusReply struct {
	RevertTo uint8
	Focus    Window
}

func readGetInputFocusReply(r *Reader, v *GetInputFocusReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.RevertTo = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Focus = Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(20)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeQueryKeymap(w *Writer) {
	w.WritePad(4)
}

type QueryKeymapReply struct {
	Keys []byte
}

func readQueryKeymapReply(r *Reader, v *QueryKeymapReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	// keys
	v.Keys = r.ReadBytes(32)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeOpenFont(w *Writer, fid Font, name string) {
	w.WritePad(4)
	w.Write4b(uint32(fid))

	nameLen := uint16(len(name))
	w.Write2b(nameLen)
	w.WritePad(2)

	w.WriteString(name)
	w.WritePad(Pad(len(name)))
}

// #WREQ
func writeCloseFont(w *Writer, font Font) {
	w.WritePad(4)
	w.Write4b(uint32(font))
}

// #WREQ
func writeQueryFont(w *Writer, font Fontable) {
	w.WritePad(4)
	w.Write4b(uint32(font))
}

type QueryFontReply struct {
	MinBounds      CharInfo
	MaxBounds      CharInfo
	MinCharOrByte2 uint16
	MaxCharOrByte2 uint16
	DefaultChar    uint16
	PropertiesLen  uint16
	DrawDirection  uint8
	MinByte1       uint8
	MaxByte1       uint8
	AllCharsExist  bool
	FontAscent     int16
	FontDescent    int16
	CharInfosLen   uint32
	Properties     []FontProp
	CharInfos      []CharInfo
}

func readQueryFontReply(r *Reader, v *QueryFontReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	err := readCharInfo(r, &v.MinBounds)
	if err != nil {
		return err
	}

	// unused
	r.ReadPad(4)
	if r.Err() != nil {
		return r.Err()
	}

	err = readCharInfo(r, &v.MaxBounds)
	if err != nil {
		return err
	}

	// unused
	r.ReadPad(4)
	if r.Err() != nil {
		return r.Err()
	}

	v.MinCharOrByte2 = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MaxCharOrByte2 = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.DefaultChar = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	propsLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.DrawDirection = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MinByte1 = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MaxByte1 = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.AllCharsExist = Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.FontAscent = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.FontDescent = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	charInfosLen := int(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	if propsLen > 0 {
		v.Properties = make([]FontProp, propsLen)
		for i := 0; i < propsLen; i++ {
			err := readFontProp(r, &v.Properties[i])
			if err != nil {
				return err
			}
		}
	}

	if charInfosLen > 0 {
		v.CharInfos = make([]CharInfo, charInfosLen)
		for i := 0; i < charInfosLen; i++ {
			err := readCharInfo(r, &v.CharInfos[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type CharInfo struct {
	LeftSideBearing  int16
	RightSideBearing int16
	CharacterWidth   int16
	Ascent           int16
	Descent          int16
	Attributes       uint16
}

func readCharInfo(r *Reader, v *CharInfo) error {
	v.LeftSideBearing = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.RightSideBearing = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.CharacterWidth = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Ascent = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Descent = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Attributes = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}
	return nil
}

type FontProp struct {
	Name  Atom
	Value uint32
}

func readFontProp(r *Reader, v *FontProp) error {
	v.Name = Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Value = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// TODO: QueryTextExtents

// #WREQ
func writeListFonts(w *Writer, maxNames uint16, pattern string) {
	w.WritePad(4)
	w.Write2b(maxNames)
	patternLen := uint16(len(pattern))
	w.Write2b(patternLen)

	w.WriteString(pattern)
	w.WritePad(Pad(len(pattern)))
}

type ListFontsReply struct {
	Names []string
}

func readListFontsReply(r *Reader, v *ListFontsReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	namesLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(22)
	if r.Err() != nil {
		return r.Err()
	}

	// names
	var n int
	if namesLen > 0 {
		names := make([]string, namesLen)
		for i := 0; i < namesLen; i++ {
			var str Str
			nn, err := readStr(r, &str)
			if err != nil {
				return err
			}
			n += nn
			names[i] = str.Value
		}
		v.Names = names
	}

	// unused
	r.ReadPad(Pad(n))
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeListFontsWithInfo(w *Writer, maxNames uint16, pattern string) {
	w.WritePad(4)
	w.Write2b(maxNames)
	patternLen := uint16(len(pattern))
	w.Write2b(patternLen)

	w.WriteString(pattern)
	w.WritePad(Pad(len(pattern)))
}

type ListFontsWithInfoReply struct {
	LastReply      bool
	MinBounds      CharInfo
	MaxBounds      CharInfo
	MinCharOrByte2 uint16
	MaxCharOrByte2 uint16
	DefaultChar    uint16
	PropertiesLen  uint16
	DrawDirection  uint8
	MinByte1       uint8
	MaxByte1       uint8
	AllCharsExist  bool
	FontAscent     int16
	FontDescent    int16
	RepliesHint    uint32
	Properties     []FontProp
	Name           string
}

func readListFontsWithInfoReply(r *Reader, v *ListFontsWithInfoReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	lastReplyIndicator := r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	replyLen := int(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	if lastReplyIndicator == 0 {
		v.LastReply = true
		r.ReadPad(52)
		if r.Err() != nil {
			return r.Err()
		}
		return nil
	} else {
		v.LastReply = false
	}

	err := readCharInfo(r, &v.MinBounds)
	if err != nil {
		return err
	}

	// unused
	r.ReadPad(4)
	if r.Err() != nil {
		return r.Err()
	}

	err = readCharInfo(r, &v.MaxBounds)
	if err != nil {
		return err
	}

	// unused
	r.ReadPad(4)
	if r.Err() != nil {
		return r.Err()
	}

	v.MinCharOrByte2 = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MaxCharOrByte2 = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.DefaultChar = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	propsLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.DrawDirection = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MinByte1 = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MaxByte1 = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.AllCharsExist = Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.FontAscent = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.FontDescent = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.RepliesHint = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	if propsLen > 0 {
		v.Properties = make([]FontProp, propsLen)
		for i := 0; i < propsLen; i++ {
			err := readFontProp(r, &v.Properties[i])
			if err != nil {
				return err
			}
		}
	}

	nameLen := (replyLen - 7 - (2 * propsLen)) * 4
	nameBytes := r.ReadBytes(nameLen)
	if r.Err() != nil {
		return r.Err()
	}

	zeroIdx := bytes.IndexByte(nameBytes, 0)
	if zeroIdx != -1 {
		nameBytes = nameBytes[:zeroIdx]
	}
	v.Name = string(nameBytes)
	return nil
}

// #WREQ
func writeSetFontPath(w *Writer, paths []string) {
	w.WritePad(4)
	pathsLen := uint16(len(paths))
	w.Write2b(pathsLen)
	w.WritePad(2)

	var n int
	for _, path := range paths {
		n += writeStr(w, path)
	}
	w.WritePad(Pad(n))
}

// #WREQ
func writeGetFontPath(w *Writer) {
	w.WritePad(4)
}

type GetFontPathReply struct {
	Paths []string
}

func readGetFontPathReply(r *Reader, v *GetFontPathReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	pathsLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(22)
	if r.Err() != nil {
		return r.Err()
	}

	var n int
	if pathsLen > 0 {
		v.Paths = make([]string, pathsLen)
		for i := 0; i < pathsLen; i++ {
			var str Str
			nn, err := readStr(r, &str)
			if err != nil {
				return err
			}
			n += nn
			v.Paths[i] = str.Value
		}
	}

	r.ReadPad(Pad(n))
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeCreatePixmap(w *Writer, depth uint8, pid Pixmap, drawable Drawable,
	width, height uint16) {
	w.WritePad(1)
	w.Write1b(depth)
	w.WritePad(2)

	w.Write4b(uint32(pid))
	w.Write4b(uint32(drawable))

	w.Write2b(width)
	w.Write2b(height)
}

// #WREQ
func writeFreePixmap(w *Writer, pixmap Pixmap) {
	w.WritePad(4)
	w.Write4b(uint32(pixmap))
}

// #WREQ
func writeCreateGC(w *Writer, cid GContext, drawable Drawable, valueMask uint32,
	valueList []uint32) {

	w.WritePad(4)
	w.Write4b(uint32(cid))
	w.Write4b(uint32(drawable))
	w.Write4b(valueMask)
	for _, value := range valueList {
		w.Write4b(value)
	}
}

// #WREQ
func writeChangeGC(w *Writer, gc GContext, valueMask uint32, valueList []uint32) {
	w.WritePad(4)
	w.Write4b(uint32(gc))
	w.Write4b(valueMask)
	for _, value := range valueList {
		w.Write4b(value)
	}
}

// #WREQ
func writeCopyGC(w *Writer, srcGC, dstGC GContext, valueMask uint32) {
	w.WritePad(4)
	w.Write4b(uint32(srcGC))
	w.Write4b(uint32(dstGC))
	w.Write4b(valueMask)
}

// #WREQ
func writeSetDashes(w *Writer, gc GContext, dashOffset uint16, dashes []uint8) {
	w.WritePad(4)
	w.Write4b(uint32(gc))
	w.Write2b(dashOffset)
	dashsLen := uint16(len(dashes))
	w.Write2b(dashsLen)
	w.WriteBytes(dashes)
	w.WritePad(Pad(len(dashes)))
}

// #WREQ
func writeSetClipRectangles(w *Writer, ordering uint8, gc GContext,
	clipXOrigin, clipYOrigin int16, rectangles []Rectangle) {

	w.WritePad(1)
	w.Write1b(ordering)
	w.WritePad(2)

	w.Write4b(uint32(gc))
	w.Write2b(uint16(clipXOrigin))
	w.Write2b(uint16(clipYOrigin))
	for _, rect := range rectangles {
		WriteRectangle(w, rect)
	}
}

// #WREQ
func writeFreeGC(w *Writer, gc GContext) {
	w.WritePad(4)
	w.Write4b(uint32(gc))
}

// #WREQ
func writeClearArea(w *Writer, exposures bool, window Window, x, y int16,
	width, height uint16) {

	w.WritePad(1)
	w.Write1b(BoolToUint8(exposures))
	w.WritePad(2)

	w.Write4b(uint32(window))
	w.Write2b(uint16(x))
	w.Write2b(uint16(y))
	w.Write2b(width)
	w.Write2b(height)
}

// #WREQ
func writeCopyArea(w *Writer, srcDrawable, dstDrawable Drawable, gc GContext,
	srcX, srcY, dstX, dstY int16, width, height uint16) {

	w.WritePad(4)
	w.Write4b(uint32(srcDrawable))
	w.Write4b(uint32(dstDrawable))
	w.Write4b(uint32(gc))
	w.Write2b(uint16(srcX))
	w.Write2b(uint16(srcY))
	w.Write2b(uint16(dstX))
	w.Write2b(uint16(dstY))
	w.Write2b(width)
	w.Write2b(height)
}

// #WREQ
func writeCopyPlane(w *Writer, srcDrawable, dstDrawable Drawable, gc GContext,
	srcX, srcY, dstX, dstY int16, width, height uint16, bitPlane uint32) {

	w.WritePad(4)
	w.Write4b(uint32(srcDrawable))
	w.Write4b(uint32(dstDrawable))
	w.Write4b(uint32(gc))
	w.Write2b(uint16(srcX))
	w.Write2b(uint16(srcY))
	w.Write2b(uint16(dstX))
	w.Write2b(uint16(dstY))
	w.Write2b(width)
	w.Write2b(height)
	w.Write4b(bitPlane)
}

type Point struct {
	X, Y int16
}

func writePoint(w *Writer, v Point) {
	w.Write2b(uint16(v.X))
	w.Write2b(uint16(v.Y))
}

// #WREQ
func writePolyPoint(w *Writer, coordinateMode uint8, drawable Drawable, gc GContext,
	points []Point) {

	w.WritePad(1)
	w.Write1b(coordinateMode)
	w.WritePad(2)
	w.Write4b(uint32(drawable))
	w.Write4b(uint32(gc))
	for _, p := range points {
		writePoint(w, p)
	}
}

// #WREQ
func writePolyLine(w *Writer, coordinateMode uint8, drawable Drawable, gc GContext,
	points []Point) {

	w.WritePad(1)
	w.Write1b(coordinateMode)
	w.WritePad(2)
	w.Write4b(uint32(drawable))
	w.Write4b(uint32(gc))
	for _, p := range points {
		writePoint(w, p)
	}
}

// #WREQ
func writePolySegment(w *Writer, drawable Drawable, gc GContext, segments []Segment) {
	w.WritePad(4)
	w.Write4b(uint32(drawable))
	w.Write4b(uint32(gc))
	for _, seg := range segments {
		writeSegment(w, seg)
	}
}

type Segment struct {
	X1, Y1, X2, Y2 int16
}

func writeSegment(w *Writer, v Segment) {
	w.Write2b(uint16(v.X1))
	w.Write2b(uint16(v.Y1))
	w.Write2b(uint16(v.X2))
	w.Write2b(uint16(v.Y2))
}

// #WREQ
func writePolyRectangle(w *Writer, drawable Drawable, gc GContext,
	rectangles []Rectangle) {

	w.WritePad(4)
	w.Write4b(uint32(drawable))
	w.Write4b(uint32(gc))
	for _, rect := range rectangles {
		WriteRectangle(w, rect)
	}
}

// #WREQ
func writePolyArc(w *Writer, drawable Drawable, gc GContext, arcs []Arc) {
	w.WritePad(4)
	w.Write4b(uint32(drawable))
	w.Write4b(uint32(gc))
	for _, arc := range arcs {
		writeArc(w, arc)
	}
}

type Arc struct {
	X, Y           int16
	Width, Height  uint16
	Angle1, Angle2 int16
}

func writeArc(w *Writer, v Arc) {
	w.Write2b(uint16(v.X))
	w.Write2b(uint16(v.Y))
	w.Write2b(v.Width)
	w.Write2b(v.Height)
	w.Write2b(uint16(v.Angle1))
	w.Write2b(uint16(v.Angle2))
}

// #WREQ
func writeFillPoly(w *Writer, drawable Drawable, gc GContext, shape uint8,
	coordinateMode uint8, points []Point) {
	w.WritePad(4)
	w.Write4b(uint32(drawable))
	w.Write4b(uint32(gc))
	w.Write1b(shape)
	w.Write1b(coordinateMode)
	w.WritePad(2)
	for _, p := range points {
		writePoint(w, p)
	}
}

// #WREQ
func writePolyFillRectangle(w *Writer, drawable Drawable, gc GContext,
	rectangles []Rectangle) {

	w.WritePad(4)
	w.Write4b(uint32(drawable))
	w.Write4b(uint32(gc))
	for _, rect := range rectangles {
		WriteRectangle(w, rect)
	}
}

// #WREQ
func writePolyFillArc(w *Writer, drawable Drawable, gc GContext, arcs []Arc) {
	w.WritePad(4)
	w.Write4b(uint32(drawable))
	w.Write4b(uint32(gc))
	for _, arc := range arcs {
		writeArc(w, arc)
	}
}

// #WREQ
func writePutImage(w *Writer, format uint8, drawable Drawable, gc GContext,
	width, height uint16, dstX, dstY int16, leftPad, depth uint8, data []byte) {

	w.WritePad(1)
	w.Write1b(format)
	w.WritePad(2)

	w.Write4b(uint32(drawable))
	w.Write4b(uint32(gc))
	w.Write2b(width)
	w.Write2b(height)
	w.Write2b(uint16(dstX))
	w.Write2b(uint16(dstY))
	w.Write1b(leftPad)
	w.Write1b(depth)
	w.WritePad(2)
	w.WriteBytes(data)
	w.WritePad(Pad(len(data)))
}

// #WREQ
func writeGetImage(w *Writer, format uint8, drawable Drawable, x, y int16,
	width, height uint16, planeMask uint32) {
	w.WritePad(1)
	w.Write1b(format)
	w.WritePad(2)

	w.Write4b(uint32(drawable))

	w.Write2b(uint16(x))
	w.Write2b(uint16(y))

	w.Write2b(width)
	w.Write2b(height)

	w.Write4b(planeMask)
}

type GetImageReply struct {
	Depth  uint8
	Visual VisualID
	Data   []byte
}

func readGetImageReply(r *Reader, v *GetImageReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Depth = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	replyLen := r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Visual = VisualID(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(20)
	if r.Err() != nil {
		return r.Err()
	}

	dataLen := int(replyLen) * 4
	v.Data = r.ReadBytes(dataLen)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeQueryExtension(w *Writer, name string) {
	w.WritePad(4)
	nameLen := len(name)
	w.Write2b(uint16(nameLen))
	w.WritePad(2)
	w.WriteString(name)
	w.WritePad(Pad(nameLen))
}

type QueryExtensionReply struct {
	Present     bool
	MajorOpcode uint8
	FirstEvent  uint8
	FirstError  uint8
}

func readQueryExtensionReply(r *Reader, v *QueryExtensionReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Present = Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.MajorOpcode = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstEvent = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstError = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(20)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeListExtensions(w *Writer) {
	w.WritePad(4)
}

type ListExtensionsReply struct {
	Names []string
}

type Str struct {
	Value string
}

func ReadStr(r *Reader) (string, error) {
	var v Str
	_, err := readStr(r, &v)
	if err != nil {
		return "", err
	}
	return v.Value, nil
}

func readStr(r *Reader, v *Str) (int, error) {
	nameLen := int(r.Read1b())
	if r.Err() != nil {
		return 0, r.Err()
	}

	v.Value = r.ReadString(nameLen)
	if r.Err() != nil {
		return 0, r.Err()
	}
	return 1 + nameLen, nil
}

func writeStr(w *Writer, value string) int {
	valueLen := len(value)
	if valueLen > 0xff {
		valueLen = 0xff
		value = value[:0xff]
	}

	w.Write1b(uint8(valueLen))
	w.WriteString(value)
	return 1 + valueLen
}

func readListExtensionsReply(r *Reader, v *ListExtensionsReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	strsLen := int(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(24)
	if r.Err() != nil {
		return r.Err()
	}

	var n int
	if strsLen > 0 {
		names := make([]string, strsLen)
		for i := 0; i < strsLen; i++ {
			var str Str
			nn, err := readStr(r, &str)
			if err != nil {
				return err
			}
			n += nn
			names[i] = str.Value
		}
		v.Names = names
	}

	// unused
	r.ReadPad(Pad(n))
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeKillClient(w *Writer, resource uint32) {
	w.WritePad(4)
	w.Write4b(resource)
}

// #WREQ
func writeGetKeyboardMapping(w *Writer, firstKeycode Keycode, count uint8) {
	w.WritePad(4)
	w.Write1b(uint8(firstKeycode))
	w.Write1b(count)
	w.WritePad(2)
}

type GetKeyboardMappingReply struct {
	KeysymsPerKeycode uint8
	Keysyms           []Keysym
}

func readGetKeyboardMappingReply(r *Reader, v *GetKeyboardMappingReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.KeysymsPerKeycode = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	keysymsLen := int(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(24)
	if r.Err() != nil {
		return r.Err()
	}

	v.Keysyms = make([]Keysym, keysymsLen)
	for i := 0; i < keysymsLen; i++ {
		v.Keysyms[i] = Keysym(r.Read4b())
		if r.Err() != nil {
			return r.Err()
		}
	}

	return nil
}

// #WREQ
func writeSetScreenSaver(w *Writer, timeout, interval int16, preferBlanking, allowExposures uint8) {
	w.WritePad(4)
	w.Write2b(uint16(timeout))
	w.Write2b(uint16(interval))

	w.Write1b(preferBlanking)
	w.Write1b(allowExposures)
	w.WritePad(2)
}

// #WREQ
func writeGetScreenSaver(w *Writer) {
	w.WritePad(4)
}

type GetScreenSaverReply struct {
	Timeout        uint16
	Interval       uint16
	PreferBlanking uint8
	AllowExposures uint8
}

func readGetScreenSaverReply(r *Reader, v *GetScreenSaverReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Timeout = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Interval = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.PreferBlanking = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.AllowExposures = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(18)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeForceScreenSaver(w *Writer, mode uint8) {
	w.WritePad(1)
	w.Write1b(mode)
	w.WritePad(2)
}

// #WREQ
func writeNoOperation(w *Writer, n int) {
	w.WritePad(4)
	for i := 0; i < n; i++ {
		w.WritePad(4)
	}
}
