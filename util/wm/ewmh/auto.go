// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package ewmh

import x "github.com/linuxdeepin/go-x11-client"
import "errors"

type GetUTF8StrCookie x.GetPropertyCookie

func (cookie GetUTF8StrCookie) Reply(c *x.Conn) (string, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return "", err
	}
	if reply.Type != getAtom(c, "UTF8_STRING") {
		return "", errors.New("bad reply")
	}
	return getUTF8StrFromReply(reply)
}

type GetUTF8StrsCookie x.GetPropertyCookie

func (cookie GetUTF8StrsCookie) Reply(c *x.Conn) ([]string, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return nil, err
	}
	if reply.Type != getAtom(c, "UTF8_STRING") {
		return nil, errors.New("bad reply")
	}
	return getUTF8StrsFromReply(reply)
}

type GetBooleanCookie x.GetPropertyCookie

func (cookie GetBooleanCookie) Reply(c *x.Conn) (bool, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return false, err
	}
	if reply.Type != x.AtomCardinal {
		return false, errors.New("bad reply")
	}
	return getBooleanFromReply(reply)
}

func getWindowFromReply(r *x.GetPropertyReply) (x.Window, error) {
	if r.Format != 32 || len(r.Value) != 4 {
		return 0, errors.New("bad reply")
	}
	return x.Window(x.Get32(r.Value)), nil
}

type GetWindowCookie x.GetPropertyCookie

func (cookie GetWindowCookie) Reply(c *x.Conn) (x.Window, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return 0, err
	}
	if reply.Type != x.AtomWindow {
		return 0, errors.New("bad reply")
	}
	return getWindowFromReply(reply)
}

func getCardinalFromReply(r *x.GetPropertyReply) (uint32, error) {
	if r.Format != 32 || len(r.Value) != 4 {
		return 0, errors.New("bad reply")
	}
	return uint32(x.Get32(r.Value)), nil
}

type GetCardinalCookie x.GetPropertyCookie

func (cookie GetCardinalCookie) Reply(c *x.Conn) (uint32, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return 0, err
	}
	if reply.Type != x.AtomCardinal {
		return 0, errors.New("bad reply")
	}
	return getCardinalFromReply(reply)
}

func getWindowsFromReply(r *x.GetPropertyReply) ([]x.Window, error) {
	if r.Format != 32 {
		return nil, errors.New("bad reply")
	}
	count := len(r.Value) / 4
	ret := make([]x.Window, count)
	rdr := x.NewReaderFromData(r.Value)
	for i := 0; i < count; i++ {
		ret[i] = x.Window(rdr.Read4b())
	}
	return ret, nil
}

type GetWindowsCookie x.GetPropertyCookie

func (cookie GetWindowsCookie) Reply(c *x.Conn) ([]x.Window, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return nil, err
	}
	if reply.Type != x.AtomWindow {
		return nil, errors.New("bad reply")
	}
	return getWindowsFromReply(reply)
}

func getAtomsFromReply(r *x.GetPropertyReply) ([]x.Atom, error) {
	if r.Format != 32 {
		return nil, errors.New("bad reply")
	}
	count := len(r.Value) / 4
	ret := make([]x.Atom, count)
	rdr := x.NewReaderFromData(r.Value)
	for i := 0; i < count; i++ {
		ret[i] = x.Atom(rdr.Read4b())
	}
	return ret, nil
}

type GetAtomsCookie x.GetPropertyCookie

func (cookie GetAtomsCookie) Reply(c *x.Conn) ([]x.Atom, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return nil, err
	}
	if reply.Type != x.AtomAtom {
		return nil, errors.New("bad reply")
	}
	return getAtomsFromReply(reply)
}

func getCardinalsFromReply(r *x.GetPropertyReply) ([]uint32, error) {
	if r.Format != 32 {
		return nil, errors.New("bad reply")
	}
	count := len(r.Value) / 4
	ret := make([]uint32, count)
	rdr := x.NewReaderFromData(r.Value)
	for i := 0; i < count; i++ {
		ret[i] = uint32(rdr.Read4b())
	}
	return ret, nil
}

type GetCardinalsCookie x.GetPropertyCookie

func (cookie GetCardinalsCookie) Reply(c *x.Conn) ([]uint32, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return nil, err
	}
	if reply.Type != x.AtomCardinal {
		return nil, errors.New("bad reply")
	}
	return getCardinalsFromReply(reply)
}

/**
 *    _NET_SUPPORTED
 */

