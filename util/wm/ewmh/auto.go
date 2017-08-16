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
	if reply.Type != x.AtomCARDINAL {
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
	if reply.Type != x.AtomWINDOW {
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
	if reply.Type != x.AtomCARDINAL {
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
	if reply.Type != x.AtomWINDOW {
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
	if reply.Type != x.AtomATOM {
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
	if reply.Type != x.AtomCARDINAL {
		return nil, errors.New("bad reply")
	}
	return getCardinalsFromReply(reply)
}

/**
 *    _NET_SUPPORTED
 */

func GetSupported(c *Conn) GetAtomsCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomATOM, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func GetSupportedUnchecked(c *Conn) GetAtomsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomATOM, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func SetSupportedChecked(c *Conn, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomATOM, 32, uint32(len(vals)), w.Bytes())
}

func SetSupported(c *Conn, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomATOM, 32, uint32(len(vals)), w.Bytes())
}

/**
 *    _NET_CLIENT_LIST
 */

func GetClientList(c *Conn) GetWindowsCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_CLIENT_LIST"), x.AtomWINDOW, 0, LENGTH_MAX)
	return GetWindowsCookie(cookie)
}

func GetClientListUnchecked(c *Conn) GetWindowsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_CLIENT_LIST"), x.AtomWINDOW, 0, LENGTH_MAX)
	return GetWindowsCookie(cookie)
}

func SetClientListChecked(c *Conn, vals []x.Window) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomWINDOW, 32, uint32(len(vals)), w.Bytes())
}

func SetClientList(c *Conn, vals []x.Window) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomWINDOW, 32, uint32(len(vals)), w.Bytes())
}

/**
 *    _NET_CLIENT_LIST_STACKING
 */

func GetClientListStacking(c *Conn) GetWindowsCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_CLIENT_LIST_STACKING"), x.AtomWINDOW, 0, LENGTH_MAX)
	return GetWindowsCookie(cookie)
}

func GetClientListStackingUnchecked(c *Conn) GetWindowsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_CLIENT_LIST_STACKING"), x.AtomWINDOW, 0, LENGTH_MAX)
	return GetWindowsCookie(cookie)
}

func SetClientListStackingChecked(c *Conn, vals []x.Window) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomWINDOW, 32, uint32(len(vals)), w.Bytes())
}

func SetClientListStacking(c *Conn, vals []x.Window) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomWINDOW, 32, uint32(len(vals)), w.Bytes())
}

/**
 *    _NET_NUMBER_OF_DESKTOPS
 */

func GetNumberOfDesktop(c *Conn) GetCardinalCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_NUMBER_OF_DESKTOPS"), x.AtomCARDINAL, 0, 1)
	return GetCardinalCookie(cookie)
}

func GetNumberOfDesktopUnchecked(c *Conn) GetCardinalCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_NUMBER_OF_DESKTOPS"), x.AtomCARDINAL, 0, 1)
	return GetCardinalCookie(cookie)
}

func SetNumberOfDesktopChecked(c *Conn, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_NUMBER_OF_DESKTOPS"), x.AtomCARDINAL, 32, 1, w.Bytes())
}

func SetNumberOfDesktop(c *Conn, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_NUMBER_OF_DESKTOPS"), x.AtomCARDINAL, 32, 1, w.Bytes())
}

/**
 *    _NET_DESKTOP_GEOMETRY
 */

func GetDesktopGeometry(c *Conn) GetDesktopGeometryCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_GEOMETRY"), x.AtomCARDINAL, 0, 2)
	return GetDesktopGeometryCookie(cookie)
}

func GetDesktopGeometryUnchecked(c *Conn) GetDesktopGeometryCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_GEOMETRY"), x.AtomCARDINAL, 0, 2)
	return GetDesktopGeometryCookie(cookie)
}

type GetDesktopGeometryCookie x.GetPropertyCookie

