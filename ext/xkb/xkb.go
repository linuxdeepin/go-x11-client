// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package xkb

import (
	"bytes"
	"math"
	"sort"

	"github.com/linuxdeepin/go-x11-client"
)

// #WREQ
func encodeUseExtension(wantedMajor, wantedMinor uint16) (b x.RequestBody) {
	b.AddBlock(1).
		Write2b(wantedMajor).
		Write2b(wantedMinor).
		End()
	return
}

type UseExtensionReply struct {
	Supported   bool
	ServerMajor uint16
	ServerMinor uint16
}

func readUseExtensionReply(r *x.Reader, v *UseExtensionReply) error {
	if !r.RemainAtLeast4b(3) {
		return x.ErrDataLenShort
	}
	supported, _ := r.ReadReplyHeader()
	v.Supported = x.Uint8ToBool(supported)

	v.ServerMajor = r.Read2b()
	v.ServerMinor = r.Read2b() // 3

	return nil
}

// #WREQ vararg:options
func encodeSelectEvents(deviceSpec DeviceSpec, options ...SelectOption) (b x.RequestBody) {
	var affectWhich, clear, selectAll, affectMap, mapDetail uint16
	w := x.NewWriter()
	sort.Sort(selectOptions(options))
	for _, option := range options {
		affectWhich |= option.affectEvent()
		switch o := option.(type) {
		case selectOptionAll:
			selectAll |= o.eventType

		case selectOptionClear:
			clear |= o.eventType

		case selectOptionDetail:
			affect, detail := o.getAffectDetail()
			switch o.eventType {
			case EventTypeNewKeyboardNotify, EventTypeStateNotify, EventTypeNamesNotify,
				EventTypeAccessXNotify, EventTypeExtensionDeviceNotify:
				w.Write2b(uint16(affect))
				w.Write2b(uint16(detail))

			case EventTypeControlsNotify, EventTypeIndicatorStateNotify,
				EventTypeIndicatorMapNotify:
				w.Write4b(uint32(affect))
				w.Write4b(uint32(detail))

			case EventTypeCompatMapNotify, EventTypeBellNotify, EventTypeActionMessage:
				w.Write1b(uint8(affect))
				w.Write1b(uint8(detail))

			case EventTypeMapNotify:
				affectMap = uint16(affect)
				mapDetail = uint16(detail)
			}
		}
	}

	b.AddBlock(3).
		Write2b(uint16(deviceSpec)).
		Write2b(affectWhich).
		Write2b(clear).
		Write2b(selectAll).
		Write2b(affectMap).
		Write2b(mapDetail).End()
	b.AddBytes(w.Bytes())
	return
}

type selectOptions []SelectOption

func (v selectOptions) Len() int {
	return len(v)
}

func (v selectOptions) Less(i, j int) bool {
	return v[i].affectEvent() < v[j].affectEvent()
}

func (v selectOptions) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

type SelectOption interface {
	affectEvent() uint16
}

type selectOptionClear struct {
	eventType uint16
}

func (v selectOptionClear) affectEvent() uint16 {
	return v.eventType
}

type selectOptionAll struct {
	eventType uint16
}

func (v selectOptionAll) affectEvent() uint16 {
	return v.eventType
}

type selectOptionDetail struct {
	eventType uint16
	detail    map[uint]bool
}

func (v selectOptionDetail) affectEvent() uint16 {
	return v.eventType
}

func (v selectOptionDetail) getAffectDetail() (affect, detail uint) {
	for key, value := range v.detail {
		affect |= key
		if value {
			detail |= key
		}
	}
	return
}

func SelectAll(eventType uint16) SelectOption {
	return selectOptionAll{eventType}
}

func SelectDetail(eventType uint16, detail map[uint]bool) SelectOption {
	return selectOptionDetail{eventType: eventType, detail: detail}
}

func SelectClear(eventType uint16) SelectOption {
	return selectOptionClear{eventType}
}

// #WREQ
func encodeBell(deviceSpec DeviceSpec, bellClass BellClassSpec, bellID IdSpec,
	percent int8, forceSound, eventOnly bool, pitch, duration int16,
	name x.Atom, window x.Window) (b x.RequestBody) {

	b.AddBlock(6).
		Write2b(uint16(deviceSpec)).
		Write2b(uint16(bellClass)). // 1
		Write2b(uint16(bellID)).
		Write1b(uint8(percent)).
		WriteBool(forceSound). // 2
		WriteBool(eventOnly).
		WritePad(1).
		Write2b(uint16(pitch)). // 3
		Write2b(uint16(duration)).
		WritePad(2). // 4
		Write4b(uint32(name)).
		Write4b(uint32(window)).
		End()
	return
}

// #WREQ
func encodeGetState(deviceSpec DeviceSpec) (b x.RequestBody) {
	b.AddBlock(1).
		Write2b(uint16(deviceSpec)).
		WritePad(2).
		End()
	return
}

type GetStateReply struct {
	DeviceID         uint8
	Mods             uint8
	BaseMods         uint8
	LatchedMods      uint8
	LockedMods       uint8
	Group            uint8
	LockedGroup      uint8
	BaseGroup        int16
	LatchedGroup     int16
	CompatState      uint8
	GrabMods         uint8
	CompatGrabMods   uint8
	LookupMods       uint8
	CompatLookupMods uint8
	PtrBtnState      uint16
}

func readGetStateReply(r *x.Reader, v *GetStateReply) error {
	if !r.RemainAtLeast4b(7) {
		return x.ErrDataLenShort
	}
	v.DeviceID, _ = r.ReadReplyHeader()

	v.Mods = r.Read1b()
	v.BaseMods = r.Read1b()
	v.LatchedMods = r.Read1b()
	v.LockedMods = r.Read1b() // 3

	v.Group = r.Read1b()
	v.LockedGroup = r.Read1b()
	v.BaseGroup = int16(r.Read2b())

	v.LatchedGroup = int16(r.Read2b())
	v.CompatState = r.Read1b()
	v.GrabMods = r.Read1b() // 5

	v.CompatGrabMods = r.Read1b()
	v.LookupMods = r.Read1b()
	v.CompatLookupMods = r.Read1b()
	r.ReadPad(1) // 6

	v.PtrBtnState = r.Read2b() // 7

	return nil
}

// #WREQ
func encodeLatchLockState(deviceSpec DeviceSpec, affectModLocks, modLocks uint8,
	lockGroup bool, groupLock, affectModLatches, modLatches uint8, latchGroup bool,
	groupLatch uint16) (b x.RequestBody) {

	b.AddBlock(3).
		Write2b(uint16(deviceSpec)).
		Write1b(affectModLocks).
		Write1b(modLocks). // 1
		WriteBool(lockGroup).
		Write1b(groupLock).
		Write1b(affectModLatches).
		Write1b(modLatches). // 2
		WritePad(1).
		WriteBool(latchGroup).
		Write2b(groupLatch). // 3
		End()
	return
}

// #WREQ
func encodeGetControls(deviceSpec DeviceSpec) (b x.RequestBody) {
	b.AddBlock(1).
		Write2b(uint16(deviceSpec)).
		WritePad(2).
		End()
	return
}

type GetControlsReply struct {
	DeviceID uint8
	Controls
}

func readGetControlsReply(r *x.Reader, v *GetControlsReply) error {
	if !r.RemainAtLeast4b(15) {
		return x.ErrDataLenShort
	}
	v.DeviceID, _ = r.ReadReplyHeader()

	v.MouseKeysDefaultBtn = r.Read1b()
	v.NumGroups = r.Read1b()
	v.GroupsWrap = r.Read1b()
	v.Internal.Mask = r.Read1b() // 3

	v.IgnoreLock.Mask = r.Read1b()
	v.Internal.RealMods = r.Read1b()
	v.IgnoreLock.RealMods = r.Read1b()
	r.ReadPad(1) // 4

	v.Internal.VirtualMods = r.Read2b()
	v.IgnoreLock.VirtualMods = r.Read2b() // 5

	v.RepeatDelay = r.Read2b()
	v.RepeatInterval = r.Read2b()

	v.SlowKeysDelay = r.Read2b()
	v.DebounceDelay = r.Read2b()

	v.MouseKeysDelay = r.Read2b()
	v.MouseKeysInterval = r.Read2b() // 8

	v.MouseKeysTimeToMax = r.Read2b()
	v.MouseKeysMaxSpeed = r.Read2b()

	v.MouseKeysCurve = int16(r.Read2b())
	v.AccessXOption = r.Read2b() // 10

	v.AccessXTimeout = r.Read2b()
	v.AccessXTimeoutOptionsMask = r.Read2b()

	v.AccessXTimeoutOptionsValues = r.Read2b()
	r.ReadPad(2) // 12

	v.AccessXTimeoutMask = r.Read4b()

	v.AccessXTimeoutValues = r.Read4b()

	v.EnabledControls = r.Read4b() // 15

	var err error
	v.PerKeyRepeat, err = r.ReadBytes(ConstPerKeyBitArraySize)
	return err
}

// size: 1 * 4b
type ModDef struct {
	// RealMods | VirtualMods mapped to real modifiers
	Mask uint8
	//  real modifier bits
	RealMods uint8
	// virtual modifier bits
	VirtualMods uint16
}

func readModDef(r *x.Reader) (v ModDef) {
	v.Mask = r.Read1b()
	v.RealMods = r.Read1b()
	v.VirtualMods = r.Read2b()
	return v
}

func writeModDef(b *x.FixedSizeBuf, v ModDef) {
	b.Write1b(v.Mask).
		Write1b(v.RealMods).
		Write2b(v.VirtualMods)
}

type Controls struct {
	MouseKeysDefaultBtn         uint8
	NumGroups                   uint8
	GroupsWrap                  uint8
	Internal                    ModDef
	IgnoreLock                  ModDef
	EnabledControls             uint32
	RepeatDelay                 uint16
	RepeatInterval              uint16
	SlowKeysDelay               uint16
	DebounceDelay               uint16
	MouseKeysDelay              uint16
	MouseKeysInterval           uint16
	MouseKeysTimeToMax          uint16
	MouseKeysMaxSpeed           uint16
	MouseKeysCurve              int16
	AccessXOption               uint16
	AccessXTimeout              uint16
	AccessXTimeoutOptionsMask   uint16
	AccessXTimeoutOptionsValues uint16

	AccessXTimeoutMask   uint32
	AccessXTimeoutValues uint32
	PerKeyRepeat         []byte // len 32
}

