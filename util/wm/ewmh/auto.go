package ewmh

import x "github.com/linuxdeepin/go-x11-client"
import "errors"

type GetUTF8StrCookie x.GetPropertyCookie

func (cookie GetUTF8StrCookie) Reply(c *Conn) (string, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return "", err
	}
	if reply.Type != c.GetAtom("UTF8_STRING") {
		return "", errors.New("bad reply")
	}
	return getUTF8StrFromReply(reply)
}

type GetUTF8StrsCookie x.GetPropertyCookie

func (cookie GetUTF8StrsCookie) Reply(c *Conn) ([]string, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return nil, err
	}
	if reply.Type != c.GetAtom("UTF8_STRING") {
		return nil, errors.New("bad reply")
	}
	return getUTF8StrsFromReply(reply)
}

type GetBooleanCookie x.GetPropertyCookie

func (cookie GetBooleanCookie) Reply(c *Conn) (bool, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
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

func (cookie GetWindowCookie) Reply(c *Conn) (x.Window, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
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

func (cookie GetCardinalCookie) Reply(c *Conn) (uint32, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
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

func (cookie GetWindowsCookie) Reply(c *Conn) ([]x.Window, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
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

func (cookie GetAtomsCookie) Reply(c *Conn) ([]x.Atom, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
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

func (cookie GetCardinalsCookie) Reply(c *Conn) ([]uint32, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
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

func (c *Conn) GetSupported() GetAtomsCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomAtom, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func (c *Conn) GetSupportedUnchecked() GetAtomsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomAtom, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func (c *Conn) SetSupportedChecked(vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomAtom, 32, uint32(len(vals)), w.Bytes())
}

func (c *Conn) SetSupported(vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomAtom, 32, uint32(len(vals)), w.Bytes())
}

/**
 *    _NET_CLIENT_LIST
 */

func (c *Conn) GetClientList() GetWindowsCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_CLIENT_LIST"), x.AtomWindow, 0, LENGTH_MAX)
	return GetWindowsCookie(cookie)
}

func (c *Conn) GetClientListUnchecked() GetWindowsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_CLIENT_LIST"), x.AtomWindow, 0, LENGTH_MAX)
	return GetWindowsCookie(cookie)
}

func (c *Conn) SetClientListChecked(vals []x.Window) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomWindow, 32, uint32(len(vals)), w.Bytes())
}

func (c *Conn) SetClientList(vals []x.Window) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomWindow, 32, uint32(len(vals)), w.Bytes())
}

/**
 *    _NET_CLIENT_LIST_STACKING
 */

func (c *Conn) GetClientListStacking() GetWindowsCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_CLIENT_LIST_STACKING"), x.AtomWindow, 0, LENGTH_MAX)
	return GetWindowsCookie(cookie)
}

func (c *Conn) GetClientListStackingUnchecked() GetWindowsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_CLIENT_LIST_STACKING"), x.AtomWindow, 0, LENGTH_MAX)
	return GetWindowsCookie(cookie)
}

func (c *Conn) SetClientListStackingChecked(vals []x.Window) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomWindow, 32, uint32(len(vals)), w.Bytes())
}

func (c *Conn) SetClientListStacking(vals []x.Window) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomWindow, 32, uint32(len(vals)), w.Bytes())
}

/**
 *    _NET_NUMBER_OF_DESKTOPS
 */

func (c *Conn) GetNumberOfDesktop() GetCardinalCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_NUMBER_OF_DESKTOPS"), x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func (c *Conn) GetNumberOfDesktopUnchecked() GetCardinalCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_NUMBER_OF_DESKTOPS"), x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func (c *Conn) SetNumberOfDesktopChecked(val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_NUMBER_OF_DESKTOPS"), x.AtomCardinal, 32, 1, w.Bytes())
}

func (c *Conn) SetNumberOfDesktop(val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_NUMBER_OF_DESKTOPS"), x.AtomCardinal, 32, 1, w.Bytes())
}

/**
 *    _NET_DESKTOP_GEOMETRY
 */