func (cookie GetDesktopGeometryCookie) Reply(c *Conn) (DesktopGeometry, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return DesktopGeometry{}, err
	}
	if reply.Type != x.AtomCARDINAL {
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

func SetDesktopGeometryChecked(c *Conn, val DesktopGeometry) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Width)
	w.Write4b(val.Height)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_GEOMETRY"), x.AtomCARDINAL, 32, 2, w.Bytes())
}

func SetDesktopGeometry(c *Conn, val DesktopGeometry) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Width)
	w.Write4b(val.Height)
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_GEOMETRY"), x.AtomCARDINAL, 32, 2, w.Bytes())
}

/**
 *    _NET_DESKTOP_VIEWPORT
 */

func GetDesktopViewport(c *Conn) GetDesktopViewportCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_VIEWPORT"), x.AtomCARDINAL, 0, LENGTH_MAX)
	return GetDesktopViewportCookie(cookie)
}

func GetDesktopViewportUnchecked(c *Conn) GetDesktopViewportCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_VIEWPORT"), x.AtomCARDINAL, 0, LENGTH_MAX)
	return GetDesktopViewportCookie(cookie)
}

type GetDesktopViewportCookie x.GetPropertyCookie

func (cookie GetDesktopViewportCookie) Reply(c *Conn) ([]ViewportLeftTopCorner, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return nil, err
	}
	if reply.Type != x.AtomCARDINAL {
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

func SetDesktopViewportChecked(c *Conn, vals []ViewportLeftTopCorner) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(val.X)
		w.Write4b(val.Y)
	}
	length := uint32(len(vals) * 2)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_VIEWPORT"), x.AtomCARDINAL, 32, length, w.Bytes())
}

func SetDesktopViewport(c *Conn, vals []ViewportLeftTopCorner) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(val.X)
		w.Write4b(val.Y)
	}
	length := uint32(len(vals) * 2)
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_VIEWPORT"), x.AtomCARDINAL, 32, length, w.Bytes())
}

/**
 *    _NET_CURRENT_DESKTOP
 */

func GetCurrentDesktop(c *Conn) GetCardinalCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_CURRENT_DESKTOP"), x.AtomCARDINAL, 0, 1)
	return GetCardinalCookie(cookie)
}

func GetCurrentDesktopUnchecked(c *Conn) GetCardinalCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_CURRENT_DESKTOP"), x.AtomCARDINAL, 0, 1)
	return GetCardinalCookie(cookie)
}

func SetCurrentDesktopChecked(c *Conn, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_CURRENT_DESKTOP"), x.AtomCARDINAL, 32, 1, w.Bytes())
}

func SetCurrentDesktop(c *Conn, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_CURRENT_DESKTOP"), x.AtomCARDINAL, 32, 1, w.Bytes())
}

/**
 *    _NET_DESKTOP_NAMES
 */

func GetDesktopNames(c *Conn) GetUTF8StrsCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_NAMES"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrsCookie(cookie)
}

func GetDesktopNamesUnchecked(c *Conn) GetUTF8StrsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_NAMES"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrsCookie(cookie)
}

func SetDesktopNamesChecked(c *Conn, vals []string) x.VoidCookie {
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

func SetDesktopNames(c *Conn, vals []string) x.VoidCookie {
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

func GetActiveWindow(c *Conn) GetWindowCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_ACTIVE_WINDOW"), x.AtomWINDOW, 0, 1)
	return GetWindowCookie(cookie)
}

func GetActiveWindowUnchecked(c *Conn) GetWindowCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_ACTIVE_WINDOW"), x.AtomWINDOW, 0, 1)
	return GetWindowCookie(cookie)
}

func SetActiveWindowChecked(c *Conn, val x.Window) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_ACTIVE_WINDOW"), x.AtomWINDOW, 32, 1, w.Bytes())
}

func SetActiveWindow(c *Conn, val x.Window) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_ACTIVE_WINDOW"), x.AtomWINDOW, 32, 1, w.Bytes())
}