func GetSupported(c *x.Conn) GetAtomsCookie {
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_SUPPORTED")
	cookie := x.GetProperty(c, false, rootWin,
		atom, x.AtomAtom, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func SetSupportedChecked(c *x.Conn, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_SUPPORTED")
	return x.ChangePropertyChecked(c, x.PropModeReplace, rootWin,
		atom, x.AtomAtom, 32, w.Bytes())
}

func SetSupported(c *x.Conn, vals []x.Atom) {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_SUPPORTED")
	x.ChangeProperty(c, x.PropModeReplace, rootWin,
		atom, x.AtomAtom, 32, w.Bytes())
}

/**
 *    _NET_CLIENT_LIST
 */

func GetClientList(c *x.Conn) GetWindowsCookie {
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_CLIENT_LIST")
	cookie := x.GetProperty(c, false, rootWin,
		atom, x.AtomWindow, 0, LENGTH_MAX)
	return GetWindowsCookie(cookie)
}

func SetClientListChecked(c *x.Conn, vals []x.Window) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_CLIENT_LIST")
	return x.ChangePropertyChecked(c, x.PropModeReplace, rootWin,
		atom, x.AtomWindow, 32, w.Bytes())
}

func SetClientList(c *x.Conn, vals []x.Window) {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_CLIENT_LIST")
	x.ChangeProperty(c, x.PropModeReplace, rootWin,
		atom, x.AtomWindow, 32, w.Bytes())
}

/**
 *    _NET_CLIENT_LIST_STACKING
 */

func GetClientListStacking(c *x.Conn) GetWindowsCookie {
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_CLIENT_LIST_STACKING")
	cookie := x.GetProperty(c, false, rootWin,
		atom, x.AtomWindow, 0, LENGTH_MAX)
	return GetWindowsCookie(cookie)
}

func SetClientListStackingChecked(c *x.Conn, vals []x.Window) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_CLIENT_LIST_STACKING")
	return x.ChangePropertyChecked(c, x.PropModeReplace, rootWin,
		atom, x.AtomWindow, 32, w.Bytes())
}

func SetClientListStacking(c *x.Conn, vals []x.Window) {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_CLIENT_LIST_STACKING")
	x.ChangeProperty(c, x.PropModeReplace, rootWin,
		atom, x.AtomWindow, 32, w.Bytes())
}

/**
 *    _NET_NUMBER_OF_DESKTOPS
 */

func GetNumberOfDesktop(c *x.Conn) GetCardinalCookie {
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_NUMBER_OF_DESKTOPS")
	cookie := x.GetProperty(c, false, rootWin,
		atom, x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func SetNumberOfDesktopChecked(c *x.Conn, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_NUMBER_OF_DESKTOPS")
	return x.ChangePropertyChecked(c, x.PropModeReplace, rootWin,
		atom, x.AtomCardinal, 32, w.Bytes())
}

func SetNumberOfDesktop(c *x.Conn, val uint32) {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_NUMBER_OF_DESKTOPS")
	x.ChangeProperty(c, x.PropModeReplace, rootWin,
		atom, x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_DESKTOP_GEOMETRY
 */

func GetDesktopGeometry(c *x.Conn) GetDesktopGeometryCookie {
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_DESKTOP_GEOMETRY")
	cookie := x.GetProperty(c, false, rootWin,
		atom, x.AtomCardinal, 0, 2)
	return GetDesktopGeometryCookie(cookie)
}

type GetDesktopGeometryCookie x.GetPropertyCookie

func (cookie GetDesktopGeometryCookie) Reply(c *x.Conn) (DesktopGeometry, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return DesktopGeometry{}, err
	}
	if reply.Type != x.AtomCardinal {
		return DesktopGeometry{}, errors.New("bad reply")
	}
	return getDesktopGeometryFromReply(reply)
}

type DesktopGeometry struct {
	Width, Height uint32
}

func getDesktopGeometryFromReply(reply *x.GetPropertyReply) (DesktopGeometry, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return DesktopGeometry{}, err
	}

	if len(list) != 2 {
		return DesktopGeometry{}, errors.New("length of list is not 2")
	}
	return DesktopGeometry{
		Width:  list[0],
		Height: list[1],
	}, nil
}

func SetDesktopGeometryChecked(c *x.Conn, val DesktopGeometry) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Width)
	w.Write4b(val.Height)
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_DESKTOP_GEOMETRY")
	return x.ChangePropertyChecked(c, x.PropModeReplace, rootWin,
		atom, x.AtomCardinal, 32, w.Bytes())
}

func SetDesktopGeometry(c *x.Conn, val DesktopGeometry) {
	w := x.NewWriter()
	w.Write4b(val.Width)
	w.Write4b(val.Height)
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_DESKTOP_GEOMETRY")
	x.ChangeProperty(c, x.PropModeReplace, rootWin,
		atom, x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_DESKTOP_VIEWPORT
 */

func GetDesktopViewport(c *x.Conn) GetDesktopViewportCookie {
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_DESKTOP_VIEWPORT")
	cookie := x.GetProperty(c, false, rootWin,
		atom, x.AtomCardinal, 0, LENGTH_MAX)
	return GetDesktopViewportCookie(cookie)
}

type GetDesktopViewportCookie x.GetPropertyCookie

func (cookie GetDesktopViewportCookie) Reply(c *x.Conn) ([]ViewportLeftTopCorner, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return nil, err
	}
	if reply.Type != x.AtomCardinal {
		return nil, errors.New("bad reply")
	}
	return getDesktopViewportFromReply(reply)
}

type ViewportLeftTopCorner struct {
	X, Y uint32
}

func getDesktopViewportFromReply(reply *x.GetPropertyReply) ([]ViewportLeftTopCorner, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return nil, err
	}
	if len(list)%2 != 0 {
		return nil, errors.New("length of list is not a multiple of 2")
	}
	ret := make([]ViewportLeftTopCorner, len(list)/2)
	idx := 0
	for i := 0; i < len(list); i += 2 {
		ret[idx] = ViewportLeftTopCorner{
			X: list[i],
			Y: list[i+1],
		}
		idx++
	}
	return ret, nil
}

func SetDesktopViewportChecked(c *x.Conn, vals []ViewportLeftTopCorner) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(val.X)
		w.Write4b(val.Y)
	}
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_DESKTOP_VIEWPORT")
	return x.ChangePropertyChecked(c, x.PropModeReplace, rootWin,
		atom, x.AtomCardinal, 32, w.Bytes())
}

