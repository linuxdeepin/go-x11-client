package composite

import x "github.com/linuxdeepin/go-x11-client"

func QueryVersion(conn *x.Conn, majorVersion, minorVersion uint32) QueryVersionCookie {
	w := x.NewWriter()
	writeQueryVersion(w, majorVersion, minorVersion)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: QueryVersionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return QueryVersionCookie(seq)
}

func (cookie QueryVersionCookie) Reply(conn *x.Conn) (*QueryVersionReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply QueryVersionReply
	err := readQueryVersionReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func RedirectWindow(conn *x.Conn, window x.Window, update uint8) {
	w := x.NewWriter()
	writeRedirectWindow(w, window, update)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  RedirectWindowOpcode,
	}
	conn.SendRequest(0, d, req)
}

func RedirectWindowChecked(conn *x.Conn, window x.Window, update uint8) x.VoidCookie {
	w := x.NewWriter()
	writeRedirectWindow(w, window, update)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  RedirectWindowOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func RedirectSubwindows(conn *x.Conn, window x.Window, update uint8) {
	w := x.NewWriter()
	writeRedirectSubwindows(w, window, update)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  RedirectSubwindowsOpcode,
	}
	conn.SendRequest(0, d, req)
}

func RedirectSubwindowsChecked(conn *x.Conn, window x.Window, update uint8) x.VoidCookie {
	w := x.NewWriter()
	writeRedirectSubwindows(w, window, update)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  RedirectSubwindowsOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func UnredirectWindow(conn *x.Conn, window x.Window) {
	w := x.NewWriter()
	writeUnredirectWindow(w, window)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  UnredirectWindowOpcode,
	}
	conn.SendRequest(0, d, req)
}

func UnredirectWindowChecked(conn *x.Conn, window x.Window) x.VoidCookie {
	w := x.NewWriter()
	writeUnredirectWindow(w, window)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  UnredirectWindowOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func UnredirectSubwindows(conn *x.Conn, window x.Window) {
	w := x.NewWriter()
	writeUnredirectSubwindows(w, window)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  UnredirectSubwindowsOpcode,
	}
	conn.SendRequest(0, d, req)
}

func UnredirectSubwindowsChecked(conn *x.Conn, window x.Window) x.VoidCookie {
	w := x.NewWriter()
	writeUnredirectSubwindows(w, window)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  UnredirectSubwindowsOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func NameWindowPixmap(conn *x.Conn, window x.Window, pixmap x.Pixmap) {
	w := x.NewWriter()
	writeNameWindowPixmap(w, window, pixmap)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  NameWindowPixmapOpcode,
	}
	conn.SendRequest(0, d, req)
}

func NameWindowPixmapChecked(conn *x.Conn, window x.Window, pixmap x.Pixmap) x.VoidCookie {
	w := x.NewWriter()
	writeNameWindowPixmap(w, window, pixmap)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  NameWindowPixmapOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}