/**
 *    _NET_WORKAREA
 */

func GetWorkarea(c *Conn) GetWorkareaCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_WORKAREA"), x.AtomCARDINAL, 0, LENGTH_MAX)
	return GetWorkareaCookie(cookie)
}

func GetWorkareaUnchecked(c *Conn) GetWorkareaCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_WORKAREA"), x.AtomCARDINAL, 0, LENGTH_MAX)
	return GetWorkareaCookie(cookie)
}

type GetWorkareaCookie x.GetPropertyCookie

func (cookie GetWorkareaCookie) Reply(c *Conn) ([]WorkareaGeometry, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return nil, err
	}
	if reply.Type != x.AtomCARDINAL {
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

func SetWorkareaChecked(c *Conn, vals []WorkareaGeometry) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(val.X)
		w.Write4b(val.Y)
		w.Write4b(val.Width)
		w.Write4b(val.Height)
	}
	length := uint32(len(vals) * 4)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_WORKAREA"), x.AtomCARDINAL, 32, length, w.Bytes())
}

func SetWorkarea(c *Conn, vals []WorkareaGeometry) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(val.X)
		w.Write4b(val.Y)
		w.Write4b(val.Width)
		w.Write4b(val.Height)
	}
	length := uint32(len(vals) * 4)
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_WORKAREA"), x.AtomCARDINAL, 32, length, w.Bytes())
}

/**
 *    _NET_SUPPORTING_WM_CHECK
 */

func GetSupportingWmCheck(c *Conn, window x.Window) GetWindowCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_SUPPORTING_WM_CHECK"), x.AtomWINDOW, 0, 1)
	return GetWindowCookie(cookie)
}

func GetSupportingWmCheckUnchecked(c *Conn, window x.Window) GetWindowCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_SUPPORTING_WM_CHECK"), x.AtomWINDOW, 0, 1)
	return GetWindowCookie(cookie)
}

func SetSupportingWmCheckChecked(c *Conn, window x.Window, val x.Window) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_SUPPORTING_WM_CHECK"), x.AtomWINDOW, 32, 1, w.Bytes())
}

func SetSupportingWmCheck(c *Conn, window x.Window, val x.Window) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_SUPPORTING_WM_CHECK"), x.AtomWINDOW, 32, 1, w.Bytes())
}

/**
 *    _NET_VIRTUAL_ROOTS
 */

func GetVirtualRoots(c *Conn) GetWindowsCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_VIRTUAL_ROOTS"), x.AtomWINDOW, 0, LENGTH_MAX)
	return GetWindowsCookie(cookie)
}

func GetVirtualRootsUnchecked(c *Conn) GetWindowsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_VIRTUAL_ROOTS"), x.AtomWINDOW, 0, LENGTH_MAX)
	return GetWindowsCookie(cookie)
}

func SetVirtualRootsChecked(c *Conn, vals []x.Window) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomWINDOW, 32, uint32(len(vals)), w.Bytes())
}

func SetVirtualRoots(c *Conn, vals []x.Window) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_SUPPORTED"), x.AtomWINDOW, 32, uint32(len(vals)), w.Bytes())
}

/**
 *    _NET_DESKTOP_LAYOUT
 */

func GetDesktopLayout(c *Conn) GetDesktopLayoutCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_LAYOUT"), x.AtomCARDINAL, 0, 4)
	return GetDesktopLayoutCookie(cookie)
}

func GetDesktopLayoutUnchecked(c *Conn) GetDesktopLayoutCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_LAYOUT"), x.AtomCARDINAL, 0, 4)
	return GetDesktopLayoutCookie(cookie)
}

type GetDesktopLayoutCookie x.GetPropertyCookie

