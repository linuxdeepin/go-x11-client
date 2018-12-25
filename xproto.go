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

	r.ReadPad(Pad(int(vendorLen)))
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

func WriteRectangle(b *FixedSizeBuf, v Rectangle) {
	b.Write2b(uint16(v.X))
	b.Write2b(uint16(v.Y))
	b.Write2b(v.Width)
	b.Write2b(v.Height)
}

// #WREQ
func encodeCreateWindow(depth uint8, wid, parent Window, x, y int16,
	width, height, borderWidth uint16, class uint16, visual VisualID,
	valueMask uint32, valueList []uint32) (hd uint8, b RequestBody) {

	hd = depth
	b0 := b.AddBlock(7 + len(valueList)).
		Write4b(uint32(wid)).
		Write4b(uint32(parent)).
		Write2b(uint16(x)).
		Write2b(uint16(y)).
		Write2b(width).
		Write2b(height).
		Write2b(borderWidth).
		Write2b(class).
		Write4b(uint32(visual)).
		Write4b(valueMask)
	for _, value := range valueList {
		b0.Write4b(value)
	}
	b0.End()
	return
}

// #WREQ
func encodeChangeWindowAttributes(window Window, valueMask uint32,
	valueList []uint32) (hd uint8, b RequestBody) {

	b0 := b.AddBlock(2 + len(valueList)).
		Write4b(uint32(window)).
		Write4b(valueMask)
	for _, value := range valueList {
		b0.Write4b(value)
	}
	b0.End()
	return
}

// #WREQ
func encodeGetWindowAttributes(window Window) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
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
func encodeDestroyWindow(window Window) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

// #WREQ
func encodeDestroySubwindows(window Window) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

// #WREQ
func encodeChangeSaveSet(mode uint8, window Window) (hd uint8, b RequestBody) {
	hd = mode
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

// #WREQ
func encodeReparentWindow(window, parent Window, x, y int16) (hd uint8, b RequestBody) {
	b.AddBlock(3).
		Write4b(uint32(window)).
		Write4b(uint32(parent)).
		Write2b(uint16(x)).
		Write2b(uint16(y)).
		End()
	return
}

// #WREQ
func encodeMapWindow(window Window) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

// #WREQ
func encodeMapSubwindows(window Window) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

// #WREQ
func encodeUnmapWindow(window Window) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

// #WREQ
func encodeUnmapSubwindows(window Window) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

// #WREQ
func encodeConfigureWindow(window Window, valueMask uint16, valueList []uint32) (hd uint8, b RequestBody) {
	b0 := b.AddBlock(2 + len(valueList)).
		Write4b(uint32(window)).
		Write2b(valueMask).
		WritePad(2)
	for _, value := range valueList {
		b0.Write4b(value)
	}
	b0.End()
	return
}

// #WREQ
func encodeCirculateWindow(direction uint8, window Window) (hd uint8, b RequestBody) {
	hd = direction
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

// #WREQ
func encodeGetGeometry(drawable Drawable) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(drawable)).
		End()
	return
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
func encodeQueryTree(window Window) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
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
func encodeInternAtom(onlyIfExists bool, name string) (hd uint8, b RequestBody) {
	hd = BoolToUint8(onlyIfExists)
	name = TruncateStr(name, math.MaxUint16)
	nameLen := len(name)
	b.AddBlock(1 + SizeIn4bWithPad(nameLen)).
		Write2b(uint16(nameLen)).
		WritePad(2).
		WriteString(name).
		WritePad(Pad(nameLen)).
		End()
	return
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
func encodeGetAtomName(atom Atom) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(atom)).
		End()
	return
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
func encodeChangeProperty(mode uint8, window Window, property, type0 Atom,
	format uint8, data []byte) (hd uint8, b RequestBody) {
	hd = mode

	dataLen := len(data)
	b.AddBlock(5).
		Write4b(uint32(window)).
		Write4b(uint32(property)).
		Write4b(uint32(type0)).
		Write1b(format).
		WritePad(3).
		Write4b(uint32(dataLen / (int(format) / 8))).
		End()
	b.AddBytes(data)
	return
}

// #WREQ
func encodeDeleteProperty(window Window, property Atom) (hd uint8, b RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(window)).
		Write4b(uint32(property)).
		End()
	return
}