// #WREQ
func encodeSetControls(deviceSpec DeviceSpec,
	affectInternalRealMods, affectIgnoreLockRealMods uint8,
	affectInternalVirtualMods, affectIgnoreLockVirtualMods uint16,
	affectEnabledControls, changeControls uint32, c *Controls) (b x.RequestBody) {

	b.AddBlock(16).
		Write2b(uint16(deviceSpec)).
		Write1b(affectInternalRealMods).
		Write1b(c.Internal.RealMods). // 1
		Write1b(affectIgnoreLockRealMods).
		Write1b(c.IgnoreLock.RealMods).
		Write2b(affectInternalVirtualMods). // 2
		Write2b(c.Internal.VirtualMods).
		Write2b(affectIgnoreLockVirtualMods). // 3
		Write2b(c.IgnoreLock.VirtualMods).
		Write1b(c.MouseKeysDefaultBtn).
		Write1b(c.GroupsWrap). // 4
		Write2b(c.AccessXOption).
		WritePad(2). // 5
		Write4b(affectEnabledControls).
		Write4b(c.EnabledControls).
		Write4b(changeControls). // 8
		Write2b(c.RepeatDelay).
		Write2b(c.RepeatInterval). // 9
		Write2b(c.SlowKeysDelay).
		Write2b(c.DebounceDelay).
		Write2b(c.MouseKeysDelay).
		Write2b(c.MouseKeysInterval).
		Write2b(c.MouseKeysTimeToMax).
		Write2b(c.MouseKeysMaxSpeed).
		Write2b(uint16(c.MouseKeysCurve)).
		Write2b(c.AccessXTimeout). // 13
		Write4b(c.AccessXTimeoutMask).
		Write4b(c.AccessXTimeoutValues).
		Write2b(c.AccessXTimeoutOptionsMask).
		Write2b(c.AccessXTimeoutOptionsValues). // 16
		End()
	b.AddBytes(c.PerKeyRepeat)
	return
}

type GetMapPartOptions struct {
	Partial   uint16
	FirstType uint8
	NTypes    uint8

	FirstKeySym x.Keycode
	NKeySyms    uint8

	FirstKeyAction x.Keycode
	NKeyActions    uint8

	FirstKeyBehavior x.Keycode
	NKeyBehavior     uint8

	VirtualMods uint16

	FirstKeyExplicit x.Keycode
	NKeyExplicit     uint8

	FirstModMapKey x.Keycode
	NModMapKeys    uint8

	FirstVModMapKey x.Keycode
	NVModMapKeys    uint8
}

// #WREQ
func encodeGetMap(deviceSpec DeviceSpec, full uint16, opt *GetMapPartOptions) (b x.RequestBody) {
	if opt == nil {
		opt = &GetMapPartOptions{}
	} else {
		if opt.NTypes != 0 || opt.FirstType != 0 {
			opt.Partial |= MapPartKeyTypes
		}

		if opt.NKeySyms != 0 || opt.FirstKeySym != 0 {
			opt.Partial |= MapPartKeySyms
		}

		if opt.NKeyActions != 0 || opt.FirstKeyAction != 0 {
			opt.Partial |= MapPartKeyActions
		}

		if opt.NKeyBehavior != 0 || opt.FirstKeyBehavior != 0 {
			opt.Partial |= MapPartKeyBehaviors
		}

		if opt.NKeyExplicit != 0 || opt.FirstKeyExplicit != 0 {
			opt.Partial |= MapPartExplicitComponents
		}

		if opt.NModMapKeys != 0 || opt.FirstModMapKey != 0 {
			opt.Partial |= MapPartModifierMap
		}

		if opt.NVModMapKeys != 0 || opt.FirstVModMapKey != 0 {
			opt.Partial |= MapPartVirtualModMap
		}

		if opt.VirtualMods != 0 {
			opt.Partial |= MapPartVirtualMods
		}
	}

	b.AddBlock(6).
		Write2b(uint16(deviceSpec)).
		Write2b(full).
		Write2b(opt.Partial).
		Write1b(opt.FirstType).
		Write1b(opt.NTypes). // 2
		Write1b(uint8(opt.FirstKeySym)).
		Write1b(opt.NKeySyms).
		Write1b(uint8(opt.FirstKeyAction)).
		Write1b(opt.NKeyActions).
		Write1b(uint8(opt.FirstKeyBehavior)).
		Write1b(opt.NKeyBehavior).
		Write2b(opt.VirtualMods). // 4
		Write1b(uint8(opt.FirstKeyExplicit)).
		Write1b(opt.NKeyExplicit).
		Write1b(uint8(opt.FirstModMapKey)).
		Write1b(opt.NVModMapKeys).
		Write1b(uint8(opt.FirstVModMapKey)).
		Write1b(opt.NVModMapKeys).
		WritePad(2). // 6
		End()
	return
}

type GetMapReply struct {
	DeviceID   uint8
	MinKeyCode x.Keycode
	MaxKeyCode x.Keycode
	Present    uint16

	FirstType  uint8
	NTypes     uint8
	TotalTypes uint8

	FirstKeySym x.Keycode
	TotalSyms   uint16
	NKeySyms    uint8

	FirstKeyAction x.Keycode
	TotalActions   uint16
	NKeyActions    uint8

	FirstKeyBehavior  x.Keycode
	NKeyBehaviors     uint8
	TotalKeyBehaviors uint8

	FirstKeyExplicit x.Keycode
	NKeyExplicit     uint8
	TotalKeyExplicit uint8

	FirstModMapKey  x.Keycode
	NModMapKeys     uint8
	TotalModMapKeys uint8

	FirstVModMapKey  x.Keycode
	NVModMapKeys     uint8
	TotalVModMapKeys uint8

	VirtualMods uint16

	Types        []KeyType
	Syms         []KeySymMap
	ActionsCount []uint8
	Actions      []Action
	Behaviors    []SetBehavior
	VMods        []uint8
	Explicit     []SetExplicit
	ModMap       []KeyModMap
	VModMap      []KeyVModMap
}

type KeyType struct {
	ModsMask    uint8
	ModsMods    uint8
	ModsVMods   uint16
	NumLevels   uint8
	NMapEntries uint8
	HasPreserve bool
	Map         []KTMapEntry
	Preserve    []ModDef
}

type SetKeyType struct {
	Mask            uint8
	RealMods        uint8
	VirtualMods     uint16
	NumLevels       uint8
	NMapEntries     uint8
	Preserve        bool
	Entries         []KTSetMapEntry
	PreserveEntries []KTSetMapEntry
}

func writeSetKeyType(w *x.FixedSizeBuf, v *SetKeyType) {
	w.Write1b(v.Mask).
		Write1b(v.RealMods).
		Write2b(v.VirtualMods).
		Write1b(v.NumLevels).
		Write1b(v.NMapEntries).
		WriteBool(v.Preserve).
		WritePad(1) // 2
	for i := 0; i < int(v.NMapEntries); i++ {
		writeKTSetMapEntry(w, v.Entries[i]) // v.NMapEntries
	}
	if v.Preserve {
		for i := 0; i < int(v.NMapEntries); i++ {
			writeKTSetMapEntry(w, v.PreserveEntries[i])
		}
	}
}

func (v *SetKeyType) sizeIn4b() int {
	n := 4 + int(v.NMapEntries)
	if v.Preserve {
		n += int(v.NMapEntries)
	}
	return n
}

func readKeyType(r *x.Reader, v *KeyType) error {
	if !r.RemainAtLeast4b(2) {
		return x.ErrDataLenShort
	}
	v.ModsMask = r.Read1b()
	v.ModsMods = r.Read1b()
	v.ModsVMods = r.Read2b()

	v.NumLevels = r.Read1b()
	v.NMapEntries = r.Read1b()
	v.HasPreserve = r.ReadBool()
	r.ReadPad(1) // 2

	if v.NMapEntries > 0 {
		if !r.RemainAtLeast4b(3 * int(v.NMapEntries)) {
			return x.ErrDataLenShort
		}
		v.Map = make([]KTMapEntry, v.NMapEntries)
		for i := 0; i < int(v.NMapEntries); i++ {
			readKTMapEntry(r, &v.Map[i])
		}

		if v.HasPreserve {
			v.Preserve = make([]ModDef, v.NMapEntries)
			for i := 0; i < int(v.NMapEntries); i++ {
				v.Preserve[i] = readModDef(r)
			}
		}
	}
	return nil
}

// size: 1 * 4b
type KTSetMapEntry struct {
	Level       uint8
	RealMods    uint8
	VirtualMods uint16
}

func writeKTSetMapEntry(w *x.FixedSizeBuf, v KTSetMapEntry) {
	w.Write1b(v.Level).
		Write1b(v.RealMods).
		Write2b(v.VirtualMods)
}

// key type map entry, size: 2 * 4b
type KTMapEntry struct {
	Active    bool
	ModsMask  uint8
	Level     uint8
	ModsMods  uint8
	ModsVMods uint16
}

func readKTMapEntry(r *x.Reader, v *KTMapEntry) {
	v.Active = r.ReadBool()
	v.ModsMask = r.Read1b()
	v.Level = r.Read1b()
	v.ModsMods = r.Read1b()

	v.ModsVMods = r.Read2b()
	r.ReadPad(2) // 2
}

type KeySymMap struct {
	KtIndex   []uint8 // length 4
	GroupInfo uint8
	Width     uint8
	NSyms     uint16
	Syms      []x.Keysym
}

func readKeySymMap(r *x.Reader, v *KeySymMap) error {
	if !r.RemainAtLeast4b(2) {
		return x.ErrDataLenShort
	}
	v.KtIndex = r.MustReadBytes(4)

	v.GroupInfo = r.Read1b()
	v.Width = r.Read1b()
	v.NSyms = r.Read2b() // 2

	if v.NSyms > 0 {
		if !r.RemainAtLeast4b(int(v.NSyms)) {
			return x.ErrDataLenShort
		}
		v.Syms = make([]x.Keysym, v.NSyms)
		for i := 0; i < int(v.NSyms); i++ {
			v.Syms[i] = x.Keysym(r.Read4b())
		}
	}
	return nil
}

func writeKeySymMap(b *x.FixedSizeBuf, v *KeySymMap) {
	b.WriteBytes(v.KtIndex).
		Write1b(v.GroupInfo).
		Write1b(v.Width).
		Write2b(v.NSyms)
	for _, sym := range v.Syms {
		b.Write4b(uint32(sym))
	}
}

func (v *KeySymMap) sizeIn4b() int {
	return 2 + len(v.Syms)
}

// size: 2 * 4b
type Action interface {
	Type() uint8 // TODO: rename ActionType
	writeTo(b *x.FixedSizeBuf)
}

var readActionFuncArray = [...]func(r *x.Reader) Action{
	readSANoAction, // 0
	readSASetMods,
	readSALatchMods,
	readSALockMods,
	readSASetGroup,
	readSALatchGroup,
	readSALockGroup, // 6
	readSAMovePtr,
	readSAPtrBtn,
	readSALockPtrBtn,
	readSASetPtrDflt, // 10
	readSAIsoLock,
	readSATerminate,
	readSASwitchScreen,
	readSASetControls,
	readSALockControls, // 15
	readSAActionMessage,
	readSARedirectKey,
	readSADeviceBtn,
	readSALockDeviceBtn,
	readSADeviceValuator, // 20
}

func readAction(r *x.Reader) Action {
	type0 := r.Read1b()
	data := r.MustReadBytes(7) // 2

	dataR := x.NewReaderFromData(data)
	if type0 <= SATypeDeviceValuator {
		action := readActionFuncArray[type0](dataR)
		return action
	}
	// out of range
	return UnknownAction{type0: type0, data: data}
}

