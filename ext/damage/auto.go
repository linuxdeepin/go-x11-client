package damage

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

// _ns.ext_name: Damage
const MajorVersion = 1
const MinorVersion = 1

var _ext = x.NewExtension("DAMAGE")

func Ext() *x.Extension {
	return _ext
}

// simple ('xcb', 'Damage', 'DAMAGE')
type Damage uint32

// enum ReportLevel
const (
	ReportLevelRawRectangles   = 0
	ReportLevelDeltaRectangles = 1
	ReportLevelBoundingBox     = 2
	ReportLevelNonEmpty        = 3
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

const CreateOpcode = 1

func CreateRequestWrite(w *x.Writer, Damage Damage, Drawable x.Drawable, Level uint8) {
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

	// write wire field Damage
	w.Write4b(uint32(Damage))
	n += 4

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field Level
	w.Write1b(Level)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func Create(conn *x.Conn, Damage Damage, Drawable x.Drawable, Level uint8) x.VoidCookie {
	w := x.NewWriter()
	CreateRequestWrite(w, Damage, Drawable, Level)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CreateChecked(conn *x.Conn, Damage Damage, Drawable x.Drawable, Level uint8) x.VoidCookie {
	w := x.NewWriter()
	CreateRequestWrite(w, Damage, Drawable, Level)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const DestroyOpcode = 2

func DestroyRequestWrite(w *x.Writer, Damage Damage) {
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

	// write wire field Damage
	w.Write4b(uint32(Damage))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func Destroy(conn *x.Conn, Damage Damage) x.VoidCookie {
	w := x.NewWriter()
	DestroyRequestWrite(w, Damage)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: DestroyOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func DestroyChecked(conn *x.Conn, Damage Damage) x.VoidCookie {
	w := x.NewWriter()
	DestroyRequestWrite(w, Damage)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: DestroyOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const SubtractOpcode = 3

func SubtractRequestWrite(w *x.Writer, Damage Damage, Repair xfixes.Region, Parts xfixes.Region) {
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

	// write wire field Damage
	w.Write4b(uint32(Damage))
	n += 4

	// write wire field Repair
	w.Write4b(uint32(Repair))
	n += 4

	// write wire field Parts
	w.Write4b(uint32(Parts))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func Subtract(conn *x.Conn, Damage Damage, Repair xfixes.Region, Parts xfixes.Region) x.VoidCookie {
	w := x.NewWriter()
	SubtractRequestWrite(w, Damage, Repair, Parts)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SubtractOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func SubtractChecked(conn *x.Conn, Damage Damage, Repair xfixes.Region, Parts xfixes.Region) x.VoidCookie {
	w := x.NewWriter()
	SubtractRequestWrite(w, Damage, Repair, Parts)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SubtractOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const AddOpcode = 4

func AddRequestWrite(w *x.Writer, Drawable x.Drawable, Region xfixes.Region) {
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

	// write wire field Drawable
	w.Write4b(uint32(Drawable))
	n += 4

	// write wire field Region
	w.Write4b(uint32(Region))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func Add(conn *x.Conn, Drawable x.Drawable, Region xfixes.Region) x.VoidCookie {
	w := x.NewWriter()
	AddRequestWrite(w, Drawable, Region)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: AddOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func AddChecked(conn *x.Conn, Drawable x.Drawable, Region xfixes.Region) x.VoidCookie {
	w := x.NewWriter()
	AddRequestWrite(w, Drawable, Region)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: AddOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

// event name: ('xcb', 'Damage', 'Notify')
// self.name: ('xcb', 'Damage', 'Notify')
const NotifyEventCode = 0

// not event copy
type NotifyEvent struct {
	ResponseType uint8
	Level        uint8
	Sequence     uint16
	Drawable     x.Drawable
	Damage       Damage
	Timestamp    x.Timestamp
	Area         x.Rectangle
	Geometry     x.Rectangle
}

func NotifyEventRead(r *x.Reader, v *NotifyEvent) (n int) {
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

	// read field Level
	v.Level = r.Read1b()
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

	// read field Drawable
	v.Drawable = x.Drawable(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Damage
	v.Damage = Damage(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Timestamp
	v.Timestamp = x.Timestamp(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Area
	n += x.RectangleRead(r, &v.Area)
	if r.Err() != nil {
		return
	}

	// read field Geometry
	n += x.RectangleRead(r, &v.Geometry)
	if r.Err() != nil {
		return
	}
	return
}
func NotifyEventWrite(w *x.Writer, v *NotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Level
	w.Write1b(v.Level)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Drawable
	w.Write4b(uint32(v.Drawable))
	n += 4

	// write field Damage
	w.Write4b(uint32(v.Damage))
	n += 4

	// write field Timestamp
	w.Write4b(uint32(v.Timestamp))
	n += 4

	// write field Area
	n += x.RectangleWrite(w, &v.Area)

	// write field Geometry
	n += x.RectangleWrite(w, &v.Geometry)
	return
}

func NewNotifyEvent(data []byte) (*NotifyEvent, error) {
	var ev NotifyEvent
	r := x.NewReaderFromData(data)
	NotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *NotifyEvent) String() string {
	return fmt.Sprintf("NotifyEvent {Level: %v, Sequence: %v, Drawable: %v, Damage: %v, Timestamp: %v, Area: %v, Geometry: %v}",
		v.Level, v.Sequence, v.Drawable, v.Damage, v.Timestamp, v.Area, v.Geometry)
}

type _NotifyEventHandler struct {
	mu  sync.Mutex
	cbs []func(*NotifyEvent)
}

func (eh *_NotifyEventHandler) Run(ge x.GenericEvent) {
	ev, err := NewNotifyEvent(ge)
	if err != nil {
		log.Println(err)
		return
	}
	eh.mu.Lock()
	cbs := eh.cbs
	eh.mu.Unlock()
	for _, cb := range cbs {
		cb(ev)
	}
}

func (_ *_NotifyEventHandler) Detach(_ uint32) {}

func (eh *_NotifyEventHandler) attach(cb func(ev *NotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs
	newCbs := make([]func(*NotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs = newCbs
	eh.mu.Unlock()
}

func ConnectNotifyEvent(conn *x.Conn, cb func(ev *NotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[NotifyEventCode]
	if !ok {
		newEh := &_NotifyEventHandler{}
		conn.EventHandlers[NotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_NotifyEventHandler).attach(cb)
}
