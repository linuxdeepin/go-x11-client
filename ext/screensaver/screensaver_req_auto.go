package screensaver

import x "github.com/linuxdeepin/go-x11-client"

func QueryVersion(conn *x.Conn, clientMajorVersion, clientMinorVersion uint8) QueryVersionCookie {
	w := x.NewWriter()
	writeQueryVersion(w, clientMajorVersion, clientMinorVersion)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: QueryVersionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return QueryVersionCookie(seq)
}

func (cookie QueryVersionCookie) Reply(conn *x.Conn) (*QueryVersionReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
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

func QueryInfo(conn *x.Conn, drawable x.Drawable) QueryInfoCookie {
	w := x.NewWriter()
	writeQueryInfo(w, drawable)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: QueryInfoOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return QueryInfoCookie(seq)
}

func (cookie QueryInfoCookie) Reply(conn *x.Conn) (*QueryInfoReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply QueryInfoReply
	err = readQueryInfoReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SelectInput(conn *x.Conn, drawable x.Drawable, eventMask uint32) {
	w := x.NewWriter()
	writeSelectInput(w, drawable, eventMask)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SelectInputOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SelectInputChecked(conn *x.Conn, drawable x.Drawable, eventMask uint32) x.VoidCookie {
	w := x.NewWriter()
	writeSelectInput(w, drawable, eventMask)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SelectInputOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func SetAttributes(conn *x.Conn, drawable x.Drawable, X, y int16, width, height, boardWidth uint16, class, depth uint8, visual x.VisualID, valueMask uint32, valueList []uint32) {
	w := x.NewWriter()
	writeSetAttributes(w, drawable, X, y, width, height, boardWidth, class, depth, visual, valueMask, valueList)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SetAttributesOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SetAttributesChecked(conn *x.Conn, drawable x.Drawable, X, y int16, width, height, boardWidth uint16, class, depth uint8, visual x.VisualID, valueMask uint32, valueList []uint32) x.VoidCookie {
	w := x.NewWriter()
	writeSetAttributes(w, drawable, X, y, width, height, boardWidth, class, depth, visual, valueMask, valueList)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SetAttributesOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func UnsetAttributes(conn *x.Conn, drawable x.Drawable) {
	w := x.NewWriter()
	writeUnsetAttributes(w, drawable)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  UnsetAttributesOpcode,
	}
	conn.SendRequest(0, d, req)
}

func UnsetAttributesChecked(conn *x.Conn, drawable x.Drawable) x.VoidCookie {
	w := x.NewWriter()
	writeUnsetAttributes(w, drawable)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  UnsetAttributesOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func Suspend(conn *x.Conn, suspend bool) {
	w := x.NewWriter()
	writeSuspend(w, suspend)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SuspendOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SuspendChecked(conn *x.Conn, suspend bool) x.VoidCookie {
	w := x.NewWriter()
	writeSuspend(w, suspend)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SuspendOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}
