package composite

import x "github.com/linuxdeepin/go-x11-client"

// module.direct_imports: [('xproto', 'xproto'), ('xfixes', 'xfixes')]
import "github.com/linuxdeepin/go-x11-client/ext/xfixes"
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

// _ns.ext_name: Composite
const MajorVersion = 0
const MinorVersion = 4

var _ext = x.NewExtension("Composite")

func Ext() *x.Extension {
	return _ext
}

// enum Redirect
const (
	RedirectAutomatic = 0
	RedirectManual    = 1
)

const QueryVersionOpcode = 0

func QueryVersionRequestWrite(w *x.Writer, ClientMajorVersion uint32, ClientMinorVersion uint32) {
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

	// write wire field ClientMajorVersion
	w.Write4b(ClientMajorVersion)
	n += 4

	// write wire field ClientMinorVersion
	w.Write4b(ClientMinorVersion)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type QueryVersionReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence     uint16
	Length       uint32
	MajorVersion uint32
	MinorVersion uint32
	// Pad1 [16]uint8
}

func QueryVersionReplyRead(r *x.Reader, v *QueryVersionReply) (n int) {
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

	// read field MajorVersion
	v.MajorVersion = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field MinorVersion
	v.MinorVersion = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(16)
	n += 16
	if r.Err() != nil {
		return
	}
	return
}

type QueryVersionCookie uint64

