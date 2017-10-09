package mousebind

import (
	x "github.com/linuxdeepin/go-x11-client"
)

func GrabPointer(conn *x.Conn, win x.Window, eventMask uint16, confineTo x.Window, cursor x.Cursor) error {
	reply, err := x.GrabPointer(conn, x.True, win, eventMask,
		x.GrabModeAsync, x.GrabModeAsync,
		confineTo, cursor, x.CurrentTime).Reply(conn)
	if err != nil {
		return err
	}
	if reply.Status == x.GrabStatusSuccess {
		return nil
	}
	return GrabPointerError{reply.Status}
}

type GrabPointerError struct {
	Status byte
}

func (err GrabPointerError) Error() string {
	const errMsgPrefix = "GrabPointer Failed status: "

	switch err.Status {
	case x.GrabStatusAlreadyGrabbed:
		return errMsgPrefix + "AlreadyGrabbed"
	case x.GrabStatusInvalidTime:
		return errMsgPrefix + "InvalidTime"
	case x.GrabStatusNotViewable:
		return errMsgPrefix + "NotViewable"
	case x.GrabStatusFrozen:
		return errMsgPrefix + "Frozen"
	default:
		return errMsgPrefix + "Unknown"
	}
}

func UngrabPointer(conn *x.Conn) error {
	return x.UngrabPointerChecked(conn, x.CurrentTime).Check(conn)
}
