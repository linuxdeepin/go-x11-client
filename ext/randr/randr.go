// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package randr

import (
	"errors"
	"math"

	"github.com/linuxdeepin/go-x11-client/ext/render"

	"github.com/linuxdeepin/go-x11-client"
)

const (
	StatusSuccess           = 0
	StatusInvalidConfigTime = 1
	StatusInvalidTime       = 2
	StatusFailed            = 3
)

// #WREQ
func encodeQueryVersion(clientMajorVersion, clientMinorVersion uint32) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(clientMajorVersion).
		Write4b(clientMinorVersion).
		End()
	return
}

type QueryVersionReply struct {
	ServerMajorVersion uint32
	ServerMinorVersion uint32
}

func readQueryVersionReply(r *x.Reader, v *QueryVersionReply) error {
	if !r.RemainAtLeast4b(4) {
		return x.ErrDataLenShort
	}
	r.ReadPad(8)

	v.ServerMajorVersion = r.Read4b()

	v.ServerMinorVersion = r.Read4b() // 4

	return nil
}

// #WREQ
func encodeSetScreenConfig(window x.Window, timestamp x.Timestamp, configTimestamp x.Timestamp,
	sizeID uint16, rotation uint16, rate uint16) (b x.RequestBody) {
	b.AddBlock(5).
		Write4b(uint32(window)).
		Write4b(uint32(timestamp)).
		Write4b(uint32(configTimestamp)). // 3
		Write2b(uint16(sizeID)).
		Write2b(rotation).
		Write2b(rate).
		WritePad(2). // 5
		End()
	return
}

type SetScreenConfigReply struct {
	Status          uint8
	NewTimestamp    x.Timestamp
	ConfigTimestamp x.Timestamp
	Root            x.Window
	SubpixelOrder   uint16
}

func readSetScreenConfigReply(r *x.Reader, v *SetScreenConfigReply) error {
	if !r.RemainAtLeast4b(6) {
		return x.ErrDataLenShort
	}
	v.Status, _ = r.ReadReplyHeader()

	v.NewTimestamp = x.Timestamp(r.Read4b())

	v.ConfigTimestamp = x.Timestamp(r.Read4b()) // 4

	v.Root = x.Window(r.Read4b())

	v.SubpixelOrder = r.Read2b() // 6
	return nil
}

// #WREQ
func encodeGetScreenInfo(window x.Window) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window))
	return
}

type GetScreenInfoReply struct {
	Rotations       uint8
	Root            x.Window
	Timestamp       x.Timestamp
	ConfigTimestamp x.Timestamp
	NSizes          uint16
	SizeID          uint16
	Rotation        uint16
	Rate            uint16
	NInfo           uint16 // useless
	Sizes           []ScreenSize
	Rates           []RefreshRates
}

