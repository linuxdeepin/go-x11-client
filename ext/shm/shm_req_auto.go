package shm

import x "github.com/linuxdeepin/go-x11-client"

func QueryVersion(conn *x.Conn) QueryVersionCookie {
	body := encodeQueryVersion()
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

func Attach(conn *x.Conn, shmSeg Seg, shmId uint32, readOnly bool) {
	body := encodeAttach(shmSeg, shmId, readOnly)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: AttachOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func AttachChecked(conn *x.Conn, shmSeg Seg, shmId uint32, readOnly bool) x.VoidCookie {
	body := encodeAttach(shmSeg, shmId, readOnly)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: AttachOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func Detach(conn *x.Conn, shmSeg Seg) {
	body := encodeDetach(shmSeg)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: DetachOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func DetachChecked(conn *x.Conn, shmSeg Seg) x.VoidCookie {
	body := encodeDetach(shmSeg)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: DetachOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func PutImage(conn *x.Conn, drawable x.Drawable, gc x.GContext, totalWidth, totalHeight, srcX, srcY, srcWidth, srcHeight uint16, dstX, dstY int16, depth, format uint8, sendEvent bool, shmSeg Seg, offset uint32) {
	body := encodePutImage(drawable, gc, totalWidth, totalHeight, srcX, srcY, srcWidth, srcHeight, dstX, dstY, depth, format, sendEvent, shmSeg, offset)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: PutImageOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func PutImageChecked(conn *x.Conn, drawable x.Drawable, gc x.GContext, totalWidth, totalHeight, srcX, srcY, srcWidth, srcHeight uint16, dstX, dstY int16, depth, format uint8, sendEvent bool, shmSeg Seg, offset uint32) x.VoidCookie {
	body := encodePutImage(drawable, gc, totalWidth, totalHeight, srcX, srcY, srcWidth, srcHeight, dstX, dstY, depth, format, sendEvent, shmSeg, offset)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: PutImageOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}

func GetImage(conn *x.Conn, drawable x.Drawable, X, y int16, width, height uint16, planeMask uint32, format uint8, shmSeg Seg, offset uint32) GetImageCookie {
	body := encodeGetImage(drawable, X, y, width, height, planeMask, format, shmSeg, offset)
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: GetImageOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return GetImageCookie(seq)
}

func (cookie GetImageCookie) Reply(conn *x.Conn) (*GetImageReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
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
	body := encodeCreatePixmap(pid, drawable, width, height, depth, shmSeg, offset)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreatePixmapOpcode,
		},
		Body: body,
	}
	conn.SendRequest(0, req)
}

func CreatePixmapChecked(conn *x.Conn, pid x.Pixmap, drawable x.Drawable, width, height uint16, depth uint8, shmSeg Seg, offset uint32) x.VoidCookie {
	body := encodeCreatePixmap(pid, drawable, width, height, depth, shmSeg, offset)
	req := &x.ProtocolRequest{
		Ext:     _ext,
		NoReply: true,
		Header: x.RequestHeader{
			Data: CreatePixmapOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return x.VoidCookie(seq)
}
