package keysyms

// The algorithm is copied from xcb-util-keysyms

import (
	"errors"
	"sync"

	x "github.com/linuxdeepin/go-x11-client"
)

func StringToKeysym(str string) (x.Keysym, bool) {
	sym, ok := EngKeysymMap[str]
	if !ok {
		return XK_VoidSymbol, false
	}
	return sym, true
}

func StringToVisibleChar(str string) (rune, bool) {
	sym, ok := EngKeysymMap[str]
	if !ok {
		return 0, false
	}
	char, ok := KeysymVisibleCharMap[sym]
	return char, ok
}

func KeysymToString(sym x.Keysym) (string, bool) {
	if sym == x.NoSymbol {
		return "NoSymbol", true
	}
	eng, ok := KeysymEngMap[sym]
	return eng, ok
}

const (
	ModMaskAlt        = x.ModMask1
	ModMaskNumLock    = x.ModMask2
	ModMaskSuper      = x.ModMask4
	ModMaskModeSwitch = x.ModMask5
	ModMaskShift      = x.ModMaskShift
	ModMaskCapsLock   = x.ModMaskLock
	ModMaskControl    = x.ModMaskControl
)

var LockMods = []uint16{
	ModMaskCapsLock,
	ModMaskNumLock,
	ModMaskCapsLock | ModMaskNumLock,
}

type KeySymbols struct {
	mu                     sync.Mutex
	conn                   *x.Conn
	minKeycode, maxKeycode x.Keycode

	cookie     x.GetKeyboardMappingCookie
	oldCookies []x.GetKeyboardMappingCookie
	kbdMap     *x.GetKeyboardMappingReply
	tag        int
}

const (
	tagCookie = 0
	tagReply  = 1
)

func NewKeySymbols(conn *x.Conn) *KeySymbols {
	setup := conn.GetSetup()
	minKeycode := setup.MinKeycode
	maxKeycode := setup.MaxKeycode
	ks := &KeySymbols{
		conn:       conn,
		minKeycode: minKeycode,
		maxKeycode: maxKeycode,
	}

	ks.cookie = x.GetKeyboardMapping(conn, minKeycode, uint8(maxKeycode-minKeycode+1))
	ks.tag = tagCookie
	return ks
}

func (ks *KeySymbols) getKbdMap() *x.GetKeyboardMappingReply {
	ks.mu.Lock()
	defer ks.mu.Unlock()

	if ks.tag == tagReply {
		return ks.kbdMap
	}
	// no reply, but cookie is 0
	if ks.cookie == 0 {
		return nil
	}

	reply, err := ks.cookie.Reply(ks.conn)
	//println("cookie reply", ks.cookie)
	// use cookie once
	ks.cookie = 0

	// discard old replies
	for _, ck := range ks.oldCookies {
		_, _ = ck.Reply(ks.conn)
	}
	ks.oldCookies = nil

	if err != nil {
		return nil
	}
	ks.tag = tagReply
	ks.kbdMap = reply
	return reply
}

func (ks *KeySymbols) RefreshKeyboardMapping(event *x.MappingNotifyEvent) bool {
	if event.Request == x.MappingKeyboard {
		ks.mu.Lock()

		if ks.cookie > 0 {
			ks.oldCookies = append(ks.oldCookies, ks.cookie)
		}
		ks.cookie = x.GetKeyboardMapping(ks.conn, ks.minKeycode, uint8(ks.maxKeycode-ks.minKeycode+1))
		//println("send GetKeyboardMapping request", ks.cookie)
		ks.tag = tagCookie
		ks.mu.Unlock()
		return true
	}
	return false
}

func (ks *KeySymbols) GetKeycodes(sym x.Keysym) []x.Keycode {
	kbdMap := ks.getKbdMap()
	if kbdMap == nil {
		return nil
	}
	return getKeycodes(kbdMap, ks.minKeycode, ks.maxKeycode, sym)
}

func (ks *KeySymbols) GetKeysym(code x.Keycode, col int) x.Keysym {
	kbdMap := ks.getKbdMap()
	if kbdMap == nil {
		return x.NoSymbol
	}
	return getKeysym(kbdMap, ks.minKeycode, ks.maxKeycode, code, col)
}

func (ks *KeySymbols) StringToKeycodes(str string) ([]x.Keycode, error) {
	sym, ok := EngKeysymMap[str]
	if !ok {
		return nil, errors.New("failed to get keysym")
	}

	return ks.GetKeycodes(sym), nil
}

func getKeycodes(kbdMapping *x.GetKeyboardMappingReply, minKeycode, maxKeycode x.Keycode, sym x.Keysym) (result []x.Keycode) {
	per := kbdMapping.KeysymsPerKeycode
	// i keycode
	// j col
	for i := minKeycode; i > 0 && i <= maxKeycode; i++ {
		//fmt.Println("i is ", i)
		for j := 0; j < int(per); j++ {
			ks := getKeysym(kbdMapping, minKeycode, maxKeycode, i, j)
			if ks == sym {
				result = append(result, i)
				break
			}
		}
	}
	return
}

