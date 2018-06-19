package x

import (
	"bytes"
	"io"
)

type Reader struct {
	r   io.Reader
	buf [8]byte
	err error
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		r: r,
	}
}

func NewReaderFromData(b []byte) *Reader {
	// TODO Implement a simple bytes reader
	return NewReader(bytes.NewReader(b))
}

func (r *Reader) Err() error {
	return r.err
}

func (r *Reader) ReadPad(n int) {
	quotient := n / 8
	//log.Println("quotient:", quotient)
	for i := 0; i < quotient; i++ {
		length, err := r.r.Read(r.buf[:])
		if err != nil {
			r.err = err
			return
		}
		if length < 8 {
			r.err = io.ErrUnexpectedEOF
			return
		}
	}
	remainder := n % 8
	//log.Println("remainder:", remainder)
	if remainder > 0 {
		length, err := r.r.Read(r.buf[:remainder])
		if err != nil {
			r.err = err
			return
		}
		if length < remainder {
			r.err = io.ErrUnexpectedEOF
			return
		}
	}
}

func (r *Reader) ReadBytes(n int) []byte {
	// TODO avoid panic makeSlice
	buf := make([]byte, n)
	_, err := io.ReadFull(r.r, buf)
	if err != nil {
		r.err = err
		return nil
	}
	return buf
}

func (r *Reader) ReadString(n int) string {
	return string(r.ReadBytes(n))
}

func (r *Reader) Read1b() uint8 {
	length, err := r.r.Read(r.buf[:1])
	if err != nil {
		r.err = err
		return 0
	}
	if length < 1 {
		r.err = io.ErrUnexpectedEOF
		return 0
	}
	return r.buf[0]
}

func (r *Reader) Read2b() uint16 {
	length, err := r.r.Read(r.buf[:2])
	if err != nil {
		r.err = err
		return 0
	}

	if length < 2 {
		r.err = io.ErrUnexpectedEOF
		return 0
	}

	v := uint16(r.buf[0])
	v |= uint16(r.buf[1]) << 8
	return v
}

func (r *Reader) Read4b() uint32 {
	length, err := r.r.Read(r.buf[:4])
	if err != nil {
		r.err = err
		return 0
	}

	if length < 4 {
		r.err = io.ErrUnexpectedEOF
		return 0
	}

	v := uint32(r.buf[0])
	v |= uint32(r.buf[1]) << 8
	v |= uint32(r.buf[2]) << 16
	v |= uint32(r.buf[3]) << 24
	return v
}

func (r *Reader) Read8b() uint64 {
	length, err := r.r.Read(r.buf[:8])
	if err != nil {
		r.err = err
		return 0
	}

	if length < 8 {
		r.err = io.ErrUnexpectedEOF
		return 0
	}

	v := uint64(r.buf[0])
	v |= uint64(r.buf[1]) << 8
	v |= uint64(r.buf[2]) << 16
	v |= uint64(r.buf[3]) << 24
	v |= uint64(r.buf[4]) << 32
	v |= uint64(r.buf[5]) << 40
	v |= uint64(r.buf[6]) << 48
	v |= uint64(r.buf[7]) << 56
	return v
}
