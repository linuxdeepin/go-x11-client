package record

import (
	"fmt"

	"github.com/linuxdeepin/go-x11-client"
)

type Range8 struct {
	First uint8
	Last  uint8
}

func writeRange8(w *x.Writer, v Range8) {
	w.Write1b(v.First)
	w.Write1b(v.Last)
}

func readRange8(r *x.Reader) (Range8, error) {
	var v Range8
	v.First = r.Read1b()
	if r.Err() != nil {
		return Range8{}, r.Err()
	}

	v.Last = r.Read1b()
	if r.Err() != nil {
		return Range8{}, r.Err()
	}
	return v, nil
}

type Range16 struct {
	First uint16
	Last  uint16
}

func readRange16(r *x.Reader) (Range16, error) {
	var v Range16
	v.First = r.Read2b()
	if r.Err() != nil {
		return Range16{}, r.Err()
	}

	v.Last = r.Read2b()
	if r.Err() != nil {
		return Range16{}, r.Err()
	}
	return v, nil
}

func writeRange16(w *x.Writer, v Range16) {
	w.Write2b(v.First)
	w.Write2b(v.Last)
}

type ExtRange struct {
	Major Range8
	Minor Range16
}

func readExtRange(r *x.Reader) (v ExtRange, err error) {
	v.Major, err = readRange8(r)
	if err != nil {
		return ExtRange{}, err
	}

	v.Minor, err = readRange16(r)
	if err != nil {
		return ExtRange{}, err
	}

	return v, nil
}

func writeExtRange(w *x.Writer, v ExtRange) {
	writeRange8(w, v.Major)
	writeRange16(w, v.Minor)
}

type Range struct {
	CoreRequests    Range8
	CoreReplies     Range8
	ExtRequests     ExtRange
	ExtReplies      ExtRange
	DeliveredEvents Range8
	DeviceEvents    Range8
	Errors          Range8
	ClientStarted   bool
	ClientDied      bool
}

func writeRange(w *x.Writer, v *Range) {
	writeRange8(w, v.CoreRequests)
	writeRange8(w, v.CoreReplies)
	writeExtRange(w, v.ExtRequests)
	writeExtRange(w, v.ExtReplies)
	writeRange8(w, v.DeliveredEvents)
	writeRange8(w, v.DeviceEvents)
	writeRange8(w, v.Errors)

	w.Write1b(x.BoolToUint8(v.ClientStarted))
	w.Write1b(x.BoolToUint8(v.ClientDied))
}

func readRange(r *x.Reader, v *Range) error {
	var err error
	v.CoreRequests, err = readRange8(r)
	if err != nil {
		return err
	}

	v.CoreReplies, err = readRange8(r)
	if err != nil {
		return err
	}

	v.ExtRequests, err = readExtRange(r)
	if err != nil {
		return err
	}

	v.ExtRequests, err = readExtRange(r)
	if err != nil {
		return err
	}

	v.DeliveredEvents, err = readRange8(r)
	if err != nil {
		return err
	}

	v.DeviceEvents, err = readRange8(r)
	if err != nil {
		return err
	}

	v.Errors, err = readRange8(r)
	if err != nil {
		return err
	}

	v.ClientStarted = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.ClientDied = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}
	return nil
}

type ClientInfo struct {
	ClientResource      ClientSpec
	InterceptedProtocol []Range
}

