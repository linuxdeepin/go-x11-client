package input

import "github.com/linuxdeepin/go-x11-client"

type DeviceError struct {
	x.ResourceErrorBase
}

func (e DeviceError) Error() string {
	return "input.DeviceError" + e.Msg()
}

func readDeviceError(r *x.Reader) x.Error {
	return DeviceError{x.ReadResourceErrorBase(r)}
}

type DeviceBusyError struct {
	x.ResourceErrorBase
}

func (e DeviceBusyError) Error() string {
	return "input.DeviceBusyError" + e.Msg()
}

func readDeviceBusyError(r *x.Reader) x.Error {
	return DeviceBusyError{x.ReadResourceErrorBase(r)}
}

type EventError struct {
	x.ResourceErrorBase
}

func (e EventError) Error() string {
	return "input.EventError" + e.Msg()
}

func readEventError(r *x.Reader) x.Error {
	return EventError{x.ReadResourceErrorBase(r)}
}

type ModeError struct {
	x.ResourceErrorBase
}

func (e ModeError) Error() string {
	return "input.ModeError" + e.Msg()
}

func readModeError(r *x.Reader) x.Error {
	return ModeError{x.ReadResourceErrorBase(r)}
}

type ClassError struct {
	x.ResourceErrorBase
}

func (e ClassError) Error() string {
	return "input.ClassError" + e.Msg()
}

func readClassError(r *x.Reader) x.Error {
	return ClassError{x.ReadResourceErrorBase(r)}
}