// #WREQ
func encodeGetProperty(delete bool, window Window, property, type0 Atom,
	longOffset, longLength uint32) (hd uint8, b RequestBody) {

	hd = BoolToUint8(delete)
	b.AddBlock(5).
		Write4b(uint32(window)).
		Write4b(uint32(property)).
		Write4b(uint32(type0)).
		Write4b(longOffset).
		Write4b(longLength).
		End()
	return
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
func encodeListProperties(window Window) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
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
func encodeSetSelectionOwner(owner Window, selection Atom, time Timestamp) (hd uint8, b RequestBody) {
	b.AddBlock(3).
		Write4b(uint32(owner)).
		Write4b(uint32(selection)).
		Write4b(uint32(time)).
		End()
	return
}

// #WREQ
func encodeGetSelectionOwner(selection Atom) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(selection)).
		End()
	return
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
func encodeConvertSelection(requestor Window, selection, target, property Atom,
	time Timestamp) (hd uint8, b RequestBody) {
	b.AddBlock(5).
		Write4b(uint32(requestor)).
		Write4b(uint32(selection)).
		Write4b(uint32(target)).
		Write4b(uint32(property)).
		Write4b(uint32(time)).
		End()
	return
}

// #WREQ
func encodeSendEvent(propagate bool, destination Window, eventMask uint32,
	event []byte) (hd uint8, b RequestBody) {
	hd = BoolToUint8(propagate)
	b.AddBlock(2).
		Write4b(uint32(destination)).
		Write4b(eventMask)
	b.AddBytes(event)
	return
}

// #WREQ
func encodeGrabPointer(ownerEvents bool, grabWindow Window,
	eventMask uint16, pointerMode, keyboardMode uint8, confineTo Window,
	cursor Cursor, time Timestamp) (hd uint8, b RequestBody) {
	hd = BoolToUint8(ownerEvents)
	b.AddBlock(5).
		Write4b(uint32(grabWindow)).
		Write2b(eventMask).
		Write1b(pointerMode).
		Write1b(keyboardMode).
		Write4b(uint32(confineTo)).
		Write4b(uint32(cursor)).
		Write4b(uint32(time)).
		End()
	return
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
func encodeUngrabPointer(time Timestamp) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(time)).
		End()
	return
}

// #WREQ
func encodeGrabButton(ownerEvents bool, grabWindow Window, eventMask uint16,
	pointerMode, keyboardMode uint8, confineTo Window, cursor Cursor, button uint8,
	modifiers uint16) (hd uint8, b RequestBody) {
	hd = BoolToUint8(ownerEvents)
	b.AddBlock(5).
		Write4b(uint32(grabWindow)).
		Write2b(eventMask).
		Write1b(pointerMode).
		Write1b(keyboardMode).
		Write4b(uint32(confineTo)).
		Write4b(uint32(cursor)).
		Write1b(button).
		WritePad(1).
		Write2b(modifiers).
		End()
	return
}

// #WREQ
func encodeUngrabButton(button uint8, grabWindow Window,
	modifiers uint16) (hd uint8, b RequestBody) {
	hd = button
	b.AddBlock(2).
		Write4b(uint32(grabWindow)).
		Write2b(modifiers).
		WritePad(2).
		End()
	return
}

// #WREQ
func encodeChangeActivePointerGrab(cursor Cursor, time Timestamp,
	eventMask uint16) (hd uint8, b RequestBody) {
	b.AddBlock(3).
		Write4b(uint32(cursor)).
		Write4b(uint32(time)).
		Write2b(eventMask).
		WritePad(2).
		End()
	return
}

// #WREQ
func encodeGrabKeyboard(ownerEvents bool, grabWindow Window, time Timestamp,
	pointerMode, keyboardMode uint8) (hd uint8, b RequestBody) {
	hd = BoolToUint8(ownerEvents)
	b.AddBlock(3).
		Write4b(uint32(grabWindow)).
		Write4b(uint32(time)).
		Write1b(pointerMode).
		Write1b(keyboardMode).
		WritePad(2).
		End()
	return
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
func encodeUngrabKeyboard(time Timestamp) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(time)).
		End()
	return
}

