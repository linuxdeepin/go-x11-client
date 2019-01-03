package dpms

import x "github.com/linuxdeepin/go-x11-client"

func GetVersion(conn *x.Conn, clientMajorVersion, clientMinorVersion uint16) GetVersionCookie {
	body := encodeGetVersion(clientMajorVersion, clientMinorVersion)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetVersionOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetVersionCookie(seq)
}

func (cookie GetVersionCookie) Reply(conn *x.Conn) (*GetVersionReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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
	body := encodeCapable()
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: CapableOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return CapableCookie(seq)
}

func (cookie CapableCookie) Reply(conn *x.Conn) (*CapableReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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
	body := encodeGetTimeouts()
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetTimeoutsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetTimeoutsCookie(seq)
}

func (cookie GetTimeoutsCookie) Reply(conn *x.Conn) (*GetTimeoutsReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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
	body := encodeSetTimeouts(standbyTimeout, suspendTimeout, offTimeout)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetTimeoutsOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetTimeoutsChecked(conn *x.Conn, standbyTimeout, suspendTimeout, offTimeout uint16) x.VoidCookie {
	body := encodeSetTimeouts(standbyTimeout, suspendTimeout, offTimeout)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetTimeoutsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func Enable(conn *x.Conn) {
	body := encodeEnable()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: EnableOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func EnableChecked(conn *x.Conn) x.VoidCookie {
	body := encodeEnable()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: EnableOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func Disable(conn *x.Conn) {
	body := encodeDisable()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: DisableOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func DisableChecked(conn *x.Conn) x.VoidCookie {
	body := encodeDisable()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: DisableOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func ForceLevel(conn *x.Conn, powerLevel uint16) {
	body := encodeForceLevel(powerLevel)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ForceLevelOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func ForceLevelChecked(conn *x.Conn, powerLevel uint16) x.VoidCookie {
	body := encodeForceLevel(powerLevel)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ForceLevelOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func Info(conn *x.Conn) InfoCookie {
	body := encodeInfo()
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: InfoOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return InfoCookie(seq)
}

func (cookie InfoCookie) Reply(conn *x.Conn) (*InfoReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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
