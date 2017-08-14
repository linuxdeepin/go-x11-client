package screensaver

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

// _ns.ext_name: ScreenSaver
const MajorVersion = 1
const MinorVersion = 1

var _ext = x.NewExtension("MIT-SCREEN-SAVER")

func Ext() *x.Extension {
	return _ext
}

// enum Kind
const (
	KindBlanked  = 0
	KindInternal = 1
	KindExternal = 2
)

func KindEnumToStr(v int) string {
	switch v {
	case 0:
		return "Blanked"
	case 1:
		return "Internal"
	case 2:
		return "External"
	default:
		return "<Unknown>"
	}
}

// enum Event
const (
	EventNotifyMask = 1
	EventCycleMask  = 2
)

// enum State
const (
	StateOff      = 0
	StateOn       = 1
	StateCycle    = 2
	StateDisabled = 3
)

func StateEnumToStr(v int) string {
	switch v {
	case 0:
		return "Off"
	case 1:
		return "On"
	case 2:
		return "Cycle"
	case 3:
		return "Disabled"
	default:
		return "<Unknown>"
	}
}

const QueryVersionOpcode = 0

func QueryVersionRequestWrite(w *x.Writer, ClientMajorVersion uint8, ClientMinorVersion uint8) {
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
	w.Write1b(ClientMajorVersion)
	n += 1

	// write wire field ClientMinorVersion
	w.Write1b(ClientMinorVersion)
	n += 1

	// write wire field Pad0
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
	Sequence           uint16
	Length             uint32
	ServerMajorVersion uint16
	ServerMinorVersion uint16
	// Pad1 [20]uint8
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

	// read field ServerMajorVersion
	v.ServerMajorVersion = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field ServerMinorVersion
	v.ServerMinorVersion = r.Read2b()
	n += 2
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

func QueryVersion(conn *x.Conn, ClientMajorVersion uint8, ClientMinorVersion uint8) QueryVersionCookie {
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

func QueryVersionUnchecked(conn *x.Conn, ClientMajorVersion uint8, ClientMinorVersion uint8) QueryVersionCookie {
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

const QueryInfoOpcode = 1

func QueryInfoRequestWrite(w *x.Writer, Drawable x.Drawable) {
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
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type QueryInfoReply struct {
	ResponseType     uint8
	State            uint8
	Sequence         uint16
	Length           uint32
	SaverWindow      x.Window
	MsUntilServer    uint32
	MsSinceUserInput uint32
	EventMask        uint32
	Kind             uint8
	// Pad0 [7]uint8
}

func QueryInfoReplyRead(r *x.Reader, v *QueryInfoReply) (n int) {
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

	// read field State
	v.State = r.Read1b()
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

	// read field SaverWindow
	v.SaverWindow = x.Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field MsUntilServer
	v.MsUntilServer = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field MsSinceUserInput
	v.MsSinceUserInput = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field EventMask
	v.EventMask = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Kind
	v.Kind = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(7)
	n += 7
	if r.Err() != nil {
		return
	}
	return
}

type QueryInfoCookie uint64

func (cookie QueryInfoCookie) Reply(conn *x.Conn) (*QueryInfoReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewGenericError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply QueryInfoReply
	QueryInfoReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func QueryInfo(conn *x.Conn, Drawable x.Drawable) QueryInfoCookie {
	w := x.NewWriter()
	QueryInfoRequestWrite(w, Drawable)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: QueryInfoOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return QueryInfoCookie(seq)
}

func QueryInfoUnchecked(conn *x.Conn, Drawable x.Drawable) QueryInfoCookie {
	w := x.NewWriter()
	QueryInfoRequestWrite(w, Drawable)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: QueryInfoOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return QueryInfoCookie(seq)
}

const SelectInputOpcode = 2

func SelectInputRequestWrite(w *x.Writer, Drawable x.Drawable, EventMask uint32) {
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

	// write wire field EventMask
	w.Write4b(EventMask)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func SelectInput(conn *x.Conn, Drawable x.Drawable, EventMask uint32) x.VoidCookie {
	w := x.NewWriter()
	SelectInputRequestWrite(w, Drawable, EventMask)
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

func SelectInputChecked(conn *x.Conn, Drawable x.Drawable, EventMask uint32) x.VoidCookie {
	w := x.NewWriter()
	SelectInputRequestWrite(w, Drawable, EventMask)
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

type SetAttributesValueList struct {
	BackgroundPixmap   x.Pixmap
	BackgroundPixel    uint32
	BorderPixmap       x.Pixmap
	BorderPixel        uint32
	BitGravity         uint32
	WinGravity         uint32
	BackingStore       uint32
	BackingPlanes      uint32
	BackingPixel       uint32
	OverrideRedirect   x.Bool32
	SaveUnder          x.Bool32
	EventMask          uint32
	DoNotPropogateMask uint32
	Colormap           x.Colormap
	Cursor             x.Cursor
}

// _go_switch_write SwitchType "xcb.ScreenSaver.SetAttributes.value_list"
func SetAttributesValueListSerialize(w *x.Writer, ValueMask uint32, _aux *SetAttributesValueList) {
	n := 0
	xgbPaddingOffset := 0
	alignTo := 0
	xgbPad := 0
	// switch expr: int(ValueMask)

	// case_field Field with type BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase1"
	if (int(ValueMask) & x.CWBackPixmap) != 0 {
		// todo
		// field Field "background_pixmap" in BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase1"
		w.Write4b(uint32(_aux.BackgroundPixmap))
		n += 4
		alignTo = int(unsafe.Alignof(x.Pixmap(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase2"
	if (int(ValueMask) & x.CWBackPixel) != 0 {
		// todo
		// field Field "background_pixel" in BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase2"
		w.Write4b(_aux.BackgroundPixel)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase3"
	if (int(ValueMask) & x.CWBorderPixmap) != 0 {
		// todo
		// field Field "border_pixmap" in BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase3"
		w.Write4b(uint32(_aux.BorderPixmap))
		n += 4
		alignTo = int(unsafe.Alignof(x.Pixmap(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase4"
	if (int(ValueMask) & x.CWBorderPixel) != 0 {
		// todo
		// field Field "border_pixel" in BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase4"
		w.Write4b(_aux.BorderPixel)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase5"
	if (int(ValueMask) & x.CWBitGravity) != 0 {
		// todo
		// field Field "bit_gravity" in BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase5"
		w.Write4b(_aux.BitGravity)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase6"
	if (int(ValueMask) & x.CWWinGravity) != 0 {
		// todo
		// field Field "win_gravity" in BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase6"
		w.Write4b(_aux.WinGravity)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase7"
	if (int(ValueMask) & x.CWBackingStore) != 0 {
		// todo
		// field Field "backing_store" in BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase7"
		w.Write4b(_aux.BackingStore)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase8"
	if (int(ValueMask) & x.CWBackingPlanes) != 0 {
		// todo
		// field Field "backing_planes" in BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase8"
		w.Write4b(_aux.BackingPlanes)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase9"
	if (int(ValueMask) & x.CWBackingPixel) != 0 {
		// todo
		// field Field "backing_pixel" in BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase9"
		w.Write4b(_aux.BackingPixel)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase10"
	if (int(ValueMask) & x.CWOverrideRedirect) != 0 {
		// todo
		// field Field "override_redirect" in BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase10"
		w.Write4b(uint32(_aux.OverrideRedirect))
		n += 4
		alignTo = int(unsafe.Alignof(x.Bool32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase11"
	if (int(ValueMask) & x.CWSaveUnder) != 0 {
		// todo
		// field Field "save_under" in BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase11"
		w.Write4b(uint32(_aux.SaveUnder))
		n += 4
		alignTo = int(unsafe.Alignof(x.Bool32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase12"
	if (int(ValueMask) & x.CWEventMask) != 0 {
		// todo
		// field Field "event_mask" in BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase12"
		w.Write4b(_aux.EventMask)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase13"
	if (int(ValueMask) & x.CWDontPropagate) != 0 {
		// todo
		// field Field "do_not_propogate_mask" in BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase13"
		w.Write4b(_aux.DoNotPropogateMask)
		n += 4
		alignTo = int(unsafe.Alignof(uint32(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase14"
	if (int(ValueMask) & x.CWColormap) != 0 {
		// todo
		// field Field "colormap" in BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase14"
		w.Write4b(uint32(_aux.Colormap))
		n += 4
		alignTo = int(unsafe.Alignof(x.Colormap(0)))
	}
	// end a case_field

	// case_field Field with type BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase15"
	if (int(ValueMask) & x.CWCursor) != 0 {
		// todo
		// field Field "cursor" in BitcaseType "xcb.ScreenSaver.SetAttributes.value_list.bitcase15"
		w.Write4b(uint32(_aux.Cursor))
		n += 4
		alignTo = int(unsafe.Alignof(x.Cursor(0)))
	}
	// end a case_field

	/* insert padding */
	xgbPad = -(n + xgbPaddingOffset) & (alignTo - 1)
	w.WritePad(xgbPad)
	xgbPad = 0
	n = 0
	xgbPaddingOffset = 0
}

const SetAttributesOpcode = 3

func SetAttributesRequestWrite(w *x.Writer, Drawable x.Drawable, X int16, Y int16, Width uint16, Height uint16, BorderWidth uint16, Class uint8, Depth uint8, Visual x.VisualID, ValueMask uint32, ValueList *SetAttributesValueList) {
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

	// write wire field X
	w.Write2b(uint16(X))
	n += 2

	// write wire field Y
	w.Write2b(uint16(Y))
	n += 2

	// write wire field Width
	w.Write2b(Width)
	n += 2

	// write wire field Height
	w.Write2b(Height)
	n += 2

	// write wire field BorderWidth
	w.Write2b(BorderWidth)
	n += 2

	// write wire field Class
	w.Write1b(Class)
	n += 1

	// write wire field Depth
	w.Write1b(Depth)
	n += 1

	// write wire field Visual
	w.Write4b(uint32(Visual))
	n += 4

	// write wire field ValueMask
	w.Write4b(ValueMask)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field ValueList
	// switch serialize
	SetAttributesValueListSerialize(w, ValueMask, ValueList)
}

func SetAttributes(conn *x.Conn, Drawable x.Drawable, X int16, Y int16, Width uint16, Height uint16, BorderWidth uint16, Class uint8, Depth uint8, Visual x.VisualID, ValueMask uint32, ValueList *SetAttributesValueList) x.VoidCookie {
	w := x.NewWriter()
	SetAttributesRequestWrite(w, Drawable, X, Y, Width, Height, BorderWidth, Class, Depth, Visual, ValueMask, ValueList)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SetAttributesOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func SetAttributesChecked(conn *x.Conn, Drawable x.Drawable, X int16, Y int16, Width uint16, Height uint16, BorderWidth uint16, Class uint8, Depth uint8, Visual x.VisualID, ValueMask uint32, ValueList *SetAttributesValueList) x.VoidCookie {
	w := x.NewWriter()
	SetAttributesRequestWrite(w, Drawable, X, Y, Width, Height, BorderWidth, Class, Depth, Visual, ValueMask, ValueList)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SetAttributesOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const UnsetAttributesOpcode = 4

func UnsetAttributesRequestWrite(w *x.Writer, Drawable x.Drawable) {
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
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func UnsetAttributes(conn *x.Conn, Drawable x.Drawable) x.VoidCookie {
	w := x.NewWriter()
	UnsetAttributesRequestWrite(w, Drawable)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: UnsetAttributesOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func UnsetAttributesChecked(conn *x.Conn, Drawable x.Drawable) x.VoidCookie {
	w := x.NewWriter()
	UnsetAttributesRequestWrite(w, Drawable)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: UnsetAttributesOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const SuspendOpcode = 5

func SuspendRequestWrite(w *x.Writer, Suspend uint8) {
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

	// write wire field Suspend
	w.Write1b(Suspend)
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func Suspend(conn *x.Conn, Suspend uint8) x.VoidCookie {
	w := x.NewWriter()
	SuspendRequestWrite(w, Suspend)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SuspendOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func SuspendChecked(conn *x.Conn, Suspend uint8) x.VoidCookie {
	w := x.NewWriter()
	SuspendRequestWrite(w, Suspend)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: SuspendOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

// event name: ('xcb', 'ScreenSaver', 'Notify')
// self.name: ('xcb', 'ScreenSaver', 'Notify')
const NotifyEventCode = 0

// not event copy
type NotifyEvent struct {
	ResponseType uint8
	State        uint8
	Sequence     uint16
	Time         x.Timestamp
	Root         x.Window
	Window       x.Window
	Kind         uint8
	Forced       uint8
	// Pad0 [14]uint8
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

	// read field State
	v.State = r.Read1b()
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

	// read field Time
	v.Time = x.Timestamp(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Root
	v.Root = x.Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Window
	v.Window = x.Window(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Kind
	v.Kind = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Forced
	v.Forced = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(14)
	n += 14
	if r.Err() != nil {
		return
	}
	return
}
func NotifyEventWrite(w *x.Writer, v *NotifyEvent) (n int) {

	// write field ResponseType
	w.Write1b(v.ResponseType)
	n += 1

	// write field State
	w.Write1b(v.State)
	n += 1

	// write field Sequence
	w.Write2b(v.Sequence)
	n += 2

	// write field Time
	w.Write4b(uint32(v.Time))
	n += 4

	// write field Root
	w.Write4b(uint32(v.Root))
	n += 4

	// write field Window
	w.Write4b(uint32(v.Window))
	n += 4

	// write field Kind
	w.Write1b(v.Kind)
	n += 1

	// write field Forced
	w.Write1b(v.Forced)
	n += 1

	// write field Pad0
	w.WritePad(14)
	n += 14
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
	return fmt.Sprintf("NotifyEvent {State: %v, Sequence: %v, Time: %v, Root: %v, Window: %v, Kind: %v, Forced: %v}",
		v.State, v.Sequence, v.Time, v.Root, v.Window, v.Kind, v.Forced)
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