// #WREQ
func encodeGrabKey(ownerEvents bool, grabWindow Window, modifiers uint16,
	key Keycode, pointerMode, keyboardMode uint8) (hd uint8, b RequestBody) {
	hd = BoolToUint8(ownerEvents)
	b.AddBlock(3).
		Write4b(uint32(grabWindow)).
		Write2b(modifiers).
		Write1b(uint8(key)).
		Write1b(pointerMode).
		Write1b(keyboardMode).
		WritePad(3).
		End()
	return
}

// #WREQ
func encodeUngrabKey(key Keycode, grabWindow Window, modifiers uint16) (hd uint8, b RequestBody) {
	hd = uint8(key)
	b.AddBlock(2).
		Write4b(uint32(grabWindow)).
		Write2b(modifiers).
		WritePad(2).
		End()
	return
}

// #WREQ
func encodeAllowEvents(mode uint8, time Timestamp) (hd uint8, b RequestBody) {
	hd = mode
	b.AddBlock(1).
		Write4b(uint32(time)).
		End()
	return
}

// #WREQ
func encodeGrabServer() (hd uint8, b RequestBody) {
	return
}

// #WREQ
func encodeUngrabServer() (hd uint8, b RequestBody) {
	return
}

// #WREQ
func encodeQueryPointer(window Window) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
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
func encodeGetMotionEvents(window Window, start, stop Timestamp) (hd uint8, b RequestBody) {
	b.AddBlock(3).
		Write4b(uint32(window)).
		Write4b(uint32(start)).
		Write4b(uint32(stop)).
		End()
	return
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
func encodeTranslateCoordinates(srcWindow, dstWindow Window,
	srcX, srcY int16) (hd uint8, b RequestBody) {
	b.AddBlock(3).
		Write4b(uint32(srcWindow)).
		Write4b(uint32(dstWindow)).
		Write2b(uint16(srcX)).
		Write2b(uint16(srcY)).
		End()
	return
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
func encodeWarpPointer(srcWindow, dstWindow Window, srcX, srcY int16,
	srcWidth, srcHeight uint16, dstX, dstY int16) (hd uint8, b RequestBody) {
	b.AddBlock(5).
		Write4b(uint32(srcWindow)).
		Write4b(uint32(dstWindow)).
		Write2b(uint16(srcX)).
		Write2b(uint16(srcY)).
		Write2b(srcWidth).
		Write2b(srcHeight).
		Write2b(uint16(dstX)).
		Write2b(uint16(dstY)).
		End()
	return
}

// #WREQ
func encodeSetInputFocus(revertTo uint8, focus Window, time Timestamp) (hd uint8, b RequestBody) {
	hd = revertTo
	b.AddBlock(2).
		Write4b(uint32(focus)).
		Write4b(uint32(time)).
		End()
	return
}

// #WREQ
func encodeGetInputFocus() (hd uint8, b RequestBody) {
	return
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
func encodeQueryKeymap() (hd uint8, b RequestBody) {
	return
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
func encodeOpenFont(fid Font, name string) (hd uint8, b RequestBody) {
	name = TruncateStr(name, math.MaxUint16)
	nameLen := len(name)
	b.AddBlock(2 + SizeIn4bWithPad(nameLen)).
		Write4b(uint32(fid)).
		Write2b(uint16(nameLen)).
		WritePad(2).
		WriteString(name).
		WritePad(Pad(nameLen)).
		End()
	return
}

// #WREQ
func encodeCloseFont(font Font) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(font)).
		End()
	return
}

// #WREQ
func encodeQueryFont(font Fontable) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(font)).
		End()
	return
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
func encodeListFonts(maxNames uint16, pattern string) (hd uint8, b RequestBody) {
	pattern = TruncateStr(pattern, math.MaxUint16)
	patternLen := len(pattern)
	b.AddBlock(1 + SizeIn4bWithPad(patternLen)).
		Write2b(maxNames).
		Write2b(uint16(patternLen)).
		WriteString(pattern).
		WritePad(Pad(patternLen)).
		End()
	return
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
func encodeListFontsWithInfo(maxNames uint16, pattern string) (hd uint8, b RequestBody) {
	pattern = TruncateStr(pattern, math.MaxUint16)
	patternLen := len(pattern)
	b.AddBlock(1 + SizeIn4bWithPad(patternLen)).
		Write2b(maxNames).
		Write2b(uint16(patternLen)).
		WriteString(pattern).
		WritePad(Pad(patternLen)).
		End()
	return
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
func encodeSetFontPath(paths []string) (hd uint8, b RequestBody) {
	pathsLen := uint16(len(paths))

	var n int
	for _, p := range paths {
		pLen := len(p)
		if pLen > math.MaxUint8 {
			pLen = math.MaxUint8
		}
		n += 1 + pLen
	}

	b0 := b.AddBlock(1 + SizeIn4bWithPad(n)).
		Write2b(pathsLen).
		WritePad(2)

	for _, p := range paths {
		writeStr(b0, p)
	}
	b0.WritePad(Pad(n)).End()
	return
}

// #WREQ
func encodeGetFontPath() (hd uint8, b RequestBody) {
	return
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
func encodeCreatePixmap(depth uint8, pid Pixmap, drawable Drawable,
	width, height uint16) (hd uint8, b RequestBody) {
	hd = depth
	b.AddBlock(3).
		Write4b(uint32(pid)).
		Write4b(uint32(drawable)).
		Write2b(width).
		Write2b(height).
		End()
	return
}

// #WREQ
func encodeFreePixmap(pixmap Pixmap) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(pixmap)).
		End()
	return
}

// #WREQ
func encodeCreateGC(cid GContext, drawable Drawable, valueMask uint32,
	valueList []uint32) (hd uint8, b RequestBody) {

	b0 := b.AddBlock(3 + len(valueList)).
		Write4b(uint32(cid)).
		Write4b(uint32(drawable)).
		Write4b(valueMask)

	for _, value := range valueList {
		b0.Write4b(value)
	}
	b0.End()
	return
}

// #WREQ
func encodeChangeGC(gc GContext, valueMask uint32, valueList []uint32) (hd uint8, b RequestBody) {
	b0 := b.AddBlock(2 + len(valueList)).
		Write4b(uint32(gc)).
		Write4b(valueMask)
	for _, value := range valueList {
		b0.Write4b(value)
	}
	b0.End()
	return
}

// #WREQ
func encodeCopyGC(srcGC, dstGC GContext, valueMask uint32) (hd uint8, b RequestBody) {
	b.AddBlock(3).
		Write4b(uint32(srcGC)).
		Write4b(uint32(dstGC)).
		Write4b(valueMask).
		End()
	return
}

// #WREQ
func encodeSetDashes(gc GContext, dashOffset uint16, dashes []uint8) (hd uint8, b RequestBody) {
	dashesLen := uint16(len(dashes))
	b.AddBlock(2).
		Write4b(uint32(gc)).
		Write2b(dashOffset).
		Write2b(dashesLen).
		End()
	b.AddBytes(dashes)
	return
}

// #WREQ
func encodeSetClipRectangles(ordering uint8, gc GContext,
	clipXOrigin, clipYOrigin int16, rectangles []Rectangle) (hd uint8, b RequestBody) {
	hd = ordering
	b0 := b.AddBlock(2 + 2*len(rectangles)).
		Write4b(uint32(gc)).
		Write2b(uint16(clipXOrigin)).
		Write2b(uint16(clipYOrigin))

	for _, rect := range rectangles {
		WriteRectangle(b0, rect)
	}
	b0.End()
	return
}

// #WREQ
func encodeFreeGC(gc GContext) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(gc)).
		End()
	return
}

