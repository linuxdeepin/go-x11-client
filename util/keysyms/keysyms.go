package keysyms

// The algorithm is copied from xcb-util-keysyms

import (
	x "github.com/linuxdeepin/go-x11-client"
)

func StringToKeysym(str string) x.Keysym {
	sym, ok := EngKeysymMap[str]
	if !ok {
		return XK_VoidSymbol
	}
	return sym
}

func KeysymToString(sym x.Keysym) string {
	return KeysymEngMap[sym]
}

type KeySymbols struct {
	conn                   *x.Conn
	minKeycode, maxKeycode x.Keycode

	cookie   x.GetKeyboardMappingCookie
	reply    *x.GetKeyboardMappingReply
	hasValue bool
}

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
	return ks
}

func (ks *KeySymbols) getReply() (*x.GetKeyboardMappingReply, error) {
	if ks.hasValue {
		return ks.reply, nil
	}
	// no reply, but cookie is 0
	if ks.cookie == 0 {
		return nil, nil
	}
	reply, err := ks.cookie.Reply(ks.conn)
	ks.cookie = 0
	if err != nil {
		return nil, err
	}
	ks.reply = reply
	ks.hasValue = true
	return reply, nil
}

func (ks *KeySymbols) RefreshKeyboardMapping(event *x.MappingNotifyEvent) bool {
	if event.Request == x.MappingKeyboard && ks != nil {
		if ks.hasValue {
			ks.cookie = x.GetKeyboardMapping(ks.conn, ks.minKeycode, uint8(ks.maxKeycode-ks.minKeycode+1))
			ks.hasValue = false
		}

		return true
	}
	return false
}

func (ks *KeySymbols) GetKeycode(sym x.Keysym) []x.Keycode {
	kbdMapping, _ := ks.getReply()
	if kbdMapping == nil {
		return nil
	}
	return getKeycode(kbdMapping, ks.minKeycode, ks.maxKeycode, sym)
}

func (ks *KeySymbols) GetKeysym(code x.Keycode, col int) x.Keysym {
	kbdMapping, _ := ks.getReply()
	if kbdMapping == nil {
		return x.NoSymbol
	}
	return getKeysym(kbdMapping, ks.minKeycode, ks.maxKeycode, code, col)
}

func getKeycode(kbdMapping *x.GetKeyboardMappingReply, minKeycode, maxKeycode x.Keycode, sym x.Keysym) (result []x.Keycode) {
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
