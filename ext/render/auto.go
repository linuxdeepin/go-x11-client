package render

import x "github.com/linuxdeepin/go-x11-client"

// module.direct_imports: [('xproto', 'xproto')]
import "unsafe"
import "log"
import "sync"
import "fmt"
import "strings"

var _ = x.Reader{}
var _ = unsafe.Alignof(0)
var _ sync.Mutex
var _ = fmt.Println
var _ = strings.Join

func init() {
	log.SetFlags(log.Lshortfile)
}

// _ns.ext_name: Render
const MajorVersion = 0
const MinorVersion = 11

var _ext = x.NewExtension("RENDER")

func Ext() *x.Extension {
	return _ext
}

// enum PictType
const (
	PictTypeIndexed = 0
	PictTypeDirect  = 1
)

// enum Picture
const (
	PictureNone = 0
)

// enum PictOp
const (
	PictOpClear               = 0
	PictOpSrc                 = 1
	PictOpDst                 = 2
	PictOpOver                = 3
	PictOpOverReverse         = 4
	PictOpIn                  = 5
	PictOpInReverse           = 6
	PictOpOut                 = 7
	PictOpOutReverse          = 8
	PictOpAtop                = 9
	PictOpAtopReverse         = 10
	PictOpXor                 = 11
	PictOpAdd                 = 12
	PictOpSaturate            = 13
	PictOpDisjointClear       = 16
	PictOpDisjointSrc         = 17
	PictOpDisjointDst         = 18
	PictOpDisjointOver        = 19
	PictOpDisjointOverReverse = 20
	PictOpDisjointIn          = 21
	PictOpDisjointInReverse   = 22
	PictOpDisjointOut         = 23
	PictOpDisjointOutReverse  = 24
	PictOpDisjointAtop        = 25
	PictOpDisjointAtopReverse = 26
	PictOpDisjointXor         = 27
	PictOpConjointClear       = 32
	PictOpConjointSrc         = 33
	PictOpConjointDst         = 34
	PictOpConjointOver        = 35
	PictOpConjointOverReverse = 36
	PictOpConjointIn          = 37
	PictOpConjointInReverse   = 38
	PictOpConjointOut         = 39
	PictOpConjointOutReverse  = 40
	PictOpConjointAtop        = 41
	PictOpConjointAtopReverse = 42
	PictOpConjointXor         = 43
	PictOpMultiply            = 48
	PictOpScreen              = 49
	PictOpOverlay             = 50
	PictOpDarken              = 51
	PictOpLighten             = 52
	PictOpColorDodge          = 53
	PictOpColorBurn           = 54
	PictOpHardLight           = 55
	PictOpSoftLight           = 56
	PictOpDifference          = 57
	PictOpExclusion           = 58
	PictOpHSLHue              = 59
	PictOpHSLSaturation       = 60
	PictOpHSLColor            = 61
	PictOpHSLLuminosity       = 62
)

// enum PolyEdge
const (
	PolyEdgeSharp  = 0
	PolyEdgeSmooth = 1
)

// enum PolyMode
const (
	PolyModePrecise   = 0
	PolyModeImprecise = 1
)

// enum CP
const (
	CPRepeat           = 1
	CPAlphaMap         = 2
	CPAlphaXOrigin     = 4
	CPAlphaYOrigin     = 8
	CPClipXOrigin      = 16
	CPClipYOrigin      = 32
	CPClipMask         = 64
	CPGraphicsExposure = 128
	CPSubwindowMode    = 256
	CPPolyEdge         = 512
	CPPolyMode         = 1024
	CPDither           = 2048
	CPComponentAlpha   = 4096
)

// enum SubPixel
const (
	SubPixelUnknown       = 0
	SubPixelHorizontalRGB = 1
	SubPixelHorizontalBGR = 2
	SubPixelVerticalRGB   = 3
	SubPixelVerticalBGR   = 4
	SubPixelNone          = 5
)

// enum Repeat
const (
	RepeatNone    = 0
	RepeatNormal  = 1
	RepeatPad     = 2
	RepeatReflect = 3
)

// simple ('xcb', 'Render', 'GLYPH')
type Glyph uint32

// simple ('xcb', 'Render', 'GLYPHSET')
type GlyphSet uint32

// simple ('xcb', 'Render', 'PICTURE')
type Picture uint32

// simple ('xcb', 'Render', 'PICTFORMAT')
type PictFormat uint32

// simple ('xcb', 'Render', 'FIXED')
type Fixed int32
type DirectFormat struct {
	RedShift   uint16
	RedMask    uint16
	GreenShift uint16
	GreenMask  uint16
	BlueShift  uint16
	BlueMask   uint16
	AlphaShift uint16
	AlphaMask  uint16
}

type CDirectFormat struct {
	RedShift   uint16
	RedMask    uint16
	GreenShift uint16
	GreenMask  uint16
	BlueShift  uint16
	BlueMask   uint16
	AlphaShift uint16
	AlphaMask  uint16
}

func DirectFormatRead(r *x.Reader, v *DirectFormat) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field RedShift
	v.RedShift = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field RedMask
	v.RedMask = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field GreenShift
	v.GreenShift = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field GreenMask
	v.GreenMask = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field BlueShift
	v.BlueShift = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field BlueMask
	v.BlueMask = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field AlphaShift
	v.AlphaShift = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field AlphaMask
	v.AlphaMask = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