func SetDesktopViewport(c *x.Conn, vals []ViewportLeftTopCorner) {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(val.X)
		w.Write4b(val.Y)
	}
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_DESKTOP_VIEWPORT")
	x.ChangeProperty(c, x.PropModeReplace, rootWin,
		atom, x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_CURRENT_DESKTOP
 */

func GetCurrentDesktop(c *x.Conn) GetCardinalCookie {
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_CURRENT_DESKTOP")
	cookie := x.GetProperty(c, false, rootWin,
		atom, x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func SetCurrentDesktopChecked(c *x.Conn, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_CURRENT_DESKTOP")
	return x.ChangePropertyChecked(c, x.PropModeReplace, rootWin,
		atom, x.AtomCardinal, 32, w.Bytes())
}

func SetCurrentDesktop(c *x.Conn, val uint32) {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_CURRENT_DESKTOP")
	x.ChangeProperty(c, x.PropModeReplace, rootWin,
		atom, x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_DESKTOP_NAMES
 */

func GetDesktopNames(c *x.Conn) GetUTF8StrsCookie {
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_DESKTOP_NAMES")
	cookie := x.GetProperty(c, false, rootWin,
		atom, getAtom(c, "UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrsCookie(cookie)
}

func SetDesktopNamesChecked(c *x.Conn, vals []string) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.WriteString(val)
		w.Write1b(0)
	}
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_DESKTOP_NAMES")
	atomUTF8String, _ := c.GetAtom("UTF8_STRING")
	return x.ChangePropertyChecked(c, x.PropModeReplace, rootWin,
		atom, atomUTF8String, 8, w.Bytes())
}

func SetDesktopNames(c *x.Conn, vals []string) {
	w := x.NewWriter()
	for _, val := range vals {
		w.WriteString(val)
		w.Write1b(0)
	}
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_DESKTOP_NAMES")
	atomUTF8String, _ := c.GetAtom("UTF8_STRING")
	x.ChangeProperty(c, x.PropModeReplace, rootWin,
		atom, atomUTF8String, 8, w.Bytes())
}

/**
 *    _NET_ACTIVE_WINDOW
 */

func GetActiveWindow(c *x.Conn) GetWindowCookie {
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_ACTIVE_WINDOW")
	cookie := x.GetProperty(c, false, rootWin,
		atom, x.AtomWindow, 0, 1)
	return GetWindowCookie(cookie)
}

func SetActiveWindowChecked(c *x.Conn, val x.Window) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_ACTIVE_WINDOW")
	return x.ChangePropertyChecked(c, x.PropModeReplace, rootWin,
		atom, x.AtomWindow, 32, w.Bytes())
}

func SetActiveWindow(c *x.Conn, val x.Window) {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_ACTIVE_WINDOW")
	x.ChangeProperty(c, x.PropModeReplace, rootWin,
		atom, x.AtomWindow, 32, w.Bytes())
}

/**
 *    _NET_WORKAREA
 */

func GetWorkarea(c *x.Conn) GetWorkareaCookie {
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_WORKAREA")
	cookie := x.GetProperty(c, false, rootWin,
		atom, x.AtomCardinal, 0, LENGTH_MAX)
	return GetWorkareaCookie(cookie)
}

type GetWorkareaCookie x.GetPropertyCookie

func (cookie GetWorkareaCookie) Reply(c *x.Conn) ([]WorkareaGeometry, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return nil, err
	}
	if reply.Type != x.AtomCardinal {
		return nil, errors.New("bad reply")
	}
	return getWorkareaFromReply(reply)
}

type WorkareaGeometry struct {
	X, Y, Width, Height uint32
}

func getWorkareaFromReply(reply *x.GetPropertyReply) ([]WorkareaGeometry, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return nil, err
	}
	if len(list)%4 != 0 {
		return nil, errors.New("length of list is not a multiple of 4")
	}
	ret := make([]WorkareaGeometry, len(list)/4)
	idx := 0
	for i := 0; i < len(list); i += 4 {
		ret[idx] = WorkareaGeometry{
			X:      list[i],
			Y:      list[i+1],
			Width:  list[i+2],
			Height: list[i+3],
		}
		idx++
	}
	return ret, nil
}

func SetWorkareaChecked(c *x.Conn, vals []WorkareaGeometry) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(val.X)
		w.Write4b(val.Y)
		w.Write4b(val.Width)
		w.Write4b(val.Height)
	}
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_WORKAREA")
	return x.ChangePropertyChecked(c, x.PropModeReplace, rootWin,
		atom, x.AtomCardinal, 32, w.Bytes())
}

func SetWorkarea(c *x.Conn, vals []WorkareaGeometry) {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(val.X)
		w.Write4b(val.Y)
		w.Write4b(val.Width)
		w.Write4b(val.Height)
	}
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_WORKAREA")
	x.ChangeProperty(c, x.PropModeReplace, rootWin,
		atom, x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_SUPPORTING_WM_CHECK
 */

func GetSupportingWMCheck(c *x.Conn) GetWindowCookie {
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_SUPPORTING_WM_CHECK")
	cookie := x.GetProperty(c, false, rootWin,
		atom, x.AtomWindow, 0, 1)
	return GetWindowCookie(cookie)
}

func SetSupportingWMCheckChecked(c *x.Conn, val x.Window) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_SUPPORTING_WM_CHECK")
	return x.ChangePropertyChecked(c, x.PropModeReplace, rootWin,
		atom, x.AtomWindow, 32, w.Bytes())
}

