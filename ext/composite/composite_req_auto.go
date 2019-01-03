package composite

import x "github.com/linuxdeepin/go-x11-client"

func QueryVersion(conn *x.Conn, majorVersion, minorVersion uint32) QueryVersionCookie {
	body := encodeQueryVersion(majorVersion, minorVersion)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: QueryVersionOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return QueryVersionCookie(seq)
}

func (cookie QueryVersionCookie) Reply(conn *x.Conn) (*QueryVersionReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply QueryVersionReply
	err = readQueryVersionReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func RedirectWindow(conn *x.Conn, window x.Window, update uint8) {
	body := encodeRedirectWindow(window, update)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: RedirectWindowOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func RedirectWindowChecked(conn *x.Conn, window x.Window, update uint8) x.VoidCookie {
	body := encodeRedirectWindow(window, update)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: RedirectWindowOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func RedirectSubwindows(conn *x.Conn, window x.Window, update uint8) {
	body := encodeRedirectSubwindows(window, update)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: RedirectSubwindowsOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func RedirectSubwindowsChecked(conn *x.Conn, window x.Window, update uint8) x.VoidCookie {
	body := encodeRedirectSubwindows(window, update)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: RedirectSubwindowsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func UnredirectWindow(conn *x.Conn, window x.Window) {
	body := encodeUnredirectWindow(window)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: UnredirectWindowOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func UnredirectWindowChecked(conn *x.Conn, window x.Window) x.VoidCookie {
	body := encodeUnredirectWindow(window)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: UnredirectWindowOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func UnredirectSubwindows(conn *x.Conn, window x.Window) {
	body := encodeUnredirectSubwindows(window)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: UnredirectSubwindowsOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func UnredirectSubwindowsChecked(conn *x.Conn, window x.Window) x.VoidCookie {
	body := encodeUnredirectSubwindows(window)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: UnredirectSubwindowsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func NameWindowPixmap(conn *x.Conn, window x.Window, pixmap x.Pixmap) {
	body := encodeNameWindowPixmap(window, pixmap)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: NameWindowPixmapOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func NameWindowPixmapChecked(conn *x.Conn, window x.Window, pixmap x.Pixmap) x.VoidCookie {
	body := encodeNameWindowPixmap(window, pixmap)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: NameWindowPixmapOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}
