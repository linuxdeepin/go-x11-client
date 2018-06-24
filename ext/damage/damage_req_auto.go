package damage

import x "github.com/linuxdeepin/go-x11-client"

func QueryVersion(conn *x.Conn, majorVersion, minorVersion uint32) QueryVersionCookie {
	w := x.NewWriter()
	writeQueryVersion(w, majorVersion, minorVersion)
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

func Create(conn *x.Conn, damage Damage, drawable x.Drawable, level uint8) {
	w := x.NewWriter()
	writeCreate(w, damage, drawable, level)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CreateChecked(conn *x.Conn, damage Damage, drawable x.Drawable, level uint8) x.VoidCookie {
	w := x.NewWriter()
	writeCreate(w, damage, drawable, level)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreateOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func Destroy(conn *x.Conn, damage Damage) {
	w := x.NewWriter()
	writeDestroy(w, damage)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  DestroyOpcode,
	}
	conn.SendRequest(0, d, req)
}

func DestroyChecked(conn *x.Conn, damage Damage) x.VoidCookie {
	w := x.NewWriter()
	writeDestroy(w, damage)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  DestroyOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}
