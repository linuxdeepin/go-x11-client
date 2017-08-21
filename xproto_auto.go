package client

// module.direct_imports: []
import "unsafe"
import "log"
import "sync"
import "fmt"
import "strings"

var _ = unsafe.Alignof(0)
var _ sync.Mutex
var _ = fmt.Println
var _ = strings.Join

func init() {
	log.SetFlags(log.Lshortfile)
}

type Char2B struct {
	Byte1 uint8
	Byte2 uint8
}

type CChar2B struct {
	Byte1 uint8
	Byte2 uint8
}

func Char2BRead(r *Reader, v *Char2B) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Byte1
	v.Byte1 = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Byte2
	v.Byte2 = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}
	return
}

func Char2BReadN(r *Reader, dest []Char2B) (n int) {
	for i := 0; i < len(dest); i++ {
		n += Char2BRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func Char2BWrite(w *Writer, v *Char2B) (n int) {

	// write field Byte1
	w.Write1b(v.Byte1)
	n += 1

	// write field Byte2
	w.Write1b(v.Byte2)
	n += 1
	return
}
func Char2BWriteN(w *Writer, src []Char2B) (n int) {
	for i := 0; i < len(src); i++ {
		n += Char2BWrite(w, &src[i])
	}
	return
}

// simple ('xcb', 'WINDOW')
type Window uint32

// simple ('xcb', 'PIXMAP')
type Pixmap uint32

// simple ('xcb', 'CURSOR')
type Cursor uint32

// simple ('xcb', 'FONT')
type Font uint32

// simple ('xcb', 'GCONTEXT')
type GContext uint32

// simple ('xcb', 'COLORMAP')
type Colormap uint32

// simple ('xcb', 'ATOM')
type Atom uint32

// simple ('xcb', 'DRAWABLE')
type Drawable uint32

// simple ('xcb', 'FONTABLE')
type Fontable uint32

// simple ('xcb', 'BOOL32')
type Bool32 uint32

// simple ('xcb', 'VISUALID')
type VisualID uint32

// simple ('xcb', 'TIMESTAMP')
type Timestamp uint32

// simple ('xcb', 'KEYSYM')
type Keysym uint32

// simple ('xcb', 'KEYCODE')
type Keycode uint8

// simple ('xcb', 'KEYCODE32')
type Keycode32 uint32

// simple ('xcb', 'BUTTON')
type Button uint8
type Point struct {
	X int16
	Y int16
}

type CPoint struct {
	X int16
	Y int16
}

func PointRead(r *Reader, v *Point) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field X
	v.X = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Y
	v.Y = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

func PointReadN(r *Reader, dest []Point) (n int) {
	for i := 0; i < len(dest); i++ {
		n += PointRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func PointWrite(w *Writer, v *Point) (n int) {

	// write field X
	w.Write2b(uint16(v.X))
	n += 2

	// write field Y
	w.Write2b(uint16(v.Y))
	n += 2
	return
}
func PointWriteN(w *Writer, src []Point) (n int) {
	for i := 0; i < len(src); i++ {
		n += PointWrite(w, &src[i])
	}
	return
}

type Rectangle struct {
	X      int16
	Y      int16
	Width  uint16
	Height uint16
}

type CRectangle struct {
	X      int16
	Y      int16
	Width  uint16
	Height uint16
}

func RectangleRead(r *Reader, v *Rectangle) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field X
	v.X = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Y
	v.Y = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Width
	v.Width = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Height
	v.Height = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

func RectangleReadN(r *Reader, dest []Rectangle) (n int) {
	for i := 0; i < len(dest); i++ {
		n += RectangleRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func RectangleWrite(w *Writer, v *Rectangle) (n int) {

	// write field X
	w.Write2b(uint16(v.X))
	n += 2

	// write field Y
	w.Write2b(uint16(v.Y))
	n += 2

	// write field Width
	w.Write2b(v.Width)
	n += 2

	// write field Height
	w.Write2b(v.Height)
	n += 2
	return
}
func RectangleWriteN(w *Writer, src []Rectangle) (n int) {
	for i := 0; i < len(src); i++ {
		n += RectangleWrite(w, &src[i])
	}
	return
}

type Arc struct {
	X      int16
	Y      int16
	Width  uint16
	Height uint16
	Angle1 int16
	Angle2 int16
}

type CArc struct {
	X      int16
	Y      int16
	Width  uint16
	Height uint16
	Angle1 int16
	Angle2 int16
}

func ArcRead(r *Reader, v *Arc) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field X
	v.X = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Y
	v.Y = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Width
	v.Width = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Height
	v.Height = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Angle1
	v.Angle1 = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Angle2
	v.Angle2 = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

func ArcReadN(r *Reader, dest []Arc) (n int) {
	for i := 0; i < len(dest); i++ {
		n += ArcRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func ArcWrite(w *Writer, v *Arc) (n int) {

	// write field X
	w.Write2b(uint16(v.X))
	n += 2

	// write field Y
	w.Write2b(uint16(v.Y))
	n += 2

	// write field Width
	w.Write2b(v.Width)
	n += 2

	// write field Height
	w.Write2b(v.Height)
	n += 2

	// write field Angle1
	w.Write2b(uint16(v.Angle1))
	n += 2

	// write field Angle2
	w.Write2b(uint16(v.Angle2))
	n += 2
	return
}
func ArcWriteN(w *Writer, src []Arc) (n int) {
	for i := 0; i < len(src); i++ {
		n += ArcWrite(w, &src[i])
	}
	return
}

type Format struct {
	Depth        uint8
	BitsPerPixel uint8
	ScanlinePad  uint8
	// Pad0 [5]uint8
}

type CFormat struct {
	Depth        uint8
	BitsPerPixel uint8
	ScanlinePad  uint8
	Pad0         [5]uint8
}

func FormatRead(r *Reader, v *Format) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Depth
	v.Depth = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field BitsPerPixel
	v.BitsPerPixel = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field ScanlinePad
	v.ScanlinePad = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(5)
	n += 5
	if r.Err() != nil {
		return
	}
	return
}

func FormatReadN(r *Reader, dest []Format) (n int) {
	for i := 0; i < len(dest); i++ {
		n += FormatRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func FormatWrite(w *Writer, v *Format) (n int) {

	// write field Depth
	w.Write1b(v.Depth)
	n += 1

	// write field BitsPerPixel
	w.Write1b(v.BitsPerPixel)
	n += 1

	// write field ScanlinePad
	w.Write1b(v.ScanlinePad)
	n += 1

	// write field Pad0
	w.WritePad(5)
	n += 5
	return
}
func FormatWriteN(w *Writer, src []Format) (n int) {
	for i := 0; i < len(src); i++ {
		n += FormatWrite(w, &src[i])
	}
	return
}

// enum VisualClass
const (
	VisualClassStaticGray  = 0
	VisualClassGrayScale   = 1
	VisualClassStaticColor = 2
	VisualClassPseudoColor = 3
	VisualClassTrueColor   = 4
	VisualClassDirectColor = 5
)

func VisualClassEnumToStr(v int) string {
	switch v {
	case 0:
		return "StaticGray"
	case 1:
		return "GrayScale"
	case 2:
		return "StaticColor"
	case 3:
		return "PseudoColor"
	case 4:
		return "TrueColor"
	case 5:
		return "DirectColor"
	default:
		return "<Unknown>"
	}
}

type VisualType struct {
	VisualId        VisualID
	Class           uint8
	BitsPerRgbValue uint8
	ColormapEntries uint16
	RedMask         uint32
	GreenMask       uint32
	BlueMask        uint32
	// Pad0 [4]uint8
}

type CVisualType struct {
	VisualId        VisualID
	Class           uint8
	BitsPerRgbValue uint8
	ColormapEntries uint16
	RedMask         uint32
	GreenMask       uint32
	BlueMask        uint32
	Pad0            [4]uint8
}

func VisualTypeRead(r *Reader, v *VisualType) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field VisualId
	v.VisualId = VisualID(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Class
	v.Class = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field BitsPerRgbValue
	v.BitsPerRgbValue = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field ColormapEntries
	v.ColormapEntries = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field RedMask
	v.RedMask = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field GreenMask
	v.GreenMask = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field BlueMask
	v.BlueMask = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(4)
	n += 4
	if r.Err() != nil {
		return
	}
	return
}

func VisualTypeReadN(r *Reader, dest []VisualType) (n int) {
	for i := 0; i < len(dest); i++ {
		n += VisualTypeRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func VisualTypeWrite(w *Writer, v *VisualType) (n int) {

	// write field VisualId
	w.Write4b(uint32(v.VisualId))
	n += 4

	// write field Class
	w.Write1b(v.Class)
	n += 1

	// write field BitsPerRgbValue
	w.Write1b(v.BitsPerRgbValue)
	n += 1

	// write field ColormapEntries
	w.Write2b(v.ColormapEntries)
	n += 2

	// write field RedMask
	w.Write4b(v.RedMask)
	n += 4

	// write field GreenMask
	w.Write4b(v.GreenMask)
	n += 4

	// write field BlueMask
	w.Write4b(v.BlueMask)
	n += 4

	// write field Pad0
	w.WritePad(4)
	n += 4
	return
}
func VisualTypeWriteN(w *Writer, src []VisualType) (n int) {
	for i := 0; i < len(src); i++ {
		n += VisualTypeWrite(w, &src[i])
	}
	return
}

type Depth struct {
	Depth uint8
	// Pad0 uint8
	VisualsLen uint16
	// Pad1 [4]uint8
	Visuals []VisualType
}

type CDepth struct {
	Depth      uint8
	Pad0       uint8
	VisualsLen uint16
	Pad1       [4]uint8
}

func DepthRead(r *Reader, v *Depth) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Depth
	v.Depth = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field VisualsLen
	v.VisualsLen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(4)
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Visuals
	v.Visuals = make([]VisualType, int(v.VisualsLen))
	blockLen += VisualTypeReadN(r, v.Visuals)
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CVisualType{}))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

func DepthReadN(r *Reader, dest []Depth) (n int) {
	for i := 0; i < len(dest); i++ {
		n += DepthRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func DepthWrite(w *Writer, v *Depth) (n int) {

	// write field Depth
	w.Write1b(v.Depth)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field VisualsLen
	w.Write2b(v.VisualsLen)
	n += 2

	// write field Pad1
	w.WritePad(4)
	n += 4

	// write field Visuals
	return
}
func DepthWriteN(w *Writer, src []Depth) (n int) {
	for i := 0; i < len(src); i++ {
		n += DepthWrite(w, &src[i])
	}
	return
}

// enum EventMask
const (
	EventMaskNoEvent              = 0
	EventMaskKeyPress             = 1
	EventMaskKeyRelease           = 2
	EventMaskButtonPress          = 4
	EventMaskButtonRelease        = 8
	EventMaskEnterWindow          = 16
	EventMaskLeaveWindow          = 32
	EventMaskPointerMotion        = 64
	EventMaskPointerMotionHint    = 128
	EventMaskButton1Motion        = 256
	EventMaskButton2Motion        = 512
	EventMaskButton3Motion        = 1024
	EventMaskButton4Motion        = 2048
	EventMaskButton5Motion        = 4096
	EventMaskButtonMotion         = 8192
	EventMaskKeymapState          = 16384
	EventMaskExposure             = 32768
	EventMaskVisibilityChange     = 65536
	EventMaskStructureNotify      = 131072
	EventMaskResizeRedirect       = 262144
	EventMaskSubstructureNotify   = 524288
	EventMaskSubstructureRedirect = 1048576
	EventMaskFocusChange          = 2097152
	EventMaskPropertyChange       = 4194304
	EventMaskColorMapChange       = 8388608
	EventMaskOwnerGrabButton      = 16777216
)

func EventMaskEnumToStr(v int) string {
	strv := []string{}
	switch {
	case v&EventMaskKeyPress > 0:
		strv = append(strv, "KeyPress")
	case v&EventMaskKeyRelease > 0:
		strv = append(strv, "KeyRelease")
	case v&EventMaskButtonPress > 0:
		strv = append(strv, "ButtonPress")
	case v&EventMaskButtonRelease > 0:
		strv = append(strv, "ButtonRelease")
	case v&EventMaskEnterWindow > 0:
		strv = append(strv, "EnterWindow")
	case v&EventMaskLeaveWindow > 0:
		strv = append(strv, "LeaveWindow")
	case v&EventMaskPointerMotion > 0:
		strv = append(strv, "PointerMotion")
	case v&EventMaskPointerMotionHint > 0:
		strv = append(strv, "PointerMotionHint")
	case v&EventMaskButton1Motion > 0:
		strv = append(strv, "Button1Motion")
	case v&EventMaskButton2Motion > 0:
		strv = append(strv, "Button2Motion")
	case v&EventMaskButton3Motion > 0:
		strv = append(strv, "Button3Motion")
	case v&EventMaskButton4Motion > 0:
		strv = append(strv, "Button4Motion")
	case v&EventMaskButton5Motion > 0:
		strv = append(strv, "Button5Motion")
	case v&EventMaskButtonMotion > 0:
		strv = append(strv, "ButtonMotion")
	case v&EventMaskKeymapState > 0:
		strv = append(strv, "KeymapState")
	case v&EventMaskExposure > 0:
		strv = append(strv, "Exposure")
	case v&EventMaskVisibilityChange > 0:
		strv = append(strv, "VisibilityChange")
	case v&EventMaskStructureNotify > 0:
		strv = append(strv, "StructureNotify")
	case v&EventMaskResizeRedirect > 0:
		strv = append(strv, "ResizeRedirect")
	case v&EventMaskSubstructureNotify > 0:
		strv = append(strv, "SubstructureNotify")
	case v&EventMaskSubstructureRedirect > 0:
		strv = append(strv, "SubstructureRedirect")
	case v&EventMaskFocusChange > 0:
		strv = append(strv, "FocusChange")
	case v&EventMaskPropertyChange > 0:
		strv = append(strv, "PropertyChange")
	case v&EventMaskColorMapChange > 0:
		strv = append(strv, "ColorMapChange")
	case v&EventMaskOwnerGrabButton > 0:
		strv = append(strv, "OwnerGrabButton")
	}
	return "[" + strings.Join(strv, "|") + "]"
}

// enum BackingStore
const (
	BackingStoreNotUseful  = 0
	BackingStoreWhenMapped = 1
	BackingStoreAlways     = 2
)

func BackingStoreEnumToStr(v int) string {
	switch v {
	case 0:
		return "NotUseful"
	case 1:
		return "WhenMapped"
	case 2:
		return "Always"
	default:
		return "<Unknown>"
	}
}

type Screen struct {
	Root                Window
	DefaultColormap     Colormap
	WhitePixel          uint32
	BlackPixel          uint32
	CurrentInputMasks   uint32
	WidthInPixels       uint16
	HeightInPixels      uint16
	WidthInMillimeters  uint16
	HeightInMillimeters uint16
	MinInstalledMaps    uint16
	MaxInstalledMaps    uint16
	RootVisual          VisualID
	BackingStores       uint8
	SaveUnders          uint8
	RootDepth           uint8
	AllowedDepthsLen    uint8
	AllowedDepths       []Depth
}

type CScreen struct {
	Root                Window
	DefaultColormap     Colormap
	WhitePixel          uint32
	BlackPixel          uint32
	CurrentInputMasks   uint32
	WidthInPixels       uint16
	HeightInPixels      uint16
	WidthInMillimeters  uint16
	HeightInMillimeters uint16
	MinInstalledMaps    uint16
	MaxInstalledMaps    uint16
	RootVisual          VisualID
	BackingStores       uint8
	SaveUnders          uint8
	RootDepth           uint8
	AllowedDepthsLen    uint8
}

func ScreenRead(r *Reader, v *Screen) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Root
	v.Root = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field DefaultColormap
	v.DefaultColormap = Colormap(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field WhitePixel
	v.WhitePixel = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field BlackPixel
	v.BlackPixel = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field CurrentInputMasks
	v.CurrentInputMasks = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field WidthInPixels
	v.WidthInPixels = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field HeightInPixels
	v.HeightInPixels = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field WidthInMillimeters
	v.WidthInMillimeters = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field HeightInMillimeters
	v.HeightInMillimeters = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field MinInstalledMaps
	v.MinInstalledMaps = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field MaxInstalledMaps
	v.MaxInstalledMaps = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field RootVisual
	v.RootVisual = VisualID(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field BackingStores
	v.BackingStores = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field SaveUnders
	v.SaveUnders = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field RootDepth
	v.RootDepth = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field AllowedDepthsLen
	v.AllowedDepthsLen = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field AllowedDepths
	v.AllowedDepths = make([]Depth, int(v.AllowedDepthsLen))
	blockLen += DepthReadN(r, v.AllowedDepths)
	n += blockLen
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CDepth{}))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

func ScreenReadN(r *Reader, dest []Screen) (n int) {
	for i := 0; i < len(dest); i++ {
		n += ScreenRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func ScreenWrite(w *Writer, v *Screen) (n int) {

	// write field Root
	w.Write4b(uint32(v.Root))
	n += 4

	// write field DefaultColormap
	w.Write4b(uint32(v.DefaultColormap))
	n += 4

	// write field WhitePixel
	w.Write4b(v.WhitePixel)
	n += 4

	// write field BlackPixel
	w.Write4b(v.BlackPixel)
	n += 4

	// write field CurrentInputMasks
	w.Write4b(v.CurrentInputMasks)
	n += 4

	// write field WidthInPixels
	w.Write2b(v.WidthInPixels)
	n += 2

	// write field HeightInPixels
	w.Write2b(v.HeightInPixels)
	n += 2

	// write field WidthInMillimeters
	w.Write2b(v.WidthInMillimeters)
	n += 2

	// write field HeightInMillimeters
	w.Write2b(v.HeightInMillimeters)
	n += 2

	// write field MinInstalledMaps
	w.Write2b(v.MinInstalledMaps)
	n += 2

	// write field MaxInstalledMaps
	w.Write2b(v.MaxInstalledMaps)
	n += 2

	// write field RootVisual
	w.Write4b(uint32(v.RootVisual))
	n += 4

	// write field BackingStores
	w.Write1b(v.BackingStores)
	n += 1

	// write field SaveUnders
	w.Write1b(v.SaveUnders)
	n += 1

	// write field RootDepth
	w.Write1b(v.RootDepth)
	n += 1

	// write field AllowedDepthsLen
	w.Write1b(v.AllowedDepthsLen)
	n += 1

	// write field AllowedDepths
	return
}
func ScreenWriteN(w *Writer, src []Screen) (n int) {
	for i := 0; i < len(src); i++ {
		n += ScreenWrite(w, &src[i])
	}
	return
}

type SetupRequest struct {
	ByteOrder uint8
	// Pad0 uint8
	ProtocolMajorVersion         uint16
	ProtocolMinorVersion         uint16
	AuthorizationProtocolNameLen uint16
	AuthorizationProtocolDataLen uint16
	// Pad1 [2]uint8
	AuthorizationProtocolName string
	AuthorizationProtocolData string
}

type CSetupRequest struct {
	ByteOrder                    uint8
	Pad0                         uint8
	ProtocolMajorVersion         uint16
	ProtocolMinorVersion         uint16
	AuthorizationProtocolNameLen uint16
	AuthorizationProtocolDataLen uint16
	Pad1                         [2]uint8
}

func SetupRequestRead(r *Reader, v *SetupRequest) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ByteOrder
	v.ByteOrder = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field ProtocolMajorVersion
	v.ProtocolMajorVersion = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field ProtocolMinorVersion
	v.ProtocolMinorVersion = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field AuthorizationProtocolNameLen
	v.AuthorizationProtocolNameLen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field AuthorizationProtocolDataLen
	v.AuthorizationProtocolDataLen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(2)
	n += 2
	if r.Err() != nil {
		return
	}

	// read field AuthorizationProtocolName
	{
		dataLen := int(int(v.AuthorizationProtocolNameLen))
		data := r.ReadString(dataLen)
		if r.Err() != nil {
			return
		}
		blockLen += dataLen
		n += blockLen
		v.AuthorizationProtocolName = data
	}
	alignTo = int(unsafe.Alignof(byte(0)))

	// read field AuthorizationProtocolData

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	{
		dataLen := int(int(v.AuthorizationProtocolDataLen))
		data := r.ReadString(dataLen)
		if r.Err() != nil {
			return
		}
		blockLen += dataLen
		n += blockLen
		v.AuthorizationProtocolData = data
	}
	alignTo = int(unsafe.Alignof(byte(0)))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

func SetupRequestReadN(r *Reader, dest []SetupRequest) (n int) {
	for i := 0; i < len(dest); i++ {
		n += SetupRequestRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func SetupRequestWrite(w *Writer, v *SetupRequest) (n int) {

	// write field ByteOrder
	w.Write1b(v.ByteOrder)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field ProtocolMajorVersion
	w.Write2b(v.ProtocolMajorVersion)
	n += 2

	// write field ProtocolMinorVersion
	w.Write2b(v.ProtocolMinorVersion)
	n += 2

	// write field AuthorizationProtocolNameLen
	w.Write2b(v.AuthorizationProtocolNameLen)
	n += 2

	// write field AuthorizationProtocolDataLen
	w.Write2b(v.AuthorizationProtocolDataLen)
	n += 2

	// write field Pad1
	w.WritePad(2)
	n += 2

	// write field AuthorizationProtocolName

	// write field AuthorizationProtocolData
	return
}
func SetupRequestWriteN(w *Writer, src []SetupRequest) (n int) {
	for i := 0; i < len(src); i++ {
		n += SetupRequestWrite(w, &src[i])
	}
	return
}

type SetupFailed struct {
	Status               uint8
	ReasonLen            uint8
	ProtocolMajorVersion uint16
	ProtocolMinorVersion uint16
	Length               uint16
	Reason               string
}

type CSetupFailed struct {
	Status               uint8
	ReasonLen            uint8
	ProtocolMajorVersion uint16
	ProtocolMinorVersion uint16
	Length               uint16
}

func SetupFailedRead(r *Reader, v *SetupFailed) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Status
	v.Status = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field ReasonLen
	v.ReasonLen = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field ProtocolMajorVersion
	v.ProtocolMajorVersion = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field ProtocolMinorVersion
	v.ProtocolMinorVersion = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Reason
	{
		dataLen := int(int(v.ReasonLen))
		data := r.ReadString(dataLen)
		if r.Err() != nil {
			return
		}
		blockLen += dataLen
		n += blockLen
		v.Reason = data
	}
	alignTo = int(unsafe.Alignof(byte(0)))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

func SetupFailedReadN(r *Reader, dest []SetupFailed) (n int) {
	for i := 0; i < len(dest); i++ {
		n += SetupFailedRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func SetupFailedWrite(w *Writer, v *SetupFailed) (n int) {

	// write field Status
	w.Write1b(v.Status)
	n += 1

	// write field ReasonLen
	w.Write1b(v.ReasonLen)
	n += 1

	// write field ProtocolMajorVersion
	w.Write2b(v.ProtocolMajorVersion)
	n += 2

	// write field ProtocolMinorVersion
	w.Write2b(v.ProtocolMinorVersion)
	n += 2

	// write field Length
	w.Write2b(v.Length)
	n += 2

	// write field Reason
	return
}
func SetupFailedWriteN(w *Writer, src []SetupFailed) (n int) {
	for i := 0; i < len(src); i++ {
		n += SetupFailedWrite(w, &src[i])
	}
	return
}

type SetupAuthenticate struct {
	Status uint8
	// Pad0 [5]uint8
	Length uint16
	Reason string
}

type CSetupAuthenticate struct {
	Status uint8
	Pad0   [5]uint8
	Length uint16
}

func SetupAuthenticateRead(r *Reader, v *SetupAuthenticate) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Status
	v.Status = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(5)
	n += 5
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Reason
	{
		dataLen := int((int(v.Length) * 4))
		data := r.ReadString(dataLen)
		if r.Err() != nil {
			return
		}
		blockLen += dataLen
		n += blockLen
		v.Reason = data
	}
	alignTo = int(unsafe.Alignof(byte(0)))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

func SetupAuthenticateReadN(r *Reader, dest []SetupAuthenticate) (n int) {
	for i := 0; i < len(dest); i++ {
		n += SetupAuthenticateRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func SetupAuthenticateWrite(w *Writer, v *SetupAuthenticate) (n int) {

	// write field Status
	w.Write1b(v.Status)
	n += 1

	// write field Pad0
	w.WritePad(5)
	n += 5

	// write field Length
	w.Write2b(v.Length)
	n += 2

	// write field Reason
	return
}
func SetupAuthenticateWriteN(w *Writer, src []SetupAuthenticate) (n int) {
	for i := 0; i < len(src); i++ {
		n += SetupAuthenticateWrite(w, &src[i])
	}
	return
}

// enum ImageOrder
const (
	ImageOrderLSBFirst = 0
	ImageOrderMSBFirst = 1
)

func ImageOrderEnumToStr(v int) string {
	switch v {
	case 0:
		return "LSBFirst"
	case 1:
		return "MSBFirst"
	default:
		return "<Unknown>"
	}
}

type Setup struct {
	Status uint8
	// Pad0 uint8
	ProtocolMajorVersion     uint16
	ProtocolMinorVersion     uint16
	Length                   uint16
	ReleaseNumber            uint32
	ResourceIdBase           uint32
	ResourceIdMask           uint32
	MotionBufferSize         uint32
	VendorLen                uint16
	MaximumRequestLength     uint16
	RootsLen                 uint8
	PixmapFormatsLen         uint8
	ImageByteOrder           uint8
	BitmapFormatBitOrder     uint8
	BitmapFormatScanlineUnit uint8
	BitmapFormatScanlinePad  uint8
	MinKeycode               Keycode
	MaxKeycode               Keycode
	// Pad1 [4]uint8
	Vendor        string
	PixmapFormats []Format
	Roots         []Screen
}

type CSetup struct {
	Status                   uint8
	Pad0                     uint8
	ProtocolMajorVersion     uint16
	ProtocolMinorVersion     uint16
	Length                   uint16
	ReleaseNumber            uint32
	ResourceIdBase           uint32
	ResourceIdMask           uint32
	MotionBufferSize         uint32
	VendorLen                uint16
	MaximumRequestLength     uint16
	RootsLen                 uint8
	PixmapFormatsLen         uint8
	ImageByteOrder           uint8
	BitmapFormatBitOrder     uint8
	BitmapFormatScanlineUnit uint8
	BitmapFormatScanlinePad  uint8
	MinKeycode               Keycode
	MaxKeycode               Keycode
	Pad1                     [4]uint8
}

func SetupRead(r *Reader, v *Setup) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Status
	v.Status = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field ProtocolMajorVersion
	v.ProtocolMajorVersion = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field ProtocolMinorVersion
	v.ProtocolMinorVersion = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field ReleaseNumber
	v.ReleaseNumber = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field ResourceIdBase
	v.ResourceIdBase = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field ResourceIdMask
	v.ResourceIdMask = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field MotionBufferSize
	v.MotionBufferSize = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field VendorLen
	v.VendorLen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field MaximumRequestLength
	v.MaximumRequestLength = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field RootsLen
	v.RootsLen = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field PixmapFormatsLen
	v.PixmapFormatsLen = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field ImageByteOrder
	v.ImageByteOrder = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field BitmapFormatBitOrder
	v.BitmapFormatBitOrder = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field BitmapFormatScanlineUnit
	v.BitmapFormatScanlineUnit = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field BitmapFormatScanlinePad
	v.BitmapFormatScanlinePad = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field MinKeycode
	v.MinKeycode = Keycode(r.Read1b())
	n += 1
	if r.Err() != nil {
		return
	}

	// read field MaxKeycode
	v.MaxKeycode = Keycode(r.Read1b())
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(4)
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Vendor
	{
		dataLen := int(int(v.VendorLen))
		data := r.ReadString(dataLen)
		if r.Err() != nil {
			return
		}
		blockLen += dataLen
		n += blockLen
		v.Vendor = data
	}
	alignTo = int(unsafe.Alignof(byte(0)))

	// read field Pad2
	alignTo = 4

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */

	// read field PixmapFormats

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	v.PixmapFormats = make([]Format, int(v.PixmapFormatsLen))
	blockLen += FormatReadN(r, v.PixmapFormats)
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CFormat{}))

	// read field Roots

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	v.Roots = make([]Screen, int(v.RootsLen))
	blockLen += ScreenReadN(r, v.Roots)
	n += blockLen
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CScreen{}))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

func SetupReadN(r *Reader, dest []Setup) (n int) {
	for i := 0; i < len(dest); i++ {
		n += SetupRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func SetupWrite(w *Writer, v *Setup) (n int) {

	// write field Status
	w.Write1b(v.Status)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field ProtocolMajorVersion
	w.Write2b(v.ProtocolMajorVersion)
	n += 2

	// write field ProtocolMinorVersion
	w.Write2b(v.ProtocolMinorVersion)
	n += 2

	// write field Length
	w.Write2b(v.Length)
	n += 2

	// write field ReleaseNumber
	w.Write4b(v.ReleaseNumber)
	n += 4

	// write field ResourceIdBase
	w.Write4b(v.ResourceIdBase)
	n += 4

	// write field ResourceIdMask
	w.Write4b(v.ResourceIdMask)
	n += 4

	// write field MotionBufferSize
	w.Write4b(v.MotionBufferSize)
	n += 4

	// write field VendorLen
	w.Write2b(v.VendorLen)
	n += 2

	// write field MaximumRequestLength
	w.Write2b(v.MaximumRequestLength)
	n += 2

	// write field RootsLen
	w.Write1b(v.RootsLen)
	n += 1

	// write field PixmapFormatsLen
	w.Write1b(v.PixmapFormatsLen)
	n += 1

	// write field ImageByteOrder
	w.Write1b(v.ImageByteOrder)
	n += 1

	// write field BitmapFormatBitOrder
	w.Write1b(v.BitmapFormatBitOrder)
	n += 1

	// write field BitmapFormatScanlineUnit
	w.Write1b(v.BitmapFormatScanlineUnit)
	n += 1

	// write field BitmapFormatScanlinePad
	w.Write1b(v.BitmapFormatScanlinePad)
	n += 1

	// write field MinKeycode
	w.Write1b(uint8(v.MinKeycode))
	n += 1

	// write field MaxKeycode
	w.Write1b(uint8(v.MaxKeycode))
	n += 1

	// write field Pad1
	w.WritePad(4)
	n += 4

	// write field Vendor

	// write field Pad2

	// write field PixmapFormats

	// write field Roots
	return
}
func SetupWriteN(w *Writer, src []Setup) (n int) {
	for i := 0; i < len(src); i++ {
		n += SetupWrite(w, &src[i])
	}
	return
}

// enum ModMask
const (
	ModMaskShift   = 1
	ModMaskLock    = 2
	ModMaskControl = 4
	ModMask1       = 8
	ModMask2       = 16
	ModMask3       = 32
	ModMask4       = 64
	ModMask5       = 128
	ModMaskAny     = 32768
)

func ModMaskEnumToStr(v int) string {
	strv := []string{}
	switch {
	case v&ModMaskShift > 0:
		strv = append(strv, "Shift")
	case v&ModMaskLock > 0:
		strv = append(strv, "Lock")
	case v&ModMaskControl > 0:
		strv = append(strv, "Control")
	case v&ModMask1 > 0:
		strv = append(strv, "1")
	case v&ModMask2 > 0:
		strv = append(strv, "2")
	case v&ModMask3 > 0:
		strv = append(strv, "3")
	case v&ModMask4 > 0:
		strv = append(strv, "4")
	case v&ModMask5 > 0:
		strv = append(strv, "5")
	case v&ModMaskAny > 0:
		strv = append(strv, "Any")
	}
	return "[" + strings.Join(strv, "|") + "]"
}

// enum KeyButMask
const (
	KeyButMaskShift   = 1
	KeyButMaskLock    = 2
	KeyButMaskControl = 4
	KeyButMaskMod1    = 8
	KeyButMaskMod2    = 16
	KeyButMaskMod3    = 32
	KeyButMaskMod4    = 64
	KeyButMaskMod5    = 128
	KeyButMaskButton1 = 256
	KeyButMaskButton2 = 512
	KeyButMaskButton3 = 1024
	KeyButMaskButton4 = 2048
	KeyButMaskButton5 = 4096
)

func KeyButMaskEnumToStr(v int) string {
	strv := []string{}
	switch {
	case v&KeyButMaskShift > 0:
		strv = append(strv, "Shift")
	case v&KeyButMaskLock > 0:
		strv = append(strv, "Lock")
	case v&KeyButMaskControl > 0:
		strv = append(strv, "Control")
	case v&KeyButMaskMod1 > 0:
		strv = append(strv, "Mod1")
	case v&KeyButMaskMod2 > 0:
		strv = append(strv, "Mod2")
	case v&KeyButMaskMod3 > 0:
		strv = append(strv, "Mod3")
	case v&KeyButMaskMod4 > 0:
		strv = append(strv, "Mod4")
	case v&KeyButMaskMod5 > 0:
		strv = append(strv, "Mod5")
	case v&KeyButMaskButton1 > 0:
		strv = append(strv, "Button1")
	case v&KeyButMaskButton2 > 0:
		strv = append(strv, "Button2")
	case v&KeyButMaskButton3 > 0:
		strv = append(strv, "Button3")
	case v&KeyButMaskButton4 > 0:
		strv = append(strv, "Button4")
	case v&KeyButMaskButton5 > 0:
		strv = append(strv, "Button5")
	}
	return "[" + strings.Join(strv, "|") + "]"
}

// enum Window
const (
	WindowNone = 0
)

// event name: ('xcb', 'KeyPress')
// self.name: ('xcb', 'KeyPress')
const KeyPressEventCode = 2

// not event copy
type KeyPressEvent struct {
	ResponseType uint8
	Detail       Keycode
	Sequence     uint16
	Time         Timestamp
	Root         Window
	Event        Window
	Child        Window
	RootX        int16
	RootY        int16
	EventX       int16
	EventY       int16
	State        uint16
	SameScreen   uint8
	// Pad0 uint8
}

func KeyPressEventRead(r *Reader, v *KeyPressEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Detail
	v.Detail = Keycode(r.Read1b())
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Time
	v.Time = Timestamp(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Root
	v.Root = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Event
	v.Event = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Child
	v.Child = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field RootX
	v.RootX = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field RootY
	v.RootY = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field EventX
	v.EventX = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field EventY
	v.EventY = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field State
	v.State = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field SameScreen
	v.SameScreen = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}
	return
}
func KeyPressEventWrite(w *Writer, v *KeyPressEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Detail
	w.Write1b(uint8(v.Detail))
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Time
	w.Write4b(uint32(v.Time))
	n += 4

	// write field Root
	w.Write4b(uint32(v.Root))
	n += 4

	// write field Event
	w.Write4b(uint32(v.Event))
	n += 4

	// write field Child
	w.Write4b(uint32(v.Child))
	n += 4

	// write field RootX
	w.Write2b(uint16(v.RootX))
	n += 2

	// write field RootY
	w.Write2b(uint16(v.RootY))
	n += 2

	// write field EventX
	w.Write2b(uint16(v.EventX))
	n += 2

	// write field EventY
	w.Write2b(uint16(v.EventY))
	n += 2

	// write field State
	w.Write2b(v.State)
	n += 2

	// write field SameScreen
	w.Write1b(v.SameScreen)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1
	return
}

func NewKeyPressEvent(data []byte) (*KeyPressEvent, error) {
	var ev KeyPressEvent
	r := NewReaderFromData(data)
	KeyPressEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *KeyPressEvent) String() string {
	return fmt.Sprintf("KeyPressEvent {Detail: %v, Sequence: %v, Time: %v, Root: %v, Event: %v, Child: %v, RootX: %v, RootY: %v, EventX: %v, EventY: %v, State: %v, SameScreen: %v}",
		v.Detail, v.Sequence, v.Time, v.Root, v.Event, v.Child, v.RootX, v.RootY, v.EventX, v.EventY, v.State, v.SameScreen)
}

// _go_event_handler_with_xid KeyPressEvent event
type _KeyPressEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*KeyPressEvent)
}

func (eh *_KeyPressEventHandler) run(xid uint32, ev *KeyPressEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_KeyPressEventHandler) Run(ge GenericEvent) {
	ev, err := NewKeyPressEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Event), ev)
}

func (eh *_KeyPressEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_KeyPressEventHandler) attach(xid uint32, cb func(ev *KeyPressEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*KeyPressEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectKeyPressEvent(conn *Conn, xid Window, cb func(ev *KeyPressEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[KeyPressEventCode]
	if !ok {
		newEh := &_KeyPressEventHandler{
			cbs: make(map[uint32][]func(*KeyPressEvent)),
		}
		conn.EventHandlers[KeyPressEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_KeyPressEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'KeyRelease')
// self.name: ('xcb', 'KeyPress')
const KeyReleaseEventCode = 3

// is event copy
type KeyReleaseEvent KeyPressEvent

func KeyReleaseEventRead(r *Reader, v *KeyReleaseEvent) (n int) {
	// source_type_name: KeyPressEvent
	return KeyPressEventRead(r, (*KeyPressEvent)(v))
}

func NewKeyReleaseEvent(data []byte) (*KeyReleaseEvent, error) {
	var ev KeyReleaseEvent
	r := NewReaderFromData(data)
	KeyReleaseEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *KeyReleaseEvent) String() string {
	return fmt.Sprintf("KeyReleaseEvent {Detail: %v, Sequence: %v, Time: %v, Root: %v, Event: %v, Child: %v, RootX: %v, RootY: %v, EventX: %v, EventY: %v, State: %v, SameScreen: %v}",
		v.Detail, v.Sequence, v.Time, v.Root, v.Event, v.Child, v.RootX, v.RootY, v.EventX, v.EventY, v.State, v.SameScreen)
}

// try event source KeyPress
// _go_event_handler_with_xid KeyReleaseEvent event
type _KeyReleaseEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*KeyReleaseEvent)
}

func (eh *_KeyReleaseEventHandler) run(xid uint32, ev *KeyReleaseEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_KeyReleaseEventHandler) Run(ge GenericEvent) {
	ev, err := NewKeyReleaseEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Event), ev)
}

func (eh *_KeyReleaseEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_KeyReleaseEventHandler) attach(xid uint32, cb func(ev *KeyReleaseEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*KeyReleaseEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectKeyReleaseEvent(conn *Conn, xid Window, cb func(ev *KeyReleaseEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[KeyReleaseEventCode]
	if !ok {
		newEh := &_KeyReleaseEventHandler{
			cbs: make(map[uint32][]func(*KeyReleaseEvent)),
		}
		conn.EventHandlers[KeyReleaseEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_KeyReleaseEventHandler).attach(uint32(xid), cb)
}

// enum ButtonMask
const (
	ButtonMask1   = 256
	ButtonMask2   = 512
	ButtonMask3   = 1024
	ButtonMask4   = 2048
	ButtonMask5   = 4096
	ButtonMaskAny = 32768
)

func ButtonMaskEnumToStr(v int) string {
	strv := []string{}
	switch {
	case v&ButtonMask1 > 0:
		strv = append(strv, "1")
	case v&ButtonMask2 > 0:
		strv = append(strv, "2")
	case v&ButtonMask3 > 0:
		strv = append(strv, "3")
	case v&ButtonMask4 > 0:
		strv = append(strv, "4")
	case v&ButtonMask5 > 0:
		strv = append(strv, "5")
	case v&ButtonMaskAny > 0:
		strv = append(strv, "Any")
	}
	return "[" + strings.Join(strv, "|") + "]"
}

// event name: ('xcb', 'ButtonPress')
// self.name: ('xcb', 'ButtonPress')
const ButtonPressEventCode = 4

// not event copy
type ButtonPressEvent struct {
	ResponseType uint8
	Detail       Button
	Sequence     uint16
	Time         Timestamp
	Root         Window
	Event        Window
	Child        Window
	RootX        int16
	RootY        int16
	EventX       int16
	EventY       int16
	State        uint16
	SameScreen   uint8
	// Pad0 uint8
}

func ButtonPressEventRead(r *Reader, v *ButtonPressEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Detail
	v.Detail = Button(r.Read1b())
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Time
	v.Time = Timestamp(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Root
	v.Root = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Event
	v.Event = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Child
	v.Child = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field RootX
	v.RootX = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field RootY
	v.RootY = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field EventX
	v.EventX = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field EventY
	v.EventY = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field State
	v.State = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field SameScreen
	v.SameScreen = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}
	return
}
func ButtonPressEventWrite(w *Writer, v *ButtonPressEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Detail
	w.Write1b(uint8(v.Detail))
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Time
	w.Write4b(uint32(v.Time))
	n += 4

	// write field Root
	w.Write4b(uint32(v.Root))
	n += 4

	// write field Event
	w.Write4b(uint32(v.Event))
	n += 4

	// write field Child
	w.Write4b(uint32(v.Child))
	n += 4

	// write field RootX
	w.Write2b(uint16(v.RootX))
	n += 2

	// write field RootY
	w.Write2b(uint16(v.RootY))
	n += 2

	// write field EventX
	w.Write2b(uint16(v.EventX))
	n += 2

	// write field EventY
	w.Write2b(uint16(v.EventY))
	n += 2

	// write field State
	w.Write2b(v.State)
	n += 2

	// write field SameScreen
	w.Write1b(v.SameScreen)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1
	return
}

func NewButtonPressEvent(data []byte) (*ButtonPressEvent, error) {
	var ev ButtonPressEvent
	r := NewReaderFromData(data)
	ButtonPressEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *ButtonPressEvent) String() string {
	return fmt.Sprintf("ButtonPressEvent {Detail: %v, Sequence: %v, Time: %v, Root: %v, Event: %v, Child: %v, RootX: %v, RootY: %v, EventX: %v, EventY: %v, State: %v, SameScreen: %v}",
		v.Detail, v.Sequence, v.Time, v.Root, v.Event, v.Child, v.RootX, v.RootY, v.EventX, v.EventY, v.State, v.SameScreen)
}

// _go_event_handler_with_xid ButtonPressEvent event
type _ButtonPressEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*ButtonPressEvent)
}

func (eh *_ButtonPressEventHandler) run(xid uint32, ev *ButtonPressEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_ButtonPressEventHandler) Run(ge GenericEvent) {
	ev, err := NewButtonPressEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Event), ev)
}

func (eh *_ButtonPressEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_ButtonPressEventHandler) attach(xid uint32, cb func(ev *ButtonPressEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*ButtonPressEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectButtonPressEvent(conn *Conn, xid Window, cb func(ev *ButtonPressEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[ButtonPressEventCode]
	if !ok {
		newEh := &_ButtonPressEventHandler{
			cbs: make(map[uint32][]func(*ButtonPressEvent)),
		}
		conn.EventHandlers[ButtonPressEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_ButtonPressEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'ButtonRelease')
// self.name: ('xcb', 'ButtonPress')
const ButtonReleaseEventCode = 5

// is event copy
type ButtonReleaseEvent ButtonPressEvent

func ButtonReleaseEventRead(r *Reader, v *ButtonReleaseEvent) (n int) {
	// source_type_name: ButtonPressEvent
	return ButtonPressEventRead(r, (*ButtonPressEvent)(v))
}

func NewButtonReleaseEvent(data []byte) (*ButtonReleaseEvent, error) {
	var ev ButtonReleaseEvent
	r := NewReaderFromData(data)
	ButtonReleaseEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *ButtonReleaseEvent) String() string {
	return fmt.Sprintf("ButtonReleaseEvent {Detail: %v, Sequence: %v, Time: %v, Root: %v, Event: %v, Child: %v, RootX: %v, RootY: %v, EventX: %v, EventY: %v, State: %v, SameScreen: %v}",
		v.Detail, v.Sequence, v.Time, v.Root, v.Event, v.Child, v.RootX, v.RootY, v.EventX, v.EventY, v.State, v.SameScreen)
}

// try event source ButtonPress
// _go_event_handler_with_xid ButtonReleaseEvent event
type _ButtonReleaseEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*ButtonReleaseEvent)
}

func (eh *_ButtonReleaseEventHandler) run(xid uint32, ev *ButtonReleaseEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_ButtonReleaseEventHandler) Run(ge GenericEvent) {
	ev, err := NewButtonReleaseEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Event), ev)
}

func (eh *_ButtonReleaseEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_ButtonReleaseEventHandler) attach(xid uint32, cb func(ev *ButtonReleaseEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*ButtonReleaseEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectButtonReleaseEvent(conn *Conn, xid Window, cb func(ev *ButtonReleaseEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[ButtonReleaseEventCode]
	if !ok {
		newEh := &_ButtonReleaseEventHandler{
			cbs: make(map[uint32][]func(*ButtonReleaseEvent)),
		}
		conn.EventHandlers[ButtonReleaseEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_ButtonReleaseEventHandler).attach(uint32(xid), cb)
}

// enum Motion
const (
	MotionNormal = 0
	MotionHint   = 1
)

func MotionEnumToStr(v int) string {
	switch v {
	case 0:
		return "Normal"
	case 1:
		return "Hint"
	default:
		return "<Unknown>"
	}
}

// event name: ('xcb', 'MotionNotify')
// self.name: ('xcb', 'MotionNotify')
const MotionNotifyEventCode = 6

// not event copy
type MotionNotifyEvent struct {
	ResponseType uint8
	Detail       uint8
	Sequence     uint16
	Time         Timestamp
	Root         Window
	Event        Window
	Child        Window
	RootX        int16
	RootY        int16
	EventX       int16
	EventY       int16
	State        uint16
	SameScreen   uint8
	// Pad0 uint8
}

func MotionNotifyEventRead(r *Reader, v *MotionNotifyEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Detail
	v.Detail = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Time
	v.Time = Timestamp(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Root
	v.Root = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Event
	v.Event = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Child
	v.Child = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field RootX
	v.RootX = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field RootY
	v.RootY = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field EventX
	v.EventX = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field EventY
	v.EventY = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field State
	v.State = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field SameScreen
	v.SameScreen = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}
	return
}
func MotionNotifyEventWrite(w *Writer, v *MotionNotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Detail
	w.Write1b(v.Detail)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Time
	w.Write4b(uint32(v.Time))
	n += 4

	// write field Root
	w.Write4b(uint32(v.Root))
	n += 4

	// write field Event
	w.Write4b(uint32(v.Event))
	n += 4

	// write field Child
	w.Write4b(uint32(v.Child))
	n += 4

	// write field RootX
	w.Write2b(uint16(v.RootX))
	n += 2

	// write field RootY
	w.Write2b(uint16(v.RootY))
	n += 2

	// write field EventX
	w.Write2b(uint16(v.EventX))
	n += 2

	// write field EventY
	w.Write2b(uint16(v.EventY))
	n += 2

	// write field State
	w.Write2b(v.State)
	n += 2

	// write field SameScreen
	w.Write1b(v.SameScreen)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1
	return
}

func NewMotionNotifyEvent(data []byte) (*MotionNotifyEvent, error) {
	var ev MotionNotifyEvent
	r := NewReaderFromData(data)
	MotionNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *MotionNotifyEvent) String() string {
	return fmt.Sprintf("MotionNotifyEvent {Detail: %v, Sequence: %v, Time: %v, Root: %v, Event: %v, Child: %v, RootX: %v, RootY: %v, EventX: %v, EventY: %v, State: %v, SameScreen: %v}",
		v.Detail, v.Sequence, v.Time, v.Root, v.Event, v.Child, v.RootX, v.RootY, v.EventX, v.EventY, v.State, v.SameScreen)
}

// _go_event_handler_with_xid MotionNotifyEvent event
type _MotionNotifyEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*MotionNotifyEvent)
}

func (eh *_MotionNotifyEventHandler) run(xid uint32, ev *MotionNotifyEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_MotionNotifyEventHandler) Run(ge GenericEvent) {
	ev, err := NewMotionNotifyEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Event), ev)
}

func (eh *_MotionNotifyEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_MotionNotifyEventHandler) attach(xid uint32, cb func(ev *MotionNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*MotionNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectMotionNotifyEvent(conn *Conn, xid Window, cb func(ev *MotionNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[MotionNotifyEventCode]
	if !ok {
		newEh := &_MotionNotifyEventHandler{
			cbs: make(map[uint32][]func(*MotionNotifyEvent)),
		}
		conn.EventHandlers[MotionNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_MotionNotifyEventHandler).attach(uint32(xid), cb)
}

// enum NotifyDetail
const (
	NotifyDetailAncestor         = 0
	NotifyDetailVirtual          = 1
	NotifyDetailInferior         = 2
	NotifyDetailNonlinear        = 3
	NotifyDetailNonlinearVirtual = 4
	NotifyDetailPointer          = 5
	NotifyDetailPointerRoot      = 6
	NotifyDetailNone             = 7
)

func NotifyDetailEnumToStr(v int) string {
	switch v {
	case 0:
		return "Ancestor"
	case 1:
		return "Virtual"
	case 2:
		return "Inferior"
	case 3:
		return "Nonlinear"
	case 4:
		return "NonlinearVirtual"
	case 5:
		return "Pointer"
	case 6:
		return "PointerRoot"
	case 7:
		return "None"
	default:
		return "<Unknown>"
	}
}

// enum NotifyMode
const (
	NotifyModeNormal       = 0
	NotifyModeGrab         = 1
	NotifyModeUngrab       = 2
	NotifyModeWhileGrabbed = 3
)

func NotifyModeEnumToStr(v int) string {
	switch v {
	case 0:
		return "Normal"
	case 1:
		return "Grab"
	case 2:
		return "Ungrab"
	case 3:
		return "WhileGrabbed"
	default:
		return "<Unknown>"
	}
}

// event name: ('xcb', 'EnterNotify')
// self.name: ('xcb', 'EnterNotify')
const EnterNotifyEventCode = 7

// not event copy
type EnterNotifyEvent struct {
	ResponseType    uint8
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

func EnterNotifyEventRead(r *Reader, v *EnterNotifyEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Detail
	v.Detail = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Time
	v.Time = Timestamp(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Root
	v.Root = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Event
	v.Event = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Child
	v.Child = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field RootX
	v.RootX = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field RootY
	v.RootY = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field EventX
	v.EventX = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field EventY
	v.EventY = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field State
	v.State = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Mode
	v.Mode = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field SameScreenFocus
	v.SameScreenFocus = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}
	return
}
func EnterNotifyEventWrite(w *Writer, v *EnterNotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Detail
	w.Write1b(v.Detail)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Time
	w.Write4b(uint32(v.Time))
	n += 4

	// write field Root
	w.Write4b(uint32(v.Root))
	n += 4

	// write field Event
	w.Write4b(uint32(v.Event))
	n += 4

	// write field Child
	w.Write4b(uint32(v.Child))
	n += 4

	// write field RootX
	w.Write2b(uint16(v.RootX))
	n += 2

	// write field RootY
	w.Write2b(uint16(v.RootY))
	n += 2

	// write field EventX
	w.Write2b(uint16(v.EventX))
	n += 2

	// write field EventY
	w.Write2b(uint16(v.EventY))
	n += 2

	// write field State
	w.Write2b(v.State)
	n += 2

	// write field Mode
	w.Write1b(v.Mode)
	n += 1

	// write field SameScreenFocus
	w.Write1b(v.SameScreenFocus)
	n += 1
	return
}

func NewEnterNotifyEvent(data []byte) (*EnterNotifyEvent, error) {
	var ev EnterNotifyEvent
	r := NewReaderFromData(data)
	EnterNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *EnterNotifyEvent) String() string {
	return fmt.Sprintf("EnterNotifyEvent {Detail: %v, Sequence: %v, Time: %v, Root: %v, Event: %v, Child: %v, RootX: %v, RootY: %v, EventX: %v, EventY: %v, State: %v, Mode: %v, SameScreenFocus: %v}",
		v.Detail, v.Sequence, v.Time, v.Root, v.Event, v.Child, v.RootX, v.RootY, v.EventX, v.EventY, v.State, v.Mode, v.SameScreenFocus)
}

// _go_event_handler_with_xid EnterNotifyEvent event
type _EnterNotifyEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*EnterNotifyEvent)
}

func (eh *_EnterNotifyEventHandler) run(xid uint32, ev *EnterNotifyEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_EnterNotifyEventHandler) Run(ge GenericEvent) {
	ev, err := NewEnterNotifyEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Event), ev)
}

func (eh *_EnterNotifyEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_EnterNotifyEventHandler) attach(xid uint32, cb func(ev *EnterNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*EnterNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectEnterNotifyEvent(conn *Conn, xid Window, cb func(ev *EnterNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[EnterNotifyEventCode]
	if !ok {
		newEh := &_EnterNotifyEventHandler{
			cbs: make(map[uint32][]func(*EnterNotifyEvent)),
		}
		conn.EventHandlers[EnterNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_EnterNotifyEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'LeaveNotify')
// self.name: ('xcb', 'EnterNotify')
const LeaveNotifyEventCode = 8

// is event copy
type LeaveNotifyEvent EnterNotifyEvent

func LeaveNotifyEventRead(r *Reader, v *LeaveNotifyEvent) (n int) {
	// source_type_name: EnterNotifyEvent
	return EnterNotifyEventRead(r, (*EnterNotifyEvent)(v))
}

func NewLeaveNotifyEvent(data []byte) (*LeaveNotifyEvent, error) {
	var ev LeaveNotifyEvent
	r := NewReaderFromData(data)
	LeaveNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *LeaveNotifyEvent) String() string {
	return fmt.Sprintf("LeaveNotifyEvent {Detail: %v, Sequence: %v, Time: %v, Root: %v, Event: %v, Child: %v, RootX: %v, RootY: %v, EventX: %v, EventY: %v, State: %v, Mode: %v, SameScreenFocus: %v}",
		v.Detail, v.Sequence, v.Time, v.Root, v.Event, v.Child, v.RootX, v.RootY, v.EventX, v.EventY, v.State, v.Mode, v.SameScreenFocus)
}

// try event source EnterNotify
// _go_event_handler_with_xid LeaveNotifyEvent event
type _LeaveNotifyEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*LeaveNotifyEvent)
}

func (eh *_LeaveNotifyEventHandler) run(xid uint32, ev *LeaveNotifyEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_LeaveNotifyEventHandler) Run(ge GenericEvent) {
	ev, err := NewLeaveNotifyEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Event), ev)
}

func (eh *_LeaveNotifyEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_LeaveNotifyEventHandler) attach(xid uint32, cb func(ev *LeaveNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*LeaveNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectLeaveNotifyEvent(conn *Conn, xid Window, cb func(ev *LeaveNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[LeaveNotifyEventCode]
	if !ok {
		newEh := &_LeaveNotifyEventHandler{
			cbs: make(map[uint32][]func(*LeaveNotifyEvent)),
		}
		conn.EventHandlers[LeaveNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_LeaveNotifyEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'FocusIn')
// self.name: ('xcb', 'FocusIn')
const FocusInEventCode = 9

// not event copy
type FocusInEvent struct {
	ResponseType uint8
	Detail       uint8
	Sequence     uint16
	Event        Window
	Mode         uint8
	// Pad0 [3]uint8
}

func FocusInEventRead(r *Reader, v *FocusInEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Detail
	v.Detail = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Event
	v.Event = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Mode
	v.Mode = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(3)
	n += 3
	if r.Err() != nil {
		return
	}
	return
}
func FocusInEventWrite(w *Writer, v *FocusInEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Detail
	w.Write1b(v.Detail)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Event
	w.Write4b(uint32(v.Event))
	n += 4

	// write field Mode
	w.Write1b(v.Mode)
	n += 1

	// write field Pad0
	w.WritePad(3)
	n += 3
	return
}

func NewFocusInEvent(data []byte) (*FocusInEvent, error) {
	var ev FocusInEvent
	r := NewReaderFromData(data)
	FocusInEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *FocusInEvent) String() string {
	return fmt.Sprintf("FocusInEvent {Detail: %v, Sequence: %v, Event: %v, Mode: %v}",
		v.Detail, v.Sequence, v.Event, v.Mode)
}

// _go_event_handler_with_xid FocusInEvent event
type _FocusInEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*FocusInEvent)
}

func (eh *_FocusInEventHandler) run(xid uint32, ev *FocusInEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_FocusInEventHandler) Run(ge GenericEvent) {
	ev, err := NewFocusInEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Event), ev)
}

func (eh *_FocusInEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_FocusInEventHandler) attach(xid uint32, cb func(ev *FocusInEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*FocusInEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectFocusInEvent(conn *Conn, xid Window, cb func(ev *FocusInEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[FocusInEventCode]
	if !ok {
		newEh := &_FocusInEventHandler{
			cbs: make(map[uint32][]func(*FocusInEvent)),
		}
		conn.EventHandlers[FocusInEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_FocusInEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'FocusOut')
// self.name: ('xcb', 'FocusIn')
const FocusOutEventCode = 10

// is event copy
type FocusOutEvent FocusInEvent

func FocusOutEventRead(r *Reader, v *FocusOutEvent) (n int) {
	// source_type_name: FocusInEvent
	return FocusInEventRead(r, (*FocusInEvent)(v))
}

func NewFocusOutEvent(data []byte) (*FocusOutEvent, error) {
	var ev FocusOutEvent
	r := NewReaderFromData(data)
	FocusOutEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *FocusOutEvent) String() string {
	return fmt.Sprintf("FocusOutEvent {Detail: %v, Sequence: %v, Event: %v, Mode: %v}",
		v.Detail, v.Sequence, v.Event, v.Mode)
}

// try event source FocusIn
// _go_event_handler_with_xid FocusOutEvent event
type _FocusOutEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*FocusOutEvent)
}

func (eh *_FocusOutEventHandler) run(xid uint32, ev *FocusOutEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_FocusOutEventHandler) Run(ge GenericEvent) {
	ev, err := NewFocusOutEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Event), ev)
}

func (eh *_FocusOutEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_FocusOutEventHandler) attach(xid uint32, cb func(ev *FocusOutEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*FocusOutEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectFocusOutEvent(conn *Conn, xid Window, cb func(ev *FocusOutEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[FocusOutEventCode]
	if !ok {
		newEh := &_FocusOutEventHandler{
			cbs: make(map[uint32][]func(*FocusOutEvent)),
		}
		conn.EventHandlers[FocusOutEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_FocusOutEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'KeymapNotify')
// self.name: ('xcb', 'KeymapNotify')
const KeymapNotifyEventCode = 11

// not event copy
type KeymapNotifyEvent struct {
	ResponseType uint8
	Keys         [31]uint8
}

func KeymapNotifyEventRead(r *Reader, v *KeymapNotifyEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Keys
	{
		dataLen := int(31)
		for i := 0; i < dataLen; i++ {
			v.Keys[i] = r.Read1b()
			if r.Err() != nil {
				return
			}
		}
		n += dataLen
	}
	return
}
func KeymapNotifyEventWrite(w *Writer, v *KeymapNotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Keys
	return
}

func NewKeymapNotifyEvent(data []byte) (*KeymapNotifyEvent, error) {
	var ev KeymapNotifyEvent
	r := NewReaderFromData(data)
	KeymapNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *KeymapNotifyEvent) String() string {
	return fmt.Sprintf("KeymapNotifyEvent {Keys: %v}",
		v.Keys)
}

// try event source KeymapNotify
type _KeymapNotifyEventHandler struct {
	mu  sync.Mutex
	cbs []func(*KeymapNotifyEvent)
}

func (eh *_KeymapNotifyEventHandler) Run(ge GenericEvent) {
	ev, err := NewKeymapNotifyEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.mu.Lock()
	cbs := eh.cbs
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (_ *_KeymapNotifyEventHandler) Detach(_ uint32) {}

func (eh *_KeymapNotifyEventHandler) attach(cb func(ev *KeymapNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs
	newCbs := make([]func(*KeymapNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs = newCbs
	eh.mu.Unlock()
}

func ConnectKeymapNotifyEvent(conn *Conn, cb func(ev *KeymapNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[KeymapNotifyEventCode]
	if !ok {
		newEh := &_KeymapNotifyEventHandler{}
		conn.EventHandlers[KeymapNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_KeymapNotifyEventHandler).attach(cb)
}

// event name: ('xcb', 'Expose')
// self.name: ('xcb', 'Expose')
const ExposeEventCode = 12

// not event copy
type ExposeEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Window   Window
	X        uint16
	Y        uint16
	Width    uint16
	Height   uint16
	Count    uint16
	// Pad1 [2]uint8
}

func ExposeEventRead(r *Reader, v *ExposeEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Window
	v.Window = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field X
	v.X = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Y
	v.Y = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Width
	v.Width = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Height
	v.Height = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Count
	v.Count = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(2)
	n += 2
	if r.Err() != nil {
		return
	}
	return
}
func ExposeEventWrite(w *Writer, v *ExposeEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4

	// write field X
	w.Write2b(v.X)
	n += 2

	// write field Y
	w.Write2b(v.Y)
	n += 2

	// write field Width
	w.Write2b(v.Width)
	n += 2

	// write field Height
	w.Write2b(v.Height)
	n += 2

	// write field Count
	w.Write2b(v.Count)
	n += 2

	// write field Pad1
	w.WritePad(2)
	n += 2
	return
}

func NewExposeEvent(data []byte) (*ExposeEvent, error) {
	var ev ExposeEvent
	r := NewReaderFromData(data)
	ExposeEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *ExposeEvent) String() string {
	return fmt.Sprintf("ExposeEvent {Sequence: %v, Window: %v, X: %v, Y: %v, Width: %v, Height: %v, Count: %v}",
		v.Sequence, v.Window, v.X, v.Y, v.Width, v.Height, v.Count)
}

// _go_event_handler_with_xid ExposeEvent window
type _ExposeEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*ExposeEvent)
}

func (eh *_ExposeEventHandler) run(xid uint32, ev *ExposeEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_ExposeEventHandler) Run(ge GenericEvent) {
	ev, err := NewExposeEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Window), ev)
}

func (eh *_ExposeEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_ExposeEventHandler) attach(xid uint32, cb func(ev *ExposeEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*ExposeEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectExposeEvent(conn *Conn, xid Window, cb func(ev *ExposeEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[ExposeEventCode]
	if !ok {
		newEh := &_ExposeEventHandler{
			cbs: make(map[uint32][]func(*ExposeEvent)),
		}
		conn.EventHandlers[ExposeEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_ExposeEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'GraphicsExposure')
// self.name: ('xcb', 'GraphicsExposure')
const GraphicsExposureEventCode = 13

// not event copy
type GraphicsExposureEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence    uint16
	Drawable    Drawable
	X           uint16
	Y           uint16
	Width       uint16
	Height      uint16
	MinorOpcode uint16
	Count       uint16
	MajorOpcode uint8
	// Pad1 [3]uint8
}

func GraphicsExposureEventRead(r *Reader, v *GraphicsExposureEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Drawable
	v.Drawable = Drawable(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field X
	v.X = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Y
	v.Y = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Width
	v.Width = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Height
	v.Height = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field MinorOpcode
	v.MinorOpcode = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Count
	v.Count = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field MajorOpcode
	v.MajorOpcode = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(3)
	n += 3
	if r.Err() != nil {
		return
	}
	return
}
func GraphicsExposureEventWrite(w *Writer, v *GraphicsExposureEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Drawable
	w.Write4b(uint32(v.Drawable))
	n += 4

	// write field X
	w.Write2b(v.X)
	n += 2

	// write field Y
	w.Write2b(v.Y)
	n += 2

	// write field Width
	w.Write2b(v.Width)
	n += 2

	// write field Height
	w.Write2b(v.Height)
	n += 2

	// write field MinorOpcode
	w.Write2b(v.MinorOpcode)
	n += 2

	// write field Count
	w.Write2b(v.Count)
	n += 2

	// write field MajorOpcode
	w.Write1b(v.MajorOpcode)
	n += 1

	// write field Pad1
	w.WritePad(3)
	n += 3
	return
}

func NewGraphicsExposureEvent(data []byte) (*GraphicsExposureEvent, error) {
	var ev GraphicsExposureEvent
	r := NewReaderFromData(data)
	GraphicsExposureEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *GraphicsExposureEvent) String() string {
	return fmt.Sprintf("GraphicsExposureEvent {Sequence: %v, Drawable: %v, X: %v, Y: %v, Width: %v, Height: %v, MinorOpcode: %v, Count: %v, MajorOpcode: %v}",
		v.Sequence, v.Drawable, v.X, v.Y, v.Width, v.Height, v.MinorOpcode, v.Count, v.MajorOpcode)
}

// _go_event_handler_with_xid GraphicsExposureEvent drawable
type _GraphicsExposureEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*GraphicsExposureEvent)
}

func (eh *_GraphicsExposureEventHandler) run(xid uint32, ev *GraphicsExposureEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_GraphicsExposureEventHandler) Run(ge GenericEvent) {
	ev, err := NewGraphicsExposureEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Drawable), ev)
}

func (eh *_GraphicsExposureEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_GraphicsExposureEventHandler) attach(xid uint32, cb func(ev *GraphicsExposureEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*GraphicsExposureEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectGraphicsExposureEvent(conn *Conn, xid Drawable, cb func(ev *GraphicsExposureEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[GraphicsExposureEventCode]
	if !ok {
		newEh := &_GraphicsExposureEventHandler{
			cbs: make(map[uint32][]func(*GraphicsExposureEvent)),
		}
		conn.EventHandlers[GraphicsExposureEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_GraphicsExposureEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'NoExposure')
// self.name: ('xcb', 'NoExposure')
const NoExposureEventCode = 14

// not event copy
type NoExposureEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence    uint16
	Drawable    Drawable
	MinorOpcode uint16
	MajorOpcode uint8
	// Pad1 uint8
}

func NoExposureEventRead(r *Reader, v *NoExposureEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Drawable
	v.Drawable = Drawable(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field MinorOpcode
	v.MinorOpcode = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field MajorOpcode
	v.MajorOpcode = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}
	return
}
func NoExposureEventWrite(w *Writer, v *NoExposureEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Drawable
	w.Write4b(uint32(v.Drawable))
	n += 4

	// write field MinorOpcode
	w.Write2b(v.MinorOpcode)
	n += 2

	// write field MajorOpcode
	w.Write1b(v.MajorOpcode)
	n += 1

	// write field Pad1
	w.WritePad(1)
	n += 1
	return
}

func NewNoExposureEvent(data []byte) (*NoExposureEvent, error) {
	var ev NoExposureEvent
	r := NewReaderFromData(data)
	NoExposureEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *NoExposureEvent) String() string {
	return fmt.Sprintf("NoExposureEvent {Sequence: %v, Drawable: %v, MinorOpcode: %v, MajorOpcode: %v}",
		v.Sequence, v.Drawable, v.MinorOpcode, v.MajorOpcode)
}

// _go_event_handler_with_xid NoExposureEvent drawable
type _NoExposureEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*NoExposureEvent)
}

func (eh *_NoExposureEventHandler) run(xid uint32, ev *NoExposureEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_NoExposureEventHandler) Run(ge GenericEvent) {
	ev, err := NewNoExposureEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Drawable), ev)
}

func (eh *_NoExposureEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_NoExposureEventHandler) attach(xid uint32, cb func(ev *NoExposureEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*NoExposureEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectNoExposureEvent(conn *Conn, xid Drawable, cb func(ev *NoExposureEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[NoExposureEventCode]
	if !ok {
		newEh := &_NoExposureEventHandler{
			cbs: make(map[uint32][]func(*NoExposureEvent)),
		}
		conn.EventHandlers[NoExposureEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_NoExposureEventHandler).attach(uint32(xid), cb)
}

// enum Visibility
const (
	VisibilityUnobscured        = 0
	VisibilityPartiallyObscured = 1
	VisibilityFullyObscured     = 2
)

func VisibilityEnumToStr(v int) string {
	switch v {
	case 0:
		return "Unobscured"
	case 1:
		return "PartiallyObscured"
	case 2:
		return "FullyObscured"
	default:
		return "<Unknown>"
	}
}

// event name: ('xcb', 'VisibilityNotify')
// self.name: ('xcb', 'VisibilityNotify')
const VisibilityNotifyEventCode = 15

// not event copy
type VisibilityNotifyEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Window   Window
	State    uint8
	// Pad1 [3]uint8
}

func VisibilityNotifyEventRead(r *Reader, v *VisibilityNotifyEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Window
	v.Window = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field State
	v.State = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(3)
	n += 3
	if r.Err() != nil {
		return
	}
	return
}
func VisibilityNotifyEventWrite(w *Writer, v *VisibilityNotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4

	// write field State
	w.Write1b(v.State)
	n += 1

	// write field Pad1
	w.WritePad(3)
	n += 3
	return
}

func NewVisibilityNotifyEvent(data []byte) (*VisibilityNotifyEvent, error) {
	var ev VisibilityNotifyEvent
	r := NewReaderFromData(data)
	VisibilityNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *VisibilityNotifyEvent) String() string {
	return fmt.Sprintf("VisibilityNotifyEvent {Sequence: %v, Window: %v, State: %v}",
		v.Sequence, v.Window, v.State)
}

// _go_event_handler_with_xid VisibilityNotifyEvent window
type _VisibilityNotifyEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*VisibilityNotifyEvent)
}

func (eh *_VisibilityNotifyEventHandler) run(xid uint32, ev *VisibilityNotifyEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_VisibilityNotifyEventHandler) Run(ge GenericEvent) {
	ev, err := NewVisibilityNotifyEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Window), ev)
}

func (eh *_VisibilityNotifyEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_VisibilityNotifyEventHandler) attach(xid uint32, cb func(ev *VisibilityNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*VisibilityNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectVisibilityNotifyEvent(conn *Conn, xid Window, cb func(ev *VisibilityNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[VisibilityNotifyEventCode]
	if !ok {
		newEh := &_VisibilityNotifyEventHandler{
			cbs: make(map[uint32][]func(*VisibilityNotifyEvent)),
		}
		conn.EventHandlers[VisibilityNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_VisibilityNotifyEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'CreateNotify')
// self.name: ('xcb', 'CreateNotify')
const CreateNotifyEventCode = 16

// not event copy
type CreateNotifyEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence         uint16
	Parent           Window
	Window           Window
	X                int16
	Y                int16
	Width            uint16
	Height           uint16
	BorderWidth      uint16
	OverrideRedirect uint8
	// Pad1 uint8
}

func CreateNotifyEventRead(r *Reader, v *CreateNotifyEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Parent
	v.Parent = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Window
	v.Window = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field X
	v.X = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Y
	v.Y = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Width
	v.Width = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Height
	v.Height = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field BorderWidth
	v.BorderWidth = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field OverrideRedirect
	v.OverrideRedirect = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}
	return
}
func CreateNotifyEventWrite(w *Writer, v *CreateNotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Parent
	w.Write4b(uint32(v.Parent))
	n += 4

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4

	// write field X
	w.Write2b(uint16(v.X))
	n += 2

	// write field Y
	w.Write2b(uint16(v.Y))
	n += 2

	// write field Width
	w.Write2b(v.Width)
	n += 2

	// write field Height
	w.Write2b(v.Height)
	n += 2

	// write field BorderWidth
	w.Write2b(v.BorderWidth)
	n += 2

	// write field OverrideRedirect
	w.Write1b(v.OverrideRedirect)
	n += 1

	// write field Pad1
	w.WritePad(1)
	n += 1
	return
}

func NewCreateNotifyEvent(data []byte) (*CreateNotifyEvent, error) {
	var ev CreateNotifyEvent
	r := NewReaderFromData(data)
	CreateNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *CreateNotifyEvent) String() string {
	return fmt.Sprintf("CreateNotifyEvent {Sequence: %v, Parent: %v, Window: %v, X: %v, Y: %v, Width: %v, Height: %v, BorderWidth: %v, OverrideRedirect: %v}",
		v.Sequence, v.Parent, v.Window, v.X, v.Y, v.Width, v.Height, v.BorderWidth, v.OverrideRedirect)
}

// _go_event_handler_with_xid CreateNotifyEvent parent
type _CreateNotifyEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*CreateNotifyEvent)
}

func (eh *_CreateNotifyEventHandler) run(xid uint32, ev *CreateNotifyEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_CreateNotifyEventHandler) Run(ge GenericEvent) {
	ev, err := NewCreateNotifyEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Parent), ev)
}

func (eh *_CreateNotifyEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_CreateNotifyEventHandler) attach(xid uint32, cb func(ev *CreateNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*CreateNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectCreateNotifyEvent(conn *Conn, xid Window, cb func(ev *CreateNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[CreateNotifyEventCode]
	if !ok {
		newEh := &_CreateNotifyEventHandler{
			cbs: make(map[uint32][]func(*CreateNotifyEvent)),
		}
		conn.EventHandlers[CreateNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_CreateNotifyEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'DestroyNotify')
// self.name: ('xcb', 'DestroyNotify')
const DestroyNotifyEventCode = 17

// not event copy
type DestroyNotifyEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Event    Window
	Window   Window
}

func DestroyNotifyEventRead(r *Reader, v *DestroyNotifyEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Event
	v.Event = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Window
	v.Window = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}
	return
}
func DestroyNotifyEventWrite(w *Writer, v *DestroyNotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Event
	w.Write4b(uint32(v.Event))
	n += 4

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4
	return
}

func NewDestroyNotifyEvent(data []byte) (*DestroyNotifyEvent, error) {
	var ev DestroyNotifyEvent
	r := NewReaderFromData(data)
	DestroyNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *DestroyNotifyEvent) String() string {
	return fmt.Sprintf("DestroyNotifyEvent {Sequence: %v, Event: %v, Window: %v}",
		v.Sequence, v.Event, v.Window)
}

// _go_event_handler_with_xid DestroyNotifyEvent window
type _DestroyNotifyEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*DestroyNotifyEvent)
}

func (eh *_DestroyNotifyEventHandler) run(xid uint32, ev *DestroyNotifyEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_DestroyNotifyEventHandler) Run(ge GenericEvent) {
	ev, err := NewDestroyNotifyEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Window), ev)
}

func (eh *_DestroyNotifyEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_DestroyNotifyEventHandler) attach(xid uint32, cb func(ev *DestroyNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*DestroyNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectDestroyNotifyEvent(conn *Conn, xid Window, cb func(ev *DestroyNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[DestroyNotifyEventCode]
	if !ok {
		newEh := &_DestroyNotifyEventHandler{
			cbs: make(map[uint32][]func(*DestroyNotifyEvent)),
		}
		conn.EventHandlers[DestroyNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_DestroyNotifyEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'UnmapNotify')
// self.name: ('xcb', 'UnmapNotify')
const UnmapNotifyEventCode = 18

// not event copy
type UnmapNotifyEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence      uint16
	Event         Window
	Window        Window
	FromConfigure uint8
	// Pad1 [3]uint8
}

func UnmapNotifyEventRead(r *Reader, v *UnmapNotifyEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Event
	v.Event = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Window
	v.Window = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field FromConfigure
	v.FromConfigure = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(3)
	n += 3
	if r.Err() != nil {
		return
	}
	return
}
func UnmapNotifyEventWrite(w *Writer, v *UnmapNotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Event
	w.Write4b(uint32(v.Event))
	n += 4

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4

	// write field FromConfigure
	w.Write1b(v.FromConfigure)
	n += 1

	// write field Pad1
	w.WritePad(3)
	n += 3
	return
}

func NewUnmapNotifyEvent(data []byte) (*UnmapNotifyEvent, error) {
	var ev UnmapNotifyEvent
	r := NewReaderFromData(data)
	UnmapNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *UnmapNotifyEvent) String() string {
	return fmt.Sprintf("UnmapNotifyEvent {Sequence: %v, Event: %v, Window: %v, FromConfigure: %v}",
		v.Sequence, v.Event, v.Window, v.FromConfigure)
}

// _go_event_handler_with_xid UnmapNotifyEvent window
type _UnmapNotifyEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*UnmapNotifyEvent)
}

func (eh *_UnmapNotifyEventHandler) run(xid uint32, ev *UnmapNotifyEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_UnmapNotifyEventHandler) Run(ge GenericEvent) {
	ev, err := NewUnmapNotifyEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Window), ev)
}

func (eh *_UnmapNotifyEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_UnmapNotifyEventHandler) attach(xid uint32, cb func(ev *UnmapNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*UnmapNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectUnmapNotifyEvent(conn *Conn, xid Window, cb func(ev *UnmapNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[UnmapNotifyEventCode]
	if !ok {
		newEh := &_UnmapNotifyEventHandler{
			cbs: make(map[uint32][]func(*UnmapNotifyEvent)),
		}
		conn.EventHandlers[UnmapNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_UnmapNotifyEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'MapNotify')
// self.name: ('xcb', 'MapNotify')
const MapNotifyEventCode = 19

// not event copy
type MapNotifyEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence         uint16
	Event            Window
	Window           Window
	OverrideRedirect uint8
	// Pad1 [3]uint8
}

func MapNotifyEventRead(r *Reader, v *MapNotifyEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Event
	v.Event = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Window
	v.Window = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field OverrideRedirect
	v.OverrideRedirect = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(3)
	n += 3
	if r.Err() != nil {
		return
	}
	return
}
func MapNotifyEventWrite(w *Writer, v *MapNotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Event
	w.Write4b(uint32(v.Event))
	n += 4

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4

	// write field OverrideRedirect
	w.Write1b(v.OverrideRedirect)
	n += 1

	// write field Pad1
	w.WritePad(3)
	n += 3
	return
}

func NewMapNotifyEvent(data []byte) (*MapNotifyEvent, error) {
	var ev MapNotifyEvent
	r := NewReaderFromData(data)
	MapNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *MapNotifyEvent) String() string {
	return fmt.Sprintf("MapNotifyEvent {Sequence: %v, Event: %v, Window: %v, OverrideRedirect: %v}",
		v.Sequence, v.Event, v.Window, v.OverrideRedirect)
}

// _go_event_handler_with_xid MapNotifyEvent window
type _MapNotifyEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*MapNotifyEvent)
}

func (eh *_MapNotifyEventHandler) run(xid uint32, ev *MapNotifyEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_MapNotifyEventHandler) Run(ge GenericEvent) {
	ev, err := NewMapNotifyEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Window), ev)
}

func (eh *_MapNotifyEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_MapNotifyEventHandler) attach(xid uint32, cb func(ev *MapNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*MapNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectMapNotifyEvent(conn *Conn, xid Window, cb func(ev *MapNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[MapNotifyEventCode]
	if !ok {
		newEh := &_MapNotifyEventHandler{
			cbs: make(map[uint32][]func(*MapNotifyEvent)),
		}
		conn.EventHandlers[MapNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_MapNotifyEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'MapRequest')
// self.name: ('xcb', 'MapRequest')
const MapRequestEventCode = 20

// not event copy
type MapRequestEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Parent   Window
	Window   Window
}

func MapRequestEventRead(r *Reader, v *MapRequestEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Parent
	v.Parent = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Window
	v.Window = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}
	return
}
func MapRequestEventWrite(w *Writer, v *MapRequestEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Parent
	w.Write4b(uint32(v.Parent))
	n += 4

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4
	return
}

func NewMapRequestEvent(data []byte) (*MapRequestEvent, error) {
	var ev MapRequestEvent
	r := NewReaderFromData(data)
	MapRequestEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *MapRequestEvent) String() string {
	return fmt.Sprintf("MapRequestEvent {Sequence: %v, Parent: %v, Window: %v}",
		v.Sequence, v.Parent, v.Window)
}

// _go_event_handler_with_xid MapRequestEvent window
type _MapRequestEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*MapRequestEvent)
}

func (eh *_MapRequestEventHandler) run(xid uint32, ev *MapRequestEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_MapRequestEventHandler) Run(ge GenericEvent) {
	ev, err := NewMapRequestEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Window), ev)
}

func (eh *_MapRequestEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_MapRequestEventHandler) attach(xid uint32, cb func(ev *MapRequestEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*MapRequestEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectMapRequestEvent(conn *Conn, xid Window, cb func(ev *MapRequestEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[MapRequestEventCode]
	if !ok {
		newEh := &_MapRequestEventHandler{
			cbs: make(map[uint32][]func(*MapRequestEvent)),
		}
		conn.EventHandlers[MapRequestEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_MapRequestEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'ReparentNotify')
// self.name: ('xcb', 'ReparentNotify')
const ReparentNotifyEventCode = 21

// not event copy
type ReparentNotifyEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence         uint16
	Event            Window
	Window           Window
	Parent           Window
	X                int16
	Y                int16
	OverrideRedirect uint8
	// Pad1 [3]uint8
}

func ReparentNotifyEventRead(r *Reader, v *ReparentNotifyEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Event
	v.Event = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Window
	v.Window = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Parent
	v.Parent = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field X
	v.X = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Y
	v.Y = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field OverrideRedirect
	v.OverrideRedirect = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(3)
	n += 3
	if r.Err() != nil {
		return
	}
	return
}
func ReparentNotifyEventWrite(w *Writer, v *ReparentNotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Event
	w.Write4b(uint32(v.Event))
	n += 4

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4

	// write field Parent
	w.Write4b(uint32(v.Parent))
	n += 4

	// write field X
	w.Write2b(uint16(v.X))
	n += 2

	// write field Y
	w.Write2b(uint16(v.Y))
	n += 2

	// write field OverrideRedirect
	w.Write1b(v.OverrideRedirect)
	n += 1

	// write field Pad1
	w.WritePad(3)
	n += 3
	return
}

func NewReparentNotifyEvent(data []byte) (*ReparentNotifyEvent, error) {
	var ev ReparentNotifyEvent
	r := NewReaderFromData(data)
	ReparentNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *ReparentNotifyEvent) String() string {
	return fmt.Sprintf("ReparentNotifyEvent {Sequence: %v, Event: %v, Window: %v, Parent: %v, X: %v, Y: %v, OverrideRedirect: %v}",
		v.Sequence, v.Event, v.Window, v.Parent, v.X, v.Y, v.OverrideRedirect)
}

// _go_event_handler_with_xid ReparentNotifyEvent window
type _ReparentNotifyEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*ReparentNotifyEvent)
}

func (eh *_ReparentNotifyEventHandler) run(xid uint32, ev *ReparentNotifyEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_ReparentNotifyEventHandler) Run(ge GenericEvent) {
	ev, err := NewReparentNotifyEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Window), ev)
}

func (eh *_ReparentNotifyEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_ReparentNotifyEventHandler) attach(xid uint32, cb func(ev *ReparentNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*ReparentNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectReparentNotifyEvent(conn *Conn, xid Window, cb func(ev *ReparentNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[ReparentNotifyEventCode]
	if !ok {
		newEh := &_ReparentNotifyEventHandler{
			cbs: make(map[uint32][]func(*ReparentNotifyEvent)),
		}
		conn.EventHandlers[ReparentNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_ReparentNotifyEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'ConfigureNotify')
// self.name: ('xcb', 'ConfigureNotify')
const ConfigureNotifyEventCode = 22

// not event copy
type ConfigureNotifyEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence         uint16
	Event            Window
	Window           Window
	AboveSibling     Window
	X                int16
	Y                int16
	Width            uint16
	Height           uint16
	BorderWidth      uint16
	OverrideRedirect uint8
	// Pad1 uint8
}

func ConfigureNotifyEventRead(r *Reader, v *ConfigureNotifyEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Event
	v.Event = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Window
	v.Window = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field AboveSibling
	v.AboveSibling = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field X
	v.X = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Y
	v.Y = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Width
	v.Width = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Height
	v.Height = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field BorderWidth
	v.BorderWidth = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field OverrideRedirect
	v.OverrideRedirect = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}
	return
}
func ConfigureNotifyEventWrite(w *Writer, v *ConfigureNotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Event
	w.Write4b(uint32(v.Event))
	n += 4

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4

	// write field AboveSibling
	w.Write4b(uint32(v.AboveSibling))
	n += 4

	// write field X
	w.Write2b(uint16(v.X))
	n += 2

	// write field Y
	w.Write2b(uint16(v.Y))
	n += 2

	// write field Width
	w.Write2b(v.Width)
	n += 2

	// write field Height
	w.Write2b(v.Height)
	n += 2

	// write field BorderWidth
	w.Write2b(v.BorderWidth)
	n += 2

	// write field OverrideRedirect
	w.Write1b(v.OverrideRedirect)
	n += 1

	// write field Pad1
	w.WritePad(1)
	n += 1
	return
}

func NewConfigureNotifyEvent(data []byte) (*ConfigureNotifyEvent, error) {
	var ev ConfigureNotifyEvent
	r := NewReaderFromData(data)
	ConfigureNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *ConfigureNotifyEvent) String() string {
	return fmt.Sprintf("ConfigureNotifyEvent {Sequence: %v, Event: %v, Window: %v, AboveSibling: %v, X: %v, Y: %v, Width: %v, Height: %v, BorderWidth: %v, OverrideRedirect: %v}",
		v.Sequence, v.Event, v.Window, v.AboveSibling, v.X, v.Y, v.Width, v.Height, v.BorderWidth, v.OverrideRedirect)
}

// _go_event_handler_with_xid ConfigureNotifyEvent window
type _ConfigureNotifyEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*ConfigureNotifyEvent)
}

func (eh *_ConfigureNotifyEventHandler) run(xid uint32, ev *ConfigureNotifyEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_ConfigureNotifyEventHandler) Run(ge GenericEvent) {
	ev, err := NewConfigureNotifyEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Window), ev)
}

func (eh *_ConfigureNotifyEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_ConfigureNotifyEventHandler) attach(xid uint32, cb func(ev *ConfigureNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*ConfigureNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectConfigureNotifyEvent(conn *Conn, xid Window, cb func(ev *ConfigureNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[ConfigureNotifyEventCode]
	if !ok {
		newEh := &_ConfigureNotifyEventHandler{
			cbs: make(map[uint32][]func(*ConfigureNotifyEvent)),
		}
		conn.EventHandlers[ConfigureNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_ConfigureNotifyEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'ConfigureRequest')
// self.name: ('xcb', 'ConfigureRequest')
const ConfigureRequestEventCode = 23

// not event copy
type ConfigureRequestEvent struct {
	ResponseType uint8
	StackMode    uint8
	Sequence     uint16
	Parent       Window
	Window       Window
	Sibling      Window
	X            int16
	Y            int16
	Width        uint16
	Height       uint16
	BorderWidth  uint16
	ValueMask    uint16
}

func ConfigureRequestEventRead(r *Reader, v *ConfigureRequestEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field StackMode
	v.StackMode = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Parent
	v.Parent = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Window
	v.Window = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Sibling
	v.Sibling = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field X
	v.X = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Y
	v.Y = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Width
	v.Width = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Height
	v.Height = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field BorderWidth
	v.BorderWidth = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field ValueMask
	v.ValueMask = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}
	return
}
func ConfigureRequestEventWrite(w *Writer, v *ConfigureRequestEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field StackMode
	w.Write1b(v.StackMode)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Parent
	w.Write4b(uint32(v.Parent))
	n += 4

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4

	// write field Sibling
	w.Write4b(uint32(v.Sibling))
	n += 4

	// write field X
	w.Write2b(uint16(v.X))
	n += 2

	// write field Y
	w.Write2b(uint16(v.Y))
	n += 2

	// write field Width
	w.Write2b(v.Width)
	n += 2

	// write field Height
	w.Write2b(v.Height)
	n += 2

	// write field BorderWidth
	w.Write2b(v.BorderWidth)
	n += 2

	// write field ValueMask
	w.Write2b(v.ValueMask)
	n += 2
	return
}

func NewConfigureRequestEvent(data []byte) (*ConfigureRequestEvent, error) {
	var ev ConfigureRequestEvent
	r := NewReaderFromData(data)
	ConfigureRequestEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *ConfigureRequestEvent) String() string {
	return fmt.Sprintf("ConfigureRequestEvent {StackMode: %v, Sequence: %v, Parent: %v, Window: %v, Sibling: %v, X: %v, Y: %v, Width: %v, Height: %v, BorderWidth: %v, ValueMask: %v}",
		v.StackMode, v.Sequence, v.Parent, v.Window, v.Sibling, v.X, v.Y, v.Width, v.Height, v.BorderWidth, v.ValueMask)
}

// _go_event_handler_with_xid ConfigureRequestEvent window
type _ConfigureRequestEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*ConfigureRequestEvent)
}

func (eh *_ConfigureRequestEventHandler) run(xid uint32, ev *ConfigureRequestEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_ConfigureRequestEventHandler) Run(ge GenericEvent) {
	ev, err := NewConfigureRequestEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Window), ev)
}

func (eh *_ConfigureRequestEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_ConfigureRequestEventHandler) attach(xid uint32, cb func(ev *ConfigureRequestEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*ConfigureRequestEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectConfigureRequestEvent(conn *Conn, xid Window, cb func(ev *ConfigureRequestEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[ConfigureRequestEventCode]
	if !ok {
		newEh := &_ConfigureRequestEventHandler{
			cbs: make(map[uint32][]func(*ConfigureRequestEvent)),
		}
		conn.EventHandlers[ConfigureRequestEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_ConfigureRequestEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'GravityNotify')
// self.name: ('xcb', 'GravityNotify')
const GravityNotifyEventCode = 24

// not event copy
type GravityNotifyEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Event    Window
	Window   Window
	X        int16
	Y        int16
}

func GravityNotifyEventRead(r *Reader, v *GravityNotifyEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Event
	v.Event = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Window
	v.Window = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field X
	v.X = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Y
	v.Y = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}
	return
}
func GravityNotifyEventWrite(w *Writer, v *GravityNotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Event
	w.Write4b(uint32(v.Event))
	n += 4

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4

	// write field X
	w.Write2b(uint16(v.X))
	n += 2

	// write field Y
	w.Write2b(uint16(v.Y))
	n += 2
	return
}

func NewGravityNotifyEvent(data []byte) (*GravityNotifyEvent, error) {
	var ev GravityNotifyEvent
	r := NewReaderFromData(data)
	GravityNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *GravityNotifyEvent) String() string {
	return fmt.Sprintf("GravityNotifyEvent {Sequence: %v, Event: %v, Window: %v, X: %v, Y: %v}",
		v.Sequence, v.Event, v.Window, v.X, v.Y)
}

// _go_event_handler_with_xid GravityNotifyEvent window
type _GravityNotifyEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*GravityNotifyEvent)
}

func (eh *_GravityNotifyEventHandler) run(xid uint32, ev *GravityNotifyEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_GravityNotifyEventHandler) Run(ge GenericEvent) {
	ev, err := NewGravityNotifyEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Window), ev)
}

func (eh *_GravityNotifyEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_GravityNotifyEventHandler) attach(xid uint32, cb func(ev *GravityNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*GravityNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectGravityNotifyEvent(conn *Conn, xid Window, cb func(ev *GravityNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[GravityNotifyEventCode]
	if !ok {
		newEh := &_GravityNotifyEventHandler{
			cbs: make(map[uint32][]func(*GravityNotifyEvent)),
		}
		conn.EventHandlers[GravityNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_GravityNotifyEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'ResizeRequest')
// self.name: ('xcb', 'ResizeRequest')
const ResizeRequestEventCode = 25

// not event copy
type ResizeRequestEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Window   Window
	Width    uint16
	Height   uint16
}

func ResizeRequestEventRead(r *Reader, v *ResizeRequestEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Window
	v.Window = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Width
	v.Width = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Height
	v.Height = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}
	return
}
func ResizeRequestEventWrite(w *Writer, v *ResizeRequestEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4

	// write field Width
	w.Write2b(v.Width)
	n += 2

	// write field Height
	w.Write2b(v.Height)
	n += 2
	return
}

func NewResizeRequestEvent(data []byte) (*ResizeRequestEvent, error) {
	var ev ResizeRequestEvent
	r := NewReaderFromData(data)
	ResizeRequestEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *ResizeRequestEvent) String() string {
	return fmt.Sprintf("ResizeRequestEvent {Sequence: %v, Window: %v, Width: %v, Height: %v}",
		v.Sequence, v.Window, v.Width, v.Height)
}

// _go_event_handler_with_xid ResizeRequestEvent window
type _ResizeRequestEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*ResizeRequestEvent)
}

func (eh *_ResizeRequestEventHandler) run(xid uint32, ev *ResizeRequestEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_ResizeRequestEventHandler) Run(ge GenericEvent) {
	ev, err := NewResizeRequestEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Window), ev)
}

func (eh *_ResizeRequestEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_ResizeRequestEventHandler) attach(xid uint32, cb func(ev *ResizeRequestEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*ResizeRequestEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectResizeRequestEvent(conn *Conn, xid Window, cb func(ev *ResizeRequestEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[ResizeRequestEventCode]
	if !ok {
		newEh := &_ResizeRequestEventHandler{
			cbs: make(map[uint32][]func(*ResizeRequestEvent)),
		}
		conn.EventHandlers[ResizeRequestEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_ResizeRequestEventHandler).attach(uint32(xid), cb)
}

// enum Place
const (
	PlaceOnTop    = 0
	PlaceOnBottom = 1
)

func PlaceEnumToStr(v int) string {
	switch v {
	case 0:
		return "OnTop"
	case 1:
		return "OnBottom"
	default:
		return "<Unknown>"
	}
}

// event name: ('xcb', 'CirculateNotify')
// self.name: ('xcb', 'CirculateNotify')
const CirculateNotifyEventCode = 26

// not event copy
type CirculateNotifyEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Event    Window
	Window   Window
	// Pad1 [4]uint8
	Place uint8
	// Pad2 [3]uint8
}

func CirculateNotifyEventRead(r *Reader, v *CirculateNotifyEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Event
	v.Event = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Window
	v.Window = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(4)
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Place
	v.Place = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad2
	r.ReadPad(3)
	n += 3
	if r.Err() != nil {
		return
	}
	return
}
func CirculateNotifyEventWrite(w *Writer, v *CirculateNotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Event
	w.Write4b(uint32(v.Event))
	n += 4

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4

	// write field Pad1
	w.WritePad(4)
	n += 4

	// write field Place
	w.Write1b(v.Place)
	n += 1

	// write field Pad2
	w.WritePad(3)
	n += 3
	return
}

func NewCirculateNotifyEvent(data []byte) (*CirculateNotifyEvent, error) {
	var ev CirculateNotifyEvent
	r := NewReaderFromData(data)
	CirculateNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *CirculateNotifyEvent) String() string {
	return fmt.Sprintf("CirculateNotifyEvent {Sequence: %v, Event: %v, Window: %v, Place: %v}",
		v.Sequence, v.Event, v.Window, v.Place)
}

// _go_event_handler_with_xid CirculateNotifyEvent window
type _CirculateNotifyEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*CirculateNotifyEvent)
}

func (eh *_CirculateNotifyEventHandler) run(xid uint32, ev *CirculateNotifyEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_CirculateNotifyEventHandler) Run(ge GenericEvent) {
	ev, err := NewCirculateNotifyEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Window), ev)
}

func (eh *_CirculateNotifyEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_CirculateNotifyEventHandler) attach(xid uint32, cb func(ev *CirculateNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*CirculateNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectCirculateNotifyEvent(conn *Conn, xid Window, cb func(ev *CirculateNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[CirculateNotifyEventCode]
	if !ok {
		newEh := &_CirculateNotifyEventHandler{
			cbs: make(map[uint32][]func(*CirculateNotifyEvent)),
		}
		conn.EventHandlers[CirculateNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_CirculateNotifyEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'CirculateRequest')
// self.name: ('xcb', 'CirculateNotify')
const CirculateRequestEventCode = 27

// is event copy
type CirculateRequestEvent CirculateNotifyEvent

func CirculateRequestEventRead(r *Reader, v *CirculateRequestEvent) (n int) {
	// source_type_name: CirculateNotifyEvent
	return CirculateNotifyEventRead(r, (*CirculateNotifyEvent)(v))
}

func NewCirculateRequestEvent(data []byte) (*CirculateRequestEvent, error) {
	var ev CirculateRequestEvent
	r := NewReaderFromData(data)
	CirculateRequestEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *CirculateRequestEvent) String() string {
	return fmt.Sprintf("CirculateRequestEvent {Sequence: %v, Event: %v, Window: %v, Place: %v}",
		v.Sequence, v.Event, v.Window, v.Place)
}

// try event source CirculateNotify
// _go_event_handler_with_xid CirculateRequestEvent window
type _CirculateRequestEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*CirculateRequestEvent)
}

func (eh *_CirculateRequestEventHandler) run(xid uint32, ev *CirculateRequestEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_CirculateRequestEventHandler) Run(ge GenericEvent) {
	ev, err := NewCirculateRequestEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Window), ev)
}

func (eh *_CirculateRequestEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_CirculateRequestEventHandler) attach(xid uint32, cb func(ev *CirculateRequestEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*CirculateRequestEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectCirculateRequestEvent(conn *Conn, xid Window, cb func(ev *CirculateRequestEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[CirculateRequestEventCode]
	if !ok {
		newEh := &_CirculateRequestEventHandler{
			cbs: make(map[uint32][]func(*CirculateRequestEvent)),
		}
		conn.EventHandlers[CirculateRequestEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_CirculateRequestEventHandler).attach(uint32(xid), cb)
}

// enum Property
const (
	PropertyNewValue = 0
	PropertyDelete   = 1
)

// event name: ('xcb', 'PropertyNotify')
// self.name: ('xcb', 'PropertyNotify')
const PropertyNotifyEventCode = 28

// not event copy
type PropertyNotifyEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Window   Window
	Atom     Atom
	Time     Timestamp
	State    uint8
	// Pad1 [3]uint8
}

func PropertyNotifyEventRead(r *Reader, v *PropertyNotifyEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Window
	v.Window = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Atom
	v.Atom = Atom(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Time
	v.Time = Timestamp(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field State
	v.State = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(3)
	n += 3
	if r.Err() != nil {
		return
	}
	return
}
func PropertyNotifyEventWrite(w *Writer, v *PropertyNotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4

	// write field Atom
	w.Write4b(uint32(v.Atom))
	n += 4

	// write field Time
	w.Write4b(uint32(v.Time))
	n += 4

	// write field State
	w.Write1b(v.State)
	n += 1

	// write field Pad1
	w.WritePad(3)
	n += 3
	return
}

func NewPropertyNotifyEvent(data []byte) (*PropertyNotifyEvent, error) {
	var ev PropertyNotifyEvent
	r := NewReaderFromData(data)
	PropertyNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *PropertyNotifyEvent) String() string {
	return fmt.Sprintf("PropertyNotifyEvent {Sequence: %v, Window: %v, Atom: %v, Time: %v, State: %v}",
		v.Sequence, v.Window, v.Atom, v.Time, v.State)
}

// _go_event_handler_with_xid PropertyNotifyEvent window
type _PropertyNotifyEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*PropertyNotifyEvent)
}

func (eh *_PropertyNotifyEventHandler) run(xid uint32, ev *PropertyNotifyEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_PropertyNotifyEventHandler) Run(ge GenericEvent) {
	ev, err := NewPropertyNotifyEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Window), ev)
}

func (eh *_PropertyNotifyEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_PropertyNotifyEventHandler) attach(xid uint32, cb func(ev *PropertyNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*PropertyNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectPropertyNotifyEvent(conn *Conn, xid Window, cb func(ev *PropertyNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[PropertyNotifyEventCode]
	if !ok {
		newEh := &_PropertyNotifyEventHandler{
			cbs: make(map[uint32][]func(*PropertyNotifyEvent)),
		}
		conn.EventHandlers[PropertyNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_PropertyNotifyEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'SelectionClear')
// self.name: ('xcb', 'SelectionClear')
const SelectionClearEventCode = 29

// not event copy
type SelectionClearEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence  uint16
	Time      Timestamp
	Owner     Window
	Selection Atom
}

func SelectionClearEventRead(r *Reader, v *SelectionClearEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Time
	v.Time = Timestamp(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Owner
	v.Owner = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Selection
	v.Selection = Atom(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}
	return
}
func SelectionClearEventWrite(w *Writer, v *SelectionClearEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Time
	w.Write4b(uint32(v.Time))
	n += 4

	// write field Owner
	w.Write4b(uint32(v.Owner))
	n += 4

	// write field Selection
	w.Write4b(uint32(v.Selection))
	n += 4
	return
}

func NewSelectionClearEvent(data []byte) (*SelectionClearEvent, error) {
	var ev SelectionClearEvent
	r := NewReaderFromData(data)
	SelectionClearEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *SelectionClearEvent) String() string {
	return fmt.Sprintf("SelectionClearEvent {Sequence: %v, Time: %v, Owner: %v, Selection: %v}",
		v.Sequence, v.Time, v.Owner, v.Selection)
}

// _go_event_handler_with_xid SelectionClearEvent owner
type _SelectionClearEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*SelectionClearEvent)
}

func (eh *_SelectionClearEventHandler) run(xid uint32, ev *SelectionClearEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_SelectionClearEventHandler) Run(ge GenericEvent) {
	ev, err := NewSelectionClearEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Owner), ev)
}

func (eh *_SelectionClearEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_SelectionClearEventHandler) attach(xid uint32, cb func(ev *SelectionClearEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*SelectionClearEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectSelectionClearEvent(conn *Conn, xid Window, cb func(ev *SelectionClearEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[SelectionClearEventCode]
	if !ok {
		newEh := &_SelectionClearEventHandler{
			cbs: make(map[uint32][]func(*SelectionClearEvent)),
		}
		conn.EventHandlers[SelectionClearEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_SelectionClearEventHandler).attach(uint32(xid), cb)
}

// enum Time
const (
	TimeCurrentTime = 0
)

// enum Atom
const (
	AtomNone               = 0
	AtomAny                = 0
	AtomPrimary            = 1
	AtomSecondary          = 2
	AtomArc                = 3
	AtomAtom               = 4
	AtomBitmap             = 5
	AtomCardinal           = 6
	AtomColormap           = 7
	AtomCursor             = 8
	AtomCutBuffer0         = 9
	AtomCutBuffer1         = 10
	AtomCutBuffer2         = 11
	AtomCutBuffer3         = 12
	AtomCutBuffer4         = 13
	AtomCutBuffer5         = 14
	AtomCutBuffer6         = 15
	AtomCutBuffer7         = 16
	AtomDrawable           = 17
	AtomFont               = 18
	AtomInteger            = 19
	AtomPixmap             = 20
	AtomPoint              = 21
	AtomRectangle          = 22
	AtomResourceManager    = 23
	AtomRGBColorMap        = 24
	AtomRGBBestMap         = 25
	AtomRGBBlueMap         = 26
	AtomRGBDefaultMap      = 27
	AtomRGBGrayMap         = 28
	AtomRGBGreenMap        = 29
	AtomRGBRedMap          = 30
	AtomString             = 31
	AtomVisualID           = 32
	AtomWindow             = 33
	AtomWMCommand          = 34
	AtomWMHints            = 35
	AtomWMClientMachine    = 36
	AtomWMIconName         = 37
	AtomWMIconSize         = 38
	AtomWMName             = 39
	AtomWMNormalHints      = 40
	AtomWMSizeHints        = 41
	AtomWMZoomHints        = 42
	AtomMinSpace           = 43
	AtomNormSpace          = 44
	AtomMaxSpace           = 45
	AtomEndSpace           = 46
	AtomSuperscriptX       = 47
	AtomSuperscriptY       = 48
	AtomSubscriptX         = 49
	AtomSubscriptY         = 50
	AtomUnderlinePosition  = 51
	AtomUnderlineThickness = 52
	AtomStrikeoutAscent    = 53
	AtomStrikeoutDescent   = 54
	AtomItalicAngle        = 55
	AtomXHeight            = 56
	AtomQuadWidth          = 57
	AtomWeight             = 58
	AtomPointSize          = 59
	AtomResolution         = 60
	AtomCopyright          = 61
	AtomNotice             = 62
	AtomFontName           = 63
	AtomFamilyName         = 64
	AtomFullName           = 65
	AtomCapHeight          = 66
	AtomWMClass            = 67
	AtomWMTransientFor     = 68
)

// event name: ('xcb', 'SelectionRequest')
// self.name: ('xcb', 'SelectionRequest')
const SelectionRequestEventCode = 30

// not event copy
type SelectionRequestEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence  uint16
	Time      Timestamp
	Owner     Window
	Requestor Window
	Selection Atom
	Target    Atom
	Property  Atom
}

func SelectionRequestEventRead(r *Reader, v *SelectionRequestEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Time
	v.Time = Timestamp(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Owner
	v.Owner = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Requestor
	v.Requestor = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Selection
	v.Selection = Atom(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Target
	v.Target = Atom(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Property
	v.Property = Atom(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}
	return
}
func SelectionRequestEventWrite(w *Writer, v *SelectionRequestEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Time
	w.Write4b(uint32(v.Time))
	n += 4

	// write field Owner
	w.Write4b(uint32(v.Owner))
	n += 4

	// write field Requestor
	w.Write4b(uint32(v.Requestor))
	n += 4

	// write field Selection
	w.Write4b(uint32(v.Selection))
	n += 4

	// write field Target
	w.Write4b(uint32(v.Target))
	n += 4

	// write field Property
	w.Write4b(uint32(v.Property))
	n += 4
	return
}

func NewSelectionRequestEvent(data []byte) (*SelectionRequestEvent, error) {
	var ev SelectionRequestEvent
	r := NewReaderFromData(data)
	SelectionRequestEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *SelectionRequestEvent) String() string {
	return fmt.Sprintf("SelectionRequestEvent {Sequence: %v, Time: %v, Owner: %v, Requestor: %v, Selection: %v, Target: %v, Property: %v}",
		v.Sequence, v.Time, v.Owner, v.Requestor, v.Selection, v.Target, v.Property)
}

// _go_event_handler_with_xid SelectionRequestEvent requestor
type _SelectionRequestEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*SelectionRequestEvent)
}

func (eh *_SelectionRequestEventHandler) run(xid uint32, ev *SelectionRequestEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_SelectionRequestEventHandler) Run(ge GenericEvent) {
	ev, err := NewSelectionRequestEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Requestor), ev)
}

func (eh *_SelectionRequestEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_SelectionRequestEventHandler) attach(xid uint32, cb func(ev *SelectionRequestEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*SelectionRequestEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectSelectionRequestEvent(conn *Conn, xid Window, cb func(ev *SelectionRequestEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[SelectionRequestEventCode]
	if !ok {
		newEh := &_SelectionRequestEventHandler{
			cbs: make(map[uint32][]func(*SelectionRequestEvent)),
		}
		conn.EventHandlers[SelectionRequestEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_SelectionRequestEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'SelectionNotify')
// self.name: ('xcb', 'SelectionNotify')
const SelectionNotifyEventCode = 31

// not event copy
type SelectionNotifyEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence  uint16
	Time      Timestamp
	Requestor Window
	Selection Atom
	Target    Atom
	Property  Atom
}

func SelectionNotifyEventRead(r *Reader, v *SelectionNotifyEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Time
	v.Time = Timestamp(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Requestor
	v.Requestor = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Selection
	v.Selection = Atom(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Target
	v.Target = Atom(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Property
	v.Property = Atom(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}
	return
}
func SelectionNotifyEventWrite(w *Writer, v *SelectionNotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Time
	w.Write4b(uint32(v.Time))
	n += 4

	// write field Requestor
	w.Write4b(uint32(v.Requestor))
	n += 4

	// write field Selection
	w.Write4b(uint32(v.Selection))
	n += 4

	// write field Target
	w.Write4b(uint32(v.Target))
	n += 4

	// write field Property
	w.Write4b(uint32(v.Property))
	n += 4
	return
}

func NewSelectionNotifyEvent(data []byte) (*SelectionNotifyEvent, error) {
	var ev SelectionNotifyEvent
	r := NewReaderFromData(data)
	SelectionNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *SelectionNotifyEvent) String() string {
	return fmt.Sprintf("SelectionNotifyEvent {Sequence: %v, Time: %v, Requestor: %v, Selection: %v, Target: %v, Property: %v}",
		v.Sequence, v.Time, v.Requestor, v.Selection, v.Target, v.Property)
}

// try event source SelectionNotify
type _SelectionNotifyEventHandler struct {
	mu  sync.Mutex
	cbs []func(*SelectionNotifyEvent)
}

func (eh *_SelectionNotifyEventHandler) Run(ge GenericEvent) {
	ev, err := NewSelectionNotifyEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.mu.Lock()
	cbs := eh.cbs
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (_ *_SelectionNotifyEventHandler) Detach(_ uint32) {}

func (eh *_SelectionNotifyEventHandler) attach(cb func(ev *SelectionNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs
	newCbs := make([]func(*SelectionNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs = newCbs
	eh.mu.Unlock()
}

func ConnectSelectionNotifyEvent(conn *Conn, cb func(ev *SelectionNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[SelectionNotifyEventCode]
	if !ok {
		newEh := &_SelectionNotifyEventHandler{}
		conn.EventHandlers[SelectionNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_SelectionNotifyEventHandler).attach(cb)
}

// enum ColormapState
const (
	ColormapStateUninstalled = 0
	ColormapStateInstalled   = 1
)

func ColormapStateEnumToStr(v int) string {
	switch v {
	case 0:
		return "Uninstalled"
	case 1:
		return "Installed"
	default:
		return "<Unknown>"
	}
}

// enum Colormap
const (
	ColormapNone = 0
)

// event name: ('xcb', 'ColormapNotify')
// self.name: ('xcb', 'ColormapNotify')
const ColormapNotifyEventCode = 32

// not event copy
type ColormapNotifyEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Window   Window
	Colormap Colormap
	New      uint8
	State    uint8
	// Pad1 [2]uint8
}

func ColormapNotifyEventRead(r *Reader, v *ColormapNotifyEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Window
	v.Window = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Colormap
	v.Colormap = Colormap(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field New
	v.New = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field State
	v.State = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(2)
	n += 2
	if r.Err() != nil {
		return
	}
	return
}
func ColormapNotifyEventWrite(w *Writer, v *ColormapNotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4

	// write field Colormap
	w.Write4b(uint32(v.Colormap))
	n += 4

	// write field New
	w.Write1b(v.New)
	n += 1

	// write field State
	w.Write1b(v.State)
	n += 1

	// write field Pad1
	w.WritePad(2)
	n += 2
	return
}

func NewColormapNotifyEvent(data []byte) (*ColormapNotifyEvent, error) {
	var ev ColormapNotifyEvent
	r := NewReaderFromData(data)
	ColormapNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *ColormapNotifyEvent) String() string {
	return fmt.Sprintf("ColormapNotifyEvent {Sequence: %v, Window: %v, Colormap: %v, New: %v, State: %v}",
		v.Sequence, v.Window, v.Colormap, v.New, v.State)
}

// _go_event_handler_with_xid ColormapNotifyEvent window
type _ColormapNotifyEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*ColormapNotifyEvent)
}

func (eh *_ColormapNotifyEventHandler) run(xid uint32, ev *ColormapNotifyEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_ColormapNotifyEventHandler) Run(ge GenericEvent) {
	ev, err := NewColormapNotifyEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Window), ev)
}

func (eh *_ColormapNotifyEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_ColormapNotifyEventHandler) attach(xid uint32, cb func(ev *ColormapNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*ColormapNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectColormapNotifyEvent(conn *Conn, xid Window, cb func(ev *ColormapNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[ColormapNotifyEventCode]
	if !ok {
		newEh := &_ColormapNotifyEventHandler{
			cbs: make(map[uint32][]func(*ColormapNotifyEvent)),
		}
		conn.EventHandlers[ColormapNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_ColormapNotifyEventHandler).attach(uint32(xid), cb)
}

// event name: ('xcb', 'ClientMessage')
// self.name: ('xcb', 'ClientMessage')
const ClientMessageEventCode = 33

// not event copy
type ClientMessageEvent struct {
	ResponseType uint8
	Format       uint8
	Sequence     uint16
	Window       Window
	Type         Atom
	Data         ClientMessageData
}

func ClientMessageEventRead(r *Reader, v *ClientMessageEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Format
	v.Format = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Window
	v.Window = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Type
	v.Type = Atom(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Data
	n += ClientMessageDataRead(r, &v.Data)
	if r.Err() != nil {
		return
	}
	return
}
func ClientMessageEventWrite(w *Writer, v *ClientMessageEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Format
	w.Write1b(v.Format)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4

	// write field Type
	w.Write4b(uint32(v.Type))
	n += 4

	// write field Data
	n += ClientMessageDataWrite(w, &v.Data)
	return
}

func NewClientMessageEvent(data []byte) (*ClientMessageEvent, error) {
	var ev ClientMessageEvent
	r := NewReaderFromData(data)
	ClientMessageEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *ClientMessageEvent) String() string {
	return fmt.Sprintf("ClientMessageEvent {Format: %v, Sequence: %v, Window: %v, Type: %v, Data: %v}",
		v.Format, v.Sequence, v.Window, v.Type, v.Data)
}

// _go_event_handler_with_xid ClientMessageEvent window
type _ClientMessageEventHandler struct {
	mu  sync.Mutex
	cbs map[uint32][]func(*ClientMessageEvent)
}

func (eh *_ClientMessageEventHandler) run(xid uint32, ev *ClientMessageEvent) {
	eh.mu.Lock()
	cbs := eh.cbs[xid]
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (eh *_ClientMessageEventHandler) Run(ge GenericEvent) {
	ev, err := NewClientMessageEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.run(uint32(ev.Window), ev)
}

func (eh *_ClientMessageEventHandler) Detach(xid uint32) {
	eh.mu.Lock()
	delete(eh.cbs, xid)
	eh.mu.Unlock()
}

func (eh *_ClientMessageEventHandler) attach(xid uint32, cb func(ev *ClientMessageEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs[xid]
	newCbs := make([]func(*ClientMessageEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs[xid] = newCbs
	eh.mu.Unlock()
}

func ConnectClientMessageEvent(conn *Conn, xid Window, cb func(ev *ClientMessageEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[ClientMessageEventCode]
	if !ok {
		newEh := &_ClientMessageEventHandler{
			cbs: make(map[uint32][]func(*ClientMessageEvent)),
		}
		conn.EventHandlers[ClientMessageEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(uint32(xid), cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_ClientMessageEventHandler).attach(uint32(xid), cb)
}

// enum Mapping
const (
	MappingModifier = 0
	MappingKeyboard = 1
	MappingPointer  = 2
)

func MappingEnumToStr(v int) string {
	switch v {
	case 0:
		return "Modifier"
	case 1:
		return "Keyboard"
	case 2:
		return "Pointer"
	default:
		return "<Unknown>"
	}
}

// event name: ('xcb', 'MappingNotify')
// self.name: ('xcb', 'MappingNotify')
const MappingNotifyEventCode = 34

// not event copy
type MappingNotifyEvent struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence     uint16
	Request      uint8
	FirstKeycode Keycode
	Count        uint8
	// Pad1 uint8
}

func MappingNotifyEventRead(r *Reader, v *MappingNotifyEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Request
	v.Request = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field FirstKeycode
	v.FirstKeycode = Keycode(r.Read1b())
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Count
	v.Count = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}
	return
}
func MappingNotifyEventWrite(w *Writer, v *MappingNotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Request
	w.Write1b(v.Request)
	n += 1

	// write field FirstKeycode
	w.Write1b(uint8(v.FirstKeycode))
	n += 1

	// write field Count
	w.Write1b(v.Count)
	n += 1

	// write field Pad1
	w.WritePad(1)
	n += 1
	return
}

func NewMappingNotifyEvent(data []byte) (*MappingNotifyEvent, error) {
	var ev MappingNotifyEvent
	r := NewReaderFromData(data)
	MappingNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *MappingNotifyEvent) String() string {
	return fmt.Sprintf("MappingNotifyEvent {Sequence: %v, Request: %v, FirstKeycode: %v, Count: %v}",
		v.Sequence, v.Request, v.FirstKeycode, v.Count)
}

// try event source MappingNotify
type _MappingNotifyEventHandler struct {
	mu  sync.Mutex
	cbs []func(*MappingNotifyEvent)
}

func (eh *_MappingNotifyEventHandler) Run(ge GenericEvent) {
	ev, err := NewMappingNotifyEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.mu.Lock()
	cbs := eh.cbs
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (_ *_MappingNotifyEventHandler) Detach(_ uint32) {}

func (eh *_MappingNotifyEventHandler) attach(cb func(ev *MappingNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs
	newCbs := make([]func(*MappingNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs = newCbs
	eh.mu.Unlock()
}

func ConnectMappingNotifyEvent(conn *Conn, cb func(ev *MappingNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[MappingNotifyEventCode]
	if !ok {
		newEh := &_MappingNotifyEventHandler{}
		conn.EventHandlers[MappingNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_MappingNotifyEventHandler).attach(cb)
}

// event name: ('xcb', 'GeGeneric')
// self.name: ('xcb', 'GeGeneric')
const GeGenericEventCode = 35

// not event copy
type GeGenericEvent struct {
	ResponseType uint8
	Extension    uint8
	Sequence     uint16
	Length       uint32
	EventType    uint16
	// Pad0 [22]uint8
}

func GeGenericEventRead(r *Reader, v *GeGenericEvent) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Extension
	v.Extension = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field EventType
	v.EventType = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(22)
	n += 22
	if r.Err() != nil {
		return
	}
	return
}
func GeGenericEventWrite(w *Writer, v *GeGenericEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Extension
	w.Write1b(v.Extension)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Length
	w.Write4b(v.Length)
	n += 4

	// write field EventType
	w.Write2b(v.EventType)
	n += 2

	// write field Pad0
	w.WritePad(22)
	n += 22
	return
}

func NewGeGenericEvent(data []byte) (*GeGenericEvent, error) {
	var ev GeGenericEvent
	r := NewReaderFromData(data)
	GeGenericEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *GeGenericEvent) String() string {
	return fmt.Sprintf("GeGenericEvent {Extension: %v, Sequence: %v, Length: %v, EventType: %v}",
		v.Extension, v.Sequence, v.Length, v.EventType)
}

// try event source GeGeneric
type _GeGenericEventHandler struct {
	mu  sync.Mutex
	cbs []func(*GeGenericEvent)
}

func (eh *_GeGenericEventHandler) Run(ge GenericEvent) {
	ev, err := NewGeGenericEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.mu.Lock()
	cbs := eh.cbs
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (_ *_GeGenericEventHandler) Detach(_ uint32) {}

func (eh *_GeGenericEventHandler) attach(cb func(ev *GeGenericEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs
	newCbs := make([]func(*GeGenericEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs = newCbs
	eh.mu.Unlock()
}

func ConnectGeGenericEvent(conn *Conn, cb func(ev *GeGenericEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[GeGenericEventCode]
	if !ok {
		newEh := &_GeGenericEventHandler{}
		conn.EventHandlers[GeGenericEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_GeGenericEventHandler).attach(cb)
}

// enum WindowClass
const (
	WindowClassCopyFromParent = 0
	WindowClassInputOutput    = 1
	WindowClassInputOnly      = 2
)

func WindowClassEnumToStr(v int) string {
	switch v {
	case 0:
		return "CopyFromParent"
	case 1:
		return "InputOutput"
	case 2:
		return "InputOnly"
	default:
		return "<Unknown>"
	}
}

// enum CW
const (
	CWBackPixmap       = 1
	CWBackPixel        = 2
	CWBorderPixmap     = 4
	CWBorderPixel      = 8
	CWBitGravity       = 16
	CWWinGravity       = 32
	CWBackingStore     = 64
	CWBackingPlanes    = 128
	CWBackingPixel     = 256
	CWOverrideRedirect = 512
	CWSaveUnder        = 1024
	CWEventMask        = 2048
	CWDontPropagate    = 4096
	CWColormap         = 8192
	CWCursor           = 16384
)

// enum BackPixmap
const (
	BackPixmapNone           = 0
	BackPixmapParentRelative = 1
)

// enum Gravity
const (
	GravityBitForget = 0
	GravityWinUnmap  = 0
	GravityNorthWest = 1
	GravityNorth     = 2
	GravityNorthEast = 3
	GravityWest      = 4
	GravityCenter    = 5
	GravityEast      = 6
	GravitySouthWest = 7
	GravitySouth     = 8
	GravitySouthEast = 9
	GravityStatic    = 10
)

func GravityEnumToStr(v int) string {
	switch v {
	case 0:
		return "BitForget or WinUnmap"
	case 1:
		return "NorthWest"
	case 2:
		return "North"
	case 3:
		return "NorthEast"
	case 4:
		return "West"
	case 5:
		return "Center"
	case 6:
		return "East"
	case 7:
		return "SouthWest"
	case 8:
		return "South"
	case 9:
		return "SouthEast"
	case 10:
		return "Static"
	default:
		return "<Unknown>"
	}
}

type CreateWindowValueList struct {
	BackgroundPixmap   Pixmap
	BackgroundPixel    uint32
	BorderPixmap       Pixmap
	BorderPixel        uint32
	BitGravity         uint32
	WinGravity         uint32
	BackingStore       uint32
	BackingPlanes      uint32
	BackingPixel       uint32
	OverrideRedirect   Bool32
	SaveUnder          Bool32
	EventMask          uint32
	DoNotPropogateMask uint32
	Colormap           Colormap
	Cursor             Cursor
}

// _go_switch_write SwitchType "xcb.CreateWindow.value_list"
func CreateWindowValueListSerialize(w *Writer, ValueMask uint32, _aux *CreateWindowValueList) {
	n := 0
	xgbPaddingOffset := 0
	alignTo := 0
	xgbPad := 0
	// switch expr: int(ValueMask)

	// case_field Field with type BitcaseType "xcb.CreateWindow.value_list.bitcase1"
	if (int(ValueMask) & CWBackPixmap) != 0 {
		// todo
		// field Field "background_pixmap" in BitcaseType "xcb.CreateWindow.value_list.bitcase1"
		w.Write4b(uint32(_aux.BackgroundPixmap))
		n += 4
		alignTo = int(unsafe.Alignof(Pixmap(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateWindow.value_list.bitcase2"
	if (int(ValueMask) & CWBackPixel) != 0 {
		// todo
		// field Field "background_pixel" in BitcaseType "xcb.CreateWindow.value_list.bitcase2"
		w.Write4b(_aux.BackgroundPixel)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateWindow.value_list.bitcase3"
	if (int(ValueMask) & CWBorderPixmap) != 0 {
		// todo
		// field Field "border_pixmap" in BitcaseType "xcb.CreateWindow.value_list.bitcase3"
		w.Write4b(uint32(_aux.BorderPixmap))
		n += 4
		alignTo = int(unsafe.Alignof(Pixmap(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateWindow.value_list.bitcase4"
	if (int(ValueMask) & CWBorderPixel) != 0 {
		// todo
		// field Field "border_pixel" in BitcaseType "xcb.CreateWindow.value_list.bitcase4"
		w.Write4b(_aux.BorderPixel)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateWindow.value_list.bitcase5"
	if (int(ValueMask) & CWBitGravity) != 0 {
		// todo
		// field Field "bit_gravity" in BitcaseType "xcb.CreateWindow.value_list.bitcase5"
		w.Write4b(_aux.BitGravity)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateWindow.value_list.bitcase6"
	if (int(ValueMask) & CWWinGravity) != 0 {
		// todo
		// field Field "win_gravity" in BitcaseType "xcb.CreateWindow.value_list.bitcase6"
		w.Write4b(_aux.WinGravity)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateWindow.value_list.bitcase7"
	if (int(ValueMask) & CWBackingStore) != 0 {
		// todo
		// field Field "backing_store" in BitcaseType "xcb.CreateWindow.value_list.bitcase7"
		w.Write4b(_aux.BackingStore)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateWindow.value_list.bitcase8"
	if (int(ValueMask) & CWBackingPlanes) != 0 {
		// todo
		// field Field "backing_planes" in BitcaseType "xcb.CreateWindow.value_list.bitcase8"
		w.Write4b(_aux.BackingPlanes)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateWindow.value_list.bitcase9"
	if (int(ValueMask) & CWBackingPixel) != 0 {
		// todo
		// field Field "backing_pixel" in BitcaseType "xcb.CreateWindow.value_list.bitcase9"
		w.Write4b(_aux.BackingPixel)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateWindow.value_list.bitcase10"
	if (int(ValueMask) & CWOverrideRedirect) != 0 {
		// todo
		// field Field "override_redirect" in BitcaseType "xcb.CreateWindow.value_list.bitcase10"
		w.Write4b(uint32(_aux.OverrideRedirect))
		n += 4
		alignTo = int(unsafe.Alignof(Bool32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateWindow.value_list.bitcase11"
	if (int(ValueMask) & CWSaveUnder) != 0 {
		// todo
		// field Field "save_under" in BitcaseType "xcb.CreateWindow.value_list.bitcase11"
		w.Write4b(uint32(_aux.SaveUnder))
		n += 4
		alignTo = int(unsafe.Alignof(Bool32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateWindow.value_list.bitcase12"
	if (int(ValueMask) & CWEventMask) != 0 {
		// todo
		// field Field "event_mask" in BitcaseType "xcb.CreateWindow.value_list.bitcase12"
		w.Write4b(_aux.EventMask)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateWindow.value_list.bitcase13"
	if (int(ValueMask) & CWDontPropagate) != 0 {
		// todo
		// field Field "do_not_propogate_mask" in BitcaseType "xcb.CreateWindow.value_list.bitcase13"
		w.Write4b(_aux.DoNotPropogateMask)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateWindow.value_list.bitcase14"
	if (int(ValueMask) & CWColormap) != 0 {
		// todo
		// field Field "colormap" in BitcaseType "xcb.CreateWindow.value_list.bitcase14"
		w.Write4b(uint32(_aux.Colormap))
		n += 4
		alignTo = int(unsafe.Alignof(Colormap(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateWindow.value_list.bitcase15"
	if (int(ValueMask) & CWCursor) != 0 {
		// todo
		// field Field "cursor" in BitcaseType "xcb.CreateWindow.value_list.bitcase15"
		w.Write4b(uint32(_aux.Cursor))
		n += 4
		alignTo = int(unsafe.Alignof(Cursor(0)))
	}
	// end a case_field

	/* insert padding */
	xgbPad = -(n + xgbPaddingOffset) & (alignTo - 1)
	w.WritePad(xgbPad)
	xgbPad = 0
	n = 0
	xgbPaddingOffset = 0
}

const CreateWindowOpcode = 1

func CreateWindowRequestWrite(w *Writer, Depth uint8, Wid Window, Parent Window, X int16, Y int16, Width uint16, Height uint16, BorderWidth uint16, Class uint16, Visual VisualID, ValueMask uint32, ValueList *CreateWindowValueList) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Depth
	w.Write1b(Depth)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Wid
	w.Write4b(uint32(Wid))
	n += 4

	// write wire field Parent
	w.Write4b(uint32(Parent))
	n += 4

	// write wire field X
	w.Write2b(uint16(X))
	n += 2

	// write wire field Y
	w.Write2b(uint16(Y))
	n += 2

	// write wire field Width
	w.Write2b(Width)
	n += 2

	// write wire field Height
	w.Write2b(Height)
	n += 2

	// write wire field BorderWidth
	w.Write2b(BorderWidth)
	n += 2

	// write wire field Class
	w.Write2b(Class)
	n += 2

	// write wire field Visual
	w.Write4b(uint32(Visual))
	n += 4

	// write wire field ValueMask
	w.Write4b(ValueMask)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field ValueList
	// switch serialize
	CreateWindowValueListSerialize(w, ValueMask, ValueList)
}

func CreateWindow(conn *Conn, Depth uint8, Wid Window, Parent Window, X int16, Y int16, Width uint16, Height uint16, BorderWidth uint16, Class uint16, Visual VisualID, ValueMask uint32, ValueList *CreateWindowValueList) VoidCookie {
	w := NewWriter()
	CreateWindowRequestWrite(w, Depth, Wid, Parent, X, Y, Width, Height, BorderWidth, Class, Visual, ValueMask, ValueList)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CreateWindowOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func CreateWindowChecked(conn *Conn, Depth uint8, Wid Window, Parent Window, X int16, Y int16, Width uint16, Height uint16, BorderWidth uint16, Class uint16, Visual VisualID, ValueMask uint32, ValueList *CreateWindowValueList) VoidCookie {
	w := NewWriter()
	CreateWindowRequestWrite(w, Depth, Wid, Parent, X, Y, Width, Height, BorderWidth, Class, Visual, ValueMask, ValueList)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CreateWindowOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

type ChangeWindowAttributesValueList struct {
	BackgroundPixmap   Pixmap
	BackgroundPixel    uint32
	BorderPixmap       Pixmap
	BorderPixel        uint32
	BitGravity         uint32
	WinGravity         uint32
	BackingStore       uint32
	BackingPlanes      uint32
	BackingPixel       uint32
	OverrideRedirect   Bool32
	SaveUnder          Bool32
	EventMask          uint32
	DoNotPropogateMask uint32
	Colormap           Colormap
	Cursor             Cursor
}

// _go_switch_write SwitchType "xcb.ChangeWindowAttributes.value_list"
func ChangeWindowAttributesValueListSerialize(w *Writer, ValueMask uint32, _aux *ChangeWindowAttributesValueList) {
	n := 0
	xgbPaddingOffset := 0
	alignTo := 0
	xgbPad := 0
	// switch expr: int(ValueMask)

	// case_field Field with type BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase1"
	if (int(ValueMask) & CWBackPixmap) != 0 {
		// todo
		// field Field "background_pixmap" in BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase1"
		w.Write4b(uint32(_aux.BackgroundPixmap))
		n += 4
		alignTo = int(unsafe.Alignof(Pixmap(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase2"
	if (int(ValueMask) & CWBackPixel) != 0 {
		// todo
		// field Field "background_pixel" in BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase2"
		w.Write4b(_aux.BackgroundPixel)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase3"
	if (int(ValueMask) & CWBorderPixmap) != 0 {
		// todo
		// field Field "border_pixmap" in BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase3"
		w.Write4b(uint32(_aux.BorderPixmap))
		n += 4
		alignTo = int(unsafe.Alignof(Pixmap(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase4"
	if (int(ValueMask) & CWBorderPixel) != 0 {
		// todo
		// field Field "border_pixel" in BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase4"
		w.Write4b(_aux.BorderPixel)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase5"
	if (int(ValueMask) & CWBitGravity) != 0 {
		// todo
		// field Field "bit_gravity" in BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase5"
		w.Write4b(_aux.BitGravity)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase6"
	if (int(ValueMask) & CWWinGravity) != 0 {
		// todo
		// field Field "win_gravity" in BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase6"
		w.Write4b(_aux.WinGravity)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase7"
	if (int(ValueMask) & CWBackingStore) != 0 {
		// todo
		// field Field "backing_store" in BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase7"
		w.Write4b(_aux.BackingStore)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase8"
	if (int(ValueMask) & CWBackingPlanes) != 0 {
		// todo
		// field Field "backing_planes" in BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase8"
		w.Write4b(_aux.BackingPlanes)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase9"
	if (int(ValueMask) & CWBackingPixel) != 0 {
		// todo
		// field Field "backing_pixel" in BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase9"
		w.Write4b(_aux.BackingPixel)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase10"
	if (int(ValueMask) & CWOverrideRedirect) != 0 {
		// todo
		// field Field "override_redirect" in BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase10"
		w.Write4b(uint32(_aux.OverrideRedirect))
		n += 4
		alignTo = int(unsafe.Alignof(Bool32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase11"
	if (int(ValueMask) & CWSaveUnder) != 0 {
		// todo
		// field Field "save_under" in BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase11"
		w.Write4b(uint32(_aux.SaveUnder))
		n += 4
		alignTo = int(unsafe.Alignof(Bool32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase12"
	if (int(ValueMask) & CWEventMask) != 0 {
		// todo
		// field Field "event_mask" in BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase12"
		w.Write4b(_aux.EventMask)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase13"
	if (int(ValueMask) & CWDontPropagate) != 0 {
		// todo
		// field Field "do_not_propogate_mask" in BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase13"
		w.Write4b(_aux.DoNotPropogateMask)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase14"
	if (int(ValueMask) & CWColormap) != 0 {
		// todo
		// field Field "colormap" in BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase14"
		w.Write4b(uint32(_aux.Colormap))
		n += 4
		alignTo = int(unsafe.Alignof(Colormap(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase15"
	if (int(ValueMask) & CWCursor) != 0 {
		// todo
		// field Field "cursor" in BitcaseType "xcb.ChangeWindowAttributes.value_list.bitcase15"
		w.Write4b(uint32(_aux.Cursor))
		n += 4
		alignTo = int(unsafe.Alignof(Cursor(0)))
	}
	// end a case_field

	/* insert padding */
	xgbPad = -(n + xgbPaddingOffset) & (alignTo - 1)
	w.WritePad(xgbPad)
	xgbPad = 0
	n = 0
	xgbPaddingOffset = 0
}

const ChangeWindowAttributesOpcode = 2

func ChangeWindowAttributesRequestWrite(w *Writer, Window Window, ValueMask uint32, ValueList *ChangeWindowAttributesValueList) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4

	// write wire field ValueMask
	w.Write4b(ValueMask)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field ValueList
	// switch serialize
	ChangeWindowAttributesValueListSerialize(w, ValueMask, ValueList)
}

func ChangeWindowAttributes(conn *Conn, Window Window, ValueMask uint32, ValueList *ChangeWindowAttributesValueList) VoidCookie {
	w := NewWriter()
	ChangeWindowAttributesRequestWrite(w, Window, ValueMask, ValueList)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ChangeWindowAttributesOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func ChangeWindowAttributesChecked(conn *Conn, Window Window, ValueMask uint32, ValueList *ChangeWindowAttributesValueList) VoidCookie {
	w := NewWriter()
	ChangeWindowAttributesRequestWrite(w, Window, ValueMask, ValueList)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ChangeWindowAttributesOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum MapState
const (
	MapStateUnmapped   = 0
	MapStateUnviewable = 1
	MapStateViewable   = 2
)

func MapStateEnumToStr(v int) string {
	switch v {
	case 0:
		return "Unmapped"
	case 1:
		return "Unviewable"
	case 2:
		return "Viewable"
	default:
		return "<Unknown>"
	}
}

const GetWindowAttributesOpcode = 3

func GetWindowAttributesRequestWrite(w *Writer, Window Window) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetWindowAttributesReply struct {
	ResponseType       uint8
	BackingStore       uint8
	Sequence           uint16
	Length             uint32
	Visual             VisualID
	Class              uint16
	BitGravity         uint8
	WinGravity         uint8
	BackingPlanes      uint32
	BackingPixel       uint32
	SaveUnder          uint8
	MapIsInstalled     uint8
	MapState           uint8
	OverrideRedirect   uint8
	Colormap           Colormap
	AllEventMasks      uint32
	YourEventMask      uint32
	DoNotPropagateMask uint16
	// Pad0 [2]uint8
}

func GetWindowAttributesReplyRead(r *Reader, v *GetWindowAttributesReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field BackingStore
	v.BackingStore = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Visual
	v.Visual = VisualID(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Class
	v.Class = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field BitGravity
	v.BitGravity = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field WinGravity
	v.WinGravity = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field BackingPlanes
	v.BackingPlanes = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field BackingPixel
	v.BackingPixel = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field SaveUnder
	v.SaveUnder = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field MapIsInstalled
	v.MapIsInstalled = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field MapState
	v.MapState = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field OverrideRedirect
	v.OverrideRedirect = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Colormap
	v.Colormap = Colormap(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field AllEventMasks
	v.AllEventMasks = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field YourEventMask
	v.YourEventMask = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field DoNotPropagateMask
	v.DoNotPropagateMask = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(2)
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

type GetWindowAttributesCookie uint64

func (cookie GetWindowAttributesCookie) Reply(conn *Conn) (*GetWindowAttributesReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply GetWindowAttributesReply
	GetWindowAttributesReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetWindowAttributes(conn *Conn, Window Window) GetWindowAttributesCookie {
	w := NewWriter()
	GetWindowAttributesRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetWindowAttributesOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return GetWindowAttributesCookie(seq)
}

func GetWindowAttributesUnchecked(conn *Conn, Window Window) GetWindowAttributesCookie {
	w := NewWriter()
	GetWindowAttributesRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetWindowAttributesOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetWindowAttributesCookie(seq)
}

const DestroyWindowOpcode = 4

func DestroyWindowRequestWrite(w *Writer, Window Window) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func DestroyWindow(conn *Conn, Window Window) VoidCookie {
	w := NewWriter()
	DestroyWindowRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: DestroyWindowOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func DestroyWindowChecked(conn *Conn, Window Window) VoidCookie {
	w := NewWriter()
	DestroyWindowRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: DestroyWindowOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const DestroySubwindowsOpcode = 5

func DestroySubwindowsRequestWrite(w *Writer, Window Window) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func DestroySubwindows(conn *Conn, Window Window) VoidCookie {
	w := NewWriter()
	DestroySubwindowsRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: DestroySubwindowsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func DestroySubwindowsChecked(conn *Conn, Window Window) VoidCookie {
	w := NewWriter()
	DestroySubwindowsRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: DestroySubwindowsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum SetMode
const (
	SetModeInsert = 0
	SetModeDelete = 1
)

const ChangeSaveSetOpcode = 6

func ChangeSaveSetRequestWrite(w *Writer, Mode uint8, Window Window) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Mode
	w.Write1b(Mode)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func ChangeSaveSet(conn *Conn, Mode uint8, Window Window) VoidCookie {
	w := NewWriter()
	ChangeSaveSetRequestWrite(w, Mode, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ChangeSaveSetOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func ChangeSaveSetChecked(conn *Conn, Mode uint8, Window Window) VoidCookie {
	w := NewWriter()
	ChangeSaveSetRequestWrite(w, Mode, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ChangeSaveSetOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const ReparentWindowOpcode = 7

func ReparentWindowRequestWrite(w *Writer, Window Window, Parent Window, X int16, Y int16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4

	// write wire field Parent
	w.Write4b(uint32(Parent))
	n += 4

	// write wire field X
	w.Write2b(uint16(X))
	n += 2

	// write wire field Y
	w.Write2b(uint16(Y))
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func ReparentWindow(conn *Conn, Window Window, Parent Window, X int16, Y int16) VoidCookie {
	w := NewWriter()
	ReparentWindowRequestWrite(w, Window, Parent, X, Y)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ReparentWindowOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func ReparentWindowChecked(conn *Conn, Window Window, Parent Window, X int16, Y int16) VoidCookie {
	w := NewWriter()
	ReparentWindowRequestWrite(w, Window, Parent, X, Y)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ReparentWindowOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const MapWindowOpcode = 8

func MapWindowRequestWrite(w *Writer, Window Window) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func MapWindow(conn *Conn, Window Window) VoidCookie {
	w := NewWriter()
	MapWindowRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: MapWindowOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func MapWindowChecked(conn *Conn, Window Window) VoidCookie {
	w := NewWriter()
	MapWindowRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: MapWindowOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const MapSubwindowsOpcode = 9

func MapSubwindowsRequestWrite(w *Writer, Window Window) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func MapSubwindows(conn *Conn, Window Window) VoidCookie {
	w := NewWriter()
	MapSubwindowsRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: MapSubwindowsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func MapSubwindowsChecked(conn *Conn, Window Window) VoidCookie {
	w := NewWriter()
	MapSubwindowsRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: MapSubwindowsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const UnmapWindowOpcode = 10

func UnmapWindowRequestWrite(w *Writer, Window Window) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func UnmapWindow(conn *Conn, Window Window) VoidCookie {
	w := NewWriter()
	UnmapWindowRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: UnmapWindowOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func UnmapWindowChecked(conn *Conn, Window Window) VoidCookie {
	w := NewWriter()
	UnmapWindowRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: UnmapWindowOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const UnmapSubwindowsOpcode = 11

func UnmapSubwindowsRequestWrite(w *Writer, Window Window) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func UnmapSubwindows(conn *Conn, Window Window) VoidCookie {
	w := NewWriter()
	UnmapSubwindowsRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: UnmapSubwindowsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func UnmapSubwindowsChecked(conn *Conn, Window Window) VoidCookie {
	w := NewWriter()
	UnmapSubwindowsRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: UnmapSubwindowsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum ConfigWindow
const (
	ConfigWindowX           = 1
	ConfigWindowY           = 2
	ConfigWindowWidth       = 4
	ConfigWindowHeight      = 8
	ConfigWindowBorderWidth = 16
	ConfigWindowSibling     = 32
	ConfigWindowStackMode   = 64
)

func ConfigWindowEnumToStr(v int) string {
	strv := []string{}
	switch {
	case v&ConfigWindowX > 0:
		strv = append(strv, "X")
	case v&ConfigWindowY > 0:
		strv = append(strv, "Y")
	case v&ConfigWindowWidth > 0:
		strv = append(strv, "Width")
	case v&ConfigWindowHeight > 0:
		strv = append(strv, "Height")
	case v&ConfigWindowBorderWidth > 0:
		strv = append(strv, "BorderWidth")
	case v&ConfigWindowSibling > 0:
		strv = append(strv, "Sibling")
	case v&ConfigWindowStackMode > 0:
		strv = append(strv, "StackMode")
	}
	return "[" + strings.Join(strv, "|") + "]"
}

// enum StackMode
const (
	StackModeAbove    = 0
	StackModeBelow    = 1
	StackModeTopIf    = 2
	StackModeBottomIf = 3
	StackModeOpposite = 4
)

func StackModeEnumToStr(v int) string {
	switch v {
	case 0:
		return "Above"
	case 1:
		return "Below"
	case 2:
		return "TopIf"
	case 3:
		return "BottomIf"
	case 4:
		return "Opposite"
	default:
		return "<Unknown>"
	}
}

type ConfigureWindowValueList struct {
	X           int32
	Y           int32
	Width       uint32
	Height      uint32
	BorderWidth uint32
	Sibling     Window
	StackMode   uint32
}

// _go_switch_write SwitchType "xcb.ConfigureWindow.value_list"
func ConfigureWindowValueListSerialize(w *Writer, ValueMask uint16, _aux *ConfigureWindowValueList) {
	n := 0
	xgbPaddingOffset := 0
	alignTo := 0
	xgbPad := 0
	// switch expr: int(ValueMask)

	// case_field Field with type BitcaseType "xcb.ConfigureWindow.value_list.bitcase1"
	if (int(ValueMask) & ConfigWindowX) != 0 {
		// todo
		// field Field "x" in BitcaseType "xcb.ConfigureWindow.value_list.bitcase1"
		w.Write4b(uint32(_aux.X))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ConfigureWindow.value_list.bitcase2"
	if (int(ValueMask) & ConfigWindowY) != 0 {
		// todo
		// field Field "y" in BitcaseType "xcb.ConfigureWindow.value_list.bitcase2"
		w.Write4b(uint32(_aux.Y))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ConfigureWindow.value_list.bitcase3"
	if (int(ValueMask) & ConfigWindowWidth) != 0 {
		// todo
		// field Field "width" in BitcaseType "xcb.ConfigureWindow.value_list.bitcase3"
		w.Write4b(_aux.Width)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ConfigureWindow.value_list.bitcase4"
	if (int(ValueMask) & ConfigWindowHeight) != 0 {
		// todo
		// field Field "height" in BitcaseType "xcb.ConfigureWindow.value_list.bitcase4"
		w.Write4b(_aux.Height)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ConfigureWindow.value_list.bitcase5"
	if (int(ValueMask) & ConfigWindowBorderWidth) != 0 {
		// todo
		// field Field "border_width" in BitcaseType "xcb.ConfigureWindow.value_list.bitcase5"
		w.Write4b(_aux.BorderWidth)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ConfigureWindow.value_list.bitcase6"
	if (int(ValueMask) & ConfigWindowSibling) != 0 {
		// todo
		// field Field "sibling" in BitcaseType "xcb.ConfigureWindow.value_list.bitcase6"
		w.Write4b(uint32(_aux.Sibling))
		n += 4
		alignTo = int(unsafe.Alignof(Window(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ConfigureWindow.value_list.bitcase7"
	if (int(ValueMask) & ConfigWindowStackMode) != 0 {
		// todo
		// field Field "stack_mode" in BitcaseType "xcb.ConfigureWindow.value_list.bitcase7"
		w.Write4b(_aux.StackMode)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	/* insert padding */
	xgbPad = -(n + xgbPaddingOffset) & (alignTo - 1)
	w.WritePad(xgbPad)
	xgbPad = 0
	n = 0
	xgbPaddingOffset = 0
}

const ConfigureWindowOpcode = 12

func ConfigureWindowRequestWrite(w *Writer, Window Window, ValueMask uint16, ValueList *ConfigureWindowValueList) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4

	// write wire field ValueMask
	w.Write2b(ValueMask)
	n += 2

	// write wire field Pad1
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field ValueList
	// switch serialize
	ConfigureWindowValueListSerialize(w, ValueMask, ValueList)
}

func ConfigureWindow(conn *Conn, Window Window, ValueMask uint16, ValueList *ConfigureWindowValueList) VoidCookie {
	w := NewWriter()
	ConfigureWindowRequestWrite(w, Window, ValueMask, ValueList)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ConfigureWindowOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func ConfigureWindowChecked(conn *Conn, Window Window, ValueMask uint16, ValueList *ConfigureWindowValueList) VoidCookie {
	w := NewWriter()
	ConfigureWindowRequestWrite(w, Window, ValueMask, ValueList)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ConfigureWindowOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum Circulate
const (
	CirculateRaiseLowest  = 0
	CirculateLowerHighest = 1
)

const CirculateWindowOpcode = 13

func CirculateWindowRequestWrite(w *Writer, Direction uint8, Window Window) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Direction
	w.Write1b(Direction)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CirculateWindow(conn *Conn, Direction uint8, Window Window) VoidCookie {
	w := NewWriter()
	CirculateWindowRequestWrite(w, Direction, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CirculateWindowOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func CirculateWindowChecked(conn *Conn, Direction uint8, Window Window) VoidCookie {
	w := NewWriter()
	CirculateWindowRequestWrite(w, Direction, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CirculateWindowOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const GetGeometryOpcode = 14

func GetGeometryRequestWrite(w *Writer, Drawable Drawable) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetGeometryReply struct {
	ResponseType uint8
	Depth        uint8
	Sequence     uint16
	Length       uint32
	Root         Window
	X            int16
	Y            int16
	Width        uint16
	Height       uint16
	BorderWidth  uint16
	// Pad0 [2]uint8
}

func GetGeometryReplyRead(r *Reader, v *GetGeometryReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Depth
	v.Depth = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Root
	v.Root = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field X
	v.X = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Y
	v.Y = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Width
	v.Width = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Height
	v.Height = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field BorderWidth
	v.BorderWidth = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(2)
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

type GetGeometryCookie uint64

func (cookie GetGeometryCookie) Reply(conn *Conn) (*GetGeometryReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply GetGeometryReply
	GetGeometryReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetGeometry(conn *Conn, Drawable Drawable) GetGeometryCookie {
	w := NewWriter()
	GetGeometryRequestWrite(w, Drawable)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetGeometryOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return GetGeometryCookie(seq)
}

func GetGeometryUnchecked(conn *Conn, Drawable Drawable) GetGeometryCookie {
	w := NewWriter()
	GetGeometryRequestWrite(w, Drawable)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetGeometryOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetGeometryCookie(seq)
}

const QueryTreeOpcode = 15

func QueryTreeRequestWrite(w *Writer, Window Window) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type QueryTreeReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence    uint16
	Length      uint32
	Root        Window
	Parent      Window
	ChildrenLen uint16
	// Pad1 [14]uint8
	Children []Window
}

func QueryTreeReplyRead(r *Reader, v *QueryTreeReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Root
	v.Root = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Parent
	v.Parent = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field ChildrenLen
	v.ChildrenLen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(14)
	n += 14
	if r.Err() != nil {
		return
	}

	// read field Children
	{
		dataLen := int(int(v.ChildrenLen))
		data := make([]Window, dataLen)
		for i := 0; i < dataLen; i++ {
			data[i] = Window(r.Read4b())
			if r.Err() != nil {
				return
			}
		}
		blockLen += dataLen * 4
		n += blockLen
		v.Children = data
	}
	alignTo = int(unsafe.Alignof(Window(0)))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type QueryTreeCookie uint64

func (cookie QueryTreeCookie) Reply(conn *Conn) (*QueryTreeReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply QueryTreeReply
	QueryTreeReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryTree(conn *Conn, Window Window) QueryTreeCookie {
	w := NewWriter()
	QueryTreeRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: QueryTreeOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return QueryTreeCookie(seq)
}

func QueryTreeUnchecked(conn *Conn, Window Window) QueryTreeCookie {
	w := NewWriter()
	QueryTreeRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: QueryTreeOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return QueryTreeCookie(seq)
}

const InternAtomOpcode = 16

func InternAtomRequestWrite(w *Writer, OnlyIfExists uint8, NameLen uint16, Name string) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field OnlyIfExists
	w.Write1b(OnlyIfExists)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field NameLen
	w.Write2b(NameLen)
	n += 2

	// write wire field Pad0
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Name
	w.WriteString(Name)
	n += len(Name)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type InternAtomReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Length   uint32
	Atom     Atom
}

func InternAtomReplyRead(r *Reader, v *InternAtomReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Atom
	v.Atom = Atom(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}
	return
}

type InternAtomCookie uint64

func (cookie InternAtomCookie) Reply(conn *Conn) (*InternAtomReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply InternAtomReply
	InternAtomReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func InternAtom(conn *Conn, OnlyIfExists uint8, NameLen uint16, Name string) InternAtomCookie {
	w := NewWriter()
	InternAtomRequestWrite(w, OnlyIfExists, NameLen, Name)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: InternAtomOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return InternAtomCookie(seq)
}

func InternAtomUnchecked(conn *Conn, OnlyIfExists uint8, NameLen uint16, Name string) InternAtomCookie {
	w := NewWriter()
	InternAtomRequestWrite(w, OnlyIfExists, NameLen, Name)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: InternAtomOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return InternAtomCookie(seq)
}

const GetAtomNameOpcode = 17

func GetAtomNameRequestWrite(w *Writer, Atom Atom) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Atom
	w.Write4b(uint32(Atom))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetAtomNameReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Length   uint32
	NameLen  uint16
	// Pad1 [22]uint8
	Name string
}

func GetAtomNameReplyRead(r *Reader, v *GetAtomNameReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field NameLen
	v.NameLen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(22)
	n += 22
	if r.Err() != nil {
		return
	}

	// read field Name
	{
		dataLen := int(int(v.NameLen))
		data := r.ReadString(dataLen)
		if r.Err() != nil {
			return
		}
		blockLen += dataLen
		n += blockLen
		v.Name = data
	}
	alignTo = int(unsafe.Alignof(byte(0)))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type GetAtomNameCookie uint64

func (cookie GetAtomNameCookie) Reply(conn *Conn) (*GetAtomNameReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply GetAtomNameReply
	GetAtomNameReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetAtomName(conn *Conn, Atom Atom) GetAtomNameCookie {
	w := NewWriter()
	GetAtomNameRequestWrite(w, Atom)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetAtomNameOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return GetAtomNameCookie(seq)
}

func GetAtomNameUnchecked(conn *Conn, Atom Atom) GetAtomNameCookie {
	w := NewWriter()
	GetAtomNameRequestWrite(w, Atom)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetAtomNameOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetAtomNameCookie(seq)
}

// enum PropMode
const (
	PropModeReplace = 0
	PropModePrepend = 1
	PropModeAppend  = 2
)

const ChangePropertyOpcode = 18

func ChangePropertyRequestWrite(w *Writer, Mode uint8, Window Window, Property Atom, Type Atom, Format uint8, DataLen uint32, Data []byte) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Mode
	w.Write1b(Mode)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4

	// write wire field Property
	w.Write4b(uint32(Property))
	n += 4

	// write wire field Type
	w.Write4b(uint32(Type))
	n += 4

	// write wire field Format
	w.Write1b(Format)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3

	// write wire field DataLen
	w.Write4b(DataLen)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Data
	{
		_dataLen := ((int(DataLen) * int(Format)) / 8)
		for i := 0; i < _dataLen; i++ {
			w.Write1b(uint8(Data[i]))
		}
		n += _dataLen * 1
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func ChangeProperty(conn *Conn, Mode uint8, Window Window, Property Atom, Type Atom, Format uint8, DataLen uint32, Data []byte) VoidCookie {
	w := NewWriter()
	ChangePropertyRequestWrite(w, Mode, Window, Property, Type, Format, DataLen, Data)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ChangePropertyOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func ChangePropertyChecked(conn *Conn, Mode uint8, Window Window, Property Atom, Type Atom, Format uint8, DataLen uint32, Data []byte) VoidCookie {
	w := NewWriter()
	ChangePropertyRequestWrite(w, Mode, Window, Property, Type, Format, DataLen, Data)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ChangePropertyOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const DeletePropertyOpcode = 19

func DeletePropertyRequestWrite(w *Writer, Window Window, Property Atom) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4

	// write wire field Property
	w.Write4b(uint32(Property))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func DeleteProperty(conn *Conn, Window Window, Property Atom) VoidCookie {
	w := NewWriter()
	DeletePropertyRequestWrite(w, Window, Property)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: DeletePropertyOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func DeletePropertyChecked(conn *Conn, Window Window, Property Atom) VoidCookie {
	w := NewWriter()
	DeletePropertyRequestWrite(w, Window, Property)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: DeletePropertyOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum GetPropertyType
const (
	GetPropertyTypeAny = 0
)

const GetPropertyOpcode = 20

func GetPropertyRequestWrite(w *Writer, Delete uint8, Window Window, Property Atom, Type Atom, LongOffset uint32, LongLength uint32) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Delete
	w.Write1b(Delete)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4

	// write wire field Property
	w.Write4b(uint32(Property))
	n += 4

	// write wire field Type
	w.Write4b(uint32(Type))
	n += 4

	// write wire field LongOffset
	w.Write4b(LongOffset)
	n += 4

	// write wire field LongLength
	w.Write4b(LongLength)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetPropertyReply struct {
	ResponseType uint8
	Format       uint8
	Sequence     uint16
	Length       uint32
	Type         Atom
	BytesAfter   uint32
	ValueLen     uint32
	// Pad0 [12]uint8
	Value []byte
}

func GetPropertyReplyRead(r *Reader, v *GetPropertyReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Format
	v.Format = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Type
	v.Type = Atom(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field BytesAfter
	v.BytesAfter = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field ValueLen
	v.ValueLen = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(12)
	n += 12
	if r.Err() != nil {
		return
	}

	// read field Value
	{
		dataLen := int((int(v.ValueLen) * (int(v.Format) / 8)))
		data := r.ReadBytes(dataLen)
		if r.Err() != nil {
			return
		}
		blockLen += dataLen
		n += blockLen
		v.Value = data
	}
	alignTo = int(unsafe.Alignof(byte(0)))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type GetPropertyCookie uint64

func (cookie GetPropertyCookie) Reply(conn *Conn) (*GetPropertyReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply GetPropertyReply
	GetPropertyReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetProperty(conn *Conn, Delete uint8, Window Window, Property Atom, Type Atom, LongOffset uint32, LongLength uint32) GetPropertyCookie {
	w := NewWriter()
	GetPropertyRequestWrite(w, Delete, Window, Property, Type, LongOffset, LongLength)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetPropertyOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return GetPropertyCookie(seq)
}

func GetPropertyUnchecked(conn *Conn, Delete uint8, Window Window, Property Atom, Type Atom, LongOffset uint32, LongLength uint32) GetPropertyCookie {
	w := NewWriter()
	GetPropertyRequestWrite(w, Delete, Window, Property, Type, LongOffset, LongLength)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetPropertyOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetPropertyCookie(seq)
}

const ListPropertiesOpcode = 21

func ListPropertiesRequestWrite(w *Writer, Window Window) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type ListPropertiesReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Length   uint32
	AtomsLen uint16
	// Pad1 [22]uint8
	Atoms []Atom
}

func ListPropertiesReplyRead(r *Reader, v *ListPropertiesReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field AtomsLen
	v.AtomsLen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(22)
	n += 22
	if r.Err() != nil {
		return
	}

	// read field Atoms
	{
		dataLen := int(int(v.AtomsLen))
		data := make([]Atom, dataLen)
		for i := 0; i < dataLen; i++ {
			data[i] = Atom(r.Read4b())
			if r.Err() != nil {
				return
			}
		}
		blockLen += dataLen * 4
		n += blockLen
		v.Atoms = data
	}
	alignTo = int(unsafe.Alignof(Atom(0)))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type ListPropertiesCookie uint64

func (cookie ListPropertiesCookie) Reply(conn *Conn) (*ListPropertiesReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply ListPropertiesReply
	ListPropertiesReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func ListProperties(conn *Conn, Window Window) ListPropertiesCookie {
	w := NewWriter()
	ListPropertiesRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: ListPropertiesOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return ListPropertiesCookie(seq)
}

func ListPropertiesUnchecked(conn *Conn, Window Window) ListPropertiesCookie {
	w := NewWriter()
	ListPropertiesRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: ListPropertiesOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return ListPropertiesCookie(seq)
}

const SetSelectionOwnerOpcode = 22

func SetSelectionOwnerRequestWrite(w *Writer, Owner Window, Selection Atom, Time Timestamp) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Owner
	w.Write4b(uint32(Owner))
	n += 4

	// write wire field Selection
	w.Write4b(uint32(Selection))
	n += 4

	// write wire field Time
	w.Write4b(uint32(Time))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SetSelectionOwner(conn *Conn, Owner Window, Selection Atom, Time Timestamp) VoidCookie {
	w := NewWriter()
	SetSelectionOwnerRequestWrite(w, Owner, Selection, Time)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: SetSelectionOwnerOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func SetSelectionOwnerChecked(conn *Conn, Owner Window, Selection Atom, Time Timestamp) VoidCookie {
	w := NewWriter()
	SetSelectionOwnerRequestWrite(w, Owner, Selection, Time)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: SetSelectionOwnerOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const GetSelectionOwnerOpcode = 23

func GetSelectionOwnerRequestWrite(w *Writer, Selection Atom) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Selection
	w.Write4b(uint32(Selection))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetSelectionOwnerReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Length   uint32
	Owner    Window
}

func GetSelectionOwnerReplyRead(r *Reader, v *GetSelectionOwnerReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Owner
	v.Owner = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}
	return
}

type GetSelectionOwnerCookie uint64

func (cookie GetSelectionOwnerCookie) Reply(conn *Conn) (*GetSelectionOwnerReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply GetSelectionOwnerReply
	GetSelectionOwnerReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetSelectionOwner(conn *Conn, Selection Atom) GetSelectionOwnerCookie {
	w := NewWriter()
	GetSelectionOwnerRequestWrite(w, Selection)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetSelectionOwnerOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return GetSelectionOwnerCookie(seq)
}

func GetSelectionOwnerUnchecked(conn *Conn, Selection Atom) GetSelectionOwnerCookie {
	w := NewWriter()
	GetSelectionOwnerRequestWrite(w, Selection)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetSelectionOwnerOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetSelectionOwnerCookie(seq)
}

const ConvertSelectionOpcode = 24

func ConvertSelectionRequestWrite(w *Writer, Requestor Window, Selection Atom, Target Atom, Property Atom, Time Timestamp) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Requestor
	w.Write4b(uint32(Requestor))
	n += 4

	// write wire field Selection
	w.Write4b(uint32(Selection))
	n += 4

	// write wire field Target
	w.Write4b(uint32(Target))
	n += 4

	// write wire field Property
	w.Write4b(uint32(Property))
	n += 4

	// write wire field Time
	w.Write4b(uint32(Time))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func ConvertSelection(conn *Conn, Requestor Window, Selection Atom, Target Atom, Property Atom, Time Timestamp) VoidCookie {
	w := NewWriter()
	ConvertSelectionRequestWrite(w, Requestor, Selection, Target, Property, Time)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ConvertSelectionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func ConvertSelectionChecked(conn *Conn, Requestor Window, Selection Atom, Target Atom, Property Atom, Time Timestamp) VoidCookie {
	w := NewWriter()
	ConvertSelectionRequestWrite(w, Requestor, Selection, Target, Property, Time)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ConvertSelectionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum SendEventDest
const (
	SendEventDestPointerWindow = 0
	SendEventDestItemFocus     = 1
)

const SendEventOpcode = 25

func SendEventRequestWrite(w *Writer, Propagate uint8, Destination Window, EventMask uint32, Event []byte) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Propagate
	w.Write1b(Propagate)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Destination
	w.Write4b(uint32(Destination))
	n += 4

	// write wire field EventMask
	w.Write4b(EventMask)
	n += 4

	// write wire field Event
	w.WriteNBytes(32, Event)
	n += 32
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SendEvent(conn *Conn, Propagate uint8, Destination Window, EventMask uint32, Event []byte) VoidCookie {
	w := NewWriter()
	SendEventRequestWrite(w, Propagate, Destination, EventMask, Event)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: SendEventOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func SendEventChecked(conn *Conn, Propagate uint8, Destination Window, EventMask uint32, Event []byte) VoidCookie {
	w := NewWriter()
	SendEventRequestWrite(w, Propagate, Destination, EventMask, Event)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: SendEventOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum GrabMode
const (
	GrabModeSync  = 0
	GrabModeAsync = 1
)

// enum GrabStatus
const (
	GrabStatusSuccess        = 0
	GrabStatusAlreadyGrabbed = 1
	GrabStatusInvalidTime    = 2
	GrabStatusNotViewable    = 3
	GrabStatusFrozen         = 4
)

func GrabStatusEnumToStr(v int) string {
	switch v {
	case 0:
		return "Success"
	case 1:
		return "AlreadyGrabbed"
	case 2:
		return "InvalidTime"
	case 3:
		return "NotViewable"
	case 4:
		return "Frozen"
	default:
		return "<Unknown>"
	}
}

// enum Cursor
const (
	CursorNone = 0
)

const GrabPointerOpcode = 26

func GrabPointerRequestWrite(w *Writer, OwnerEvents uint8, GrabWindow Window, EventMask uint16, PointerMode uint8, KeyboardMode uint8, ConfineTo Window, Cursor Cursor, Time Timestamp) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field OwnerEvents
	w.Write1b(OwnerEvents)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field GrabWindow
	w.Write4b(uint32(GrabWindow))
	n += 4

	// write wire field EventMask
	w.Write2b(EventMask)
	n += 2

	// write wire field PointerMode
	w.Write1b(PointerMode)
	n += 1

	// write wire field KeyboardMode
	w.Write1b(KeyboardMode)
	n += 1

	// write wire field ConfineTo
	w.Write4b(uint32(ConfineTo))
	n += 4

	// write wire field Cursor
	w.Write4b(uint32(Cursor))
	n += 4

	// write wire field Time
	w.Write4b(uint32(Time))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GrabPointerReply struct {
	ResponseType uint8
	Status       uint8
	Sequence     uint16
	Length       uint32
}

func GrabPointerReplyRead(r *Reader, v *GrabPointerReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Status
	v.Status = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}
	return
}

type GrabPointerCookie uint64

func (cookie GrabPointerCookie) Reply(conn *Conn) (*GrabPointerReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply GrabPointerReply
	GrabPointerReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GrabPointer(conn *Conn, OwnerEvents uint8, GrabWindow Window, EventMask uint16, PointerMode uint8, KeyboardMode uint8, ConfineTo Window, Cursor Cursor, Time Timestamp) GrabPointerCookie {
	w := NewWriter()
	GrabPointerRequestWrite(w, OwnerEvents, GrabWindow, EventMask, PointerMode, KeyboardMode, ConfineTo, Cursor, Time)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GrabPointerOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return GrabPointerCookie(seq)
}

func GrabPointerUnchecked(conn *Conn, OwnerEvents uint8, GrabWindow Window, EventMask uint16, PointerMode uint8, KeyboardMode uint8, ConfineTo Window, Cursor Cursor, Time Timestamp) GrabPointerCookie {
	w := NewWriter()
	GrabPointerRequestWrite(w, OwnerEvents, GrabWindow, EventMask, PointerMode, KeyboardMode, ConfineTo, Cursor, Time)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GrabPointerOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GrabPointerCookie(seq)
}

const UngrabPointerOpcode = 27

func UngrabPointerRequestWrite(w *Writer, Time Timestamp) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Time
	w.Write4b(uint32(Time))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func UngrabPointer(conn *Conn, Time Timestamp) VoidCookie {
	w := NewWriter()
	UngrabPointerRequestWrite(w, Time)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: UngrabPointerOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func UngrabPointerChecked(conn *Conn, Time Timestamp) VoidCookie {
	w := NewWriter()
	UngrabPointerRequestWrite(w, Time)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: UngrabPointerOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum ButtonIndex
const (
	ButtonIndexAny = 0
	ButtonIndex1   = 1
	ButtonIndex2   = 2
	ButtonIndex3   = 3
	ButtonIndex4   = 4
	ButtonIndex5   = 5
)

const GrabButtonOpcode = 28

func GrabButtonRequestWrite(w *Writer, OwnerEvents uint8, GrabWindow Window, EventMask uint16, PointerMode uint8, KeyboardMode uint8, ConfineTo Window, Cursor Cursor, Button uint8, Modifiers uint16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field OwnerEvents
	w.Write1b(OwnerEvents)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field GrabWindow
	w.Write4b(uint32(GrabWindow))
	n += 4

	// write wire field EventMask
	w.Write2b(EventMask)
	n += 2

	// write wire field PointerMode
	w.Write1b(PointerMode)
	n += 1

	// write wire field KeyboardMode
	w.Write1b(KeyboardMode)
	n += 1

	// write wire field ConfineTo
	w.Write4b(uint32(ConfineTo))
	n += 4

	// write wire field Cursor
	w.Write4b(uint32(Cursor))
	n += 4

	// write wire field Button
	w.Write1b(Button)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Modifiers
	w.Write2b(Modifiers)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func GrabButton(conn *Conn, OwnerEvents uint8, GrabWindow Window, EventMask uint16, PointerMode uint8, KeyboardMode uint8, ConfineTo Window, Cursor Cursor, Button uint8, Modifiers uint16) VoidCookie {
	w := NewWriter()
	GrabButtonRequestWrite(w, OwnerEvents, GrabWindow, EventMask, PointerMode, KeyboardMode, ConfineTo, Cursor, Button, Modifiers)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: GrabButtonOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func GrabButtonChecked(conn *Conn, OwnerEvents uint8, GrabWindow Window, EventMask uint16, PointerMode uint8, KeyboardMode uint8, ConfineTo Window, Cursor Cursor, Button uint8, Modifiers uint16) VoidCookie {
	w := NewWriter()
	GrabButtonRequestWrite(w, OwnerEvents, GrabWindow, EventMask, PointerMode, KeyboardMode, ConfineTo, Cursor, Button, Modifiers)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: GrabButtonOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const UngrabButtonOpcode = 29

func UngrabButtonRequestWrite(w *Writer, Button uint8, GrabWindow Window, Modifiers uint16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Button
	w.Write1b(Button)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field GrabWindow
	w.Write4b(uint32(GrabWindow))
	n += 4

	// write wire field Modifiers
	w.Write2b(Modifiers)
	n += 2

	// write wire field Pad0
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func UngrabButton(conn *Conn, Button uint8, GrabWindow Window, Modifiers uint16) VoidCookie {
	w := NewWriter()
	UngrabButtonRequestWrite(w, Button, GrabWindow, Modifiers)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: UngrabButtonOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func UngrabButtonChecked(conn *Conn, Button uint8, GrabWindow Window, Modifiers uint16) VoidCookie {
	w := NewWriter()
	UngrabButtonRequestWrite(w, Button, GrabWindow, Modifiers)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: UngrabButtonOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const ChangeActivePointerGrabOpcode = 30

func ChangeActivePointerGrabRequestWrite(w *Writer, Cursor Cursor, Time Timestamp, EventMask uint16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cursor
	w.Write4b(uint32(Cursor))
	n += 4

	// write wire field Time
	w.Write4b(uint32(Time))
	n += 4

	// write wire field EventMask
	w.Write2b(EventMask)
	n += 2

	// write wire field Pad1
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func ChangeActivePointerGrab(conn *Conn, Cursor Cursor, Time Timestamp, EventMask uint16) VoidCookie {
	w := NewWriter()
	ChangeActivePointerGrabRequestWrite(w, Cursor, Time, EventMask)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ChangeActivePointerGrabOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func ChangeActivePointerGrabChecked(conn *Conn, Cursor Cursor, Time Timestamp, EventMask uint16) VoidCookie {
	w := NewWriter()
	ChangeActivePointerGrabRequestWrite(w, Cursor, Time, EventMask)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ChangeActivePointerGrabOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const GrabKeyboardOpcode = 31

func GrabKeyboardRequestWrite(w *Writer, OwnerEvents uint8, GrabWindow Window, Time Timestamp, PointerMode uint8, KeyboardMode uint8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field OwnerEvents
	w.Write1b(OwnerEvents)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field GrabWindow
	w.Write4b(uint32(GrabWindow))
	n += 4

	// write wire field Time
	w.Write4b(uint32(Time))
	n += 4

	// write wire field PointerMode
	w.Write1b(PointerMode)
	n += 1

	// write wire field KeyboardMode
	w.Write1b(KeyboardMode)
	n += 1

	// write wire field Pad0
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GrabKeyboardReply struct {
	ResponseType uint8
	Status       uint8
	Sequence     uint16
	Length       uint32
}

func GrabKeyboardReplyRead(r *Reader, v *GrabKeyboardReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Status
	v.Status = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}
	return
}

type GrabKeyboardCookie uint64

func (cookie GrabKeyboardCookie) Reply(conn *Conn) (*GrabKeyboardReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply GrabKeyboardReply
	GrabKeyboardReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GrabKeyboard(conn *Conn, OwnerEvents uint8, GrabWindow Window, Time Timestamp, PointerMode uint8, KeyboardMode uint8) GrabKeyboardCookie {
	w := NewWriter()
	GrabKeyboardRequestWrite(w, OwnerEvents, GrabWindow, Time, PointerMode, KeyboardMode)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GrabKeyboardOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return GrabKeyboardCookie(seq)
}

func GrabKeyboardUnchecked(conn *Conn, OwnerEvents uint8, GrabWindow Window, Time Timestamp, PointerMode uint8, KeyboardMode uint8) GrabKeyboardCookie {
	w := NewWriter()
	GrabKeyboardRequestWrite(w, OwnerEvents, GrabWindow, Time, PointerMode, KeyboardMode)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GrabKeyboardOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GrabKeyboardCookie(seq)
}

const UngrabKeyboardOpcode = 32

func UngrabKeyboardRequestWrite(w *Writer, Time Timestamp) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Time
	w.Write4b(uint32(Time))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func UngrabKeyboard(conn *Conn, Time Timestamp) VoidCookie {
	w := NewWriter()
	UngrabKeyboardRequestWrite(w, Time)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: UngrabKeyboardOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func UngrabKeyboardChecked(conn *Conn, Time Timestamp) VoidCookie {
	w := NewWriter()
	UngrabKeyboardRequestWrite(w, Time)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: UngrabKeyboardOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum Grab
const (
	GrabAny = 0
)

const GrabKeyOpcode = 33

func GrabKeyRequestWrite(w *Writer, OwnerEvents uint8, GrabWindow Window, Modifiers uint16, Key Keycode, PointerMode uint8, KeyboardMode uint8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field OwnerEvents
	w.Write1b(OwnerEvents)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field GrabWindow
	w.Write4b(uint32(GrabWindow))
	n += 4

	// write wire field Modifiers
	w.Write2b(Modifiers)
	n += 2

	// write wire field Key
	w.Write1b(uint8(Key))
	n += 1

	// write wire field PointerMode
	w.Write1b(PointerMode)
	n += 1

	// write wire field KeyboardMode
	w.Write1b(KeyboardMode)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func GrabKey(conn *Conn, OwnerEvents uint8, GrabWindow Window, Modifiers uint16, Key Keycode, PointerMode uint8, KeyboardMode uint8) VoidCookie {
	w := NewWriter()
	GrabKeyRequestWrite(w, OwnerEvents, GrabWindow, Modifiers, Key, PointerMode, KeyboardMode)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: GrabKeyOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func GrabKeyChecked(conn *Conn, OwnerEvents uint8, GrabWindow Window, Modifiers uint16, Key Keycode, PointerMode uint8, KeyboardMode uint8) VoidCookie {
	w := NewWriter()
	GrabKeyRequestWrite(w, OwnerEvents, GrabWindow, Modifiers, Key, PointerMode, KeyboardMode)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: GrabKeyOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const UngrabKeyOpcode = 34

func UngrabKeyRequestWrite(w *Writer, Key Keycode, GrabWindow Window, Modifiers uint16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Key
	w.Write1b(uint8(Key))
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field GrabWindow
	w.Write4b(uint32(GrabWindow))
	n += 4

	// write wire field Modifiers
	w.Write2b(Modifiers)
	n += 2

	// write wire field Pad0
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func UngrabKey(conn *Conn, Key Keycode, GrabWindow Window, Modifiers uint16) VoidCookie {
	w := NewWriter()
	UngrabKeyRequestWrite(w, Key, GrabWindow, Modifiers)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: UngrabKeyOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func UngrabKeyChecked(conn *Conn, Key Keycode, GrabWindow Window, Modifiers uint16) VoidCookie {
	w := NewWriter()
	UngrabKeyRequestWrite(w, Key, GrabWindow, Modifiers)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: UngrabKeyOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum Allow
const (
	AllowAsyncPointer   = 0
	AllowSyncPointer    = 1
	AllowReplayPointer  = 2
	AllowAsyncKeyboard  = 3
	AllowSyncKeyboard   = 4
	AllowReplayKeyboard = 5
	AllowAsyncBoth      = 6
	AllowSyncBoth       = 7
)

const AllowEventsOpcode = 35

func AllowEventsRequestWrite(w *Writer, Mode uint8, Time Timestamp) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Mode
	w.Write1b(Mode)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Time
	w.Write4b(uint32(Time))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func AllowEvents(conn *Conn, Mode uint8, Time Timestamp) VoidCookie {
	w := NewWriter()
	AllowEventsRequestWrite(w, Mode, Time)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: AllowEventsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func AllowEventsChecked(conn *Conn, Mode uint8, Time Timestamp) VoidCookie {
	w := NewWriter()
	AllowEventsRequestWrite(w, Mode, Time)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: AllowEventsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const GrabServerOpcode = 36

func GrabServerRequestWrite(w *Writer) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func GrabServer(conn *Conn) VoidCookie {
	w := NewWriter()
	GrabServerRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: GrabServerOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func GrabServerChecked(conn *Conn) VoidCookie {
	w := NewWriter()
	GrabServerRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: GrabServerOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const UngrabServerOpcode = 37

func UngrabServerRequestWrite(w *Writer) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func UngrabServer(conn *Conn) VoidCookie {
	w := NewWriter()
	UngrabServerRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: UngrabServerOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func UngrabServerChecked(conn *Conn) VoidCookie {
	w := NewWriter()
	UngrabServerRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: UngrabServerOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const QueryPointerOpcode = 38

func QueryPointerRequestWrite(w *Writer, Window Window) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type QueryPointerReply struct {
	ResponseType uint8
	SameScreen   uint8
	Sequence     uint16
	Length       uint32
	Root         Window
	Child        Window
	RootX        int16
	RootY        int16
	WinX         int16
	WinY         int16
	Mask         uint16
	// Pad0 [2]uint8
}

func QueryPointerReplyRead(r *Reader, v *QueryPointerReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field SameScreen
	v.SameScreen = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Root
	v.Root = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Child
	v.Child = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field RootX
	v.RootX = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field RootY
	v.RootY = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field WinX
	v.WinX = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field WinY
	v.WinY = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Mask
	v.Mask = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(2)
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

type QueryPointerCookie uint64

func (cookie QueryPointerCookie) Reply(conn *Conn) (*QueryPointerReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply QueryPointerReply
	QueryPointerReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryPointer(conn *Conn, Window Window) QueryPointerCookie {
	w := NewWriter()
	QueryPointerRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: QueryPointerOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return QueryPointerCookie(seq)
}

func QueryPointerUnchecked(conn *Conn, Window Window) QueryPointerCookie {
	w := NewWriter()
	QueryPointerRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: QueryPointerOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return QueryPointerCookie(seq)
}

type TimeCoord struct {
	Time Timestamp
	X    int16
	Y    int16
}

type CTimeCoord struct {
	Time Timestamp
	X    int16
	Y    int16
}

func TimeCoordRead(r *Reader, v *TimeCoord) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Time
	v.Time = Timestamp(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field X
	v.X = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Y
	v.Y = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

func TimeCoordReadN(r *Reader, dest []TimeCoord) (n int) {
	for i := 0; i < len(dest); i++ {
		n += TimeCoordRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func TimeCoordWrite(w *Writer, v *TimeCoord) (n int) {

	// write field Time
	w.Write4b(uint32(v.Time))
	n += 4

	// write field X
	w.Write2b(uint16(v.X))
	n += 2

	// write field Y
	w.Write2b(uint16(v.Y))
	n += 2
	return
}
func TimeCoordWriteN(w *Writer, src []TimeCoord) (n int) {
	for i := 0; i < len(src); i++ {
		n += TimeCoordWrite(w, &src[i])
	}
	return
}

const GetMotionEventsOpcode = 39

func GetMotionEventsRequestWrite(w *Writer, Window Window, Start Timestamp, Stop Timestamp) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4

	// write wire field Start
	w.Write4b(uint32(Start))
	n += 4

	// write wire field Stop
	w.Write4b(uint32(Stop))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetMotionEventsReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence  uint16
	Length    uint32
	EventsLen uint32
	// Pad1 [20]uint8
	Events []TimeCoord
}

func GetMotionEventsReplyRead(r *Reader, v *GetMotionEventsReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field EventsLen
	v.EventsLen = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(20)
	n += 20
	if r.Err() != nil {
		return
	}

	// read field Events
	v.Events = make([]TimeCoord, int(v.EventsLen))
	blockLen += TimeCoordReadN(r, v.Events)
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CTimeCoord{}))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type GetMotionEventsCookie uint64

func (cookie GetMotionEventsCookie) Reply(conn *Conn) (*GetMotionEventsReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply GetMotionEventsReply
	GetMotionEventsReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetMotionEvents(conn *Conn, Window Window, Start Timestamp, Stop Timestamp) GetMotionEventsCookie {
	w := NewWriter()
	GetMotionEventsRequestWrite(w, Window, Start, Stop)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetMotionEventsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return GetMotionEventsCookie(seq)
}

func GetMotionEventsUnchecked(conn *Conn, Window Window, Start Timestamp, Stop Timestamp) GetMotionEventsCookie {
	w := NewWriter()
	GetMotionEventsRequestWrite(w, Window, Start, Stop)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetMotionEventsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetMotionEventsCookie(seq)
}

const TranslateCoordinatesOpcode = 40

func TranslateCoordinatesRequestWrite(w *Writer, SrcWindow Window, DstWindow Window, SrcX int16, SrcY int16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field SrcWindow
	w.Write4b(uint32(SrcWindow))
	n += 4

	// write wire field DstWindow
	w.Write4b(uint32(DstWindow))
	n += 4

	// write wire field SrcX
	w.Write2b(uint16(SrcX))
	n += 2

	// write wire field SrcY
	w.Write2b(uint16(SrcY))
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type TranslateCoordinatesReply struct {
	ResponseType uint8
	SameScreen   uint8
	Sequence     uint16
	Length       uint32
	Child        Window
	DstX         int16
	DstY         int16
}

func TranslateCoordinatesReplyRead(r *Reader, v *TranslateCoordinatesReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field SameScreen
	v.SameScreen = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Child
	v.Child = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field DstX
	v.DstX = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field DstY
	v.DstY = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

type TranslateCoordinatesCookie uint64

func (cookie TranslateCoordinatesCookie) Reply(conn *Conn) (*TranslateCoordinatesReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply TranslateCoordinatesReply
	TranslateCoordinatesReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func TranslateCoordinates(conn *Conn, SrcWindow Window, DstWindow Window, SrcX int16, SrcY int16) TranslateCoordinatesCookie {
	w := NewWriter()
	TranslateCoordinatesRequestWrite(w, SrcWindow, DstWindow, SrcX, SrcY)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: TranslateCoordinatesOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return TranslateCoordinatesCookie(seq)
}

func TranslateCoordinatesUnchecked(conn *Conn, SrcWindow Window, DstWindow Window, SrcX int16, SrcY int16) TranslateCoordinatesCookie {
	w := NewWriter()
	TranslateCoordinatesRequestWrite(w, SrcWindow, DstWindow, SrcX, SrcY)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: TranslateCoordinatesOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return TranslateCoordinatesCookie(seq)
}

const WarpPointerOpcode = 41

func WarpPointerRequestWrite(w *Writer, SrcWindow Window, DstWindow Window, SrcX int16, SrcY int16, SrcWidth uint16, SrcHeight uint16, DstX int16, DstY int16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field SrcWindow
	w.Write4b(uint32(SrcWindow))
	n += 4

	// write wire field DstWindow
	w.Write4b(uint32(DstWindow))
	n += 4

	// write wire field SrcX
	w.Write2b(uint16(SrcX))
	n += 2

	// write wire field SrcY
	w.Write2b(uint16(SrcY))
	n += 2

	// write wire field SrcWidth
	w.Write2b(SrcWidth)
	n += 2

	// write wire field SrcHeight
	w.Write2b(SrcHeight)
	n += 2

	// write wire field DstX
	w.Write2b(uint16(DstX))
	n += 2

	// write wire field DstY
	w.Write2b(uint16(DstY))
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func WarpPointer(conn *Conn, SrcWindow Window, DstWindow Window, SrcX int16, SrcY int16, SrcWidth uint16, SrcHeight uint16, DstX int16, DstY int16) VoidCookie {
	w := NewWriter()
	WarpPointerRequestWrite(w, SrcWindow, DstWindow, SrcX, SrcY, SrcWidth, SrcHeight, DstX, DstY)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: WarpPointerOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func WarpPointerChecked(conn *Conn, SrcWindow Window, DstWindow Window, SrcX int16, SrcY int16, SrcWidth uint16, SrcHeight uint16, DstX int16, DstY int16) VoidCookie {
	w := NewWriter()
	WarpPointerRequestWrite(w, SrcWindow, DstWindow, SrcX, SrcY, SrcWidth, SrcHeight, DstX, DstY)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: WarpPointerOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum InputFocus
const (
	InputFocusNone           = 0
	InputFocusPointerRoot    = 1
	InputFocusParent         = 2
	InputFocusFollowKeyboard = 3
)

const SetInputFocusOpcode = 42

func SetInputFocusRequestWrite(w *Writer, RevertTo uint8, Focus Window, Time Timestamp) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field RevertTo
	w.Write1b(RevertTo)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Focus
	w.Write4b(uint32(Focus))
	n += 4

	// write wire field Time
	w.Write4b(uint32(Time))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SetInputFocus(conn *Conn, RevertTo uint8, Focus Window, Time Timestamp) VoidCookie {
	w := NewWriter()
	SetInputFocusRequestWrite(w, RevertTo, Focus, Time)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: SetInputFocusOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func SetInputFocusChecked(conn *Conn, RevertTo uint8, Focus Window, Time Timestamp) VoidCookie {
	w := NewWriter()
	SetInputFocusRequestWrite(w, RevertTo, Focus, Time)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: SetInputFocusOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const GetInputFocusOpcode = 43

func GetInputFocusRequestWrite(w *Writer) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetInputFocusReply struct {
	ResponseType uint8
	RevertTo     uint8
	Sequence     uint16
	Length       uint32
	Focus        Window
}

func GetInputFocusReplyRead(r *Reader, v *GetInputFocusReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field RevertTo
	v.RevertTo = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Focus
	v.Focus = Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}
	return
}

type GetInputFocusCookie uint64

func (cookie GetInputFocusCookie) Reply(conn *Conn) (*GetInputFocusReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply GetInputFocusReply
	GetInputFocusReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetInputFocus(conn *Conn) GetInputFocusCookie {
	w := NewWriter()
	GetInputFocusRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetInputFocusOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return GetInputFocusCookie(seq)
}

func GetInputFocusUnchecked(conn *Conn) GetInputFocusCookie {
	w := NewWriter()
	GetInputFocusRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetInputFocusOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetInputFocusCookie(seq)
}

const QueryKeymapOpcode = 44

func QueryKeymapRequestWrite(w *Writer) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type QueryKeymapReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Length   uint32
	Keys     [32]uint8
}

func QueryKeymapReplyRead(r *Reader, v *QueryKeymapReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Keys
	{
		dataLen := int(32)
		for i := 0; i < dataLen; i++ {
			v.Keys[i] = r.Read1b()
			if r.Err() != nil {
				return
			}
		}
		n += dataLen
	}
	return
}

type QueryKeymapCookie uint64

func (cookie QueryKeymapCookie) Reply(conn *Conn) (*QueryKeymapReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply QueryKeymapReply
	QueryKeymapReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryKeymap(conn *Conn) QueryKeymapCookie {
	w := NewWriter()
	QueryKeymapRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: QueryKeymapOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return QueryKeymapCookie(seq)
}

func QueryKeymapUnchecked(conn *Conn) QueryKeymapCookie {
	w := NewWriter()
	QueryKeymapRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: QueryKeymapOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return QueryKeymapCookie(seq)
}

const OpenFontOpcode = 45

func OpenFontRequestWrite(w *Writer, Fid Font, NameLen uint16, Name string) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Fid
	w.Write4b(uint32(Fid))
	n += 4

	// write wire field NameLen
	w.Write2b(NameLen)
	n += 2

	// write wire field Pad1
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Name
	w.WriteString(Name)
	n += len(Name)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func OpenFont(conn *Conn, Fid Font, NameLen uint16, Name string) VoidCookie {
	w := NewWriter()
	OpenFontRequestWrite(w, Fid, NameLen, Name)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: OpenFontOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func OpenFontChecked(conn *Conn, Fid Font, NameLen uint16, Name string) VoidCookie {
	w := NewWriter()
	OpenFontRequestWrite(w, Fid, NameLen, Name)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: OpenFontOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const CloseFontOpcode = 46

func CloseFontRequestWrite(w *Writer, Font Font) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Font
	w.Write4b(uint32(Font))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CloseFont(conn *Conn, Font Font) VoidCookie {
	w := NewWriter()
	CloseFontRequestWrite(w, Font)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CloseFontOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func CloseFontChecked(conn *Conn, Font Font) VoidCookie {
	w := NewWriter()
	CloseFontRequestWrite(w, Font)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CloseFontOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum FontDraw
const (
	FontDrawLeftToRight = 0
	FontDrawRightToLeft = 1
)

func FontDrawEnumToStr(v int) string {
	switch v {
	case 0:
		return "LeftToRight"
	case 1:
		return "RightToLeft"
	default:
		return "<Unknown>"
	}
}

type FontProp struct {
	Name  Atom
	Value uint32
}

type CFontProp struct {
	Name  Atom
	Value uint32
}

func FontPropRead(r *Reader, v *FontProp) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Name
	v.Name = Atom(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Value
	v.Value = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}
	return
}

func FontPropReadN(r *Reader, dest []FontProp) (n int) {
	for i := 0; i < len(dest); i++ {
		n += FontPropRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func FontPropWrite(w *Writer, v *FontProp) (n int) {

	// write field Name
	w.Write4b(uint32(v.Name))
	n += 4

	// write field Value
	w.Write4b(v.Value)
	n += 4
	return
}
func FontPropWriteN(w *Writer, src []FontProp) (n int) {
	for i := 0; i < len(src); i++ {
		n += FontPropWrite(w, &src[i])
	}
	return
}

type CharInfo struct {
	LeftSideBearing  int16
	RightSideBearing int16
	CharacterWidth   int16
	Ascent           int16
	Descent          int16
	Attributes       uint16
}

type CCharInfo struct {
	LeftSideBearing  int16
	RightSideBearing int16
	CharacterWidth   int16
	Ascent           int16
	Descent          int16
	Attributes       uint16
}

func CharInfoRead(r *Reader, v *CharInfo) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field LeftSideBearing
	v.LeftSideBearing = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field RightSideBearing
	v.RightSideBearing = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field CharacterWidth
	v.CharacterWidth = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Ascent
	v.Ascent = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Descent
	v.Descent = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Attributes
	v.Attributes = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

func CharInfoReadN(r *Reader, dest []CharInfo) (n int) {
	for i := 0; i < len(dest); i++ {
		n += CharInfoRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func CharInfoWrite(w *Writer, v *CharInfo) (n int) {

	// write field LeftSideBearing
	w.Write2b(uint16(v.LeftSideBearing))
	n += 2

	// write field RightSideBearing
	w.Write2b(uint16(v.RightSideBearing))
	n += 2

	// write field CharacterWidth
	w.Write2b(uint16(v.CharacterWidth))
	n += 2

	// write field Ascent
	w.Write2b(uint16(v.Ascent))
	n += 2

	// write field Descent
	w.Write2b(uint16(v.Descent))
	n += 2

	// write field Attributes
	w.Write2b(v.Attributes)
	n += 2
	return
}
func CharInfoWriteN(w *Writer, src []CharInfo) (n int) {
	for i := 0; i < len(src); i++ {
		n += CharInfoWrite(w, &src[i])
	}
	return
}

const QueryFontOpcode = 47

func QueryFontRequestWrite(w *Writer, Font Fontable) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Font
	w.Write4b(uint32(Font))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type QueryFontReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence  uint16
	Length    uint32
	MinBounds CharInfo
	// Pad1 [4]uint8
	MaxBounds CharInfo
	// Pad2 [4]uint8
	MinCharOrByte2 uint16
	MaxCharOrByte2 uint16
	DefaultChar    uint16
	PropertiesLen  uint16
	DrawDirection  uint8
	MinByte1       uint8
	MaxByte1       uint8
	AllCharsExist  uint8
	FontAscent     int16
	FontDescent    int16
	CharInfosLen   uint32
	Properties     []FontProp
	CharInfos      []CharInfo
}

func QueryFontReplyRead(r *Reader, v *QueryFontReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field MinBounds
	n += CharInfoRead(r, &v.MinBounds)
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(4)
	n += 4
	if r.Err() != nil {
		return
	}

	// read field MaxBounds
	n += CharInfoRead(r, &v.MaxBounds)
	if r.Err() != nil {
		return
	}

	// read field Pad2
	r.ReadPad(4)
	n += 4
	if r.Err() != nil {
		return
	}

	// read field MinCharOrByte2
	v.MinCharOrByte2 = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field MaxCharOrByte2
	v.MaxCharOrByte2 = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field DefaultChar
	v.DefaultChar = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field PropertiesLen
	v.PropertiesLen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field DrawDirection
	v.DrawDirection = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field MinByte1
	v.MinByte1 = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field MaxByte1
	v.MaxByte1 = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field AllCharsExist
	v.AllCharsExist = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field FontAscent
	v.FontAscent = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field FontDescent
	v.FontDescent = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field CharInfosLen
	v.CharInfosLen = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Properties
	v.Properties = make([]FontProp, int(v.PropertiesLen))
	blockLen += FontPropReadN(r, v.Properties)
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CFontProp{}))

	// read field CharInfos

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	v.CharInfos = make([]CharInfo, int(v.CharInfosLen))
	blockLen += CharInfoReadN(r, v.CharInfos)
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CCharInfo{}))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type QueryFontCookie uint64

func (cookie QueryFontCookie) Reply(conn *Conn) (*QueryFontReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply QueryFontReply
	QueryFontReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryFont(conn *Conn, Font Fontable) QueryFontCookie {
	w := NewWriter()
	QueryFontRequestWrite(w, Font)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: QueryFontOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return QueryFontCookie(seq)
}

func QueryFontUnchecked(conn *Conn, Font Fontable) QueryFontCookie {
	w := NewWriter()
	QueryFontRequestWrite(w, Font)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: QueryFontOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return QueryFontCookie(seq)
}

const QueryTextExtentsOpcode = 48

func QueryTextExtentsRequestWrite(w *Writer, Font Fontable, StringLen uint32, String []Char2B) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field OddLength

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Font
	w.Write4b(uint32(Font))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field String
	n += Char2BWriteN(w, String)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type QueryTextExtentsReply struct {
	ResponseType   uint8
	DrawDirection  uint8
	Sequence       uint16
	Length         uint32
	FontAscent     int16
	FontDescent    int16
	OverallAscent  int16
	OverallDescent int16
	OverallWidth   int32
	OverallLeft    int32
	OverallRight   int32
}

func QueryTextExtentsReplyRead(r *Reader, v *QueryTextExtentsReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field DrawDirection
	v.DrawDirection = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field FontAscent
	v.FontAscent = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field FontDescent
	v.FontDescent = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field OverallAscent
	v.OverallAscent = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field OverallDescent
	v.OverallDescent = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field OverallWidth
	v.OverallWidth = int32(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field OverallLeft
	v.OverallLeft = int32(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field OverallRight
	v.OverallRight = int32(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}
	return
}

type QueryTextExtentsCookie uint64

func (cookie QueryTextExtentsCookie) Reply(conn *Conn) (*QueryTextExtentsReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply QueryTextExtentsReply
	QueryTextExtentsReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryTextExtents(conn *Conn, Font Fontable, StringLen uint32, String []Char2B) QueryTextExtentsCookie {
	w := NewWriter()
	QueryTextExtentsRequestWrite(w, Font, StringLen, String)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: QueryTextExtentsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return QueryTextExtentsCookie(seq)
}

func QueryTextExtentsUnchecked(conn *Conn, Font Fontable, StringLen uint32, String []Char2B) QueryTextExtentsCookie {
	w := NewWriter()
	QueryTextExtentsRequestWrite(w, Font, StringLen, String)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: QueryTextExtentsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return QueryTextExtentsCookie(seq)
}

type Str struct {
	NameLen uint8
	Name    string
}

type CStr struct {
	NameLen uint8
}

func StrRead(r *Reader, v *Str) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field NameLen
	v.NameLen = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Name
	{
		dataLen := int(int(v.NameLen))
		data := r.ReadString(dataLen)
		if r.Err() != nil {
			return
		}
		blockLen += dataLen
		n += blockLen
		v.Name = data
	}
	alignTo = int(unsafe.Alignof(byte(0)))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

func StrReadN(r *Reader, dest []Str) (n int) {
	for i := 0; i < len(dest); i++ {
		n += StrRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func StrWrite(w *Writer, v *Str) (n int) {

	// write field NameLen
	w.Write1b(v.NameLen)
	n += 1

	// write field Name
	return
}
func StrWriteN(w *Writer, src []Str) (n int) {
	for i := 0; i < len(src); i++ {
		n += StrWrite(w, &src[i])
	}
	return
}

const ListFontsOpcode = 49

func ListFontsRequestWrite(w *Writer, MaxNames uint16, PatternLen uint16, Pattern string) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field MaxNames
	w.Write2b(MaxNames)
	n += 2

	// write wire field PatternLen
	w.Write2b(PatternLen)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Pattern
	w.WriteString(Pattern)
	n += len(Pattern)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type ListFontsReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Length   uint32
	NamesLen uint16
	// Pad1 [22]uint8
	Names []Str
}

func ListFontsReplyRead(r *Reader, v *ListFontsReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field NamesLen
	v.NamesLen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(22)
	n += 22
	if r.Err() != nil {
		return
	}

	// read field Names
	v.Names = make([]Str, int(v.NamesLen))
	blockLen += StrReadN(r, v.Names)
	n += blockLen
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CStr{}))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type ListFontsCookie uint64

func (cookie ListFontsCookie) Reply(conn *Conn) (*ListFontsReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply ListFontsReply
	ListFontsReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func ListFonts(conn *Conn, MaxNames uint16, PatternLen uint16, Pattern string) ListFontsCookie {
	w := NewWriter()
	ListFontsRequestWrite(w, MaxNames, PatternLen, Pattern)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: ListFontsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return ListFontsCookie(seq)
}

func ListFontsUnchecked(conn *Conn, MaxNames uint16, PatternLen uint16, Pattern string) ListFontsCookie {
	w := NewWriter()
	ListFontsRequestWrite(w, MaxNames, PatternLen, Pattern)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: ListFontsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return ListFontsCookie(seq)
}

const ListFontsWithInfoOpcode = 50

func ListFontsWithInfoRequestWrite(w *Writer, MaxNames uint16, PatternLen uint16, Pattern string) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field MaxNames
	w.Write2b(MaxNames)
	n += 2

	// write wire field PatternLen
	w.Write2b(PatternLen)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Pattern
	w.WriteString(Pattern)
	n += len(Pattern)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type ListFontsWithInfoReply struct {
	ResponseType uint8
	NameLen      uint8
	Sequence     uint16
	Length       uint32
	MinBounds    CharInfo
	// Pad0 [4]uint8
	MaxBounds CharInfo
	// Pad1 [4]uint8
	MinCharOrByte2 uint16
	MaxCharOrByte2 uint16
	DefaultChar    uint16
	PropertiesLen  uint16
	DrawDirection  uint8
	MinByte1       uint8
	MaxByte1       uint8
	AllCharsExist  uint8
	FontAscent     int16
	FontDescent    int16
	RepliesHint    uint32
	Properties     []FontProp
	Name           string
}

func ListFontsWithInfoReplyRead(r *Reader, v *ListFontsWithInfoReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field NameLen
	v.NameLen = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field MinBounds
	n += CharInfoRead(r, &v.MinBounds)
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(4)
	n += 4
	if r.Err() != nil {
		return
	}

	// read field MaxBounds
	n += CharInfoRead(r, &v.MaxBounds)
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(4)
	n += 4
	if r.Err() != nil {
		return
	}

	// read field MinCharOrByte2
	v.MinCharOrByte2 = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field MaxCharOrByte2
	v.MaxCharOrByte2 = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field DefaultChar
	v.DefaultChar = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field PropertiesLen
	v.PropertiesLen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field DrawDirection
	v.DrawDirection = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field MinByte1
	v.MinByte1 = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field MaxByte1
	v.MaxByte1 = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field AllCharsExist
	v.AllCharsExist = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field FontAscent
	v.FontAscent = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field FontDescent
	v.FontDescent = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field RepliesHint
	v.RepliesHint = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Properties
	v.Properties = make([]FontProp, int(v.PropertiesLen))
	blockLen += FontPropReadN(r, v.Properties)
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CFontProp{}))

	// read field Name

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	{
		dataLen := int(int(v.NameLen))
		data := r.ReadString(dataLen)
		if r.Err() != nil {
			return
		}
		blockLen += dataLen
		n += blockLen
		v.Name = data
	}
	alignTo = int(unsafe.Alignof(byte(0)))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type ListFontsWithInfoCookie uint64

func (cookie ListFontsWithInfoCookie) Reply(conn *Conn) (*ListFontsWithInfoReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply ListFontsWithInfoReply
	ListFontsWithInfoReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func ListFontsWithInfo(conn *Conn, MaxNames uint16, PatternLen uint16, Pattern string) ListFontsWithInfoCookie {
	w := NewWriter()
	ListFontsWithInfoRequestWrite(w, MaxNames, PatternLen, Pattern)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: ListFontsWithInfoOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return ListFontsWithInfoCookie(seq)
}

func ListFontsWithInfoUnchecked(conn *Conn, MaxNames uint16, PatternLen uint16, Pattern string) ListFontsWithInfoCookie {
	w := NewWriter()
	ListFontsWithInfoRequestWrite(w, MaxNames, PatternLen, Pattern)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: ListFontsWithInfoOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return ListFontsWithInfoCookie(seq)
}

const SetFontPathOpcode = 51

func SetFontPathRequestWrite(w *Writer, FontQty uint16, Font []Str) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field FontQty
	w.Write2b(FontQty)
	n += 2

	// write wire field Pad1
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Font
}

func SetFontPath(conn *Conn, FontQty uint16, Font []Str) VoidCookie {
	w := NewWriter()
	SetFontPathRequestWrite(w, FontQty, Font)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: SetFontPathOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func SetFontPathChecked(conn *Conn, FontQty uint16, Font []Str) VoidCookie {
	w := NewWriter()
	SetFontPathRequestWrite(w, FontQty, Font)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: SetFontPathOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const GetFontPathOpcode = 52

func GetFontPathRequestWrite(w *Writer) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetFontPathReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Length   uint32
	PathLen  uint16
	// Pad1 [22]uint8
	Path []Str
}

func GetFontPathReplyRead(r *Reader, v *GetFontPathReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field PathLen
	v.PathLen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(22)
	n += 22
	if r.Err() != nil {
		return
	}

	// read field Path
	v.Path = make([]Str, int(v.PathLen))
	blockLen += StrReadN(r, v.Path)
	n += blockLen
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CStr{}))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type GetFontPathCookie uint64

func (cookie GetFontPathCookie) Reply(conn *Conn) (*GetFontPathReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply GetFontPathReply
	GetFontPathReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetFontPath(conn *Conn) GetFontPathCookie {
	w := NewWriter()
	GetFontPathRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetFontPathOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return GetFontPathCookie(seq)
}

func GetFontPathUnchecked(conn *Conn) GetFontPathCookie {
	w := NewWriter()
	GetFontPathRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetFontPathOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetFontPathCookie(seq)
}

const CreatePixmapOpcode = 53

func CreatePixmapRequestWrite(w *Writer, Depth uint8, Pid Pixmap, Drawable Drawable, Width uint16, Height uint16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Depth
	w.Write1b(Depth)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Pid
	w.Write4b(uint32(Pid))
	n += 4

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field Width
	w.Write2b(Width)
	n += 2

	// write wire field Height
	w.Write2b(Height)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CreatePixmap(conn *Conn, Depth uint8, Pid Pixmap, Drawable Drawable, Width uint16, Height uint16) VoidCookie {
	w := NewWriter()
	CreatePixmapRequestWrite(w, Depth, Pid, Drawable, Width, Height)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CreatePixmapOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func CreatePixmapChecked(conn *Conn, Depth uint8, Pid Pixmap, Drawable Drawable, Width uint16, Height uint16) VoidCookie {
	w := NewWriter()
	CreatePixmapRequestWrite(w, Depth, Pid, Drawable, Width, Height)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CreatePixmapOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const FreePixmapOpcode = 54

func FreePixmapRequestWrite(w *Writer, Pixmap Pixmap) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Pixmap
	w.Write4b(uint32(Pixmap))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func FreePixmap(conn *Conn, Pixmap Pixmap) VoidCookie {
	w := NewWriter()
	FreePixmapRequestWrite(w, Pixmap)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: FreePixmapOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func FreePixmapChecked(conn *Conn, Pixmap Pixmap) VoidCookie {
	w := NewWriter()
	FreePixmapRequestWrite(w, Pixmap)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: FreePixmapOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum GC
const (
	GCFunction           = 1
	GCPlaneMask          = 2
	GCForeground         = 4
	GCBackground         = 8
	GCLineWidth          = 16
	GCLineStyle          = 32
	GCCapStyle           = 64
	GCJoinStyle          = 128
	GCFillStyle          = 256
	GCFillRule           = 512
	GCTile               = 1024
	GCStipple            = 2048
	GCTileStippleOriginX = 4096
	GCTileStippleOriginY = 8192
	GCFont               = 16384
	GCSubwindowMode      = 32768
	GCGraphicsExposures  = 65536
	GCClipOriginX        = 131072
	GCClipOriginY        = 262144
	GCClipMask           = 524288
	GCDashOffset         = 1048576
	GCDashList           = 2097152
	GCArcMode            = 4194304
)

// enum GX
const (
	GXClear        = 0
	GXAnd          = 1
	GXAndReverse   = 2
	GXCopy         = 3
	GXAndInverted  = 4
	GXNoop         = 5
	GXXor          = 6
	GXOr           = 7
	GXNor          = 8
	GXEquiv        = 9
	GXInvert       = 10
	GXOrReverse    = 11
	GXCopyInverted = 12
	GXOrInverted   = 13
	GXNand         = 14
	GXSet          = 15
)

// enum LineStyle
const (
	LineStyleSolid      = 0
	LineStyleOnOffDash  = 1
	LineStyleDoubleDash = 2
)

// enum CapStyle
const (
	CapStyleNotLast    = 0
	CapStyleButt       = 1
	CapStyleRound      = 2
	CapStyleProjecting = 3
)

// enum JoinStyle
const (
	JoinStyleMiter = 0
	JoinStyleRound = 1
	JoinStyleBevel = 2
)

// enum FillStyle
const (
	FillStyleSolid          = 0
	FillStyleTiled          = 1
	FillStyleStippled       = 2
	FillStyleOpaqueStippled = 3
)

// enum FillRule
const (
	FillRuleEvenOdd = 0
	FillRuleWinding = 1
)

// enum SubwindowMode
const (
	SubwindowModeClipByChildren   = 0
	SubwindowModeIncludeInferiors = 1
)

// enum ArcMode
const (
	ArcModeChord    = 0
	ArcModePieSlice = 1
)

type CreateGcValueList struct {
	Function           uint32
	PlaneMask          uint32
	Foreground         uint32
	Background         uint32
	LineWidth          uint32
	LineStyle          uint32
	CapStyle           uint32
	JoinStyle          uint32
	FillStyle          uint32
	FillRule           uint32
	Tile               Pixmap
	Stipple            Pixmap
	TileStippleXOrigin int32
	TileStippleYOrigin int32
	Font               Font
	SubwindowMode      uint32
	GraphicsExposures  Bool32
	ClipXOrigin        int32
	ClipYOrigin        int32
	ClipMask           Pixmap
	DashOffset         uint32
	Dashes             uint32
	ArcMode            uint32
}

// _go_switch_write SwitchType "xcb.CreateGC.value_list"
func CreateGcValueListSerialize(w *Writer, ValueMask uint32, _aux *CreateGcValueList) {
	n := 0
	xgbPaddingOffset := 0
	alignTo := 0
	xgbPad := 0
	// switch expr: int(ValueMask)

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase1"
	if (int(ValueMask) & GCFunction) != 0 {
		// todo
		// field Field "function" in BitcaseType "xcb.CreateGC.value_list.bitcase1"
		w.Write4b(_aux.Function)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase2"
	if (int(ValueMask) & GCPlaneMask) != 0 {
		// todo
		// field Field "plane_mask" in BitcaseType "xcb.CreateGC.value_list.bitcase2"
		w.Write4b(_aux.PlaneMask)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase3"
	if (int(ValueMask) & GCForeground) != 0 {
		// todo
		// field Field "foreground" in BitcaseType "xcb.CreateGC.value_list.bitcase3"
		w.Write4b(_aux.Foreground)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase4"
	if (int(ValueMask) & GCBackground) != 0 {
		// todo
		// field Field "background" in BitcaseType "xcb.CreateGC.value_list.bitcase4"
		w.Write4b(_aux.Background)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase5"
	if (int(ValueMask) & GCLineWidth) != 0 {
		// todo
		// field Field "line_width" in BitcaseType "xcb.CreateGC.value_list.bitcase5"
		w.Write4b(_aux.LineWidth)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase6"
	if (int(ValueMask) & GCLineStyle) != 0 {
		// todo
		// field Field "line_style" in BitcaseType "xcb.CreateGC.value_list.bitcase6"
		w.Write4b(_aux.LineStyle)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase7"
	if (int(ValueMask) & GCCapStyle) != 0 {
		// todo
		// field Field "cap_style" in BitcaseType "xcb.CreateGC.value_list.bitcase7"
		w.Write4b(_aux.CapStyle)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase8"
	if (int(ValueMask) & GCJoinStyle) != 0 {
		// todo
		// field Field "join_style" in BitcaseType "xcb.CreateGC.value_list.bitcase8"
		w.Write4b(_aux.JoinStyle)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase9"
	if (int(ValueMask) & GCFillStyle) != 0 {
		// todo
		// field Field "fill_style" in BitcaseType "xcb.CreateGC.value_list.bitcase9"
		w.Write4b(_aux.FillStyle)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase10"
	if (int(ValueMask) & GCFillRule) != 0 {
		// todo
		// field Field "fill_rule" in BitcaseType "xcb.CreateGC.value_list.bitcase10"
		w.Write4b(_aux.FillRule)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase11"
	if (int(ValueMask) & GCTile) != 0 {
		// todo
		// field Field "tile" in BitcaseType "xcb.CreateGC.value_list.bitcase11"
		w.Write4b(uint32(_aux.Tile))
		n += 4
		alignTo = int(unsafe.Alignof(Pixmap(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase12"
	if (int(ValueMask) & GCStipple) != 0 {
		// todo
		// field Field "stipple" in BitcaseType "xcb.CreateGC.value_list.bitcase12"
		w.Write4b(uint32(_aux.Stipple))
		n += 4
		alignTo = int(unsafe.Alignof(Pixmap(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase13"
	if (int(ValueMask) & GCTileStippleOriginX) != 0 {
		// todo
		// field Field "tile_stipple_x_origin" in BitcaseType "xcb.CreateGC.value_list.bitcase13"
		w.Write4b(uint32(_aux.TileStippleXOrigin))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase14"
	if (int(ValueMask) & GCTileStippleOriginY) != 0 {
		// todo
		// field Field "tile_stipple_y_origin" in BitcaseType "xcb.CreateGC.value_list.bitcase14"
		w.Write4b(uint32(_aux.TileStippleYOrigin))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase15"
	if (int(ValueMask) & GCFont) != 0 {
		// todo
		// field Field "font" in BitcaseType "xcb.CreateGC.value_list.bitcase15"
		w.Write4b(uint32(_aux.Font))
		n += 4
		alignTo = int(unsafe.Alignof(Font(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase16"
	if (int(ValueMask) & GCSubwindowMode) != 0 {
		// todo
		// field Field "subwindow_mode" in BitcaseType "xcb.CreateGC.value_list.bitcase16"
		w.Write4b(_aux.SubwindowMode)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase17"
	if (int(ValueMask) & GCGraphicsExposures) != 0 {
		// todo
		// field Field "graphics_exposures" in BitcaseType "xcb.CreateGC.value_list.bitcase17"
		w.Write4b(uint32(_aux.GraphicsExposures))
		n += 4
		alignTo = int(unsafe.Alignof(Bool32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase18"
	if (int(ValueMask) & GCClipOriginX) != 0 {
		// todo
		// field Field "clip_x_origin" in BitcaseType "xcb.CreateGC.value_list.bitcase18"
		w.Write4b(uint32(_aux.ClipXOrigin))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase19"
	if (int(ValueMask) & GCClipOriginY) != 0 {
		// todo
		// field Field "clip_y_origin" in BitcaseType "xcb.CreateGC.value_list.bitcase19"
		w.Write4b(uint32(_aux.ClipYOrigin))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase20"
	if (int(ValueMask) & GCClipMask) != 0 {
		// todo
		// field Field "clip_mask" in BitcaseType "xcb.CreateGC.value_list.bitcase20"
		w.Write4b(uint32(_aux.ClipMask))
		n += 4
		alignTo = int(unsafe.Alignof(Pixmap(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase21"
	if (int(ValueMask) & GCDashOffset) != 0 {
		// todo
		// field Field "dash_offset" in BitcaseType "xcb.CreateGC.value_list.bitcase21"
		w.Write4b(_aux.DashOffset)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase22"
	if (int(ValueMask) & GCDashList) != 0 {
		// todo
		// field Field "dashes" in BitcaseType "xcb.CreateGC.value_list.bitcase22"
		w.Write4b(_aux.Dashes)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.CreateGC.value_list.bitcase23"
	if (int(ValueMask) & GCArcMode) != 0 {
		// todo
		// field Field "arc_mode" in BitcaseType "xcb.CreateGC.value_list.bitcase23"
		w.Write4b(_aux.ArcMode)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	/* insert padding */
	xgbPad = -(n + xgbPaddingOffset) & (alignTo - 1)
	w.WritePad(xgbPad)
	xgbPad = 0
	n = 0
	xgbPaddingOffset = 0
}

const CreateGCOpcode = 55

func CreateGcRequestWrite(w *Writer, Cid GContext, Drawable Drawable, ValueMask uint32, ValueList *CreateGcValueList) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cid
	w.Write4b(uint32(Cid))
	n += 4

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field ValueMask
	w.Write4b(ValueMask)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field ValueList
	// switch serialize
	CreateGcValueListSerialize(w, ValueMask, ValueList)
}

func CreateGC(conn *Conn, Cid GContext, Drawable Drawable, ValueMask uint32, ValueList *CreateGcValueList) VoidCookie {
	w := NewWriter()
	CreateGcRequestWrite(w, Cid, Drawable, ValueMask, ValueList)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CreateGCOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func CreateGCChecked(conn *Conn, Cid GContext, Drawable Drawable, ValueMask uint32, ValueList *CreateGcValueList) VoidCookie {
	w := NewWriter()
	CreateGcRequestWrite(w, Cid, Drawable, ValueMask, ValueList)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CreateGCOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

type ChangeGcValueList struct {
	Function           uint32
	PlaneMask          uint32
	Foreground         uint32
	Background         uint32
	LineWidth          uint32
	LineStyle          uint32
	CapStyle           uint32
	JoinStyle          uint32
	FillStyle          uint32
	FillRule           uint32
	Tile               Pixmap
	Stipple            Pixmap
	TileStippleXOrigin int32
	TileStippleYOrigin int32
	Font               Font
	SubwindowMode      uint32
	GraphicsExposures  Bool32
	ClipXOrigin        int32
	ClipYOrigin        int32
	ClipMask           Pixmap
	DashOffset         uint32
	Dashes             uint32
	ArcMode            uint32
}

// _go_switch_write SwitchType "xcb.ChangeGC.value_list"
func ChangeGcValueListSerialize(w *Writer, ValueMask uint32, _aux *ChangeGcValueList) {
	n := 0
	xgbPaddingOffset := 0
	alignTo := 0
	xgbPad := 0
	// switch expr: int(ValueMask)

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase1"
	if (int(ValueMask) & GCFunction) != 0 {
		// todo
		// field Field "function" in BitcaseType "xcb.ChangeGC.value_list.bitcase1"
		w.Write4b(_aux.Function)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase2"
	if (int(ValueMask) & GCPlaneMask) != 0 {
		// todo
		// field Field "plane_mask" in BitcaseType "xcb.ChangeGC.value_list.bitcase2"
		w.Write4b(_aux.PlaneMask)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase3"
	if (int(ValueMask) & GCForeground) != 0 {
		// todo
		// field Field "foreground" in BitcaseType "xcb.ChangeGC.value_list.bitcase3"
		w.Write4b(_aux.Foreground)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase4"
	if (int(ValueMask) & GCBackground) != 0 {
		// todo
		// field Field "background" in BitcaseType "xcb.ChangeGC.value_list.bitcase4"
		w.Write4b(_aux.Background)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase5"
	if (int(ValueMask) & GCLineWidth) != 0 {
		// todo
		// field Field "line_width" in BitcaseType "xcb.ChangeGC.value_list.bitcase5"
		w.Write4b(_aux.LineWidth)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase6"
	if (int(ValueMask) & GCLineStyle) != 0 {
		// todo
		// field Field "line_style" in BitcaseType "xcb.ChangeGC.value_list.bitcase6"
		w.Write4b(_aux.LineStyle)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase7"
	if (int(ValueMask) & GCCapStyle) != 0 {
		// todo
		// field Field "cap_style" in BitcaseType "xcb.ChangeGC.value_list.bitcase7"
		w.Write4b(_aux.CapStyle)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase8"
	if (int(ValueMask) & GCJoinStyle) != 0 {
		// todo
		// field Field "join_style" in BitcaseType "xcb.ChangeGC.value_list.bitcase8"
		w.Write4b(_aux.JoinStyle)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase9"
	if (int(ValueMask) & GCFillStyle) != 0 {
		// todo
		// field Field "fill_style" in BitcaseType "xcb.ChangeGC.value_list.bitcase9"
		w.Write4b(_aux.FillStyle)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase10"
	if (int(ValueMask) & GCFillRule) != 0 {
		// todo
		// field Field "fill_rule" in BitcaseType "xcb.ChangeGC.value_list.bitcase10"
		w.Write4b(_aux.FillRule)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase11"
	if (int(ValueMask) & GCTile) != 0 {
		// todo
		// field Field "tile" in BitcaseType "xcb.ChangeGC.value_list.bitcase11"
		w.Write4b(uint32(_aux.Tile))
		n += 4
		alignTo = int(unsafe.Alignof(Pixmap(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase12"
	if (int(ValueMask) & GCStipple) != 0 {
		// todo
		// field Field "stipple" in BitcaseType "xcb.ChangeGC.value_list.bitcase12"
		w.Write4b(uint32(_aux.Stipple))
		n += 4
		alignTo = int(unsafe.Alignof(Pixmap(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase13"
	if (int(ValueMask) & GCTileStippleOriginX) != 0 {
		// todo
		// field Field "tile_stipple_x_origin" in BitcaseType "xcb.ChangeGC.value_list.bitcase13"
		w.Write4b(uint32(_aux.TileStippleXOrigin))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase14"
	if (int(ValueMask) & GCTileStippleOriginY) != 0 {
		// todo
		// field Field "tile_stipple_y_origin" in BitcaseType "xcb.ChangeGC.value_list.bitcase14"
		w.Write4b(uint32(_aux.TileStippleYOrigin))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase15"
	if (int(ValueMask) & GCFont) != 0 {
		// todo
		// field Field "font" in BitcaseType "xcb.ChangeGC.value_list.bitcase15"
		w.Write4b(uint32(_aux.Font))
		n += 4
		alignTo = int(unsafe.Alignof(Font(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase16"
	if (int(ValueMask) & GCSubwindowMode) != 0 {
		// todo
		// field Field "subwindow_mode" in BitcaseType "xcb.ChangeGC.value_list.bitcase16"
		w.Write4b(_aux.SubwindowMode)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase17"
	if (int(ValueMask) & GCGraphicsExposures) != 0 {
		// todo
		// field Field "graphics_exposures" in BitcaseType "xcb.ChangeGC.value_list.bitcase17"
		w.Write4b(uint32(_aux.GraphicsExposures))
		n += 4
		alignTo = int(unsafe.Alignof(Bool32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase18"
	if (int(ValueMask) & GCClipOriginX) != 0 {
		// todo
		// field Field "clip_x_origin" in BitcaseType "xcb.ChangeGC.value_list.bitcase18"
		w.Write4b(uint32(_aux.ClipXOrigin))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase19"
	if (int(ValueMask) & GCClipOriginY) != 0 {
		// todo
		// field Field "clip_y_origin" in BitcaseType "xcb.ChangeGC.value_list.bitcase19"
		w.Write4b(uint32(_aux.ClipYOrigin))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase20"
	if (int(ValueMask) & GCClipMask) != 0 {
		// todo
		// field Field "clip_mask" in BitcaseType "xcb.ChangeGC.value_list.bitcase20"
		w.Write4b(uint32(_aux.ClipMask))
		n += 4
		alignTo = int(unsafe.Alignof(Pixmap(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase21"
	if (int(ValueMask) & GCDashOffset) != 0 {
		// todo
		// field Field "dash_offset" in BitcaseType "xcb.ChangeGC.value_list.bitcase21"
		w.Write4b(_aux.DashOffset)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase22"
	if (int(ValueMask) & GCDashList) != 0 {
		// todo
		// field Field "dashes" in BitcaseType "xcb.ChangeGC.value_list.bitcase22"
		w.Write4b(_aux.Dashes)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeGC.value_list.bitcase23"
	if (int(ValueMask) & GCArcMode) != 0 {
		// todo
		// field Field "arc_mode" in BitcaseType "xcb.ChangeGC.value_list.bitcase23"
		w.Write4b(_aux.ArcMode)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	/* insert padding */
	xgbPad = -(n + xgbPaddingOffset) & (alignTo - 1)
	w.WritePad(xgbPad)
	xgbPad = 0
	n = 0
	xgbPaddingOffset = 0
}

const ChangeGCOpcode = 56

func ChangeGcRequestWrite(w *Writer, Gc GContext, ValueMask uint32, ValueList *ChangeGcValueList) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4

	// write wire field ValueMask
	w.Write4b(ValueMask)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field ValueList
	// switch serialize
	ChangeGcValueListSerialize(w, ValueMask, ValueList)
}

func ChangeGC(conn *Conn, Gc GContext, ValueMask uint32, ValueList *ChangeGcValueList) VoidCookie {
	w := NewWriter()
	ChangeGcRequestWrite(w, Gc, ValueMask, ValueList)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ChangeGCOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func ChangeGCChecked(conn *Conn, Gc GContext, ValueMask uint32, ValueList *ChangeGcValueList) VoidCookie {
	w := NewWriter()
	ChangeGcRequestWrite(w, Gc, ValueMask, ValueList)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ChangeGCOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const CopyGCOpcode = 57

func CopyGcRequestWrite(w *Writer, SrcGc GContext, DstGc GContext, ValueMask uint32) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field SrcGc
	w.Write4b(uint32(SrcGc))
	n += 4

	// write wire field DstGc
	w.Write4b(uint32(DstGc))
	n += 4

	// write wire field ValueMask
	w.Write4b(ValueMask)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CopyGC(conn *Conn, SrcGc GContext, DstGc GContext, ValueMask uint32) VoidCookie {
	w := NewWriter()
	CopyGcRequestWrite(w, SrcGc, DstGc, ValueMask)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CopyGCOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func CopyGCChecked(conn *Conn, SrcGc GContext, DstGc GContext, ValueMask uint32) VoidCookie {
	w := NewWriter()
	CopyGcRequestWrite(w, SrcGc, DstGc, ValueMask)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CopyGCOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const SetDashesOpcode = 58

func SetDashesRequestWrite(w *Writer, Gc GContext, DashOffset uint16, DashesLen uint16, Dashes []uint8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4

	// write wire field DashOffset
	w.Write2b(DashOffset)
	n += 2

	// write wire field DashesLen
	w.Write2b(DashesLen)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Dashes
	{
		_dataLen := int(DashesLen)
		for i := 0; i < _dataLen; i++ {
			w.Write1b(uint8(Dashes[i]))
		}
		n += _dataLen * 1
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SetDashes(conn *Conn, Gc GContext, DashOffset uint16, DashesLen uint16, Dashes []uint8) VoidCookie {
	w := NewWriter()
	SetDashesRequestWrite(w, Gc, DashOffset, DashesLen, Dashes)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: SetDashesOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func SetDashesChecked(conn *Conn, Gc GContext, DashOffset uint16, DashesLen uint16, Dashes []uint8) VoidCookie {
	w := NewWriter()
	SetDashesRequestWrite(w, Gc, DashOffset, DashesLen, Dashes)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: SetDashesOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum ClipOrdering
const (
	ClipOrderingUnsorted = 0
	ClipOrderingYSorted  = 1
	ClipOrderingYXSorted = 2
	ClipOrderingYXBanded = 3
)

const SetClipRectanglesOpcode = 59

func SetClipRectanglesRequestWrite(w *Writer, Ordering uint8, Gc GContext, ClipXOrigin int16, ClipYOrigin int16, RectanglesLen uint32, Rectangles []Rectangle) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Ordering
	w.Write1b(Ordering)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4

	// write wire field ClipXOrigin
	w.Write2b(uint16(ClipXOrigin))
	n += 2

	// write wire field ClipYOrigin
	w.Write2b(uint16(ClipYOrigin))
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Rectangles
	n += RectangleWriteN(w, Rectangles)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SetClipRectangles(conn *Conn, Ordering uint8, Gc GContext, ClipXOrigin int16, ClipYOrigin int16, RectanglesLen uint32, Rectangles []Rectangle) VoidCookie {
	w := NewWriter()
	SetClipRectanglesRequestWrite(w, Ordering, Gc, ClipXOrigin, ClipYOrigin, RectanglesLen, Rectangles)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: SetClipRectanglesOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func SetClipRectanglesChecked(conn *Conn, Ordering uint8, Gc GContext, ClipXOrigin int16, ClipYOrigin int16, RectanglesLen uint32, Rectangles []Rectangle) VoidCookie {
	w := NewWriter()
	SetClipRectanglesRequestWrite(w, Ordering, Gc, ClipXOrigin, ClipYOrigin, RectanglesLen, Rectangles)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: SetClipRectanglesOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const FreeGCOpcode = 60

func FreeGcRequestWrite(w *Writer, Gc GContext) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func FreeGC(conn *Conn, Gc GContext) VoidCookie {
	w := NewWriter()
	FreeGcRequestWrite(w, Gc)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: FreeGCOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func FreeGCChecked(conn *Conn, Gc GContext) VoidCookie {
	w := NewWriter()
	FreeGcRequestWrite(w, Gc)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: FreeGCOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const ClearAreaOpcode = 61

func ClearAreaRequestWrite(w *Writer, Exposures uint8, Window Window, X int16, Y int16, Width uint16, Height uint16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Exposures
	w.Write1b(Exposures)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4

	// write wire field X
	w.Write2b(uint16(X))
	n += 2

	// write wire field Y
	w.Write2b(uint16(Y))
	n += 2

	// write wire field Width
	w.Write2b(Width)
	n += 2

	// write wire field Height
	w.Write2b(Height)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func ClearArea(conn *Conn, Exposures uint8, Window Window, X int16, Y int16, Width uint16, Height uint16) VoidCookie {
	w := NewWriter()
	ClearAreaRequestWrite(w, Exposures, Window, X, Y, Width, Height)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ClearAreaOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func ClearAreaChecked(conn *Conn, Exposures uint8, Window Window, X int16, Y int16, Width uint16, Height uint16) VoidCookie {
	w := NewWriter()
	ClearAreaRequestWrite(w, Exposures, Window, X, Y, Width, Height)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ClearAreaOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const CopyAreaOpcode = 62

func CopyAreaRequestWrite(w *Writer, SrcDrawable Drawable, DstDrawable Drawable, Gc GContext, SrcX int16, SrcY int16, DstX int16, DstY int16, Width uint16, Height uint16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field SrcDrawable
	w.Write4b(uint32(SrcDrawable))
	n += 4

	// write wire field DstDrawable
	w.Write4b(uint32(DstDrawable))
	n += 4

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4

	// write wire field SrcX
	w.Write2b(uint16(SrcX))
	n += 2

	// write wire field SrcY
	w.Write2b(uint16(SrcY))
	n += 2

	// write wire field DstX
	w.Write2b(uint16(DstX))
	n += 2

	// write wire field DstY
	w.Write2b(uint16(DstY))
	n += 2

	// write wire field Width
	w.Write2b(Width)
	n += 2

	// write wire field Height
	w.Write2b(Height)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CopyArea(conn *Conn, SrcDrawable Drawable, DstDrawable Drawable, Gc GContext, SrcX int16, SrcY int16, DstX int16, DstY int16, Width uint16, Height uint16) VoidCookie {
	w := NewWriter()
	CopyAreaRequestWrite(w, SrcDrawable, DstDrawable, Gc, SrcX, SrcY, DstX, DstY, Width, Height)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CopyAreaOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func CopyAreaChecked(conn *Conn, SrcDrawable Drawable, DstDrawable Drawable, Gc GContext, SrcX int16, SrcY int16, DstX int16, DstY int16, Width uint16, Height uint16) VoidCookie {
	w := NewWriter()
	CopyAreaRequestWrite(w, SrcDrawable, DstDrawable, Gc, SrcX, SrcY, DstX, DstY, Width, Height)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CopyAreaOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const CopyPlaneOpcode = 63

func CopyPlaneRequestWrite(w *Writer, SrcDrawable Drawable, DstDrawable Drawable, Gc GContext, SrcX int16, SrcY int16, DstX int16, DstY int16, Width uint16, Height uint16, BitPlane uint32) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field SrcDrawable
	w.Write4b(uint32(SrcDrawable))
	n += 4

	// write wire field DstDrawable
	w.Write4b(uint32(DstDrawable))
	n += 4

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4

	// write wire field SrcX
	w.Write2b(uint16(SrcX))
	n += 2

	// write wire field SrcY
	w.Write2b(uint16(SrcY))
	n += 2

	// write wire field DstX
	w.Write2b(uint16(DstX))
	n += 2

	// write wire field DstY
	w.Write2b(uint16(DstY))
	n += 2

	// write wire field Width
	w.Write2b(Width)
	n += 2

	// write wire field Height
	w.Write2b(Height)
	n += 2

	// write wire field BitPlane
	w.Write4b(BitPlane)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CopyPlane(conn *Conn, SrcDrawable Drawable, DstDrawable Drawable, Gc GContext, SrcX int16, SrcY int16, DstX int16, DstY int16, Width uint16, Height uint16, BitPlane uint32) VoidCookie {
	w := NewWriter()
	CopyPlaneRequestWrite(w, SrcDrawable, DstDrawable, Gc, SrcX, SrcY, DstX, DstY, Width, Height, BitPlane)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CopyPlaneOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func CopyPlaneChecked(conn *Conn, SrcDrawable Drawable, DstDrawable Drawable, Gc GContext, SrcX int16, SrcY int16, DstX int16, DstY int16, Width uint16, Height uint16, BitPlane uint32) VoidCookie {
	w := NewWriter()
	CopyPlaneRequestWrite(w, SrcDrawable, DstDrawable, Gc, SrcX, SrcY, DstX, DstY, Width, Height, BitPlane)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CopyPlaneOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum CoordMode
const (
	CoordModeOrigin   = 0
	CoordModePrevious = 1
)

const PolyPointOpcode = 64

func PolyPointRequestWrite(w *Writer, CoordinateMode uint8, Drawable Drawable, Gc GContext, PointsLen uint32, Points []Point) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field CoordinateMode
	w.Write1b(CoordinateMode)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Points
	n += PointWriteN(w, Points)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func PolyPoint(conn *Conn, CoordinateMode uint8, Drawable Drawable, Gc GContext, PointsLen uint32, Points []Point) VoidCookie {
	w := NewWriter()
	PolyPointRequestWrite(w, CoordinateMode, Drawable, Gc, PointsLen, Points)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PolyPointOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func PolyPointChecked(conn *Conn, CoordinateMode uint8, Drawable Drawable, Gc GContext, PointsLen uint32, Points []Point) VoidCookie {
	w := NewWriter()
	PolyPointRequestWrite(w, CoordinateMode, Drawable, Gc, PointsLen, Points)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PolyPointOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const PolyLineOpcode = 65

func PolyLineRequestWrite(w *Writer, CoordinateMode uint8, Drawable Drawable, Gc GContext, PointsLen uint32, Points []Point) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field CoordinateMode
	w.Write1b(CoordinateMode)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Points
	n += PointWriteN(w, Points)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func PolyLine(conn *Conn, CoordinateMode uint8, Drawable Drawable, Gc GContext, PointsLen uint32, Points []Point) VoidCookie {
	w := NewWriter()
	PolyLineRequestWrite(w, CoordinateMode, Drawable, Gc, PointsLen, Points)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PolyLineOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func PolyLineChecked(conn *Conn, CoordinateMode uint8, Drawable Drawable, Gc GContext, PointsLen uint32, Points []Point) VoidCookie {
	w := NewWriter()
	PolyLineRequestWrite(w, CoordinateMode, Drawable, Gc, PointsLen, Points)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PolyLineOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

type Segment struct {
	X1 int16
	Y1 int16
	X2 int16
	Y2 int16
}

type CSegment struct {
	X1 int16
	Y1 int16
	X2 int16
	Y2 int16
}

func SegmentRead(r *Reader, v *Segment) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field X1
	v.X1 = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Y1
	v.Y1 = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field X2
	v.X2 = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Y2
	v.Y2 = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

func SegmentReadN(r *Reader, dest []Segment) (n int) {
	for i := 0; i < len(dest); i++ {
		n += SegmentRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func SegmentWrite(w *Writer, v *Segment) (n int) {

	// write field X1
	w.Write2b(uint16(v.X1))
	n += 2

	// write field Y1
	w.Write2b(uint16(v.Y1))
	n += 2

	// write field X2
	w.Write2b(uint16(v.X2))
	n += 2

	// write field Y2
	w.Write2b(uint16(v.Y2))
	n += 2
	return
}
func SegmentWriteN(w *Writer, src []Segment) (n int) {
	for i := 0; i < len(src); i++ {
		n += SegmentWrite(w, &src[i])
	}
	return
}

const PolySegmentOpcode = 66

func PolySegmentRequestWrite(w *Writer, Drawable Drawable, Gc GContext, SegmentsLen uint32, Segments []Segment) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Segments
	n += SegmentWriteN(w, Segments)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func PolySegment(conn *Conn, Drawable Drawable, Gc GContext, SegmentsLen uint32, Segments []Segment) VoidCookie {
	w := NewWriter()
	PolySegmentRequestWrite(w, Drawable, Gc, SegmentsLen, Segments)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PolySegmentOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func PolySegmentChecked(conn *Conn, Drawable Drawable, Gc GContext, SegmentsLen uint32, Segments []Segment) VoidCookie {
	w := NewWriter()
	PolySegmentRequestWrite(w, Drawable, Gc, SegmentsLen, Segments)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PolySegmentOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const PolyRectangleOpcode = 67

func PolyRectangleRequestWrite(w *Writer, Drawable Drawable, Gc GContext, RectanglesLen uint32, Rectangles []Rectangle) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Rectangles
	n += RectangleWriteN(w, Rectangles)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func PolyRectangle(conn *Conn, Drawable Drawable, Gc GContext, RectanglesLen uint32, Rectangles []Rectangle) VoidCookie {
	w := NewWriter()
	PolyRectangleRequestWrite(w, Drawable, Gc, RectanglesLen, Rectangles)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PolyRectangleOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func PolyRectangleChecked(conn *Conn, Drawable Drawable, Gc GContext, RectanglesLen uint32, Rectangles []Rectangle) VoidCookie {
	w := NewWriter()
	PolyRectangleRequestWrite(w, Drawable, Gc, RectanglesLen, Rectangles)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PolyRectangleOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const PolyArcOpcode = 68

func PolyArcRequestWrite(w *Writer, Drawable Drawable, Gc GContext, ArcsLen uint32, Arcs []Arc) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Arcs
	n += ArcWriteN(w, Arcs)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func PolyArc(conn *Conn, Drawable Drawable, Gc GContext, ArcsLen uint32, Arcs []Arc) VoidCookie {
	w := NewWriter()
	PolyArcRequestWrite(w, Drawable, Gc, ArcsLen, Arcs)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PolyArcOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func PolyArcChecked(conn *Conn, Drawable Drawable, Gc GContext, ArcsLen uint32, Arcs []Arc) VoidCookie {
	w := NewWriter()
	PolyArcRequestWrite(w, Drawable, Gc, ArcsLen, Arcs)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PolyArcOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum PolyShape
const (
	PolyShapeComplex   = 0
	PolyShapeNonconvex = 1
	PolyShapeConvex    = 2
)

const FillPolyOpcode = 69

func FillPolyRequestWrite(w *Writer, Drawable Drawable, Gc GContext, Shape uint8, CoordinateMode uint8, PointsLen uint32, Points []Point) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4

	// write wire field Shape
	w.Write1b(Shape)
	n += 1

	// write wire field CoordinateMode
	w.Write1b(CoordinateMode)
	n += 1

	// write wire field Pad1
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Points
	n += PointWriteN(w, Points)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func FillPoly(conn *Conn, Drawable Drawable, Gc GContext, Shape uint8, CoordinateMode uint8, PointsLen uint32, Points []Point) VoidCookie {
	w := NewWriter()
	FillPolyRequestWrite(w, Drawable, Gc, Shape, CoordinateMode, PointsLen, Points)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: FillPolyOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func FillPolyChecked(conn *Conn, Drawable Drawable, Gc GContext, Shape uint8, CoordinateMode uint8, PointsLen uint32, Points []Point) VoidCookie {
	w := NewWriter()
	FillPolyRequestWrite(w, Drawable, Gc, Shape, CoordinateMode, PointsLen, Points)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: FillPolyOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const PolyFillRectangleOpcode = 70

func PolyFillRectangleRequestWrite(w *Writer, Drawable Drawable, Gc GContext, RectanglesLen uint32, Rectangles []Rectangle) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Rectangles
	n += RectangleWriteN(w, Rectangles)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func PolyFillRectangle(conn *Conn, Drawable Drawable, Gc GContext, RectanglesLen uint32, Rectangles []Rectangle) VoidCookie {
	w := NewWriter()
	PolyFillRectangleRequestWrite(w, Drawable, Gc, RectanglesLen, Rectangles)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PolyFillRectangleOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func PolyFillRectangleChecked(conn *Conn, Drawable Drawable, Gc GContext, RectanglesLen uint32, Rectangles []Rectangle) VoidCookie {
	w := NewWriter()
	PolyFillRectangleRequestWrite(w, Drawable, Gc, RectanglesLen, Rectangles)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PolyFillRectangleOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const PolyFillArcOpcode = 71

func PolyFillArcRequestWrite(w *Writer, Drawable Drawable, Gc GContext, ArcsLen uint32, Arcs []Arc) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Arcs
	n += ArcWriteN(w, Arcs)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func PolyFillArc(conn *Conn, Drawable Drawable, Gc GContext, ArcsLen uint32, Arcs []Arc) VoidCookie {
	w := NewWriter()
	PolyFillArcRequestWrite(w, Drawable, Gc, ArcsLen, Arcs)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PolyFillArcOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func PolyFillArcChecked(conn *Conn, Drawable Drawable, Gc GContext, ArcsLen uint32, Arcs []Arc) VoidCookie {
	w := NewWriter()
	PolyFillArcRequestWrite(w, Drawable, Gc, ArcsLen, Arcs)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PolyFillArcOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum ImageFormat
const (
	ImageFormatXYBitmap = 0
	ImageFormatXYPixmap = 1
	ImageFormatZPixmap  = 2
)

const PutImageOpcode = 72

func PutImageRequestWrite(w *Writer, Format uint8, Drawable Drawable, Gc GContext, Width uint16, Height uint16, DstX int16, DstY int16, LeftPad uint8, Depth uint8, DataLen uint32, Data []uint8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Format
	w.Write1b(Format)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4

	// write wire field Width
	w.Write2b(Width)
	n += 2

	// write wire field Height
	w.Write2b(Height)
	n += 2

	// write wire field DstX
	w.Write2b(uint16(DstX))
	n += 2

	// write wire field DstY
	w.Write2b(uint16(DstY))
	n += 2

	// write wire field LeftPad
	w.Write1b(LeftPad)
	n += 1

	// write wire field Depth
	w.Write1b(Depth)
	n += 1

	// write wire field Pad0
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Data
	{
		_dataLen := int(DataLen)
		for i := 0; i < _dataLen; i++ {
			w.Write1b(uint8(Data[i]))
		}
		n += _dataLen * 1
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func PutImage(conn *Conn, Format uint8, Drawable Drawable, Gc GContext, Width uint16, Height uint16, DstX int16, DstY int16, LeftPad uint8, Depth uint8, DataLen uint32, Data []uint8) VoidCookie {
	w := NewWriter()
	PutImageRequestWrite(w, Format, Drawable, Gc, Width, Height, DstX, DstY, LeftPad, Depth, DataLen, Data)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PutImageOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func PutImageChecked(conn *Conn, Format uint8, Drawable Drawable, Gc GContext, Width uint16, Height uint16, DstX int16, DstY int16, LeftPad uint8, Depth uint8, DataLen uint32, Data []uint8) VoidCookie {
	w := NewWriter()
	PutImageRequestWrite(w, Format, Drawable, Gc, Width, Height, DstX, DstY, LeftPad, Depth, DataLen, Data)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PutImageOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const GetImageOpcode = 73

func GetImageRequestWrite(w *Writer, Format uint8, Drawable Drawable, X int16, Y int16, Width uint16, Height uint16, PlaneMask uint32) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Format
	w.Write1b(Format)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field X
	w.Write2b(uint16(X))
	n += 2

	// write wire field Y
	w.Write2b(uint16(Y))
	n += 2

	// write wire field Width
	w.Write2b(Width)
	n += 2

	// write wire field Height
	w.Write2b(Height)
	n += 2

	// write wire field PlaneMask
	w.Write4b(PlaneMask)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetImageReply struct {
	ResponseType uint8
	Depth        uint8
	Sequence     uint16
	Length       uint32
	Visual       VisualID
	// Pad0 [20]uint8
	Data []uint8
}

func GetImageReplyRead(r *Reader, v *GetImageReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Depth
	v.Depth = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Visual
	v.Visual = VisualID(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(20)
	n += 20
	if r.Err() != nil {
		return
	}

	// read field Data
	{
		dataLen := int((int(v.Length) * 4))
		data := make([]uint8, dataLen)
		for i := 0; i < dataLen; i++ {
			data[i] = uint8(r.Read1b())
			if r.Err() != nil {
				return
			}
		}
		blockLen += dataLen
		n += blockLen
		v.Data = data
	}
	alignTo = int(unsafe.Alignof(uint8(0)))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type GetImageCookie uint64

func (cookie GetImageCookie) Reply(conn *Conn) (*GetImageReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply GetImageReply
	GetImageReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetImage(conn *Conn, Format uint8, Drawable Drawable, X int16, Y int16, Width uint16, Height uint16, PlaneMask uint32) GetImageCookie {
	w := NewWriter()
	GetImageRequestWrite(w, Format, Drawable, X, Y, Width, Height, PlaneMask)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetImageOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return GetImageCookie(seq)
}

func GetImageUnchecked(conn *Conn, Format uint8, Drawable Drawable, X int16, Y int16, Width uint16, Height uint16, PlaneMask uint32) GetImageCookie {
	w := NewWriter()
	GetImageRequestWrite(w, Format, Drawable, X, Y, Width, Height, PlaneMask)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetImageOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetImageCookie(seq)
}

const PolyText8Opcode = 74

func PolyText8RequestWrite(w *Writer, Drawable Drawable, Gc GContext, X int16, Y int16, ItemsLen uint32, Items []uint8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4

	// write wire field X
	w.Write2b(uint16(X))
	n += 2

	// write wire field Y
	w.Write2b(uint16(Y))
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Items
	{
		_dataLen := int(ItemsLen)
		for i := 0; i < _dataLen; i++ {
			w.Write1b(uint8(Items[i]))
		}
		n += _dataLen * 1
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func PolyText8(conn *Conn, Drawable Drawable, Gc GContext, X int16, Y int16, ItemsLen uint32, Items []uint8) VoidCookie {
	w := NewWriter()
	PolyText8RequestWrite(w, Drawable, Gc, X, Y, ItemsLen, Items)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PolyText8Opcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func PolyText8Checked(conn *Conn, Drawable Drawable, Gc GContext, X int16, Y int16, ItemsLen uint32, Items []uint8) VoidCookie {
	w := NewWriter()
	PolyText8RequestWrite(w, Drawable, Gc, X, Y, ItemsLen, Items)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PolyText8Opcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const PolyText16Opcode = 75

func PolyText16RequestWrite(w *Writer, Drawable Drawable, Gc GContext, X int16, Y int16, ItemsLen uint32, Items []uint8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4

	// write wire field X
	w.Write2b(uint16(X))
	n += 2

	// write wire field Y
	w.Write2b(uint16(Y))
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Items
	{
		_dataLen := int(ItemsLen)
		for i := 0; i < _dataLen; i++ {
			w.Write1b(uint8(Items[i]))
		}
		n += _dataLen * 1
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func PolyText16(conn *Conn, Drawable Drawable, Gc GContext, X int16, Y int16, ItemsLen uint32, Items []uint8) VoidCookie {
	w := NewWriter()
	PolyText16RequestWrite(w, Drawable, Gc, X, Y, ItemsLen, Items)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PolyText16Opcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func PolyText16Checked(conn *Conn, Drawable Drawable, Gc GContext, X int16, Y int16, ItemsLen uint32, Items []uint8) VoidCookie {
	w := NewWriter()
	PolyText16RequestWrite(w, Drawable, Gc, X, Y, ItemsLen, Items)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: PolyText16Opcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const ImageText8Opcode = 76

func ImageText8RequestWrite(w *Writer, StringLen uint8, Drawable Drawable, Gc GContext, X int16, Y int16, String string) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field StringLen
	w.Write1b(StringLen)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4

	// write wire field X
	w.Write2b(uint16(X))
	n += 2

	// write wire field Y
	w.Write2b(uint16(Y))
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field String
	w.WriteString(String)
	n += len(String)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func ImageText8(conn *Conn, StringLen uint8, Drawable Drawable, Gc GContext, X int16, Y int16, String string) VoidCookie {
	w := NewWriter()
	ImageText8RequestWrite(w, StringLen, Drawable, Gc, X, Y, String)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ImageText8Opcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func ImageText8Checked(conn *Conn, StringLen uint8, Drawable Drawable, Gc GContext, X int16, Y int16, String string) VoidCookie {
	w := NewWriter()
	ImageText8RequestWrite(w, StringLen, Drawable, Gc, X, Y, String)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ImageText8Opcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const ImageText16Opcode = 77

func ImageText16RequestWrite(w *Writer, StringLen uint8, Drawable Drawable, Gc GContext, X int16, Y int16, String []Char2B) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field StringLen
	w.Write1b(StringLen)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4

	// write wire field X
	w.Write2b(uint16(X))
	n += 2

	// write wire field Y
	w.Write2b(uint16(Y))
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field String
	n += Char2BWriteN(w, String)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func ImageText16(conn *Conn, StringLen uint8, Drawable Drawable, Gc GContext, X int16, Y int16, String []Char2B) VoidCookie {
	w := NewWriter()
	ImageText16RequestWrite(w, StringLen, Drawable, Gc, X, Y, String)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ImageText16Opcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func ImageText16Checked(conn *Conn, StringLen uint8, Drawable Drawable, Gc GContext, X int16, Y int16, String []Char2B) VoidCookie {
	w := NewWriter()
	ImageText16RequestWrite(w, StringLen, Drawable, Gc, X, Y, String)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ImageText16Opcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum ColormapAlloc
const (
	ColormapAllocNone = 0
	ColormapAllocAll  = 1
)

const CreateColormapOpcode = 78

func CreateColormapRequestWrite(w *Writer, Alloc uint8, Mid Colormap, Window Window, Visual VisualID) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Alloc
	w.Write1b(Alloc)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Mid
	w.Write4b(uint32(Mid))
	n += 4

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4

	// write wire field Visual
	w.Write4b(uint32(Visual))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CreateColormap(conn *Conn, Alloc uint8, Mid Colormap, Window Window, Visual VisualID) VoidCookie {
	w := NewWriter()
	CreateColormapRequestWrite(w, Alloc, Mid, Window, Visual)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CreateColormapOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func CreateColormapChecked(conn *Conn, Alloc uint8, Mid Colormap, Window Window, Visual VisualID) VoidCookie {
	w := NewWriter()
	CreateColormapRequestWrite(w, Alloc, Mid, Window, Visual)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CreateColormapOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const FreeColormapOpcode = 79

func FreeColormapRequestWrite(w *Writer, Cmap Colormap) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cmap
	w.Write4b(uint32(Cmap))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func FreeColormap(conn *Conn, Cmap Colormap) VoidCookie {
	w := NewWriter()
	FreeColormapRequestWrite(w, Cmap)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: FreeColormapOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func FreeColormapChecked(conn *Conn, Cmap Colormap) VoidCookie {
	w := NewWriter()
	FreeColormapRequestWrite(w, Cmap)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: FreeColormapOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const CopyColormapAndFreeOpcode = 80

func CopyColormapAndFreeRequestWrite(w *Writer, Mid Colormap, SrcCmap Colormap) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Mid
	w.Write4b(uint32(Mid))
	n += 4

	// write wire field SrcCmap
	w.Write4b(uint32(SrcCmap))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CopyColormapAndFree(conn *Conn, Mid Colormap, SrcCmap Colormap) VoidCookie {
	w := NewWriter()
	CopyColormapAndFreeRequestWrite(w, Mid, SrcCmap)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CopyColormapAndFreeOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func CopyColormapAndFreeChecked(conn *Conn, Mid Colormap, SrcCmap Colormap) VoidCookie {
	w := NewWriter()
	CopyColormapAndFreeRequestWrite(w, Mid, SrcCmap)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CopyColormapAndFreeOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const InstallColormapOpcode = 81

func InstallColormapRequestWrite(w *Writer, Cmap Colormap) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cmap
	w.Write4b(uint32(Cmap))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func InstallColormap(conn *Conn, Cmap Colormap) VoidCookie {
	w := NewWriter()
	InstallColormapRequestWrite(w, Cmap)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: InstallColormapOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func InstallColormapChecked(conn *Conn, Cmap Colormap) VoidCookie {
	w := NewWriter()
	InstallColormapRequestWrite(w, Cmap)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: InstallColormapOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const UninstallColormapOpcode = 82

func UninstallColormapRequestWrite(w *Writer, Cmap Colormap) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cmap
	w.Write4b(uint32(Cmap))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func UninstallColormap(conn *Conn, Cmap Colormap) VoidCookie {
	w := NewWriter()
	UninstallColormapRequestWrite(w, Cmap)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: UninstallColormapOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func UninstallColormapChecked(conn *Conn, Cmap Colormap) VoidCookie {
	w := NewWriter()
	UninstallColormapRequestWrite(w, Cmap)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: UninstallColormapOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const ListInstalledColormapsOpcode = 83

func ListInstalledColormapsRequestWrite(w *Writer, Window Window) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type ListInstalledColormapsReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Length   uint32
	CmapsLen uint16
	// Pad1 [22]uint8
	Cmaps []Colormap
}

func ListInstalledColormapsReplyRead(r *Reader, v *ListInstalledColormapsReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field CmapsLen
	v.CmapsLen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(22)
	n += 22
	if r.Err() != nil {
		return
	}

	// read field Cmaps
	{
		dataLen := int(int(v.CmapsLen))
		data := make([]Colormap, dataLen)
		for i := 0; i < dataLen; i++ {
			data[i] = Colormap(r.Read4b())
			if r.Err() != nil {
				return
			}
		}
		blockLen += dataLen * 4
		n += blockLen
		v.Cmaps = data
	}
	alignTo = int(unsafe.Alignof(Colormap(0)))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type ListInstalledColormapsCookie uint64

func (cookie ListInstalledColormapsCookie) Reply(conn *Conn) (*ListInstalledColormapsReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply ListInstalledColormapsReply
	ListInstalledColormapsReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func ListInstalledColormaps(conn *Conn, Window Window) ListInstalledColormapsCookie {
	w := NewWriter()
	ListInstalledColormapsRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: ListInstalledColormapsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return ListInstalledColormapsCookie(seq)
}

func ListInstalledColormapsUnchecked(conn *Conn, Window Window) ListInstalledColormapsCookie {
	w := NewWriter()
	ListInstalledColormapsRequestWrite(w, Window)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: ListInstalledColormapsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return ListInstalledColormapsCookie(seq)
}

const AllocColorOpcode = 84

func AllocColorRequestWrite(w *Writer, Cmap Colormap, Red uint16, Green uint16, Blue uint16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cmap
	w.Write4b(uint32(Cmap))
	n += 4

	// write wire field Red
	w.Write2b(Red)
	n += 2

	// write wire field Green
	w.Write2b(Green)
	n += 2

	// write wire field Blue
	w.Write2b(Blue)
	n += 2

	// write wire field Pad1
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type AllocColorReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Length   uint32
	Red      uint16
	Green    uint16
	Blue     uint16
	// Pad1 [2]uint8
	Pixel uint32
}

func AllocColorReplyRead(r *Reader, v *AllocColorReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Red
	v.Red = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Green
	v.Green = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Blue
	v.Blue = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(2)
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pixel
	v.Pixel = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}
	return
}

type AllocColorCookie uint64

func (cookie AllocColorCookie) Reply(conn *Conn) (*AllocColorReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply AllocColorReply
	AllocColorReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func AllocColor(conn *Conn, Cmap Colormap, Red uint16, Green uint16, Blue uint16) AllocColorCookie {
	w := NewWriter()
	AllocColorRequestWrite(w, Cmap, Red, Green, Blue)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: AllocColorOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return AllocColorCookie(seq)
}

func AllocColorUnchecked(conn *Conn, Cmap Colormap, Red uint16, Green uint16, Blue uint16) AllocColorCookie {
	w := NewWriter()
	AllocColorRequestWrite(w, Cmap, Red, Green, Blue)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: AllocColorOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return AllocColorCookie(seq)
}

const AllocNamedColorOpcode = 85

func AllocNamedColorRequestWrite(w *Writer, Cmap Colormap, NameLen uint16, Name string) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cmap
	w.Write4b(uint32(Cmap))
	n += 4

	// write wire field NameLen
	w.Write2b(NameLen)
	n += 2

	// write wire field Pad1
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Name
	w.WriteString(Name)
	n += len(Name)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type AllocNamedColorReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence    uint16
	Length      uint32
	Pixel       uint32
	ExactRed    uint16
	ExactGreen  uint16
	ExactBlue   uint16
	VisualRed   uint16
	VisualGreen uint16
	VisualBlue  uint16
}

func AllocNamedColorReplyRead(r *Reader, v *AllocNamedColorReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pixel
	v.Pixel = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field ExactRed
	v.ExactRed = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field ExactGreen
	v.ExactGreen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field ExactBlue
	v.ExactBlue = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field VisualRed
	v.VisualRed = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field VisualGreen
	v.VisualGreen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field VisualBlue
	v.VisualBlue = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

type AllocNamedColorCookie uint64

func (cookie AllocNamedColorCookie) Reply(conn *Conn) (*AllocNamedColorReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply AllocNamedColorReply
	AllocNamedColorReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func AllocNamedColor(conn *Conn, Cmap Colormap, NameLen uint16, Name string) AllocNamedColorCookie {
	w := NewWriter()
	AllocNamedColorRequestWrite(w, Cmap, NameLen, Name)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: AllocNamedColorOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return AllocNamedColorCookie(seq)
}

func AllocNamedColorUnchecked(conn *Conn, Cmap Colormap, NameLen uint16, Name string) AllocNamedColorCookie {
	w := NewWriter()
	AllocNamedColorRequestWrite(w, Cmap, NameLen, Name)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: AllocNamedColorOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return AllocNamedColorCookie(seq)
}

const AllocColorCellsOpcode = 86

func AllocColorCellsRequestWrite(w *Writer, Contiguous uint8, Cmap Colormap, Colors uint16, Planes uint16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Contiguous
	w.Write1b(Contiguous)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cmap
	w.Write4b(uint32(Cmap))
	n += 4

	// write wire field Colors
	w.Write2b(Colors)
	n += 2

	// write wire field Planes
	w.Write2b(Planes)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type AllocColorCellsReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence  uint16
	Length    uint32
	PixelsLen uint16
	MasksLen  uint16
	// Pad1 [20]uint8
	Pixels []uint32
	Masks  []uint32
}

func AllocColorCellsReplyRead(r *Reader, v *AllocColorCellsReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field PixelsLen
	v.PixelsLen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field MasksLen
	v.MasksLen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(20)
	n += 20
	if r.Err() != nil {
		return
	}

	// read field Pixels
	{
		dataLen := int(int(v.PixelsLen))
		data := make([]uint32, dataLen)
		for i := 0; i < dataLen; i++ {
			data[i] = uint32(r.Read4b())
			if r.Err() != nil {
				return
			}
		}
		blockLen += dataLen * 4
		n += blockLen
		v.Pixels = data
	}
	alignTo = int(unsafe.Alignof(uint32(0)))

	// read field Masks

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	{
		dataLen := int(int(v.MasksLen))
		data := make([]uint32, dataLen)
		for i := 0; i < dataLen; i++ {
			data[i] = uint32(r.Read4b())
			if r.Err() != nil {
				return
			}
		}
		blockLen += dataLen * 4
		n += blockLen
		v.Masks = data
	}
	alignTo = int(unsafe.Alignof(uint32(0)))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type AllocColorCellsCookie uint64

func (cookie AllocColorCellsCookie) Reply(conn *Conn) (*AllocColorCellsReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply AllocColorCellsReply
	AllocColorCellsReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func AllocColorCells(conn *Conn, Contiguous uint8, Cmap Colormap, Colors uint16, Planes uint16) AllocColorCellsCookie {
	w := NewWriter()
	AllocColorCellsRequestWrite(w, Contiguous, Cmap, Colors, Planes)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: AllocColorCellsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return AllocColorCellsCookie(seq)
}

func AllocColorCellsUnchecked(conn *Conn, Contiguous uint8, Cmap Colormap, Colors uint16, Planes uint16) AllocColorCellsCookie {
	w := NewWriter()
	AllocColorCellsRequestWrite(w, Contiguous, Cmap, Colors, Planes)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: AllocColorCellsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return AllocColorCellsCookie(seq)
}

const AllocColorPlanesOpcode = 87

func AllocColorPlanesRequestWrite(w *Writer, Contiguous uint8, Cmap Colormap, Colors uint16, Reds uint16, Greens uint16, Blues uint16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Contiguous
	w.Write1b(Contiguous)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cmap
	w.Write4b(uint32(Cmap))
	n += 4

	// write wire field Colors
	w.Write2b(Colors)
	n += 2

	// write wire field Reds
	w.Write2b(Reds)
	n += 2

	// write wire field Greens
	w.Write2b(Greens)
	n += 2

	// write wire field Blues
	w.Write2b(Blues)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type AllocColorPlanesReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence  uint16
	Length    uint32
	PixelsLen uint16
	// Pad1 [2]uint8
	RedMask   uint32
	GreenMask uint32
	BlueMask  uint32
	// Pad2 [8]uint8
	Pixels []uint32
}

func AllocColorPlanesReplyRead(r *Reader, v *AllocColorPlanesReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field PixelsLen
	v.PixelsLen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(2)
	n += 2
	if r.Err() != nil {
		return
	}

	// read field RedMask
	v.RedMask = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field GreenMask
	v.GreenMask = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field BlueMask
	v.BlueMask = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pad2
	r.ReadPad(8)
	n += 8
	if r.Err() != nil {
		return
	}

	// read field Pixels
	{
		dataLen := int(int(v.PixelsLen))
		data := make([]uint32, dataLen)
		for i := 0; i < dataLen; i++ {
			data[i] = uint32(r.Read4b())
			if r.Err() != nil {
				return
			}
		}
		blockLen += dataLen * 4
		n += blockLen
		v.Pixels = data
	}
	alignTo = int(unsafe.Alignof(uint32(0)))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type AllocColorPlanesCookie uint64

func (cookie AllocColorPlanesCookie) Reply(conn *Conn) (*AllocColorPlanesReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply AllocColorPlanesReply
	AllocColorPlanesReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func AllocColorPlanes(conn *Conn, Contiguous uint8, Cmap Colormap, Colors uint16, Reds uint16, Greens uint16, Blues uint16) AllocColorPlanesCookie {
	w := NewWriter()
	AllocColorPlanesRequestWrite(w, Contiguous, Cmap, Colors, Reds, Greens, Blues)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: AllocColorPlanesOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return AllocColorPlanesCookie(seq)
}

func AllocColorPlanesUnchecked(conn *Conn, Contiguous uint8, Cmap Colormap, Colors uint16, Reds uint16, Greens uint16, Blues uint16) AllocColorPlanesCookie {
	w := NewWriter()
	AllocColorPlanesRequestWrite(w, Contiguous, Cmap, Colors, Reds, Greens, Blues)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: AllocColorPlanesOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return AllocColorPlanesCookie(seq)
}

const FreeColorsOpcode = 88

func FreeColorsRequestWrite(w *Writer, Cmap Colormap, PlaneMask uint32, PixelsLen uint32, Pixels []uint32) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cmap
	w.Write4b(uint32(Cmap))
	n += 4

	// write wire field PlaneMask
	w.Write4b(PlaneMask)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Pixels
	{
		_dataLen := int(PixelsLen)
		for i := 0; i < _dataLen; i++ {
			w.Write4b(uint32(Pixels[i]))
		}
		n += _dataLen * 4
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func FreeColors(conn *Conn, Cmap Colormap, PlaneMask uint32, PixelsLen uint32, Pixels []uint32) VoidCookie {
	w := NewWriter()
	FreeColorsRequestWrite(w, Cmap, PlaneMask, PixelsLen, Pixels)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: FreeColorsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func FreeColorsChecked(conn *Conn, Cmap Colormap, PlaneMask uint32, PixelsLen uint32, Pixels []uint32) VoidCookie {
	w := NewWriter()
	FreeColorsRequestWrite(w, Cmap, PlaneMask, PixelsLen, Pixels)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: FreeColorsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum ColorFlag
const (
	ColorFlagRed   = 1
	ColorFlagGreen = 2
	ColorFlagBlue  = 4
)

type ColorItem struct {
	Pixel uint32
	Red   uint16
	Green uint16
	Blue  uint16
	Flags uint8
	// Pad0 uint8
}

type CColorItem struct {
	Pixel uint32
	Red   uint16
	Green uint16
	Blue  uint16
	Flags uint8
	Pad0  uint8
}

func ColorItemRead(r *Reader, v *ColorItem) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Pixel
	v.Pixel = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Red
	v.Red = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Green
	v.Green = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Blue
	v.Blue = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Flags
	v.Flags = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}
	return
}

func ColorItemReadN(r *Reader, dest []ColorItem) (n int) {
	for i := 0; i < len(dest); i++ {
		n += ColorItemRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func ColorItemWrite(w *Writer, v *ColorItem) (n int) {

	// write field Pixel
	w.Write4b(v.Pixel)
	n += 4

	// write field Red
	w.Write2b(v.Red)
	n += 2

	// write field Green
	w.Write2b(v.Green)
	n += 2

	// write field Blue
	w.Write2b(v.Blue)
	n += 2

	// write field Flags
	w.Write1b(v.Flags)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1
	return
}
func ColorItemWriteN(w *Writer, src []ColorItem) (n int) {
	for i := 0; i < len(src); i++ {
		n += ColorItemWrite(w, &src[i])
	}
	return
}

const StoreColorsOpcode = 89

func StoreColorsRequestWrite(w *Writer, Cmap Colormap, ItemsLen uint32, Items []ColorItem) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cmap
	w.Write4b(uint32(Cmap))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Items
	n += ColorItemWriteN(w, Items)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func StoreColors(conn *Conn, Cmap Colormap, ItemsLen uint32, Items []ColorItem) VoidCookie {
	w := NewWriter()
	StoreColorsRequestWrite(w, Cmap, ItemsLen, Items)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: StoreColorsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func StoreColorsChecked(conn *Conn, Cmap Colormap, ItemsLen uint32, Items []ColorItem) VoidCookie {
	w := NewWriter()
	StoreColorsRequestWrite(w, Cmap, ItemsLen, Items)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: StoreColorsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const StoreNamedColorOpcode = 90

func StoreNamedColorRequestWrite(w *Writer, Flags uint8, Cmap Colormap, Pixel uint32, NameLen uint16, Name string) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Flags
	w.Write1b(Flags)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cmap
	w.Write4b(uint32(Cmap))
	n += 4

	// write wire field Pixel
	w.Write4b(Pixel)
	n += 4

	// write wire field NameLen
	w.Write2b(NameLen)
	n += 2

	// write wire field Pad0
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Name
	w.WriteString(Name)
	n += len(Name)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func StoreNamedColor(conn *Conn, Flags uint8, Cmap Colormap, Pixel uint32, NameLen uint16, Name string) VoidCookie {
	w := NewWriter()
	StoreNamedColorRequestWrite(w, Flags, Cmap, Pixel, NameLen, Name)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: StoreNamedColorOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func StoreNamedColorChecked(conn *Conn, Flags uint8, Cmap Colormap, Pixel uint32, NameLen uint16, Name string) VoidCookie {
	w := NewWriter()
	StoreNamedColorRequestWrite(w, Flags, Cmap, Pixel, NameLen, Name)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: StoreNamedColorOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

type RGB struct {
	Red   uint16
	Green uint16
	Blue  uint16
	// Pad0 [2]uint8
}

type CRGB struct {
	Red   uint16
	Green uint16
	Blue  uint16
	Pad0  [2]uint8
}

func RGBRead(r *Reader, v *RGB) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Red
	v.Red = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Green
	v.Green = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Blue
	v.Blue = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(2)
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

func RGBReadN(r *Reader, dest []RGB) (n int) {
	for i := 0; i < len(dest); i++ {
		n += RGBRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func RGBWrite(w *Writer, v *RGB) (n int) {

	// write field Red
	w.Write2b(v.Red)
	n += 2

	// write field Green
	w.Write2b(v.Green)
	n += 2

	// write field Blue
	w.Write2b(v.Blue)
	n += 2

	// write field Pad0
	w.WritePad(2)
	n += 2
	return
}
func RGBWriteN(w *Writer, src []RGB) (n int) {
	for i := 0; i < len(src); i++ {
		n += RGBWrite(w, &src[i])
	}
	return
}

const QueryColorsOpcode = 91

func QueryColorsRequestWrite(w *Writer, Cmap Colormap, PixelsLen uint32, Pixels []uint32) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cmap
	w.Write4b(uint32(Cmap))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Pixels
	{
		_dataLen := int(PixelsLen)
		for i := 0; i < _dataLen; i++ {
			w.Write4b(uint32(Pixels[i]))
		}
		n += _dataLen * 4
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type QueryColorsReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence  uint16
	Length    uint32
	ColorsLen uint16
	// Pad1 [22]uint8
	Colors []RGB
}

func QueryColorsReplyRead(r *Reader, v *QueryColorsReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field ColorsLen
	v.ColorsLen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(22)
	n += 22
	if r.Err() != nil {
		return
	}

	// read field Colors
	v.Colors = make([]RGB, int(v.ColorsLen))
	blockLen += RGBReadN(r, v.Colors)
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CRGB{}))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type QueryColorsCookie uint64

func (cookie QueryColorsCookie) Reply(conn *Conn) (*QueryColorsReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply QueryColorsReply
	QueryColorsReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryColors(conn *Conn, Cmap Colormap, PixelsLen uint32, Pixels []uint32) QueryColorsCookie {
	w := NewWriter()
	QueryColorsRequestWrite(w, Cmap, PixelsLen, Pixels)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: QueryColorsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return QueryColorsCookie(seq)
}

func QueryColorsUnchecked(conn *Conn, Cmap Colormap, PixelsLen uint32, Pixels []uint32) QueryColorsCookie {
	w := NewWriter()
	QueryColorsRequestWrite(w, Cmap, PixelsLen, Pixels)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: QueryColorsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return QueryColorsCookie(seq)
}

const LookupColorOpcode = 92

func LookupColorRequestWrite(w *Writer, Cmap Colormap, NameLen uint16, Name string) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cmap
	w.Write4b(uint32(Cmap))
	n += 4

	// write wire field NameLen
	w.Write2b(NameLen)
	n += 2

	// write wire field Pad1
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Name
	w.WriteString(Name)
	n += len(Name)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type LookupColorReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence    uint16
	Length      uint32
	ExactRed    uint16
	ExactGreen  uint16
	ExactBlue   uint16
	VisualRed   uint16
	VisualGreen uint16
	VisualBlue  uint16
}

func LookupColorReplyRead(r *Reader, v *LookupColorReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field ExactRed
	v.ExactRed = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field ExactGreen
	v.ExactGreen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field ExactBlue
	v.ExactBlue = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field VisualRed
	v.VisualRed = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field VisualGreen
	v.VisualGreen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field VisualBlue
	v.VisualBlue = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

type LookupColorCookie uint64

func (cookie LookupColorCookie) Reply(conn *Conn) (*LookupColorReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply LookupColorReply
	LookupColorReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func LookupColor(conn *Conn, Cmap Colormap, NameLen uint16, Name string) LookupColorCookie {
	w := NewWriter()
	LookupColorRequestWrite(w, Cmap, NameLen, Name)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: LookupColorOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return LookupColorCookie(seq)
}

func LookupColorUnchecked(conn *Conn, Cmap Colormap, NameLen uint16, Name string) LookupColorCookie {
	w := NewWriter()
	LookupColorRequestWrite(w, Cmap, NameLen, Name)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: LookupColorOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return LookupColorCookie(seq)
}

// enum Pixmap
const (
	PixmapNone = 0
)

const CreateCursorOpcode = 93

func CreateCursorRequestWrite(w *Writer, Cid Cursor, Source Pixmap, Mask Pixmap, ForeRed uint16, ForeGreen uint16, ForeBlue uint16, BackRed uint16, BackGreen uint16, BackBlue uint16, X uint16, Y uint16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cid
	w.Write4b(uint32(Cid))
	n += 4

	// write wire field Source
	w.Write4b(uint32(Source))
	n += 4

	// write wire field Mask
	w.Write4b(uint32(Mask))
	n += 4

	// write wire field ForeRed
	w.Write2b(ForeRed)
	n += 2

	// write wire field ForeGreen
	w.Write2b(ForeGreen)
	n += 2

	// write wire field ForeBlue
	w.Write2b(ForeBlue)
	n += 2

	// write wire field BackRed
	w.Write2b(BackRed)
	n += 2

	// write wire field BackGreen
	w.Write2b(BackGreen)
	n += 2

	// write wire field BackBlue
	w.Write2b(BackBlue)
	n += 2

	// write wire field X
	w.Write2b(X)
	n += 2

	// write wire field Y
	w.Write2b(Y)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CreateCursor(conn *Conn, Cid Cursor, Source Pixmap, Mask Pixmap, ForeRed uint16, ForeGreen uint16, ForeBlue uint16, BackRed uint16, BackGreen uint16, BackBlue uint16, X uint16, Y uint16) VoidCookie {
	w := NewWriter()
	CreateCursorRequestWrite(w, Cid, Source, Mask, ForeRed, ForeGreen, ForeBlue, BackRed, BackGreen, BackBlue, X, Y)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CreateCursorOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func CreateCursorChecked(conn *Conn, Cid Cursor, Source Pixmap, Mask Pixmap, ForeRed uint16, ForeGreen uint16, ForeBlue uint16, BackRed uint16, BackGreen uint16, BackBlue uint16, X uint16, Y uint16) VoidCookie {
	w := NewWriter()
	CreateCursorRequestWrite(w, Cid, Source, Mask, ForeRed, ForeGreen, ForeBlue, BackRed, BackGreen, BackBlue, X, Y)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CreateCursorOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum Font
const (
	FontNone = 0
)

const CreateGlyphCursorOpcode = 94

func CreateGlyphCursorRequestWrite(w *Writer, Cid Cursor, SourceFont Font, MaskFont Font, SourceChar uint16, MaskChar uint16, ForeRed uint16, ForeGreen uint16, ForeBlue uint16, BackRed uint16, BackGreen uint16, BackBlue uint16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cid
	w.Write4b(uint32(Cid))
	n += 4

	// write wire field SourceFont
	w.Write4b(uint32(SourceFont))
	n += 4

	// write wire field MaskFont
	w.Write4b(uint32(MaskFont))
	n += 4

	// write wire field SourceChar
	w.Write2b(SourceChar)
	n += 2

	// write wire field MaskChar
	w.Write2b(MaskChar)
	n += 2

	// write wire field ForeRed
	w.Write2b(ForeRed)
	n += 2

	// write wire field ForeGreen
	w.Write2b(ForeGreen)
	n += 2

	// write wire field ForeBlue
	w.Write2b(ForeBlue)
	n += 2

	// write wire field BackRed
	w.Write2b(BackRed)
	n += 2

	// write wire field BackGreen
	w.Write2b(BackGreen)
	n += 2

	// write wire field BackBlue
	w.Write2b(BackBlue)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CreateGlyphCursor(conn *Conn, Cid Cursor, SourceFont Font, MaskFont Font, SourceChar uint16, MaskChar uint16, ForeRed uint16, ForeGreen uint16, ForeBlue uint16, BackRed uint16, BackGreen uint16, BackBlue uint16) VoidCookie {
	w := NewWriter()
	CreateGlyphCursorRequestWrite(w, Cid, SourceFont, MaskFont, SourceChar, MaskChar, ForeRed, ForeGreen, ForeBlue, BackRed, BackGreen, BackBlue)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CreateGlyphCursorOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func CreateGlyphCursorChecked(conn *Conn, Cid Cursor, SourceFont Font, MaskFont Font, SourceChar uint16, MaskChar uint16, ForeRed uint16, ForeGreen uint16, ForeBlue uint16, BackRed uint16, BackGreen uint16, BackBlue uint16) VoidCookie {
	w := NewWriter()
	CreateGlyphCursorRequestWrite(w, Cid, SourceFont, MaskFont, SourceChar, MaskChar, ForeRed, ForeGreen, ForeBlue, BackRed, BackGreen, BackBlue)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: CreateGlyphCursorOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const FreeCursorOpcode = 95

func FreeCursorRequestWrite(w *Writer, Cursor Cursor) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cursor
	w.Write4b(uint32(Cursor))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func FreeCursor(conn *Conn, Cursor Cursor) VoidCookie {
	w := NewWriter()
	FreeCursorRequestWrite(w, Cursor)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: FreeCursorOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func FreeCursorChecked(conn *Conn, Cursor Cursor) VoidCookie {
	w := NewWriter()
	FreeCursorRequestWrite(w, Cursor)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: FreeCursorOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const RecolorCursorOpcode = 96

func RecolorCursorRequestWrite(w *Writer, Cursor Cursor, ForeRed uint16, ForeGreen uint16, ForeBlue uint16, BackRed uint16, BackGreen uint16, BackBlue uint16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cursor
	w.Write4b(uint32(Cursor))
	n += 4

	// write wire field ForeRed
	w.Write2b(ForeRed)
	n += 2

	// write wire field ForeGreen
	w.Write2b(ForeGreen)
	n += 2

	// write wire field ForeBlue
	w.Write2b(ForeBlue)
	n += 2

	// write wire field BackRed
	w.Write2b(BackRed)
	n += 2

	// write wire field BackGreen
	w.Write2b(BackGreen)
	n += 2

	// write wire field BackBlue
	w.Write2b(BackBlue)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func RecolorCursor(conn *Conn, Cursor Cursor, ForeRed uint16, ForeGreen uint16, ForeBlue uint16, BackRed uint16, BackGreen uint16, BackBlue uint16) VoidCookie {
	w := NewWriter()
	RecolorCursorRequestWrite(w, Cursor, ForeRed, ForeGreen, ForeBlue, BackRed, BackGreen, BackBlue)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: RecolorCursorOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func RecolorCursorChecked(conn *Conn, Cursor Cursor, ForeRed uint16, ForeGreen uint16, ForeBlue uint16, BackRed uint16, BackGreen uint16, BackBlue uint16) VoidCookie {
	w := NewWriter()
	RecolorCursorRequestWrite(w, Cursor, ForeRed, ForeGreen, ForeBlue, BackRed, BackGreen, BackBlue)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: RecolorCursorOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum QueryShapeOf
const (
	QueryShapeOfLargestCursor  = 0
	QueryShapeOfFastestTile    = 1
	QueryShapeOfFastestStipple = 2
)

const QueryBestSizeOpcode = 97

func QueryBestSizeRequestWrite(w *Writer, Class uint8, Drawable Drawable, Width uint16, Height uint16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Class
	w.Write1b(Class)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field Width
	w.Write2b(Width)
	n += 2

	// write wire field Height
	w.Write2b(Height)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type QueryBestSizeReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Length   uint32
	Width    uint16
	Height   uint16
}

func QueryBestSizeReplyRead(r *Reader, v *QueryBestSizeReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Width
	v.Width = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Height
	v.Height = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

type QueryBestSizeCookie uint64

func (cookie QueryBestSizeCookie) Reply(conn *Conn) (*QueryBestSizeReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply QueryBestSizeReply
	QueryBestSizeReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryBestSize(conn *Conn, Class uint8, Drawable Drawable, Width uint16, Height uint16) QueryBestSizeCookie {
	w := NewWriter()
	QueryBestSizeRequestWrite(w, Class, Drawable, Width, Height)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: QueryBestSizeOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return QueryBestSizeCookie(seq)
}

func QueryBestSizeUnchecked(conn *Conn, Class uint8, Drawable Drawable, Width uint16, Height uint16) QueryBestSizeCookie {
	w := NewWriter()
	QueryBestSizeRequestWrite(w, Class, Drawable, Width, Height)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: QueryBestSizeOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return QueryBestSizeCookie(seq)
}

const QueryExtensionOpcode = 98

func QueryExtensionRequestWrite(w *Writer, NameLen uint16, Name string) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field NameLen
	w.Write2b(NameLen)
	n += 2

	// write wire field Pad1
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Name
	w.WriteString(Name)
	n += len(Name)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type QueryExtensionReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence    uint16
	Length      uint32
	Present     uint8
	MajorOpcode uint8
	FirstEvent  uint8
	FirstError  uint8
}

func QueryExtensionReplyRead(r *Reader, v *QueryExtensionReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Present
	v.Present = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field MajorOpcode
	v.MajorOpcode = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field FirstEvent
	v.FirstEvent = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field FirstError
	v.FirstError = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}
	return
}

type QueryExtensionCookie uint64

func (cookie QueryExtensionCookie) Reply(conn *Conn) (*QueryExtensionReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply QueryExtensionReply
	QueryExtensionReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryExtension(conn *Conn, NameLen uint16, Name string) QueryExtensionCookie {
	w := NewWriter()
	QueryExtensionRequestWrite(w, NameLen, Name)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: QueryExtensionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return QueryExtensionCookie(seq)
}

func QueryExtensionUnchecked(conn *Conn, NameLen uint16, Name string) QueryExtensionCookie {
	w := NewWriter()
	QueryExtensionRequestWrite(w, NameLen, Name)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: QueryExtensionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return QueryExtensionCookie(seq)
}

const ListExtensionsOpcode = 99

func ListExtensionsRequestWrite(w *Writer) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type ListExtensionsReply struct {
	ResponseType uint8
	NamesLen     uint8
	Sequence     uint16
	Length       uint32
	// Pad0 [24]uint8
	Names []Str
}

func ListExtensionsReplyRead(r *Reader, v *ListExtensionsReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field NamesLen
	v.NamesLen = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(24)
	n += 24
	if r.Err() != nil {
		return
	}

	// read field Names
	v.Names = make([]Str, int(v.NamesLen))
	blockLen += StrReadN(r, v.Names)
	n += blockLen
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CStr{}))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type ListExtensionsCookie uint64

func (cookie ListExtensionsCookie) Reply(conn *Conn) (*ListExtensionsReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply ListExtensionsReply
	ListExtensionsReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func ListExtensions(conn *Conn) ListExtensionsCookie {
	w := NewWriter()
	ListExtensionsRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: ListExtensionsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return ListExtensionsCookie(seq)
}

func ListExtensionsUnchecked(conn *Conn) ListExtensionsCookie {
	w := NewWriter()
	ListExtensionsRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: ListExtensionsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return ListExtensionsCookie(seq)
}

const ChangeKeyboardMappingOpcode = 100

func ChangeKeyboardMappingRequestWrite(w *Writer, KeycodeCount uint8, FirstKeycode Keycode, KeysymsPerKeycode uint8, Keysyms []Keysym) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field KeycodeCount
	w.Write1b(KeycodeCount)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field FirstKeycode
	w.Write1b(uint8(FirstKeycode))
	n += 1

	// write wire field KeysymsPerKeycode
	w.Write1b(KeysymsPerKeycode)
	n += 1

	// write wire field Pad0
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Keysyms
	{
		_dataLen := (int(KeycodeCount) * int(KeysymsPerKeycode))
		for i := 0; i < _dataLen; i++ {
			w.Write4b(uint32(Keysyms[i]))
		}
		n += _dataLen * 4
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func ChangeKeyboardMapping(conn *Conn, KeycodeCount uint8, FirstKeycode Keycode, KeysymsPerKeycode uint8, Keysyms []Keysym) VoidCookie {
	w := NewWriter()
	ChangeKeyboardMappingRequestWrite(w, KeycodeCount, FirstKeycode, KeysymsPerKeycode, Keysyms)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ChangeKeyboardMappingOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func ChangeKeyboardMappingChecked(conn *Conn, KeycodeCount uint8, FirstKeycode Keycode, KeysymsPerKeycode uint8, Keysyms []Keysym) VoidCookie {
	w := NewWriter()
	ChangeKeyboardMappingRequestWrite(w, KeycodeCount, FirstKeycode, KeysymsPerKeycode, Keysyms)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ChangeKeyboardMappingOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const GetKeyboardMappingOpcode = 101

func GetKeyboardMappingRequestWrite(w *Writer, FirstKeycode Keycode, Count uint8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field FirstKeycode
	w.Write1b(uint8(FirstKeycode))
	n += 1

	// write wire field Count
	w.Write1b(Count)
	n += 1
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetKeyboardMappingReply struct {
	ResponseType      uint8
	KeysymsPerKeycode uint8
	Sequence          uint16
	Length            uint32
	// Pad0 [24]uint8
	Keysyms []Keysym
}

func GetKeyboardMappingReplyRead(r *Reader, v *GetKeyboardMappingReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field KeysymsPerKeycode
	v.KeysymsPerKeycode = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(24)
	n += 24
	if r.Err() != nil {
		return
	}

	// read field Keysyms
	{
		dataLen := int(int(v.Length))
		data := make([]Keysym, dataLen)
		for i := 0; i < dataLen; i++ {
			data[i] = Keysym(r.Read4b())
			if r.Err() != nil {
				return
			}
		}
		blockLen += dataLen * 4
		n += blockLen
		v.Keysyms = data
	}
	alignTo = int(unsafe.Alignof(Keysym(0)))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type GetKeyboardMappingCookie uint64

func (cookie GetKeyboardMappingCookie) Reply(conn *Conn) (*GetKeyboardMappingReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply GetKeyboardMappingReply
	GetKeyboardMappingReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetKeyboardMapping(conn *Conn, FirstKeycode Keycode, Count uint8) GetKeyboardMappingCookie {
	w := NewWriter()
	GetKeyboardMappingRequestWrite(w, FirstKeycode, Count)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetKeyboardMappingOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return GetKeyboardMappingCookie(seq)
}

func GetKeyboardMappingUnchecked(conn *Conn, FirstKeycode Keycode, Count uint8) GetKeyboardMappingCookie {
	w := NewWriter()
	GetKeyboardMappingRequestWrite(w, FirstKeycode, Count)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetKeyboardMappingOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetKeyboardMappingCookie(seq)
}

// enum KB
const (
	KBKeyClickPercent = 1
	KBBellPercent     = 2
	KBBellPitch       = 4
	KBBellDuration    = 8
	KBLed             = 16
	KBLedMode         = 32
	KBKey             = 64
	KBAutoRepeatMode  = 128
)

// enum LedMode
const (
	LedModeOff = 0
	LedModeOn  = 1
)

// enum AutoRepeatMode
const (
	AutoRepeatModeOff     = 0
	AutoRepeatModeOn      = 1
	AutoRepeatModeDefault = 2
)

type ChangeKeyboardControlValueList struct {
	KeyClickPercent int32
	BellPercent     int32
	BellPitch       int32
	BellDuration    int32
	Led             uint32
	LedMode         uint32
	Key             Keycode32
	AutoRepeatMode  uint32
}

// _go_switch_write SwitchType "xcb.ChangeKeyboardControl.value_list"
func ChangeKeyboardControlValueListSerialize(w *Writer, ValueMask uint32, _aux *ChangeKeyboardControlValueList) {
	n := 0
	xgbPaddingOffset := 0
	alignTo := 0
	xgbPad := 0
	// switch expr: int(ValueMask)

	// case_field Field with type BitcaseType "xcb.ChangeKeyboardControl.value_list.bitcase1"
	if (int(ValueMask) & KBKeyClickPercent) != 0 {
		// todo
		// field Field "key_click_percent" in BitcaseType "xcb.ChangeKeyboardControl.value_list.bitcase1"
		w.Write4b(uint32(_aux.KeyClickPercent))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeKeyboardControl.value_list.bitcase2"
	if (int(ValueMask) & KBBellPercent) != 0 {
		// todo
		// field Field "bell_percent" in BitcaseType "xcb.ChangeKeyboardControl.value_list.bitcase2"
		w.Write4b(uint32(_aux.BellPercent))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeKeyboardControl.value_list.bitcase3"
	if (int(ValueMask) & KBBellPitch) != 0 {
		// todo
		// field Field "bell_pitch" in BitcaseType "xcb.ChangeKeyboardControl.value_list.bitcase3"
		w.Write4b(uint32(_aux.BellPitch))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeKeyboardControl.value_list.bitcase4"
	if (int(ValueMask) & KBBellDuration) != 0 {
		// todo
		// field Field "bell_duration" in BitcaseType "xcb.ChangeKeyboardControl.value_list.bitcase4"
		w.Write4b(uint32(_aux.BellDuration))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeKeyboardControl.value_list.bitcase5"
	if (int(ValueMask) & KBLed) != 0 {
		// todo
		// field Field "led" in BitcaseType "xcb.ChangeKeyboardControl.value_list.bitcase5"
		w.Write4b(_aux.Led)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeKeyboardControl.value_list.bitcase6"
	if (int(ValueMask) & KBLedMode) != 0 {
		// todo
		// field Field "led_mode" in BitcaseType "xcb.ChangeKeyboardControl.value_list.bitcase6"
		w.Write4b(_aux.LedMode)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeKeyboardControl.value_list.bitcase7"
	if (int(ValueMask) & KBKey) != 0 {
		// todo
		// field Field "key" in BitcaseType "xcb.ChangeKeyboardControl.value_list.bitcase7"
		w.Write4b(uint32(_aux.Key))
		n += 4
		alignTo = int(unsafe.Alignof(Keycode32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ChangeKeyboardControl.value_list.bitcase8"
	if (int(ValueMask) & KBAutoRepeatMode) != 0 {
		// todo
		// field Field "auto_repeat_mode" in BitcaseType "xcb.ChangeKeyboardControl.value_list.bitcase8"
		w.Write4b(_aux.AutoRepeatMode)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	/* insert padding */
	xgbPad = -(n + xgbPaddingOffset) & (alignTo - 1)
	w.WritePad(xgbPad)
	xgbPad = 0
	n = 0
	xgbPaddingOffset = 0
}

const ChangeKeyboardControlOpcode = 102

func ChangeKeyboardControlRequestWrite(w *Writer, ValueMask uint32, ValueList *ChangeKeyboardControlValueList) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field ValueMask
	w.Write4b(ValueMask)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field ValueList
	// switch serialize
	ChangeKeyboardControlValueListSerialize(w, ValueMask, ValueList)
}

func ChangeKeyboardControl(conn *Conn, ValueMask uint32, ValueList *ChangeKeyboardControlValueList) VoidCookie {
	w := NewWriter()
	ChangeKeyboardControlRequestWrite(w, ValueMask, ValueList)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ChangeKeyboardControlOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func ChangeKeyboardControlChecked(conn *Conn, ValueMask uint32, ValueList *ChangeKeyboardControlValueList) VoidCookie {
	w := NewWriter()
	ChangeKeyboardControlRequestWrite(w, ValueMask, ValueList)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ChangeKeyboardControlOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const GetKeyboardControlOpcode = 103

func GetKeyboardControlRequestWrite(w *Writer) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetKeyboardControlReply struct {
	ResponseType     uint8
	GlobalAutoRepeat uint8
	Sequence         uint16
	Length           uint32
	LedMask          uint32
	KeyClickPercent  uint8
	BellPercent      uint8
	BellPitch        uint16
	BellDuration     uint16
	// Pad0 [2]uint8
	AutoRepeats [32]uint8
}

func GetKeyboardControlReplyRead(r *Reader, v *GetKeyboardControlReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field GlobalAutoRepeat
	v.GlobalAutoRepeat = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field LedMask
	v.LedMask = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field KeyClickPercent
	v.KeyClickPercent = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field BellPercent
	v.BellPercent = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field BellPitch
	v.BellPitch = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field BellDuration
	v.BellDuration = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(2)
	n += 2
	if r.Err() != nil {
		return
	}

	// read field AutoRepeats
	{
		dataLen := int(32)
		for i := 0; i < dataLen; i++ {
			v.AutoRepeats[i] = r.Read1b()
			if r.Err() != nil {
				return
			}
		}
		n += dataLen
	}
	return
}

type GetKeyboardControlCookie uint64

func (cookie GetKeyboardControlCookie) Reply(conn *Conn) (*GetKeyboardControlReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply GetKeyboardControlReply
	GetKeyboardControlReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetKeyboardControl(conn *Conn) GetKeyboardControlCookie {
	w := NewWriter()
	GetKeyboardControlRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetKeyboardControlOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return GetKeyboardControlCookie(seq)
}

func GetKeyboardControlUnchecked(conn *Conn) GetKeyboardControlCookie {
	w := NewWriter()
	GetKeyboardControlRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetKeyboardControlOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetKeyboardControlCookie(seq)
}

const BellOpcode = 104

func BellRequestWrite(w *Writer, Percent int8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Percent
	w.Write1b(uint8(Percent))
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func Bell(conn *Conn, Percent int8) VoidCookie {
	w := NewWriter()
	BellRequestWrite(w, Percent)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: BellOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func BellChecked(conn *Conn, Percent int8) VoidCookie {
	w := NewWriter()
	BellRequestWrite(w, Percent)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: BellOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const ChangePointerControlOpcode = 105

func ChangePointerControlRequestWrite(w *Writer, AccelerationNumerator int16, AccelerationDenominator int16, Threshold int16, DoAcceleration uint8, DoThreshold uint8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field AccelerationNumerator
	w.Write2b(uint16(AccelerationNumerator))
	n += 2

	// write wire field AccelerationDenominator
	w.Write2b(uint16(AccelerationDenominator))
	n += 2

	// write wire field Threshold
	w.Write2b(uint16(Threshold))
	n += 2

	// write wire field DoAcceleration
	w.Write1b(DoAcceleration)
	n += 1

	// write wire field DoThreshold
	w.Write1b(DoThreshold)
	n += 1
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func ChangePointerControl(conn *Conn, AccelerationNumerator int16, AccelerationDenominator int16, Threshold int16, DoAcceleration uint8, DoThreshold uint8) VoidCookie {
	w := NewWriter()
	ChangePointerControlRequestWrite(w, AccelerationNumerator, AccelerationDenominator, Threshold, DoAcceleration, DoThreshold)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ChangePointerControlOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func ChangePointerControlChecked(conn *Conn, AccelerationNumerator int16, AccelerationDenominator int16, Threshold int16, DoAcceleration uint8, DoThreshold uint8) VoidCookie {
	w := NewWriter()
	ChangePointerControlRequestWrite(w, AccelerationNumerator, AccelerationDenominator, Threshold, DoAcceleration, DoThreshold)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ChangePointerControlOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const GetPointerControlOpcode = 106

func GetPointerControlRequestWrite(w *Writer) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetPointerControlReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence                uint16
	Length                  uint32
	AccelerationNumerator   uint16
	AccelerationDenominator uint16
	Threshold               uint16
	// Pad1 [18]uint8
}

func GetPointerControlReplyRead(r *Reader, v *GetPointerControlReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field AccelerationNumerator
	v.AccelerationNumerator = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field AccelerationDenominator
	v.AccelerationDenominator = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Threshold
	v.Threshold = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(18)
	n += 18
	if r.Err() != nil {
		return
	}
	return
}

type GetPointerControlCookie uint64

func (cookie GetPointerControlCookie) Reply(conn *Conn) (*GetPointerControlReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply GetPointerControlReply
	GetPointerControlReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetPointerControl(conn *Conn) GetPointerControlCookie {
	w := NewWriter()
	GetPointerControlRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetPointerControlOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return GetPointerControlCookie(seq)
}

func GetPointerControlUnchecked(conn *Conn) GetPointerControlCookie {
	w := NewWriter()
	GetPointerControlRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetPointerControlOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetPointerControlCookie(seq)
}

// enum Blanking
const (
	BlankingNotPreferred = 0
	BlankingPreferred    = 1
	BlankingDefault      = 2
)

func BlankingEnumToStr(v int) string {
	switch v {
	case 0:
		return "NotPreferred"
	case 1:
		return "Preferred"
	case 2:
		return "Default"
	default:
		return "<Unknown>"
	}
}

// enum Exposures
const (
	ExposuresNotAllowed = 0
	ExposuresAllowed    = 1
	ExposuresDefault    = 2
)

func ExposuresEnumToStr(v int) string {
	switch v {
	case 0:
		return "NotAllowed"
	case 1:
		return "Allowed"
	case 2:
		return "Default"
	default:
		return "<Unknown>"
	}
}

const SetScreenSaverOpcode = 107

func SetScreenSaverRequestWrite(w *Writer, Timeout int16, Interval int16, PreferBlanking uint8, AllowExposures uint8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Timeout
	w.Write2b(uint16(Timeout))
	n += 2

	// write wire field Interval
	w.Write2b(uint16(Interval))
	n += 2

	// write wire field PreferBlanking
	w.Write1b(PreferBlanking)
	n += 1

	// write wire field AllowExposures
	w.Write1b(AllowExposures)
	n += 1
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SetScreenSaver(conn *Conn, Timeout int16, Interval int16, PreferBlanking uint8, AllowExposures uint8) VoidCookie {
	w := NewWriter()
	SetScreenSaverRequestWrite(w, Timeout, Interval, PreferBlanking, AllowExposures)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: SetScreenSaverOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func SetScreenSaverChecked(conn *Conn, Timeout int16, Interval int16, PreferBlanking uint8, AllowExposures uint8) VoidCookie {
	w := NewWriter()
	SetScreenSaverRequestWrite(w, Timeout, Interval, PreferBlanking, AllowExposures)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: SetScreenSaverOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const GetScreenSaverOpcode = 108

func GetScreenSaverRequestWrite(w *Writer) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetScreenSaverReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence       uint16
	Length         uint32
	Timeout        uint16
	Interval       uint16
	PreferBlanking uint8
	AllowExposures uint8
	// Pad1 [18]uint8
}

func GetScreenSaverReplyRead(r *Reader, v *GetScreenSaverReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Timeout
	v.Timeout = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Interval
	v.Interval = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field PreferBlanking
	v.PreferBlanking = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field AllowExposures
	v.AllowExposures = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(18)
	n += 18
	if r.Err() != nil {
		return
	}
	return
}

type GetScreenSaverCookie uint64

func (cookie GetScreenSaverCookie) Reply(conn *Conn) (*GetScreenSaverReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply GetScreenSaverReply
	GetScreenSaverReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetScreenSaver(conn *Conn) GetScreenSaverCookie {
	w := NewWriter()
	GetScreenSaverRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetScreenSaverOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return GetScreenSaverCookie(seq)
}

func GetScreenSaverUnchecked(conn *Conn) GetScreenSaverCookie {
	w := NewWriter()
	GetScreenSaverRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetScreenSaverOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetScreenSaverCookie(seq)
}

// enum HostMode
const (
	HostModeInsert = 0
	HostModeDelete = 1
)

// enum Family
const (
	FamilyInternet          = 0
	FamilyDECnet            = 1
	FamilyChaos             = 2
	FamilyServerInterpreted = 5
	FamilyInternet6         = 6
)

const ChangeHostsOpcode = 109

func ChangeHostsRequestWrite(w *Writer, Mode uint8, Family uint8, AddressLen uint16, Address []uint8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Mode
	w.Write1b(Mode)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Family
	w.Write1b(Family)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field AddressLen
	w.Write2b(AddressLen)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Address
	{
		_dataLen := int(AddressLen)
		for i := 0; i < _dataLen; i++ {
			w.Write1b(uint8(Address[i]))
		}
		n += _dataLen * 1
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func ChangeHosts(conn *Conn, Mode uint8, Family uint8, AddressLen uint16, Address []uint8) VoidCookie {
	w := NewWriter()
	ChangeHostsRequestWrite(w, Mode, Family, AddressLen, Address)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ChangeHostsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func ChangeHostsChecked(conn *Conn, Mode uint8, Family uint8, AddressLen uint16, Address []uint8) VoidCookie {
	w := NewWriter()
	ChangeHostsRequestWrite(w, Mode, Family, AddressLen, Address)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ChangeHostsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

type Host struct {
	Family uint8
	// Pad0 uint8
	AddressLen uint16
	Address    []uint8
}

type CHost struct {
	Family     uint8
	Pad0       uint8
	AddressLen uint16
}

func HostRead(r *Reader, v *Host) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Family
	v.Family = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field AddressLen
	v.AddressLen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Address
	{
		dataLen := int(int(v.AddressLen))
		data := make([]uint8, dataLen)
		for i := 0; i < dataLen; i++ {
			data[i] = uint8(r.Read1b())
			if r.Err() != nil {
				return
			}
		}
		blockLen += dataLen
		n += blockLen
		v.Address = data
	}
	alignTo = int(unsafe.Alignof(uint8(0)))

	// read field Pad1
	alignTo = 4

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

func HostReadN(r *Reader, dest []Host) (n int) {
	for i := 0; i < len(dest); i++ {
		n += HostRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func HostWrite(w *Writer, v *Host) (n int) {

	// write field Family
	w.Write1b(v.Family)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field AddressLen
	w.Write2b(v.AddressLen)
	n += 2

	// write field Address

	// write field Pad1
	return
}
func HostWriteN(w *Writer, src []Host) (n int) {
	for i := 0; i < len(src); i++ {
		n += HostWrite(w, &src[i])
	}
	return
}

const ListHostsOpcode = 110

func ListHostsRequestWrite(w *Writer) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type ListHostsReply struct {
	ResponseType uint8
	Mode         uint8
	Sequence     uint16
	Length       uint32
	HostsLen     uint16
	// Pad0 [22]uint8
	Hosts []Host
}

func ListHostsReplyRead(r *Reader, v *ListHostsReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Mode
	v.Mode = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field HostsLen
	v.HostsLen = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(22)
	n += 22
	if r.Err() != nil {
		return
	}

	// read field Hosts
	v.Hosts = make([]Host, int(v.HostsLen))
	blockLen += HostReadN(r, v.Hosts)
	n += blockLen
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CHost{}))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type ListHostsCookie uint64

func (cookie ListHostsCookie) Reply(conn *Conn) (*ListHostsReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply ListHostsReply
	ListHostsReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func ListHosts(conn *Conn) ListHostsCookie {
	w := NewWriter()
	ListHostsRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: ListHostsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return ListHostsCookie(seq)
}

func ListHostsUnchecked(conn *Conn) ListHostsCookie {
	w := NewWriter()
	ListHostsRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: ListHostsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return ListHostsCookie(seq)
}

// enum AccessControl
const (
	AccessControlDisable = 0
	AccessControlEnable  = 1
)

const SetAccessControlOpcode = 111

func SetAccessControlRequestWrite(w *Writer, Mode uint8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Mode
	w.Write1b(Mode)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SetAccessControl(conn *Conn, Mode uint8) VoidCookie {
	w := NewWriter()
	SetAccessControlRequestWrite(w, Mode)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: SetAccessControlOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func SetAccessControlChecked(conn *Conn, Mode uint8) VoidCookie {
	w := NewWriter()
	SetAccessControlRequestWrite(w, Mode)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: SetAccessControlOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum CloseDown
const (
	CloseDownDestroyAll      = 0
	CloseDownRetainPermanent = 1
	CloseDownRetainTemporary = 2
)

const SetCloseDownModeOpcode = 112

func SetCloseDownModeRequestWrite(w *Writer, Mode uint8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Mode
	w.Write1b(Mode)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SetCloseDownMode(conn *Conn, Mode uint8) VoidCookie {
	w := NewWriter()
	SetCloseDownModeRequestWrite(w, Mode)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: SetCloseDownModeOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func SetCloseDownModeChecked(conn *Conn, Mode uint8) VoidCookie {
	w := NewWriter()
	SetCloseDownModeRequestWrite(w, Mode)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: SetCloseDownModeOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum Kill
const (
	KillAllTemporary = 0
)

const KillClientOpcode = 113

func KillClientRequestWrite(w *Writer, Resource uint32) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Resource
	w.Write4b(Resource)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func KillClient(conn *Conn, Resource uint32) VoidCookie {
	w := NewWriter()
	KillClientRequestWrite(w, Resource)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: KillClientOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func KillClientChecked(conn *Conn, Resource uint32) VoidCookie {
	w := NewWriter()
	KillClientRequestWrite(w, Resource)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: KillClientOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

const RotatePropertiesOpcode = 114

func RotatePropertiesRequestWrite(w *Writer, Window Window, AtomsLen uint16, Delta int16, Atoms []Atom) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4

	// write wire field AtomsLen
	w.Write2b(AtomsLen)
	n += 2

	// write wire field Delta
	w.Write2b(uint16(Delta))
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Atoms
	{
		_dataLen := int(AtomsLen)
		for i := 0; i < _dataLen; i++ {
			w.Write4b(uint32(Atoms[i]))
		}
		n += _dataLen * 4
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func RotateProperties(conn *Conn, Window Window, AtomsLen uint16, Delta int16, Atoms []Atom) VoidCookie {
	w := NewWriter()
	RotatePropertiesRequestWrite(w, Window, AtomsLen, Delta, Atoms)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: RotatePropertiesOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func RotatePropertiesChecked(conn *Conn, Window Window, AtomsLen uint16, Delta int16, Atoms []Atom) VoidCookie {
	w := NewWriter()
	RotatePropertiesRequestWrite(w, Window, AtomsLen, Delta, Atoms)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: RotatePropertiesOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum ScreenSaver
const (
	ScreenSaverReset  = 0
	ScreenSaverActive = 1
)

const ForceScreenSaverOpcode = 115

func ForceScreenSaverRequestWrite(w *Writer, Mode uint8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Mode
	w.Write1b(Mode)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func ForceScreenSaver(conn *Conn, Mode uint8) VoidCookie {
	w := NewWriter()
	ForceScreenSaverRequestWrite(w, Mode)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ForceScreenSaverOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func ForceScreenSaverChecked(conn *Conn, Mode uint8) VoidCookie {
	w := NewWriter()
	ForceScreenSaverRequestWrite(w, Mode)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: ForceScreenSaverOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}

// enum MappingStatus
const (
	MappingStatusSuccess = 0
	MappingStatusBusy    = 1
	MappingStatusFailure = 2
)

func MappingStatusEnumToStr(v int) string {
	switch v {
	case 0:
		return "Success"
	case 1:
		return "Busy"
	case 2:
		return "Failure"
	default:
		return "<Unknown>"
	}
}

const SetPointerMappingOpcode = 116

func SetPointerMappingRequestWrite(w *Writer, MapLen uint8, Map []uint8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MapLen
	w.Write1b(MapLen)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Map
	{
		_dataLen := int(MapLen)
		for i := 0; i < _dataLen; i++ {
			w.Write1b(uint8(Map[i]))
		}
		n += _dataLen * 1
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type SetPointerMappingReply struct {
	ResponseType uint8
	Status       uint8
	Sequence     uint16
	Length       uint32
}

func SetPointerMappingReplyRead(r *Reader, v *SetPointerMappingReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Status
	v.Status = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}
	return
}

type SetPointerMappingCookie uint64

func (cookie SetPointerMappingCookie) Reply(conn *Conn) (*SetPointerMappingReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply SetPointerMappingReply
	SetPointerMappingReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetPointerMapping(conn *Conn, MapLen uint8, Map []uint8) SetPointerMappingCookie {
	w := NewWriter()
	SetPointerMappingRequestWrite(w, MapLen, Map)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: SetPointerMappingOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return SetPointerMappingCookie(seq)
}

func SetPointerMappingUnchecked(conn *Conn, MapLen uint8, Map []uint8) SetPointerMappingCookie {
	w := NewWriter()
	SetPointerMappingRequestWrite(w, MapLen, Map)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: SetPointerMappingOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return SetPointerMappingCookie(seq)
}

const GetPointerMappingOpcode = 117

func GetPointerMappingRequestWrite(w *Writer) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetPointerMappingReply struct {
	ResponseType uint8
	MapLen       uint8
	Sequence     uint16
	Length       uint32
	// Pad0 [24]uint8
	Map []uint8
}

func GetPointerMappingReplyRead(r *Reader, v *GetPointerMappingReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field MapLen
	v.MapLen = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(24)
	n += 24
	if r.Err() != nil {
		return
	}

	// read field Map
	{
		dataLen := int(int(v.MapLen))
		data := make([]uint8, dataLen)
		for i := 0; i < dataLen; i++ {
			data[i] = uint8(r.Read1b())
			if r.Err() != nil {
				return
			}
		}
		blockLen += dataLen
		n += blockLen
		v.Map = data
	}
	alignTo = int(unsafe.Alignof(uint8(0)))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type GetPointerMappingCookie uint64

func (cookie GetPointerMappingCookie) Reply(conn *Conn) (*GetPointerMappingReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply GetPointerMappingReply
	GetPointerMappingReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetPointerMapping(conn *Conn) GetPointerMappingCookie {
	w := NewWriter()
	GetPointerMappingRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetPointerMappingOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return GetPointerMappingCookie(seq)
}

func GetPointerMappingUnchecked(conn *Conn) GetPointerMappingCookie {
	w := NewWriter()
	GetPointerMappingRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetPointerMappingOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetPointerMappingCookie(seq)
}

// enum MapIndex
const (
	MapIndexShift   = 0
	MapIndexLock    = 1
	MapIndexControl = 2
	MapIndex1       = 3
	MapIndex2       = 4
	MapIndex3       = 5
	MapIndex4       = 6
	MapIndex5       = 7
)

const SetModifierMappingOpcode = 118

func SetModifierMappingRequestWrite(w *Writer, KeycodesPerModifier uint8, Keycodes []Keycode) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field KeycodesPerModifier
	w.Write1b(KeycodesPerModifier)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Keycodes
	{
		_dataLen := (int(KeycodesPerModifier) * 8)
		for i := 0; i < _dataLen; i++ {
			w.Write1b(uint8(Keycodes[i]))
		}
		n += _dataLen * 1
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type SetModifierMappingReply struct {
	ResponseType uint8
	Status       uint8
	Sequence     uint16
	Length       uint32
}

func SetModifierMappingReplyRead(r *Reader, v *SetModifierMappingReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Status
	v.Status = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}
	return
}

type SetModifierMappingCookie uint64

func (cookie SetModifierMappingCookie) Reply(conn *Conn) (*SetModifierMappingReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply SetModifierMappingReply
	SetModifierMappingReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetModifierMapping(conn *Conn, KeycodesPerModifier uint8, Keycodes []Keycode) SetModifierMappingCookie {
	w := NewWriter()
	SetModifierMappingRequestWrite(w, KeycodesPerModifier, Keycodes)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: SetModifierMappingOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return SetModifierMappingCookie(seq)
}

func SetModifierMappingUnchecked(conn *Conn, KeycodesPerModifier uint8, Keycodes []Keycode) SetModifierMappingCookie {
	w := NewWriter()
	SetModifierMappingRequestWrite(w, KeycodesPerModifier, Keycodes)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: SetModifierMappingOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return SetModifierMappingCookie(seq)
}

const GetModifierMappingOpcode = 119

func GetModifierMappingRequestWrite(w *Writer) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetModifierMappingReply struct {
	ResponseType        uint8
	KeycodesPerModifier uint8
	Sequence            uint16
	Length              uint32
	// Pad0 [24]uint8
	Keycodes []Keycode
}

func GetModifierMappingReplyRead(r *Reader, v *GetModifierMappingReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field KeycodesPerModifier
	v.KeycodesPerModifier = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(24)
	n += 24
	if r.Err() != nil {
		return
	}

	// read field Keycodes
	{
		dataLen := int((int(v.KeycodesPerModifier) * 8))
		data := make([]Keycode, dataLen)
		for i := 0; i < dataLen; i++ {
			data[i] = Keycode(r.Read1b())
			if r.Err() != nil {
				return
			}
		}
		blockLen += dataLen
		n += blockLen
		v.Keycodes = data
	}
	alignTo = int(unsafe.Alignof(Keycode(0)))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type GetModifierMappingCookie uint64

func (cookie GetModifierMappingCookie) Reply(conn *Conn) (*GetModifierMappingReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply GetModifierMappingReply
	GetModifierMappingReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetModifierMapping(conn *Conn) GetModifierMappingCookie {
	w := NewWriter()
	GetModifierMappingRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetModifierMappingOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return GetModifierMappingCookie(seq)
}

func GetModifierMappingUnchecked(conn *Conn) GetModifierMappingCookie {
	w := NewWriter()
	GetModifierMappingRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: false,
		Opcode: GetModifierMappingOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetModifierMappingCookie(seq)
}

const NoOperationOpcode = 127

func NoOperationRequestWrite(w *Writer) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func NoOperation(conn *Conn) VoidCookie {
	w := NewWriter()
	NoOperationRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: NoOperationOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return VoidCookie(seq)
}

func NoOperationChecked(conn *Conn) VoidCookie {
	w := NewWriter()
	NoOperationRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		IsVoid: true,
		Opcode: NoOperationOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return VoidCookie(seq)
}
