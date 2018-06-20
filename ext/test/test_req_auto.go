package test

import x "github.com/linuxdeepin/go-x11-client"

func GetVersion(conn *x.Conn, majorVersion uint8, minorVersion uint16) GetVersionCookie {
	w := x.NewWriter()
	writeGetVersion(w, majorVersion, minorVersion)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: GetVersionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return GetVersionCookie(seq)
}

func (cookie GetVersionCookie) Reply(conn *x.Conn) (*GetVersionReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, conn.NewError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetVersionReply
	err := readGetVersionReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func CompareCursor(conn *x.Conn, window x.Window, cursor x.Cursor) CompareCursorCookie {
	w := x.NewWriter()
	writeCompareCursor(w, window, cursor)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: CompareCursorOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return CompareCursorCookie(seq)
}

func (cookie CompareCursorCookie) Reply(conn *x.Conn) (*CompareCursorReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, conn.NewError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply CompareCursorReply
	err := readCompareCursorReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func FakeInput(conn *x.Conn, evType uint8, detail uint8, time x.Timestamp, root x.Window, rootX, rootY int16, deviceId uint8) {
	w := x.NewWriter()
	writeFakeInput(w, evType, detail, time, root, rootX, rootY, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  FakeInputOpcode,
	}
	conn.SendRequest(0, d, req)
}

func FakeInputChecked(conn *x.Conn, evType uint8, detail uint8, time x.Timestamp, root x.Window, rootX, rootY int16, deviceId uint8) x.VoidCookie {
	w := x.NewWriter()
	writeFakeInput(w, evType, detail, time, root, rootX, rootY, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  FakeInputOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func GrabControl(conn *x.Conn, impervious bool) {
	w := x.NewWriter()
	writeGrabControl(w, impervious)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  GrabControlOpcode,
	}
	conn.SendRequest(0, d, req)
}

func GrabControlChecked(conn *x.Conn, impervious bool) x.VoidCookie {
	w := x.NewWriter()
	writeGrabControl(w, impervious)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  GrabControlOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}
