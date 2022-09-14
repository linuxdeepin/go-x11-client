// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package screensaver

import x "github.com/linuxdeepin/go-x11-client"

func QueryVersion(conn *x.Conn, clientMajorVersion, clientMinorVersion uint8) QueryVersionCookie {
	body := encodeQueryVersion(clientMajorVersion, clientMinorVersion)
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

func QueryInfo(conn *x.Conn, drawable x.Drawable) QueryInfoCookie {
	body := encodeQueryInfo(drawable)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: QueryInfoOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return QueryInfoCookie(seq)
}

func (cookie QueryInfoCookie) Reply(conn *x.Conn) (*QueryInfoReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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
	body := encodeSelectInput(drawable, eventMask)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SelectInputOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SelectInputChecked(conn *x.Conn, drawable x.Drawable, eventMask uint32) x.VoidCookie {
	body := encodeSelectInput(drawable, eventMask)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SelectInputOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func SetAttributes(conn *x.Conn, drawable x.Drawable, X, y int16, width, height, boardWidth uint16, class, depth uint8, visual x.VisualID, valueMask uint32, valueList []uint32) {
	body := encodeSetAttributes(drawable, X, y, width, height, boardWidth, class, depth, visual, valueMask, valueList)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetAttributesOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetAttributesChecked(conn *x.Conn, drawable x.Drawable, X, y int16, width, height, boardWidth uint16, class, depth uint8, visual x.VisualID, valueMask uint32, valueList []uint32) x.VoidCookie {
	body := encodeSetAttributes(drawable, X, y, width, height, boardWidth, class, depth, visual, valueMask, valueList)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetAttributesOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func UnsetAttributes(conn *x.Conn, drawable x.Drawable) {
	body := encodeUnsetAttributes(drawable)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: UnsetAttributesOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func UnsetAttributesChecked(conn *x.Conn, drawable x.Drawable) x.VoidCookie {
	body := encodeUnsetAttributes(drawable)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: UnsetAttributesOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func Suspend(conn *x.Conn, suspend bool) {
	body := encodeSuspend(suspend)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SuspendOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SuspendChecked(conn *x.Conn, suspend bool) x.VoidCookie {
	body := encodeSuspend(suspend)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SuspendOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}