func readClientInfo(r *x.Reader, v *ClientInfo) error {
	v.ClientResource = ClientSpec(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	interceptedProtocolLen := int(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	if interceptedProtocolLen > 0 {
		v.InterceptedProtocol = make([]Range, interceptedProtocolLen)
		for i := 0; i < interceptedProtocolLen; i++ {
			err := readRange(r, &v.InterceptedProtocol[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type ContextError struct {
	Code     uint8
	Sequence uint16
	Bad      Context
}

func (e ContextError) GetCode() uint8 {
	return e.Code
}

func (e ContextError) GetSequenceNumber() uint16 {
	return e.Sequence
}

func (e ContextError) GetMinorOpcode() uint16 {
	return 0
}

func (e ContextError) GetMajorOpcode() uint8 {
	return 0
}

func (e ContextError) Error() string {
	return fmt.Sprintf("record.ContextError{code=%d, seq=%d, bad=%d}",
		e.Code, e.Sequence, e.Bad)
}

func readBadContextError(r *x.Reader) x.Error {
	var e ContextError
	// Error
	r.Read1b()
	e.Code = r.Read1b()
	e.Sequence = r.Read2b()

	e.Bad = Context(r.Read4b())

	r.ReadPad(24)
	return e
}

// #WREQ
func writeQueryVersion(w *x.Writer, majorVersion, minorVersion uint16) {
	w.WritePad(4)
	w.Write2b(majorVersion)
	w.Write2b(minorVersion)
}

type QueryVersionReply struct {
	MajorVersion uint16
	MinorVersion uint16
}

func readQueryVersionReply(r *x.Reader, v *QueryVersionReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MajorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MinorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeCreateContext(w *x.Writer, context Context, elementHeader ElementHeader,
	clientSpecs []ClientSpec, ranges []Range) {
	w.WritePad(4)
	w.Write4b(uint32(context))

	w.Write1b(uint8(elementHeader))
	w.WritePad(3)

	clientSpecsLen := len(clientSpecs)
	w.Write4b(uint32(clientSpecsLen))

	rangesLen := len(ranges)
	w.Write4b(uint32(rangesLen))

	for _, clientSpec := range clientSpecs {
		w.Write4b(uint32(clientSpec))
	}

	for i := 0; i < rangesLen; i++ {
		writeRange(w, &ranges[i])
	}
}

// #WREQ
func writeRegisterClients(w *x.Writer, context Context, elementHeader ElementHeader,
	clientSpecs []ClientSpec, ranges []Range) {
	w.WritePad(4)
	w.Write4b(uint32(context))

	w.Write1b(uint8(elementHeader))
	w.WritePad(3)

	clientSpecsLen := len(clientSpecs)
	w.Write4b(uint32(clientSpecsLen))

	rangesLen := len(ranges)
	w.Write4b(uint32(rangesLen))

	for _, clientSpec := range clientSpecs {
		w.Write4b(uint32(clientSpec))
	}

	for i := 0; i < rangesLen; i++ {
		writeRange(w, &ranges[i])
	}
}

// #WREQ
func writeUnregisterClients(w *x.Writer, context Context, clientSpecs []ClientSpec) {
	w.WritePad(4)
	w.Write4b(uint32(context))
	clientSpecsLen := len(clientSpecs)
	w.Write4b(uint32(clientSpecsLen))
	for _, clientSpec := range clientSpecs {
		w.Write4b(uint32(clientSpec))
	}
}

// #WREQ
func writeGetContext(w *x.Writer, context Context) {
	w.WritePad(4)
	w.Write4b(uint32(context))
}

type GetContextReply struct {
	Enabled            bool
	ElementHeader      ElementHeader
	InterceptedClients []ClientInfo
}

func readGetContextReply(r *x.Reader, v *GetContextReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Enabled = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// len
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ElementHeader = ElementHeader(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(3)
	if r.Err() != nil {
		return r.Err()
	}

	interceptedClientsLen := int(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	if interceptedClientsLen > 0 {
		v.InterceptedClients = make([]ClientInfo, interceptedClientsLen)
		for i := 0; i < interceptedClientsLen; i++ {
			err := readClientInfo(r, &v.InterceptedClients[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// #WREQ
func writeEnableContext(w *x.Writer, context Context) {
	w.WritePad(4)
	w.Write4b(uint32(context))
}

type EnableContextReply struct {
	Category       uint8
	ElementHeader  ElementHeader
	ClientSwapped  bool
	XidBase        uint32
	ServerTime     x.Timestamp
	RecSequenceNum uint32
	Data           []uint8
}

func readEnableContextReply(r *x.Reader, v *EnableContextReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Category = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// len
	replyLen := r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ElementHeader = ElementHeader(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.ClientSwapped = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(2)
	if r.Err() != nil {
		return r.Err()
	}

	v.XidBase = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ServerTime = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.RecSequenceNum = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(8)

	dataLen := 4 * int(replyLen)
	v.Data = r.ReadBytes(dataLen)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeDisableContext(w *x.Writer, context Context) {
	w.WritePad(4)
	w.Write4b(uint32(context))
}

// #WREQ
func writeFreeContext(w *x.Writer, context Context) {
	w.WritePad(4)
	w.Write4b(uint32(context))
}
