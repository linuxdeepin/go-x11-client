package input

import x "github.com/linuxdeepin/go-x11-client"

func XIQueryVersion(conn *x.Conn, majorVersion, minorVersion uint16) XIQueryVersionCookie {
	w := x.NewWriter()
	writeXIQueryVersion(w, majorVersion, minorVersion)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: XIQueryVersionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
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
	w := x.NewWriter()
	writeXIQueryDevice(w, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: XIQueryDeviceOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
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
	w := x.NewWriter()
	writeXISelectEvents(w, window, masks)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XISelectEventsOpcode,
	}
	conn.SendRequest(0, d, req)
}

func XISelectEventsChecked(conn *x.Conn, window x.Window, masks []EventMask) x.VoidCookie {
	w := x.NewWriter()
	writeXISelectEvents(w, window, masks)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XISelectEventsOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func XIGetSelectedEvents(conn *x.Conn, window x.Window) XIGetSelectedEventsCookie {
	w := x.NewWriter()
	writeXIGetSelectedEvents(w, window)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: XIGetSelectedEventsOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
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
	w := x.NewWriter()
	writeXIQueryPointer(w, window, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: XIQueryPointerOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
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
	w := x.NewWriter()
	writeXIWarpPointer(w, srcWin, dstWin, srcX, srcY, srcWidth, srcHeight, dstX, dstY, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XIWarpPointerOpcode,
	}
	conn.SendRequest(0, d, req)
}

func XIWarpPointerChecked(conn *x.Conn, srcWin, dstWin x.Window, srcX, srcY float32, srcWidth, srcHeight uint16, dstX, dstY float32, deviceId DeviceId) x.VoidCookie {
	w := x.NewWriter()
	writeXIWarpPointer(w, srcWin, dstWin, srcX, srcY, srcWidth, srcHeight, dstX, dstY, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XIWarpPointerOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func XIChangeCursor(conn *x.Conn, window x.Window, cursor x.Cursor, deviceId DeviceId) {
	w := x.NewWriter()
	writeXIChangeCursor(w, window, cursor, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XIChangeCursorOpcode,
	}
	conn.SendRequest(0, d, req)
}

func XIChangeCursorChecked(conn *x.Conn, window x.Window, cursor x.Cursor, deviceId DeviceId) x.VoidCookie {
	w := x.NewWriter()
	writeXIChangeCursor(w, window, cursor, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XIChangeCursorOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func XIChangeHierarchy(conn *x.Conn, changes []HierarchyChange) {
	w := x.NewWriter()
	writeXIChangeHierarchy(w, changes)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XIChangeHierarchyOpcode,
	}
	conn.SendRequest(0, d, req)
}

func XIChangeHierarchyChecked(conn *x.Conn, changes []HierarchyChange) x.VoidCookie {
	w := x.NewWriter()
	writeXIChangeHierarchy(w, changes)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XIChangeHierarchyOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func XISetClientPointer(conn *x.Conn, window x.Window, deviceId DeviceId) {
	w := x.NewWriter()
	writeXISetClientPointer(w, window, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XISetClientPointerOpcode,
	}
	conn.SendRequest(0, d, req)
}

func XISetClientPointerChecked(conn *x.Conn, window x.Window, deviceId DeviceId) x.VoidCookie {
	w := x.NewWriter()
	writeXISetClientPointer(w, window, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XISetClientPointerOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func XIGetClientPointer(conn *x.Conn, window x.Window) XIGetClientPointerCookie {
	w := x.NewWriter()
	writeXIGetClientPointer(w, window)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: XIGetClientPointerOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
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
	w := x.NewWriter()
	writeXISetFocus(w, focus, time, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XISetFocusOpcode,
	}
	conn.SendRequest(0, d, req)
}

func XISetFocusChecked(conn *x.Conn, focus x.Window, time x.Timestamp, deviceId DeviceId) x.VoidCookie {
	w := x.NewWriter()
	writeXISetFocus(w, focus, time, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XISetFocusOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func XIGetFocus(conn *x.Conn, deviceId DeviceId) XIGetFocusCookie {
	w := x.NewWriter()
	writeXIGetFocus(w, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: XIGetFocusOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
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
	w := x.NewWriter()
	writeXIGrabDevice(w, window, time, cursor, deviceId, mode, pairedDeviceMode, ownerEvents, masks)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: XIGrabDeviceOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
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
	w := x.NewWriter()
	writeXIUngrabDevice(w, time, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XIUngrabDeviceOpcode,
	}
	conn.SendRequest(0, d, req)
}

func XIUngrabDeviceChecked(conn *x.Conn, time x.Timestamp, deviceId DeviceId) x.VoidCookie {
	w := x.NewWriter()
	writeXIUngrabDevice(w, time, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XIUngrabDeviceOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func XIAllowEvents(conn *x.Conn, time x.Timestamp, deviceId DeviceId, eventMode uint8, touchId uint32, grabWindow x.Window) {
	w := x.NewWriter()
	writeXIAllowEvents(w, time, deviceId, eventMode, touchId, grabWindow)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XIAllowEventsOpcode,
	}
	conn.SendRequest(0, d, req)
}

func XIAllowEventsChecked(conn *x.Conn, time x.Timestamp, deviceId DeviceId, eventMode uint8, touchId uint32, grabWindow x.Window) x.VoidCookie {
	w := x.NewWriter()
	writeXIAllowEvents(w, time, deviceId, eventMode, touchId, grabWindow)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XIAllowEventsOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func XIPassiveGrabDevice(conn *x.Conn, grabWindow x.Window, cursor x.Cursor, detail uint32, deviceId DeviceId, grabType, grabMode, pairedDeviceMode uint8, ownerEvents bool, masks []uint32, modifiers []uint32) XIPassiveGrabDeviceCookie {
	w := x.NewWriter()
	writeXIPassiveGrabDevice(w, grabWindow, cursor, detail, deviceId, grabType, grabMode, pairedDeviceMode, ownerEvents, masks, modifiers)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: XIPassiveGrabDeviceOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
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
	w := x.NewWriter()
	writeXIPassiveUngrabDevice(w, grabWindow, detail, deviceId, grabType, modifiers)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XIPassiveUngrabDeviceOpcode,
	}
	conn.SendRequest(0, d, req)
}

func XIPassiveUngrabDeviceChecked(conn *x.Conn, grabWindow x.Window, detail uint32, deviceId DeviceId, grabType uint8, modifiers []uint32) x.VoidCookie {
	w := x.NewWriter()
	writeXIPassiveUngrabDevice(w, grabWindow, detail, deviceId, grabType, modifiers)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XIPassiveUngrabDeviceOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func XIListProperties(conn *x.Conn, deviceId DeviceId) XIListPropertiesCookie {
	w := x.NewWriter()
	writeXIListProperties(w, deviceId)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: XIListPropertiesOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
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
	w := x.NewWriter()
	writeXIChangeProperty(w, deviceId, mode, format, property, Type, data)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XIChangePropertyOpcode,
	}
	conn.SendRequest(0, d, req)
}

func XIChangePropertyChecked(conn *x.Conn, deviceId DeviceId, mode uint8, format uint8, property x.Atom, Type x.Atom, data []byte) x.VoidCookie {
	w := x.NewWriter()
	writeXIChangeProperty(w, deviceId, mode, format, property, Type, data)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XIChangePropertyOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func XIDeleteProperty(conn *x.Conn, deviceId DeviceId, property x.Atom) {
	w := x.NewWriter()
	writeXIDeleteProperty(w, deviceId, property)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XIDeletePropertyOpcode,
	}
	conn.SendRequest(0, d, req)
}

func XIDeletePropertyChecked(conn *x.Conn, deviceId DeviceId, property x.Atom) x.VoidCookie {
	w := x.NewWriter()
	writeXIDeleteProperty(w, deviceId, property)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  XIDeletePropertyOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func XIGetProperty(conn *x.Conn, deviceId DeviceId, delete bool, property x.Atom, Type x.Atom, offset, len uint32) XIGetPropertyCookie {
	w := x.NewWriter()
	writeXIGetProperty(w, deviceId, delete, property, Type, offset, len)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: XIGetPropertyOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
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