func readGetScreenInfoReply(r *x.Reader, v *GetScreenInfoReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	v.Rotations, _ = r.ReadReplyHeader()

	v.Root = x.Window(r.Read4b())

	v.Timestamp = x.Timestamp(r.Read4b())

	v.ConfigTimestamp = x.Timestamp(r.Read4b()) // 5

	v.NSizes = r.Read2b()
	v.SizeID = r.Read2b()

	v.Rotation = r.Read2b()
	v.Rate = r.Read2b()

	v.NInfo = r.Read2b()
	r.ReadPad(2) // 8

	if v.NSizes > 0 {
		if !r.RemainAtLeast4b(2 * int(v.NSizes)) {
			return x.ErrDataLenShort
		}
		v.Sizes = make([]ScreenSize, v.NSizes)
		for i := 0; i < int(v.NSizes); i++ {
			v.Sizes[i] = readScreenSize(r)
		}
	}

	nRefreshRates := int(v.NSizes)
	if nRefreshRates > 0 {
		v.Rates = make([]RefreshRates, nRefreshRates)
		for i := 0; i < nRefreshRates; i++ {
			err := readRefreshRates(r, &v.Rates[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// size: 2 * 4b
type ScreenSize struct {
	Width  uint16 // pixels
	Height uint16

	MWidth  uint16 // millimeters
	MHeight uint16
}

func readScreenSize(r *x.Reader) (v ScreenSize) {
	v.Width = r.Read2b()
	v.Height = r.Read2b()

	v.MWidth = r.Read2b()
	v.MHeight = r.Read2b()
	return
}

// size: ?
type RefreshRates struct {
	NRates uint16
	Rates  []uint16
}

func readRefreshRates(r *x.Reader, v *RefreshRates) error {
	if !r.RemainAtLeast(2) {
		return x.ErrDataLenShort
	}
	v.NRates = r.Read2b()
	if v.NRates > 0 {
		if !r.RemainAtLeast(2 * int(v.NRates)) {
			return x.ErrDataLenShort
		}

		v.Rates = make([]uint16, v.NRates)
		for i := 0; i < int(v.NRates); i++ {
			v.Rates[i] = r.Read2b()
		}
	}
	return nil
}

// #WREQ
func encodeGetScreenSizeRange(window x.Window) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

type GetScreenSizeRangeReply struct {
	MinWidth  uint16
	MinHeight uint16
	MaxWidth  uint16
	MaxHeight uint16
}

func readGetScreenSizeRangeReply(r *x.Reader, v *GetScreenSizeRangeReply) error {
	if !r.RemainAtLeast4b(4) {
		return x.ErrDataLenShort
	}
	r.ReadReplyHeader()

	v.MinWidth = r.Read2b()
	v.MinHeight = r.Read2b()

	v.MaxWidth = r.Read2b()
	v.MaxHeight = r.Read2b() // 4

	return nil
}

// #WREQ
func encodeSetScreenSize(window x.Window, width, height uint16,
	mmWidth, mmHeight uint32) (b x.RequestBody) {
	b.AddBlock(4).
		Write4b(uint32(window)).
		Write2b(width).
		Write2b(height).
		Write4b(mmWidth).
		Write4b(mmHeight).
		End()
	return
}

// #WREQ
func encodeGetCrtcInfo(crtc Crtc, configTimestamp x.Timestamp) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(crtc)).
		Write4b(uint32(configTimestamp)).
		End()
	return
}

type GetCrtcInfoReply struct {
	Status          uint8
	Timestamp       x.Timestamp
	X               int16
	Y               int16
	Width           uint16
	Height          uint16
	Mode            Mode
	Rotation        uint16
	Rotations       uint16
	Outputs         []Output
	PossibleOutputs []Output
}

func readGetCrtcInfoReply(r *x.Reader, v *GetCrtcInfoReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	v.Status, _ = r.ReadReplyHeader()

	v.Timestamp = x.Timestamp(r.Read4b())

	v.X = int16(r.Read2b())
	v.Y = int16(r.Read2b())

	v.Width = r.Read2b()
	v.Height = r.Read2b() // 5

	v.Mode = Mode(r.Read4b())

	v.Rotation = r.Read2b()
	v.Rotations = r.Read2b()

	outputsLen := int(r.Read2b())
	possibleOutputsLen := int(r.Read2b()) // 8

	if outputsLen > 0 {
		if !r.RemainAtLeast4b(outputsLen) {
			return x.ErrDataLenShort
		}
		v.Outputs = make([]Output, outputsLen)
		for i := 0; i < outputsLen; i++ {
			v.Outputs[i] = Output(r.Read4b())
		}
	}

	if possibleOutputsLen > 0 {
		if !r.RemainAtLeast4b(possibleOutputsLen) {
			return x.ErrDataLenShort
		}
		v.PossibleOutputs = make([]Output, possibleOutputsLen)
		for i := 0; i < possibleOutputsLen; i++ {
			v.PossibleOutputs[i] = Output(r.Read4b())
		}
	}

	return nil
}

// #WREQ
func encodeGetOutputInfo(output Output, configTimestamp x.Timestamp) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(output)).
		Write4b(uint32(configTimestamp)).
		End()
	return
}

type GetOutputInfoReply struct {
	Status        uint8
	Timestamp     x.Timestamp
	Crtc          Crtc
	MmWidth       uint32
	MmHeight      uint32
	Connection    uint8
	SubPixelOrder uint8
	NumPreferred  uint16
	Crtcs         []Crtc
	Modes         []Mode
	Clones        []Output
	Name          string
}

func readGetOutputInfoReply(r *x.Reader, v *GetOutputInfoReply) error {
	if !r.RemainAtLeast4b(9) {
		return x.ErrDataLenShort
	}
	v.Status, _ = r.ReadReplyHeader()

	v.Timestamp = x.Timestamp(r.Read4b())

	v.Crtc = Crtc(r.Read4b())

	v.MmWidth = r.Read4b() // 5

	v.MmHeight = r.Read4b()

	v.Connection = r.Read1b()
	v.SubPixelOrder = r.Read1b()
	crtcsLen := int(r.Read2b())

	modesLen := int(r.Read2b())
	v.NumPreferred = r.Read2b() // 8

	if int(v.NumPreferred) > modesLen {
		return errors.New("numPreferred > modesLen")
	}

	clonesLen := int(r.Read2b())
	nameLen := int(r.Read2b()) // 9

	if crtcsLen > 0 {
		if !r.RemainAtLeast4b(crtcsLen) {
			return x.ErrDataLenShort
		}
		v.Crtcs = make([]Crtc, crtcsLen)
		for i := 0; i < crtcsLen; i++ {
			v.Crtcs[i] = Crtc(r.Read4b())
		}
	}

	if modesLen > 0 {
		if !r.RemainAtLeast4b(modesLen) {
			return x.ErrDataLenShort
		}
		v.Modes = make([]Mode, modesLen)
		for i := 0; i < modesLen; i++ {
			v.Modes[i] = Mode(r.Read4b())
		}
	}

	if clonesLen > 0 {
		if !r.RemainAtLeast4b(clonesLen) {
			return x.ErrDataLenShort
		}
		v.Clones = make([]Output, clonesLen)
		for i := 0; i < clonesLen; i++ {
			v.Clones[i] = Output(r.Read4b())
		}
	}

	var err error
	v.Name, err = r.ReadString(nameLen)
	return err
}

func (r *GetOutputInfoReply) GetPreferredMode() Mode {
	if r.NumPreferred == 0 || len(r.Modes) == 0 {
		return 0
	}
	return r.Modes[r.NumPreferred-1]
}

// #WREQ
func encodeGetScreenResources(window x.Window) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

type GetScreenResourcesReply struct {
	Timestamp       x.Timestamp
	ConfigTimestamp x.Timestamp
	Crtcs           []Crtc
	Outputs         []Output // size: xgb.Pad((int(NumOutputs) * 4))
	Modes           []ModeInfo
}

// size: 8 * 4b
type ModeInfo struct {
	Id         uint32
	Width      uint16
	Height     uint16
	DotClock   uint32
	HSyncStart uint16
	HSyncEnd   uint16
	HTotal     uint16
	HSkew      uint16
	VSyncStart uint16
	VSyncEnd   uint16
	VTotal     uint16
	nameLen    uint16
	Name       string
	ModeFlags  uint32
}

func readModeInfo(r *x.Reader, v *ModeInfo) {
	v.Id = r.Read4b()

	v.Width = r.Read2b()
	v.Height = r.Read2b()

	v.DotClock = r.Read4b()

	v.HSyncStart = r.Read2b()
	v.HSyncEnd = r.Read2b()

	v.HTotal = r.Read2b()
	v.HSkew = r.Read2b() // 5

	v.VSyncStart = r.Read2b()
	v.VSyncEnd = r.Read2b()

	v.VTotal = r.Read2b()
	v.nameLen = r.Read2b()

	v.ModeFlags = r.Read4b() // 8
}

func writeModeInfo(b *x.FixedSizeBuf, v *ModeInfo) {
	b.Write4b(v.Id).
		Write2b(v.Width).
		Write2b(v.Height).
		Write4b(v.DotClock).
		Write2b(v.HSyncStart).
		Write2b(v.HSyncEnd).
		Write2b(v.HTotal).
		Write2b(v.HSkew).
		Write2b(v.VSyncStart).
		Write2b(v.VSyncEnd).
		Write2b(v.VTotal).
		Write2b(v.nameLen).
		Write4b(v.ModeFlags)
}

func readGetScreenResourcesReply(r *x.Reader, v *GetScreenResourcesReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	r.ReadPad(8)

	v.Timestamp = x.Timestamp(r.Read4b())

	v.ConfigTimestamp = x.Timestamp(r.Read4b())

	crtcsLen := int(r.Read2b())
	outputsLen := int(r.Read2b()) // 5

	modesLen := int(r.Read2b())
	// namesLen
	modeNamesLen := int(r.Read2b())

	// unused
	r.ReadPad(8) // 8

	if crtcsLen > 0 {
		if !r.RemainAtLeast4b(crtcsLen) {
			return x.ErrDataLenShort
		}
		v.Crtcs = make([]Crtc, crtcsLen)
		for i := 0; i < crtcsLen; i++ {
			v.Crtcs[i] = Crtc(r.Read4b())
		}
	}

	if outputsLen > 0 {
		if !r.RemainAtLeast4b(outputsLen) {
			return x.ErrDataLenShort
		}
		v.Outputs = make([]Output, outputsLen)
		for i := 0; i < outputsLen; i++ {
			v.Outputs[i] = Output(r.Read4b())
		}
	}

	if modesLen > 0 {
		if !r.RemainAtLeast4b(8 * modesLen) {
			return x.ErrDataLenShort
		}
		v.Modes = make([]ModeInfo, modesLen)
		for i := 0; i < modesLen; i++ {
			readModeInfo(r, &v.Modes[i])
		}
	}

	var b int
	for i := 0; i < modesLen; i++ {
		modeNameLen := int(v.Modes[i].nameLen)
		b += modeNameLen
		var err error
		v.Modes[i].Name, err = r.ReadString(modeNameLen)
		if err != nil {
			return err
		}
	}

	if b != modeNamesLen {
		return errors.New("mode names len not equal")
	}

	return nil
}

// #WREQ
func encodeGetScreenResourcesCurrent(window x.Window) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

type GetScreenResourcesCurrentReply GetScreenResourcesReply

func readGetScreenResourcesCurrentReply(r *x.Reader, v *GetScreenResourcesCurrentReply) error {
	return readGetScreenResourcesReply(r, (*GetScreenResourcesReply)(v))
}

// #WREQ
func encodeSetCrtcTransform(crtc Crtc, transform *render.Transform, filter string,
	filterParams []render.Fixed) (b x.RequestBody) {

	filter = x.TruncateStr(filter, math.MaxUint16)

	buf := b.AddBlock(11 + x.SizeIn4bWithPad(len(filter)) + len(filterParams)).
		Write4b(uint32(crtc))

	render.WriteTransform(buf, transform) // 10

	buf.Write2b(uint16(len(filter))).
		WritePad(2). // 11
		WriteString(filter).
		WritePad(x.Pad(len(filter)))
	for _, value := range filterParams {
		buf.Write4b(uint32(value)) // +len(filterParams)
	}
	buf.End()
	return
}

// #WREQ
func encodeGetCrtcTransform(crtc Crtc) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(crtc)).
		End()
	return
}