func (cookie QueryVersionCookie) Reply(conn *x.Conn) (*QueryVersionReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewGenericError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply QueryVersionReply
	QueryVersionReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryVersion(conn *x.Conn, ClientMajorVersion uint32, ClientMinorVersion uint32) QueryVersionCookie {
	w := x.NewWriter()
	QueryVersionRequestWrite(w, ClientMajorVersion, ClientMinorVersion)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: QueryVersionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return QueryVersionCookie(seq)
}

func QueryVersionUnchecked(conn *x.Conn, ClientMajorVersion uint32, ClientMinorVersion uint32) QueryVersionCookie {
	w := x.NewWriter()
	QueryVersionRequestWrite(w, ClientMajorVersion, ClientMinorVersion)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: QueryVersionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return QueryVersionCookie(seq)
}

const RedirectWindowOpcode = 1

func RedirectWindowRequestWrite(w *x.Writer, Window x.Window, Update uint8) {
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

	// write wire field Update
	w.Write1b(Update)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func RedirectWindow(conn *x.Conn, Window x.Window, Update uint8) x.VoidCookie {
	w := x.NewWriter()
	RedirectWindowRequestWrite(w, Window, Update)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: RedirectWindowOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func RedirectWindowChecked(conn *x.Conn, Window x.Window, Update uint8) x.VoidCookie {
	w := x.NewWriter()
	RedirectWindowRequestWrite(w, Window, Update)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: RedirectWindowOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const RedirectSubwindowsOpcode = 2

func RedirectSubwindowsRequestWrite(w *x.Writer, Window x.Window, Update uint8) {
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

	// write wire field Update
	w.Write1b(Update)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func RedirectSubwindows(conn *x.Conn, Window x.Window, Update uint8) x.VoidCookie {
	w := x.NewWriter()
	RedirectSubwindowsRequestWrite(w, Window, Update)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: RedirectSubwindowsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func RedirectSubwindowsChecked(conn *x.Conn, Window x.Window, Update uint8) x.VoidCookie {
	w := x.NewWriter()
	RedirectSubwindowsRequestWrite(w, Window, Update)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: RedirectSubwindowsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const UnredirectWindowOpcode = 3

func UnredirectWindowRequestWrite(w *x.Writer, Window x.Window, Update uint8) {
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

	// write wire field Update
	w.Write1b(Update)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func UnredirectWindow(conn *x.Conn, Window x.Window, Update uint8) x.VoidCookie {
	w := x.NewWriter()
	UnredirectWindowRequestWrite(w, Window, Update)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: UnredirectWindowOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func UnredirectWindowChecked(conn *x.Conn, Window x.Window, Update uint8) x.VoidCookie {
	w := x.NewWriter()
	UnredirectWindowRequestWrite(w, Window, Update)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: UnredirectWindowOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const UnredirectSubwindowsOpcode = 4

func UnredirectSubwindowsRequestWrite(w *x.Writer, Window x.Window, Update uint8) {
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

	// write wire field Update
	w.Write1b(Update)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func UnredirectSubwindows(conn *x.Conn, Window x.Window, Update uint8) x.VoidCookie {
	w := x.NewWriter()
	UnredirectSubwindowsRequestWrite(w, Window, Update)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: UnredirectSubwindowsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func UnredirectSubwindowsChecked(conn *x.Conn, Window x.Window, Update uint8) x.VoidCookie {
	w := x.NewWriter()
	UnredirectSubwindowsRequestWrite(w, Window, Update)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: UnredirectSubwindowsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const CreateRegionFromBorderClipOpcode = 5

func CreateRegionFromBorderClipRequestWrite(w *x.Writer, Region xfixes.Region, Window x.Window) {
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

	// write wire field Region
	w.Write4b(uint32(Region))
	n += 4

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CreateRegionFromBorderClip(conn *x.Conn, Region xfixes.Region, Window x.Window) x.VoidCookie {
	w := x.NewWriter()
	CreateRegionFromBorderClipRequestWrite(w, Region, Window)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateRegionFromBorderClipOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CreateRegionFromBorderClipChecked(conn *x.Conn, Region xfixes.Region, Window x.Window) x.VoidCookie {
	w := x.NewWriter()
	CreateRegionFromBorderClipRequestWrite(w, Region, Window)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateRegionFromBorderClipOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const NameWindowPixmapOpcode = 6

func NameWindowPixmapRequestWrite(w *x.Writer, Window x.Window, Pixmap x.Pixmap) {
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

	// write wire field Pixmap
	w.Write4b(uint32(Pixmap))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func NameWindowPixmap(conn *x.Conn, Window x.Window, Pixmap x.Pixmap) x.VoidCookie {
	w := x.NewWriter()
	NameWindowPixmapRequestWrite(w, Window, Pixmap)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: NameWindowPixmapOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func NameWindowPixmapChecked(conn *x.Conn, Window x.Window, Pixmap x.Pixmap) x.VoidCookie {
	w := x.NewWriter()
	NameWindowPixmapRequestWrite(w, Window, Pixmap)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: NameWindowPixmapOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const GetOverlayWindowOpcode = 7

func GetOverlayWindowRequestWrite(w *x.Writer, Window x.Window) {
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
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetOverlayWindowReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence   uint16
	Length     uint32
	OverlayWin x.Window
	// Pad1 [20]uint8
}

func GetOverlayWindowReplyRead(r *x.Reader, v *GetOverlayWindowReply) (n int) {
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

	// read field OverlayWin
	v.OverlayWin = x.Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(20)
	n += 20
	if r.Err() != nil {
		return
	}
	return
}

type GetOverlayWindowCookie uint64

func (cookie GetOverlayWindowCookie) Reply(conn *x.Conn) (*GetOverlayWindowReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewGenericError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetOverlayWindowReply
	GetOverlayWindowReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetOverlayWindow(conn *x.Conn, Window x.Window) GetOverlayWindowCookie {
	w := x.NewWriter()
	GetOverlayWindowRequestWrite(w, Window)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: GetOverlayWindowOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return GetOverlayWindowCookie(seq)
}

func GetOverlayWindowUnchecked(conn *x.Conn, Window x.Window) GetOverlayWindowCookie {
	w := x.NewWriter()
	GetOverlayWindowRequestWrite(w, Window)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: GetOverlayWindowOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetOverlayWindowCookie(seq)
}

const ReleaseOverlayWindowOpcode = 8

func ReleaseOverlayWindowRequestWrite(w *x.Writer, Window x.Window) {
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
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func ReleaseOverlayWindow(conn *x.Conn, Window x.Window) x.VoidCookie {
	w := x.NewWriter()
	ReleaseOverlayWindowRequestWrite(w, Window)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: ReleaseOverlayWindowOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func ReleaseOverlayWindowChecked(conn *x.Conn, Window x.Window) x.VoidCookie {
	w := x.NewWriter()
	ReleaseOverlayWindowRequestWrite(w, Window)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: ReleaseOverlayWindowOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}
