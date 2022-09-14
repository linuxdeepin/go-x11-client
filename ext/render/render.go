// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package render

import (
	"math"

	"github.com/linuxdeepin/go-x11-client"
)

const fixedOne = 1 << 16

func ToFixed(v float64) Fixed {
	return Fixed(v * fixedOne)
}

func (f Fixed) ToFloat64() float64 {
	return float64(f) / fixedOne
}

// #WREQ
func encodeQueryVersion(majorVersion, minorVersion uint32) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(majorVersion).
		Write4b(minorVersion).
		End()
	return
}

type QueryVersionReply struct {
	MajorVersion, MinorVersion uint32
}

func readQueryVersionReply(r *x.Reader, v *QueryVersionReply) error {
	if !r.RemainAtLeast4b(4) {
		return x.ErrDataLenShort
	}
	r.ReadPad(8)

	v.MajorVersion = r.Read4b()

	v.MinorVersion = r.Read4b() // 4
	return nil
}

// #WREQ
func encodeQueryPictFormats() (b x.RequestBody) {
	return
}

type QueryPictFormatsReply struct {
	Formats    []PictFormatInfo
	Screens    []PictScreen
	NumDepths  uint32
	NumVisuals uint32
	SubPixels  []uint32
}

// size: 7 * 4b
type PictFormatInfo struct {
	Id       PictFormat
	Type     uint8
	Depth    uint8
	Direct   DirectFormat
	Colormap x.Colormap
}

func readPictFormatInfo(r *x.Reader, v *PictFormatInfo) {
	v.Id = PictFormat(r.Read4b())

	v.Type = r.Read1b()
	v.Depth = r.Read1b()
	r.ReadPad(2) // 2

	readDirectFormat(r, &v.Direct)

	v.Colormap = x.Colormap(r.Read4b()) // 7
}

// size: 4 * 4b
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

func readDirectFormat(r *x.Reader, v *DirectFormat) {
	v.RedShift = r.Read2b()
	v.RedMask = r.Read2b()

	v.GreenShift = r.Read2b()
	v.GreenMask = r.Read2b()

	v.BlueShift = r.Read2b()
	v.BlueMask = r.Read2b()

	v.AlphaShift = r.Read2b()
	v.AlphaMask = r.Read2b()
}

type PictScreen struct {
	Fallback PictFormat
	Depths   []PictDepth
}