func SetSupportingWMCheck(c *x.Conn, val x.Window) {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_SUPPORTING_WM_CHECK")
	x.ChangeProperty(c, x.PropModeReplace, rootWin,
		atom, x.AtomWindow, 32, w.Bytes())
}

/**
 *    _NET_VIRTUAL_ROOTS
 */

func GetVirtualRoots(c *x.Conn) GetWindowsCookie {
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_VIRTUAL_ROOTS")
	cookie := x.GetProperty(c, false, rootWin,
		atom, x.AtomWindow, 0, LENGTH_MAX)
	return GetWindowsCookie(cookie)
}

func SetVirtualRootsChecked(c *x.Conn, vals []x.Window) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_VIRTUAL_ROOTS")
	return x.ChangePropertyChecked(c, x.PropModeReplace, rootWin,
		atom, x.AtomWindow, 32, w.Bytes())
}

func SetVirtualRoots(c *x.Conn, vals []x.Window) {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_VIRTUAL_ROOTS")
	x.ChangeProperty(c, x.PropModeReplace, rootWin,
		atom, x.AtomWindow, 32, w.Bytes())
}

/**
 *    _NET_DESKTOP_LAYOUT
 */

func GetDesktopLayout(c *x.Conn) GetDesktopLayoutCookie {
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_DESKTOP_LAYOUT")
	cookie := x.GetProperty(c, false, rootWin,
		atom, x.AtomCardinal, 0, 4)
	return GetDesktopLayoutCookie(cookie)
}

type GetDesktopLayoutCookie x.GetPropertyCookie

func (cookie GetDesktopLayoutCookie) Reply(c *x.Conn) (DesktopLayout, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return DesktopLayout{}, err
	}
	if reply.Type != x.AtomCardinal {
		return DesktopLayout{}, errors.New("bad reply")
	}
	return getDesktopLayoutFromReply(reply)
}

type DesktopLayout struct {
	Orientation, Columns, Rows, StartingCorner uint32
}

func getDesktopLayoutFromReply(reply *x.GetPropertyReply) (DesktopLayout, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return DesktopLayout{}, err
	}

	if len(list) != 4 {
		return DesktopLayout{}, errors.New("length of list is not 4")
	}
	return DesktopLayout{
		Orientation:    list[0],
		Columns:        list[1],
		Rows:           list[2],
		StartingCorner: list[3],
	}, nil
}

func SetDesktopLayoutChecked(c *x.Conn, val DesktopLayout) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Orientation)
	w.Write4b(val.Columns)
	w.Write4b(val.Rows)
	w.Write4b(val.StartingCorner)
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_DESKTOP_LAYOUT")
	return x.ChangePropertyChecked(c, x.PropModeReplace, rootWin,
		atom, x.AtomCardinal, 32, w.Bytes())
}

