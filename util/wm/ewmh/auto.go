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
	cookie := x.GetProperty(c.conn, false, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomAtom, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func (c *Conn) SetSupportedChecked(vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomAtom, 32, w.Bytes())
}

func (c *Conn) SetSupported(vals []x.Atom) {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomAtom, 32, w.Bytes())
}

/**
 *    _NET_CLIENT_LIST
 */

func (c *Conn) GetClientList() GetWindowsCookie {
	cookie := x.GetProperty(c.conn, false, c.GetRootWin(),
		c.GetAtom("_NET_CLIENT_LIST"), x.AtomWindow, 0, LENGTH_MAX)
	return GetWindowsCookie(cookie)
}

func (c *Conn) SetClientListChecked(vals []x.Window) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_CLIENT_LIST"), x.AtomWindow, 32, w.Bytes())
}

func (c *Conn) SetClientList(vals []x.Window) {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_CLIENT_LIST"), x.AtomWindow, 32, w.Bytes())
}

/**
 *    _NET_CLIENT_LIST_STACKING
 */

func (c *Conn) GetClientListStacking() GetWindowsCookie {
	cookie := x.GetProperty(c.conn, false, c.GetRootWin(),
		c.GetAtom("_NET_CLIENT_LIST_STACKING"), x.AtomWindow, 0, LENGTH_MAX)
	return GetWindowsCookie(cookie)
}

func (c *Conn) SetClientListStackingChecked(vals []x.Window) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_CLIENT_LIST_STACKING"), x.AtomWindow, 32, w.Bytes())
}

func (c *Conn) SetClientListStacking(vals []x.Window) {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_CLIENT_LIST_STACKING"), x.AtomWindow, 32, w.Bytes())
}

/**
 *    _NET_NUMBER_OF_DESKTOPS
 */

