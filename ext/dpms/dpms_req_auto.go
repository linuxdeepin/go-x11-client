package dpms

import x "github.com/linuxdeepin/go-x11-client"

func GetVersion(conn *x.Conn, clientMajorVersion, clientMinorVersion uint16) GetVersionCookie {
	w := x.NewWriter()
	writeGetVersion(w, clientMajorVersion, clientMinorVersion)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: GetVersionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return GetVersionCookie(seq)
}

func (cookie GetVersionCookie) Reply(conn *x.Conn) (*GetVersionReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetVersionReply
	err = readGetVersionReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func Capable(conn *x.Conn) CapableCookie {
	w := x.NewWriter()
	writeCapable(w)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: CapableOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return CapableCookie(seq)
}

func (cookie CapableCookie) Reply(conn *x.Conn) (*CapableReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply CapableReply
	err = readCapableReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetTimeouts(conn *x.Conn) GetTimeoutsCookie {
	w := x.NewWriter()
	writeGetTimeouts(w)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: GetTimeoutsOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return GetTimeoutsCookie(seq)
}

func (cookie GetTimeoutsCookie) Reply(conn *x.Conn) (*GetTimeoutsReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetTimeoutsReply
	err = readGetTimeoutsReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetTimeouts(conn *x.Conn, standbyTimeout, suspendTimeout, offTimeout uint16) {
	w := x.NewWriter()
	writeSetTimeouts(w, standbyTimeout, suspendTimeout, offTimeout)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SetTimeoutsOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SetTimeoutsChecked(conn *x.Conn, standbyTimeout, suspendTimeout, offTimeout uint16) x.VoidCookie {
	w := x.NewWriter()
	writeSetTimeouts(w, standbyTimeout, suspendTimeout, offTimeout)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SetTimeoutsOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func Enable(conn *x.Conn) {
	w := x.NewWriter()
	writeEnable(w)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  EnableOpcode,
	}
	conn.SendRequest(0, d, req)
}

func EnableChecked(conn *x.Conn) x.VoidCookie {
	w := x.NewWriter()
	writeEnable(w)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  EnableOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func Disable(conn *x.Conn) {
	w := x.NewWriter()
	writeDisable(w)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  DisableOpcode,
	}
	conn.SendRequest(0, d, req)
}

func DisableChecked(conn *x.Conn) x.VoidCookie {
	w := x.NewWriter()
	writeDisable(w)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  DisableOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func ForceLevel(conn *x.Conn, powerLevel uint16) {
	w := x.NewWriter()
	writeForceLevel(w, powerLevel)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  ForceLevelOpcode,
	}
	conn.SendRequest(0, d, req)
}

func ForceLevelChecked(conn *x.Conn, powerLevel uint16) x.VoidCookie {
	w := x.NewWriter()
	writeForceLevel(w, powerLevel)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  ForceLevelOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func Info(conn *x.Conn) InfoCookie {
	w := x.NewWriter()
	writeInfo(w)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: InfoOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return InfoCookie(seq)
}

func (cookie InfoCookie) Reply(conn *x.Conn) (*InfoReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply InfoReply
	err = readInfoReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