func readPictScreen(r *x.Reader, v *PictScreen) error {
	if !r.RemainAtLeast4b(2) {
		return x.ErrDataLenShort
	}
	depthsLen := int(r.Read4b())

	v.Fallback = PictFormat(r.Read4b()) // 2

	if depthsLen > 0 {
		v.Depths = make([]PictDepth, depthsLen)
		for i := 0; i < depthsLen; i++ {
			err := readPictDepth(r, &v.Depths[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type PictDepth struct {
	Depth   uint8
	Visuals []PictVisual
}

func readPictDepth(r *x.Reader, v *PictDepth) error {
	v.Depth = r.Read1b()
	r.ReadPad(1)
	visualsLen := int(r.Read2b())

	r.ReadPad(4) // 2

	if visualsLen > 0 {
		if !r.RemainAtLeast4b(2 * visualsLen) {
			return x.ErrDataLenShort
		}
		v.Visuals = make([]PictVisual, visualsLen)
		for i := 0; i < visualsLen; i++ {
			v.Visuals[i] = readPictVisual(r)
		}
	}

	return nil
}

// size: 2 * 4b
type PictVisual struct {
	Visual x.VisualID
	Format PictFormat
}

func readPictVisual(r *x.Reader) PictVisual {
	var v PictVisual
	v.Visual = x.VisualID(r.Read4b())

	v.Format = PictFormat(r.Read4b())
	return v
}

func readQueryPictFormatsReply(r *x.Reader, v *QueryPictFormatsReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	r.ReadPad(8)

	formatsLen := int(r.Read4b())

	screensLen := int(r.Read4b())

	v.NumDepths = r.Read4b() // 5

	v.NumVisuals = r.Read4b()

	subPixelsLen := int(r.Read4b())

	r.ReadPad(4) // 8

	if formatsLen > 0 {
		if !r.RemainAtLeast4b(7 * formatsLen) {
			return x.ErrDataLenShort
		}
		v.Formats = make([]PictFormatInfo, formatsLen)
		for i := 0; i < formatsLen; i++ {
			readPictFormatInfo(r, &v.Formats[i])
		}
	}

	if screensLen > 0 {
		v.Screens = make([]PictScreen, screensLen)
		for i := 0; i < screensLen; i++ {
			err := readPictScreen(r, &v.Screens[i])
			if err != nil {
				return err
			}
		}
	}

	if subPixelsLen > 0 {
		if !r.RemainAtLeast4b(subPixelsLen) {
			return x.ErrDataLenShort
		}
		v.SubPixels = make([]uint32, subPixelsLen)
		for i := 0; i < subPixelsLen; i++ {
			v.SubPixels[i] = r.Read4b()
		}
	}

	return nil
}

// #WREQ
func encodeQueryPictIndexValues(format PictFormat) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(format)).
		End()
	return
}

// size: 3 * 4b
type IndexValue struct {
	Pixel uint32
	Red   uint16
	Green uint16
	Blue  uint16
	Alpha uint16
}

func readIndexValue(r *x.Reader, v *IndexValue) {
	v.Pixel = r.Read4b()

	v.Red = r.Read2b()
	v.Green = r.Read2b()

	v.Blue = r.Read2b()
	v.Alpha = r.Read2b()
}

type QueryPictIndexValuesReply struct {
	Values []IndexValue
}

func readQueryPictIndexValuesReply(r *x.Reader, v *QueryPictIndexValuesReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	r.ReadPad(8)

	valuesLen := int(r.Read4b())

	r.ReadPad(20) // 8

	if valuesLen > 0 {
		if !r.RemainAtLeast4b(3 * valuesLen) {
			return x.ErrDataLenShort
		}
		v.Values = make([]IndexValue, valuesLen)
		for i := 0; i < valuesLen; i++ {
			readIndexValue(r, &v.Values[i])
		}
	}

	return nil
}

// #WREQ
func encodeQueryFilters(drawable x.Drawable) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(drawable)).
		End()
	return
}

type QueryFiltersReply struct {
	Aliases []uint16
	Filters []string
}

func readQueryFiltersReply(r *x.Reader, v *QueryFiltersReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	r.ReadPad(8)

	aliasesLen := int(r.Read4b())

	filtersLen := int(r.Read4b())

	r.ReadPad(16) // 8

	if aliasesLen > 0 {
		pad := x.Pad(2 * aliasesLen)
		if !r.RemainAtLeast(2*aliasesLen + pad) {
			return x.ErrDataLenShort
		}
		v.Aliases = make([]uint16, aliasesLen)
		for i := 0; i < aliasesLen; i++ {
			v.Aliases[i] = r.Read2b()
		}
		r.ReadPad(pad)
	}

	if filtersLen > 0 {
		v.Filters = make([]string, filtersLen)
		var err error
		for i := 0; i < filtersLen; i++ {
			v.Filters[i], err = x.ReadStr(r)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// #WREQ
func encodeCreatePicture(pid Picture, drawable x.Drawable, format PictFormat,
	valueMask uint32, valueList []uint32) (b x.RequestBody) {
	b0 := b.AddBlock(4 + len(valueList)).
		Write4b(uint32(pid)).
		Write4b(uint32(drawable)).
		Write4b(uint32(format)).
		Write4b(valueMask)
	for _, value := range valueList {
		b0.Write4b(value)
	}
	b0.End()
	return
}

// #WREQ
func encodeChangePicture(picture Picture, valueMask uint32, valueList []uint32) (b x.RequestBody) {
	b0 := b.AddBlock(2 + len(valueList)).
		Write4b(uint32(picture)).
		Write4b(valueMask)
	for _, value := range valueList {
		b0.Write4b(value)
	}
	b0.End()
	return
}

// #WREQ
func encodeSetPictureClipRectangles(picture Picture, clipXOrigin,
	clipYOrigin int16, rectangles []x.Rectangle) (b x.RequestBody) {
	b0 := b.AddBlock(2 + 2*len(rectangles)).
		Write4b(uint32(picture)).
		Write2b(uint16(clipXOrigin)).
		Write2b(uint16(clipYOrigin))
	for _, rect := range rectangles {
		x.WriteRectangle(b0, rect)
	}
	b0.End()
	return
}

// size: 9 * 4b
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

func WriteTransform(b *x.FixedSizeBuf, v *Transform) {
	b.Write4b(uint32(v.Matrix11))
	b.Write4b(uint32(v.Matrix12))
	b.Write4b(uint32(v.Matrix13))

	b.Write4b(uint32(v.Matrix21))
	b.Write4b(uint32(v.Matrix22))
	b.Write4b(uint32(v.Matrix23))

	b.Write4b(uint32(v.Matrix31))
	b.Write4b(uint32(v.Matrix32))
	b.Write4b(uint32(v.Matrix33))
}

func ReadTransform(r *x.Reader, v *Transform) {
	v.Matrix11 = Fixed(r.Read4b())
	v.Matrix12 = Fixed(r.Read4b())
	v.Matrix13 = Fixed(r.Read4b())

	v.Matrix21 = Fixed(r.Read4b())
	v.Matrix22 = Fixed(r.Read4b())
	v.Matrix23 = Fixed(r.Read4b())

	v.Matrix31 = Fixed(r.Read4b())
	v.Matrix32 = Fixed(r.Read4b())
	v.Matrix33 = Fixed(r.Read4b())
}

// #WREQ
func encodeSetPictureTransform(picture Picture, transform *Transform) (b x.RequestBody) {
	b0 := b.AddBlock(10).
		Write4b(uint32(picture))
	WriteTransform(b0, transform)
	b0.End()
	return
}

// #WREQ
func encodeSetPictureFilter(picture Picture, filter string, values []Fixed) (b x.RequestBody) {
	filter = x.TruncateStr(filter, math.MaxUint16)
	filterLen := len(filter)
	b0 := b.AddBlock(2 + x.SizeIn4bWithPad(filterLen) + len(values)).
		Write4b(uint32(picture)).
		Write2b(uint16(filterLen)).
		WritePad(2).
		WriteString(filter).
		WritePad(x.Pad(filterLen))

	for _, value := range values {
		b0.Write4b(uint32(value))
	}
	b0.End()
	return
}

// #WREQ
func encodeFreePicture(picture Picture) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(picture)).
		End()
	return
}

// #WREQ
func encodeComposite(op uint8, src, mask, dst Picture, srcX, srcY,
	maskX, maskY, dstX, dstY int16, width, height uint16) (b x.RequestBody) {
	b.AddBlock(8).
		Write1b(op).
		WritePad(3).
		Write4b(uint32(src)).
		Write4b(uint32(mask)).
		Write4b(uint32(dst)).
		Write2b(uint16(srcX)).
		Write2b(uint16(srcY)).
		Write2b(uint16(maskX)).
		Write2b(uint16(maskY)).
		Write2b(uint16(dstX)).
		Write2b(uint16(dstY)).
		Write2b(width).
		Write2b(height).
		End()
	return
}

// size: 2 * 4b
type Color struct {
	Red, Green, Blue, Alpha uint16
}

func writeColor(b *x.FixedSizeBuf, v Color) {
	b.Write2b(v.Red)
	b.Write2b(v.Green)
	b.Write2b(v.Blue)
	b.Write2b(v.Alpha)
}

// #WREQ
func encodeFillRectangles(op uint8, dst Picture, color Color,
	rects []x.Rectangle) (b x.RequestBody) {
	b0 := b.AddBlock(4 + 2*len(rects)).
		Write1b(op).
		WritePad(3).
		Write4b(uint32(dst))

	writeColor(b0, color)
	for _, rect := range rects {
		x.WriteRectangle(b0, rect)
	}
	b0.End()
	return
}

// size: 10 * 4b
type Trapezoid struct {
	Top    Fixed
	Bottom Fixed
	Left   LineFixed
	Right  LineFixed
}

func writeTrapezoid(b *x.FixedSizeBuf, v *Trapezoid) {
	b.Write4b(uint32(v.Top))
	b.Write4b(uint32(v.Bottom))
	writeLineFixed(b, v.Left)
	writeLineFixed(b, v.Right)
}

// size: 4 * 4b
type LineFixed struct {
	P1 PointFixed
	P2 PointFixed
}

func writeLineFixed(b *x.FixedSizeBuf, v LineFixed) {
	writePointFixed(b, v.P1)
	writePointFixed(b, v.P2)
}

// size: 2 * 4b
type PointFixed struct {
	X Fixed
	Y Fixed
}

func writePointFixed(b *x.FixedSizeBuf, v PointFixed) {
	b.Write4b(uint32(v.X))
	b.Write4b(uint32(v.Y))
}

// #WREQ
func encodeTrapezoids(op uint8, src, dst Picture, maskFormat PictFormat,
	srcX, srcY int16, traps []Trapezoid) (b x.RequestBody) {

	b0 := b.AddBlock(5 + 10*len(traps)).
		Write1b(op).
		WritePad(3).
		Write4b(uint32(src)).
		Write4b(uint32(dst)).
		Write4b(uint32(maskFormat)).
		Write2b(uint16(srcX)).
		Write2b(uint16(srcY))

	for _, trap := range traps {
		writeTrapezoid(b0, &trap)
	}
	b0.End()
	return
}

// size: 6 * 4b
type Triangle struct {
	P1 PointFixed
	P2 PointFixed
	P3 PointFixed
}

func writeTriangle(b *x.FixedSizeBuf, v Triangle) {
	writePointFixed(b, v.P1)
	writePointFixed(b, v.P2)
	writePointFixed(b, v.P3)
}

// #WREQ
func encodeTriangles(op uint8, src, dst Picture, maskFormat PictFormat,
	srcX, srcY int16, triangles []Triangle) (b x.RequestBody) {

	b0 := b.AddBlock(5 + 6*len(triangles)).
		Write1b(op).
		WritePad(3).
		Write4b(uint32(src)).
		Write4b(uint32(dst)).
		Write4b(uint32(maskFormat)).
		Write2b(uint16(srcX)).
		Write2b(uint16(srcY))

	for _, triangle := range triangles {
		writeTriangle(b0, triangle)
	}
	b0.End()
	return
}

// #WREQ
func encodeTriStrip(op uint8, src, dst Picture, maskFormat PictFormat,
	srcX, srcY int16, points []PointFixed) (b x.RequestBody) {

	b0 := b.AddBlock(5 + len(points)*2).
		Write1b(op).
		WritePad(3).
		Write4b(uint32(src)).
		Write4b(uint32(dst)).
		Write4b(uint32(maskFormat)).
		Write2b(uint16(srcX)).
		Write2b(uint16(srcY))

	for _, p := range points {
		writePointFixed(b0, p)
	}
	b0.End()
	return
}

// #WREQ
func encodeTriFan(op uint8, src, dst Picture, maskFormat PictFormat,
	srcX, srcY int16, points []PointFixed) (b x.RequestBody) {

	b0 := b.AddBlock(5 + len(points)*2).
		Write1b(op).
		WritePad(3).
		Write4b(uint32(src)).
		Write4b(uint32(dst)).
		Write4b(uint32(maskFormat)).
		Write2b(uint16(srcX)).
		Write2b(uint16(srcY))

	for _, p := range points {
		writePointFixed(b0, p)
	}
	b0.End()
	return
}

// #WREQ
func encodeCreateGlyphSet(gsId GlyphSet, format PictFormat) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(gsId)).
		Write4b(uint32(format)).
		End()
	return
}