func (c *Conn) GetNumberOfDesktop() GetCardinalCookie {
	cookie := x.GetProperty(c.conn, false, c.GetRootWin(),
		c.GetAtom("_NET_NUMBER_OF_DESKTOPS"), x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func (c *Conn) SetNumberOfDesktopChecked(val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_NUMBER_OF_DESKTOPS"), x.AtomCardinal, 32, w.Bytes())
}

func (c *Conn) SetNumberOfDesktop(val uint32) {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_NUMBER_OF_DESKTOPS"), x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_DESKTOP_GEOMETRY
 */

func (c *Conn) GetDesktopGeometry() GetDesktopGeometryCookie {
	cookie := x.GetProperty(c.conn, false, c.GetRootWin(),
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
		c.GetAtom("_NET_DESKTOP_GEOMETRY"), x.AtomCardinal, 32, w.Bytes())
}

func (c *Conn) SetDesktopGeometry(val DesktopGeometry) {
	w := x.NewWriter()
	w.Write4b(val.Width)
	w.Write4b(val.Height)
	x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_GEOMETRY"), x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_DESKTOP_VIEWPORT
 */

func (c *Conn) GetDesktopViewport() GetDesktopViewportCookie {
	cookie := x.GetProperty(c.conn, false, c.GetRootWin(),
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
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_VIEWPORT"), x.AtomCardinal, 32, w.Bytes())
}

func (c *Conn) SetDesktopViewport(vals []ViewportLeftTopCorner) {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(val.X)
		w.Write4b(val.Y)
	}
	x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_VIEWPORT"), x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_CURRENT_DESKTOP
 */

func (c *Conn) GetCurrentDesktop() GetCardinalCookie {
	cookie := x.GetProperty(c.conn, false, c.GetRootWin(),
		c.GetAtom("_NET_CURRENT_DESKTOP"), x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func (c *Conn) SetCurrentDesktopChecked(val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_CURRENT_DESKTOP"), x.AtomCardinal, 32, w.Bytes())
}

func (c *Conn) SetCurrentDesktop(val uint32) {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_CURRENT_DESKTOP"), x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_DESKTOP_NAMES
 */

func (c *Conn) GetDesktopNames() GetUTF8StrsCookie {
	cookie := x.GetProperty(c.conn, false, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_NAMES"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrsCookie(cookie)
}

func (c *Conn) SetDesktopNamesChecked(vals []string) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.WriteString(val)
		w.Write1b(0)
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_NAMES"), c.GetAtom("UTF8_STRING"), 8, w.Bytes())
}

func (c *Conn) SetDesktopNames(vals []string) {
	w := x.NewWriter()
	for _, val := range vals {
		w.WriteString(val)
		w.Write1b(0)
	}
	x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_NAMES"), c.GetAtom("UTF8_STRING"), 8, w.Bytes())
}

/**
 *    _NET_ACTIVE_WINDOW
 */

func (c *Conn) GetActiveWindow() GetWindowCookie {
	cookie := x.GetProperty(c.conn, false, c.GetRootWin(),
		c.GetAtom("_NET_ACTIVE_WINDOW"), x.AtomWindow, 0, 1)
	return GetWindowCookie(cookie)
}

func (c *Conn) SetActiveWindowChecked(val x.Window) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_ACTIVE_WINDOW"), x.AtomWindow, 32, w.Bytes())
}

func (c *Conn) SetActiveWindow(val x.Window) {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_ACTIVE_WINDOW"), x.AtomWindow, 32, w.Bytes())
}

/**
 *    _NET_WORKAREA
 */

func (c *Conn) GetWorkarea() GetWorkareaCookie {
	cookie := x.GetProperty(c.conn, false, c.GetRootWin(),
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
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_WORKAREA"), x.AtomCardinal, 32, w.Bytes())
}

func (c *Conn) SetWorkarea(vals []WorkareaGeometry) {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(val.X)
		w.Write4b(val.Y)
		w.Write4b(val.Width)
		w.Write4b(val.Height)
	}
	x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_WORKAREA"), x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_SUPPORTING_WM_CHECK
 */

func (c *Conn) GetSupportingWMCheck(window x.Window) GetWindowCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_SUPPORTING_WM_CHECK"), x.AtomWindow, 0, 1)
	return GetWindowCookie(cookie)
}

func (c *Conn) SetSupportingWMCheckChecked(window x.Window, val x.Window) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_SUPPORTING_WM_CHECK"), x.AtomWindow, 32, w.Bytes())
}

func (c *Conn) SetSupportingWMCheck(window x.Window, val x.Window) {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_SUPPORTING_WM_CHECK"), x.AtomWindow, 32, w.Bytes())
}

/**
 *    _NET_VIRTUAL_ROOTS
 */

func (c *Conn) GetVirtualRoots() GetWindowsCookie {
	cookie := x.GetProperty(c.conn, false, c.GetRootWin(),
		c.GetAtom("_NET_VIRTUAL_ROOTS"), x.AtomWindow, 0, LENGTH_MAX)
	return GetWindowsCookie(cookie)
}

func (c *Conn) SetVirtualRootsChecked(vals []x.Window) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_VIRTUAL_ROOTS"), x.AtomWindow, 32, w.Bytes())
}

func (c *Conn) SetVirtualRoots(vals []x.Window) {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_VIRTUAL_ROOTS"), x.AtomWindow, 32, w.Bytes())
}

/**
 *    _NET_DESKTOP_LAYOUT
 */

func (c *Conn) GetDesktopLayout() GetDesktopLayoutCookie {
	cookie := x.GetProperty(c.conn, false, c.GetRootWin(),
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
		c.GetAtom("_NET_DESKTOP_LAYOUT"), x.AtomCardinal, 32, w.Bytes())
}

func (c *Conn) SetDesktopLayout(val DesktopLayout) {
	w := x.NewWriter()
	w.Write4b(val.Orientation)
	w.Write4b(val.Columns)
	w.Write4b(val.Rows)
	w.Write4b(val.StartingCorner)
	x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_LAYOUT"), x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_SHOWING_DESKTOP
 */

func (c *Conn) GetShowingDesktop() GetBooleanCookie {
	cookie := x.GetProperty(c.conn, false, c.GetRootWin(),
		c.GetAtom("_NET_SHOWING_DESKTOP"), x.AtomCardinal, 0, 1)
	return GetBooleanCookie(cookie)
}

/**
 *    _NET_WM_NAME
 */

func (c *Conn) GetWMName(window x.Window) GetUTF8StrCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_WM_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func (c *Conn) SetWMNameChecked(window x.Window, val string) x.VoidCookie {
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_NAME"), c.GetAtom("UTF8_STRING"), 8, []byte(val))
}

func (c *Conn) SetWMName(window x.Window, val string) {
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_NAME"), c.GetAtom("UTF8_STRING"), 8, []byte(val))
}

/**
 *    _NET_WM_VISIBLE_NAME
 */

func (c *Conn) GetWMVisibleName(window x.Window) GetUTF8StrCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_WM_VISIBLE_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func (c *Conn) SetWMVisibleNameChecked(window x.Window, val string) x.VoidCookie {
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_VISIBLE_NAME"), c.GetAtom("UTF8_STRING"), 8, []byte(val))
}

func (c *Conn) SetWMVisibleName(window x.Window, val string) {
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_VISIBLE_NAME"), c.GetAtom("UTF8_STRING"), 8, []byte(val))
}

/**
 *    _NET_WM_ICON_NAME
 */

func (c *Conn) GetWMIconName(window x.Window) GetUTF8StrCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_WM_ICON_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func (c *Conn) SetWMIconNameChecked(window x.Window, val string) x.VoidCookie {
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_ICON_NAME"), c.GetAtom("UTF8_STRING"), 8, []byte(val))
}

func (c *Conn) SetWMIconName(window x.Window, val string) {
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_ICON_NAME"), c.GetAtom("UTF8_STRING"), 8, []byte(val))
}

/**
 *    _NET_WM_VISIBLE_ICON_NAME
 */

func (c *Conn) GetWMVisibleIconName(window x.Window) GetUTF8StrCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_WM_VISIBLE_ICON_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func (c *Conn) SetWMVisibleIconNameChecked(window x.Window, val string) x.VoidCookie {
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_VISIBLE_ICON_NAME"), c.GetAtom("UTF8_STRING"), 8, []byte(val))
}

func (c *Conn) SetWMVisibleIconName(window x.Window, val string) {
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_VISIBLE_ICON_NAME"), c.GetAtom("UTF8_STRING"), 8, []byte(val))
}

/**
 *    _NET_WM_DESKTOP
 */

func (c *Conn) GetWMDesktop(window x.Window) GetCardinalCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_WM_DESKTOP"), x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func (c *Conn) SetWMDesktopChecked(window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_DESKTOP"), x.AtomCardinal, 32, w.Bytes())
}

func (c *Conn) SetWMDesktop(window x.Window, val uint32) {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_DESKTOP"), x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_WM_WINDOW_TYPE
 */

func (c *Conn) GetWMWindowType(window x.Window) GetAtomsCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_WM_WINDOW_TYPE"), x.AtomAtom, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func (c *Conn) SetWMWindowTypeChecked(window x.Window, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_WINDOW_TYPE"), x.AtomAtom, 32, w.Bytes())
}

func (c *Conn) SetWMWindowType(window x.Window, vals []x.Atom) {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_WINDOW_TYPE"), x.AtomAtom, 32, w.Bytes())
}

/**
 *    _NET_WM_STATE
 */

func (c *Conn) GetWMState(window x.Window) GetAtomsCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_WM_STATE"), x.AtomAtom, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func (c *Conn) SetWMStateChecked(window x.Window, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_STATE"), x.AtomAtom, 32, w.Bytes())
}

