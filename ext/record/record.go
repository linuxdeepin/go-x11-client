package record

import (
	"fmt"

	"github.com/linuxdeepin/go-x11-client"
)

// size: 2b
type Range8 struct {
	First uint8
	Last  uint8
}

func writeRange8(b *x.FixedSizeBuf, v Range8) {
	b.Write1b(v.First)
	b.Write1b(v.Last)
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

// size: 4b
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

func writeRange16(b *x.FixedSizeBuf, v Range16) {
	b.Write2b(v.First)
	b.Write2b(v.Last)
}

// size: 6b
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

func writeExtRange(b *x.FixedSizeBuf, v ExtRange) {
	writeRange8(b, v.Major)
	writeRange16(b, v.Minor)
}

// size: 6 * 4b
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

func writeRange(b *x.FixedSizeBuf, v *Range) {
	writeRange8(b, v.CoreRequests)
	writeRange8(b, v.CoreReplies)
	writeExtRange(b, v.ExtRequests)
	writeExtRange(b, v.ExtReplies)
	writeRange8(b, v.DeliveredEvents)
	writeRange8(b, v.DeviceEvents)
	writeRange8(b, v.Errors)

	b.Write1b(x.BoolToUint8(v.ClientStarted))
	b.Write1b(x.BoolToUint8(v.ClientDied))
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
func encodeQueryVersion(majorVersion, minorVersion uint16) (b x.RequestBody) {
	b.AddBlock(1).
		Write2b(majorVersion).
		Write2b(minorVersion).
		End()
	return
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
func encodeCreateContext(context Context, elementHeader ElementHeader,
	clientSpecs []ClientSpec, ranges []Range) (b x.RequestBody) {

	clientSpecsLen := len(clientSpecs)
	rangesLen := len(ranges)
	b0 := b.AddBlock(4 + clientSpecsLen + rangesLen*6).
		Write4b(uint32(context)).
		Write1b(uint8(elementHeader)).
		WritePad(3).
		Write4b(uint32(clientSpecsLen)).
		Write4b(uint32(rangesLen))

	for _, clientSpec := range clientSpecs {
		b0.Write4b(uint32(clientSpec))
	}

	for i := 0; i < rangesLen; i++ {
		writeRange(b0, &ranges[i])
	}
	b0.End()
	return
}

// #WREQ
func encodeRegisterClients(context Context, elementHeader ElementHeader,
	clientSpecs []ClientSpec, ranges []Range) (b x.RequestBody) {

	clientSpecsLen := len(clientSpecs)
	rangesLen := len(ranges)

	b0 := b.AddBlock(4 + clientSpecsLen + rangesLen*6).
		Write4b(uint32(context)).
		Write1b(uint8(elementHeader)).
		WritePad(3).
		Write4b(uint32(clientSpecsLen)).
		Write4b(uint32(rangesLen))

	for _, clientSpec := range clientSpecs {
		b0.Write4b(uint32(clientSpec))
	}

	for i := 0; i < rangesLen; i++ {
		writeRange(b0, &ranges[i])
	}
	b0.End()
	return
}

// #WREQ
func encodeUnregisterClients(context Context, clientSpecs []ClientSpec) (b x.RequestBody) {
	clientSpecsLen := len(clientSpecs)
	b0 := b.AddBlock(2 + clientSpecsLen).
		Write4b(uint32(context)).
		Write4b(uint32(clientSpecsLen))

	for _, clientSpec := range clientSpecs {
		b0.Write4b(uint32(clientSpec))
	}
	b0.End()
	return
}

// #WREQ
func encodeGetContext(context Context) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(context)).
		End()
	return
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
func encodeEnableContext(context Context) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(context)).
		End()
	return
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
func encodeDisableContext(context Context) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(context)).
		End()
	return
}

// #WREQ
func encodeFreeContext(context Context) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(context)).
		End()
	return
}