func (c *Conn) GetDesktopGeometry() GetDesktopGeometryCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_GEOMETRY"), x.AtomCardinal, 0, 2)
	return GetDesktopGeometryCookie(cookie)
}

func (c *Conn) GetDesktopGeometryUnchecked() GetDesktopGeometryCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_GEOMETRY"), x.AtomCardinal, 0, 2)
	return GetDesktopGeometryCookie(cookie)
}

type GetDesktopGeometryCookie x.GetPropertyCookie

func (cookie GetDesktopGeometryCookie) Reply(c *Conn) (DesktopGeometry, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
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

func (c *Conn) SetDesktopGeometryChecked(val DesktopGeometry) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Width)
	w.Write4b(val.Height)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_GEOMETRY"), x.AtomCardinal, 32, 2, w.Bytes())
}

func (c *Conn) SetDesktopGeometry(val DesktopGeometry) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Width)
	w.Write4b(val.Height)
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_GEOMETRY"), x.AtomCardinal, 32, 2, w.Bytes())
}

/**
 *    _NET_DESKTOP_VIEWPORT
 */

func (c *Conn) GetDesktopViewport() GetDesktopViewportCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_VIEWPORT"), x.AtomCardinal, 0, LENGTH_MAX)
	return GetDesktopViewportCookie(cookie)
}

func (c *Conn) GetDesktopViewportUnchecked() GetDesktopViewportCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_VIEWPORT"), x.AtomCardinal, 0, LENGTH_MAX)
	return GetDesktopViewportCookie(cookie)
}

type GetDesktopViewportCookie x.GetPropertyCookie

func (cookie GetDesktopViewportCookie) Reply(c *Conn) ([]ViewportLeftTopCorner, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
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

func (c *Conn) SetDesktopViewportChecked(vals []ViewportLeftTopCorner) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(val.X)
		w.Write4b(val.Y)
	}
	length := uint32(len(vals) * 2)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_VIEWPORT"), x.AtomCardinal, 32, length, w.Bytes())
}

func (c *Conn) SetDesktopViewport(vals []ViewportLeftTopCorner) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(val.X)
		w.Write4b(val.Y)
	}
	length := uint32(len(vals) * 2)
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_VIEWPORT"), x.AtomCardinal, 32, length, w.Bytes())
}

/**
 *    _NET_CURRENT_DESKTOP
 */

func (c *Conn) GetCurrentDesktop() GetCardinalCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_CURRENT_DESKTOP"), x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func (c *Conn) GetCurrentDesktopUnchecked() GetCardinalCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_CURRENT_DESKTOP"), x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func (c *Conn) SetCurrentDesktopChecked(val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_CURRENT_DESKTOP"), x.AtomCardinal, 32, 1, w.Bytes())
}

func (c *Conn) SetCurrentDesktop(val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_CURRENT_DESKTOP"), x.AtomCardinal, 32, 1, w.Bytes())
}

/**
 *    _NET_DESKTOP_NAMES
 */

func (c *Conn) GetDesktopNames() GetUTF8StrsCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_NAMES"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrsCookie(cookie)
}

func (c *Conn) GetDesktopNamesUnchecked() GetUTF8StrsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_NAMES"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrsCookie(cookie)
}

func (c *Conn) SetDesktopNamesChecked(vals []string) x.VoidCookie {
	w := x.NewWriter()
	length := 0
	for _, val := range vals {
		w.WriteString(val)
		w.Write1b(0)
		length += len(val) + 1
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_NAMES"), c.GetAtom("UTF8_STRING"), 8, uint32(length), w.Bytes())
}

func (c *Conn) SetDesktopNames(vals []string) x.VoidCookie {
	w := x.NewWriter()
	length := 0
	for _, val := range vals {
		w.WriteString(val)
		w.Write1b(0)
		length += len(val) + 1
	}
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_NAMES"), c.GetAtom("UTF8_STRING"), 8, uint32(length), w.Bytes())
}

/**
 *    _NET_ACTIVE_WINDOW
 */

func (c *Conn) GetActiveWindow() GetWindowCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_ACTIVE_WINDOW"), x.AtomWindow, 0, 1)
	return GetWindowCookie(cookie)
}