func (c *Conn) SetWMState(window x.Window, vals []x.Atom) {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_STATE"), x.AtomAtom, 32, w.Bytes())
}

/**
 *    _NET_WM_ALLOWED_ACTIONS
 */

func (c *Conn) GetWMAllowedActions(window x.Window) GetAtomsCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_WM_ALLOWED_ACTIONS"), x.AtomAtom, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func (c *Conn) SetWMAllowedActionsChecked(window x.Window, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_ALLOWED_ACTIONS"), x.AtomAtom, 32, w.Bytes())
}

func (c *Conn) SetWMAllowedActions(window x.Window, vals []x.Atom) {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_ALLOWED_ACTIONS"), x.AtomAtom, 32, w.Bytes())
}

/**
 *    _NET_WM_STRUT
 */

func (c *Conn) GetWMStrut(window x.Window) GetWMStrutCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_WM_STRUT"), x.AtomCardinal, 0, 4)
	return GetWMStrutCookie(cookie)
}

type GetWMStrutCookie x.GetPropertyCookie

func (cookie GetWMStrutCookie) Reply(c *Conn) (WMStrut, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
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

func (c *Conn) SetWMStrutChecked(window x.Window, val WMStrut) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_STRUT"), x.AtomCardinal, 32, w.Bytes())
}

