// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package record

import (
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

func readRange8(r *x.Reader) Range8 {
	var v Range8
	v.First = r.Read1b()
	v.Last = r.Read1b()
	return v
}

// size: 4b
type Range16 struct {
	First uint16
	Last  uint16
}

func readRange16(r *x.Reader) Range16 {
	var v Range16
	v.First = r.Read2b()
	v.Last = r.Read2b()
	return v
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

func readExtRange(r *x.Reader) (v ExtRange) {
	v.Major = readRange8(r)
	v.Minor = readRange16(r)
	return
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

	b.WriteBool(v.ClientStarted).
		WriteBool(v.ClientDied)
}

func readRange(r *x.Reader, v *Range) {
	v.CoreRequests = readRange8(r)
	v.CoreReplies = readRange8(r)

	v.ExtRequests = readExtRange(r)

	v.ExtRequests = readExtRange(r)

	v.DeliveredEvents = readRange8(r)
	v.DeviceEvents = readRange8(r)

	v.Errors = readRange8(r)
	v.ClientStarted = r.ReadBool()
	v.ClientDied = r.ReadBool()
}

type ClientInfo struct {
	ClientResource      ClientSpec
	InterceptedProtocol []Range
}

func readClientInfo(r *x.Reader, v *ClientInfo) error {
	if !r.RemainAtLeast4b(2) {
		return x.ErrDataLenShort
	}
	v.ClientResource = ClientSpec(r.Read4b())

	interceptedProtocolLen := int(r.Read4b()) // 2

	if interceptedProtocolLen > 0 {
		if !r.RemainAtLeast4b(6 * interceptedProtocolLen) {
			return x.ErrDataLenShort
		}
		v.InterceptedProtocol = make([]Range, interceptedProtocolLen)
		for i := 0; i < interceptedProtocolLen; i++ {
			readRange(r, &v.InterceptedProtocol[i])
		}
	}

	return nil
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
	if !r.RemainAtLeast4b(3) {
		return x.ErrDataLenShort
	}
	r.ReadPad(8)

	v.MajorVersion = r.Read2b()
	v.MinorVersion = r.Read2b() // 3

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
	if !r.RemainAtLeast4b(4) {
		return x.ErrDataLenShort
	}
	enabled, _ := r.ReadReplyHeader()
	v.Enabled = x.Uint8ToBool(enabled)

	v.ElementHeader = ElementHeader(r.Read1b())
	r.ReadPad(3)

	interceptedClientsLen := int(r.Read4b()) // 4

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
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	var replyLen uint32
	v.Category, replyLen = r.ReadReplyHeader() // 2

	v.ElementHeader = ElementHeader(r.Read1b())
	v.ClientSwapped = r.ReadBool()
	r.ReadPad(2) // 3

	v.XidBase = r.Read4b()

	v.ServerTime = x.Timestamp(r.Read4b())

	v.RecSequenceNum = r.Read4b() // 6

	// unused
	r.ReadPad(8) // 8

	dataLen := 4 * int(replyLen)
	var err error
	v.Data, err = r.ReadBytes(dataLen)
	return err
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