func (cookie GetDesktopLayoutCookie) Reply(c *Conn) (DesktopLayout, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return DesktopLayout{}, err
	}
	if reply.Type != x.AtomCARDINAL {
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

func SetDesktopLayoutChecked(c *Conn, val DesktopLayout) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Orientation)
	w.Write4b(val.Columns)
	w.Write4b(val.Rows)
	w.Write4b(val.StartingCorner)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_LAYOUT"), x.AtomCARDINAL, 32, 4, w.Bytes())
}

func SetDesktopLayout(c *Conn, val DesktopLayout) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Orientation)
	w.Write4b(val.Columns)
	w.Write4b(val.Rows)
	w.Write4b(val.StartingCorner)
	return x.ChangeProperty(c.conn, x.PropModeReplace, c.GetRootWin(),
		c.GetAtom("_NET_DESKTOP_LAYOUT"), x.AtomCARDINAL, 32, 4, w.Bytes())
}

/**
 *    _NET_SHOWING_DESKTOP
 */

func GetShowingDesktop(c *Conn) GetBooleanCookie {
	cookie := x.GetProperty(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_SHOWING_DESKTOP"), x.AtomCARDINAL, 0, 1)
	return GetBooleanCookie(cookie)
}

func GetShowingDesktopUnchecked(c *Conn) GetBooleanCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, c.GetRootWin(),
		c.GetAtom("_NET_SHOWING_DESKTOP"), x.AtomCARDINAL, 0, 1)
	return GetBooleanCookie(cookie)
}

/**
 *    _NET_WM_NAME
 */

func GetWmName(c *Conn, window x.Window) GetUTF8StrCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func GetWmNameUnchecked(c *Conn, window x.Window) GetUTF8StrCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func SetWmNameChecked(c *Conn, window x.Window, val string) x.VoidCookie {
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_NAME"), c.GetAtom("UTF8_STRING"), 8, uint32(len(val)), []byte(val))
}

func SetWmName(c *Conn, window x.Window, val string) x.VoidCookie {
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_NAME"), c.GetAtom("UTF8_STRING"), 8, uint32(len(val)), []byte(val))
}

/**
 *    _NET_WM_VISIBLE_NAME
 */

func GetWmVisibleName(c *Conn, window x.Window) GetUTF8StrCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_VISIBLE_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func GetWmVisibleNameUnchecked(c *Conn, window x.Window) GetUTF8StrCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_VISIBLE_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func SetWmVisibleNameChecked(c *Conn, window x.Window, val string) x.VoidCookie {
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_VISIBLE_NAME"), c.GetAtom("UTF8_STRING"), 8, uint32(len(val)), []byte(val))
}

func SetWmVisibleName(c *Conn, window x.Window, val string) x.VoidCookie {
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_VISIBLE_NAME"), c.GetAtom("UTF8_STRING"), 8, uint32(len(val)), []byte(val))
}

/**
 *    _NET_WM_ICON_NAME
 */

func GetWmIconName(c *Conn, window x.Window) GetUTF8StrCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_ICON_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func GetWmIconNameUnchecked(c *Conn, window x.Window) GetUTF8StrCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_ICON_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func SetWmIconNameChecked(c *Conn, window x.Window, val string) x.VoidCookie {
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_ICON_NAME"), c.GetAtom("UTF8_STRING"), 8, uint32(len(val)), []byte(val))
}

func SetWmIconName(c *Conn, window x.Window, val string) x.VoidCookie {
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_ICON_NAME"), c.GetAtom("UTF8_STRING"), 8, uint32(len(val)), []byte(val))
}

/**
 *    _NET_WM_VISIBLE_ICON_NAME
 */

func GetWmVisibleIconName(c *Conn, window x.Window) GetUTF8StrCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_VISIBLE_ICON_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func GetWmVisibleIconNameUnchecked(c *Conn, window x.Window) GetUTF8StrCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_VISIBLE_ICON_NAME"), c.GetAtom("UTF8_STRING"), 0, LENGTH_MAX)
	return GetUTF8StrCookie(cookie)
}

