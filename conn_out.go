package client

import (
//"sync"
)

func (c *Conn) write(data []byte) error {
	Logger.Printf("conn write %x\n", data)
	_, err := c.conn.Write(data)
	if err != nil {
		Logger.Println("write error:", err)
	}
	return err
}

func (c *Conn) writeLoop() {
	for data := range c.writeChan {
		c.write(data)
	}
}

// getInputFocusRequest writes the raw bytes to a buffer.
// It is duplicated from xproto/xproto.go.
func getInputFocusRequest() []byte {
	const size = 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 43 // request opcode
	b += 1

	b += 1                         // padding
	Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

func (c *Conn) sendSync() {
	data := getInputFocusRequest()
	c.sendRequest(false, 0, RequestDiscardReply, data)
}

func (c *Conn) SendSync() {
	c.iolock.Lock()
	c.sendSync()
	c.iolock.Unlock()
}

type Out struct {
	request uint64
	//mu sync.Mutex
	//requestWritten uint64 // ?
}

func NewOut() *Out {
	return new(Out)
}

func (c *Conn) sendRequest(isVoid bool, workaround uint, flags uint, data []byte) {
	c.out.request++
	if !isVoid {
		// not void, has reply
		c.in.requestExpected = c.out.request
	}

	if workaround != 0 || flags != 0 {
		c.in.expectReply(c.out.request, workaround, flags)
	}
	c.writeChan <- data
}

type ProtocolRequest struct {
	Ext    *Extension
	Opcode uint8
	IsVoid bool
}

// return sequencue id
func (c *Conn) SendRequest(flags uint, data []byte, req *ProtocolRequest) uint64 {
	// process data auto field
	// set the major opcode, and the minor opcode for extensions
	if req.Ext != nil {
		extension := c.GetExtensionData(req.Ext)
		if !(extension != nil && extension.Present != 0) {
			// TODO
			// return 0
			panic("ext not supported")
		}

		data[0] = extension.MajorOpcode
		data[1] = req.Opcode
	} else {
		data[0] = req.Opcode
	}

	var shortlen uint16 = uint16(len(data))
	if shortlen&3 != 0 {
		panic("assert shorlen & 3 == 0 falied")
	}
	shortlen >>= 2
	// set length field
	Put16(data[2:], shortlen)

	c.iolock.Lock()
	c.sendRequest(req.IsVoid, 0, flags, data)
	request := c.out.request
	c.iolock.Unlock()
	return request
}