// #WREQ
func encodeClearArea(exposures bool, window Window, x, y int16,
	width, height uint16) (hd uint8, b RequestBody) {
	hd = BoolToUint8(exposures)
	b.AddBlock(3).
		Write4b(uint32(window)).
		Write2b(uint16(x)).
		Write2b(uint16(y)).
		Write2b(width).
		Write2b(height).
		End()
	return
}

// #WREQ
func encodeCopyArea(srcDrawable, dstDrawable Drawable, gc GContext,
	srcX, srcY, dstX, dstY int16, width, height uint16) (hd uint8, b RequestBody) {
	b.AddBlock(6).
		Write4b(uint32(srcDrawable)).
		Write4b(uint32(dstDrawable)).
		Write4b(uint32(gc)).
		Write2b(uint16(srcX)).
		Write2b(uint16(srcY)).
		Write2b(uint16(dstX)).
		Write2b(uint16(dstY)).
		Write2b(width).
		Write2b(height).
		End()
	return
}

// #WREQ
func encodeCopyPlane(srcDrawable, dstDrawable Drawable, gc GContext,
	srcX, srcY, dstX, dstY int16, width, height uint16, bitPlane uint32) (hd uint8, b RequestBody) {

	b.AddBlock(7).
		Write4b(uint32(srcDrawable)).
		Write4b(uint32(dstDrawable)).
		Write4b(uint32(gc)).
		Write2b(uint16(srcX)).
		Write2b(uint16(srcY)).
		Write2b(uint16(dstX)).
		Write2b(uint16(dstY)).
		Write2b(width).
		Write2b(height).
		Write4b(bitPlane).
		End()
	return
}