func writeAction(b *x.FixedSizeBuf, act Action) {
	b.Write1b(act.Type())
	act.writeTo(b)
}

// Action NoAction
type SANoAction struct {
}

func (SANoAction) Type() uint8 {
	return SATypeNoAction
}

func (SANoAction) writeTo(b *x.FixedSizeBuf) {
	b.WritePad(7)
}

func readSANoAction(r *x.Reader) Action {
	return SANoAction{}
}

// Action SetMods
type SASetMods struct {
	Flags    uint8
	Mask     uint8
	RealMods uint8
	VMods    uint16
}

func (SASetMods) Type() uint8 {
	return SATypeSetMods
}

func readSASetMods(r *x.Reader) Action {
	var v SASetMods
	readSASetModsAux(r, &v)
	return v
}

func readSASetModsAux(r *x.Reader, v *SASetMods) {
	v.Flags = r.Read1b()
	v.Mask = r.Read1b()
	v.RealMods = r.Read1b()
	vModsHigh := r.Read1b()
	vModsLow := r.Read1b()
	v.VMods = uint16(vModsHigh)<<8 | uint16(vModsLow)
}

func (v SASetMods) writeTo(b *x.FixedSizeBuf) {
	vmodsHigh := uint8(v.VMods >> 8)
	vmodsLow := uint8(v.VMods)
	b.Write1b(v.Flags).
		Write1b(v.Mask).
		Write1b(v.RealMods).
		Write1b(vmodsHigh).
		Write1b(vmodsLow). // 5b
		WritePad(2)
}

// Action LatchMods
type SALatchMods SASetMods

func (SALatchMods) Type() uint8 {
	return SATypeLatchMods
}

func readSALatchMods(r *x.Reader) Action {
	var v SALatchMods
	readSASetModsAux(r, (*SASetMods)(&v))
	return v
}

func (v SALatchMods) writeTo(b *x.FixedSizeBuf) {
	SASetMods(v).writeTo(b)
}

// Action LockMods
type SALockMods SASetMods

func (SALockMods) Type() uint8 {
	return SATypeLockMods
}

func readSALockMods(r *x.Reader) Action {
	var v SALockMods
	readSASetModsAux(r, (*SASetMods)(&v))
	return v
}

func (v SALockMods) writeTo(b *x.FixedSizeBuf) {
	SASetMods(v).writeTo(b)
}

// Action SetGroup
type SASetGroup struct {
	Flags uint8
	Group int8
}

func (SASetGroup) Type() uint8 {
	return SATypeSetGroup
}

func readSASetGroup(r *x.Reader) Action {
	var v SASetGroup
	readSASetGroupAux(r, &v)
	return v
}

func readSASetGroupAux(r *x.Reader, v *SASetGroup) {
	v.Flags = r.Read1b()
	v.Group = int8(r.Read1b())
}

func (v SASetGroup) writeTo(b *x.FixedSizeBuf) {
	b.Write1b(v.Flags).
		Write1b(uint8(v.Group)).
		WritePad(5)
}

// Action LatchGroup
type SALatchGroup SASetGroup

func (SALatchGroup) Type() uint8 {
	return SATypeLatchGroup
}

func readSALatchGroup(r *x.Reader) Action {
	var v SALatchGroup
	readSASetGroupAux(r, (*SASetGroup)(&v))
	return v
}

func (v SALatchGroup) writeTo(b *x.FixedSizeBuf) {
	SASetGroup(v).writeTo(b)
}

// Action LockGroup
type SALockGroup SASetGroup

func (SALockGroup) Type() uint8 {
	return SATypeLockGroup
}

func readSALockGroup(r *x.Reader) Action {
	var v SALockGroup
	readSASetGroupAux(r, (*SASetGroup)(&v))
	return v
}

func (v SALockGroup) writeTo(b *x.FixedSizeBuf) {
	SASetGroup(v).writeTo(b)
}

// Action MovePtr
type SAMovePtr struct {
	Flags uint8
	X     int16
	Y     int16
}

func (SAMovePtr) Type() uint8 {
	return SATypeMovePtr
}

func readSAMovePtr(r *x.Reader) Action {
	var v SAMovePtr
	v.Flags = r.Read1b()

	xHigh := int8(r.Read1b())
	xLow := r.Read1b()
	v.X = int16(xHigh)<<8 | int16(xLow)

	yHigh := int8(r.Read1b())
	yLow := r.Read1b()
	v.Y = int16(yHigh)<<8 | int16(yLow)
	return v
}

func (v SAMovePtr) writeTo(b *x.FixedSizeBuf) {
	xHigh := int8(v.X >> 8)
	xLow := uint8(v.X)

	yHigh := int8(v.Y >> 8)
	yLow := uint8(v.Y)

	b.Write1b(v.Flags).
		Write1b(uint8(xHigh)).
		Write1b(xLow).
		Write1b(uint8(yHigh)).
		Write1b(yLow). // 5
		WritePad(2)
}

// Action PtrBtn
type SAPtrBtn struct {
	Flags  uint8
	Count  uint8
	Button uint8
}

func (SAPtrBtn) Type() uint8 {
	return SATypePtrBtn
}

func readSAPtrBtn(r *x.Reader) Action {
	var v SAPtrBtn
	v.Flags = r.Read1b()
	v.Count = r.Read1b()
	v.Button = r.Read1b()
	return v
}

func (v SAPtrBtn) writeTo(b *x.FixedSizeBuf) {
	b.Write1b(v.Flags).
		Write1b(v.Count).
		Write1b(v.Button).
		WritePad(4)
}

// Action LockPtrBtn
type SALockPtrBtn struct {
	Flags  uint8
	Button uint8
}

func (SALockPtrBtn) Type() uint8 {
	return SATypeLockPtrBtn
}

func readSALockPtrBtn(r *x.Reader) Action {
	var v SALockPtrBtn
	v.Flags = r.Read1b()
	r.ReadPad(1)
	v.Button = r.Read1b()
	return v
}

func (v SALockPtrBtn) writeTo(b *x.FixedSizeBuf) {
	b.Write1b(v.Flags).
		Write1b(v.Button).
		WritePad(5)
}

// Action SetPtrDflt
type SASetPtrDflt struct {
	Flags  uint8
	Affect uint8
	Value  int8
}

func (SASetPtrDflt) Type() uint8 {
	return SATypeSetPtrDflt
}

func readSASetPtrDflt(r *x.Reader) Action {
	var v SASetPtrDflt
	v.Flags = r.Read1b()
	v.Affect = r.Read1b()
	v.Value = int8(r.Read1b())
	return v
}

func (v SASetPtrDflt) writeTo(b *x.FixedSizeBuf) {
	b.Write1b(v.Flags).
		Write1b(v.Affect).
		Write1b(uint8(v.Value)).
		WritePad(4)
}

// Action Iso Lock
type SAIsoLock struct {
	Flags    uint8
	Mask     uint8
	RealMods uint8
	Group    int8
	Affect   uint8
	VMods    uint16
}

func (SAIsoLock) Type() uint8 {
	return SATypeIsoLock
}

func readSAIsoLock(r *x.Reader) Action {
	var v SAIsoLock
	v.Flags = r.Read1b()
	v.Mask = r.Read1b()
	v.RealMods = r.Read1b()
	v.Group = int8(r.Read1b())
	v.Affect = r.Read1b()

	vModsHigh := r.Read1b()
	vModsLow := r.Read1b()
	v.VMods = uint16(vModsHigh)<<8 | uint16(vModsLow)
	return v
}

func (v SAIsoLock) writeTo(b *x.FixedSizeBuf) {
	vModsHigh := uint8(v.VMods >> 8)
	vModsLow := uint8(v.VMods)
	b.Write1b(v.Flags).
		Write1b(v.Mask).
		Write1b(v.RealMods).
		Write1b(uint8(v.Group)).
		Write1b(v.Affect).
		Write1b(vModsHigh).
		Write1b(vModsLow) // 7b
}

// Action Terminate
type SATerminate struct {
}

func (SATerminate) Type() uint8 {
	return SATypeTerminate
}

func readSATerminate(r *x.Reader) Action {
	return SATerminate{}
}

func (SATerminate) writeTo(b *x.FixedSizeBuf) {
	b.WritePad(7)
}

// Action Switch Screen
type SASwitchScreen struct {
	Flags     uint8
	NewScreen int8
}

func (SASwitchScreen) Type() uint8 {
	return SATypeSwitchScreen
}

func readSASwitchScreen(r *x.Reader) Action {
	var v SASwitchScreen
	v.Flags = r.Read1b()
	v.NewScreen = int8(r.Read1b())
	return v
}

func (v SASwitchScreen) writeTo(b *x.FixedSizeBuf) {
	b.Write1b(v.Flags).
		Write1b(uint8(v.NewScreen)).
		WritePad(5)
}

// Action SetControls
type SASetControls struct {
	BoolCtrls uint16
}

func (SASetControls) Type() uint8 {
	return SATypeSetControls
}

func readSASetControls(r *x.Reader) Action {
	var v SASetControls
	readSASetControlsAux(r, &v)
	return v
}

func readSASetControlsAux(r *x.Reader, v *SASetControls) {
	r.ReadPad(3)
	boolCtrlsHigh := r.Read1b()
	boolCtrlsLow := r.Read1b()
	v.BoolCtrls = uint16(boolCtrlsHigh)<<8 | uint16(boolCtrlsLow)
}

func (v SASetControls) writeTo(b *x.FixedSizeBuf) {
	boolCtrlsHigh := uint8(v.BoolCtrls >> 8)
	boolCtrlsLow := uint8(v.BoolCtrls)
	b.WritePad(3).
		Write1b(boolCtrlsHigh).
		Write1b(boolCtrlsLow).
		WritePad(2)
}

// Action LockControls
type SALockControls SASetControls

func (SALockControls) Type() uint8 {
	return SATypeLockControls
}

func readSALockControls(r *x.Reader) Action {
	var v SALockControls
	readSASetControlsAux(r, (*SASetControls)(&v))
	return v
}

func (v SALockControls) writeTo(b *x.FixedSizeBuf) {
	SASetControls(v).writeTo(b)
}

// Action ActionMessage
type SAActionMessage struct {
	Flags   uint8
	Message []byte // length 6
}

func (SAActionMessage) Type() uint8 {
	return SATypeActionMessage
}

func readSAActionMessage(r *x.Reader) Action {
	var v SAActionMessage
	v.Flags = r.Read1b()
	v.Message = r.MustReadBytes(6)
	return v
}

func (v SAActionMessage) writeTo(b *x.FixedSizeBuf) {
	if len(v.Message) != 6 {
		panic("length of message is not 6")
	}
	b.Write1b(v.Flags).
		WriteBytes(v.Message)
}

// Action RedirectKey
type SARedirectKey struct {
	NewKey        x.Keycode
	Mask          uint8
	RealModifiers uint8
	VModsMask     uint16
	VMods         uint16
}

func (SARedirectKey) Type() uint8 {
	return SATypeRedirectKey
}