func SetWmVisibleIconNameChecked(c *Conn, window x.Window, val string) x.VoidCookie {
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_VISIBLE_ICON_NAME"), c.GetAtom("UTF8_STRING"), 8, uint32(len(val)), []byte(val))
}

func SetWmVisibleIconName(c *Conn, window x.Window, val string) x.VoidCookie {
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_VISIBLE_ICON_NAME"), c.GetAtom("UTF8_STRING"), 8, uint32(len(val)), []byte(val))
}

/**
 *    _NET_WM_DESKTOP
 */

func GetWmDesktop(c *Conn, window x.Window) GetCardinalCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_DESKTOP"), x.AtomCARDINAL, 0, 1)
	return GetCardinalCookie(cookie)
}

func GetWmDesktopUnchecked(c *Conn, window x.Window) GetCardinalCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_DESKTOP"), x.AtomCARDINAL, 0, 1)
	return GetCardinalCookie(cookie)
}

func SetWmDesktopChecked(c *Conn, window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_DESKTOP"), x.AtomCARDINAL, 32, 1, w.Bytes())
}

func SetWmDesktop(c *Conn, window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_DESKTOP"), x.AtomCARDINAL, 32, 1, w.Bytes())
}

/**
 *    _NET_WM_WINDOW_TYPE
 */

func GetWmWindowType(c *Conn, window x.Window) GetAtomsCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_WINDOW_TYPE"), x.AtomATOM, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func GetWmWindowTypeUnchecked(c *Conn, window x.Window) GetAtomsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_WINDOW_TYPE"), x.AtomATOM, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func SetWmWindowTypeChecked(c *Conn, window x.Window, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_SUPPORTED"), x.AtomATOM, 32, uint32(len(vals)), w.Bytes())
}

func SetWmWindowType(c *Conn, window x.Window, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_SUPPORTED"), x.AtomATOM, 32, uint32(len(vals)), w.Bytes())
}

/**
 *    _NET_WM_STATE
 */

func GetWmState(c *Conn, window x.Window) GetAtomsCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_STATE"), x.AtomATOM, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func GetWmStateUnchecked(c *Conn, window x.Window) GetAtomsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_STATE"), x.AtomATOM, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func SetWmStateChecked(c *Conn, window x.Window, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_SUPPORTED"), x.AtomATOM, 32, uint32(len(vals)), w.Bytes())
}

func SetWmState(c *Conn, window x.Window, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_SUPPORTED"), x.AtomATOM, 32, uint32(len(vals)), w.Bytes())
}

/**
 *    _NET_WM_ALLOWED_ACTIONS
 */

func GetWmAllowedActions(c *Conn, window x.Window) GetAtomsCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_ALLOWED_ACTIONS"), x.AtomATOM, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func GetWmAllowedActionsUnchecked(c *Conn, window x.Window) GetAtomsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_ALLOWED_ACTIONS"), x.AtomATOM, 0, LENGTH_MAX)
	return GetAtomsCookie(cookie)
}

func SetWmAllowedActionsChecked(c *Conn, window x.Window, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_SUPPORTED"), x.AtomATOM, 32, uint32(len(vals)), w.Bytes())
}

func SetWmAllowedActions(c *Conn, window x.Window, vals []x.Atom) x.VoidCookie {
	w := x.NewWriter()
	for _, val := range vals {
		w.Write4b(uint32(val))
	}
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_SUPPORTED"), x.AtomATOM, 32, uint32(len(vals)), w.Bytes())
}

/**
 *    _NET_WM_STRUT
 */

func GetWmStrut(c *Conn, window x.Window) GetWmStrutCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_STRUT"), x.AtomCARDINAL, 0, 4)
	return GetWmStrutCookie(cookie)
}

func GetWmStrutUnchecked(c *Conn, window x.Window) GetWmStrutCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_STRUT"), x.AtomCARDINAL, 0, 4)
	return GetWmStrutCookie(cookie)
}

type GetWmStrutCookie x.GetPropertyCookie