type GetCrtcTransformReply struct {
	PendingTransform render.Transform
	HasTransform     bool
	CurrentTransform render.Transform

	PendingLen     uint16
	PendingNParams uint16

	CurrentLen     uint16
	CurrentNParams uint16

	PendingFilter string
	PendingParams []render.Fixed

	CurrentFilter string
	CurrentParams []render.Fixed
}

func readGetCrtcTransformReply(r *x.Reader, v *GetCrtcTransformReply) error {
	if !r.RemainAtLeast4b(24) {
		return x.ErrDataLenShort
	}
	r.ReadReplyHeader()
	render.ReadTransform(r, &v.PendingTransform) // 11

	v.HasTransform = r.ReadBool()
	r.ReadPad(3) // 12

	render.ReadTransform(r, &v.CurrentTransform) // 21

	r.ReadPad(4) // 22

	v.PendingLen = r.Read2b()
	v.PendingNParams = r.Read2b() // 23

	v.CurrentLen = r.Read2b()
	v.CurrentNParams = r.Read2b() // 24

	var err error
	v.PendingFilter, err = r.ReadStrWithPad(int(v.PendingLen))
	if err != nil {
		return err
	}

	if v.PendingNParams > 0 {
		if !r.RemainAtLeast4b(int(v.PendingNParams)) {
			return x.ErrDataLenShort
		}
		v.PendingParams = make([]render.Fixed, v.PendingNParams)
		for i := 0; i < int(v.PendingNParams); i++ {
			v.PendingParams[i] = render.Fixed(r.Read4b())
		}
	}

	v.CurrentFilter, err = r.ReadStrWithPad(int(v.CurrentLen))
	if err != nil {
		return err
	}

	if v.CurrentNParams > 0 {
		if !r.RemainAtLeast4b(int(v.CurrentNParams)) {
			return x.ErrDataLenShort
		}
		v.CurrentParams = make([]render.Fixed, v.CurrentNParams)
		for i := 0; i < int(v.CurrentNParams); i++ {
			v.CurrentParams[i] = render.Fixed(r.Read4b())
		}
	}
	return nil
}

