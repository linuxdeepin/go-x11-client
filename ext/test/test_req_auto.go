// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package test

import x "github.com/linuxdeepin/go-x11-client"

func GetVersion(conn *x.Conn, majorVersion uint8, minorVersion uint16) GetVersionCookie {
	body := encodeGetVersion(majorVersion, minorVersion)
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

func CompareCursor(conn *x.Conn, window x.Window, cursor x.Cursor) CompareCursorCookie {
	body := encodeCompareCursor(window, cursor)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: CompareCursorOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return CompareCursorCookie(seq)
}

func (cookie CompareCursorCookie) Reply(conn *x.Conn) (*CompareCursorReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply CompareCursorReply
	err = readCompareCursorReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func FakeInput(conn *x.Conn, evType uint8, detail uint8, time x.Timestamp, root x.Window, rootX, rootY int16, deviceId uint8) {
	body := encodeFakeInput(evType, detail, time, root, rootX, rootY, deviceId)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: FakeInputOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func FakeInputChecked(conn *x.Conn, evType uint8, detail uint8, time x.Timestamp, root x.Window, rootX, rootY int16, deviceId uint8) x.VoidCookie {
	body := encodeFakeInput(evType, detail, time, root, rootX, rootY, deviceId)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: FakeInputOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func GrabControl(conn *x.Conn, impervious bool) {
	body := encodeGrabControl(impervious)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: GrabControlOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func GrabControlChecked(conn *x.Conn, impervious bool) x.VoidCookie {
	body := encodeGrabControl(impervious)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: GrabControlOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}
