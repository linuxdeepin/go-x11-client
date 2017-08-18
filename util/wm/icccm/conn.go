package icccm

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
		"WM_STATE",
		"WM_PROTOCOLS",
		"WM_COLORMAP_WINDOWS",
	}

	length := len(atomNames)
	cookies := make([]x.InternAtomCookie, length)
	for i := 0; i < length; i++ {
		name := atomNames[i]
		cookies[i] = x.InternAtom(c.conn, x.False, uint16(len(name)), name)
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
