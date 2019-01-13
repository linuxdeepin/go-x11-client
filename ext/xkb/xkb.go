package xkb

import (
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
	data, _ := r.ReadReplyHeader()
	if r.Err() != nil {
		return r.Err()
	}
	v.Supported = x.Uint8ToBool(data)

	v.ServerMajor = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ServerMinor = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(20)
	return r.Err()
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
		Write1b(x.BoolToUint8(forceSound)). // 2
		Write1b(x.BoolToUint8(eventOnly)).
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
	v.DeviceID, _ = r.ReadReplyHeader()
	if r.Err() != nil {
		return r.Err()
	}

	v.Mods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.BaseMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.LatchedMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.LockedMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Group = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.LockedGroup = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.BaseGroup = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.LatchedGroup = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.CompatState = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.GrabMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.CompatGrabMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.LookupMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.CompatLookupMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.PtrBtnState = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(6)
	return r.Err()
}

// #WREQ
func encodeLatchLockState(deviceSpec DeviceSpec, affectModLocks, modLocks uint8,
	lockGroup bool, groupLock, affectModLatches, modLatches uint8, latchGroup bool,
	groupLatch uint16) (b x.RequestBody) {

	b.AddBlock(3).
		Write2b(uint16(deviceSpec)).
		Write1b(affectModLocks).
		Write1b(modLocks). // 1
		Write1b(x.BoolToUint8(lockGroup)).
		Write1b(groupLock).
		Write1b(affectModLatches).
		Write1b(modLatches). // 2
		WritePad(1).
		Write1b(x.BoolToUint8(latchGroup)).
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
	v.DeviceID, _ = r.ReadReplyHeader()
	if r.Err() != nil {
		return r.Err()
	}

	v.MouseKeysDefaultBtn = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NumGroups = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.GroupsWrap = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Internal.Mask = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.IgnoreLock.Mask = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Internal.RealMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.IgnoreLock.RealMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.Internal.VirtualMods = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.IgnoreLock.VirtualMods = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.RepeatDelay = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.RepeatInterval = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SlowKeysDelay = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.DebounceDelay = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MouseKeysDelay = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MouseKeysInterval = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MouseKeysTimeToMax = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MouseKeysMaxSpeed = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MouseKeysCurve = int16(r.Read2b())
	if r.Err() != nil {
		return r.Err()
	}

	v.AccessXOption = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.AccessXTimeout = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.AccessXTimeoutOptionsMask = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.AccessXTimeoutOptionsValues = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(2)
	if r.Err() != nil {
		return r.Err()
	}

	v.AccessXTimeoutMask = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.AccessXTimeoutValues = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.EnabledControls = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	var err error
	v.PerKeyRepeat, err = r.ReadBytes(ConstPerKeyBitArraySize)
	return err
}

