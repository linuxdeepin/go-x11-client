package ewmh

import (
	x "github.com/linuxdeepin/go-x11-client"
)

type Conn struct {
	conn          *x.Conn
	defaultScreen *x.Screen
	atoms         map[string]x.Atom
}

func NewConn(conn *x.Conn) (*Conn, error) {
	c := &Conn{
		conn:          conn,
		defaultScreen: conn.GetDefaultScreen(),
		atoms:         make(map[string]x.Atom),
	}

	err := c.initAtoms()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Conn) initAtoms() error {
	atomNames := []string{
		"UTF8_STRING",
		"WM_PROTOCOLS",
		"_NET_ACTIVE_WINDOW",
		"_NET_CLIENT_LIST",
		"_NET_CLIENT_LIST_STACKING",
		"_NET_CLOSE_WINDOW",
		"_NET_CURRENT_DESKTOP",
		"_NET_DESKTOP_GEOMETRY",
		"_NET_DESKTOP_LAYOUT",
		"_NET_DESKTOP_NAMES",
		"_NET_DESKTOP_VIEWPORT",
		"_NET_FRAME_EXTENTS",
		"_NET_MOVERESIZE_WINDOW",
		"_NET_NUMBER_OF_DESKTOPS",
		"_NET_RESTACK_WINDOW",
		"_NET_SHOWING_DESKTOP",
		"_NET_SUPPORTED",
		"_NET_SUPPORTING_WM_CHECK",
		"_NET_VIRTUAL_ROOTS",
		"_NET_WM_ALLOWED_ACTIONS",
		"_NET_WM_DESKTOP",
		"_NET_WM_FULLSCREEN_MONITORS",
		"_NET_WM_HANDLED_ICONS",
		"_NET_WM_ICON",
		"_NET_WM_ICON_GEOMETRY",
		"_NET_WM_ICON_NAME",
		"_NET_WM_MOVERESIZE",
		"_NET_WM_NAME",
		"_NET_WM_PID",
		"_NET_WM_PING",
		"_NET_WM_STATE",
		"_NET_WM_STRUT",
		"_NET_WM_STRUT_PARTIAL",
		"_NET_WM_SYNC_REQUEST",
		"_NET_WM_SYNC_REQUEST_COUNTER",
		"_NET_WM_USER_TIME",
		"_NET_WM_USER_TIME_WINDOW",
		"_NET_WM_VISIBLE_ICON_NAME",
		"_NET_WM_VISIBLE_NAME",
		"_NET_WM_WINDOW_TYPE",
		"_NET_WORKAREA",

		// window types
		"_NET_WM_WINDOW_TYPE_DESKTOP",
		"_NET_WM_WINDOW_TYPE_DOCK",
		"_NET_WM_WINDOW_TYPE_TOOLBAR",
		"_NET_WM_WINDOW_TYPE_MENU",
		"_NET_WM_WINDOW_TYPE_UTILITY",
		"_NET_WM_WINDOW_TYPE_SPLASH",
		"_NET_WM_WINDOW_TYPE_DIALOG",
		"_NET_WM_WINDOW_TYPE_DROPDOWN_MENU",
		"_NET_WM_WINDOW_TYPE_POPUP_MENU",
		"_NET_WM_WINDOW_TYPE_TOOLTIP",
		"_NET_WM_WINDOW_TYPE_NOTIFICATION",
		"_NET_WM_WINDOW_TYPE_COMBO",
		"_NET_WM_WINDOW_TYPE_DND",
		"_NET_WM_WINDOW_TYPE_NORMAL",

		// states
		"_NET_WM_STATE_MODAL",
		"_NET_WM_STATE_STICKY",
		"_NET_WM_STATE_MAXIMIZED_VERT",
		"_NET_WM_STATE_MAXIMIZED_HORZ",
		"_NET_WM_STATE_SHADED",
		"_NET_WM_STATE_SKIP_TASKBAR",
		"_NET_WM_STATE_SKIP_PAGER",
		"_NET_WM_STATE_HIDDEN",
		"_NET_WM_STATE_FULLSCREEN",
		"_NET_WM_STATE_ABOVE",
		"_NET_WM_STATE_BELOW",
		"_NET_WM_STATE_DEMANDS_ATTENTION",

		// allowed actions
		"_NET_WM_ACTION_MOVE",
		"_NET_WM_ACTION_RESIZE",
		"_NET_WM_ACTION_MINIMIZE",
		"_NET_WM_ACTION_SHADE",
		"_NET_WM_ACTION_STICK",
		"_NET_WM_ACTION_MAXIMIZE_HORZ",
		"_NET_WM_ACTION_MAXIMIZE_VERT",
		"_NET_WM_ACTION_FULLSCREEN",
		"_NET_WM_ACTION_CHANGE_DESKTOP",
		"_NET_WM_ACTION_CLOSE",
		"_NET_WM_ACTION_ABOVE",
		"_NET_WM_ACTION_BELOW",
	}

	length := len(atomNames)
	cookies := make([]x.InternAtomCookie, length)
	for i := 0; i < length; i++ {
		name := atomNames[i]
		cookies[i] = x.InternAtom(c.conn, false, name)
	}

	for i := 0; i < length; i++ {
		reply, err := cookies[i].Reply(c.conn)
		if err != nil {
			return err
		}
		// TODO
		//if reply.Atom == 0 {
		//return er
		//}
		c.atoms[atomNames[i]] = reply.Atom
	}

	return nil
}

func (c *Conn) GetAtom(name string) x.Atom {
	a, ok := c.atoms[name]
	if !ok {
		panic("atom not found" + name)
	}
	return a
}

func (c *Conn) GetRootWin() x.Window {
	return c.defaultScreen.Root
}