func DirectFormatReadN(r *x.Reader, dest []DirectFormat) (n int) {
	for i := 0; i < len(dest); i++ {
		n += DirectFormatRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func DirectFormatWrite(w *x.Writer, v *DirectFormat) (n int) {

	// write field RedShift
	w.Write2b(v.RedShift)
	n += 2

	// write field RedMask
	w.Write2b(v.RedMask)
	n += 2

	// write field GreenShift
	w.Write2b(v.GreenShift)
	n += 2

	// write field GreenMask
	w.Write2b(v.GreenMask)
	n += 2

	// write field BlueShift
	w.Write2b(v.BlueShift)
	n += 2

	// write field BlueMask
	w.Write2b(v.BlueMask)
	n += 2

	// write field AlphaShift
	w.Write2b(v.AlphaShift)
	n += 2

	// write field AlphaMask
	w.Write2b(v.AlphaMask)
	n += 2
	return
}
func DirectFormatWriteN(w *x.Writer, src []DirectFormat) (n int) {
	for i := 0; i < len(src); i++ {
		n += DirectFormatWrite(w, &src[i])
	}
	return
}

type PictFormInfo struct {
	Id    PictFormat
	Type  uint8
	Depth uint8
	// Pad0 [2]uint8
	Direct   DirectFormat
	Colormap x.Colormap
}

type CPictFormInfo struct {
	Id       PictFormat
	Type     uint8
	Depth    uint8
	Pad0     [2]uint8
	Direct   DirectFormat
	Colormap x.Colormap
}

func PictFormInfoRead(r *x.Reader, v *PictFormInfo) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Id
	v.Id = PictFormat(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Type
	v.Type = r.Read1b()
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

	// read field Pad0
	r.ReadPad(2)
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Direct
	n += DirectFormatRead(r, &v.Direct)
	if r.Err() != nil {
		return
	}

	// read field Colormap
	v.Colormap = x.Colormap(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}
	return
}

func PictFormInfoReadN(r *x.Reader, dest []PictFormInfo) (n int) {
	for i := 0; i < len(dest); i++ {
		n += PictFormInfoRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func PictFormInfoWrite(w *x.Writer, v *PictFormInfo) (n int) {

	// write field Id
	w.Write4b(uint32(v.Id))
	n += 4

	// write field Type
	w.Write1b(v.Type)
	n += 1

	// write field Depth
	w.Write1b(v.Depth)
	n += 1

	// write field Pad0
	w.WritePad(2)
	n += 2

	// write field Direct
	n += DirectFormatWrite(w, &v.Direct)

	// write field Colormap
	w.Write4b(uint32(v.Colormap))
	n += 4
	return
}
func PictFormInfoWriteN(w *x.Writer, src []PictFormInfo) (n int) {
	for i := 0; i < len(src); i++ {
		n += PictFormInfoWrite(w, &src[i])
	}
	return
}

type PictVisual struct {
	Visual x.VisualID
	Format PictFormat
}

type CPictVisual struct {
	Visual x.VisualID
	Format PictFormat
}

func PictVisualRead(r *x.Reader, v *PictVisual) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Visual
	v.Visual = x.VisualID(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Format
	v.Format = PictFormat(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}
	return
}

func PictVisualReadN(r *x.Reader, dest []PictVisual) (n int) {
	for i := 0; i < len(dest); i++ {
		n += PictVisualRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func PictVisualWrite(w *x.Writer, v *PictVisual) (n int) {

	// write field Visual
	w.Write4b(uint32(v.Visual))
	n += 4

	// write field Format
	w.Write4b(uint32(v.Format))
	n += 4
	return
}
func PictVisualWriteN(w *x.Writer, src []PictVisual) (n int) {
	for i := 0; i < len(src); i++ {
		n += PictVisualWrite(w, &src[i])
	}
	return
}

type PictDepth struct {
	Depth uint8
	// Pad0 uint8
	NumVisuals uint16
	// Pad1 [4]uint8
	Visuals []PictVisual
}

type CPictDepth struct {
	Depth      uint8
	Pad0       uint8
	NumVisuals uint16
	Pad1       [4]uint8
}

func PictDepthRead(r *x.Reader, v *PictDepth) (n int) {
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

	// read field NumVisuals
	v.NumVisuals = r.Read2b()
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
	v.Visuals = make([]PictVisual, int(v.NumVisuals))
	blockLen += PictVisualReadN(r, v.Visuals)
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CPictVisual{}))

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

func PictDepthReadN(r *x.Reader, dest []PictDepth) (n int) {
	for i := 0; i < len(dest); i++ {
		n += PictDepthRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func PictDepthWrite(w *x.Writer, v *PictDepth) (n int) {

	// write field Depth
	w.Write1b(v.Depth)
	n += 1

	// write field Pad0
	w.WritePad(1)
	n += 1

	// write field NumVisuals
	w.Write2b(v.NumVisuals)
	n += 2

	// write field Pad1
	w.WritePad(4)
	n += 4

	// write field Visuals
	return
}
func PictDepthWriteN(w *x.Writer, src []PictDepth) (n int) {
	for i := 0; i < len(src); i++ {
		n += PictDepthWrite(w, &src[i])
	}
	return
}

type PictScreen struct {
	NumDepths uint32
	Fallback  PictFormat
	Depths    []PictDepth
}

type CPictScreen struct {
	NumDepths uint32
	Fallback  PictFormat
}

func PictScreenRead(r *x.Reader, v *PictScreen) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field NumDepths
	v.NumDepths = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Fallback
	v.Fallback = PictFormat(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Depths
	v.Depths = make([]PictDepth, int(v.NumDepths))
	blockLen += PictDepthReadN(r, v.Depths)
	n += blockLen
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CPictDepth{}))

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

func PictScreenReadN(r *x.Reader, dest []PictScreen) (n int) {
	for i := 0; i < len(dest); i++ {
		n += PictScreenRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func PictScreenWrite(w *x.Writer, v *PictScreen) (n int) {

	// write field NumDepths
	w.Write4b(v.NumDepths)
	n += 4

	// write field Fallback
	w.Write4b(uint32(v.Fallback))
	n += 4

	// write field Depths
	return
}
func PictScreenWriteN(w *x.Writer, src []PictScreen) (n int) {
	for i := 0; i < len(src); i++ {
		n += PictScreenWrite(w, &src[i])
	}
	return
}

type IndexValue struct {
	Pixel uint32
	Red   uint16
	Green uint16
	Blue  uint16
	Alpha uint16
}

type CIndexValue struct {
	Pixel uint32
	Red   uint16
	Green uint16
	Blue  uint16
	Alpha uint16
}

func IndexValueRead(r *x.Reader, v *IndexValue) (n int) {
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

	// read field Alpha
	v.Alpha = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

func IndexValueReadN(r *x.Reader, dest []IndexValue) (n int) {
	for i := 0; i < len(dest); i++ {
		n += IndexValueRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func IndexValueWrite(w *x.Writer, v *IndexValue) (n int) {

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

	// write field Alpha
	w.Write2b(v.Alpha)
	n += 2
	return
}
func IndexValueWriteN(w *x.Writer, src []IndexValue) (n int) {
	for i := 0; i < len(src); i++ {
		n += IndexValueWrite(w, &src[i])
	}
	return
}

type Color struct {
	Red   uint16
	Green uint16
	Blue  uint16
	Alpha uint16
}

type CColor struct {
	Red   uint16
	Green uint16
	Blue  uint16
	Alpha uint16
}

func ColorRead(r *x.Reader, v *Color) (n int) {
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

	// read field Alpha
	v.Alpha = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

func ColorReadN(r *x.Reader, dest []Color) (n int) {
	for i := 0; i < len(dest); i++ {
		n += ColorRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func ColorWrite(w *x.Writer, v *Color) (n int) {

	// write field Red
	w.Write2b(v.Red)
	n += 2

	// write field Green
	w.Write2b(v.Green)
	n += 2

	// write field Blue
	w.Write2b(v.Blue)
	n += 2

	// write field Alpha
	w.Write2b(v.Alpha)
	n += 2
	return
}
func ColorWriteN(w *x.Writer, src []Color) (n int) {
	for i := 0; i < len(src); i++ {
		n += ColorWrite(w, &src[i])
	}
	return
}

type PointFix struct {
	X Fixed
	Y Fixed
}

type CPointFix struct {
	X Fixed
	Y Fixed
}

func PointFixRead(r *x.Reader, v *PointFix) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field X
	v.X = Fixed(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Y
	v.Y = Fixed(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}
	return
}

func PointFixReadN(r *x.Reader, dest []PointFix) (n int) {
	for i := 0; i < len(dest); i++ {
		n += PointFixRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func PointFixWrite(w *x.Writer, v *PointFix) (n int) {

	// write field X
	w.Write4b(uint32(v.X))
	n += 4

	// write field Y
	w.Write4b(uint32(v.Y))
	n += 4
	return
}
func PointFixWriteN(w *x.Writer, src []PointFix) (n int) {
	for i := 0; i < len(src); i++ {
		n += PointFixWrite(w, &src[i])
	}
	return
}

type LineFix struct {
	P1 PointFix
	P2 PointFix
}

type CLineFix struct {
	P1 PointFix
	P2 PointFix
}

func LineFixRead(r *x.Reader, v *LineFix) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field P1
	n += PointFixRead(r, &v.P1)
	if r.Err() != nil {
		return
	}

	// read field P2
	n += PointFixRead(r, &v.P2)
	if r.Err() != nil {
		return
	}
	return
}

func LineFixReadN(r *x.Reader, dest []LineFix) (n int) {
	for i := 0; i < len(dest); i++ {
		n += LineFixRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func LineFixWrite(w *x.Writer, v *LineFix) (n int) {

	// write field P1
	n += PointFixWrite(w, &v.P1)

	// write field P2
	n += PointFixWrite(w, &v.P2)
	return
}
func LineFixWriteN(w *x.Writer, src []LineFix) (n int) {
	for i := 0; i < len(src); i++ {
		n += LineFixWrite(w, &src[i])
	}
	return
}

type Triangle struct {
	P1 PointFix
	P2 PointFix
	P3 PointFix
}

type CTriangle struct {
	P1 PointFix
	P2 PointFix
	P3 PointFix
}

func TriangleRead(r *x.Reader, v *Triangle) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field P1
	n += PointFixRead(r, &v.P1)
	if r.Err() != nil {
		return
	}

	// read field P2
	n += PointFixRead(r, &v.P2)
	if r.Err() != nil {
		return
	}

	// read field P3
	n += PointFixRead(r, &v.P3)
	if r.Err() != nil {
		return
	}
	return
}

func TriangleReadN(r *x.Reader, dest []Triangle) (n int) {
	for i := 0; i < len(dest); i++ {
		n += TriangleRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func TriangleWrite(w *x.Writer, v *Triangle) (n int) {

	// write field P1
	n += PointFixWrite(w, &v.P1)

	// write field P2
	n += PointFixWrite(w, &v.P2)

	// write field P3
	n += PointFixWrite(w, &v.P3)
	return
}
func TriangleWriteN(w *x.Writer, src []Triangle) (n int) {
	for i := 0; i < len(src); i++ {
		n += TriangleWrite(w, &src[i])
	}
	return
}

type Trapezoid struct {
	Top    Fixed
	Bottom Fixed
	Left   LineFix
	Right  LineFix
}

type CTrapezoid struct {
	Top    Fixed
	Bottom Fixed
	Left   LineFix
	Right  LineFix
}

func TrapezoidRead(r *x.Reader, v *Trapezoid) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Top
	v.Top = Fixed(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Bottom
	v.Bottom = Fixed(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Left
	n += LineFixRead(r, &v.Left)
	if r.Err() != nil {
		return
	}

	// read field Right
	n += LineFixRead(r, &v.Right)
	if r.Err() != nil {
		return
	}
	return
}

func TrapezoidReadN(r *x.Reader, dest []Trapezoid) (n int) {
	for i := 0; i < len(dest); i++ {
		n += TrapezoidRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func TrapezoidWrite(w *x.Writer, v *Trapezoid) (n int) {

	// write field Top
	w.Write4b(uint32(v.Top))
	n += 4

	// write field Bottom
	w.Write4b(uint32(v.Bottom))
	n += 4

	// write field Left
	n += LineFixWrite(w, &v.Left)

	// write field Right
	n += LineFixWrite(w, &v.Right)
	return
}
func TrapezoidWriteN(w *x.Writer, src []Trapezoid) (n int) {
	for i := 0; i < len(src); i++ {
		n += TrapezoidWrite(w, &src[i])
	}
	return
}

type GlyphInfo struct {
	Width  uint16
	Height uint16
	X      int16
	Y      int16
	XOff   int16
	YOff   int16
}

type CGlyphInfo struct {
	Width  uint16
	Height uint16
	X      int16
	Y      int16
	XOff   int16
	YOff   int16
}

func GlyphInfoRead(r *x.Reader, v *GlyphInfo) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

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

	// read field XOff
	v.XOff = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field YOff
	v.YOff = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

func GlyphInfoReadN(r *x.Reader, dest []GlyphInfo) (n int) {
	for i := 0; i < len(dest); i++ {
		n += GlyphInfoRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func GlyphInfoWrite(w *x.Writer, v *GlyphInfo) (n int) {

	// write field Width
	w.Write2b(v.Width)
	n += 2

	// write field Height
	w.Write2b(v.Height)
	n += 2

	// write field X
	w.Write2b(uint16(v.X))
	n += 2

	// write field Y
	w.Write2b(uint16(v.Y))
	n += 2

	// write field XOff
	w.Write2b(uint16(v.XOff))
	n += 2

	// write field YOff
	w.Write2b(uint16(v.YOff))
	n += 2
	return
}
func GlyphInfoWriteN(w *x.Writer, src []GlyphInfo) (n int) {
	for i := 0; i < len(src); i++ {
		n += GlyphInfoWrite(w, &src[i])
	}
	return
}

const QueryVersionOpcode = 0

func QueryVersionRequestWrite(w *x.Writer, ClientMajorVersion uint32, ClientMinorVersion uint32) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field ClientMajorVersion
	w.Write4b(ClientMajorVersion)
	n += 4

	// write wire field ClientMinorVersion
	w.Write4b(ClientMinorVersion)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type QueryVersionReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence     uint16
	Length       uint32
	MajorVersion uint32
	MinorVersion uint32
	// Pad1 [16]uint8
}

func QueryVersionReplyRead(r *x.Reader, v *QueryVersionReply) (n int) {
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

	// read field MajorVersion
	v.MajorVersion = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field MinorVersion
	v.MinorVersion = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(16)
	n += 16
	if r.Err() != nil {
		return
	}
	return
}

type QueryVersionCookie uint64

func (cookie QueryVersionCookie) Reply(conn *x.Conn) (*QueryVersionReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewGenericError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply QueryVersionReply
	QueryVersionReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryVersion(conn *x.Conn, ClientMajorVersion uint32, ClientMinorVersion uint32) QueryVersionCookie {
	w := x.NewWriter()
	QueryVersionRequestWrite(w, ClientMajorVersion, ClientMinorVersion)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: QueryVersionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return QueryVersionCookie(seq)
}

func QueryVersionUnchecked(conn *x.Conn, ClientMajorVersion uint32, ClientMinorVersion uint32) QueryVersionCookie {
	w := x.NewWriter()
	QueryVersionRequestWrite(w, ClientMajorVersion, ClientMinorVersion)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: QueryVersionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return QueryVersionCookie(seq)
}

const QueryPictFormatsOpcode = 1

func QueryPictFormatsRequestWrite(w *x.Writer) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
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

type QueryPictFormatsReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence    uint16
	Length      uint32
	NumFormats  uint32
	NumScreens  uint32
	NumDepths   uint32
	NumVisuals  uint32
	NumSubpixel uint32
	// Pad1 [4]uint8
	Formats   []PictFormInfo
	Screens   []PictScreen
	Subpixels []uint32
}

func QueryPictFormatsReplyRead(r *x.Reader, v *QueryPictFormatsReply) (n int) {
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

	// read field NumFormats
	v.NumFormats = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field NumScreens
	v.NumScreens = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field NumDepths
	v.NumDepths = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field NumVisuals
	v.NumVisuals = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field NumSubpixel
	v.NumSubpixel = r.Read4b()
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

	// read field Formats
	v.Formats = make([]PictFormInfo, int(v.NumFormats))
	blockLen += PictFormInfoReadN(r, v.Formats)
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CPictFormInfo{}))

	// read field Screens

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
	v.Screens = make([]PictScreen, int(v.NumScreens))
	blockLen += PictScreenReadN(r, v.Screens)
	n += blockLen
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CPictScreen{}))

	// read field Subpixels

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
		dataLen := int(int(v.NumSubpixel))
		data := make([]uint32, dataLen)
		for i := 0; i < dataLen; i++ {
			data[i] = uint32(r.Read4b())
			if r.Err() != nil {
				return
			}
		}
		blockLen += dataLen * 4
		n += blockLen
		v.Subpixels = data
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

type QueryPictFormatsCookie uint64

func (cookie QueryPictFormatsCookie) Reply(conn *x.Conn) (*QueryPictFormatsReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewGenericError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply QueryPictFormatsReply
	QueryPictFormatsReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryPictFormats(conn *x.Conn) QueryPictFormatsCookie {
	w := x.NewWriter()
	QueryPictFormatsRequestWrite(w)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: QueryPictFormatsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return QueryPictFormatsCookie(seq)
}

func QueryPictFormatsUnchecked(conn *x.Conn) QueryPictFormatsCookie {
	w := x.NewWriter()
	QueryPictFormatsRequestWrite(w)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: QueryPictFormatsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return QueryPictFormatsCookie(seq)
}

const QueryPictIndexValuesOpcode = 2

func QueryPictIndexValuesRequestWrite(w *x.Writer, Format PictFormat) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Format
	w.Write4b(uint32(Format))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type QueryPictIndexValuesReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence  uint16
	Length    uint32
	NumValues uint32
	// Pad1 [20]uint8
	Values []IndexValue
}

func QueryPictIndexValuesReplyRead(r *x.Reader, v *QueryPictIndexValuesReply) (n int) {
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

	// read field NumValues
	v.NumValues = r.Read4b()
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

	// read field Values
	v.Values = make([]IndexValue, int(v.NumValues))
	blockLen += IndexValueReadN(r, v.Values)
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CIndexValue{}))

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

type QueryPictIndexValuesCookie uint64

func (cookie QueryPictIndexValuesCookie) Reply(conn *x.Conn) (*QueryPictIndexValuesReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewGenericError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply QueryPictIndexValuesReply
	QueryPictIndexValuesReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryPictIndexValues(conn *x.Conn, Format PictFormat) QueryPictIndexValuesCookie {
	w := x.NewWriter()
	QueryPictIndexValuesRequestWrite(w, Format)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: QueryPictIndexValuesOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return QueryPictIndexValuesCookie(seq)
}

func QueryPictIndexValuesUnchecked(conn *x.Conn, Format PictFormat) QueryPictIndexValuesCookie {
	w := x.NewWriter()
	QueryPictIndexValuesRequestWrite(w, Format)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: QueryPictIndexValuesOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return QueryPictIndexValuesCookie(seq)
}

type CreatePictureValueList struct {
	Repeat           uint32
	Alphamap         Picture
	Alphaxorigin     int32
	Alphayorigin     int32
	Clipxorigin      int32
	Clipyorigin      int32
	Clipmask         x.Pixmap
	Graphicsexposure uint32
	Subwindowmode    uint32
	Polyedge         uint32
	Polymode         uint32
	Dither           x.Atom
	Componentalpha   uint32
}

// _go_switch_write SwitchType "xcb.Render.CreatePicture.value_list"
func CreatePictureValueListSerialize(w *x.Writer, ValueMask uint32, _aux *CreatePictureValueList) {
	n := 0
	xgbPaddingOffset := 0
	alignTo := 0
	xgbPad := 0
	// switch expr: int(ValueMask)

	// case_field Field with type BitcaseType "xcb.Render.CreatePicture.value_list.bitcase1"
	if (int(ValueMask) & CPRepeat) != 0 {
		// todo
		// field Field "repeat" in BitcaseType "xcb.Render.CreatePicture.value_list.bitcase1"
		w.Write4b(_aux.Repeat)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.CreatePicture.value_list.bitcase2"
	if (int(ValueMask) & CPAlphaMap) != 0 {
		// todo
		// field Field "alphamap" in BitcaseType "xcb.Render.CreatePicture.value_list.bitcase2"
		w.Write4b(uint32(_aux.Alphamap))
		n += 4
		alignTo = int(unsafe.Alignof(Picture(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.CreatePicture.value_list.bitcase3"
	if (int(ValueMask) & CPAlphaXOrigin) != 0 {
		// todo
		// field Field "alphaxorigin" in BitcaseType "xcb.Render.CreatePicture.value_list.bitcase3"
		w.Write4b(uint32(_aux.Alphaxorigin))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.CreatePicture.value_list.bitcase4"
	if (int(ValueMask) & CPAlphaYOrigin) != 0 {
		// todo
		// field Field "alphayorigin" in BitcaseType "xcb.Render.CreatePicture.value_list.bitcase4"
		w.Write4b(uint32(_aux.Alphayorigin))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.CreatePicture.value_list.bitcase5"
	if (int(ValueMask) & CPClipXOrigin) != 0 {
		// todo
		// field Field "clipxorigin" in BitcaseType "xcb.Render.CreatePicture.value_list.bitcase5"
		w.Write4b(uint32(_aux.Clipxorigin))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.CreatePicture.value_list.bitcase6"
	if (int(ValueMask) & CPClipYOrigin) != 0 {
		// todo
		// field Field "clipyorigin" in BitcaseType "xcb.Render.CreatePicture.value_list.bitcase6"
		w.Write4b(uint32(_aux.Clipyorigin))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.CreatePicture.value_list.bitcase7"
	if (int(ValueMask) & CPClipMask) != 0 {
		// todo
		// field Field "clipmask" in BitcaseType "xcb.Render.CreatePicture.value_list.bitcase7"
		w.Write4b(uint32(_aux.Clipmask))
		n += 4
		alignTo = int(unsafe.Alignof(x.Pixmap(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.CreatePicture.value_list.bitcase8"
	if (int(ValueMask) & CPGraphicsExposure) != 0 {
		// todo
		// field Field "graphicsexposure" in BitcaseType "xcb.Render.CreatePicture.value_list.bitcase8"
		w.Write4b(_aux.Graphicsexposure)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.CreatePicture.value_list.bitcase9"
	if (int(ValueMask) & CPSubwindowMode) != 0 {
		// todo
		// field Field "subwindowmode" in BitcaseType "xcb.Render.CreatePicture.value_list.bitcase9"
		w.Write4b(_aux.Subwindowmode)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.CreatePicture.value_list.bitcase10"
	if (int(ValueMask) & CPPolyEdge) != 0 {
		// todo
		// field Field "polyedge" in BitcaseType "xcb.Render.CreatePicture.value_list.bitcase10"
		w.Write4b(_aux.Polyedge)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.CreatePicture.value_list.bitcase11"
	if (int(ValueMask) & CPPolyMode) != 0 {
		// todo
		// field Field "polymode" in BitcaseType "xcb.Render.CreatePicture.value_list.bitcase11"
		w.Write4b(_aux.Polymode)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.CreatePicture.value_list.bitcase12"
	if (int(ValueMask) & CPDither) != 0 {
		// todo
		// field Field "dither" in BitcaseType "xcb.Render.CreatePicture.value_list.bitcase12"
		w.Write4b(uint32(_aux.Dither))
		n += 4
		alignTo = int(unsafe.Alignof(x.Atom(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.CreatePicture.value_list.bitcase13"
	if (int(ValueMask) & CPComponentAlpha) != 0 {
		// todo
		// field Field "componentalpha" in BitcaseType "xcb.Render.CreatePicture.value_list.bitcase13"
		w.Write4b(_aux.Componentalpha)
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

const CreatePictureOpcode = 4

func CreatePictureRequestWrite(w *x.Writer, Pid Picture, Drawable x.Drawable, Format PictFormat, ValueMask uint32, ValueList *CreatePictureValueList) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
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

	// write wire field Format
	w.Write4b(uint32(Format))
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
	CreatePictureValueListSerialize(w, ValueMask, ValueList)
}

func CreatePicture(conn *x.Conn, Pid Picture, Drawable x.Drawable, Format PictFormat, ValueMask uint32, ValueList *CreatePictureValueList) x.VoidCookie {
	w := x.NewWriter()
	CreatePictureRequestWrite(w, Pid, Drawable, Format, ValueMask, ValueList)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreatePictureOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CreatePictureChecked(conn *x.Conn, Pid Picture, Drawable x.Drawable, Format PictFormat, ValueMask uint32, ValueList *CreatePictureValueList) x.VoidCookie {
	w := x.NewWriter()
	CreatePictureRequestWrite(w, Pid, Drawable, Format, ValueMask, ValueList)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreatePictureOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

type ChangePictureValueList struct {
	Repeat           uint32
	Alphamap         Picture
	Alphaxorigin     int32
	Alphayorigin     int32
	Clipxorigin      int32
	Clipyorigin      int32
	Clipmask         x.Pixmap
	Graphicsexposure uint32
	Subwindowmode    uint32
	Polyedge         uint32
	Polymode         uint32
	Dither           x.Atom
	Componentalpha   uint32
}

// _go_switch_write SwitchType "xcb.Render.ChangePicture.value_list"
func ChangePictureValueListSerialize(w *x.Writer, ValueMask uint32, _aux *ChangePictureValueList) {
	n := 0
	xgbPaddingOffset := 0
	alignTo := 0
	xgbPad := 0
	// switch expr: int(ValueMask)

	// case_field Field with type BitcaseType "xcb.Render.ChangePicture.value_list.bitcase1"
	if (int(ValueMask) & CPRepeat) != 0 {
		// todo
		// field Field "repeat" in BitcaseType "xcb.Render.ChangePicture.value_list.bitcase1"
		w.Write4b(_aux.Repeat)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.ChangePicture.value_list.bitcase2"
	if (int(ValueMask) & CPAlphaMap) != 0 {
		// todo
		// field Field "alphamap" in BitcaseType "xcb.Render.ChangePicture.value_list.bitcase2"
		w.Write4b(uint32(_aux.Alphamap))
		n += 4
		alignTo = int(unsafe.Alignof(Picture(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.ChangePicture.value_list.bitcase3"
	if (int(ValueMask) & CPAlphaXOrigin) != 0 {
		// todo
		// field Field "alphaxorigin" in BitcaseType "xcb.Render.ChangePicture.value_list.bitcase3"
		w.Write4b(uint32(_aux.Alphaxorigin))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.ChangePicture.value_list.bitcase4"
	if (int(ValueMask) & CPAlphaYOrigin) != 0 {
		// todo
		// field Field "alphayorigin" in BitcaseType "xcb.Render.ChangePicture.value_list.bitcase4"
		w.Write4b(uint32(_aux.Alphayorigin))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.ChangePicture.value_list.bitcase5"
	if (int(ValueMask) & CPClipXOrigin) != 0 {
		// todo
		// field Field "clipxorigin" in BitcaseType "xcb.Render.ChangePicture.value_list.bitcase5"
		w.Write4b(uint32(_aux.Clipxorigin))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.ChangePicture.value_list.bitcase6"
	if (int(ValueMask) & CPClipYOrigin) != 0 {
		// todo
		// field Field "clipyorigin" in BitcaseType "xcb.Render.ChangePicture.value_list.bitcase6"
		w.Write4b(uint32(_aux.Clipyorigin))
		n += 4
		alignTo = int(unsafe.Alignof(int32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.ChangePicture.value_list.bitcase7"
	if (int(ValueMask) & CPClipMask) != 0 {
		// todo
		// field Field "clipmask" in BitcaseType "xcb.Render.ChangePicture.value_list.bitcase7"
		w.Write4b(uint32(_aux.Clipmask))
		n += 4
		alignTo = int(unsafe.Alignof(x.Pixmap(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.ChangePicture.value_list.bitcase8"
	if (int(ValueMask) & CPGraphicsExposure) != 0 {
		// todo
		// field Field "graphicsexposure" in BitcaseType "xcb.Render.ChangePicture.value_list.bitcase8"
		w.Write4b(_aux.Graphicsexposure)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.ChangePicture.value_list.bitcase9"
	if (int(ValueMask) & CPSubwindowMode) != 0 {
		// todo
		// field Field "subwindowmode" in BitcaseType "xcb.Render.ChangePicture.value_list.bitcase9"
		w.Write4b(_aux.Subwindowmode)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.ChangePicture.value_list.bitcase10"
	if (int(ValueMask) & CPPolyEdge) != 0 {
		// todo
		// field Field "polyedge" in BitcaseType "xcb.Render.ChangePicture.value_list.bitcase10"
		w.Write4b(_aux.Polyedge)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.ChangePicture.value_list.bitcase11"
	if (int(ValueMask) & CPPolyMode) != 0 {
		// todo
		// field Field "polymode" in BitcaseType "xcb.Render.ChangePicture.value_list.bitcase11"
		w.Write4b(_aux.Polymode)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.ChangePicture.value_list.bitcase12"
	if (int(ValueMask) & CPDither) != 0 {
		// todo
		// field Field "dither" in BitcaseType "xcb.Render.ChangePicture.value_list.bitcase12"
		w.Write4b(uint32(_aux.Dither))
		n += 4
		alignTo = int(unsafe.Alignof(x.Atom(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.Render.ChangePicture.value_list.bitcase13"
	if (int(ValueMask) & CPComponentAlpha) != 0 {
		// todo
		// field Field "componentalpha" in BitcaseType "xcb.Render.ChangePicture.value_list.bitcase13"
		w.Write4b(_aux.Componentalpha)
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

const ChangePictureOpcode = 5

func ChangePictureRequestWrite(w *x.Writer, Picture Picture, ValueMask uint32, ValueList *ChangePictureValueList) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Picture
	w.Write4b(uint32(Picture))
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
	ChangePictureValueListSerialize(w, ValueMask, ValueList)
}

func ChangePicture(conn *x.Conn, Picture Picture, ValueMask uint32, ValueList *ChangePictureValueList) x.VoidCookie {
	w := x.NewWriter()
	ChangePictureRequestWrite(w, Picture, ValueMask, ValueList)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: ChangePictureOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func ChangePictureChecked(conn *x.Conn, Picture Picture, ValueMask uint32, ValueList *ChangePictureValueList) x.VoidCookie {
	w := x.NewWriter()
	ChangePictureRequestWrite(w, Picture, ValueMask, ValueList)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: ChangePictureOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const SetPictureClipRectanglesOpcode = 6

func SetPictureClipRectanglesRequestWrite(w *x.Writer, Picture Picture, ClipXOrigin int16, ClipYOrigin int16, RectanglesLen uint32, Rectangles []x.Rectangle) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Picture
	w.Write4b(uint32(Picture))
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
	n += x.RectangleWriteN(w, Rectangles)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SetPictureClipRectangles(conn *x.Conn, Picture Picture, ClipXOrigin int16, ClipYOrigin int16, RectanglesLen uint32, Rectangles []x.Rectangle) x.VoidCookie {
	w := x.NewWriter()
	SetPictureClipRectanglesRequestWrite(w, Picture, ClipXOrigin, ClipYOrigin, RectanglesLen, Rectangles)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SetPictureClipRectanglesOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func SetPictureClipRectanglesChecked(conn *x.Conn, Picture Picture, ClipXOrigin int16, ClipYOrigin int16, RectanglesLen uint32, Rectangles []x.Rectangle) x.VoidCookie {
	w := x.NewWriter()
	SetPictureClipRectanglesRequestWrite(w, Picture, ClipXOrigin, ClipYOrigin, RectanglesLen, Rectangles)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SetPictureClipRectanglesOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const FreePictureOpcode = 7

func FreePictureRequestWrite(w *x.Writer, Picture Picture) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Picture
	w.Write4b(uint32(Picture))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func FreePicture(conn *x.Conn, Picture Picture) x.VoidCookie {
	w := x.NewWriter()
	FreePictureRequestWrite(w, Picture)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: FreePictureOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func FreePictureChecked(conn *x.Conn, Picture Picture) x.VoidCookie {
	w := x.NewWriter()
	FreePictureRequestWrite(w, Picture)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: FreePictureOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const CompositeOpcode = 8

func CompositeRequestWrite(w *x.Writer, Op uint8, Src Picture, Mask Picture, Dst Picture, SrcX int16, SrcY int16, MaskX int16, MaskY int16, DstX int16, DstY int16, Width uint16, Height uint16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Op
	w.Write1b(Op)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3

	// write wire field Src
	w.Write4b(uint32(Src))
	n += 4

	// write wire field Mask
	w.Write4b(uint32(Mask))
	n += 4

	// write wire field Dst
	w.Write4b(uint32(Dst))
	n += 4

	// write wire field SrcX
	w.Write2b(uint16(SrcX))
	n += 2

	// write wire field SrcY
	w.Write2b(uint16(SrcY))
	n += 2

	// write wire field MaskX
	w.Write2b(uint16(MaskX))
	n += 2

	// write wire field MaskY
	w.Write2b(uint16(MaskY))
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

func Composite(conn *x.Conn, Op uint8, Src Picture, Mask Picture, Dst Picture, SrcX int16, SrcY int16, MaskX int16, MaskY int16, DstX int16, DstY int16, Width uint16, Height uint16) x.VoidCookie {
	w := x.NewWriter()
	CompositeRequestWrite(w, Op, Src, Mask, Dst, SrcX, SrcY, MaskX, MaskY, DstX, DstY, Width, Height)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CompositeOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CompositeChecked(conn *x.Conn, Op uint8, Src Picture, Mask Picture, Dst Picture, SrcX int16, SrcY int16, MaskX int16, MaskY int16, DstX int16, DstY int16, Width uint16, Height uint16) x.VoidCookie {
	w := x.NewWriter()
	CompositeRequestWrite(w, Op, Src, Mask, Dst, SrcX, SrcY, MaskX, MaskY, DstX, DstY, Width, Height)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CompositeOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const TrapezoidsOpcode = 10

func TrapezoidsRequestWrite(w *x.Writer, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, SrcX int16, SrcY int16, TrapsLen uint32, Traps []Trapezoid) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Op
	w.Write1b(Op)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3

	// write wire field Src
	w.Write4b(uint32(Src))
	n += 4

	// write wire field Dst
	w.Write4b(uint32(Dst))
	n += 4

	// write wire field MaskFormat
	w.Write4b(uint32(MaskFormat))
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

	// write param field Traps
	n += TrapezoidWriteN(w, Traps)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func Trapezoids(conn *x.Conn, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, SrcX int16, SrcY int16, TrapsLen uint32, Traps []Trapezoid) x.VoidCookie {
	w := x.NewWriter()
	TrapezoidsRequestWrite(w, Op, Src, Dst, MaskFormat, SrcX, SrcY, TrapsLen, Traps)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: TrapezoidsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func TrapezoidsChecked(conn *x.Conn, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, SrcX int16, SrcY int16, TrapsLen uint32, Traps []Trapezoid) x.VoidCookie {
	w := x.NewWriter()
	TrapezoidsRequestWrite(w, Op, Src, Dst, MaskFormat, SrcX, SrcY, TrapsLen, Traps)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: TrapezoidsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const TrianglesOpcode = 11

func TrianglesRequestWrite(w *x.Writer, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, SrcX int16, SrcY int16, TrianglesLen uint32, Triangles []Triangle) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Op
	w.Write1b(Op)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3

	// write wire field Src
	w.Write4b(uint32(Src))
	n += 4

	// write wire field Dst
	w.Write4b(uint32(Dst))
	n += 4

	// write wire field MaskFormat
	w.Write4b(uint32(MaskFormat))
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

	// write param field Triangles
	n += TriangleWriteN(w, Triangles)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func Triangles(conn *x.Conn, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, SrcX int16, SrcY int16, TrianglesLen uint32, Triangles []Triangle) x.VoidCookie {
	w := x.NewWriter()
	TrianglesRequestWrite(w, Op, Src, Dst, MaskFormat, SrcX, SrcY, TrianglesLen, Triangles)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: TrianglesOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func TrianglesChecked(conn *x.Conn, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, SrcX int16, SrcY int16, TrianglesLen uint32, Triangles []Triangle) x.VoidCookie {
	w := x.NewWriter()
	TrianglesRequestWrite(w, Op, Src, Dst, MaskFormat, SrcX, SrcY, TrianglesLen, Triangles)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: TrianglesOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const TriStripOpcode = 12

func TriStripRequestWrite(w *x.Writer, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, SrcX int16, SrcY int16, PointsLen uint32, Points []PointFix) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Op
	w.Write1b(Op)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3

	// write wire field Src
	w.Write4b(uint32(Src))
	n += 4

	// write wire field Dst
	w.Write4b(uint32(Dst))
	n += 4

	// write wire field MaskFormat
	w.Write4b(uint32(MaskFormat))
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

	// write param field Points
	n += PointFixWriteN(w, Points)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func TriStrip(conn *x.Conn, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, SrcX int16, SrcY int16, PointsLen uint32, Points []PointFix) x.VoidCookie {
	w := x.NewWriter()
	TriStripRequestWrite(w, Op, Src, Dst, MaskFormat, SrcX, SrcY, PointsLen, Points)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: TriStripOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func TriStripChecked(conn *x.Conn, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, SrcX int16, SrcY int16, PointsLen uint32, Points []PointFix) x.VoidCookie {
	w := x.NewWriter()
	TriStripRequestWrite(w, Op, Src, Dst, MaskFormat, SrcX, SrcY, PointsLen, Points)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: TriStripOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const TriFanOpcode = 13

func TriFanRequestWrite(w *x.Writer, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, SrcX int16, SrcY int16, PointsLen uint32, Points []PointFix) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Op
	w.Write1b(Op)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3

	// write wire field Src
	w.Write4b(uint32(Src))
	n += 4

	// write wire field Dst
	w.Write4b(uint32(Dst))
	n += 4

	// write wire field MaskFormat
	w.Write4b(uint32(MaskFormat))
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

	// write param field Points
	n += PointFixWriteN(w, Points)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func TriFan(conn *x.Conn, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, SrcX int16, SrcY int16, PointsLen uint32, Points []PointFix) x.VoidCookie {
	w := x.NewWriter()
	TriFanRequestWrite(w, Op, Src, Dst, MaskFormat, SrcX, SrcY, PointsLen, Points)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: TriFanOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func TriFanChecked(conn *x.Conn, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, SrcX int16, SrcY int16, PointsLen uint32, Points []PointFix) x.VoidCookie {
	w := x.NewWriter()
	TriFanRequestWrite(w, Op, Src, Dst, MaskFormat, SrcX, SrcY, PointsLen, Points)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: TriFanOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const CreateGlyphSetOpcode = 17

func CreateGlyphSetRequestWrite(w *x.Writer, Gsid GlyphSet, Format PictFormat) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Gsid
	w.Write4b(uint32(Gsid))
	n += 4

	// write wire field Format
	w.Write4b(uint32(Format))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CreateGlyphSet(conn *x.Conn, Gsid GlyphSet, Format PictFormat) x.VoidCookie {
	w := x.NewWriter()
	CreateGlyphSetRequestWrite(w, Gsid, Format)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateGlyphSetOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CreateGlyphSetChecked(conn *x.Conn, Gsid GlyphSet, Format PictFormat) x.VoidCookie {
	w := x.NewWriter()
	CreateGlyphSetRequestWrite(w, Gsid, Format)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateGlyphSetOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const ReferenceGlyphSetOpcode = 18

func ReferenceGlyphSetRequestWrite(w *x.Writer, Gsid GlyphSet, Existing GlyphSet) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Gsid
	w.Write4b(uint32(Gsid))
	n += 4

	// write wire field Existing
	w.Write4b(uint32(Existing))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func ReferenceGlyphSet(conn *x.Conn, Gsid GlyphSet, Existing GlyphSet) x.VoidCookie {
	w := x.NewWriter()
	ReferenceGlyphSetRequestWrite(w, Gsid, Existing)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: ReferenceGlyphSetOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func ReferenceGlyphSetChecked(conn *x.Conn, Gsid GlyphSet, Existing GlyphSet) x.VoidCookie {
	w := x.NewWriter()
	ReferenceGlyphSetRequestWrite(w, Gsid, Existing)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: ReferenceGlyphSetOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const FreeGlyphSetOpcode = 19

func FreeGlyphSetRequestWrite(w *x.Writer, Glyphset GlyphSet) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Glyphset
	w.Write4b(uint32(Glyphset))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func FreeGlyphSet(conn *x.Conn, Glyphset GlyphSet) x.VoidCookie {
	w := x.NewWriter()
	FreeGlyphSetRequestWrite(w, Glyphset)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: FreeGlyphSetOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func FreeGlyphSetChecked(conn *x.Conn, Glyphset GlyphSet) x.VoidCookie {
	w := x.NewWriter()
	FreeGlyphSetRequestWrite(w, Glyphset)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: FreeGlyphSetOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const AddGlyphsOpcode = 20

func AddGlyphsRequestWrite(w *x.Writer, Glyphset GlyphSet, GlyphsLen uint32, Glyphids []uint32, Glyphs []GlyphInfo, DataLen uint32, Data []uint8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Glyphset
	w.Write4b(uint32(Glyphset))
	n += 4

	// write wire field GlyphsLen
	w.Write4b(GlyphsLen)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Glyphids
	{
		_dataLen := int(GlyphsLen)
		for i := 0; i < _dataLen; i++ {
			w.Write4b(uint32(Glyphids[i]))
		}
		n += _dataLen * 4
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Glyphs
	n += GlyphInfoWriteN(w, Glyphs)
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

func AddGlyphs(conn *x.Conn, Glyphset GlyphSet, GlyphsLen uint32, Glyphids []uint32, Glyphs []GlyphInfo, DataLen uint32, Data []uint8) x.VoidCookie {
	w := x.NewWriter()
	AddGlyphsRequestWrite(w, Glyphset, GlyphsLen, Glyphids, Glyphs, DataLen, Data)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: AddGlyphsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func AddGlyphsChecked(conn *x.Conn, Glyphset GlyphSet, GlyphsLen uint32, Glyphids []uint32, Glyphs []GlyphInfo, DataLen uint32, Data []uint8) x.VoidCookie {
	w := x.NewWriter()
	AddGlyphsRequestWrite(w, Glyphset, GlyphsLen, Glyphids, Glyphs, DataLen, Data)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: AddGlyphsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const FreeGlyphsOpcode = 22

func FreeGlyphsRequestWrite(w *x.Writer, Glyphset GlyphSet, GlyphsLen uint32, Glyphs []Glyph) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Glyphset
	w.Write4b(uint32(Glyphset))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Glyphs
	{
		_dataLen := int(GlyphsLen)
		for i := 0; i < _dataLen; i++ {
			w.Write4b(uint32(Glyphs[i]))
		}
		n += _dataLen * 4
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func FreeGlyphs(conn *x.Conn, Glyphset GlyphSet, GlyphsLen uint32, Glyphs []Glyph) x.VoidCookie {
	w := x.NewWriter()
	FreeGlyphsRequestWrite(w, Glyphset, GlyphsLen, Glyphs)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: FreeGlyphsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func FreeGlyphsChecked(conn *x.Conn, Glyphset GlyphSet, GlyphsLen uint32, Glyphs []Glyph) x.VoidCookie {
	w := x.NewWriter()
	FreeGlyphsRequestWrite(w, Glyphset, GlyphsLen, Glyphs)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: FreeGlyphsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const CompositeGlyphs8Opcode = 23

func CompositeGlyphs8RequestWrite(w *x.Writer, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, Glyphset GlyphSet, SrcX int16, SrcY int16, GlyphcmdsLen uint32, Glyphcmds []uint8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Op
	w.Write1b(Op)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3

	// write wire field Src
	w.Write4b(uint32(Src))
	n += 4

	// write wire field Dst
	w.Write4b(uint32(Dst))
	n += 4

	// write wire field MaskFormat
	w.Write4b(uint32(MaskFormat))
	n += 4

	// write wire field Glyphset
	w.Write4b(uint32(Glyphset))
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

	// write param field Glyphcmds
	{
		_dataLen := int(GlyphcmdsLen)
		for i := 0; i < _dataLen; i++ {
			w.Write1b(uint8(Glyphcmds[i]))
		}
		n += _dataLen * 1
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CompositeGlyphs8(conn *x.Conn, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, Glyphset GlyphSet, SrcX int16, SrcY int16, GlyphcmdsLen uint32, Glyphcmds []uint8) x.VoidCookie {
	w := x.NewWriter()
	CompositeGlyphs8RequestWrite(w, Op, Src, Dst, MaskFormat, Glyphset, SrcX, SrcY, GlyphcmdsLen, Glyphcmds)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CompositeGlyphs8Opcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CompositeGlyphs8Checked(conn *x.Conn, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, Glyphset GlyphSet, SrcX int16, SrcY int16, GlyphcmdsLen uint32, Glyphcmds []uint8) x.VoidCookie {
	w := x.NewWriter()
	CompositeGlyphs8RequestWrite(w, Op, Src, Dst, MaskFormat, Glyphset, SrcX, SrcY, GlyphcmdsLen, Glyphcmds)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CompositeGlyphs8Opcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const CompositeGlyphs16Opcode = 24

func CompositeGlyphs16RequestWrite(w *x.Writer, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, Glyphset GlyphSet, SrcX int16, SrcY int16, GlyphcmdsLen uint32, Glyphcmds []uint8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Op
	w.Write1b(Op)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3

	// write wire field Src
	w.Write4b(uint32(Src))
	n += 4

	// write wire field Dst
	w.Write4b(uint32(Dst))
	n += 4

	// write wire field MaskFormat
	w.Write4b(uint32(MaskFormat))
	n += 4

	// write wire field Glyphset
	w.Write4b(uint32(Glyphset))
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

	// write param field Glyphcmds
	{
		_dataLen := int(GlyphcmdsLen)
		for i := 0; i < _dataLen; i++ {
			w.Write1b(uint8(Glyphcmds[i]))
		}
		n += _dataLen * 1
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CompositeGlyphs16(conn *x.Conn, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, Glyphset GlyphSet, SrcX int16, SrcY int16, GlyphcmdsLen uint32, Glyphcmds []uint8) x.VoidCookie {
	w := x.NewWriter()
	CompositeGlyphs16RequestWrite(w, Op, Src, Dst, MaskFormat, Glyphset, SrcX, SrcY, GlyphcmdsLen, Glyphcmds)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CompositeGlyphs16Opcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CompositeGlyphs16Checked(conn *x.Conn, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, Glyphset GlyphSet, SrcX int16, SrcY int16, GlyphcmdsLen uint32, Glyphcmds []uint8) x.VoidCookie {
	w := x.NewWriter()
	CompositeGlyphs16RequestWrite(w, Op, Src, Dst, MaskFormat, Glyphset, SrcX, SrcY, GlyphcmdsLen, Glyphcmds)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CompositeGlyphs16Opcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const CompositeGlyphs32Opcode = 25

func CompositeGlyphs32RequestWrite(w *x.Writer, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, Glyphset GlyphSet, SrcX int16, SrcY int16, GlyphcmdsLen uint32, Glyphcmds []uint8) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Op
	w.Write1b(Op)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3

	// write wire field Src
	w.Write4b(uint32(Src))
	n += 4

	// write wire field Dst
	w.Write4b(uint32(Dst))
	n += 4

	// write wire field MaskFormat
	w.Write4b(uint32(MaskFormat))
	n += 4

	// write wire field Glyphset
	w.Write4b(uint32(Glyphset))
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

	// write param field Glyphcmds
	{
		_dataLen := int(GlyphcmdsLen)
		for i := 0; i < _dataLen; i++ {
			w.Write1b(uint8(Glyphcmds[i]))
		}
		n += _dataLen * 1
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CompositeGlyphs32(conn *x.Conn, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, Glyphset GlyphSet, SrcX int16, SrcY int16, GlyphcmdsLen uint32, Glyphcmds []uint8) x.VoidCookie {
	w := x.NewWriter()
	CompositeGlyphs32RequestWrite(w, Op, Src, Dst, MaskFormat, Glyphset, SrcX, SrcY, GlyphcmdsLen, Glyphcmds)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CompositeGlyphs32Opcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CompositeGlyphs32Checked(conn *x.Conn, Op uint8, Src Picture, Dst Picture, MaskFormat PictFormat, Glyphset GlyphSet, SrcX int16, SrcY int16, GlyphcmdsLen uint32, Glyphcmds []uint8) x.VoidCookie {
	w := x.NewWriter()
	CompositeGlyphs32RequestWrite(w, Op, Src, Dst, MaskFormat, Glyphset, SrcX, SrcY, GlyphcmdsLen, Glyphcmds)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CompositeGlyphs32Opcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const FillRectanglesOpcode = 26

func FillRectanglesRequestWrite(w *x.Writer, Op uint8, Dst Picture, Color Color, RectsLen uint32, Rects []x.Rectangle) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Op
	w.Write1b(Op)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3

	// write wire field Dst
	w.Write4b(uint32(Dst))
	n += 4

	// write wire field Color
	// TODO RequestWriteDefine container_FS
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Rects
	n += x.RectangleWriteN(w, Rects)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func FillRectangles(conn *x.Conn, Op uint8, Dst Picture, Color Color, RectsLen uint32, Rects []x.Rectangle) x.VoidCookie {
	w := x.NewWriter()
	FillRectanglesRequestWrite(w, Op, Dst, Color, RectsLen, Rects)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: FillRectanglesOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func FillRectanglesChecked(conn *x.Conn, Op uint8, Dst Picture, Color Color, RectsLen uint32, Rects []x.Rectangle) x.VoidCookie {
	w := x.NewWriter()
	FillRectanglesRequestWrite(w, Op, Dst, Color, RectsLen, Rects)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: FillRectanglesOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const CreateCursorOpcode = 27

func CreateCursorRequestWrite(w *x.Writer, Cid x.Cursor, Source Picture, X uint16, Y uint16) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
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

func CreateCursor(conn *x.Conn, Cid x.Cursor, Source Picture, X uint16, Y uint16) x.VoidCookie {
	w := x.NewWriter()
	CreateCursorRequestWrite(w, Cid, Source, X, Y)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateCursorOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CreateCursorChecked(conn *x.Conn, Cid x.Cursor, Source Picture, X uint16, Y uint16) x.VoidCookie {
	w := x.NewWriter()
	CreateCursorRequestWrite(w, Cid, Source, X, Y)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateCursorOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

type Transform struct {
	Matrix11 Fixed
	Matrix12 Fixed
	Matrix13 Fixed
	Matrix21 Fixed
	Matrix22 Fixed
	Matrix23 Fixed
	Matrix31 Fixed
	Matrix32 Fixed
	Matrix33 Fixed
}

type CTransform struct {
	Matrix11 Fixed
	Matrix12 Fixed
	Matrix13 Fixed
	Matrix21 Fixed
	Matrix22 Fixed
	Matrix23 Fixed
	Matrix31 Fixed
	Matrix32 Fixed
	Matrix33 Fixed
}

func TransformRead(r *x.Reader, v *Transform) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Matrix11
	v.Matrix11 = Fixed(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Matrix12
	v.Matrix12 = Fixed(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Matrix13
	v.Matrix13 = Fixed(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Matrix21
	v.Matrix21 = Fixed(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Matrix22
	v.Matrix22 = Fixed(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Matrix23
	v.Matrix23 = Fixed(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Matrix31
	v.Matrix31 = Fixed(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Matrix32
	v.Matrix32 = Fixed(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Matrix33
	v.Matrix33 = Fixed(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}
	return
}

func TransformReadN(r *x.Reader, dest []Transform) (n int) {
	for i := 0; i < len(dest); i++ {
		n += TransformRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func TransformWrite(w *x.Writer, v *Transform) (n int) {

	// write field Matrix11
	w.Write4b(uint32(v.Matrix11))
	n += 4

	// write field Matrix12
	w.Write4b(uint32(v.Matrix12))
	n += 4

	// write field Matrix13
	w.Write4b(uint32(v.Matrix13))
	n += 4

	// write field Matrix21
	w.Write4b(uint32(v.Matrix21))
	n += 4

	// write field Matrix22
	w.Write4b(uint32(v.Matrix22))
	n += 4

	// write field Matrix23
	w.Write4b(uint32(v.Matrix23))
	n += 4

	// write field Matrix31
	w.Write4b(uint32(v.Matrix31))
	n += 4

	// write field Matrix32
	w.Write4b(uint32(v.Matrix32))
	n += 4

	// write field Matrix33
	w.Write4b(uint32(v.Matrix33))
	n += 4
	return
}
func TransformWriteN(w *x.Writer, src []Transform) (n int) {
	for i := 0; i < len(src); i++ {
		n += TransformWrite(w, &src[i])
	}
	return
}

const SetPictureTransformOpcode = 28

func SetPictureTransformRequestWrite(w *x.Writer, Picture Picture, Transform Transform) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Picture
	w.Write4b(uint32(Picture))
	n += 4

	// write wire field Transform
	// TODO RequestWriteDefine container_FS
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SetPictureTransform(conn *x.Conn, Picture Picture, Transform Transform) x.VoidCookie {
	w := x.NewWriter()
	SetPictureTransformRequestWrite(w, Picture, Transform)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SetPictureTransformOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func SetPictureTransformChecked(conn *x.Conn, Picture Picture, Transform Transform) x.VoidCookie {
	w := x.NewWriter()
	SetPictureTransformRequestWrite(w, Picture, Transform)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SetPictureTransformOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const QueryFiltersOpcode = 29

func QueryFiltersRequestWrite(w *x.Writer, Drawable x.Drawable) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
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

type QueryFiltersReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence   uint16
	Length     uint32
	NumAliases uint32
	NumFilters uint32
	// Pad1 [16]uint8
	Aliases []uint16
	Filters []x.Str
}

func QueryFiltersReplyRead(r *x.Reader, v *QueryFiltersReply) (n int) {
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

	// read field NumAliases
	v.NumAliases = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field NumFilters
	v.NumFilters = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(16)
	n += 16
	if r.Err() != nil {
		return
	}

	// read field Aliases
	{
		dataLen := int(int(v.NumAliases))
		data := make([]uint16, dataLen)
		for i := 0; i < dataLen; i++ {
			data[i] = uint16(r.Read2b())
			if r.Err() != nil {
				return
			}
		}
		blockLen += dataLen * 2
		n += blockLen
		v.Aliases = data
	}
	alignTo = int(unsafe.Alignof(uint16(0)))

	// read field Filters

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
	v.Filters = make([]x.Str, int(v.NumFilters))
	blockLen += x.StrReadN(r, v.Filters)
	n += blockLen
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(x.CStr{}))

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

type QueryFiltersCookie uint64

func (cookie QueryFiltersCookie) Reply(conn *x.Conn) (*QueryFiltersReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewGenericError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply QueryFiltersReply
	QueryFiltersReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryFilters(conn *x.Conn, Drawable x.Drawable) QueryFiltersCookie {
	w := x.NewWriter()
	QueryFiltersRequestWrite(w, Drawable)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: QueryFiltersOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return QueryFiltersCookie(seq)
}

func QueryFiltersUnchecked(conn *x.Conn, Drawable x.Drawable) QueryFiltersCookie {
	w := x.NewWriter()
	QueryFiltersRequestWrite(w, Drawable)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: QueryFiltersOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return QueryFiltersCookie(seq)
}

const SetPictureFilterOpcode = 30

func SetPictureFilterRequestWrite(w *x.Writer, Picture Picture, FilterLen uint16, Filter string, ValuesLen uint32, Values []Fixed) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Picture
	w.Write4b(uint32(Picture))
	n += 4

	// write wire field FilterLen
	w.Write2b(FilterLen)
	n += 2

	// write wire field Pad0
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Filter
	w.WriteString(Filter)
	n += len(Filter)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Values
	{
		_dataLen := int(ValuesLen)
		for i := 0; i < _dataLen; i++ {
			w.Write4b(uint32(Values[i]))
		}
		n += _dataLen * 4
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SetPictureFilter(conn *x.Conn, Picture Picture, FilterLen uint16, Filter string, ValuesLen uint32, Values []Fixed) x.VoidCookie {
	w := x.NewWriter()
	SetPictureFilterRequestWrite(w, Picture, FilterLen, Filter, ValuesLen, Values)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SetPictureFilterOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func SetPictureFilterChecked(conn *x.Conn, Picture Picture, FilterLen uint16, Filter string, ValuesLen uint32, Values []Fixed) x.VoidCookie {
	w := x.NewWriter()
	SetPictureFilterRequestWrite(w, Picture, FilterLen, Filter, ValuesLen, Values)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SetPictureFilterOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

type AnimCursorElt struct {
	Cursor x.Cursor
	Delay  uint32
}

type CAnimCursorElt struct {
	Cursor x.Cursor
	Delay  uint32
}

func AnimCursorEltRead(r *x.Reader, v *AnimCursorElt) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Cursor
	v.Cursor = x.Cursor(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Delay
	v.Delay = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}
	return
}

func AnimCursorEltReadN(r *x.Reader, dest []AnimCursorElt) (n int) {
	for i := 0; i < len(dest); i++ {
		n += AnimCursorEltRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func AnimCursorEltWrite(w *x.Writer, v *AnimCursorElt) (n int) {

	// write field Cursor
	w.Write4b(uint32(v.Cursor))
	n += 4

	// write field Delay
	w.Write4b(v.Delay)
	n += 4
	return
}
func AnimCursorEltWriteN(w *x.Writer, src []AnimCursorElt) (n int) {
	for i := 0; i < len(src); i++ {
		n += AnimCursorEltWrite(w, &src[i])
	}
	return
}

const CreateAnimCursorOpcode = 31

func CreateAnimCursorRequestWrite(w *x.Writer, Cid x.Cursor, CursorsLen uint32, Cursors []AnimCursorElt) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Cid
	w.Write4b(uint32(Cid))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Cursors
	n += AnimCursorEltWriteN(w, Cursors)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CreateAnimCursor(conn *x.Conn, Cid x.Cursor, CursorsLen uint32, Cursors []AnimCursorElt) x.VoidCookie {
	w := x.NewWriter()
	CreateAnimCursorRequestWrite(w, Cid, CursorsLen, Cursors)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateAnimCursorOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CreateAnimCursorChecked(conn *x.Conn, Cid x.Cursor, CursorsLen uint32, Cursors []AnimCursorElt) x.VoidCookie {
	w := x.NewWriter()
	CreateAnimCursorRequestWrite(w, Cid, CursorsLen, Cursors)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateAnimCursorOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

type SpanFix struct {
	L Fixed
	R Fixed
	Y Fixed
}

type CSpanFix struct {
	L Fixed
	R Fixed
	Y Fixed
}

func SpanFixRead(r *x.Reader, v *SpanFix) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field L
	v.L = Fixed(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field R
	v.R = Fixed(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Y
	v.Y = Fixed(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}
	return
}

func SpanFixReadN(r *x.Reader, dest []SpanFix) (n int) {
	for i := 0; i < len(dest); i++ {
		n += SpanFixRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func SpanFixWrite(w *x.Writer, v *SpanFix) (n int) {

	// write field L
	w.Write4b(uint32(v.L))
	n += 4

	// write field R
	w.Write4b(uint32(v.R))
	n += 4

	// write field Y
	w.Write4b(uint32(v.Y))
	n += 4
	return
}
func SpanFixWriteN(w *x.Writer, src []SpanFix) (n int) {
	for i := 0; i < len(src); i++ {
		n += SpanFixWrite(w, &src[i])
	}
	return
}

type Trap struct {
	Top SpanFix
	Bot SpanFix
}

type CTrap struct {
	Top SpanFix
	Bot SpanFix
}

func TrapRead(r *x.Reader, v *Trap) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Top
	n += SpanFixRead(r, &v.Top)
	if r.Err() != nil {
		return
	}

	// read field Bot
	n += SpanFixRead(r, &v.Bot)
	if r.Err() != nil {
		return
	}
	return
}

func TrapReadN(r *x.Reader, dest []Trap) (n int) {
	for i := 0; i < len(dest); i++ {
		n += TrapRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func TrapWrite(w *x.Writer, v *Trap) (n int) {

	// write field Top
	n += SpanFixWrite(w, &v.Top)

	// write field Bot
	n += SpanFixWrite(w, &v.Bot)
	return
}
func TrapWriteN(w *x.Writer, src []Trap) (n int) {
	for i := 0; i < len(src); i++ {
		n += TrapWrite(w, &src[i])
	}
	return
}

const AddTrapsOpcode = 32

func AddTrapsRequestWrite(w *x.Writer, Picture Picture, XOff int16, YOff int16, TrapsLen uint32, Traps []Trap) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Picture
	w.Write4b(uint32(Picture))
	n += 4

	// write wire field XOff
	w.Write2b(uint16(XOff))
	n += 2

	// write wire field YOff
	w.Write2b(uint16(YOff))
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Traps
	n += TrapWriteN(w, Traps)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func AddTraps(conn *x.Conn, Picture Picture, XOff int16, YOff int16, TrapsLen uint32, Traps []Trap) x.VoidCookie {
	w := x.NewWriter()
	AddTrapsRequestWrite(w, Picture, XOff, YOff, TrapsLen, Traps)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: AddTrapsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func AddTrapsChecked(conn *x.Conn, Picture Picture, XOff int16, YOff int16, TrapsLen uint32, Traps []Trap) x.VoidCookie {
	w := x.NewWriter()
	AddTrapsRequestWrite(w, Picture, XOff, YOff, TrapsLen, Traps)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: AddTrapsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const CreateSolidFillOpcode = 33

func CreateSolidFillRequestWrite(w *x.Writer, Picture Picture, Color Color) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Picture
	w.Write4b(uint32(Picture))
	n += 4

	// write wire field Color
	// TODO RequestWriteDefine container_FS
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CreateSolidFill(conn *x.Conn, Picture Picture, Color Color) x.VoidCookie {
	w := x.NewWriter()
	CreateSolidFillRequestWrite(w, Picture, Color)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateSolidFillOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CreateSolidFillChecked(conn *x.Conn, Picture Picture, Color Color) x.VoidCookie {
	w := x.NewWriter()
	CreateSolidFillRequestWrite(w, Picture, Color)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateSolidFillOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const CreateLinearGradientOpcode = 34

func CreateLinearGradientRequestWrite(w *x.Writer, Picture Picture, P1 PointFix, P2 PointFix, NumStops uint32, Stops []Fixed, Colors []Color) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Picture
	w.Write4b(uint32(Picture))
	n += 4

	// write wire field P1
	// TODO RequestWriteDefine container_FS

	// write wire field P2
	// TODO RequestWriteDefine container_FS

	// write wire field NumStops
	w.Write4b(NumStops)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Stops
	{
		_dataLen := int(NumStops)
		for i := 0; i < _dataLen; i++ {
			w.Write4b(uint32(Stops[i]))
		}
		n += _dataLen * 4
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Colors
	n += ColorWriteN(w, Colors)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CreateLinearGradient(conn *x.Conn, Picture Picture, P1 PointFix, P2 PointFix, NumStops uint32, Stops []Fixed, Colors []Color) x.VoidCookie {
	w := x.NewWriter()
	CreateLinearGradientRequestWrite(w, Picture, P1, P2, NumStops, Stops, Colors)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateLinearGradientOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CreateLinearGradientChecked(conn *x.Conn, Picture Picture, P1 PointFix, P2 PointFix, NumStops uint32, Stops []Fixed, Colors []Color) x.VoidCookie {
	w := x.NewWriter()
	CreateLinearGradientRequestWrite(w, Picture, P1, P2, NumStops, Stops, Colors)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateLinearGradientOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const CreateRadialGradientOpcode = 35

func CreateRadialGradientRequestWrite(w *x.Writer, Picture Picture, Inner PointFix, Outer PointFix, InnerRadius Fixed, OuterRadius Fixed, NumStops uint32, Stops []Fixed, Colors []Color) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Picture
	w.Write4b(uint32(Picture))
	n += 4

	// write wire field Inner
	// TODO RequestWriteDefine container_FS

	// write wire field Outer
	// TODO RequestWriteDefine container_FS

	// write wire field InnerRadius
	w.Write4b(uint32(InnerRadius))
	n += 4

	// write wire field OuterRadius
	w.Write4b(uint32(OuterRadius))
	n += 4

	// write wire field NumStops
	w.Write4b(NumStops)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Stops
	{
		_dataLen := int(NumStops)
		for i := 0; i < _dataLen; i++ {
			w.Write4b(uint32(Stops[i]))
		}
		n += _dataLen * 4
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Colors
	n += ColorWriteN(w, Colors)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CreateRadialGradient(conn *x.Conn, Picture Picture, Inner PointFix, Outer PointFix, InnerRadius Fixed, OuterRadius Fixed, NumStops uint32, Stops []Fixed, Colors []Color) x.VoidCookie {
	w := x.NewWriter()
	CreateRadialGradientRequestWrite(w, Picture, Inner, Outer, InnerRadius, OuterRadius, NumStops, Stops, Colors)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateRadialGradientOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CreateRadialGradientChecked(conn *x.Conn, Picture Picture, Inner PointFix, Outer PointFix, InnerRadius Fixed, OuterRadius Fixed, NumStops uint32, Stops []Fixed, Colors []Color) x.VoidCookie {
	w := x.NewWriter()
	CreateRadialGradientRequestWrite(w, Picture, Inner, Outer, InnerRadius, OuterRadius, NumStops, Stops, Colors)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateRadialGradientOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const CreateConicalGradientOpcode = 36

func CreateConicalGradientRequestWrite(w *x.Writer, Picture Picture, Center PointFix, Angle Fixed, NumStops uint32, Stops []Fixed, Colors []Color) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2

	// write wire field Picture
	w.Write4b(uint32(Picture))
	n += 4

	// write wire field Center
	// TODO RequestWriteDefine container_FS

	// write wire field Angle
	w.Write4b(uint32(Angle))
	n += 4

	// write wire field NumStops
	w.Write4b(NumStops)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Stops
	{
		_dataLen := int(NumStops)
		for i := 0; i < _dataLen; i++ {
			w.Write4b(uint32(Stops[i]))
		}
		n += _dataLen * 4
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Colors
	n += ColorWriteN(w, Colors)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CreateConicalGradient(conn *x.Conn, Picture Picture, Center PointFix, Angle Fixed, NumStops uint32, Stops []Fixed, Colors []Color) x.VoidCookie {
	w := x.NewWriter()
	CreateConicalGradientRequestWrite(w, Picture, Center, Angle, NumStops, Stops, Colors)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateConicalGradientOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CreateConicalGradientChecked(conn *x.Conn, Picture Picture, Center PointFix, Angle Fixed, NumStops uint32, Stops []Fixed, Colors []Color) x.VoidCookie {
	w := x.NewWriter()
	CreateConicalGradientRequestWrite(w, Picture, Center, Angle, NumStops, Stops, Colors)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateConicalGradientOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}