func (c *Conn) SetWMStrut(window x.Window, val WMStrut) {
	w := x.NewWriter()
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_STRUT"), x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_WM_STRUT_PARTIAL
 */

func (c *Conn) GetWMStrutPartial(window x.Window) GetWMStrutPartialCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_WM_STRUT_PARTIAL"), x.AtomCardinal, 0, 12)
	return GetWMStrutPartialCookie(cookie)
}

type GetWMStrutPartialCookie x.GetPropertyCookie

func (cookie GetWMStrutPartialCookie) Reply(c *Conn) (WMStrutPartial, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
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

func (c *Conn) SetWMStrutPartialChecked(window x.Window, val WMStrutPartial) x.VoidCookie {
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
		c.GetAtom("_NET_WM_STRUT_PARTIAL"), x.AtomCardinal, 32, w.Bytes())
}

func (c *Conn) SetWMStrutPartial(window x.Window, val WMStrutPartial) {
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
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_STRUT_PARTIAL"), x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_WM_ICON_GEOMETRY
 */

func (c *Conn) GetWMIconGeometry(window x.Window) GetWMIconGeometryCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_WM_ICON_GEOMETRY"), x.AtomCardinal, 0, 4)
	return GetWMIconGeometryCookie(cookie)
}

type GetWMIconGeometryCookie x.GetPropertyCookie

func (cookie GetWMIconGeometryCookie) Reply(c *Conn) (WMIconGeometry, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
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

func (c *Conn) SetWMIconGeometryChecked(window x.Window, val WMIconGeometry) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.X)
	w.Write4b(val.Y)
	w.Write4b(val.Width)
	w.Write4b(val.Height)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_ICON_GEOMETRY"), x.AtomCardinal, 32, w.Bytes())
}

func (c *Conn) SetWMIconGeometry(window x.Window, val WMIconGeometry) {
	w := x.NewWriter()
	w.Write4b(val.X)
	w.Write4b(val.Y)
	w.Write4b(val.Width)
	w.Write4b(val.Height)
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_ICON_GEOMETRY"), x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_WM_ICON
 */

func (c *Conn) GetWMIcon(window x.Window) GetWMIconCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_WM_ICON"), x.AtomCardinal, 0, LENGTH_MAX)
	return GetWMIconCookie(cookie)
}

type GetWMIconCookie x.GetPropertyCookie

func (cookie GetWMIconCookie) Reply(c *Conn) ([]WMIcon, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
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

func (c *Conn) GetWMPid(window x.Window) GetCardinalCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_WM_PID"), x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func (c *Conn) SetWMPidChecked(window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_PID"), x.AtomCardinal, 32, w.Bytes())
}

func (c *Conn) SetWMPid(window x.Window, val uint32) {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_PID"), x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_WM_HANDLED_ICONS
 */

func (c *Conn) GetWMHandledIcons(window x.Window) GetCardinalCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_WM_HANDLED_ICONS"), x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func (c *Conn) SetWMHandledIconsChecked(window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_HANDLED_ICONS"), x.AtomCardinal, 32, w.Bytes())
}

