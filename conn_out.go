package x

import (
	"io"
	"net"
)

func (c *Conn) Flush() {
	// TODO: check error
	c.ioMu.Lock()
	c.out.flushTo(c.out.request)
	c.ioMu.Unlock()
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
	c.ioMu.Lock()
	c.sendSync()
	c.ioMu.Unlock()
}

type out struct {
	request        uint64
	requestWritten uint64

	err error
	buf []byte
	n   int
	wr  io.Writer
}

func newOut(conn net.Conn) *out {
	out := &out{
		buf: make([]byte, 4096),
		wr:  conn,
	}
	return out
}

func (c *Conn) sendRequest(noReply bool, workaround uint, flags uint, data []byte) {
	c.out.request++
	if !noReply {
		// has reply
		c.in.requestExpected = c.out.request
	}

	if workaround != 0 || flags != 0 {
		c.in.expectReply(c.out.request, workaround, flags)
	}
	logPrintln("sendRequest seq:", c.out.request)
	_, err := c.out.write(c.out.request, data)
	if err != nil {
		logPrintln("write error:", err)
	}
}

type ProtocolRequest struct {
	Ext     *Extension
	Opcode  uint8
	NoReply bool
}

// return sequence id
func (c *Conn) SendRequest(flags uint, data []byte, req *ProtocolRequest) uint64 {
	// process data auto field
	// set the major opcode, and the minor opcode for extensions
	if req.Ext != nil {
		extension := c.GetExtensionData(req.Ext)
		if !(extension != nil && extension.Present) {
			// TODO
			// return 0
			panic("ext not supported")
		}

		data[0] = extension.MajorOpcode
		data[1] = req.Opcode
	} else {
		data[0] = req.Opcode
	}

	var requestLen = uint16(len(data))
	if requestLen&3 != 0 {
		panic("assert requestLen & 3 == 0 failed")
	}
	requestLen >>= 2
	// set length field
	Put16(data[2:], requestLen)

	c.ioMu.Lock()
	c.sendRequest(req.NoReply, 0, flags, data)
	request := c.out.request
	c.ioMu.Unlock()
	return request
}

func (o *out) available() int {
	return len(o.buf) - o.n
}

func (o *out) buffered() int {
	return o.n
}

func (o *out) write(request uint64, p []byte) (nn int, err error) {
	for len(p) > o.available() && o.err == nil {
		var n int
		if o.buffered() == 0 {
			// Large write, empty buffer.
			// Write directly from p to avoid copy.
			n, o.err = o.wr.Write(p)
		} else {
			n = copy(o.buf[o.n:], p)
			o.n += n
			o.flush()
		}
		nn += n
		p = p[n:]
	}
	if o.err != nil {
		return nn, o.err
	}
	n := copy(o.buf[o.n:], p)
	if n == 0 {
		//	完全写入到conn
		o.requestWritten = o.request
	}

	o.n += n
	nn += n

	return nn, nil
}

func (o *out) flushTo(request uint64) {
	if !(request <= o.request) {
		panic("assert request < o.request failed")
	}

	if o.requestWritten >= request {
		return
	}

	o.flush()
	o.requestWritten = o.request
}

// flush writes any buffered data to the underlying io.Writer.
func (o *out) flush() error {
	if o.err != nil {
		return o.err
	}
	if o.n == 0 {
		return nil
	}
	n, err := o.wr.Write(o.buf[0:o.n])
	if n < o.n && err == nil {
		err = io.ErrShortWrite
	}
	if err != nil {
		if n > 0 && n < o.n {
			copy(o.buf[0:o.n-n], o.buf[n:o.n])
		}
		o.n -= n
		o.err = err
		return err
	}

	o.n = 0

	logPrintln("flush done")
	return nil
}