func readSARedirectKey(r *x.Reader) Action {
	var v SARedirectKey
	v.NewKey = x.Keycode(r.Read1b())
	v.Mask = r.Read1b()
	v.RealModifiers = r.Read1b()
	vModsMaskHigh := r.Read1b()
	vModsMaskLow := r.Read1b()
	v.VModsMask = uint16(vModsMaskHigh)<<8 | uint16(vModsMaskLow)
	vModsHigh := r.Read1b()
	vModsLow := r.Read1b()
	v.VMods = uint16(vModsHigh)<<8 | uint16(vModsLow)
	return v
}

func (v SARedirectKey) writeTo(b *x.FixedSizeBuf) {
	vModsMaskHigh := uint8(v.VModsMask >> 8)
	vModsMaskLow := uint8(v.VModsMask)

	vModsHigh := uint8(v.VMods >> 8)
	vModsLow := uint8(v.VMods)

	b.Write1b(uint8(v.NewKey)).
		Write1b(v.Mask).
		Write1b(v.RealModifiers).
		Write1b(vModsMaskHigh).
		Write1b(vModsMaskLow).
		Write1b(vModsHigh).
		Write1b(vModsLow) // 7b
}

// Action DeviceBtn
type SADeviceBtn struct {
	Flags  uint8
	Count  uint8
	Button uint8
	Device uint8
}

func (SADeviceBtn) Type() uint8 {
	return SATypeDeviceBtn
}

func readSADeviceBtn(r *x.Reader) Action {
	var v SADeviceBtn
	v.Flags = r.Read1b()
	v.Count = r.Read1b()
	v.Button = r.Read1b()
	v.Device = r.Read1b()
	return v
}

func (v SADeviceBtn) writeTo(b *x.FixedSizeBuf) {
	b.Write1b(v.Flags).
		Write1b(v.Count).
		Write1b(v.Button).
		Write1b(v.Device).
		WritePad(3)
}

// Action LockDeviceBtn
type SALockDeviceBtn struct {
	Flags  uint8
	Button uint8
	Device uint8
}

func (SALockDeviceBtn) Type() uint8 {
	return SATypeLockDeviceBtn
}

func readSALockDeviceBtn(r *x.Reader) Action {
	var v SALockDeviceBtn
	v.Flags = r.Read1b()
	r.ReadPad(1)
	v.Button = r.Read1b()
	v.Device = r.Read1b()
	return v
}

func (v SALockDeviceBtn) writeTo(b *x.FixedSizeBuf) {
	b.Write1b(v.Flags).
		WritePad(1).
		Write1b(v.Button).
		Write1b(v.Device). // 4b
		WritePad(3)
}

// Action DeviceValuator
type SADeviceValuator struct {
	Device    uint8
	Val1What  uint8
	Val1Index uint8
	Val1Value uint8

	Val2What  uint8
	Val2Index uint8
	Val2Value uint8
}

func (SADeviceValuator) Type() uint8 {
	return SATypeDeviceValuator
}

func readSADeviceValuator(r *x.Reader) Action {
	var v SADeviceValuator
	v.Device = r.Read1b()
	v.Val1What = r.Read1b()
	v.Val1Index = r.Read1b()
	v.Val1Value = r.Read1b()

	v.Val2What = r.Read1b()
	v.Val2Index = r.Read1b()
	v.Val2Value = r.Read1b()
	return v
}

func (v SADeviceValuator) writeTo(b *x.FixedSizeBuf) {
	b.Write1b(v.Device).
		Write1b(v.Val1What).
		Write1b(v.Val1Index).
		Write1b(v.Val1Value).
		Write1b(v.Val2What).
		Write1b(v.Val2Index).
		Write1b(v.Val2Value) // 7b
}

type UnknownAction struct {
	type0 uint8
	data  []byte // length 7
}

func (v UnknownAction) Type() uint8 {
	return v.type0
}

func (v UnknownAction) writeTo(b *x.FixedSizeBuf) {
	if len(v.data) != 7 {
		panic("length of data is not 7")
	}
	b.WriteBytes(v.data)
}

// size: 1 * 4b
type SetBehavior struct {
	Keycode  x.Keycode
	Behavior Behavior
}

func readSetBehavior(r *x.Reader, v *SetBehavior) {
	v.Keycode = x.Keycode(r.Read1b())
	v.Behavior = readBehavior(r)
	r.ReadPad(1) // 1
}

func writeSetBehavior(b *x.FixedSizeBuf, v SetBehavior) {
	b.Write1b(uint8(v.Keycode))
	writeBehavior(b, v.Behavior)
	b.WritePad(1)
}

// size: 2b
type Behavior interface {
	Type() uint8
	writeTo(b *x.FixedSizeBuf)
}

func readBehavior(r *x.Reader) Behavior {
	type0 := r.Read1b()
	data := r.Read1b()

	switch type0 {
	case BehaviorTypeDefault:
		return DefaultBehavior{}
	case BehaviorTypeLock:
		return LockBehavior{}
	case BehaviorTypeRadioGroup:
		return RadioGroupBehavior{data}
	case BehaviorTypeOverlay1:
		return Overlay1Behavior{x.Keycode(data)}
	case BehaviorTypeOverlay2:
		return Overlay2Behavior{x.Keycode(data)}
	case BehaviorTypePermamentLock:
		return PermamentLockBehavior{}
	case BehaviorTypePermamentRadioGroup:
		return PermamentRadioGroupBehavior{data}
	case BehaviorTypePermamentOverlay1:
		return PermamentOverlay1Behavior{}
	case BehaviorTypePermamentOverlay2:
		return PermamentOverlay1Behavior{}
	default:
		return UnknownBehavior{type0, data}
	}
}

func writeBehavior(b *x.FixedSizeBuf, v Behavior) {
	b.Write1b(v.Type())
	v.writeTo(b)
}

type UnknownBehavior struct {
	type0 uint8
	Data  uint8
}

func (v UnknownBehavior) writeTo(b *x.FixedSizeBuf) {
	b.Write1b(v.Data)
}

func (v UnknownBehavior) Type() uint8 {
	return v.type0
}

type DefaultBehavior struct {
}

func (DefaultBehavior) writeTo(b *x.FixedSizeBuf) {
	b.WritePad(1)
}

func (DefaultBehavior) Type() uint8 {
	return BehaviorTypeDefault
}

type LockBehavior struct {
}

func (LockBehavior) writeTo(b *x.FixedSizeBuf) {
	b.WritePad(1)
}

func (LockBehavior) Type() uint8 {
	return BehaviorTypeLock
}

type RadioGroupBehavior struct {
	Group uint8
}

func (v RadioGroupBehavior) writeTo(b *x.FixedSizeBuf) {
	b.Write1b(v.Group)
}

func (RadioGroupBehavior) Type() uint8 {
	return BehaviorTypeRadioGroup
}

type Overlay1Behavior struct {
	Key x.Keycode
}

func (v Overlay1Behavior) writeTo(b *x.FixedSizeBuf) {
	b.Write1b(uint8(v.Key))
}

func (Overlay1Behavior) Type() uint8 {
	return BehaviorTypeOverlay1
}

type Overlay2Behavior Overlay1Behavior

func (v Overlay2Behavior) writeTo(b *x.FixedSizeBuf) {
	b.Write1b(uint8(v.Key))
}

func (Overlay2Behavior) Type() uint8 {
	return BehaviorTypeOverlay2
}

type PermamentLockBehavior LockBehavior

func (PermamentLockBehavior) writeTo(b *x.FixedSizeBuf) {
	b.WritePad(1)
}

func (PermamentLockBehavior) Type() uint8 {
	return BehaviorTypePermamentLock
}

type PermamentRadioGroupBehavior RadioGroupBehavior

func (v PermamentRadioGroupBehavior) writeTo(b *x.FixedSizeBuf) {
	b.Write1b(v.Group)
}

func (PermamentRadioGroupBehavior) Type() uint8 {
	return BehaviorTypePermamentRadioGroup
}

type PermamentOverlay1Behavior Overlay1Behavior

func (v PermamentOverlay1Behavior) writeTo(b *x.FixedSizeBuf) {
	b.Write1b(uint8(v.Key))
}

func (PermamentOverlay1Behavior) Type() uint8 {
	return BehaviorTypePermamentOverlay1
}

type PermamentOverlay2Behavior Overlay1Behavior

func (v PermamentOverlay2Behavior) writeTo(b *x.FixedSizeBuf) {
	b.Write1b(uint8(v.Key))
}

func (PermamentOverlay2Behavior) Type() uint8 {
	return BehaviorTypePermamentOverlay2
}

// size: 2b
type SetExplicit struct {
	Keycode  x.Keycode
	Explicit uint8
}

func readSetExplicit(r *x.Reader) SetExplicit {
	var v SetExplicit
	v.Keycode = x.Keycode(r.Read1b())
	v.Explicit = r.Read1b()
	return v
}

func writeSetExplicit(b *x.FixedSizeBuf, v SetExplicit) {
	b.Write1b(uint8(v.Keycode))
	b.Write1b(v.Explicit)
}

// size: 2b
type KeyModMap struct {
	Keycode x.Keycode
	Mods    uint8
}

func readKeyModMap(r *x.Reader) KeyModMap {
	var v KeyModMap
	v.Keycode = x.Keycode(r.Read1b())
	v.Mods = r.Read1b()
	return v
}

func writeKeyModMap(b *x.FixedSizeBuf, v KeyModMap) {
	b.Write1b(uint8(v.Keycode)).
		Write1b(v.Mods)
}

// size: 1 * 4b
type KeyVModMap struct {
	Keycode x.Keycode
	VMods   uint16
}

func readKeyVModMap(r *x.Reader) KeyVModMap {
	var v KeyVModMap
	v.Keycode = x.Keycode(r.Read1b())
	r.ReadPad(1)
	v.VMods = r.Read2b()
	return v
}

func writeKeyVModMap(b *x.FixedSizeBuf, v KeyVModMap) {
	b.Write1b(uint8(v.Keycode)).
		WritePad(1).
		Write2b(v.VMods)
}

