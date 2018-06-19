package icccm

import (
	"errors"
	"strings"

	x "github.com/linuxdeepin/go-x11-client"
)

const getPropertyMaxLength = 0xffff

func getCardinalsFromReply(r *x.GetPropertyReply) ([]uint32, error) {
	if r.Format != 32 {
		return nil, errors.New("bad reply")
	}
	count := len(r.Value) / 4
	ret := make([]uint32, count)
	rdr := x.NewReaderFromData(r.Value)
	for i := 0; i < count; i++ {
		ret[i] = uint32(rdr.Read4b())
	}
	return ret, nil
}

func getWindowFromReply(r *x.GetPropertyReply) (x.Window, error) {
	if r.Format != 32 || len(r.Value) != 4 {
		return 0, errors.New("bad reply")
	}
	return x.Window(x.Get32(r.Value)), nil
}

type GetWindowCookie x.GetPropertyCookie

func (cookie GetWindowCookie) Reply(c *Conn) (x.Window, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return 0, err
	}
	if reply.Type != x.AtomWindow {
		return 0, errors.New("bad reply")
	}
	return getWindowFromReply(reply)
}

func getWindowsFromReply(r *x.GetPropertyReply) ([]x.Window, error) {
	if r.Format != 32 {
		return nil, errors.New("bad reply")
	}
	count := len(r.Value) / 4
	ret := make([]x.Window, count)
	rdr := x.NewReaderFromData(r.Value)
	for i := 0; i < count; i++ {
		ret[i] = x.Window(rdr.Read4b())
	}
	return ret, nil
}

type GetWindowsCookie x.GetPropertyCookie

func (cookie GetWindowsCookie) Reply(c *Conn) ([]x.Window, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return nil, err
	}
	if reply.Type != x.AtomWindow {
		return nil, errors.New("bad reply")
	}
	return getWindowsFromReply(reply)
}

func getAtomsFromReply(r *x.GetPropertyReply) ([]x.Atom, error) {
	if r.Format != 32 {
		return nil, errors.New("bad reply")
	}
	count := len(r.Value) / 4
	ret := make([]x.Atom, count)
	rdr := x.NewReaderFromData(r.Value)
	for i := 0; i < count; i++ {
		ret[i] = x.Atom(rdr.Read4b())
	}
	return ret, nil
}

type GetAtomsCookie x.GetPropertyCookie

func (cookie GetAtomsCookie) Reply(c *Conn) ([]x.Atom, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return nil, err
	}
	if reply.Type != x.AtomAtom {
		return nil, errors.New("bad reply")
	}
	return getAtomsFromReply(reply)
}

/* WM_NAME */
func (c *Conn) GetWMName(window x.Window) GetTextCookie {
	return c.getTextProperty(window, x.AtomWMName)
}

func (c *Conn) GetWMNameUnchecked(window x.Window) GetTextCookie {
	return c.getTextPropertyUnchecked(window, x.AtomWMName)
}

/* WM_ICON_NAME */
func (c *Conn) GetWMIconName(window x.Window) GetTextCookie {
	return c.getTextProperty(window, x.AtomWMIconName)
}

func (c *Conn) GetWMIconNameUnchecked(window x.Window) GetTextCookie {
	return c.getTextPropertyUnchecked(window, x.AtomWMIconName)
}

/* WM_COLORMAP_WINDOWS */
func (c *Conn) GetWMColormapWindows(window x.Window) GetWindowsCookie {
	cookie := x.GetProperty(c.conn, false, window, c.GetAtom("WM_COLORMAP_WINDOWS"),
		x.AtomWindow, 0, getPropertyMaxLength)
	return GetWindowsCookie(cookie)
}

/* WM_CLIENT_MACHINE  */
func (c *Conn) GetWMClientMachine(window x.Window) GetTextCookie {
	return c.getTextProperty(window, x.AtomWMClientMachine)
}

func (c *Conn) GetWMClientMachineUnchecked(window x.Window) GetTextCookie {
	return c.getTextPropertyUnchecked(window, x.AtomWMClientMachine)
}

/* WM_CLASS */
type GetWMClassCookie x.GetPropertyCookie

func (cookie GetWMClassCookie) Reply(c *Conn) (WMClass, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return WMClass{}, err
	}
	return getWMClassFromReply(reply)
}

func getWMClassFromReply(reply *x.GetPropertyReply) (WMClass, error) {
	if reply.Type != x.AtomString || reply.Format != 8 {
		return WMClass{}, errors.New("bad reply")
	}

	data := reply.Value
	var parts [][]byte
	sstart := 0
	for i, c := range data {
		if c == 0 {
			parts = append(parts, data[sstart:i])
			sstart = i + 1
		}
	}
	if sstart < len(data) {
		parts = append(parts, data[sstart:])
	}

	if len(parts) != 2 {
		return WMClass{}, errors.New("length of parts is not 2")
	}

	// convert parts to utf8str
	instance, err := convertLatin1ToUTF8(parts[0])
	if err != nil {
		return WMClass{}, err
	}
	class, err := convertLatin1ToUTF8(parts[1])
	if err != nil {
		return WMClass{}, err
	}
	return WMClass{
		Instance: instance,
		Class:    class,
	}, nil
}