type ModDef struct {
	// RealMods | VirtualMods mapped to real modifiers
	Mask uint8
	//  real modifier bits
	RealMods uint8
	// virtual modifier bits
	VirtualMods uint16
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

func readKeyType(r *x.Reader, v *KeyType) error {
	v.ModsMask = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ModsMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ModsVMods = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NumLevels = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NMapEntries = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.HasPreserve = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	if v.NMapEntries > 0 {
		v.Map = make([]KTMapEntry, v.NMapEntries)
		for i := 0; i < int(v.NMapEntries); i++ {
			err := readKTMapEntry(r, &v.Map[i])
			if err != nil {
				return err
			}
		}

		if v.HasPreserve {
			v.Preserve = make([]ModDef, v.NMapEntries)
			for i := 0; i < int(v.NMapEntries); i++ {
				err := readModDef(r, &v.Preserve[i])
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// key type map entry
type KTMapEntry struct {
	Active    bool
	ModsMask  uint8
	Level     uint8
	ModsMods  uint8
	ModsVMods uint16
}

func readModDef(r *x.Reader, v *ModDef) error {
	v.Mask = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.RealMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.VirtualMods = r.Read2b()
	return r.Err()
}

func readKTMapEntry(r *x.Reader, v *KTMapEntry) error {
	v.Active = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.ModsMask = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Level = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ModsMods = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ModsVMods = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(2)
	return r.Err()
}

type KeySymMap struct {
	KtIndex   []uint8 // length 4
	GroupInfo uint8
	Width     uint8
	NSyms     uint16
	Syms      []x.Keysym
}

func readKeySymMap(r *x.Reader, v *KeySymMap) error {
	v.KtIndex = r.MustReadBytes(4)
	if r.Err() != nil {
		return r.Err()
	}

	v.GroupInfo = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Width = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NSyms = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	if v.NSyms > 0 {
		v.Syms = make([]x.Keysym, v.NSyms)
		for i := 0; i < int(v.NSyms); i++ {
			v.Syms[i] = x.Keysym(r.Read4b())
			if r.Err() != nil {
				return r.Err()
			}
		}
	}
	return nil
}

type Action interface {
	Type() uint8
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

func readAction(r *x.Reader) (Action, error) {
	type0 := r.Read1b()
	if r.Err() != nil {
		return nil, r.Err()
	}

	data := r.MustReadBytes(7)
	if r.Err() != nil {
		return nil, r.Err()
	}

	dataR := x.NewReaderFromData(data)
	if type0 <= SATypeDeviceValuator {
		action := readActionFuncArray[type0](dataR)
		return action, nil
	} else {
		// out of range
		return UnknownAction{type0: type0, data: data}, nil
	}
}

type SANoAction struct {
}

func (SANoAction) Type() uint8 {
	return SATypeNoAction
}

func readSANoAction(r *x.Reader) Action {
	return SANoAction{}
}

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

type SALatchMods SASetMods

func (SALatchMods) Type() uint8 {
	return SATypeLatchMods
}

func readSALatchMods(r *x.Reader) Action {
	var v SALatchMods
	readSASetModsAux(r, (*SASetMods)(&v))
	return v
}

type SALockMods SASetMods

func (SALockMods) Type() uint8 {
	return SATypeLockMods
}

func readSALockMods(r *x.Reader) Action {
	var v SALockMods
	readSASetModsAux(r, (*SASetMods)(&v))
	return v
}

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

type SALatchGroup SASetGroup

func (SALatchGroup) Type() uint8 {
	return SATypeLatchGroup
}

func readSALatchGroup(r *x.Reader) Action {
	var v SALatchGroup
	readSASetGroupAux(r, (*SASetGroup)(&v))
	return v
}

type SALockGroup SASetGroup

func (SALockGroup) Type() uint8 {
	return SATypeLockGroup
}

func readSALockGroup(r *x.Reader) Action {
	var v SALockGroup
	readSASetGroupAux(r, (*SASetGroup)(&v))
	return v
}

type SAMovePtr struct {
	Flags uint8
	X     int16
	Y     int16
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

func (SAMovePtr) Type() uint8 {
	return SATypeMovePtr
}

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

type SATerminate struct {
}

func (SATerminate) Type() uint8 {
	return SATypeTerminate
}

func readSATerminate(r *x.Reader) Action {
	return SATerminate{}
}

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

type SALockControls SASetControls

func readSALockControls(r *x.Reader) Action {
	var v SALockControls
	readSASetControlsAux(r, (*SASetControls)(&v))
	return v
}

func (SALockControls) Type() uint8 {
	return SATypeLockControls
}

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

type UnknownAction struct {
	type0 uint8
	data  []byte // length 7
}

func (v UnknownAction) Type() uint8 {
	return v.type0
}

type SetBehavior struct {
	Keycode  x.Keycode
	Behavior Behavior
}

func readSetBehavior(r *x.Reader, v *SetBehavior) error {
	v.Keycode = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	var err error
	v.Behavior, err = readBehavior(r)
	if err != nil {
		return err
	}

	r.ReadPad(1)
	return r.Err()
}

type Behavior interface {
	Type() uint8
}

func readBehavior(r *x.Reader) (Behavior, error) {
	type0 := r.Read1b()
	data := r.Read1b()

	switch type0 {
	case BehaviorTypeDefault:
		return DefaultBehavior{}, nil
	case BehaviorTypeLock:
		return LockBehavior{}, nil
	case BehaviorTypeRadioGroup:
		return RadioGroupBehavior{data}, nil
	case BehaviorTypeOverlay1:
		return Overlay1Behavior{x.Keycode(data)}, nil
	case BehaviorTypeOverlay2:
		return Overlay2Behavior{x.Keycode(data)}, nil
	case BehaviorTypePermamentLock:
		return PermamentLockBehavior{}, nil
	case BehaviorTypePermamentRadioGroup:
		return PermamentRadioGroupBehavior{data}, nil
	case BehaviorTypePermamentOverlay1:
		return PermamentOverlay1Behavior{}, nil
	case BehaviorTypePermamentOverlay2:
		return PermamentOverlay1Behavior{}, nil
	default:
		return UnknownBehavior{type0, data}, nil
	}
}

type UnknownBehavior struct {
	type0 uint8
	Data  uint8
}

func (v UnknownBehavior) Type() uint8 {
	return v.type0
}

type DefaultBehavior struct {
}

func (DefaultBehavior) Type() uint8 {
	return BehaviorTypeDefault
}

type LockBehavior struct {
}

func (LockBehavior) Type() uint8 {
	return BehaviorTypeLock
}

type RadioGroupBehavior struct {
	Group uint8
}

func (RadioGroupBehavior) Type() uint8 {
	return BehaviorTypeRadioGroup
}

type Overlay1Behavior struct {
	Key x.Keycode
}

func (Overlay1Behavior) Type() uint8 {
	return BehaviorTypeOverlay1
}

type Overlay2Behavior Overlay1Behavior

func (Overlay2Behavior) Type() uint8 {
	return BehaviorTypeOverlay2
}

type PermamentLockBehavior LockBehavior

func (PermamentLockBehavior) Type() uint8 {
	return BehaviorTypePermamentLock
}

type PermamentRadioGroupBehavior RadioGroupBehavior

func (PermamentRadioGroupBehavior) Type() uint8 {
	return BehaviorTypePermamentRadioGroup
}

type PermamentOverlay1Behavior Overlay1Behavior

func (PermamentOverlay1Behavior) Type() uint8 {
	return BehaviorTypePermamentOverlay1
}

type PermamentOverlay2Behavior Overlay1Behavior

func (PermamentOverlay2Behavior) Type() uint8 {
	return BehaviorTypePermamentOverlay2
}

type SetExplicit struct {
	Keycode  x.Keycode
	Explicit uint8
}

func readSetExplicit(r *x.Reader) (SetExplicit, error) {
	var v SetExplicit
	v.Keycode = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return SetExplicit{}, r.Err()
	}

	v.Explicit = r.Read1b()
	if r.Err() != nil {
		return SetExplicit{}, r.Err()
	}

	return v, nil
}

type KeyModMap struct {
	Keycode x.Keycode
	Mods    uint8
}

func readKeyModMap(r *x.Reader) (KeyModMap, error) {
	var v KeyModMap
	v.Keycode = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return KeyModMap{}, r.Err()
	}

	v.Mods = r.Read1b()
	if r.Err() != nil {
		return KeyModMap{}, r.Err()
	}

	return v, nil
}

type KeyVModMap struct {
	Keycode x.Keycode
	VMods   uint16
}

func readKeyVModMap(r *x.Reader) (KeyVModMap, error) {
	var v KeyVModMap
	v.Keycode = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return KeyVModMap{}, r.Err()
	}

	r.ReadPad(1)
	if r.Err() != nil {
		return KeyVModMap{}, r.Err()
	}

	v.VMods = r.Read2b()
	if r.Err() != nil {
		return KeyVModMap{}, r.Err()
	}

	return v, nil
}

func readGetMapReply(r *x.Reader, v *GetMapReply) error {
	v.DeviceID, _ = r.ReadReplyHeader()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(2)
	if r.Err() != nil {
		return r.Err()
	}

	v.MinKeyCode = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.MaxKeyCode = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Present = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstType = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NTypes = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.TotalTypes = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstKeySym = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.TotalSyms = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NKeySyms = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstKeyAction = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.TotalActions = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.NKeyActions = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstKeyBehavior = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.NKeyBehaviors = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.TotalKeyBehaviors = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstKeyExplicit = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.NKeyExplicit = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.TotalKeyBehaviors = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstModMapKey = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.NModMapKeys = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.TotalModMapKeys = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.FirstVModMapKey = x.Keycode(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	v.NVModMapKeys = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.TotalVModMapKeys = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.VirtualMods = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

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
			_, v.ActionsCount, err = r.ReadBytesWithPad(int(v.NKeyActions))
			if err != nil {
				return err
			}
		}

		if v.TotalActions > 0 {
			v.Actions = make([]Action, v.TotalActions)
			for i := 0; i < int(v.TotalActions); i++ {
				var err error
				v.Actions[i], err = readAction(r)
				if err != nil {
					return err
				}
			}
		}
	}

	if v.Present&MapPartKeyBehaviors != 0 && v.TotalKeyBehaviors > 0 {
		v.Behaviors = make([]SetBehavior, v.TotalKeyBehaviors)
		for i := 0; i < int(v.TotalKeyBehaviors); i++ {
			err := readSetBehavior(r, &v.Behaviors[i])
			if err != nil {
				return err
			}
		}
	}

	if v.Present&MapPartVirtualMods != 0 {
		length := x.PopCount(int(v.VirtualMods))
		if length > 0 {
			var err error
			_, v.VMods, err = r.ReadBytesWithPad(length)
			if err != nil {
				return err
			}

		}
	}

	if v.Present&MapPartExplicitComponents != 0 && v.TotalKeyExplicit > 0 {
		v.Explicit = make([]SetExplicit, v.TotalKeyExplicit)
		for i := 0; i < int(v.TotalKeyExplicit); i++ {
			var err error
			v.Explicit[i], err = readSetExplicit(r)
			if err != nil {
				return err
			}
		}

		r.ReadPad(x.Pad(int(v.TotalKeyExplicit) * 2))
		if r.Err() != nil {
			return r.Err()
		}
	}

	if v.Present&MapPartModifierMap != 0 && v.TotalModMapKeys > 0 {
		v.ModMap = make([]KeyModMap, v.TotalModMapKeys)
		for i := 0; i < int(v.TotalModMapKeys); i++ {
			var err error
			v.ModMap[i], err = readKeyModMap(r)
			if err != nil {
				return err
			}
		}

		r.ReadPad(x.Pad(int(v.TotalModMapKeys) * 2))
		if r.Err() != nil {
			return r.Err()
		}
	}

	if v.Present&MapPartVirtualModMap != 0 && v.TotalVModMapKeys > 0 {
		v.VModMap = make([]KeyVModMap, v.TotalVModMapKeys)
		for i := 0; i < int(v.TotalVModMapKeys); i++ {
			var err error
			v.VModMap[i], err = readKeyVModMap(r)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