func (ks *KeySymbols) LookupString(keycode x.Keycode, modifier uint16) (string, bool) {
	sym := ks.translateKey(keycode, modifier)
	return KeysymToString(sym)
}

func (ks *KeySymbols) translateKey(keycode x.Keycode, modifiers uint16) (result x.Keysym) {

	if (keycode < ks.minKeycode) || (keycode > ks.maxKeycode) {
		return x.NoSymbol
	}

	kbdMap := ks.getKbdMap()
	if kbdMap == nil {
		return x.NoSymbol
	}

	syms := kbdMap.Keysyms
	per := int(kbdMap.KeysymsPerKeycode)
	syms = syms[int(keycode-ks.minKeycode)*per:]

	for per > 2 && syms[per-1] == x.NoSymbol {
		per--
	}
	if (per > 2) && (modifiers&ModMaskModeSwitch) != 0 {
		syms = syms[2:]
		per -= 2
	}

	shiftOn := (modifiers & ModMaskShift) != 0
	capsLockOn := (modifiers & ModMaskCapsLock) != 0

	if (modifiers&ModMaskNumLock) != 0 &&
		(per > 1 && (IsKeypadKey(syms[1]) || IsPrivateKeypadKey(syms[1]))) {
		// num_lock is on and key is keypad key
		if shiftOn {
			result = syms[0]
		} else {
			// shift is off
			// syms[1] is number
			result = syms[1]
		}
	} else {
		symArr := [2]x.Keysym{syms[0]}
		if per == 1 || syms[1] == x.NoSymbol {
			_, usym := ConvertCase(syms[0])
			symArr[1] = usym
		} else {
			symArr[1] = syms[1]
		}

		idx := 0
		onlyCaseDiff := false
		_, upperSym := ConvertCase(symArr[0])
		if upperSym == symArr[1] {
			onlyCaseDiff = true
		}

		if capsLockOn && onlyCaseDiff {
			idx = 1
		}
		if shiftOn {
			if idx == 0 {
				idx = 1
			} else {
				idx = 0
			}
		}
		result = symArr[idx]
	}

	if result == XK_VoidSymbol {
		result = x.NoSymbol
	}
	return result
}

func getKeysym(kbdMapping *x.GetKeyboardMappingReply, minKeycode, maxKeycode, keycode x.Keycode, col int) x.Keysym {
	syms := kbdMapping.Keysyms
	per := kbdMapping.KeysymsPerKeycode

	if col < 0 || (col >= int(per) && col > 3) ||
		keycode < minKeycode ||
		keycode > maxKeycode {
		return x.NoSymbol
	}

	syms = syms[int(keycode-minKeycode)*int(per):]

	if col < 4 {
		// col is in 0,1,2,3
		if col > 1 {
			for per > 2 && syms[per-1] == x.NoSymbol {
				per--
			}
			if per < 3 {
				col -= 2
			}
		}

		if (int(per) <= (col | 1)) || (syms[col|1] == x.NoSymbol) {
			lsym, usym := ConvertCase(syms[col&^1])
			if (col & 1) == 0 {
				return lsym
			} else if usym == lsym {
				return x.NoSymbol
			} else {
				return usym
			}
		}
	}
	return syms[col]
}

