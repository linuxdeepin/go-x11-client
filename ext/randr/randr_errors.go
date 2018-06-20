package randr

import "github.com/linuxdeepin/go-x11-client"

type BadOutputError struct {
	x.ResourceErrorBase
}

func readBadOutputError(r *x.Reader) x.Error {
	return BadOutputError{x.ReadResourceErrorBase(r)}
}

func (e BadOutputError) Error() string {
	return "BadOutputError" + e.Msg()
}

type BadCrtcError struct {
	x.ResourceErrorBase
}

func readBadCrtcError(r *x.Reader) x.Error {
	return BadCrtcError{x.ReadResourceErrorBase(r)}
}

func (e BadCrtcError) Error() string {
	return "BadCrtcError" + e.Msg()
}

type BadModeError struct {
	x.ResourceErrorBase
}

func readBadModeError(r *x.Reader) x.Error {
	return BadModeError{x.ReadResourceErrorBase(r)}
}

func (e BadModeError) Error() string {
	return "BadModeError" + e.Msg()
}

type BadProviderError struct {
	x.ResourceErrorBase
}

func readBadProviderError(r *x.Reader) x.Error {
	return BadProviderError{x.ReadResourceErrorBase(r)}
}

func (e BadProviderError) Error() string {
	return "BadProviderError" + e.Msg()
}