func (c *Conn) GetActiveWindowUnchecked() GetWindowCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_ACTIVE_WINDOW"), x.AtomWindow, 0, 1)
	return GetWindowCookie(cookie)
}

func (c *Conn) SetActiveWindowChecked(val x.Window) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_ACTIVE_WINDOW"), x.AtomWindow, 32, 1, w.Bytes())
}

func (c *Conn) SetActiveWindow(val x.Window) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_ACTIVE_WINDOW"), x.AtomWindow, 32, 1, w.Bytes())
}

/**
 *    _NET_WORKAREA
 */

func (c *Conn) GetWorkarea() GetWorkareaCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_WORKAREA"), x.AtomCardinal, 0, LENGTH_MAX)
	return GetWorkareaCookie(cookie)
}

func (c *Conn) GetWorkareaUnchecked() GetWorkareaCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_WORKAREA"), x.AtomCardinal, 0, LENGTH_MAX)
	return GetWorkareaCookie(cookie)
}

type GetWorkareaCookie x.GetPropertyCookie

func (cookie GetWorkareaCookie) Reply(c *Conn) ([]WorkareaGeometry, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
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

func (c *Conn) SetWorkareaChecked(vals []WorkareaGeometry) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(val.X)
		w.Write4b(val.Y)
		w.Write4b(val.Width)
		w.Write4b(val.Height)
	}
	length := uint32(len(vals) * 4)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_WORKAREA"), x.AtomCardinal, 32, length, w.Bytes())
}

func (c *Conn) SetWorkarea(vals []WorkareaGeometry) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(val.X)
		w.Write4b(val.Y)
		w.Write4b(val.Width)
		w.Write4b(val.Height)
	}
	length := uint32(len(vals) * 4)
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_WORKAREA"), x.AtomCardinal, 32, length, w.Bytes())
}

/**
 *    _NET_SUPPORTING_WM_CHECK
 */

func (c *Conn) GetSupportingWmCheck(window x.Window) GetWindowCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_SUPPORTING_WM_CHECK"), x.AtomWindow, 0, 1)
	return GetWindowCookie(cookie)
}

func (c *Conn) GetSupportingWmCheckUnchecked(window x.Window) GetWindowCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_SUPPORTING_WM_CHECK"), x.AtomWindow, 0, 1)
	return GetWindowCookie(cookie)
}

func (c *Conn) SetSupportingWmCheckChecked(window x.Window, val x.Window) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_SUPPORTING_WM_CHECK"), x.AtomWindow, 32, 1, w.Bytes())
}

func (c *Conn) SetSupportingWmCheck(window x.Window, val x.Window) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_SUPPORTING_WM_CHECK"), x.AtomWindow, 32, 1, w.Bytes())
}

/**
 *    _NET_VIRTUAL_ROOTS
 */

func (c *Conn) GetVirtualRoots() GetWindowsCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_VIRTUAL_ROOTS"), x.AtomWindow, 0, LENGTH_MAX)
	return GetWindowsCookie(cookie)
}

func (c *Conn) GetVirtualRootsUnchecked() GetWindowsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_VIRTUAL_ROOTS"), x.AtomWindow, 0, LENGTH_MAX)
	return GetWindowsCookie(cookie)
}

func (c *Conn) SetVirtualRootsChecked(vals []x.Window) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomWindow, 32, uint32(len(vals)), w.Bytes())
}

func (c *Conn) SetVirtualRoots(vals []x.Window) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomWindow, 32, uint32(len(vals)), w.Bytes())
}

/**
 *    _NET_DESKTOP_LAYOUT
 */

func (c *Conn) GetDesktopLayout() GetDesktopLayoutCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_LAYOUT"), x.AtomCardinal, 0, 4)
	return GetDesktopLayoutCookie(cookie)
}

func (c *Conn) GetDesktopLayoutUnchecked() GetDesktopLayoutCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_LAYOUT"), x.AtomCardinal, 0, 4)
	return GetDesktopLayoutCookie(cookie)
}

type GetDesktopLayoutCookie x.GetPropertyCookie