// #WREQ
func encodeReferenceGlyphSet(gsId, existing GlyphSet) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(gsId)).
		Write4b(uint32(existing)).
		End()
	return
}

// #WREQ
func encodeFreeGlyphSet(glyphSet GlyphSet) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(glyphSet)).
		End()
	return
}

// size: 3 * 4b
type GlyphInfo struct {
	Width  uint16
	Height uint16
	X      int16
	Y      int16
	XOff   int16
	YOff   int16
}

func writeGlyphInfo(b *x.FixedSizeBuf, v GlyphInfo) {
	b.Write2b(v.Width)
	b.Write2b(v.Height)
	b.Write2b(uint16(v.X))
	b.Write2b(uint16(v.Y))
	b.Write2b(uint16(v.XOff))
	b.Write2b(uint16(v.YOff))
}

// #WREQ
func encodeAddGlyphs(glyphSet GlyphSet, glyphIds []uint32, glyphs []GlyphInfo,
	data []byte) (b x.RequestBody) {

	b0 := b.AddBlock(2 + len(glyphs)*3).
		Write4b(uint32(glyphSet)).
		Write4b(uint32(len(glyphIds)))

	for _, gi := range glyphs {
		writeGlyphInfo(b0, gi)
	}
	b0.End()

	b.AddBytes(data)
	return
}