func readGetMapReply(r *x.Reader, v *GetMapReply) error {
	if !r.RemainAtLeast4b(10) {
		return x.ErrDataLenShort
	}
	v.DeviceID, _ = r.ReadReplyHeader()

	r.ReadPad(2)
	v.MinKeyCode = x.Keycode(r.Read1b())
	v.MaxKeyCode = x.Keycode(r.Read1b()) // 3

	v.Present = r.Read2b()
	v.FirstType = r.Read1b()
	v.NTypes = r.Read1b()

	v.TotalTypes = r.Read1b()
	v.FirstKeySym = x.Keycode(r.Read1b())
	v.TotalSyms = r.Read2b() // 5

	v.NKeySyms = r.Read1b()
	v.FirstKeyAction = x.Keycode(r.Read1b())
	v.TotalActions = r.Read2b()

	v.NKeyActions = r.Read1b()
	v.FirstKeyBehavior = x.Keycode(r.Read1b())
	v.NKeyBehaviors = r.Read1b()
	v.TotalKeyBehaviors = r.Read1b()

	v.FirstKeyExplicit = x.Keycode(r.Read1b())
	v.NKeyExplicit = r.Read1b()
	v.TotalKeyBehaviors = r.Read1b()
	v.FirstModMapKey = x.Keycode(r.Read1b()) // 8

	v.NModMapKeys = r.Read1b()
	v.TotalModMapKeys = r.Read1b()
	v.FirstVModMapKey = x.Keycode(r.Read1b())
	v.NVModMapKeys = r.Read1b()

	v.TotalVModMapKeys = r.Read1b()
	r.ReadPad(1)
	v.VirtualMods = r.Read2b() // 10

	if v.Present&MapPartKeyTypes != 0 && v.NTypes > 0 {
		v.Types = make([]KeyType, v.NTypes)
		for i := 0; i < int(v.NTypes); i++ {
			err := readKeyType(r, &v.Types[i])
			if err != nil {
				return err
			}
		}
	}

	if v.Present&MapPartKeySyms != 0 && v.NKeySyms > 0 {
		v.Syms = make([]KeySymMap, v.NKeySyms)
		for i := 0; i < int(v.NKeySyms); i++ {
			err := readKeySymMap(r, &v.Syms[i])
			if err != nil {
				return err
			}
		}
	}

	if v.Present&MapPartKeyActions != 0 {
		if v.NKeyActions > 0 {
			var err error
			v.ActionsCount, err = r.ReadBytesWithPad(int(v.NKeyActions))
			if err != nil {
				return err
			}
		}

		if v.TotalActions > 0 {
			if !r.RemainAtLeast4b(2 * int(v.TotalActions)) {
				return x.ErrDataLenShort
			}
			v.Actions = make([]Action, v.TotalActions)
			for i := 0; i < int(v.TotalActions); i++ {
				v.Actions[i] = readAction(r)
			}
		}
	}

	if v.Present&MapPartKeyBehaviors != 0 && v.TotalKeyBehaviors > 0 {
		if !r.RemainAtLeast4b(int(v.TotalKeyBehaviors)) {
			return x.ErrDataLenShort
		}
		v.Behaviors = make([]SetBehavior, v.TotalKeyBehaviors)
		for i := 0; i < int(v.TotalKeyBehaviors); i++ {
			readSetBehavior(r, &v.Behaviors[i])
		}
	}

	if v.Present&MapPartVirtualMods != 0 {
		length := x.PopCount(int(v.VirtualMods))
		if length > 0 {
			var err error
			v.VMods, err = r.ReadBytesWithPad(length)
			if err != nil {
				return err
			}
		}
	}

	if v.Present&MapPartExplicitComponents != 0 && v.TotalKeyExplicit > 0 {
		n := int(v.TotalKeyExplicit) * 2
		pad := x.Pad(n)
		if !r.RemainAtLeast(n + pad) {
			return x.ErrDataLenShort
		}
		v.Explicit = make([]SetExplicit, v.TotalKeyExplicit)
		for i := 0; i < int(v.TotalKeyExplicit); i++ {
			v.Explicit[i] = readSetExplicit(r)
		}
		r.ReadPad(pad)
	}

	if v.Present&MapPartModifierMap != 0 && v.TotalModMapKeys > 0 {
		n := int(v.TotalModMapKeys) * 2
		pad := x.Pad(n)
		if !r.RemainAtLeast(n + pad) {
			return x.ErrDataLenShort
		}
		v.ModMap = make([]KeyModMap, v.TotalModMapKeys)
		for i := 0; i < int(v.TotalModMapKeys); i++ {
			v.ModMap[i] = readKeyModMap(r)
		}
		r.ReadPad(pad)
	}

	if v.Present&MapPartVirtualModMap != 0 && v.TotalVModMapKeys > 0 {
		if !r.RemainAtLeast4b(int(v.TotalVModMapKeys)) {
			return x.ErrDataLenShort
		}
		v.VModMap = make([]KeyVModMap, v.TotalVModMapKeys)
		for i := 0; i < int(v.TotalVModMapKeys); i++ {
			v.VModMap[i] = readKeyVModMap(r)
		}
	}

	return nil
}

type SetMapValue struct {
	MinKeyCode x.Keycode
	MaxKeyCode x.Keycode
	//Present    uint16

	FirstType  uint8
	NTypes     uint8 // Types
	TotalTypes uint8

	FirstKeySym x.Keycode
	TotalSyms   uint16
	NKeySyms    uint8 // Syms

	FirstKeyAction x.Keycode
	TotalActions   uint16 // Actions
	NKeyActions    uint8  // ActionsCount

	FirstKeyBehavior  x.Keycode
	NKeyBehaviors     uint8
	TotalKeyBehaviors uint8 // Behaviors

	FirstKeyExplicit x.Keycode
	NKeyExplicit     uint8
	TotalKeyExplicit uint8 // Explicit

	FirstModMapKey  x.Keycode
	NModMapKeys     uint8
	TotalModMapKeys uint8 // modMap

	FirstVModMapKey  x.Keycode
	NVModMapKeys     uint8
	TotalVModMapKeys uint8 // vModMap

	VirtualMods uint16 // VMods

	Types        []SetKeyType
	Syms         []KeySymMap
	ActionsCount []uint8
	Actions      []Action
	Behaviors    []SetBehavior
	VMods        []uint8
	Explicit     []SetExplicit
	ModMap       []KeyModMap
	VModMap      []KeyVModMap
}

// #WREQ
func encodeSetMap(deviceSpec DeviceSpec, flags uint16, m *SetMapValue) (b x.RequestBody) {
	var present uint16

	if m.NTypes > 0 {
		present |= MapPartKeyTypes
	}
	if m.NKeySyms > 0 {
		present |= MapPartKeySyms
	}
	if m.NKeyActions > 0 || m.TotalActions > 0 {
		present |= MapPartKeyActions
	}
	if m.TotalKeyBehaviors > 0 {
		present |= MapPartKeyBehaviors
	}
	if m.VirtualMods != 0 {
		present |= MapPartVirtualMods
	}
	if m.TotalKeyExplicit > 0 {
		present |= MapPartExplicitComponents
	}
	if m.TotalModMapKeys > 0 {
		present |= MapPartModifierMap
	}
	if m.TotalVModMapKeys > 0 {
		present |= MapPartVirtualModMap
	}

	// block size in 4b
	size := 8
	for i := 0; i < int(m.NTypes); i++ {
		size += m.Types[i].sizeIn4b()
	}
	for i := 0; i < int(m.NKeySyms); i++ {
		size += m.Syms[i].sizeIn4b()
	}
	size += x.SizeIn4bWithPad(len(m.ActionsCount))
	size += 2 * len(m.Actions)
	size += len(m.Behaviors)
	size += x.SizeIn4bWithPad(len(m.VMods))
	size += x.SizeIn4bWithPad(len(m.Explicit) * 2)
	size += x.SizeIn4bWithPad(len(m.ModMap) * 2)
	size += len(m.VModMap)

	b0 := b.AddBlock(size).
		Write2b(uint16(deviceSpec)).
		Write2b(present).
		Write2b(flags).
		Write1b(uint8(m.MinKeyCode)).
		Write1b(uint8(m.MaxKeyCode)). // 2
		Write1b(m.FirstType).
		Write1b(m.NTypes).
		Write1b(uint8(m.FirstKeySym)).
		Write1b(m.NKeySyms). // 3
		Write2b(m.TotalSyms).
		Write1b(uint8(m.FirstKeyAction)).
		Write1b(m.NKeyActions).
		Write2b(m.TotalActions).
		Write1b(uint8(m.FirstKeyBehavior)).
		Write1b(m.NKeyBehaviors). // 5
		Write1b(m.TotalKeyBehaviors).
		Write1b(uint8(m.FirstKeyExplicit)).
		Write1b(m.NKeyExplicit).
		Write1b(m.TotalKeyExplicit).
		Write1b(uint8(m.FirstModMapKey)).
		Write1b(m.NModMapKeys).
		Write1b(m.TotalModMapKeys).
		Write1b(uint8(m.FirstVModMapKey)). // 7
		Write1b(m.NVModMapKeys).
		Write1b(m.TotalVModMapKeys).
		Write2b(m.VirtualMods) // 8

	for i := 0; i < int(m.NTypes); i++ {
		writeSetKeyType(b0, &m.Types[i])
	}

	for i := 0; i < int(m.NKeySyms); i++ {
		writeKeySymMap(b0, &m.Syms[i])
	}

	b0.WriteBytes(m.ActionsCount)
	b0.WritePad(x.Pad(len(m.ActionsCount)))

	for _, act := range m.Actions {
		writeAction(b0, act)
	}

	for _, b := range m.Behaviors {
		writeSetBehavior(b0, b)
	}

	b0.WriteBytes(m.VMods)
	b0.WritePad(x.Pad(len(m.VMods)))

	for _, e := range m.Explicit {
		writeSetExplicit(b0, e)
	}
	b0.WritePad(x.Pad(len(m.Explicit) * 2))

	for _, m := range m.ModMap {
		writeKeyModMap(b0, m)
	}
	b0.WritePad(x.Pad(len(m.ModMap) * 2))

	for _, v := range m.VModMap {
		writeKeyVModMap(b0, v)
	}
	b0.End()
	return
}

// #WREQ
func encodeGetCompatMap(deviceSpec DeviceSpec, groups uint8, getAllSI bool,
	firstSI, NSI uint16) (b x.RequestBody) {
	b.AddBlock(2).
		Write2b(uint16(deviceSpec)).
		Write1b(groups).
		WriteBool(getAllSI).
		Write2b(firstSI).
		Write2b(NSI).
		End()
	return
}

type GetCompatMapReply struct {
	DeviceID    uint8
	GroupsRtrn  uint8
	FirstSIRtrn uint16
	NSIRtrn     uint16
	NTotalSI    uint16
	SIRtrn      []SymInterpret
	GroupRtrn   []ModDef
}

func readGetCompatMapReply(r *x.Reader, v *GetCompatMapReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	v.DeviceID, _ = r.ReadReplyHeader()

	v.GroupsRtrn = r.Read1b()
	r.ReadPad(1)
	v.FirstSIRtrn = r.Read2b()

	v.NSIRtrn = r.Read2b()
	v.NTotalSI = r.Read2b() // 4

	r.ReadPad(16) // 8

	if v.NSIRtrn > 0 {
		if !r.RemainAtLeast4b(4 * int(v.NSIRtrn)) {
			return x.ErrDataLenShort
		}
		v.SIRtrn = make([]SymInterpret, v.NSIRtrn)
		for i := 0; i < int(v.NSIRtrn); i++ {
			readSymInterpret(r, &v.SIRtrn[i])
		}
	}

	if v.GroupsRtrn > 0 {
		length := x.PopCount(int(v.GroupsRtrn))
		if !r.RemainAtLeast4b(length) {
			return x.ErrDataLenShort
		}
		v.GroupRtrn = make([]ModDef, length)
		for i := 0; i < length; i++ {
			v.GroupRtrn[i] = readModDef(r)
		}
	}

	return nil
}

// size: 4 * 4b
type SymInterpret struct {
	Sym        x.Keysym
	Mods       uint8
	Match      uint8
	VirtualMod uint8
	Flags      uint8
	Action     Action
}