func (cookie GetDesktopLayoutCookie) Reply(c *Conn) (DesktopLayout, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
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

func (c *Conn) SetDesktopLayoutChecked(val DesktopLayout) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Orientation)
	w.Write4b(val.Columns)
	w.Write4b(val.Rows)
	w.Write4b(val.StartingCorner)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_LAYOUT"), x.AtomCardinal, 32, 4, w.Bytes())
}

func (c *Conn) SetDesktopLayout(val DesktopLayout) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Orientation)
	w.Write4b(val.Columns)
	w.Write4b(val.Rows)
	w.Write4b(val.StartingCorner)
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_LAYOUT"), x.AtomCardinal, 32, 4, w.Bytes())
}

/**
 *    _NET_SHOWING_DESKTOP
 */

func (c *Conn) GetShowingDesktop() GetBooleanCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_SHOWING_DESKTOP"), x.AtomCardinal, 0, 1)
	return GetBooleanCookie(cookie)
}

func (c *Conn) GetShowingDesktopUnchecked() GetBooleanCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_SHOWING_DESKTOP"), x.AtomCardinal, 0, 1)
	return GetBooleanCookie(cookie)
}

/**
 *    _NET_WM_NAME
 */

func (c *Conn) GetWmName(window x.Window) GetUTF8StrCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func (c *Conn) GetWmNameUnchecked(window x.Window) GetUTF8StrCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func (c *Conn) SetWmNameChecked(window x.Window, val string) x.VoidCookie {
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_NAME"), c.GetAtom("UTF8_STRING"), 8, uint32(len(val)), []byte(val))
}

func (c *Conn) SetWmName(window x.Window, val string) x.VoidCookie {
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_NAME"), c.GetAtom("UTF8_STRING"), 8, uint32(len(val)), []byte(val))
}

/**
 *    _NET_WM_VISIBLE_NAME
 */

func (c *Conn) GetWmVisibleName(window x.Window) GetUTF8StrCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_VISIBLE_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func (c *Conn) GetWmVisibleNameUnchecked(window x.Window) GetUTF8StrCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_VISIBLE_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func (c *Conn) SetWmVisibleNameChecked(window x.Window, val string) x.VoidCookie {
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_VISIBLE_NAME"), c.GetAtom("UTF8_STRING"), 8, uint32(len(val)), []byte(val))
}

func (c *Conn) SetWmVisibleName(window x.Window, val string) x.VoidCookie {
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_VISIBLE_NAME"), c.GetAtom("UTF8_STRING"), 8, uint32(len(val)), []byte(val))
}

/**
 *    _NET_WM_ICON_NAME
 */

func (c *Conn) GetWmIconName(window x.Window) GetUTF8StrCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_ICON_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func (c *Conn) GetWmIconNameUnchecked(window x.Window) GetUTF8StrCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_ICON_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func (c *Conn) SetWmIconNameChecked(window x.Window, val string) x.VoidCookie {
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_ICON_NAME"), c.GetAtom("UTF8_STRING"), 8, uint32(len(val)), []byte(val))
}

func (c *Conn) SetWmIconName(window x.Window, val string) x.VoidCookie {
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_ICON_NAME"), c.GetAtom("UTF8_STRING"), 8, uint32(len(val)), []byte(val))
}

/**
 *    _NET_WM_VISIBLE_ICON_NAME
 */

func (c *Conn) GetWmVisibleIconName(window x.Window) GetUTF8StrCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_VISIBLE_ICON_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func (c *Conn) GetWmVisibleIconNameUnchecked(window x.Window) GetUTF8StrCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_VISIBLE_ICON_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func (c *Conn) SetWmVisibleIconNameChecked(window x.Window, val string) x.VoidCookie {
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_VISIBLE_ICON_NAME"), c.GetAtom("UTF8_STRING"), 8, uint32(len(val)), []byte(val))
}

func (c *Conn) SetWmVisibleIconName(window x.Window, val string) x.VoidCookie {
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_VISIBLE_ICON_NAME"), c.GetAtom("UTF8_STRING"), 8, uint32(len(val)), []byte(val))
}

/**
 *    _NET_WM_DESKTOP
 */