func (c *Conn) SetWMHandledIcons(window x.Window, val uint32) {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_HANDLED_ICONS"), x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_WM_USER_TIME
 */

func (c *Conn) GetWMUserTime(window x.Window) GetCardinalCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_WM_USER_TIME"), x.AtomCardinal, 0, 1)
	return GetCardinalCookie(cookie)
}

func (c *Conn) SetWMUserTimeChecked(window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_USER_TIME"), x.AtomCardinal, 32, w.Bytes())
}

func (c *Conn) SetWMUserTime(window x.Window, val uint32) {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_USER_TIME"), x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_WM_USER_TIME_WINDOW
 */

func (c *Conn) GetWMUserTimeWindow(window x.Window) GetWindowCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_WM_USER_TIME_WINDOW"), x.AtomWindow, 0, 1)
	return GetWindowCookie(cookie)
}

func (c *Conn) SetWMUserTimeWindowChecked(window x.Window, val x.Window) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_USER_TIME_WINDOW"), x.AtomWindow, 32, w.Bytes())
}

func (c *Conn) SetWMUserTimeWindow(window x.Window, val x.Window) {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_USER_TIME_WINDOW"), x.AtomWindow, 32, w.Bytes())
}

/**
 *    _NET_FRAME_EXTENTS
 */

func (c *Conn) GetFrameExtents(window x.Window) GetFrameExtentsCookie {
	cookie := x.GetProperty(c.conn, false, window,
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
		c.GetAtom("_NET_FRAME_EXTENTS"), x.AtomCardinal, 32, w.Bytes())
}

func (c *Conn) SetFrameExtents(window x.Window, val FrameExtents) {
	w := x.NewWriter()
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_FRAME_EXTENTS"), x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_WM_SYNC_REQUEST_COUNTER
 */

func (c *Conn) GetWMSyncRequestCounter(window x.Window) GetWMSyncRequestCounterCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_WM_SYNC_REQUEST_COUNTER"), x.AtomCardinal, 0, 2)
	return GetWMSyncRequestCounterCookie(cookie)
}

type GetWMSyncRequestCounterCookie x.GetPropertyCookie

func (cookie GetWMSyncRequestCounterCookie) Reply(c *Conn) (WMSyncRequestCounter, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
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

func (c *Conn) SetWMSyncRequestCounterChecked(window x.Window, val WMSyncRequestCounter) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Low)
	w.Write4b(val.High)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_SYNC_REQUEST_COUNTER"), x.AtomCardinal, 32, w.Bytes())
}

func (c *Conn) SetWMSyncRequestCounter(window x.Window, val WMSyncRequestCounter) {
	w := x.NewWriter()
	w.Write4b(val.Low)
	w.Write4b(val.High)
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_SYNC_REQUEST_COUNTER"), x.AtomCardinal, 32, w.Bytes())
}

/**
 *    _NET_WM_FULLSCREEN_MONITORS
 */

func (c *Conn) GetWMFullscreenMonitors(window x.Window) GetWMFullscreenMonitorsCookie {
	cookie := x.GetProperty(c.conn, false, window,
		c.GetAtom("_NET_WM_FULLSCREEN_MONITORS"), x.AtomCardinal, 0, 4)
	return GetWMFullscreenMonitorsCookie(cookie)
}

type GetWMFullscreenMonitorsCookie x.GetPropertyCookie

func (cookie GetWMFullscreenMonitorsCookie) Reply(c *Conn) (WMFullscreenMonitors, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
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

func (c *Conn) SetWMFullscreenMonitorsChecked(window x.Window, val WMFullscreenMonitors) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_FULLSCREEN_MONITORS"), x.AtomCardinal, 32, w.Bytes())
}

func (c *Conn) SetWMFullscreenMonitors(window x.Window, val WMFullscreenMonitors) {
	w := x.NewWriter()
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_FULLSCREEN_MONITORS"), x.AtomCardinal, 32, w.Bytes())
}
