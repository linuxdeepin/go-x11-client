package record

import x "github.com/linuxdeepin/go-x11-client"

// module.direct_imports: []
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

// _ns.ext_name: Record
const MajorVersion = 1
const MinorVersion = 13

var _ext = x.NewExtension("RECORD")

func Ext() *x.Extension {
	return _ext
}

// simple ('xcb', 'Record', 'CONTEXT')
type Context uint32
type Range8 struct {
	First uint8
	Last  uint8
}

type CRange8 struct {
	First uint8
	Last  uint8
}

func Range8Read(r *x.Reader, v *Range8) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field First
	v.First = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Last
	v.Last = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}
	return
}

func Range8ReadN(r *x.Reader, dest []Range8) (n int) {
	for i := 0; i < len(dest); i++ {
		n += Range8Read(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func Range8Write(w *x.Writer, v *Range8) (n int) {

	// write field First
	w.Write1b(v.First)
	n += 1

	// write field Last
	w.Write1b(v.Last)
	n += 1
	return
}
func Range8WriteN(w *x.Writer, src []Range8) (n int) {
	for i := 0; i < len(src); i++ {
		n += Range8Write(w, &src[i])
	}
	return
}

type Range16 struct {
	First uint16
	Last  uint16
}

type CRange16 struct {
	First uint16
	Last  uint16
}

func Range16Read(r *x.Reader, v *Range16) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field First
	v.First = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}

	// read field Last
	v.Last = r.Read2b()
	n += 2
	if r.Err() != nil {
		return
	}
	return
}

func Range16ReadN(r *x.Reader, dest []Range16) (n int) {
	for i := 0; i < len(dest); i++ {
		n += Range16Read(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func Range16Write(w *x.Writer, v *Range16) (n int) {

	// write field First
	w.Write2b(v.First)
	n += 2

	// write field Last
	w.Write2b(v.Last)
	n += 2
	return
}
func Range16WriteN(w *x.Writer, src []Range16) (n int) {
	for i := 0; i < len(src); i++ {
		n += Range16Write(w, &src[i])
	}
	return
}

type ExtRange struct {
	Major Range8
	Minor Range16
}

type CExtRange struct {
	Major Range8
	Minor Range16
}

func ExtRangeRead(r *x.Reader, v *ExtRange) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field Major
	n += Range8Read(r, &v.Major)
	if r.Err() != nil {
		return
	}

	// read field Minor
	n += Range16Read(r, &v.Minor)
	if r.Err() != nil {
		return
	}
	return
}

func ExtRangeReadN(r *x.Reader, dest []ExtRange) (n int) {
	for i := 0; i < len(dest); i++ {
		n += ExtRangeRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func ExtRangeWrite(w *x.Writer, v *ExtRange) (n int) {

	// write field Major
	n += Range8Write(w, &v.Major)

	// write field Minor
	n += Range16Write(w, &v.Minor)
	return
}
func ExtRangeWriteN(w *x.Writer, src []ExtRange) (n int) {
	for i := 0; i < len(src); i++ {
		n += ExtRangeWrite(w, &src[i])
	}
	return
}

type Range struct {
	CoreRequests    Range8
	CoreReplies     Range8
	ExtRequests     ExtRange
	ExtReplies      ExtRange
	DeliveredEvents Range8
	DeviceEvents    Range8
	Errors          Range8
	ClientStarted   uint8
	ClientDied      uint8
}

type CRange struct {
	CoreRequests    Range8
	CoreReplies     Range8
	ExtRequests     ExtRange
	ExtReplies      ExtRange
	DeliveredEvents Range8
	DeviceEvents    Range8
	Errors          Range8
	ClientStarted   uint8
	ClientDied      uint8
}

func RangeRead(r *x.Reader, v *Range) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field CoreRequests
	n += Range8Read(r, &v.CoreRequests)
	if r.Err() != nil {
		return
	}

	// read field CoreReplies
	n += Range8Read(r, &v.CoreReplies)
	if r.Err() != nil {
		return
	}

	// read field ExtRequests
	n += ExtRangeRead(r, &v.ExtRequests)
	if r.Err() != nil {
		return
	}

	// read field ExtReplies
	n += ExtRangeRead(r, &v.ExtReplies)
	if r.Err() != nil {
		return
	}

	// read field DeliveredEvents
	n += Range8Read(r, &v.DeliveredEvents)
	if r.Err() != nil {
		return
	}

	// read field DeviceEvents
	n += Range8Read(r, &v.DeviceEvents)
	if r.Err() != nil {
		return
	}

	// read field Errors
	n += Range8Read(r, &v.Errors)
	if r.Err() != nil {
		return
	}

	// read field ClientStarted
	v.ClientStarted = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field ClientDied
	v.ClientDied = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}
	return
}

func RangeReadN(r *x.Reader, dest []Range) (n int) {
	for i := 0; i < len(dest); i++ {
		n += RangeRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func RangeWrite(w *x.Writer, v *Range) (n int) {

	// write field CoreRequests
	n += Range8Write(w, &v.CoreRequests)

	// write field CoreReplies
	n += Range8Write(w, &v.CoreReplies)

	// write field ExtRequests
	n += ExtRangeWrite(w, &v.ExtRequests)

	// write field ExtReplies
	n += ExtRangeWrite(w, &v.ExtReplies)

	// write field DeliveredEvents
	n += Range8Write(w, &v.DeliveredEvents)

	// write field DeviceEvents
	n += Range8Write(w, &v.DeviceEvents)

	// write field Errors
	n += Range8Write(w, &v.Errors)

	// write field ClientStarted
	w.Write1b(v.ClientStarted)
	n += 1

	// write field ClientDied
	w.Write1b(v.ClientDied)
	n += 1
	return
}
func RangeWriteN(w *x.Writer, src []Range) (n int) {
	for i := 0; i < len(src); i++ {
		n += RangeWrite(w, &src[i])
	}
	return
}

// simple ('xcb', 'Record', 'ElementHeader')
type ElementHeader uint8

// enum HType
const (
	HTypeFromServerTime     = 1
	HTypeFromClientTime     = 2
	HTypeFromClientSequence = 4
)

// simple ('xcb', 'Record', 'ClientSpec')
type ClientSpec uint32

// enum CS
const (
	CSCurrentClients = 1
	CSFutureClients  = 2
	CSAllClients     = 3
)

type ClientInfo struct {
	ClientResource ClientSpec
	NumRanges      uint32
	Ranges         []Range
}

type CClientInfo struct {
	ClientResource ClientSpec
	NumRanges      uint32
}

func ClientInfoRead(r *x.Reader, v *ClientInfo) (n int) {
	var pad int
	var alignTo int
	var blockLen int
	_ = pad
	_ = alignTo
	_ = blockLen

	// read field ClientResource
	v.ClientResource = ClientSpec(r.Read4b())
	n += 4
	if r.Err() != nil {
		return
	}

	// read field NumRanges
	v.NumRanges = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field Ranges
	v.Ranges = make([]Range, int(v.NumRanges))
	blockLen += RangeReadN(r, v.Ranges)
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CRange{}))

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

func ClientInfoReadN(r *x.Reader, dest []ClientInfo) (n int) {
	for i := 0; i < len(dest); i++ {
		n += ClientInfoRead(r, &dest[i])
		if r.Err() != nil {
			return
		}
	}
	return
}
func ClientInfoWrite(w *x.Writer, v *ClientInfo) (n int) {

	// write field ClientResource
	w.Write4b(uint32(v.ClientResource))
	n += 4

	// write field NumRanges
	w.Write4b(v.NumRanges)
	n += 4

	// write field Ranges
	return
}
func ClientInfoWriteN(w *x.Writer, src []ClientInfo) (n int) {
	for i := 0; i < len(src); i++ {
		n += ClientInfoWrite(w, &src[i])
	}
	return
}

const QueryVersionOpcode = 0

func QueryVersionRequestWrite(w *x.Writer, MajorVersion uint16, MinorVersion uint16) {
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
	w.Write2b(MajorVersion)
	n += 2

	// write wire field MinorVersion
	w.Write2b(MinorVersion)
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

func QueryVersion(conn *x.Conn, MajorVersion uint16, MinorVersion uint16) QueryVersionCookie {
	w := x.NewWriter()
	QueryVersionRequestWrite(w, MajorVersion, MinorVersion)
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

func QueryVersionUnchecked(conn *x.Conn, MajorVersion uint16, MinorVersion uint16) QueryVersionCookie {
	w := x.NewWriter()
	QueryVersionRequestWrite(w, MajorVersion, MinorVersion)
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

const CreateContextOpcode = 1

func CreateContextRequestWrite(w *x.Writer, Context Context, ElementHeader ElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []ClientSpec, Ranges []Range) {
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

	// write wire field Context
	w.Write4b(uint32(Context))
	n += 4

	// write wire field ElementHeader
	w.Write1b(uint8(ElementHeader))
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3

	// write wire field NumClientSpecs
	w.Write4b(NumClientSpecs)
	n += 4

	// write wire field NumRanges
	w.Write4b(NumRanges)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field ClientSpecs
	{
		_dataLen := int(NumClientSpecs)
		for i := 0; i < _dataLen; i++ {
			w.Write4b(uint32(ClientSpecs[i]))
		}
		n += _dataLen * 4
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Ranges
	n += RangeWriteN(w, Ranges)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func CreateContext(conn *x.Conn, Context Context, ElementHeader ElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []ClientSpec, Ranges []Range) x.VoidCookie {
	w := x.NewWriter()
	CreateContextRequestWrite(w, Context, ElementHeader, NumClientSpecs, NumRanges, ClientSpecs, Ranges)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateContextOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func CreateContextChecked(conn *x.Conn, Context Context, ElementHeader ElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []ClientSpec, Ranges []Range) x.VoidCookie {
	w := x.NewWriter()
	CreateContextRequestWrite(w, Context, ElementHeader, NumClientSpecs, NumRanges, ClientSpecs, Ranges)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: CreateContextOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const RegisterClientsOpcode = 2

func RegisterClientsRequestWrite(w *x.Writer, Context Context, ElementHeader ElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []ClientSpec, Ranges []Range) {
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

	// write wire field Context
	w.Write4b(uint32(Context))
	n += 4

	// write wire field ElementHeader
	w.Write1b(uint8(ElementHeader))
	n += 1

	// write wire field Pad0
	w.WritePad(3)
	n += 3

	// write wire field NumClientSpecs
	w.Write4b(NumClientSpecs)
	n += 4

	// write wire field NumRanges
	w.Write4b(NumRanges)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field ClientSpecs
	{
		_dataLen := int(NumClientSpecs)
		for i := 0; i < _dataLen; i++ {
			w.Write4b(uint32(ClientSpecs[i]))
		}
		n += _dataLen * 4
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field Ranges
	n += RangeWriteN(w, Ranges)
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func RegisterClients(conn *x.Conn, Context Context, ElementHeader ElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []ClientSpec, Ranges []Range) x.VoidCookie {
	w := x.NewWriter()
	RegisterClientsRequestWrite(w, Context, ElementHeader, NumClientSpecs, NumRanges, ClientSpecs, Ranges)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: RegisterClientsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func RegisterClientsChecked(conn *x.Conn, Context Context, ElementHeader ElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []ClientSpec, Ranges []Range) x.VoidCookie {
	w := x.NewWriter()
	RegisterClientsRequestWrite(w, Context, ElementHeader, NumClientSpecs, NumRanges, ClientSpecs, Ranges)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: RegisterClientsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const UnregisterClientsOpcode = 3

func UnregisterClientsRequestWrite(w *x.Writer, Context Context, NumClientSpecs uint32, ClientSpecs []ClientSpec) {
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

	// write wire field Context
	w.Write4b(uint32(Context))
	n += 4

	// write wire field NumClientSpecs
	w.Write4b(NumClientSpecs)
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0

	// write param field ClientSpecs
	{
		_dataLen := int(NumClientSpecs)
		for i := 0; i < _dataLen; i++ {
			w.Write4b(uint32(ClientSpecs[i]))
		}
		n += _dataLen * 4
	}
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func UnregisterClients(conn *x.Conn, Context Context, NumClientSpecs uint32, ClientSpecs []ClientSpec) x.VoidCookie {
	w := x.NewWriter()
	UnregisterClientsRequestWrite(w, Context, NumClientSpecs, ClientSpecs)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: UnregisterClientsOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func UnregisterClientsChecked(conn *x.Conn, Context Context, NumClientSpecs uint32, ClientSpecs []ClientSpec) x.VoidCookie {
	w := x.NewWriter()
	UnregisterClientsRequestWrite(w, Context, NumClientSpecs, ClientSpecs)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: UnregisterClientsOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const GetContextOpcode = 4

func GetContextRequestWrite(w *x.Writer, Context Context) {
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

	// write wire field Context
	w.Write4b(uint32(Context))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type GetContextReply struct {
	ResponseType  uint8
	Enabled       uint8
	Sequence      uint16
	Length        uint32
	ElementHeader ElementHeader
	// Pad0 [3]uint8
	NumInterceptedClients uint32
	// Pad1 [16]uint8
	InterceptedClients []ClientInfo
}

func GetContextReplyRead(r *x.Reader, v *GetContextReply) (n int) {
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

	// read field ElementHeader
	v.ElementHeader = ElementHeader(r.Read1b())
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(3)
	n += 3
	if r.Err() != nil {
		return
	}

	// read field NumInterceptedClients
	v.NumInterceptedClients = r.Read4b()
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

	// read field InterceptedClients
	v.InterceptedClients = make([]ClientInfo, int(v.NumInterceptedClients))
	blockLen += ClientInfoReadN(r, v.InterceptedClients)
	n += blockLen
	if r.Err() != nil {
		return
	}
	alignTo = int(unsafe.Alignof(CClientInfo{}))

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

type GetContextCookie uint64

func (cookie GetContextCookie) Reply(conn *x.Conn) (*GetContextReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewGenericError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply GetContextReply
	GetContextReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func GetContext(conn *x.Conn, Context Context) GetContextCookie {
	w := x.NewWriter()
	GetContextRequestWrite(w, Context)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: GetContextOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return GetContextCookie(seq)
}

func GetContextUnchecked(conn *x.Conn, Context Context) GetContextCookie {
	w := x.NewWriter()
	GetContextRequestWrite(w, Context)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: GetContextOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return GetContextCookie(seq)
}

const EnableContextOpcode = 5

func EnableContextRequestWrite(w *x.Writer, Context Context) {
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

	// write wire field Context
	w.Write4b(uint32(Context))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

type EnableContextReply struct {
	ResponseType  uint8
	Category      uint8
	Sequence      uint16
	Length        uint32
	ElementHeader ElementHeader
	ClientSwapped uint8
	// Pad0 [2]uint8
	XidBase        uint32
	ServerTime     uint32
	RecSequenceNum uint32
	// Pad1 [8]uint8
	Data []uint8
}

func EnableContextReplyRead(r *x.Reader, v *EnableContextReply) (n int) {
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

	// read field Category
	v.Category = r.Read1b()
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

	// read field ElementHeader
	v.ElementHeader = ElementHeader(r.Read1b())
	n += 1
	if r.Err() != nil {
		return
	}

	// read field ClientSwapped
	v.ClientSwapped = r.Read1b()
	n += 1
	if r.Err() != nil {
		return
	}

	// read field Pad0
	r.ReadPad(2)
	n += 2
	if r.Err() != nil {
		return
	}

	// read field XidBase
	v.XidBase = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field ServerTime
	v.ServerTime = r.Read4b()
	n += 4
	if r.Err() != nil {
		return
	}

	// read field RecSequenceNum
	v.RecSequenceNum = r.Read4b()
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

	// read field Data
	{
		dataLen := int((int(v.Length) * 4))
		data := make([]uint8, dataLen)
		for i := 0; i < dataLen; i++ {
			data[i] = uint8(r.Read1b())
			if r.Err() != nil {
				return
			}
		}
		blockLen += dataLen
		n += blockLen
		v.Data = data
	}
	alignTo = int(unsafe.Alignof(uint8(0)))

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

type EnableContextCookie uint64

func (cookie EnableContextCookie) Reply(conn *x.Conn) (*EnableContextReply, error) {
	replyBuf, isErr := conn.WaitForReply(uint64(cookie))
	if isErr {
		return nil, x.NewGenericError(replyBuf)
	}
	r := x.NewReaderFromData(replyBuf)
	var reply EnableContextReply
	EnableContextReplyRead(r, &reply)
	if err := r.Err(); err != nil {
		return nil, err
	}
	return &reply, nil
}

func EnableContext(conn *x.Conn, Context Context) EnableContextCookie {
	w := x.NewWriter()
	EnableContextRequestWrite(w, Context)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: EnableContextOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return EnableContextCookie(seq)
}

func EnableContextUnchecked(conn *x.Conn, Context Context) EnableContextCookie {
	w := x.NewWriter()
	EnableContextRequestWrite(w, Context)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: false,
		Opcode: EnableContextOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return EnableContextCookie(seq)
}

const DisableContextOpcode = 6

func DisableContextRequestWrite(w *x.Writer, Context Context) {
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

	// write wire field Context
	w.Write4b(uint32(Context))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func DisableContext(conn *x.Conn, Context Context) x.VoidCookie {
	w := x.NewWriter()
	DisableContextRequestWrite(w, Context)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: DisableContextOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func DisableContextChecked(conn *x.Conn, Context Context) x.VoidCookie {
	w := x.NewWriter()
	DisableContextRequestWrite(w, Context)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: DisableContextOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}

const FreeContextOpcode = 7

func FreeContextRequestWrite(w *x.Writer, Context Context) {
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

	// write wire field Context
	w.Write4b(uint32(Context))
	n += 4
	/* end fixed size part */
	pad = (-n) & 3
	w.WritePad(pad)
	n = 0
}

func FreeContext(conn *x.Conn, Context Context) x.VoidCookie {
	w := x.NewWriter()
	FreeContextRequestWrite(w, Context)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: FreeContextOpcode,
	}
	// no check: True
	seq := conn.SendRequest(0, data, req)
	return x.VoidCookie(seq)
}

func FreeContextChecked(conn *x.Conn, Context Context) x.VoidCookie {
	w := x.NewWriter()
	FreeContextRequestWrite(w, Context)
	data := w.Bytes()
	req := &x.ProtocolRequest{
		Ext:    _ext,
		IsVoid: true,
		Opcode: FreeContextOpcode,
	}
	// no check: False
	seq := conn.SendRequest(x.RequestChecked, data, req)
	return x.VoidCookie(seq)
}
