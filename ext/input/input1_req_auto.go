// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package input

import x "github.com/linuxdeepin/go-x11-client"

func OpenDevice(conn *x.Conn, deviceId uint8) OpenDeviceCookie {
	body := encodeOpenDevice(deviceId)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: OpenDeviceOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return OpenDeviceCookie(seq)
}

func (cookie OpenDeviceCookie) Reply(conn *x.Conn) (*OpenDeviceReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply OpenDeviceReply
	err = readOpenDeviceReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func CloseDevice(conn *x.Conn, deviceId uint8) {
	body := encodeCloseDevice(deviceId)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CloseDeviceOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CloseDeviceChecked(conn *x.Conn, deviceId uint8) x.VoidCookie {
	body := encodeCloseDevice(deviceId)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CloseDeviceOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func SelectExtensionEvent(conn *x.Conn, window x.Window, classes []EventClass) {
	body := encodeSelectExtensionEvent(window, classes)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SelectExtensionEventOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SelectExtensionEventChecked(conn *x.Conn, window x.Window, classes []EventClass) x.VoidCookie {
	body := encodeSelectExtensionEvent(window, classes)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SelectExtensionEventOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}