func (cookie GetWmStrutCookie) Reply(c *Conn) (WmStrut, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return WmStrut{}, err
	}
	if reply.Type != x.AtomCARDINAL {
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

func SetWmStrutChecked(c *Conn, window x.Window, val WmStrut) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_STRUT"), x.AtomCARDINAL, 32, 4, w.Bytes())
}

func SetWmStrut(c *Conn, window x.Window, val WmStrut) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_STRUT"), x.AtomCARDINAL, 32, 4, w.Bytes())
}

/**
 *    _NET_WM_STRUT_PARTIAL
 */

func GetWmStrutPartial(c *Conn, window x.Window) GetWmStrutPartialCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_STRUT_PARTIAL"), x.AtomCARDINAL, 0, 12)
	return GetWmStrutPartialCookie(cookie)
}

func GetWmStrutPartialUnchecked(c *Conn, window x.Window) GetWmStrutPartialCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_STRUT_PARTIAL"), x.AtomCARDINAL, 0, 12)
	return GetWmStrutPartialCookie(cookie)
}

type GetWmStrutPartialCookie x.GetPropertyCookie

func (cookie GetWmStrutPartialCookie) Reply(c *Conn) (WmStrutPartial, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return WmStrutPartial{}, err
	}
	if reply.Type != x.AtomCARDINAL {
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

func SetWmStrutPartialChecked(c *Conn, window x.Window, val WmStrutPartial) x.VoidCookie {
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
		c.GetAtom("_NET_WM_STRUT_PARTIAL"), x.AtomCARDINAL, 32, 12, w.Bytes())
}

func SetWmStrutPartial(c *Conn, window x.Window, val WmStrutPartial) x.VoidCookie {
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
		c.GetAtom("_NET_WM_STRUT_PARTIAL"), x.AtomCARDINAL, 32, 12, w.Bytes())
}

/**
 *    _NET_WM_ICON_GEOMETRY
 */

func GetWmIconGeometry(c *Conn, window x.Window) GetWmIconGeometryCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_ICON_GEOMETRY"), x.AtomCARDINAL, 0, 4)
	return GetWmIconGeometryCookie(cookie)
}

func GetWmIconGeometryUnchecked(c *Conn, window x.Window) GetWmIconGeometryCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_ICON_GEOMETRY"), x.AtomCARDINAL, 0, 4)
	return GetWmIconGeometryCookie(cookie)
}

type GetWmIconGeometryCookie x.GetPropertyCookie

func (cookie GetWmIconGeometryCookie) Reply(c *Conn) (WmIconGeometry, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return WmIconGeometry{}, err
	}
	if reply.Type != x.AtomCARDINAL {
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

func SetWmIconGeometryChecked(c *Conn, window x.Window, val WmIconGeometry) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.X)
	w.Write4b(val.Y)
	w.Write4b(val.Width)
	w.Write4b(val.Height)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_ICON_GEOMETRY"), x.AtomCARDINAL, 32, 4, w.Bytes())
}

func SetWmIconGeometry(c *Conn, window x.Window, val WmIconGeometry) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.X)
	w.Write4b(val.Y)
	w.Write4b(val.Width)
	w.Write4b(val.Height)
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_ICON_GEOMETRY"), x.AtomCARDINAL, 32, 4, w.Bytes())
}

/**
 *    _NET_WM_ICON
 */

func GetWmIcon(c *Conn, window x.Window) GetWmIconCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_ICON"), x.AtomCARDINAL, 0, LENGTH_MAX)
	return GetWmIconCookie(cookie)
}

func GetWmIconUnchecked(c *Conn, window x.Window) GetWmIconCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_ICON"), x.AtomCARDINAL, 0, LENGTH_MAX)
	return GetWmIconCookie(cookie)
}

type GetWmIconCookie x.GetPropertyCookie

func (cookie GetWmIconCookie) Reply(c *Conn) ([]WmIcon, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return nil, err
	}
	if reply.Type != x.AtomCARDINAL {
		return nil, errors.New("bad reply")
	}
	return getWmIconFromReply(reply)
}