// #WREQ
func encodeFreeGlyphs(glyphSet GlyphSet, glyphs []Glyph) (b x.RequestBody) {
	b0 := b.AddBlock(1 + len(glyphs)).
		Write4b(uint32(glyphSet))
	for _, g := range glyphs {
		b0.Write4b(uint32(g))
	}
	b0.End()
	return
}

func encodeCompositeGlyphsN(op uint8, src, dst Picture, maskFormat PictFormat,
	glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) (b x.RequestBody) {

	b.AddBlock(6).
		Write1b(op).
		WritePad(3).
		Write4b(uint32(src)).
		Write4b(uint32(dst)).
		Write4b(uint32(maskFormat)).
		Write4b(uint32(glyphSet)).
		Write2b(uint16(srcX)).
		Write2b(uint16(srcY)).
		End()

	b.AddBytes(glyphCmds)
	return
}

// #WREQ
func encodeCompositeGlyphs8(op uint8, src, dst Picture, maskFormat PictFormat,
	glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) (b x.RequestBody) {
	return encodeCompositeGlyphsN(op, src, dst, maskFormat, glyphSet, srcX, srcY, glyphCmds)
}

// #WREQ
func encodeCompositeGlyphs16(op uint8, src, dst Picture, maskFormat PictFormat,
	glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) (b x.RequestBody) {
	return encodeCompositeGlyphsN(op, src, dst, maskFormat, glyphSet, srcX, srcY, glyphCmds)
}

