package xfixes

import x "github.com/linuxdeepin/go-x11-client"

// module.direct_imports: [('xproto', 'xproto'), ('render', 'render'), ('shape', 'shape')]
import "github.com/linuxdeepin/go-x11-client/ext/render"
import "github.com/linuxdeepin/go-x11-client/ext/shape"
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

// _ns.ext_name: XFixes
const MajorVersion = 5
const MinorVersion = 0

var _ext = x.NewExtension("XFIXES")

func Ext() *x.Extension {
	return _ext
}

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

// enum SaveSetMode
const (
	SaveSetModeInsert = 0
	SaveSetModeDelete = 1
)

// enum SaveSetTarget
const (
	SaveSetTargetNearest = 0
	SaveSetTargetRoot    = 1
)

// enum SaveSetMapping
const (
	SaveSetMappingMap   = 0
	SaveSetMappingUnmap = 1
)

const ChangeSaveSetOpcode = 1

func ChangeSaveSetRequestWrite(w *x.Writer, Mode uint8, Target uint8, Map uint8, Window x.Window) {
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

	// write wire field Mode
	w.Write1b(Mode)
	n += 1

	// write wire field Target
	w.Write1b(Target)
	n += 1

	// write wire field Map
	w.Write1b(Map)
	n += 1

	// write wire field Pad0
	w.WritePad(1)
	n += 1

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func ChangeSaveSet(conn *x.Conn, Mode uint8, Target uint8, Map uint8, Window x.Window) x.VoidCookie {
	w := x.NewWriter()
	ChangeSaveSetRequestWrite(w, Mode, Target, Map, Window)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: ChangeSaveSetOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func ChangeSaveSetChecked(conn *x.Conn, Mode uint8, Target uint8, Map uint8, Window x.Window) x.VoidCookie {
	w := x.NewWriter()
	ChangeSaveSetRequestWrite(w, Mode, Target, Map, Window)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: ChangeSaveSetOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

// enum SelectionEvent
const (
	SelectionEventSetSelectionOwner      = 0
	SelectionEventSelectionWindowDestroy = 1
	SelectionEventSelectionClientClose   = 2
)

// enum SelectionEventMask
const (
	SelectionEventMaskSetSelectionOwner      = 1
	SelectionEventMaskSelectionWindowDestroy = 2
	SelectionEventMaskSelectionClientClose   = 4
)

// event name: ('xcb', 'XFixes', 'SelectionNotify')
// self.name: ('xcb', 'XFixes', 'SelectionNotify')
const SelectionNotifyEventCode = 0

// not event copy
type SelectionNotifyEvent struct {
	ResponseType       uint8
	Subtype            uint8
	Sequence           uint16
	Window             x.Window
	Owner              x.Window
	Selection          x.Atom
	Timestamp          x.Timestamp
	SelectionTimestamp x.Timestamp
	// Pad0 [8]uint8
}

func SelectionNotifyEventRead(r *x.Reader, v *SelectionNotifyEvent) (n int) {
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

	// read field Subtype
	v.Subtype = r.Read1b()
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

	// read field Window
	v.Window = x.Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Owner
	v.Owner = x.Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Selection
	v.Selection = x.Atom(r.Read4b())
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

	// read field SelectionTimestamp
	v.SelectionTimestamp = x.Timestamp(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(8)
	n += 8
	if r.Err() != nil {
		return
	}
	return
}
func SelectionNotifyEventWrite(w *x.Writer, v *SelectionNotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Subtype
	w.Write1b(v.Subtype)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4

	// write field Owner
	w.Write4b(uint32(v.Owner))
	n += 4

	// write field Selection
	w.Write4b(uint32(v.Selection))
	n += 4

	// write field Timestamp
	w.Write4b(uint32(v.Timestamp))
	n += 4

	// write field SelectionTimestamp
	w.Write4b(uint32(v.SelectionTimestamp))
	n += 4

	// write field Pad0
	w.WritePad(8)
	n += 8
	return
}

func NewSelectionNotifyEvent(data []byte) (*SelectionNotifyEvent, error) {
	var ev SelectionNotifyEvent
	r := x.NewReaderFromData(data)
	SelectionNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *SelectionNotifyEvent) String() string {
	return fmt.Sprintf("SelectionNotifyEvent {Subtype: %v, Sequence: %v, Window: %v, Owner: %v, Selection: %v, Timestamp: %v, SelectionTimestamp: %v}",
		v.Subtype, v.Sequence, v.Window, v.Owner, v.Selection, v.Timestamp, v.SelectionTimestamp)
}

type _SelectionNotifyEventHandler struct {
	mu  sync.Mutex
	cbs []func(*SelectionNotifyEvent)
}

func (eh *_SelectionNotifyEventHandler) Run(ge x.GenericEvent) {
	ev, err := NewSelectionNotifyEvent(ge)
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

func (_ *_SelectionNotifyEventHandler) Detach(_ uint32) {}

func (eh *_SelectionNotifyEventHandler) attach(cb func(ev *SelectionNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs
	newCbs := make([]func(*SelectionNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs = newCbs
	eh.mu.Unlock()
}

func ConnectSelectionNotifyEvent(conn *x.Conn, cb func(ev *SelectionNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[SelectionNotifyEventCode]
	if !ok {
		newEh := &_SelectionNotifyEventHandler{}
		conn.EventHandlers[SelectionNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_SelectionNotifyEventHandler).attach(cb)
}

const SelectSelectionInputOpcode = 2

func SelectSelectionInputRequestWrite(w *x.Writer, Window x.Window, Selection x.Atom, EventMask uint32) {
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

	// write wire field Selection
	w.Write4b(uint32(Selection))
	n += 4

	// write wire field EventMask
	w.Write4b(EventMask)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SelectSelectionInput(conn *x.Conn, Window x.Window, Selection x.Atom, EventMask uint32) x.VoidCookie {
	w := x.NewWriter()
	SelectSelectionInputRequestWrite(w, Window, Selection, EventMask)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SelectSelectionInputOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func SelectSelectionInputChecked(conn *x.Conn, Window x.Window, Selection x.Atom, EventMask uint32) x.VoidCookie {
	w := x.NewWriter()
	SelectSelectionInputRequestWrite(w, Window, Selection, EventMask)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SelectSelectionInputOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

// enum CursorNotify
const (
	CursorNotifyDisplayCursor = 0
)

// enum CursorNotifyMask
const (
	CursorNotifyMaskDisplayCursor = 1
)

// event name: ('xcb', 'XFixes', 'CursorNotify')
// self.name: ('xcb', 'XFixes', 'CursorNotify')
const CursorNotifyEventCode = 1

// not event copy
type CursorNotifyEvent struct {
	ResponseType uint8
	Subtype      uint8
	Sequence     uint16
	Window       x.Window
	CursorSerial uint32
	Timestamp    x.Timestamp
	Name         x.Atom
	// Pad0 [12]uint8
}

func CursorNotifyEventRead(r *x.Reader, v *CursorNotifyEvent) (n int) {
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

	// read field Subtype
	v.Subtype = r.Read1b()
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

	// read field Window
	v.Window = x.Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field CursorSerial
	v.CursorSerial = r.Read4b()
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

	// read field Name
	v.Name = x.Atom(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(12)
	n += 12
	if r.Err() != nil {
		return
	}
	return
}
func CursorNotifyEventWrite(w *x.Writer, v *CursorNotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field Subtype
	w.Write1b(v.Subtype)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4

	// write field CursorSerial
	w.Write4b(v.CursorSerial)
	n += 4

	// write field Timestamp
	w.Write4b(uint32(v.Timestamp))
	n += 4

	// write field Name
	w.Write4b(uint32(v.Name))
	n += 4

	// write field Pad0
	w.WritePad(12)
	n += 12
	return
}

func NewCursorNotifyEvent(data []byte) (*CursorNotifyEvent, error) {
	var ev CursorNotifyEvent
	r := x.NewReaderFromData(data)
	CursorNotifyEventRead(r, &ev)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &ev, nil
}

func (v *CursorNotifyEvent) String() string {
	return fmt.Sprintf("CursorNotifyEvent {Subtype: %v, Sequence: %v, Window: %v, CursorSerial: %v, Timestamp: %v, Name: %v}",
		v.Subtype, v.Sequence, v.Window, v.CursorSerial, v.Timestamp, v.Name)
}

type _CursorNotifyEventHandler struct {
	mu  sync.Mutex
	cbs []func(*CursorNotifyEvent)
}

func (eh *_CursorNotifyEventHandler) Run(ge x.GenericEvent) {
	ev, err := NewCursorNotifyEvent(ge)
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

func (_ *_CursorNotifyEventHandler) Detach(_ uint32) {}

func (eh *_CursorNotifyEventHandler) attach(cb func(ev *CursorNotifyEvent)) {
	eh.mu.Lock()
	oldCbs := eh.cbs
	newCbs := make([]func(*CursorNotifyEvent), len(oldCbs)+1)
	copy(newCbs, oldCbs)
	newCbs[len(newCbs)-1] = cb
	eh.cbs = newCbs
	eh.mu.Unlock()
}

func ConnectCursorNotifyEvent(conn *x.Conn, cb func(ev *CursorNotifyEvent)) {
	conn.EventHandlersMu.Lock()
	eh, ok := conn.EventHandlers[CursorNotifyEventCode]
	if !ok {
		newEh := &_CursorNotifyEventHandler{}
		conn.EventHandlers[CursorNotifyEventCode] = newEh
		conn.EventHandlersMu.Unlock()
		newEh.attach(cb)
		return
	}
	conn.EventHandlersMu.Unlock()
	eh.(*_CursorNotifyEventHandler).attach(cb)
}

const SelectCursorInputOpcode = 3

func SelectCursorInputRequestWrite(w *x.Writer, Window x.Window, EventMask uint32) {
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

	// write wire field EventMask
	w.Write4b(EventMask)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SelectCursorInput(conn *x.Conn, Window x.Window, EventMask uint32) x.VoidCookie {
	w := x.NewWriter()
	SelectCursorInputRequestWrite(w, Window, EventMask)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SelectCursorInputOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func SelectCursorInputChecked(conn *x.Conn, Window x.Window, EventMask uint32) x.VoidCookie {
	w := x.NewWriter()
	SelectCursorInputRequestWrite(w, Window, EventMask)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SelectCursorInputOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const GetCursorImageOpcode = 4

func GetCursorImageRequestWrite(w *x.Writer) {
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

type GetCursorImageReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence     uint16
	Length       uint32
	X            int16
	Y            int16
	Width        uint16
	Height       uint16
	Xhot         uint16
	Yhot         uint16
	CursorSerial uint32
	// Pad1 [8]uint8
	CursorImage []uint32
}

func GetCursorImageReplyRead(r *x.Reader, v *GetCursorImageReply) (n int) {
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

	// read field X
	v.X = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Y
	v.Y = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Width
	v.Width = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Height
	v.Height = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Xhot
	v.Xhot = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Yhot
	v.Yhot = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field CursorSerial
	v.CursorSerial = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(8)
	n += 8
	if r.Err() != nil {
		return
	}

	// read field CursorImage
	{
		dataLen := int((int(v.Width) * int(v.Height)))
		data := make([]uint32, dataLen)
		for i := 0; i < dataLen; i++ {
			data[i] = uint32(r.Read4b())
			if r.Err() != nil {
				return
			}
		}
		blockLen += dataLen * 4
		n += blockLen
		v.CursorImage = data
	}
	alignTo = int(unsafe.Alignof(uint32(0)))

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

type GetCursorImageCookie uint64

func (cookie GetCursorImageCookie) Reply(conn *x.Conn) (*GetCursorImageReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewGenericError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetCursorImageReply
	GetCursorImageReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetCursorImage(conn *x.Conn) GetCursorImageCookie {
	w := x.NewWriter()
	GetCursorImageRequestWrite(w)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: GetCursorImageOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return GetCursorImageCookie(seq)
}

func GetCursorImageUnchecked(conn *x.Conn) GetCursorImageCookie {
	w := x.NewWriter()
	GetCursorImageRequestWrite(w)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: GetCursorImageOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetCursorImageCookie(seq)
}

// simple ('xcb', 'XFixes', 'REGION')
type Region uint32

// enum Region
const (
	RegionNone = 0
)

const CreateRegionOpcode = 5

func CreateRegionRequestWrite(w *x.Writer, Region Region, RectanglesLen uint32, Rectangles []x.Rectangle) {
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

func CreateRegion(conn *x.Conn, Region Region, RectanglesLen uint32, Rectangles []x.Rectangle) x.VoidCookie {
	w := x.NewWriter()
	CreateRegionRequestWrite(w, Region, RectanglesLen, Rectangles)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateRegionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CreateRegionChecked(conn *x.Conn, Region Region, RectanglesLen uint32, Rectangles []x.Rectangle) x.VoidCookie {
	w := x.NewWriter()
	CreateRegionRequestWrite(w, Region, RectanglesLen, Rectangles)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateRegionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const CreateRegionFromBitmapOpcode = 6

func CreateRegionFromBitmapRequestWrite(w *x.Writer, Region Region, Bitmap x.Pixmap) {
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

	// write wire field Bitmap
	w.Write4b(uint32(Bitmap))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CreateRegionFromBitmap(conn *x.Conn, Region Region, Bitmap x.Pixmap) x.VoidCookie {
	w := x.NewWriter()
	CreateRegionFromBitmapRequestWrite(w, Region, Bitmap)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateRegionFromBitmapOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CreateRegionFromBitmapChecked(conn *x.Conn, Region Region, Bitmap x.Pixmap) x.VoidCookie {
	w := x.NewWriter()
	CreateRegionFromBitmapRequestWrite(w, Region, Bitmap)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateRegionFromBitmapOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const CreateRegionFromWindowOpcode = 7

func CreateRegionFromWindowRequestWrite(w *x.Writer, Region Region, Window x.Window, Kind shape.Kind) {
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

	// write wire field Kind
	w.Write1b(uint8(Kind))
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CreateRegionFromWindow(conn *x.Conn, Region Region, Window x.Window, Kind shape.Kind) x.VoidCookie {
	w := x.NewWriter()
	CreateRegionFromWindowRequestWrite(w, Region, Window, Kind)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateRegionFromWindowOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CreateRegionFromWindowChecked(conn *x.Conn, Region Region, Window x.Window, Kind shape.Kind) x.VoidCookie {
	w := x.NewWriter()
	CreateRegionFromWindowRequestWrite(w, Region, Window, Kind)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateRegionFromWindowOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const CreateRegionFromGCOpcode = 8

func CreateRegionFromGcRequestWrite(w *x.Writer, Region Region, Gc x.GContext) {
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

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CreateRegionFromGC(conn *x.Conn, Region Region, Gc x.GContext) x.VoidCookie {
	w := x.NewWriter()
	CreateRegionFromGcRequestWrite(w, Region, Gc)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateRegionFromGCOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CreateRegionFromGCChecked(conn *x.Conn, Region Region, Gc x.GContext) x.VoidCookie {
	w := x.NewWriter()
	CreateRegionFromGcRequestWrite(w, Region, Gc)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateRegionFromGCOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const CreateRegionFromPictureOpcode = 9

func CreateRegionFromPictureRequestWrite(w *x.Writer, Region Region, Picture render.Picture) {
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

	// write wire field Picture
	w.Write4b(uint32(Picture))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CreateRegionFromPicture(conn *x.Conn, Region Region, Picture render.Picture) x.VoidCookie {
	w := x.NewWriter()
	CreateRegionFromPictureRequestWrite(w, Region, Picture)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateRegionFromPictureOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CreateRegionFromPictureChecked(conn *x.Conn, Region Region, Picture render.Picture) x.VoidCookie {
	w := x.NewWriter()
	CreateRegionFromPictureRequestWrite(w, Region, Picture)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateRegionFromPictureOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const DestroyRegionOpcode = 10

func DestroyRegionRequestWrite(w *x.Writer, Region Region) {
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
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func DestroyRegion(conn *x.Conn, Region Region) x.VoidCookie {
	w := x.NewWriter()
	DestroyRegionRequestWrite(w, Region)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: DestroyRegionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func DestroyRegionChecked(conn *x.Conn, Region Region) x.VoidCookie {
	w := x.NewWriter()
	DestroyRegionRequestWrite(w, Region)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: DestroyRegionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const SetRegionOpcode = 11

func SetRegionRequestWrite(w *x.Writer, Region Region, RectanglesLen uint32, Rectangles []x.Rectangle) {
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

func SetRegion(conn *x.Conn, Region Region, RectanglesLen uint32, Rectangles []x.Rectangle) x.VoidCookie {
	w := x.NewWriter()
	SetRegionRequestWrite(w, Region, RectanglesLen, Rectangles)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SetRegionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func SetRegionChecked(conn *x.Conn, Region Region, RectanglesLen uint32, Rectangles []x.Rectangle) x.VoidCookie {
	w := x.NewWriter()
	SetRegionRequestWrite(w, Region, RectanglesLen, Rectangles)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SetRegionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const CopyRegionOpcode = 12

func CopyRegionRequestWrite(w *x.Writer, Source Region, Destination Region) {
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

	// write wire field Source
	w.Write4b(uint32(Source))
	n += 4

	// write wire field Destination
	w.Write4b(uint32(Destination))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CopyRegion(conn *x.Conn, Source Region, Destination Region) x.VoidCookie {
	w := x.NewWriter()
	CopyRegionRequestWrite(w, Source, Destination)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CopyRegionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CopyRegionChecked(conn *x.Conn, Source Region, Destination Region) x.VoidCookie {
	w := x.NewWriter()
	CopyRegionRequestWrite(w, Source, Destination)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CopyRegionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const UnionRegionOpcode = 13

func UnionRegionRequestWrite(w *x.Writer, Source1 Region, Source2 Region, Destination Region) {
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

	// write wire field Source1
	w.Write4b(uint32(Source1))
	n += 4

	// write wire field Source2
	w.Write4b(uint32(Source2))
	n += 4

	// write wire field Destination
	w.Write4b(uint32(Destination))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func UnionRegion(conn *x.Conn, Source1 Region, Source2 Region, Destination Region) x.VoidCookie {
	w := x.NewWriter()
	UnionRegionRequestWrite(w, Source1, Source2, Destination)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: UnionRegionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func UnionRegionChecked(conn *x.Conn, Source1 Region, Source2 Region, Destination Region) x.VoidCookie {
	w := x.NewWriter()
	UnionRegionRequestWrite(w, Source1, Source2, Destination)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: UnionRegionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const IntersectRegionOpcode = 14

func IntersectRegionRequestWrite(w *x.Writer, Source1 Region, Source2 Region, Destination Region) {
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

	// write wire field Source1
	w.Write4b(uint32(Source1))
	n += 4

	// write wire field Source2
	w.Write4b(uint32(Source2))
	n += 4

	// write wire field Destination
	w.Write4b(uint32(Destination))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func IntersectRegion(conn *x.Conn, Source1 Region, Source2 Region, Destination Region) x.VoidCookie {
	w := x.NewWriter()
	IntersectRegionRequestWrite(w, Source1, Source2, Destination)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: IntersectRegionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func IntersectRegionChecked(conn *x.Conn, Source1 Region, Source2 Region, Destination Region) x.VoidCookie {
	w := x.NewWriter()
	IntersectRegionRequestWrite(w, Source1, Source2, Destination)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: IntersectRegionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const SubtractRegionOpcode = 15

func SubtractRegionRequestWrite(w *x.Writer, Source1 Region, Source2 Region, Destination Region) {
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

	// write wire field Source1
	w.Write4b(uint32(Source1))
	n += 4

	// write wire field Source2
	w.Write4b(uint32(Source2))
	n += 4

	// write wire field Destination
	w.Write4b(uint32(Destination))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SubtractRegion(conn *x.Conn, Source1 Region, Source2 Region, Destination Region) x.VoidCookie {
	w := x.NewWriter()
	SubtractRegionRequestWrite(w, Source1, Source2, Destination)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SubtractRegionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func SubtractRegionChecked(conn *x.Conn, Source1 Region, Source2 Region, Destination Region) x.VoidCookie {
	w := x.NewWriter()
	SubtractRegionRequestWrite(w, Source1, Source2, Destination)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SubtractRegionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const InvertRegionOpcode = 16

func InvertRegionRequestWrite(w *x.Writer, Source Region, Bounds x.Rectangle, Destination Region) {
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

	// write wire field Source
	w.Write4b(uint32(Source))
	n += 4

	// write wire field Bounds
	// TODO RequestWriteDefine container_FS

	// write wire field Destination
	w.Write4b(uint32(Destination))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func InvertRegion(conn *x.Conn, Source Region, Bounds x.Rectangle, Destination Region) x.VoidCookie {
	w := x.NewWriter()
	InvertRegionRequestWrite(w, Source, Bounds, Destination)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: InvertRegionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func InvertRegionChecked(conn *x.Conn, Source Region, Bounds x.Rectangle, Destination Region) x.VoidCookie {
	w := x.NewWriter()
	InvertRegionRequestWrite(w, Source, Bounds, Destination)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: InvertRegionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const TranslateRegionOpcode = 17

func TranslateRegionRequestWrite(w *x.Writer, Region Region, Dx int16, Dy int16) {
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

	// write wire field Dx
	w.Write2b(uint16(Dx))
	n += 2

	// write wire field Dy
	w.Write2b(uint16(Dy))
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func TranslateRegion(conn *x.Conn, Region Region, Dx int16, Dy int16) x.VoidCookie {
	w := x.NewWriter()
	TranslateRegionRequestWrite(w, Region, Dx, Dy)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: TranslateRegionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func TranslateRegionChecked(conn *x.Conn, Region Region, Dx int16, Dy int16) x.VoidCookie {
	w := x.NewWriter()
	TranslateRegionRequestWrite(w, Region, Dx, Dy)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: TranslateRegionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const RegionExtentsOpcode = 18

func RegionExtentsRequestWrite(w *x.Writer, Source Region, Destination Region) {
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

	// write wire field Source
	w.Write4b(uint32(Source))
	n += 4

	// write wire field Destination
	w.Write4b(uint32(Destination))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func RegionExtents(conn *x.Conn, Source Region, Destination Region) x.VoidCookie {
	w := x.NewWriter()
	RegionExtentsRequestWrite(w, Source, Destination)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: RegionExtentsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func RegionExtentsChecked(conn *x.Conn, Source Region, Destination Region) x.VoidCookie {
	w := x.NewWriter()
	RegionExtentsRequestWrite(w, Source, Destination)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: RegionExtentsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const FetchRegionOpcode = 19

func FetchRegionRequestWrite(w *x.Writer, Region Region) {
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
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type FetchRegionReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Length   uint32
	Extents  x.Rectangle
	// Pad1 [16]uint8
	Rectangles []x.Rectangle
}

func FetchRegionReplyRead(r *x.Reader, v *FetchRegionReply) (n int) {
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

	// read field Extents
	n += x.RectangleRead(r, &v.Extents)
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(16)
	n += 16
	if r.Err() != nil {
		return
	}

	// read field Rectangles
	v.Rectangles = make([]x.Rectangle, (int(v.Length) / 2))
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

type FetchRegionCookie uint64

func (cookie FetchRegionCookie) Reply(conn *x.Conn) (*FetchRegionReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewGenericError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply FetchRegionReply
	FetchRegionReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func FetchRegion(conn *x.Conn, Region Region) FetchRegionCookie {
	w := x.NewWriter()
	FetchRegionRequestWrite(w, Region)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: FetchRegionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return FetchRegionCookie(seq)
}

func FetchRegionUnchecked(conn *x.Conn, Region Region) FetchRegionCookie {
	w := x.NewWriter()
	FetchRegionRequestWrite(w, Region)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: FetchRegionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return FetchRegionCookie(seq)
}

const SetGCClipRegionOpcode = 20

func SetGcClipRegionRequestWrite(w *x.Writer, Gc x.GContext, Region Region, XOrigin int16, YOrigin int16) {
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

	// write wire field Gc
	w.Write4b(uint32(Gc))
	n += 4

	// write wire field Region
	w.Write4b(uint32(Region))
	n += 4

	// write wire field XOrigin
	w.Write2b(uint16(XOrigin))
	n += 2

	// write wire field YOrigin
	w.Write2b(uint16(YOrigin))
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SetGCClipRegion(conn *x.Conn, Gc x.GContext, Region Region, XOrigin int16, YOrigin int16) x.VoidCookie {
	w := x.NewWriter()
	SetGcClipRegionRequestWrite(w, Gc, Region, XOrigin, YOrigin)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SetGCClipRegionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func SetGCClipRegionChecked(conn *x.Conn, Gc x.GContext, Region Region, XOrigin int16, YOrigin int16) x.VoidCookie {
	w := x.NewWriter()
	SetGcClipRegionRequestWrite(w, Gc, Region, XOrigin, YOrigin)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SetGCClipRegionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const SetWindowShapeRegionOpcode = 21

func SetWindowShapeRegionRequestWrite(w *x.Writer, Dest x.Window, DestKind shape.Kind, XOffset int16, YOffset int16, Region Region) {
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

	// write wire field Dest
	w.Write4b(uint32(Dest))
	n += 4

	// write wire field DestKind
	w.Write1b(uint8(DestKind))
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3

	// write wire field XOffset
	w.Write2b(uint16(XOffset))
	n += 2

	// write wire field YOffset
	w.Write2b(uint16(YOffset))
	n += 2

	// write wire field Region
	w.Write4b(uint32(Region))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SetWindowShapeRegion(conn *x.Conn, Dest x.Window, DestKind shape.Kind, XOffset int16, YOffset int16, Region Region) x.VoidCookie {
	w := x.NewWriter()
	SetWindowShapeRegionRequestWrite(w, Dest, DestKind, XOffset, YOffset, Region)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SetWindowShapeRegionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func SetWindowShapeRegionChecked(conn *x.Conn, Dest x.Window, DestKind shape.Kind, XOffset int16, YOffset int16, Region Region) x.VoidCookie {
	w := x.NewWriter()
	SetWindowShapeRegionRequestWrite(w, Dest, DestKind, XOffset, YOffset, Region)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SetWindowShapeRegionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const SetPictureClipRegionOpcode = 22

func SetPictureClipRegionRequestWrite(w *x.Writer, Picture render.Picture, Region Region, XOrigin int16, YOrigin int16) {
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

	// write wire field Picture
	w.Write4b(uint32(Picture))
	n += 4

	// write wire field Region
	w.Write4b(uint32(Region))
	n += 4

	// write wire field XOrigin
	w.Write2b(uint16(XOrigin))
	n += 2

	// write wire field YOrigin
	w.Write2b(uint16(YOrigin))
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SetPictureClipRegion(conn *x.Conn, Picture render.Picture, Region Region, XOrigin int16, YOrigin int16) x.VoidCookie {
	w := x.NewWriter()
	SetPictureClipRegionRequestWrite(w, Picture, Region, XOrigin, YOrigin)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SetPictureClipRegionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func SetPictureClipRegionChecked(conn *x.Conn, Picture render.Picture, Region Region, XOrigin int16, YOrigin int16) x.VoidCookie {
	w := x.NewWriter()
	SetPictureClipRegionRequestWrite(w, Picture, Region, XOrigin, YOrigin)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SetPictureClipRegionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const SetCursorNameOpcode = 23

func SetCursorNameRequestWrite(w *x.Writer, Cursor x.Cursor, Nbytes uint16, Name string) {
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

	// write wire field Cursor
	w.Write4b(uint32(Cursor))
	n += 4

	// write wire field Nbytes
	w.Write2b(Nbytes)
	n += 2

	// write wire field Pad0
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Name
	w.WriteString(Name)
	n += len(Name)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SetCursorName(conn *x.Conn, Cursor x.Cursor, Nbytes uint16, Name string) x.VoidCookie {
	w := x.NewWriter()
	SetCursorNameRequestWrite(w, Cursor, Nbytes, Name)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SetCursorNameOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func SetCursorNameChecked(conn *x.Conn, Cursor x.Cursor, Nbytes uint16, Name string) x.VoidCookie {
	w := x.NewWriter()
	SetCursorNameRequestWrite(w, Cursor, Nbytes, Name)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SetCursorNameOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const GetCursorNameOpcode = 24

func GetCursorNameRequestWrite(w *x.Writer, Cursor x.Cursor) {
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

	// write wire field Cursor
	w.Write4b(uint32(Cursor))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetCursorNameReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence uint16
	Length   uint32
	Atom     x.Atom
	Nbytes   uint16
	// Pad1 [18]uint8
	Name string
}

func GetCursorNameReplyRead(r *x.Reader, v *GetCursorNameReply) (n int) {
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

	// read field Atom
	v.Atom = x.Atom(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Nbytes
	v.Nbytes = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(18)
	n += 18
	if r.Err() != nil {
		return
	}

	// read field Name
	{
		dataLen := int(int(v.Nbytes))
		data := r.ReadString(dataLen)
		if r.Err() != nil {
			return
		}
		blockLen += dataLen
		n += blockLen
		v.Name = data
	}
	alignTo = int(unsafe.Alignof(byte(0)))

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

type GetCursorNameCookie uint64

func (cookie GetCursorNameCookie) Reply(conn *x.Conn) (*GetCursorNameReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewGenericError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetCursorNameReply
	GetCursorNameReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetCursorName(conn *x.Conn, Cursor x.Cursor) GetCursorNameCookie {
	w := x.NewWriter()
	GetCursorNameRequestWrite(w, Cursor)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: GetCursorNameOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return GetCursorNameCookie(seq)
}

func GetCursorNameUnchecked(conn *x.Conn, Cursor x.Cursor) GetCursorNameCookie {
	w := x.NewWriter()
	GetCursorNameRequestWrite(w, Cursor)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: GetCursorNameOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetCursorNameCookie(seq)
}

const GetCursorImageAndNameOpcode = 25

func GetCursorImageAndNameRequestWrite(w *x.Writer) {
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

type GetCursorImageAndNameReply struct {
	ResponseType uint8
	// Pad0 uint8
	Sequence     uint16
	Length       uint32
	X            int16
	Y            int16
	Width        uint16
	Height       uint16
	Xhot         uint16
	Yhot         uint16
	CursorSerial uint32
	CursorAtom   x.Atom
	Nbytes       uint16
	// Pad1 [2]uint8
	CursorImage []uint32
	Name        string
}

func GetCursorImageAndNameReplyRead(r *x.Reader, v *GetCursorImageAndNameReply) (n int) {
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

	// read field X
	v.X = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Y
	v.Y = int16(r.Read2b())
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Width
	v.Width = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Height
	v.Height = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Xhot
	v.Xhot = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Yhot
	v.Yhot = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field CursorSerial
	v.CursorSerial = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field CursorAtom
	v.CursorAtom = x.Atom(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Nbytes
	v.Nbytes = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Pad1
	r.ReadPad(2)
	n += 2
	if r.Err() != nil {
		return
	}

	// read field CursorImage
	{
		dataLen := int((int(v.Width) * int(v.Height)))
		data := make([]uint32, dataLen)
		for i := 0; i < dataLen; i++ {
			data[i] = uint32(r.Read4b())
			if r.Err() != nil {
				return
			}
		}
		blockLen += dataLen * 4
		n += blockLen
		v.CursorImage = data
	}
	alignTo = int(unsafe.Alignof(uint32(0)))

	// read field Name

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
	{
		dataLen := int(int(v.Nbytes))
		data := r.ReadString(dataLen)
		if r.Err() != nil {
			return
		}
		blockLen += dataLen
		n += blockLen
		v.Name = data
	}
	alignTo = int(unsafe.Alignof(byte(0)))

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

type GetCursorImageAndNameCookie uint64

func (cookie GetCursorImageAndNameCookie) Reply(conn *x.Conn) (*GetCursorImageAndNameReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewGenericError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetCursorImageAndNameReply
	GetCursorImageAndNameReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetCursorImageAndName(conn *x.Conn) GetCursorImageAndNameCookie {
	w := x.NewWriter()
	GetCursorImageAndNameRequestWrite(w)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: GetCursorImageAndNameOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return GetCursorImageAndNameCookie(seq)
}

func GetCursorImageAndNameUnchecked(conn *x.Conn) GetCursorImageAndNameCookie {
	w := x.NewWriter()
	GetCursorImageAndNameRequestWrite(w)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: GetCursorImageAndNameOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetCursorImageAndNameCookie(seq)
}

const ChangeCursorOpcode = 26

func ChangeCursorRequestWrite(w *x.Writer, Source x.Cursor, Destination x.Cursor) {
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

	// write wire field Source
	w.Write4b(uint32(Source))
	n += 4

	// write wire field Destination
	w.Write4b(uint32(Destination))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func ChangeCursor(conn *x.Conn, Source x.Cursor, Destination x.Cursor) x.VoidCookie {
	w := x.NewWriter()
	ChangeCursorRequestWrite(w, Source, Destination)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: ChangeCursorOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func ChangeCursorChecked(conn *x.Conn, Source x.Cursor, Destination x.Cursor) x.VoidCookie {
	w := x.NewWriter()
	ChangeCursorRequestWrite(w, Source, Destination)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: ChangeCursorOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const ChangeCursorByNameOpcode = 27

func ChangeCursorByNameRequestWrite(w *x.Writer, Src x.Cursor, Nbytes uint16, Name string) {
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

	// write wire field Src
	w.Write4b(uint32(Src))
	n += 4

	// write wire field Nbytes
	w.Write2b(Nbytes)
	n += 2

	// write wire field Pad0
	w.WritePad(2)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Name
	w.WriteString(Name)
	n += len(Name)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func ChangeCursorByName(conn *x.Conn, Src x.Cursor, Nbytes uint16, Name string) x.VoidCookie {
	w := x.NewWriter()
	ChangeCursorByNameRequestWrite(w, Src, Nbytes, Name)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: ChangeCursorByNameOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func ChangeCursorByNameChecked(conn *x.Conn, Src x.Cursor, Nbytes uint16, Name string) x.VoidCookie {
	w := x.NewWriter()
	ChangeCursorByNameRequestWrite(w, Src, Nbytes, Name)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: ChangeCursorByNameOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const ExpandRegionOpcode = 28

func ExpandRegionRequestWrite(w *x.Writer, Source Region, Destination Region, Left uint16, Right uint16, Top uint16, Bottom uint16) {
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

	// write wire field Source
	w.Write4b(uint32(Source))
	n += 4

	// write wire field Destination
	w.Write4b(uint32(Destination))
	n += 4

	// write wire field Left
	w.Write2b(Left)
	n += 2

	// write wire field Right
	w.Write2b(Right)
	n += 2

	// write wire field Top
	w.Write2b(Top)
	n += 2

	// write wire field Bottom
	w.Write2b(Bottom)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func ExpandRegion(conn *x.Conn, Source Region, Destination Region, Left uint16, Right uint16, Top uint16, Bottom uint16) x.VoidCookie {
	w := x.NewWriter()
	ExpandRegionRequestWrite(w, Source, Destination, Left, Right, Top, Bottom)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: ExpandRegionOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func ExpandRegionChecked(conn *x.Conn, Source Region, Destination Region, Left uint16, Right uint16, Top uint16, Bottom uint16) x.VoidCookie {
	w := x.NewWriter()
	ExpandRegionRequestWrite(w, Source, Destination, Left, Right, Top, Bottom)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: ExpandRegionOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const HideCursorOpcode = 29

func HideCursorRequestWrite(w *x.Writer, Window x.Window) {
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

func HideCursor(conn *x.Conn, Window x.Window) x.VoidCookie {
	w := x.NewWriter()
	HideCursorRequestWrite(w, Window)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: HideCursorOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func HideCursorChecked(conn *x.Conn, Window x.Window) x.VoidCookie {
	w := x.NewWriter()
	HideCursorRequestWrite(w, Window)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: HideCursorOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const ShowCursorOpcode = 30

func ShowCursorRequestWrite(w *x.Writer, Window x.Window) {
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

func ShowCursor(conn *x.Conn, Window x.Window) x.VoidCookie {
	w := x.NewWriter()
	ShowCursorRequestWrite(w, Window)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: ShowCursorOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func ShowCursorChecked(conn *x.Conn, Window x.Window) x.VoidCookie {
	w := x.NewWriter()
	ShowCursorRequestWrite(w, Window)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: ShowCursorOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

// simple ('xcb', 'XFixes', 'BARRIER')
type Barrier uint32

// enum BarrierDirections
const (
	BarrierDirectionsPositiveX = 1
	BarrierDirectionsPositiveY = 2
	BarrierDirectionsNegativeX = 4
	BarrierDirectionsNegativeY = 8
)

const CreatePointerBarrierOpcode = 31

func CreatePointerBarrierRequestWrite(w *x.Writer, Barrier Barrier, Window x.Window, X1 uint16, Y1 uint16, X2 uint16, Y2 uint16, Directions uint32, NumDevices uint16, Devices []uint16) {
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

	// write wire field Barrier
	w.Write4b(uint32(Barrier))
	n += 4

	// write wire field Window
	w.Write4b(uint32(Window))
	n += 4

	// write wire field X1
	w.Write2b(X1)
	n += 2

	// write wire field Y1
	w.Write2b(Y1)
	n += 2

	// write wire field X2
	w.Write2b(X2)
	n += 2

	// write wire field Y2
	w.Write2b(Y2)
	n += 2

	// write wire field Directions
	w.Write4b(Directions)
	n += 4

	// write wire field Pad0
	w.WritePad(2)
	n += 2

	// write wire field NumDevices
	w.Write2b(NumDevices)
	n += 2
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Devices
	{
		_dataLen := int(NumDevices)
		for i := 0; i < _dataLen; i++ {
			w.Write2b(uint16(Devices[i]))
		}
		n += _dataLen * 2
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CreatePointerBarrier(conn *x.Conn, Barrier Barrier, Window x.Window, X1 uint16, Y1 uint16, X2 uint16, Y2 uint16, Directions uint32, NumDevices uint16, Devices []uint16) x.VoidCookie {
	w := x.NewWriter()
	CreatePointerBarrierRequestWrite(w, Barrier, Window, X1, Y1, X2, Y2, Directions, NumDevices, Devices)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreatePointerBarrierOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CreatePointerBarrierChecked(conn *x.Conn, Barrier Barrier, Window x.Window, X1 uint16, Y1 uint16, X2 uint16, Y2 uint16, Directions uint32, NumDevices uint16, Devices []uint16) x.VoidCookie {
	w := x.NewWriter()
	CreatePointerBarrierRequestWrite(w, Barrier, Window, X1, Y1, X2, Y2, Directions, NumDevices, Devices)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreatePointerBarrierOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const DeletePointerBarrierOpcode = 32

func DeletePointerBarrierRequestWrite(w *x.Writer, Barrier Barrier) {
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

	// write wire field Barrier
	w.Write4b(uint32(Barrier))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func DeletePointerBarrier(conn *x.Conn, Barrier Barrier) x.VoidCookie {
	w := x.NewWriter()
	DeletePointerBarrierRequestWrite(w, Barrier)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: DeletePointerBarrierOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func DeletePointerBarrierChecked(conn *x.Conn, Barrier Barrier) x.VoidCookie {
	w := x.NewWriter()
	DeletePointerBarrierRequestWrite(w, Barrier)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: DeletePointerBarrierOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}