type Point struct {
	X, Y int16
}

func writePoint(b *FixedSizeBuf, v Point) {
	b.Write2b(uint16(v.X))
	b.Write2b(uint16(v.Y))
}

// #WREQ
func encodePolyPoint(coordinateMode uint8, drawable Drawable, gc GContext,
	points []Point) (hd uint8, b RequestBody) {

	hd = coordinateMode
	b0 := b.AddBlock(2 + len(points)).
		Write4b(uint32(drawable)).
		Write4b(uint32(gc))
	for _, p := range points {
		writePoint(b0, p)
	}
	b0.End()
	return
}

// #WREQ
func encodePolyLine(coordinateMode uint8, drawable Drawable, gc GContext,
	points []Point) (hd uint8, b RequestBody) {
	hd = coordinateMode
	b0 := b.AddBlock(2 + len(points)).
		Write4b(uint32(drawable)).
		Write4b(uint32(gc))

	for _, p := range points {
		writePoint(b0, p)
	}
	b0.End()
	return
}

// #WREQ
func encodePolySegment(drawable Drawable, gc GContext,
	segments []Segment) (hd uint8, b RequestBody) {

	b0 := b.AddBlock(2 + 2*len(segments)).
		Write4b(uint32(drawable)).
		Write4b(uint32(gc))
	for _, seg := range segments {
		writeSegment(b0, seg)
	}
	b0.End()
	return
}

type Segment struct {
	X1, Y1, X2, Y2 int16
}

func writeSegment(w *FixedSizeBuf, v Segment) {
	w.Write2b(uint16(v.X1))
	w.Write2b(uint16(v.Y1))
	w.Write2b(uint16(v.X2))
	w.Write2b(uint16(v.Y2))
}

// #WREQ
func encodePolyRectangle(drawable Drawable, gc GContext,
	rectangles []Rectangle) (hd uint8, b RequestBody) {

	b0 := b.AddBlock(2 + 2*len(rectangles)).
		Write4b(uint32(drawable)).
		Write4b(uint32(gc))
	for _, rect := range rectangles {
		WriteRectangle(b0, rect)
	}
	b0.End()
	return
}

// #WREQ
func encodePolyArc(drawable Drawable, gc GContext, arcs []Arc) (hd uint8, b RequestBody) {
	b0 := b.AddBlock(2 + 3*len(arcs)).
		Write4b(uint32(drawable)).
		Write4b(uint32(gc))
	for _, arc := range arcs {
		writeArc(b0, arc)
	}
	b0.End()
	return
}

type Arc struct {
	X, Y           int16
	Width, Height  uint16
	Angle1, Angle2 int16
}

func writeArc(w *FixedSizeBuf, v Arc) {
	w.Write2b(uint16(v.X))
	w.Write2b(uint16(v.Y))
	w.Write2b(v.Width)
	w.Write2b(v.Height)
	w.Write2b(uint16(v.Angle1))
	w.Write2b(uint16(v.Angle2))
}