// #WREQ
func encodeCompositeGlyphs32(op uint8, src, dst Picture, maskFormat PictFormat,
	glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) (b x.RequestBody) {
	return encodeCompositeGlyphsN(op, src, dst, maskFormat, glyphSet, srcX, srcY, glyphCmds)
}

// #WREQ
func encodeCreateCursor(cid x.Cursor, source Picture, X, y uint16) (b x.RequestBody) {
	b.AddBlock(3).
		Write4b(uint32(cid)).
		Write4b(uint32(source)).
		Write2b(X).
		Write2b(y).
		End()
	return
}

// size: 2 * 4b
type AnimCursorElt struct {
	Cursor x.Cursor
	Delay  uint32
}

func writeAnimCursorElt(b *x.FixedSizeBuf, v AnimCursorElt) {
	b.Write4b(uint32(v.Cursor))
	b.Write4b(v.Delay)
}

// #WREQ
func encodeCreateAnimCursor(cid x.Cursor, cursors []AnimCursorElt) (b x.RequestBody) {
	b0 := b.AddBlock(1 + len(cursors)*2).
		Write4b(uint32(cid))
	for _, c := range cursors {
		writeAnimCursorElt(b0, c)
	}
	b0.End()
	return
}

// size: 3 * 4b
type SpanFixed struct {
	L, R, Y Fixed
}

// size: 6 * 4b
type Trap struct {
	Top SpanFixed
	Bot SpanFixed
}

func writeSpanFixed(b *x.FixedSizeBuf, v SpanFixed) {
	b.Write4b(uint32(v.L))
	b.Write4b(uint32(v.R))
	b.Write4b(uint32(v.Y))
}

func writeTrap(b *x.FixedSizeBuf, v *Trap) {
	writeSpanFixed(b, v.Top)
	writeSpanFixed(b, v.Bot)
}

// #WREQ
func encodeAddTraps(picture Picture, xOff, yOff int16, traps []Trap) (b x.RequestBody) {
	b0 := b.AddBlock(2 + len(traps)*6).
		Write4b(uint32(picture)).
		Write2b(uint16(xOff)).
		Write2b(uint16(yOff))

	for _, trap := range traps {
		writeTrap(b0, &trap)
	}
	b0.End()
	return
}

// #WREQ
func encodeCreateSolidFill(pid Picture, color Color) (b x.RequestBody) {
	b0 := b.AddBlock(3).
		Write4b(uint32(pid))
	writeColor(b0, color)
	b0.End()
	return
}

// #WREQ
func encodeCreateLinearGradient(picture Picture, p1, p2 PointFixed,
	stops []Fixed, colors []Color) (b x.RequestBody) {
	b0 := b.AddBlock(6 + len(stops) + 2*len(colors)).
		Write4b(uint32(picture))
	writePointFixed(b0, p1)
	writePointFixed(b0, p2)

	b0.Write4b(uint32(len(stops)))

	for _, s := range stops {
		b0.Write4b(uint32(s))
	}
	for _, c := range colors {
		writeColor(b0, c)
	}
	b0.End()
	return
}

// #WREQ
func encodeCreateRadialGradient(picture Picture, inner, outer PointFixed,
	innerRadius, outerRadius Fixed, stops []Fixed, colors []Color) (b x.RequestBody) {
	b0 := b.AddBlock(8 + len(stops) + 2*len(colors)).
		Write4b(uint32(picture))
	writePointFixed(b0, inner)
	writePointFixed(b0, outer)
	b0.Write4b(uint32(innerRadius)).
		Write4b(uint32(outerRadius)).
		Write4b(uint32(len(stops)))

	for _, s := range stops {
		b0.Write4b(uint32(s))
	}
	for _, c := range colors {
		writeColor(b0, c)
	}
	b0.End()
	return
}

// #WREQ
func encodeCreateConicalGradient(picture Picture, center PointFixed,
	angle Fixed, stops []Fixed, colors []Color) (b x.RequestBody) {

	b0 := b.AddBlock(4 + len(stops) + 2*len(colors)).
		Write4b(uint32(picture))

	writePointFixed(b0, center)
	b0.Write4b(uint32(angle))

	for _, s := range stops {
		b0.Write4b(uint32(s))
	}
	for _, c := range colors {
		writeColor(b0, c)
	}
	b0.End()
	return
}