type WMClass struct {
	Class, Instance string
}

func (c *Conn) GetWMClass(window x.Window) GetWMClassCookie {
	cookie := x.GetProperty(c.conn, false, window, x.AtomWMClass, x.AtomString, 0, 2048)
	return GetWMClassCookie(cookie)
}

/* WM_TRANSIENT_FOR */
func (c *Conn) GetWMTransientFor(window x.Window) GetWindowCookie {
	cookie := x.GetProperty(c.conn, false, window, x.AtomWMTransientFor, x.AtomWindow, 0, 1)
	return GetWindowCookie(cookie)
}

/* WM_SIZE_HINTS */
type WMSizeHints struct {
	Flags                 SizeHintsFlags
	X                     int32
	Y                     int32
	Width, Height         uint32
	MinWidth, MinHeight   uint32
	MaxWidth, MaxHeight   uint32
	WidthInc, HeightInc   uint32
	MinAspect             AspectRatio
	MaxAspect             AspectRatio
	BaseWidth, BaseHeight uint32
	WinGravity            uint32
}

type AspectRatio struct {
	Numerator, Denominator uint32
}

// WM_SIZE_HINTS.flags bit definitions
const (
	SizeHintsFlagUSPosition = 1 << iota
	SizeHintsFlagUSSize
	SizeHintsFlagPPosition
	SizeHintsFlagPSize
	SizeHintsFlagPMinSize
	SizeHintsFlagPMaxSize
	SizeHintsFlagPResizeInc
	SizeHintsFlagPAspect
	SizeHintsFlagPBaseSize
	SizeHintsFlagPWinGravity
)

type SizeHintsFlags uint32

func (flags SizeHintsFlags) String() string {
	names := []string{
		"USPosition",
		"USSize",
		"PPostion",
		"PSize",
		"PMinSize",
		"PMaxSize",
		"PResizeInc",
		"PAspect",
		"PBaseSize",
		"PWinGravity",
	}
	var ret []string
	for i, name := range names {
		if flags&(1<<uint32(i)) != 0 {
			ret = append(ret, name)
		}
	}
	return "[" + strings.Join(ret, ",") + "]"
}

type GetWMSizeHintsCookie x.GetPropertyCookie

func (cookie GetWMSizeHintsCookie) Reply(c *Conn) (*WMSizeHints, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return nil, err
	}

	if reply.Type != x.AtomWMSizeHints {
		return nil, errors.New("bad reply")
	}
	return getWMSizeHintsFromReply(reply)
}

func getWMSizeHintsFromReply(reply *x.GetPropertyReply) (*WMSizeHints, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return nil, err
	}

	if len(list) != wmSizeHintsElements {
		return nil, errors.New("length of list is incorrect")
	}

	var v WMSizeHints
	v.Flags = SizeHintsFlags(list[0])

	v.X = int32(list[1])
	v.Y = int32(list[2])
	v.Width = list[3]
	v.Height = list[4]

	v.MinWidth = list[5]
	v.MinHeight = list[6]

	v.MaxWidth = list[7]
	v.MaxHeight = list[8]

	v.WidthInc = list[9]
	v.HeightInc = list[10]

	v.MinAspect.Numerator = list[11]
	v.MinAspect.Denominator = list[12]

	v.MaxAspect.Numerator = list[13]
	v.MaxAspect.Denominator = list[14]

	v.BaseWidth = list[15]
	v.BaseHeight = list[16]

	v.WinGravity = list[17]
	return &v, nil
}

const wmSizeHintsElements = 18

/* WM_NORMAL_HINTS */
func (c *Conn) GetWMNormalHints(window x.Window) GetWMSizeHintsCookie {
	cookie := x.GetProperty(c.conn, false, window, x.AtomWMNormalHints, x.AtomWMSizeHints, 0, wmSizeHintsElements)
	return GetWMSizeHintsCookie(cookie)
}

/* WM_HINTS */
type WMHints struct {
	Flags        HintsFlags
	Input        uint32
	InitialState uint32
	IconPixmap   x.Pixmap
	IconWindow   x.Window
	IconX, IconY int32
	IconMask     x.Pixmap
	WindowGroup  x.Window
}

type HintsFlags uint32

const (
	HintsFlagInput = 1 << iota
	HintsFlagState
	HintsFlagIconPixmap
	HintsFlagIconWindow
	HintsFlagIconPostion
	HintsFlagIconMask
	HinstFlagWindowGroup
	HintsFlagMessage // this bit is obsolete
	HinstFlagUrgency
)

func (flags HintsFlags) String() string {
	names := []string{
		"Input",
		"State",
		"IconPixmap",
		"IconWindow",
		"IconPosition",
		"IconMask",
		"WindowGroup",
		"Message",
		"Urgency",
	}
	var ret []string
	for i, name := range names {
		if flags&(1<<uint32(i)) != 0 {
			ret = append(ret, name)
		}
	}
	return "[" + strings.Join(ret, ",") + "]"
}