func (c *Conn) GetWmDesktop(window x.Window) GetCardinalCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_DESKTOP"), x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func (c *Conn) GetWmDesktopUnchecked(window x.Window) GetCardinalCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_DESKTOP"), x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func (c *Conn) SetWmDesktopChecked(window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_DESKTOP"), x.AtomCardinal, 32, 1, w.Bytes())
}

func (c *Conn) SetWmDesktop(window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_DESKTOP"), x.AtomCardinal, 32, 1, w.Bytes())
}

/**
 *    _NET_WM_WINDOW_TYPE
 */

func (c *Conn) GetWmWindowType(window x.Window) GetAtomsCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_WINDOW_TYPE"), x.AtomAtom, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func (c *Conn) GetWmWindowTypeUnchecked(window x.Window) GetAtomsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_WINDOW_TYPE"), x.AtomAtom, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func (c *Conn) SetWmWindowTypeChecked(window x.Window, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_SUPPORTED"), x.AtomAtom, 32, uint32(len(vals)), w.Bytes())
}

func (c *Conn) SetWmWindowType(window x.Window, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_SUPPORTED"), x.AtomAtom, 32, uint32(len(vals)), w.Bytes())
}

/**
 *    _NET_WM_STATE
 */

func (c *Conn) GetWmState(window x.Window) GetAtomsCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_STATE"), x.AtomAtom, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func (c *Conn) GetWmStateUnchecked(window x.Window) GetAtomsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_STATE"), x.AtomAtom, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func (c *Conn) SetWmStateChecked(window x.Window, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_SUPPORTED"), x.AtomAtom, 32, uint32(len(vals)), w.Bytes())
}

func (c *Conn) SetWmState(window x.Window, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_SUPPORTED"), x.AtomAtom, 32, uint32(len(vals)), w.Bytes())
}

/**
 *    _NET_WM_ALLOWED_ACTIONS
 */

func (c *Conn) GetWmAllowedActions(window x.Window) GetAtomsCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_ALLOWED_ACTIONS"), x.AtomAtom, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func (c *Conn) GetWmAllowedActionsUnchecked(window x.Window) GetAtomsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_ALLOWED_ACTIONS"), x.AtomAtom, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func (c *Conn) SetWmAllowedActionsChecked(window x.Window, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_SUPPORTED"), x.AtomAtom, 32, uint32(len(vals)), w.Bytes())
}

func (c *Conn) SetWmAllowedActions(window x.Window, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_SUPPORTED"), x.AtomAtom, 32, uint32(len(vals)), w.Bytes())
}

/**
 *    _NET_WM_STRUT
 */

func (c *Conn) GetWmStrut(window x.Window) GetWmStrutCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_STRUT"), x.AtomCardinal, 0, 4)
	return GetWmStrutCookie(cookie)
}

func (c *Conn) GetWmStrutUnchecked(window x.Window) GetWmStrutCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_STRUT"), x.AtomCardinal, 0, 4)
	return GetWmStrutCookie(cookie)
}

type GetWmStrutCookie x.GetPropertyCookie

func (cookie GetWmStrutCookie) Reply(c *Conn) (WmStrut, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return WmStrut{}, err
	}
	if reply.Type != x.AtomCardinal {
		return WmStrut{}, errors.New("bad reply")
	}
	return getWmStrutFromReply(reply)
}

type WmStrut struct {
	Left, Right, Top, Bottom uint32
}

func getWmStrutFromReply(reply *x.GetPropertyReply) (WmStrut, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return WmStrut{}, err
	}

	if len(list) != 4 {
		return WmStrut{}, errors.New("length of list is not 4")
	}
	return WmStrut{
		Left:   list[0],
		Right:  list[1],
		Top:    list[2],
		Bottom: list[3],
	}, nil
}

func (c *Conn) SetWmStrutChecked(window x.Window, val WmStrut) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_STRUT"), x.AtomCardinal, 32, 4, w.Bytes())
}

func (c *Conn) SetWmStrut(window x.Window, val WmStrut) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_STRUT"), x.AtomCardinal, 32, 4, w.Bytes())
}

