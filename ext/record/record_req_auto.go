package record

import x "github.com/linuxdeepin/go-x11-client"

func QueryVersion(conn *x.Conn, majorVersion, minorVersion uint16) QueryVersionCookie {
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

func CreateContext(conn *x.Conn, context Context, elementHeader ElementHeader, clientSpecs []ClientSpec, ranges []Range) {
	w := x.NewWriter()
	writeCreateContext(w, context, elementHeader, clientSpecs, ranges)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateContextOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CreateContextChecked(conn *x.Conn, context Context, elementHeader ElementHeader, clientSpecs []ClientSpec, ranges []Range) x.VoidCookie {
	w := x.NewWriter()
	writeCreateContext(w, context, elementHeader, clientSpecs, ranges)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateContextOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func RegisterClients(conn *x.Conn, context Context, elementHeader ElementHeader, clientSpecs []ClientSpec, ranges []Range) {
	w := x.NewWriter()
	writeRegisterClients(w, context, elementHeader, clientSpecs, ranges)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  RegisterClientsOpcode,
	}
	conn.SendRequest(0, d, req)
}

func RegisterClientsChecked(conn *x.Conn, context Context, elementHeader ElementHeader, clientSpecs []ClientSpec, ranges []Range) x.VoidCookie {
	w := x.NewWriter()
	writeRegisterClients(w, context, elementHeader, clientSpecs, ranges)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  RegisterClientsOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func UnregisterClients(conn *x.Conn, context Context, clientSpecs []ClientSpec) {
	w := x.NewWriter()
	writeUnregisterClients(w, context, clientSpecs)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  UnregisterClientsOpcode,
	}
	conn.SendRequest(0, d, req)
}

func UnregisterClientsChecked(conn *x.Conn, context Context, clientSpecs []ClientSpec) x.VoidCookie {
	w := x.NewWriter()
	writeUnregisterClients(w, context, clientSpecs)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  UnregisterClientsOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func GetContext(conn *x.Conn, context Context) GetContextCookie {
	w := x.NewWriter()
	writeGetContext(w, context)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: GetContextOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return GetContextCookie(seq)
}

func (cookie GetContextCookie) Reply(conn *x.Conn) (*GetContextReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetContextReply
	err = readGetContextReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func EnableContext(conn *x.Conn, context Context) EnableContextCookie {
	w := x.NewWriter()
	writeEnableContext(w, context)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: EnableContextOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return EnableContextCookie(seq)
}

func (cookie EnableContextCookie) Reply(conn *x.Conn) (*EnableContextReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply EnableContextReply
	err = readEnableContextReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func DisableContext(conn *x.Conn, context Context) {
	w := x.NewWriter()
	writeDisableContext(w, context)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  DisableContextOpcode,
	}
	conn.SendRequest(0, d, req)
}

func DisableContextChecked(conn *x.Conn, context Context) x.VoidCookie {
	w := x.NewWriter()
	writeDisableContext(w, context)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  DisableContextOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func FreeContext(conn *x.Conn, context Context) {
	w := x.NewWriter()
	writeFreeContext(w, context)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  FreeContextOpcode,
	}
	conn.SendRequest(0, d, req)
}

func FreeContextChecked(conn *x.Conn, context Context) x.VoidCookie {
	w := x.NewWriter()
	writeFreeContext(w, context)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  FreeContextOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}
