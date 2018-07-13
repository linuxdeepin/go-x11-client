package shm

import x "github.com/linuxdeepin/go-x11-client"

func QueryVersion(conn *x.Conn) QueryVersionCookie {
	w := x.NewWriter()
	writeQueryVersion(w)
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

func Attach(conn *x.Conn, shmSeg Seg, shmId uint32, readOnly bool) {
	w := x.NewWriter()
	writeAttach(w, shmSeg, shmId, readOnly)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  AttachOpcode,
	}
	conn.SendRequest(0, d, req)
}

func AttachChecked(conn *x.Conn, shmSeg Seg, shmId uint32, readOnly bool) x.VoidCookie {
	w := x.NewWriter()
	writeAttach(w, shmSeg, shmId, readOnly)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  AttachOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func Detach(conn *x.Conn, shmSeg Seg) {
	w := x.NewWriter()
	writeDetach(w, shmSeg)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  DetachOpcode,
	}
	conn.SendRequest(0, d, req)
}

func DetachChecked(conn *x.Conn, shmSeg Seg) x.VoidCookie {
	w := x.NewWriter()
	writeDetach(w, shmSeg)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  DetachOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func PutImage(conn *x.Conn, drawable x.Drawable, gc x.GContext, totalWidth, totalHeight, srcX, srcY, srcWidth, srcHeight uint16, dstX, dstY int16, depth, format uint8, sendEvent bool, shmSeg Seg, offset uint32) {
	w := x.NewWriter()
	writePutImage(w, drawable, gc, totalWidth, totalHeight, srcX, srcY, srcWidth, srcHeight, dstX, dstY, depth, format, sendEvent, shmSeg, offset)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  PutImageOpcode,
	}
	conn.SendRequest(0, d, req)
}

func PutImageChecked(conn *x.Conn, drawable x.Drawable, gc x.GContext, totalWidth, totalHeight, srcX, srcY, srcWidth, srcHeight uint16, dstX, dstY int16, depth, format uint8, sendEvent bool, shmSeg Seg, offset uint32) x.VoidCookie {
	w := x.NewWriter()
	writePutImage(w, drawable, gc, totalWidth, totalHeight, srcX, srcY, srcWidth, srcHeight, dstX, dstY, depth, format, sendEvent, shmSeg, offset)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  PutImageOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}

func GetImage(conn *x.Conn, drawable x.Drawable, X, y int16, width, height uint16, planeMask uint32, format uint8, shmSeg Seg, offset uint32) GetImageCookie {
	w := x.NewWriter()
	writeGetImage(w, drawable, X, y, width, height, planeMask, format, shmSeg, offset)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		Opcode: GetImageOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return GetImageCookie(seq)
}

func (cookie GetImageCookie) Reply(conn *x.Conn) (*GetImageReply, error) {
	replyBuf, err := conn.WaitForReply(uint64(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetImageReply
	err = readGetImageReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}

func CreatePixmap(conn *x.Conn, pid x.Pixmap, drawable x.Drawable, width, height uint16, depth uint8, shmSeg Seg, offset uint32) {
	w := x.NewWriter()
	writeCreatePixmap(w, pid, drawable, width, height, depth, shmSeg, offset)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreatePixmapOpcode,
	}
	conn.SendRequest(0, d, req)
}

func CreatePixmapChecked(conn *x.Conn, pid x.Pixmap, drawable x.Drawable, width, height uint16, depth uint8, shmSeg Seg, offset uint32) x.VoidCookie {
	w := x.NewWriter()
	writeCreatePixmap(w, pid, drawable, width, height, depth, shmSeg, offset)
	d := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Opcode:  CreatePixmapOpcode,
	}
	seq := conn.SendRequest(x.RequestChecked, d, req)
	return x.VoidCookie(seq)
}
