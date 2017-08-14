package client

var bigReqExt = NewExtension("BIG-REQUESTS")

const EnableBigReqOpcode = 0

func EnableBigReqRequestWrite(w *Writer) {
	n := 0
	pad := 0

	// write wire field MajorOpcode
	// auto field MajorOpcode
	w.WritePad(1)
	n += 1

	// write wire field MinorOpcode
	// auto field MinorOpcode
	w.WritePad(1)
	n += 1

	// write wire field Length
	// auto field Length
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type EnableBigReqReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence             uint16
	Length               uint32
	MaximumRequestLength uint32
}

func EnableBigReqReplyRead(r *Reader, v *EnableBigReqReply) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ResponseType
	v.ResponseType = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(1)
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Sequence
	v.Sequence = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Length
	v.Length = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field MaximumRequestLength
	v.MaximumRequestLength = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}
	return
}

type EnableBigReqCookie uint64

func (cookie EnableBigReqCookie) Reply(conn *Conn) (*EnableBigReqReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, NewGenericError(replyBuf)
	}
	r := NewReaderFromData(replyBuf)
	var reply EnableBigReqReply
	EnableBigReqReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func EnableBigReq(conn *Conn) EnableBigReqCookie {
	w := NewWriter()
	EnableBigReqRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		Ext:    bigReqExt,
		IsVoid: false,
		Opcode: EnableBigReqOpcode,
	}
	// no check: False
	seq := conn.SendRequest(RequestChecked, data, req)
	return EnableBigReqCookie(seq)
}

func EnableBigReqUnchecked(conn *Conn) EnableBigReqCookie {
	w := NewWriter()
	EnableBigReqRequestWrite(w)
	data := w.Bytes()
	req := &ProtocolRequest{
		Ext:    bigReqExt,
		IsVoid: false,
		Opcode: EnableBigReqOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return EnableBigReqCookie(seq)
}
