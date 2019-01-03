package bigrequests

import x "github.com/linuxdeepin/go-x11-client"

func Enable(conn *x.Conn) EnableCookie {
	body := encodeEnable()
	req := &x.ProtocolRequest{
		Ext: _ext,
		Header: x.RequestHeader{
			Data: EnableOpcode,
		},
		Body: body,
	}
	seq := conn.SendRequest(x.RequestChecked, req)
	return EnableCookie(seq)
}

func (cookie EnableCookie) Reply(conn *x.Conn) (*EnableReply, error) {
	replyBuf, err := conn.WaitForReply(x.SeqNum(cookie))
	if err != nil {
		return nil, err
	}
	r := x.NewReaderFromData(replyBuf)
	var reply EnableReply
	err = readEnableReply(r, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