func readSymInterpret(r *x.Reader, v *SymInterpret) {
	v.Sym = x.Keysym(r.Read4b())
	v.Mods = r.Read1b()
	v.Match = r.Read1b()
	v.VirtualMod = r.Read1b()
	v.Flags = r.Read1b()
	v.Action = readAction(r)
}

func writeSymInterpret(b *x.FixedSizeBuf, v *SymInterpret) {
	b.Write4b(uint32(v.Sym)).
		Write1b(v.Mods).
		Write1b(v.Match).
		Write1b(v.VirtualMod).
		Write1b(v.Flags)
	writeAction(b, v.Action)
}

type SetCompatMapValue struct {
	RecomputeActions bool
	TruncateSI       bool
	Groups           uint8
	FirstSI          uint16
	NSI              uint16
	SI               []SymInterpret
	GroupMaps        []ModDef
}

// #WREQ
func encodeSetCompatMap(deviceSpec DeviceSpec, v *SetCompatMapValue) (b x.RequestBody) {
	v.NSI = uint16(len(v.SI))
	b0 := b.AddBlock(3 + 4*len(v.SI)).
		Write2b(uint16(deviceSpec)).
		WritePad(1).
		WriteBool(v.RecomputeActions).
		WriteBool(v.TruncateSI).
		Write1b(v.Groups).
		Write2b(v.FirstSI).
		Write2b(v.NSI).
		WritePad(2) // 3

	for i := 0; i < len(v.SI); i++ {
		writeSymInterpret(b0, &v.SI[i])
	}

	nGroup := x.PopCount(int(v.Groups))
	for i := 0; i < nGroup; i++ {
		writeModDef(b0, v.GroupMaps[i])
	}
	b0.End()
	return b
}

// #WREQ
func encodeGetIndicatorState(deviceSpec DeviceSpec) (b x.RequestBody) {
	b.AddBlock(1).
		Write2b(uint16(deviceSpec)).
		WritePad(2).
		End()
	return
}

type GetIndicatorStateReply struct {
	DeviceID uint8
	State    uint32
}

func readGetIndicatorStateReply(r *x.Reader, v *GetIndicatorStateReply) error {
	if !r.RemainAtLeast4b(3) {
		return x.ErrDataLenShort
	}
	v.DeviceID, _ = r.ReadReplyHeader()
	v.State = r.Read4b()
	return nil
}

// #WREQ
func encodeGetIndicatorMap(deviceSpec DeviceSpec, which uint32) (b x.RequestBody) {
	b.AddBlock(2).
		Write2b(uint16(deviceSpec)).
		WritePad(2).
		Write4b(which)
	return
}

type GetIndicatorMapReply struct {
	DeviceID       uint8
	Which          uint32
	RealIndicators uint32
	NIndicators    uint8
	Maps           []IndicatorMap
}

// size: 3 * 4b
type IndicatorMap struct {
	Flags       uint8
	WhichGroups uint8
	Groups      uint8
	WhichMods   uint8

	Mods     uint8
	RealMods uint8
	VMods    uint16

	Ctrls uint32
}

func readIndicatorMap(r *x.Reader, v *IndicatorMap) {
	v.Flags = r.Read1b()
	v.WhichGroups = r.Read1b()
	v.Groups = r.Read1b()
	v.WhichMods = r.Read1b()

	v.Mods = r.Read1b()
	v.RealMods = r.Read1b()
	v.VMods = r.Read2b()

	v.Ctrls = r.Read4b()
}

func writeIndicatorMap(b *x.FixedSizeBuf, v *IndicatorMap) {
	b.Write1b(v.Flags).
		Write1b(v.WhichGroups).
		Write1b(v.Groups).
		Write1b(v.WhichGroups).
		Write1b(v.Mods).
		Write1b(v.RealMods).
		Write2b(v.VMods).
		Write4b(v.Ctrls)
}

func readGetIndicatorMapReply(r *x.Reader, v *GetIndicatorMapReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	v.DeviceID, _ = r.ReadReplyHeader()
	v.Which = r.Read4b()
	v.RealIndicators = r.Read4b() // 4
	v.NIndicators = r.Read1b()
	r.ReadPad(15) // 8

	mapsLen := x.PopCount(int(v.Which))
	if mapsLen > 0 {
		if !r.RemainAtLeast4b(3 * mapsLen) {
			return x.ErrDataLenShort
		}
		v.Maps = make([]IndicatorMap, mapsLen)
		for i := 0; i < mapsLen; i++ {
			readIndicatorMap(r, &v.Maps[i])
		}
	}

	return nil
}

// #WREQ
func encodeSetIndicatorMap(deviceSpec DeviceSpec, which uint32,
	maps []IndicatorMap) (b x.RequestBody) {

	mapsLen := x.PopCount(int(which))
	b0 := b.AddBlock(2 + 3*mapsLen).
		Write2b(uint16(deviceSpec)).
		WritePad(2).
		Write4b(which)

	for i := 0; i < mapsLen; i++ {
		writeIndicatorMap(b0, &maps[i])
	}
	b0.End()
	return
}

// #WREQ
func encodeGetNamedIndicator(deviceSpec DeviceSpec, ledClass LedClassSpec,
	ledID IdSpec, indicator x.Atom) (b x.RequestBody) {
	b.AddBlock(3).
		Write2b(uint16(deviceSpec)).
		Write2b(uint16(ledClass)).
		Write2b(uint16(ledID)).
		WritePad(2).
		Write4b(uint32(indicator)).
		End()
	return
}

type GetNamedIndicatorReply struct {
	DeviceID      uint8
	Indicator     x.Atom
	Found         bool
	On            bool
	RealIndicator bool
	Ndx           uint8
	Map           IndicatorMap
	Supported     bool
}

func readGetNamedIndicatorReply(r *x.Reader, v *GetNamedIndicatorReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	v.DeviceID, _ = r.ReadReplyHeader()

	v.Indicator = x.Atom(r.Read4b()) // 3

	v.Found = r.ReadBool()
	v.On = r.ReadBool()
	v.RealIndicator = r.ReadBool()
	v.Ndx = r.Read1b() // 4

	readIndicatorMap(r, &v.Map) // 7

	v.Supported = r.ReadBool() // 8

	return nil
}

// #WREQ
func encodeSetNamedIndicator(deviceSpec DeviceSpec, ledClass LedClassSpec,
	ledID IdSpec, indicator x.Atom, setState, on, setMap, createMap bool,
	map0 *IndicatorMap) (b x.RequestBody) {

	b.AddBlock(7).
		Write2b(uint16(deviceSpec)).
		Write2b(uint16(ledClass)).
		Write2b(uint16(ledID)).
		WritePad(2). // 2
		Write4b(uint32(indicator)).
		WriteBool(setState).
		WriteBool(on).
		WriteBool(setMap).
		WriteBool(createMap). // 4
		WritePad(1).
		Write1b(map0.Flags).
		Write1b(map0.WhichGroups).
		Write1b(map0.Groups). // 5
		Write1b(map0.WhichMods).
		Write1b(map0.RealMods).
		Write2b(map0.VMods).
		Write4b(map0.Ctrls). // 7
		End()
	// NOTE: map0.Mods is not used
	return
}

// #WREQ
func encodeGetNames(deviceSpec DeviceSpec, which uint32) (b x.RequestBody) {
	b.AddBlock(2).
		Write2b(uint16(deviceSpec)).
		WritePad(2).
		Write4b(which).
		End()
	return
}

type GetNamesReply struct {
	DeviceID     uint8
	Which        uint32
	MinKeyCode   x.Keycode
	MaxKeyCode   x.Keycode
	NTypes       uint8
	Groups       uint8
	VirtualMods  uint16
	FirstKey     x.Keycode
	NKeys        uint8
	Indicators   uint32
	NRadioGroups uint8
	NKeyAliases  uint8
	NKTLevels    uint16

	KeycodesName    x.Atom
	GeometryName    x.Atom
	SymbolsName     x.Atom
	PhysSymbolsName x.Atom
	TypesName       x.Atom
	CompatName      x.Atom
	TypeNames       []x.Atom
	NLevelsPerType  []uint8
	KTLevelNames    []x.Atom
	IndicatorNames  []x.Atom
	VirtualModNames []x.Atom
	GroupNames      []x.Atom
	KeyNames        []string
	KeyAliases      []KeyAlias
	RadioGroupNames []x.Atom
}

// size: 1 * 4b
//type KeyName struct {
//	Name []byte // length 4
//}

func readKeyName(r *x.Reader) string {
	return bytesToStr(r.MustReadBytes(4))
}

func writeKeyName(b *x.FixedSizeBuf, str string) {
	writeStr4b(b, str)
}

func writeStr4b(b *x.FixedSizeBuf, str string) {
	if len(str) >= 4 {
		b.WriteBytes([]byte(str[:4]))
	} else {
		b.WriteBytes([]byte(str))
		b.WritePad(4 - len(str))
	}
}

func bytesToStr(data []byte) string {
	idx := bytes.IndexByte(data, 0)
	if idx != -1 {
		return string(data[:idx])
	}
	return string(data)
}

// size: 2 * 4b
type KeyAlias struct {
	Real  string
	Alias string
}

func readKeyAlias(r *x.Reader) KeyAlias {
	var v KeyAlias
	v.Real = bytesToStr(r.MustReadBytes(4))
	v.Alias = bytesToStr(r.MustReadBytes(4))
	return v
}

func writeKeyAlias(b *x.FixedSizeBuf, v KeyAlias) {
	writeStr4b(b, v.Real)
	writeStr4b(b, v.Alias)
}

//type KeyAlias struct {
//	Real []byte // length 4
//	Alias []byte // length 4
//}

