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

func SetMap(conn *x.Conn, deviceSpec DeviceSpec, flags uint16, m *SetMapValue) {
	body := encodeSetMap(deviceSpec, flags, m)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetMapOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetMapChecked(conn *x.Conn, deviceSpec DeviceSpec, flags uint16, m *SetMapValue) x.VoidCookie {
	body := encodeSetMap(deviceSpec, flags, m)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetMapOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func GetCompatMap(conn *x.Conn, deviceSpec DeviceSpec, groups uint8, getAllSI bool, firstSI, NSI uint16) GetCompatMapCookie {
	body := encodeGetCompatMap(deviceSpec, groups, getAllSI, firstSI, NSI)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetCompatMapOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetCompatMapCookie(seq)
}

func (cookie GetCompatMapCookie) Reply(conn *x.Conn) (*GetCompatMapReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetCompatMapReply
	err = readGetCompatMapReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetCompatMap(conn *x.Conn, deviceSpec DeviceSpec, v *SetCompatMapValue) {
	body := encodeSetCompatMap(deviceSpec, v)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetCompatMapOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetCompatMapChecked(conn *x.Conn, deviceSpec DeviceSpec, v *SetCompatMapValue) x.VoidCookie {
	body := encodeSetCompatMap(deviceSpec, v)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetCompatMapOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func GetIndicatorState(conn *x.Conn, deviceSpec DeviceSpec) GetIndicatorStateCookie {
	body := encodeGetIndicatorState(deviceSpec)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetIndicatorStateOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetIndicatorStateCookie(seq)
}

func (cookie GetIndicatorStateCookie) Reply(conn *x.Conn) (*GetIndicatorStateReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetIndicatorStateReply
	err = readGetIndicatorStateReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetIndicatorMap(conn *x.Conn, deviceSpec DeviceSpec, which uint32) GetIndicatorMapCookie {
	body := encodeGetIndicatorMap(deviceSpec, which)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetIndicatorMapOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetIndicatorMapCookie(seq)
}

func (cookie GetIndicatorMapCookie) Reply(conn *x.Conn) (*GetIndicatorMapReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetIndicatorMapReply
	err = readGetIndicatorMapReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetIndicatorMap(conn *x.Conn, deviceSpec DeviceSpec, which uint32, maps []IndicatorMap) {
	body := encodeSetIndicatorMap(deviceSpec, which, maps)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetIndicatorMapOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetIndicatorMapChecked(conn *x.Conn, deviceSpec DeviceSpec, which uint32, maps []IndicatorMap) x.VoidCookie {
	body := encodeSetIndicatorMap(deviceSpec, which, maps)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetIndicatorMapOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func GetNamedIndicator(conn *x.Conn, deviceSpec DeviceSpec, ledClass LedClassSpec, ledID IdSpec, indicator x.Atom) GetNamedIndicatorCookie {
	body := encodeGetNamedIndicator(deviceSpec, ledClass, ledID, indicator)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetNamedIndicatorOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetNamedIndicatorCookie(seq)
}

func (cookie GetNamedIndicatorCookie) Reply(conn *x.Conn) (*GetNamedIndicatorReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetNamedIndicatorReply
	err = readGetNamedIndicatorReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetNamedIndicator(conn *x.Conn, deviceSpec DeviceSpec, ledClass LedClassSpec, ledID IdSpec, indicator x.Atom, setState, on, setMap, createMap bool, map0 *IndicatorMap) {
	body := encodeSetNamedIndicator(deviceSpec, ledClass, ledID, indicator, setState, on, setMap, createMap, map0)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetNamedIndicatorOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetNamedIndicatorChecked(conn *x.Conn, deviceSpec DeviceSpec, ledClass LedClassSpec, ledID IdSpec, indicator x.Atom, setState, on, setMap, createMap bool, map0 *IndicatorMap) x.VoidCookie {
	body := encodeSetNamedIndicator(deviceSpec, ledClass, ledID, indicator, setState, on, setMap, createMap, map0)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetNamedIndicatorOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func GetNames(conn *x.Conn, deviceSpec DeviceSpec, which uint32) GetNamesCookie {
	body := encodeGetNames(deviceSpec, which)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetNamesOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetNamesCookie(seq)
}

func (cookie GetNamesCookie) Reply(conn *x.Conn) (*GetNamesReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetNamesReply
	err = readGetNamesReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetNames(conn *x.Conn, r *SetNamesRequest) {
	body := encodeSetNames(r)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetNamesOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetNamesChecked(conn *x.Conn, r *SetNamesRequest) x.VoidCookie {
	body := encodeSetNames(r)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetNamesOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func PerClientFlags(conn *x.Conn, deviceSpec DeviceSpec, change, value, ctrlsToChange, autoCtrls, autoCtrlsValues uint32) PerClientFlagsCookie {
	body := encodePerClientFlags(deviceSpec, change, value, ctrlsToChange, autoCtrls, autoCtrlsValues)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: PerClientFlagsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return PerClientFlagsCookie(seq)
}

func (cookie PerClientFlagsCookie) Reply(conn *x.Conn) (*PerClientFlagsReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply PerClientFlagsReply
	err = readPerClientFlagsReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func ListComponents(conn *x.Conn, deviceSpec DeviceSpec, maxNames uint16, cn *ComponentNames) ListComponentsCookie {
	body := encodeListComponents(deviceSpec, maxNames, cn)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: ListComponentsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return ListComponentsCookie(seq)
}

func (cookie ListComponentsCookie) Reply(conn *x.Conn) (*ListComponentsReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply ListComponentsReply
	err = readListComponentsReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetDeviceInfo(conn *x.Conn, deviceSpec DeviceSpec, wanted uint16, allButtons bool, FirstButton, nButtons uint8, ledClass LedClassSpec, ledID IdSpec) GetDeviceInfoCookie {
	body := encodeGetDeviceInfo(deviceSpec, wanted, allButtons, FirstButton, nButtons, ledClass, ledID)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetDeviceInfoOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetDeviceInfoCookie(seq)
}

func (cookie GetDeviceInfoCookie) Reply(conn *x.Conn) (*GetDeviceInfoReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetDeviceInfoReply
	err = readGetDeviceInfoReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetDeviceInfo(conn *x.Conn, deviceSpec DeviceSpec, firstBtn uint8, change uint16, btnActions []Action, leds []DeviceLedInfo) {
	body := encodeSetDeviceInfo(deviceSpec, firstBtn, change, btnActions, leds)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetDeviceInfoOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetDeviceInfoChecked(conn *x.Conn, deviceSpec DeviceSpec, firstBtn uint8, change uint16, btnActions []Action, leds []DeviceLedInfo) x.VoidCookie {
	body := encodeSetDeviceInfo(deviceSpec, firstBtn, change, btnActions, leds)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetDeviceInfoOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func SetDebuggingFlags(conn *x.Conn, affectFlags, flags, affectCtrls, ctrls uint32, message string) SetDebuggingFlagsCookie {
	body := encodeSetDebuggingFlags(affectFlags, flags, affectCtrls, ctrls, message)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: SetDebuggingFlagsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return SetDebuggingFlagsCookie(seq)
}

func (cookie SetDebuggingFlagsCookie) Reply(conn *x.Conn) (*SetDebuggingFlagsReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply SetDebuggingFlagsReply
	err = readSetDebuggingFlagsReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