// size: 6 * 4b
type Panning struct {
	Left, Top, Width, Height                         uint16
	TrackLeft, TrackTop, TrackWidth, TrackHeight     uint16
	BorderLeft, BorderTop, BorderRight, BorderBottom int16
}

func readPanning(r *x.Reader, v *Panning) {
	v.Left = r.Read2b()
	v.Top = r.Read2b()

	v.Width = r.Read2b()
	v.Height = r.Read2b()

	v.TrackLeft = r.Read2b()
	v.TrackTop = r.Read2b()

	v.TrackWidth = r.Read2b()
	v.TrackHeight = r.Read2b()

	v.BorderLeft = int16(r.Read2b())
	v.BorderTop = int16(r.Read2b())
	v.BorderRight = int16(r.Read2b())
	v.BorderBottom = int16(r.Read2b())
}

func writePanning(b *x.FixedSizeBuf, v *Panning) {
	b.Write2b(v.Left)
	b.Write2b(v.Top)

	b.Write2b(v.Width)
	b.Write2b(v.Height)

	b.Write2b(v.TrackLeft)
	b.Write2b(v.TrackTop)

	b.Write2b(v.TrackWidth)
	b.Write2b(v.TrackHeight)

	b.Write2b(uint16(v.BorderLeft))
	b.Write2b(uint16(v.TrackTop))

	b.Write2b(uint16(v.BorderRight))
	b.Write2b(uint16(v.BorderBottom))
}