/**
 *    _NET_WM_STRUT_PARTIAL
 */

func (c *Conn) GetWmStrutPartial(window x.Window) GetWmStrutPartialCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_STRUT_PARTIAL"), x.AtomCardinal, 0, 12)
	return GetWmStrutPartialCookie(cookie)
}

func (c *Conn) GetWmStrutPartialUnchecked(window x.Window) GetWmStrutPartialCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_STRUT_PARTIAL"), x.AtomCardinal, 0, 12)
	return GetWmStrutPartialCookie(cookie)
}

type GetWmStrutPartialCookie x.GetPropertyCookie

func (cookie GetWmStrutPartialCookie) Reply(c *Conn) (WmStrutPartial, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return WmStrutPartial{}, err
	}
	if reply.Type != x.AtomCardinal {
		return WmStrutPartial{}, errors.New("bad reply")
	}
	return getWmStrutPartialFromReply(reply)
}

type WmStrutPartial struct {
	Left, Right, Top, Bottom, LeftStartY, LeftEndY, RightStartY, RightEndY, TopStartX, TopEndX, BottomStartX, BottomEndX uint32
}

func getWmStrutPartialFromReply(reply *x.GetPropertyReply) (WmStrutPartial, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return WmStrutPartial{}, err
	}

	if len(list) != 12 {
		return WmStrutPartial{}, errors.New("length of list is not 12")
	}
	return WmStrutPartial{
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

func (c *Conn) SetWmStrutPartialChecked(window x.Window, val WmStrutPartial) x.VoidCookie {
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
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_STRUT_PARTIAL"), x.AtomCardinal, 32, 12, w.Bytes())
}

func (c *Conn) SetWmStrutPartial(window x.Window, val WmStrutPartial) x.VoidCookie {
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
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_STRUT_PARTIAL"), x.AtomCardinal, 32, 12, w.Bytes())
}

/**
 *    _NET_WM_ICON_GEOMETRY
 */

func (c *Conn) GetWmIconGeometry(window x.Window) GetWmIconGeometryCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_ICON_GEOMETRY"), x.AtomCardinal, 0, 4)
	return GetWmIconGeometryCookie(cookie)
}

func (c *Conn) GetWmIconGeometryUnchecked(window x.Window) GetWmIconGeometryCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_ICON_GEOMETRY"), x.AtomCardinal, 0, 4)
	return GetWmIconGeometryCookie(cookie)
}

type GetWmIconGeometryCookie x.GetPropertyCookie

func (cookie GetWmIconGeometryCookie) Reply(c *Conn) (WmIconGeometry, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return WmIconGeometry{}, err
	}
	if reply.Type != x.AtomCardinal {
		return WmIconGeometry{}, errors.New("bad reply")
	}
	return getWmIconGeometryFromReply(reply)
}

type WmIconGeometry struct {
	X, Y, Width, Height uint32
}

func getWmIconGeometryFromReply(reply *x.GetPropertyReply) (WmIconGeometry, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return WmIconGeometry{}, err
	}

	if len(list) != 4 {
		return WmIconGeometry{}, errors.New("length of list is not 4")
	}
	return WmIconGeometry{
		X:      list[0],
		Y:      list[1],
		Width:  list[2],
		Height: list[3],
	}, nil
}

func (c *Conn) SetWmIconGeometryChecked(window x.Window, val WmIconGeometry) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.X)
	w.Write4b(val.Y)
	w.Write4b(val.Width)
	w.Write4b(val.Height)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_ICON_GEOMETRY"), x.AtomCardinal, 32, 4, w.Bytes())
}

func (c *Conn) SetWmIconGeometry(window x.Window, val WmIconGeometry) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.X)
	w.Write4b(val.Y)
	w.Write4b(val.Width)
	w.Write4b(val.Height)
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_ICON_GEOMETRY"), x.AtomCardinal, 32, 4, w.Bytes())
}

/**
 *    _NET_WM_ICON
 */

func (c *Conn) GetWmIcon(window x.Window) GetWmIconCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_ICON"), x.AtomCardinal, 0, LENGTH_MAX)
	return GetWmIconCookie(cookie)
}