func readGetNamesReply(r *x.Reader, v *GetNamesReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	v.DeviceID, _ = r.ReadReplyHeader()

	v.Which = r.Read4b()

	v.MinKeyCode = x.Keycode(r.Read1b())
	v.MaxKeyCode = x.Keycode(r.Read1b())
	v.NTypes = r.Read1b()
	v.Groups = r.Read1b() // 4

	v.VirtualMods = r.Read2b()
	v.FirstKey = x.Keycode(r.Read1b())
	v.NKeys = r.Read1b()

	v.Indicators = r.Read4b()

	v.NRadioGroups = r.Read1b()
	v.NKeyAliases = r.Read1b()
	v.NKTLevels = r.Read2b()

	r.ReadPad(4) // 8

	// in 4b
	var valueListSize int
	if v.Which&NameDetailKeycodes != 0 {
		valueListSize++
	}
	if v.Which&NameDetailGeometry != 0 {
		valueListSize++
	}
	if v.Which&NameDetailSymbols != 0 {
		valueListSize++
	}
	if v.Which&NameDetailPhysSymbols != 0 {
		valueListSize++
	}
	if v.Which&NameDetailTypes != 0 {
		valueListSize++
	}
	if v.Which&NameDetailCompat != 0 {
		valueListSize++
	}
	if v.Which&NameDetailKeyTypeNames != 0 {
		valueListSize += int(v.NTypes)
	}
	if !r.RemainAtLeast4b(valueListSize) {
		return x.ErrDataLenShort
	}

	if v.Which&NameDetailKeycodes != 0 {
		v.KeycodesName = x.Atom(r.Read4b())
	}
	if v.Which&NameDetailGeometry != 0 {
		v.GeometryName = x.Atom(r.Read4b())
	}
	if v.Which&NameDetailSymbols != 0 {
		v.SymbolsName = x.Atom(r.Read4b())
	}
	if v.Which&NameDetailPhysSymbols != 0 {
		v.PhysSymbolsName = x.Atom(r.Read4b())
	}
	if v.Which&NameDetailTypes != 0 {
		v.TypesName = x.Atom(r.Read4b())
	}
	if v.Which&NameDetailCompat != 0 {
		v.CompatName = x.Atom(r.Read4b())
	}
	if v.Which&NameDetailKeyTypeNames != 0 && v.NTypes > 0 {
		v.TypeNames = make([]x.Atom, v.NTypes)
		for i := 0; i < int(v.NTypes); i++ {
			v.TypeNames[i] = x.Atom(r.Read4b())
		}
	}

	if v.Which&NameDetailKtLevelNames != 0 {
		var err error
		v.NLevelsPerType, err = r.ReadBytesWithPad(int(v.NTypes))
		if err != nil {
			return err
		}

		var sum int
		for _, value := range v.NLevelsPerType {
			sum += int(value)
		}
		if sum > 0 {
			if !r.RemainAtLeast4b(sum) {
				return x.ErrDataLenShort
			}
			v.KTLevelNames = make([]x.Atom, sum)
			for i := 0; i < sum; i++ {
				v.KTLevelNames[i] = x.Atom(r.Read4b())
			}
		}
	}

	valueListSize = 0
	if v.Which&NameDetailIndicatorNames != 0 {
		valueListSize += x.PopCount(int(v.Indicators))
	}
	if v.Which&NameDetailVirtualModNames != 0 {
		valueListSize += x.PopCount(int(v.VirtualMods))
	}
	if v.Which&NameDetailGroupNames != 0 {
		valueListSize += x.PopCount(int(v.Groups))
	}
	if v.Which&NameDetailKeyNames != 0 {
		valueListSize += int(v.NKeys)
	}
	if v.Which&NameDetailKeyAliases != 0 {
		valueListSize += 2 * int(v.NKeyAliases)
	}
	if v.Which&NameDetailRgNames != 0 {
		valueListSize += int(v.NRadioGroups)
	}
	if !r.RemainAtLeast4b(valueListSize) {
		return x.ErrDataLenShort
	}

	if v.Which&NameDetailIndicatorNames != 0 {
		n := x.PopCount(int(v.Indicators))
		if n > 0 {
			v.IndicatorNames = make([]x.Atom, n)
			for i := 0; i < n; i++ {
				v.IndicatorNames[i] = x.Atom(r.Read4b())
			}
		}
	}
	if v.Which&NameDetailVirtualModNames != 0 {
		n := x.PopCount(int(v.VirtualMods))
		if n > 0 {
			v.VirtualModNames = make([]x.Atom, n)
			for i := 0; i < n; i++ {
				v.VirtualModNames[i] = x.Atom(r.Read4b())
			}
		}
	}
	if v.Which&NameDetailGroupNames != 0 {
		n := x.PopCount(int(v.Groups))
		if n > 0 {
			v.GroupNames = make([]x.Atom, n)
			for i := 0; i < n; i++ {
				v.GroupNames[i] = x.Atom(r.Read4b())
			}
		}
	}
	if v.Which&NameDetailKeyNames != 0 && v.NKeys > 0 {
		v.KeyNames = make([]string, v.NKeys)
		for i := 0; i < int(v.NKeys); i++ {
			v.KeyNames[i] = readKeyName(r)
		}
	}
	if v.Which&NameDetailKeyAliases != 0 && v.NKeyAliases > 0 {
		v.KeyAliases = make([]KeyAlias, v.NKeyAliases)
		for i := 0; i < int(v.NKeyAliases); i++ {
			v.KeyAliases[i] = readKeyAlias(r)
		}
	}
	if v.Which&NameDetailRgNames != 0 && v.NRadioGroups > 0 {
		v.RadioGroupNames = make([]x.Atom, v.NRadioGroups)
		for i := 0; i < int(v.NRadioGroups); i++ {
			v.RadioGroupNames[i] = x.Atom(r.Read4b())
		}
	}
	return nil
}

type SetNamesRequest struct {
	DeviceSpec        DeviceSpec
	VirtualMods       uint16
	Which             uint32
	FirstType         uint8
	NTypes            uint8
	FirstKTLevel      uint8
	NKTLevels         uint8
	Indicators        uint32
	Groups            uint8
	NRadioGroups      uint8
	FirstKey          x.Keycode
	NKeys             uint8
	NKeyAliases       uint8
	TotalKTLevelNames uint16

	KeycodesName    x.Atom
	GeometryName    x.Atom
	SymbolsName     x.Atom
	PhysSymbolsName x.Atom
	TypesName       x.Atom
	CompatName      x.Atom
	TypeNames       []x.Atom
	NLevelsPerType  []uint8
	KTLevelNames    []x.Atom
	IndicatorNames  []x.Atom
	VirtualModNames []x.Atom
	GroupNames      []x.Atom
	KeyNames        []string
	KeyAliases      []KeyAlias
	RadioGroupNames []x.Atom
}

// #WREQ
func encodeSetNames(r *SetNamesRequest) (b x.RequestBody) {
	// in 4b
	valueListSize := 0

	if r.Which&NameDetailKeycodes != 0 {
		valueListSize++
	}
	if r.Which&NameDetailGeometry != 0 {
		valueListSize++
	}
	if r.Which&NameDetailSymbols != 0 {
		valueListSize++
	}
	if r.Which&NameDetailPhysSymbols != 0 {
		valueListSize++
	}
	if r.Which&NameDetailTypes != 0 {
		valueListSize++
	}
	if r.Which&NameDetailCompat != 0 {
		valueListSize++
	}
	if r.Which&NameDetailKeyTypeNames != 0 {
		valueListSize += int(r.NTypes)
	}
	if r.Which&NameDetailKtLevelNames != 0 {
		sum := 0
		for _, value := range r.NLevelsPerType {
			sum += int(value)
		}
		valueListSize += x.SizeIn4bWithPad(len(r.NLevelsPerType)) + sum
	}
	if r.Which&NameDetailIndicatorNames != 0 {
		valueListSize += x.PopCount(int(r.Indicators))
	}
	if r.Which&NameDetailVirtualModNames != 0 {
		valueListSize += x.PopCount(int(r.VirtualMods))
	}
	if r.Which&NameDetailGroupNames != 0 {
		valueListSize += x.PopCount(int(r.Groups))
	}
	if r.Which&NameDetailKeyNames != 0 {
		valueListSize += int(r.NKeys)
	}
	if r.Which&NameDetailKeyAliases != 0 {
		valueListSize += 3 * int(r.NKeyAliases)
	}
	if r.Which&NameDetailRgNames != 0 {
		valueListSize += int(r.NRadioGroups)
	}

	b0 := b.AddBlock(6 + valueListSize).
		Write2b(uint16(r.DeviceSpec)).
		Write2b(r.VirtualMods).
		Write4b(r.Which). // 2
		Write1b(r.FirstType).
		Write1b(r.NTypes).
		Write1b(r.FirstKTLevel).
		Write1b(r.NKTLevels). // 3
		Write4b(r.Indicators).
		Write1b(r.Groups).
		Write1b(r.NRadioGroups).
		Write1b(uint8(r.FirstKey)).
		Write1b(r.NKeys). // 5
		Write1b(r.NKeyAliases).
		WritePad(1).
		Write2b(r.TotalKTLevelNames) // 6

	if r.Which&NameDetailKeycodes != 0 {
		b0.Write4b(uint32(r.KeycodesName))
	}
	if r.Which&NameDetailGeometry != 0 {
		b0.Write4b(uint32(r.GeometryName))
	}
	if r.Which&NameDetailSymbols != 0 {
		b0.Write4b(uint32(r.SymbolsName))
	}
	if r.Which&NameDetailPhysSymbols != 0 {
		b0.Write4b(uint32(r.PhysSymbolsName))
	}
	if r.Which&NameDetailTypes != 0 {
		b0.Write4b(uint32(r.TypesName))
	}
	if r.Which&NameDetailCompat != 0 {
		b0.Write4b(uint32(r.CompatName))
	}
	if r.Which&NameDetailKeyTypeNames != 0 {
		for i := 0; i < int(r.NTypes); i++ {
			b0.Write4b(uint32(r.TypeNames[i]))
		}
	}
	if r.Which&NameDetailKtLevelNames != 0 {
		b0.WriteBytes(r.NLevelsPerType)
		b0.WritePad(x.Pad(len(r.NLevelsPerType)))

		sum := 0
		for _, value := range r.NLevelsPerType {
			sum += int(value)
		}
		for i := 0; i < sum; i++ {
			b0.Write4b(uint32(r.KTLevelNames[i]))
		}
	}
	if r.Which&NameDetailIndicatorNames != 0 {
		n := x.PopCount(int(r.Indicators))
		for i := 0; i < n; i++ {
			b0.Write4b(uint32(r.IndicatorNames[i]))
		}
	}
	if r.Which&NameDetailVirtualModNames != 0 {
		n := x.PopCount(int(r.VirtualMods))
		for i := 0; i < n; i++ {
			b0.Write4b(uint32(r.VirtualModNames[i]))
		}
	}
	if r.Which&NameDetailGroupNames != 0 {
		n := x.PopCount(int(r.Groups))
		for i := 0; i < n; i++ {
			b0.Write4b(uint32(r.GroupNames[i]))
		}
	}
	if r.Which&NameDetailKeyNames != 0 {
		for i := 0; i < int(r.NKeys); i++ {
			writeKeyName(b0, r.KeyNames[i])
		}
	}
	if r.Which&NameDetailKeyAliases != 0 {
		for i := 0; i < int(r.NKeyAliases); i++ {
			writeKeyAlias(b0, r.KeyAliases[i])
		}
	}
	if r.Which&NameDetailRgNames != 0 {
		for i := 0; i < int(r.NRadioGroups); i++ {
			b0.Write4b(uint32(r.RadioGroupNames[i]))
		}
	}

	return
}

// TODO: GetGeometry

// TODO: SetGeometry

// #WREQ
func encodePerClientFlags(deviceSpec DeviceSpec, change, value, ctrlsToChange,
	autoCtrls, autoCtrlsValues uint32) (b x.RequestBody) {

	b.AddBlock(6).
		Write2b(uint16(deviceSpec)).
		WritePad(2).
		Write4b(change).
		Write4b(value). // 3
		Write4b(ctrlsToChange).
		Write4b(autoCtrls).
		Write4b(autoCtrlsValues). // 6
		End()
	return
}

type PerClientFlagsReply struct {
	DeviceID        uint8
	Supported       uint32
	Value           uint32
	AutoCtrls       uint32
	AutoCtrlsValues uint32
}