// #WREQ
func encodeGetPanning(crtc Crtc) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(crtc)).
		End()
	return
}

type GetPanningReply struct {
	Status    uint8
	Timestamp x.Timestamp
	Panning
}

func readGetPanningReply(r *x.Reader, v *GetPanningReply) error {
	if !r.RemainAtLeast4b(9) {
		return x.ErrDataLenShort
	}
	v.Status, _ = r.ReadReplyHeader()

	v.Timestamp = x.Timestamp(r.Read4b()) // 3

	readPanning(r, &v.Panning) // 9
	return nil
}

// #WREQ
func encodeSetPanning(crtc Crtc, timestamp x.Timestamp, panning *Panning) (b x.RequestBody) {
	buf := b.AddBlock(8).
		Write4b(uint32(crtc)).
		Write4b(uint32(timestamp))
	writePanning(buf, panning)
	return
}

type SetPanningReply struct {
	Status    uint8
	Timestamp x.Timestamp
}

func readSetPanningReply(r *x.Reader, v *SetPanningReply) error {
	if !r.RemainAtLeast4b(3) {
		return x.ErrDataLenShort
	}
	v.Status, _ = r.ReadReplyHeader()
	v.Timestamp = x.Timestamp(r.Read4b()) // 3
	return nil
}

// #WREQ
func encodeSetOutputPrimary(window x.Window, output Output) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(window)).
		Write4b(uint32(output)).
		End()
	return
}

// #WREQ
func encodeGetOutputPrimary(window x.Window) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

type GetOutputPrimaryReply struct {
	Output Output
}

func readGetOutputPrimaryReply(r *x.Reader, v *GetOutputPrimaryReply) error {
	if !r.RemainAtLeast4b(3) {
		return x.ErrDataLenShort
	}
	r.ReadPad(8)

	v.Output = Output(r.Read4b()) // 3

	return nil
}

// #WREQ
func encodeListOutputProperties(output Output) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(output)).
		End()
	return
}

type ListOutputPropertiesReply struct {
	NumAtoms uint16
	Atoms    []x.Atom
}

func readListOutputPropertiesReply(r *x.Reader, v *ListOutputPropertiesReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	r.ReadReplyHeader()

	v.NumAtoms = r.Read2b()
	r.ReadPad(22) // 8

	if v.NumAtoms > 0 {
		if !r.RemainAtLeast4b(int(v.NumAtoms)) {
			return x.ErrDataLenShort
		}
		v.Atoms = make([]x.Atom, v.NumAtoms)
		for i := 0; i < int(v.NumAtoms); i++ {
			v.Atoms[i] = x.Atom(r.Read4b())
		}
	}
	return nil
}

// #WREQ
func encodeQueryOutputProperty(output Output, property x.Atom) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(output)).
		Write4b(uint32(property)).
		End()
	return
}

type QueryOutputPropertyReply struct {
	Pending     bool
	Range       bool
	Immutable   bool
	ValidValues []int32
}

func readQueryOutputPropertyReply(r *x.Reader, v *QueryOutputPropertyReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	_, length := r.ReadReplyHeader()

	v.Pending = r.ReadBool()
	v.Range = r.ReadBool()
	v.Immutable = r.ReadBool()
	r.ReadPad(21) // 8

	if length > 0 {
		if !r.RemainAtLeast4b(int(length)) {
			return x.ErrDataLenShort
		}
		v.ValidValues = make([]int32, length)
		for i := 0; i < int(length); i++ {
			v.ValidValues[i] = int32(r.Read4b())
		}
	}
	return nil
}

// #WREQ
func encodeConfigureOutputProperty(output Output, property x.Atom, pending, range0 bool,
	values []int32) (b x.RequestBody) {
	buf := b.AddBlock(3 + len(values)).
		Write4b(uint32(output)).
		Write4b(uint32(property)).
		WriteBool(pending).
		WriteBool(range0).
		WritePad(2) // 3
	for _, value := range values {
		buf.Write4b(uint32(value))
	}
	return
}

// #WREQ
func encodeChangeOutputProperty(output Output, property x.Atom, type0 x.Atom,
	format, mode uint8, data []byte) (b x.RequestBody) {
	numUnits := uint32(len(data) / (int(format) / 8))
	b.AddBlock(5).
		Write4b(uint32(output)).
		Write4b(uint32(property)).
		Write4b(uint32(type0)). // 3
		Write1b(format).
		Write1b(mode).
		WritePad(2).
		Write4b(numUnits). // 5
		End()
	b.AddBytes(data)
	return
}

// #WREQ
func encodeDeleteOutputProperty(output Output, property x.Atom) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(output)).
		Write4b(uint32(property)).
		End()
	return
}