func (c *Conn) GetWmIconUnchecked(window x.Window) GetWmIconCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_ICON"), x.AtomCardinal, 0, LENGTH_MAX)
	return GetWmIconCookie(cookie)
}

type GetWmIconCookie x.GetPropertyCookie

func (cookie GetWmIconCookie) Reply(c *Conn) ([]WmIcon, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return nil, err
	}
	if reply.Type != x.AtomCardinal {
		return nil, errors.New("bad reply")
	}
	return getWmIconFromReply(reply)
}

/**
 *    _NET_WM_PID
 */

func (c *Conn) GetWmPid(window x.Window) GetCardinalCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_PID"), x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func (c *Conn) GetWmPidUnchecked(window x.Window) GetCardinalCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_PID"), x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func (c *Conn) SetWmPidChecked(window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_PID"), x.AtomCardinal, 32, 1, w.Bytes())
}

func (c *Conn) SetWmPid(window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_PID"), x.AtomCardinal, 32, 1, w.Bytes())
}

/**
 *    _NET_WM_HANDLED_ICONS
 */

func (c *Conn) GetWmHandledIcons(window x.Window) GetCardinalCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_HANDLED_ICONS"), x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func (c *Conn) GetWmHandledIconsUnchecked(window x.Window) GetCardinalCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_HANDLED_ICONS"), x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func (c *Conn) SetWmHandledIconsChecked(window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_HANDLED_ICONS"), x.AtomCardinal, 32, 1, w.Bytes())
}

func (c *Conn) SetWmHandledIcons(window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_HANDLED_ICONS"), x.AtomCardinal, 32, 1, w.Bytes())
}

/**
 *    _NET_WM_USER_TIME
 */

func (c *Conn) GetWmUserTime(window x.Window) GetCardinalCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_USER_TIME"), x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func (c *Conn) GetWmUserTimeUnchecked(window x.Window) GetCardinalCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_USER_TIME"), x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func (c *Conn) SetWmUserTimeChecked(window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_USER_TIME"), x.AtomCardinal, 32, 1, w.Bytes())
}

func (c *Conn) SetWmUserTime(window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_USER_TIME"), x.AtomCardinal, 32, 1, w.Bytes())
}

/**
 *    _NET_WM_USER_TIME_WINDOW
 */

func (c *Conn) GetWmUserTimeWindow(window x.Window) GetWindowCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_USER_TIME_WINDOW"), x.AtomWindow, 0, 1)
	return GetWindowCookie(cookie)
}

func (c *Conn) GetWmUserTimeWindowUnchecked(window x.Window) GetWindowCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_USER_TIME_WINDOW"), x.AtomWindow, 0, 1)
	return GetWindowCookie(cookie)
}

func (c *Conn) SetWmUserTimeWindowChecked(window x.Window, val x.Window) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_USER_TIME_WINDOW"), x.AtomWindow, 32, 1, w.Bytes())
}

func (c *Conn) SetWmUserTimeWindow(window x.Window, val x.Window) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_USER_TIME_WINDOW"), x.AtomWindow, 32, 1, w.Bytes())
}

/**
 *    _NET_FRAME_EXTENTS
 */

func (c *Conn) GetFrameExtents(window x.Window) GetFrameExtentsCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_FRAME_EXTENTS"), x.AtomCardinal, 0, 4)
	return GetFrameExtentsCookie(cookie)
}

func (c *Conn) GetFrameExtentsUnchecked(window x.Window) GetFrameExtentsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_FRAME_EXTENTS"), x.AtomCardinal, 0, 4)
	return GetFrameExtentsCookie(cookie)
}

type GetFrameExtentsCookie x.GetPropertyCookie

func (cookie GetFrameExtentsCookie) Reply(c *Conn) (FrameExtents, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
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

func (c *Conn) SetFrameExtentsChecked(window x.Window, val FrameExtents) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_FRAME_EXTENTS"), x.AtomCardinal, 32, 4, w.Bytes())
}

func (c *Conn) SetFrameExtents(window x.Window, val FrameExtents) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_FRAME_EXTENTS"), x.AtomCardinal, 32, 4, w.Bytes())
}

