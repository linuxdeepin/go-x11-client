package test

import x "github.com/linuxdeepin/go-x11-client"

// module.direct_imports: [('xproto', 'xproto')]
import "unsafe"
import "log"
import "sync"
import "fmt"
import "strings"

var _ = x.Reader{}
var _ = unsafe.Alignof(0)
var _ sync.Mutex
var _ = fmt.Println
var _ = strings.Join

func init() {
	log.SetFlags(log.Lshortfile)
}

// _ns.ext_name: Test
const MajorVersion = 2
const MinorVersion = 2

var _ext = x.NewExtension("XTEST")

func Ext() *x.Extension {
	return _ext
}

const GetVersionOpcode = 0

func GetVersionRequestWrite(w *x.Writer, MajorVersion uint8, MinorVersion uint16) {
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

	// write wire field MajorVersion
	w.Write1b(MajorVersion)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field MinorVersion
	w.Write2b(MinorVersion)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetVersionReply struct {
	ResponseType uint8
	MajorVersion uint8
	Sequence     uint16
	Length       uint32
	MinorVersion uint16
}

func GetVersionReplyRead(r *x.Reader, v *GetVersionReply) (n int) {
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

	// read field MajorVersion
	v.MajorVersion = r.Read1b()
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

	// read field MinorVersion
	v.MinorVersion = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

type GetVersionCookie uint64

func (cookie GetVersionCookie) Reply(conn *x.Conn) (*GetVersionReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewGenericError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetVersionReply
	GetVersionReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetVersion(conn *x.Conn, MajorVersion uint8, MinorVersion uint16) GetVersionCookie {
	w := x.NewWriter()
	GetVersionRequestWrite(w, MajorVersion, MinorVersion)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: GetVersionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return GetVersionCookie(seq)
}

func GetVersionUnchecked(conn *x.Conn, MajorVersion uint8, MinorVersion uint16) GetVersionCookie {
	w := x.NewWriter()
	GetVersionRequestWrite(w, MajorVersion, MinorVersion)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: GetVersionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetVersionCookie(seq)
}

// enum Cursor
const (
	CursorNone    = 0
	CursorCurrent = 1
)

const CompareCursorOpcode = 1

func CompareCursorRequestWrite(w *x.Writer, Window x.Window, Cursor x.Cursor) {
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

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4

	// write wire field Cursor
	w.Write4b(uint32(Cursor))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type CompareCursorReply struct {
	ResponseType uint8
	Same         uint8
	Sequence     uint16
	Length       uint32
}

func CompareCursorReplyRead(r *x.Reader, v *CompareCursorReply) (n int) {
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

	// read field Same
	v.Same = r.Read1b()
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
	return
}

type CompareCursorCookie uint64

func (cookie CompareCursorCookie) Reply(conn *x.Conn) (*CompareCursorReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewGenericError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply CompareCursorReply
	CompareCursorReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func CompareCursor(conn *x.Conn, Window x.Window, Cursor x.Cursor) CompareCursorCookie {
	w := x.NewWriter()
	CompareCursorRequestWrite(w, Window, Cursor)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: CompareCursorOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return CompareCursorCookie(seq)
}

func CompareCursorUnchecked(conn *x.Conn, Window x.Window, Cursor x.Cursor) CompareCursorCookie {
	w := x.NewWriter()
	CompareCursorRequestWrite(w, Window, Cursor)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: CompareCursorOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return CompareCursorCookie(seq)
}

const FakeInputOpcode = 2

func FakeInputRequestWrite(w *x.Writer, Type uint8, Detail uint8, Time uint32, Root x.Window, Rootx int16, Rooty int16, Deviceid uint8) {
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

	// write wire field Type
	w.Write1b(Type)
	n += 1

	// write wire field Detail
	w.Write1b(Detail)
	n += 1

	// write wire field Pad0
	w.WritePad(2)
	n += 2

	// write wire field Time
	w.Write4b(Time)
	n += 4

	// write wire field Root
	w.Write4b(uint32(Root))
	n += 4

	// write wire field Pad1
	w.WritePad(8)
	n += 8

	// write wire field Rootx
	w.Write2b(uint16(Rootx))
	n += 2

	// write wire field Rooty
	w.Write2b(uint16(Rooty))
	n += 2

	// write wire field Pad2
	w.WritePad(7)
	n += 7

	// write wire field Deviceid
	w.Write1b(Deviceid)
	n += 1
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func FakeInput(conn *x.Conn, Type uint8, Detail uint8, Time uint32, Root x.Window, Rootx int16, Rooty int16, Deviceid uint8) x.VoidCookie {
	w := x.NewWriter()
	FakeInputRequestWrite(w, Type, Detail, Time, Root, Rootx, Rooty, Deviceid)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: FakeInputOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func FakeInputChecked(conn *x.Conn, Type uint8, Detail uint8, Time uint32, Root x.Window, Rootx int16, Rooty int16, Deviceid uint8) x.VoidCookie {
	w := x.NewWriter()
	FakeInputRequestWrite(w, Type, Detail, Time, Root, Rootx, Rooty, Deviceid)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: FakeInputOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const GrabControlOpcode = 3

func GrabControlRequestWrite(w *x.Writer, Impervious uint8) {
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

	// write wire field Impervious
	w.Write1b(Impervious)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func GrabControl(conn *x.Conn, Impervious uint8) x.VoidCookie {
	w := x.NewWriter()
	GrabControlRequestWrite(w, Impervious)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: GrabControlOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func GrabControlChecked(conn *x.Conn, Impervious uint8) x.VoidCookie {
	w := x.NewWriter()
	GrabControlRequestWrite(w, Impervious)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: GrabControlOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}