// #WREQ
func encodeGetOutputProperty(output Output, property, Type x.Atom,
	longOffset, longLength uint32, delete, pending bool) (b x.RequestBody) {
	b.AddBlock(6).
		Write4b(uint32(output)).
		Write4b(uint32(property)).
		Write4b(uint32(Type)).
		Write4b(longOffset).
		Write4b(longLength). // 5
		WriteBool(delete).
		WriteBool(pending).
		WritePad(2). // 6
		End()
	return
}

type GetOutputPropertyReply struct {
	Format     byte
	Type       x.Atom
	BytesAfter uint32
	ValueLen   uint32
	Value      []byte
}

func readGetOutputPropertyReply(r *x.Reader, v *GetOutputPropertyReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	v.Format, _ = r.ReadReplyHeader()

	v.Type = x.Atom(r.Read4b())

	v.BytesAfter = r.Read4b()

	v.ValueLen = r.Read4b() // 5

	// unused
	r.ReadPad(12) // 8

	var err error
	n := int(v.ValueLen) * int(v.Format/8)
	v.Value, err = r.ReadBytes(n)
	return err
}

// #WREQ
func encodeCreateMode(window x.Window, modeInfo *ModeInfo) (b x.RequestBody) {
	modeInfo.Name = x.TruncateStr(modeInfo.Name, math.MaxUint16)
	modeInfo.nameLen = uint16(len(modeInfo.Name))

	buf := b.AddBlock(9).
		Write4b(uint32(window))
	writeModeInfo(buf, modeInfo)
	buf.End()
	b.AddBytes([]byte(modeInfo.Name))
	return
}

type CreateModeReply struct {
	Mode Mode
}

func readCreateModeReply(r *x.Reader, v *CreateModeReply) error {
	if !r.RemainAtLeast4b(3) {
		return x.ErrDataLenShort
	}
	r.ReadReplyHeader()
	v.Mode = Mode(r.Read4b())
	return nil
}

// #WREQ
func encodeDestroyMode(mode Mode) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(mode)).
		End()
	return
}

// #WREQ
func encodeAddOutputMode(output Output, mode Mode) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(output)).
		Write4b(uint32(mode)).
		End()
	return
}

// #WREQ
func encodeDeleteOutputMode(output Output, mode Mode) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(output)).
		Write4b(uint32(mode)).
		End()
	return
}

// #WREQ
func encodeSetCrtcConfig(crtc Crtc, timestamp, configTimestamp x.Timestamp,
	X, y int16, mode Mode, rotation uint16, outputs []Output) (b x.RequestBody) {
	b0 := b.AddBlock(6 + len(outputs)).
		Write4b(uint32(crtc)).
		Write4b(uint32(timestamp)).
		Write4b(uint32(configTimestamp)).
		Write2b(uint16(X)).
		Write2b(uint16(y)).
		Write4b(uint32(mode)).
		Write2b(rotation).
		WritePad(2)

	for _, output := range outputs {
		b0.Write4b(uint32(output))
	}
	b0.End()
	return
}

type SetCrtcConfigReply struct {
	Status    uint8
	Timestamp x.Timestamp
}

func readSetCrtcConfigReply(r *x.Reader, v *SetCrtcConfigReply) error {
	if !r.RemainAtLeast4b(3) {
		return x.ErrDataLenShort
	}
	v.Status, _ = r.ReadReplyHeader()

	v.Timestamp = x.Timestamp(r.Read4b()) // 3

	return nil
}

// #WREQ
func encodeSelectInput(window x.Window, enable uint16) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(window)).
		Write2b(enable).
		WritePad(2).
		End()
	return
}

// #WREQ
func encodeGetCrtcGammaSize(crtc Crtc) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(crtc)).
		End()
	return
}

type GetCrtcGammaSizeReply struct {
	Size uint16
}

func readGetCrtcGammaSizeReply(r *x.Reader, v *GetCrtcGammaSizeReply) error {
	r.Read1b()

	r.Read1b()

	// seq
	r.Read2b()

	// length
	r.Read4b()

	v.Size = r.Read2b()

	r.ReadPad(22)

	return nil
}

// #WREQ
func encodeGetCrtcGamma(crtc Crtc) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(crtc)).
		End()
	return
}

type GetCrtcGammaReply struct {
	Size  uint16
	Red   []uint16
	Green []uint16
	Blue  []uint16
}