func readPerClientFlagsReply(r *x.Reader, v *PerClientFlagsReply) error {
	if !r.RemainAtLeast4b(6) {
		return x.ErrDataLenShort
	}
	v.DeviceID, _ = r.ReadReplyHeader()

	v.Supported = r.Read4b()

	v.Value = r.Read4b() // 4

	v.AutoCtrls = r.Read4b()

	v.AutoCtrlsValues = r.Read4b() // 6
	return nil
}

type ComponentNames struct {
	KeymapsSpec   string
	KeycodesSpec  string
	TypesSpec     string
	CompatMapSpec string
	SymbolsSpec   string
	GeometrySpec  string
}

// #WREQ
func encodeListComponents(deviceSpec DeviceSpec, maxNames uint16,
	cn *ComponentNames) (b x.RequestBody) {

	keymapsSpec := x.TruncateStr(cn.KeymapsSpec, math.MaxUint8)
	keycodesSpec := x.TruncateStr(cn.KeycodesSpec, math.MaxUint8)
	typesSpec := x.TruncateStr(cn.TypesSpec, math.MaxUint8)
	compatMapSpec := x.TruncateStr(cn.CompatMapSpec, math.MaxUint8)
	symbolsSpec := x.TruncateStr(cn.SymbolsSpec, math.MaxUint8)
	geometrySpec := x.TruncateStr(cn.GeometrySpec, math.MaxUint8)

	specsLen := 6 +
		len(keymapsSpec) + len(keycodesSpec) + len(typesSpec) +
		len(compatMapSpec) + len(symbolsSpec) + len(geometrySpec)
	b.AddBlock(1 + x.SizeIn4bWithPad(specsLen)).
		Write2b(uint16(deviceSpec)).
		Write2b(maxNames).
		Write1b(uint8(len(keymapsSpec))).
		WriteString(keymapsSpec).
		Write1b(uint8(len(keycodesSpec))).
		WriteString(keycodesSpec).
		Write1b(uint8(len(typesSpec))).
		WriteString(typesSpec).
		Write1b(uint8(len(compatMapSpec))).
		WriteString(compatMapSpec).
		Write1b(uint8(len(symbolsSpec))).
		WriteString(symbolsSpec).
		Write1b(uint8(len(geometrySpec))).
		WriteString(geometrySpec).
		WritePad(x.Pad(specsLen)).
		End()
	return
}

type ListComponentsReply struct {
	DeviceID    uint8
	NKeymaps    uint16
	NKeycodes   uint16
	NTypes      uint16
	NCompatMaps uint16
	NSymbols    uint16
	NGeometries uint16
	Extra       uint16

	Keymaps    []int
	Keycodes   []int
	Types      []int
	CompatMaps []int
	Symbols    []int
	Geometries []int
}

func readListComponentsReply(r *x.Reader, v *ListComponentsReply) error {
	if !r.RemainAtLeast4b(8) {
		return x.ErrDataLenShort
	}
	v.DeviceID, _ = r.ReadReplyHeader()

	v.NKeymaps = r.Read2b()
	v.NKeycodes = r.Read2b()

	v.NTypes = r.Read2b()
	v.NCompatMaps = r.Read2b()

	v.NSymbols = r.Read2b()
	v.NGeometries = r.Read2b() // 5

	v.Extra = r.Read2b()
	r.ReadPad(10) // 8

	// TODO:
	return nil
}

// TODO: GetKbdByName

// #WREQ
func encodeGetDeviceInfo(deviceSpec DeviceSpec, wanted uint16, allButtons bool,
	FirstButton, nButtons uint8, ledClass LedClassSpec, ledID IdSpec) (b x.RequestBody) {

	b.AddBlock(3).
		Write2b(uint16(deviceSpec)).
		Write2b(wanted).
		WriteBool(allButtons).
		Write1b(FirstButton).
		Write1b(nButtons).
		WritePad(1). // 2
		Write2b(uint16(ledClass)).
		Write2b(uint16(ledID)). // 3
		End()
	return
}

type GetDeviceInfoReply struct {
	DeviceID       uint8
	Present        uint16
	Supported      uint16
	Unsupported    uint16
	NDeviceLedFBs  uint16
	FirstBtnWanted uint8
	NBtnsWanted    uint8
	FirstBtnRtrn   uint8
	NBtnsRtrn      uint8
	TotalBtns      uint8
	HasOwnState    bool
	DfltKbdFB      uint16
	DfltLedFB      uint16
	DevType        x.Atom
	NameLen        uint16
	Name           string
	BtnActions     []Action
	Leds           []DeviceLedInfo
}

type DeviceLedInfo struct {
	LedClass       LedClassSpec
	LedID          IdSpec
	NamesPresent   uint32
	MapsPresent    uint32
	PhysIndicators uint32
	State          uint32
	Names          []x.Atom
	Maps           []IndicatorMap
}

func readDeviceLedInfo(r *x.Reader, v *DeviceLedInfo) error {
	if !r.RemainAtLeast4b(5) {
		return x.ErrDataLenShort
	}
	v.LedClass = LedClassSpec(r.Read2b())
	v.LedID = IdSpec(r.Read2b())

	v.NamesPresent = r.Read4b()

	v.MapsPresent = r.Read4b()

	v.PhysIndicators = r.Read4b()

	v.State = r.Read4b() // 5

	if v.NamesPresent != 0 {
		n := x.PopCount(int(v.NamesPresent))
		if !r.RemainAtLeast4b(n) {
			return x.ErrDataLenShort
		}
		v.Names = make([]x.Atom, n)
		for i := 0; i < n; i++ {
			v.Names[i] = x.Atom(r.Read4b())
		}
	}

	if v.MapsPresent != 0 {
		n := x.PopCount(int(v.MapsPresent))
		if !r.RemainAtLeast4b(3 * n) {
			return x.ErrDataLenShort
		}
		v.Maps = make([]IndicatorMap, n)
		for i := 0; i < n; i++ {
			readIndicatorMap(r, &v.Maps[i])
		}
	}

	return nil
}

func (v *DeviceLedInfo) sizeIn4b() int {
	return 5 + x.PopCount(int(v.NamesPresent)) + 3*x.PopCount(int(v.MapsPresent))
}

func writeDeviceLedInfo(b *x.FixedSizeBuf, v *DeviceLedInfo) {
	b.Write2b(uint16(v.LedClass)).
		Write2b(uint16(v.LedID)). // 1
		Write4b(v.NamesPresent).
		Write4b(v.MapsPresent).
		Write4b(v.PhysIndicators).
		Write4b(v.State) // 5

	if v.NamesPresent != 0 {
		n := x.PopCount(int(v.NamesPresent))
		for i := 0; i < n; i++ {
			b.Write4b(uint32(v.Names[i]))
		}
	}

	if v.MapsPresent != 0 {
		n := x.PopCount(int(v.MapsPresent))
		for i := 0; i < n; i++ {
			writeIndicatorMap(b, &v.Maps[i])
		}
	}
}

func readGetDeviceInfoReply(r *x.Reader, v *GetDeviceInfoReply) error {
	if !r.RemainAtLeast(34) {
		return x.ErrDataLenShort
	}
	v.DeviceID, _ = r.ReadReplyHeader()

	v.Present = r.Read2b()
	v.Supported = r.Read2b()

	v.Unsupported = r.Read2b()
	v.NDeviceLedFBs = r.Read2b() // 4

	v.FirstBtnWanted = r.Read1b()
	v.NBtnsWanted = r.Read1b()
	v.FirstBtnRtrn = r.Read1b()
	v.NBtnsRtrn = r.Read1b()

	v.TotalBtns = r.Read1b()
	v.HasOwnState = r.ReadBool()
	v.DfltKbdFB = r.Read2b() // 6

	v.DfltLedFB = r.Read2b()
	r.ReadPad(2)

	v.DevType = x.Atom(r.Read4b()) // 8

	v.NameLen = r.Read2b() // 34b

	var err error
	v.Name, err = r.ReadStrWithPad(int(v.NameLen))
	if err != nil {
		return err
	}

	if v.NBtnsRtrn > 0 {
		if !r.RemainAtLeast4b(2 * int(v.NBtnsRtrn)) {
			return x.ErrDataLenShort
		}
		v.BtnActions = make([]Action, v.NBtnsRtrn)
		for i := 0; i < int(v.NBtnsRtrn); i++ {
			v.BtnActions[i] = readAction(r)
		}
	}

	if v.NDeviceLedFBs > 0 {
		v.Leds = make([]DeviceLedInfo, v.NDeviceLedFBs)
		for i := 0; i < int(v.NDeviceLedFBs); i++ {
			err = readDeviceLedInfo(r, &v.Leds[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// #WREQ
func encodeSetDeviceInfo(deviceSpec DeviceSpec, firstBtn uint8,
	change uint16, btnActions []Action, leds []DeviceLedInfo) (b x.RequestBody) {

	var nBtns uint8
	if len(btnActions) > math.MaxUint8 {
		nBtns = math.MaxUint8
		btnActions = btnActions[:math.MaxUint8]
	}

	var nDeviceLedFB uint16
	if len(leds) > math.MaxUint16 {
		nDeviceLedFB = math.MaxUint16
		leds = leds[:math.MaxUint16]
	}

	ledsSize := 0
	for i := 0; i < len(leds); i++ {
		ledsSize += leds[i].sizeIn4b()
	}

	b0 := b.AddBlock(2 + 2*int(nBtns) + ledsSize).
		Write2b(uint16(deviceSpec)).
		Write1b(firstBtn).
		Write1b(nBtns).
		Write2b(change).
		Write2b(nDeviceLedFB) // 2

	for _, a := range btnActions {
		writeAction(b0, a)
	}

	for i := 0; i < len(leds); i++ {
		writeDeviceLedInfo(b0, &leds[i])
	}
	return
}

// #WREQ
func encodeSetDebuggingFlags(affectFlags, flags, affectCtrls, ctrls uint32,
	message string) (b x.RequestBody) {
	message = x.TruncateStr(message, math.MaxUint16)
	msgLength := len(message)

	b.AddBlock(5 + x.SizeIn4bWithPad(msgLength)).
		Write2b(uint16(msgLength)).
		WritePad(2).
		Write4b(affectCtrls). // 2
		Write4b(flags).
		Write4b(affectCtrls).
		Write4b(ctrls). // 5
		WriteString(message).
		WritePad(x.Pad(msgLength)).
		End()
	return
}

type SetDebuggingFlagsReply struct {
	CurrentFlags   uint32
	CurrentCtrls   uint32
	SupportedFlags uint32
	SupportedCtrls uint32
}

func readSetDebuggingFlagsReply(r *x.Reader, v *SetDebuggingFlagsReply) error {
	if !r.RemainAtLeast4b(6) {
		return x.ErrDataLenShort
	}
	r.ReadPad(8)

	v.CurrentFlags = r.Read4b() // 3

	v.CurrentCtrls = r.Read4b()

	v.SupportedFlags = r.Read4b()

	v.SupportedCtrls = r.Read4b() // 6

	return nil
}