/**
 *    _NET_WM_SYNC_REQUEST_COUNTER
 */

func (c *Conn) GetWmSyncRequestCounter(window x.Window) GetWmSyncRequestCounterCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_SYNC_REQUEST_COUNTER"), x.AtomCardinal, 0, 2)
	return GetWmSyncRequestCounterCookie(cookie)
}

func (c *Conn) GetWmSyncRequestCounterUnchecked(window x.Window) GetWmSyncRequestCounterCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_SYNC_REQUEST_COUNTER"), x.AtomCardinal, 0, 2)
	return GetWmSyncRequestCounterCookie(cookie)
}

type GetWmSyncRequestCounterCookie x.GetPropertyCookie

func (cookie GetWmSyncRequestCounterCookie) Reply(c *Conn) (WmSyncRequestCounter, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return WmSyncRequestCounter{}, err
	}
	if reply.Type != x.AtomCardinal {
		return WmSyncRequestCounter{}, errors.New("bad reply")
	}
	return getWmSyncRequestCounterFromReply(reply)
}

type WmSyncRequestCounter struct {
	Low, High uint32
}

func getWmSyncRequestCounterFromReply(reply *x.GetPropertyReply) (WmSyncRequestCounter, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return WmSyncRequestCounter{}, err
	}

	if len(list) != 2 {
		return WmSyncRequestCounter{}, errors.New("length of list is not 2")
	}
	return WmSyncRequestCounter{
		Low:  list[0],
		High: list[1],
	}, nil
}

func (c *Conn) SetWmSyncRequestCounterChecked(window x.Window, val WmSyncRequestCounter) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Low)
	w.Write4b(val.High)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_SYNC_REQUEST_COUNTER"), x.AtomCardinal, 32, 2, w.Bytes())
}

func (c *Conn) SetWmSyncRequestCounter(window x.Window, val WmSyncRequestCounter) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Low)
	w.Write4b(val.High)
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_SYNC_REQUEST_COUNTER"), x.AtomCardinal, 32, 2, w.Bytes())
}

/**
 *    _NET_WM_FULLSCREEN_MONITORS
 */

func (c *Conn) GetWmFullscreenMonitors(window x.Window) GetWmFullscreenMonitorsCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_FULLSCREEN_MONITORS"), x.AtomCardinal, 0, 4)
	return GetWmFullscreenMonitorsCookie(cookie)
}

func (c *Conn) GetWmFullscreenMonitorsUnchecked(window x.Window) GetWmFullscreenMonitorsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_FULLSCREEN_MONITORS"), x.AtomCardinal, 0, 4)
	return GetWmFullscreenMonitorsCookie(cookie)
}

type GetWmFullscreenMonitorsCookie x.GetPropertyCookie

func (cookie GetWmFullscreenMonitorsCookie) Reply(c *Conn) (WmFullscreenMonitors, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return WmFullscreenMonitors{}, err
	}
	if reply.Type != x.AtomCardinal {
		return WmFullscreenMonitors{}, errors.New("bad reply")
	}
	return getWmFullscreenMonitorsFromReply(reply)
}

type WmFullscreenMonitors struct {
	Top, Bottom, Left, Right uint32
}

func getWmFullscreenMonitorsFromReply(reply *x.GetPropertyReply) (WmFullscreenMonitors, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return WmFullscreenMonitors{}, err
	}

	if len(list) != 4 {
		return WmFullscreenMonitors{}, errors.New("length of list is not 4")
	}
	return WmFullscreenMonitors{
		Top:    list[0],
		Bottom: list[1],
		Left:   list[2],
		Right:  list[3],
	}, nil
}

func (c *Conn) SetWmFullscreenMonitorsChecked(window x.Window, val WmFullscreenMonitors) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_FULLSCREEN_MONITORS"), x.AtomCardinal, 32, 4, w.Bytes())
}

func (c *Conn) SetWmFullscreenMonitors(window x.Window, val WmFullscreenMonitors) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_FULLSCREEN_MONITORS"), x.AtomCardinal, 32, 4, w.Bytes())
}
