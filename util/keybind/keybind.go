// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package keybind

import (
	x "github.com/linuxdeepin/go-x11-client"
	"github.com/linuxdeepin/go-x11-client/util/keysyms"
)

var grabMods []uint16

func init() {
	grabMods = make([]uint16, len(keysyms.LockMods)+1)
	copy(grabMods, keysyms.LockMods)
}

func Grab(conn *x.Conn, win x.Window, mods uint16, key x.Keycode) {
	for _, m := range grabMods {
		x.GrabKey(conn, true, win, mods|m,
			key, x.GrabModeAsync, x.GrabModeAsync)
	}
}

func Ungrab(conn *x.Conn, win x.Window, mods uint16, key x.Keycode) {
	for _, m := range grabMods {
		_ = x.UngrabKeyChecked(conn, key, win, mods|m).Check(conn)
	}
}

func GrabChecked(conn *x.Conn, win x.Window, mods uint16, key x.Keycode) error {
	return GrabCheckedV2(conn, win, mods, key, x.GrabModeAsync, x.GrabModeAsync)
}

func GrabCheckedV2(conn *x.Conn, win x.Window, mods uint16, key x.Keycode, pointerMode uint8, keyboardMode uint8) error {
	for _, m := range grabMods {
		err := x.GrabKeyChecked(conn, true, win, mods|m,
			key, pointerMode, keyboardMode).Check(conn)
		if err != nil {
			return err
		}
	}
	return nil
}

// GrabKeyboard grabs the entire keyboard.
func GrabKeyboard(conn *x.Conn, win x.Window) error {
	reply, err := x.GrabKeyboard(conn, true, win, x.CurrentTime,
		x.GrabModeAsync, x.GrabModeAsync).Reply(conn)
	if err != nil {
		return err
	}

	if reply.Status == x.GrabStatusSuccess {
		// successful
		return nil
	}
	return GrabKeyboardError{reply.Status}
}

type GrabKeyboardError struct {
	Status byte
}

func (err GrabKeyboardError) Error() string {
	const errMsgPrefix = "GrabKeyboard Failed status: "

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

// UngrabKeyboard undoes GrabKeyboard.
func UngrabKeyboard(conn *x.Conn) error {
	return x.UngrabKeyboardChecked(conn, x.CurrentTime).Check(conn)
}
