package atom

import (
	x "github.com/linuxdeepin/go-x11-client"
)

func GetVal(conn *x.Conn, name string) (x.Atom, error) {
	reply, err := x.InternAtom(conn, false, name).Reply(conn)
	if err != nil {
		return 0, err
	}
	return reply.Atom, nil
}

func GetExistingVal(conn *x.Conn, name string) (x.Atom, error) {
	reply, err := x.InternAtom(conn, true, name).Reply(conn)
	if err != nil {
		return 0, err
	}
	return reply.Atom, nil
}

func GetName(conn *x.Conn, val x.Atom) (string, error) {
	reply, err := x.GetAtomName(conn, val).Reply(conn)
	if err != nil {
		return "", err
	}
	return reply.Name, nil
}