func SetDesktopLayout(c *x.Conn, val DesktopLayout) {
	w := x.NewWriter()
	w.Write4b(val.Orientation)
	w.Write4b(val.Columns)
	w.Write4b(val.Rows)
	w.Write4b(val.StartingCorner)
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_DESKTOP_LAYOUT")
	x.ChangeProperty(c, x.PropModeReplace, rootWin,
		atom, x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_SHOWING_DESKTOP
 */

func GetShowingDesktop(c *x.Conn) GetBooleanCookie {
	rootWin := c.GetDefaultScreen().Root
	atom, _ := c.GetAtom("_NET_SHOWING_DESKTOP")
	cookie := x.GetProperty(c, false, rootWin,
		atom, x.AtomCardinal, 0, 1)
	return GetBooleanCookie(cookie)
}

/**
 *    _NET_WM_NAME
 */

func GetWMName(c *x.Conn, window x.Window) GetUTF8StrCookie {
	atom, _ := c.GetAtom("_NET_WM_NAME")
	cookie := x.GetProperty(c, false, window,
		atom, getAtom(c, "UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func SetWMNameChecked(c *x.Conn, window x.Window, val string) x.VoidCookie {
	atom, _ := c.GetAtom("_NET_WM_NAME")
	atomUTF8String, _ := c.GetAtom("UTF8_STRING")
	return x.ChangePropertyChecked(c, x.PropModeReplace, window,
		atom, atomUTF8String, 8, []byte(val))
}

func SetWMName(c *x.Conn, window x.Window, val string) {
	atom, _ := c.GetAtom("_NET_WM_NAME")
	atomUTF8String, _ := c.GetAtom("UTF8_STRING")
	x.ChangeProperty(c, x.PropModeReplace, window,
		atom, atomUTF8String, 8, []byte(val))
}

/**
 *    _NET_WM_VISIBLE_NAME
 */

func GetWMVisibleName(c *x.Conn, window x.Window) GetUTF8StrCookie {
	atom, _ := c.GetAtom("_NET_WM_VISIBLE_NAME")
	cookie := x.GetProperty(c, false, window,
		atom, getAtom(c, "UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func SetWMVisibleNameChecked(c *x.Conn, window x.Window, val string) x.VoidCookie {
	atom, _ := c.GetAtom("_NET_WM_VISIBLE_NAME")
	atomUTF8String, _ := c.GetAtom("UTF8_STRING")
	return x.ChangePropertyChecked(c, x.PropModeReplace, window,
		atom, atomUTF8String, 8, []byte(val))
}

func SetWMVisibleName(c *x.Conn, window x.Window, val string) {
	atom, _ := c.GetAtom("_NET_WM_VISIBLE_NAME")
	atomUTF8String, _ := c.GetAtom("UTF8_STRING")
	x.ChangeProperty(c, x.PropModeReplace, window,
		atom, atomUTF8String, 8, []byte(val))
}

/**
 *    _NET_WM_ICON_NAME
 */

func GetWMIconName(c *x.Conn, window x.Window) GetUTF8StrCookie {
	atom, _ := c.GetAtom("_NET_WM_ICON_NAME")
	cookie := x.GetProperty(c, false, window,
		atom, getAtom(c, "UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func SetWMIconNameChecked(c *x.Conn, window x.Window, val string) x.VoidCookie {
	atom, _ := c.GetAtom("_NET_WM_ICON_NAME")
	atomUTF8String, _ := c.GetAtom("UTF8_STRING")
	return x.ChangePropertyChecked(c, x.PropModeReplace, window,
		atom, atomUTF8String, 8, []byte(val))
}

func SetWMIconName(c *x.Conn, window x.Window, val string) {
	atom, _ := c.GetAtom("_NET_WM_ICON_NAME")
	atomUTF8String, _ := c.GetAtom("UTF8_STRING")
	x.ChangeProperty(c, x.PropModeReplace, window,
		atom, atomUTF8String, 8, []byte(val))
}

/**
 *    _NET_WM_VISIBLE_ICON_NAME
 */

func GetWMVisibleIconName(c *x.Conn, window x.Window) GetUTF8StrCookie {
	atom, _ := c.GetAtom("_NET_WM_VISIBLE_ICON_NAME")
	cookie := x.GetProperty(c, false, window,
		atom, getAtom(c, "UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func SetWMVisibleIconNameChecked(c *x.Conn, window x.Window, val string) x.VoidCookie {
	atom, _ := c.GetAtom("_NET_WM_VISIBLE_ICON_NAME")
	atomUTF8String, _ := c.GetAtom("UTF8_STRING")
	return x.ChangePropertyChecked(c, x.PropModeReplace, window,
		atom, atomUTF8String, 8, []byte(val))
}

func SetWMVisibleIconName(c *x.Conn, window x.Window, val string) {
	atom, _ := c.GetAtom("_NET_WM_VISIBLE_ICON_NAME")
	atomUTF8String, _ := c.GetAtom("UTF8_STRING")
	x.ChangeProperty(c, x.PropModeReplace, window,
		atom, atomUTF8String, 8, []byte(val))
}

/**
 *    _NET_WM_DESKTOP
 */

func GetWMDesktop(c *x.Conn, window x.Window) GetCardinalCookie {
	atom, _ := c.GetAtom("_NET_WM_DESKTOP")
	cookie := x.GetProperty(c, false, window,
		atom, x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func SetWMDesktopChecked(c *x.Conn, window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	atom, _ := c.GetAtom("_NET_WM_DESKTOP")
	return x.ChangePropertyChecked(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

func SetWMDesktop(c *x.Conn, window x.Window, val uint32) {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	atom, _ := c.GetAtom("_NET_WM_DESKTOP")
	x.ChangeProperty(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_WM_WINDOW_TYPE
 */

func GetWMWindowType(c *x.Conn, window x.Window) GetAtomsCookie {
	atom, _ := c.GetAtom("_NET_WM_WINDOW_TYPE")
	cookie := x.GetProperty(c, false, window,
		atom, x.AtomAtom, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func SetWMWindowTypeChecked(c *x.Conn, window x.Window, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	atom, _ := c.GetAtom("_NET_WM_WINDOW_TYPE")
	return x.ChangePropertyChecked(c, x.PropModeReplace, window,
		atom, x.AtomAtom, 32, w.Bytes())
}

func SetWMWindowType(c *x.Conn, window x.Window, vals []x.Atom) {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	atom, _ := c.GetAtom("_NET_WM_WINDOW_TYPE")
	x.ChangeProperty(c, x.PropModeReplace, window,
		atom, x.AtomAtom, 32, w.Bytes())
}

/**
 *    _NET_WM_STATE
 */

func GetWMState(c *x.Conn, window x.Window) GetAtomsCookie {
	atom, _ := c.GetAtom("_NET_WM_STATE")
	cookie := x.GetProperty(c, false, window,
		atom, x.AtomAtom, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func SetWMStateChecked(c *x.Conn, window x.Window, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	atom, _ := c.GetAtom("_NET_WM_STATE")
	return x.ChangePropertyChecked(c, x.PropModeReplace, window,
		atom, x.AtomAtom, 32, w.Bytes())
}

func SetWMState(c *x.Conn, window x.Window, vals []x.Atom) {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	atom, _ := c.GetAtom("_NET_WM_STATE")
	x.ChangeProperty(c, x.PropModeReplace, window,
		atom, x.AtomAtom, 32, w.Bytes())
}

/**
 *    _NET_WM_ALLOWED_ACTIONS
 */

func GetWMAllowedActions(c *x.Conn, window x.Window) GetAtomsCookie {
	atom, _ := c.GetAtom("_NET_WM_ALLOWED_ACTIONS")
	cookie := x.GetProperty(c, false, window,
		atom, x.AtomAtom, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func SetWMAllowedActionsChecked(c *x.Conn, window x.Window, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	atom, _ := c.GetAtom("_NET_WM_ALLOWED_ACTIONS")
	return x.ChangePropertyChecked(c, x.PropModeReplace, window,
		atom, x.AtomAtom, 32, w.Bytes())
}

func SetWMAllowedActions(c *x.Conn, window x.Window, vals []x.Atom) {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	atom, _ := c.GetAtom("_NET_WM_ALLOWED_ACTIONS")
	x.ChangeProperty(c, x.PropModeReplace, window,
		atom, x.AtomAtom, 32, w.Bytes())
}

/**
 *    _NET_WM_STRUT
 */

func GetWMStrut(c *x.Conn, window x.Window) GetWMStrutCookie {
	atom, _ := c.GetAtom("_NET_WM_STRUT")
	cookie := x.GetProperty(c, false, window,
		atom, x.AtomCardinal, 0, 4)
	return GetWMStrutCookie(cookie)
}

type GetWMStrutCookie x.GetPropertyCookie

func (cookie GetWMStrutCookie) Reply(c *x.Conn) (WMStrut, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return WMStrut{}, err
	}
	if reply.Type != x.AtomCardinal {
		return WMStrut{}, errors.New("bad reply")
	}
	return getWMStrutFromReply(reply)
}

type WMStrut struct {
	Left, Right, Top, Bottom uint32
}

func getWMStrutFromReply(reply *x.GetPropertyReply) (WMStrut, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return WMStrut{}, err
	}

	if len(list) != 4 {
		return WMStrut{}, errors.New("length of list is not 4")
	}
	return WMStrut{
		Left:   list[0],
		Right:  list[1],
		Top:    list[2],
		Bottom: list[3],
	}, nil
}

func SetWMStrutChecked(c *x.Conn, window x.Window, val WMStrut) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	atom, _ := c.GetAtom("_NET_WM_STRUT")
	return x.ChangePropertyChecked(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

func SetWMStrut(c *x.Conn, window x.Window, val WMStrut) {
	w := x.NewWriter()
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	atom, _ := c.GetAtom("_NET_WM_STRUT")
	x.ChangeProperty(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_WM_STRUT_PARTIAL
 */

func GetWMStrutPartial(c *x.Conn, window x.Window) GetWMStrutPartialCookie {
	atom, _ := c.GetAtom("_NET_WM_STRUT_PARTIAL")
	cookie := x.GetProperty(c, false, window,
		atom, x.AtomCardinal, 0, 12)
	return GetWMStrutPartialCookie(cookie)
}

type GetWMStrutPartialCookie x.GetPropertyCookie

func (cookie GetWMStrutPartialCookie) Reply(c *x.Conn) (WMStrutPartial, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return WMStrutPartial{}, err
	}
	if reply.Type != x.AtomCardinal {
		return WMStrutPartial{}, errors.New("bad reply")
	}
	return getWMStrutPartialFromReply(reply)
}

type WMStrutPartial struct {
	Left, Right, Top, Bottom, LeftStartY, LeftEndY, RightStartY, RightEndY, TopStartX, TopEndX, BottomStartX, BottomEndX uint32
}

func getWMStrutPartialFromReply(reply *x.GetPropertyReply) (WMStrutPartial, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return WMStrutPartial{}, err
	}

	if len(list) != 12 {
		return WMStrutPartial{}, errors.New("length of list is not 12")
	}
	return WMStrutPartial{
		Left:         list[0],
		Right:        list[1],
		Top:          list[2],
		Bottom:       list[3],
		LeftStartY:   list[4],
		LeftEndY:     list[5],
		RightStartY:  list[6],
		RightEndY:    list[7],
		TopStartX:    list[8],
		TopEndX:      list[9],
		BottomStartX: list[10],
		BottomEndX:   list[11],
	}, nil
}

func SetWMStrutPartialChecked(c *x.Conn, window x.Window, val WMStrutPartial) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	w.Write4b(val.LeftStartY)
	w.Write4b(val.LeftEndY)
	w.Write4b(val.RightStartY)
	w.Write4b(val.RightEndY)
	w.Write4b(val.TopStartX)
	w.Write4b(val.TopEndX)
	w.Write4b(val.BottomStartX)
	w.Write4b(val.BottomEndX)
	atom, _ := c.GetAtom("_NET_WM_STRUT_PARTIAL")
	return x.ChangePropertyChecked(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

func SetWMStrutPartial(c *x.Conn, window x.Window, val WMStrutPartial) {
	w := x.NewWriter()
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	w.Write4b(val.LeftStartY)
	w.Write4b(val.LeftEndY)
	w.Write4b(val.RightStartY)
	w.Write4b(val.RightEndY)
	w.Write4b(val.TopStartX)
	w.Write4b(val.TopEndX)
	w.Write4b(val.BottomStartX)
	w.Write4b(val.BottomEndX)
	atom, _ := c.GetAtom("_NET_WM_STRUT_PARTIAL")
	x.ChangeProperty(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_WM_ICON_GEOMETRY
 */

func GetWMIconGeometry(c *x.Conn, window x.Window) GetWMIconGeometryCookie {
	atom, _ := c.GetAtom("_NET_WM_ICON_GEOMETRY")
	cookie := x.GetProperty(c, false, window,
		atom, x.AtomCardinal, 0, 4)
	return GetWMIconGeometryCookie(cookie)
}

type GetWMIconGeometryCookie x.GetPropertyCookie

func (cookie GetWMIconGeometryCookie) Reply(c *x.Conn) (WMIconGeometry, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return WMIconGeometry{}, err
	}
	if reply.Type != x.AtomCardinal {
		return WMIconGeometry{}, errors.New("bad reply")
	}
	return getWMIconGeometryFromReply(reply)
}

type WMIconGeometry struct {
	X, Y, Width, Height uint32
}

func getWMIconGeometryFromReply(reply *x.GetPropertyReply) (WMIconGeometry, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return WMIconGeometry{}, err
	}

	if len(list) != 4 {
		return WMIconGeometry{}, errors.New("length of list is not 4")
	}
	return WMIconGeometry{
		X:      list[0],
		Y:      list[1],
		Width:  list[2],
		Height: list[3],
	}, nil
}

func SetWMIconGeometryChecked(c *x.Conn, window x.Window, val WMIconGeometry) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.X)
	w.Write4b(val.Y)
	w.Write4b(val.Width)
	w.Write4b(val.Height)
	atom, _ := c.GetAtom("_NET_WM_ICON_GEOMETRY")
	return x.ChangePropertyChecked(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

func SetWMIconGeometry(c *x.Conn, window x.Window, val WMIconGeometry) {
	w := x.NewWriter()
	w.Write4b(val.X)
	w.Write4b(val.Y)
	w.Write4b(val.Width)
	w.Write4b(val.Height)
	atom, _ := c.GetAtom("_NET_WM_ICON_GEOMETRY")
	x.ChangeProperty(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_WM_ICON
 */

func GetWMIcon(c *x.Conn, window x.Window) GetWMIconCookie {
	atom, _ := c.GetAtom("_NET_WM_ICON")
	cookie := x.GetProperty(c, false, window,
		atom, x.AtomCardinal, 0, LENGTH_MAX)
	return GetWMIconCookie(cookie)
}

type GetWMIconCookie x.GetPropertyCookie

func (cookie GetWMIconCookie) Reply(c *x.Conn) ([]WMIcon, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return nil, err
	}
	if reply.Type != x.AtomCardinal {
		return nil, errors.New("bad reply")
	}
	return getWMIconFromReply(reply)
}

/**
 *    _NET_WM_PID
 */

func GetWMPid(c *x.Conn, window x.Window) GetCardinalCookie {
	atom, _ := c.GetAtom("_NET_WM_PID")
	cookie := x.GetProperty(c, false, window,
		atom, x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func SetWMPidChecked(c *x.Conn, window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	atom, _ := c.GetAtom("_NET_WM_PID")
	return x.ChangePropertyChecked(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

func SetWMPid(c *x.Conn, window x.Window, val uint32) {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	atom, _ := c.GetAtom("_NET_WM_PID")
	x.ChangeProperty(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_WM_HANDLED_ICONS
 */

func GetWMHandledIcons(c *x.Conn, window x.Window) GetCardinalCookie {
	atom, _ := c.GetAtom("_NET_WM_HANDLED_ICONS")
	cookie := x.GetProperty(c, false, window,
		atom, x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func SetWMHandledIconsChecked(c *x.Conn, window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	atom, _ := c.GetAtom("_NET_WM_HANDLED_ICONS")
	return x.ChangePropertyChecked(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

func SetWMHandledIcons(c *x.Conn, window x.Window, val uint32) {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	atom, _ := c.GetAtom("_NET_WM_HANDLED_ICONS")
	x.ChangeProperty(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_WM_USER_TIME
 */

func GetWMUserTime(c *x.Conn, window x.Window) GetCardinalCookie {
	atom, _ := c.GetAtom("_NET_WM_USER_TIME")
	cookie := x.GetProperty(c, false, window,
		atom, x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func SetWMUserTimeChecked(c *x.Conn, window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	atom, _ := c.GetAtom("_NET_WM_USER_TIME")
	return x.ChangePropertyChecked(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

func SetWMUserTime(c *x.Conn, window x.Window, val uint32) {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	atom, _ := c.GetAtom("_NET_WM_USER_TIME")
	x.ChangeProperty(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_WM_USER_TIME_WINDOW
 */

func GetWMUserTimeWindow(c *x.Conn, window x.Window) GetWindowCookie {
	atom, _ := c.GetAtom("_NET_WM_USER_TIME_WINDOW")
	cookie := x.GetProperty(c, false, window,
		atom, x.AtomWindow, 0, 1)
	return GetWindowCookie(cookie)
}

func SetWMUserTimeWindowChecked(c *x.Conn, window x.Window, val x.Window) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	atom, _ := c.GetAtom("_NET_WM_USER_TIME_WINDOW")
	return x.ChangePropertyChecked(c, x.PropModeReplace, window,
		atom, x.AtomWindow, 32, w.Bytes())
}

func SetWMUserTimeWindow(c *x.Conn, window x.Window, val x.Window) {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	atom, _ := c.GetAtom("_NET_WM_USER_TIME_WINDOW")
	x.ChangeProperty(c, x.PropModeReplace, window,
		atom, x.AtomWindow, 32, w.Bytes())
}

/**
 *    _NET_FRAME_EXTENTS
 */

func GetFrameExtents(c *x.Conn, window x.Window) GetFrameExtentsCookie {
	atom, _ := c.GetAtom("_NET_FRAME_EXTENTS")
	cookie := x.GetProperty(c, false, window,
		atom, x.AtomCardinal, 0, 4)
	return GetFrameExtentsCookie(cookie)
}

type GetFrameExtentsCookie x.GetPropertyCookie

func (cookie GetFrameExtentsCookie) Reply(c *x.Conn) (FrameExtents, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return FrameExtents{}, err
	}
	if reply.Type != x.AtomCardinal {
		return FrameExtents{}, errors.New("bad reply")
	}
	return getFrameExtentsFromReply(reply)
}

type FrameExtents struct {
	Left, Right, Top, Bottom uint32
}

func getFrameExtentsFromReply(reply *x.GetPropertyReply) (FrameExtents, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return FrameExtents{}, err
	}

	if len(list) != 4 {
		return FrameExtents{}, errors.New("length of list is not 4")
	}
	return FrameExtents{
		Left:   list[0],
		Right:  list[1],
		Top:    list[2],
		Bottom: list[3],
	}, nil
}

func SetFrameExtentsChecked(c *x.Conn, window x.Window, val FrameExtents) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	atom, _ := c.GetAtom("_NET_FRAME_EXTENTS")
	return x.ChangePropertyChecked(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

func SetFrameExtents(c *x.Conn, window x.Window, val FrameExtents) {
	w := x.NewWriter()
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	atom, _ := c.GetAtom("_NET_FRAME_EXTENTS")
	x.ChangeProperty(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_WM_SYNC_REQUEST_COUNTER
 */

func GetWMSyncRequestCounter(c *x.Conn, window x.Window) GetWMSyncRequestCounterCookie {
	atom, _ := c.GetAtom("_NET_WM_SYNC_REQUEST_COUNTER")
	cookie := x.GetProperty(c, false, window,
		atom, x.AtomCardinal, 0, 2)
	return GetWMSyncRequestCounterCookie(cookie)
}

type GetWMSyncRequestCounterCookie x.GetPropertyCookie

func (cookie GetWMSyncRequestCounterCookie) Reply(c *x.Conn) (WMSyncRequestCounter, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return WMSyncRequestCounter{}, err
	}
	if reply.Type != x.AtomCardinal {
		return WMSyncRequestCounter{}, errors.New("bad reply")
	}
	return getWMSyncRequestCounterFromReply(reply)
}

type WMSyncRequestCounter struct {
	Low, High uint32
}

func getWMSyncRequestCounterFromReply(reply *x.GetPropertyReply) (WMSyncRequestCounter, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return WMSyncRequestCounter{}, err
	}

	if len(list) != 2 {
		return WMSyncRequestCounter{}, errors.New("length of list is not 2")
	}
	return WMSyncRequestCounter{
		Low:  list[0],
		High: list[1],
	}, nil
}

func SetWMSyncRequestCounterChecked(c *x.Conn, window x.Window, val WMSyncRequestCounter) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Low)
	w.Write4b(val.High)
	atom, _ := c.GetAtom("_NET_WM_SYNC_REQUEST_COUNTER")
	return x.ChangePropertyChecked(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

func SetWMSyncRequestCounter(c *x.Conn, window x.Window, val WMSyncRequestCounter) {
	w := x.NewWriter()
	w.Write4b(val.Low)
	w.Write4b(val.High)
	atom, _ := c.GetAtom("_NET_WM_SYNC_REQUEST_COUNTER")
	x.ChangeProperty(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_WM_FULLSCREEN_MONITORS
 */

func GetWMFullscreenMonitors(c *x.Conn, window x.Window) GetWMFullscreenMonitorsCookie {
	atom, _ := c.GetAtom("_NET_WM_FULLSCREEN_MONITORS")
	cookie := x.GetProperty(c, false, window,
		atom, x.AtomCardinal, 0, 4)
	return GetWMFullscreenMonitorsCookie(cookie)
}

type GetWMFullscreenMonitorsCookie x.GetPropertyCookie

func (cookie GetWMFullscreenMonitorsCookie) Reply(c *x.Conn) (WMFullscreenMonitors, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c)
	if err != nil {
		return WMFullscreenMonitors{}, err
	}
	if reply.Type != x.AtomCardinal {
		return WMFullscreenMonitors{}, errors.New("bad reply")
	}
	return getWMFullscreenMonitorsFromReply(reply)
}

type WMFullscreenMonitors struct {
	Top, Bottom, Left, Right uint32
}

func getWMFullscreenMonitorsFromReply(reply *x.GetPropertyReply) (WMFullscreenMonitors, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return WMFullscreenMonitors{}, err
	}

	if len(list) != 4 {
		return WMFullscreenMonitors{}, errors.New("length of list is not 4")
	}
	return WMFullscreenMonitors{
		Top:    list[0],
		Bottom: list[1],
		Left:   list[2],
		Right:  list[3],
	}, nil
}

func SetWMFullscreenMonitorsChecked(c *x.Conn, window x.Window, val WMFullscreenMonitors) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	atom, _ := c.GetAtom("_NET_WM_FULLSCREEN_MONITORS")
	return x.ChangePropertyChecked(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}

func SetWMFullscreenMonitors(c *x.Conn, window x.Window, val WMFullscreenMonitors) {
	w := x.NewWriter()
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	atom, _ := c.GetAtom("_NET_WM_FULLSCREEN_MONITORS")
	x.ChangeProperty(c, x.PropModeReplace, window,
		atom, x.AtomCardinal, 32, w.Bytes())
}
