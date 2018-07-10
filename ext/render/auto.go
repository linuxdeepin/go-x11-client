package render

import x "github.com/linuxdeepin/go-x11-client"

// _ns.ext_name: Render
const MajorVersion = 0
const MinorVersion = 11

var _ext *x.Extension

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

const PictFormatErrorCode = 0
const PictureErrorCode = 1
const PictOpErrorCode = 2
const GlyphSetErrorCode = 3
const GlyphErrorCode = 4
const QueryVersionOpcode = 0

type QueryVersionCookie uint64

const QueryPictFormatsOpcode = 1

type QueryPictFormatsCookie uint64

const QueryPictIndexValuesOpcode = 2

type QueryPictIndexValuesCookie uint64

const CreatePictureOpcode = 4
const ChangePictureOpcode = 5
const SetPictureClipRectanglesOpcode = 6
const FreePictureOpcode = 7
const CompositeOpcode = 8
const TrapezoidsOpcode = 10
const TrianglesOpcode = 11
const TriStripOpcode = 12
const TriFanOpcode = 13
const CreateGlyphSetOpcode = 17
const ReferenceGlyphSetOpcode = 18
const FreeGlyphSetOpcode = 19
const AddGlyphsOpcode = 20
const FreeGlyphsOpcode = 22
const CompositeGlyphs8Opcode = 23
const CompositeGlyphs16Opcode = 24
const CompositeGlyphs32Opcode = 25
const FillRectanglesOpcode = 26
const CreateCursorOpcode = 27
const SetPictureTransformOpcode = 28
const QueryFiltersOpcode = 29

type QueryFiltersCookie uint64

const SetPictureFilterOpcode = 30
const CreateAnimCursorOpcode = 31
const AddTrapsOpcode = 32
const CreateSolidFillOpcode = 33
const CreateLinearGradientOpcode = 34
const CreateRadialGradientOpcode = 35
const CreateConicalGradientOpcode = 36

var readErrorFuncMap = make(map[uint8]x.ReadErrorFunc, 5)

func init() {
	readErrorFuncMap[PictFormatErrorCode] = readPictFormatError
	readErrorFuncMap[PictureErrorCode] = readPictureError
	readErrorFuncMap[PictOpErrorCode] = readPictOpError
	readErrorFuncMap[GlyphSetErrorCode] = readGlyphSetError
	readErrorFuncMap[GlyphErrorCode] = readGlyphError
	_ext = x.NewExtension("RENDER", 4, readErrorFuncMap)
}
