// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package record

import x "github.com/linuxdeepin/go-x11-client"

func QueryVersion(conn *x.Conn, majorVersion, minorVersion uint16) QueryVersionCookie {
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

func CreateContext(conn *x.Conn, context Context, elementHeader ElementHeader, clientSpecs []ClientSpec, ranges []Range) {
	body := encodeCreateContext(context, elementHeader, clientSpecs, ranges)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateContextOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CreateContextChecked(conn *x.Conn, context Context, elementHeader ElementHeader, clientSpecs []ClientSpec, ranges []Range) x.VoidCookie {
	body := encodeCreateContext(context, elementHeader, clientSpecs, ranges)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreateContextOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func RegisterClients(conn *x.Conn, context Context, elementHeader ElementHeader, clientSpecs []ClientSpec, ranges []Range) {
	body := encodeRegisterClients(context, elementHeader, clientSpecs, ranges)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: RegisterClientsOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func RegisterClientsChecked(conn *x.Conn, context Context, elementHeader ElementHeader, clientSpecs []ClientSpec, ranges []Range) x.VoidCookie {
	body := encodeRegisterClients(context, elementHeader, clientSpecs, ranges)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: RegisterClientsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func UnregisterClients(conn *x.Conn, context Context, clientSpecs []ClientSpec) {
	body := encodeUnregisterClients(context, clientSpecs)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: UnregisterClientsOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func UnregisterClientsChecked(conn *x.Conn, context Context, clientSpecs []ClientSpec) x.VoidCookie {
	body := encodeUnregisterClients(context, clientSpecs)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: UnregisterClientsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func GetContext(conn *x.Conn, context Context) GetContextCookie {
	body := encodeGetContext(context)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetContextOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetContextCookie(seq)
}

func (cookie GetContextCookie) Reply(conn *x.Conn) (*GetContextReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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
	body := encodeEnableContext(context)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: EnableContextOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return EnableContextCookie(seq)
}

func (cookie EnableContextCookie) Reply(conn *x.Conn) (*EnableContextReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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
	body := encodeDisableContext(context)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: DisableContextOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func DisableContextChecked(conn *x.Conn, context Context) x.VoidCookie {
	body := encodeDisableContext(context)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: DisableContextOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func FreeContext(conn *x.Conn, context Context) {
	body := encodeFreeContext(context)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: FreeContextOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func FreeContextChecked(conn *x.Conn, context Context) x.VoidCookie {
	body := encodeFreeContext(context)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: FreeContextOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}
