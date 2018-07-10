package render

import "github.com/linuxdeepin/go-x11-client"

type PictFormatError struct {
	x.ResourceErrorBase
}

func (e PictFormatError) Error() string {
	return "render.PictFormatError" + e.Msg()
}

func readPictFormatError(r *x.Reader) x.Error {
	return PictFormatError{x.ReadResourceErrorBase(r)}
}

type PictureError struct {
	x.ResourceErrorBase
}

func (e PictureError) Error() string {
	return "render.PictureError" + e.Msg()
}

func readPictureError(r *x.Reader) x.Error {
	return PictureError{x.ReadResourceErrorBase(r)}
}

type PictOpError struct {
	x.ResourceErrorBase
}

func (e PictOpError) Error() string {
	return "render.PictOpError" + e.Msg()
}

func readPictOpError(r *x.Reader) x.Error {
	return PictOpError{x.ReadResourceErrorBase(r)}
}

type GlyphError struct {
	x.ResourceErrorBase
}

func (e GlyphError) Error() string {
	return "render.GlyphError" + e.Msg()
}

func readGlyphError(r *x.Reader) x.Error {
	return GlyphError{x.ReadResourceErrorBase(r)}
}

type GlyphSetError struct {
	x.ResourceErrorBase
}

func (e GlyphSetError) Error() string {
	return "render.GlyphSetError" + e.Msg()
}

func readGlyphSetError(r *x.Reader) x.Error {
	return GlyphSetError{x.ReadResourceErrorBase(r)}
}
