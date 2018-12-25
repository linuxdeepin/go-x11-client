package input

import x "github.com/linuxdeepin/go-x11-client"

func XIQueryVersion(conn *x.Conn, majorVersion, minorVersion uint16) XIQueryVersionCookie {
	body := encodeXIQueryVersion(majorVersion, minorVersion)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: XIQueryVersionOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return XIQueryVersionCookie(seq)
}

func (cookie XIQueryVersionCookie) Reply(conn *x.Conn) (*XIQueryVersionReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply XIQueryVersionReply
	err = readXIQueryVersionReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func XIQueryDevice(conn *x.Conn, deviceId DeviceId) XIQueryDeviceCookie {
	body := encodeXIQueryDevice(deviceId)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: XIQueryDeviceOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return XIQueryDeviceCookie(seq)
}

func (cookie XIQueryDeviceCookie) Reply(conn *x.Conn) (*XIQueryDeviceReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply XIQueryDeviceReply
	err = readXIQueryDeviceReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func XISelectEvents(conn *x.Conn, window x.Window, masks []EventMask) {
	body := encodeXISelectEvents(window, masks)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XISelectEventsOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func XISelectEventsChecked(conn *x.Conn, window x.Window, masks []EventMask) x.VoidCookie {
	body := encodeXISelectEvents(window, masks)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XISelectEventsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func XIGetSelectedEvents(conn *x.Conn, window x.Window) XIGetSelectedEventsCookie {
	body := encodeXIGetSelectedEvents(window)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: XIGetSelectedEventsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return XIGetSelectedEventsCookie(seq)
}

func (cookie XIGetSelectedEventsCookie) Reply(conn *x.Conn) (*XIGetSelectedEventsReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply XIGetSelectedEventsReply
	err = readXIGetSelectedEventsReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func XIQueryPointer(conn *x.Conn, window x.Window, deviceId DeviceId) XIQueryPointerCookie {
	body := encodeXIQueryPointer(window, deviceId)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: XIQueryPointerOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return XIQueryPointerCookie(seq)
}

func (cookie XIQueryPointerCookie) Reply(conn *x.Conn) (*XIQueryPointerReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply XIQueryPointerReply
	err = readXIQueryPointerReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func XIWarpPointer(conn *x.Conn, srcWin, dstWin x.Window, srcX, srcY float32, srcWidth, srcHeight uint16, dstX, dstY float32, deviceId DeviceId) {
	body := encodeXIWarpPointer(srcWin, dstWin, srcX, srcY, srcWidth, srcHeight, dstX, dstY, deviceId)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XIWarpPointerOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func XIWarpPointerChecked(conn *x.Conn, srcWin, dstWin x.Window, srcX, srcY float32, srcWidth, srcHeight uint16, dstX, dstY float32, deviceId DeviceId) x.VoidCookie {
	body := encodeXIWarpPointer(srcWin, dstWin, srcX, srcY, srcWidth, srcHeight, dstX, dstY, deviceId)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XIWarpPointerOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func XIChangeCursor(conn *x.Conn, window x.Window, cursor x.Cursor, deviceId DeviceId) {
	body := encodeXIChangeCursor(window, cursor, deviceId)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XIChangeCursorOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func XIChangeCursorChecked(conn *x.Conn, window x.Window, cursor x.Cursor, deviceId DeviceId) x.VoidCookie {
	body := encodeXIChangeCursor(window, cursor, deviceId)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XIChangeCursorOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func XIChangeHierarchy(conn *x.Conn, changes []HierarchyChange) {
	body := encodeXIChangeHierarchy(changes)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XIChangeHierarchyOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func XIChangeHierarchyChecked(conn *x.Conn, changes []HierarchyChange) x.VoidCookie {
	body := encodeXIChangeHierarchy(changes)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XIChangeHierarchyOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func XISetClientPointer(conn *x.Conn, window x.Window, deviceId DeviceId) {
	body := encodeXISetClientPointer(window, deviceId)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XISetClientPointerOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func XISetClientPointerChecked(conn *x.Conn, window x.Window, deviceId DeviceId) x.VoidCookie {
	body := encodeXISetClientPointer(window, deviceId)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XISetClientPointerOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func XIGetClientPointer(conn *x.Conn, window x.Window) XIGetClientPointerCookie {
	body := encodeXIGetClientPointer(window)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: XIGetClientPointerOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return XIGetClientPointerCookie(seq)
}

func (cookie XIGetClientPointerCookie) Reply(conn *x.Conn) (*XIGetClientPointerReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply XIGetClientPointerReply
	err = readXIGetClientPointerReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func XISetFocus(conn *x.Conn, focus x.Window, time x.Timestamp, deviceId DeviceId) {
	body := encodeXISetFocus(focus, time, deviceId)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XISetFocusOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func XISetFocusChecked(conn *x.Conn, focus x.Window, time x.Timestamp, deviceId DeviceId) x.VoidCookie {
	body := encodeXISetFocus(focus, time, deviceId)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XISetFocusOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func XIGetFocus(conn *x.Conn, deviceId DeviceId) XIGetFocusCookie {
	body := encodeXIGetFocus(deviceId)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: XIGetFocusOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return XIGetFocusCookie(seq)
}

func (cookie XIGetFocusCookie) Reply(conn *x.Conn) (*XIGetFocusReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply XIGetFocusReply
	err = readXIGetFocusReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func XIGrabDevice(conn *x.Conn, window x.Window, time x.Timestamp, cursor x.Cursor, deviceId DeviceId, mode, pairedDeviceMode uint8, ownerEvents bool, masks []EventMask) XIGrabDeviceCookie {
	body := encodeXIGrabDevice(window, time, cursor, deviceId, mode, pairedDeviceMode, ownerEvents, masks)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: XIGrabDeviceOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return XIGrabDeviceCookie(seq)
}

func (cookie XIGrabDeviceCookie) Reply(conn *x.Conn) (*XIGrabDeviceReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply XIGrabDeviceReply
	err = readXIGrabDeviceReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func XIUngrabDevice(conn *x.Conn, time x.Timestamp, deviceId DeviceId) {
	body := encodeXIUngrabDevice(time, deviceId)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XIUngrabDeviceOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func XIUngrabDeviceChecked(conn *x.Conn, time x.Timestamp, deviceId DeviceId) x.VoidCookie {
	body := encodeXIUngrabDevice(time, deviceId)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XIUngrabDeviceOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func XIAllowEvents(conn *x.Conn, time x.Timestamp, deviceId DeviceId, eventMode uint8, touchId uint32, grabWindow x.Window) {
	body := encodeXIAllowEvents(time, deviceId, eventMode, touchId, grabWindow)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XIAllowEventsOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func XIAllowEventsChecked(conn *x.Conn, time x.Timestamp, deviceId DeviceId, eventMode uint8, touchId uint32, grabWindow x.Window) x.VoidCookie {
	body := encodeXIAllowEvents(time, deviceId, eventMode, touchId, grabWindow)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XIAllowEventsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func XIPassiveGrabDevice(conn *x.Conn, grabWindow x.Window, cursor x.Cursor, detail uint32, deviceId DeviceId, grabType, grabMode, pairedDeviceMode uint8, ownerEvents bool, masks []uint32, modifiers []uint32) XIPassiveGrabDeviceCookie {
	body := encodeXIPassiveGrabDevice(grabWindow, cursor, detail, deviceId, grabType, grabMode, pairedDeviceMode, ownerEvents, masks, modifiers)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: XIPassiveGrabDeviceOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return XIPassiveGrabDeviceCookie(seq)
}

func (cookie XIPassiveGrabDeviceCookie) Reply(conn *x.Conn) (*XIPassiveGrabDeviceReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply XIPassiveGrabDeviceReply
	err = readXIPassiveGrabDeviceReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func XIPassiveUngrabDevice(conn *x.Conn, grabWindow x.Window, detail uint32, deviceId DeviceId, grabType uint8, modifiers []uint32) {
	body := encodeXIPassiveUngrabDevice(grabWindow, detail, deviceId, grabType, modifiers)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XIPassiveUngrabDeviceOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func XIPassiveUngrabDeviceChecked(conn *x.Conn, grabWindow x.Window, detail uint32, deviceId DeviceId, grabType uint8, modifiers []uint32) x.VoidCookie {
	body := encodeXIPassiveUngrabDevice(grabWindow, detail, deviceId, grabType, modifiers)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XIPassiveUngrabDeviceOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func XIListProperties(conn *x.Conn, deviceId DeviceId) XIListPropertiesCookie {
	body := encodeXIListProperties(deviceId)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: XIListPropertiesOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return XIListPropertiesCookie(seq)
}

func (cookie XIListPropertiesCookie) Reply(conn *x.Conn) (*XIListPropertiesReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply XIListPropertiesReply
	err = readXIListPropertiesReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func XIChangeProperty(conn *x.Conn, deviceId DeviceId, mode uint8, format uint8, property x.Atom, Type x.Atom, data []byte) {
	body := encodeXIChangeProperty(deviceId, mode, format, property, Type, data)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XIChangePropertyOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func XIChangePropertyChecked(conn *x.Conn, deviceId DeviceId, mode uint8, format uint8, property x.Atom, Type x.Atom, data []byte) x.VoidCookie {
	body := encodeXIChangeProperty(deviceId, mode, format, property, Type, data)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XIChangePropertyOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func XIDeleteProperty(conn *x.Conn, deviceId DeviceId, property x.Atom) {
	body := encodeXIDeleteProperty(deviceId, property)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XIDeletePropertyOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func XIDeletePropertyChecked(conn *x.Conn, deviceId DeviceId, property x.Atom) x.VoidCookie {
	body := encodeXIDeleteProperty(deviceId, property)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: XIDeletePropertyOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func XIGetProperty(conn *x.Conn, deviceId DeviceId, delete bool, property x.Atom, Type x.Atom, offset, len uint32) XIGetPropertyCookie {
	body := encodeXIGetProperty(deviceId, delete, property, Type, offset, len)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: XIGetPropertyOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return XIGetPropertyCookie(seq)
}

func (cookie XIGetPropertyCookie) Reply(conn *x.Conn) (*XIGetPropertyReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply XIGetPropertyReply
	err = readXIGetPropertyReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
