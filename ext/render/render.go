package render

import "github.com/linuxdeepin/go-x11-client"

// #WREQ
func writeQueryVersion(w *x.Writer, majorVersion, minorVersion uint32) {
	w.WritePad(4)
	w.Write4b(majorVersion)
	w.Write4b(minorVersion)
}

type QueryVersionReply struct {
	MajorVersion, MinorVersion uint32
}

func readQueryVersionReply(r *x.Reader, v *QueryVersionReply) error {
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

	v.MajorVersion = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MinorVersion = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}
	return nil
}

// #WREQ
func writeQueryPictFormats(w *x.Writer) {
	w.WritePad(4)
}

type QueryPictFormatsReply struct {
	Formats    []PictFormatInfo
	Screens    []PictScreen
	NumDepths  uint32
	NumVisuals uint32
	SubPixels  []uint32
}

type PictFormatInfo struct {
	Id       PictFormat
	Type     uint8
	Depth    uint8
	Direct   DirectFormat
	Colormap x.Colormap
}

func readPictFormatInfo(r *x.Reader, v *PictFormatInfo) error {
	v.Id = PictFormat(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Type = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Depth = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(2)
	if r.Err() != nil {
		return r.Err()
	}

	err := readDirectFormat(r, &v.Direct)
	if err != nil {
		return err
	}

	v.Colormap = x.Colormap(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

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

func readDirectFormat(r *x.Reader, v *DirectFormat) error {
	v.RedShift = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.RedMask = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.GreenShift = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.GreenMask = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.BlueShift = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.BlueMask = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.AlphaShift = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.AlphaMask = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}
	return nil
}

type PictScreen struct {
	Fallback PictFormat
	Depths   []PictDepth
}

func readPictScreen(r *x.Reader, v *PictScreen) error {
	depthsLen := int(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Fallback = PictFormat(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

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
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	visualsLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(4)
	if r.Err() != nil {
		return r.Err()
	}

	if visualsLen > 0 {
		v.Visuals = make([]PictVisual, visualsLen)
		var err error
		for i := 0; i < visualsLen; i++ {
			v.Visuals[i], err = readPictVisual(r)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type PictVisual struct {
	Visual x.VisualID
	Format PictFormat
}

func readPictVisual(r *x.Reader) (PictVisual, error) {
	var v PictVisual
	v.Visual = x.VisualID(r.Read4b())
	if r.Err() != nil {
		return PictVisual{}, r.Err()
	}

	v.Format = PictFormat(r.Read4b())
	if r.Err() != nil {
		return PictVisual{}, r.Err()
	}

	return v, nil
}

func readQueryPictFormatsReply(r *x.Reader, v *QueryPictFormatsReply) error {
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

	formatsLen := int(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	screensLen := int(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.NumDepths = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NumVisuals = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	subPixelsLen := int(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(4)
	if r.Err() != nil {
		return r.Err()
	}

	if formatsLen > 0 {
		v.Formats = make([]PictFormatInfo, formatsLen)
		for i := 0; i < formatsLen; i++ {
			err := readPictFormatInfo(r, &v.Formats[i])
			if err != nil {
				return err
			}
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
		v.SubPixels = make([]uint32, subPixelsLen)
		for i := 0; i < subPixelsLen; i++ {
			v.SubPixels[i] = r.Read4b()
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	return nil
}

// #WREQ
func writeQueryPictIndexValues(w *x.Writer, format PictFormat) {
	w.WritePad(4)
	w.Write4b(uint32(format))
}

type IndexValue struct {
	Pixel uint32
	Red   uint16
	Green uint16
	Blue  uint16
	Alpha uint16
}

func readIndexValue(r *x.Reader, v *IndexValue) error {
	v.Pixel = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Red = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Green = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Blue = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Alpha = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type QueryPictIndexValuesReply struct {
	Values []IndexValue
}

func readQueryPictIndexValuesReply(r *x.Reader, v *QueryPictIndexValuesReply) error {
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

	valuesLen := int(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(20)
	if r.Err() != nil {
		return r.Err()
	}

	if valuesLen > 0 {
		v.Values = make([]IndexValue, valuesLen)
		for i := 0; i < valuesLen; i++ {
			err := readIndexValue(r, &v.Values[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// #WREQ
func writeQueryFilters(w *x.Writer, drawable x.Drawable) {
	w.WritePad(4)
	w.Write4b(uint32(drawable))
}

type QueryFiltersReply struct {
	Aliases []uint16
	Filters []string
}

func readQueryFiltersReply(r *x.Reader, v *QueryFiltersReply) error {
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

	aliasesLen := int(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	filtersLen := int(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(16)
	if r.Err() != nil {
		return r.Err()
	}

	if aliasesLen > 0 {
		v.Aliases = make([]uint16, aliasesLen)
		for i := 0; i < aliasesLen; i++ {
			v.Aliases[i] = r.Read2b()
			if r.Err() != nil {
				return r.Err()
			}
		}
		r.ReadPad(x.Pad(aliasesLen * 2))
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
func writeCreatePicture(w *x.Writer, pid Picture, drawable x.Drawable, format PictFormat,
	valueMask uint32, valueList []uint32) {
	w.WritePad(4)
	w.Write4b(uint32(pid))
	w.Write4b(uint32(drawable))
	w.Write4b(uint32(format))
	w.Write4b(valueMask)
	for _, value := range valueList {
		w.Write4b(value)
	}
}

// #WREQ
func writeChangePicture(w *x.Writer, picture Picture, valueMask uint32, valueList []uint32) {
	w.WritePad(4)
	w.Write4b(uint32(picture))
	w.Write4b(valueMask)
	for _, value := range valueList {
		w.Write4b(value)
	}
}

// #WREQ
func writeSetPictureClipRectangles(w *x.Writer, picture Picture, clipXOrigin,
	clipYOrigin int16, rectangles []x.Rectangle) {
	w.WritePad(4)
	w.Write4b(uint32(picture))
	w.Write2b(uint16(clipXOrigin))
	w.Write2b(uint16(clipYOrigin))
	for _, rect := range rectangles {
		x.WriteRectangle(w, rect)
	}
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

func writeTransform(w *x.Writer, v *Transform) {
	w.Write4b(uint32(v.Matrix11))
	w.Write4b(uint32(v.Matrix12))
	w.Write4b(uint32(v.Matrix13))

	w.Write4b(uint32(v.Matrix21))
	w.Write4b(uint32(v.Matrix22))
	w.Write4b(uint32(v.Matrix23))

	w.Write4b(uint32(v.Matrix31))
	w.Write4b(uint32(v.Matrix32))
	w.Write4b(uint32(v.Matrix33))
}

// #WREQ
func writeSetPictureTransform(w *x.Writer, picture Picture, transform *Transform) {
	w.WritePad(4)
	w.Write4b(uint32(picture))
	writeTransform(w, transform)
}

// #WREQ
func writeSetPictureFilter(w *x.Writer, picture Picture, filter string, values []Fixed) {
	w.WritePad(4)
	w.Write4b(uint32(picture))

	filterLen := uint16(len(filter))
	w.Write2b(filterLen)
	w.WritePad(2)

	w.WriteString(filter)
	w.WritePad(x.Pad(len(filter)))

	for _, value := range values {
		w.Write4b(uint32(value))
	}
}

// #WREQ
func writeFreePicture(w *x.Writer, picture Picture) {
	w.WritePad(4)
	w.Write4b(uint32(picture))
}

// #WREQ
func writeComposite(w *x.Writer, op uint8, src, mask, dst Picture, srcX, srcY,
	maskX, maskY, dstX, dstY int16, width, height uint16) {
	w.WritePad(4)
	w.Write1b(op)
	w.WritePad(3)
	w.Write4b(uint32(src))
	w.Write4b(uint32(mask))
	w.Write4b(uint32(dst))
	w.Write2b(uint16(srcX))
	w.Write2b(uint16(srcY))
	w.Write2b(uint16(maskX))
	w.Write2b(uint16(maskY))
	w.Write2b(uint16(dstX))
	w.Write2b(uint16(dstY))
	w.Write2b(width)
	w.Write2b(height)
}

type Color struct {
	Red, Green, Blue, Alpha uint16
}

func writeColor(w *x.Writer, v Color) {
	w.Write2b(v.Red)
	w.Write2b(v.Green)
	w.Write2b(v.Blue)
	w.Write2b(v.Alpha)
}

// #WREQ
func writeFillRectangles(w *x.Writer, op uint8, dst Picture, color Color, rects []x.Rectangle) {
	w.WritePad(4)
	w.Write1b(op)
	w.WritePad(3)
	w.Write4b(uint32(dst))
	writeColor(w, color)
	for _, rect := range rects {
		x.WriteRectangle(w, rect)
	}
}

type Trapezoid struct {
	Top    Fixed
	Bottom Fixed
	Left   LineFixed
	Right  LineFixed
}

func writeTrapezoid(w *x.Writer, v *Trapezoid) {
	w.Write4b(uint32(v.Top))
	w.Write4b(uint32(v.Bottom))
	writeLineFixed(w, v.Left)
	writeLineFixed(w, v.Right)
}

type LineFixed struct {
	P1 PointFixed
	P2 PointFixed
}

func writeLineFixed(w *x.Writer, v LineFixed) {
	writePointFixed(w, v.P1)
	writePointFixed(w, v.P2)
}

type PointFixed struct {
	X Fixed
	Y Fixed
}

func writePointFixed(w *x.Writer, v PointFixed) {
	w.Write4b(uint32(v.X))
	w.Write4b(uint32(v.Y))
}

// #WREQ
func writeTrapezoids(w *x.Writer, op uint8, src, dst Picture, maskFormat PictFormat,
	srcX, srcY int16, traps []Trapezoid) {
	w.WritePad(4)
	w.Write1b(op)
	w.WritePad(3)

	w.Write4b(uint32(src))
	w.Write4b(uint32(dst))
	w.Write4b(uint32(maskFormat))

	w.Write2b(uint16(srcX))
	w.Write2b(uint16(srcY))

	for _, trap := range traps {
		writeTrapezoid(w, &trap)
	}
}

type Triangle struct {
	P1 PointFixed
	P2 PointFixed
	P3 PointFixed
}

func writeTriangle(w *x.Writer, v Triangle) {
	writePointFixed(w, v.P1)
	writePointFixed(w, v.P2)
	writePointFixed(w, v.P3)
}

// #WREQ
func writeTriangles(w *x.Writer, op uint8, src, dst Picture, maskFormat PictFormat,
	srcX, srcY int16, triangles []Triangle) {

	w.WritePad(4)
	w.Write1b(op)
	w.WritePad(3)

	w.Write4b(uint32(src))
	w.Write4b(uint32(dst))
	w.Write4b(uint32(maskFormat))

	w.Write2b(uint16(srcX))
	w.Write2b(uint16(srcY))

	for _, triangle := range triangles {
		writeTriangle(w, triangle)
	}
}

// #WREQ
func writeTriStrip(w *x.Writer, op uint8, src, dst Picture, maskFormat PictFormat,
	srcX, srcY int16, points []PointFixed) {

	w.WritePad(4)
	w.Write1b(op)
	w.WritePad(3)

	w.Write4b(uint32(src))
	w.Write4b(uint32(dst))
	w.Write4b(uint32(maskFormat))

	w.Write2b(uint16(srcX))
	w.Write2b(uint16(srcY))

	for _, p := range points {
		writePointFixed(w, p)
	}
}

// #WREQ
func writeTriFan(w *x.Writer, op uint8, src, dst Picture, maskFormat PictFormat,
	srcX, srcY int16, points []PointFixed) {

	w.WritePad(4)
	w.Write1b(op)
	w.WritePad(3)

	w.Write4b(uint32(src))
	w.Write4b(uint32(dst))
	w.Write4b(uint32(maskFormat))

	w.Write2b(uint16(srcX))
	w.Write2b(uint16(srcY))

	for _, p := range points {
		writePointFixed(w, p)
	}
}

// #WREQ
func writeCreateGlyphSet(w *x.Writer, gsId GlyphSet, format PictFormat) {
	w.WritePad(4)
	w.Write4b(uint32(gsId))
	w.Write4b(uint32(format))
}

// #WREQ
func writeReferenceGlyphSet(w *x.Writer, gsId, existing GlyphSet) {
	w.WritePad(4)
	w.Write4b(uint32(gsId))
	w.Write4b(uint32(existing))
}

// #WREQ
func writeFreeGlyphSet(w *x.Writer, glyphSet GlyphSet) {
	w.WritePad(4)
	w.Write4b(uint32(glyphSet))
}

type GlyphInfo struct {
	Width  uint16
	Height uint16
	X      int16
	Y      int16
	XOff   int16
	YOff   int16
}

func writeGlyphInfo(w *x.Writer, v GlyphInfo) {
	w.Write2b(v.Width)
	w.Write2b(v.Height)
	w.Write2b(uint16(v.X))
	w.Write2b(uint16(v.Y))
	w.Write2b(uint16(v.XOff))
	w.Write2b(uint16(v.YOff))
}

// #WREQ
func writeAddGlyphs(w *x.Writer, glyphSet GlyphSet, glyphIds []uint32, glyphs []GlyphInfo,
	data []byte) {
	w.WritePad(4)
	w.Write4b(uint32(glyphSet))
	w.Write4b(uint32(len(glyphIds)))
	for _, gi := range glyphs {
		writeGlyphInfo(w, gi)
	}
	w.WriteBytes(data)
	w.WritePad(x.Pad(len(data)))
}

// #WREQ
func writeFreeGlyphs(w *x.Writer, glyphSet GlyphSet, glyphs []Glyph) {
	w.WritePad(4)
	w.Write4b(uint32(glyphSet))
	for _, g := range glyphs {
		w.Write4b(uint32(g))
	}
}

func writeCompositeGlyphsN(w *x.Writer, op uint8, src, dst Picture, maskFormat PictFormat,
	glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) {

	w.WritePad(4)
	w.Write1b(op)
	w.WritePad(3)
	w.Write4b(uint32(src))
	w.Write4b(uint32(dst))
	w.Write4b(uint32(maskFormat))
	w.Write4b(uint32(glyphSet))
	w.Write2b(uint16(srcX))
	w.Write2b(uint16(srcY))

	w.WriteBytes(glyphCmds)
	w.WritePad(x.Pad(len(glyphCmds)))
}

// #WREQ
func writeCompositeGlyphs8(w *x.Writer, op uint8, src, dst Picture, maskFormat PictFormat,
	glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) {
	writeCompositeGlyphsN(w, op, src, dst, maskFormat, glyphSet, srcX, srcY, glyphCmds)
}

// #WREQ
func writeCompositeGlyphs16(w *x.Writer, op uint8, src, dst Picture, maskFormat PictFormat,
	glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) {
	writeCompositeGlyphsN(w, op, src, dst, maskFormat, glyphSet, srcX, srcY, glyphCmds)
}

// #WREQ
func writeCompositeGlyphs32(w *x.Writer, op uint8, src, dst Picture, maskFormat PictFormat,
	glyphSet GlyphSet, srcX, srcY int16, glyphCmds []byte) {
	writeCompositeGlyphsN(w, op, src, dst, maskFormat, glyphSet, srcX, srcY, glyphCmds)
}

// #WREQ
func writeCreateCursor(w *x.Writer, cid x.Cursor, source Picture, X, y uint16) {
	w.WritePad(4)
	w.Write4b(uint32(cid))
	w.Write4b(uint32(source))
	w.Write2b(X)
	w.Write2b(y)
}

type AnimCursorElt struct {
	Cursor x.Cursor
	Delay  uint32
}

func writeAnimCursorElt(w *x.Writer, v AnimCursorElt) {
	w.Write4b(uint32(v.Cursor))
	w.Write4b(v.Delay)
}

// #WREQ
func writeCreateAnimCursor(w *x.Writer, cid x.Cursor, cursors []AnimCursorElt) {
	w.WritePad(4)
	w.Write4b(uint32(cid))
	for _, c := range cursors {
		writeAnimCursorElt(w, c)
	}
}

type SpanFixed struct {
	L, R, Y Fixed
}

type Trap struct {
	Top SpanFixed
	Bot SpanFixed
}

func writeSpanFixed(w *x.Writer, v SpanFixed) {
	w.Write4b(uint32(v.L))
	w.Write4b(uint32(v.R))
	w.Write4b(uint32(v.Y))
}

func writeTrap(w *x.Writer, v *Trap) {
	writeSpanFixed(w, v.Top)
	writeSpanFixed(w, v.Bot)
}

// #WREQ
func writeAddTraps(w *x.Writer, picture Picture, xOff, yOff int16, traps []Trap) {
	w.WritePad(4)
	w.Write4b(uint32(picture))
	w.Write2b(uint16(xOff))
	w.Write2b(uint16(yOff))
	for _, trap := range traps {
		writeTrap(w, &trap)
	}
}

// #WREQ
func writeCreateSolidFill(w *x.Writer, pid Picture, color Color) {
	w.WritePad(4)
	w.Write4b(uint32(pid))
	writeColor(w, color)
}

// #WREQ
func writeCreateLinearGradient(w *x.Writer, picture Picture, p1, p2 PointFixed,
	stops []Fixed, colors []Color) {
	w.WritePad(4)
	w.Write4b(uint32(picture))
	writePointFixed(w, p1)
	writePointFixed(w, p2)

	w.Write4b(uint32(len(stops)))

	for _, s := range stops {
		w.Write4b(uint32(s))
	}
	for _, c := range colors {
		writeColor(w, c)
	}
}

// #WREQ
func writeCreateRadialGradient(w *x.Writer, picture Picture, inner, outer PointFixed,
	innerRadius, outerRadius Fixed, stops []Fixed, colors []Color) {
	w.WritePad(4)
	w.Write4b(uint32(picture))
	writePointFixed(w, inner)
	writePointFixed(w, outer)
	w.Write4b(uint32(innerRadius))
	w.Write4b(uint32(outerRadius))

	w.Write4b(uint32(len(stops)))

	for _, s := range stops {
		w.Write4b(uint32(s))
	}
	for _, c := range colors {
		writeColor(w, c)
	}
}

// #WREQ
func writeCreateConicalGradient(w *x.Writer, picture Picture, center PointFixed,
	angle Fixed, stops []Fixed, colors []Color) {

	w.WritePad(4)
	w.Write4b(uint32(picture))

	writePointFixed(w, center)
	w.Write4b(uint32(angle))

	for _, s := range stops {
		w.Write4b(uint32(s))
	}
	for _, c := range colors {
		writeColor(w, c)
	}
}