/**
 *    _NET_WM_PID
 */

func GetWmPid(c *Conn, window x.Window) GetCardinalCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_PID"), x.AtomCARDINAL, 0, 1)
	return GetCardinalCookie(cookie)
}

func GetWmPidUnchecked(c *Conn, window x.Window) GetCardinalCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_PID"), x.AtomCARDINAL, 0, 1)
	return GetCardinalCookie(cookie)
}

func SetWmPidChecked(c *Conn, window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_PID"), x.AtomCARDINAL, 32, 1, w.Bytes())
}

func SetWmPid(c *Conn, window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_PID"), x.AtomCARDINAL, 32, 1, w.Bytes())
}

/**
 *    _NET_WM_HANDLED_ICONS
 */

func GetWmHandledIcons(c *Conn, window x.Window) GetCardinalCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_HANDLED_ICONS"), x.AtomCARDINAL, 0, 1)
	return GetCardinalCookie(cookie)
}

func GetWmHandledIconsUnchecked(c *Conn, window x.Window) GetCardinalCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_HANDLED_ICONS"), x.AtomCARDINAL, 0, 1)
	return GetCardinalCookie(cookie)
}

func SetWmHandledIconsChecked(c *Conn, window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_HANDLED_ICONS"), x.AtomCARDINAL, 32, 1, w.Bytes())
}

func SetWmHandledIcons(c *Conn, window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_HANDLED_ICONS"), x.AtomCARDINAL, 32, 1, w.Bytes())
}

/**
 *    _NET_WM_USER_TIME
 */

func GetWmUserTime(c *Conn, window x.Window) GetCardinalCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_USER_TIME"), x.AtomCARDINAL, 0, 1)
	return GetCardinalCookie(cookie)
}

func GetWmUserTimeUnchecked(c *Conn, window x.Window) GetCardinalCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_USER_TIME"), x.AtomCARDINAL, 0, 1)
	return GetCardinalCookie(cookie)
}

func SetWmUserTimeChecked(c *Conn, window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_USER_TIME"), x.AtomCARDINAL, 32, 1, w.Bytes())
}

func SetWmUserTime(c *Conn, window x.Window, val uint32) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_USER_TIME"), x.AtomCARDINAL, 32, 1, w.Bytes())
}

/**
 *    _NET_WM_USER_TIME_WINDOW
 */

func GetWmUserTimeWindow(c *Conn, window x.Window) GetWindowCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_USER_TIME_WINDOW"), x.AtomWINDOW, 0, 1)
	return GetWindowCookie(cookie)
}

func GetWmUserTimeWindowUnchecked(c *Conn, window x.Window) GetWindowCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_USER_TIME_WINDOW"), x.AtomWINDOW, 0, 1)
	return GetWindowCookie(cookie)
}

func SetWmUserTimeWindowChecked(c *Conn, window x.Window, val x.Window) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_USER_TIME_WINDOW"), x.AtomWINDOW, 32, 1, w.Bytes())
}

func SetWmUserTimeWindow(c *Conn, window x.Window, val x.Window) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(uint32(val))
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_USER_TIME_WINDOW"), x.AtomWINDOW, 32, 1, w.Bytes())
}

/**
 *    _NET_FRAME_EXTENTS
 */

func GetFrameExtents(c *Conn, window x.Window) GetFrameExtentsCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_FRAME_EXTENTS"), x.AtomCARDINAL, 0, 4)
	return GetFrameExtentsCookie(cookie)
}

func GetFrameExtentsUnchecked(c *Conn, window x.Window) GetFrameExtentsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_FRAME_EXTENTS"), x.AtomCARDINAL, 0, 4)
	return GetFrameExtentsCookie(cookie)
}

type GetFrameExtentsCookie x.GetPropertyCookie

