package randr

import (
	"errors"

	"github.com/linuxdeepin/go-x11-client"
)

// #WREQ
func writeQueryVersion(w *x.Writer, clientMajorVersion, clientMinorVersion uint32) {
	w.WritePad(4)
	w.Write4b(clientMajorVersion)
	w.Write4b(clientMinorVersion)
}

type QueryVersionReply struct {
	ServerMajorVersion uint32
	ServerMinorVersion uint32
}

func readQueryVersionReply(r *x.Reader, v *QueryVersionReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

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

	v.ServerMajorVersion = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ServerMinorVersion = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeGetCrtcInfo(w *x.Writer, crtc Crtc, configTimestamp x.Timestamp) {
	w.WritePad(4)
	w.Write4b(uint32(crtc))
	w.Write4b(uint32(configTimestamp))
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
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Status = r.Read1b()
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

	v.Timestamp = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.X = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Y = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Width = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Height = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Mode = Mode(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Rotation = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Rotations = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	outputsLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	possibleOutputsLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	if outputsLen > 0 {
		v.Outputs = make([]Output, outputsLen)
		for i := 0; i < outputsLen; i++ {
			v.Outputs[i] = Output(r.Read4b())
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	if possibleOutputsLen > 0 {
		v.PossibleOutputs = make([]Output, possibleOutputsLen)
		for i := 0; i < possibleOutputsLen; i++ {
			v.PossibleOutputs[i] = Output(r.Read4b())
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	return nil
}

// #WREQ
func writeGetOutputInfo(w *x.Writer, output Output, configTimestamp x.Timestamp) {
	w.WritePad(4)
	w.Write4b(uint32(output))
	w.Write4b(uint32(configTimestamp))
}

type GetOutputInfoReply struct {
	Status        uint8
	Timestamp     x.Timestamp
	Crtc          Crtc
	MmWidth       uint32
	MmHeight      uint32
	Connection    uint8
	SubpixelOrder uint8
	NumPreferred  uint16
	Crtcs         []Crtc
	Modes         []Mode
	Clones        []Output
	Name          string
}

func readGetOutputInfoReply(r *x.Reader, v *GetOutputInfoReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Status = r.Read1b()
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

	v.Timestamp = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Crtc = Crtc(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.MmWidth = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MmHeight = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Connection = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SubpixelOrder = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	crtcsLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	modesLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.NumPreferred = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	if int(v.NumPreferred) > modesLen {
		return errors.New("numPreferred > modesLen")
	}

	clonesLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	nameLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	if crtcsLen > 0 {
		v.Crtcs = make([]Crtc, crtcsLen)
		for i := 0; i < crtcsLen; i++ {
			v.Crtcs[i] = Crtc(r.Read4b())
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	if modesLen > 0 {
		v.Modes = make([]Mode, modesLen)
		for i := 0; i < modesLen; i++ {
			v.Modes[i] = Mode(r.Read4b())
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	if clonesLen > 0 {
		v.Clones = make([]Output, clonesLen)
		for i := 0; i < clonesLen; i++ {
			v.Clones[i] = Output(r.Read4b())
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	v.Name = r.ReadString(nameLen)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

func (r *GetOutputInfoReply) GetPreferredMode() Mode {
	return r.Modes[r.NumPreferred-1]
}

// #WREQ
func writeGetScreenResources(w *x.Writer, window x.Window) {
	w.WritePad(4)
	w.Write4b(uint32(window))
}

type GetScreenResourcesReply struct {
	Timestamp       x.Timestamp
	ConfigTimestamp x.Timestamp
	Crtcs           []Crtc
	Outputs         []Output // size: xgb.Pad((int(NumOutputs) * 4))
	Modes           []ModeInfo
}

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

func readModeInfo(r *x.Reader, v *ModeInfo) error {
	v.Id = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Width = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Height = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.DotClock = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.HSyncStart = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.HSyncEnd = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.HTotal = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.HSkew = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.VSyncStart = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.VSyncEnd = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.VTotal = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.nameLen = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ModeFlags = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

func readGetScreenResourcesReply(r *x.Reader, v *GetScreenResourcesReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

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

	v.Timestamp = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.ConfigTimestamp = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	crtcsLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	outputsLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	modesLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	// namesLen
	modeNamesLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(8)
	if r.Err() != nil {
		return r.Err()
	}

	if crtcsLen > 0 {
		v.Crtcs = make([]Crtc, crtcsLen)
		for i := 0; i < crtcsLen; i++ {
			v.Crtcs[i] = Crtc(r.Read4b())
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	if outputsLen > 0 {
		v.Outputs = make([]Output, outputsLen)
		for i := 0; i < outputsLen; i++ {
			v.Outputs[i] = Output(r.Read4b())
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	if modesLen > 0 {
		v.Modes = make([]ModeInfo, modesLen)
		for i := 0; i < modesLen; i++ {
			err := readModeInfo(r, &v.Modes[i])
			if err != nil {
				return err
			}
		}
	}

	var b int
	for i := 0; i < modesLen; i++ {
		modeNameLen := int(v.Modes[i].nameLen)
		b += modeNameLen
		v.Modes[i].Name = r.ReadString(modeNameLen)
		if r.Err() != nil {
			return r.Err()
		}
	}

	if b != modeNamesLen {
		return errors.New("mode names len not equal")
	}

	r.ReadPad(x.Pad(b))
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeGetOutputPrimary(w *x.Writer, window x.Window) {
	w.WritePad(4)
	w.Write4b(uint32(window))
}

type GetOutputPrimaryReply struct {
	Output Output
}

func readGetOutputPrimaryReply(r *x.Reader, v *GetOutputPrimaryReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

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

	v.Output = Output(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(16)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeGetOutputProperty(w *x.Writer, output Output, property, type0 x.Atom,
	longOffset, longLength uint32, delete, pending bool) {
	w.WritePad(4)
	w.Write4b(uint32(output))
	w.Write4b(uint32(property))
	w.Write4b(uint32(type0))
	w.Write4b(longOffset)
	w.Write4b(longLength)
	w.Write1b(x.BoolToUint8(delete))
	w.Write1b(x.BoolToUint8(pending))
	w.WritePad(2)
}

type GetOutputPropertyReply struct {
	Format     byte
	Type       x.Atom
	BytesAfter uint32
	ValueLen   uint32
	Value      []byte
}

func readGetOutputPropertyReply(r *x.Reader, v *GetOutputPropertyReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Format = r.Read1b()
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

	v.Type = x.Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.BytesAfter = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ValueLen = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(12)
	if r.Err() != nil {
		return r.Err()
	}

	n := int(v.ValueLen) * int(v.Format/8)
	v.Value = r.ReadBytes(n)
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(x.Pad(n))
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeSetCrtcConfig(w *x.Writer, crtc Crtc, timestamp, configTimestamp x.Timestamp,
	X, y int16, mode Mode, rotation uint16, outputs []Output) {
	w.WritePad(4)
	w.Write4b(uint32(crtc))
	w.Write4b(uint32(timestamp))
	w.Write4b(uint32(configTimestamp))
	w.Write2b(uint16(X))
	w.Write2b(uint16(y))
	w.Write4b(uint32(mode))
	w.Write2b(rotation)
	w.WritePad(2)
	for _, output := range outputs {
		w.Write4b(uint32(output))
	}
}

type SetCrtcConfigReply struct {
	Status    uint8
	Timestamp x.Timestamp
}

func readSetCrtcConfigReply(r *x.Reader, v *SetCrtcConfigReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Status = r.Read1b()
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

	v.Timestamp = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(20)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeSelectInput(w *x.Writer, window x.Window, enable uint16) {
	w.WritePad(4)
	w.Write4b(uint32(window))
	w.Write2b(enable)
	w.WritePad(2)
}

// #WREQ
func writeGetCrtcGammaSize(w *x.Writer, crtc Crtc) {
	w.WritePad(4)
	w.Write4b(uint32(crtc))
}

type GetCrtcGammaSizeReply struct {
	Size uint16
}

func readGetCrtcGammaSizeReply(r *x.Reader, v *GetCrtcGammaSizeReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

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

	v.Size = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(22)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func writeSetCrtcGamma(w *x.Writer, crtc Crtc, red, green, blue []uint16) {
	size := len(red)
	if len(green) != size {
		panic("assert len(green) != size failed")
	}
	if len(blue) != size {
		panic("assert len(blue) != size failed")
	}

	w.WritePad(4)
	w.Write4b(uint32(crtc))

	w.Write2b(uint16(size))
	w.WritePad(2)

	for _, value := range red {
		w.Write2b(value)
	}

	for _, value := range green {
		w.Write2b(value)
	}

	for _, value := range blue {
		w.Write2b(value)
	}
	w.WritePad(x.Pad(6 * size))
}
