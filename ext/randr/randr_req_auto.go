package randr

import x "github.com/linuxdeepin/go-x11-client"
import "github.com/linuxdeepin/go-x11-client/ext/render"

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

func SetScreenConfig(conn *x.Conn, window x.Window, timestamp x.Timestamp, configTimestamp x.Timestamp, sizeID uint16, rotation uint16, rate uint16) SetScreenConfigCookie {
	body := encodeSetScreenConfig(window, timestamp, configTimestamp, sizeID, rotation, rate)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: SetScreenConfigOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return SetScreenConfigCookie(seq)
}

func (cookie SetScreenConfigCookie) Reply(conn *x.Conn) (*SetScreenConfigReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply SetScreenConfigReply
	err = readSetScreenConfigReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetScreenInfo(conn *x.Conn, window x.Window) GetScreenInfoCookie {
	body := encodeGetScreenInfo(window)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetScreenInfoOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetScreenInfoCookie(seq)
}

func (cookie GetScreenInfoCookie) Reply(conn *x.Conn) (*GetScreenInfoReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetScreenInfoReply
	err = readGetScreenInfoReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetScreenSizeRange(conn *x.Conn, window x.Window) GetScreenSizeRangeCookie {
	body := encodeGetScreenSizeRange(window)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetScreenSizeRangeOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetScreenSizeRangeCookie(seq)
}

func (cookie GetScreenSizeRangeCookie) Reply(conn *x.Conn) (*GetScreenSizeRangeReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetScreenSizeRangeReply
	err = readGetScreenSizeRangeReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetScreenSize(conn *x.Conn, window x.Window, width, height uint16, mmWidth, mmHeight uint32) {
	body := encodeSetScreenSize(window, width, height, mmWidth, mmHeight)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetScreenSizeOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetScreenSizeChecked(conn *x.Conn, window x.Window, width, height uint16, mmWidth, mmHeight uint32) x.VoidCookie {
	body := encodeSetScreenSize(window, width, height, mmWidth, mmHeight)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetScreenSizeOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
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
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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

func GetScreenResourcesCurrent(conn *x.Conn, window x.Window) GetScreenResourcesCurrentCookie {
	body := encodeGetScreenResourcesCurrent(window)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetScreenResourcesCurrentOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetScreenResourcesCurrentCookie(seq)
}

func (cookie GetScreenResourcesCurrentCookie) Reply(conn *x.Conn) (*GetScreenResourcesCurrentReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetScreenResourcesCurrentReply
	err = readGetScreenResourcesCurrentReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetCrtcTransform(conn *x.Conn, crtc Crtc, transform *render.Transform, filter string, filterParams []render.Fixed) {
	body := encodeSetCrtcTransform(crtc, transform, filter, filterParams)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetCrtcTransformOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetCrtcTransformChecked(conn *x.Conn, crtc Crtc, transform *render.Transform, filter string, filterParams []render.Fixed) x.VoidCookie {
	body := encodeSetCrtcTransform(crtc, transform, filter, filterParams)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetCrtcTransformOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func GetCrtcTransform(conn *x.Conn, crtc Crtc) GetCrtcTransformCookie {
	body := encodeGetCrtcTransform(crtc)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetCrtcTransformOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetCrtcTransformCookie(seq)
}

func (cookie GetCrtcTransformCookie) Reply(conn *x.Conn) (*GetCrtcTransformReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetCrtcTransformReply
	err = readGetCrtcTransformReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetPanning(conn *x.Conn, crtc Crtc) GetPanningCookie {
	body := encodeGetPanning(crtc)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetPanningOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetPanningCookie(seq)
}

func (cookie GetPanningCookie) Reply(conn *x.Conn) (*GetPanningReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetPanningReply
	err = readGetPanningReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetPanning(conn *x.Conn, crtc Crtc, timestamp x.Timestamp, panning *Panning) SetPanningCookie {
	body := encodeSetPanning(crtc, timestamp, panning)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: SetPanningOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return SetPanningCookie(seq)
}

func (cookie SetPanningCookie) Reply(conn *x.Conn) (*SetPanningReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply SetPanningReply
	err = readSetPanningReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetOutputPrimary(conn *x.Conn, window x.Window, output Output) {
	body := encodeSetOutputPrimary(window, output)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetOutputPrimaryOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetOutputPrimaryChecked(conn *x.Conn, window x.Window, output Output) x.VoidCookie {
	body := encodeSetOutputPrimary(window, output)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetOutputPrimaryOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
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
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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

func ListOutputProperties(conn *x.Conn, output Output) ListOutputPropertiesCookie {
	body := encodeListOutputProperties(output)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: ListOutputPropertiesOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return ListOutputPropertiesCookie(seq)
}

func (cookie ListOutputPropertiesCookie) Reply(conn *x.Conn) (*ListOutputPropertiesReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply ListOutputPropertiesReply
	err = readListOutputPropertiesReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryOutputProperty(conn *x.Conn, output Output, property x.Atom) QueryOutputPropertyCookie {
	body := encodeQueryOutputProperty(output, property)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: QueryOutputPropertyOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return QueryOutputPropertyCookie(seq)
}

func (cookie QueryOutputPropertyCookie) Reply(conn *x.Conn) (*QueryOutputPropertyReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply QueryOutputPropertyReply
	err = readQueryOutputPropertyReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func ConfigureOutputProperty(conn *x.Conn, output Output, property x.Atom, pending, range0 bool, values []int32) {
	body := encodeConfigureOutputProperty(output, property, pending, range0, values)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ConfigureOutputPropertyOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func ConfigureOutputPropertyChecked(conn *x.Conn, output Output, property x.Atom, pending, range0 bool, values []int32) x.VoidCookie {
	body := encodeConfigureOutputProperty(output, property, pending, range0, values)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ConfigureOutputPropertyOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func ChangeOutputProperty(conn *x.Conn, output Output, property x.Atom, type0 x.Atom, format, mode uint8, data []byte) {
	body := encodeChangeOutputProperty(output, property, type0, format, mode, data)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ChangeOutputPropertyOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func ChangeOutputPropertyChecked(conn *x.Conn, output Output, property x.Atom, type0 x.Atom, format, mode uint8, data []byte) x.VoidCookie {
	body := encodeChangeOutputProperty(output, property, type0, format, mode, data)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: ChangeOutputPropertyOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func DeleteOutputProperty(conn *x.Conn, output Output, property x.Atom) {
	body := encodeDeleteOutputProperty(output, property)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: DeleteOutputPropertyOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func DeleteOutputPropertyChecked(conn *x.Conn, output Output, property x.Atom) x.VoidCookie {
	body := encodeDeleteOutputProperty(output, property)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: DeleteOutputPropertyOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
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
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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

func CreateMode(conn *x.Conn, window x.Window, modeInfo *ModeInfo) CreateModeCookie {
	body := encodeCreateMode(window, modeInfo)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: CreateModeOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return CreateModeCookie(seq)
}

func (cookie CreateModeCookie) Reply(conn *x.Conn) (*CreateModeReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply CreateModeReply
	err = readCreateModeReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func DestroyMode(conn *x.Conn, mode Mode) {
	body := encodeDestroyMode(mode)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: DestroyModeOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func DestroyModeChecked(conn *x.Conn, mode Mode) x.VoidCookie {
	body := encodeDestroyMode(mode)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: DestroyModeOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func AddOutputMode(conn *x.Conn, output Output, mode Mode) {
	body := encodeAddOutputMode(output, mode)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: AddOutputModeOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func AddOutputModeChecked(conn *x.Conn, output Output, mode Mode) x.VoidCookie {
	body := encodeAddOutputMode(output, mode)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: AddOutputModeOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func DeleteOutputMode(conn *x.Conn, output Output, mode Mode) {
	body := encodeDeleteOutputMode(output, mode)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: DeleteOutputModeOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func DeleteOutputModeChecked(conn *x.Conn, output Output, mode Mode) x.VoidCookie {
	body := encodeDeleteOutputMode(output, mode)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: DeleteOutputModeOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
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
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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

func GetCrtcGamma(conn *x.Conn, crtc Crtc) GetCrtcGammaCookie {
	body := encodeGetCrtcGamma(crtc)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetCrtcGammaOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetCrtcGammaCookie(seq)
}

func (cookie GetCrtcGammaCookie) Reply(conn *x.Conn) (*GetCrtcGammaReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetCrtcGammaReply
	err = readGetCrtcGammaReply(r, &reply)
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

func GetProviders(conn *x.Conn, window x.Window) GetProvidersCookie {
	body := encodeGetProviders(window)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetProvidersOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetProvidersCookie(seq)
}

func (cookie GetProvidersCookie) Reply(conn *x.Conn) (*GetProvidersReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetProvidersReply
	err = readGetProvidersReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetProviderInfo(conn *x.Conn, provider Provider, configTimestamp x.Timestamp) GetProviderInfoCookie {
	body := encodeGetProviderInfo(provider, configTimestamp)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetProviderInfoOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetProviderInfoCookie(seq)
}

func (cookie GetProviderInfoCookie) Reply(conn *x.Conn) (*GetProviderInfoReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetProviderInfoReply
	err = readGetProviderInfoReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetProviderOffloadSink(conn *x.Conn, provider, sinkProvider Provider, configTimestamp x.Timestamp) {
	body := encodeSetProviderOffloadSink(provider, sinkProvider, configTimestamp)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetProviderOffloadSinkOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetProviderOffloadSinkChecked(conn *x.Conn, provider, sinkProvider Provider, configTimestamp x.Timestamp) x.VoidCookie {
	body := encodeSetProviderOffloadSink(provider, sinkProvider, configTimestamp)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetProviderOffloadSinkOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func SetProviderOutputSource(conn *x.Conn, provider, sourceProvider Provider, configTimestamp x.Timestamp) {
	body := encodeSetProviderOutputSource(provider, sourceProvider, configTimestamp)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetProviderOutputSourceOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetProviderOutputSourceChecked(conn *x.Conn, provider, sourceProvider Provider, configTimestamp x.Timestamp) x.VoidCookie {
	body := encodeSetProviderOutputSource(provider, sourceProvider, configTimestamp)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetProviderOutputSourceOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func ListProviderProperties(conn *x.Conn, provider Provider) ListProviderPropertiesCookie {
	body := encodeListProviderProperties(provider)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: ListProviderPropertiesOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return ListProviderPropertiesCookie(seq)
}

func (cookie ListProviderPropertiesCookie) Reply(conn *x.Conn) (*ListProviderPropertiesReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply ListProviderPropertiesReply
	err = readListProviderPropertiesReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetMonitors(conn *x.Conn, window x.Window, getActive bool) GetMonitorsCookie {
	body := encodeGetMonitors(window, getActive)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetMonitorsOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetMonitorsCookie(seq)
}

func (cookie GetMonitorsCookie) Reply(conn *x.Conn) (*GetMonitorsReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetMonitorsReply
	err = readGetMonitorsReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func SetMonitor(conn *x.Conn, window x.Window, monitorInfo *MonitorInfo) {
	body := encodeSetMonitor(window, monitorInfo)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetMonitorOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func SetMonitorChecked(conn *x.Conn, window x.Window, monitorInfo *MonitorInfo) x.VoidCookie {
	body := encodeSetMonitor(window, monitorInfo)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: SetMonitorOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func DeleteMonitor(conn *x.Conn, window x.Window, name x.Atom) {
	body := encodeDeleteMonitor(window, name)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: DeleteMonitorOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func DeleteMonitorChecked(conn *x.Conn, window x.Window, name x.Atom) x.VoidCookie {
	body := encodeDeleteMonitor(window, name)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: DeleteMonitorOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}
