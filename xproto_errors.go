package x

import "fmt"

type Error interface {
	GetCode() uint8
	GetSequenceNumber() uint16
	GetMinorOpcode() uint16
	GetMajorOpcode() uint8
	error
}

type ErrorBase struct {
	Code        uint8
	Sequence    uint16
	MinorOpcode uint16
	MajorOpcode uint8
}

func readErrorBase(r *Reader) ErrorBase {
	var e ErrorBase
	// Error
	r.Read1b()
	e.Code = r.Read1b()
	e.Sequence = r.Read2b()

	r.ReadPad(4)

	e.MinorOpcode = r.Read2b()
	e.MajorOpcode = r.Read1b()

	r.ReadPad(21)
	return e
}

func (e ErrorBase) GetCode() uint8 {
	return e.Code
}

func (e ErrorBase) GetSequenceNumber() uint16 {
	return e.Sequence
}

func (e ErrorBase) GetMinorOpcode() uint16 {
	return e.MinorOpcode
}

func (e ErrorBase) GetMajorOpcode() uint8 {
	return e.MajorOpcode
}

func (e ErrorBase) msg() string {
	return fmt.Sprintf("{code=%d, seq=%d, major=%d, minor=%d}",
		e.Code, e.Sequence, e.MajorOpcode, e.MinorOpcode)
}

type ResourceErrorBase struct {
	Code        uint8
	Sequence    uint16
	Bad         uint32
	MinorOpcode uint16
	MajorOpcode uint8
}

func ReadResourceErrorBase(r *Reader) ResourceErrorBase {
	var e ResourceErrorBase
	// Error
	r.Read1b()
	e.Code = r.Read1b()
	e.Sequence = r.Read2b()

	e.Bad = r.Read4b()

	e.MinorOpcode = r.Read2b()
	e.MajorOpcode = r.Read1b()

	r.ReadPad(21)
	return e
}

func (e ResourceErrorBase) GetCode() uint8 {
	return e.Code
}

func (e ResourceErrorBase) GetSequenceNumber() uint16 {
	return e.Sequence
}

func (e ResourceErrorBase) GetMinorOpcode() uint16 {
	return e.MinorOpcode
}

func (e ResourceErrorBase) GetMajorOpcode() uint8 {
	return e.MajorOpcode
}

func (e ResourceErrorBase) Msg() string {
	return fmt.Sprintf("{code=%d, seq=%d, major=%d, minor=%d, bad=%d}",
		e.Code, e.Sequence, e.MajorOpcode, e.MinorOpcode, e.Bad)
}

type RequestError struct {
	ErrorBase
}

func readRequestError(r *Reader) Error {
	return RequestError{readErrorBase(r)}
}

func (e RequestError) Error() string {
	return "ReqeustError" + e.msg()
}

type ValueError struct {
	ResourceErrorBase
}

func readValueError(r *Reader) Error {
	return ValueError{ReadResourceErrorBase(r)}
}

func (e ValueError) Error() string {
	return "ValueError" + e.Msg()
}

type WindowError struct {
	ResourceErrorBase
}

func readWindowError(r *Reader) Error {
	return WindowError{ReadResourceErrorBase(r)}
}

func (e WindowError) Error() string {
	return "WindowError" + e.Msg()
}

type PixmapError struct {
	ResourceErrorBase
}

func readPixmapError(r *Reader) Error {
	return PixmapError{ReadResourceErrorBase(r)}
}

func (e PixmapError) Error() string {
	return "PixmapError" + e.Msg()
}

type AtomError struct {
	ResourceErrorBase
}

func readAtomError(r *Reader) Error {
	return AtomError{ReadResourceErrorBase(r)}
}

func (e AtomError) Error() string {
	return "AtomError" + e.Msg()
}

type CursorError struct {
	ResourceErrorBase
}

func readCursorError(r *Reader) Error {
	return CursorError{ReadResourceErrorBase(r)}
}

func (e CursorError) Error() string {
	return "CursorError" + e.Msg()
}

type FontError struct {
	ResourceErrorBase
}

func readFontError(r *Reader) Error {
	return FontError{ReadResourceErrorBase(r)}
}

func (e FontError) Error() string {
	return "FontError" + e.Msg()
}

type MatchError struct {
	ErrorBase
}

func readMatchError(r *Reader) Error {
	return MatchError{readErrorBase(r)}
}

func (e MatchError) Error() string {
	return "MatchError" + e.msg()
}

type DrawableError struct {
	ResourceErrorBase
}

func readDrawableError(r *Reader) Error {
	return DrawableError{ReadResourceErrorBase(r)}
}

func (e DrawableError) Error() string {
	return "DrawableError" + e.Msg()
}

type AccessError struct {
	ErrorBase
}

func readAccessError(r *Reader) Error {
	return AccessError{readErrorBase(r)}
}

func (e AccessError) Error() string {
	return "AccessError" + e.msg()
}

type AllocError struct {
	ErrorBase
}

func readAllocError(r *Reader) Error {
	return AllocError{readErrorBase(r)}
}

func (e AllocError) Error() string {
	return "AllocError" + e.msg()
}

type ColormapError struct {
	ResourceErrorBase
}

func readColormapError(r *Reader) Error {
	return ColormapError{ReadResourceErrorBase(r)}
}

func (e ColormapError) Error() string {
	return "ColormapError" + e.Msg()
}

type GContextError struct {
	ResourceErrorBase
}

func readGContextError(r *Reader) Error {
	return GContextError{ReadResourceErrorBase(r)}
}

func (e GContextError) Error() string {
	return "GContextError" + e.Msg()
}

type IdChoiceError struct {
	ResourceErrorBase
}

func readIdChoiceError(r *Reader) Error {
	return IdChoiceError{ReadResourceErrorBase(r)}
}

func (e IdChoiceError) Error() string {
	return "IdChoiceError" + e.Msg()
}

type NameError struct {
	ErrorBase
}

func (e NameError) Error() string {
	return "NameError" + e.msg()
}

func readNameError(r *Reader) Error {
	return NameError{readErrorBase(r)}
}

type LengthError struct {
	ErrorBase
}

func readLengthError(r *Reader) Error {
	return LengthError{readErrorBase(r)}
}

func (e LengthError) Error() string {
	return "LengthError" + e.msg()
}

type ImplementationError struct {
	ErrorBase
}

func readImplementationError(r *Reader) Error {
	return ImplementationError{readErrorBase(r)}
}

func (e ImplementationError) Error() string {
	return "ImplementationError" + e.msg()
}
