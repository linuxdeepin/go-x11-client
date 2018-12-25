package randr

import x "github.com/linuxdeepin/go-x11-client"

func QueryVersion(conn *x.Conn, clientMajorVersion, clientMinorVersion uint32) QueryVersionCookie {
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
	replyBuf, err := conn.WaitForReply(uint64(cookie))
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

func GetCrtcInfo(conn *x.Conn, crtc Crtc, configTimestamp x.Timestamp) GetCrtcInfoCookie {
	body := encodeGetCrtcInfo(crtc, configTimestamp)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetCrtcInfoOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetCrtcInfoCookie(seq)
}

func (cookie GetCrtcInfoCookie) Reply(conn *x.Conn) (*GetCrtcInfoReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetCrtcInfoReply
	err = readGetCrtcInfoReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetOutputInfo(conn *x.Conn, output Output, configTimestamp x.Timestamp) GetOutputInfoCookie {
	body := encodeGetOutputInfo(output, configTimestamp)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetOutputInfoOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetOutputInfoCookie(seq)
}

func (cookie GetOutputInfoCookie) Reply(conn *x.Conn) (*GetOutputInfoReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetOutputInfoReply
	err = readGetOutputInfoReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetScreenResources(conn *x.Conn, window x.Window) GetScreenResourcesCookie {
	body := encodeGetScreenResources(window)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetScreenResourcesOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetScreenResourcesCookie(seq)
}

func (cookie GetScreenResourcesCookie) Reply(conn *x.Conn) (*GetScreenResourcesReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetScreenResourcesReply
	err = readGetScreenResourcesReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetOutputPrimary(conn *x.Conn, window x.Window) GetOutputPrimaryCookie {
	body := encodeGetOutputPrimary(window)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetOutputPrimaryOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetOutputPrimaryCookie(seq)
}

func (cookie GetOutputPrimaryCookie) Reply(conn *x.Conn) (*GetOutputPrimaryReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetOutputPrimaryReply
	err = readGetOutputPrimaryReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetOutputProperty(conn *x.Conn, output Output, property, Type x.Atom, longOffset, longLength uint32, delete, pending bool) GetOutputPropertyCookie {
	body := encodeGetOutputProperty(output, property, Type, longOffset, longLength, delete, pending)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetOutputPropertyOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetOutputPropertyCookie(seq)
}

func (cookie GetOutputPropertyCookie) Reply(conn *x.Conn) (*GetOutputPropertyReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetOutputPropertyReply
	err = readGetOutputPropertyReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetCrtcConfig(conn *x.Conn, crtc Crtc, timestamp, configTimestamp x.Timestamp, X, y int16, mode Mode, rotation uint16, outputs []Output) SetCrtcConfigCookie {
	body := encodeSetCrtcConfig(crtc, timestamp, configTimestamp, X, y, mode, rotation, outputs)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: SetCrtcConfigOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return SetCrtcConfigCookie(seq)
}

func (cookie SetCrtcConfigCookie) Reply(conn *x.Conn) (*SetCrtcConfigReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply SetCrtcConfigReply
	err = readSetCrtcConfigReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SelectInput(conn *x.Conn, window x.Window, enable uint16) {
	body := encodeSelectInput(window, enable)
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

func SelectInputChecked(conn *x.Conn, window x.Window, enable uint16) x.VoidCookie {
	body := encodeSelectInput(window, enable)
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

func GetCrtcGammaSize(conn *x.Conn, crtc Crtc) GetCrtcGammaSizeCookie {
	body := encodeGetCrtcGammaSize(crtc)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetCrtcGammaSizeOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetCrtcGammaSizeCookie(seq)
}

func (cookie GetCrtcGammaSizeCookie) Reply(conn *x.Conn) (*GetCrtcGammaSizeReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetCrtcGammaSizeReply
	err = readGetCrtcGammaSizeReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetCrtcGamma(conn *x.Conn, crtc Crtc, red, green, blue []uint16) {
	body := encodeSetCrtcGamma(crtc, red, green, blue)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetCrtcGammaOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetCrtcGammaChecked(conn *x.Conn, crtc Crtc, red, green, blue []uint16) x.VoidCookie {
	body := encodeSetCrtcGamma(crtc, red, green, blue)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetCrtcGammaOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}