const wmHintsElements = 9

type GetWMHintsCookie x.GetPropertyCookie

func (cookie GetWMHintsCookie) Reply(c *Conn) (*WMHints, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return nil, err
	}

	if reply.Type != x.AtomWMHints {
		return nil, errors.New("bad reply")
	}
	return getWMHintsFromReply(reply)
}

func getWMHintsFromReply(reply *x.GetPropertyReply) (*WMHints, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return nil, err
	}

	if len(list) != wmHintsElements {
		return nil, errors.New("length of list is incorrect")
	}

	var v WMHints
	v.Flags = HintsFlags(list[0])
	v.Input = list[1]
	v.InitialState = list[2]
	v.IconPixmap = x.Pixmap(list[3])
	v.IconWindow = x.Window(list[4])
	v.IconX = int32(list[5])
	v.IconY = int32(list[6])
	v.IconMask = x.Pixmap(list[7])
	v.WindowGroup = x.Window(list[8])
	return &v, nil
}

func (c *Conn) GetWMHints(window x.Window) GetWMHintsCookie {
	cookie := x.GetProperty(c.conn, false, window, x.AtomWMHints, x.AtomWMHints,
		0, wmHintsElements)
	return GetWMHintsCookie(cookie)
}

/* WM_PROTOCOLS */
func (c *Conn) GetWMProtocols(window x.Window) GetAtomsCookie {
	cookie := x.GetProperty(c.conn, false, window, c.GetAtom("WM_PROTOCOLS"), x.AtomAtom,
		0, getPropertyMaxLength)
	return GetAtomsCookie(cookie)
}

/* WM_STATE */
type WMState struct {
	State  uint32
	Window x.Window
}

const wmStateElements = 2

const (
	StateWithdrawn = 0
	StateNormal    = 1
	StateIconic    = 3
)

type GetWMStateCookie x.GetPropertyCookie

func (cookie GetWMStateCookie) Reply(c *Conn) (WMState, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return WMState{}, err
	}
	return getWMStateFromReply(reply)
}

func getWMStateFromReply(reply *x.GetPropertyReply) (WMState, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return WMState{}, err
	}

	if len(list) != wmStateElements {
		return WMState{}, errors.New("length of list is incorrect")
	}

	var v WMState
	v.State = list[0]
	v.Window = x.Window(list[1])
	return v, nil
}

func (c *Conn) GetWMState(window x.Window) GetWMStateCookie {
	atomWMState := c.GetAtom("WM_STATE")
	cookie := x.GetProperty(c.conn, false, window, atomWMState, atomWMState,
		0, wmStateElements)
	return GetWMStateCookie(cookie)
}

func (c *Conn) sendClientMessage(win, dest x.Window, msgType x.Atom, pArray *[5]uint32) x.VoidCookie {
	var data x.ClientMessageData
	data.SetData32(pArray)
	event := x.ClientMessageEvent{
		Format: 32,
		Window: win,
		Type:   msgType,
		Data:   data,
	}
	w := x.NewWriter()
	x.WriteClientMessageEvent(w, &event)
	const evMask = x.EventMaskSubstructureNotify | x.EventMaskSubstructureRedirect
	return x.SendEventChecked(c.conn, false, dest, evMask, w.Bytes())
}

func (c *Conn) RequestChangeWMState(window x.Window, state uint32) x.VoidCookie {
	array := [5]uint32{state}
	return c.sendClientMessage(window, c.GetRootWin(), c.GetAtom("WM_CHANGE_STATE"), &array)
}

/* WM_ICON_SIZE */
type WMIconSize struct {
	MinWidth, MinHeight uint32
	MaxWidth, MaxHeight uint32
	WidthInc, HeightInc uint32
}

const wmIconSizeElements = 6

type GetWMIconSizeCookie x.GetPropertyCookie

func (cookie GetWMIconSizeCookie) Reply(c *Conn) (*WMIconSize, error) {
	reply, err := x.GetPropertyCookie(cookie).Reply(c.conn)
	if err != nil {
		return nil, err
	}
	return getWMIconSizeFromReply(reply)
}

func getWMIconSizeFromReply(reply *x.GetPropertyReply) (*WMIconSize, error) {
	list, err := getCardinalsFromReply(reply)
	if err != nil {
		return nil, err
	}

	if len(list) != wmIconSizeElements {
		return nil, errors.New("length of list is incorrect")
	}

	var v WMIconSize
	v.MinWidth = list[0]
	v.MinHeight = list[1]
	v.MaxWidth = list[2]
	v.MaxHeight = list[3]
	v.WidthInc = list[4]
	v.HeightInc = list[5]
	return &v, nil
}

func (c *Conn) GetWMIconSize() GetWMIconSizeCookie {
	cookie := x.GetProperty(c.conn, false, c.GetRootWin(), x.AtomWMIconSize, x.AtomCardinal, 0, wmIconSizeElements)
	return GetWMIconSizeCookie(cookie)
}