func readGetCrtcGammaReply(r *x.Reader, v *GetCrtcGammaReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	r.ReadReplyHeader()

	v.Size = r.Read2b()
	r.ReadPad(22) // 8

	if v.Size > 0 {
		if !r.RemainAtLeast(3 * (2*int(v.Size) + x.Pad(2*int(v.Size)))) {
			return x.ErrDataLenShort
		}
		v.Red = make([]uint16, v.Size)
		for i := 0; i < int(v.Size); i++ {
			v.Red[i] = r.Read2b()
		}
		r.ReadPad(x.Pad(int(v.Size) * 2))

		v.Green = make([]uint16, v.Size)
		for i := 0; i < int(v.Size); i++ {
			v.Green[i] = r.Read2b()
		}
		r.ReadPad(x.Pad(int(v.Size) * 2))

		v.Blue = make([]uint16, v.Size)
		for i := 0; i < int(v.Size); i++ {
			v.Blue[i] = r.Read2b()
		}
	}

	return nil
}

// #WREQ
func encodeSetCrtcGamma(crtc Crtc, red, green, blue []uint16) (b x.RequestBody) {
	size := len(red)
	if len(green) != size {
		panic("assert len(green) != size failed")
	}
	if len(blue) != size {
		panic("assert len(blue) != size failed")
	}

	b0 := b.AddBlock(2 + x.SizeIn4bWithPad(6*size)).
		Write4b(uint32(crtc)).
		Write2b(uint16(size)).
		WritePad(2)

	for _, value := range red {
		b0.Write2b(value)
	}

	for _, value := range green {
		b0.Write2b(value)
	}

	for _, value := range blue {
		b0.Write2b(value)
	}
	b0.WritePad(x.Pad(6 * size))
	return
}

// #WREQ
func encodeGetProviders(window x.Window) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

type GetProvidersReply struct {
	Timestamp    x.Timestamp
	NumProviders uint16
	Providers    []Provider
}

func readGetProvidersReply(r *x.Reader, v *GetProvidersReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}

	r.ReadReplyHeader()

	v.Timestamp = x.Timestamp(r.Read4b()) // 3

	v.NumProviders = r.Read2b()
	r.ReadPad(18) // 8

	if v.NumProviders > 0 {
		if !r.RemainAtLeast4b(int(v.NumProviders)) {
			return x.ErrDataLenShort
		}
		v.Providers = make([]Provider, v.NumProviders)
		for i := 0; i < int(v.NumProviders); i++ {
			v.Providers[i] = Provider(r.Read4b())
		}
	}

	return nil
}

// #WREQ
func encodeGetProviderInfo(provider Provider, configTimestamp x.Timestamp) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(provider)).
		Write4b(uint32(configTimestamp)).
		End()
	return
}

type GetProviderInfoReply struct {
	Status                 uint8
	Timestamp              x.Timestamp
	Capabilities           uint32
	NumCrtcs               uint16
	NumOutputs             uint16
	NumAssociatedProviders uint16
	NameLen                uint16
	Crtcs                  []Crtc
	Outputs                []Output
	AssociatedProviders    []Provider
	AssociatedCapability   []uint32
	Name                   string
}

func readGetProviderInfoReply(r *x.Reader, v *GetProviderInfoReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	v.Status, _ = r.ReadReplyHeader()

	v.Timestamp = x.Timestamp(r.Read4b())

	v.Capabilities = r.Read4b()

	v.NumCrtcs = r.Read2b()
	v.NumOutputs = r.Read2b()

	v.NumAssociatedProviders = r.Read2b()
	v.NameLen = r.Read2b() // 6

	r.ReadPad(8) // 8

	if v.NumCrtcs > 0 {
		if !r.RemainAtLeast4b(int(v.NumCrtcs)) {
			return x.ErrDataLenShort
		}
		v.Crtcs = make([]Crtc, v.NumCrtcs)
		for i := 0; i < int(v.NumCrtcs); i++ {
			v.Crtcs[i] = Crtc(r.Read4b())
		}
	}

	if v.NumOutputs > 0 {
		if !r.RemainAtLeast4b(int(v.NumOutputs)) {
			return x.ErrDataLenShort
		}
		v.Outputs = make([]Output, v.NumOutputs)
		for i := 0; i < int(v.NumOutputs); i++ {
			v.Outputs[i] = Output(r.Read4b())
		}
	}

	if v.NumAssociatedProviders > 0 {
		if !r.RemainAtLeast4b(2 * int(v.NumAssociatedProviders)) {
			return x.ErrDataLenShort
		}
		v.AssociatedProviders = make([]Provider, v.NumAssociatedProviders)
		for i := 0; i < int(v.NumAssociatedProviders); i++ {
			v.AssociatedProviders[i] = Provider(r.Read4b())
		}

		v.AssociatedCapability = make([]uint32, v.NumAssociatedProviders)
		for i := 0; i < int(v.NumAssociatedProviders); i++ {
			v.AssociatedCapability[i] = r.Read4b()
		}
	}

	if v.NameLen > 0 {
		var err error
		v.Name, err = r.ReadString(int(v.NameLen))
		if err != nil {
			return err
		}
	}

	return nil
}

