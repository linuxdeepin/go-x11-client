package input

import x "github.com/linuxdeepin/go-x11-client"

func OpenDevice(conn *x.Conn, deviceId uint8) OpenDeviceCookie {
	w := x.NewWriter()
	writeOpenDevice(w, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: OpenDeviceOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return OpenDeviceCookie(seq)
}

func (cookie OpenDeviceCookie) Reply(conn *x.Conn) (*OpenDeviceReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
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
	w := x.NewWriter()
	writeCloseDevice(w, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CloseDeviceOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CloseDeviceChecked(conn *x.Conn, deviceId uint8) x.VoidCookie {
	w := x.NewWriter()
	writeCloseDevice(w, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CloseDeviceOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func SelectExtensionEvent(conn *x.Conn, window x.Window, classes []EventClass) {
	w := x.NewWriter()
	writeSelectExtensionEvent(w, window, classes)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SelectExtensionEventOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SelectExtensionEventChecked(conn *x.Conn, window x.Window, classes []EventClass) x.VoidCookie {
	w := x.NewWriter()
	writeSelectExtensionEvent(w, window, classes)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SelectExtensionEventOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}
