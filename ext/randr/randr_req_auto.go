package randr

import x "github.com/linuxdeepin/go-x11-client"

func QueryVersion(conn *x.Conn, clientMajorVersion, clientMinorVersion uint32) QueryVersionCookie {
	w := x.NewWriter()
	writeQueryVersion(w, clientMajorVersion, clientMinorVersion)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: QueryVersionOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
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
	w := x.NewWriter()
	writeGetCrtcInfo(w, crtc, configTimestamp)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: GetCrtcInfoOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
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
	w := x.NewWriter()
	writeGetOutputInfo(w, output, configTimestamp)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: GetOutputInfoOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
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
	w := x.NewWriter()
	writeGetScreenResources(w, window)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: GetScreenResourcesOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
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
	w := x.NewWriter()
	writeGetOutputPrimary(w, window)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: GetOutputPrimaryOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
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
	w := x.NewWriter()
	writeGetOutputProperty(w, output, property, Type, longOffset, longLength, delete, pending)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: GetOutputPropertyOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
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
	w := x.NewWriter()
	writeSetCrtcConfig(w, crtc, timestamp, configTimestamp, X, y, mode, rotation, outputs)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: SetCrtcConfigOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
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
	w := x.NewWriter()
	writeSelectInput(w, window, enable)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SelectInputOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SelectInputChecked(conn *x.Conn, window x.Window, enable uint16) x.VoidCookie {
	w := x.NewWriter()
	writeSelectInput(w, window, enable)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SelectInputOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func GetCrtcGammaSize(conn *x.Conn, crtc Crtc) GetCrtcGammaSizeCookie {
	w := x.NewWriter()
	writeGetCrtcGammaSize(w, crtc)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: GetCrtcGammaSizeOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
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
	w := x.NewWriter()
	writeSetCrtcGamma(w, crtc, red, green, blue)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SetCrtcGammaOpcode,
	}
	conn.SendRequest(0, d, req)
}

func SetCrtcGammaChecked(conn *x.Conn, crtc Crtc, red, green, blue []uint16) x.VoidCookie {
	w := x.NewWriter()
	writeSetCrtcGamma(w, crtc, red, green, blue)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  SetCrtcGammaOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}
