package xkb

import x "github.com/linuxdeepin/go-x11-client"

func UseExtension(conn *x.Conn, wantedMajor, wantedMinor uint16) UseExtensionCookie {
	body := encodeUseExtension(wantedMajor, wantedMinor)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: UseExtensionOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return UseExtensionCookie(seq)
}

func (cookie UseExtensionCookie) Reply(conn *x.Conn) (*UseExtensionReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply UseExtensionReply
	err = readUseExtensionReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SelectEvents(conn *x.Conn, deviceSpec DeviceSpec, options ...SelectOption) {
	body := encodeSelectEvents(deviceSpec, options...)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SelectEventsOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SelectEventsChecked(conn *x.Conn, deviceSpec DeviceSpec, options ...SelectOption) x.VoidCookie {
	body := encodeSelectEvents(deviceSpec, options...)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SelectEventsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func Bell(conn *x.Conn, deviceSpec DeviceSpec, bellClass BellClassSpec, bellID IdSpec, percent int8, forceSound, eventOnly bool, pitch, duration int16, name x.Atom, window x.Window) {
	body := encodeBell(deviceSpec, bellClass, bellID, percent, forceSound, eventOnly, pitch, duration, name, window)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: BellOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func BellChecked(conn *x.Conn, deviceSpec DeviceSpec, bellClass BellClassSpec, bellID IdSpec, percent int8, forceSound, eventOnly bool, pitch, duration int16, name x.Atom, window x.Window) x.VoidCookie {
	body := encodeBell(deviceSpec, bellClass, bellID, percent, forceSound, eventOnly, pitch, duration, name, window)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: BellOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func GetState(conn *x.Conn, deviceSpec DeviceSpec) GetStateCookie {
	body := encodeGetState(deviceSpec)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetStateOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetStateCookie(seq)
}

func (cookie GetStateCookie) Reply(conn *x.Conn) (*GetStateReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetStateReply
	err = readGetStateReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func LatchLockState(conn *x.Conn, deviceSpec DeviceSpec, affectModLocks, modLocks uint8, lockGroup bool, groupLock, affectModLatches, modLatches uint8, latchGroup bool, groupLatch uint16) {
	body := encodeLatchLockState(deviceSpec, affectModLocks, modLocks, lockGroup, groupLock, affectModLatches, modLatches, latchGroup, groupLatch)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: LatchLockStateOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func LatchLockStateChecked(conn *x.Conn, deviceSpec DeviceSpec, affectModLocks, modLocks uint8, lockGroup bool, groupLock, affectModLatches, modLatches uint8, latchGroup bool, groupLatch uint16) x.VoidCookie {
	body := encodeLatchLockState(deviceSpec, affectModLocks, modLocks, lockGroup, groupLock, affectModLatches, modLatches, latchGroup, groupLatch)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: LatchLockStateOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func GetControls(conn *x.Conn, deviceSpec DeviceSpec) GetControlsCookie {
	body := encodeGetControls(deviceSpec)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetControlsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetControlsCookie(seq)
}

func (cookie GetControlsCookie) Reply(conn *x.Conn) (*GetControlsReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetControlsReply
	err = readGetControlsReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetControls(conn *x.Conn, deviceSpec DeviceSpec, affectInternalRealMods, affectIgnoreLockRealMods uint8, affectInternalVirtualMods, affectIgnoreLockVirtualMods uint16, affectEnabledControls, changeControls uint32, c *Controls) {
	body := encodeSetControls(deviceSpec, affectInternalRealMods, affectIgnoreLockRealMods, affectInternalVirtualMods, affectIgnoreLockVirtualMods, affectEnabledControls, changeControls, c)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetControlsOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetControlsChecked(conn *x.Conn, deviceSpec DeviceSpec, affectInternalRealMods, affectIgnoreLockRealMods uint8, affectInternalVirtualMods, affectIgnoreLockVirtualMods uint16, affectEnabledControls, changeControls uint32, c *Controls) x.VoidCookie {
	body := encodeSetControls(deviceSpec, affectInternalRealMods, affectIgnoreLockRealMods, affectInternalVirtualMods, affectIgnoreLockVirtualMods, affectEnabledControls, changeControls, c)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetControlsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func GetMap(conn *x.Conn, deviceSpec DeviceSpec, full uint16, opt *GetMapPartOptions) GetMapCookie {
	body := encodeGetMap(deviceSpec, full, opt)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetMapOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetMapCookie(seq)
}

func (cookie GetMapCookie) Reply(conn *x.Conn) (*GetMapReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetMapReply
	err = readGetMapReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