func (cookie GetFrameExtentsCookie) Reply(c *Conn) (FrameExtents, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return FrameExtents{}, err
	}
	if reply.Type != x.AtomCARDINAL {
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

func SetFrameExtentsChecked(c *Conn, window x.Window, val FrameExtents) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_FRAME_EXTENTS"), x.AtomCARDINAL, 32, 4, w.Bytes())
}

func SetFrameExtents(c *Conn, window x.Window, val FrameExtents) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_FRAME_EXTENTS"), x.AtomCARDINAL, 32, 4, w.Bytes())
}

/**
 *    _NET_WM_SYNC_REQUEST_COUNTER
 */

func GetWmSyncRequestCounter(c *Conn, window x.Window) GetWmSyncRequestCounterCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_SYNC_REQUEST_COUNTER"), x.AtomCARDINAL, 0, 2)
	return GetWmSyncRequestCounterCookie(cookie)
}

func GetWmSyncRequestCounterUnchecked(c *Conn, window x.Window) GetWmSyncRequestCounterCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_SYNC_REQUEST_COUNTER"), x.AtomCARDINAL, 0, 2)
	return GetWmSyncRequestCounterCookie(cookie)
}

type GetWmSyncRequestCounterCookie x.GetPropertyCookie

func (cookie GetWmSyncRequestCounterCookie) Reply(c *Conn) (WmSyncRequestCounter, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return WmSyncRequestCounter{}, err
	}
	if reply.Type != x.AtomCARDINAL {
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

func SetWmSyncRequestCounterChecked(c *Conn, window x.Window, val WmSyncRequestCounter) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Low)
	w.Write4b(val.High)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_SYNC_REQUEST_COUNTER"), x.AtomCARDINAL, 32, 2, w.Bytes())
}

func SetWmSyncRequestCounter(c *Conn, window x.Window, val WmSyncRequestCounter) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Low)
	w.Write4b(val.High)
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_SYNC_REQUEST_COUNTER"), x.AtomCARDINAL, 32, 2, w.Bytes())
}

/**
 *    _NET_WM_FULLSCREEN_MONITORS
 */

func GetWmFullscreenMonitors(c *Conn, window x.Window) GetWmFullscreenMonitorsCookie {
	cookie := x.GetProperty(c.conn, x.False, window,
		c.GetAtom("_NET_WM_FULLSCREEN_MONITORS"), x.AtomCARDINAL, 0, 4)
	return GetWmFullscreenMonitorsCookie(cookie)
}

func GetWmFullscreenMonitorsUnchecked(c *Conn, window x.Window) GetWmFullscreenMonitorsCookie {
	cookie := x.GetPropertyUnchecked(c.conn, x.False, window,
		c.GetAtom("_NET_WM_FULLSCREEN_MONITORS"), x.AtomCARDINAL, 0, 4)
	return GetWmFullscreenMonitorsCookie(cookie)
}

type GetWmFullscreenMonitorsCookie x.GetPropertyCookie

func (cookie GetWmFullscreenMonitorsCookie) Reply(c *Conn) (WmFullscreenMonitors, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return WmFullscreenMonitors{}, err
	}
	if reply.Type != x.AtomCARDINAL {
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

func SetWmFullscreenMonitorsChecked(c *Conn, window x.Window, val WmFullscreenMonitors) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	return x.ChangePropertyChecked(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_FULLSCREEN_MONITORS"), x.AtomCARDINAL, 32, 4, w.Bytes())
}

func SetWmFullscreenMonitors(c *Conn, window x.Window, val WmFullscreenMonitors) x.VoidCookie {
	w := x.NewWriter()
	w.Write4b(val.Top)
	w.Write4b(val.Bottom)
	w.Write4b(val.Left)
	w.Write4b(val.Right)
	return x.ChangeProperty(c.conn, x.PropModeReplace, window,
		c.GetAtom("_NET_WM_FULLSCREEN_MONITORS"), x.AtomCARDINAL, 32, 4, w.Bytes())
}