func ConvertCase(sym x.Keysym) (lower, upper x.Keysym) {
	lower = sym
	upper = sym

	switch sym >> 8 {
	case 0: /* Latin 1 */
		if sym >= XK_A && sym <= XK_Z {
			lower += XK_a - XK_A
		} else if sym >= XK_a && sym <= XK_z {
			upper -= XK_a - XK_A
		} else if sym >= XK_Agrave && sym <= XK_Odiaeresis {
			lower += XK_agrave - XK_Agrave
		} else if sym >= XK_agrave && sym <= XK_odiaeresis {
			upper -= XK_agrave - XK_Agrave
		} else if sym >= XK_Ooblique && sym <= XK_THORN {
			lower += XK_oslash - XK_Ooblique
		} else if sym >= XK_oslash && sym <= XK_thorn {
			upper -= XK_oslash - XK_Ooblique
		}
	case 1: /* Latin 2 */
		/* Assume the KeySym is a legal value (ignore discontinuities) */
		if sym == XK_Aogonek {
			lower = XK_aogonek
		} else if sym >= XK_Lstroke && sym <= XK_Sacute {
			lower += XK_lstroke - XK_Lstroke
		} else if sym >= XK_Scaron && sym <= XK_Zacute {
			lower += XK_scaron - XK_Scaron
		} else if sym >= XK_Zcaron && sym <= XK_Zabovedot {
			lower += XK_zcaron - XK_Zcaron
		} else if sym == XK_aogonek {
			upper = XK_Aogonek
		} else if sym >= XK_lstroke && sym <= XK_sacute {
			upper -= XK_lstroke - XK_Lstroke
		} else if sym >= XK_scaron && sym <= XK_zacute {
			upper -= XK_scaron - XK_Scaron
		} else if sym >= XK_zcaron && sym <= XK_zabovedot {
			upper -= XK_zcaron - XK_Zcaron
		} else if sym >= XK_Racute && sym <= XK_Tcedilla {
			lower += XK_racute - XK_Racute
		} else if sym >= XK_racute && sym <= XK_tcedilla {
			upper -= XK_racute - XK_Racute
		}
	case 2: /* Latin 3 */
		/* Assume the KeySym is a legal value (ignore discontinuities) */
		if sym >= XK_Hstroke && sym <= XK_Hcircumflex {
			lower += XK_hstroke - XK_Hstroke
		} else if sym >= XK_Gbreve && sym <= XK_Jcircumflex {
			lower += XK_gbreve - XK_Gbreve
		} else if sym >= XK_hstroke && sym <= XK_hcircumflex {
			upper -= XK_hstroke - XK_Hstroke
		} else if sym >= XK_gbreve && sym <= XK_jcircumflex {
			upper -= XK_gbreve - XK_Gbreve
		} else if sym >= XK_Cabovedot && sym <= XK_Scircumflex {
			lower += XK_cabovedot - XK_Cabovedot
		} else if sym >= XK_cabovedot && sym <= XK_scircumflex {
			upper -= XK_cabovedot - XK_Cabovedot
		}

	case 3: /* Latin 4 */
		/* Assume the KeySym is a legal value (ignore discontinuities) */
		if sym >= XK_Rcedilla && sym <= XK_Tslash {
			lower += XK_rcedilla - XK_Rcedilla
		} else if sym >= XK_rcedilla && sym <= XK_tslash {
			upper -= XK_rcedilla - XK_Rcedilla
		} else if sym == XK_ENG {
			lower = XK_eng
		} else if sym == XK_eng {
			upper = XK_ENG
		} else if sym >= XK_Amacron && sym <= XK_Umacron {
			lower += XK_amacron - XK_Amacron
		} else if sym >= XK_amacron && sym <= XK_umacron {
			upper -= XK_amacron - XK_Amacron
		}

	case 6: /* Cyrillic */
		/* Assume the KeySym is a legal value (ignore discontinuities) */
		if sym >= XK_Serbian_DJE && sym <= XK_Serbian_DZE {
			lower -= XK_Serbian_DJE - XK_Serbian_dje
		} else if sym >= XK_Serbian_dje && sym <= XK_Serbian_dze {
			upper += XK_Serbian_DJE - XK_Serbian_dje
		} else if sym >= XK_Cyrillic_YU && sym <= XK_Cyrillic_HARDSIGN {
			lower -= XK_Cyrillic_YU - XK_Cyrillic_yu
		} else if sym >= XK_Cyrillic_yu && sym <= XK_Cyrillic_hardsign {
			upper += XK_Cyrillic_YU - XK_Cyrillic_yu
		}

	case 7: /* Greek */
		if sym >= XK_Greek_ALPHAaccent && sym <= XK_Greek_OMEGAaccent {
			lower += XK_Greek_alphaaccent - XK_Greek_ALPHAaccent
		} else if sym >= XK_Greek_alphaaccent && sym <= XK_Greek_omegaaccent &&
			sym != XK_Greek_iotaaccentdieresis && sym != XK_Greek_upsilonaccentdieresis {
			upper -= XK_Greek_alphaaccent - XK_Greek_ALPHAaccent
		} else if sym >= XK_Greek_ALPHA && sym <= XK_Greek_OMEGA {
			lower += XK_Greek_alpha - XK_Greek_ALPHA
		} else if sym >= XK_Greek_alpha && sym <= XK_Greek_omega &&
			sym != XK_Greek_finalsmallsigma {
			upper -= XK_Greek_alpha - XK_Greek_ALPHA
		}
	case 0x14: /* Armenian */
		if sym >= XK_Armenian_AYB && sym <= XK_Armenian_fe {
			lower = sym | 1
			upper = sym &^ 1
		}
	}

	return
}

/* Tests for classes of symbols */

func IsKeypadKey(sym x.Keysym) bool {
	return sym >= XK_KP_Space && sym <= XK_KP_Equal
}

func IsPrivateKeypadKey(sym x.Keysym) bool {
	return (sym >= 0x11000000) && (sym <= 0x1100FFFF)
}

func IsCursorKey(sym x.Keysym) bool {
	return sym >= XK_Home && sym <= XK_Begin
}

func IsKeypadFuncationKey(sym x.Keysym) bool {
	return sym >= XK_KP_F1 && sym <= XK_KP_F4
}

func IsFunctionKey(sym x.Keysym) bool {
	return sym >= XK_F1 && sym <= XK_F35
}

func IsMiscFunctionKey(sym x.Keysym) bool {
	return sym >= XK_Select && sym <= XK_Break
}

func IsModifierKey(sym x.Keysym) bool {
	return ((sym >= XK_Shift_L) && (sym <= XK_Hyper_R)) ||
		((sym >= XK_ISO_Lock) && (sym <= XK_ISO_Level5_Lock)) ||
		(sym == XK_Mode_switch) ||
		(sym == XK_Num_Lock)
}