// #WREQ
func encodeFillPoly(drawable Drawable, gc GContext, shape uint8,
	coordinateMode uint8, points []Point) (hd uint8, b RequestBody) {
	b0 := b.AddBlock(3 + len(points)).
		Write4b(uint32(drawable)).
		Write4b(uint32(gc)).
		Write1b(shape).
		Write1b(coordinateMode).
		WritePad(2)
	for _, p := range points {
		writePoint(b0, p)
	}
	b0.End()
	return
}

// #WREQ
func encodePolyFillRectangle(drawable Drawable, gc GContext,
	rectangles []Rectangle) (hd uint8, b RequestBody) {

	b0 := b.AddBlock(2 + 2*len(rectangles)).
		Write4b(uint32(drawable)).
		Write4b(uint32(gc))
	for _, rect := range rectangles {
		WriteRectangle(b0, rect)
	}
	b0.End()
	return
}

// #WREQ
func encodePolyFillArc(drawable Drawable, gc GContext, arcs []Arc) (hd uint8, b RequestBody) {
	b0 := b.AddBlock(2 + 3*len(arcs)).
		Write4b(uint32(drawable)).
		Write4b(uint32(gc))

	for _, arc := range arcs {
		writeArc(b0, arc)
	}
	b0.End()
	return
}

// #WREQ
func encodePutImage(format uint8, drawable Drawable, gc GContext, width, height uint16,
	dstX, dstY int16, leftPad, depth uint8, data []byte) (hd uint8, b RequestBody) {
	hd = format
	b.AddBlock(5).
		Write4b(uint32(drawable)).
		Write4b(uint32(gc)).
		Write2b(width).
		Write2b(height).
		Write2b(uint16(dstX)).
		Write2b(uint16(dstY)).
		Write1b(leftPad).
		Write1b(depth).
		WritePad(2).
		End()
	b.AddBytes(data)
	return
}

// #WREQ
func encodeGetImage(format uint8, drawable Drawable, x, y int16,
	width, height uint16, planeMask uint32) (hd uint8, b RequestBody) {
	hd = format
	b.AddBlock(4).
		Write4b(uint32(drawable)).
		Write2b(uint16(x)).
		Write2b(uint16(y)).
		Write2b(width).
		Write2b(height).
		Write4b(planeMask).
		End()
	return
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
func encodeQueryExtension(name string) (hd uint8, b RequestBody) {
	name = TruncateStr(name, math.MaxUint16)
	nameLen := len(name)
	b.AddBlock(1 + SizeIn4bWithPad(nameLen)).
		Write2b(uint16(nameLen)).
		WritePad(2).
		WriteString(name).
		WritePad(Pad(nameLen)).
		End()
	return
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
func encodeListExtensions() (hd uint8, b RequestBody) {
	return
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

func writeStr(b *FixedSizeBuf, str string) {
	str = TruncateStr(str, math.MaxUint8)
	b.Write1b(uint8(len(str)))
	b.WriteString(str)
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
func encodeKillClient(resource uint32) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write4b(resource).
		End()
	return
}

// #WREQ
func encodeGetKeyboardMapping(firstKeycode Keycode, count uint8) (hd uint8, b RequestBody) {
	b.AddBlock(1).
		Write1b(uint8(firstKeycode)).
		Write1b(count).
		WritePad(2).
		End()
	return
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
func encodeSetScreenSaver(timeout, interval int16, preferBlanking,
	allowExposures uint8) (hd uint8, b RequestBody) {
	b.AddBlock(2).
		Write2b(uint16(timeout)).
		Write2b(uint16(interval)).
		Write1b(preferBlanking).
		Write1b(allowExposures).
		WritePad(2).
		End()
	return
}

// #WREQ
func encodeGetScreenSaver() (hd uint8, b RequestBody) {
	return
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
func encodeForceScreenSaver(mode uint8) (hd uint8, b RequestBody) {
	hd = mode
	return
}

// #WREQ
func encodeNoOperation(n int) (hd uint8, b RequestBody) {
	b.AddBlock(n)
	return
}