// #WREQ
func encodeSetProviderOffloadSink(provider, sinkProvider Provider,
	configTimestamp x.Timestamp) (b x.RequestBody) {
	b.AddBlock(3).
		Write4b(uint32(provider)).
		Write4b(uint32(sinkProvider)).
		Write4b(uint32(configTimestamp)).
		End()
	return
}

// #WREQ
func encodeSetProviderOutputSource(provider, sourceProvider Provider,
	configTimestamp x.Timestamp) (b x.RequestBody) {
	b.AddBlock(3).
		Write4b(uint32(provider)).
		Write4b(uint32(sourceProvider)).
		Write4b(uint32(configTimestamp)).
		End()
	return
}

// #WREQ
func encodeListProviderProperties(provider Provider) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(provider)).
		End()
	return
}

type ListProviderPropertiesReply struct {
	NumAtoms uint16
	Atoms    []x.Atom
}

func readListProviderPropertiesReply(r *x.Reader, v *ListProviderPropertiesReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	r.ReadReplyHeader()

	v.NumAtoms = r.Read2b()
	r.ReadPad(22) // 8

	if v.NumAtoms > 0 {
		if !r.RemainAtLeast4b(int(v.NumAtoms)) {
			return x.ErrDataLenShort
		}
		v.Atoms = make([]x.Atom, v.NumAtoms)
		for i := 0; i < int(v.NumAtoms); i++ {
			v.Atoms[i] = x.Atom(r.Read4b())
		}
	}
	return nil
}

// #WREQ
func encodeGetMonitors(window x.Window, getActive bool) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(window)).
		WriteBool(getActive).
		WritePad(3).
		End()
	return
}

type GetMonitorsReply struct {
	Timestamp x.Timestamp
	NMonitors uint32
	NOutputs  uint32
	Monitors  []MonitorInfo
}

// size: (6 + NOutputs) * 4b
type MonitorInfo struct {
	Name              x.Atom
	Primary           bool
	Automatic         bool
	NOutputs          uint16
	X, Y              int16
	Width, Height     uint16 // pixels
	MmWidth, MmHeight uint32 // in millimeters
	Outputs           []Output
}

func readMonitorInfo(r *x.Reader, v *MonitorInfo) error {
	if !r.RemainAtLeast4b(6) {
		return x.ErrDataLenShort
	}
	v.Name = x.Atom(r.Read4b())

	v.Primary = r.ReadBool()
	v.Automatic = r.ReadBool()
	v.NOutputs = r.Read2b()

	v.X = int16(r.Read2b())
	v.Y = int16(r.Read2b())

	v.Width = r.Read2b()
	v.Height = r.Read2b()

	v.MmWidth = r.Read4b()

	v.MmHeight = r.Read4b() // 6

	if v.NOutputs > 0 {
		v.Outputs = make([]Output, v.NOutputs)
		for i := 0; i < int(v.NOutputs); i++ {
			v.Outputs[i] = Output(r.Read4b())
		}
	}
	return nil
}

func writeMonitorInfo(b *x.FixedSizeBuf, v *MonitorInfo) {
	v.NOutputs = uint16(len(v.Outputs))
	b.Write4b(uint32(v.Name)).
		WriteBool(v.Primary).
		WriteBool(v.Automatic).
		Write2b(v.NOutputs).
		Write2b(uint16(v.X)).
		Write2b(uint16(v.Y)).
		Write2b(v.Width).
		Write2b(v.Height).
		Write4b(v.MmWidth).
		Write4b(v.MmHeight) // 6
	for _, o := range v.Outputs {
		b.Write4b(uint32(o))
	}
}

func readGetMonitorsReply(r *x.Reader, v *GetMonitorsReply) error {
	r.ReadReplyHeader()

	v.Timestamp = x.Timestamp(r.Read4b())

	v.NMonitors = r.Read4b()

	v.NOutputs = r.Read4b()

	r.ReadPad(12)

	if v.NMonitors > 0 {
		v.Monitors = make([]MonitorInfo, v.NMonitors)
		for i := 0; i < int(v.NMonitors); i++ {
			err := readMonitorInfo(r, &v.Monitors[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// #WREQ
func encodeSetMonitor(window x.Window, monitorInfo *MonitorInfo) (b x.RequestBody) {
	sizeIn4b := 7 + len(monitorInfo.Outputs)
	buf := b.AddBlock(sizeIn4b).
		Write4b(uint32(window))
	writeMonitorInfo(buf, monitorInfo)
	buf.End()
	return
}

// #WREQ
func encodeDeleteMonitor(window x.Window, name x.Atom) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(window)).
		Write4b(uint32(name)).
		End()
	return
}
