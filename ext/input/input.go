package input

import (
	"github.com/linuxdeepin/go-x11-client"
)

func readFP1616(r *x.Reader) (float32, error) {
	value := int32(r.Read4b())
	ret := float32(value) / float32(1<<16)
	return ret, nil
}

func writeFP1616(b *x.FixedSizeBuf, value float32) {
	b.Write4b(uint32(int32(value * float32(1<<16))))
}

func readFP3232(r *x.Reader) (float64, error) {
	integral := int32(r.Read4b())
	if r.Err() != nil {
		return 0, r.Err()
	}
	frac := r.Read4b()
	if r.Err() != nil {
		return 0, r.Err()
	}
	ret := float64(integral) + (float64(frac) * (1.0 / (1 << 32)))
	return ret, nil
}

// #WREQ
func encodeXIQueryVersion(majorVersion, minorVersion uint16) (b x.RequestBody) {
	b.AddBlock(1).
		Write2b(majorVersion).
		Write2b(minorVersion).
		End()
	return
}

type XIQueryVersionReply struct {
	MajorVersion uint16
	MinorVersion uint16
}

func readXIQueryVersionReply(r *x.Reader, v *XIQueryVersionReply) error {
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

	v.MajorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MinorVersion = r.Read2b()
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
func encodeXIQueryDevice(deviceId DeviceId) (b x.RequestBody) {
	b.AddBlock(1).
		Write2b(uint16(deviceId)).
		WritePad(2).
		End()
	return
}

type XIQueryDeviceReply struct {
	Infos []XIDeviceInfo
}

func readXIQueryDeviceReply(r *x.Reader, v *XIQueryDeviceReply) error {
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

	infosLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(22)
	if r.Err() != nil {
		return r.Err()
	}

	if infosLen > 0 {
		v.Infos = make([]XIDeviceInfo, infosLen)
		for i := 0; i < infosLen; i++ {
			err := readXIDeviceInfo(r, &v.Infos[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type XIDeviceInfo struct {
	Id         DeviceId
	Type       uint16
	Attachment DeviceId
	Enabled    bool
	Name       string
	Classes    []DeviceClass
}

func readXIDeviceInfo(r *x.Reader, v *XIDeviceInfo) error {
	v.Id = DeviceId(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Type = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Attachment = DeviceId(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	classesLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	nameLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Enabled = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	var err error
	v.Name, err = r.ReadString(nameLen)
	if err != nil {
		return err
	}

	r.ReadPad(x.Pad(nameLen))
	if r.Err() != nil {
		return r.Err()
	}

	// classes
	if classesLen > 0 {
		v.Classes = make([]DeviceClass, classesLen)
		for i := 0; i < classesLen; i++ {
			v.Classes[i], err = readDeviceClass(r)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type DeviceClass interface {
	GetDeviceClassType() uint16
	GetSourceId() DeviceId
}

func readDeviceClass(r *x.Reader) (DeviceClass, error) {
	var gdc GenericDeviceClass
	err := readGenericDeviceClass(r, &gdc)
	if err != nil {
		return nil, err
	}

	dataReader := x.NewReaderFromData(gdc.Data)
	var ret DeviceClass = &gdc
	switch gdc.Type {
	case DeviceClassTypeKey:
		var c KeyClass
		c.SourceId = gdc.SourceId
		err := readKeyClass(dataReader, &c)
		if err != nil {
			return nil, err
		}
		ret = &c
	case DeviceClassTypeButton:
		var c ButtonClass
		c.SourceId = gdc.SourceId
		err := readButtonClass(dataReader, &c)
		if err != nil {
			return nil, err
		}
		ret = &c
	case DeviceClassTypeValuator:
		var c ValuatorClass
		c.SourceId = gdc.SourceId
		err := readValuatorClass(dataReader, &c)
		if err != nil {
			return nil, err
		}
		ret = &c
	case DeviceClassTypeScroll:
		var c ScrollClass
		c.SourceId = gdc.SourceId
		err := readScrollClass(dataReader, &c)
		if err != nil {
			return nil, err
		}
		ret = &c
	case DeviceClassTypeTouch:
		var c TouchClass
		c.SourceId = gdc.SourceId
		err := readTouchClass(dataReader, &c)
		if err != nil {
			return nil, err
		}
		ret = &c
	}

	return ret, nil
}

type GenericDeviceClass struct {
	Type     uint16
	Len      uint16
	SourceId DeviceId
	Data     []byte
}

func (dc *GenericDeviceClass) GetDeviceClassType() uint16 {
	return dc.Type
}

func (dc *GenericDeviceClass) GetSourceId() DeviceId {
	return dc.SourceId
}

func readGenericDeviceClass(r *x.Reader, v *GenericDeviceClass) error {
	v.Type = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Len = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SourceId = DeviceId(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	var err error
	v.Data, err = r.ReadBytes((int(v.Len) * 4) - 6)
	if err != nil {
		return err
	}
	return nil
}

type KeyClass struct {
	SourceId DeviceId
	NumKeys  uint16
	Keys     []uint32
}

func (*KeyClass) GetDeviceClassType() uint16 {
	return DeviceClassTypeKey
}

func (kc *KeyClass) GetSourceId() DeviceId {
	return kc.SourceId
}

func readKeyClass(r *x.Reader, v *KeyClass) error {
	v.NumKeys = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	numKeys := int(v.NumKeys)
	if numKeys > 0 {
		v.Keys = make([]uint32, numKeys)
		for i := 0; i < numKeys; i++ {
			v.Keys[i] = r.Read4b()
			if r.Err() != nil {
				return r.Err()
			}
		}
	}
	return nil
}

type ButtonClass struct {
	SourceId   DeviceId
	NumButtons uint16
	State      []uint32
	Labels     []x.Atom
}

func (*ButtonClass) GetDeviceClassType() uint16 {
	return DeviceClassTypeButton
}

func (bc *ButtonClass) GetSourceId() DeviceId {
	return bc.SourceId
}

func readButtonClass(r *x.Reader, v *ButtonClass) error {
	v.NumButtons = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	numButtons := int(v.NumButtons)
	stateLen := (numButtons + 31) / 32
	if stateLen > 0 {
		v.State = make([]uint32, stateLen)
		for i := 0; i < stateLen; i++ {
			v.State[i] = r.Read4b()
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	if numButtons > 0 {
		v.Labels = make([]x.Atom, numButtons)
		for i := 0; i < numButtons; i++ {
			v.Labels[i] = x.Atom(r.Read4b())
			if r.Err() != nil {
				return r.Err()
			}
		}
	}
	return nil
}

// also called AxisClass
type ValuatorClass struct {
	SourceId   DeviceId
	Number     uint16
	Label      x.Atom
	Min        float64
	Max        float64
	Value      float64
	Resolution uint32
	Mode       uint8
}

func (*ValuatorClass) GetDeviceClassType() uint16 {
	return DeviceClassTypeValuator
}

func (vc *ValuatorClass) GetSourceId() DeviceId {
	return vc.SourceId
}

func readValuatorClass(r *x.Reader, v *ValuatorClass) error {
	v.Number = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Label = x.Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	var err error
	v.Min, err = readFP3232(r)
	if err != nil {
		return err
	}

	v.Max, err = readFP3232(r)
	if err != nil {
		return err
	}

	v.Value, err = readFP3232(r)
	if err != nil {
		return err
	}

	v.Resolution = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Mode = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(3)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type ScrollClass struct {
	SourceId   DeviceId
	Number     uint16
	ScrollType uint16
	Flags      uint32
	Increment  float64
}

func (*ScrollClass) GetDeviceClassType() uint16 {
	return DeviceClassTypeScroll
}

func (sc *ScrollClass) GetSourceId() DeviceId {
	return sc.SourceId
}

func readScrollClass(r *x.Reader, v *ScrollClass) error {
	v.Number = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ScrollType = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(2)
	if r.Err() != nil {
		return r.Err()
	}

	v.Flags = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	var err error
	v.Increment, err = readFP3232(r)
	if err != nil {
		return err
	}

	return nil
}

type TouchClass struct {
	SourceId   DeviceId
	Mode       uint8
	NumTouches uint8
}

func (*TouchClass) GetDeviceClassType() uint16 {
	return DeviceClassTypeTouch
}

func (tc *TouchClass) GetSourceId() DeviceId {
	return tc.SourceId
}

func readTouchClass(r *x.Reader, v *TouchClass) error {
	v.Mode = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NumTouches = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func encodeXISelectEvents(window x.Window, masks []EventMask) (b x.RequestBody) {
	var masksSize int
	for _, mask := range masks {
		masksSize += mask.getSizeIn4b()
	}
	b0 := b.AddBlock(2 + masksSize).
		Write4b(uint32(window)).
		Write2b(uint16(len(masks))).
		WritePad(2)

	for _, mask := range masks {
		writeEventMask(b0, mask)
	}
	b0.End()
	return
}

type EventMask struct {
	DeviceId DeviceId
	Mask     []uint32
}

func (ev *EventMask) getSizeIn4b() int {
	return 1 + len(ev.Mask)
}

func writeEventMask(b *x.FixedSizeBuf, v EventMask) {
	maskLen := len(v.Mask)
	b.Write2b(uint16(v.DeviceId)).
		Write2b(uint16(maskLen))
	for _, mask := range v.Mask {
		b.Write4b(mask)
	}
}

func readEventMask(r *x.Reader, v *EventMask) error {
	v.DeviceId = DeviceId(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	maskLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	if maskLen > 0 {
		v.Mask = make([]uint32, maskLen)
		for i := 0; i < maskLen; i++ {
			v.Mask[i] = r.Read4b()
			if r.Err() != nil {
				return r.Err()
			}
		}
	}
	return nil
}

// #WREQ
func encodeXIGetSelectedEvents(window x.Window) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

type XIGetSelectedEventsReply struct {
	Masks []EventMask
}

func readXIGetSelectedEventsReply(r *x.Reader, v *XIGetSelectedEventsReply) error {
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

	masksLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(22)
	if r.Err() != nil {
		return r.Err()
	}

	if masksLen > 0 {
		v.Masks = make([]EventMask, masksLen)
		for i := 0; i < masksLen; i++ {
			err := readEventMask(r, &v.Masks[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// #WREQ
func encodeXIQueryPointer(window x.Window, deviceId DeviceId) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(window)).
		Write2b(uint16(deviceId)).
		WritePad(2).
		End()
	return
}

type XIQueryPointerReply struct {
	Root       x.Window
	Child      x.Window
	RootX      float32
	RootY      float32
	WinX       float32
	WinY       float32
	SameScreen bool
	Buttons    []uint32
	Mods       ModifierInfo
	Group      GroupInfo
}

func readXIQueryPointerReply(r *x.Reader, v *XIQueryPointerReply) error {
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

	v.Root = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Child = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	var err error
	v.RootX, err = readFP1616(r)
	if err != nil {
		return err
	}

	v.RootY, err = readFP1616(r)
	if err != nil {
		return err
	}

	v.WinX, err = readFP1616(r)
	if err != nil {
		return err
	}

	v.WinY, err = readFP1616(r)
	if err != nil {
		return err
	}

	v.SameScreen = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	buttonsLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	// mods
	err = readModifierInfo(r, &v.Mods)
	if err != nil {
		return err
	}

	// group
	err = readGroupInfo(r, &v.Group)
	if err != nil {
		return err
	}

	if buttonsLen > 0 {
		v.Buttons = make([]uint32, buttonsLen)
		for i := 0; i < buttonsLen; i++ {
			v.Buttons[i] = r.Read4b()
			if r.Err() != nil {
				return r.Err()
			}
		}
	}

	return nil
}

type ModifierInfo struct {
	Base      uint32
	Latched   uint32
	Locked    uint32
	Effective uint32
}

func readModifierInfo(r *x.Reader, v *ModifierInfo) error {
	v.Base = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Latched = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Locked = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Effective = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

type GroupInfo struct {
	Base      uint8
	Latched   uint8
	Locked    uint8
	Effective uint8
}

func readGroupInfo(r *x.Reader, v *GroupInfo) error {
	v.Base = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Latched = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Locked = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Effective = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func encodeXIWarpPointer(srcWin, dstWin x.Window, srcX, srcY float32, srcWidth,
	srcHeight uint16, dstX, dstY float32, deviceId DeviceId) (b x.RequestBody) {

	b0 := b.AddBlock(8).
		Write4b(uint32(srcWin)).
		Write4b(uint32(dstWin))
	writeFP1616(b0, srcX)
	writeFP1616(b0, srcY)
	b0.Write2b(srcWidth).
		Write2b(srcHeight)
	writeFP1616(b0, dstX)
	writeFP1616(b0, dstY)
	b0.Write2b(uint16(deviceId)).
		WritePad(2).
		End()
	return
}

// #WREQ
func encodeXIChangeCursor(window x.Window, cursor x.Cursor, deviceId DeviceId) (b x.RequestBody) {
	b.AddBlock(3).
		Write4b(uint32(window)).
		Write4b(uint32(cursor)).
		Write2b(uint16(deviceId)).
		WritePad(2).
		End()
	return
}

// #WREQ
func encodeXIChangeHierarchy(changes []HierarchyChange) (b x.RequestBody) {
	var changesSize int
	for _, change := range changes {
		changesSize += int(change.getLen())
	}
	b0 := b.AddBlock(1 + changesSize).
		Write1b(uint8(len(changes))).
		WritePad(3)

	for _, change := range changes {
		writeHierarchyChange(b0, change)
	}
	b0.End()
	return
}

func writeHierarchyChange(b *x.FixedSizeBuf, v HierarchyChange) {
	b.Write2b(v.getHierarchyType())
	b.Write2b(v.getLen() + 1)
	v.writeTo(b)
}

type HierarchyChange interface {
	writeTo(b *x.FixedSizeBuf)
	getHierarchyType() uint16
	getLen() uint16 // Length in 4 byte units.
}

// size (1 + (len(Name) + pad(len(name))) / 4) * 4 byte
type AddMaster struct {
	SendCore bool
	Enable   bool
	Name     string
}

func (am *AddMaster) writeTo(b *x.FixedSizeBuf) {
	nameLen := len(am.Name)
	b.Write2b(uint16(nameLen))
	b.Write1b(x.BoolToUint8(am.SendCore))
	b.Write1b(x.BoolToUint8(am.Enable))

	b.WriteString(am.Name)
	b.WritePad(x.Pad(nameLen))
}

func (*AddMaster) getHierarchyType() uint16 {
	return HierarchyChangeTypeAddMaster
}

func (am *AddMaster) getLen() uint16 {
	nameLen := len(am.Name)
	return uint16(1 + (nameLen+x.Pad(nameLen))/4)
}

// size 2 * 4byte
type RemoveMaster struct {
	DeviceId       DeviceId
	ReturnMode     uint8
	ReturnPointer  DeviceId
	ReturnKeyboard DeviceId
}

func (rm *RemoveMaster) writeTo(b *x.FixedSizeBuf) {
	b.Write2b(uint16(rm.DeviceId))
	b.Write1b(rm.ReturnMode)
	b.WritePad(1)

	b.Write2b(uint16(rm.ReturnPointer))
	b.Write2b(uint16(rm.ReturnKeyboard))
}

func (*RemoveMaster) getHierarchyType() uint16 {
	return HierarchyChangeTypeRemoveMaster
}

func (rm *RemoveMaster) getLen() uint16 {
	return 2
}

// size 1 * 4 byte
type AttachSlave struct {
	DeviceId DeviceId
	Master   DeviceId
}

func (as *AttachSlave) writeTo(b *x.FixedSizeBuf) {
	b.Write2b(uint16(as.DeviceId))
	b.Write2b(uint16(as.Master))
}

func (*AttachSlave) getHierarchyType() uint16 {
	return HierarchyChangeTypeAttachSlave
}

func (as *AttachSlave) getLen() uint16 {
	return 1
}

// size 1 * 4 byte
type DetachSlave struct {
	DeviceId DeviceId
}

func (ds *DetachSlave) writeTo(b *x.FixedSizeBuf) {
	b.Write2b(uint16(ds.DeviceId))
	b.WritePad(2)
}

func (*DetachSlave) getHierarchyType() uint16 {
	return HierarchyChangeTypeDetachSlave
}

func (ds *DetachSlave) getLen() uint16 {
	return 1
}

// #WREQ
func encodeXISetClientPointer(window x.Window, deviceId DeviceId) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(window)).
		Write2b(uint16(deviceId)).
		WritePad(2).
		End()
	return
}

// #WREQ
func encodeXIGetClientPointer(window x.Window) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(window)).
		End()
	return
}

type XIGetClientPointerReply struct {
	Set      bool
	DeviceId DeviceId
}

func readXIGetClientPointerReply(r *x.Reader, v *XIGetClientPointerReply) error {
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

	v.Set = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.DeviceId = DeviceId(r.Read2b())
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
func encodeXISetFocus(focus x.Window, time x.Timestamp, deviceId DeviceId) (b x.RequestBody) {
	b.AddBlock(3).
		Write4b(uint32(focus)).
		Write4b(uint32(time)).
		Write2b(uint16(deviceId)).
		WritePad(2).
		End()
	return
}

// #WREQ
func encodeXIGetFocus(deviceId DeviceId) (b x.RequestBody) {
	b.AddBlock(1).
		Write2b(uint16(deviceId)).
		WritePad(2).
		End()
	return
}

type XIGetFocusReply struct {
	Focus x.Window
}

func readXIGetFocusReply(r *x.Reader, v *XIGetFocusReply) error {
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

	v.Focus = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func encodeXIGrabDevice(window x.Window, time x.Timestamp, cursor x.Cursor, deviceId DeviceId,
	mode, pairedDeviceMode uint8, ownerEvents bool, masks []EventMask) (b x.RequestBody) {
	var masksSize int
	for _, mask := range masks {
		masksSize += mask.getSizeIn4b()
	}
	b0 := b.AddBlock(5 + masksSize).
		Write4b(uint32(window)).
		Write4b(uint32(time)).
		Write4b(uint32(cursor)).
		Write2b(uint16(deviceId)).
		Write1b(mode).
		Write1b(pairedDeviceMode).
		Write1b(x.BoolToUint8(ownerEvents)).
		WritePad(1).
		Write2b(uint16(len(masks)))

	for _, mask := range masks {
		writeEventMask(b0, mask)
	}
	b0.End()
	return
}

type XIGrabDeviceReply struct {
	Status uint8
}

func readXIGrabDeviceReply(r *x.Reader, v *XIGrabDeviceReply) error {
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

	v.Status = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(23)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func encodeXIUngrabDevice(time x.Timestamp, deviceId DeviceId) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(time)).
		Write2b(uint16(deviceId)).
		WritePad(2).
		End()
	return
}

// #WREQ
func encodeXIAllowEvents(time x.Timestamp, deviceId DeviceId, eventMode uint8,
	touchId uint32, grabWindow x.Window) (b x.RequestBody) {
	b.AddBlock(4).
		Write4b(uint32(time)).
		Write2b(uint16(deviceId)).
		Write1b(eventMode).
		WritePad(1).
		Write4b(touchId).
		Write4b(uint32(grabWindow)).
		End()
	return
}

// #WREQ
func encodeXIPassiveGrabDevice(grabWindow x.Window, cursor x.Cursor, detail uint32,
	deviceId DeviceId, grabType, grabMode, pairedDeviceMode uint8, ownerEvents bool,
	masks []uint32, modifiers []uint32) (b x.RequestBody) {

	b0 := b.AddBlock(7 + len(masks) + len(modifiers)).
		WritePad(4). // unused time field
		Write4b(uint32(grabWindow)).
		Write4b(uint32(cursor)).
		Write4b(detail).
		Write2b(uint16(deviceId)).
		Write2b(uint16(len(modifiers))).
		Write2b(uint16(len(masks))).
		Write1b(grabType).
		Write1b(grabMode).
		Write1b(pairedDeviceMode).
		Write1b(x.BoolToUint8(ownerEvents)).
		WritePad(2)

	for _, mask := range masks {
		b0.Write4b(mask)
	}

	for _, mod := range modifiers {
		b0.Write4b(mod)
	}
	b0.End()
	return
}

type XIPassiveGrabDeviceReply struct {
	Modifiers []GrabModifierInfo
}

func readXIPassiveGrabDeviceReply(r *x.Reader, v *XIPassiveGrabDeviceReply) error {
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

	modifiersLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(22)
	if r.Err() != nil {
		return r.Err()
	}

	if modifiersLen > 0 {
		v.Modifiers = make([]GrabModifierInfo, modifiersLen)
		for i := 0; i < modifiersLen; i++ {
			err := readGrabModifierInfo(r, &v.Modifiers[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

type GrabModifierInfo struct {
	Modifiers uint32
	Status    uint8
}

func readGrabModifierInfo(r *x.Reader, v *GrabModifierInfo) error {
	v.Modifiers = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Status = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(3)
	if r.Err() != nil {
		return r.Err()
	}
	return nil
}

// #WREQ
func encodeXIPassiveUngrabDevice(grabWindow x.Window, detail uint32,
	deviceId DeviceId, grabType uint8, modifiers []uint32) (b x.RequestBody) {

	b0 := b.AddBlock(4 + len(modifiers)).
		Write4b(uint32(grabWindow)).
		Write4b(detail).
		Write2b(uint16(deviceId)).
		Write2b(uint16(len(modifiers))).
		Write1b(grabType).
		WritePad(3)

	for _, mod := range modifiers {
		b0.Write4b(mod)
	}
	b0.End()
	return
}

// #WREQ
func encodeXIListProperties(deviceId DeviceId) (b x.RequestBody) {
	b.AddBlock(1).
		Write2b(uint16(deviceId)).
		WritePad(2).
		End()
	return
}

type XIListPropertiesReply struct {
	Properties []x.Atom
}

func readXIListPropertiesReply(r *x.Reader, v *XIListPropertiesReply) error {
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

	propertiesLen := int(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(22)
	if r.Err() != nil {
		return r.Err()
	}

	if propertiesLen > 0 {
		v.Properties = make([]x.Atom, propertiesLen)
		for i := 0; i < propertiesLen; i++ {
			v.Properties[i] = x.Atom(r.Read4b())
			if r.Err() != nil {
				return r.Err()
			}
		}
	}
	return nil
}

// #WREQ
func encodeXIChangeProperty(deviceId DeviceId, mode uint8, format uint8,
	property x.Atom, Type x.Atom, data []byte) (b x.RequestBody) {

	dataLen := len(data)
	b.AddBlock(4).
		Write2b(uint16(deviceId)).
		Write1b(mode).
		Write1b(format).
		Write4b(uint32(property)).
		Write4b(uint32(Type)).
		Write4b(uint32(dataLen / (int(format) / 8))).
		End()
	b.AddBytes(data)
	return
}

// #WREQ
func encodeXIDeleteProperty(deviceId DeviceId, property x.Atom) (b x.RequestBody) {
	b.AddBlock(2).
		Write2b(uint16(deviceId)).
		WritePad(2).
		Write4b(uint32(property)).
		End()
	return
}

// #WREQ
func encodeXIGetProperty(deviceId DeviceId, delete bool, property x.Atom,
	Type x.Atom, offset, len uint32) (b x.RequestBody) {

	b.AddBlock(5).
		Write2b(uint16(deviceId)).
		Write1b(x.BoolToUint8(delete)).
		WritePad(1).
		Write4b(uint32(property)).
		Write4b(uint32(Type)).
		Write4b(offset).
		Write4b(len).
		End()
	return
}

type XIGetPropertyReply struct {
	Type       x.Atom
	BytesAfter uint32
	NumItems   uint32
	Format     uint8
	Data       []byte
}

func readXIGetPropertyReply(r *x.Reader, v *XIGetPropertyReply) error {
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

	v.Type = x.Atom(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.BytesAfter = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NumItems = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Format = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// unused
	r.ReadPad(11)
	if r.Err() != nil {
		return r.Err()
	}

	n := int(v.NumItems) * int(v.Format/8)
	var err error
	v.Data, err = r.ReadBytes(n)
	if err != nil {
		return err
	}

	return nil
}
