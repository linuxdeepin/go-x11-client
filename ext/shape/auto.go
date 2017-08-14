package shape

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

// _ns.ext_name: Shape
const MajorVersion = 1
const MinorVersion = 1

var _ext = x.NewExtension("SHAPE")

func Ext() *x.Extension {
	return _ext
}

// simple ('xcb', 'Shape', 'OP')
type Op uint8

// simple ('xcb', 'Shape', 'KIND')
type Kind uint8

// enum SO
const (
	SOSet       = 0
	SOUnion     = 1
	SOIntersect = 2
	SOSubtract  = 3
	SOInvert    = 4
)

// enum SK
const (
	SKBounding = 0
	SKClip     = 1
	SKInput    = 2
)

// event name: ('xcb', 'Shape', 'Notify')
// self.name: ('xcb', 'Shape', 'Notify')
const NotifyEventCode = 0

// not event copy
type NotifyEvent struct {
	ResponseType   uint8
	ShapeKind      Kind
	Sequence       uint16
	AffectedWindow x.Window
	ExtentsX       int16
	ExtentsY       int16
	ExtentsWidth   uint16
	ExtentsHeight  uint16
	ServerTime     x.Timestamp
	Shaped         uint8
	// Pad0 [11]uint8
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

	// read field ShapeKind
	v.ShapeKind = Kind(r.Read1b())
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

	// read field AffectedWindow
	v.AffectedWindow = x.Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field ExtentsX
	v.ExtentsX = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field ExtentsY
	v.ExtentsY = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field ExtentsWidth
	v.ExtentsWidth = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field ExtentsHeight
	v.ExtentsHeight = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field ServerTime
	v.ServerTime = x.Timestamp(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Shaped
	v.Shaped = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(11)
	n += 11
	if r.Err() != nil {
		return
	}
	return
}
func NotifyEventWrite(w *x.Writer, v *NotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field ShapeKind
	w.Write1b(uint8(v.ShapeKind))
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field AffectedWindow
	w.Write4b(uint32(v.AffectedWindow))
	n += 4

	// write field ExtentsX
	w.Write2b(uint16(v.ExtentsX))
	n += 2

	// write field ExtentsY
	w.Write2b(uint16(v.ExtentsY))
	n += 2

	// write field ExtentsWidth
	w.Write2b(v.ExtentsWidth)
	n += 2

	// write field ExtentsHeight
	w.Write2b(v.ExtentsHeight)
	n += 2

	// write field ServerTime
	w.Write4b(uint32(v.ServerTime))
	n += 4

	// write field Shaped
	w.Write1b(v.Shaped)
	n += 1

	// write field Pad0
	w.WritePad(11)
	n += 11
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
	return fmt.Sprintf("NotifyEvent {ShapeKind: %v, Sequence: %v, AffectedWindow: %v, ExtentsX: %v, ExtentsY: %v, ExtentsWidth: %v, ExtentsHeight: %v, ServerTime: %v, Shaped: %v}",
		v.ShapeKind, v.Sequence, v.AffectedWindow, v.ExtentsX, v.ExtentsY, v.ExtentsWidth, v.ExtentsHeight, v.ServerTime, v.Shaped)
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

const QueryVersionOpcode = 0

func QueryVersionRequestWrite(w *x.Writer) {
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
	MajorVersion uint16
	MinorVersion uint16
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
	v.MajorVersion = r.Read2b()
	n += 2
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

func QueryVersion(conn *x.Conn) QueryVersionCookie {
	w := x.NewWriter()
	QueryVersionRequestWrite(w)
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

func QueryVersionUnchecked(conn *x.Conn) QueryVersionCookie {
	w := x.NewWriter()
	QueryVersionRequestWrite(w)
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

const RectanglesOpcode = 1

func RectanglesRequestWrite(w *x.Writer, Operation Op, DestinationKind Kind, Ordering uint8, DestinationWindow x.Window, XOffset int16, YOffset int16, RectanglesLen uint32, Rectangles []x.Rectangle) {
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

	// write wire field Operation
	w.Write1b(uint8(Operation))
	n += 1

	// write wire field DestinationKind
	w.Write1b(uint8(DestinationKind))
	n += 1

	// write wire field Ordering
	w.Write1b(Ordering)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field DestinationWindow
	w.Write4b(uint32(DestinationWindow))
	n += 4

	// write wire field XOffset
	w.Write2b(uint16(XOffset))
	n += 2

	// write wire field YOffset
	w.Write2b(uint16(YOffset))
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Rectangles
	n += x.RectangleWriteN(w, Rectangles)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func Rectangles(conn *x.Conn, Operation Op, DestinationKind Kind, Ordering uint8, DestinationWindow x.Window, XOffset int16, YOffset int16, RectanglesLen uint32, Rectangles []x.Rectangle) x.VoidCookie {
	w := x.NewWriter()
	RectanglesRequestWrite(w, Operation, DestinationKind, Ordering, DestinationWindow, XOffset, YOffset, RectanglesLen, Rectangles)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: RectanglesOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func RectanglesChecked(conn *x.Conn, Operation Op, DestinationKind Kind, Ordering uint8, DestinationWindow x.Window, XOffset int16, YOffset int16, RectanglesLen uint32, Rectangles []x.Rectangle) x.VoidCookie {
	w := x.NewWriter()
	RectanglesRequestWrite(w, Operation, DestinationKind, Ordering, DestinationWindow, XOffset, YOffset, RectanglesLen, Rectangles)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: RectanglesOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const MaskOpcode = 2

func MaskRequestWrite(w *x.Writer, Operation Op, DestinationKind Kind, DestinationWindow x.Window, XOffset int16, YOffset int16, SourceBitmap x.Pixmap) {
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

	// write wire field Operation
	w.Write1b(uint8(Operation))
	n += 1

	// write wire field DestinationKind
	w.Write1b(uint8(DestinationKind))
	n += 1

	// write wire field Pad0
	w.WritePad(2)
	n += 2

	// write wire field DestinationWindow
	w.Write4b(uint32(DestinationWindow))
	n += 4

	// write wire field XOffset
	w.Write2b(uint16(XOffset))
	n += 2

	// write wire field YOffset
	w.Write2b(uint16(YOffset))
	n += 2

	// write wire field SourceBitmap
	w.Write4b(uint32(SourceBitmap))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func Mask(conn *x.Conn, Operation Op, DestinationKind Kind, DestinationWindow x.Window, XOffset int16, YOffset int16, SourceBitmap x.Pixmap) x.VoidCookie {
	w := x.NewWriter()
	MaskRequestWrite(w, Operation, DestinationKind, DestinationWindow, XOffset, YOffset, SourceBitmap)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: MaskOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func MaskChecked(conn *x.Conn, Operation Op, DestinationKind Kind, DestinationWindow x.Window, XOffset int16, YOffset int16, SourceBitmap x.Pixmap) x.VoidCookie {
	w := x.NewWriter()
	MaskRequestWrite(w, Operation, DestinationKind, DestinationWindow, XOffset, YOffset, SourceBitmap)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: MaskOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const CombineOpcode = 3

func CombineRequestWrite(w *x.Writer, Operation Op, DestinationKind Kind, SourceKind Kind, DestinationWindow x.Window, XOffset int16, YOffset int16, SourceWindow x.Window) {
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

	// write wire field Operation
	w.Write1b(uint8(Operation))
	n += 1

	// write wire field DestinationKind
	w.Write1b(uint8(DestinationKind))
	n += 1

	// write wire field SourceKind
	w.Write1b(uint8(SourceKind))
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field DestinationWindow
	w.Write4b(uint32(DestinationWindow))
	n += 4

	// write wire field XOffset
	w.Write2b(uint16(XOffset))
	n += 2

	// write wire field YOffset
	w.Write2b(uint16(YOffset))
	n += 2

	// write wire field SourceWindow
	w.Write4b(uint32(SourceWindow))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func Combine(conn *x.Conn, Operation Op, DestinationKind Kind, SourceKind Kind, DestinationWindow x.Window, XOffset int16, YOffset int16, SourceWindow x.Window) x.VoidCookie {
	w := x.NewWriter()
	CombineRequestWrite(w, Operation, DestinationKind, SourceKind, DestinationWindow, XOffset, YOffset, SourceWindow)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CombineOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CombineChecked(conn *x.Conn, Operation Op, DestinationKind Kind, SourceKind Kind, DestinationWindow x.Window, XOffset int16, YOffset int16, SourceWindow x.Window) x.VoidCookie {
	w := x.NewWriter()
	CombineRequestWrite(w, Operation, DestinationKind, SourceKind, DestinationWindow, XOffset, YOffset, SourceWindow)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CombineOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const OffsetOpcode = 4

func OffsetRequestWrite(w *x.Writer, DestinationKind Kind, DestinationWindow x.Window, XOffset int16, YOffset int16) {
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

	// write wire field DestinationKind
	w.Write1b(uint8(DestinationKind))
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3

	// write wire field DestinationWindow
	w.Write4b(uint32(DestinationWindow))
	n += 4

	// write wire field XOffset
	w.Write2b(uint16(XOffset))
	n += 2

	// write wire field YOffset
	w.Write2b(uint16(YOffset))
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func Offset(conn *x.Conn, DestinationKind Kind, DestinationWindow x.Window, XOffset int16, YOffset int16) x.VoidCookie {
	w := x.NewWriter()
	OffsetRequestWrite(w, DestinationKind, DestinationWindow, XOffset, YOffset)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: OffsetOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func OffsetChecked(conn *x.Conn, DestinationKind Kind, DestinationWindow x.Window, XOffset int16, YOffset int16) x.VoidCookie {
	w := x.NewWriter()
	OffsetRequestWrite(w, DestinationKind, DestinationWindow, XOffset, YOffset)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: OffsetOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const QueryExtentsOpcode = 5

func QueryExtentsRequestWrite(w *x.Writer, DestinationWindow x.Window) {
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

	// write wire field DestinationWindow
	w.Write4b(uint32(DestinationWindow))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type QueryExtentsReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence       uint16
	Length         uint32
	BoundingShaped uint8
	ClipShaped     uint8
	// Pad1 [2]uint8
	BoundingShapeExtentsX      int16
	BoundingShapeExtentsY      int16
	BoundingShapeExtentsWidth  uint16
	BoundingShapeExtentsHeight uint16
	ClipShapeExtentsX          int16
	ClipShapeExtentsY          int16
	ClipShapeExtentsWidth      uint16
	ClipShapeExtentsHeight     uint16
}

func QueryExtentsReplyRead(r *x.Reader, v *QueryExtentsReply) (n int) {
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

	// read field BoundingShaped
	v.BoundingShaped = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field ClipShaped
	v.ClipShaped = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(2)
	n += 2
	if r.Err() != nil {
		return
	}

	// read field BoundingShapeExtentsX
	v.BoundingShapeExtentsX = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field BoundingShapeExtentsY
	v.BoundingShapeExtentsY = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field BoundingShapeExtentsWidth
	v.BoundingShapeExtentsWidth = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field BoundingShapeExtentsHeight
	v.BoundingShapeExtentsHeight = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field ClipShapeExtentsX
	v.ClipShapeExtentsX = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field ClipShapeExtentsY
	v.ClipShapeExtentsY = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field ClipShapeExtentsWidth
	v.ClipShapeExtentsWidth = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field ClipShapeExtentsHeight
	v.ClipShapeExtentsHeight = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

type QueryExtentsCookie uint64

func (cookie QueryExtentsCookie) Reply(conn *x.Conn) (*QueryExtentsReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewGenericError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply QueryExtentsReply
	QueryExtentsReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryExtents(conn *x.Conn, DestinationWindow x.Window) QueryExtentsCookie {
	w := x.NewWriter()
	QueryExtentsRequestWrite(w, DestinationWindow)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: QueryExtentsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return QueryExtentsCookie(seq)
}

func QueryExtentsUnchecked(conn *x.Conn, DestinationWindow x.Window) QueryExtentsCookie {
	w := x.NewWriter()
	QueryExtentsRequestWrite(w, DestinationWindow)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: QueryExtentsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return QueryExtentsCookie(seq)
}

const SelectInputOpcode = 6

func SelectInputRequestWrite(w *x.Writer, DestinationWindow x.Window, Enable uint8) {
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

	// write wire field DestinationWindow
	w.Write4b(uint32(DestinationWindow))
	n += 4

	// write wire field Enable
	w.Write1b(Enable)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SelectInput(conn *x.Conn, DestinationWindow x.Window, Enable uint8) x.VoidCookie {
	w := x.NewWriter()
	SelectInputRequestWrite(w, DestinationWindow, Enable)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SelectInputOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func SelectInputChecked(conn *x.Conn, DestinationWindow x.Window, Enable uint8) x.VoidCookie {
	w := x.NewWriter()
	SelectInputRequestWrite(w, DestinationWindow, Enable)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SelectInputOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const InputSelectedOpcode = 7

func InputSelectedRequestWrite(w *x.Writer, DestinationWindow x.Window) {
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

	// write wire field DestinationWindow
	w.Write4b(uint32(DestinationWindow))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type InputSelectedReply struct {
	ResponseType uint8
	Enabled      uint8
	Sequence     uint16
	Length       uint32
}

func InputSelectedReplyRead(r *x.Reader, v *InputSelectedReply) (n int) {
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

	// read field Enabled
	v.Enabled = r.Read1b()
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

type InputSelectedCookie uint64

func (cookie InputSelectedCookie) Reply(conn *x.Conn) (*InputSelectedReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewGenericError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply InputSelectedReply
	InputSelectedReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func InputSelected(conn *x.Conn, DestinationWindow x.Window) InputSelectedCookie {
	w := x.NewWriter()
	InputSelectedRequestWrite(w, DestinationWindow)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: InputSelectedOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return InputSelectedCookie(seq)
}

func InputSelectedUnchecked(conn *x.Conn, DestinationWindow x.Window) InputSelectedCookie {
	w := x.NewWriter()
	InputSelectedRequestWrite(w, DestinationWindow)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: InputSelectedOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return InputSelectedCookie(seq)
}

const GetRectanglesOpcode = 8

func GetRectanglesRequestWrite(w *x.Writer, Window x.Window, SourceKind Kind) {
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

	// write wire field SourceKind
	w.Write1b(uint8(SourceKind))
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetRectanglesReply struct {
	ResponseType  uint8
	Ordering      uint8
	Sequence      uint16
	Length        uint32
	RectanglesLen uint32
	// Pad0 [20]uint8
	Rectangles []x.Rectangle
}

func GetRectanglesReplyRead(r *x.Reader, v *GetRectanglesReply) (n int) {
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

	// read field Ordering
	v.Ordering = r.Read1b()
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

	// read field RectanglesLen
	v.RectanglesLen = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(20)
	n += 20
	if r.Err() != nil {
		return
	}

	// read field Rectangles
	v.Rectangles = make([]x.Rectangle, int(v.RectanglesLen))
	blockLen += x.RectangleReadN(r, v.Rectangles)
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(x.CRectangle{}))

	/* insert padding begin */
	pad = -blockLen & (alignTo - 1)
	r.ReadPad(pad)
	n += pad
	if r.Err() != nil {
		return
	}
	pad = 0
	blockLen = 0
	/* insert padding end */
	return
}

type GetRectanglesCookie uint64

func (cookie GetRectanglesCookie) Reply(conn *x.Conn) (*GetRectanglesReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewGenericError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetRectanglesReply
	GetRectanglesReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetRectangles(conn *x.Conn, Window x.Window, SourceKind Kind) GetRectanglesCookie {
	w := x.NewWriter()
	GetRectanglesRequestWrite(w, Window, SourceKind)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: GetRectanglesOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return GetRectanglesCookie(seq)
}

func GetRectanglesUnchecked(conn *x.Conn, Window x.Window, SourceKind Kind) GetRectanglesCookie {
	w := x.NewWriter()
	GetRectanglesRequestWrite(w, Window, SourceKind)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: GetRectanglesOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetRectanglesCookie(seq)
}
