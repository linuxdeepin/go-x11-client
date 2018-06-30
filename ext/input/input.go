package input

import (
	"github.com/linuxdeepin/go-x11-client"
)

func readFP1616(r *x.Reader) (float32, error) {
	value := int32(r.Read4b())
	ret := float32(value) / float32(1<<16)
	return ret, nil
}

func writeFP1616(w *x.Writer, value float32) {
	w.Write4b(uint32(int32(value * float32(1<<16))))
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
func writeXIQueryVersion(w *x.Writer, majorVersion, minorVersion uint16) {
	w.WritePad(4)
	w.Write2b(majorVersion)
	w.Write2b(minorVersion)
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
func writeXIQueryDevice(w *x.Writer, deviceId DeviceId) {
	w.WritePad(4)
	w.Write2b(uint16(deviceId))
	w.WritePad(2)
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

	v.Name = r.ReadString(nameLen)
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(x.Pad(nameLen))
	if r.Err() != nil {
		return r.Err()
	}

	// classes
	var err error
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

	v.Data = r.ReadBytes((int(v.Len) * 4) - 6)
	if r.Err() != nil {
		return r.Err()
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
func writeXISelectEvents(w *x.Writer, window x.Window, masks []EventMask) {
	w.WritePad(4)
	w.Write4b(uint32(window))

	w.Write2b(uint16(len(masks)))
	w.WritePad(2)

	for _, mask := range masks {
		writeEventMask(w, mask)
	}
}

type EventMask struct {
	DeviceId DeviceId
	Mask     []uint32
}

func writeEventMask(w *x.Writer, v EventMask) {
	w.Write2b(uint16(v.DeviceId))
	maskLen := len(v.Mask)
	w.Write2b(uint16(maskLen))
	for _, mask := range v.Mask {
		w.Write4b(mask)
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
func writeXIGetSelectedEvents(w *x.Writer, window x.Window) {
	w.WritePad(4)
	w.Write4b(uint32(window))
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
func writeXIQueryPointer(w *x.Writer, window x.Window, deviceId DeviceId) {
	w.WritePad(4)
	w.Write4b(uint32(window))
	w.Write2b(uint16(deviceId))
	w.WritePad(2)
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
func writeXIWarpPointer(w *x.Writer, srcWin, dstWin x.Window, srcX, srcY float32,
	srcWidth, srcHeight uint16, dstX, dstY float32, deviceId DeviceId) {

	w.WritePad(4)
	w.Write4b(uint32(srcWin))
	w.Write4b(uint32(dstWin))
	writeFP1616(w, srcX)
	writeFP1616(w, srcY)
	w.Write2b(srcWidth)
	w.Write2b(srcHeight)
	writeFP1616(w, dstX)
	writeFP1616(w, dstY)
	w.Write2b(uint16(deviceId))
	w.WritePad(2)
}

// #WREQ
func writeXIChangeCursor(w *x.Writer, window x.Window, cursor x.Cursor, deviceId DeviceId) {
	w.WritePad(4)
	w.Write4b(uint32(window))
	w.Write4b(uint32(cursor))
	w.Write2b(uint16(deviceId))
	w.WritePad(2)
}

// #WREQ
func writeXIChangeHierarchy(w *x.Writer, changes []HierarchyChange) {
	w.WritePad(4)
	w.Write1b(uint8(len(changes)))
	w.WritePad(3)
	for _, change := range changes {
		writeHierarchyChange(w, change)
	}
}

func writeHierarchyChange(w *x.Writer, v HierarchyChange) {
	w.Write2b(v.getHierarchyType())
	w.Write2b(v.getLen() + 1)
	v.writeTo(w)
}

type HierarchyChange interface {
	writeTo(w *x.Writer)
	getHierarchyType() uint16
	getLen() uint16 // Length in 4 byte units.
}

// size (1 + (len(Name) + pad(len(name))) / 4) * 4 byte
type AddMaster struct {
	SendCore bool
	Enable   bool
	Name     string
}

func (am *AddMaster) writeTo(w *x.Writer) {
	nameLen := len(am.Name)
	w.Write2b(uint16(nameLen))
	w.Write1b(x.BoolToUint8(am.SendCore))
	w.Write1b(x.BoolToUint8(am.Enable))

	w.WriteString(am.Name)
	w.WritePad(x.Pad(nameLen))
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

func (rm *RemoveMaster) writeTo(w *x.Writer) {
	w.Write2b(uint16(rm.DeviceId))
	w.Write1b(rm.ReturnMode)
	w.WritePad(1)

	w.Write2b(uint16(rm.ReturnPointer))
	w.Write2b(uint16(rm.ReturnKeyboard))
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

func (as *AttachSlave) writeTo(w *x.Writer) {
	w.Write2b(uint16(as.DeviceId))
	w.Write2b(uint16(as.Master))
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

func (ds *DetachSlave) writeTo(w *x.Writer) {
	w.Write2b(uint16(ds.DeviceId))
	w.WritePad(2)
}

func (*DetachSlave) getHierarchyType() uint16 {
	return HierarchyChangeTypeDetachSlave
}

func (ds *DetachSlave) getLen() uint16 {
	return 1
}

// #WREQ
func writeXISetClientPointer(w *x.Writer, window x.Window, deviceId DeviceId) {
	w.WritePad(4)
	w.Write4b(uint32(window))
	w.Write2b(uint16(deviceId))
	w.WritePad(2)
}

// #WREQ
func writeXIGetClientPointer(w *x.Writer, window x.Window) {
	w.WritePad(4)
	w.Write4b(uint32(window))
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
func writeXISetFocus(w *x.Writer, focus x.Window, time x.Timestamp, deviceId DeviceId) {
	w.WritePad(4)
	w.Write4b(uint32(focus))
	w.Write4b(uint32(time))
	w.Write2b(uint16(deviceId))
	w.WritePad(2)
}

// #WREQ
func writeXIGetFocus(w *x.Writer, deviceId DeviceId) {
	w.WritePad(4)
	w.Write2b(uint16(deviceId))
	w.WritePad(2)
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
func writeXIGrabDevice(w *x.Writer, window x.Window, time x.Timestamp, cursor x.Cursor,
	deviceId DeviceId, mode, pairedDeviceMode uint8, ownerEvents bool, masks []EventMask) {
	w.WritePad(4)
	w.Write4b(uint32(window))
	w.Write4b(uint32(time))
	w.Write4b(uint32(cursor))

	w.Write2b(uint16(deviceId))
	w.Write1b(mode)
	w.Write1b(pairedDeviceMode)

	w.Write1b(x.BoolToUint8(ownerEvents))
	w.WritePad(1)
	w.Write2b(uint16(len(masks)))

	for _, mask := range masks {
		writeEventMask(w, mask)
	}
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
func writeXIUngrabDevice(w *x.Writer, time x.Timestamp, deviceId DeviceId) {
	w.WritePad(4)
	w.Write4b(uint32(time))
	w.Write2b(uint16(deviceId))
	w.WritePad(2)
}

// #WREQ
func writeXIAllowEvents(w *x.Writer, time x.Timestamp, deviceId DeviceId, eventMode uint8,
	touchId uint32, grabWindow x.Window) {
	w.WritePad(4)
	w.Write4b(uint32(time))

	w.Write2b(uint16(deviceId))
	w.Write1b(eventMode)
	w.WritePad(1)

	w.Write4b(touchId)
	w.Write4b(uint32(grabWindow))
}

// #WREQ
func writeXIPassiveGrabDevice(w *x.Writer, grabWindow x.Window, cursor x.Cursor,
	detail uint32, deviceId DeviceId, grabType, grabMode, pairedDeviceMode uint8,
	ownerEvents bool, masks []uint32, modifiers []uint32) {
	w.WritePad(4)
	w.WritePad(4) // unused time field
	w.Write4b(uint32(grabWindow))
	w.Write4b(uint32(cursor))
	w.Write4b(detail)

	w.Write2b(uint16(deviceId))
	w.Write2b(uint16(len(modifiers)))

	w.Write2b(uint16(len(masks)))
	w.Write1b(grabType)
	w.Write1b(grabMode)

	w.Write1b(pairedDeviceMode)
	w.Write1b(x.BoolToUint8(ownerEvents))
	w.WritePad(2)

	for _, mask := range masks {
		w.Write4b(mask)
	}

	for _, mod := range modifiers {
		w.Write4b(mod)
	}
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
func writeXIPassiveUngrabDevice(w *x.Writer, grabWindow x.Window, detail uint32,
	deviceId DeviceId, grabType uint8, modifiers []uint32) {
	w.WritePad(4)
	w.Write4b(uint32(grabWindow))
	w.Write4b(detail)

	w.Write2b(uint16(deviceId))
	w.Write2b(uint16(len(modifiers)))

	w.Write1b(grabType)
	w.WritePad(3)

	for _, mod := range modifiers {
		w.Write4b(mod)
	}
}

// #WREQ
func writeXIListProperties(w *x.Writer, deviceId DeviceId) {
	w.WritePad(4)
	w.Write2b(uint16(deviceId))
	w.WritePad(2)
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
func writeXIChangeProperty(w *x.Writer, deviceId DeviceId, mode uint8, format uint8,
	property x.Atom, Type x.Atom, data []byte) {

	w.WritePad(4)

	w.Write2b(uint16(deviceId))
	w.Write1b(mode)
	w.Write1b(format)

	w.Write4b(uint32(property))
	w.Write4b(uint32(Type))

	dataLen := len(data)
	w.Write4b(uint32(dataLen / (int(format) / 8)))
	w.WriteBytes(data)
	w.WritePad(x.Pad(dataLen))
}

// #WREQ
func writeXIDeleteProperty(w *x.Writer, deviceId DeviceId, property x.Atom) {

	w.WritePad(4)

	w.Write2b(uint16(deviceId))
	w.WritePad(2)

	w.Write4b(uint32(property))
}

// #WREQ
func writeXIGetProperty(w *x.Writer, deviceId DeviceId, delete bool, property x.Atom,
	Type x.Atom, offset, len uint32) {
	w.WritePad(4)

	w.Write2b(uint16(deviceId))
	w.Write1b(x.BoolToUint8(delete))
	w.WritePad(1)

	w.Write4b(uint32(property))
	w.Write4b(uint32(Type))
	w.Write4b(offset)
	w.Write4b(len)
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
	v.Data = r.ReadBytes(n)
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
