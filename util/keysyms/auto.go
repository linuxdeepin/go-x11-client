// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package keysyms

import x "github.com/linuxdeepin/go-x11-client"

const (
	//  Void symbol
	XK_VoidSymbol = 0xffffff
	//  Back space, back char
	XK_BackSpace = 0xff08
	XK_Tab       = 0xff09
	//  Linefeed, LF
	XK_Linefeed = 0xff0a
	XK_Clear    = 0xff0b
	//  Return, enter
	XK_Return = 0xff0d
	//  Pause, hold
	XK_Pause       = 0xff13
	XK_Scroll_Lock = 0xff14
	XK_Sys_Req     = 0xff15
	XK_Escape      = 0xff1b
	//  Delete, rubout
	XK_Delete = 0xffff
	//  Multi-key character compose
	XK_Multi_key         = 0xff20
	XK_Codeinput         = 0xff37
	XK_SingleCandidate   = 0xff3c
	XK_MultipleCandidate = 0xff3d
	XK_PreviousCandidate = 0xff3e
	//  Kanji, Kanji convert
	XK_Kanji = 0xff21
	//  Cancel Conversion
	XK_Muhenkan = 0xff22
	//  Start/Stop Conversion
	XK_Henkan_Mode = 0xff23
	//  Alias for Henkan_Mode
	XK_Henkan = 0xff23
	//  to Romaji
	XK_Romaji = 0xff24
	//  to Hiragana
	XK_Hiragana = 0xff25
	//  to Katakana
	XK_Katakana = 0xff26
	//  Hiragana/Katakana toggle
	XK_Hiragana_Katakana = 0xff27
	//  to Zenkaku
	XK_Zenkaku = 0xff28
	//  to Hankaku
	XK_Hankaku = 0xff29
	//  Zenkaku/Hankaku toggle
	XK_Zenkaku_Hankaku = 0xff2a
	//  Add to Dictionary
	XK_Touroku = 0xff2b
	//  Delete from Dictionary
	XK_Massyo = 0xff2c
	//  Kana Lock
	XK_Kana_Lock = 0xff2d
	//  Kana Shift
	XK_Kana_Shift = 0xff2e
	//  Alphanumeric Shift
	XK_Eisu_Shift = 0xff2f
	//  Alphanumeric toggle
	XK_Eisu_toggle = 0xff30
	//  Codeinput
	XK_Kanji_Bangou = 0xff37
	//  Multiple/All Candidate(s)
	XK_Zen_Koho = 0xff3d
	//  Previous Candidate
	XK_Mae_Koho = 0xff3e
	XK_Home     = 0xff50
	//  Move left, left arrow
	XK_Left = 0xff51
	//  Move up, up arrow
	XK_Up = 0xff52
	//  Move right, right arrow
	XK_Right = 0xff53
	//  Move down, down arrow
	XK_Down = 0xff54
	//  Prior, previous
	XK_Prior   = 0xff55
	XK_Page_Up = 0xff55
	//  Next
	XK_Next      = 0xff56
	XK_Page_Down = 0xff56
	//  EOL
	XK_End = 0xff57
	//  BOL
	XK_Begin = 0xff58
	//  Select, mark
	XK_Select = 0xff60
	XK_Print  = 0xff61
	//  Execute, run, do
	XK_Execute = 0xff62
	//  Insert, insert here
	XK_Insert = 0xff63
	XK_Undo   = 0xff65
	//  Redo, again
	XK_Redo = 0xff66
	XK_Menu = 0xff67
	//  Find, search
	XK_Find = 0xff68
	//  Cancel, stop, abort, exit
	XK_Cancel = 0xff69
	//  Help
	XK_Help  = 0xff6a
	XK_Break = 0xff6b
	//  Character set switch
	XK_Mode_switch = 0xff7e
	//  Alias for mode_switch
	XK_script_switch = 0xff7e
	XK_Num_Lock      = 0xff7f
	//  Space
	XK_KP_Space = 0xff80
	XK_KP_Tab   = 0xff89
	//  Enter
	XK_KP_Enter = 0xff8d
	//  PF1, KP_A, ...
	XK_KP_F1        = 0xff91
	XK_KP_F2        = 0xff92
	XK_KP_F3        = 0xff93
	XK_KP_F4        = 0xff94
	XK_KP_Home      = 0xff95
	XK_KP_Left      = 0xff96
	XK_KP_Up        = 0xff97
	XK_KP_Right     = 0xff98
	XK_KP_Down      = 0xff99
	XK_KP_Prior     = 0xff9a
	XK_KP_Page_Up   = 0xff9a
	XK_KP_Next      = 0xff9b
	XK_KP_Page_Down = 0xff9b
	XK_KP_End       = 0xff9c
	XK_KP_Begin     = 0xff9d
	XK_KP_Insert    = 0xff9e
	XK_KP_Delete    = 0xff9f
	//  Equals
	XK_KP_Equal    = 0xffbd
	XK_KP_Multiply = 0xffaa
	XK_KP_Add      = 0xffab
	//  Separator, often comma
	XK_KP_Separator = 0xffac
	XK_KP_Subtract  = 0xffad
	XK_KP_Decimal   = 0xffae
	XK_KP_Divide    = 0xffaf
	XK_KP_0         = 0xffb0
	XK_KP_1         = 0xffb1
	XK_KP_2         = 0xffb2
	XK_KP_3         = 0xffb3
	XK_KP_4         = 0xffb4
	XK_KP_5         = 0xffb5
	XK_KP_6         = 0xffb6
	XK_KP_7         = 0xffb7
	XK_KP_8         = 0xffb8
	XK_KP_9         = 0xffb9
	XK_F1           = 0xffbe
	XK_F2           = 0xffbf
	XK_F3           = 0xffc0
	XK_F4           = 0xffc1
	XK_F5           = 0xffc2
	XK_F6           = 0xffc3
	XK_F7           = 0xffc4
	XK_F8           = 0xffc5
	XK_F9           = 0xffc6
	XK_F10          = 0xffc7
	XK_F11          = 0xffc8
	XK_L1           = 0xffc8
	XK_F12          = 0xffc9
	XK_L2           = 0xffc9
	XK_F13          = 0xffca
	XK_L3           = 0xffca
	XK_F14          = 0xffcb
	XK_L4           = 0xffcb
	XK_F15          = 0xffcc
	XK_L5           = 0xffcc
	XK_F16          = 0xffcd
	XK_L6           = 0xffcd
	XK_F17          = 0xffce
	XK_L7           = 0xffce
	XK_F18          = 0xffcf
	XK_L8           = 0xffcf
	XK_F19          = 0xffd0
	XK_L9           = 0xffd0
	XK_F20          = 0xffd1
	XK_L10          = 0xffd1
	XK_F21          = 0xffd2
	XK_R1           = 0xffd2
	XK_F22          = 0xffd3
	XK_R2           = 0xffd3
	XK_F23          = 0xffd4
	XK_R3           = 0xffd4
	XK_F24          = 0xffd5
	XK_R4           = 0xffd5
	XK_F25          = 0xffd6
	XK_R5           = 0xffd6
	XK_F26          = 0xffd7
	XK_R6           = 0xffd7
	XK_F27          = 0xffd8
	XK_R7           = 0xffd8
	XK_F28          = 0xffd9
	XK_R8           = 0xffd9
	XK_F29          = 0xffda
	XK_R9           = 0xffda
	XK_F30          = 0xffdb
	XK_R10          = 0xffdb
	XK_F31          = 0xffdc
	XK_R11          = 0xffdc
	XK_F32          = 0xffdd
	XK_R12          = 0xffdd
	XK_F33          = 0xffde
	XK_R13          = 0xffde
	XK_F34          = 0xffdf
	XK_R14          = 0xffdf
	XK_F35          = 0xffe0
	XK_R15          = 0xffe0
	//  Left shift
	XK_Shift_L = 0xffe1
	//  Right shift
	XK_Shift_R = 0xffe2
	//  Left control
	XK_Control_L = 0xffe3
	//  Right control
	XK_Control_R = 0xffe4
	//  Caps lock
	XK_Caps_Lock = 0xffe5
	//  Shift lock
	XK_Shift_Lock = 0xffe6
	//  Left meta
	XK_Meta_L = 0xffe7
	//  Right meta
	XK_Meta_R = 0xffe8
	//  Left alt
	XK_Alt_L = 0xffe9
	//  Right alt
	XK_Alt_R = 0xffea
	//  Left super
	XK_Super_L = 0xffeb
	//  Right super
	XK_Super_R = 0xffec
	//  Left hyper
	XK_Hyper_L = 0xffed
	//  Right hyper
	XK_Hyper_R          = 0xffee
	XK_ISO_Lock         = 0xfe01
	XK_ISO_Level2_Latch = 0xfe02
	XK_ISO_Level3_Shift = 0xfe03
	XK_ISO_Level3_Latch = 0xfe04
	XK_ISO_Level3_Lock  = 0xfe05
	XK_ISO_Level5_Shift = 0xfe11
	XK_ISO_Level5_Latch = 0xfe12
	XK_ISO_Level5_Lock  = 0xfe13
	//  Alias for mode_switch
	XK_ISO_Group_Shift             = 0xff7e
	XK_ISO_Group_Latch             = 0xfe06
	XK_ISO_Group_Lock              = 0xfe07
	XK_ISO_Next_Group              = 0xfe08
	XK_ISO_Next_Group_Lock         = 0xfe09
	XK_ISO_Prev_Group              = 0xfe0a
	XK_ISO_Prev_Group_Lock         = 0xfe0b
	XK_ISO_First_Group             = 0xfe0c
	XK_ISO_First_Group_Lock        = 0xfe0d
	XK_ISO_Last_Group              = 0xfe0e
	XK_ISO_Last_Group_Lock         = 0xfe0f
	XK_ISO_Left_Tab                = 0xfe20
	XK_ISO_Move_Line_Up            = 0xfe21
	XK_ISO_Move_Line_Down          = 0xfe22
	XK_ISO_Partial_Line_Up         = 0xfe23
	XK_ISO_Partial_Line_Down       = 0xfe24
	XK_ISO_Partial_Space_Left      = 0xfe25
	XK_ISO_Partial_Space_Right     = 0xfe26
	XK_ISO_Set_Margin_Left         = 0xfe27
	XK_ISO_Set_Margin_Right        = 0xfe28
	XK_ISO_Release_Margin_Left     = 0xfe29
	XK_ISO_Release_Margin_Right    = 0xfe2a
	XK_ISO_Release_Both_Margins    = 0xfe2b
	XK_ISO_Fast_Cursor_Left        = 0xfe2c
	XK_ISO_Fast_Cursor_Right       = 0xfe2d
	XK_ISO_Fast_Cursor_Up          = 0xfe2e
	XK_ISO_Fast_Cursor_Down        = 0xfe2f
	XK_ISO_Continuous_Underline    = 0xfe30
	XK_ISO_Discontinuous_Underline = 0xfe31
	XK_ISO_Emphasize               = 0xfe32
	XK_ISO_Center_Object           = 0xfe33
	XK_ISO_Enter                   = 0xfe34
	XK_dead_grave                  = 0xfe50
	XK_dead_acute                  = 0xfe51
	XK_dead_circumflex             = 0xfe52
	XK_dead_tilde                  = 0xfe53
	//  alias for dead_tilde
	XK_dead_perispomeni      = 0xfe53
	XK_dead_macron           = 0xfe54
	XK_dead_breve            = 0xfe55
	XK_dead_abovedot         = 0xfe56
	XK_dead_diaeresis        = 0xfe57
	XK_dead_abovering        = 0xfe58
	XK_dead_doubleacute      = 0xfe59
	XK_dead_caron            = 0xfe5a
	XK_dead_cedilla          = 0xfe5b
	XK_dead_ogonek           = 0xfe5c
	XK_dead_iota             = 0xfe5d
	XK_dead_voiced_sound     = 0xfe5e
	XK_dead_semivoiced_sound = 0xfe5f
	XK_dead_belowdot         = 0xfe60
	XK_dead_hook             = 0xfe61
	XK_dead_horn             = 0xfe62
	XK_dead_stroke           = 0xfe63
	XK_dead_abovecomma       = 0xfe64
	//  alias for dead_abovecomma
	XK_dead_psili              = 0xfe64
	XK_dead_abovereversedcomma = 0xfe65
	//  alias for dead_abovereversedcomma
	XK_dead_dasia              = 0xfe65
	XK_dead_doublegrave        = 0xfe66
	XK_dead_belowring          = 0xfe67
	XK_dead_belowmacron        = 0xfe68
	XK_dead_belowcircumflex    = 0xfe69
	XK_dead_belowtilde         = 0xfe6a
	XK_dead_belowbreve         = 0xfe6b
	XK_dead_belowdiaeresis     = 0xfe6c
	XK_dead_invertedbreve      = 0xfe6d
	XK_dead_belowcomma         = 0xfe6e
	XK_dead_currency           = 0xfe6f
	XK_dead_lowline            = 0xfe90
	XK_dead_aboveverticalline  = 0xfe91
	XK_dead_belowverticalline  = 0xfe92
	XK_dead_longsolidusoverlay = 0xfe93
	XK_dead_a                  = 0xfe80
	XK_dead_A                  = 0xfe81
	XK_dead_e                  = 0xfe82
	XK_dead_E                  = 0xfe83
	XK_dead_i                  = 0xfe84
	XK_dead_I                  = 0xfe85
	XK_dead_o                  = 0xfe86
	XK_dead_O                  = 0xfe87
	XK_dead_u                  = 0xfe88
	XK_dead_U                  = 0xfe89
	XK_dead_small_schwa        = 0xfe8a
	XK_dead_capital_schwa      = 0xfe8b
	XK_dead_greek              = 0xfe8c
	XK_First_Virtual_Screen    = 0xfed0
	XK_Prev_Virtual_Screen     = 0xfed1
	XK_Next_Virtual_Screen     = 0xfed2
	XK_Last_Virtual_Screen     = 0xfed4
	XK_Terminate_Server        = 0xfed5
	XK_AccessX_Enable          = 0xfe70
	XK_AccessX_Feedback_Enable = 0xfe71
	XK_RepeatKeys_Enable       = 0xfe72
	XK_SlowKeys_Enable         = 0xfe73
	XK_BounceKeys_Enable       = 0xfe74
	XK_StickyKeys_Enable       = 0xfe75
	XK_MouseKeys_Enable        = 0xfe76
	XK_MouseKeys_Accel_Enable  = 0xfe77
	XK_Overlay1_Enable         = 0xfe78
	XK_Overlay2_Enable         = 0xfe79
	XK_AudibleBell_Enable      = 0xfe7a
	XK_Pointer_Left            = 0xfee0
	XK_Pointer_Right           = 0xfee1
	XK_Pointer_Up              = 0xfee2
	XK_Pointer_Down            = 0xfee3
	XK_Pointer_UpLeft          = 0xfee4
	XK_Pointer_UpRight         = 0xfee5
	XK_Pointer_DownLeft        = 0xfee6
	XK_Pointer_DownRight       = 0xfee7
	XK_Pointer_Button_Dflt     = 0xfee8
	XK_Pointer_Button1         = 0xfee9
	XK_Pointer_Button2         = 0xfeea
	XK_Pointer_Button3         = 0xfeeb
	XK_Pointer_Button4         = 0xfeec
	XK_Pointer_Button5         = 0xfeed
	XK_Pointer_DblClick_Dflt   = 0xfeee
	XK_Pointer_DblClick1       = 0xfeef
	XK_Pointer_DblClick2       = 0xfef0
	XK_Pointer_DblClick3       = 0xfef1
	XK_Pointer_DblClick4       = 0xfef2
	XK_Pointer_DblClick5       = 0xfef3
	XK_Pointer_Drag_Dflt       = 0xfef4
	XK_Pointer_Drag1           = 0xfef5
	XK_Pointer_Drag2           = 0xfef6
	XK_Pointer_Drag3           = 0xfef7
	XK_Pointer_Drag4           = 0xfef8
	XK_Pointer_Drag5           = 0xfefd
	XK_Pointer_EnableKeys      = 0xfef9
	XK_Pointer_Accelerate      = 0xfefa
	XK_Pointer_DfltBtnNext     = 0xfefb
	XK_Pointer_DfltBtnPrev     = 0xfefc
	XK_ch                      = 0xfea0
	XK_Ch                      = 0xfea1
	XK_CH                      = 0xfea2
	XK_c_h                     = 0xfea3
	XK_C_h                     = 0xfea4
	XK_C_H                     = 0xfea5
	//  U+0020 SPACE
	XK_space = 0x0020
	//  U+0021 EXCLAMATION MARK
	XK_exclam = 0x0021
	//  U+0022 QUOTATION MARK
	XK_quotedbl = 0x0022
	//  U+0023 NUMBER SIGN
	XK_numbersign = 0x0023
	//  U+0024 DOLLAR SIGN
	XK_dollar = 0x0024
	//  U+0025 PERCENT SIGN
	XK_percent = 0x0025
	//  U+0026 AMPERSAND
	XK_ampersand = 0x0026
	//  U+0027 APOSTROPHE
	XK_apostrophe = 0x0027
	//  deprecated
	XK_quoteright = 0x0027
	//  U+0028 LEFT PARENTHESIS
	XK_parenleft = 0x0028
	//  U+0029 RIGHT PARENTHESIS
	XK_parenright = 0x0029
	//  U+002A ASTERISK
	XK_asterisk = 0x002a
	//  U+002B PLUS SIGN
	XK_plus = 0x002b
	//  U+002C COMMA
	XK_comma = 0x002c
	//  U+002D HYPHEN-MINUS
	XK_minus = 0x002d
	//  U+002E FULL STOP
	XK_period = 0x002e
	//  U+002F SOLIDUS
	XK_slash = 0x002f
	//  U+0030 DIGIT ZERO
	XK_0 = 0x0030
	//  U+0031 DIGIT ONE
	XK_1 = 0x0031
	//  U+0032 DIGIT TWO
	XK_2 = 0x0032
	//  U+0033 DIGIT THREE
	XK_3 = 0x0033
	//  U+0034 DIGIT FOUR
	XK_4 = 0x0034
	//  U+0035 DIGIT FIVE
	XK_5 = 0x0035
	//  U+0036 DIGIT SIX
	XK_6 = 0x0036
	//  U+0037 DIGIT SEVEN
	XK_7 = 0x0037
	//  U+0038 DIGIT EIGHT
	XK_8 = 0x0038
	//  U+0039 DIGIT NINE
	XK_9 = 0x0039
	//  U+003A COLON
	XK_colon = 0x003a
	//  U+003B SEMICOLON
	XK_semicolon = 0x003b
	//  U+003C LESS-THAN SIGN
	XK_less = 0x003c
	//  U+003D EQUALS SIGN
	XK_equal = 0x003d
	//  U+003E GREATER-THAN SIGN
	XK_greater = 0x003e
	//  U+003F QUESTION MARK
	XK_question = 0x003f
	//  U+0040 COMMERCIAL AT
	XK_at = 0x0040
	//  U+0041 LATIN CAPITAL LETTER A
	XK_A = 0x0041
	//  U+0042 LATIN CAPITAL LETTER B
	XK_B = 0x0042
	//  U+0043 LATIN CAPITAL LETTER C
	XK_C = 0x0043
	//  U+0044 LATIN CAPITAL LETTER D
	XK_D = 0x0044
	//  U+0045 LATIN CAPITAL LETTER E
	XK_E = 0x0045
	//  U+0046 LATIN CAPITAL LETTER F
	XK_F = 0x0046
	//  U+0047 LATIN CAPITAL LETTER G
	XK_G = 0x0047
	//  U+0048 LATIN CAPITAL LETTER H
	XK_H = 0x0048
	//  U+0049 LATIN CAPITAL LETTER I
	XK_I = 0x0049
	//  U+004A LATIN CAPITAL LETTER J
	XK_J = 0x004a
	//  U+004B LATIN CAPITAL LETTER K
	XK_K = 0x004b
	//  U+004C LATIN CAPITAL LETTER L
	XK_L = 0x004c
	//  U+004D LATIN CAPITAL LETTER M
	XK_M = 0x004d
	//  U+004E LATIN CAPITAL LETTER N
	XK_N = 0x004e
	//  U+004F LATIN CAPITAL LETTER O
	XK_O = 0x004f
	//  U+0050 LATIN CAPITAL LETTER P
	XK_P = 0x0050
	//  U+0051 LATIN CAPITAL LETTER Q
	XK_Q = 0x0051
	//  U+0052 LATIN CAPITAL LETTER R
	XK_R = 0x0052
	//  U+0053 LATIN CAPITAL LETTER S
	XK_S = 0x0053
	//  U+0054 LATIN CAPITAL LETTER T
	XK_T = 0x0054
	//  U+0055 LATIN CAPITAL LETTER U
	XK_U = 0x0055
	//  U+0056 LATIN CAPITAL LETTER V
	XK_V = 0x0056
	//  U+0057 LATIN CAPITAL LETTER W
	XK_W = 0x0057
	//  U+0058 LATIN CAPITAL LETTER X
	XK_X = 0x0058
	//  U+0059 LATIN CAPITAL LETTER Y
	XK_Y = 0x0059
	//  U+005A LATIN CAPITAL LETTER Z
	XK_Z = 0x005a
	//  U+005B LEFT SQUARE BRACKET
	XK_bracketleft = 0x005b
	//  U+005C REVERSE SOLIDUS
	XK_backslash = 0x005c
	//  U+005D RIGHT SQUARE BRACKET
	XK_bracketright = 0x005d
	//  U+005E CIRCUMFLEX ACCENT
	XK_asciicircum = 0x005e
	//  U+005F LOW LINE
	XK_underscore = 0x005f
	//  U+0060 GRAVE ACCENT
	XK_grave = 0x0060
	//  deprecated
	XK_quoteleft = 0x0060
	//  U+0061 LATIN SMALL LETTER A
	XK_a = 0x0061
	//  U+0062 LATIN SMALL LETTER B
	XK_b = 0x0062
	//  U+0063 LATIN SMALL LETTER C
	XK_c = 0x0063
	//  U+0064 LATIN SMALL LETTER D
	XK_d = 0x0064
	//  U+0065 LATIN SMALL LETTER E
	XK_e = 0x0065
	//  U+0066 LATIN SMALL LETTER F
	XK_f = 0x0066
	//  U+0067 LATIN SMALL LETTER G
	XK_g = 0x0067
	//  U+0068 LATIN SMALL LETTER H
	XK_h = 0x0068
	//  U+0069 LATIN SMALL LETTER I
	XK_i = 0x0069
	//  U+006A LATIN SMALL LETTER J
	XK_j = 0x006a
	//  U+006B LATIN SMALL LETTER K
	XK_k = 0x006b
	//  U+006C LATIN SMALL LETTER L
	XK_l = 0x006c
	//  U+006D LATIN SMALL LETTER M
	XK_m = 0x006d
	//  U+006E LATIN SMALL LETTER N
	XK_n = 0x006e
	//  U+006F LATIN SMALL LETTER O
	XK_o = 0x006f
	//  U+0070 LATIN SMALL LETTER P
	XK_p = 0x0070
	//  U+0071 LATIN SMALL LETTER Q
	XK_q = 0x0071
	//  U+0072 LATIN SMALL LETTER R
	XK_r = 0x0072
	//  U+0073 LATIN SMALL LETTER S
	XK_s = 0x0073
	//  U+0074 LATIN SMALL LETTER T
	XK_t = 0x0074
	//  U+0075 LATIN SMALL LETTER U
	XK_u = 0x0075
	//  U+0076 LATIN SMALL LETTER V
	XK_v = 0x0076
	//  U+0077 LATIN SMALL LETTER W
	XK_w = 0x0077
	//  U+0078 LATIN SMALL LETTER X
	XK_x = 0x0078
	//  U+0079 LATIN SMALL LETTER Y
	XK_y = 0x0079
	//  U+007A LATIN SMALL LETTER Z
	XK_z = 0x007a
	//  U+007B LEFT CURLY BRACKET
	XK_braceleft = 0x007b
	//  U+007C VERTICAL LINE
	XK_bar = 0x007c
	//  U+007D RIGHT CURLY BRACKET
	XK_braceright = 0x007d
	//  U+007E TILDE
	XK_asciitilde = 0x007e
	//  U+00A0 NO-BREAK SPACE
	XK_nobreakspace = 0x00a0
	//  U+00A1 INVERTED EXCLAMATION MARK
	XK_exclamdown = 0x00a1
	//  U+00A2 CENT SIGN
	XK_cent = 0x00a2
	//  U+00A3 POUND SIGN
	XK_sterling = 0x00a3
	//  U+00A4 CURRENCY SIGN
	XK_currency = 0x00a4
	//  U+00A5 YEN SIGN
	XK_yen = 0x00a5
	//  U+00A6 BROKEN BAR
	XK_brokenbar = 0x00a6
	//  U+00A7 SECTION SIGN
	XK_section = 0x00a7
	//  U+00A8 DIAERESIS
	XK_diaeresis = 0x00a8
	//  U+00A9 COPYRIGHT SIGN
	XK_copyright = 0x00a9
	//  U+00AA FEMININE ORDINAL INDICATOR
	XK_ordfeminine = 0x00aa
	//  U+00AB LEFT-POINTING DOUBLE ANGLE QUOTATION MARK
	XK_guillemotleft = 0x00ab
	//  U+00AC NOT SIGN
	XK_notsign = 0x00ac
	//  U+00AD SOFT HYPHEN
	XK_hyphen = 0x00ad
	//  U+00AE REGISTERED SIGN
	XK_registered = 0x00ae
	//  U+00AF MACRON
	XK_macron = 0x00af
	//  U+00B0 DEGREE SIGN
	XK_degree = 0x00b0
	//  U+00B1 PLUS-MINUS SIGN
	XK_plusminus = 0x00b1
	//  U+00B2 SUPERSCRIPT TWO
	XK_twosuperior = 0x00b2
	//  U+00B3 SUPERSCRIPT THREE
	XK_threesuperior = 0x00b3
	//  U+00B4 ACUTE ACCENT
	XK_acute = 0x00b4
	//  U+00B5 MICRO SIGN
	XK_mu = 0x00b5
	//  U+00B6 PILCROW SIGN
	XK_paragraph = 0x00b6
	//  U+00B7 MIDDLE DOT
	XK_periodcentered = 0x00b7
	//  U+00B8 CEDILLA
	XK_cedilla = 0x00b8
	//  U+00B9 SUPERSCRIPT ONE
	XK_onesuperior = 0x00b9
	//  U+00BA MASCULINE ORDINAL INDICATOR
	XK_masculine = 0x00ba
	//  U+00BB RIGHT-POINTING DOUBLE ANGLE QUOTATION MARK
	XK_guillemotright = 0x00bb
	//  U+00BC VULGAR FRACTION ONE QUARTER
	XK_onequarter = 0x00bc
	//  U+00BD VULGAR FRACTION ONE HALF
	XK_onehalf = 0x00bd
	//  U+00BE VULGAR FRACTION THREE QUARTERS
	XK_threequarters = 0x00be
	//  U+00BF INVERTED QUESTION MARK
	XK_questiondown = 0x00bf
	//  U+00C0 LATIN CAPITAL LETTER A WITH GRAVE
	XK_Agrave = 0x00c0
	//  U+00C1 LATIN CAPITAL LETTER A WITH ACUTE
	XK_Aacute = 0x00c1
	//  U+00C2 LATIN CAPITAL LETTER A WITH CIRCUMFLEX
	XK_Acircumflex = 0x00c2
	//  U+00C3 LATIN CAPITAL LETTER A WITH TILDE
	XK_Atilde = 0x00c3
	//  U+00C4 LATIN CAPITAL LETTER A WITH DIAERESIS
	XK_Adiaeresis = 0x00c4
	//  U+00C5 LATIN CAPITAL LETTER A WITH RING ABOVE
	XK_Aring = 0x00c5
	//  U+00C6 LATIN CAPITAL LETTER AE
	XK_AE = 0x00c6
	//  U+00C7 LATIN CAPITAL LETTER C WITH CEDILLA
	XK_Ccedilla = 0x00c7
	//  U+00C8 LATIN CAPITAL LETTER E WITH GRAVE
	XK_Egrave = 0x00c8
	//  U+00C9 LATIN CAPITAL LETTER E WITH ACUTE
	XK_Eacute = 0x00c9
	//  U+00CA LATIN CAPITAL LETTER E WITH CIRCUMFLEX
	XK_Ecircumflex = 0x00ca
	//  U+00CB LATIN CAPITAL LETTER E WITH DIAERESIS
	XK_Ediaeresis = 0x00cb
	//  U+00CC LATIN CAPITAL LETTER I WITH GRAVE
	XK_Igrave = 0x00cc
	//  U+00CD LATIN CAPITAL LETTER I WITH ACUTE
	XK_Iacute = 0x00cd
	//  U+00CE LATIN CAPITAL LETTER I WITH CIRCUMFLEX
	XK_Icircumflex = 0x00ce
	//  U+00CF LATIN CAPITAL LETTER I WITH DIAERESIS
	XK_Idiaeresis = 0x00cf
	//  U+00D0 LATIN CAPITAL LETTER ETH
	XK_ETH = 0x00d0
	//  deprecated
	XK_Eth = 0x00d0
	//  U+00D1 LATIN CAPITAL LETTER N WITH TILDE
	XK_Ntilde = 0x00d1
	//  U+00D2 LATIN CAPITAL LETTER O WITH GRAVE
	XK_Ograve = 0x00d2
	//  U+00D3 LATIN CAPITAL LETTER O WITH ACUTE
	XK_Oacute = 0x00d3
	//  U+00D4 LATIN CAPITAL LETTER O WITH CIRCUMFLEX
	XK_Ocircumflex = 0x00d4
	//  U+00D5 LATIN CAPITAL LETTER O WITH TILDE
	XK_Otilde = 0x00d5
	//  U+00D6 LATIN CAPITAL LETTER O WITH DIAERESIS
	XK_Odiaeresis = 0x00d6
	//  U+00D7 MULTIPLICATION SIGN
	XK_multiply = 0x00d7
	//  U+00D8 LATIN CAPITAL LETTER O WITH STROKE
	XK_Oslash = 0x00d8
	//  U+00D8 LATIN CAPITAL LETTER O WITH STROKE
	XK_Ooblique = 0x00d8
	//  U+00D9 LATIN CAPITAL LETTER U WITH GRAVE
	XK_Ugrave = 0x00d9
	//  U+00DA LATIN CAPITAL LETTER U WITH ACUTE
	XK_Uacute = 0x00da
	//  U+00DB LATIN CAPITAL LETTER U WITH CIRCUMFLEX
	XK_Ucircumflex = 0x00db
	//  U+00DC LATIN CAPITAL LETTER U WITH DIAERESIS
	XK_Udiaeresis = 0x00dc
	//  U+00DD LATIN CAPITAL LETTER Y WITH ACUTE
	XK_Yacute = 0x00dd
	//  U+00DE LATIN CAPITAL LETTER THORN
	XK_THORN = 0x00de
	//  deprecated
	XK_Thorn = 0x00de
	//  U+00DF LATIN SMALL LETTER SHARP S
	XK_ssharp = 0x00df
	//  U+00E0 LATIN SMALL LETTER A WITH GRAVE
	XK_agrave = 0x00e0
	//  U+00E1 LATIN SMALL LETTER A WITH ACUTE
	XK_aacute = 0x00e1
	//  U+00E2 LATIN SMALL LETTER A WITH CIRCUMFLEX
	XK_acircumflex = 0x00e2
	//  U+00E3 LATIN SMALL LETTER A WITH TILDE
	XK_atilde = 0x00e3
	//  U+00E4 LATIN SMALL LETTER A WITH DIAERESIS
	XK_adiaeresis = 0x00e4
	//  U+00E5 LATIN SMALL LETTER A WITH RING ABOVE
	XK_aring = 0x00e5
	//  U+00E6 LATIN SMALL LETTER AE
	XK_ae = 0x00e6
	//  U+00E7 LATIN SMALL LETTER C WITH CEDILLA
	XK_ccedilla = 0x00e7
	//  U+00E8 LATIN SMALL LETTER E WITH GRAVE
	XK_egrave = 0x00e8
	//  U+00E9 LATIN SMALL LETTER E WITH ACUTE
	XK_eacute = 0x00e9
	//  U+00EA LATIN SMALL LETTER E WITH CIRCUMFLEX
	XK_ecircumflex = 0x00ea
	//  U+00EB LATIN SMALL LETTER E WITH DIAERESIS
	XK_ediaeresis = 0x00eb
	//  U+00EC LATIN SMALL LETTER I WITH GRAVE
	XK_igrave = 0x00ec
	//  U+00ED LATIN SMALL LETTER I WITH ACUTE
	XK_iacute = 0x00ed
	//  U+00EE LATIN SMALL LETTER I WITH CIRCUMFLEX
	XK_icircumflex = 0x00ee
	//  U+00EF LATIN SMALL LETTER I WITH DIAERESIS
	XK_idiaeresis = 0x00ef
	//  U+00F0 LATIN SMALL LETTER ETH
	XK_eth = 0x00f0
	//  U+00F1 LATIN SMALL LETTER N WITH TILDE
	XK_ntilde = 0x00f1
	//  U+00F2 LATIN SMALL LETTER O WITH GRAVE
	XK_ograve = 0x00f2
	//  U+00F3 LATIN SMALL LETTER O WITH ACUTE
	XK_oacute = 0x00f3
	//  U+00F4 LATIN SMALL LETTER O WITH CIRCUMFLEX
	XK_ocircumflex = 0x00f4
	//  U+00F5 LATIN SMALL LETTER O WITH TILDE
	XK_otilde = 0x00f5
	//  U+00F6 LATIN SMALL LETTER O WITH DIAERESIS
	XK_odiaeresis = 0x00f6
	//  U+00F7 DIVISION SIGN
	XK_division = 0x00f7
	//  U+00F8 LATIN SMALL LETTER O WITH STROKE
	XK_oslash = 0x00f8
	//  U+00F8 LATIN SMALL LETTER O WITH STROKE
	XK_ooblique = 0x00f8
	//  U+00F9 LATIN SMALL LETTER U WITH GRAVE
	XK_ugrave = 0x00f9
	//  U+00FA LATIN SMALL LETTER U WITH ACUTE
	XK_uacute = 0x00fa
	//  U+00FB LATIN SMALL LETTER U WITH CIRCUMFLEX
	XK_ucircumflex = 0x00fb
	//  U+00FC LATIN SMALL LETTER U WITH DIAERESIS
	XK_udiaeresis = 0x00fc
	//  U+00FD LATIN SMALL LETTER Y WITH ACUTE
	XK_yacute = 0x00fd
	//  U+00FE LATIN SMALL LETTER THORN
	XK_thorn = 0x00fe
	//  U+00FF LATIN SMALL LETTER Y WITH DIAERESIS
	XK_ydiaeresis = 0x00ff
	//  U+0104 LATIN CAPITAL LETTER A WITH OGONEK
	XK_Aogonek = 0x01a1
	//  U+02D8 BREVE
	XK_breve = 0x01a2
	//  U+0141 LATIN CAPITAL LETTER L WITH STROKE
	XK_Lstroke = 0x01a3
	//  U+013D LATIN CAPITAL LETTER L WITH CARON
	XK_Lcaron = 0x01a5
	//  U+015A LATIN CAPITAL LETTER S WITH ACUTE
	XK_Sacute = 0x01a6
	//  U+0160 LATIN CAPITAL LETTER S WITH CARON
	XK_Scaron = 0x01a9
	//  U+015E LATIN CAPITAL LETTER S WITH CEDILLA
	XK_Scedilla = 0x01aa
	//  U+0164 LATIN CAPITAL LETTER T WITH CARON
	XK_Tcaron = 0x01ab
	//  U+0179 LATIN CAPITAL LETTER Z WITH ACUTE
	XK_Zacute = 0x01ac
	//  U+017D LATIN CAPITAL LETTER Z WITH CARON
	XK_Zcaron = 0x01ae
	//  U+017B LATIN CAPITAL LETTER Z WITH DOT ABOVE
	XK_Zabovedot = 0x01af
	//  U+0105 LATIN SMALL LETTER A WITH OGONEK
	XK_aogonek = 0x01b1
	//  U+02DB OGONEK
	XK_ogonek = 0x01b2
	//  U+0142 LATIN SMALL LETTER L WITH STROKE
	XK_lstroke = 0x01b3
	//  U+013E LATIN SMALL LETTER L WITH CARON
	XK_lcaron = 0x01b5
	//  U+015B LATIN SMALL LETTER S WITH ACUTE
	XK_sacute = 0x01b6
	//  U+02C7 CARON
	XK_caron = 0x01b7
	//  U+0161 LATIN SMALL LETTER S WITH CARON
	XK_scaron = 0x01b9
	//  U+015F LATIN SMALL LETTER S WITH CEDILLA
	XK_scedilla = 0x01ba
	//  U+0165 LATIN SMALL LETTER T WITH CARON
	XK_tcaron = 0x01bb
	//  U+017A LATIN SMALL LETTER Z WITH ACUTE
	XK_zacute = 0x01bc
	//  U+02DD DOUBLE ACUTE ACCENT
	XK_doubleacute = 0x01bd
	//  U+017E LATIN SMALL LETTER Z WITH CARON
	XK_zcaron = 0x01be
	//  U+017C LATIN SMALL LETTER Z WITH DOT ABOVE
	XK_zabovedot = 0x01bf
	//  U+0154 LATIN CAPITAL LETTER R WITH ACUTE
	XK_Racute = 0x01c0
	//  U+0102 LATIN CAPITAL LETTER A WITH BREVE
	XK_Abreve = 0x01c3
	//  U+0139 LATIN CAPITAL LETTER L WITH ACUTE
	XK_Lacute = 0x01c5
	//  U+0106 LATIN CAPITAL LETTER C WITH ACUTE
	XK_Cacute = 0x01c6
	//  U+010C LATIN CAPITAL LETTER C WITH CARON
	XK_Ccaron = 0x01c8
	//  U+0118 LATIN CAPITAL LETTER E WITH OGONEK
	XK_Eogonek = 0x01ca
	//  U+011A LATIN CAPITAL LETTER E WITH CARON
	XK_Ecaron = 0x01cc
	//  U+010E LATIN CAPITAL LETTER D WITH CARON
	XK_Dcaron = 0x01cf
	//  U+0110 LATIN CAPITAL LETTER D WITH STROKE
	XK_Dstroke = 0x01d0
	//  U+0143 LATIN CAPITAL LETTER N WITH ACUTE
	XK_Nacute = 0x01d1
	//  U+0147 LATIN CAPITAL LETTER N WITH CARON
	XK_Ncaron = 0x01d2
	//  U+0150 LATIN CAPITAL LETTER O WITH DOUBLE ACUTE
	XK_Odoubleacute = 0x01d5
	//  U+0158 LATIN CAPITAL LETTER R WITH CARON
	XK_Rcaron = 0x01d8
	//  U+016E LATIN CAPITAL LETTER U WITH RING ABOVE
	XK_Uring = 0x01d9
	//  U+0170 LATIN CAPITAL LETTER U WITH DOUBLE ACUTE
	XK_Udoubleacute = 0x01db
	//  U+0162 LATIN CAPITAL LETTER T WITH CEDILLA
	XK_Tcedilla = 0x01de
	//  U+0155 LATIN SMALL LETTER R WITH ACUTE
	XK_racute = 0x01e0
	//  U+0103 LATIN SMALL LETTER A WITH BREVE
	XK_abreve = 0x01e3
	//  U+013A LATIN SMALL LETTER L WITH ACUTE
	XK_lacute = 0x01e5
	//  U+0107 LATIN SMALL LETTER C WITH ACUTE
	XK_cacute = 0x01e6
	//  U+010D LATIN SMALL LETTER C WITH CARON
	XK_ccaron = 0x01e8
	//  U+0119 LATIN SMALL LETTER E WITH OGONEK
	XK_eogonek = 0x01ea
	//  U+011B LATIN SMALL LETTER E WITH CARON
	XK_ecaron = 0x01ec
	//  U+010F LATIN SMALL LETTER D WITH CARON
	XK_dcaron = 0x01ef
	//  U+0111 LATIN SMALL LETTER D WITH STROKE
	XK_dstroke = 0x01f0
	//  U+0144 LATIN SMALL LETTER N WITH ACUTE
	XK_nacute = 0x01f1
	//  U+0148 LATIN SMALL LETTER N WITH CARON
	XK_ncaron = 0x01f2
	//  U+0151 LATIN SMALL LETTER O WITH DOUBLE ACUTE
	XK_odoubleacute = 0x01f5
	//  U+0159 LATIN SMALL LETTER R WITH CARON
	XK_rcaron = 0x01f8
	//  U+016F LATIN SMALL LETTER U WITH RING ABOVE
	XK_uring = 0x01f9
	//  U+0171 LATIN SMALL LETTER U WITH DOUBLE ACUTE
	XK_udoubleacute = 0x01fb
	//  U+0163 LATIN SMALL LETTER T WITH CEDILLA
	XK_tcedilla = 0x01fe
	//  U+02D9 DOT ABOVE
	XK_abovedot = 0x01ff
	//  U+0126 LATIN CAPITAL LETTER H WITH STROKE
	XK_Hstroke = 0x02a1
	//  U+0124 LATIN CAPITAL LETTER H WITH CIRCUMFLEX
	XK_Hcircumflex = 0x02a6
	//  U+0130 LATIN CAPITAL LETTER I WITH DOT ABOVE
	XK_Iabovedot = 0x02a9
	//  U+011E LATIN CAPITAL LETTER G WITH BREVE
	XK_Gbreve = 0x02ab
	//  U+0134 LATIN CAPITAL LETTER J WITH CIRCUMFLEX
	XK_Jcircumflex = 0x02ac
	//  U+0127 LATIN SMALL LETTER H WITH STROKE
	XK_hstroke = 0x02b1
	//  U+0125 LATIN SMALL LETTER H WITH CIRCUMFLEX
	XK_hcircumflex = 0x02b6
	//  U+0131 LATIN SMALL LETTER DOTLESS I
	XK_idotless = 0x02b9
	//  U+011F LATIN SMALL LETTER G WITH BREVE
	XK_gbreve = 0x02bb
	//  U+0135 LATIN SMALL LETTER J WITH CIRCUMFLEX
	XK_jcircumflex = 0x02bc
	//  U+010A LATIN CAPITAL LETTER C WITH DOT ABOVE
	XK_Cabovedot = 0x02c5
	//  U+0108 LATIN CAPITAL LETTER C WITH CIRCUMFLEX
	XK_Ccircumflex = 0x02c6
	//  U+0120 LATIN CAPITAL LETTER G WITH DOT ABOVE
	XK_Gabovedot = 0x02d5
	//  U+011C LATIN CAPITAL LETTER G WITH CIRCUMFLEX
	XK_Gcircumflex = 0x02d8
	//  U+016C LATIN CAPITAL LETTER U WITH BREVE
	XK_Ubreve = 0x02dd
	//  U+015C LATIN CAPITAL LETTER S WITH CIRCUMFLEX
	XK_Scircumflex = 0x02de
	//  U+010B LATIN SMALL LETTER C WITH DOT ABOVE
	XK_cabovedot = 0x02e5
	//  U+0109 LATIN SMALL LETTER C WITH CIRCUMFLEX
	XK_ccircumflex = 0x02e6
	//  U+0121 LATIN SMALL LETTER G WITH DOT ABOVE
	XK_gabovedot = 0x02f5
	//  U+011D LATIN SMALL LETTER G WITH CIRCUMFLEX
	XK_gcircumflex = 0x02f8
	//  U+016D LATIN SMALL LETTER U WITH BREVE
	XK_ubreve = 0x02fd
	//  U+015D LATIN SMALL LETTER S WITH CIRCUMFLEX
	XK_scircumflex = 0x02fe
	//  U+0138 LATIN SMALL LETTER KRA
	XK_kra = 0x03a2
	//  deprecated
	XK_kappa = 0x03a2
	//  U+0156 LATIN CAPITAL LETTER R WITH CEDILLA
	XK_Rcedilla = 0x03a3
	//  U+0128 LATIN CAPITAL LETTER I WITH TILDE
	XK_Itilde = 0x03a5
	//  U+013B LATIN CAPITAL LETTER L WITH CEDILLA
	XK_Lcedilla = 0x03a6
	//  U+0112 LATIN CAPITAL LETTER E WITH MACRON
	XK_Emacron = 0x03aa
	//  U+0122 LATIN CAPITAL LETTER G WITH CEDILLA
	XK_Gcedilla = 0x03ab
	//  U+0166 LATIN CAPITAL LETTER T WITH STROKE
	XK_Tslash = 0x03ac
	//  U+0157 LATIN SMALL LETTER R WITH CEDILLA
	XK_rcedilla = 0x03b3
	//  U+0129 LATIN SMALL LETTER I WITH TILDE
	XK_itilde = 0x03b5
	//  U+013C LATIN SMALL LETTER L WITH CEDILLA
	XK_lcedilla = 0x03b6
	//  U+0113 LATIN SMALL LETTER E WITH MACRON
	XK_emacron = 0x03ba
	//  U+0123 LATIN SMALL LETTER G WITH CEDILLA
	XK_gcedilla = 0x03bb
	//  U+0167 LATIN SMALL LETTER T WITH STROKE
	XK_tslash = 0x03bc
	//  U+014A LATIN CAPITAL LETTER ENG
	XK_ENG = 0x03bd
	//  U+014B LATIN SMALL LETTER ENG
	XK_eng = 0x03bf
	//  U+0100 LATIN CAPITAL LETTER A WITH MACRON
	XK_Amacron = 0x03c0
	//  U+012E LATIN CAPITAL LETTER I WITH OGONEK
	XK_Iogonek = 0x03c7
	//  U+0116 LATIN CAPITAL LETTER E WITH DOT ABOVE
	XK_Eabovedot = 0x03cc
	//  U+012A LATIN CAPITAL LETTER I WITH MACRON
	XK_Imacron = 0x03cf
	//  U+0145 LATIN CAPITAL LETTER N WITH CEDILLA
	XK_Ncedilla = 0x03d1
	//  U+014C LATIN CAPITAL LETTER O WITH MACRON
	XK_Omacron = 0x03d2
	//  U+0136 LATIN CAPITAL LETTER K WITH CEDILLA
	XK_Kcedilla = 0x03d3
	//  U+0172 LATIN CAPITAL LETTER U WITH OGONEK
	XK_Uogonek = 0x03d9
	//  U+0168 LATIN CAPITAL LETTER U WITH TILDE
	XK_Utilde = 0x03dd
	//  U+016A LATIN CAPITAL LETTER U WITH MACRON
	XK_Umacron = 0x03de
	//  U+0101 LATIN SMALL LETTER A WITH MACRON
	XK_amacron = 0x03e0
	//  U+012F LATIN SMALL LETTER I WITH OGONEK
	XK_iogonek = 0x03e7
	//  U+0117 LATIN SMALL LETTER E WITH DOT ABOVE
	XK_eabovedot = 0x03ec
	//  U+012B LATIN SMALL LETTER I WITH MACRON
	XK_imacron = 0x03ef
	//  U+0146 LATIN SMALL LETTER N WITH CEDILLA
	XK_ncedilla = 0x03f1
	//  U+014D LATIN SMALL LETTER O WITH MACRON
	XK_omacron = 0x03f2
	//  U+0137 LATIN SMALL LETTER K WITH CEDILLA
	XK_kcedilla = 0x03f3
	//  U+0173 LATIN SMALL LETTER U WITH OGONEK
	XK_uogonek = 0x03f9
	//  U+0169 LATIN SMALL LETTER U WITH TILDE
	XK_utilde = 0x03fd
	//  U+016B LATIN SMALL LETTER U WITH MACRON
	XK_umacron = 0x03fe
	//  U+0174 LATIN CAPITAL LETTER W WITH CIRCUMFLEX
	XK_Wcircumflex = 0x1000174
	//  U+0175 LATIN SMALL LETTER W WITH CIRCUMFLEX
	XK_wcircumflex = 0x1000175
	//  U+0176 LATIN CAPITAL LETTER Y WITH CIRCUMFLEX
	XK_Ycircumflex = 0x1000176
	//  U+0177 LATIN SMALL LETTER Y WITH CIRCUMFLEX
	XK_ycircumflex = 0x1000177
	//  U+1E02 LATIN CAPITAL LETTER B WITH DOT ABOVE
	XK_Babovedot = 0x1001e02
	//  U+1E03 LATIN SMALL LETTER B WITH DOT ABOVE
	XK_babovedot = 0x1001e03
	//  U+1E0A LATIN CAPITAL LETTER D WITH DOT ABOVE
	XK_Dabovedot = 0x1001e0a
	//  U+1E0B LATIN SMALL LETTER D WITH DOT ABOVE
	XK_dabovedot = 0x1001e0b
	//  U+1E1E LATIN CAPITAL LETTER F WITH DOT ABOVE
	XK_Fabovedot = 0x1001e1e
	//  U+1E1F LATIN SMALL LETTER F WITH DOT ABOVE
	XK_fabovedot = 0x1001e1f
	//  U+1E40 LATIN CAPITAL LETTER M WITH DOT ABOVE
	XK_Mabovedot = 0x1001e40
	//  U+1E41 LATIN SMALL LETTER M WITH DOT ABOVE
	XK_mabovedot = 0x1001e41
	//  U+1E56 LATIN CAPITAL LETTER P WITH DOT ABOVE
	XK_Pabovedot = 0x1001e56
	//  U+1E57 LATIN SMALL LETTER P WITH DOT ABOVE
	XK_pabovedot = 0x1001e57
	//  U+1E60 LATIN CAPITAL LETTER S WITH DOT ABOVE
	XK_Sabovedot = 0x1001e60
	//  U+1E61 LATIN SMALL LETTER S WITH DOT ABOVE
	XK_sabovedot = 0x1001e61
	//  U+1E6A LATIN CAPITAL LETTER T WITH DOT ABOVE
	XK_Tabovedot = 0x1001e6a
	//  U+1E6B LATIN SMALL LETTER T WITH DOT ABOVE
	XK_tabovedot = 0x1001e6b
	//  U+1E80 LATIN CAPITAL LETTER W WITH GRAVE
	XK_Wgrave = 0x1001e80
	//  U+1E81 LATIN SMALL LETTER W WITH GRAVE
	XK_wgrave = 0x1001e81
	//  U+1E82 LATIN CAPITAL LETTER W WITH ACUTE
	XK_Wacute = 0x1001e82
	//  U+1E83 LATIN SMALL LETTER W WITH ACUTE
	XK_wacute = 0x1001e83
	//  U+1E84 LATIN CAPITAL LETTER W WITH DIAERESIS
	XK_Wdiaeresis = 0x1001e84
	//  U+1E85 LATIN SMALL LETTER W WITH DIAERESIS
	XK_wdiaeresis = 0x1001e85
	//  U+1EF2 LATIN CAPITAL LETTER Y WITH GRAVE
	XK_Ygrave = 0x1001ef2
	//  U+1EF3 LATIN SMALL LETTER Y WITH GRAVE
	XK_ygrave = 0x1001ef3
	//  U+0152 LATIN CAPITAL LIGATURE OE
	XK_OE = 0x13bc
	//  U+0153 LATIN SMALL LIGATURE OE
	XK_oe = 0x13bd
	//  U+0178 LATIN CAPITAL LETTER Y WITH DIAERESIS
	XK_Ydiaeresis = 0x13be
	//  U+203E OVERLINE
	XK_overline = 0x047e
	//  U+3002 IDEOGRAPHIC FULL STOP
	XK_kana_fullstop = 0x04a1
	//  U+300C LEFT CORNER BRACKET
	XK_kana_openingbracket = 0x04a2
	//  U+300D RIGHT CORNER BRACKET
	XK_kana_closingbracket = 0x04a3
	//  U+3001 IDEOGRAPHIC COMMA
	XK_kana_comma = 0x04a4
	//  U+30FB KATAKANA MIDDLE DOT
	XK_kana_conjunctive = 0x04a5
	//  deprecated
	XK_kana_middledot = 0x04a5
	//  U+30F2 KATAKANA LETTER WO
	XK_kana_WO = 0x04a6
	//  U+30A1 KATAKANA LETTER SMALL A
	XK_kana_a = 0x04a7
	//  U+30A3 KATAKANA LETTER SMALL I
	XK_kana_i = 0x04a8
	//  U+30A5 KATAKANA LETTER SMALL U
	XK_kana_u = 0x04a9
	//  U+30A7 KATAKANA LETTER SMALL E
	XK_kana_e = 0x04aa
	//  U+30A9 KATAKANA LETTER SMALL O
	XK_kana_o = 0x04ab
	//  U+30E3 KATAKANA LETTER SMALL YA
	XK_kana_ya = 0x04ac
	//  U+30E5 KATAKANA LETTER SMALL YU
	XK_kana_yu = 0x04ad
	//  U+30E7 KATAKANA LETTER SMALL YO
	XK_kana_yo = 0x04ae
	//  U+30C3 KATAKANA LETTER SMALL TU
	XK_kana_tsu = 0x04af
	//  deprecated
	XK_kana_tu = 0x04af
	//  U+30FC KATAKANA-HIRAGANA PROLONGED SOUND MARK
	XK_prolongedsound = 0x04b0
	//  U+30A2 KATAKANA LETTER A
	XK_kana_A = 0x04b1
	//  U+30A4 KATAKANA LETTER I
	XK_kana_I = 0x04b2
	//  U+30A6 KATAKANA LETTER U
	XK_kana_U = 0x04b3
	//  U+30A8 KATAKANA LETTER E
	XK_kana_E = 0x04b4
	//  U+30AA KATAKANA LETTER O
	XK_kana_O = 0x04b5
	//  U+30AB KATAKANA LETTER KA
	XK_kana_KA = 0x04b6
	//  U+30AD KATAKANA LETTER KI
	XK_kana_KI = 0x04b7
	//  U+30AF KATAKANA LETTER KU
	XK_kana_KU = 0x04b8
	//  U+30B1 KATAKANA LETTER KE
	XK_kana_KE = 0x04b9
	//  U+30B3 KATAKANA LETTER KO
	XK_kana_KO = 0x04ba
	//  U+30B5 KATAKANA LETTER SA
	XK_kana_SA = 0x04bb
	//  U+30B7 KATAKANA LETTER SI
	XK_kana_SHI = 0x04bc
	//  U+30B9 KATAKANA LETTER SU
	XK_kana_SU = 0x04bd
	//  U+30BB KATAKANA LETTER SE
	XK_kana_SE = 0x04be
	//  U+30BD KATAKANA LETTER SO
	XK_kana_SO = 0x04bf
	//  U+30BF KATAKANA LETTER TA
	XK_kana_TA = 0x04c0
	//  U+30C1 KATAKANA LETTER TI
	XK_kana_CHI = 0x04c1
	//  deprecated
	XK_kana_TI = 0x04c1
	//  U+30C4 KATAKANA LETTER TU
	XK_kana_TSU = 0x04c2
	//  deprecated
	XK_kana_TU = 0x04c2
	//  U+30C6 KATAKANA LETTER TE
	XK_kana_TE = 0x04c3
	//  U+30C8 KATAKANA LETTER TO
	XK_kana_TO = 0x04c4
	//  U+30CA KATAKANA LETTER NA
	XK_kana_NA = 0x04c5
	//  U+30CB KATAKANA LETTER NI
	XK_kana_NI = 0x04c6
	//  U+30CC KATAKANA LETTER NU
	XK_kana_NU = 0x04c7
	//  U+30CD KATAKANA LETTER NE
	XK_kana_NE = 0x04c8
	//  U+30CE KATAKANA LETTER NO
	XK_kana_NO = 0x04c9
	//  U+30CF KATAKANA LETTER HA
	XK_kana_HA = 0x04ca
	//  U+30D2 KATAKANA LETTER HI
	XK_kana_HI = 0x04cb
	//  U+30D5 KATAKANA LETTER HU
	XK_kana_FU = 0x04cc
	//  deprecated
	XK_kana_HU = 0x04cc
	//  U+30D8 KATAKANA LETTER HE
	XK_kana_HE = 0x04cd
	//  U+30DB KATAKANA LETTER HO
	XK_kana_HO = 0x04ce
	//  U+30DE KATAKANA LETTER MA
	XK_kana_MA = 0x04cf
	//  U+30DF KATAKANA LETTER MI
	XK_kana_MI = 0x04d0
	//  U+30E0 KATAKANA LETTER MU
	XK_kana_MU = 0x04d1
	//  U+30E1 KATAKANA LETTER ME
	XK_kana_ME = 0x04d2
	//  U+30E2 KATAKANA LETTER MO
	XK_kana_MO = 0x04d3
	//  U+30E4 KATAKANA LETTER YA
	XK_kana_YA = 0x04d4
	//  U+30E6 KATAKANA LETTER YU
	XK_kana_YU = 0x04d5
	//  U+30E8 KATAKANA LETTER YO
	XK_kana_YO = 0x04d6
	//  U+30E9 KATAKANA LETTER RA
	XK_kana_RA = 0x04d7
	//  U+30EA KATAKANA LETTER RI
	XK_kana_RI = 0x04d8
	//  U+30EB KATAKANA LETTER RU
	XK_kana_RU = 0x04d9
	//  U+30EC KATAKANA LETTER RE
	XK_kana_RE = 0x04da
	//  U+30ED KATAKANA LETTER RO
	XK_kana_RO = 0x04db
	//  U+30EF KATAKANA LETTER WA
	XK_kana_WA = 0x04dc
	//  U+30F3 KATAKANA LETTER N
	XK_kana_N = 0x04dd
	//  U+309B KATAKANA-HIRAGANA VOICED SOUND MARK
	XK_voicedsound = 0x04de
	//  U+309C KATAKANA-HIRAGANA SEMI-VOICED SOUND MARK
	XK_semivoicedsound = 0x04df
	//  Alias for mode_switch
	XK_kana_switch = 0xff7e
	//  U+06F0 EXTENDED ARABIC-INDIC DIGIT ZERO
	XK_Farsi_0 = 0x10006f0
	//  U+06F1 EXTENDED ARABIC-INDIC DIGIT ONE
	XK_Farsi_1 = 0x10006f1
	//  U+06F2 EXTENDED ARABIC-INDIC DIGIT TWO
	XK_Farsi_2 = 0x10006f2
	//  U+06F3 EXTENDED ARABIC-INDIC DIGIT THREE
	XK_Farsi_3 = 0x10006f3
	//  U+06F4 EXTENDED ARABIC-INDIC DIGIT FOUR
	XK_Farsi_4 = 0x10006f4
	//  U+06F5 EXTENDED ARABIC-INDIC DIGIT FIVE
	XK_Farsi_5 = 0x10006f5
	//  U+06F6 EXTENDED ARABIC-INDIC DIGIT SIX
	XK_Farsi_6 = 0x10006f6
	//  U+06F7 EXTENDED ARABIC-INDIC DIGIT SEVEN
	XK_Farsi_7 = 0x10006f7
	//  U+06F8 EXTENDED ARABIC-INDIC DIGIT EIGHT
	XK_Farsi_8 = 0x10006f8
	//  U+06F9 EXTENDED ARABIC-INDIC DIGIT NINE
	XK_Farsi_9 = 0x10006f9
	//  U+066A ARABIC PERCENT SIGN
	XK_Arabic_percent = 0x100066a
	//  U+0670 ARABIC LETTER SUPERSCRIPT ALEF
	XK_Arabic_superscript_alef = 0x1000670
	//  U+0679 ARABIC LETTER TTEH
	XK_Arabic_tteh = 0x1000679
	//  U+067E ARABIC LETTER PEH
	XK_Arabic_peh = 0x100067e
	//  U+0686 ARABIC LETTER TCHEH
	XK_Arabic_tcheh = 0x1000686
	//  U+0688 ARABIC LETTER DDAL
	XK_Arabic_ddal = 0x1000688
	//  U+0691 ARABIC LETTER RREH
	XK_Arabic_rreh = 0x1000691
	//  U+060C ARABIC COMMA
	XK_Arabic_comma = 0x05ac
	//  U+06D4 ARABIC FULL STOP
	XK_Arabic_fullstop = 0x10006d4
	//  U+0660 ARABIC-INDIC DIGIT ZERO
	XK_Arabic_0 = 0x1000660
	//  U+0661 ARABIC-INDIC DIGIT ONE
	XK_Arabic_1 = 0x1000661
	//  U+0662 ARABIC-INDIC DIGIT TWO
	XK_Arabic_2 = 0x1000662
	//  U+0663 ARABIC-INDIC DIGIT THREE
	XK_Arabic_3 = 0x1000663
	//  U+0664 ARABIC-INDIC DIGIT FOUR
	XK_Arabic_4 = 0x1000664
	//  U+0665 ARABIC-INDIC DIGIT FIVE
	XK_Arabic_5 = 0x1000665
	//  U+0666 ARABIC-INDIC DIGIT SIX
	XK_Arabic_6 = 0x1000666
	//  U+0667 ARABIC-INDIC DIGIT SEVEN
	XK_Arabic_7 = 0x1000667
	//  U+0668 ARABIC-INDIC DIGIT EIGHT
	XK_Arabic_8 = 0x1000668
	//  U+0669 ARABIC-INDIC DIGIT NINE
	XK_Arabic_9 = 0x1000669
	//  U+061B ARABIC SEMICOLON
	XK_Arabic_semicolon = 0x05bb
	//  U+061F ARABIC QUESTION MARK
	XK_Arabic_question_mark = 0x05bf
	//  U+0621 ARABIC LETTER HAMZA
	XK_Arabic_hamza = 0x05c1
	//  U+0622 ARABIC LETTER ALEF WITH MADDA ABOVE
	XK_Arabic_maddaonalef = 0x05c2
	//  U+0623 ARABIC LETTER ALEF WITH HAMZA ABOVE
	XK_Arabic_hamzaonalef = 0x05c3
	//  U+0624 ARABIC LETTER WAW WITH HAMZA ABOVE
	XK_Arabic_hamzaonwaw = 0x05c4
	//  U+0625 ARABIC LETTER ALEF WITH HAMZA BELOW
	XK_Arabic_hamzaunderalef = 0x05c5
	//  U+0626 ARABIC LETTER YEH WITH HAMZA ABOVE
	XK_Arabic_hamzaonyeh = 0x05c6
	//  U+0627 ARABIC LETTER ALEF
	XK_Arabic_alef = 0x05c7
	//  U+0628 ARABIC LETTER BEH
	XK_Arabic_beh = 0x05c8
	//  U+0629 ARABIC LETTER TEH MARBUTA
	XK_Arabic_tehmarbuta = 0x05c9
	//  U+062A ARABIC LETTER TEH
	XK_Arabic_teh = 0x05ca
	//  U+062B ARABIC LETTER THEH
	XK_Arabic_theh = 0x05cb
	//  U+062C ARABIC LETTER JEEM
	XK_Arabic_jeem = 0x05cc
	//  U+062D ARABIC LETTER HAH
	XK_Arabic_hah = 0x05cd
	//  U+062E ARABIC LETTER KHAH
	XK_Arabic_khah = 0x05ce
	//  U+062F ARABIC LETTER DAL
	XK_Arabic_dal = 0x05cf
	//  U+0630 ARABIC LETTER THAL
	XK_Arabic_thal = 0x05d0
	//  U+0631 ARABIC LETTER REH
	XK_Arabic_ra = 0x05d1
	//  U+0632 ARABIC LETTER ZAIN
	XK_Arabic_zain = 0x05d2
	//  U+0633 ARABIC LETTER SEEN
	XK_Arabic_seen = 0x05d3
	//  U+0634 ARABIC LETTER SHEEN
	XK_Arabic_sheen = 0x05d4
	//  U+0635 ARABIC LETTER SAD
	XK_Arabic_sad = 0x05d5
	//  U+0636 ARABIC LETTER DAD
	XK_Arabic_dad = 0x05d6
	//  U+0637 ARABIC LETTER TAH
	XK_Arabic_tah = 0x05d7
	//  U+0638 ARABIC LETTER ZAH
	XK_Arabic_zah = 0x05d8
	//  U+0639 ARABIC LETTER AIN
	XK_Arabic_ain = 0x05d9
	//  U+063A ARABIC LETTER GHAIN
	XK_Arabic_ghain = 0x05da
	//  U+0640 ARABIC TATWEEL
	XK_Arabic_tatweel = 0x05e0
	//  U+0641 ARABIC LETTER FEH
	XK_Arabic_feh = 0x05e1
	//  U+0642 ARABIC LETTER QAF
	XK_Arabic_qaf = 0x05e2
	//  U+0643 ARABIC LETTER KAF
	XK_Arabic_kaf = 0x05e3
	//  U+0644 ARABIC LETTER LAM
	XK_Arabic_lam = 0x05e4
	//  U+0645 ARABIC LETTER MEEM
	XK_Arabic_meem = 0x05e5
	//  U+0646 ARABIC LETTER NOON
	XK_Arabic_noon = 0x05e6
	//  U+0647 ARABIC LETTER HEH
	XK_Arabic_ha = 0x05e7
	//  deprecated
	XK_Arabic_heh = 0x05e7
	//  U+0648 ARABIC LETTER WAW
	XK_Arabic_waw = 0x05e8
	//  U+0649 ARABIC LETTER ALEF MAKSURA
	XK_Arabic_alefmaksura = 0x05e9
	//  U+064A ARABIC LETTER YEH
	XK_Arabic_yeh = 0x05ea
	//  U+064B ARABIC FATHATAN
	XK_Arabic_fathatan = 0x05eb
	//  U+064C ARABIC DAMMATAN
	XK_Arabic_dammatan = 0x05ec
	//  U+064D ARABIC KASRATAN
	XK_Arabic_kasratan = 0x05ed
	//  U+064E ARABIC FATHA
	XK_Arabic_fatha = 0x05ee
	//  U+064F ARABIC DAMMA
	XK_Arabic_damma = 0x05ef
	//  U+0650 ARABIC KASRA
	XK_Arabic_kasra = 0x05f0
	//  U+0651 ARABIC SHADDA
	XK_Arabic_shadda = 0x05f1
	//  U+0652 ARABIC SUKUN
	XK_Arabic_sukun = 0x05f2
	//  U+0653 ARABIC MADDAH ABOVE
	XK_Arabic_madda_above = 0x1000653
	//  U+0654 ARABIC HAMZA ABOVE
	XK_Arabic_hamza_above = 0x1000654
	//  U+0655 ARABIC HAMZA BELOW
	XK_Arabic_hamza_below = 0x1000655
	//  U+0698 ARABIC LETTER JEH
	XK_Arabic_jeh = 0x1000698
	//  U+06A4 ARABIC LETTER VEH
	XK_Arabic_veh = 0x10006a4
	//  U+06A9 ARABIC LETTER KEHEH
	XK_Arabic_keheh = 0x10006a9
	//  U+06AF ARABIC LETTER GAF
	XK_Arabic_gaf = 0x10006af
	//  U+06BA ARABIC LETTER NOON GHUNNA
	XK_Arabic_noon_ghunna = 0x10006ba
	//  U+06BE ARABIC LETTER HEH DOACHASHMEE
	XK_Arabic_heh_doachashmee = 0x10006be
	//  U+06CC ARABIC LETTER FARSI YEH
	XK_Farsi_yeh = 0x10006cc
	//  U+06CC ARABIC LETTER FARSI YEH
	XK_Arabic_farsi_yeh = 0x10006cc
	//  U+06D2 ARABIC LETTER YEH BARREE
	XK_Arabic_yeh_baree = 0x10006d2
	//  U+06C1 ARABIC LETTER HEH GOAL
	XK_Arabic_heh_goal = 0x10006c1
	//  Alias for mode_switch
	XK_Arabic_switch = 0xff7e
	//  U+0492 CYRILLIC CAPITAL LETTER GHE WITH STROKE
	XK_Cyrillic_GHE_bar = 0x1000492
	//  U+0493 CYRILLIC SMALL LETTER GHE WITH STROKE
	XK_Cyrillic_ghe_bar = 0x1000493
	//  U+0496 CYRILLIC CAPITAL LETTER ZHE WITH DESCENDER
	XK_Cyrillic_ZHE_descender = 0x1000496
	//  U+0497 CYRILLIC SMALL LETTER ZHE WITH DESCENDER
	XK_Cyrillic_zhe_descender = 0x1000497
	//  U+049A CYRILLIC CAPITAL LETTER KA WITH DESCENDER
	XK_Cyrillic_KA_descender = 0x100049a
	//  U+049B CYRILLIC SMALL LETTER KA WITH DESCENDER
	XK_Cyrillic_ka_descender = 0x100049b
	//  U+049C CYRILLIC CAPITAL LETTER KA WITH VERTICAL STROKE
	XK_Cyrillic_KA_vertstroke = 0x100049c
	//  U+049D CYRILLIC SMALL LETTER KA WITH VERTICAL STROKE
	XK_Cyrillic_ka_vertstroke = 0x100049d
	//  U+04A2 CYRILLIC CAPITAL LETTER EN WITH DESCENDER
	XK_Cyrillic_EN_descender = 0x10004a2
	//  U+04A3 CYRILLIC SMALL LETTER EN WITH DESCENDER
	XK_Cyrillic_en_descender = 0x10004a3
	//  U+04AE CYRILLIC CAPITAL LETTER STRAIGHT U
	XK_Cyrillic_U_straight = 0x10004ae
	//  U+04AF CYRILLIC SMALL LETTER STRAIGHT U
	XK_Cyrillic_u_straight = 0x10004af
	//  U+04B0 CYRILLIC CAPITAL LETTER STRAIGHT U WITH STROKE
	XK_Cyrillic_U_straight_bar = 0x10004b0
	//  U+04B1 CYRILLIC SMALL LETTER STRAIGHT U WITH STROKE
	XK_Cyrillic_u_straight_bar = 0x10004b1
	//  U+04B2 CYRILLIC CAPITAL LETTER HA WITH DESCENDER
	XK_Cyrillic_HA_descender = 0x10004b2
	//  U+04B3 CYRILLIC SMALL LETTER HA WITH DESCENDER
	XK_Cyrillic_ha_descender = 0x10004b3
	//  U+04B6 CYRILLIC CAPITAL LETTER CHE WITH DESCENDER
	XK_Cyrillic_CHE_descender = 0x10004b6
	//  U+04B7 CYRILLIC SMALL LETTER CHE WITH DESCENDER
	XK_Cyrillic_che_descender = 0x10004b7
	//  U+04B8 CYRILLIC CAPITAL LETTER CHE WITH VERTICAL STROKE
	XK_Cyrillic_CHE_vertstroke = 0x10004b8
	//  U+04B9 CYRILLIC SMALL LETTER CHE WITH VERTICAL STROKE
	XK_Cyrillic_che_vertstroke = 0x10004b9
	//  U+04BA CYRILLIC CAPITAL LETTER SHHA
	XK_Cyrillic_SHHA = 0x10004ba
	//  U+04BB CYRILLIC SMALL LETTER SHHA
	XK_Cyrillic_shha = 0x10004bb
	//  U+04D8 CYRILLIC CAPITAL LETTER SCHWA
	XK_Cyrillic_SCHWA = 0x10004d8
	//  U+04D9 CYRILLIC SMALL LETTER SCHWA
	XK_Cyrillic_schwa = 0x10004d9
	//  U+04E2 CYRILLIC CAPITAL LETTER I WITH MACRON
	XK_Cyrillic_I_macron = 0x10004e2
	//  U+04E3 CYRILLIC SMALL LETTER I WITH MACRON
	XK_Cyrillic_i_macron = 0x10004e3
	//  U+04E8 CYRILLIC CAPITAL LETTER BARRED O
	XK_Cyrillic_O_bar = 0x10004e8
	//  U+04E9 CYRILLIC SMALL LETTER BARRED O
	XK_Cyrillic_o_bar = 0x10004e9
	//  U+04EE CYRILLIC CAPITAL LETTER U WITH MACRON
	XK_Cyrillic_U_macron = 0x10004ee
	//  U+04EF CYRILLIC SMALL LETTER U WITH MACRON
	XK_Cyrillic_u_macron = 0x10004ef
	//  U+0452 CYRILLIC SMALL LETTER DJE
	XK_Serbian_dje = 0x06a1
	//  U+0453 CYRILLIC SMALL LETTER GJE
	XK_Macedonia_gje = 0x06a2
	//  U+0451 CYRILLIC SMALL LETTER IO
	XK_Cyrillic_io = 0x06a3
	//  U+0454 CYRILLIC SMALL LETTER UKRAINIAN IE
	XK_Ukrainian_ie = 0x06a4
	//  deprecated
	XK_Ukranian_je = 0x06a4
	//  U+0455 CYRILLIC SMALL LETTER DZE
	XK_Macedonia_dse = 0x06a5
	//  U+0456 CYRILLIC SMALL LETTER BYELORUSSIAN-UKRAINIAN I
	XK_Ukrainian_i = 0x06a6
	//  deprecated
	XK_Ukranian_i = 0x06a6
	//  U+0457 CYRILLIC SMALL LETTER YI
	XK_Ukrainian_yi = 0x06a7
	//  deprecated
	XK_Ukranian_yi = 0x06a7
	//  U+0458 CYRILLIC SMALL LETTER JE
	XK_Cyrillic_je = 0x06a8
	//  deprecated
	XK_Serbian_je = 0x06a8
	//  U+0459 CYRILLIC SMALL LETTER LJE
	XK_Cyrillic_lje = 0x06a9
	//  deprecated
	XK_Serbian_lje = 0x06a9
	//  U+045A CYRILLIC SMALL LETTER NJE
	XK_Cyrillic_nje = 0x06aa
	//  deprecated
	XK_Serbian_nje = 0x06aa
	//  U+045B CYRILLIC SMALL LETTER TSHE
	XK_Serbian_tshe = 0x06ab
	//  U+045C CYRILLIC SMALL LETTER KJE
	XK_Macedonia_kje = 0x06ac
	//  U+0491 CYRILLIC SMALL LETTER GHE WITH UPTURN
	XK_Ukrainian_ghe_with_upturn = 0x06ad
	//  U+045E CYRILLIC SMALL LETTER SHORT U
	XK_Byelorussian_shortu = 0x06ae
	//  U+045F CYRILLIC SMALL LETTER DZHE
	XK_Cyrillic_dzhe = 0x06af
	//  deprecated
	XK_Serbian_dze = 0x06af
	//  U+2116 NUMERO SIGN
	XK_numerosign = 0x06b0
	//  U+0402 CYRILLIC CAPITAL LETTER DJE
	XK_Serbian_DJE = 0x06b1
	//  U+0403 CYRILLIC CAPITAL LETTER GJE
	XK_Macedonia_GJE = 0x06b2
	//  U+0401 CYRILLIC CAPITAL LETTER IO
	XK_Cyrillic_IO = 0x06b3
	//  U+0404 CYRILLIC CAPITAL LETTER UKRAINIAN IE
	XK_Ukrainian_IE = 0x06b4
	//  deprecated
	XK_Ukranian_JE = 0x06b4
	//  U+0405 CYRILLIC CAPITAL LETTER DZE
	XK_Macedonia_DSE = 0x06b5
	//  U+0406 CYRILLIC CAPITAL LETTER BYELORUSSIAN-UKRAINIAN I
	XK_Ukrainian_I = 0x06b6
	//  deprecated
	XK_Ukranian_I = 0x06b6
	//  U+0407 CYRILLIC CAPITAL LETTER YI
	XK_Ukrainian_YI = 0x06b7
	//  deprecated
	XK_Ukranian_YI = 0x06b7
	//  U+0408 CYRILLIC CAPITAL LETTER JE
	XK_Cyrillic_JE = 0x06b8
	//  deprecated
	XK_Serbian_JE = 0x06b8
	//  U+0409 CYRILLIC CAPITAL LETTER LJE
	XK_Cyrillic_LJE = 0x06b9
	//  deprecated
	XK_Serbian_LJE = 0x06b9
	//  U+040A CYRILLIC CAPITAL LETTER NJE
	XK_Cyrillic_NJE = 0x06ba
	//  deprecated
	XK_Serbian_NJE = 0x06ba
	//  U+040B CYRILLIC CAPITAL LETTER TSHE
	XK_Serbian_TSHE = 0x06bb
	//  U+040C CYRILLIC CAPITAL LETTER KJE
	XK_Macedonia_KJE = 0x06bc
	//  U+0490 CYRILLIC CAPITAL LETTER GHE WITH UPTURN
	XK_Ukrainian_GHE_WITH_UPTURN = 0x06bd
	//  U+040E CYRILLIC CAPITAL LETTER SHORT U
	XK_Byelorussian_SHORTU = 0x06be
	//  U+040F CYRILLIC CAPITAL LETTER DZHE
	XK_Cyrillic_DZHE = 0x06bf
	//  deprecated
	XK_Serbian_DZE = 0x06bf
	//  U+044E CYRILLIC SMALL LETTER YU
	XK_Cyrillic_yu = 0x06c0
	//  U+0430 CYRILLIC SMALL LETTER A
	XK_Cyrillic_a = 0x06c1
	//  U+0431 CYRILLIC SMALL LETTER BE
	XK_Cyrillic_be = 0x06c2
	//  U+0446 CYRILLIC SMALL LETTER TSE
	XK_Cyrillic_tse = 0x06c3
	//  U+0434 CYRILLIC SMALL LETTER DE
	XK_Cyrillic_de = 0x06c4
	//  U+0435 CYRILLIC SMALL LETTER IE
	XK_Cyrillic_ie = 0x06c5
	//  U+0444 CYRILLIC SMALL LETTER EF
	XK_Cyrillic_ef = 0x06c6
	//  U+0433 CYRILLIC SMALL LETTER GHE
	XK_Cyrillic_ghe = 0x06c7
	//  U+0445 CYRILLIC SMALL LETTER HA
	XK_Cyrillic_ha = 0x06c8
	//  U+0438 CYRILLIC SMALL LETTER I
	XK_Cyrillic_i = 0x06c9
	//  U+0439 CYRILLIC SMALL LETTER SHORT I
	XK_Cyrillic_shorti = 0x06ca
	//  U+043A CYRILLIC SMALL LETTER KA
	XK_Cyrillic_ka = 0x06cb
	//  U+043B CYRILLIC SMALL LETTER EL
	XK_Cyrillic_el = 0x06cc
	//  U+043C CYRILLIC SMALL LETTER EM
	XK_Cyrillic_em = 0x06cd
	//  U+043D CYRILLIC SMALL LETTER EN
	XK_Cyrillic_en = 0x06ce
	//  U+043E CYRILLIC SMALL LETTER O
	XK_Cyrillic_o = 0x06cf
	//  U+043F CYRILLIC SMALL LETTER PE
	XK_Cyrillic_pe = 0x06d0
	//  U+044F CYRILLIC SMALL LETTER YA
	XK_Cyrillic_ya = 0x06d1
	//  U+0440 CYRILLIC SMALL LETTER ER
	XK_Cyrillic_er = 0x06d2
	//  U+0441 CYRILLIC SMALL LETTER ES
	XK_Cyrillic_es = 0x06d3
	//  U+0442 CYRILLIC SMALL LETTER TE
	XK_Cyrillic_te = 0x06d4
	//  U+0443 CYRILLIC SMALL LETTER U
	XK_Cyrillic_u = 0x06d5
	//  U+0436 CYRILLIC SMALL LETTER ZHE
	XK_Cyrillic_zhe = 0x06d6
	//  U+0432 CYRILLIC SMALL LETTER VE
	XK_Cyrillic_ve = 0x06d7
	//  U+044C CYRILLIC SMALL LETTER SOFT SIGN
	XK_Cyrillic_softsign = 0x06d8
	//  U+044B CYRILLIC SMALL LETTER YERU
	XK_Cyrillic_yeru = 0x06d9
	//  U+0437 CYRILLIC SMALL LETTER ZE
	XK_Cyrillic_ze = 0x06da
	//  U+0448 CYRILLIC SMALL LETTER SHA
	XK_Cyrillic_sha = 0x06db
	//  U+044D CYRILLIC SMALL LETTER E
	XK_Cyrillic_e = 0x06dc
	//  U+0449 CYRILLIC SMALL LETTER SHCHA
	XK_Cyrillic_shcha = 0x06dd
	//  U+0447 CYRILLIC SMALL LETTER CHE
	XK_Cyrillic_che = 0x06de
	//  U+044A CYRILLIC SMALL LETTER HARD SIGN
	XK_Cyrillic_hardsign = 0x06df
	//  U+042E CYRILLIC CAPITAL LETTER YU
	XK_Cyrillic_YU = 0x06e0
	//  U+0410 CYRILLIC CAPITAL LETTER A
	XK_Cyrillic_A = 0x06e1
	//  U+0411 CYRILLIC CAPITAL LETTER BE
	XK_Cyrillic_BE = 0x06e2
	//  U+0426 CYRILLIC CAPITAL LETTER TSE
	XK_Cyrillic_TSE = 0x06e3
	//  U+0414 CYRILLIC CAPITAL LETTER DE
	XK_Cyrillic_DE = 0x06e4
	//  U+0415 CYRILLIC CAPITAL LETTER IE
	XK_Cyrillic_IE = 0x06e5
	//  U+0424 CYRILLIC CAPITAL LETTER EF
	XK_Cyrillic_EF = 0x06e6
	//  U+0413 CYRILLIC CAPITAL LETTER GHE
	XK_Cyrillic_GHE = 0x06e7
	//  U+0425 CYRILLIC CAPITAL LETTER HA
	XK_Cyrillic_HA = 0x06e8
	//  U+0418 CYRILLIC CAPITAL LETTER I
	XK_Cyrillic_I = 0x06e9
	//  U+0419 CYRILLIC CAPITAL LETTER SHORT I
	XK_Cyrillic_SHORTI = 0x06ea
	//  U+041A CYRILLIC CAPITAL LETTER KA
	XK_Cyrillic_KA = 0x06eb
	//  U+041B CYRILLIC CAPITAL LETTER EL
	XK_Cyrillic_EL = 0x06ec
	//  U+041C CYRILLIC CAPITAL LETTER EM
	XK_Cyrillic_EM = 0x06ed
	//  U+041D CYRILLIC CAPITAL LETTER EN
	XK_Cyrillic_EN = 0x06ee
	//  U+041E CYRILLIC CAPITAL LETTER O
	XK_Cyrillic_O = 0x06ef
	//  U+041F CYRILLIC CAPITAL LETTER PE
	XK_Cyrillic_PE = 0x06f0
	//  U+042F CYRILLIC CAPITAL LETTER YA
	XK_Cyrillic_YA = 0x06f1
	//  U+0420 CYRILLIC CAPITAL LETTER ER
	XK_Cyrillic_ER = 0x06f2
	//  U+0421 CYRILLIC CAPITAL LETTER ES
	XK_Cyrillic_ES = 0x06f3
	//  U+0422 CYRILLIC CAPITAL LETTER TE
	XK_Cyrillic_TE = 0x06f4
	//  U+0423 CYRILLIC CAPITAL LETTER U
	XK_Cyrillic_U = 0x06f5
	//  U+0416 CYRILLIC CAPITAL LETTER ZHE
	XK_Cyrillic_ZHE = 0x06f6
	//  U+0412 CYRILLIC CAPITAL LETTER VE
	XK_Cyrillic_VE = 0x06f7
	//  U+042C CYRILLIC CAPITAL LETTER SOFT SIGN
	XK_Cyrillic_SOFTSIGN = 0x06f8
	//  U+042B CYRILLIC CAPITAL LETTER YERU
	XK_Cyrillic_YERU = 0x06f9
	//  U+0417 CYRILLIC CAPITAL LETTER ZE
	XK_Cyrillic_ZE = 0x06fa
	//  U+0428 CYRILLIC CAPITAL LETTER SHA
	XK_Cyrillic_SHA = 0x06fb
	//  U+042D CYRILLIC CAPITAL LETTER E
	XK_Cyrillic_E = 0x06fc
	//  U+0429 CYRILLIC CAPITAL LETTER SHCHA
	XK_Cyrillic_SHCHA = 0x06fd
	//  U+0427 CYRILLIC CAPITAL LETTER CHE
	XK_Cyrillic_CHE = 0x06fe
	//  U+042A CYRILLIC CAPITAL LETTER HARD SIGN
	XK_Cyrillic_HARDSIGN = 0x06ff
	//  U+0386 GREEK CAPITAL LETTER ALPHA WITH TONOS
	XK_Greek_ALPHAaccent = 0x07a1
	//  U+0388 GREEK CAPITAL LETTER EPSILON WITH TONOS
	XK_Greek_EPSILONaccent = 0x07a2
	//  U+0389 GREEK CAPITAL LETTER ETA WITH TONOS
	XK_Greek_ETAaccent = 0x07a3
	//  U+038A GREEK CAPITAL LETTER IOTA WITH TONOS
	XK_Greek_IOTAaccent = 0x07a4
	//  U+03AA GREEK CAPITAL LETTER IOTA WITH DIALYTIKA
	XK_Greek_IOTAdieresis = 0x07a5
	//  old typo
	XK_Greek_IOTAdiaeresis = 0x07a5
	//  U+038C GREEK CAPITAL LETTER OMICRON WITH TONOS
	XK_Greek_OMICRONaccent = 0x07a7
	//  U+038E GREEK CAPITAL LETTER UPSILON WITH TONOS
	XK_Greek_UPSILONaccent = 0x07a8
	//  U+03AB GREEK CAPITAL LETTER UPSILON WITH DIALYTIKA
	XK_Greek_UPSILONdieresis = 0x07a9
	//  U+038F GREEK CAPITAL LETTER OMEGA WITH TONOS
	XK_Greek_OMEGAaccent = 0x07ab
	//  U+0385 GREEK DIALYTIKA TONOS
	XK_Greek_accentdieresis = 0x07ae
	//  U+2015 HORIZONTAL BAR
	XK_Greek_horizbar = 0x07af
	//  U+03AC GREEK SMALL LETTER ALPHA WITH TONOS
	XK_Greek_alphaaccent = 0x07b1
	//  U+03AD GREEK SMALL LETTER EPSILON WITH TONOS
	XK_Greek_epsilonaccent = 0x07b2
	//  U+03AE GREEK SMALL LETTER ETA WITH TONOS
	XK_Greek_etaaccent = 0x07b3
	//  U+03AF GREEK SMALL LETTER IOTA WITH TONOS
	XK_Greek_iotaaccent = 0x07b4
	//  U+03CA GREEK SMALL LETTER IOTA WITH DIALYTIKA
	XK_Greek_iotadieresis = 0x07b5
	//  U+0390 GREEK SMALL LETTER IOTA WITH DIALYTIKA AND TONOS
	XK_Greek_iotaaccentdieresis = 0x07b6
	//  U+03CC GREEK SMALL LETTER OMICRON WITH TONOS
	XK_Greek_omicronaccent = 0x07b7
	//  U+03CD GREEK SMALL LETTER UPSILON WITH TONOS
	XK_Greek_upsilonaccent = 0x07b8
	//  U+03CB GREEK SMALL LETTER UPSILON WITH DIALYTIKA
	XK_Greek_upsilondieresis = 0x07b9
	//  U+03B0 GREEK SMALL LETTER UPSILON WITH DIALYTIKA AND TONOS
	XK_Greek_upsilonaccentdieresis = 0x07ba
	//  U+03CE GREEK SMALL LETTER OMEGA WITH TONOS
	XK_Greek_omegaaccent = 0x07bb
	//  U+0391 GREEK CAPITAL LETTER ALPHA
	XK_Greek_ALPHA = 0x07c1
	//  U+0392 GREEK CAPITAL LETTER BETA
	XK_Greek_BETA = 0x07c2
	//  U+0393 GREEK CAPITAL LETTER GAMMA
	XK_Greek_GAMMA = 0x07c3
	//  U+0394 GREEK CAPITAL LETTER DELTA
	XK_Greek_DELTA = 0x07c4
	//  U+0395 GREEK CAPITAL LETTER EPSILON
	XK_Greek_EPSILON = 0x07c5
	//  U+0396 GREEK CAPITAL LETTER ZETA
	XK_Greek_ZETA = 0x07c6
	//  U+0397 GREEK CAPITAL LETTER ETA
	XK_Greek_ETA = 0x07c7
	//  U+0398 GREEK CAPITAL LETTER THETA
	XK_Greek_THETA = 0x07c8
	//  U+0399 GREEK CAPITAL LETTER IOTA
	XK_Greek_IOTA = 0x07c9
	//  U+039A GREEK CAPITAL LETTER KAPPA
	XK_Greek_KAPPA = 0x07ca
	//  U+039B GREEK CAPITAL LETTER LAMDA
	XK_Greek_LAMDA = 0x07cb
	//  U+039B GREEK CAPITAL LETTER LAMDA
	XK_Greek_LAMBDA = 0x07cb
	//  U+039C GREEK CAPITAL LETTER MU
	XK_Greek_MU = 0x07cc
	//  U+039D GREEK CAPITAL LETTER NU
	XK_Greek_NU = 0x07cd
	//  U+039E GREEK CAPITAL LETTER XI
	XK_Greek_XI = 0x07ce
	//  U+039F GREEK CAPITAL LETTER OMICRON
	XK_Greek_OMICRON = 0x07cf
	//  U+03A0 GREEK CAPITAL LETTER PI
	XK_Greek_PI = 0x07d0
	//  U+03A1 GREEK CAPITAL LETTER RHO
	XK_Greek_RHO = 0x07d1
	//  U+03A3 GREEK CAPITAL LETTER SIGMA
	XK_Greek_SIGMA = 0x07d2
	//  U+03A4 GREEK CAPITAL LETTER TAU
	XK_Greek_TAU = 0x07d4
	//  U+03A5 GREEK CAPITAL LETTER UPSILON
	XK_Greek_UPSILON = 0x07d5
	//  U+03A6 GREEK CAPITAL LETTER PHI
	XK_Greek_PHI = 0x07d6
	//  U+03A7 GREEK CAPITAL LETTER CHI
	XK_Greek_CHI = 0x07d7
	//  U+03A8 GREEK CAPITAL LETTER PSI
	XK_Greek_PSI = 0x07d8
	//  U+03A9 GREEK CAPITAL LETTER OMEGA
	XK_Greek_OMEGA = 0x07d9
	//  U+03B1 GREEK SMALL LETTER ALPHA
	XK_Greek_alpha = 0x07e1
	//  U+03B2 GREEK SMALL LETTER BETA
	XK_Greek_beta = 0x07e2
	//  U+03B3 GREEK SMALL LETTER GAMMA
	XK_Greek_gamma = 0x07e3
	//  U+03B4 GREEK SMALL LETTER DELTA
	XK_Greek_delta = 0x07e4
	//  U+03B5 GREEK SMALL LETTER EPSILON
	XK_Greek_epsilon = 0x07e5
	//  U+03B6 GREEK SMALL LETTER ZETA
	XK_Greek_zeta = 0x07e6
	//  U+03B7 GREEK SMALL LETTER ETA
	XK_Greek_eta = 0x07e7
	//  U+03B8 GREEK SMALL LETTER THETA
	XK_Greek_theta = 0x07e8
	//  U+03B9 GREEK SMALL LETTER IOTA
	XK_Greek_iota = 0x07e9
	//  U+03BA GREEK SMALL LETTER KAPPA
	XK_Greek_kappa = 0x07ea
	//  U+03BB GREEK SMALL LETTER LAMDA
	XK_Greek_lamda = 0x07eb
	//  U+03BB GREEK SMALL LETTER LAMDA
	XK_Greek_lambda = 0x07eb
	//  U+03BC GREEK SMALL LETTER MU
	XK_Greek_mu = 0x07ec
	//  U+03BD GREEK SMALL LETTER NU
	XK_Greek_nu = 0x07ed
	//  U+03BE GREEK SMALL LETTER XI
	XK_Greek_xi = 0x07ee
	//  U+03BF GREEK SMALL LETTER OMICRON
	XK_Greek_omicron = 0x07ef
	//  U+03C0 GREEK SMALL LETTER PI
	XK_Greek_pi = 0x07f0
	//  U+03C1 GREEK SMALL LETTER RHO
	XK_Greek_rho = 0x07f1
	//  U+03C3 GREEK SMALL LETTER SIGMA
	XK_Greek_sigma = 0x07f2
	//  U+03C2 GREEK SMALL LETTER FINAL SIGMA
	XK_Greek_finalsmallsigma = 0x07f3
	//  U+03C4 GREEK SMALL LETTER TAU
	XK_Greek_tau = 0x07f4
	//  U+03C5 GREEK SMALL LETTER UPSILON
	XK_Greek_upsilon = 0x07f5
	//  U+03C6 GREEK SMALL LETTER PHI
	XK_Greek_phi = 0x07f6
	//  U+03C7 GREEK SMALL LETTER CHI
	XK_Greek_chi = 0x07f7
	//  U+03C8 GREEK SMALL LETTER PSI
	XK_Greek_psi = 0x07f8
	//  U+03C9 GREEK SMALL LETTER OMEGA
	XK_Greek_omega = 0x07f9
	//  Alias for mode_switch
	XK_Greek_switch = 0xff7e
	//  U+2017 DOUBLE LOW LINE
	XK_hebrew_doublelowline = 0x0cdf
	//  U+05D0 HEBREW LETTER ALEF
	XK_hebrew_aleph = 0x0ce0
	//  U+05D1 HEBREW LETTER BET
	XK_hebrew_bet = 0x0ce1
	//  deprecated
	XK_hebrew_beth = 0x0ce1
	//  U+05D2 HEBREW LETTER GIMEL
	XK_hebrew_gimel = 0x0ce2
	//  deprecated
	XK_hebrew_gimmel = 0x0ce2
	//  U+05D3 HEBREW LETTER DALET
	XK_hebrew_dalet = 0x0ce3
	//  deprecated
	XK_hebrew_daleth = 0x0ce3
	//  U+05D4 HEBREW LETTER HE
	XK_hebrew_he = 0x0ce4
	//  U+05D5 HEBREW LETTER VAV
	XK_hebrew_waw = 0x0ce5
	//  U+05D6 HEBREW LETTER ZAYIN
	XK_hebrew_zain = 0x0ce6
	//  deprecated
	XK_hebrew_zayin = 0x0ce6
	//  U+05D7 HEBREW LETTER HET
	XK_hebrew_chet = 0x0ce7
	//  deprecated
	XK_hebrew_het = 0x0ce7
	//  U+05D8 HEBREW LETTER TET
	XK_hebrew_tet = 0x0ce8
	//  deprecated
	XK_hebrew_teth = 0x0ce8
	//  U+05D9 HEBREW LETTER YOD
	XK_hebrew_yod = 0x0ce9
	//  U+05DA HEBREW LETTER FINAL KAF
	XK_hebrew_finalkaph = 0x0cea
	//  U+05DB HEBREW LETTER KAF
	XK_hebrew_kaph = 0x0ceb
	//  U+05DC HEBREW LETTER LAMED
	XK_hebrew_lamed = 0x0cec
	//  U+05DD HEBREW LETTER FINAL MEM
	XK_hebrew_finalmem = 0x0ced
	//  U+05DE HEBREW LETTER MEM
	XK_hebrew_mem = 0x0cee
	//  U+05DF HEBREW LETTER FINAL NUN
	XK_hebrew_finalnun = 0x0cef
	//  U+05E0 HEBREW LETTER NUN
	XK_hebrew_nun = 0x0cf0
	//  U+05E1 HEBREW LETTER SAMEKH
	XK_hebrew_samech = 0x0cf1
	//  deprecated
	XK_hebrew_samekh = 0x0cf1
	//  U+05E2 HEBREW LETTER AYIN
	XK_hebrew_ayin = 0x0cf2
	//  U+05E3 HEBREW LETTER FINAL PE
	XK_hebrew_finalpe = 0x0cf3
	//  U+05E4 HEBREW LETTER PE
	XK_hebrew_pe = 0x0cf4
	//  U+05E5 HEBREW LETTER FINAL TSADI
	XK_hebrew_finalzade = 0x0cf5
	//  deprecated
	XK_hebrew_finalzadi = 0x0cf5
	//  U+05E6 HEBREW LETTER TSADI
	XK_hebrew_zade = 0x0cf6
	//  deprecated
	XK_hebrew_zadi = 0x0cf6
	//  U+05E7 HEBREW LETTER QOF
	XK_hebrew_qoph = 0x0cf7
	//  deprecated
	XK_hebrew_kuf = 0x0cf7
	//  U+05E8 HEBREW LETTER RESH
	XK_hebrew_resh = 0x0cf8
	//  U+05E9 HEBREW LETTER SHIN
	XK_hebrew_shin = 0x0cf9
	//  U+05EA HEBREW LETTER TAV
	XK_hebrew_taw = 0x0cfa
	//  deprecated
	XK_hebrew_taf = 0x0cfa
	//  Alias for mode_switch
	XK_Hebrew_switch = 0xff7e
	//  U+0E01 THAI CHARACTER KO KAI
	XK_Thai_kokai = 0x0da1
	//  U+0E02 THAI CHARACTER KHO KHAI
	XK_Thai_khokhai = 0x0da2
	//  U+0E03 THAI CHARACTER KHO KHUAT
	XK_Thai_khokhuat = 0x0da3
	//  U+0E04 THAI CHARACTER KHO KHWAI
	XK_Thai_khokhwai = 0x0da4
	//  U+0E05 THAI CHARACTER KHO KHON
	XK_Thai_khokhon = 0x0da5
	//  U+0E06 THAI CHARACTER KHO RAKHANG
	XK_Thai_khorakhang = 0x0da6
	//  U+0E07 THAI CHARACTER NGO NGU
	XK_Thai_ngongu = 0x0da7
	//  U+0E08 THAI CHARACTER CHO CHAN
	XK_Thai_chochan = 0x0da8
	//  U+0E09 THAI CHARACTER CHO CHING
	XK_Thai_choching = 0x0da9
	//  U+0E0A THAI CHARACTER CHO CHANG
	XK_Thai_chochang = 0x0daa
	//  U+0E0B THAI CHARACTER SO SO
	XK_Thai_soso = 0x0dab
	//  U+0E0C THAI CHARACTER CHO CHOE
	XK_Thai_chochoe = 0x0dac
	//  U+0E0D THAI CHARACTER YO YING
	XK_Thai_yoying = 0x0dad
	//  U+0E0E THAI CHARACTER DO CHADA
	XK_Thai_dochada = 0x0dae
	//  U+0E0F THAI CHARACTER TO PATAK
	XK_Thai_topatak = 0x0daf
	//  U+0E10 THAI CHARACTER THO THAN
	XK_Thai_thothan = 0x0db0
	//  U+0E11 THAI CHARACTER THO NANGMONTHO
	XK_Thai_thonangmontho = 0x0db1
	//  U+0E12 THAI CHARACTER THO PHUTHAO
	XK_Thai_thophuthao = 0x0db2
	//  U+0E13 THAI CHARACTER NO NEN
	XK_Thai_nonen = 0x0db3
	//  U+0E14 THAI CHARACTER DO DEK
	XK_Thai_dodek = 0x0db4
	//  U+0E15 THAI CHARACTER TO TAO
	XK_Thai_totao = 0x0db5
	//  U+0E16 THAI CHARACTER THO THUNG
	XK_Thai_thothung = 0x0db6
	//  U+0E17 THAI CHARACTER THO THAHAN
	XK_Thai_thothahan = 0x0db7
	//  U+0E18 THAI CHARACTER THO THONG
	XK_Thai_thothong = 0x0db8
	//  U+0E19 THAI CHARACTER NO NU
	XK_Thai_nonu = 0x0db9
	//  U+0E1A THAI CHARACTER BO BAIMAI
	XK_Thai_bobaimai = 0x0dba
	//  U+0E1B THAI CHARACTER PO PLA
	XK_Thai_popla = 0x0dbb
	//  U+0E1C THAI CHARACTER PHO PHUNG
	XK_Thai_phophung = 0x0dbc
	//  U+0E1D THAI CHARACTER FO FA
	XK_Thai_fofa = 0x0dbd
	//  U+0E1E THAI CHARACTER PHO PHAN
	XK_Thai_phophan = 0x0dbe
	//  U+0E1F THAI CHARACTER FO FAN
	XK_Thai_fofan = 0x0dbf
	//  U+0E20 THAI CHARACTER PHO SAMPHAO
	XK_Thai_phosamphao = 0x0dc0
	//  U+0E21 THAI CHARACTER MO MA
	XK_Thai_moma = 0x0dc1
	//  U+0E22 THAI CHARACTER YO YAK
	XK_Thai_yoyak = 0x0dc2
	//  U+0E23 THAI CHARACTER RO RUA
	XK_Thai_rorua = 0x0dc3
	//  U+0E24 THAI CHARACTER RU
	XK_Thai_ru = 0x0dc4
	//  U+0E25 THAI CHARACTER LO LING
	XK_Thai_loling = 0x0dc5
	//  U+0E26 THAI CHARACTER LU
	XK_Thai_lu = 0x0dc6
	//  U+0E27 THAI CHARACTER WO WAEN
	XK_Thai_wowaen = 0x0dc7
	//  U+0E28 THAI CHARACTER SO SALA
	XK_Thai_sosala = 0x0dc8
	//  U+0E29 THAI CHARACTER SO RUSI
	XK_Thai_sorusi = 0x0dc9
	//  U+0E2A THAI CHARACTER SO SUA
	XK_Thai_sosua = 0x0dca
	//  U+0E2B THAI CHARACTER HO HIP
	XK_Thai_hohip = 0x0dcb
	//  U+0E2C THAI CHARACTER LO CHULA
	XK_Thai_lochula = 0x0dcc
	//  U+0E2D THAI CHARACTER O ANG
	XK_Thai_oang = 0x0dcd
	//  U+0E2E THAI CHARACTER HO NOKHUK
	XK_Thai_honokhuk = 0x0dce
	//  U+0E2F THAI CHARACTER PAIYANNOI
	XK_Thai_paiyannoi = 0x0dcf
	//  U+0E30 THAI CHARACTER SARA A
	XK_Thai_saraa = 0x0dd0
	//  U+0E31 THAI CHARACTER MAI HAN-AKAT
	XK_Thai_maihanakat = 0x0dd1
	//  U+0E32 THAI CHARACTER SARA AA
	XK_Thai_saraaa = 0x0dd2
	//  U+0E33 THAI CHARACTER SARA AM
	XK_Thai_saraam = 0x0dd3
	//  U+0E34 THAI CHARACTER SARA I
	XK_Thai_sarai = 0x0dd4
	//  U+0E35 THAI CHARACTER SARA II
	XK_Thai_saraii = 0x0dd5
	//  U+0E36 THAI CHARACTER SARA UE
	XK_Thai_saraue = 0x0dd6
	//  U+0E37 THAI CHARACTER SARA UEE
	XK_Thai_sarauee = 0x0dd7
	//  U+0E38 THAI CHARACTER SARA U
	XK_Thai_sarau = 0x0dd8
	//  U+0E39 THAI CHARACTER SARA UU
	XK_Thai_sarauu = 0x0dd9
	//  U+0E3A THAI CHARACTER PHINTHU
	XK_Thai_phinthu           = 0x0dda
	XK_Thai_maihanakat_maitho = 0x0dde
	//  U+0E3F THAI CURRENCY SYMBOL BAHT
	XK_Thai_baht = 0x0ddf
	//  U+0E40 THAI CHARACTER SARA E
	XK_Thai_sarae = 0x0de0
	//  U+0E41 THAI CHARACTER SARA AE
	XK_Thai_saraae = 0x0de1
	//  U+0E42 THAI CHARACTER SARA O
	XK_Thai_sarao = 0x0de2
	//  U+0E43 THAI CHARACTER SARA AI MAIMUAN
	XK_Thai_saraaimaimuan = 0x0de3
	//  U+0E44 THAI CHARACTER SARA AI MAIMALAI
	XK_Thai_saraaimaimalai = 0x0de4
	//  U+0E45 THAI CHARACTER LAKKHANGYAO
	XK_Thai_lakkhangyao = 0x0de5
	//  U+0E46 THAI CHARACTER MAIYAMOK
	XK_Thai_maiyamok = 0x0de6
	//  U+0E47 THAI CHARACTER MAITAIKHU
	XK_Thai_maitaikhu = 0x0de7
	//  U+0E48 THAI CHARACTER MAI EK
	XK_Thai_maiek = 0x0de8
	//  U+0E49 THAI CHARACTER MAI THO
	XK_Thai_maitho = 0x0de9
	//  U+0E4A THAI CHARACTER MAI TRI
	XK_Thai_maitri = 0x0dea
	//  U+0E4B THAI CHARACTER MAI CHATTAWA
	XK_Thai_maichattawa = 0x0deb
	//  U+0E4C THAI CHARACTER THANTHAKHAT
	XK_Thai_thanthakhat = 0x0dec
	//  U+0E4D THAI CHARACTER NIKHAHIT
	XK_Thai_nikhahit = 0x0ded
	//  U+0E50 THAI DIGIT ZERO
	XK_Thai_leksun = 0x0df0
	//  U+0E51 THAI DIGIT ONE
	XK_Thai_leknung = 0x0df1
	//  U+0E52 THAI DIGIT TWO
	XK_Thai_leksong = 0x0df2
	//  U+0E53 THAI DIGIT THREE
	XK_Thai_leksam = 0x0df3
	//  U+0E54 THAI DIGIT FOUR
	XK_Thai_leksi = 0x0df4
	//  U+0E55 THAI DIGIT FIVE
	XK_Thai_lekha = 0x0df5
	//  U+0E56 THAI DIGIT SIX
	XK_Thai_lekhok = 0x0df6
	//  U+0E57 THAI DIGIT SEVEN
	XK_Thai_lekchet = 0x0df7
	//  U+0E58 THAI DIGIT EIGHT
	XK_Thai_lekpaet = 0x0df8
	//  U+0E59 THAI DIGIT NINE
	XK_Thai_lekkao = 0x0df9
	//  Hangul start/stop(toggle)
	XK_Hangul = 0xff31
	//  Hangul start
	XK_Hangul_Start = 0xff32
	//  Hangul end, English start
	XK_Hangul_End = 0xff33
	//  Start Hangul->Hanja Conversion
	XK_Hangul_Hanja = 0xff34
	//  Hangul Jamo mode
	XK_Hangul_Jamo = 0xff35
	//  Hangul Romaja mode
	XK_Hangul_Romaja = 0xff36
	//  Hangul code input mode
	XK_Hangul_Codeinput = 0xff37
	//  Jeonja mode
	XK_Hangul_Jeonja = 0xff38
	//  Banja mode
	XK_Hangul_Banja = 0xff39
	//  Pre Hanja conversion
	XK_Hangul_PreHanja = 0xff3a
	//  Post Hanja conversion
	XK_Hangul_PostHanja = 0xff3b
	//  Single candidate
	XK_Hangul_SingleCandidate = 0xff3c
	//  Multiple candidate
	XK_Hangul_MultipleCandidate = 0xff3d
	//  Previous candidate
	XK_Hangul_PreviousCandidate = 0xff3e
	//  Special symbols
	XK_Hangul_Special = 0xff3f
	//  Alias for mode_switch
	XK_Hangul_switch              = 0xff7e
	XK_Hangul_Kiyeog              = 0x0ea1
	XK_Hangul_SsangKiyeog         = 0x0ea2
	XK_Hangul_KiyeogSios          = 0x0ea3
	XK_Hangul_Nieun               = 0x0ea4
	XK_Hangul_NieunJieuj          = 0x0ea5
	XK_Hangul_NieunHieuh          = 0x0ea6
	XK_Hangul_Dikeud              = 0x0ea7
	XK_Hangul_SsangDikeud         = 0x0ea8
	XK_Hangul_Rieul               = 0x0ea9
	XK_Hangul_RieulKiyeog         = 0x0eaa
	XK_Hangul_RieulMieum          = 0x0eab
	XK_Hangul_RieulPieub          = 0x0eac
	XK_Hangul_RieulSios           = 0x0ead
	XK_Hangul_RieulTieut          = 0x0eae
	XK_Hangul_RieulPhieuf         = 0x0eaf
	XK_Hangul_RieulHieuh          = 0x0eb0
	XK_Hangul_Mieum               = 0x0eb1
	XK_Hangul_Pieub               = 0x0eb2
	XK_Hangul_SsangPieub          = 0x0eb3
	XK_Hangul_PieubSios           = 0x0eb4
	XK_Hangul_Sios                = 0x0eb5
	XK_Hangul_SsangSios           = 0x0eb6
	XK_Hangul_Ieung               = 0x0eb7
	XK_Hangul_Jieuj               = 0x0eb8
	XK_Hangul_SsangJieuj          = 0x0eb9
	XK_Hangul_Cieuc               = 0x0eba
	XK_Hangul_Khieuq              = 0x0ebb
	XK_Hangul_Tieut               = 0x0ebc
	XK_Hangul_Phieuf              = 0x0ebd
	XK_Hangul_Hieuh               = 0x0ebe
	XK_Hangul_A                   = 0x0ebf
	XK_Hangul_AE                  = 0x0ec0
	XK_Hangul_YA                  = 0x0ec1
	XK_Hangul_YAE                 = 0x0ec2
	XK_Hangul_EO                  = 0x0ec3
	XK_Hangul_E                   = 0x0ec4
	XK_Hangul_YEO                 = 0x0ec5
	XK_Hangul_YE                  = 0x0ec6
	XK_Hangul_O                   = 0x0ec7
	XK_Hangul_WA                  = 0x0ec8
	XK_Hangul_WAE                 = 0x0ec9
	XK_Hangul_OE                  = 0x0eca
	XK_Hangul_YO                  = 0x0ecb
	XK_Hangul_U                   = 0x0ecc
	XK_Hangul_WEO                 = 0x0ecd
	XK_Hangul_WE                  = 0x0ece
	XK_Hangul_WI                  = 0x0ecf
	XK_Hangul_YU                  = 0x0ed0
	XK_Hangul_EU                  = 0x0ed1
	XK_Hangul_YI                  = 0x0ed2
	XK_Hangul_I                   = 0x0ed3
	XK_Hangul_J_Kiyeog            = 0x0ed4
	XK_Hangul_J_SsangKiyeog       = 0x0ed5
	XK_Hangul_J_KiyeogSios        = 0x0ed6
	XK_Hangul_J_Nieun             = 0x0ed7
	XK_Hangul_J_NieunJieuj        = 0x0ed8
	XK_Hangul_J_NieunHieuh        = 0x0ed9
	XK_Hangul_J_Dikeud            = 0x0eda
	XK_Hangul_J_Rieul             = 0x0edb
	XK_Hangul_J_RieulKiyeog       = 0x0edc
	XK_Hangul_J_RieulMieum        = 0x0edd
	XK_Hangul_J_RieulPieub        = 0x0ede
	XK_Hangul_J_RieulSios         = 0x0edf
	XK_Hangul_J_RieulTieut        = 0x0ee0
	XK_Hangul_J_RieulPhieuf       = 0x0ee1
	XK_Hangul_J_RieulHieuh        = 0x0ee2
	XK_Hangul_J_Mieum             = 0x0ee3
	XK_Hangul_J_Pieub             = 0x0ee4
	XK_Hangul_J_PieubSios         = 0x0ee5
	XK_Hangul_J_Sios              = 0x0ee6
	XK_Hangul_J_SsangSios         = 0x0ee7
	XK_Hangul_J_Ieung             = 0x0ee8
	XK_Hangul_J_Jieuj             = 0x0ee9
	XK_Hangul_J_Cieuc             = 0x0eea
	XK_Hangul_J_Khieuq            = 0x0eeb
	XK_Hangul_J_Tieut             = 0x0eec
	XK_Hangul_J_Phieuf            = 0x0eed
	XK_Hangul_J_Hieuh             = 0x0eee
	XK_Hangul_RieulYeorinHieuh    = 0x0eef
	XK_Hangul_SunkyeongeumMieum   = 0x0ef0
	XK_Hangul_SunkyeongeumPieub   = 0x0ef1
	XK_Hangul_PanSios             = 0x0ef2
	XK_Hangul_KkogjiDalrinIeung   = 0x0ef3
	XK_Hangul_SunkyeongeumPhieuf  = 0x0ef4
	XK_Hangul_YeorinHieuh         = 0x0ef5
	XK_Hangul_AraeA               = 0x0ef6
	XK_Hangul_AraeAE              = 0x0ef7
	XK_Hangul_J_PanSios           = 0x0ef8
	XK_Hangul_J_KkogjiDalrinIeung = 0x0ef9
	XK_Hangul_J_YeorinHieuh       = 0x0efa
	//  U+20A9 WON SIGN
	XK_Korean_Won = 0x0eff
	//  U+0587 ARMENIAN SMALL LIGATURE ECH YIWN
	XK_Armenian_ligature_ew = 0x1000587
	//  U+0589 ARMENIAN FULL STOP
	XK_Armenian_full_stop = 0x1000589
	//  U+0589 ARMENIAN FULL STOP
	XK_Armenian_verjaket = 0x1000589
	//  U+055D ARMENIAN COMMA
	XK_Armenian_separation_mark = 0x100055d
	//  U+055D ARMENIAN COMMA
	XK_Armenian_but = 0x100055d
	//  U+058A ARMENIAN HYPHEN
	XK_Armenian_hyphen = 0x100058a
	//  U+058A ARMENIAN HYPHEN
	XK_Armenian_yentamna = 0x100058a
	//  U+055C ARMENIAN EXCLAMATION MARK
	XK_Armenian_exclam = 0x100055c
	//  U+055C ARMENIAN EXCLAMATION MARK
	XK_Armenian_amanak = 0x100055c
	//  U+055B ARMENIAN EMPHASIS MARK
	XK_Armenian_accent = 0x100055b
	//  U+055B ARMENIAN EMPHASIS MARK
	XK_Armenian_shesht = 0x100055b
	//  U+055E ARMENIAN QUESTION MARK
	XK_Armenian_question = 0x100055e
	//  U+055E ARMENIAN QUESTION MARK
	XK_Armenian_paruyk = 0x100055e
	//  U+0531 ARMENIAN CAPITAL LETTER AYB
	XK_Armenian_AYB = 0x1000531
	//  U+0561 ARMENIAN SMALL LETTER AYB
	XK_Armenian_ayb = 0x1000561
	//  U+0532 ARMENIAN CAPITAL LETTER BEN
	XK_Armenian_BEN = 0x1000532
	//  U+0562 ARMENIAN SMALL LETTER BEN
	XK_Armenian_ben = 0x1000562
	//  U+0533 ARMENIAN CAPITAL LETTER GIM
	XK_Armenian_GIM = 0x1000533
	//  U+0563 ARMENIAN SMALL LETTER GIM
	XK_Armenian_gim = 0x1000563
	//  U+0534 ARMENIAN CAPITAL LETTER DA
	XK_Armenian_DA = 0x1000534
	//  U+0564 ARMENIAN SMALL LETTER DA
	XK_Armenian_da = 0x1000564
	//  U+0535 ARMENIAN CAPITAL LETTER ECH
	XK_Armenian_YECH = 0x1000535
	//  U+0565 ARMENIAN SMALL LETTER ECH
	XK_Armenian_yech = 0x1000565
	//  U+0536 ARMENIAN CAPITAL LETTER ZA
	XK_Armenian_ZA = 0x1000536
	//  U+0566 ARMENIAN SMALL LETTER ZA
	XK_Armenian_za = 0x1000566
	//  U+0537 ARMENIAN CAPITAL LETTER EH
	XK_Armenian_E = 0x1000537
	//  U+0567 ARMENIAN SMALL LETTER EH
	XK_Armenian_e = 0x1000567
	//  U+0538 ARMENIAN CAPITAL LETTER ET
	XK_Armenian_AT = 0x1000538
	//  U+0568 ARMENIAN SMALL LETTER ET
	XK_Armenian_at = 0x1000568
	//  U+0539 ARMENIAN CAPITAL LETTER TO
	XK_Armenian_TO = 0x1000539
	//  U+0569 ARMENIAN SMALL LETTER TO
	XK_Armenian_to = 0x1000569
	//  U+053A ARMENIAN CAPITAL LETTER ZHE
	XK_Armenian_ZHE = 0x100053a
	//  U+056A ARMENIAN SMALL LETTER ZHE
	XK_Armenian_zhe = 0x100056a
	//  U+053B ARMENIAN CAPITAL LETTER INI
	XK_Armenian_INI = 0x100053b
	//  U+056B ARMENIAN SMALL LETTER INI
	XK_Armenian_ini = 0x100056b
	//  U+053C ARMENIAN CAPITAL LETTER LIWN
	XK_Armenian_LYUN = 0x100053c
	//  U+056C ARMENIAN SMALL LETTER LIWN
	XK_Armenian_lyun = 0x100056c
	//  U+053D ARMENIAN CAPITAL LETTER XEH
	XK_Armenian_KHE = 0x100053d
	//  U+056D ARMENIAN SMALL LETTER XEH
	XK_Armenian_khe = 0x100056d
	//  U+053E ARMENIAN CAPITAL LETTER CA
	XK_Armenian_TSA = 0x100053e
	//  U+056E ARMENIAN SMALL LETTER CA
	XK_Armenian_tsa = 0x100056e
	//  U+053F ARMENIAN CAPITAL LETTER KEN
	XK_Armenian_KEN = 0x100053f
	//  U+056F ARMENIAN SMALL LETTER KEN
	XK_Armenian_ken = 0x100056f
	//  U+0540 ARMENIAN CAPITAL LETTER HO
	XK_Armenian_HO = 0x1000540
	//  U+0570 ARMENIAN SMALL LETTER HO
	XK_Armenian_ho = 0x1000570
	//  U+0541 ARMENIAN CAPITAL LETTER JA
	XK_Armenian_DZA = 0x1000541
	//  U+0571 ARMENIAN SMALL LETTER JA
	XK_Armenian_dza = 0x1000571
	//  U+0542 ARMENIAN CAPITAL LETTER GHAD
	XK_Armenian_GHAT = 0x1000542
	//  U+0572 ARMENIAN SMALL LETTER GHAD
	XK_Armenian_ghat = 0x1000572
	//  U+0543 ARMENIAN CAPITAL LETTER CHEH
	XK_Armenian_TCHE = 0x1000543
	//  U+0573 ARMENIAN SMALL LETTER CHEH
	XK_Armenian_tche = 0x1000573
	//  U+0544 ARMENIAN CAPITAL LETTER MEN
	XK_Armenian_MEN = 0x1000544
	//  U+0574 ARMENIAN SMALL LETTER MEN
	XK_Armenian_men = 0x1000574
	//  U+0545 ARMENIAN CAPITAL LETTER YI
	XK_Armenian_HI = 0x1000545
	//  U+0575 ARMENIAN SMALL LETTER YI
	XK_Armenian_hi = 0x1000575
	//  U+0546 ARMENIAN CAPITAL LETTER NOW
	XK_Armenian_NU = 0x1000546
	//  U+0576 ARMENIAN SMALL LETTER NOW
	XK_Armenian_nu = 0x1000576
	//  U+0547 ARMENIAN CAPITAL LETTER SHA
	XK_Armenian_SHA = 0x1000547
	//  U+0577 ARMENIAN SMALL LETTER SHA
	XK_Armenian_sha = 0x1000577
	//  U+0548 ARMENIAN CAPITAL LETTER VO
	XK_Armenian_VO = 0x1000548
	//  U+0578 ARMENIAN SMALL LETTER VO
	XK_Armenian_vo = 0x1000578
	//  U+0549 ARMENIAN CAPITAL LETTER CHA
	XK_Armenian_CHA = 0x1000549
	//  U+0579 ARMENIAN SMALL LETTER CHA
	XK_Armenian_cha = 0x1000579
	//  U+054A ARMENIAN CAPITAL LETTER PEH
	XK_Armenian_PE = 0x100054a
	//  U+057A ARMENIAN SMALL LETTER PEH
	XK_Armenian_pe = 0x100057a
	//  U+054B ARMENIAN CAPITAL LETTER JHEH
	XK_Armenian_JE = 0x100054b
	//  U+057B ARMENIAN SMALL LETTER JHEH
	XK_Armenian_je = 0x100057b
	//  U+054C ARMENIAN CAPITAL LETTER RA
	XK_Armenian_RA = 0x100054c
	//  U+057C ARMENIAN SMALL LETTER RA
	XK_Armenian_ra = 0x100057c
	//  U+054D ARMENIAN CAPITAL LETTER SEH
	XK_Armenian_SE = 0x100054d
	//  U+057D ARMENIAN SMALL LETTER SEH
	XK_Armenian_se = 0x100057d
	//  U+054E ARMENIAN CAPITAL LETTER VEW
	XK_Armenian_VEV = 0x100054e
	//  U+057E ARMENIAN SMALL LETTER VEW
	XK_Armenian_vev = 0x100057e
	//  U+054F ARMENIAN CAPITAL LETTER TIWN
	XK_Armenian_TYUN = 0x100054f
	//  U+057F ARMENIAN SMALL LETTER TIWN
	XK_Armenian_tyun = 0x100057f
	//  U+0550 ARMENIAN CAPITAL LETTER REH
	XK_Armenian_RE = 0x1000550
	//  U+0580 ARMENIAN SMALL LETTER REH
	XK_Armenian_re = 0x1000580
	//  U+0551 ARMENIAN CAPITAL LETTER CO
	XK_Armenian_TSO = 0x1000551
	//  U+0581 ARMENIAN SMALL LETTER CO
	XK_Armenian_tso = 0x1000581
	//  U+0552 ARMENIAN CAPITAL LETTER YIWN
	XK_Armenian_VYUN = 0x1000552
	//  U+0582 ARMENIAN SMALL LETTER YIWN
	XK_Armenian_vyun = 0x1000582
	//  U+0553 ARMENIAN CAPITAL LETTER PIWR
	XK_Armenian_PYUR = 0x1000553
	//  U+0583 ARMENIAN SMALL LETTER PIWR
	XK_Armenian_pyur = 0x1000583
	//  U+0554 ARMENIAN CAPITAL LETTER KEH
	XK_Armenian_KE = 0x1000554
	//  U+0584 ARMENIAN SMALL LETTER KEH
	XK_Armenian_ke = 0x1000584
	//  U+0555 ARMENIAN CAPITAL LETTER OH
	XK_Armenian_O = 0x1000555
	//  U+0585 ARMENIAN SMALL LETTER OH
	XK_Armenian_o = 0x1000585
	//  U+0556 ARMENIAN CAPITAL LETTER FEH
	XK_Armenian_FE = 0x1000556
	//  U+0586 ARMENIAN SMALL LETTER FEH
	XK_Armenian_fe = 0x1000586
	//  U+055A ARMENIAN APOSTROPHE
	XK_Armenian_apostrophe = 0x100055a
	//  U+10D0 GEORGIAN LETTER AN
	XK_Georgian_an = 0x10010d0
	//  U+10D1 GEORGIAN LETTER BAN
	XK_Georgian_ban = 0x10010d1
	//  U+10D2 GEORGIAN LETTER GAN
	XK_Georgian_gan = 0x10010d2
	//  U+10D3 GEORGIAN LETTER DON
	XK_Georgian_don = 0x10010d3
	//  U+10D4 GEORGIAN LETTER EN
	XK_Georgian_en = 0x10010d4
	//  U+10D5 GEORGIAN LETTER VIN
	XK_Georgian_vin = 0x10010d5
	//  U+10D6 GEORGIAN LETTER ZEN
	XK_Georgian_zen = 0x10010d6
	//  U+10D7 GEORGIAN LETTER TAN
	XK_Georgian_tan = 0x10010d7
	//  U+10D8 GEORGIAN LETTER IN
	XK_Georgian_in = 0x10010d8
	//  U+10D9 GEORGIAN LETTER KAN
	XK_Georgian_kan = 0x10010d9
	//  U+10DA GEORGIAN LETTER LAS
	XK_Georgian_las = 0x10010da
	//  U+10DB GEORGIAN LETTER MAN
	XK_Georgian_man = 0x10010db
	//  U+10DC GEORGIAN LETTER NAR
	XK_Georgian_nar = 0x10010dc
	//  U+10DD GEORGIAN LETTER ON
	XK_Georgian_on = 0x10010dd
	//  U+10DE GEORGIAN LETTER PAR
	XK_Georgian_par = 0x10010de
	//  U+10DF GEORGIAN LETTER ZHAR
	XK_Georgian_zhar = 0x10010df
	//  U+10E0 GEORGIAN LETTER RAE
	XK_Georgian_rae = 0x10010e0
	//  U+10E1 GEORGIAN LETTER SAN
	XK_Georgian_san = 0x10010e1
	//  U+10E2 GEORGIAN LETTER TAR
	XK_Georgian_tar = 0x10010e2
	//  U+10E3 GEORGIAN LETTER UN
	XK_Georgian_un = 0x10010e3
	//  U+10E4 GEORGIAN LETTER PHAR
	XK_Georgian_phar = 0x10010e4
	//  U+10E5 GEORGIAN LETTER KHAR
	XK_Georgian_khar = 0x10010e5
	//  U+10E6 GEORGIAN LETTER GHAN
	XK_Georgian_ghan = 0x10010e6
	//  U+10E7 GEORGIAN LETTER QAR
	XK_Georgian_qar = 0x10010e7
	//  U+10E8 GEORGIAN LETTER SHIN
	XK_Georgian_shin = 0x10010e8
	//  U+10E9 GEORGIAN LETTER CHIN
	XK_Georgian_chin = 0x10010e9
	//  U+10EA GEORGIAN LETTER CAN
	XK_Georgian_can = 0x10010ea
	//  U+10EB GEORGIAN LETTER JIL
	XK_Georgian_jil = 0x10010eb
	//  U+10EC GEORGIAN LETTER CIL
	XK_Georgian_cil = 0x10010ec
	//  U+10ED GEORGIAN LETTER CHAR
	XK_Georgian_char = 0x10010ed
	//  U+10EE GEORGIAN LETTER XAN
	XK_Georgian_xan = 0x10010ee
	//  U+10EF GEORGIAN LETTER JHAN
	XK_Georgian_jhan = 0x10010ef
	//  U+10F0 GEORGIAN LETTER HAE
	XK_Georgian_hae = 0x10010f0
	//  U+10F1 GEORGIAN LETTER HE
	XK_Georgian_he = 0x10010f1
	//  U+10F2 GEORGIAN LETTER HIE
	XK_Georgian_hie = 0x10010f2
	//  U+10F3 GEORGIAN LETTER WE
	XK_Georgian_we = 0x10010f3
	//  U+10F4 GEORGIAN LETTER HAR
	XK_Georgian_har = 0x10010f4
	//  U+10F5 GEORGIAN LETTER HOE
	XK_Georgian_hoe = 0x10010f5
	//  U+10F6 GEORGIAN LETTER FI
	XK_Georgian_fi = 0x10010f6
	//  U+1E8A LATIN CAPITAL LETTER X WITH DOT ABOVE
	XK_Xabovedot = 0x1001e8a
	//  U+012C LATIN CAPITAL LETTER I WITH BREVE
	XK_Ibreve = 0x100012c
	//  U+01B5 LATIN CAPITAL LETTER Z WITH STROKE
	XK_Zstroke = 0x10001b5
	//  U+01E6 LATIN CAPITAL LETTER G WITH CARON
	XK_Gcaron = 0x10001e6
	//  U+01D2 LATIN CAPITAL LETTER O WITH CARON
	XK_Ocaron = 0x10001d1
	//  U+019F LATIN CAPITAL LETTER O WITH MIDDLE TILDE
	XK_Obarred = 0x100019f
	//  U+1E8B LATIN SMALL LETTER X WITH DOT ABOVE
	XK_xabovedot = 0x1001e8b
	//  U+012D LATIN SMALL LETTER I WITH BREVE
	XK_ibreve = 0x100012d
	//  U+01B6 LATIN SMALL LETTER Z WITH STROKE
	XK_zstroke = 0x10001b6
	//  U+01E7 LATIN SMALL LETTER G WITH CARON
	XK_gcaron = 0x10001e7
	//  U+01D2 LATIN SMALL LETTER O WITH CARON
	XK_ocaron = 0x10001d2
	//  U+0275 LATIN SMALL LETTER BARRED O
	XK_obarred = 0x1000275
	//  U+018F LATIN CAPITAL LETTER SCHWA
	XK_SCHWA = 0x100018f
	//  U+0259 LATIN SMALL LETTER SCHWA
	XK_schwa = 0x1000259
	//  U+01B7 LATIN CAPITAL LETTER EZH
	XK_EZH = 0x10001b7
	//  U+0292 LATIN SMALL LETTER EZH
	XK_ezh = 0x1000292
	//  U+1E36 LATIN CAPITAL LETTER L WITH DOT BELOW
	XK_Lbelowdot = 0x1001e36
	//  U+1E37 LATIN SMALL LETTER L WITH DOT BELOW
	XK_lbelowdot = 0x1001e37
	//  U+1EA0 LATIN CAPITAL LETTER A WITH DOT BELOW
	XK_Abelowdot = 0x1001ea0
	//  U+1EA1 LATIN SMALL LETTER A WITH DOT BELOW
	XK_abelowdot = 0x1001ea1
	//  U+1EA2 LATIN CAPITAL LETTER A WITH HOOK ABOVE
	XK_Ahook = 0x1001ea2
	//  U+1EA3 LATIN SMALL LETTER A WITH HOOK ABOVE
	XK_ahook = 0x1001ea3
	//  U+1EA4 LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND ACUTE
	XK_Acircumflexacute = 0x1001ea4
	//  U+1EA5 LATIN SMALL LETTER A WITH CIRCUMFLEX AND ACUTE
	XK_acircumflexacute = 0x1001ea5
	//  U+1EA6 LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND GRAVE
	XK_Acircumflexgrave = 0x1001ea6
	//  U+1EA7 LATIN SMALL LETTER A WITH CIRCUMFLEX AND GRAVE
	XK_acircumflexgrave = 0x1001ea7
	//  U+1EA8 LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND HOOK ABOVE
	XK_Acircumflexhook = 0x1001ea8
	//  U+1EA9 LATIN SMALL LETTER A WITH CIRCUMFLEX AND HOOK ABOVE
	XK_acircumflexhook = 0x1001ea9
	//  U+1EAA LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND TILDE
	XK_Acircumflextilde = 0x1001eaa
	//  U+1EAB LATIN SMALL LETTER A WITH CIRCUMFLEX AND TILDE
	XK_acircumflextilde = 0x1001eab
	//  U+1EAC LATIN CAPITAL LETTER A WITH CIRCUMFLEX AND DOT BELOW
	XK_Acircumflexbelowdot = 0x1001eac
	//  U+1EAD LATIN SMALL LETTER A WITH CIRCUMFLEX AND DOT BELOW
	XK_acircumflexbelowdot = 0x1001ead
	//  U+1EAE LATIN CAPITAL LETTER A WITH BREVE AND ACUTE
	XK_Abreveacute = 0x1001eae
	//  U+1EAF LATIN SMALL LETTER A WITH BREVE AND ACUTE
	XK_abreveacute = 0x1001eaf
	//  U+1EB0 LATIN CAPITAL LETTER A WITH BREVE AND GRAVE
	XK_Abrevegrave = 0x1001eb0
	//  U+1EB1 LATIN SMALL LETTER A WITH BREVE AND GRAVE
	XK_abrevegrave = 0x1001eb1
	//  U+1EB2 LATIN CAPITAL LETTER A WITH BREVE AND HOOK ABOVE
	XK_Abrevehook = 0x1001eb2
	//  U+1EB3 LATIN SMALL LETTER A WITH BREVE AND HOOK ABOVE
	XK_abrevehook = 0x1001eb3
	//  U+1EB4 LATIN CAPITAL LETTER A WITH BREVE AND TILDE
	XK_Abrevetilde = 0x1001eb4
	//  U+1EB5 LATIN SMALL LETTER A WITH BREVE AND TILDE
	XK_abrevetilde = 0x1001eb5
	//  U+1EB6 LATIN CAPITAL LETTER A WITH BREVE AND DOT BELOW
	XK_Abrevebelowdot = 0x1001eb6
	//  U+1EB7 LATIN SMALL LETTER A WITH BREVE AND DOT BELOW
	XK_abrevebelowdot = 0x1001eb7
	//  U+1EB8 LATIN CAPITAL LETTER E WITH DOT BELOW
	XK_Ebelowdot = 0x1001eb8
	//  U+1EB9 LATIN SMALL LETTER E WITH DOT BELOW
	XK_ebelowdot = 0x1001eb9
	//  U+1EBA LATIN CAPITAL LETTER E WITH HOOK ABOVE
	XK_Ehook = 0x1001eba
	//  U+1EBB LATIN SMALL LETTER E WITH HOOK ABOVE
	XK_ehook = 0x1001ebb
	//  U+1EBC LATIN CAPITAL LETTER E WITH TILDE
	XK_Etilde = 0x1001ebc
	//  U+1EBD LATIN SMALL LETTER E WITH TILDE
	XK_etilde = 0x1001ebd
	//  U+1EBE LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND ACUTE
	XK_Ecircumflexacute = 0x1001ebe
	//  U+1EBF LATIN SMALL LETTER E WITH CIRCUMFLEX AND ACUTE
	XK_ecircumflexacute = 0x1001ebf
	//  U+1EC0 LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND GRAVE
	XK_Ecircumflexgrave = 0x1001ec0
	//  U+1EC1 LATIN SMALL LETTER E WITH CIRCUMFLEX AND GRAVE
	XK_ecircumflexgrave = 0x1001ec1
	//  U+1EC2 LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND HOOK ABOVE
	XK_Ecircumflexhook = 0x1001ec2
	//  U+1EC3 LATIN SMALL LETTER E WITH CIRCUMFLEX AND HOOK ABOVE
	XK_ecircumflexhook = 0x1001ec3
	//  U+1EC4 LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND TILDE
	XK_Ecircumflextilde = 0x1001ec4
	//  U+1EC5 LATIN SMALL LETTER E WITH CIRCUMFLEX AND TILDE
	XK_ecircumflextilde = 0x1001ec5
	//  U+1EC6 LATIN CAPITAL LETTER E WITH CIRCUMFLEX AND DOT BELOW
	XK_Ecircumflexbelowdot = 0x1001ec6
	//  U+1EC7 LATIN SMALL LETTER E WITH CIRCUMFLEX AND DOT BELOW
	XK_ecircumflexbelowdot = 0x1001ec7
	//  U+1EC8 LATIN CAPITAL LETTER I WITH HOOK ABOVE
	XK_Ihook = 0x1001ec8
	//  U+1EC9 LATIN SMALL LETTER I WITH HOOK ABOVE
	XK_ihook = 0x1001ec9
	//  U+1ECA LATIN CAPITAL LETTER I WITH DOT BELOW
	XK_Ibelowdot = 0x1001eca
	//  U+1ECB LATIN SMALL LETTER I WITH DOT BELOW
	XK_ibelowdot = 0x1001ecb
	//  U+1ECC LATIN CAPITAL LETTER O WITH DOT BELOW
	XK_Obelowdot = 0x1001ecc
	//  U+1ECD LATIN SMALL LETTER O WITH DOT BELOW
	XK_obelowdot = 0x1001ecd
	//  U+1ECE LATIN CAPITAL LETTER O WITH HOOK ABOVE
	XK_Ohook = 0x1001ece
	//  U+1ECF LATIN SMALL LETTER O WITH HOOK ABOVE
	XK_ohook = 0x1001ecf
	//  U+1ED0 LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND ACUTE
	XK_Ocircumflexacute = 0x1001ed0
	//  U+1ED1 LATIN SMALL LETTER O WITH CIRCUMFLEX AND ACUTE
	XK_ocircumflexacute = 0x1001ed1
	//  U+1ED2 LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND GRAVE
	XK_Ocircumflexgrave = 0x1001ed2
	//  U+1ED3 LATIN SMALL LETTER O WITH CIRCUMFLEX AND GRAVE
	XK_ocircumflexgrave = 0x1001ed3
	//  U+1ED4 LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND HOOK ABOVE
	XK_Ocircumflexhook = 0x1001ed4
	//  U+1ED5 LATIN SMALL LETTER O WITH CIRCUMFLEX AND HOOK ABOVE
	XK_ocircumflexhook = 0x1001ed5
	//  U+1ED6 LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND TILDE
	XK_Ocircumflextilde = 0x1001ed6
	//  U+1ED7 LATIN SMALL LETTER O WITH CIRCUMFLEX AND TILDE
	XK_ocircumflextilde = 0x1001ed7
	//  U+1ED8 LATIN CAPITAL LETTER O WITH CIRCUMFLEX AND DOT BELOW
	XK_Ocircumflexbelowdot = 0x1001ed8
	//  U+1ED9 LATIN SMALL LETTER O WITH CIRCUMFLEX AND DOT BELOW
	XK_ocircumflexbelowdot = 0x1001ed9
	//  U+1EDA LATIN CAPITAL LETTER O WITH HORN AND ACUTE
	XK_Ohornacute = 0x1001eda
	//  U+1EDB LATIN SMALL LETTER O WITH HORN AND ACUTE
	XK_ohornacute = 0x1001edb
	//  U+1EDC LATIN CAPITAL LETTER O WITH HORN AND GRAVE
	XK_Ohorngrave = 0x1001edc
	//  U+1EDD LATIN SMALL LETTER O WITH HORN AND GRAVE
	XK_ohorngrave = 0x1001edd
	//  U+1EDE LATIN CAPITAL LETTER O WITH HORN AND HOOK ABOVE
	XK_Ohornhook = 0x1001ede
	//  U+1EDF LATIN SMALL LETTER O WITH HORN AND HOOK ABOVE
	XK_ohornhook = 0x1001edf
	//  U+1EE0 LATIN CAPITAL LETTER O WITH HORN AND TILDE
	XK_Ohorntilde = 0x1001ee0
	//  U+1EE1 LATIN SMALL LETTER O WITH HORN AND TILDE
	XK_ohorntilde = 0x1001ee1
	//  U+1EE2 LATIN CAPITAL LETTER O WITH HORN AND DOT BELOW
	XK_Ohornbelowdot = 0x1001ee2
	//  U+1EE3 LATIN SMALL LETTER O WITH HORN AND DOT BELOW
	XK_ohornbelowdot = 0x1001ee3
	//  U+1EE4 LATIN CAPITAL LETTER U WITH DOT BELOW
	XK_Ubelowdot = 0x1001ee4
	//  U+1EE5 LATIN SMALL LETTER U WITH DOT BELOW
	XK_ubelowdot = 0x1001ee5
	//  U+1EE6 LATIN CAPITAL LETTER U WITH HOOK ABOVE
	XK_Uhook = 0x1001ee6
	//  U+1EE7 LATIN SMALL LETTER U WITH HOOK ABOVE
	XK_uhook = 0x1001ee7
	//  U+1EE8 LATIN CAPITAL LETTER U WITH HORN AND ACUTE
	XK_Uhornacute = 0x1001ee8
	//  U+1EE9 LATIN SMALL LETTER U WITH HORN AND ACUTE
	XK_uhornacute = 0x1001ee9
	//  U+1EEA LATIN CAPITAL LETTER U WITH HORN AND GRAVE
	XK_Uhorngrave = 0x1001eea
	//  U+1EEB LATIN SMALL LETTER U WITH HORN AND GRAVE
	XK_uhorngrave = 0x1001eeb
	//  U+1EEC LATIN CAPITAL LETTER U WITH HORN AND HOOK ABOVE
	XK_Uhornhook = 0x1001eec
	//  U+1EED LATIN SMALL LETTER U WITH HORN AND HOOK ABOVE
	XK_uhornhook = 0x1001eed
	//  U+1EEE LATIN CAPITAL LETTER U WITH HORN AND TILDE
	XK_Uhorntilde = 0x1001eee
	//  U+1EEF LATIN SMALL LETTER U WITH HORN AND TILDE
	XK_uhorntilde = 0x1001eef
	//  U+1EF0 LATIN CAPITAL LETTER U WITH HORN AND DOT BELOW
	XK_Uhornbelowdot = 0x1001ef0
	//  U+1EF1 LATIN SMALL LETTER U WITH HORN AND DOT BELOW
	XK_uhornbelowdot = 0x1001ef1
	//  U+1EF4 LATIN CAPITAL LETTER Y WITH DOT BELOW
	XK_Ybelowdot = 0x1001ef4
	//  U+1EF5 LATIN SMALL LETTER Y WITH DOT BELOW
	XK_ybelowdot = 0x1001ef5
	//  U+1EF6 LATIN CAPITAL LETTER Y WITH HOOK ABOVE
	XK_Yhook = 0x1001ef6
	//  U+1EF7 LATIN SMALL LETTER Y WITH HOOK ABOVE
	XK_yhook = 0x1001ef7
	//  U+1EF8 LATIN CAPITAL LETTER Y WITH TILDE
	XK_Ytilde = 0x1001ef8
	//  U+1EF9 LATIN SMALL LETTER Y WITH TILDE
	XK_ytilde = 0x1001ef9
	//  U+01A0 LATIN CAPITAL LETTER O WITH HORN
	XK_Ohorn = 0x10001a0
	//  U+01A1 LATIN SMALL LETTER O WITH HORN
	XK_ohorn = 0x10001a1
	//  U+01AF LATIN CAPITAL LETTER U WITH HORN
	XK_Uhorn = 0x10001af
	//  U+01B0 LATIN SMALL LETTER U WITH HORN
	XK_uhorn = 0x10001b0
	//  U+20A0 EURO-CURRENCY SIGN
	XK_EcuSign = 0x10020a0
	//  U+20A1 COLON SIGN
	XK_ColonSign = 0x10020a1
	//  U+20A2 CRUZEIRO SIGN
	XK_CruzeiroSign = 0x10020a2
	//  U+20A3 FRENCH FRANC SIGN
	XK_FFrancSign = 0x10020a3
	//  U+20A4 LIRA SIGN
	XK_LiraSign = 0x10020a4
	//  U+20A5 MILL SIGN
	XK_MillSign = 0x10020a5
	//  U+20A6 NAIRA SIGN
	XK_NairaSign = 0x10020a6
	//  U+20A7 PESETA SIGN
	XK_PesetaSign = 0x10020a7
	//  U+20A8 RUPEE SIGN
	XK_RupeeSign = 0x10020a8
	//  U+20A9 WON SIGN
	XK_WonSign = 0x10020a9
	//  U+20AA NEW SHEQEL SIGN
	XK_NewSheqelSign = 0x10020aa
	//  U+20AB DONG SIGN
	XK_DongSign = 0x10020ab
	//  U+20AC EURO SIGN
	XK_EuroSign = 0x20ac
	//  U+2070 SUPERSCRIPT ZERO
	XK_zerosuperior = 0x1002070
	//  U+2074 SUPERSCRIPT FOUR
	XK_foursuperior = 0x1002074
	//  U+2075 SUPERSCRIPT FIVE
	XK_fivesuperior = 0x1002075
	//  U+2076 SUPERSCRIPT SIX
	XK_sixsuperior = 0x1002076
	//  U+2077 SUPERSCRIPT SEVEN
	XK_sevensuperior = 0x1002077
	//  U+2078 SUPERSCRIPT EIGHT
	XK_eightsuperior = 0x1002078
	//  U+2079 SUPERSCRIPT NINE
	XK_ninesuperior = 0x1002079
	//  U+2080 SUBSCRIPT ZERO
	XK_zerosubscript = 0x1002080
	//  U+2081 SUBSCRIPT ONE
	XK_onesubscript = 0x1002081
	//  U+2082 SUBSCRIPT TWO
	XK_twosubscript = 0x1002082
	//  U+2083 SUBSCRIPT THREE
	XK_threesubscript = 0x1002083
	//  U+2084 SUBSCRIPT FOUR
	XK_foursubscript = 0x1002084
	//  U+2085 SUBSCRIPT FIVE
	XK_fivesubscript = 0x1002085
	//  U+2086 SUBSCRIPT SIX
	XK_sixsubscript = 0x1002086
	//  U+2087 SUBSCRIPT SEVEN
	XK_sevensubscript = 0x1002087
	//  U+2088 SUBSCRIPT EIGHT
	XK_eightsubscript = 0x1002088
	//  U+2089 SUBSCRIPT NINE
	XK_ninesubscript = 0x1002089
	//  U+2202 PARTIAL DIFFERENTIAL
	XK_partdifferential = 0x1002202
	//  U+2205 NULL SET
	XK_emptyset = 0x1002205
	//  U+2208 ELEMENT OF
	XK_elementof = 0x1002208
	//  U+2209 NOT AN ELEMENT OF
	XK_notelementof = 0x1002209
	XK_containsas   = 0x100220
	XK_squareroot   = 0x100221
	XK_cuberoot     = 0x100221
	XK_fourthroot   = 0x100221
	XK_dintegral    = 0x100222
	XK_tintegral    = 0x100222
	//  U+2235 BECAUSE
	XK_because = 0x1002235
	//  U+2245 ALMOST EQUAL TO
	XK_approxeq = 0x1002248
	//  U+2247 NOT ALMOST EQUAL TO
	XK_notapproxeq = 0x1002247
	//  U+2262 NOT IDENTICAL TO
	XK_notidentical = 0x1002262
	//  U+2263 STRICTLY EQUIVALENT TO
	XK_stricteq       = 0x1002263
	XK_braille_dot_1  = 0xfff1
	XK_braille_dot_2  = 0xfff2
	XK_braille_dot_3  = 0xfff3
	XK_braille_dot_4  = 0xfff4
	XK_braille_dot_5  = 0xfff5
	XK_braille_dot_6  = 0xfff6
	XK_braille_dot_7  = 0xfff7
	XK_braille_dot_8  = 0xfff8
	XK_braille_dot_9  = 0xfff9
	XK_braille_dot_10 = 0xfffa
	//  U+2800 BRAILLE PATTERN BLANK
	XK_braille_blank = 0x1002800
	//  U+2801 BRAILLE PATTERN DOTS-1
	XK_braille_dots_1 = 0x1002801
	//  U+2802 BRAILLE PATTERN DOTS-2
	XK_braille_dots_2 = 0x1002802
	//  U+2803 BRAILLE PATTERN DOTS-12
	XK_braille_dots_12 = 0x1002803
	//  U+2804 BRAILLE PATTERN DOTS-3
	XK_braille_dots_3 = 0x1002804
	//  U+2805 BRAILLE PATTERN DOTS-13
	XK_braille_dots_13 = 0x1002805
	//  U+2806 BRAILLE PATTERN DOTS-23
	XK_braille_dots_23 = 0x1002806
	//  U+2807 BRAILLE PATTERN DOTS-123
	XK_braille_dots_123 = 0x1002807
	//  U+2808 BRAILLE PATTERN DOTS-4
	XK_braille_dots_4 = 0x1002808
	//  U+2809 BRAILLE PATTERN DOTS-14
	XK_braille_dots_14 = 0x1002809
	//  U+280a BRAILLE PATTERN DOTS-24
	XK_braille_dots_24 = 0x100280a
	//  U+280b BRAILLE PATTERN DOTS-124
	XK_braille_dots_124 = 0x100280b
	//  U+280c BRAILLE PATTERN DOTS-34
	XK_braille_dots_34 = 0x100280c
	//  U+280d BRAILLE PATTERN DOTS-134
	XK_braille_dots_134 = 0x100280d
	//  U+280e BRAILLE PATTERN DOTS-234
	XK_braille_dots_234 = 0x100280e
	//  U+280f BRAILLE PATTERN DOTS-1234
	XK_braille_dots_1234 = 0x100280f
	//  U+2810 BRAILLE PATTERN DOTS-5
	XK_braille_dots_5 = 0x1002810
	//  U+2811 BRAILLE PATTERN DOTS-15
	XK_braille_dots_15 = 0x1002811
	//  U+2812 BRAILLE PATTERN DOTS-25
	XK_braille_dots_25 = 0x1002812
	//  U+2813 BRAILLE PATTERN DOTS-125
	XK_braille_dots_125 = 0x1002813
	//  U+2814 BRAILLE PATTERN DOTS-35
	XK_braille_dots_35 = 0x1002814
	//  U+2815 BRAILLE PATTERN DOTS-135
	XK_braille_dots_135 = 0x1002815
	//  U+2816 BRAILLE PATTERN DOTS-235
	XK_braille_dots_235 = 0x1002816
	//  U+2817 BRAILLE PATTERN DOTS-1235
	XK_braille_dots_1235 = 0x1002817
	//  U+2818 BRAILLE PATTERN DOTS-45
	XK_braille_dots_45 = 0x1002818
	//  U+2819 BRAILLE PATTERN DOTS-145
	XK_braille_dots_145 = 0x1002819
	//  U+281a BRAILLE PATTERN DOTS-245
	XK_braille_dots_245 = 0x100281a
	//  U+281b BRAILLE PATTERN DOTS-1245
	XK_braille_dots_1245 = 0x100281b
	//  U+281c BRAILLE PATTERN DOTS-345
	XK_braille_dots_345 = 0x100281c
	//  U+281d BRAILLE PATTERN DOTS-1345
	XK_braille_dots_1345 = 0x100281d
	//  U+281e BRAILLE PATTERN DOTS-2345
	XK_braille_dots_2345 = 0x100281e
	//  U+281f BRAILLE PATTERN DOTS-12345
	XK_braille_dots_12345 = 0x100281f
	//  U+2820 BRAILLE PATTERN DOTS-6
	XK_braille_dots_6 = 0x1002820
	//  U+2821 BRAILLE PATTERN DOTS-16
	XK_braille_dots_16 = 0x1002821
	//  U+2822 BRAILLE PATTERN DOTS-26
	XK_braille_dots_26 = 0x1002822
	//  U+2823 BRAILLE PATTERN DOTS-126
	XK_braille_dots_126 = 0x1002823
	//  U+2824 BRAILLE PATTERN DOTS-36
	XK_braille_dots_36 = 0x1002824
	//  U+2825 BRAILLE PATTERN DOTS-136
	XK_braille_dots_136 = 0x1002825
	//  U+2826 BRAILLE PATTERN DOTS-236
	XK_braille_dots_236 = 0x1002826
	//  U+2827 BRAILLE PATTERN DOTS-1236
	XK_braille_dots_1236 = 0x1002827
	//  U+2828 BRAILLE PATTERN DOTS-46
	XK_braille_dots_46 = 0x1002828
	//  U+2829 BRAILLE PATTERN DOTS-146
	XK_braille_dots_146 = 0x1002829
	//  U+282a BRAILLE PATTERN DOTS-246
	XK_braille_dots_246 = 0x100282a
	//  U+282b BRAILLE PATTERN DOTS-1246
	XK_braille_dots_1246 = 0x100282b
	//  U+282c BRAILLE PATTERN DOTS-346
	XK_braille_dots_346 = 0x100282c
	//  U+282d BRAILLE PATTERN DOTS-1346
	XK_braille_dots_1346 = 0x100282d
	//  U+282e BRAILLE PATTERN DOTS-2346
	XK_braille_dots_2346 = 0x100282e
	//  U+282f BRAILLE PATTERN DOTS-12346
	XK_braille_dots_12346 = 0x100282f
	//  U+2830 BRAILLE PATTERN DOTS-56
	XK_braille_dots_56 = 0x1002830
	//  U+2831 BRAILLE PATTERN DOTS-156
	XK_braille_dots_156 = 0x1002831
	//  U+2832 BRAILLE PATTERN DOTS-256
	XK_braille_dots_256 = 0x1002832
	//  U+2833 BRAILLE PATTERN DOTS-1256
	XK_braille_dots_1256 = 0x1002833
	//  U+2834 BRAILLE PATTERN DOTS-356
	XK_braille_dots_356 = 0x1002834
	//  U+2835 BRAILLE PATTERN DOTS-1356
	XK_braille_dots_1356 = 0x1002835
	//  U+2836 BRAILLE PATTERN DOTS-2356
	XK_braille_dots_2356 = 0x1002836
	//  U+2837 BRAILLE PATTERN DOTS-12356
	XK_braille_dots_12356 = 0x1002837
	//  U+2838 BRAILLE PATTERN DOTS-456
	XK_braille_dots_456 = 0x1002838
	//  U+2839 BRAILLE PATTERN DOTS-1456
	XK_braille_dots_1456 = 0x1002839
	//  U+283a BRAILLE PATTERN DOTS-2456
	XK_braille_dots_2456 = 0x100283a
	//  U+283b BRAILLE PATTERN DOTS-12456
	XK_braille_dots_12456 = 0x100283b
	//  U+283c BRAILLE PATTERN DOTS-3456
	XK_braille_dots_3456 = 0x100283c
	//  U+283d BRAILLE PATTERN DOTS-13456
	XK_braille_dots_13456 = 0x100283d
	//  U+283e BRAILLE PATTERN DOTS-23456
	XK_braille_dots_23456 = 0x100283e
	//  U+283f BRAILLE PATTERN DOTS-123456
	XK_braille_dots_123456 = 0x100283f
	//  U+2840 BRAILLE PATTERN DOTS-7
	XK_braille_dots_7 = 0x1002840
	//  U+2841 BRAILLE PATTERN DOTS-17
	XK_braille_dots_17 = 0x1002841
	//  U+2842 BRAILLE PATTERN DOTS-27
	XK_braille_dots_27 = 0x1002842
	//  U+2843 BRAILLE PATTERN DOTS-127
	XK_braille_dots_127 = 0x1002843
	//  U+2844 BRAILLE PATTERN DOTS-37
	XK_braille_dots_37 = 0x1002844
	//  U+2845 BRAILLE PATTERN DOTS-137
	XK_braille_dots_137 = 0x1002845
	//  U+2846 BRAILLE PATTERN DOTS-237
	XK_braille_dots_237 = 0x1002846
	//  U+2847 BRAILLE PATTERN DOTS-1237
	XK_braille_dots_1237 = 0x1002847
	//  U+2848 BRAILLE PATTERN DOTS-47
	XK_braille_dots_47 = 0x1002848
	//  U+2849 BRAILLE PATTERN DOTS-147
	XK_braille_dots_147 = 0x1002849
	//  U+284a BRAILLE PATTERN DOTS-247
	XK_braille_dots_247 = 0x100284a
	//  U+284b BRAILLE PATTERN DOTS-1247
	XK_braille_dots_1247 = 0x100284b
	//  U+284c BRAILLE PATTERN DOTS-347
	XK_braille_dots_347 = 0x100284c
	//  U+284d BRAILLE PATTERN DOTS-1347
	XK_braille_dots_1347 = 0x100284d
	//  U+284e BRAILLE PATTERN DOTS-2347
	XK_braille_dots_2347 = 0x100284e
	//  U+284f BRAILLE PATTERN DOTS-12347
	XK_braille_dots_12347 = 0x100284f
	//  U+2850 BRAILLE PATTERN DOTS-57
	XK_braille_dots_57 = 0x1002850
	//  U+2851 BRAILLE PATTERN DOTS-157
	XK_braille_dots_157 = 0x1002851
	//  U+2852 BRAILLE PATTERN DOTS-257
	XK_braille_dots_257 = 0x1002852
	//  U+2853 BRAILLE PATTERN DOTS-1257
	XK_braille_dots_1257 = 0x1002853
	//  U+2854 BRAILLE PATTERN DOTS-357
	XK_braille_dots_357 = 0x1002854
	//  U+2855 BRAILLE PATTERN DOTS-1357
	XK_braille_dots_1357 = 0x1002855
	//  U+2856 BRAILLE PATTERN DOTS-2357
	XK_braille_dots_2357 = 0x1002856
	//  U+2857 BRAILLE PATTERN DOTS-12357
	XK_braille_dots_12357 = 0x1002857
	//  U+2858 BRAILLE PATTERN DOTS-457
	XK_braille_dots_457 = 0x1002858
	//  U+2859 BRAILLE PATTERN DOTS-1457
	XK_braille_dots_1457 = 0x1002859
	//  U+285a BRAILLE PATTERN DOTS-2457
	XK_braille_dots_2457 = 0x100285a
	//  U+285b BRAILLE PATTERN DOTS-12457
	XK_braille_dots_12457 = 0x100285b
	//  U+285c BRAILLE PATTERN DOTS-3457
	XK_braille_dots_3457 = 0x100285c
	//  U+285d BRAILLE PATTERN DOTS-13457
	XK_braille_dots_13457 = 0x100285d
	//  U+285e BRAILLE PATTERN DOTS-23457
	XK_braille_dots_23457 = 0x100285e
	//  U+285f BRAILLE PATTERN DOTS-123457
	XK_braille_dots_123457 = 0x100285f
	//  U+2860 BRAILLE PATTERN DOTS-67
	XK_braille_dots_67 = 0x1002860
	//  U+2861 BRAILLE PATTERN DOTS-167
	XK_braille_dots_167 = 0x1002861
	//  U+2862 BRAILLE PATTERN DOTS-267
	XK_braille_dots_267 = 0x1002862
	//  U+2863 BRAILLE PATTERN DOTS-1267
	XK_braille_dots_1267 = 0x1002863
	//  U+2864 BRAILLE PATTERN DOTS-367
	XK_braille_dots_367 = 0x1002864
	//  U+2865 BRAILLE PATTERN DOTS-1367
	XK_braille_dots_1367 = 0x1002865
	//  U+2866 BRAILLE PATTERN DOTS-2367
	XK_braille_dots_2367 = 0x1002866
	//  U+2867 BRAILLE PATTERN DOTS-12367
	XK_braille_dots_12367 = 0x1002867
	//  U+2868 BRAILLE PATTERN DOTS-467
	XK_braille_dots_467 = 0x1002868
	//  U+2869 BRAILLE PATTERN DOTS-1467
	XK_braille_dots_1467 = 0x1002869
	//  U+286a BRAILLE PATTERN DOTS-2467
	XK_braille_dots_2467 = 0x100286a
	//  U+286b BRAILLE PATTERN DOTS-12467
	XK_braille_dots_12467 = 0x100286b
	//  U+286c BRAILLE PATTERN DOTS-3467
	XK_braille_dots_3467 = 0x100286c
	//  U+286d BRAILLE PATTERN DOTS-13467
	XK_braille_dots_13467 = 0x100286d
	//  U+286e BRAILLE PATTERN DOTS-23467
	XK_braille_dots_23467 = 0x100286e
	//  U+286f BRAILLE PATTERN DOTS-123467
	XK_braille_dots_123467 = 0x100286f
	//  U+2870 BRAILLE PATTERN DOTS-567
	XK_braille_dots_567 = 0x1002870
	//  U+2871 BRAILLE PATTERN DOTS-1567
	XK_braille_dots_1567 = 0x1002871
	//  U+2872 BRAILLE PATTERN DOTS-2567
	XK_braille_dots_2567 = 0x1002872
	//  U+2873 BRAILLE PATTERN DOTS-12567
	XK_braille_dots_12567 = 0x1002873
	//  U+2874 BRAILLE PATTERN DOTS-3567
	XK_braille_dots_3567 = 0x1002874
	//  U+2875 BRAILLE PATTERN DOTS-13567
	XK_braille_dots_13567 = 0x1002875
	//  U+2876 BRAILLE PATTERN DOTS-23567
	XK_braille_dots_23567 = 0x1002876
	//  U+2877 BRAILLE PATTERN DOTS-123567
	XK_braille_dots_123567 = 0x1002877
	//  U+2878 BRAILLE PATTERN DOTS-4567
	XK_braille_dots_4567 = 0x1002878
	//  U+2879 BRAILLE PATTERN DOTS-14567
	XK_braille_dots_14567 = 0x1002879
	//  U+287a BRAILLE PATTERN DOTS-24567
	XK_braille_dots_24567 = 0x100287a
	//  U+287b BRAILLE PATTERN DOTS-124567
	XK_braille_dots_124567 = 0x100287b
	//  U+287c BRAILLE PATTERN DOTS-34567
	XK_braille_dots_34567 = 0x100287c
	//  U+287d BRAILLE PATTERN DOTS-134567
	XK_braille_dots_134567 = 0x100287d
	//  U+287e BRAILLE PATTERN DOTS-234567
	XK_braille_dots_234567 = 0x100287e
	//  U+287f BRAILLE PATTERN DOTS-1234567
	XK_braille_dots_1234567 = 0x100287f
	//  U+2880 BRAILLE PATTERN DOTS-8
	XK_braille_dots_8 = 0x1002880
	//  U+2881 BRAILLE PATTERN DOTS-18
	XK_braille_dots_18 = 0x1002881
	//  U+2882 BRAILLE PATTERN DOTS-28
	XK_braille_dots_28 = 0x1002882
	//  U+2883 BRAILLE PATTERN DOTS-128
	XK_braille_dots_128 = 0x1002883
	//  U+2884 BRAILLE PATTERN DOTS-38
	XK_braille_dots_38 = 0x1002884
	//  U+2885 BRAILLE PATTERN DOTS-138
	XK_braille_dots_138 = 0x1002885
	//  U+2886 BRAILLE PATTERN DOTS-238
	XK_braille_dots_238 = 0x1002886
	//  U+2887 BRAILLE PATTERN DOTS-1238
	XK_braille_dots_1238 = 0x1002887
	//  U+2888 BRAILLE PATTERN DOTS-48
	XK_braille_dots_48 = 0x1002888
	//  U+2889 BRAILLE PATTERN DOTS-148
	XK_braille_dots_148 = 0x1002889
	//  U+288a BRAILLE PATTERN DOTS-248
	XK_braille_dots_248 = 0x100288a
	//  U+288b BRAILLE PATTERN DOTS-1248
	XK_braille_dots_1248 = 0x100288b
	//  U+288c BRAILLE PATTERN DOTS-348
	XK_braille_dots_348 = 0x100288c
	//  U+288d BRAILLE PATTERN DOTS-1348
	XK_braille_dots_1348 = 0x100288d
	//  U+288e BRAILLE PATTERN DOTS-2348
	XK_braille_dots_2348 = 0x100288e
	//  U+288f BRAILLE PATTERN DOTS-12348
	XK_braille_dots_12348 = 0x100288f
	//  U+2890 BRAILLE PATTERN DOTS-58
	XK_braille_dots_58 = 0x1002890
	//  U+2891 BRAILLE PATTERN DOTS-158
	XK_braille_dots_158 = 0x1002891
	//  U+2892 BRAILLE PATTERN DOTS-258
	XK_braille_dots_258 = 0x1002892
	//  U+2893 BRAILLE PATTERN DOTS-1258
	XK_braille_dots_1258 = 0x1002893
	//  U+2894 BRAILLE PATTERN DOTS-358
	XK_braille_dots_358 = 0x1002894
	//  U+2895 BRAILLE PATTERN DOTS-1358
	XK_braille_dots_1358 = 0x1002895
	//  U+2896 BRAILLE PATTERN DOTS-2358
	XK_braille_dots_2358 = 0x1002896
	//  U+2897 BRAILLE PATTERN DOTS-12358
	XK_braille_dots_12358 = 0x1002897
	//  U+2898 BRAILLE PATTERN DOTS-458
	XK_braille_dots_458 = 0x1002898
	//  U+2899 BRAILLE PATTERN DOTS-1458
	XK_braille_dots_1458 = 0x1002899
	//  U+289a BRAILLE PATTERN DOTS-2458
	XK_braille_dots_2458 = 0x100289a
	//  U+289b BRAILLE PATTERN DOTS-12458
	XK_braille_dots_12458 = 0x100289b
	//  U+289c BRAILLE PATTERN DOTS-3458
	XK_braille_dots_3458 = 0x100289c
	//  U+289d BRAILLE PATTERN DOTS-13458
	XK_braille_dots_13458 = 0x100289d
	//  U+289e BRAILLE PATTERN DOTS-23458
	XK_braille_dots_23458 = 0x100289e
	//  U+289f BRAILLE PATTERN DOTS-123458
	XK_braille_dots_123458 = 0x100289f
	//  U+28a0 BRAILLE PATTERN DOTS-68
	XK_braille_dots_68 = 0x10028a0
	//  U+28a1 BRAILLE PATTERN DOTS-168
	XK_braille_dots_168 = 0x10028a1
	//  U+28a2 BRAILLE PATTERN DOTS-268
	XK_braille_dots_268 = 0x10028a2
	//  U+28a3 BRAILLE PATTERN DOTS-1268
	XK_braille_dots_1268 = 0x10028a3
	//  U+28a4 BRAILLE PATTERN DOTS-368
	XK_braille_dots_368 = 0x10028a4
	//  U+28a5 BRAILLE PATTERN DOTS-1368
	XK_braille_dots_1368 = 0x10028a5
	//  U+28a6 BRAILLE PATTERN DOTS-2368
	XK_braille_dots_2368 = 0x10028a6
	//  U+28a7 BRAILLE PATTERN DOTS-12368
	XK_braille_dots_12368 = 0x10028a7
	//  U+28a8 BRAILLE PATTERN DOTS-468
	XK_braille_dots_468 = 0x10028a8
	//  U+28a9 BRAILLE PATTERN DOTS-1468
	XK_braille_dots_1468 = 0x10028a9
	//  U+28aa BRAILLE PATTERN DOTS-2468
	XK_braille_dots_2468 = 0x10028aa
	//  U+28ab BRAILLE PATTERN DOTS-12468
	XK_braille_dots_12468 = 0x10028ab
	//  U+28ac BRAILLE PATTERN DOTS-3468
	XK_braille_dots_3468 = 0x10028ac
	//  U+28ad BRAILLE PATTERN DOTS-13468
	XK_braille_dots_13468 = 0x10028ad
	//  U+28ae BRAILLE PATTERN DOTS-23468
	XK_braille_dots_23468 = 0x10028ae
	//  U+28af BRAILLE PATTERN DOTS-123468
	XK_braille_dots_123468 = 0x10028af
	//  U+28b0 BRAILLE PATTERN DOTS-568
	XK_braille_dots_568 = 0x10028b0
	//  U+28b1 BRAILLE PATTERN DOTS-1568
	XK_braille_dots_1568 = 0x10028b1
	//  U+28b2 BRAILLE PATTERN DOTS-2568
	XK_braille_dots_2568 = 0x10028b2
	//  U+28b3 BRAILLE PATTERN DOTS-12568
	XK_braille_dots_12568 = 0x10028b3
	//  U+28b4 BRAILLE PATTERN DOTS-3568
	XK_braille_dots_3568 = 0x10028b4
	//  U+28b5 BRAILLE PATTERN DOTS-13568
	XK_braille_dots_13568 = 0x10028b5
	//  U+28b6 BRAILLE PATTERN DOTS-23568
	XK_braille_dots_23568 = 0x10028b6
	//  U+28b7 BRAILLE PATTERN DOTS-123568
	XK_braille_dots_123568 = 0x10028b7
	//  U+28b8 BRAILLE PATTERN DOTS-4568
	XK_braille_dots_4568 = 0x10028b8
	//  U+28b9 BRAILLE PATTERN DOTS-14568
	XK_braille_dots_14568 = 0x10028b9
	//  U+28ba BRAILLE PATTERN DOTS-24568
	XK_braille_dots_24568 = 0x10028ba
	//  U+28bb BRAILLE PATTERN DOTS-124568
	XK_braille_dots_124568 = 0x10028bb
	//  U+28bc BRAILLE PATTERN DOTS-34568
	XK_braille_dots_34568 = 0x10028bc
	//  U+28bd BRAILLE PATTERN DOTS-134568
	XK_braille_dots_134568 = 0x10028bd
	//  U+28be BRAILLE PATTERN DOTS-234568
	XK_braille_dots_234568 = 0x10028be
	//  U+28bf BRAILLE PATTERN DOTS-1234568
	XK_braille_dots_1234568 = 0x10028bf
	//  U+28c0 BRAILLE PATTERN DOTS-78
	XK_braille_dots_78 = 0x10028c0
	//  U+28c1 BRAILLE PATTERN DOTS-178
	XK_braille_dots_178 = 0x10028c1
	//  U+28c2 BRAILLE PATTERN DOTS-278
	XK_braille_dots_278 = 0x10028c2
	//  U+28c3 BRAILLE PATTERN DOTS-1278
	XK_braille_dots_1278 = 0x10028c3
	//  U+28c4 BRAILLE PATTERN DOTS-378
	XK_braille_dots_378 = 0x10028c4
	//  U+28c5 BRAILLE PATTERN DOTS-1378
	XK_braille_dots_1378 = 0x10028c5
	//  U+28c6 BRAILLE PATTERN DOTS-2378
	XK_braille_dots_2378 = 0x10028c6
	//  U+28c7 BRAILLE PATTERN DOTS-12378
	XK_braille_dots_12378 = 0x10028c7
	//  U+28c8 BRAILLE PATTERN DOTS-478
	XK_braille_dots_478 = 0x10028c8
	//  U+28c9 BRAILLE PATTERN DOTS-1478
	XK_braille_dots_1478 = 0x10028c9
	//  U+28ca BRAILLE PATTERN DOTS-2478
	XK_braille_dots_2478 = 0x10028ca
	//  U+28cb BRAILLE PATTERN DOTS-12478
	XK_braille_dots_12478 = 0x10028cb
	//  U+28cc BRAILLE PATTERN DOTS-3478
	XK_braille_dots_3478 = 0x10028cc
	//  U+28cd BRAILLE PATTERN DOTS-13478
	XK_braille_dots_13478 = 0x10028cd
	//  U+28ce BRAILLE PATTERN DOTS-23478
	XK_braille_dots_23478 = 0x10028ce
	//  U+28cf BRAILLE PATTERN DOTS-123478
	XK_braille_dots_123478 = 0x10028cf
	//  U+28d0 BRAILLE PATTERN DOTS-578
	XK_braille_dots_578 = 0x10028d0
	//  U+28d1 BRAILLE PATTERN DOTS-1578
	XK_braille_dots_1578 = 0x10028d1
	//  U+28d2 BRAILLE PATTERN DOTS-2578
	XK_braille_dots_2578 = 0x10028d2
	//  U+28d3 BRAILLE PATTERN DOTS-12578
	XK_braille_dots_12578 = 0x10028d3
	//  U+28d4 BRAILLE PATTERN DOTS-3578
	XK_braille_dots_3578 = 0x10028d4
	//  U+28d5 BRAILLE PATTERN DOTS-13578
	XK_braille_dots_13578 = 0x10028d5
	//  U+28d6 BRAILLE PATTERN DOTS-23578
	XK_braille_dots_23578 = 0x10028d6
	//  U+28d7 BRAILLE PATTERN DOTS-123578
	XK_braille_dots_123578 = 0x10028d7
	//  U+28d8 BRAILLE PATTERN DOTS-4578
	XK_braille_dots_4578 = 0x10028d8
	//  U+28d9 BRAILLE PATTERN DOTS-14578
	XK_braille_dots_14578 = 0x10028d9
	//  U+28da BRAILLE PATTERN DOTS-24578
	XK_braille_dots_24578 = 0x10028da
	//  U+28db BRAILLE PATTERN DOTS-124578
	XK_braille_dots_124578 = 0x10028db
	//  U+28dc BRAILLE PATTERN DOTS-34578
	XK_braille_dots_34578 = 0x10028dc
	//  U+28dd BRAILLE PATTERN DOTS-134578
	XK_braille_dots_134578 = 0x10028dd
	//  U+28de BRAILLE PATTERN DOTS-234578
	XK_braille_dots_234578 = 0x10028de
	//  U+28df BRAILLE PATTERN DOTS-1234578
	XK_braille_dots_1234578 = 0x10028df
	//  U+28e0 BRAILLE PATTERN DOTS-678
	XK_braille_dots_678 = 0x10028e0
	//  U+28e1 BRAILLE PATTERN DOTS-1678
	XK_braille_dots_1678 = 0x10028e1
	//  U+28e2 BRAILLE PATTERN DOTS-2678
	XK_braille_dots_2678 = 0x10028e2
	//  U+28e3 BRAILLE PATTERN DOTS-12678
	XK_braille_dots_12678 = 0x10028e3
	//  U+28e4 BRAILLE PATTERN DOTS-3678
	XK_braille_dots_3678 = 0x10028e4
	//  U+28e5 BRAILLE PATTERN DOTS-13678
	XK_braille_dots_13678 = 0x10028e5
	//  U+28e6 BRAILLE PATTERN DOTS-23678
	XK_braille_dots_23678 = 0x10028e6
	//  U+28e7 BRAILLE PATTERN DOTS-123678
	XK_braille_dots_123678 = 0x10028e7
	//  U+28e8 BRAILLE PATTERN DOTS-4678
	XK_braille_dots_4678 = 0x10028e8
	//  U+28e9 BRAILLE PATTERN DOTS-14678
	XK_braille_dots_14678 = 0x10028e9
	//  U+28ea BRAILLE PATTERN DOTS-24678
	XK_braille_dots_24678 = 0x10028ea
	//  U+28eb BRAILLE PATTERN DOTS-124678
	XK_braille_dots_124678 = 0x10028eb
	//  U+28ec BRAILLE PATTERN DOTS-34678
	XK_braille_dots_34678 = 0x10028ec
	//  U+28ed BRAILLE PATTERN DOTS-134678
	XK_braille_dots_134678 = 0x10028ed
	//  U+28ee BRAILLE PATTERN DOTS-234678
	XK_braille_dots_234678 = 0x10028ee
	//  U+28ef BRAILLE PATTERN DOTS-1234678
	XK_braille_dots_1234678 = 0x10028ef
	//  U+28f0 BRAILLE PATTERN DOTS-5678
	XK_braille_dots_5678 = 0x10028f0
	//  U+28f1 BRAILLE PATTERN DOTS-15678
	XK_braille_dots_15678 = 0x10028f1
	//  U+28f2 BRAILLE PATTERN DOTS-25678
	XK_braille_dots_25678 = 0x10028f2
	//  U+28f3 BRAILLE PATTERN DOTS-125678
	XK_braille_dots_125678 = 0x10028f3
	//  U+28f4 BRAILLE PATTERN DOTS-35678
	XK_braille_dots_35678 = 0x10028f4
	//  U+28f5 BRAILLE PATTERN DOTS-135678
	XK_braille_dots_135678 = 0x10028f5
	//  U+28f6 BRAILLE PATTERN DOTS-235678
	XK_braille_dots_235678 = 0x10028f6
	//  U+28f7 BRAILLE PATTERN DOTS-1235678
	XK_braille_dots_1235678 = 0x10028f7
	//  U+28f8 BRAILLE PATTERN DOTS-45678
	XK_braille_dots_45678 = 0x10028f8
	//  U+28f9 BRAILLE PATTERN DOTS-145678
	XK_braille_dots_145678 = 0x10028f9
	//  U+28fa BRAILLE PATTERN DOTS-245678
	XK_braille_dots_245678 = 0x10028fa
	//  U+28fb BRAILLE PATTERN DOTS-1245678
	XK_braille_dots_1245678 = 0x10028fb
	//  U+28fc BRAILLE PATTERN DOTS-345678
	XK_braille_dots_345678 = 0x10028fc
	//  U+28fd BRAILLE PATTERN DOTS-1345678
	XK_braille_dots_1345678 = 0x10028fd
	//  U+28fe BRAILLE PATTERN DOTS-2345678
	XK_braille_dots_2345678 = 0x10028fe
	//  U+28ff BRAILLE PATTERN DOTS-12345678
	XK_braille_dots_12345678 = 0x10028ff
	//  U+0D82 SINHALA ANUSVARAYA
	XK_Sinh_ng = 0x1000d82
	//  U+0D83 SINHALA VISARGAYA
	XK_Sinh_h2 = 0x1000d83
	//  U+0D85 SINHALA AYANNA
	XK_Sinh_a = 0x1000d85
	//  U+0D86 SINHALA AAYANNA
	XK_Sinh_aa = 0x1000d86
	//  U+0D87 SINHALA AEYANNA
	XK_Sinh_ae = 0x1000d87
	//  U+0D88 SINHALA AEEYANNA
	XK_Sinh_aee = 0x1000d88
	//  U+0D89 SINHALA IYANNA
	XK_Sinh_i = 0x1000d89
	//  U+0D8A SINHALA IIYANNA
	XK_Sinh_ii = 0x1000d8a
	//  U+0D8B SINHALA UYANNA
	XK_Sinh_u = 0x1000d8b
	//  U+0D8C SINHALA UUYANNA
	XK_Sinh_uu = 0x1000d8c
	//  U+0D8D SINHALA IRUYANNA
	XK_Sinh_ri = 0x1000d8d
	//  U+0D8E SINHALA IRUUYANNA
	XK_Sinh_rii = 0x1000d8e
	//  U+0D8F SINHALA ILUYANNA
	XK_Sinh_lu = 0x1000d8f
	//  U+0D90 SINHALA ILUUYANNA
	XK_Sinh_luu = 0x1000d90
	//  U+0D91 SINHALA EYANNA
	XK_Sinh_e = 0x1000d91
	//  U+0D92 SINHALA EEYANNA
	XK_Sinh_ee = 0x1000d92
	//  U+0D93 SINHALA AIYANNA
	XK_Sinh_ai = 0x1000d93
	//  U+0D94 SINHALA OYANNA
	XK_Sinh_o = 0x1000d94
	//  U+0D95 SINHALA OOYANNA
	XK_Sinh_oo = 0x1000d95
	//  U+0D96 SINHALA AUYANNA
	XK_Sinh_au = 0x1000d96
	//  U+0D9A SINHALA KAYANNA
	XK_Sinh_ka = 0x1000d9a
	//  U+0D9B SINHALA MAHA. KAYANNA
	XK_Sinh_kha = 0x1000d9b
	//  U+0D9C SINHALA GAYANNA
	XK_Sinh_ga = 0x1000d9c
	//  U+0D9D SINHALA MAHA. GAYANNA
	XK_Sinh_gha = 0x1000d9d
	//  U+0D9E SINHALA KANTAJA NAASIKYAYA
	XK_Sinh_ng2 = 0x1000d9e
	//  U+0D9F SINHALA SANYAKA GAYANNA
	XK_Sinh_nga = 0x1000d9f
	//  U+0DA0 SINHALA CAYANNA
	XK_Sinh_ca = 0x1000da0
	//  U+0DA1 SINHALA MAHA. CAYANNA
	XK_Sinh_cha = 0x1000da1
	//  U+0DA2 SINHALA JAYANNA
	XK_Sinh_ja = 0x1000da2
	//  U+0DA3 SINHALA MAHA. JAYANNA
	XK_Sinh_jha = 0x1000da3
	//  U+0DA4 SINHALA TAALUJA NAASIKYAYA
	XK_Sinh_nya = 0x1000da4
	//  U+0DA5 SINHALA TAALUJA SANYOOGA NAASIKYAYA
	XK_Sinh_jnya = 0x1000da5
	//  U+0DA6 SINHALA SANYAKA JAYANNA
	XK_Sinh_nja = 0x1000da6
	//  U+0DA7 SINHALA TTAYANNA
	XK_Sinh_tta = 0x1000da7
	//  U+0DA8 SINHALA MAHA. TTAYANNA
	XK_Sinh_ttha = 0x1000da8
	//  U+0DA9 SINHALA DDAYANNA
	XK_Sinh_dda = 0x1000da9
	//  U+0DAA SINHALA MAHA. DDAYANNA
	XK_Sinh_ddha = 0x1000daa
	//  U+0DAB SINHALA MUURDHAJA NAYANNA
	XK_Sinh_nna = 0x1000dab
	//  U+0DAC SINHALA SANYAKA DDAYANNA
	XK_Sinh_ndda = 0x1000dac
	//  U+0DAD SINHALA TAYANNA
	XK_Sinh_tha = 0x1000dad
	//  U+0DAE SINHALA MAHA. TAYANNA
	XK_Sinh_thha = 0x1000dae
	//  U+0DAF SINHALA DAYANNA
	XK_Sinh_dha = 0x1000daf
	//  U+0DB0 SINHALA MAHA. DAYANNA
	XK_Sinh_dhha = 0x1000db0
	//  U+0DB1 SINHALA DANTAJA NAYANNA
	XK_Sinh_na = 0x1000db1
	//  U+0DB3 SINHALA SANYAKA DAYANNA
	XK_Sinh_ndha = 0x1000db3
	//  U+0DB4 SINHALA PAYANNA
	XK_Sinh_pa = 0x1000db4
	//  U+0DB5 SINHALA MAHA. PAYANNA
	XK_Sinh_pha = 0x1000db5
	//  U+0DB6 SINHALA BAYANNA
	XK_Sinh_ba = 0x1000db6
	//  U+0DB7 SINHALA MAHA. BAYANNA
	XK_Sinh_bha = 0x1000db7
	//  U+0DB8 SINHALA MAYANNA
	XK_Sinh_ma = 0x1000db8
	//  U+0DB9 SINHALA AMBA BAYANNA
	XK_Sinh_mba = 0x1000db9
	//  U+0DBA SINHALA YAYANNA
	XK_Sinh_ya = 0x1000dba
	//  U+0DBB SINHALA RAYANNA
	XK_Sinh_ra = 0x1000dbb
	//  U+0DBD SINHALA DANTAJA LAYANNA
	XK_Sinh_la = 0x1000dbd
	//  U+0DC0 SINHALA VAYANNA
	XK_Sinh_va = 0x1000dc0
	//  U+0DC1 SINHALA TAALUJA SAYANNA
	XK_Sinh_sha = 0x1000dc1
	//  U+0DC2 SINHALA MUURDHAJA SAYANNA
	XK_Sinh_ssha = 0x1000dc2
	//  U+0DC3 SINHALA DANTAJA SAYANNA
	XK_Sinh_sa = 0x1000dc3
	//  U+0DC4 SINHALA HAYANNA
	XK_Sinh_ha = 0x1000dc4
	//  U+0DC5 SINHALA MUURDHAJA LAYANNA
	XK_Sinh_lla = 0x1000dc5
	//  U+0DC6 SINHALA FAYANNA
	XK_Sinh_fa = 0x1000dc6
	//  U+0DCA SINHALA AL-LAKUNA
	XK_Sinh_al = 0x1000dca
	//  U+0DCF SINHALA AELA-PILLA
	XK_Sinh_aa2 = 0x1000dcf
	//  U+0DD0 SINHALA AEDA-PILLA
	XK_Sinh_ae2 = 0x1000dd0
	//  U+0DD1 SINHALA DIGA AEDA-PILLA
	XK_Sinh_aee2 = 0x1000dd1
	//  U+0DD2 SINHALA IS-PILLA
	XK_Sinh_i2 = 0x1000dd2
	//  U+0DD3 SINHALA DIGA IS-PILLA
	XK_Sinh_ii2 = 0x1000dd3
	//  U+0DD4 SINHALA PAA-PILLA
	XK_Sinh_u2 = 0x1000dd4
	//  U+0DD6 SINHALA DIGA PAA-PILLA
	XK_Sinh_uu2 = 0x1000dd6
	//  U+0DD8 SINHALA GAETTA-PILLA
	XK_Sinh_ru2 = 0x1000dd8
	//  U+0DD9 SINHALA KOMBUVA
	XK_Sinh_e2 = 0x1000dd9
	//  U+0DDA SINHALA DIGA KOMBUVA
	XK_Sinh_ee2 = 0x1000dda
	//  U+0DDB SINHALA KOMBU DEKA
	XK_Sinh_ai2 = 0x1000ddb
	//  U+0DDC SINHALA KOMBUVA HAA AELA-PILLA
	XK_Sinh_o2 = 0x1000ddc
	//  U+0DDD SINHALA KOMBUVA HAA DIGA AELA-PILLA
	XK_Sinh_oo2 = 0x1000ddd
	//  U+0DDE SINHALA KOMBUVA HAA GAYANUKITTA
	XK_Sinh_au2 = 0x1000dde
	//  U+0DDF SINHALA GAYANUKITTA
	XK_Sinh_lu2 = 0x1000ddf
	//  U+0DF2 SINHALA DIGA GAETTA-PILLA
	XK_Sinh_ruu2 = 0x1000df2
	//  U+0DF3 SINHALA DIGA GAYANUKITTA
	XK_Sinh_luu2 = 0x1000df3
	//  U+0DF4 SINHALA KUNDDALIYA
	XK_Sinh_kunddaliya = 0x1000df4
	//  Mode Switch Lock
	XF86XK_ModeLock = 0x1008FF01
	//  Monitor/panel brightness
	XF86XK_MonBrightnessUp = 0x1008FF02
	//  Monitor/panel brightness
	XF86XK_MonBrightnessDown = 0x1008FF03
	//  Keyboards may be lit
	XF86XK_KbdLightOnOff = 0x1008FF04
	//  Keyboards may be lit
	XF86XK_KbdBrightnessUp = 0x1008FF05
	//  Keyboards may be lit
	XF86XK_KbdBrightnessDown = 0x1008FF06
	//  System into standby mode
	XF86XK_Standby = 0x1008FF10
	//  Volume control down
	XF86XK_AudioLowerVolume = 0x1008FF11
	//  Mute sound from the system
	XF86XK_AudioMute = 0x1008FF12
	//  Volume control up
	XF86XK_AudioRaiseVolume = 0x1008FF13
	//  Start playing of audio >
	XF86XK_AudioPlay = 0x1008FF14
	//  Stop playing audio
	XF86XK_AudioStop = 0x1008FF15
	//  Previous track
	XF86XK_AudioPrev = 0x1008FF16
	//  Next track
	XF86XK_AudioNext = 0x1008FF17
	//  Display user's home page
	XF86XK_HomePage = 0x1008FF18
	//  Invoke user's mail program
	XF86XK_Mail = 0x1008FF19
	//  Start application
	XF86XK_Start = 0x1008FF1A
	//  Search
	XF86XK_Search = 0x1008FF1B
	//  Record audio application
	XF86XK_AudioRecord = 0x1008FF1C
	//  Invoke calculator program
	XF86XK_Calculator = 0x1008FF1D
	//  Invoke Memo taking program
	XF86XK_Memo = 0x1008FF1E
	//  Invoke To Do List program
	XF86XK_ToDoList = 0x1008FF1F
	//  Invoke Calendar program
	XF86XK_Calendar = 0x1008FF20
	//  Deep sleep the system
	XF86XK_PowerDown = 0x1008FF21
	//  Adjust screen contrast
	XF86XK_ContrastAdjust = 0x1008FF22
	//  Rocker switches exist up
	XF86XK_RockerUp = 0x1008FF23
	//  and down
	XF86XK_RockerDown = 0x1008FF24
	//  and let you press them
	XF86XK_RockerEnter = 0x1008FF25
	//  Like back on a browser
	XF86XK_Back = 0x1008FF26
	//  Like forward on a browser
	XF86XK_Forward = 0x1008FF27
	//  Stop current operation
	XF86XK_Stop = 0x1008FF28
	//  Refresh the page
	XF86XK_Refresh = 0x1008FF29
	//  Power off system entirely
	XF86XK_PowerOff = 0x1008FF2A
	//  Wake up system from sleep
	XF86XK_WakeUp = 0x1008FF2B
	//  Eject device (e.g. DVD)
	XF86XK_Eject = 0x1008FF2C
	//  Invoke screensaver
	XF86XK_ScreenSaver = 0x1008FF2D
	//  Invoke web browser
	XF86XK_WWW = 0x1008FF2E
	//  Put system to sleep
	XF86XK_Sleep = 0x1008FF2F
	//  Show favorite locations
	XF86XK_Favorites = 0x1008FF30
	//  Pause audio playing
	XF86XK_AudioPause = 0x1008FF31
	//  Launch media collection app
	XF86XK_AudioMedia = 0x1008FF32
	//  Display "My Computer" window
	XF86XK_MyComputer = 0x1008FF33
	//  Display vendor home web site
	XF86XK_VendorHome = 0x1008FF34
	//  Light bulb keys exist
	XF86XK_LightBulb = 0x1008FF35
	//  Display shopping web site
	XF86XK_Shop = 0x1008FF36
	//  Show history of web surfing
	XF86XK_History = 0x1008FF37
	//  Open selected URL
	XF86XK_OpenURL = 0x1008FF38
	//  Add URL to favorites list
	XF86XK_AddFavorite = 0x1008FF39
	//  Show "hot" links
	XF86XK_HotLinks = 0x1008FF3A
	//  Invoke brightness adj. UI
	XF86XK_BrightnessAdjust = 0x1008FF3B
	//  Display financial site
	XF86XK_Finance = 0x1008FF3C
	//  Display user's community
	XF86XK_Community = 0x1008FF3D
	//  "rewind" audio track
	XF86XK_AudioRewind = 0x1008FF3E
	//  ???
	XF86XK_BackForward = 0x1008FF3F
	//  Launch Application
	XF86XK_Launch0 = 0x1008FF40
	//  Launch Application
	XF86XK_Launch1 = 0x1008FF41
	//  Launch Application
	XF86XK_Launch2 = 0x1008FF42
	//  Launch Application
	XF86XK_Launch3 = 0x1008FF43
	//  Launch Application
	XF86XK_Launch4 = 0x1008FF44
	//  Launch Application
	XF86XK_Launch5 = 0x1008FF45
	//  Launch Application
	XF86XK_Launch6 = 0x1008FF46
	//  Launch Application
	XF86XK_Launch7 = 0x1008FF47
	//  Launch Application
	XF86XK_Launch8 = 0x1008FF48
	//  Launch Application
	XF86XK_Launch9 = 0x1008FF49
	//  Launch Application
	XF86XK_LaunchA = 0x1008FF4A
	//  Launch Application
	XF86XK_LaunchB = 0x1008FF4B
	//  Launch Application
	XF86XK_LaunchC = 0x1008FF4C
	//  Launch Application
	XF86XK_LaunchD = 0x1008FF4D
	//  Launch Application
	XF86XK_LaunchE = 0x1008FF4E
	//  Launch Application
	XF86XK_LaunchF = 0x1008FF4F
	//  switch to application, left
	XF86XK_ApplicationLeft = 0x1008FF50
	//  switch to application, right
	XF86XK_ApplicationRight = 0x1008FF51
	//  Launch bookreader
	XF86XK_Book = 0x1008FF52
	//  Launch CD/DVD player
	XF86XK_CD = 0x1008FF53
	//  Launch Calculater
	XF86XK_Calculater = 0x1008FF54
	//  Clear window, screen
	XF86XK_Clear = 0x1008FF55
	//  Close window
	XF86XK_Close = 0x1008FF56
	//  Copy selection
	XF86XK_Copy = 0x1008FF57
	//  Cut selection
	XF86XK_Cut = 0x1008FF58
	//  Output switch key
	XF86XK_Display = 0x1008FF59
	//  Launch DOS (emulation)
	XF86XK_DOS = 0x1008FF5A
	//  Open documents window
	XF86XK_Documents = 0x1008FF5B
	//  Launch spread sheet
	XF86XK_Excel = 0x1008FF5C
	//  Launch file explorer
	XF86XK_Explorer = 0x1008FF5D
	//  Launch game
	XF86XK_Game = 0x1008FF5E
	//  Go to URL
	XF86XK_Go = 0x1008FF5F
	//  Logitch iTouch- don't use
	XF86XK_iTouch = 0x1008FF60
	//  Log off system
	XF86XK_LogOff = 0x1008FF61
	//  ??
	XF86XK_Market = 0x1008FF62
	//  enter meeting in calendar
	XF86XK_Meeting = 0x1008FF63
	//  distingush keyboard from PB
	XF86XK_MenuKB = 0x1008FF65
	//  distinuish PB from keyboard
	XF86XK_MenuPB = 0x1008FF66
	//  Favourites
	XF86XK_MySites = 0x1008FF67
	//  New (folder, document...
	XF86XK_New = 0x1008FF68
	//  News
	XF86XK_News = 0x1008FF69
	//  Office home (old Staroffice)
	XF86XK_OfficeHome = 0x1008FF6A
	//  Open
	XF86XK_Open = 0x1008FF6B
	//  ??
	XF86XK_Option = 0x1008FF6C
	//  Paste
	XF86XK_Paste = 0x1008FF6D
	//  Launch phone; dial number
	XF86XK_Phone = 0x1008FF6E
	//  Compaq's Q - don't use
	XF86XK_Q = 0x1008FF70
	//  Reply e.g., mail
	XF86XK_Reply = 0x1008FF72
	//  Reload web page, file, etc.
	XF86XK_Reload = 0x1008FF73
	//  Rotate windows e.g. xrandr
	XF86XK_RotateWindows = 0x1008FF74
	//  don't use
	XF86XK_RotationPB = 0x1008FF75
	//  don't use
	XF86XK_RotationKB = 0x1008FF76
	//  Save (file, document, state
	XF86XK_Save = 0x1008FF77
	//  Scroll window/contents up
	XF86XK_ScrollUp = 0x1008FF78
	//  Scrool window/contentd down
	XF86XK_ScrollDown = 0x1008FF79
	//  Use XKB mousekeys instead
	XF86XK_ScrollClick = 0x1008FF7A
	//  Send mail, file, object
	XF86XK_Send = 0x1008FF7B
	//  Spell checker
	XF86XK_Spell = 0x1008FF7C
	//  Split window or screen
	XF86XK_SplitScreen = 0x1008FF7D
	//  Get support (??)
	XF86XK_Support = 0x1008FF7E
	//  Show tasks
	XF86XK_TaskPane = 0x1008FF7F
	//  Launch terminal emulator
	XF86XK_Terminal = 0x1008FF80
	//  toolbox of desktop/app.
	XF86XK_Tools = 0x1008FF81
	//  ??
	XF86XK_Travel = 0x1008FF82
	//  ??
	XF86XK_UserPB = 0x1008FF84
	//  ??
	XF86XK_User1KB = 0x1008FF85
	//  ??
	XF86XK_User2KB = 0x1008FF86
	//  Launch video player
	XF86XK_Video = 0x1008FF87
	//  button from a mouse wheel
	XF86XK_WheelButton = 0x1008FF88
	//  Launch word processor
	XF86XK_Word = 0x1008FF89
	XF86XK_Xfer = 0x1008FF8A
	//  zoom in view, map, etc.
	XF86XK_ZoomIn = 0x1008FF8B
	//  zoom out view, map, etc.
	XF86XK_ZoomOut = 0x1008FF8C
	//  mark yourself as away
	XF86XK_Away = 0x1008FF8D
	//  as in instant messaging
	XF86XK_Messenger = 0x1008FF8E
	//  Launch web camera app.
	XF86XK_WebCam = 0x1008FF8F
	//  Forward in mail
	XF86XK_MailForward = 0x1008FF90
	//  Show pictures
	XF86XK_Pictures = 0x1008FF91
	//  Launch music application
	XF86XK_Music = 0x1008FF92
	//  Display battery information
	XF86XK_Battery = 0x1008FF93
	//  Enable/disable Bluetooth
	XF86XK_Bluetooth = 0x1008FF94
	//  Enable/disable WLAN
	XF86XK_WLAN = 0x1008FF95
	//  Enable/disable UWB
	XF86XK_UWB = 0x1008FF96
	//  fast-forward audio track
	XF86XK_AudioForward = 0x1008FF97
	//  toggle repeat mode
	XF86XK_AudioRepeat = 0x1008FF98
	//  toggle shuffle mode
	XF86XK_AudioRandomPlay = 0x1008FF99
	//  cycle through subtitle
	XF86XK_Subtitle = 0x1008FF9A
	//  cycle through audio tracks
	XF86XK_AudioCycleTrack = 0x1008FF9B
	//  cycle through angles
	XF86XK_CycleAngle = 0x1008FF9C
	//  video: go one frame back
	XF86XK_FrameBack = 0x1008FF9D
	//  video: go one frame forward
	XF86XK_FrameForward = 0x1008FF9E
	//  display, or shows an entry for time seeking
	XF86XK_Time = 0x1008FF9F
	//  Select button on joypads and remotes
	XF86XK_Select = 0x1008FFA0
	//  Show a view options/properties
	XF86XK_View = 0x1008FFA1
	//  Go to a top-level menu in a video
	XF86XK_TopMenu = 0x1008FFA2
	//  Red button
	XF86XK_Red = 0x1008FFA3
	//  Green button
	XF86XK_Green = 0x1008FFA4
	//  Yellow button
	XF86XK_Yellow = 0x1008FFA5
	//  Blue button
	XF86XK_Blue = 0x1008FFA6
	//  Sleep to RAM
	XF86XK_Suspend = 0x1008FFA7
	//  Sleep to disk
	XF86XK_Hibernate = 0x1008FFA8
	//  Toggle between touchpad/trackstick
	XF86XK_TouchpadToggle = 0x1008FFA9
	//  The touchpad got switched on
	XF86XK_TouchpadOn = 0x1008FFB0
	//  The touchpad got switched off
	XF86XK_TouchpadOff = 0x1008FFB1
	//  Mute the Mic from the system
	XF86XK_AudioMicMute = 0x1008FFB2
	//  User defined keyboard related action
	XF86XK_Keyboard = 0x1008FFB3
	//  Toggle WWAN (LTE, UMTS, etc.) radio
	XF86XK_WWAN = 0x1008FFB4
	//  Toggle radios on/off
	XF86XK_RFKill = 0x1008FFB5
	//  Select equalizer preset, e.g. theatre-mode
	XF86XK_AudioPreset  = 0x1008FFB6
	XF86XK_Switch_VT_1  = 0x1008FE01
	XF86XK_Switch_VT_2  = 0x1008FE02
	XF86XK_Switch_VT_3  = 0x1008FE03
	XF86XK_Switch_VT_4  = 0x1008FE04
	XF86XK_Switch_VT_5  = 0x1008FE05
	XF86XK_Switch_VT_6  = 0x1008FE06
	XF86XK_Switch_VT_7  = 0x1008FE07
	XF86XK_Switch_VT_8  = 0x1008FE08
	XF86XK_Switch_VT_9  = 0x1008FE09
	XF86XK_Switch_VT_10 = 0x1008FE0A
	XF86XK_Switch_VT_11 = 0x1008FE0B
	XF86XK_Switch_VT_12 = 0x1008FE0C
	//  force ungrab
	XF86XK_Ungrab = 0x1008FE20
	//  kill application with grab
	XF86XK_ClearGrab = 0x1008FE21
	//  next video mode available
	XF86XK_Next_VMode = 0x1008FE22
	//  prev. video mode available
	XF86XK_Prev_VMode = 0x1008FE23
	//  print window tree to log
	XF86XK_LogWindowTree = 0x1008FE24
	//  print all active grabs to log
	XF86XK_LogGrabInfo = 0x1008FE25
)

var KeysymEngMap = map[x.Keysym]string{
	XK_VoidSymbol:        "VoidSymbol",
	XK_BackSpace:         "BackSpace",
	XK_Tab:               "Tab",
	XK_Linefeed:          "Linefeed",
	XK_Clear:             "Clear",
	XK_Return:            "Return",
	XK_Pause:             "Pause",
	XK_Scroll_Lock:       "Scroll_Lock",
	XK_Sys_Req:           "Sys_Req",
	XK_Escape:            "Escape",
	XK_Delete:            "Delete",
	XK_Multi_key:         "Multi_key",
	XK_Codeinput:         "Codeinput",
	XK_SingleCandidate:   "SingleCandidate",
	XK_MultipleCandidate: "MultipleCandidate",
	XK_PreviousCandidate: "PreviousCandidate",
	XK_Kanji:             "Kanji",
	XK_Muhenkan:          "Muhenkan",
	XK_Henkan_Mode:       "Henkan_Mode",
	// XK_Henkan == XK_Henkan_Mode # Alias for Henkan_Mode
	XK_Romaji:            "Romaji",
	XK_Hiragana:          "Hiragana",
	XK_Katakana:          "Katakana",
	XK_Hiragana_Katakana: "Hiragana_Katakana",
	XK_Zenkaku:           "Zenkaku",
	XK_Hankaku:           "Hankaku",
	XK_Zenkaku_Hankaku:   "Zenkaku_Hankaku",
	XK_Touroku:           "Touroku",
	XK_Massyo:            "Massyo",
	XK_Kana_Lock:         "Kana_Lock",
	XK_Kana_Shift:        "Kana_Shift",
	XK_Eisu_Shift:        "Eisu_Shift",
	XK_Eisu_toggle:       "Eisu_toggle",
	// XK_Kanji_Bangou == XK_Codeinput # Codeinput
	// XK_Zen_Koho == XK_MultipleCandidate # Multiple/All Candidate(s)
	// XK_Mae_Koho == XK_PreviousCandidate # Previous Candidate
	XK_Home:  "Home",
	XK_Left:  "Left",
	XK_Up:    "Up",
	XK_Right: "Right",
	XK_Down:  "Down",
	XK_Prior: "Prior",
	// XK_Page_Up == XK_Prior #
	XK_Next: "Next",
	// XK_Page_Down == XK_Next #
	XK_End:         "End",
	XK_Begin:       "Begin",
	XK_Select:      "Select",
	XK_Print:       "Print",
	XK_Execute:     "Execute",
	XK_Insert:      "Insert",
	XK_Undo:        "Undo",
	XK_Redo:        "Redo",
	XK_Menu:        "Menu",
	XK_Find:        "Find",
	XK_Cancel:      "Cancel",
	XK_Help:        "Help",
	XK_Break:       "Break",
	XK_Mode_switch: "Mode_switch",
	// XK_script_switch == XK_Mode_switch # Alias for mode_switch
	XK_Num_Lock: "Num_Lock",
	XK_KP_Space: "KP_Space",
	XK_KP_Tab:   "KP_Tab",
	XK_KP_Enter: "KP_Enter",
	XK_KP_F1:    "KP_F1",
	XK_KP_F2:    "KP_F2",
	XK_KP_F3:    "KP_F3",
	XK_KP_F4:    "KP_F4",
	XK_KP_Home:  "KP_Home",
	XK_KP_Left:  "KP_Left",
	XK_KP_Up:    "KP_Up",
	XK_KP_Right: "KP_Right",
	XK_KP_Down:  "KP_Down",
	XK_KP_Prior: "KP_Prior",
	// XK_KP_Page_Up == XK_KP_Prior #
	XK_KP_Next: "KP_Next",
	// XK_KP_Page_Down == XK_KP_Next #
	XK_KP_End:       "KP_End",
	XK_KP_Begin:     "KP_Begin",
	XK_KP_Insert:    "KP_Insert",
	XK_KP_Delete:    "KP_Delete",
	XK_KP_Equal:     "KP_Equal",
	XK_KP_Multiply:  "KP_Multiply",
	XK_KP_Add:       "KP_Add",
	XK_KP_Separator: "KP_Separator",
	XK_KP_Subtract:  "KP_Subtract",
	XK_KP_Decimal:   "KP_Decimal",
	XK_KP_Divide:    "KP_Divide",
	XK_KP_0:         "KP_0",
	XK_KP_1:         "KP_1",
	XK_KP_2:         "KP_2",
	XK_KP_3:         "KP_3",
	XK_KP_4:         "KP_4",
	XK_KP_5:         "KP_5",
	XK_KP_6:         "KP_6",
	XK_KP_7:         "KP_7",
	XK_KP_8:         "KP_8",
	XK_KP_9:         "KP_9",
	XK_F1:           "F1",
	XK_F2:           "F2",
	XK_F3:           "F3",
	XK_F4:           "F4",
	XK_F5:           "F5",
	XK_F6:           "F6",
	XK_F7:           "F7",
	XK_F8:           "F8",
	XK_F9:           "F9",
	XK_F10:          "F10",
	XK_F11:          "F11",
	// XK_L1 == XK_F11 #
	XK_F12: "F12",
	// XK_L2 == XK_F12 #
	XK_F13: "F13",
	// XK_L3 == XK_F13 #
	XK_F14: "F14",
	// XK_L4 == XK_F14 #
	XK_F15: "F15",
	// XK_L5 == XK_F15 #
	XK_F16: "F16",
	// XK_L6 == XK_F16 #
	XK_F17: "F17",
	// XK_L7 == XK_F17 #
	XK_F18: "F18",
	// XK_L8 == XK_F18 #
	XK_F19: "F19",
	// XK_L9 == XK_F19 #
	XK_F20: "F20",
	// XK_L10 == XK_F20 #
	XK_F21: "F21",
	// XK_R1 == XK_F21 #
	XK_F22: "F22",
	// XK_R2 == XK_F22 #
	XK_F23: "F23",
	// XK_R3 == XK_F23 #
	XK_F24: "F24",
	// XK_R4 == XK_F24 #
	XK_F25: "F25",
	// XK_R5 == XK_F25 #
	XK_F26: "F26",
	// XK_R6 == XK_F26 #
	XK_F27: "F27",
	// XK_R7 == XK_F27 #
	XK_F28: "F28",
	// XK_R8 == XK_F28 #
	XK_F29: "F29",
	// XK_R9 == XK_F29 #
	XK_F30: "F30",
	// XK_R10 == XK_F30 #
	XK_F31: "F31",
	// XK_R11 == XK_F31 #
	XK_F32: "F32",
	// XK_R12 == XK_F32 #
	XK_F33: "F33",
	// XK_R13 == XK_F33 #
	XK_F34: "F34",
	// XK_R14 == XK_F34 #
	XK_F35: "F35",
	// XK_R15 == XK_F35 #
	XK_Shift_L:          "Shift_L",
	XK_Shift_R:          "Shift_R",
	XK_Control_L:        "Control_L",
	XK_Control_R:        "Control_R",
	XK_Caps_Lock:        "Caps_Lock",
	XK_Shift_Lock:       "Shift_Lock",
	XK_Meta_L:           "Meta_L",
	XK_Meta_R:           "Meta_R",
	XK_Alt_L:            "Alt_L",
	XK_Alt_R:            "Alt_R",
	XK_Super_L:          "Super_L",
	XK_Super_R:          "Super_R",
	XK_Hyper_L:          "Hyper_L",
	XK_Hyper_R:          "Hyper_R",
	XK_ISO_Lock:         "ISO_Lock",
	XK_ISO_Level2_Latch: "ISO_Level2_Latch",
	XK_ISO_Level3_Shift: "ISO_Level3_Shift",
	XK_ISO_Level3_Latch: "ISO_Level3_Latch",
	XK_ISO_Level3_Lock:  "ISO_Level3_Lock",
	XK_ISO_Level5_Shift: "ISO_Level5_Shift",
	XK_ISO_Level5_Latch: "ISO_Level5_Latch",
	XK_ISO_Level5_Lock:  "ISO_Level5_Lock",
	// XK_ISO_Group_Shift == XK_Mode_switch # Alias for mode_switch
	XK_ISO_Group_Latch:             "ISO_Group_Latch",
	XK_ISO_Group_Lock:              "ISO_Group_Lock",
	XK_ISO_Next_Group:              "ISO_Next_Group",
	XK_ISO_Next_Group_Lock:         "ISO_Next_Group_Lock",
	XK_ISO_Prev_Group:              "ISO_Prev_Group",
	XK_ISO_Prev_Group_Lock:         "ISO_Prev_Group_Lock",
	XK_ISO_First_Group:             "ISO_First_Group",
	XK_ISO_First_Group_Lock:        "ISO_First_Group_Lock",
	XK_ISO_Last_Group:              "ISO_Last_Group",
	XK_ISO_Last_Group_Lock:         "ISO_Last_Group_Lock",
	XK_ISO_Left_Tab:                "ISO_Left_Tab",
	XK_ISO_Move_Line_Up:            "ISO_Move_Line_Up",
	XK_ISO_Move_Line_Down:          "ISO_Move_Line_Down",
	XK_ISO_Partial_Line_Up:         "ISO_Partial_Line_Up",
	XK_ISO_Partial_Line_Down:       "ISO_Partial_Line_Down",
	XK_ISO_Partial_Space_Left:      "ISO_Partial_Space_Left",
	XK_ISO_Partial_Space_Right:     "ISO_Partial_Space_Right",
	XK_ISO_Set_Margin_Left:         "ISO_Set_Margin_Left",
	XK_ISO_Set_Margin_Right:        "ISO_Set_Margin_Right",
	XK_ISO_Release_Margin_Left:     "ISO_Release_Margin_Left",
	XK_ISO_Release_Margin_Right:    "ISO_Release_Margin_Right",
	XK_ISO_Release_Both_Margins:    "ISO_Release_Both_Margins",
	XK_ISO_Fast_Cursor_Left:        "ISO_Fast_Cursor_Left",
	XK_ISO_Fast_Cursor_Right:       "ISO_Fast_Cursor_Right",
	XK_ISO_Fast_Cursor_Up:          "ISO_Fast_Cursor_Up",
	XK_ISO_Fast_Cursor_Down:        "ISO_Fast_Cursor_Down",
	XK_ISO_Continuous_Underline:    "ISO_Continuous_Underline",
	XK_ISO_Discontinuous_Underline: "ISO_Discontinuous_Underline",
	XK_ISO_Emphasize:               "ISO_Emphasize",
	XK_ISO_Center_Object:           "ISO_Center_Object",
	XK_ISO_Enter:                   "ISO_Enter",
	XK_dead_grave:                  "dead_grave",
	XK_dead_acute:                  "dead_acute",
	XK_dead_circumflex:             "dead_circumflex",
	XK_dead_tilde:                  "dead_tilde",
	// XK_dead_perispomeni == XK_dead_tilde # alias for dead_tilde
	XK_dead_macron:           "dead_macron",
	XK_dead_breve:            "dead_breve",
	XK_dead_abovedot:         "dead_abovedot",
	XK_dead_diaeresis:        "dead_diaeresis",
	XK_dead_abovering:        "dead_abovering",
	XK_dead_doubleacute:      "dead_doubleacute",
	XK_dead_caron:            "dead_caron",
	XK_dead_cedilla:          "dead_cedilla",
	XK_dead_ogonek:           "dead_ogonek",
	XK_dead_iota:             "dead_iota",
	XK_dead_voiced_sound:     "dead_voiced_sound",
	XK_dead_semivoiced_sound: "dead_semivoiced_sound",
	XK_dead_belowdot:         "dead_belowdot",
	XK_dead_hook:             "dead_hook",
	XK_dead_horn:             "dead_horn",
	XK_dead_stroke:           "dead_stroke",
	XK_dead_abovecomma:       "dead_abovecomma",
	// XK_dead_psili == XK_dead_abovecomma # alias for dead_abovecomma
	XK_dead_abovereversedcomma: "dead_abovereversedcomma",
	// XK_dead_dasia == XK_dead_abovereversedcomma # alias for dead_abovereversedcomma
	XK_dead_doublegrave:        "dead_doublegrave",
	XK_dead_belowring:          "dead_belowring",
	XK_dead_belowmacron:        "dead_belowmacron",
	XK_dead_belowcircumflex:    "dead_belowcircumflex",
	XK_dead_belowtilde:         "dead_belowtilde",
	XK_dead_belowbreve:         "dead_belowbreve",
	XK_dead_belowdiaeresis:     "dead_belowdiaeresis",
	XK_dead_invertedbreve:      "dead_invertedbreve",
	XK_dead_belowcomma:         "dead_belowcomma",
	XK_dead_currency:           "dead_currency",
	XK_dead_lowline:            "dead_lowline",
	XK_dead_aboveverticalline:  "dead_aboveverticalline",
	XK_dead_belowverticalline:  "dead_belowverticalline",
	XK_dead_longsolidusoverlay: "dead_longsolidusoverlay",
	XK_dead_a:                  "dead_a",
	XK_dead_A:                  "dead_A",
	XK_dead_e:                  "dead_e",
	XK_dead_E:                  "dead_E",
	XK_dead_i:                  "dead_i",
	XK_dead_I:                  "dead_I",
	XK_dead_o:                  "dead_o",
	XK_dead_O:                  "dead_O",
	XK_dead_u:                  "dead_u",
	XK_dead_U:                  "dead_U",
	XK_dead_small_schwa:        "dead_small_schwa",
	XK_dead_capital_schwa:      "dead_capital_schwa",
	XK_dead_greek:              "dead_greek",
	XK_First_Virtual_Screen:    "First_Virtual_Screen",
	XK_Prev_Virtual_Screen:     "Prev_Virtual_Screen",
	XK_Next_Virtual_Screen:     "Next_Virtual_Screen",
	XK_Last_Virtual_Screen:     "Last_Virtual_Screen",
	XK_Terminate_Server:        "Terminate_Server",
	XK_AccessX_Enable:          "AccessX_Enable",
	XK_AccessX_Feedback_Enable: "AccessX_Feedback_Enable",
	XK_RepeatKeys_Enable:       "RepeatKeys_Enable",
	XK_SlowKeys_Enable:         "SlowKeys_Enable",
	XK_BounceKeys_Enable:       "BounceKeys_Enable",
	XK_StickyKeys_Enable:       "StickyKeys_Enable",
	XK_MouseKeys_Enable:        "MouseKeys_Enable",
	XK_MouseKeys_Accel_Enable:  "MouseKeys_Accel_Enable",
	XK_Overlay1_Enable:         "Overlay1_Enable",
	XK_Overlay2_Enable:         "Overlay2_Enable",
	XK_AudibleBell_Enable:      "AudibleBell_Enable",
	XK_Pointer_Left:            "Pointer_Left",
	XK_Pointer_Right:           "Pointer_Right",
	XK_Pointer_Up:              "Pointer_Up",
	XK_Pointer_Down:            "Pointer_Down",
	XK_Pointer_UpLeft:          "Pointer_UpLeft",
	XK_Pointer_UpRight:         "Pointer_UpRight",
	XK_Pointer_DownLeft:        "Pointer_DownLeft",
	XK_Pointer_DownRight:       "Pointer_DownRight",
	XK_Pointer_Button_Dflt:     "Pointer_Button_Dflt",
	XK_Pointer_Button1:         "Pointer_Button1",
	XK_Pointer_Button2:         "Pointer_Button2",
	XK_Pointer_Button3:         "Pointer_Button3",
	XK_Pointer_Button4:         "Pointer_Button4",
	XK_Pointer_Button5:         "Pointer_Button5",
	XK_Pointer_DblClick_Dflt:   "Pointer_DblClick_Dflt",
	XK_Pointer_DblClick1:       "Pointer_DblClick1",
	XK_Pointer_DblClick2:       "Pointer_DblClick2",
	XK_Pointer_DblClick3:       "Pointer_DblClick3",
	XK_Pointer_DblClick4:       "Pointer_DblClick4",
	XK_Pointer_DblClick5:       "Pointer_DblClick5",
	XK_Pointer_Drag_Dflt:       "Pointer_Drag_Dflt",
	XK_Pointer_Drag1:           "Pointer_Drag1",
	XK_Pointer_Drag2:           "Pointer_Drag2",
	XK_Pointer_Drag3:           "Pointer_Drag3",
	XK_Pointer_Drag4:           "Pointer_Drag4",
	XK_Pointer_Drag5:           "Pointer_Drag5",
	XK_Pointer_EnableKeys:      "Pointer_EnableKeys",
	XK_Pointer_Accelerate:      "Pointer_Accelerate",
	XK_Pointer_DfltBtnNext:     "Pointer_DfltBtnNext",
	XK_Pointer_DfltBtnPrev:     "Pointer_DfltBtnPrev",
	XK_ch:                      "ch",
	XK_Ch:                      "Ch",
	XK_CH:                      "CH",
	XK_c_h:                     "c_h",
	XK_C_h:                     "C_h",
	XK_C_H:                     "C_H",
	XK_space:                   "space",
	XK_exclam:                  "exclam",
	XK_quotedbl:                "quotedbl",
	XK_numbersign:              "numbersign",
	XK_dollar:                  "dollar",
	XK_percent:                 "percent",
	XK_ampersand:               "ampersand",
	XK_apostrophe:              "apostrophe",
	// XK_quoteright == XK_apostrophe # deprecated
	XK_parenleft:    "parenleft",
	XK_parenright:   "parenright",
	XK_asterisk:     "asterisk",
	XK_plus:         "plus",
	XK_comma:        "comma",
	XK_minus:        "minus",
	XK_period:       "period",
	XK_slash:        "slash",
	XK_0:            "0",
	XK_1:            "1",
	XK_2:            "2",
	XK_3:            "3",
	XK_4:            "4",
	XK_5:            "5",
	XK_6:            "6",
	XK_7:            "7",
	XK_8:            "8",
	XK_9:            "9",
	XK_colon:        "colon",
	XK_semicolon:    "semicolon",
	XK_less:         "less",
	XK_equal:        "equal",
	XK_greater:      "greater",
	XK_question:     "question",
	XK_at:           "at",
	XK_A:            "A",
	XK_B:            "B",
	XK_C:            "C",
	XK_D:            "D",
	XK_E:            "E",
	XK_F:            "F",
	XK_G:            "G",
	XK_H:            "H",
	XK_I:            "I",
	XK_J:            "J",
	XK_K:            "K",
	XK_L:            "L",
	XK_M:            "M",
	XK_N:            "N",
	XK_O:            "O",
	XK_P:            "P",
	XK_Q:            "Q",
	XK_R:            "R",
	XK_S:            "S",
	XK_T:            "T",
	XK_U:            "U",
	XK_V:            "V",
	XK_W:            "W",
	XK_X:            "X",
	XK_Y:            "Y",
	XK_Z:            "Z",
	XK_bracketleft:  "bracketleft",
	XK_backslash:    "backslash",
	XK_bracketright: "bracketright",
	XK_asciicircum:  "asciicircum",
	XK_underscore:   "underscore",
	XK_grave:        "grave",
	// XK_quoteleft == XK_grave # deprecated
	XK_a:              "a",
	XK_b:              "b",
	XK_c:              "c",
	XK_d:              "d",
	XK_e:              "e",
	XK_f:              "f",
	XK_g:              "g",
	XK_h:              "h",
	XK_i:              "i",
	XK_j:              "j",
	XK_k:              "k",
	XK_l:              "l",
	XK_m:              "m",
	XK_n:              "n",
	XK_o:              "o",
	XK_p:              "p",
	XK_q:              "q",
	XK_r:              "r",
	XK_s:              "s",
	XK_t:              "t",
	XK_u:              "u",
	XK_v:              "v",
	XK_w:              "w",
	XK_x:              "x",
	XK_y:              "y",
	XK_z:              "z",
	XK_braceleft:      "braceleft",
	XK_bar:            "bar",
	XK_braceright:     "braceright",
	XK_asciitilde:     "asciitilde",
	XK_nobreakspace:   "nobreakspace",
	XK_exclamdown:     "exclamdown",
	XK_cent:           "cent",
	XK_sterling:       "sterling",
	XK_currency:       "currency",
	XK_yen:            "yen",
	XK_brokenbar:      "brokenbar",
	XK_section:        "section",
	XK_diaeresis:      "diaeresis",
	XK_copyright:      "copyright",
	XK_ordfeminine:    "ordfeminine",
	XK_guillemotleft:  "guillemotleft",
	XK_notsign:        "notsign",
	XK_hyphen:         "hyphen",
	XK_registered:     "registered",
	XK_macron:         "macron",
	XK_degree:         "degree",
	XK_plusminus:      "plusminus",
	XK_twosuperior:    "twosuperior",
	XK_threesuperior:  "threesuperior",
	XK_acute:          "acute",
	XK_mu:             "mu",
	XK_paragraph:      "paragraph",
	XK_periodcentered: "periodcentered",
	XK_cedilla:        "cedilla",
	XK_onesuperior:    "onesuperior",
	XK_masculine:      "masculine",
	XK_guillemotright: "guillemotright",
	XK_onequarter:     "onequarter",
	XK_onehalf:        "onehalf",
	XK_threequarters:  "threequarters",
	XK_questiondown:   "questiondown",
	XK_Agrave:         "Agrave",
	XK_Aacute:         "Aacute",
	XK_Acircumflex:    "Acircumflex",
	XK_Atilde:         "Atilde",
	XK_Adiaeresis:     "Adiaeresis",
	XK_Aring:          "Aring",
	XK_AE:             "AE",
	XK_Ccedilla:       "Ccedilla",
	XK_Egrave:         "Egrave",
	XK_Eacute:         "Eacute",
	XK_Ecircumflex:    "Ecircumflex",
	XK_Ediaeresis:     "Ediaeresis",
	XK_Igrave:         "Igrave",
	XK_Iacute:         "Iacute",
	XK_Icircumflex:    "Icircumflex",
	XK_Idiaeresis:     "Idiaeresis",
	XK_ETH:            "ETH",
	// XK_Eth == XK_ETH # deprecated
	XK_Ntilde:      "Ntilde",
	XK_Ograve:      "Ograve",
	XK_Oacute:      "Oacute",
	XK_Ocircumflex: "Ocircumflex",
	XK_Otilde:      "Otilde",
	XK_Odiaeresis:  "Odiaeresis",
	XK_multiply:    "multiply",
	XK_Oslash:      "Oslash",
	// XK_Ooblique == XK_Oslash # U+00D8 LATIN CAPITAL LETTER O WITH STROKE
	XK_Ugrave:      "Ugrave",
	XK_Uacute:      "Uacute",
	XK_Ucircumflex: "Ucircumflex",
	XK_Udiaeresis:  "Udiaeresis",
	XK_Yacute:      "Yacute",
	XK_THORN:       "THORN",
	// XK_Thorn == XK_THORN # deprecated
	XK_ssharp:      "ssharp",
	XK_agrave:      "agrave",
	XK_aacute:      "aacute",
	XK_acircumflex: "acircumflex",
	XK_atilde:      "atilde",
	XK_adiaeresis:  "adiaeresis",
	XK_aring:       "aring",
	XK_ae:          "ae",
	XK_ccedilla:    "ccedilla",
	XK_egrave:      "egrave",
	XK_eacute:      "eacute",
	XK_ecircumflex: "ecircumflex",
	XK_ediaeresis:  "ediaeresis",
	XK_igrave:      "igrave",
	XK_iacute:      "iacute",
	XK_icircumflex: "icircumflex",
	XK_idiaeresis:  "idiaeresis",
	XK_eth:         "eth",
	XK_ntilde:      "ntilde",
	XK_ograve:      "ograve",
	XK_oacute:      "oacute",
	XK_ocircumflex: "ocircumflex",
	XK_otilde:      "otilde",
	XK_odiaeresis:  "odiaeresis",
	XK_division:    "division",
	XK_oslash:      "oslash",
	// XK_ooblique == XK_oslash # U+00F8 LATIN SMALL LETTER O WITH STROKE
	XK_ugrave:       "ugrave",
	XK_uacute:       "uacute",
	XK_ucircumflex:  "ucircumflex",
	XK_udiaeresis:   "udiaeresis",
	XK_yacute:       "yacute",
	XK_thorn:        "thorn",
	XK_ydiaeresis:   "ydiaeresis",
	XK_Aogonek:      "Aogonek",
	XK_breve:        "breve",
	XK_Lstroke:      "Lstroke",
	XK_Lcaron:       "Lcaron",
	XK_Sacute:       "Sacute",
	XK_Scaron:       "Scaron",
	XK_Scedilla:     "Scedilla",
	XK_Tcaron:       "Tcaron",
	XK_Zacute:       "Zacute",
	XK_Zcaron:       "Zcaron",
	XK_Zabovedot:    "Zabovedot",
	XK_aogonek:      "aogonek",
	XK_ogonek:       "ogonek",
	XK_lstroke:      "lstroke",
	XK_lcaron:       "lcaron",
	XK_sacute:       "sacute",
	XK_caron:        "caron",
	XK_scaron:       "scaron",
	XK_scedilla:     "scedilla",
	XK_tcaron:       "tcaron",
	XK_zacute:       "zacute",
	XK_doubleacute:  "doubleacute",
	XK_zcaron:       "zcaron",
	XK_zabovedot:    "zabovedot",
	XK_Racute:       "Racute",
	XK_Abreve:       "Abreve",
	XK_Lacute:       "Lacute",
	XK_Cacute:       "Cacute",
	XK_Ccaron:       "Ccaron",
	XK_Eogonek:      "Eogonek",
	XK_Ecaron:       "Ecaron",
	XK_Dcaron:       "Dcaron",
	XK_Dstroke:      "Dstroke",
	XK_Nacute:       "Nacute",
	XK_Ncaron:       "Ncaron",
	XK_Odoubleacute: "Odoubleacute",
	XK_Rcaron:       "Rcaron",
	XK_Uring:        "Uring",
	XK_Udoubleacute: "Udoubleacute",
	XK_Tcedilla:     "Tcedilla",
	XK_racute:       "racute",
	XK_abreve:       "abreve",
	XK_lacute:       "lacute",
	XK_cacute:       "cacute",
	XK_ccaron:       "ccaron",
	XK_eogonek:      "eogonek",
	XK_ecaron:       "ecaron",
	XK_dcaron:       "dcaron",
	XK_dstroke:      "dstroke",
	XK_nacute:       "nacute",
	XK_ncaron:       "ncaron",
	XK_odoubleacute: "odoubleacute",
	XK_rcaron:       "rcaron",
	XK_uring:        "uring",
	XK_udoubleacute: "udoubleacute",
	XK_tcedilla:     "tcedilla",
	XK_abovedot:     "abovedot",
	XK_Hstroke:      "Hstroke",
	XK_Hcircumflex:  "Hcircumflex",
	XK_Iabovedot:    "Iabovedot",
	XK_Gbreve:       "Gbreve",
	XK_Jcircumflex:  "Jcircumflex",
	XK_hstroke:      "hstroke",
	XK_hcircumflex:  "hcircumflex",
	XK_idotless:     "idotless",
	XK_gbreve:       "gbreve",
	XK_jcircumflex:  "jcircumflex",
	XK_Cabovedot:    "Cabovedot",
	XK_Ccircumflex:  "Ccircumflex",
	XK_Gabovedot:    "Gabovedot",
	XK_Gcircumflex:  "Gcircumflex",
	XK_Ubreve:       "Ubreve",
	XK_Scircumflex:  "Scircumflex",
	XK_cabovedot:    "cabovedot",
	XK_ccircumflex:  "ccircumflex",
	XK_gabovedot:    "gabovedot",
	XK_gcircumflex:  "gcircumflex",
	XK_ubreve:       "ubreve",
	XK_scircumflex:  "scircumflex",
	XK_kra:          "kra",
	// XK_kappa == XK_kra # deprecated
	XK_Rcedilla:            "Rcedilla",
	XK_Itilde:              "Itilde",
	XK_Lcedilla:            "Lcedilla",
	XK_Emacron:             "Emacron",
	XK_Gcedilla:            "Gcedilla",
	XK_Tslash:              "Tslash",
	XK_rcedilla:            "rcedilla",
	XK_itilde:              "itilde",
	XK_lcedilla:            "lcedilla",
	XK_emacron:             "emacron",
	XK_gcedilla:            "gcedilla",
	XK_tslash:              "tslash",
	XK_ENG:                 "ENG",
	XK_eng:                 "eng",
	XK_Amacron:             "Amacron",
	XK_Iogonek:             "Iogonek",
	XK_Eabovedot:           "Eabovedot",
	XK_Imacron:             "Imacron",
	XK_Ncedilla:            "Ncedilla",
	XK_Omacron:             "Omacron",
	XK_Kcedilla:            "Kcedilla",
	XK_Uogonek:             "Uogonek",
	XK_Utilde:              "Utilde",
	XK_Umacron:             "Umacron",
	XK_amacron:             "amacron",
	XK_iogonek:             "iogonek",
	XK_eabovedot:           "eabovedot",
	XK_imacron:             "imacron",
	XK_ncedilla:            "ncedilla",
	XK_omacron:             "omacron",
	XK_kcedilla:            "kcedilla",
	XK_uogonek:             "uogonek",
	XK_utilde:              "utilde",
	XK_umacron:             "umacron",
	XK_Wcircumflex:         "Wcircumflex",
	XK_wcircumflex:         "wcircumflex",
	XK_Ycircumflex:         "Ycircumflex",
	XK_ycircumflex:         "ycircumflex",
	XK_Babovedot:           "Babovedot",
	XK_babovedot:           "babovedot",
	XK_Dabovedot:           "Dabovedot",
	XK_dabovedot:           "dabovedot",
	XK_Fabovedot:           "Fabovedot",
	XK_fabovedot:           "fabovedot",
	XK_Mabovedot:           "Mabovedot",
	XK_mabovedot:           "mabovedot",
	XK_Pabovedot:           "Pabovedot",
	XK_pabovedot:           "pabovedot",
	XK_Sabovedot:           "Sabovedot",
	XK_sabovedot:           "sabovedot",
	XK_Tabovedot:           "Tabovedot",
	XK_tabovedot:           "tabovedot",
	XK_Wgrave:              "Wgrave",
	XK_wgrave:              "wgrave",
	XK_Wacute:              "Wacute",
	XK_wacute:              "wacute",
	XK_Wdiaeresis:          "Wdiaeresis",
	XK_wdiaeresis:          "wdiaeresis",
	XK_Ygrave:              "Ygrave",
	XK_ygrave:              "ygrave",
	XK_OE:                  "OE",
	XK_oe:                  "oe",
	XK_Ydiaeresis:          "Ydiaeresis",
	XK_overline:            "overline",
	XK_kana_fullstop:       "kana_fullstop",
	XK_kana_openingbracket: "kana_openingbracket",
	XK_kana_closingbracket: "kana_closingbracket",
	XK_kana_comma:          "kana_comma",
	XK_kana_conjunctive:    "kana_conjunctive",
	// XK_kana_middledot == XK_kana_conjunctive # deprecated
	XK_kana_WO:  "kana_WO",
	XK_kana_a:   "kana_a",
	XK_kana_i:   "kana_i",
	XK_kana_u:   "kana_u",
	XK_kana_e:   "kana_e",
	XK_kana_o:   "kana_o",
	XK_kana_ya:  "kana_ya",
	XK_kana_yu:  "kana_yu",
	XK_kana_yo:  "kana_yo",
	XK_kana_tsu: "kana_tsu",
	// XK_kana_tu == XK_kana_tsu # deprecated
	XK_prolongedsound: "prolongedsound",
	XK_kana_A:         "kana_A",
	XK_kana_I:         "kana_I",
	XK_kana_U:         "kana_U",
	XK_kana_E:         "kana_E",
	XK_kana_O:         "kana_O",
	XK_kana_KA:        "kana_KA",
	XK_kana_KI:        "kana_KI",
	XK_kana_KU:        "kana_KU",
	XK_kana_KE:        "kana_KE",
	XK_kana_KO:        "kana_KO",
	XK_kana_SA:        "kana_SA",
	XK_kana_SHI:       "kana_SHI",
	XK_kana_SU:        "kana_SU",
	XK_kana_SE:        "kana_SE",
	XK_kana_SO:        "kana_SO",
	XK_kana_TA:        "kana_TA",
	XK_kana_CHI:       "kana_CHI",
	// XK_kana_TI == XK_kana_CHI # deprecated
	XK_kana_TSU: "kana_TSU",
	// XK_kana_TU == XK_kana_TSU # deprecated
	XK_kana_TE: "kana_TE",
	XK_kana_TO: "kana_TO",
	XK_kana_NA: "kana_NA",
	XK_kana_NI: "kana_NI",
	XK_kana_NU: "kana_NU",
	XK_kana_NE: "kana_NE",
	XK_kana_NO: "kana_NO",
	XK_kana_HA: "kana_HA",
	XK_kana_HI: "kana_HI",
	XK_kana_FU: "kana_FU",
	// XK_kana_HU == XK_kana_FU # deprecated
	XK_kana_HE:         "kana_HE",
	XK_kana_HO:         "kana_HO",
	XK_kana_MA:         "kana_MA",
	XK_kana_MI:         "kana_MI",
	XK_kana_MU:         "kana_MU",
	XK_kana_ME:         "kana_ME",
	XK_kana_MO:         "kana_MO",
	XK_kana_YA:         "kana_YA",
	XK_kana_YU:         "kana_YU",
	XK_kana_YO:         "kana_YO",
	XK_kana_RA:         "kana_RA",
	XK_kana_RI:         "kana_RI",
	XK_kana_RU:         "kana_RU",
	XK_kana_RE:         "kana_RE",
	XK_kana_RO:         "kana_RO",
	XK_kana_WA:         "kana_WA",
	XK_kana_N:          "kana_N",
	XK_voicedsound:     "voicedsound",
	XK_semivoicedsound: "semivoicedsound",
	// XK_kana_switch == XK_Mode_switch # Alias for mode_switch
	XK_Farsi_0:                 "Farsi_0",
	XK_Farsi_1:                 "Farsi_1",
	XK_Farsi_2:                 "Farsi_2",
	XK_Farsi_3:                 "Farsi_3",
	XK_Farsi_4:                 "Farsi_4",
	XK_Farsi_5:                 "Farsi_5",
	XK_Farsi_6:                 "Farsi_6",
	XK_Farsi_7:                 "Farsi_7",
	XK_Farsi_8:                 "Farsi_8",
	XK_Farsi_9:                 "Farsi_9",
	XK_Arabic_percent:          "Arabic_percent",
	XK_Arabic_superscript_alef: "Arabic_superscript_alef",
	XK_Arabic_tteh:             "Arabic_tteh",
	XK_Arabic_peh:              "Arabic_peh",
	XK_Arabic_tcheh:            "Arabic_tcheh",
	XK_Arabic_ddal:             "Arabic_ddal",
	XK_Arabic_rreh:             "Arabic_rreh",
	XK_Arabic_comma:            "Arabic_comma",
	XK_Arabic_fullstop:         "Arabic_fullstop",
	XK_Arabic_0:                "Arabic_0",
	XK_Arabic_1:                "Arabic_1",
	XK_Arabic_2:                "Arabic_2",
	XK_Arabic_3:                "Arabic_3",
	XK_Arabic_4:                "Arabic_4",
	XK_Arabic_5:                "Arabic_5",
	XK_Arabic_6:                "Arabic_6",
	XK_Arabic_7:                "Arabic_7",
	XK_Arabic_8:                "Arabic_8",
	XK_Arabic_9:                "Arabic_9",
	XK_Arabic_semicolon:        "Arabic_semicolon",
	XK_Arabic_question_mark:    "Arabic_question_mark",
	XK_Arabic_hamza:            "Arabic_hamza",
	XK_Arabic_maddaonalef:      "Arabic_maddaonalef",
	XK_Arabic_hamzaonalef:      "Arabic_hamzaonalef",
	XK_Arabic_hamzaonwaw:       "Arabic_hamzaonwaw",
	XK_Arabic_hamzaunderalef:   "Arabic_hamzaunderalef",
	XK_Arabic_hamzaonyeh:       "Arabic_hamzaonyeh",
	XK_Arabic_alef:             "Arabic_alef",
	XK_Arabic_beh:              "Arabic_beh",
	XK_Arabic_tehmarbuta:       "Arabic_tehmarbuta",
	XK_Arabic_teh:              "Arabic_teh",
	XK_Arabic_theh:             "Arabic_theh",
	XK_Arabic_jeem:             "Arabic_jeem",
	XK_Arabic_hah:              "Arabic_hah",
	XK_Arabic_khah:             "Arabic_khah",
	XK_Arabic_dal:              "Arabic_dal",
	XK_Arabic_thal:             "Arabic_thal",
	XK_Arabic_ra:               "Arabic_ra",
	XK_Arabic_zain:             "Arabic_zain",
	XK_Arabic_seen:             "Arabic_seen",
	XK_Arabic_sheen:            "Arabic_sheen",
	XK_Arabic_sad:              "Arabic_sad",
	XK_Arabic_dad:              "Arabic_dad",
	XK_Arabic_tah:              "Arabic_tah",
	XK_Arabic_zah:              "Arabic_zah",
	XK_Arabic_ain:              "Arabic_ain",
	XK_Arabic_ghain:            "Arabic_ghain",
	XK_Arabic_tatweel:          "Arabic_tatweel",
	XK_Arabic_feh:              "Arabic_feh",
	XK_Arabic_qaf:              "Arabic_qaf",
	XK_Arabic_kaf:              "Arabic_kaf",
	XK_Arabic_lam:              "Arabic_lam",
	XK_Arabic_meem:             "Arabic_meem",
	XK_Arabic_noon:             "Arabic_noon",
	XK_Arabic_ha:               "Arabic_ha",
	// XK_Arabic_heh == XK_Arabic_ha # deprecated
	XK_Arabic_waw:             "Arabic_waw",
	XK_Arabic_alefmaksura:     "Arabic_alefmaksura",
	XK_Arabic_yeh:             "Arabic_yeh",
	XK_Arabic_fathatan:        "Arabic_fathatan",
	XK_Arabic_dammatan:        "Arabic_dammatan",
	XK_Arabic_kasratan:        "Arabic_kasratan",
	XK_Arabic_fatha:           "Arabic_fatha",
	XK_Arabic_damma:           "Arabic_damma",
	XK_Arabic_kasra:           "Arabic_kasra",
	XK_Arabic_shadda:          "Arabic_shadda",
	XK_Arabic_sukun:           "Arabic_sukun",
	XK_Arabic_madda_above:     "Arabic_madda_above",
	XK_Arabic_hamza_above:     "Arabic_hamza_above",
	XK_Arabic_hamza_below:     "Arabic_hamza_below",
	XK_Arabic_jeh:             "Arabic_jeh",
	XK_Arabic_veh:             "Arabic_veh",
	XK_Arabic_keheh:           "Arabic_keheh",
	XK_Arabic_gaf:             "Arabic_gaf",
	XK_Arabic_noon_ghunna:     "Arabic_noon_ghunna",
	XK_Arabic_heh_doachashmee: "Arabic_heh_doachashmee",
	XK_Farsi_yeh:              "Farsi_yeh",
	// XK_Arabic_farsi_yeh == XK_Farsi_yeh # U+06CC ARABIC LETTER FARSI YEH
	XK_Arabic_yeh_baree: "Arabic_yeh_baree",
	XK_Arabic_heh_goal:  "Arabic_heh_goal",
	// XK_Arabic_switch == XK_Mode_switch # Alias for mode_switch
	XK_Cyrillic_GHE_bar:        "Cyrillic_GHE_bar",
	XK_Cyrillic_ghe_bar:        "Cyrillic_ghe_bar",
	XK_Cyrillic_ZHE_descender:  "Cyrillic_ZHE_descender",
	XK_Cyrillic_zhe_descender:  "Cyrillic_zhe_descender",
	XK_Cyrillic_KA_descender:   "Cyrillic_KA_descender",
	XK_Cyrillic_ka_descender:   "Cyrillic_ka_descender",
	XK_Cyrillic_KA_vertstroke:  "Cyrillic_KA_vertstroke",
	XK_Cyrillic_ka_vertstroke:  "Cyrillic_ka_vertstroke",
	XK_Cyrillic_EN_descender:   "Cyrillic_EN_descender",
	XK_Cyrillic_en_descender:   "Cyrillic_en_descender",
	XK_Cyrillic_U_straight:     "Cyrillic_U_straight",
	XK_Cyrillic_u_straight:     "Cyrillic_u_straight",
	XK_Cyrillic_U_straight_bar: "Cyrillic_U_straight_bar",
	XK_Cyrillic_u_straight_bar: "Cyrillic_u_straight_bar",
	XK_Cyrillic_HA_descender:   "Cyrillic_HA_descender",
	XK_Cyrillic_ha_descender:   "Cyrillic_ha_descender",
	XK_Cyrillic_CHE_descender:  "Cyrillic_CHE_descender",
	XK_Cyrillic_che_descender:  "Cyrillic_che_descender",
	XK_Cyrillic_CHE_vertstroke: "Cyrillic_CHE_vertstroke",
	XK_Cyrillic_che_vertstroke: "Cyrillic_che_vertstroke",
	XK_Cyrillic_SHHA:           "Cyrillic_SHHA",
	XK_Cyrillic_shha:           "Cyrillic_shha",
	XK_Cyrillic_SCHWA:          "Cyrillic_SCHWA",
	XK_Cyrillic_schwa:          "Cyrillic_schwa",
	XK_Cyrillic_I_macron:       "Cyrillic_I_macron",
	XK_Cyrillic_i_macron:       "Cyrillic_i_macron",
	XK_Cyrillic_O_bar:          "Cyrillic_O_bar",
	XK_Cyrillic_o_bar:          "Cyrillic_o_bar",
	XK_Cyrillic_U_macron:       "Cyrillic_U_macron",
	XK_Cyrillic_u_macron:       "Cyrillic_u_macron",
	XK_Serbian_dje:             "Serbian_dje",
	XK_Macedonia_gje:           "Macedonia_gje",
	XK_Cyrillic_io:             "Cyrillic_io",
	XK_Ukrainian_ie:            "Ukrainian_ie",
	// XK_Ukranian_je == XK_Ukrainian_ie # deprecated
	XK_Macedonia_dse: "Macedonia_dse",
	XK_Ukrainian_i:   "Ukrainian_i",
	// XK_Ukranian_i == XK_Ukrainian_i # deprecated
	XK_Ukrainian_yi: "Ukrainian_yi",
	// XK_Ukranian_yi == XK_Ukrainian_yi # deprecated
	XK_Cyrillic_je: "Cyrillic_je",
	// XK_Serbian_je == XK_Cyrillic_je # deprecated
	XK_Cyrillic_lje: "Cyrillic_lje",
	// XK_Serbian_lje == XK_Cyrillic_lje # deprecated
	XK_Cyrillic_nje: "Cyrillic_nje",
	// XK_Serbian_nje == XK_Cyrillic_nje # deprecated
	XK_Serbian_tshe:              "Serbian_tshe",
	XK_Macedonia_kje:             "Macedonia_kje",
	XK_Ukrainian_ghe_with_upturn: "Ukrainian_ghe_with_upturn",
	XK_Byelorussian_shortu:       "Byelorussian_shortu",
	XK_Cyrillic_dzhe:             "Cyrillic_dzhe",
	// XK_Serbian_dze == XK_Cyrillic_dzhe # deprecated
	XK_numerosign:    "numerosign",
	XK_Serbian_DJE:   "Serbian_DJE",
	XK_Macedonia_GJE: "Macedonia_GJE",
	XK_Cyrillic_IO:   "Cyrillic_IO",
	XK_Ukrainian_IE:  "Ukrainian_IE",
	// XK_Ukranian_JE == XK_Ukrainian_IE # deprecated
	XK_Macedonia_DSE: "Macedonia_DSE",
	XK_Ukrainian_I:   "Ukrainian_I",
	// XK_Ukranian_I == XK_Ukrainian_I # deprecated
	XK_Ukrainian_YI: "Ukrainian_YI",
	// XK_Ukranian_YI == XK_Ukrainian_YI # deprecated
	XK_Cyrillic_JE: "Cyrillic_JE",
	// XK_Serbian_JE == XK_Cyrillic_JE # deprecated
	XK_Cyrillic_LJE: "Cyrillic_LJE",
	// XK_Serbian_LJE == XK_Cyrillic_LJE # deprecated
	XK_Cyrillic_NJE: "Cyrillic_NJE",
	// XK_Serbian_NJE == XK_Cyrillic_NJE # deprecated
	XK_Serbian_TSHE:              "Serbian_TSHE",
	XK_Macedonia_KJE:             "Macedonia_KJE",
	XK_Ukrainian_GHE_WITH_UPTURN: "Ukrainian_GHE_WITH_UPTURN",
	XK_Byelorussian_SHORTU:       "Byelorussian_SHORTU",
	XK_Cyrillic_DZHE:             "Cyrillic_DZHE",
	// XK_Serbian_DZE == XK_Cyrillic_DZHE # deprecated
	XK_Cyrillic_yu:         "Cyrillic_yu",
	XK_Cyrillic_a:          "Cyrillic_a",
	XK_Cyrillic_be:         "Cyrillic_be",
	XK_Cyrillic_tse:        "Cyrillic_tse",
	XK_Cyrillic_de:         "Cyrillic_de",
	XK_Cyrillic_ie:         "Cyrillic_ie",
	XK_Cyrillic_ef:         "Cyrillic_ef",
	XK_Cyrillic_ghe:        "Cyrillic_ghe",
	XK_Cyrillic_ha:         "Cyrillic_ha",
	XK_Cyrillic_i:          "Cyrillic_i",
	XK_Cyrillic_shorti:     "Cyrillic_shorti",
	XK_Cyrillic_ka:         "Cyrillic_ka",
	XK_Cyrillic_el:         "Cyrillic_el",
	XK_Cyrillic_em:         "Cyrillic_em",
	XK_Cyrillic_en:         "Cyrillic_en",
	XK_Cyrillic_o:          "Cyrillic_o",
	XK_Cyrillic_pe:         "Cyrillic_pe",
	XK_Cyrillic_ya:         "Cyrillic_ya",
	XK_Cyrillic_er:         "Cyrillic_er",
	XK_Cyrillic_es:         "Cyrillic_es",
	XK_Cyrillic_te:         "Cyrillic_te",
	XK_Cyrillic_u:          "Cyrillic_u",
	XK_Cyrillic_zhe:        "Cyrillic_zhe",
	XK_Cyrillic_ve:         "Cyrillic_ve",
	XK_Cyrillic_softsign:   "Cyrillic_softsign",
	XK_Cyrillic_yeru:       "Cyrillic_yeru",
	XK_Cyrillic_ze:         "Cyrillic_ze",
	XK_Cyrillic_sha:        "Cyrillic_sha",
	XK_Cyrillic_e:          "Cyrillic_e",
	XK_Cyrillic_shcha:      "Cyrillic_shcha",
	XK_Cyrillic_che:        "Cyrillic_che",
	XK_Cyrillic_hardsign:   "Cyrillic_hardsign",
	XK_Cyrillic_YU:         "Cyrillic_YU",
	XK_Cyrillic_A:          "Cyrillic_A",
	XK_Cyrillic_BE:         "Cyrillic_BE",
	XK_Cyrillic_TSE:        "Cyrillic_TSE",
	XK_Cyrillic_DE:         "Cyrillic_DE",
	XK_Cyrillic_IE:         "Cyrillic_IE",
	XK_Cyrillic_EF:         "Cyrillic_EF",
	XK_Cyrillic_GHE:        "Cyrillic_GHE",
	XK_Cyrillic_HA:         "Cyrillic_HA",
	XK_Cyrillic_I:          "Cyrillic_I",
	XK_Cyrillic_SHORTI:     "Cyrillic_SHORTI",
	XK_Cyrillic_KA:         "Cyrillic_KA",
	XK_Cyrillic_EL:         "Cyrillic_EL",
	XK_Cyrillic_EM:         "Cyrillic_EM",
	XK_Cyrillic_EN:         "Cyrillic_EN",
	XK_Cyrillic_O:          "Cyrillic_O",
	XK_Cyrillic_PE:         "Cyrillic_PE",
	XK_Cyrillic_YA:         "Cyrillic_YA",
	XK_Cyrillic_ER:         "Cyrillic_ER",
	XK_Cyrillic_ES:         "Cyrillic_ES",
	XK_Cyrillic_TE:         "Cyrillic_TE",
	XK_Cyrillic_U:          "Cyrillic_U",
	XK_Cyrillic_ZHE:        "Cyrillic_ZHE",
	XK_Cyrillic_VE:         "Cyrillic_VE",
	XK_Cyrillic_SOFTSIGN:   "Cyrillic_SOFTSIGN",
	XK_Cyrillic_YERU:       "Cyrillic_YERU",
	XK_Cyrillic_ZE:         "Cyrillic_ZE",
	XK_Cyrillic_SHA:        "Cyrillic_SHA",
	XK_Cyrillic_E:          "Cyrillic_E",
	XK_Cyrillic_SHCHA:      "Cyrillic_SHCHA",
	XK_Cyrillic_CHE:        "Cyrillic_CHE",
	XK_Cyrillic_HARDSIGN:   "Cyrillic_HARDSIGN",
	XK_Greek_ALPHAaccent:   "Greek_ALPHAaccent",
	XK_Greek_EPSILONaccent: "Greek_EPSILONaccent",
	XK_Greek_ETAaccent:     "Greek_ETAaccent",
	XK_Greek_IOTAaccent:    "Greek_IOTAaccent",
	XK_Greek_IOTAdieresis:  "Greek_IOTAdieresis",
	// XK_Greek_IOTAdiaeresis == XK_Greek_IOTAdieresis # old typo
	XK_Greek_OMICRONaccent:         "Greek_OMICRONaccent",
	XK_Greek_UPSILONaccent:         "Greek_UPSILONaccent",
	XK_Greek_UPSILONdieresis:       "Greek_UPSILONdieresis",
	XK_Greek_OMEGAaccent:           "Greek_OMEGAaccent",
	XK_Greek_accentdieresis:        "Greek_accentdieresis",
	XK_Greek_horizbar:              "Greek_horizbar",
	XK_Greek_alphaaccent:           "Greek_alphaaccent",
	XK_Greek_epsilonaccent:         "Greek_epsilonaccent",
	XK_Greek_etaaccent:             "Greek_etaaccent",
	XK_Greek_iotaaccent:            "Greek_iotaaccent",
	XK_Greek_iotadieresis:          "Greek_iotadieresis",
	XK_Greek_iotaaccentdieresis:    "Greek_iotaaccentdieresis",
	XK_Greek_omicronaccent:         "Greek_omicronaccent",
	XK_Greek_upsilonaccent:         "Greek_upsilonaccent",
	XK_Greek_upsilondieresis:       "Greek_upsilondieresis",
	XK_Greek_upsilonaccentdieresis: "Greek_upsilonaccentdieresis",
	XK_Greek_omegaaccent:           "Greek_omegaaccent",
	XK_Greek_ALPHA:                 "Greek_ALPHA",
	XK_Greek_BETA:                  "Greek_BETA",
	XK_Greek_GAMMA:                 "Greek_GAMMA",
	XK_Greek_DELTA:                 "Greek_DELTA",
	XK_Greek_EPSILON:               "Greek_EPSILON",
	XK_Greek_ZETA:                  "Greek_ZETA",
	XK_Greek_ETA:                   "Greek_ETA",
	XK_Greek_THETA:                 "Greek_THETA",
	XK_Greek_IOTA:                  "Greek_IOTA",
	XK_Greek_KAPPA:                 "Greek_KAPPA",
	XK_Greek_LAMDA:                 "Greek_LAMDA",
	// XK_Greek_LAMBDA == XK_Greek_LAMDA # U+039B GREEK CAPITAL LETTER LAMDA
	XK_Greek_MU:      "Greek_MU",
	XK_Greek_NU:      "Greek_NU",
	XK_Greek_XI:      "Greek_XI",
	XK_Greek_OMICRON: "Greek_OMICRON",
	XK_Greek_PI:      "Greek_PI",
	XK_Greek_RHO:     "Greek_RHO",
	XK_Greek_SIGMA:   "Greek_SIGMA",
	XK_Greek_TAU:     "Greek_TAU",
	XK_Greek_UPSILON: "Greek_UPSILON",
	XK_Greek_PHI:     "Greek_PHI",
	XK_Greek_CHI:     "Greek_CHI",
	XK_Greek_PSI:     "Greek_PSI",
	XK_Greek_OMEGA:   "Greek_OMEGA",
	XK_Greek_alpha:   "Greek_alpha",
	XK_Greek_beta:    "Greek_beta",
	XK_Greek_gamma:   "Greek_gamma",
	XK_Greek_delta:   "Greek_delta",
	XK_Greek_epsilon: "Greek_epsilon",
	XK_Greek_zeta:    "Greek_zeta",
	XK_Greek_eta:     "Greek_eta",
	XK_Greek_theta:   "Greek_theta",
	XK_Greek_iota:    "Greek_iota",
	XK_Greek_kappa:   "Greek_kappa",
	XK_Greek_lamda:   "Greek_lamda",
	// XK_Greek_lambda == XK_Greek_lamda # U+03BB GREEK SMALL LETTER LAMDA
	XK_Greek_mu:              "Greek_mu",
	XK_Greek_nu:              "Greek_nu",
	XK_Greek_xi:              "Greek_xi",
	XK_Greek_omicron:         "Greek_omicron",
	XK_Greek_pi:              "Greek_pi",
	XK_Greek_rho:             "Greek_rho",
	XK_Greek_sigma:           "Greek_sigma",
	XK_Greek_finalsmallsigma: "Greek_finalsmallsigma",
	XK_Greek_tau:             "Greek_tau",
	XK_Greek_upsilon:         "Greek_upsilon",
	XK_Greek_phi:             "Greek_phi",
	XK_Greek_chi:             "Greek_chi",
	XK_Greek_psi:             "Greek_psi",
	XK_Greek_omega:           "Greek_omega",
	// XK_Greek_switch == XK_Mode_switch # Alias for mode_switch
	XK_hebrew_doublelowline: "hebrew_doublelowline",
	XK_hebrew_aleph:         "hebrew_aleph",
	XK_hebrew_bet:           "hebrew_bet",
	// XK_hebrew_beth == XK_hebrew_bet # deprecated
	XK_hebrew_gimel: "hebrew_gimel",
	// XK_hebrew_gimmel == XK_hebrew_gimel # deprecated
	XK_hebrew_dalet: "hebrew_dalet",
	// XK_hebrew_daleth == XK_hebrew_dalet # deprecated
	XK_hebrew_he:   "hebrew_he",
	XK_hebrew_waw:  "hebrew_waw",
	XK_hebrew_zain: "hebrew_zain",
	// XK_hebrew_zayin == XK_hebrew_zain # deprecated
	XK_hebrew_chet: "hebrew_chet",
	// XK_hebrew_het == XK_hebrew_chet # deprecated
	XK_hebrew_tet: "hebrew_tet",
	// XK_hebrew_teth == XK_hebrew_tet # deprecated
	XK_hebrew_yod:       "hebrew_yod",
	XK_hebrew_finalkaph: "hebrew_finalkaph",
	XK_hebrew_kaph:      "hebrew_kaph",
	XK_hebrew_lamed:     "hebrew_lamed",
	XK_hebrew_finalmem:  "hebrew_finalmem",
	XK_hebrew_mem:       "hebrew_mem",
	XK_hebrew_finalnun:  "hebrew_finalnun",
	XK_hebrew_nun:       "hebrew_nun",
	XK_hebrew_samech:    "hebrew_samech",
	// XK_hebrew_samekh == XK_hebrew_samech # deprecated
	XK_hebrew_ayin:      "hebrew_ayin",
	XK_hebrew_finalpe:   "hebrew_finalpe",
	XK_hebrew_pe:        "hebrew_pe",
	XK_hebrew_finalzade: "hebrew_finalzade",
	// XK_hebrew_finalzadi == XK_hebrew_finalzade # deprecated
	XK_hebrew_zade: "hebrew_zade",
	// XK_hebrew_zadi == XK_hebrew_zade # deprecated
	XK_hebrew_qoph: "hebrew_qoph",
	// XK_hebrew_kuf == XK_hebrew_qoph # deprecated
	XK_hebrew_resh: "hebrew_resh",
	XK_hebrew_shin: "hebrew_shin",
	XK_hebrew_taw:  "hebrew_taw",
	// XK_hebrew_taf == XK_hebrew_taw # deprecated
	// XK_Hebrew_switch == XK_Mode_switch # Alias for mode_switch
	XK_Thai_kokai:             "Thai_kokai",
	XK_Thai_khokhai:           "Thai_khokhai",
	XK_Thai_khokhuat:          "Thai_khokhuat",
	XK_Thai_khokhwai:          "Thai_khokhwai",
	XK_Thai_khokhon:           "Thai_khokhon",
	XK_Thai_khorakhang:        "Thai_khorakhang",
	XK_Thai_ngongu:            "Thai_ngongu",
	XK_Thai_chochan:           "Thai_chochan",
	XK_Thai_choching:          "Thai_choching",
	XK_Thai_chochang:          "Thai_chochang",
	XK_Thai_soso:              "Thai_soso",
	XK_Thai_chochoe:           "Thai_chochoe",
	XK_Thai_yoying:            "Thai_yoying",
	XK_Thai_dochada:           "Thai_dochada",
	XK_Thai_topatak:           "Thai_topatak",
	XK_Thai_thothan:           "Thai_thothan",
	XK_Thai_thonangmontho:     "Thai_thonangmontho",
	XK_Thai_thophuthao:        "Thai_thophuthao",
	XK_Thai_nonen:             "Thai_nonen",
	XK_Thai_dodek:             "Thai_dodek",
	XK_Thai_totao:             "Thai_totao",
	XK_Thai_thothung:          "Thai_thothung",
	XK_Thai_thothahan:         "Thai_thothahan",
	XK_Thai_thothong:          "Thai_thothong",
	XK_Thai_nonu:              "Thai_nonu",
	XK_Thai_bobaimai:          "Thai_bobaimai",
	XK_Thai_popla:             "Thai_popla",
	XK_Thai_phophung:          "Thai_phophung",
	XK_Thai_fofa:              "Thai_fofa",
	XK_Thai_phophan:           "Thai_phophan",
	XK_Thai_fofan:             "Thai_fofan",
	XK_Thai_phosamphao:        "Thai_phosamphao",
	XK_Thai_moma:              "Thai_moma",
	XK_Thai_yoyak:             "Thai_yoyak",
	XK_Thai_rorua:             "Thai_rorua",
	XK_Thai_ru:                "Thai_ru",
	XK_Thai_loling:            "Thai_loling",
	XK_Thai_lu:                "Thai_lu",
	XK_Thai_wowaen:            "Thai_wowaen",
	XK_Thai_sosala:            "Thai_sosala",
	XK_Thai_sorusi:            "Thai_sorusi",
	XK_Thai_sosua:             "Thai_sosua",
	XK_Thai_hohip:             "Thai_hohip",
	XK_Thai_lochula:           "Thai_lochula",
	XK_Thai_oang:              "Thai_oang",
	XK_Thai_honokhuk:          "Thai_honokhuk",
	XK_Thai_paiyannoi:         "Thai_paiyannoi",
	XK_Thai_saraa:             "Thai_saraa",
	XK_Thai_maihanakat:        "Thai_maihanakat",
	XK_Thai_saraaa:            "Thai_saraaa",
	XK_Thai_saraam:            "Thai_saraam",
	XK_Thai_sarai:             "Thai_sarai",
	XK_Thai_saraii:            "Thai_saraii",
	XK_Thai_saraue:            "Thai_saraue",
	XK_Thai_sarauee:           "Thai_sarauee",
	XK_Thai_sarau:             "Thai_sarau",
	XK_Thai_sarauu:            "Thai_sarauu",
	XK_Thai_phinthu:           "Thai_phinthu",
	XK_Thai_maihanakat_maitho: "Thai_maihanakat_maitho",
	XK_Thai_baht:              "Thai_baht",
	XK_Thai_sarae:             "Thai_sarae",
	XK_Thai_saraae:            "Thai_saraae",
	XK_Thai_sarao:             "Thai_sarao",
	XK_Thai_saraaimaimuan:     "Thai_saraaimaimuan",
	XK_Thai_saraaimaimalai:    "Thai_saraaimaimalai",
	XK_Thai_lakkhangyao:       "Thai_lakkhangyao",
	XK_Thai_maiyamok:          "Thai_maiyamok",
	XK_Thai_maitaikhu:         "Thai_maitaikhu",
	XK_Thai_maiek:             "Thai_maiek",
	XK_Thai_maitho:            "Thai_maitho",
	XK_Thai_maitri:            "Thai_maitri",
	XK_Thai_maichattawa:       "Thai_maichattawa",
	XK_Thai_thanthakhat:       "Thai_thanthakhat",
	XK_Thai_nikhahit:          "Thai_nikhahit",
	XK_Thai_leksun:            "Thai_leksun",
	XK_Thai_leknung:           "Thai_leknung",
	XK_Thai_leksong:           "Thai_leksong",
	XK_Thai_leksam:            "Thai_leksam",
	XK_Thai_leksi:             "Thai_leksi",
	XK_Thai_lekha:             "Thai_lekha",
	XK_Thai_lekhok:            "Thai_lekhok",
	XK_Thai_lekchet:           "Thai_lekchet",
	XK_Thai_lekpaet:           "Thai_lekpaet",
	XK_Thai_lekkao:            "Thai_lekkao",
	XK_Hangul:                 "Hangul",
	XK_Hangul_Start:           "Hangul_Start",
	XK_Hangul_End:             "Hangul_End",
	XK_Hangul_Hanja:           "Hangul_Hanja",
	XK_Hangul_Jamo:            "Hangul_Jamo",
	XK_Hangul_Romaja:          "Hangul_Romaja",
	// XK_Hangul_Codeinput == XK_Codeinput # Hangul code input mode
	XK_Hangul_Jeonja:    "Hangul_Jeonja",
	XK_Hangul_Banja:     "Hangul_Banja",
	XK_Hangul_PreHanja:  "Hangul_PreHanja",
	XK_Hangul_PostHanja: "Hangul_PostHanja",
	// XK_Hangul_SingleCandidate == XK_SingleCandidate # Single candidate
	// XK_Hangul_MultipleCandidate == XK_MultipleCandidate # Multiple candidate
	// XK_Hangul_PreviousCandidate == XK_PreviousCandidate # Previous candidate
	XK_Hangul_Special: "Hangul_Special",
	// XK_Hangul_switch == XK_Mode_switch # Alias for mode_switch
	XK_Hangul_Kiyeog:              "Hangul_Kiyeog",
	XK_Hangul_SsangKiyeog:         "Hangul_SsangKiyeog",
	XK_Hangul_KiyeogSios:          "Hangul_KiyeogSios",
	XK_Hangul_Nieun:               "Hangul_Nieun",
	XK_Hangul_NieunJieuj:          "Hangul_NieunJieuj",
	XK_Hangul_NieunHieuh:          "Hangul_NieunHieuh",
	XK_Hangul_Dikeud:              "Hangul_Dikeud",
	XK_Hangul_SsangDikeud:         "Hangul_SsangDikeud",
	XK_Hangul_Rieul:               "Hangul_Rieul",
	XK_Hangul_RieulKiyeog:         "Hangul_RieulKiyeog",
	XK_Hangul_RieulMieum:          "Hangul_RieulMieum",
	XK_Hangul_RieulPieub:          "Hangul_RieulPieub",
	XK_Hangul_RieulSios:           "Hangul_RieulSios",
	XK_Hangul_RieulTieut:          "Hangul_RieulTieut",
	XK_Hangul_RieulPhieuf:         "Hangul_RieulPhieuf",
	XK_Hangul_RieulHieuh:          "Hangul_RieulHieuh",
	XK_Hangul_Mieum:               "Hangul_Mieum",
	XK_Hangul_Pieub:               "Hangul_Pieub",
	XK_Hangul_SsangPieub:          "Hangul_SsangPieub",
	XK_Hangul_PieubSios:           "Hangul_PieubSios",
	XK_Hangul_Sios:                "Hangul_Sios",
	XK_Hangul_SsangSios:           "Hangul_SsangSios",
	XK_Hangul_Ieung:               "Hangul_Ieung",
	XK_Hangul_Jieuj:               "Hangul_Jieuj",
	XK_Hangul_SsangJieuj:          "Hangul_SsangJieuj",
	XK_Hangul_Cieuc:               "Hangul_Cieuc",
	XK_Hangul_Khieuq:              "Hangul_Khieuq",
	XK_Hangul_Tieut:               "Hangul_Tieut",
	XK_Hangul_Phieuf:              "Hangul_Phieuf",
	XK_Hangul_Hieuh:               "Hangul_Hieuh",
	XK_Hangul_A:                   "Hangul_A",
	XK_Hangul_AE:                  "Hangul_AE",
	XK_Hangul_YA:                  "Hangul_YA",
	XK_Hangul_YAE:                 "Hangul_YAE",
	XK_Hangul_EO:                  "Hangul_EO",
	XK_Hangul_E:                   "Hangul_E",
	XK_Hangul_YEO:                 "Hangul_YEO",
	XK_Hangul_YE:                  "Hangul_YE",
	XK_Hangul_O:                   "Hangul_O",
	XK_Hangul_WA:                  "Hangul_WA",
	XK_Hangul_WAE:                 "Hangul_WAE",
	XK_Hangul_OE:                  "Hangul_OE",
	XK_Hangul_YO:                  "Hangul_YO",
	XK_Hangul_U:                   "Hangul_U",
	XK_Hangul_WEO:                 "Hangul_WEO",
	XK_Hangul_WE:                  "Hangul_WE",
	XK_Hangul_WI:                  "Hangul_WI",
	XK_Hangul_YU:                  "Hangul_YU",
	XK_Hangul_EU:                  "Hangul_EU",
	XK_Hangul_YI:                  "Hangul_YI",
	XK_Hangul_I:                   "Hangul_I",
	XK_Hangul_J_Kiyeog:            "Hangul_J_Kiyeog",
	XK_Hangul_J_SsangKiyeog:       "Hangul_J_SsangKiyeog",
	XK_Hangul_J_KiyeogSios:        "Hangul_J_KiyeogSios",
	XK_Hangul_J_Nieun:             "Hangul_J_Nieun",
	XK_Hangul_J_NieunJieuj:        "Hangul_J_NieunJieuj",
	XK_Hangul_J_NieunHieuh:        "Hangul_J_NieunHieuh",
	XK_Hangul_J_Dikeud:            "Hangul_J_Dikeud",
	XK_Hangul_J_Rieul:             "Hangul_J_Rieul",
	XK_Hangul_J_RieulKiyeog:       "Hangul_J_RieulKiyeog",
	XK_Hangul_J_RieulMieum:        "Hangul_J_RieulMieum",
	XK_Hangul_J_RieulPieub:        "Hangul_J_RieulPieub",
	XK_Hangul_J_RieulSios:         "Hangul_J_RieulSios",
	XK_Hangul_J_RieulTieut:        "Hangul_J_RieulTieut",
	XK_Hangul_J_RieulPhieuf:       "Hangul_J_RieulPhieuf",
	XK_Hangul_J_RieulHieuh:        "Hangul_J_RieulHieuh",
	XK_Hangul_J_Mieum:             "Hangul_J_Mieum",
	XK_Hangul_J_Pieub:             "Hangul_J_Pieub",
	XK_Hangul_J_PieubSios:         "Hangul_J_PieubSios",
	XK_Hangul_J_Sios:              "Hangul_J_Sios",
	XK_Hangul_J_SsangSios:         "Hangul_J_SsangSios",
	XK_Hangul_J_Ieung:             "Hangul_J_Ieung",
	XK_Hangul_J_Jieuj:             "Hangul_J_Jieuj",
	XK_Hangul_J_Cieuc:             "Hangul_J_Cieuc",
	XK_Hangul_J_Khieuq:            "Hangul_J_Khieuq",
	XK_Hangul_J_Tieut:             "Hangul_J_Tieut",
	XK_Hangul_J_Phieuf:            "Hangul_J_Phieuf",
	XK_Hangul_J_Hieuh:             "Hangul_J_Hieuh",
	XK_Hangul_RieulYeorinHieuh:    "Hangul_RieulYeorinHieuh",
	XK_Hangul_SunkyeongeumMieum:   "Hangul_SunkyeongeumMieum",
	XK_Hangul_SunkyeongeumPieub:   "Hangul_SunkyeongeumPieub",
	XK_Hangul_PanSios:             "Hangul_PanSios",
	XK_Hangul_KkogjiDalrinIeung:   "Hangul_KkogjiDalrinIeung",
	XK_Hangul_SunkyeongeumPhieuf:  "Hangul_SunkyeongeumPhieuf",
	XK_Hangul_YeorinHieuh:         "Hangul_YeorinHieuh",
	XK_Hangul_AraeA:               "Hangul_AraeA",
	XK_Hangul_AraeAE:              "Hangul_AraeAE",
	XK_Hangul_J_PanSios:           "Hangul_J_PanSios",
	XK_Hangul_J_KkogjiDalrinIeung: "Hangul_J_KkogjiDalrinIeung",
	XK_Hangul_J_YeorinHieuh:       "Hangul_J_YeorinHieuh",
	XK_Korean_Won:                 "Korean_Won",
	XK_Armenian_ligature_ew:       "Armenian_ligature_ew",
	XK_Armenian_full_stop:         "Armenian_full_stop",
	// XK_Armenian_verjaket == XK_Armenian_full_stop # U+0589 ARMENIAN FULL STOP
	XK_Armenian_separation_mark: "Armenian_separation_mark",
	// XK_Armenian_but == XK_Armenian_separation_mark # U+055D ARMENIAN COMMA
	XK_Armenian_hyphen: "Armenian_hyphen",
	// XK_Armenian_yentamna == XK_Armenian_hyphen # U+058A ARMENIAN HYPHEN
	XK_Armenian_exclam: "Armenian_exclam",
	// XK_Armenian_amanak == XK_Armenian_exclam # U+055C ARMENIAN EXCLAMATION MARK
	XK_Armenian_accent: "Armenian_accent",
	// XK_Armenian_shesht == XK_Armenian_accent # U+055B ARMENIAN EMPHASIS MARK
	XK_Armenian_question: "Armenian_question",
	// XK_Armenian_paruyk == XK_Armenian_question # U+055E ARMENIAN QUESTION MARK
	XK_Armenian_AYB:        "Armenian_AYB",
	XK_Armenian_ayb:        "Armenian_ayb",
	XK_Armenian_BEN:        "Armenian_BEN",
	XK_Armenian_ben:        "Armenian_ben",
	XK_Armenian_GIM:        "Armenian_GIM",
	XK_Armenian_gim:        "Armenian_gim",
	XK_Armenian_DA:         "Armenian_DA",
	XK_Armenian_da:         "Armenian_da",
	XK_Armenian_YECH:       "Armenian_YECH",
	XK_Armenian_yech:       "Armenian_yech",
	XK_Armenian_ZA:         "Armenian_ZA",
	XK_Armenian_za:         "Armenian_za",
	XK_Armenian_E:          "Armenian_E",
	XK_Armenian_e:          "Armenian_e",
	XK_Armenian_AT:         "Armenian_AT",
	XK_Armenian_at:         "Armenian_at",
	XK_Armenian_TO:         "Armenian_TO",
	XK_Armenian_to:         "Armenian_to",
	XK_Armenian_ZHE:        "Armenian_ZHE",
	XK_Armenian_zhe:        "Armenian_zhe",
	XK_Armenian_INI:        "Armenian_INI",
	XK_Armenian_ini:        "Armenian_ini",
	XK_Armenian_LYUN:       "Armenian_LYUN",
	XK_Armenian_lyun:       "Armenian_lyun",
	XK_Armenian_KHE:        "Armenian_KHE",
	XK_Armenian_khe:        "Armenian_khe",
	XK_Armenian_TSA:        "Armenian_TSA",
	XK_Armenian_tsa:        "Armenian_tsa",
	XK_Armenian_KEN:        "Armenian_KEN",
	XK_Armenian_ken:        "Armenian_ken",
	XK_Armenian_HO:         "Armenian_HO",
	XK_Armenian_ho:         "Armenian_ho",
	XK_Armenian_DZA:        "Armenian_DZA",
	XK_Armenian_dza:        "Armenian_dza",
	XK_Armenian_GHAT:       "Armenian_GHAT",
	XK_Armenian_ghat:       "Armenian_ghat",
	XK_Armenian_TCHE:       "Armenian_TCHE",
	XK_Armenian_tche:       "Armenian_tche",
	XK_Armenian_MEN:        "Armenian_MEN",
	XK_Armenian_men:        "Armenian_men",
	XK_Armenian_HI:         "Armenian_HI",
	XK_Armenian_hi:         "Armenian_hi",
	XK_Armenian_NU:         "Armenian_NU",
	XK_Armenian_nu:         "Armenian_nu",
	XK_Armenian_SHA:        "Armenian_SHA",
	XK_Armenian_sha:        "Armenian_sha",
	XK_Armenian_VO:         "Armenian_VO",
	XK_Armenian_vo:         "Armenian_vo",
	XK_Armenian_CHA:        "Armenian_CHA",
	XK_Armenian_cha:        "Armenian_cha",
	XK_Armenian_PE:         "Armenian_PE",
	XK_Armenian_pe:         "Armenian_pe",
	XK_Armenian_JE:         "Armenian_JE",
	XK_Armenian_je:         "Armenian_je",
	XK_Armenian_RA:         "Armenian_RA",
	XK_Armenian_ra:         "Armenian_ra",
	XK_Armenian_SE:         "Armenian_SE",
	XK_Armenian_se:         "Armenian_se",
	XK_Armenian_VEV:        "Armenian_VEV",
	XK_Armenian_vev:        "Armenian_vev",
	XK_Armenian_TYUN:       "Armenian_TYUN",
	XK_Armenian_tyun:       "Armenian_tyun",
	XK_Armenian_RE:         "Armenian_RE",
	XK_Armenian_re:         "Armenian_re",
	XK_Armenian_TSO:        "Armenian_TSO",
	XK_Armenian_tso:        "Armenian_tso",
	XK_Armenian_VYUN:       "Armenian_VYUN",
	XK_Armenian_vyun:       "Armenian_vyun",
	XK_Armenian_PYUR:       "Armenian_PYUR",
	XK_Armenian_pyur:       "Armenian_pyur",
	XK_Armenian_KE:         "Armenian_KE",
	XK_Armenian_ke:         "Armenian_ke",
	XK_Armenian_O:          "Armenian_O",
	XK_Armenian_o:          "Armenian_o",
	XK_Armenian_FE:         "Armenian_FE",
	XK_Armenian_fe:         "Armenian_fe",
	XK_Armenian_apostrophe: "Armenian_apostrophe",
	XK_Georgian_an:         "Georgian_an",
	XK_Georgian_ban:        "Georgian_ban",
	XK_Georgian_gan:        "Georgian_gan",
	XK_Georgian_don:        "Georgian_don",
	XK_Georgian_en:         "Georgian_en",
	XK_Georgian_vin:        "Georgian_vin",
	XK_Georgian_zen:        "Georgian_zen",
	XK_Georgian_tan:        "Georgian_tan",
	XK_Georgian_in:         "Georgian_in",
	XK_Georgian_kan:        "Georgian_kan",
	XK_Georgian_las:        "Georgian_las",
	XK_Georgian_man:        "Georgian_man",
	XK_Georgian_nar:        "Georgian_nar",
	XK_Georgian_on:         "Georgian_on",
	XK_Georgian_par:        "Georgian_par",
	XK_Georgian_zhar:       "Georgian_zhar",
	XK_Georgian_rae:        "Georgian_rae",
	XK_Georgian_san:        "Georgian_san",
	XK_Georgian_tar:        "Georgian_tar",
	XK_Georgian_un:         "Georgian_un",
	XK_Georgian_phar:       "Georgian_phar",
	XK_Georgian_khar:       "Georgian_khar",
	XK_Georgian_ghan:       "Georgian_ghan",
	XK_Georgian_qar:        "Georgian_qar",
	XK_Georgian_shin:       "Georgian_shin",
	XK_Georgian_chin:       "Georgian_chin",
	XK_Georgian_can:        "Georgian_can",
	XK_Georgian_jil:        "Georgian_jil",
	XK_Georgian_cil:        "Georgian_cil",
	XK_Georgian_char:       "Georgian_char",
	XK_Georgian_xan:        "Georgian_xan",
	XK_Georgian_jhan:       "Georgian_jhan",
	XK_Georgian_hae:        "Georgian_hae",
	XK_Georgian_he:         "Georgian_he",
	XK_Georgian_hie:        "Georgian_hie",
	XK_Georgian_we:         "Georgian_we",
	XK_Georgian_har:        "Georgian_har",
	XK_Georgian_hoe:        "Georgian_hoe",
	XK_Georgian_fi:         "Georgian_fi",
	XK_Xabovedot:           "Xabovedot",
	XK_Ibreve:              "Ibreve",
	XK_Zstroke:             "Zstroke",
	XK_Gcaron:              "Gcaron",
	XK_Ocaron:              "Ocaron",
	XK_Obarred:             "Obarred",
	XK_xabovedot:           "xabovedot",
	XK_ibreve:              "ibreve",
	XK_zstroke:             "zstroke",
	XK_gcaron:              "gcaron",
	XK_ocaron:              "ocaron",
	XK_obarred:             "obarred",
	XK_SCHWA:               "SCHWA",
	XK_schwa:               "schwa",
	XK_EZH:                 "EZH",
	XK_ezh:                 "ezh",
	XK_Lbelowdot:           "Lbelowdot",
	XK_lbelowdot:           "lbelowdot",
	XK_Abelowdot:           "Abelowdot",
	XK_abelowdot:           "abelowdot",
	XK_Ahook:               "Ahook",
	XK_ahook:               "ahook",
	XK_Acircumflexacute:    "Acircumflexacute",
	XK_acircumflexacute:    "acircumflexacute",
	XK_Acircumflexgrave:    "Acircumflexgrave",
	XK_acircumflexgrave:    "acircumflexgrave",
	XK_Acircumflexhook:     "Acircumflexhook",
	XK_acircumflexhook:     "acircumflexhook",
	XK_Acircumflextilde:    "Acircumflextilde",
	XK_acircumflextilde:    "acircumflextilde",
	XK_Acircumflexbelowdot: "Acircumflexbelowdot",
	XK_acircumflexbelowdot: "acircumflexbelowdot",
	XK_Abreveacute:         "Abreveacute",
	XK_abreveacute:         "abreveacute",
	XK_Abrevegrave:         "Abrevegrave",
	XK_abrevegrave:         "abrevegrave",
	XK_Abrevehook:          "Abrevehook",
	XK_abrevehook:          "abrevehook",
	XK_Abrevetilde:         "Abrevetilde",
	XK_abrevetilde:         "abrevetilde",
	XK_Abrevebelowdot:      "Abrevebelowdot",
	XK_abrevebelowdot:      "abrevebelowdot",
	XK_Ebelowdot:           "Ebelowdot",
	XK_ebelowdot:           "ebelowdot",
	XK_Ehook:               "Ehook",
	XK_ehook:               "ehook",
	XK_Etilde:              "Etilde",
	XK_etilde:              "etilde",
	XK_Ecircumflexacute:    "Ecircumflexacute",
	XK_ecircumflexacute:    "ecircumflexacute",
	XK_Ecircumflexgrave:    "Ecircumflexgrave",
	XK_ecircumflexgrave:    "ecircumflexgrave",
	XK_Ecircumflexhook:     "Ecircumflexhook",
	XK_ecircumflexhook:     "ecircumflexhook",
	XK_Ecircumflextilde:    "Ecircumflextilde",
	XK_ecircumflextilde:    "ecircumflextilde",
	XK_Ecircumflexbelowdot: "Ecircumflexbelowdot",
	XK_ecircumflexbelowdot: "ecircumflexbelowdot",
	XK_Ihook:               "Ihook",
	XK_ihook:               "ihook",
	XK_Ibelowdot:           "Ibelowdot",
	XK_ibelowdot:           "ibelowdot",
	XK_Obelowdot:           "Obelowdot",
	XK_obelowdot:           "obelowdot",
	XK_Ohook:               "Ohook",
	XK_ohook:               "ohook",
	XK_Ocircumflexacute:    "Ocircumflexacute",
	XK_ocircumflexacute:    "ocircumflexacute",
	XK_Ocircumflexgrave:    "Ocircumflexgrave",
	XK_ocircumflexgrave:    "ocircumflexgrave",
	XK_Ocircumflexhook:     "Ocircumflexhook",
	XK_ocircumflexhook:     "ocircumflexhook",
	XK_Ocircumflextilde:    "Ocircumflextilde",
	XK_ocircumflextilde:    "ocircumflextilde",
	XK_Ocircumflexbelowdot: "Ocircumflexbelowdot",
	XK_ocircumflexbelowdot: "ocircumflexbelowdot",
	XK_Ohornacute:          "Ohornacute",
	XK_ohornacute:          "ohornacute",
	XK_Ohorngrave:          "Ohorngrave",
	XK_ohorngrave:          "ohorngrave",
	XK_Ohornhook:           "Ohornhook",
	XK_ohornhook:           "ohornhook",
	XK_Ohorntilde:          "Ohorntilde",
	XK_ohorntilde:          "ohorntilde",
	XK_Ohornbelowdot:       "Ohornbelowdot",
	XK_ohornbelowdot:       "ohornbelowdot",
	XK_Ubelowdot:           "Ubelowdot",
	XK_ubelowdot:           "ubelowdot",
	XK_Uhook:               "Uhook",
	XK_uhook:               "uhook",
	XK_Uhornacute:          "Uhornacute",
	XK_uhornacute:          "uhornacute",
	XK_Uhorngrave:          "Uhorngrave",
	XK_uhorngrave:          "uhorngrave",
	XK_Uhornhook:           "Uhornhook",
	XK_uhornhook:           "uhornhook",
	XK_Uhorntilde:          "Uhorntilde",
	XK_uhorntilde:          "uhorntilde",
	XK_Uhornbelowdot:       "Uhornbelowdot",
	XK_uhornbelowdot:       "uhornbelowdot",
	XK_Ybelowdot:           "Ybelowdot",
	XK_ybelowdot:           "ybelowdot",
	XK_Yhook:               "Yhook",
	XK_yhook:               "yhook",
	XK_Ytilde:              "Ytilde",
	XK_ytilde:              "ytilde",
	XK_Ohorn:               "Ohorn",
	XK_ohorn:               "ohorn",
	XK_Uhorn:               "Uhorn",
	XK_uhorn:               "uhorn",
	XK_EcuSign:             "EcuSign",
	XK_ColonSign:           "ColonSign",
	XK_CruzeiroSign:        "CruzeiroSign",
	XK_FFrancSign:          "FFrancSign",
	XK_LiraSign:            "LiraSign",
	XK_MillSign:            "MillSign",
	XK_NairaSign:           "NairaSign",
	XK_PesetaSign:          "PesetaSign",
	XK_RupeeSign:           "RupeeSign",
	XK_WonSign:             "WonSign",
	XK_NewSheqelSign:       "NewSheqelSign",
	XK_DongSign:            "DongSign",
	XK_EuroSign:            "EuroSign",
	XK_zerosuperior:        "zerosuperior",
	XK_foursuperior:        "foursuperior",
	XK_fivesuperior:        "fivesuperior",
	XK_sixsuperior:         "sixsuperior",
	XK_sevensuperior:       "sevensuperior",
	XK_eightsuperior:       "eightsuperior",
	XK_ninesuperior:        "ninesuperior",
	XK_zerosubscript:       "zerosubscript",
	XK_onesubscript:        "onesubscript",
	XK_twosubscript:        "twosubscript",
	XK_threesubscript:      "threesubscript",
	XK_foursubscript:       "foursubscript",
	XK_fivesubscript:       "fivesubscript",
	XK_sixsubscript:        "sixsubscript",
	XK_sevensubscript:      "sevensubscript",
	XK_eightsubscript:      "eightsubscript",
	XK_ninesubscript:       "ninesubscript",
	XK_partdifferential:    "partdifferential",
	XK_emptyset:            "emptyset",
	XK_elementof:           "elementof",
	XK_notelementof:        "notelementof",
	XK_containsas:          "containsas",
	XK_squareroot:          "squareroot",
	// XK_cuberoot == XK_squareroot #
	// XK_fourthroot == XK_squareroot #
	XK_dintegral: "dintegral",
	// XK_tintegral == XK_dintegral #
	XK_because:               "because",
	XK_approxeq:              "approxeq",
	XK_notapproxeq:           "notapproxeq",
	XK_notidentical:          "notidentical",
	XK_stricteq:              "stricteq",
	XK_braille_dot_1:         "braille_dot_1",
	XK_braille_dot_2:         "braille_dot_2",
	XK_braille_dot_3:         "braille_dot_3",
	XK_braille_dot_4:         "braille_dot_4",
	XK_braille_dot_5:         "braille_dot_5",
	XK_braille_dot_6:         "braille_dot_6",
	XK_braille_dot_7:         "braille_dot_7",
	XK_braille_dot_8:         "braille_dot_8",
	XK_braille_dot_9:         "braille_dot_9",
	XK_braille_dot_10:        "braille_dot_10",
	XK_braille_blank:         "braille_blank",
	XK_braille_dots_1:        "braille_dots_1",
	XK_braille_dots_2:        "braille_dots_2",
	XK_braille_dots_12:       "braille_dots_12",
	XK_braille_dots_3:        "braille_dots_3",
	XK_braille_dots_13:       "braille_dots_13",
	XK_braille_dots_23:       "braille_dots_23",
	XK_braille_dots_123:      "braille_dots_123",
	XK_braille_dots_4:        "braille_dots_4",
	XK_braille_dots_14:       "braille_dots_14",
	XK_braille_dots_24:       "braille_dots_24",
	XK_braille_dots_124:      "braille_dots_124",
	XK_braille_dots_34:       "braille_dots_34",
	XK_braille_dots_134:      "braille_dots_134",
	XK_braille_dots_234:      "braille_dots_234",
	XK_braille_dots_1234:     "braille_dots_1234",
	XK_braille_dots_5:        "braille_dots_5",
	XK_braille_dots_15:       "braille_dots_15",
	XK_braille_dots_25:       "braille_dots_25",
	XK_braille_dots_125:      "braille_dots_125",
	XK_braille_dots_35:       "braille_dots_35",
	XK_braille_dots_135:      "braille_dots_135",
	XK_braille_dots_235:      "braille_dots_235",
	XK_braille_dots_1235:     "braille_dots_1235",
	XK_braille_dots_45:       "braille_dots_45",
	XK_braille_dots_145:      "braille_dots_145",
	XK_braille_dots_245:      "braille_dots_245",
	XK_braille_dots_1245:     "braille_dots_1245",
	XK_braille_dots_345:      "braille_dots_345",
	XK_braille_dots_1345:     "braille_dots_1345",
	XK_braille_dots_2345:     "braille_dots_2345",
	XK_braille_dots_12345:    "braille_dots_12345",
	XK_braille_dots_6:        "braille_dots_6",
	XK_braille_dots_16:       "braille_dots_16",
	XK_braille_dots_26:       "braille_dots_26",
	XK_braille_dots_126:      "braille_dots_126",
	XK_braille_dots_36:       "braille_dots_36",
	XK_braille_dots_136:      "braille_dots_136",
	XK_braille_dots_236:      "braille_dots_236",
	XK_braille_dots_1236:     "braille_dots_1236",
	XK_braille_dots_46:       "braille_dots_46",
	XK_braille_dots_146:      "braille_dots_146",
	XK_braille_dots_246:      "braille_dots_246",
	XK_braille_dots_1246:     "braille_dots_1246",
	XK_braille_dots_346:      "braille_dots_346",
	XK_braille_dots_1346:     "braille_dots_1346",
	XK_braille_dots_2346:     "braille_dots_2346",
	XK_braille_dots_12346:    "braille_dots_12346",
	XK_braille_dots_56:       "braille_dots_56",
	XK_braille_dots_156:      "braille_dots_156",
	XK_braille_dots_256:      "braille_dots_256",
	XK_braille_dots_1256:     "braille_dots_1256",
	XK_braille_dots_356:      "braille_dots_356",
	XK_braille_dots_1356:     "braille_dots_1356",
	XK_braille_dots_2356:     "braille_dots_2356",
	XK_braille_dots_12356:    "braille_dots_12356",
	XK_braille_dots_456:      "braille_dots_456",
	XK_braille_dots_1456:     "braille_dots_1456",
	XK_braille_dots_2456:     "braille_dots_2456",
	XK_braille_dots_12456:    "braille_dots_12456",
	XK_braille_dots_3456:     "braille_dots_3456",
	XK_braille_dots_13456:    "braille_dots_13456",
	XK_braille_dots_23456:    "braille_dots_23456",
	XK_braille_dots_123456:   "braille_dots_123456",
	XK_braille_dots_7:        "braille_dots_7",
	XK_braille_dots_17:       "braille_dots_17",
	XK_braille_dots_27:       "braille_dots_27",
	XK_braille_dots_127:      "braille_dots_127",
	XK_braille_dots_37:       "braille_dots_37",
	XK_braille_dots_137:      "braille_dots_137",
	XK_braille_dots_237:      "braille_dots_237",
	XK_braille_dots_1237:     "braille_dots_1237",
	XK_braille_dots_47:       "braille_dots_47",
	XK_braille_dots_147:      "braille_dots_147",
	XK_braille_dots_247:      "braille_dots_247",
	XK_braille_dots_1247:     "braille_dots_1247",
	XK_braille_dots_347:      "braille_dots_347",
	XK_braille_dots_1347:     "braille_dots_1347",
	XK_braille_dots_2347:     "braille_dots_2347",
	XK_braille_dots_12347:    "braille_dots_12347",
	XK_braille_dots_57:       "braille_dots_57",
	XK_braille_dots_157:      "braille_dots_157",
	XK_braille_dots_257:      "braille_dots_257",
	XK_braille_dots_1257:     "braille_dots_1257",
	XK_braille_dots_357:      "braille_dots_357",
	XK_braille_dots_1357:     "braille_dots_1357",
	XK_braille_dots_2357:     "braille_dots_2357",
	XK_braille_dots_12357:    "braille_dots_12357",
	XK_braille_dots_457:      "braille_dots_457",
	XK_braille_dots_1457:     "braille_dots_1457",
	XK_braille_dots_2457:     "braille_dots_2457",
	XK_braille_dots_12457:    "braille_dots_12457",
	XK_braille_dots_3457:     "braille_dots_3457",
	XK_braille_dots_13457:    "braille_dots_13457",
	XK_braille_dots_23457:    "braille_dots_23457",
	XK_braille_dots_123457:   "braille_dots_123457",
	XK_braille_dots_67:       "braille_dots_67",
	XK_braille_dots_167:      "braille_dots_167",
	XK_braille_dots_267:      "braille_dots_267",
	XK_braille_dots_1267:     "braille_dots_1267",
	XK_braille_dots_367:      "braille_dots_367",
	XK_braille_dots_1367:     "braille_dots_1367",
	XK_braille_dots_2367:     "braille_dots_2367",
	XK_braille_dots_12367:    "braille_dots_12367",
	XK_braille_dots_467:      "braille_dots_467",
	XK_braille_dots_1467:     "braille_dots_1467",
	XK_braille_dots_2467:     "braille_dots_2467",
	XK_braille_dots_12467:    "braille_dots_12467",
	XK_braille_dots_3467:     "braille_dots_3467",
	XK_braille_dots_13467:    "braille_dots_13467",
	XK_braille_dots_23467:    "braille_dots_23467",
	XK_braille_dots_123467:   "braille_dots_123467",
	XK_braille_dots_567:      "braille_dots_567",
	XK_braille_dots_1567:     "braille_dots_1567",
	XK_braille_dots_2567:     "braille_dots_2567",
	XK_braille_dots_12567:    "braille_dots_12567",
	XK_braille_dots_3567:     "braille_dots_3567",
	XK_braille_dots_13567:    "braille_dots_13567",
	XK_braille_dots_23567:    "braille_dots_23567",
	XK_braille_dots_123567:   "braille_dots_123567",
	XK_braille_dots_4567:     "braille_dots_4567",
	XK_braille_dots_14567:    "braille_dots_14567",
	XK_braille_dots_24567:    "braille_dots_24567",
	XK_braille_dots_124567:   "braille_dots_124567",
	XK_braille_dots_34567:    "braille_dots_34567",
	XK_braille_dots_134567:   "braille_dots_134567",
	XK_braille_dots_234567:   "braille_dots_234567",
	XK_braille_dots_1234567:  "braille_dots_1234567",
	XK_braille_dots_8:        "braille_dots_8",
	XK_braille_dots_18:       "braille_dots_18",
	XK_braille_dots_28:       "braille_dots_28",
	XK_braille_dots_128:      "braille_dots_128",
	XK_braille_dots_38:       "braille_dots_38",
	XK_braille_dots_138:      "braille_dots_138",
	XK_braille_dots_238:      "braille_dots_238",
	XK_braille_dots_1238:     "braille_dots_1238",
	XK_braille_dots_48:       "braille_dots_48",
	XK_braille_dots_148:      "braille_dots_148",
	XK_braille_dots_248:      "braille_dots_248",
	XK_braille_dots_1248:     "braille_dots_1248",
	XK_braille_dots_348:      "braille_dots_348",
	XK_braille_dots_1348:     "braille_dots_1348",
	XK_braille_dots_2348:     "braille_dots_2348",
	XK_braille_dots_12348:    "braille_dots_12348",
	XK_braille_dots_58:       "braille_dots_58",
	XK_braille_dots_158:      "braille_dots_158",
	XK_braille_dots_258:      "braille_dots_258",
	XK_braille_dots_1258:     "braille_dots_1258",
	XK_braille_dots_358:      "braille_dots_358",
	XK_braille_dots_1358:     "braille_dots_1358",
	XK_braille_dots_2358:     "braille_dots_2358",
	XK_braille_dots_12358:    "braille_dots_12358",
	XK_braille_dots_458:      "braille_dots_458",
	XK_braille_dots_1458:     "braille_dots_1458",
	XK_braille_dots_2458:     "braille_dots_2458",
	XK_braille_dots_12458:    "braille_dots_12458",
	XK_braille_dots_3458:     "braille_dots_3458",
	XK_braille_dots_13458:    "braille_dots_13458",
	XK_braille_dots_23458:    "braille_dots_23458",
	XK_braille_dots_123458:   "braille_dots_123458",
	XK_braille_dots_68:       "braille_dots_68",
	XK_braille_dots_168:      "braille_dots_168",
	XK_braille_dots_268:      "braille_dots_268",
	XK_braille_dots_1268:     "braille_dots_1268",
	XK_braille_dots_368:      "braille_dots_368",
	XK_braille_dots_1368:     "braille_dots_1368",
	XK_braille_dots_2368:     "braille_dots_2368",
	XK_braille_dots_12368:    "braille_dots_12368",
	XK_braille_dots_468:      "braille_dots_468",
	XK_braille_dots_1468:     "braille_dots_1468",
	XK_braille_dots_2468:     "braille_dots_2468",
	XK_braille_dots_12468:    "braille_dots_12468",
	XK_braille_dots_3468:     "braille_dots_3468",
	XK_braille_dots_13468:    "braille_dots_13468",
	XK_braille_dots_23468:    "braille_dots_23468",
	XK_braille_dots_123468:   "braille_dots_123468",
	XK_braille_dots_568:      "braille_dots_568",
	XK_braille_dots_1568:     "braille_dots_1568",
	XK_braille_dots_2568:     "braille_dots_2568",
	XK_braille_dots_12568:    "braille_dots_12568",
	XK_braille_dots_3568:     "braille_dots_3568",
	XK_braille_dots_13568:    "braille_dots_13568",
	XK_braille_dots_23568:    "braille_dots_23568",
	XK_braille_dots_123568:   "braille_dots_123568",
	XK_braille_dots_4568:     "braille_dots_4568",
	XK_braille_dots_14568:    "braille_dots_14568",
	XK_braille_dots_24568:    "braille_dots_24568",
	XK_braille_dots_124568:   "braille_dots_124568",
	XK_braille_dots_34568:    "braille_dots_34568",
	XK_braille_dots_134568:   "braille_dots_134568",
	XK_braille_dots_234568:   "braille_dots_234568",
	XK_braille_dots_1234568:  "braille_dots_1234568",
	XK_braille_dots_78:       "braille_dots_78",
	XK_braille_dots_178:      "braille_dots_178",
	XK_braille_dots_278:      "braille_dots_278",
	XK_braille_dots_1278:     "braille_dots_1278",
	XK_braille_dots_378:      "braille_dots_378",
	XK_braille_dots_1378:     "braille_dots_1378",
	XK_braille_dots_2378:     "braille_dots_2378",
	XK_braille_dots_12378:    "braille_dots_12378",
	XK_braille_dots_478:      "braille_dots_478",
	XK_braille_dots_1478:     "braille_dots_1478",
	XK_braille_dots_2478:     "braille_dots_2478",
	XK_braille_dots_12478:    "braille_dots_12478",
	XK_braille_dots_3478:     "braille_dots_3478",
	XK_braille_dots_13478:    "braille_dots_13478",
	XK_braille_dots_23478:    "braille_dots_23478",
	XK_braille_dots_123478:   "braille_dots_123478",
	XK_braille_dots_578:      "braille_dots_578",
	XK_braille_dots_1578:     "braille_dots_1578",
	XK_braille_dots_2578:     "braille_dots_2578",
	XK_braille_dots_12578:    "braille_dots_12578",
	XK_braille_dots_3578:     "braille_dots_3578",
	XK_braille_dots_13578:    "braille_dots_13578",
	XK_braille_dots_23578:    "braille_dots_23578",
	XK_braille_dots_123578:   "braille_dots_123578",
	XK_braille_dots_4578:     "braille_dots_4578",
	XK_braille_dots_14578:    "braille_dots_14578",
	XK_braille_dots_24578:    "braille_dots_24578",
	XK_braille_dots_124578:   "braille_dots_124578",
	XK_braille_dots_34578:    "braille_dots_34578",
	XK_braille_dots_134578:   "braille_dots_134578",
	XK_braille_dots_234578:   "braille_dots_234578",
	XK_braille_dots_1234578:  "braille_dots_1234578",
	XK_braille_dots_678:      "braille_dots_678",
	XK_braille_dots_1678:     "braille_dots_1678",
	XK_braille_dots_2678:     "braille_dots_2678",
	XK_braille_dots_12678:    "braille_dots_12678",
	XK_braille_dots_3678:     "braille_dots_3678",
	XK_braille_dots_13678:    "braille_dots_13678",
	XK_braille_dots_23678:    "braille_dots_23678",
	XK_braille_dots_123678:   "braille_dots_123678",
	XK_braille_dots_4678:     "braille_dots_4678",
	XK_braille_dots_14678:    "braille_dots_14678",
	XK_braille_dots_24678:    "braille_dots_24678",
	XK_braille_dots_124678:   "braille_dots_124678",
	XK_braille_dots_34678:    "braille_dots_34678",
	XK_braille_dots_134678:   "braille_dots_134678",
	XK_braille_dots_234678:   "braille_dots_234678",
	XK_braille_dots_1234678:  "braille_dots_1234678",
	XK_braille_dots_5678:     "braille_dots_5678",
	XK_braille_dots_15678:    "braille_dots_15678",
	XK_braille_dots_25678:    "braille_dots_25678",
	XK_braille_dots_125678:   "braille_dots_125678",
	XK_braille_dots_35678:    "braille_dots_35678",
	XK_braille_dots_135678:   "braille_dots_135678",
	XK_braille_dots_235678:   "braille_dots_235678",
	XK_braille_dots_1235678:  "braille_dots_1235678",
	XK_braille_dots_45678:    "braille_dots_45678",
	XK_braille_dots_145678:   "braille_dots_145678",
	XK_braille_dots_245678:   "braille_dots_245678",
	XK_braille_dots_1245678:  "braille_dots_1245678",
	XK_braille_dots_345678:   "braille_dots_345678",
	XK_braille_dots_1345678:  "braille_dots_1345678",
	XK_braille_dots_2345678:  "braille_dots_2345678",
	XK_braille_dots_12345678: "braille_dots_12345678",
	XK_Sinh_ng:               "Sinh_ng",
	XK_Sinh_h2:               "Sinh_h2",
	XK_Sinh_a:                "Sinh_a",
	XK_Sinh_aa:               "Sinh_aa",
	XK_Sinh_ae:               "Sinh_ae",
	XK_Sinh_aee:              "Sinh_aee",
	XK_Sinh_i:                "Sinh_i",
	XK_Sinh_ii:               "Sinh_ii",
	XK_Sinh_u:                "Sinh_u",
	XK_Sinh_uu:               "Sinh_uu",
	XK_Sinh_ri:               "Sinh_ri",
	XK_Sinh_rii:              "Sinh_rii",
	XK_Sinh_lu:               "Sinh_lu",
	XK_Sinh_luu:              "Sinh_luu",
	XK_Sinh_e:                "Sinh_e",
	XK_Sinh_ee:               "Sinh_ee",
	XK_Sinh_ai:               "Sinh_ai",
	XK_Sinh_o:                "Sinh_o",
	XK_Sinh_oo:               "Sinh_oo",
	XK_Sinh_au:               "Sinh_au",
	XK_Sinh_ka:               "Sinh_ka",
	XK_Sinh_kha:              "Sinh_kha",
	XK_Sinh_ga:               "Sinh_ga",
	XK_Sinh_gha:              "Sinh_gha",
	XK_Sinh_ng2:              "Sinh_ng2",
	XK_Sinh_nga:              "Sinh_nga",
	XK_Sinh_ca:               "Sinh_ca",
	XK_Sinh_cha:              "Sinh_cha",
	XK_Sinh_ja:               "Sinh_ja",
	XK_Sinh_jha:              "Sinh_jha",
	XK_Sinh_nya:              "Sinh_nya",
	XK_Sinh_jnya:             "Sinh_jnya",
	XK_Sinh_nja:              "Sinh_nja",
	XK_Sinh_tta:              "Sinh_tta",
	XK_Sinh_ttha:             "Sinh_ttha",
	XK_Sinh_dda:              "Sinh_dda",
	XK_Sinh_ddha:             "Sinh_ddha",
	XK_Sinh_nna:              "Sinh_nna",
	XK_Sinh_ndda:             "Sinh_ndda",
	XK_Sinh_tha:              "Sinh_tha",
	XK_Sinh_thha:             "Sinh_thha",
	XK_Sinh_dha:              "Sinh_dha",
	XK_Sinh_dhha:             "Sinh_dhha",
	XK_Sinh_na:               "Sinh_na",
	XK_Sinh_ndha:             "Sinh_ndha",
	XK_Sinh_pa:               "Sinh_pa",
	XK_Sinh_pha:              "Sinh_pha",
	XK_Sinh_ba:               "Sinh_ba",
	XK_Sinh_bha:              "Sinh_bha",
	XK_Sinh_ma:               "Sinh_ma",
	XK_Sinh_mba:              "Sinh_mba",
	XK_Sinh_ya:               "Sinh_ya",
	XK_Sinh_ra:               "Sinh_ra",
	XK_Sinh_la:               "Sinh_la",
	XK_Sinh_va:               "Sinh_va",
	XK_Sinh_sha:              "Sinh_sha",
	XK_Sinh_ssha:             "Sinh_ssha",
	XK_Sinh_sa:               "Sinh_sa",
	XK_Sinh_ha:               "Sinh_ha",
	XK_Sinh_lla:              "Sinh_lla",
	XK_Sinh_fa:               "Sinh_fa",
	XK_Sinh_al:               "Sinh_al",
	XK_Sinh_aa2:              "Sinh_aa2",
	XK_Sinh_ae2:              "Sinh_ae2",
	XK_Sinh_aee2:             "Sinh_aee2",
	XK_Sinh_i2:               "Sinh_i2",
	XK_Sinh_ii2:              "Sinh_ii2",
	XK_Sinh_u2:               "Sinh_u2",
	XK_Sinh_uu2:              "Sinh_uu2",
	XK_Sinh_ru2:              "Sinh_ru2",
	XK_Sinh_e2:               "Sinh_e2",
	XK_Sinh_ee2:              "Sinh_ee2",
	XK_Sinh_ai2:              "Sinh_ai2",
	XK_Sinh_o2:               "Sinh_o2",
	XK_Sinh_oo2:              "Sinh_oo2",
	XK_Sinh_au2:              "Sinh_au2",
	XK_Sinh_lu2:              "Sinh_lu2",
	XK_Sinh_ruu2:             "Sinh_ruu2",
	XK_Sinh_luu2:             "Sinh_luu2",
	XK_Sinh_kunddaliya:       "Sinh_kunddaliya",
	XF86XK_ModeLock:          "XF86ModeLock",
	XF86XK_MonBrightnessUp:   "XF86MonBrightnessUp",
	XF86XK_MonBrightnessDown: "XF86MonBrightnessDown",
	XF86XK_KbdLightOnOff:     "XF86KbdLightOnOff",
	XF86XK_KbdBrightnessUp:   "XF86KbdBrightnessUp",
	XF86XK_KbdBrightnessDown: "XF86KbdBrightnessDown",
	XF86XK_Standby:           "XF86Standby",
	XF86XK_AudioLowerVolume:  "XF86AudioLowerVolume",
	XF86XK_AudioMute:         "XF86AudioMute",
	XF86XK_AudioRaiseVolume:  "XF86AudioRaiseVolume",
	XF86XK_AudioPlay:         "XF86AudioPlay",
	XF86XK_AudioStop:         "XF86AudioStop",
	XF86XK_AudioPrev:         "XF86AudioPrev",
	XF86XK_AudioNext:         "XF86AudioNext",
	XF86XK_HomePage:          "XF86HomePage",
	XF86XK_Mail:              "XF86Mail",
	XF86XK_Start:             "XF86Start",
	XF86XK_Search:            "XF86Search",
	XF86XK_AudioRecord:       "XF86AudioRecord",
	XF86XK_Calculator:        "XF86Calculator",
	XF86XK_Memo:              "XF86Memo",
	XF86XK_ToDoList:          "XF86ToDoList",
	XF86XK_Calendar:          "XF86Calendar",
	XF86XK_PowerDown:         "XF86PowerDown",
	XF86XK_ContrastAdjust:    "XF86ContrastAdjust",
	XF86XK_RockerUp:          "XF86RockerUp",
	XF86XK_RockerDown:        "XF86RockerDown",
	XF86XK_RockerEnter:       "XF86RockerEnter",
	XF86XK_Back:              "XF86Back",
	XF86XK_Forward:           "XF86Forward",
	XF86XK_Stop:              "XF86Stop",
	XF86XK_Refresh:           "XF86Refresh",
	XF86XK_PowerOff:          "XF86PowerOff",
	XF86XK_WakeUp:            "XF86WakeUp",
	XF86XK_Eject:             "XF86Eject",
	XF86XK_ScreenSaver:       "XF86ScreenSaver",
	XF86XK_WWW:               "XF86WWW",
	XF86XK_Sleep:             "XF86Sleep",
	XF86XK_Favorites:         "XF86Favorites",
	XF86XK_AudioPause:        "XF86AudioPause",
	XF86XK_AudioMedia:        "XF86AudioMedia",
	XF86XK_MyComputer:        "XF86MyComputer",
	XF86XK_VendorHome:        "XF86VendorHome",
	XF86XK_LightBulb:         "XF86LightBulb",
	XF86XK_Shop:              "XF86Shop",
	XF86XK_History:           "XF86History",
	XF86XK_OpenURL:           "XF86OpenURL",
	XF86XK_AddFavorite:       "XF86AddFavorite",
	XF86XK_HotLinks:          "XF86HotLinks",
	XF86XK_BrightnessAdjust:  "XF86BrightnessAdjust",
	XF86XK_Finance:           "XF86Finance",
	XF86XK_Community:         "XF86Community",
	XF86XK_AudioRewind:       "XF86AudioRewind",
	XF86XK_BackForward:       "XF86BackForward",
	XF86XK_Launch0:           "XF86Launch0",
	XF86XK_Launch1:           "XF86Launch1",
	XF86XK_Launch2:           "XF86Launch2",
	XF86XK_Launch3:           "XF86Launch3",
	XF86XK_Launch4:           "XF86Launch4",
	XF86XK_Launch5:           "XF86Launch5",
	XF86XK_Launch6:           "XF86Launch6",
	XF86XK_Launch7:           "XF86Launch7",
	XF86XK_Launch8:           "XF86Launch8",
	XF86XK_Launch9:           "XF86Launch9",
	XF86XK_LaunchA:           "XF86LaunchA",
	XF86XK_LaunchB:           "XF86LaunchB",
	XF86XK_LaunchC:           "XF86LaunchC",
	XF86XK_LaunchD:           "XF86LaunchD",
	XF86XK_LaunchE:           "XF86LaunchE",
	XF86XK_LaunchF:           "XF86LaunchF",
	XF86XK_ApplicationLeft:   "XF86ApplicationLeft",
	XF86XK_ApplicationRight:  "XF86ApplicationRight",
	XF86XK_Book:              "XF86Book",
	XF86XK_CD:                "XF86CD",
	XF86XK_Calculater:        "XF86Calculater",
	XF86XK_Clear:             "XF86Clear",
	XF86XK_Close:             "XF86Close",
	XF86XK_Copy:              "XF86Copy",
	XF86XK_Cut:               "XF86Cut",
	XF86XK_Display:           "XF86Display",
	XF86XK_DOS:               "XF86DOS",
	XF86XK_Documents:         "XF86Documents",
	XF86XK_Excel:             "XF86Excel",
	XF86XK_Explorer:          "XF86Explorer",
	XF86XK_Game:              "XF86Game",
	XF86XK_Go:                "XF86Go",
	XF86XK_iTouch:            "XF86iTouch",
	XF86XK_LogOff:            "XF86LogOff",
	XF86XK_Market:            "XF86Market",
	XF86XK_Meeting:           "XF86Meeting",
	XF86XK_MenuKB:            "XF86MenuKB",
	XF86XK_MenuPB:            "XF86MenuPB",
	XF86XK_MySites:           "XF86MySites",
	XF86XK_New:               "XF86New",
	XF86XK_News:              "XF86News",
	XF86XK_OfficeHome:        "XF86OfficeHome",
	XF86XK_Open:              "XF86Open",
	XF86XK_Option:            "XF86Option",
	XF86XK_Paste:             "XF86Paste",
	XF86XK_Phone:             "XF86Phone",
	XF86XK_Q:                 "XF86Q",
	XF86XK_Reply:             "XF86Reply",
	XF86XK_Reload:            "XF86Reload",
	XF86XK_RotateWindows:     "XF86RotateWindows",
	XF86XK_RotationPB:        "XF86RotationPB",
	XF86XK_RotationKB:        "XF86RotationKB",
	XF86XK_Save:              "XF86Save",
	XF86XK_ScrollUp:          "XF86ScrollUp",
	XF86XK_ScrollDown:        "XF86ScrollDown",
	XF86XK_ScrollClick:       "XF86ScrollClick",
	XF86XK_Send:              "XF86Send",
	XF86XK_Spell:             "XF86Spell",
	XF86XK_SplitScreen:       "XF86SplitScreen",
	XF86XK_Support:           "XF86Support",
	XF86XK_TaskPane:          "XF86TaskPane",
	XF86XK_Terminal:          "XF86Terminal",
	XF86XK_Tools:             "XF86Tools",
	XF86XK_Travel:            "XF86Travel",
	XF86XK_UserPB:            "XF86UserPB",
	XF86XK_User1KB:           "XF86User1KB",
	XF86XK_User2KB:           "XF86User2KB",
	XF86XK_Video:             "XF86Video",
	XF86XK_WheelButton:       "XF86WheelButton",
	XF86XK_Word:              "XF86Word",
	XF86XK_Xfer:              "XF86Xfer",
	XF86XK_ZoomIn:            "XF86ZoomIn",
	XF86XK_ZoomOut:           "XF86ZoomOut",
	XF86XK_Away:              "XF86Away",
	XF86XK_Messenger:         "XF86Messenger",
	XF86XK_WebCam:            "XF86WebCam",
	XF86XK_MailForward:       "XF86MailForward",
	XF86XK_Pictures:          "XF86Pictures",
	XF86XK_Music:             "XF86Music",
	XF86XK_Battery:           "XF86Battery",
	XF86XK_Bluetooth:         "XF86Bluetooth",
	XF86XK_WLAN:              "XF86WLAN",
	XF86XK_UWB:               "XF86UWB",
	XF86XK_AudioForward:      "XF86AudioForward",
	XF86XK_AudioRepeat:       "XF86AudioRepeat",
	XF86XK_AudioRandomPlay:   "XF86AudioRandomPlay",
	XF86XK_Subtitle:          "XF86Subtitle",
	XF86XK_AudioCycleTrack:   "XF86AudioCycleTrack",
	XF86XK_CycleAngle:        "XF86CycleAngle",
	XF86XK_FrameBack:         "XF86FrameBack",
	XF86XK_FrameForward:      "XF86FrameForward",
	XF86XK_Time:              "XF86Time",
	XF86XK_Select:            "XF86Select",
	XF86XK_View:              "XF86View",
	XF86XK_TopMenu:           "XF86TopMenu",
	XF86XK_Red:               "XF86Red",
	XF86XK_Green:             "XF86Green",
	XF86XK_Yellow:            "XF86Yellow",
	XF86XK_Blue:              "XF86Blue",
	XF86XK_Suspend:           "XF86Suspend",
	XF86XK_Hibernate:         "XF86Hibernate",
	XF86XK_TouchpadToggle:    "XF86TouchpadToggle",
	XF86XK_TouchpadOn:        "XF86TouchpadOn",
	XF86XK_TouchpadOff:       "XF86TouchpadOff",
	XF86XK_AudioMicMute:      "XF86AudioMicMute",
	XF86XK_Keyboard:          "XF86Keyboard",
	XF86XK_WWAN:              "XF86WWAN",
	XF86XK_RFKill:            "XF86RFKill",
	XF86XK_AudioPreset:       "XF86AudioPreset",
	XF86XK_Switch_VT_1:       "XF86Switch_VT_1",
	XF86XK_Switch_VT_2:       "XF86Switch_VT_2",
	XF86XK_Switch_VT_3:       "XF86Switch_VT_3",
	XF86XK_Switch_VT_4:       "XF86Switch_VT_4",
	XF86XK_Switch_VT_5:       "XF86Switch_VT_5",
	XF86XK_Switch_VT_6:       "XF86Switch_VT_6",
	XF86XK_Switch_VT_7:       "XF86Switch_VT_7",
	XF86XK_Switch_VT_8:       "XF86Switch_VT_8",
	XF86XK_Switch_VT_9:       "XF86Switch_VT_9",
	XF86XK_Switch_VT_10:      "XF86Switch_VT_10",
	XF86XK_Switch_VT_11:      "XF86Switch_VT_11",
	XF86XK_Switch_VT_12:      "XF86Switch_VT_12",
	XF86XK_Ungrab:            "XF86Ungrab",
	XF86XK_ClearGrab:         "XF86ClearGrab",
	XF86XK_Next_VMode:        "XF86Next_VMode",
	XF86XK_Prev_VMode:        "XF86Prev_VMode",
	XF86XK_LogWindowTree:     "XF86LogWindowTree",
	XF86XK_LogGrabInfo:       "XF86LogGrabInfo",
}
var EngKeysymMap = map[string]x.Keysym{
	"VoidSymbol":                  XK_VoidSymbol,
	"BackSpace":                   XK_BackSpace,
	"Tab":                         XK_Tab,
	"Linefeed":                    XK_Linefeed,
	"Clear":                       XK_Clear,
	"Return":                      XK_Return,
	"Pause":                       XK_Pause,
	"Scroll_Lock":                 XK_Scroll_Lock,
	"Sys_Req":                     XK_Sys_Req,
	"Escape":                      XK_Escape,
	"Delete":                      XK_Delete,
	"Multi_key":                   XK_Multi_key,
	"Codeinput":                   XK_Codeinput,
	"SingleCandidate":             XK_SingleCandidate,
	"MultipleCandidate":           XK_MultipleCandidate,
	"PreviousCandidate":           XK_PreviousCandidate,
	"Kanji":                       XK_Kanji,
	"Muhenkan":                    XK_Muhenkan,
	"Henkan_Mode":                 XK_Henkan_Mode,
	"Henkan":                      XK_Henkan,
	"Romaji":                      XK_Romaji,
	"Hiragana":                    XK_Hiragana,
	"Katakana":                    XK_Katakana,
	"Hiragana_Katakana":           XK_Hiragana_Katakana,
	"Zenkaku":                     XK_Zenkaku,
	"Hankaku":                     XK_Hankaku,
	"Zenkaku_Hankaku":             XK_Zenkaku_Hankaku,
	"Touroku":                     XK_Touroku,
	"Massyo":                      XK_Massyo,
	"Kana_Lock":                   XK_Kana_Lock,
	"Kana_Shift":                  XK_Kana_Shift,
	"Eisu_Shift":                  XK_Eisu_Shift,
	"Eisu_toggle":                 XK_Eisu_toggle,
	"Kanji_Bangou":                XK_Kanji_Bangou,
	"Zen_Koho":                    XK_Zen_Koho,
	"Mae_Koho":                    XK_Mae_Koho,
	"Home":                        XK_Home,
	"Left":                        XK_Left,
	"Up":                          XK_Up,
	"Right":                       XK_Right,
	"Down":                        XK_Down,
	"Prior":                       XK_Prior,
	"Page_Up":                     XK_Page_Up,
	"Next":                        XK_Next,
	"Page_Down":                   XK_Page_Down,
	"End":                         XK_End,
	"Begin":                       XK_Begin,
	"Select":                      XK_Select,
	"Print":                       XK_Print,
	"Execute":                     XK_Execute,
	"Insert":                      XK_Insert,
	"Undo":                        XK_Undo,
	"Redo":                        XK_Redo,
	"Menu":                        XK_Menu,
	"Find":                        XK_Find,
	"Cancel":                      XK_Cancel,
	"Help":                        XK_Help,
	"Break":                       XK_Break,
	"Mode_switch":                 XK_Mode_switch,
	"script_switch":               XK_script_switch,
	"Num_Lock":                    XK_Num_Lock,
	"KP_Space":                    XK_KP_Space,
	"KP_Tab":                      XK_KP_Tab,
	"KP_Enter":                    XK_KP_Enter,
	"KP_F1":                       XK_KP_F1,
	"KP_F2":                       XK_KP_F2,
	"KP_F3":                       XK_KP_F3,
	"KP_F4":                       XK_KP_F4,
	"KP_Home":                     XK_KP_Home,
	"KP_Left":                     XK_KP_Left,
	"KP_Up":                       XK_KP_Up,
	"KP_Right":                    XK_KP_Right,
	"KP_Down":                     XK_KP_Down,
	"KP_Prior":                    XK_KP_Prior,
	"KP_Page_Up":                  XK_KP_Page_Up,
	"KP_Next":                     XK_KP_Next,
	"KP_Page_Down":                XK_KP_Page_Down,
	"KP_End":                      XK_KP_End,
	"KP_Begin":                    XK_KP_Begin,
	"KP_Insert":                   XK_KP_Insert,
	"KP_Delete":                   XK_KP_Delete,
	"KP_Equal":                    XK_KP_Equal,
	"KP_Multiply":                 XK_KP_Multiply,
	"KP_Add":                      XK_KP_Add,
	"KP_Separator":                XK_KP_Separator,
	"KP_Subtract":                 XK_KP_Subtract,
	"KP_Decimal":                  XK_KP_Decimal,
	"KP_Divide":                   XK_KP_Divide,
	"KP_0":                        XK_KP_0,
	"KP_1":                        XK_KP_1,
	"KP_2":                        XK_KP_2,
	"KP_3":                        XK_KP_3,
	"KP_4":                        XK_KP_4,
	"KP_5":                        XK_KP_5,
	"KP_6":                        XK_KP_6,
	"KP_7":                        XK_KP_7,
	"KP_8":                        XK_KP_8,
	"KP_9":                        XK_KP_9,
	"F1":                          XK_F1,
	"F2":                          XK_F2,
	"F3":                          XK_F3,
	"F4":                          XK_F4,
	"F5":                          XK_F5,
	"F6":                          XK_F6,
	"F7":                          XK_F7,
	"F8":                          XK_F8,
	"F9":                          XK_F9,
	"F10":                         XK_F10,
	"F11":                         XK_F11,
	"L1":                          XK_L1,
	"F12":                         XK_F12,
	"L2":                          XK_L2,
	"F13":                         XK_F13,
	"L3":                          XK_L3,
	"F14":                         XK_F14,
	"L4":                          XK_L4,
	"F15":                         XK_F15,
	"L5":                          XK_L5,
	"F16":                         XK_F16,
	"L6":                          XK_L6,
	"F17":                         XK_F17,
	"L7":                          XK_L7,
	"F18":                         XK_F18,
	"L8":                          XK_L8,
	"F19":                         XK_F19,
	"L9":                          XK_L9,
	"F20":                         XK_F20,
	"L10":                         XK_L10,
	"F21":                         XK_F21,
	"R1":                          XK_R1,
	"F22":                         XK_F22,
	"R2":                          XK_R2,
	"F23":                         XK_F23,
	"R3":                          XK_R3,
	"F24":                         XK_F24,
	"R4":                          XK_R4,
	"F25":                         XK_F25,
	"R5":                          XK_R5,
	"F26":                         XK_F26,
	"R6":                          XK_R6,
	"F27":                         XK_F27,
	"R7":                          XK_R7,
	"F28":                         XK_F28,
	"R8":                          XK_R8,
	"F29":                         XK_F29,
	"R9":                          XK_R9,
	"F30":                         XK_F30,
	"R10":                         XK_R10,
	"F31":                         XK_F31,
	"R11":                         XK_R11,
	"F32":                         XK_F32,
	"R12":                         XK_R12,
	"F33":                         XK_F33,
	"R13":                         XK_R13,
	"F34":                         XK_F34,
	"R14":                         XK_R14,
	"F35":                         XK_F35,
	"R15":                         XK_R15,
	"Shift_L":                     XK_Shift_L,
	"Shift_R":                     XK_Shift_R,
	"Control_L":                   XK_Control_L,
	"Control_R":                   XK_Control_R,
	"Caps_Lock":                   XK_Caps_Lock,
	"Shift_Lock":                  XK_Shift_Lock,
	"Meta_L":                      XK_Meta_L,
	"Meta_R":                      XK_Meta_R,
	"Alt_L":                       XK_Alt_L,
	"Alt_R":                       XK_Alt_R,
	"Super_L":                     XK_Super_L,
	"Super_R":                     XK_Super_R,
	"Hyper_L":                     XK_Hyper_L,
	"Hyper_R":                     XK_Hyper_R,
	"ISO_Lock":                    XK_ISO_Lock,
	"ISO_Level2_Latch":            XK_ISO_Level2_Latch,
	"ISO_Level3_Shift":            XK_ISO_Level3_Shift,
	"ISO_Level3_Latch":            XK_ISO_Level3_Latch,
	"ISO_Level3_Lock":             XK_ISO_Level3_Lock,
	"ISO_Level5_Shift":            XK_ISO_Level5_Shift,
	"ISO_Level5_Latch":            XK_ISO_Level5_Latch,
	"ISO_Level5_Lock":             XK_ISO_Level5_Lock,
	"ISO_Group_Shift":             XK_ISO_Group_Shift,
	"ISO_Group_Latch":             XK_ISO_Group_Latch,
	"ISO_Group_Lock":              XK_ISO_Group_Lock,
	"ISO_Next_Group":              XK_ISO_Next_Group,
	"ISO_Next_Group_Lock":         XK_ISO_Next_Group_Lock,
	"ISO_Prev_Group":              XK_ISO_Prev_Group,
	"ISO_Prev_Group_Lock":         XK_ISO_Prev_Group_Lock,
	"ISO_First_Group":             XK_ISO_First_Group,
	"ISO_First_Group_Lock":        XK_ISO_First_Group_Lock,
	"ISO_Last_Group":              XK_ISO_Last_Group,
	"ISO_Last_Group_Lock":         XK_ISO_Last_Group_Lock,
	"ISO_Left_Tab":                XK_ISO_Left_Tab,
	"ISO_Move_Line_Up":            XK_ISO_Move_Line_Up,
	"ISO_Move_Line_Down":          XK_ISO_Move_Line_Down,
	"ISO_Partial_Line_Up":         XK_ISO_Partial_Line_Up,
	"ISO_Partial_Line_Down":       XK_ISO_Partial_Line_Down,
	"ISO_Partial_Space_Left":      XK_ISO_Partial_Space_Left,
	"ISO_Partial_Space_Right":     XK_ISO_Partial_Space_Right,
	"ISO_Set_Margin_Left":         XK_ISO_Set_Margin_Left,
	"ISO_Set_Margin_Right":        XK_ISO_Set_Margin_Right,
	"ISO_Release_Margin_Left":     XK_ISO_Release_Margin_Left,
	"ISO_Release_Margin_Right":    XK_ISO_Release_Margin_Right,
	"ISO_Release_Both_Margins":    XK_ISO_Release_Both_Margins,
	"ISO_Fast_Cursor_Left":        XK_ISO_Fast_Cursor_Left,
	"ISO_Fast_Cursor_Right":       XK_ISO_Fast_Cursor_Right,
	"ISO_Fast_Cursor_Up":          XK_ISO_Fast_Cursor_Up,
	"ISO_Fast_Cursor_Down":        XK_ISO_Fast_Cursor_Down,
	"ISO_Continuous_Underline":    XK_ISO_Continuous_Underline,
	"ISO_Discontinuous_Underline": XK_ISO_Discontinuous_Underline,
	"ISO_Emphasize":               XK_ISO_Emphasize,
	"ISO_Center_Object":           XK_ISO_Center_Object,
	"ISO_Enter":                   XK_ISO_Enter,
	"dead_grave":                  XK_dead_grave,
	"dead_acute":                  XK_dead_acute,
	"dead_circumflex":             XK_dead_circumflex,
	"dead_tilde":                  XK_dead_tilde,
	"dead_perispomeni":            XK_dead_perispomeni,
	"dead_macron":                 XK_dead_macron,
	"dead_breve":                  XK_dead_breve,
	"dead_abovedot":               XK_dead_abovedot,
	"dead_diaeresis":              XK_dead_diaeresis,
	"dead_abovering":              XK_dead_abovering,
	"dead_doubleacute":            XK_dead_doubleacute,
	"dead_caron":                  XK_dead_caron,
	"dead_cedilla":                XK_dead_cedilla,
	"dead_ogonek":                 XK_dead_ogonek,
	"dead_iota":                   XK_dead_iota,
	"dead_voiced_sound":           XK_dead_voiced_sound,
	"dead_semivoiced_sound":       XK_dead_semivoiced_sound,
	"dead_belowdot":               XK_dead_belowdot,
	"dead_hook":                   XK_dead_hook,
	"dead_horn":                   XK_dead_horn,
	"dead_stroke":                 XK_dead_stroke,
	"dead_abovecomma":             XK_dead_abovecomma,
	"dead_psili":                  XK_dead_psili,
	"dead_abovereversedcomma":     XK_dead_abovereversedcomma,
	"dead_dasia":                  XK_dead_dasia,
	"dead_doublegrave":            XK_dead_doublegrave,
	"dead_belowring":              XK_dead_belowring,
	"dead_belowmacron":            XK_dead_belowmacron,
	"dead_belowcircumflex":        XK_dead_belowcircumflex,
	"dead_belowtilde":             XK_dead_belowtilde,
	"dead_belowbreve":             XK_dead_belowbreve,
	"dead_belowdiaeresis":         XK_dead_belowdiaeresis,
	"dead_invertedbreve":          XK_dead_invertedbreve,
	"dead_belowcomma":             XK_dead_belowcomma,
	"dead_currency":               XK_dead_currency,
	"dead_lowline":                XK_dead_lowline,
	"dead_aboveverticalline":      XK_dead_aboveverticalline,
	"dead_belowverticalline":      XK_dead_belowverticalline,
	"dead_longsolidusoverlay":     XK_dead_longsolidusoverlay,
	"dead_a":                      XK_dead_a,
	"dead_A":                      XK_dead_A,
	"dead_e":                      XK_dead_e,
	"dead_E":                      XK_dead_E,
	"dead_i":                      XK_dead_i,
	"dead_I":                      XK_dead_I,
	"dead_o":                      XK_dead_o,
	"dead_O":                      XK_dead_O,
	"dead_u":                      XK_dead_u,
	"dead_U":                      XK_dead_U,
	"dead_small_schwa":            XK_dead_small_schwa,
	"dead_capital_schwa":          XK_dead_capital_schwa,
	"dead_greek":                  XK_dead_greek,
	"First_Virtual_Screen":        XK_First_Virtual_Screen,
	"Prev_Virtual_Screen":         XK_Prev_Virtual_Screen,
	"Next_Virtual_Screen":         XK_Next_Virtual_Screen,
	"Last_Virtual_Screen":         XK_Last_Virtual_Screen,
	"Terminate_Server":            XK_Terminate_Server,
	"AccessX_Enable":              XK_AccessX_Enable,
	"AccessX_Feedback_Enable":     XK_AccessX_Feedback_Enable,
	"RepeatKeys_Enable":           XK_RepeatKeys_Enable,
	"SlowKeys_Enable":             XK_SlowKeys_Enable,
	"BounceKeys_Enable":           XK_BounceKeys_Enable,
	"StickyKeys_Enable":           XK_StickyKeys_Enable,
	"MouseKeys_Enable":            XK_MouseKeys_Enable,
	"MouseKeys_Accel_Enable":      XK_MouseKeys_Accel_Enable,
	"Overlay1_Enable":             XK_Overlay1_Enable,
	"Overlay2_Enable":             XK_Overlay2_Enable,
	"AudibleBell_Enable":          XK_AudibleBell_Enable,
	"Pointer_Left":                XK_Pointer_Left,
	"Pointer_Right":               XK_Pointer_Right,
	"Pointer_Up":                  XK_Pointer_Up,
	"Pointer_Down":                XK_Pointer_Down,
	"Pointer_UpLeft":              XK_Pointer_UpLeft,
	"Pointer_UpRight":             XK_Pointer_UpRight,
	"Pointer_DownLeft":            XK_Pointer_DownLeft,
	"Pointer_DownRight":           XK_Pointer_DownRight,
	"Pointer_Button_Dflt":         XK_Pointer_Button_Dflt,
	"Pointer_Button1":             XK_Pointer_Button1,
	"Pointer_Button2":             XK_Pointer_Button2,
	"Pointer_Button3":             XK_Pointer_Button3,
	"Pointer_Button4":             XK_Pointer_Button4,
	"Pointer_Button5":             XK_Pointer_Button5,
	"Pointer_DblClick_Dflt":       XK_Pointer_DblClick_Dflt,
	"Pointer_DblClick1":           XK_Pointer_DblClick1,
	"Pointer_DblClick2":           XK_Pointer_DblClick2,
	"Pointer_DblClick3":           XK_Pointer_DblClick3,
	"Pointer_DblClick4":           XK_Pointer_DblClick4,
	"Pointer_DblClick5":           XK_Pointer_DblClick5,
	"Pointer_Drag_Dflt":           XK_Pointer_Drag_Dflt,
	"Pointer_Drag1":               XK_Pointer_Drag1,
	"Pointer_Drag2":               XK_Pointer_Drag2,
	"Pointer_Drag3":               XK_Pointer_Drag3,
	"Pointer_Drag4":               XK_Pointer_Drag4,
	"Pointer_Drag5":               XK_Pointer_Drag5,
	"Pointer_EnableKeys":          XK_Pointer_EnableKeys,
	"Pointer_Accelerate":          XK_Pointer_Accelerate,
	"Pointer_DfltBtnNext":         XK_Pointer_DfltBtnNext,
	"Pointer_DfltBtnPrev":         XK_Pointer_DfltBtnPrev,
	"ch":                          XK_ch,
	"Ch":                          XK_Ch,
	"CH":                          XK_CH,
	"c_h":                         XK_c_h,
	"C_h":                         XK_C_h,
	"C_H":                         XK_C_H,
	"space":                       XK_space,
	"exclam":                      XK_exclam,
	"quotedbl":                    XK_quotedbl,
	"numbersign":                  XK_numbersign,
	"dollar":                      XK_dollar,
	"percent":                     XK_percent,
	"ampersand":                   XK_ampersand,
	"apostrophe":                  XK_apostrophe,
	"quoteright":                  XK_quoteright,
	"parenleft":                   XK_parenleft,
	"parenright":                  XK_parenright,
	"asterisk":                    XK_asterisk,
	"plus":                        XK_plus,
	"comma":                       XK_comma,
	"minus":                       XK_minus,
	"period":                      XK_period,
	"slash":                       XK_slash,
	"0":                           XK_0,
	"1":                           XK_1,
	"2":                           XK_2,
	"3":                           XK_3,
	"4":                           XK_4,
	"5":                           XK_5,
	"6":                           XK_6,
	"7":                           XK_7,
	"8":                           XK_8,
	"9":                           XK_9,
	"colon":                       XK_colon,
	"semicolon":                   XK_semicolon,
	"less":                        XK_less,
	"equal":                       XK_equal,
	"greater":                     XK_greater,
	"question":                    XK_question,
	"at":                          XK_at,
	"A":                           XK_A,
	"B":                           XK_B,
	"C":                           XK_C,
	"D":                           XK_D,
	"E":                           XK_E,
	"F":                           XK_F,
	"G":                           XK_G,
	"H":                           XK_H,
	"I":                           XK_I,
	"J":                           XK_J,
	"K":                           XK_K,
	"L":                           XK_L,
	"M":                           XK_M,
	"N":                           XK_N,
	"O":                           XK_O,
	"P":                           XK_P,
	"Q":                           XK_Q,
	"R":                           XK_R,
	"S":                           XK_S,
	"T":                           XK_T,
	"U":                           XK_U,
	"V":                           XK_V,
	"W":                           XK_W,
	"X":                           XK_X,
	"Y":                           XK_Y,
	"Z":                           XK_Z,
	"bracketleft":                 XK_bracketleft,
	"backslash":                   XK_backslash,
	"bracketright":                XK_bracketright,
	"asciicircum":                 XK_asciicircum,
	"underscore":                  XK_underscore,
	"grave":                       XK_grave,
	"quoteleft":                   XK_quoteleft,
	"a":                           XK_a,
	"b":                           XK_b,
	"c":                           XK_c,
	"d":                           XK_d,
	"e":                           XK_e,
	"f":                           XK_f,
	"g":                           XK_g,
	"h":                           XK_h,
	"i":                           XK_i,
	"j":                           XK_j,
	"k":                           XK_k,
	"l":                           XK_l,
	"m":                           XK_m,
	"n":                           XK_n,
	"o":                           XK_o,
	"p":                           XK_p,
	"q":                           XK_q,
	"r":                           XK_r,
	"s":                           XK_s,
	"t":                           XK_t,
	"u":                           XK_u,
	"v":                           XK_v,
	"w":                           XK_w,
	"x":                           XK_x,
	"y":                           XK_y,
	"z":                           XK_z,
	"braceleft":                   XK_braceleft,
	"bar":                         XK_bar,
	"braceright":                  XK_braceright,
	"asciitilde":                  XK_asciitilde,
	"nobreakspace":                XK_nobreakspace,
	"exclamdown":                  XK_exclamdown,
	"cent":                        XK_cent,
	"sterling":                    XK_sterling,
	"currency":                    XK_currency,
	"yen":                         XK_yen,
	"brokenbar":                   XK_brokenbar,
	"section":                     XK_section,
	"diaeresis":                   XK_diaeresis,
	"copyright":                   XK_copyright,
	"ordfeminine":                 XK_ordfeminine,
	"guillemotleft":               XK_guillemotleft,
	"notsign":                     XK_notsign,
	"hyphen":                      XK_hyphen,
	"registered":                  XK_registered,
	"macron":                      XK_macron,
	"degree":                      XK_degree,
	"plusminus":                   XK_plusminus,
	"twosuperior":                 XK_twosuperior,
	"threesuperior":               XK_threesuperior,
	"acute":                       XK_acute,
	"mu":                          XK_mu,
	"paragraph":                   XK_paragraph,
	"periodcentered":              XK_periodcentered,
	"cedilla":                     XK_cedilla,
	"onesuperior":                 XK_onesuperior,
	"masculine":                   XK_masculine,
	"guillemotright":              XK_guillemotright,
	"onequarter":                  XK_onequarter,
	"onehalf":                     XK_onehalf,
	"threequarters":               XK_threequarters,
	"questiondown":                XK_questiondown,
	"Agrave":                      XK_Agrave,
	"Aacute":                      XK_Aacute,
	"Acircumflex":                 XK_Acircumflex,
	"Atilde":                      XK_Atilde,
	"Adiaeresis":                  XK_Adiaeresis,
	"Aring":                       XK_Aring,
	"AE":                          XK_AE,
	"Ccedilla":                    XK_Ccedilla,
	"Egrave":                      XK_Egrave,
	"Eacute":                      XK_Eacute,
	"Ecircumflex":                 XK_Ecircumflex,
	"Ediaeresis":                  XK_Ediaeresis,
	"Igrave":                      XK_Igrave,
	"Iacute":                      XK_Iacute,
	"Icircumflex":                 XK_Icircumflex,
	"Idiaeresis":                  XK_Idiaeresis,
	"ETH":                         XK_ETH,
	"Eth":                         XK_Eth,
	"Ntilde":                      XK_Ntilde,
	"Ograve":                      XK_Ograve,
	"Oacute":                      XK_Oacute,
	"Ocircumflex":                 XK_Ocircumflex,
	"Otilde":                      XK_Otilde,
	"Odiaeresis":                  XK_Odiaeresis,
	"multiply":                    XK_multiply,
	"Oslash":                      XK_Oslash,
	"Ooblique":                    XK_Ooblique,
	"Ugrave":                      XK_Ugrave,
	"Uacute":                      XK_Uacute,
	"Ucircumflex":                 XK_Ucircumflex,
	"Udiaeresis":                  XK_Udiaeresis,
	"Yacute":                      XK_Yacute,
	"THORN":                       XK_THORN,
	"Thorn":                       XK_Thorn,
	"ssharp":                      XK_ssharp,
	"agrave":                      XK_agrave,
	"aacute":                      XK_aacute,
	"acircumflex":                 XK_acircumflex,
	"atilde":                      XK_atilde,
	"adiaeresis":                  XK_adiaeresis,
	"aring":                       XK_aring,
	"ae":                          XK_ae,
	"ccedilla":                    XK_ccedilla,
	"egrave":                      XK_egrave,
	"eacute":                      XK_eacute,
	"ecircumflex":                 XK_ecircumflex,
	"ediaeresis":                  XK_ediaeresis,
	"igrave":                      XK_igrave,
	"iacute":                      XK_iacute,
	"icircumflex":                 XK_icircumflex,
	"idiaeresis":                  XK_idiaeresis,
	"eth":                         XK_eth,
	"ntilde":                      XK_ntilde,
	"ograve":                      XK_ograve,
	"oacute":                      XK_oacute,
	"ocircumflex":                 XK_ocircumflex,
	"otilde":                      XK_otilde,
	"odiaeresis":                  XK_odiaeresis,
	"division":                    XK_division,
	"oslash":                      XK_oslash,
	"ooblique":                    XK_ooblique,
	"ugrave":                      XK_ugrave,
	"uacute":                      XK_uacute,
	"ucircumflex":                 XK_ucircumflex,
	"udiaeresis":                  XK_udiaeresis,
	"yacute":                      XK_yacute,
	"thorn":                       XK_thorn,
	"ydiaeresis":                  XK_ydiaeresis,
	"Aogonek":                     XK_Aogonek,
	"breve":                       XK_breve,
	"Lstroke":                     XK_Lstroke,
	"Lcaron":                      XK_Lcaron,
	"Sacute":                      XK_Sacute,
	"Scaron":                      XK_Scaron,
	"Scedilla":                    XK_Scedilla,
	"Tcaron":                      XK_Tcaron,
	"Zacute":                      XK_Zacute,
	"Zcaron":                      XK_Zcaron,
	"Zabovedot":                   XK_Zabovedot,
	"aogonek":                     XK_aogonek,
	"ogonek":                      XK_ogonek,
	"lstroke":                     XK_lstroke,
	"lcaron":                      XK_lcaron,
	"sacute":                      XK_sacute,
	"caron":                       XK_caron,
	"scaron":                      XK_scaron,
	"scedilla":                    XK_scedilla,
	"tcaron":                      XK_tcaron,
	"zacute":                      XK_zacute,
	"doubleacute":                 XK_doubleacute,
	"zcaron":                      XK_zcaron,
	"zabovedot":                   XK_zabovedot,
	"Racute":                      XK_Racute,
	"Abreve":                      XK_Abreve,
	"Lacute":                      XK_Lacute,
	"Cacute":                      XK_Cacute,
	"Ccaron":                      XK_Ccaron,
	"Eogonek":                     XK_Eogonek,
	"Ecaron":                      XK_Ecaron,
	"Dcaron":                      XK_Dcaron,
	"Dstroke":                     XK_Dstroke,
	"Nacute":                      XK_Nacute,
	"Ncaron":                      XK_Ncaron,
	"Odoubleacute":                XK_Odoubleacute,
	"Rcaron":                      XK_Rcaron,
	"Uring":                       XK_Uring,
	"Udoubleacute":                XK_Udoubleacute,
	"Tcedilla":                    XK_Tcedilla,
	"racute":                      XK_racute,
	"abreve":                      XK_abreve,
	"lacute":                      XK_lacute,
	"cacute":                      XK_cacute,
	"ccaron":                      XK_ccaron,
	"eogonek":                     XK_eogonek,
	"ecaron":                      XK_ecaron,
	"dcaron":                      XK_dcaron,
	"dstroke":                     XK_dstroke,
	"nacute":                      XK_nacute,
	"ncaron":                      XK_ncaron,
	"odoubleacute":                XK_odoubleacute,
	"rcaron":                      XK_rcaron,
	"uring":                       XK_uring,
	"udoubleacute":                XK_udoubleacute,
	"tcedilla":                    XK_tcedilla,
	"abovedot":                    XK_abovedot,
	"Hstroke":                     XK_Hstroke,
	"Hcircumflex":                 XK_Hcircumflex,
	"Iabovedot":                   XK_Iabovedot,
	"Gbreve":                      XK_Gbreve,
	"Jcircumflex":                 XK_Jcircumflex,
	"hstroke":                     XK_hstroke,
	"hcircumflex":                 XK_hcircumflex,
	"idotless":                    XK_idotless,
	"gbreve":                      XK_gbreve,
	"jcircumflex":                 XK_jcircumflex,
	"Cabovedot":                   XK_Cabovedot,
	"Ccircumflex":                 XK_Ccircumflex,
	"Gabovedot":                   XK_Gabovedot,
	"Gcircumflex":                 XK_Gcircumflex,
	"Ubreve":                      XK_Ubreve,
	"Scircumflex":                 XK_Scircumflex,
	"cabovedot":                   XK_cabovedot,
	"ccircumflex":                 XK_ccircumflex,
	"gabovedot":                   XK_gabovedot,
	"gcircumflex":                 XK_gcircumflex,
	"ubreve":                      XK_ubreve,
	"scircumflex":                 XK_scircumflex,
	"kra":                         XK_kra,
	"kappa":                       XK_kappa,
	"Rcedilla":                    XK_Rcedilla,
	"Itilde":                      XK_Itilde,
	"Lcedilla":                    XK_Lcedilla,
	"Emacron":                     XK_Emacron,
	"Gcedilla":                    XK_Gcedilla,
	"Tslash":                      XK_Tslash,
	"rcedilla":                    XK_rcedilla,
	"itilde":                      XK_itilde,
	"lcedilla":                    XK_lcedilla,
	"emacron":                     XK_emacron,
	"gcedilla":                    XK_gcedilla,
	"tslash":                      XK_tslash,
	"ENG":                         XK_ENG,
	"eng":                         XK_eng,
	"Amacron":                     XK_Amacron,
	"Iogonek":                     XK_Iogonek,
	"Eabovedot":                   XK_Eabovedot,
	"Imacron":                     XK_Imacron,
	"Ncedilla":                    XK_Ncedilla,
	"Omacron":                     XK_Omacron,
	"Kcedilla":                    XK_Kcedilla,
	"Uogonek":                     XK_Uogonek,
	"Utilde":                      XK_Utilde,
	"Umacron":                     XK_Umacron,
	"amacron":                     XK_amacron,
	"iogonek":                     XK_iogonek,
	"eabovedot":                   XK_eabovedot,
	"imacron":                     XK_imacron,
	"ncedilla":                    XK_ncedilla,
	"omacron":                     XK_omacron,
	"kcedilla":                    XK_kcedilla,
	"uogonek":                     XK_uogonek,
	"utilde":                      XK_utilde,
	"umacron":                     XK_umacron,
	"Wcircumflex":                 XK_Wcircumflex,
	"wcircumflex":                 XK_wcircumflex,
	"Ycircumflex":                 XK_Ycircumflex,
	"ycircumflex":                 XK_ycircumflex,
	"Babovedot":                   XK_Babovedot,
	"babovedot":                   XK_babovedot,
	"Dabovedot":                   XK_Dabovedot,
	"dabovedot":                   XK_dabovedot,
	"Fabovedot":                   XK_Fabovedot,
	"fabovedot":                   XK_fabovedot,
	"Mabovedot":                   XK_Mabovedot,
	"mabovedot":                   XK_mabovedot,
	"Pabovedot":                   XK_Pabovedot,
	"pabovedot":                   XK_pabovedot,
	"Sabovedot":                   XK_Sabovedot,
	"sabovedot":                   XK_sabovedot,
	"Tabovedot":                   XK_Tabovedot,
	"tabovedot":                   XK_tabovedot,
	"Wgrave":                      XK_Wgrave,
	"wgrave":                      XK_wgrave,
	"Wacute":                      XK_Wacute,
	"wacute":                      XK_wacute,
	"Wdiaeresis":                  XK_Wdiaeresis,
	"wdiaeresis":                  XK_wdiaeresis,
	"Ygrave":                      XK_Ygrave,
	"ygrave":                      XK_ygrave,
	"OE":                          XK_OE,
	"oe":                          XK_oe,
	"Ydiaeresis":                  XK_Ydiaeresis,
	"overline":                    XK_overline,
	"kana_fullstop":               XK_kana_fullstop,
	"kana_openingbracket":         XK_kana_openingbracket,
	"kana_closingbracket":         XK_kana_closingbracket,
	"kana_comma":                  XK_kana_comma,
	"kana_conjunctive":            XK_kana_conjunctive,
	"kana_middledot":              XK_kana_middledot,
	"kana_WO":                     XK_kana_WO,
	"kana_a":                      XK_kana_a,
	"kana_i":                      XK_kana_i,
	"kana_u":                      XK_kana_u,
	"kana_e":                      XK_kana_e,
	"kana_o":                      XK_kana_o,
	"kana_ya":                     XK_kana_ya,
	"kana_yu":                     XK_kana_yu,
	"kana_yo":                     XK_kana_yo,
	"kana_tsu":                    XK_kana_tsu,
	"kana_tu":                     XK_kana_tu,
	"prolongedsound":              XK_prolongedsound,
	"kana_A":                      XK_kana_A,
	"kana_I":                      XK_kana_I,
	"kana_U":                      XK_kana_U,
	"kana_E":                      XK_kana_E,
	"kana_O":                      XK_kana_O,
	"kana_KA":                     XK_kana_KA,
	"kana_KI":                     XK_kana_KI,
	"kana_KU":                     XK_kana_KU,
	"kana_KE":                     XK_kana_KE,
	"kana_KO":                     XK_kana_KO,
	"kana_SA":                     XK_kana_SA,
	"kana_SHI":                    XK_kana_SHI,
	"kana_SU":                     XK_kana_SU,
	"kana_SE":                     XK_kana_SE,
	"kana_SO":                     XK_kana_SO,
	"kana_TA":                     XK_kana_TA,
	"kana_CHI":                    XK_kana_CHI,
	"kana_TI":                     XK_kana_TI,
	"kana_TSU":                    XK_kana_TSU,
	"kana_TU":                     XK_kana_TU,
	"kana_TE":                     XK_kana_TE,
	"kana_TO":                     XK_kana_TO,
	"kana_NA":                     XK_kana_NA,
	"kana_NI":                     XK_kana_NI,
	"kana_NU":                     XK_kana_NU,
	"kana_NE":                     XK_kana_NE,
	"kana_NO":                     XK_kana_NO,
	"kana_HA":                     XK_kana_HA,
	"kana_HI":                     XK_kana_HI,
	"kana_FU":                     XK_kana_FU,
	"kana_HU":                     XK_kana_HU,
	"kana_HE":                     XK_kana_HE,
	"kana_HO":                     XK_kana_HO,
	"kana_MA":                     XK_kana_MA,
	"kana_MI":                     XK_kana_MI,
	"kana_MU":                     XK_kana_MU,
	"kana_ME":                     XK_kana_ME,
	"kana_MO":                     XK_kana_MO,
	"kana_YA":                     XK_kana_YA,
	"kana_YU":                     XK_kana_YU,
	"kana_YO":                     XK_kana_YO,
	"kana_RA":                     XK_kana_RA,
	"kana_RI":                     XK_kana_RI,
	"kana_RU":                     XK_kana_RU,
	"kana_RE":                     XK_kana_RE,
	"kana_RO":                     XK_kana_RO,
	"kana_WA":                     XK_kana_WA,
	"kana_N":                      XK_kana_N,
	"voicedsound":                 XK_voicedsound,
	"semivoicedsound":             XK_semivoicedsound,
	"kana_switch":                 XK_kana_switch,
	"Farsi_0":                     XK_Farsi_0,
	"Farsi_1":                     XK_Farsi_1,
	"Farsi_2":                     XK_Farsi_2,
	"Farsi_3":                     XK_Farsi_3,
	"Farsi_4":                     XK_Farsi_4,
	"Farsi_5":                     XK_Farsi_5,
	"Farsi_6":                     XK_Farsi_6,
	"Farsi_7":                     XK_Farsi_7,
	"Farsi_8":                     XK_Farsi_8,
	"Farsi_9":                     XK_Farsi_9,
	"Arabic_percent":              XK_Arabic_percent,
	"Arabic_superscript_alef":     XK_Arabic_superscript_alef,
	"Arabic_tteh":                 XK_Arabic_tteh,
	"Arabic_peh":                  XK_Arabic_peh,
	"Arabic_tcheh":                XK_Arabic_tcheh,
	"Arabic_ddal":                 XK_Arabic_ddal,
	"Arabic_rreh":                 XK_Arabic_rreh,
	"Arabic_comma":                XK_Arabic_comma,
	"Arabic_fullstop":             XK_Arabic_fullstop,
	"Arabic_0":                    XK_Arabic_0,
	"Arabic_1":                    XK_Arabic_1,
	"Arabic_2":                    XK_Arabic_2,
	"Arabic_3":                    XK_Arabic_3,
	"Arabic_4":                    XK_Arabic_4,
	"Arabic_5":                    XK_Arabic_5,
	"Arabic_6":                    XK_Arabic_6,
	"Arabic_7":                    XK_Arabic_7,
	"Arabic_8":                    XK_Arabic_8,
	"Arabic_9":                    XK_Arabic_9,
	"Arabic_semicolon":            XK_Arabic_semicolon,
	"Arabic_question_mark":        XK_Arabic_question_mark,
	"Arabic_hamza":                XK_Arabic_hamza,
	"Arabic_maddaonalef":          XK_Arabic_maddaonalef,
	"Arabic_hamzaonalef":          XK_Arabic_hamzaonalef,
	"Arabic_hamzaonwaw":           XK_Arabic_hamzaonwaw,
	"Arabic_hamzaunderalef":       XK_Arabic_hamzaunderalef,
	"Arabic_hamzaonyeh":           XK_Arabic_hamzaonyeh,
	"Arabic_alef":                 XK_Arabic_alef,
	"Arabic_beh":                  XK_Arabic_beh,
	"Arabic_tehmarbuta":           XK_Arabic_tehmarbuta,
	"Arabic_teh":                  XK_Arabic_teh,
	"Arabic_theh":                 XK_Arabic_theh,
	"Arabic_jeem":                 XK_Arabic_jeem,
	"Arabic_hah":                  XK_Arabic_hah,
	"Arabic_khah":                 XK_Arabic_khah,
	"Arabic_dal":                  XK_Arabic_dal,
	"Arabic_thal":                 XK_Arabic_thal,
	"Arabic_ra":                   XK_Arabic_ra,
	"Arabic_zain":                 XK_Arabic_zain,
	"Arabic_seen":                 XK_Arabic_seen,
	"Arabic_sheen":                XK_Arabic_sheen,
	"Arabic_sad":                  XK_Arabic_sad,
	"Arabic_dad":                  XK_Arabic_dad,
	"Arabic_tah":                  XK_Arabic_tah,
	"Arabic_zah":                  XK_Arabic_zah,
	"Arabic_ain":                  XK_Arabic_ain,
	"Arabic_ghain":                XK_Arabic_ghain,
	"Arabic_tatweel":              XK_Arabic_tatweel,
	"Arabic_feh":                  XK_Arabic_feh,
	"Arabic_qaf":                  XK_Arabic_qaf,
	"Arabic_kaf":                  XK_Arabic_kaf,
	"Arabic_lam":                  XK_Arabic_lam,
	"Arabic_meem":                 XK_Arabic_meem,
	"Arabic_noon":                 XK_Arabic_noon,
	"Arabic_ha":                   XK_Arabic_ha,
	"Arabic_heh":                  XK_Arabic_heh,
	"Arabic_waw":                  XK_Arabic_waw,
	"Arabic_alefmaksura":          XK_Arabic_alefmaksura,
	"Arabic_yeh":                  XK_Arabic_yeh,
	"Arabic_fathatan":             XK_Arabic_fathatan,
	"Arabic_dammatan":             XK_Arabic_dammatan,
	"Arabic_kasratan":             XK_Arabic_kasratan,
	"Arabic_fatha":                XK_Arabic_fatha,
	"Arabic_damma":                XK_Arabic_damma,
	"Arabic_kasra":                XK_Arabic_kasra,
	"Arabic_shadda":               XK_Arabic_shadda,
	"Arabic_sukun":                XK_Arabic_sukun,
	"Arabic_madda_above":          XK_Arabic_madda_above,
	"Arabic_hamza_above":          XK_Arabic_hamza_above,
	"Arabic_hamza_below":          XK_Arabic_hamza_below,
	"Arabic_jeh":                  XK_Arabic_jeh,
	"Arabic_veh":                  XK_Arabic_veh,
	"Arabic_keheh":                XK_Arabic_keheh,
	"Arabic_gaf":                  XK_Arabic_gaf,
	"Arabic_noon_ghunna":          XK_Arabic_noon_ghunna,
	"Arabic_heh_doachashmee":      XK_Arabic_heh_doachashmee,
	"Farsi_yeh":                   XK_Farsi_yeh,
	"Arabic_farsi_yeh":            XK_Arabic_farsi_yeh,
	"Arabic_yeh_baree":            XK_Arabic_yeh_baree,
	"Arabic_heh_goal":             XK_Arabic_heh_goal,
	"Arabic_switch":               XK_Arabic_switch,
	"Cyrillic_GHE_bar":            XK_Cyrillic_GHE_bar,
	"Cyrillic_ghe_bar":            XK_Cyrillic_ghe_bar,
	"Cyrillic_ZHE_descender":      XK_Cyrillic_ZHE_descender,
	"Cyrillic_zhe_descender":      XK_Cyrillic_zhe_descender,
	"Cyrillic_KA_descender":       XK_Cyrillic_KA_descender,
	"Cyrillic_ka_descender":       XK_Cyrillic_ka_descender,
	"Cyrillic_KA_vertstroke":      XK_Cyrillic_KA_vertstroke,
	"Cyrillic_ka_vertstroke":      XK_Cyrillic_ka_vertstroke,
	"Cyrillic_EN_descender":       XK_Cyrillic_EN_descender,
	"Cyrillic_en_descender":       XK_Cyrillic_en_descender,
	"Cyrillic_U_straight":         XK_Cyrillic_U_straight,
	"Cyrillic_u_straight":         XK_Cyrillic_u_straight,
	"Cyrillic_U_straight_bar":     XK_Cyrillic_U_straight_bar,
	"Cyrillic_u_straight_bar":     XK_Cyrillic_u_straight_bar,
	"Cyrillic_HA_descender":       XK_Cyrillic_HA_descender,
	"Cyrillic_ha_descender":       XK_Cyrillic_ha_descender,
	"Cyrillic_CHE_descender":      XK_Cyrillic_CHE_descender,
	"Cyrillic_che_descender":      XK_Cyrillic_che_descender,
	"Cyrillic_CHE_vertstroke":     XK_Cyrillic_CHE_vertstroke,
	"Cyrillic_che_vertstroke":     XK_Cyrillic_che_vertstroke,
	"Cyrillic_SHHA":               XK_Cyrillic_SHHA,
	"Cyrillic_shha":               XK_Cyrillic_shha,
	"Cyrillic_SCHWA":              XK_Cyrillic_SCHWA,
	"Cyrillic_schwa":              XK_Cyrillic_schwa,
	"Cyrillic_I_macron":           XK_Cyrillic_I_macron,
	"Cyrillic_i_macron":           XK_Cyrillic_i_macron,
	"Cyrillic_O_bar":              XK_Cyrillic_O_bar,
	"Cyrillic_o_bar":              XK_Cyrillic_o_bar,
	"Cyrillic_U_macron":           XK_Cyrillic_U_macron,
	"Cyrillic_u_macron":           XK_Cyrillic_u_macron,
	"Serbian_dje":                 XK_Serbian_dje,
	"Macedonia_gje":               XK_Macedonia_gje,
	"Cyrillic_io":                 XK_Cyrillic_io,
	"Ukrainian_ie":                XK_Ukrainian_ie,
	"Ukranian_je":                 XK_Ukranian_je,
	"Macedonia_dse":               XK_Macedonia_dse,
	"Ukrainian_i":                 XK_Ukrainian_i,
	"Ukranian_i":                  XK_Ukranian_i,
	"Ukrainian_yi":                XK_Ukrainian_yi,
	"Ukranian_yi":                 XK_Ukranian_yi,
	"Cyrillic_je":                 XK_Cyrillic_je,
	"Serbian_je":                  XK_Serbian_je,
	"Cyrillic_lje":                XK_Cyrillic_lje,
	"Serbian_lje":                 XK_Serbian_lje,
	"Cyrillic_nje":                XK_Cyrillic_nje,
	"Serbian_nje":                 XK_Serbian_nje,
	"Serbian_tshe":                XK_Serbian_tshe,
	"Macedonia_kje":               XK_Macedonia_kje,
	"Ukrainian_ghe_with_upturn":   XK_Ukrainian_ghe_with_upturn,
	"Byelorussian_shortu":         XK_Byelorussian_shortu,
	"Cyrillic_dzhe":               XK_Cyrillic_dzhe,
	"Serbian_dze":                 XK_Serbian_dze,
	"numerosign":                  XK_numerosign,
	"Serbian_DJE":                 XK_Serbian_DJE,
	"Macedonia_GJE":               XK_Macedonia_GJE,
	"Cyrillic_IO":                 XK_Cyrillic_IO,
	"Ukrainian_IE":                XK_Ukrainian_IE,
	"Ukranian_JE":                 XK_Ukranian_JE,
	"Macedonia_DSE":               XK_Macedonia_DSE,
	"Ukrainian_I":                 XK_Ukrainian_I,
	"Ukranian_I":                  XK_Ukranian_I,
	"Ukrainian_YI":                XK_Ukrainian_YI,
	"Ukranian_YI":                 XK_Ukranian_YI,
	"Cyrillic_JE":                 XK_Cyrillic_JE,
	"Serbian_JE":                  XK_Serbian_JE,
	"Cyrillic_LJE":                XK_Cyrillic_LJE,
	"Serbian_LJE":                 XK_Serbian_LJE,
	"Cyrillic_NJE":                XK_Cyrillic_NJE,
	"Serbian_NJE":                 XK_Serbian_NJE,
	"Serbian_TSHE":                XK_Serbian_TSHE,
	"Macedonia_KJE":               XK_Macedonia_KJE,
	"Ukrainian_GHE_WITH_UPTURN":   XK_Ukrainian_GHE_WITH_UPTURN,
	"Byelorussian_SHORTU":         XK_Byelorussian_SHORTU,
	"Cyrillic_DZHE":               XK_Cyrillic_DZHE,
	"Serbian_DZE":                 XK_Serbian_DZE,
	"Cyrillic_yu":                 XK_Cyrillic_yu,
	"Cyrillic_a":                  XK_Cyrillic_a,
	"Cyrillic_be":                 XK_Cyrillic_be,
	"Cyrillic_tse":                XK_Cyrillic_tse,
	"Cyrillic_de":                 XK_Cyrillic_de,
	"Cyrillic_ie":                 XK_Cyrillic_ie,
	"Cyrillic_ef":                 XK_Cyrillic_ef,
	"Cyrillic_ghe":                XK_Cyrillic_ghe,
	"Cyrillic_ha":                 XK_Cyrillic_ha,
	"Cyrillic_i":                  XK_Cyrillic_i,
	"Cyrillic_shorti":             XK_Cyrillic_shorti,
	"Cyrillic_ka":                 XK_Cyrillic_ka,
	"Cyrillic_el":                 XK_Cyrillic_el,
	"Cyrillic_em":                 XK_Cyrillic_em,
	"Cyrillic_en":                 XK_Cyrillic_en,
	"Cyrillic_o":                  XK_Cyrillic_o,
	"Cyrillic_pe":                 XK_Cyrillic_pe,
	"Cyrillic_ya":                 XK_Cyrillic_ya,
	"Cyrillic_er":                 XK_Cyrillic_er,
	"Cyrillic_es":                 XK_Cyrillic_es,
	"Cyrillic_te":                 XK_Cyrillic_te,
	"Cyrillic_u":                  XK_Cyrillic_u,
	"Cyrillic_zhe":                XK_Cyrillic_zhe,
	"Cyrillic_ve":                 XK_Cyrillic_ve,
	"Cyrillic_softsign":           XK_Cyrillic_softsign,
	"Cyrillic_yeru":               XK_Cyrillic_yeru,
	"Cyrillic_ze":                 XK_Cyrillic_ze,
	"Cyrillic_sha":                XK_Cyrillic_sha,
	"Cyrillic_e":                  XK_Cyrillic_e,
	"Cyrillic_shcha":              XK_Cyrillic_shcha,
	"Cyrillic_che":                XK_Cyrillic_che,
	"Cyrillic_hardsign":           XK_Cyrillic_hardsign,
	"Cyrillic_YU":                 XK_Cyrillic_YU,
	"Cyrillic_A":                  XK_Cyrillic_A,
	"Cyrillic_BE":                 XK_Cyrillic_BE,
	"Cyrillic_TSE":                XK_Cyrillic_TSE,
	"Cyrillic_DE":                 XK_Cyrillic_DE,
	"Cyrillic_IE":                 XK_Cyrillic_IE,
	"Cyrillic_EF":                 XK_Cyrillic_EF,
	"Cyrillic_GHE":                XK_Cyrillic_GHE,
	"Cyrillic_HA":                 XK_Cyrillic_HA,
	"Cyrillic_I":                  XK_Cyrillic_I,
	"Cyrillic_SHORTI":             XK_Cyrillic_SHORTI,
	"Cyrillic_KA":                 XK_Cyrillic_KA,
	"Cyrillic_EL":                 XK_Cyrillic_EL,
	"Cyrillic_EM":                 XK_Cyrillic_EM,
	"Cyrillic_EN":                 XK_Cyrillic_EN,
	"Cyrillic_O":                  XK_Cyrillic_O,
	"Cyrillic_PE":                 XK_Cyrillic_PE,
	"Cyrillic_YA":                 XK_Cyrillic_YA,
	"Cyrillic_ER":                 XK_Cyrillic_ER,
	"Cyrillic_ES":                 XK_Cyrillic_ES,
	"Cyrillic_TE":                 XK_Cyrillic_TE,
	"Cyrillic_U":                  XK_Cyrillic_U,
	"Cyrillic_ZHE":                XK_Cyrillic_ZHE,
	"Cyrillic_VE":                 XK_Cyrillic_VE,
	"Cyrillic_SOFTSIGN":           XK_Cyrillic_SOFTSIGN,
	"Cyrillic_YERU":               XK_Cyrillic_YERU,
	"Cyrillic_ZE":                 XK_Cyrillic_ZE,
	"Cyrillic_SHA":                XK_Cyrillic_SHA,
	"Cyrillic_E":                  XK_Cyrillic_E,
	"Cyrillic_SHCHA":              XK_Cyrillic_SHCHA,
	"Cyrillic_CHE":                XK_Cyrillic_CHE,
	"Cyrillic_HARDSIGN":           XK_Cyrillic_HARDSIGN,
	"Greek_ALPHAaccent":           XK_Greek_ALPHAaccent,
	"Greek_EPSILONaccent":         XK_Greek_EPSILONaccent,
	"Greek_ETAaccent":             XK_Greek_ETAaccent,
	"Greek_IOTAaccent":            XK_Greek_IOTAaccent,
	"Greek_IOTAdieresis":          XK_Greek_IOTAdieresis,
	"Greek_IOTAdiaeresis":         XK_Greek_IOTAdiaeresis,
	"Greek_OMICRONaccent":         XK_Greek_OMICRONaccent,
	"Greek_UPSILONaccent":         XK_Greek_UPSILONaccent,
	"Greek_UPSILONdieresis":       XK_Greek_UPSILONdieresis,
	"Greek_OMEGAaccent":           XK_Greek_OMEGAaccent,
	"Greek_accentdieresis":        XK_Greek_accentdieresis,
	"Greek_horizbar":              XK_Greek_horizbar,
	"Greek_alphaaccent":           XK_Greek_alphaaccent,
	"Greek_epsilonaccent":         XK_Greek_epsilonaccent,
	"Greek_etaaccent":             XK_Greek_etaaccent,
	"Greek_iotaaccent":            XK_Greek_iotaaccent,
	"Greek_iotadieresis":          XK_Greek_iotadieresis,
	"Greek_iotaaccentdieresis":    XK_Greek_iotaaccentdieresis,
	"Greek_omicronaccent":         XK_Greek_omicronaccent,
	"Greek_upsilonaccent":         XK_Greek_upsilonaccent,
	"Greek_upsilondieresis":       XK_Greek_upsilondieresis,
	"Greek_upsilonaccentdieresis": XK_Greek_upsilonaccentdieresis,
	"Greek_omegaaccent":           XK_Greek_omegaaccent,
	"Greek_ALPHA":                 XK_Greek_ALPHA,
	"Greek_BETA":                  XK_Greek_BETA,
	"Greek_GAMMA":                 XK_Greek_GAMMA,
	"Greek_DELTA":                 XK_Greek_DELTA,
	"Greek_EPSILON":               XK_Greek_EPSILON,
	"Greek_ZETA":                  XK_Greek_ZETA,
	"Greek_ETA":                   XK_Greek_ETA,
	"Greek_THETA":                 XK_Greek_THETA,
	"Greek_IOTA":                  XK_Greek_IOTA,
	"Greek_KAPPA":                 XK_Greek_KAPPA,
	"Greek_LAMDA":                 XK_Greek_LAMDA,
	"Greek_LAMBDA":                XK_Greek_LAMBDA,
	"Greek_MU":                    XK_Greek_MU,
	"Greek_NU":                    XK_Greek_NU,
	"Greek_XI":                    XK_Greek_XI,
	"Greek_OMICRON":               XK_Greek_OMICRON,
	"Greek_PI":                    XK_Greek_PI,
	"Greek_RHO":                   XK_Greek_RHO,
	"Greek_SIGMA":                 XK_Greek_SIGMA,
	"Greek_TAU":                   XK_Greek_TAU,
	"Greek_UPSILON":               XK_Greek_UPSILON,
	"Greek_PHI":                   XK_Greek_PHI,
	"Greek_CHI":                   XK_Greek_CHI,
	"Greek_PSI":                   XK_Greek_PSI,
	"Greek_OMEGA":                 XK_Greek_OMEGA,
	"Greek_alpha":                 XK_Greek_alpha,
	"Greek_beta":                  XK_Greek_beta,
	"Greek_gamma":                 XK_Greek_gamma,
	"Greek_delta":                 XK_Greek_delta,
	"Greek_epsilon":               XK_Greek_epsilon,
	"Greek_zeta":                  XK_Greek_zeta,
	"Greek_eta":                   XK_Greek_eta,
	"Greek_theta":                 XK_Greek_theta,
	"Greek_iota":                  XK_Greek_iota,
	"Greek_kappa":                 XK_Greek_kappa,
	"Greek_lamda":                 XK_Greek_lamda,
	"Greek_lambda":                XK_Greek_lambda,
	"Greek_mu":                    XK_Greek_mu,
	"Greek_nu":                    XK_Greek_nu,
	"Greek_xi":                    XK_Greek_xi,
	"Greek_omicron":               XK_Greek_omicron,
	"Greek_pi":                    XK_Greek_pi,
	"Greek_rho":                   XK_Greek_rho,
	"Greek_sigma":                 XK_Greek_sigma,
	"Greek_finalsmallsigma":       XK_Greek_finalsmallsigma,
	"Greek_tau":                   XK_Greek_tau,
	"Greek_upsilon":               XK_Greek_upsilon,
	"Greek_phi":                   XK_Greek_phi,
	"Greek_chi":                   XK_Greek_chi,
	"Greek_psi":                   XK_Greek_psi,
	"Greek_omega":                 XK_Greek_omega,
	"Greek_switch":                XK_Greek_switch,
	"hebrew_doublelowline":        XK_hebrew_doublelowline,
	"hebrew_aleph":                XK_hebrew_aleph,
	"hebrew_bet":                  XK_hebrew_bet,
	"hebrew_beth":                 XK_hebrew_beth,
	"hebrew_gimel":                XK_hebrew_gimel,
	"hebrew_gimmel":               XK_hebrew_gimmel,
	"hebrew_dalet":                XK_hebrew_dalet,
	"hebrew_daleth":               XK_hebrew_daleth,
	"hebrew_he":                   XK_hebrew_he,
	"hebrew_waw":                  XK_hebrew_waw,
	"hebrew_zain":                 XK_hebrew_zain,
	"hebrew_zayin":                XK_hebrew_zayin,
	"hebrew_chet":                 XK_hebrew_chet,
	"hebrew_het":                  XK_hebrew_het,
	"hebrew_tet":                  XK_hebrew_tet,
	"hebrew_teth":                 XK_hebrew_teth,
	"hebrew_yod":                  XK_hebrew_yod,
	"hebrew_finalkaph":            XK_hebrew_finalkaph,
	"hebrew_kaph":                 XK_hebrew_kaph,
	"hebrew_lamed":                XK_hebrew_lamed,
	"hebrew_finalmem":             XK_hebrew_finalmem,
	"hebrew_mem":                  XK_hebrew_mem,
	"hebrew_finalnun":             XK_hebrew_finalnun,
	"hebrew_nun":                  XK_hebrew_nun,
	"hebrew_samech":               XK_hebrew_samech,
	"hebrew_samekh":               XK_hebrew_samekh,
	"hebrew_ayin":                 XK_hebrew_ayin,
	"hebrew_finalpe":              XK_hebrew_finalpe,
	"hebrew_pe":                   XK_hebrew_pe,
	"hebrew_finalzade":            XK_hebrew_finalzade,
	"hebrew_finalzadi":            XK_hebrew_finalzadi,
	"hebrew_zade":                 XK_hebrew_zade,
	"hebrew_zadi":                 XK_hebrew_zadi,
	"hebrew_qoph":                 XK_hebrew_qoph,
	"hebrew_kuf":                  XK_hebrew_kuf,
	"hebrew_resh":                 XK_hebrew_resh,
	"hebrew_shin":                 XK_hebrew_shin,
	"hebrew_taw":                  XK_hebrew_taw,
	"hebrew_taf":                  XK_hebrew_taf,
	"Hebrew_switch":               XK_Hebrew_switch,
	"Thai_kokai":                  XK_Thai_kokai,
	"Thai_khokhai":                XK_Thai_khokhai,
	"Thai_khokhuat":               XK_Thai_khokhuat,
	"Thai_khokhwai":               XK_Thai_khokhwai,
	"Thai_khokhon":                XK_Thai_khokhon,
	"Thai_khorakhang":             XK_Thai_khorakhang,
	"Thai_ngongu":                 XK_Thai_ngongu,
	"Thai_chochan":                XK_Thai_chochan,
	"Thai_choching":               XK_Thai_choching,
	"Thai_chochang":               XK_Thai_chochang,
	"Thai_soso":                   XK_Thai_soso,
	"Thai_chochoe":                XK_Thai_chochoe,
	"Thai_yoying":                 XK_Thai_yoying,
	"Thai_dochada":                XK_Thai_dochada,
	"Thai_topatak":                XK_Thai_topatak,
	"Thai_thothan":                XK_Thai_thothan,
	"Thai_thonangmontho":          XK_Thai_thonangmontho,
	"Thai_thophuthao":             XK_Thai_thophuthao,
	"Thai_nonen":                  XK_Thai_nonen,
	"Thai_dodek":                  XK_Thai_dodek,
	"Thai_totao":                  XK_Thai_totao,
	"Thai_thothung":               XK_Thai_thothung,
	"Thai_thothahan":              XK_Thai_thothahan,
	"Thai_thothong":               XK_Thai_thothong,
	"Thai_nonu":                   XK_Thai_nonu,
	"Thai_bobaimai":               XK_Thai_bobaimai,
	"Thai_popla":                  XK_Thai_popla,
	"Thai_phophung":               XK_Thai_phophung,
	"Thai_fofa":                   XK_Thai_fofa,
	"Thai_phophan":                XK_Thai_phophan,
	"Thai_fofan":                  XK_Thai_fofan,
	"Thai_phosamphao":             XK_Thai_phosamphao,
	"Thai_moma":                   XK_Thai_moma,
	"Thai_yoyak":                  XK_Thai_yoyak,
	"Thai_rorua":                  XK_Thai_rorua,
	"Thai_ru":                     XK_Thai_ru,
	"Thai_loling":                 XK_Thai_loling,
	"Thai_lu":                     XK_Thai_lu,
	"Thai_wowaen":                 XK_Thai_wowaen,
	"Thai_sosala":                 XK_Thai_sosala,
	"Thai_sorusi":                 XK_Thai_sorusi,
	"Thai_sosua":                  XK_Thai_sosua,
	"Thai_hohip":                  XK_Thai_hohip,
	"Thai_lochula":                XK_Thai_lochula,
	"Thai_oang":                   XK_Thai_oang,
	"Thai_honokhuk":               XK_Thai_honokhuk,
	"Thai_paiyannoi":              XK_Thai_paiyannoi,
	"Thai_saraa":                  XK_Thai_saraa,
	"Thai_maihanakat":             XK_Thai_maihanakat,
	"Thai_saraaa":                 XK_Thai_saraaa,
	"Thai_saraam":                 XK_Thai_saraam,
	"Thai_sarai":                  XK_Thai_sarai,
	"Thai_saraii":                 XK_Thai_saraii,
	"Thai_saraue":                 XK_Thai_saraue,
	"Thai_sarauee":                XK_Thai_sarauee,
	"Thai_sarau":                  XK_Thai_sarau,
	"Thai_sarauu":                 XK_Thai_sarauu,
	"Thai_phinthu":                XK_Thai_phinthu,
	"Thai_maihanakat_maitho":      XK_Thai_maihanakat_maitho,
	"Thai_baht":                   XK_Thai_baht,
	"Thai_sarae":                  XK_Thai_sarae,
	"Thai_saraae":                 XK_Thai_saraae,
	"Thai_sarao":                  XK_Thai_sarao,
	"Thai_saraaimaimuan":          XK_Thai_saraaimaimuan,
	"Thai_saraaimaimalai":         XK_Thai_saraaimaimalai,
	"Thai_lakkhangyao":            XK_Thai_lakkhangyao,
	"Thai_maiyamok":               XK_Thai_maiyamok,
	"Thai_maitaikhu":              XK_Thai_maitaikhu,
	"Thai_maiek":                  XK_Thai_maiek,
	"Thai_maitho":                 XK_Thai_maitho,
	"Thai_maitri":                 XK_Thai_maitri,
	"Thai_maichattawa":            XK_Thai_maichattawa,
	"Thai_thanthakhat":            XK_Thai_thanthakhat,
	"Thai_nikhahit":               XK_Thai_nikhahit,
	"Thai_leksun":                 XK_Thai_leksun,
	"Thai_leknung":                XK_Thai_leknung,
	"Thai_leksong":                XK_Thai_leksong,
	"Thai_leksam":                 XK_Thai_leksam,
	"Thai_leksi":                  XK_Thai_leksi,
	"Thai_lekha":                  XK_Thai_lekha,
	"Thai_lekhok":                 XK_Thai_lekhok,
	"Thai_lekchet":                XK_Thai_lekchet,
	"Thai_lekpaet":                XK_Thai_lekpaet,
	"Thai_lekkao":                 XK_Thai_lekkao,
	"Hangul":                      XK_Hangul,
	"Hangul_Start":                XK_Hangul_Start,
	"Hangul_End":                  XK_Hangul_End,
	"Hangul_Hanja":                XK_Hangul_Hanja,
	"Hangul_Jamo":                 XK_Hangul_Jamo,
	"Hangul_Romaja":               XK_Hangul_Romaja,
	"Hangul_Codeinput":            XK_Hangul_Codeinput,
	"Hangul_Jeonja":               XK_Hangul_Jeonja,
	"Hangul_Banja":                XK_Hangul_Banja,
	"Hangul_PreHanja":             XK_Hangul_PreHanja,
	"Hangul_PostHanja":            XK_Hangul_PostHanja,
	"Hangul_SingleCandidate":      XK_Hangul_SingleCandidate,
	"Hangul_MultipleCandidate":    XK_Hangul_MultipleCandidate,
	"Hangul_PreviousCandidate":    XK_Hangul_PreviousCandidate,
	"Hangul_Special":              XK_Hangul_Special,
	"Hangul_switch":               XK_Hangul_switch,
	"Hangul_Kiyeog":               XK_Hangul_Kiyeog,
	"Hangul_SsangKiyeog":          XK_Hangul_SsangKiyeog,
	"Hangul_KiyeogSios":           XK_Hangul_KiyeogSios,
	"Hangul_Nieun":                XK_Hangul_Nieun,
	"Hangul_NieunJieuj":           XK_Hangul_NieunJieuj,
	"Hangul_NieunHieuh":           XK_Hangul_NieunHieuh,
	"Hangul_Dikeud":               XK_Hangul_Dikeud,
	"Hangul_SsangDikeud":          XK_Hangul_SsangDikeud,
	"Hangul_Rieul":                XK_Hangul_Rieul,
	"Hangul_RieulKiyeog":          XK_Hangul_RieulKiyeog,
	"Hangul_RieulMieum":           XK_Hangul_RieulMieum,
	"Hangul_RieulPieub":           XK_Hangul_RieulPieub,
	"Hangul_RieulSios":            XK_Hangul_RieulSios,
	"Hangul_RieulTieut":           XK_Hangul_RieulTieut,
	"Hangul_RieulPhieuf":          XK_Hangul_RieulPhieuf,
	"Hangul_RieulHieuh":           XK_Hangul_RieulHieuh,
	"Hangul_Mieum":                XK_Hangul_Mieum,
	"Hangul_Pieub":                XK_Hangul_Pieub,
	"Hangul_SsangPieub":           XK_Hangul_SsangPieub,
	"Hangul_PieubSios":            XK_Hangul_PieubSios,
	"Hangul_Sios":                 XK_Hangul_Sios,
	"Hangul_SsangSios":            XK_Hangul_SsangSios,
	"Hangul_Ieung":                XK_Hangul_Ieung,
	"Hangul_Jieuj":                XK_Hangul_Jieuj,
	"Hangul_SsangJieuj":           XK_Hangul_SsangJieuj,
	"Hangul_Cieuc":                XK_Hangul_Cieuc,
	"Hangul_Khieuq":               XK_Hangul_Khieuq,
	"Hangul_Tieut":                XK_Hangul_Tieut,
	"Hangul_Phieuf":               XK_Hangul_Phieuf,
	"Hangul_Hieuh":                XK_Hangul_Hieuh,
	"Hangul_A":                    XK_Hangul_A,
	"Hangul_AE":                   XK_Hangul_AE,
	"Hangul_YA":                   XK_Hangul_YA,
	"Hangul_YAE":                  XK_Hangul_YAE,
	"Hangul_EO":                   XK_Hangul_EO,
	"Hangul_E":                    XK_Hangul_E,
	"Hangul_YEO":                  XK_Hangul_YEO,
	"Hangul_YE":                   XK_Hangul_YE,
	"Hangul_O":                    XK_Hangul_O,
	"Hangul_WA":                   XK_Hangul_WA,
	"Hangul_WAE":                  XK_Hangul_WAE,
	"Hangul_OE":                   XK_Hangul_OE,
	"Hangul_YO":                   XK_Hangul_YO,
	"Hangul_U":                    XK_Hangul_U,
	"Hangul_WEO":                  XK_Hangul_WEO,
	"Hangul_WE":                   XK_Hangul_WE,
	"Hangul_WI":                   XK_Hangul_WI,
	"Hangul_YU":                   XK_Hangul_YU,
	"Hangul_EU":                   XK_Hangul_EU,
	"Hangul_YI":                   XK_Hangul_YI,
	"Hangul_I":                    XK_Hangul_I,
	"Hangul_J_Kiyeog":             XK_Hangul_J_Kiyeog,
	"Hangul_J_SsangKiyeog":        XK_Hangul_J_SsangKiyeog,
	"Hangul_J_KiyeogSios":         XK_Hangul_J_KiyeogSios,
	"Hangul_J_Nieun":              XK_Hangul_J_Nieun,
	"Hangul_J_NieunJieuj":         XK_Hangul_J_NieunJieuj,
	"Hangul_J_NieunHieuh":         XK_Hangul_J_NieunHieuh,
	"Hangul_J_Dikeud":             XK_Hangul_J_Dikeud,
	"Hangul_J_Rieul":              XK_Hangul_J_Rieul,
	"Hangul_J_RieulKiyeog":        XK_Hangul_J_RieulKiyeog,
	"Hangul_J_RieulMieum":         XK_Hangul_J_RieulMieum,
	"Hangul_J_RieulPieub":         XK_Hangul_J_RieulPieub,
	"Hangul_J_RieulSios":          XK_Hangul_J_RieulSios,
	"Hangul_J_RieulTieut":         XK_Hangul_J_RieulTieut,
	"Hangul_J_RieulPhieuf":        XK_Hangul_J_RieulPhieuf,
	"Hangul_J_RieulHieuh":         XK_Hangul_J_RieulHieuh,
	"Hangul_J_Mieum":              XK_Hangul_J_Mieum,
	"Hangul_J_Pieub":              XK_Hangul_J_Pieub,
	"Hangul_J_PieubSios":          XK_Hangul_J_PieubSios,
	"Hangul_J_Sios":               XK_Hangul_J_Sios,
	"Hangul_J_SsangSios":          XK_Hangul_J_SsangSios,
	"Hangul_J_Ieung":              XK_Hangul_J_Ieung,
	"Hangul_J_Jieuj":              XK_Hangul_J_Jieuj,
	"Hangul_J_Cieuc":              XK_Hangul_J_Cieuc,
	"Hangul_J_Khieuq":             XK_Hangul_J_Khieuq,
	"Hangul_J_Tieut":              XK_Hangul_J_Tieut,
	"Hangul_J_Phieuf":             XK_Hangul_J_Phieuf,
	"Hangul_J_Hieuh":              XK_Hangul_J_Hieuh,
	"Hangul_RieulYeorinHieuh":     XK_Hangul_RieulYeorinHieuh,
	"Hangul_SunkyeongeumMieum":    XK_Hangul_SunkyeongeumMieum,
	"Hangul_SunkyeongeumPieub":    XK_Hangul_SunkyeongeumPieub,
	"Hangul_PanSios":              XK_Hangul_PanSios,
	"Hangul_KkogjiDalrinIeung":    XK_Hangul_KkogjiDalrinIeung,
	"Hangul_SunkyeongeumPhieuf":   XK_Hangul_SunkyeongeumPhieuf,
	"Hangul_YeorinHieuh":          XK_Hangul_YeorinHieuh,
	"Hangul_AraeA":                XK_Hangul_AraeA,
	"Hangul_AraeAE":               XK_Hangul_AraeAE,
	"Hangul_J_PanSios":            XK_Hangul_J_PanSios,
	"Hangul_J_KkogjiDalrinIeung":  XK_Hangul_J_KkogjiDalrinIeung,
	"Hangul_J_YeorinHieuh":        XK_Hangul_J_YeorinHieuh,
	"Korean_Won":                  XK_Korean_Won,
	"Armenian_ligature_ew":        XK_Armenian_ligature_ew,
	"Armenian_full_stop":          XK_Armenian_full_stop,
	"Armenian_verjaket":           XK_Armenian_verjaket,
	"Armenian_separation_mark":    XK_Armenian_separation_mark,
	"Armenian_but":                XK_Armenian_but,
	"Armenian_hyphen":             XK_Armenian_hyphen,
	"Armenian_yentamna":           XK_Armenian_yentamna,
	"Armenian_exclam":             XK_Armenian_exclam,
	"Armenian_amanak":             XK_Armenian_amanak,
	"Armenian_accent":             XK_Armenian_accent,
	"Armenian_shesht":             XK_Armenian_shesht,
	"Armenian_question":           XK_Armenian_question,
	"Armenian_paruyk":             XK_Armenian_paruyk,
	"Armenian_AYB":                XK_Armenian_AYB,
	"Armenian_ayb":                XK_Armenian_ayb,
	"Armenian_BEN":                XK_Armenian_BEN,
	"Armenian_ben":                XK_Armenian_ben,
	"Armenian_GIM":                XK_Armenian_GIM,
	"Armenian_gim":                XK_Armenian_gim,
	"Armenian_DA":                 XK_Armenian_DA,
	"Armenian_da":                 XK_Armenian_da,
	"Armenian_YECH":               XK_Armenian_YECH,
	"Armenian_yech":               XK_Armenian_yech,
	"Armenian_ZA":                 XK_Armenian_ZA,
	"Armenian_za":                 XK_Armenian_za,
	"Armenian_E":                  XK_Armenian_E,
	"Armenian_e":                  XK_Armenian_e,
	"Armenian_AT":                 XK_Armenian_AT,
	"Armenian_at":                 XK_Armenian_at,
	"Armenian_TO":                 XK_Armenian_TO,
	"Armenian_to":                 XK_Armenian_to,
	"Armenian_ZHE":                XK_Armenian_ZHE,
	"Armenian_zhe":                XK_Armenian_zhe,
	"Armenian_INI":                XK_Armenian_INI,
	"Armenian_ini":                XK_Armenian_ini,
	"Armenian_LYUN":               XK_Armenian_LYUN,
	"Armenian_lyun":               XK_Armenian_lyun,
	"Armenian_KHE":                XK_Armenian_KHE,
	"Armenian_khe":                XK_Armenian_khe,
	"Armenian_TSA":                XK_Armenian_TSA,
	"Armenian_tsa":                XK_Armenian_tsa,
	"Armenian_KEN":                XK_Armenian_KEN,
	"Armenian_ken":                XK_Armenian_ken,
	"Armenian_HO":                 XK_Armenian_HO,
	"Armenian_ho":                 XK_Armenian_ho,
	"Armenian_DZA":                XK_Armenian_DZA,
	"Armenian_dza":                XK_Armenian_dza,
	"Armenian_GHAT":               XK_Armenian_GHAT,
	"Armenian_ghat":               XK_Armenian_ghat,
	"Armenian_TCHE":               XK_Armenian_TCHE,
	"Armenian_tche":               XK_Armenian_tche,
	"Armenian_MEN":                XK_Armenian_MEN,
	"Armenian_men":                XK_Armenian_men,
	"Armenian_HI":                 XK_Armenian_HI,
	"Armenian_hi":                 XK_Armenian_hi,
	"Armenian_NU":                 XK_Armenian_NU,
	"Armenian_nu":                 XK_Armenian_nu,
	"Armenian_SHA":                XK_Armenian_SHA,
	"Armenian_sha":                XK_Armenian_sha,
	"Armenian_VO":                 XK_Armenian_VO,
	"Armenian_vo":                 XK_Armenian_vo,
	"Armenian_CHA":                XK_Armenian_CHA,
	"Armenian_cha":                XK_Armenian_cha,
	"Armenian_PE":                 XK_Armenian_PE,
	"Armenian_pe":                 XK_Armenian_pe,
	"Armenian_JE":                 XK_Armenian_JE,
	"Armenian_je":                 XK_Armenian_je,
	"Armenian_RA":                 XK_Armenian_RA,
	"Armenian_ra":                 XK_Armenian_ra,
	"Armenian_SE":                 XK_Armenian_SE,
	"Armenian_se":                 XK_Armenian_se,
	"Armenian_VEV":                XK_Armenian_VEV,
	"Armenian_vev":                XK_Armenian_vev,
	"Armenian_TYUN":               XK_Armenian_TYUN,
	"Armenian_tyun":               XK_Armenian_tyun,
	"Armenian_RE":                 XK_Armenian_RE,
	"Armenian_re":                 XK_Armenian_re,
	"Armenian_TSO":                XK_Armenian_TSO,
	"Armenian_tso":                XK_Armenian_tso,
	"Armenian_VYUN":               XK_Armenian_VYUN,
	"Armenian_vyun":               XK_Armenian_vyun,
	"Armenian_PYUR":               XK_Armenian_PYUR,
	"Armenian_pyur":               XK_Armenian_pyur,
	"Armenian_KE":                 XK_Armenian_KE,
	"Armenian_ke":                 XK_Armenian_ke,
	"Armenian_O":                  XK_Armenian_O,
	"Armenian_o":                  XK_Armenian_o,
	"Armenian_FE":                 XK_Armenian_FE,
	"Armenian_fe":                 XK_Armenian_fe,
	"Armenian_apostrophe":         XK_Armenian_apostrophe,
	"Georgian_an":                 XK_Georgian_an,
	"Georgian_ban":                XK_Georgian_ban,
	"Georgian_gan":                XK_Georgian_gan,
	"Georgian_don":                XK_Georgian_don,
	"Georgian_en":                 XK_Georgian_en,
	"Georgian_vin":                XK_Georgian_vin,
	"Georgian_zen":                XK_Georgian_zen,
	"Georgian_tan":                XK_Georgian_tan,
	"Georgian_in":                 XK_Georgian_in,
	"Georgian_kan":                XK_Georgian_kan,
	"Georgian_las":                XK_Georgian_las,
	"Georgian_man":                XK_Georgian_man,
	"Georgian_nar":                XK_Georgian_nar,
	"Georgian_on":                 XK_Georgian_on,
	"Georgian_par":                XK_Georgian_par,
	"Georgian_zhar":               XK_Georgian_zhar,
	"Georgian_rae":                XK_Georgian_rae,
	"Georgian_san":                XK_Georgian_san,
	"Georgian_tar":                XK_Georgian_tar,
	"Georgian_un":                 XK_Georgian_un,
	"Georgian_phar":               XK_Georgian_phar,
	"Georgian_khar":               XK_Georgian_khar,
	"Georgian_ghan":               XK_Georgian_ghan,
	"Georgian_qar":                XK_Georgian_qar,
	"Georgian_shin":               XK_Georgian_shin,
	"Georgian_chin":               XK_Georgian_chin,
	"Georgian_can":                XK_Georgian_can,
	"Georgian_jil":                XK_Georgian_jil,
	"Georgian_cil":                XK_Georgian_cil,
	"Georgian_char":               XK_Georgian_char,
	"Georgian_xan":                XK_Georgian_xan,
	"Georgian_jhan":               XK_Georgian_jhan,
	"Georgian_hae":                XK_Georgian_hae,
	"Georgian_he":                 XK_Georgian_he,
	"Georgian_hie":                XK_Georgian_hie,
	"Georgian_we":                 XK_Georgian_we,
	"Georgian_har":                XK_Georgian_har,
	"Georgian_hoe":                XK_Georgian_hoe,
	"Georgian_fi":                 XK_Georgian_fi,
	"Xabovedot":                   XK_Xabovedot,
	"Ibreve":                      XK_Ibreve,
	"Zstroke":                     XK_Zstroke,
	"Gcaron":                      XK_Gcaron,
	"Ocaron":                      XK_Ocaron,
	"Obarred":                     XK_Obarred,
	"xabovedot":                   XK_xabovedot,
	"ibreve":                      XK_ibreve,
	"zstroke":                     XK_zstroke,
	"gcaron":                      XK_gcaron,
	"ocaron":                      XK_ocaron,
	"obarred":                     XK_obarred,
	"SCHWA":                       XK_SCHWA,
	"schwa":                       XK_schwa,
	"EZH":                         XK_EZH,
	"ezh":                         XK_ezh,
	"Lbelowdot":                   XK_Lbelowdot,
	"lbelowdot":                   XK_lbelowdot,
	"Abelowdot":                   XK_Abelowdot,
	"abelowdot":                   XK_abelowdot,
	"Ahook":                       XK_Ahook,
	"ahook":                       XK_ahook,
	"Acircumflexacute":            XK_Acircumflexacute,
	"acircumflexacute":            XK_acircumflexacute,
	"Acircumflexgrave":            XK_Acircumflexgrave,
	"acircumflexgrave":            XK_acircumflexgrave,
	"Acircumflexhook":             XK_Acircumflexhook,
	"acircumflexhook":             XK_acircumflexhook,
	"Acircumflextilde":            XK_Acircumflextilde,
	"acircumflextilde":            XK_acircumflextilde,
	"Acircumflexbelowdot":         XK_Acircumflexbelowdot,
	"acircumflexbelowdot":         XK_acircumflexbelowdot,
	"Abreveacute":                 XK_Abreveacute,
	"abreveacute":                 XK_abreveacute,
	"Abrevegrave":                 XK_Abrevegrave,
	"abrevegrave":                 XK_abrevegrave,
	"Abrevehook":                  XK_Abrevehook,
	"abrevehook":                  XK_abrevehook,
	"Abrevetilde":                 XK_Abrevetilde,
	"abrevetilde":                 XK_abrevetilde,
	"Abrevebelowdot":              XK_Abrevebelowdot,
	"abrevebelowdot":              XK_abrevebelowdot,
	"Ebelowdot":                   XK_Ebelowdot,
	"ebelowdot":                   XK_ebelowdot,
	"Ehook":                       XK_Ehook,
	"ehook":                       XK_ehook,
	"Etilde":                      XK_Etilde,
	"etilde":                      XK_etilde,
	"Ecircumflexacute":            XK_Ecircumflexacute,
	"ecircumflexacute":            XK_ecircumflexacute,
	"Ecircumflexgrave":            XK_Ecircumflexgrave,
	"ecircumflexgrave":            XK_ecircumflexgrave,
	"Ecircumflexhook":             XK_Ecircumflexhook,
	"ecircumflexhook":             XK_ecircumflexhook,
	"Ecircumflextilde":            XK_Ecircumflextilde,
	"ecircumflextilde":            XK_ecircumflextilde,
	"Ecircumflexbelowdot":         XK_Ecircumflexbelowdot,
	"ecircumflexbelowdot":         XK_ecircumflexbelowdot,
	"Ihook":                       XK_Ihook,
	"ihook":                       XK_ihook,
	"Ibelowdot":                   XK_Ibelowdot,
	"ibelowdot":                   XK_ibelowdot,
	"Obelowdot":                   XK_Obelowdot,
	"obelowdot":                   XK_obelowdot,
	"Ohook":                       XK_Ohook,
	"ohook":                       XK_ohook,
	"Ocircumflexacute":            XK_Ocircumflexacute,
	"ocircumflexacute":            XK_ocircumflexacute,
	"Ocircumflexgrave":            XK_Ocircumflexgrave,
	"ocircumflexgrave":            XK_ocircumflexgrave,
	"Ocircumflexhook":             XK_Ocircumflexhook,
	"ocircumflexhook":             XK_ocircumflexhook,
	"Ocircumflextilde":            XK_Ocircumflextilde,
	"ocircumflextilde":            XK_ocircumflextilde,
	"Ocircumflexbelowdot":         XK_Ocircumflexbelowdot,
	"ocircumflexbelowdot":         XK_ocircumflexbelowdot,
	"Ohornacute":                  XK_Ohornacute,
	"ohornacute":                  XK_ohornacute,
	"Ohorngrave":                  XK_Ohorngrave,
	"ohorngrave":                  XK_ohorngrave,
	"Ohornhook":                   XK_Ohornhook,
	"ohornhook":                   XK_ohornhook,
	"Ohorntilde":                  XK_Ohorntilde,
	"ohorntilde":                  XK_ohorntilde,
	"Ohornbelowdot":               XK_Ohornbelowdot,
	"ohornbelowdot":               XK_ohornbelowdot,
	"Ubelowdot":                   XK_Ubelowdot,
	"ubelowdot":                   XK_ubelowdot,
	"Uhook":                       XK_Uhook,
	"uhook":                       XK_uhook,
	"Uhornacute":                  XK_Uhornacute,
	"uhornacute":                  XK_uhornacute,
	"Uhorngrave":                  XK_Uhorngrave,
	"uhorngrave":                  XK_uhorngrave,
	"Uhornhook":                   XK_Uhornhook,
	"uhornhook":                   XK_uhornhook,
	"Uhorntilde":                  XK_Uhorntilde,
	"uhorntilde":                  XK_uhorntilde,
	"Uhornbelowdot":               XK_Uhornbelowdot,
	"uhornbelowdot":               XK_uhornbelowdot,
	"Ybelowdot":                   XK_Ybelowdot,
	"ybelowdot":                   XK_ybelowdot,
	"Yhook":                       XK_Yhook,
	"yhook":                       XK_yhook,
	"Ytilde":                      XK_Ytilde,
	"ytilde":                      XK_ytilde,
	"Ohorn":                       XK_Ohorn,
	"ohorn":                       XK_ohorn,
	"Uhorn":                       XK_Uhorn,
	"uhorn":                       XK_uhorn,
	"EcuSign":                     XK_EcuSign,
	"ColonSign":                   XK_ColonSign,
	"CruzeiroSign":                XK_CruzeiroSign,
	"FFrancSign":                  XK_FFrancSign,
	"LiraSign":                    XK_LiraSign,
	"MillSign":                    XK_MillSign,
	"NairaSign":                   XK_NairaSign,
	"PesetaSign":                  XK_PesetaSign,
	"RupeeSign":                   XK_RupeeSign,
	"WonSign":                     XK_WonSign,
	"NewSheqelSign":               XK_NewSheqelSign,
	"DongSign":                    XK_DongSign,
	"EuroSign":                    XK_EuroSign,
	"zerosuperior":                XK_zerosuperior,
	"foursuperior":                XK_foursuperior,
	"fivesuperior":                XK_fivesuperior,
	"sixsuperior":                 XK_sixsuperior,
	"sevensuperior":               XK_sevensuperior,
	"eightsuperior":               XK_eightsuperior,
	"ninesuperior":                XK_ninesuperior,
	"zerosubscript":               XK_zerosubscript,
	"onesubscript":                XK_onesubscript,
	"twosubscript":                XK_twosubscript,
	"threesubscript":              XK_threesubscript,
	"foursubscript":               XK_foursubscript,
	"fivesubscript":               XK_fivesubscript,
	"sixsubscript":                XK_sixsubscript,
	"sevensubscript":              XK_sevensubscript,
	"eightsubscript":              XK_eightsubscript,
	"ninesubscript":               XK_ninesubscript,
	"partdifferential":            XK_partdifferential,
	"emptyset":                    XK_emptyset,
	"elementof":                   XK_elementof,
	"notelementof":                XK_notelementof,
	"containsas":                  XK_containsas,
	"squareroot":                  XK_squareroot,
	"cuberoot":                    XK_cuberoot,
	"fourthroot":                  XK_fourthroot,
	"dintegral":                   XK_dintegral,
	"tintegral":                   XK_tintegral,
	"because":                     XK_because,
	"approxeq":                    XK_approxeq,
	"notapproxeq":                 XK_notapproxeq,
	"notidentical":                XK_notidentical,
	"stricteq":                    XK_stricteq,
	"braille_dot_1":               XK_braille_dot_1,
	"braille_dot_2":               XK_braille_dot_2,
	"braille_dot_3":               XK_braille_dot_3,
	"braille_dot_4":               XK_braille_dot_4,
	"braille_dot_5":               XK_braille_dot_5,
	"braille_dot_6":               XK_braille_dot_6,
	"braille_dot_7":               XK_braille_dot_7,
	"braille_dot_8":               XK_braille_dot_8,
	"braille_dot_9":               XK_braille_dot_9,
	"braille_dot_10":              XK_braille_dot_10,
	"braille_blank":               XK_braille_blank,
	"braille_dots_1":              XK_braille_dots_1,
	"braille_dots_2":              XK_braille_dots_2,
	"braille_dots_12":             XK_braille_dots_12,
	"braille_dots_3":              XK_braille_dots_3,
	"braille_dots_13":             XK_braille_dots_13,
	"braille_dots_23":             XK_braille_dots_23,
	"braille_dots_123":            XK_braille_dots_123,
	"braille_dots_4":              XK_braille_dots_4,
	"braille_dots_14":             XK_braille_dots_14,
	"braille_dots_24":             XK_braille_dots_24,
	"braille_dots_124":            XK_braille_dots_124,
	"braille_dots_34":             XK_braille_dots_34,
	"braille_dots_134":            XK_braille_dots_134,
	"braille_dots_234":            XK_braille_dots_234,
	"braille_dots_1234":           XK_braille_dots_1234,
	"braille_dots_5":              XK_braille_dots_5,
	"braille_dots_15":             XK_braille_dots_15,
	"braille_dots_25":             XK_braille_dots_25,
	"braille_dots_125":            XK_braille_dots_125,
	"braille_dots_35":             XK_braille_dots_35,
	"braille_dots_135":            XK_braille_dots_135,
	"braille_dots_235":            XK_braille_dots_235,
	"braille_dots_1235":           XK_braille_dots_1235,
	"braille_dots_45":             XK_braille_dots_45,
	"braille_dots_145":            XK_braille_dots_145,
	"braille_dots_245":            XK_braille_dots_245,
	"braille_dots_1245":           XK_braille_dots_1245,
	"braille_dots_345":            XK_braille_dots_345,
	"braille_dots_1345":           XK_braille_dots_1345,
	"braille_dots_2345":           XK_braille_dots_2345,
	"braille_dots_12345":          XK_braille_dots_12345,
	"braille_dots_6":              XK_braille_dots_6,
	"braille_dots_16":             XK_braille_dots_16,
	"braille_dots_26":             XK_braille_dots_26,
	"braille_dots_126":            XK_braille_dots_126,
	"braille_dots_36":             XK_braille_dots_36,
	"braille_dots_136":            XK_braille_dots_136,
	"braille_dots_236":            XK_braille_dots_236,
	"braille_dots_1236":           XK_braille_dots_1236,
	"braille_dots_46":             XK_braille_dots_46,
	"braille_dots_146":            XK_braille_dots_146,
	"braille_dots_246":            XK_braille_dots_246,
	"braille_dots_1246":           XK_braille_dots_1246,
	"braille_dots_346":            XK_braille_dots_346,
	"braille_dots_1346":           XK_braille_dots_1346,
	"braille_dots_2346":           XK_braille_dots_2346,
	"braille_dots_12346":          XK_braille_dots_12346,
	"braille_dots_56":             XK_braille_dots_56,
	"braille_dots_156":            XK_braille_dots_156,
	"braille_dots_256":            XK_braille_dots_256,
	"braille_dots_1256":           XK_braille_dots_1256,
	"braille_dots_356":            XK_braille_dots_356,
	"braille_dots_1356":           XK_braille_dots_1356,
	"braille_dots_2356":           XK_braille_dots_2356,
	"braille_dots_12356":          XK_braille_dots_12356,
	"braille_dots_456":            XK_braille_dots_456,
	"braille_dots_1456":           XK_braille_dots_1456,
	"braille_dots_2456":           XK_braille_dots_2456,
	"braille_dots_12456":          XK_braille_dots_12456,
	"braille_dots_3456":           XK_braille_dots_3456,
	"braille_dots_13456":          XK_braille_dots_13456,
	"braille_dots_23456":          XK_braille_dots_23456,
	"braille_dots_123456":         XK_braille_dots_123456,
	"braille_dots_7":              XK_braille_dots_7,
	"braille_dots_17":             XK_braille_dots_17,
	"braille_dots_27":             XK_braille_dots_27,
	"braille_dots_127":            XK_braille_dots_127,
	"braille_dots_37":             XK_braille_dots_37,
	"braille_dots_137":            XK_braille_dots_137,
	"braille_dots_237":            XK_braille_dots_237,
	"braille_dots_1237":           XK_braille_dots_1237,
	"braille_dots_47":             XK_braille_dots_47,
	"braille_dots_147":            XK_braille_dots_147,
	"braille_dots_247":            XK_braille_dots_247,
	"braille_dots_1247":           XK_braille_dots_1247,
	"braille_dots_347":            XK_braille_dots_347,
	"braille_dots_1347":           XK_braille_dots_1347,
	"braille_dots_2347":           XK_braille_dots_2347,
	"braille_dots_12347":          XK_braille_dots_12347,
	"braille_dots_57":             XK_braille_dots_57,
	"braille_dots_157":            XK_braille_dots_157,
	"braille_dots_257":            XK_braille_dots_257,
	"braille_dots_1257":           XK_braille_dots_1257,
	"braille_dots_357":            XK_braille_dots_357,
	"braille_dots_1357":           XK_braille_dots_1357,
	"braille_dots_2357":           XK_braille_dots_2357,
	"braille_dots_12357":          XK_braille_dots_12357,
	"braille_dots_457":            XK_braille_dots_457,
	"braille_dots_1457":           XK_braille_dots_1457,
	"braille_dots_2457":           XK_braille_dots_2457,
	"braille_dots_12457":          XK_braille_dots_12457,
	"braille_dots_3457":           XK_braille_dots_3457,
	"braille_dots_13457":          XK_braille_dots_13457,
	"braille_dots_23457":          XK_braille_dots_23457,
	"braille_dots_123457":         XK_braille_dots_123457,
	"braille_dots_67":             XK_braille_dots_67,
	"braille_dots_167":            XK_braille_dots_167,
	"braille_dots_267":            XK_braille_dots_267,
	"braille_dots_1267":           XK_braille_dots_1267,
	"braille_dots_367":            XK_braille_dots_367,
	"braille_dots_1367":           XK_braille_dots_1367,
	"braille_dots_2367":           XK_braille_dots_2367,
	"braille_dots_12367":          XK_braille_dots_12367,
	"braille_dots_467":            XK_braille_dots_467,
	"braille_dots_1467":           XK_braille_dots_1467,
	"braille_dots_2467":           XK_braille_dots_2467,
	"braille_dots_12467":          XK_braille_dots_12467,
	"braille_dots_3467":           XK_braille_dots_3467,
	"braille_dots_13467":          XK_braille_dots_13467,
	"braille_dots_23467":          XK_braille_dots_23467,
	"braille_dots_123467":         XK_braille_dots_123467,
	"braille_dots_567":            XK_braille_dots_567,
	"braille_dots_1567":           XK_braille_dots_1567,
	"braille_dots_2567":           XK_braille_dots_2567,
	"braille_dots_12567":          XK_braille_dots_12567,
	"braille_dots_3567":           XK_braille_dots_3567,
	"braille_dots_13567":          XK_braille_dots_13567,
	"braille_dots_23567":          XK_braille_dots_23567,
	"braille_dots_123567":         XK_braille_dots_123567,
	"braille_dots_4567":           XK_braille_dots_4567,
	"braille_dots_14567":          XK_braille_dots_14567,
	"braille_dots_24567":          XK_braille_dots_24567,
	"braille_dots_124567":         XK_braille_dots_124567,
	"braille_dots_34567":          XK_braille_dots_34567,
	"braille_dots_134567":         XK_braille_dots_134567,
	"braille_dots_234567":         XK_braille_dots_234567,
	"braille_dots_1234567":        XK_braille_dots_1234567,
	"braille_dots_8":              XK_braille_dots_8,
	"braille_dots_18":             XK_braille_dots_18,
	"braille_dots_28":             XK_braille_dots_28,
	"braille_dots_128":            XK_braille_dots_128,
	"braille_dots_38":             XK_braille_dots_38,
	"braille_dots_138":            XK_braille_dots_138,
	"braille_dots_238":            XK_braille_dots_238,
	"braille_dots_1238":           XK_braille_dots_1238,
	"braille_dots_48":             XK_braille_dots_48,
	"braille_dots_148":            XK_braille_dots_148,
	"braille_dots_248":            XK_braille_dots_248,
	"braille_dots_1248":           XK_braille_dots_1248,
	"braille_dots_348":            XK_braille_dots_348,
	"braille_dots_1348":           XK_braille_dots_1348,
	"braille_dots_2348":           XK_braille_dots_2348,
	"braille_dots_12348":          XK_braille_dots_12348,
	"braille_dots_58":             XK_braille_dots_58,
	"braille_dots_158":            XK_braille_dots_158,
	"braille_dots_258":            XK_braille_dots_258,
	"braille_dots_1258":           XK_braille_dots_1258,
	"braille_dots_358":            XK_braille_dots_358,
	"braille_dots_1358":           XK_braille_dots_1358,
	"braille_dots_2358":           XK_braille_dots_2358,
	"braille_dots_12358":          XK_braille_dots_12358,
	"braille_dots_458":            XK_braille_dots_458,
	"braille_dots_1458":           XK_braille_dots_1458,
	"braille_dots_2458":           XK_braille_dots_2458,
	"braille_dots_12458":          XK_braille_dots_12458,
	"braille_dots_3458":           XK_braille_dots_3458,
	"braille_dots_13458":          XK_braille_dots_13458,
	"braille_dots_23458":          XK_braille_dots_23458,
	"braille_dots_123458":         XK_braille_dots_123458,
	"braille_dots_68":             XK_braille_dots_68,
	"braille_dots_168":            XK_braille_dots_168,
	"braille_dots_268":            XK_braille_dots_268,
	"braille_dots_1268":           XK_braille_dots_1268,
	"braille_dots_368":            XK_braille_dots_368,
	"braille_dots_1368":           XK_braille_dots_1368,
	"braille_dots_2368":           XK_braille_dots_2368,
	"braille_dots_12368":          XK_braille_dots_12368,
	"braille_dots_468":            XK_braille_dots_468,
	"braille_dots_1468":           XK_braille_dots_1468,
	"braille_dots_2468":           XK_braille_dots_2468,
	"braille_dots_12468":          XK_braille_dots_12468,
	"braille_dots_3468":           XK_braille_dots_3468,
	"braille_dots_13468":          XK_braille_dots_13468,
	"braille_dots_23468":          XK_braille_dots_23468,
	"braille_dots_123468":         XK_braille_dots_123468,
	"braille_dots_568":            XK_braille_dots_568,
	"braille_dots_1568":           XK_braille_dots_1568,
	"braille_dots_2568":           XK_braille_dots_2568,
	"braille_dots_12568":          XK_braille_dots_12568,
	"braille_dots_3568":           XK_braille_dots_3568,
	"braille_dots_13568":          XK_braille_dots_13568,
	"braille_dots_23568":          XK_braille_dots_23568,
	"braille_dots_123568":         XK_braille_dots_123568,
	"braille_dots_4568":           XK_braille_dots_4568,
	"braille_dots_14568":          XK_braille_dots_14568,
	"braille_dots_24568":          XK_braille_dots_24568,
	"braille_dots_124568":         XK_braille_dots_124568,
	"braille_dots_34568":          XK_braille_dots_34568,
	"braille_dots_134568":         XK_braille_dots_134568,
	"braille_dots_234568":         XK_braille_dots_234568,
	"braille_dots_1234568":        XK_braille_dots_1234568,
	"braille_dots_78":             XK_braille_dots_78,
	"braille_dots_178":            XK_braille_dots_178,
	"braille_dots_278":            XK_braille_dots_278,
	"braille_dots_1278":           XK_braille_dots_1278,
	"braille_dots_378":            XK_braille_dots_378,
	"braille_dots_1378":           XK_braille_dots_1378,
	"braille_dots_2378":           XK_braille_dots_2378,
	"braille_dots_12378":          XK_braille_dots_12378,
	"braille_dots_478":            XK_braille_dots_478,
	"braille_dots_1478":           XK_braille_dots_1478,
	"braille_dots_2478":           XK_braille_dots_2478,
	"braille_dots_12478":          XK_braille_dots_12478,
	"braille_dots_3478":           XK_braille_dots_3478,
	"braille_dots_13478":          XK_braille_dots_13478,
	"braille_dots_23478":          XK_braille_dots_23478,
	"braille_dots_123478":         XK_braille_dots_123478,
	"braille_dots_578":            XK_braille_dots_578,
	"braille_dots_1578":           XK_braille_dots_1578,
	"braille_dots_2578":           XK_braille_dots_2578,
	"braille_dots_12578":          XK_braille_dots_12578,
	"braille_dots_3578":           XK_braille_dots_3578,
	"braille_dots_13578":          XK_braille_dots_13578,
	"braille_dots_23578":          XK_braille_dots_23578,
	"braille_dots_123578":         XK_braille_dots_123578,
	"braille_dots_4578":           XK_braille_dots_4578,
	"braille_dots_14578":          XK_braille_dots_14578,
	"braille_dots_24578":          XK_braille_dots_24578,
	"braille_dots_124578":         XK_braille_dots_124578,
	"braille_dots_34578":          XK_braille_dots_34578,
	"braille_dots_134578":         XK_braille_dots_134578,
	"braille_dots_234578":         XK_braille_dots_234578,
	"braille_dots_1234578":        XK_braille_dots_1234578,
	"braille_dots_678":            XK_braille_dots_678,
	"braille_dots_1678":           XK_braille_dots_1678,
	"braille_dots_2678":           XK_braille_dots_2678,
	"braille_dots_12678":          XK_braille_dots_12678,
	"braille_dots_3678":           XK_braille_dots_3678,
	"braille_dots_13678":          XK_braille_dots_13678,
	"braille_dots_23678":          XK_braille_dots_23678,
	"braille_dots_123678":         XK_braille_dots_123678,
	"braille_dots_4678":           XK_braille_dots_4678,
	"braille_dots_14678":          XK_braille_dots_14678,
	"braille_dots_24678":          XK_braille_dots_24678,
	"braille_dots_124678":         XK_braille_dots_124678,
	"braille_dots_34678":          XK_braille_dots_34678,
	"braille_dots_134678":         XK_braille_dots_134678,
	"braille_dots_234678":         XK_braille_dots_234678,
	"braille_dots_1234678":        XK_braille_dots_1234678,
	"braille_dots_5678":           XK_braille_dots_5678,
	"braille_dots_15678":          XK_braille_dots_15678,
	"braille_dots_25678":          XK_braille_dots_25678,
	"braille_dots_125678":         XK_braille_dots_125678,
	"braille_dots_35678":          XK_braille_dots_35678,
	"braille_dots_135678":         XK_braille_dots_135678,
	"braille_dots_235678":         XK_braille_dots_235678,
	"braille_dots_1235678":        XK_braille_dots_1235678,
	"braille_dots_45678":          XK_braille_dots_45678,
	"braille_dots_145678":         XK_braille_dots_145678,
	"braille_dots_245678":         XK_braille_dots_245678,
	"braille_dots_1245678":        XK_braille_dots_1245678,
	"braille_dots_345678":         XK_braille_dots_345678,
	"braille_dots_1345678":        XK_braille_dots_1345678,
	"braille_dots_2345678":        XK_braille_dots_2345678,
	"braille_dots_12345678":       XK_braille_dots_12345678,
	"Sinh_ng":                     XK_Sinh_ng,
	"Sinh_h2":                     XK_Sinh_h2,
	"Sinh_a":                      XK_Sinh_a,
	"Sinh_aa":                     XK_Sinh_aa,
	"Sinh_ae":                     XK_Sinh_ae,
	"Sinh_aee":                    XK_Sinh_aee,
	"Sinh_i":                      XK_Sinh_i,
	"Sinh_ii":                     XK_Sinh_ii,
	"Sinh_u":                      XK_Sinh_u,
	"Sinh_uu":                     XK_Sinh_uu,
	"Sinh_ri":                     XK_Sinh_ri,
	"Sinh_rii":                    XK_Sinh_rii,
	"Sinh_lu":                     XK_Sinh_lu,
	"Sinh_luu":                    XK_Sinh_luu,
	"Sinh_e":                      XK_Sinh_e,
	"Sinh_ee":                     XK_Sinh_ee,
	"Sinh_ai":                     XK_Sinh_ai,
	"Sinh_o":                      XK_Sinh_o,
	"Sinh_oo":                     XK_Sinh_oo,
	"Sinh_au":                     XK_Sinh_au,
	"Sinh_ka":                     XK_Sinh_ka,
	"Sinh_kha":                    XK_Sinh_kha,
	"Sinh_ga":                     XK_Sinh_ga,
	"Sinh_gha":                    XK_Sinh_gha,
	"Sinh_ng2":                    XK_Sinh_ng2,
	"Sinh_nga":                    XK_Sinh_nga,
	"Sinh_ca":                     XK_Sinh_ca,
	"Sinh_cha":                    XK_Sinh_cha,
	"Sinh_ja":                     XK_Sinh_ja,
	"Sinh_jha":                    XK_Sinh_jha,
	"Sinh_nya":                    XK_Sinh_nya,
	"Sinh_jnya":                   XK_Sinh_jnya,
	"Sinh_nja":                    XK_Sinh_nja,
	"Sinh_tta":                    XK_Sinh_tta,
	"Sinh_ttha":                   XK_Sinh_ttha,
	"Sinh_dda":                    XK_Sinh_dda,
	"Sinh_ddha":                   XK_Sinh_ddha,
	"Sinh_nna":                    XK_Sinh_nna,
	"Sinh_ndda":                   XK_Sinh_ndda,
	"Sinh_tha":                    XK_Sinh_tha,
	"Sinh_thha":                   XK_Sinh_thha,
	"Sinh_dha":                    XK_Sinh_dha,
	"Sinh_dhha":                   XK_Sinh_dhha,
	"Sinh_na":                     XK_Sinh_na,
	"Sinh_ndha":                   XK_Sinh_ndha,
	"Sinh_pa":                     XK_Sinh_pa,
	"Sinh_pha":                    XK_Sinh_pha,
	"Sinh_ba":                     XK_Sinh_ba,
	"Sinh_bha":                    XK_Sinh_bha,
	"Sinh_ma":                     XK_Sinh_ma,
	"Sinh_mba":                    XK_Sinh_mba,
	"Sinh_ya":                     XK_Sinh_ya,
	"Sinh_ra":                     XK_Sinh_ra,
	"Sinh_la":                     XK_Sinh_la,
	"Sinh_va":                     XK_Sinh_va,
	"Sinh_sha":                    XK_Sinh_sha,
	"Sinh_ssha":                   XK_Sinh_ssha,
	"Sinh_sa":                     XK_Sinh_sa,
	"Sinh_ha":                     XK_Sinh_ha,
	"Sinh_lla":                    XK_Sinh_lla,
	"Sinh_fa":                     XK_Sinh_fa,
	"Sinh_al":                     XK_Sinh_al,
	"Sinh_aa2":                    XK_Sinh_aa2,
	"Sinh_ae2":                    XK_Sinh_ae2,
	"Sinh_aee2":                   XK_Sinh_aee2,
	"Sinh_i2":                     XK_Sinh_i2,
	"Sinh_ii2":                    XK_Sinh_ii2,
	"Sinh_u2":                     XK_Sinh_u2,
	"Sinh_uu2":                    XK_Sinh_uu2,
	"Sinh_ru2":                    XK_Sinh_ru2,
	"Sinh_e2":                     XK_Sinh_e2,
	"Sinh_ee2":                    XK_Sinh_ee2,
	"Sinh_ai2":                    XK_Sinh_ai2,
	"Sinh_o2":                     XK_Sinh_o2,
	"Sinh_oo2":                    XK_Sinh_oo2,
	"Sinh_au2":                    XK_Sinh_au2,
	"Sinh_lu2":                    XK_Sinh_lu2,
	"Sinh_ruu2":                   XK_Sinh_ruu2,
	"Sinh_luu2":                   XK_Sinh_luu2,
	"Sinh_kunddaliya":             XK_Sinh_kunddaliya,
	"XF86ModeLock":                XF86XK_ModeLock,
	"XF86MonBrightnessUp":         XF86XK_MonBrightnessUp,
	"XF86MonBrightnessDown":       XF86XK_MonBrightnessDown,
	"XF86KbdLightOnOff":           XF86XK_KbdLightOnOff,
	"XF86KbdBrightnessUp":         XF86XK_KbdBrightnessUp,
	"XF86KbdBrightnessDown":       XF86XK_KbdBrightnessDown,
	"XF86Standby":                 XF86XK_Standby,
	"XF86AudioLowerVolume":        XF86XK_AudioLowerVolume,
	"XF86AudioMute":               XF86XK_AudioMute,
	"XF86AudioRaiseVolume":        XF86XK_AudioRaiseVolume,
	"XF86AudioPlay":               XF86XK_AudioPlay,
	"XF86AudioStop":               XF86XK_AudioStop,
	"XF86AudioPrev":               XF86XK_AudioPrev,
	"XF86AudioNext":               XF86XK_AudioNext,
	"XF86HomePage":                XF86XK_HomePage,
	"XF86Mail":                    XF86XK_Mail,
	"XF86Start":                   XF86XK_Start,
	"XF86Search":                  XF86XK_Search,
	"XF86AudioRecord":             XF86XK_AudioRecord,
	"XF86Calculator":              XF86XK_Calculator,
	"XF86Memo":                    XF86XK_Memo,
	"XF86ToDoList":                XF86XK_ToDoList,
	"XF86Calendar":                XF86XK_Calendar,
	"XF86PowerDown":               XF86XK_PowerDown,
	"XF86ContrastAdjust":          XF86XK_ContrastAdjust,
	"XF86RockerUp":                XF86XK_RockerUp,
	"XF86RockerDown":              XF86XK_RockerDown,
	"XF86RockerEnter":             XF86XK_RockerEnter,
	"XF86Back":                    XF86XK_Back,
	"XF86Forward":                 XF86XK_Forward,
	"XF86Stop":                    XF86XK_Stop,
	"XF86Refresh":                 XF86XK_Refresh,
	"XF86PowerOff":                XF86XK_PowerOff,
	"XF86WakeUp":                  XF86XK_WakeUp,
	"XF86Eject":                   XF86XK_Eject,
	"XF86ScreenSaver":             XF86XK_ScreenSaver,
	"XF86WWW":                     XF86XK_WWW,
	"XF86Sleep":                   XF86XK_Sleep,
	"XF86Favorites":               XF86XK_Favorites,
	"XF86AudioPause":              XF86XK_AudioPause,
	"XF86AudioMedia":              XF86XK_AudioMedia,
	"XF86MyComputer":              XF86XK_MyComputer,
	"XF86VendorHome":              XF86XK_VendorHome,
	"XF86LightBulb":               XF86XK_LightBulb,
	"XF86Shop":                    XF86XK_Shop,
	"XF86History":                 XF86XK_History,
	"XF86OpenURL":                 XF86XK_OpenURL,
	"XF86AddFavorite":             XF86XK_AddFavorite,
	"XF86HotLinks":                XF86XK_HotLinks,
	"XF86BrightnessAdjust":        XF86XK_BrightnessAdjust,
	"XF86Finance":                 XF86XK_Finance,
	"XF86Community":               XF86XK_Community,
	"XF86AudioRewind":             XF86XK_AudioRewind,
	"XF86BackForward":             XF86XK_BackForward,
	"XF86Launch0":                 XF86XK_Launch0,
	"XF86Launch1":                 XF86XK_Launch1,
	"XF86Launch2":                 XF86XK_Launch2,
	"XF86Launch3":                 XF86XK_Launch3,
	"XF86Launch4":                 XF86XK_Launch4,
	"XF86Launch5":                 XF86XK_Launch5,
	"XF86Launch6":                 XF86XK_Launch6,
	"XF86Launch7":                 XF86XK_Launch7,
	"XF86Launch8":                 XF86XK_Launch8,
	"XF86Launch9":                 XF86XK_Launch9,
	"XF86LaunchA":                 XF86XK_LaunchA,
	"XF86LaunchB":                 XF86XK_LaunchB,
	"XF86LaunchC":                 XF86XK_LaunchC,
	"XF86LaunchD":                 XF86XK_LaunchD,
	"XF86LaunchE":                 XF86XK_LaunchE,
	"XF86LaunchF":                 XF86XK_LaunchF,
	"XF86ApplicationLeft":         XF86XK_ApplicationLeft,
	"XF86ApplicationRight":        XF86XK_ApplicationRight,
	"XF86Book":                    XF86XK_Book,
	"XF86CD":                      XF86XK_CD,
	"XF86Calculater":              XF86XK_Calculater,
	"XF86Clear":                   XF86XK_Clear,
	"XF86Close":                   XF86XK_Close,
	"XF86Copy":                    XF86XK_Copy,
	"XF86Cut":                     XF86XK_Cut,
	"XF86Display":                 XF86XK_Display,
	"XF86DOS":                     XF86XK_DOS,
	"XF86Documents":               XF86XK_Documents,
	"XF86Excel":                   XF86XK_Excel,
	"XF86Explorer":                XF86XK_Explorer,
	"XF86Game":                    XF86XK_Game,
	"XF86Go":                      XF86XK_Go,
	"XF86iTouch":                  XF86XK_iTouch,
	"XF86LogOff":                  XF86XK_LogOff,
	"XF86Market":                  XF86XK_Market,
	"XF86Meeting":                 XF86XK_Meeting,
	"XF86MenuKB":                  XF86XK_MenuKB,
	"XF86MenuPB":                  XF86XK_MenuPB,
	"XF86MySites":                 XF86XK_MySites,
	"XF86New":                     XF86XK_New,
	"XF86News":                    XF86XK_News,
	"XF86OfficeHome":              XF86XK_OfficeHome,
	"XF86Open":                    XF86XK_Open,
	"XF86Option":                  XF86XK_Option,
	"XF86Paste":                   XF86XK_Paste,
	"XF86Phone":                   XF86XK_Phone,
	"XF86Q":                       XF86XK_Q,
	"XF86Reply":                   XF86XK_Reply,
	"XF86Reload":                  XF86XK_Reload,
	"XF86RotateWindows":           XF86XK_RotateWindows,
	"XF86RotationPB":              XF86XK_RotationPB,
	"XF86RotationKB":              XF86XK_RotationKB,
	"XF86Save":                    XF86XK_Save,
	"XF86ScrollUp":                XF86XK_ScrollUp,
	"XF86ScrollDown":              XF86XK_ScrollDown,
	"XF86ScrollClick":             XF86XK_ScrollClick,
	"XF86Send":                    XF86XK_Send,
	"XF86Spell":                   XF86XK_Spell,
	"XF86SplitScreen":             XF86XK_SplitScreen,
	"XF86Support":                 XF86XK_Support,
	"XF86TaskPane":                XF86XK_TaskPane,
	"XF86Terminal":                XF86XK_Terminal,
	"XF86Tools":                   XF86XK_Tools,
	"XF86Travel":                  XF86XK_Travel,
	"XF86UserPB":                  XF86XK_UserPB,
	"XF86User1KB":                 XF86XK_User1KB,
	"XF86User2KB":                 XF86XK_User2KB,
	"XF86Video":                   XF86XK_Video,
	"XF86WheelButton":             XF86XK_WheelButton,
	"XF86Word":                    XF86XK_Word,
	"XF86Xfer":                    XF86XK_Xfer,
	"XF86ZoomIn":                  XF86XK_ZoomIn,
	"XF86ZoomOut":                 XF86XK_ZoomOut,
	"XF86Away":                    XF86XK_Away,
	"XF86Messenger":               XF86XK_Messenger,
	"XF86WebCam":                  XF86XK_WebCam,
	"XF86MailForward":             XF86XK_MailForward,
	"XF86Pictures":                XF86XK_Pictures,
	"XF86Music":                   XF86XK_Music,
	"XF86Battery":                 XF86XK_Battery,
	"XF86Bluetooth":               XF86XK_Bluetooth,
	"XF86WLAN":                    XF86XK_WLAN,
	"XF86UWB":                     XF86XK_UWB,
	"XF86AudioForward":            XF86XK_AudioForward,
	"XF86AudioRepeat":             XF86XK_AudioRepeat,
	"XF86AudioRandomPlay":         XF86XK_AudioRandomPlay,
	"XF86Subtitle":                XF86XK_Subtitle,
	"XF86AudioCycleTrack":         XF86XK_AudioCycleTrack,
	"XF86CycleAngle":              XF86XK_CycleAngle,
	"XF86FrameBack":               XF86XK_FrameBack,
	"XF86FrameForward":            XF86XK_FrameForward,
	"XF86Time":                    XF86XK_Time,
	"XF86Select":                  XF86XK_Select,
	"XF86View":                    XF86XK_View,
	"XF86TopMenu":                 XF86XK_TopMenu,
	"XF86Red":                     XF86XK_Red,
	"XF86Green":                   XF86XK_Green,
	"XF86Yellow":                  XF86XK_Yellow,
	"XF86Blue":                    XF86XK_Blue,
	"XF86Suspend":                 XF86XK_Suspend,
	"XF86Hibernate":               XF86XK_Hibernate,
	"XF86TouchpadToggle":          XF86XK_TouchpadToggle,
	"XF86TouchpadOn":              XF86XK_TouchpadOn,
	"XF86TouchpadOff":             XF86XK_TouchpadOff,
	"XF86AudioMicMute":            XF86XK_AudioMicMute,
	"XF86Keyboard":                XF86XK_Keyboard,
	"XF86WWAN":                    XF86XK_WWAN,
	"XF86RFKill":                  XF86XK_RFKill,
	"XF86AudioPreset":             XF86XK_AudioPreset,
	"XF86Switch_VT_1":             XF86XK_Switch_VT_1,
	"XF86Switch_VT_2":             XF86XK_Switch_VT_2,
	"XF86Switch_VT_3":             XF86XK_Switch_VT_3,
	"XF86Switch_VT_4":             XF86XK_Switch_VT_4,
	"XF86Switch_VT_5":             XF86XK_Switch_VT_5,
	"XF86Switch_VT_6":             XF86XK_Switch_VT_6,
	"XF86Switch_VT_7":             XF86XK_Switch_VT_7,
	"XF86Switch_VT_8":             XF86XK_Switch_VT_8,
	"XF86Switch_VT_9":             XF86XK_Switch_VT_9,
	"XF86Switch_VT_10":            XF86XK_Switch_VT_10,
	"XF86Switch_VT_11":            XF86XK_Switch_VT_11,
	"XF86Switch_VT_12":            XF86XK_Switch_VT_12,
	"XF86Ungrab":                  XF86XK_Ungrab,
	"XF86ClearGrab":               XF86XK_ClearGrab,
	"XF86Next_VMode":              XF86XK_Next_VMode,
	"XF86Prev_VMode":              XF86XK_Prev_VMode,
	"XF86LogWindowTree":           XF86XK_LogWindowTree,
	"XF86LogGrabInfo":             XF86XK_LogGrabInfo,
}
var KeysymVisibleCharMap = map[x.Keysym]rune{
	XK_space:          '\u0020', //
	XK_exclam:         '\u0021', // !
	XK_quotedbl:       '\u0022', // "
	XK_numbersign:     '\u0023', // #
	XK_dollar:         '\u0024', // $
	XK_percent:        '\u0025', // %
	XK_ampersand:      '\u0026', // &
	XK_apostrophe:     '\u0027', // '
	XK_parenleft:      '\u0028', // (
	XK_parenright:     '\u0029', // )
	XK_asterisk:       '\u002a', // *
	XK_plus:           '\u002b', // +
	XK_comma:          '\u002c', // ,
	XK_minus:          '\u002d', // -
	XK_period:         '\u002e', // .
	XK_slash:          '\u002f', // /
	XK_0:              '\u0030', // 0
	XK_1:              '\u0031', // 1
	XK_2:              '\u0032', // 2
	XK_3:              '\u0033', // 3
	XK_4:              '\u0034', // 4
	XK_5:              '\u0035', // 5
	XK_6:              '\u0036', // 6
	XK_7:              '\u0037', // 7
	XK_8:              '\u0038', // 8
	XK_9:              '\u0039', // 9
	XK_colon:          '\u003a', // :
	XK_semicolon:      '\u003b', // ;
	XK_less:           '\u003c', // <
	XK_equal:          '\u003d', // =
	XK_greater:        '\u003e', // >
	XK_question:       '\u003f', // ?
	XK_at:             '\u0040', // @
	XK_A:              '\u0041', // A
	XK_B:              '\u0042', // B
	XK_C:              '\u0043', // C
	XK_D:              '\u0044', // D
	XK_E:              '\u0045', // E
	XK_F:              '\u0046', // F
	XK_G:              '\u0047', // G
	XK_H:              '\u0048', // H
	XK_I:              '\u0049', // I
	XK_J:              '\u004a', // J
	XK_K:              '\u004b', // K
	XK_L:              '\u004c', // L
	XK_M:              '\u004d', // M
	XK_N:              '\u004e', // N
	XK_O:              '\u004f', // O
	XK_P:              '\u0050', // P
	XK_Q:              '\u0051', // Q
	XK_R:              '\u0052', // R
	XK_S:              '\u0053', // S
	XK_T:              '\u0054', // T
	XK_U:              '\u0055', // U
	XK_V:              '\u0056', // V
	XK_W:              '\u0057', // W
	XK_X:              '\u0058', // X
	XK_Y:              '\u0059', // Y
	XK_Z:              '\u005a', // Z
	XK_bracketleft:    '\u005b', // [
	XK_backslash:      '\u005c', // \
	XK_bracketright:   '\u005d', // ]
	XK_asciicircum:    '\u005e', // ^
	XK_underscore:     '\u005f', // _
	XK_grave:          '\u0060', // `
	XK_a:              '\u0061', // a
	XK_b:              '\u0062', // b
	XK_c:              '\u0063', // c
	XK_d:              '\u0064', // d
	XK_e:              '\u0065', // e
	XK_f:              '\u0066', // f
	XK_g:              '\u0067', // g
	XK_h:              '\u0068', // h
	XK_i:              '\u0069', // i
	XK_j:              '\u006a', // j
	XK_k:              '\u006b', // k
	XK_l:              '\u006c', // l
	XK_m:              '\u006d', // m
	XK_n:              '\u006e', // n
	XK_o:              '\u006f', // o
	XK_p:              '\u0070', // p
	XK_q:              '\u0071', // q
	XK_r:              '\u0072', // r
	XK_s:              '\u0073', // s
	XK_t:              '\u0074', // t
	XK_u:              '\u0075', // u
	XK_v:              '\u0076', // v
	XK_w:              '\u0077', // w
	XK_x:              '\u0078', // x
	XK_y:              '\u0079', // y
	XK_z:              '\u007a', // z
	XK_braceleft:      '\u007b', // {
	XK_bar:            '\u007c', // |
	XK_braceright:     '\u007d', // }
	XK_asciitilde:     '\u007e', // ~
	XK_nobreakspace:   '\u00a0', //
	XK_exclamdown:     '\u00a1', // ¡
	XK_cent:           '\u00a2', // ¢
	XK_sterling:       '\u00a3', // £
	XK_currency:       '\u00a4', // ¤
	XK_yen:            '\u00a5', // ¥
	XK_brokenbar:      '\u00a6', // ¦
	XK_section:        '\u00a7', // §
	XK_diaeresis:      '\u00a8', // ¨
	XK_copyright:      '\u00a9', // ©
	XK_ordfeminine:    '\u00aa', // ª
	XK_guillemotleft:  '\u00ab', // «
	XK_notsign:        '\u00ac', // ¬
	XK_hyphen:         '\u00ad', // ­
	XK_registered:     '\u00ae', // ®
	XK_macron:         '\u00af', // ¯
	XK_degree:         '\u00b0', // °
	XK_plusminus:      '\u00b1', // ±
	XK_twosuperior:    '\u00b2', // ²
	XK_threesuperior:  '\u00b3', // ³
	XK_acute:          '\u00b4', // ´
	XK_mu:             '\u00b5', // µ
	XK_paragraph:      '\u00b6', // ¶
	XK_periodcentered: '\u00b7', // ·
	XK_cedilla:        '\u00b8', // ¸
	XK_onesuperior:    '\u00b9', // ¹
	XK_masculine:      '\u00ba', // º
	XK_guillemotright: '\u00bb', // »
	XK_onequarter:     '\u00bc', // ¼
	XK_onehalf:        '\u00bd', // ½
	XK_threequarters:  '\u00be', // ¾
	XK_questiondown:   '\u00bf', // ¿
	XK_Agrave:         '\u00c0', // À
	XK_Aacute:         '\u00c1', // Á
	XK_Acircumflex:    '\u00c2', // Â
	XK_Atilde:         '\u00c3', // Ã
	XK_Adiaeresis:     '\u00c4', // Ä
	XK_Aring:          '\u00c5', // Å
	XK_AE:             '\u00c6', // Æ
	XK_Ccedilla:       '\u00c7', // Ç
	XK_Egrave:         '\u00c8', // È
	XK_Eacute:         '\u00c9', // É
	XK_Ecircumflex:    '\u00ca', // Ê
	XK_Ediaeresis:     '\u00cb', // Ë
	XK_Igrave:         '\u00cc', // Ì
	XK_Iacute:         '\u00cd', // Í
	XK_Icircumflex:    '\u00ce', // Î
	XK_Idiaeresis:     '\u00cf', // Ï
	XK_ETH:            '\u00d0', // Ð
	XK_Ntilde:         '\u00d1', // Ñ
	XK_Ograve:         '\u00d2', // Ò
	XK_Oacute:         '\u00d3', // Ó
	XK_Ocircumflex:    '\u00d4', // Ô
	XK_Otilde:         '\u00d5', // Õ
	XK_Odiaeresis:     '\u00d6', // Ö
	XK_multiply:       '\u00d7', // ×
	XK_Oslash:         '\u00d8', // Ø
	// XK_Ooblique == XK_Oslash # U+00D8 LATIN CAPITAL LETTER O WITH STROKE
	XK_Ugrave:      '\u00d9', // Ù
	XK_Uacute:      '\u00da', // Ú
	XK_Ucircumflex: '\u00db', // Û
	XK_Udiaeresis:  '\u00dc', // Ü
	XK_Yacute:      '\u00dd', // Ý
	XK_THORN:       '\u00de', // Þ
	XK_ssharp:      '\u00df', // ß
	XK_agrave:      '\u00e0', // à
	XK_aacute:      '\u00e1', // á
	XK_acircumflex: '\u00e2', // â
	XK_atilde:      '\u00e3', // ã
	XK_adiaeresis:  '\u00e4', // ä
	XK_aring:       '\u00e5', // å
	XK_ae:          '\u00e6', // æ
	XK_ccedilla:    '\u00e7', // ç
	XK_egrave:      '\u00e8', // è
	XK_eacute:      '\u00e9', // é
	XK_ecircumflex: '\u00ea', // ê
	XK_ediaeresis:  '\u00eb', // ë
	XK_igrave:      '\u00ec', // ì
	XK_iacute:      '\u00ed', // í
	XK_icircumflex: '\u00ee', // î
	XK_idiaeresis:  '\u00ef', // ï
	XK_eth:         '\u00f0', // ð
	XK_ntilde:      '\u00f1', // ñ
	XK_ograve:      '\u00f2', // ò
	XK_oacute:      '\u00f3', // ó
	XK_ocircumflex: '\u00f4', // ô
	XK_otilde:      '\u00f5', // õ
	XK_odiaeresis:  '\u00f6', // ö
	XK_division:    '\u00f7', // ÷
	XK_oslash:      '\u00f8', // ø
	// XK_ooblique == XK_oslash # U+00F8 LATIN SMALL LETTER O WITH STROKE
	XK_ugrave:                  '\u00f9', // ù
	XK_uacute:                  '\u00fa', // ú
	XK_ucircumflex:             '\u00fb', // û
	XK_udiaeresis:              '\u00fc', // ü
	XK_yacute:                  '\u00fd', // ý
	XK_thorn:                   '\u00fe', // þ
	XK_ydiaeresis:              '\u00ff', // ÿ
	XK_Aogonek:                 '\u0104', // Ą
	XK_breve:                   '\u02d8', // ˘
	XK_Lstroke:                 '\u0141', // Ł
	XK_Lcaron:                  '\u013d', // Ľ
	XK_Sacute:                  '\u015a', // Ś
	XK_Scaron:                  '\u0160', // Š
	XK_Scedilla:                '\u015e', // Ş
	XK_Tcaron:                  '\u0164', // Ť
	XK_Zacute:                  '\u0179', // Ź
	XK_Zcaron:                  '\u017d', // Ž
	XK_Zabovedot:               '\u017b', // Ż
	XK_aogonek:                 '\u0105', // ą
	XK_ogonek:                  '\u02db', // ˛
	XK_lstroke:                 '\u0142', // ł
	XK_lcaron:                  '\u013e', // ľ
	XK_sacute:                  '\u015b', // ś
	XK_caron:                   '\u02c7', // ˇ
	XK_scaron:                  '\u0161', // š
	XK_scedilla:                '\u015f', // ş
	XK_tcaron:                  '\u0165', // ť
	XK_zacute:                  '\u017a', // ź
	XK_doubleacute:             '\u02dd', // ˝
	XK_zcaron:                  '\u017e', // ž
	XK_zabovedot:               '\u017c', // ż
	XK_Racute:                  '\u0154', // Ŕ
	XK_Abreve:                  '\u0102', // Ă
	XK_Lacute:                  '\u0139', // Ĺ
	XK_Cacute:                  '\u0106', // Ć
	XK_Ccaron:                  '\u010c', // Č
	XK_Eogonek:                 '\u0118', // Ę
	XK_Ecaron:                  '\u011a', // Ě
	XK_Dcaron:                  '\u010e', // Ď
	XK_Dstroke:                 '\u0110', // Đ
	XK_Nacute:                  '\u0143', // Ń
	XK_Ncaron:                  '\u0147', // Ň
	XK_Odoubleacute:            '\u0150', // Ő
	XK_Rcaron:                  '\u0158', // Ř
	XK_Uring:                   '\u016e', // Ů
	XK_Udoubleacute:            '\u0170', // Ű
	XK_Tcedilla:                '\u0162', // Ţ
	XK_racute:                  '\u0155', // ŕ
	XK_abreve:                  '\u0103', // ă
	XK_lacute:                  '\u013a', // ĺ
	XK_cacute:                  '\u0107', // ć
	XK_ccaron:                  '\u010d', // č
	XK_eogonek:                 '\u0119', // ę
	XK_ecaron:                  '\u011b', // ě
	XK_dcaron:                  '\u010f', // ď
	XK_dstroke:                 '\u0111', // đ
	XK_nacute:                  '\u0144', // ń
	XK_ncaron:                  '\u0148', // ň
	XK_odoubleacute:            '\u0151', // ő
	XK_rcaron:                  '\u0159', // ř
	XK_uring:                   '\u016f', // ů
	XK_udoubleacute:            '\u0171', // ű
	XK_tcedilla:                '\u0163', // ţ
	XK_abovedot:                '\u02d9', // ˙
	XK_Hstroke:                 '\u0126', // Ħ
	XK_Hcircumflex:             '\u0124', // Ĥ
	XK_Iabovedot:               '\u0130', // İ
	XK_Gbreve:                  '\u011e', // Ğ
	XK_Jcircumflex:             '\u0134', // Ĵ
	XK_hstroke:                 '\u0127', // ħ
	XK_hcircumflex:             '\u0125', // ĥ
	XK_idotless:                '\u0131', // ı
	XK_gbreve:                  '\u011f', // ğ
	XK_jcircumflex:             '\u0135', // ĵ
	XK_Cabovedot:               '\u010a', // Ċ
	XK_Ccircumflex:             '\u0108', // Ĉ
	XK_Gabovedot:               '\u0120', // Ġ
	XK_Gcircumflex:             '\u011c', // Ĝ
	XK_Ubreve:                  '\u016c', // Ŭ
	XK_Scircumflex:             '\u015c', // Ŝ
	XK_cabovedot:               '\u010b', // ċ
	XK_ccircumflex:             '\u0109', // ĉ
	XK_gabovedot:               '\u0121', // ġ
	XK_gcircumflex:             '\u011d', // ĝ
	XK_ubreve:                  '\u016d', // ŭ
	XK_scircumflex:             '\u015d', // ŝ
	XK_kra:                     '\u0138', // ĸ
	XK_Rcedilla:                '\u0156', // Ŗ
	XK_Itilde:                  '\u0128', // Ĩ
	XK_Lcedilla:                '\u013b', // Ļ
	XK_Emacron:                 '\u0112', // Ē
	XK_Gcedilla:                '\u0122', // Ģ
	XK_Tslash:                  '\u0166', // Ŧ
	XK_rcedilla:                '\u0157', // ŗ
	XK_itilde:                  '\u0129', // ĩ
	XK_lcedilla:                '\u013c', // ļ
	XK_emacron:                 '\u0113', // ē
	XK_gcedilla:                '\u0123', // ģ
	XK_tslash:                  '\u0167', // ŧ
	XK_ENG:                     '\u014a', // Ŋ
	XK_eng:                     '\u014b', // ŋ
	XK_Amacron:                 '\u0100', // Ā
	XK_Iogonek:                 '\u012e', // Į
	XK_Eabovedot:               '\u0116', // Ė
	XK_Imacron:                 '\u012a', // Ī
	XK_Ncedilla:                '\u0145', // Ņ
	XK_Omacron:                 '\u014c', // Ō
	XK_Kcedilla:                '\u0136', // Ķ
	XK_Uogonek:                 '\u0172', // Ų
	XK_Utilde:                  '\u0168', // Ũ
	XK_Umacron:                 '\u016a', // Ū
	XK_amacron:                 '\u0101', // ā
	XK_iogonek:                 '\u012f', // į
	XK_eabovedot:               '\u0117', // ė
	XK_imacron:                 '\u012b', // ī
	XK_ncedilla:                '\u0146', // ņ
	XK_omacron:                 '\u014d', // ō
	XK_kcedilla:                '\u0137', // ķ
	XK_uogonek:                 '\u0173', // ų
	XK_utilde:                  '\u0169', // ũ
	XK_umacron:                 '\u016b', // ū
	XK_Wcircumflex:             '\u0174', // Ŵ
	XK_wcircumflex:             '\u0175', // ŵ
	XK_Ycircumflex:             '\u0176', // Ŷ
	XK_ycircumflex:             '\u0177', // ŷ
	XK_Babovedot:               '\u1e02', // Ḃ
	XK_babovedot:               '\u1e03', // ḃ
	XK_Dabovedot:               '\u1e0a', // Ḋ
	XK_dabovedot:               '\u1e0b', // ḋ
	XK_Fabovedot:               '\u1e1e', // Ḟ
	XK_fabovedot:               '\u1e1f', // ḟ
	XK_Mabovedot:               '\u1e40', // Ṁ
	XK_mabovedot:               '\u1e41', // ṁ
	XK_Pabovedot:               '\u1e56', // Ṗ
	XK_pabovedot:               '\u1e57', // ṗ
	XK_Sabovedot:               '\u1e60', // Ṡ
	XK_sabovedot:               '\u1e61', // ṡ
	XK_Tabovedot:               '\u1e6a', // Ṫ
	XK_tabovedot:               '\u1e6b', // ṫ
	XK_Wgrave:                  '\u1e80', // Ẁ
	XK_wgrave:                  '\u1e81', // ẁ
	XK_Wacute:                  '\u1e82', // Ẃ
	XK_wacute:                  '\u1e83', // ẃ
	XK_Wdiaeresis:              '\u1e84', // Ẅ
	XK_wdiaeresis:              '\u1e85', // ẅ
	XK_Ygrave:                  '\u1ef2', // Ỳ
	XK_ygrave:                  '\u1ef3', // ỳ
	XK_OE:                      '\u0152', // Œ
	XK_oe:                      '\u0153', // œ
	XK_Ydiaeresis:              '\u0178', // Ÿ
	XK_overline:                '\u203e', // ‾
	XK_kana_fullstop:           '\u3002', // 。
	XK_kana_openingbracket:     '\u300c', // 「
	XK_kana_closingbracket:     '\u300d', // 」
	XK_kana_comma:              '\u3001', // 、
	XK_kana_conjunctive:        '\u30fb', // ・
	XK_kana_WO:                 '\u30f2', // ヲ
	XK_kana_a:                  '\u30a1', // ァ
	XK_kana_i:                  '\u30a3', // ィ
	XK_kana_u:                  '\u30a5', // ゥ
	XK_kana_e:                  '\u30a7', // ェ
	XK_kana_o:                  '\u30a9', // ォ
	XK_kana_ya:                 '\u30e3', // ャ
	XK_kana_yu:                 '\u30e5', // ュ
	XK_kana_yo:                 '\u30e7', // ョ
	XK_kana_tsu:                '\u30c3', // ッ
	XK_prolongedsound:          '\u30fc', // ー
	XK_kana_A:                  '\u30a2', // ア
	XK_kana_I:                  '\u30a4', // イ
	XK_kana_U:                  '\u30a6', // ウ
	XK_kana_E:                  '\u30a8', // エ
	XK_kana_O:                  '\u30aa', // オ
	XK_kana_KA:                 '\u30ab', // カ
	XK_kana_KI:                 '\u30ad', // キ
	XK_kana_KU:                 '\u30af', // ク
	XK_kana_KE:                 '\u30b1', // ケ
	XK_kana_KO:                 '\u30b3', // コ
	XK_kana_SA:                 '\u30b5', // サ
	XK_kana_SHI:                '\u30b7', // シ
	XK_kana_SU:                 '\u30b9', // ス
	XK_kana_SE:                 '\u30bb', // セ
	XK_kana_SO:                 '\u30bd', // ソ
	XK_kana_TA:                 '\u30bf', // タ
	XK_kana_CHI:                '\u30c1', // チ
	XK_kana_TSU:                '\u30c4', // ツ
	XK_kana_TE:                 '\u30c6', // テ
	XK_kana_TO:                 '\u30c8', // ト
	XK_kana_NA:                 '\u30ca', // ナ
	XK_kana_NI:                 '\u30cb', // ニ
	XK_kana_NU:                 '\u30cc', // ヌ
	XK_kana_NE:                 '\u30cd', // ネ
	XK_kana_NO:                 '\u30ce', // ノ
	XK_kana_HA:                 '\u30cf', // ハ
	XK_kana_HI:                 '\u30d2', // ヒ
	XK_kana_FU:                 '\u30d5', // フ
	XK_kana_HE:                 '\u30d8', // ヘ
	XK_kana_HO:                 '\u30db', // ホ
	XK_kana_MA:                 '\u30de', // マ
	XK_kana_MI:                 '\u30df', // ミ
	XK_kana_MU:                 '\u30e0', // ム
	XK_kana_ME:                 '\u30e1', // メ
	XK_kana_MO:                 '\u30e2', // モ
	XK_kana_YA:                 '\u30e4', // ヤ
	XK_kana_YU:                 '\u30e6', // ユ
	XK_kana_YO:                 '\u30e8', // ヨ
	XK_kana_RA:                 '\u30e9', // ラ
	XK_kana_RI:                 '\u30ea', // リ
	XK_kana_RU:                 '\u30eb', // ル
	XK_kana_RE:                 '\u30ec', // レ
	XK_kana_RO:                 '\u30ed', // ロ
	XK_kana_WA:                 '\u30ef', // ワ
	XK_kana_N:                  '\u30f3', // ン
	XK_voicedsound:             '\u309b', // ゛
	XK_semivoicedsound:         '\u309c', // ゜
	XK_Farsi_0:                 '\u06f0', // ۰
	XK_Farsi_1:                 '\u06f1', // ۱
	XK_Farsi_2:                 '\u06f2', // ۲
	XK_Farsi_3:                 '\u06f3', // ۳
	XK_Farsi_4:                 '\u06f4', // ۴
	XK_Farsi_5:                 '\u06f5', // ۵
	XK_Farsi_6:                 '\u06f6', // ۶
	XK_Farsi_7:                 '\u06f7', // ۷
	XK_Farsi_8:                 '\u06f8', // ۸
	XK_Farsi_9:                 '\u06f9', // ۹
	XK_Arabic_percent:          '\u066a', // ٪
	XK_Arabic_superscript_alef: '\u0670', // ٰ
	XK_Arabic_tteh:             '\u0679', // ٹ
	XK_Arabic_peh:              '\u067e', // پ
	XK_Arabic_tcheh:            '\u0686', // چ
	XK_Arabic_ddal:             '\u0688', // ڈ
	XK_Arabic_rreh:             '\u0691', // ڑ
	XK_Arabic_comma:            '\u060c', // ،
	XK_Arabic_fullstop:         '\u06d4', // ۔
	XK_Arabic_0:                '\u0660', // ٠
	XK_Arabic_1:                '\u0661', // ١
	XK_Arabic_2:                '\u0662', // ٢
	XK_Arabic_3:                '\u0663', // ٣
	XK_Arabic_4:                '\u0664', // ٤
	XK_Arabic_5:                '\u0665', // ٥
	XK_Arabic_6:                '\u0666', // ٦
	XK_Arabic_7:                '\u0667', // ٧
	XK_Arabic_8:                '\u0668', // ٨
	XK_Arabic_9:                '\u0669', // ٩
	XK_Arabic_semicolon:        '\u061b', // ؛
	XK_Arabic_question_mark:    '\u061f', // ؟
	XK_Arabic_hamza:            '\u0621', // ء
	XK_Arabic_maddaonalef:      '\u0622', // آ
	XK_Arabic_hamzaonalef:      '\u0623', // أ
	XK_Arabic_hamzaonwaw:       '\u0624', // ؤ
	XK_Arabic_hamzaunderalef:   '\u0625', // إ
	XK_Arabic_hamzaonyeh:       '\u0626', // ئ
	XK_Arabic_alef:             '\u0627', // ا
	XK_Arabic_beh:              '\u0628', // ب
	XK_Arabic_tehmarbuta:       '\u0629', // ة
	XK_Arabic_teh:              '\u062a', // ت
	XK_Arabic_theh:             '\u062b', // ث
	XK_Arabic_jeem:             '\u062c', // ج
	XK_Arabic_hah:              '\u062d', // ح
	XK_Arabic_khah:             '\u062e', // خ
	XK_Arabic_dal:              '\u062f', // د
	XK_Arabic_thal:             '\u0630', // ذ
	XK_Arabic_ra:               '\u0631', // ر
	XK_Arabic_zain:             '\u0632', // ز
	XK_Arabic_seen:             '\u0633', // س
	XK_Arabic_sheen:            '\u0634', // ش
	XK_Arabic_sad:              '\u0635', // ص
	XK_Arabic_dad:              '\u0636', // ض
	XK_Arabic_tah:              '\u0637', // ط
	XK_Arabic_zah:              '\u0638', // ظ
	XK_Arabic_ain:              '\u0639', // ع
	XK_Arabic_ghain:            '\u063a', // غ
	XK_Arabic_tatweel:          '\u0640', // ـ
	XK_Arabic_feh:              '\u0641', // ف
	XK_Arabic_qaf:              '\u0642', // ق
	XK_Arabic_kaf:              '\u0643', // ك
	XK_Arabic_lam:              '\u0644', // ل
	XK_Arabic_meem:             '\u0645', // م
	XK_Arabic_noon:             '\u0646', // ن
	XK_Arabic_ha:               '\u0647', // ه
	XK_Arabic_waw:              '\u0648', // و
	XK_Arabic_alefmaksura:      '\u0649', // ى
	XK_Arabic_yeh:              '\u064a', // ي
	XK_Arabic_fathatan:         '\u064b', // ً
	XK_Arabic_dammatan:         '\u064c', // ٌ
	XK_Arabic_kasratan:         '\u064d', // ٍ
	XK_Arabic_fatha:            '\u064e', // َ
	XK_Arabic_damma:            '\u064f', // ُ
	XK_Arabic_kasra:            '\u0650', // ِ
	XK_Arabic_shadda:           '\u0651', // ّ
	XK_Arabic_sukun:            '\u0652', // ْ
	XK_Arabic_madda_above:      '\u0653', // ٓ
	XK_Arabic_hamza_above:      '\u0654', // ٔ
	XK_Arabic_hamza_below:      '\u0655', // ٕ
	XK_Arabic_jeh:              '\u0698', // ژ
	XK_Arabic_veh:              '\u06a4', // ڤ
	XK_Arabic_keheh:            '\u06a9', // ک
	XK_Arabic_gaf:              '\u06af', // گ
	XK_Arabic_noon_ghunna:      '\u06ba', // ں
	XK_Arabic_heh_doachashmee:  '\u06be', // ھ
	XK_Farsi_yeh:               '\u06cc', // ی
	// XK_Arabic_farsi_yeh == XK_Farsi_yeh # U+06CC ARABIC LETTER FARSI YEH
	XK_Arabic_yeh_baree:            '\u06d2', // ے
	XK_Arabic_heh_goal:             '\u06c1', // ہ
	XK_Cyrillic_GHE_bar:            '\u0492', // Ғ
	XK_Cyrillic_ghe_bar:            '\u0493', // ғ
	XK_Cyrillic_ZHE_descender:      '\u0496', // Җ
	XK_Cyrillic_zhe_descender:      '\u0497', // җ
	XK_Cyrillic_KA_descender:       '\u049a', // Қ
	XK_Cyrillic_ka_descender:       '\u049b', // қ
	XK_Cyrillic_KA_vertstroke:      '\u049c', // Ҝ
	XK_Cyrillic_ka_vertstroke:      '\u049d', // ҝ
	XK_Cyrillic_EN_descender:       '\u04a2', // Ң
	XK_Cyrillic_en_descender:       '\u04a3', // ң
	XK_Cyrillic_U_straight:         '\u04ae', // Ү
	XK_Cyrillic_u_straight:         '\u04af', // ү
	XK_Cyrillic_U_straight_bar:     '\u04b0', // Ұ
	XK_Cyrillic_u_straight_bar:     '\u04b1', // ұ
	XK_Cyrillic_HA_descender:       '\u04b2', // Ҳ
	XK_Cyrillic_ha_descender:       '\u04b3', // ҳ
	XK_Cyrillic_CHE_descender:      '\u04b6', // Ҷ
	XK_Cyrillic_che_descender:      '\u04b7', // ҷ
	XK_Cyrillic_CHE_vertstroke:     '\u04b8', // Ҹ
	XK_Cyrillic_che_vertstroke:     '\u04b9', // ҹ
	XK_Cyrillic_SHHA:               '\u04ba', // Һ
	XK_Cyrillic_shha:               '\u04bb', // һ
	XK_Cyrillic_SCHWA:              '\u04d8', // Ә
	XK_Cyrillic_schwa:              '\u04d9', // ә
	XK_Cyrillic_I_macron:           '\u04e2', // Ӣ
	XK_Cyrillic_i_macron:           '\u04e3', // ӣ
	XK_Cyrillic_O_bar:              '\u04e8', // Ө
	XK_Cyrillic_o_bar:              '\u04e9', // ө
	XK_Cyrillic_U_macron:           '\u04ee', // Ӯ
	XK_Cyrillic_u_macron:           '\u04ef', // ӯ
	XK_Serbian_dje:                 '\u0452', // ђ
	XK_Macedonia_gje:               '\u0453', // ѓ
	XK_Cyrillic_io:                 '\u0451', // ё
	XK_Ukrainian_ie:                '\u0454', // є
	XK_Macedonia_dse:               '\u0455', // ѕ
	XK_Ukrainian_i:                 '\u0456', // і
	XK_Ukrainian_yi:                '\u0457', // ї
	XK_Cyrillic_je:                 '\u0458', // ј
	XK_Cyrillic_lje:                '\u0459', // љ
	XK_Cyrillic_nje:                '\u045a', // њ
	XK_Serbian_tshe:                '\u045b', // ћ
	XK_Macedonia_kje:               '\u045c', // ќ
	XK_Ukrainian_ghe_with_upturn:   '\u0491', // ґ
	XK_Byelorussian_shortu:         '\u045e', // ў
	XK_Cyrillic_dzhe:               '\u045f', // џ
	XK_numerosign:                  '\u2116', // №
	XK_Serbian_DJE:                 '\u0402', // Ђ
	XK_Macedonia_GJE:               '\u0403', // Ѓ
	XK_Cyrillic_IO:                 '\u0401', // Ё
	XK_Ukrainian_IE:                '\u0404', // Є
	XK_Macedonia_DSE:               '\u0405', // Ѕ
	XK_Ukrainian_I:                 '\u0406', // І
	XK_Ukrainian_YI:                '\u0407', // Ї
	XK_Cyrillic_JE:                 '\u0408', // Ј
	XK_Cyrillic_LJE:                '\u0409', // Љ
	XK_Cyrillic_NJE:                '\u040a', // Њ
	XK_Serbian_TSHE:                '\u040b', // Ћ
	XK_Macedonia_KJE:               '\u040c', // Ќ
	XK_Ukrainian_GHE_WITH_UPTURN:   '\u0490', // Ґ
	XK_Byelorussian_SHORTU:         '\u040e', // Ў
	XK_Cyrillic_DZHE:               '\u040f', // Џ
	XK_Cyrillic_yu:                 '\u044e', // ю
	XK_Cyrillic_a:                  '\u0430', // а
	XK_Cyrillic_be:                 '\u0431', // б
	XK_Cyrillic_tse:                '\u0446', // ц
	XK_Cyrillic_de:                 '\u0434', // д
	XK_Cyrillic_ie:                 '\u0435', // е
	XK_Cyrillic_ef:                 '\u0444', // ф
	XK_Cyrillic_ghe:                '\u0433', // г
	XK_Cyrillic_ha:                 '\u0445', // х
	XK_Cyrillic_i:                  '\u0438', // и
	XK_Cyrillic_shorti:             '\u0439', // й
	XK_Cyrillic_ka:                 '\u043a', // к
	XK_Cyrillic_el:                 '\u043b', // л
	XK_Cyrillic_em:                 '\u043c', // м
	XK_Cyrillic_en:                 '\u043d', // н
	XK_Cyrillic_o:                  '\u043e', // о
	XK_Cyrillic_pe:                 '\u043f', // п
	XK_Cyrillic_ya:                 '\u044f', // я
	XK_Cyrillic_er:                 '\u0440', // р
	XK_Cyrillic_es:                 '\u0441', // с
	XK_Cyrillic_te:                 '\u0442', // т
	XK_Cyrillic_u:                  '\u0443', // у
	XK_Cyrillic_zhe:                '\u0436', // ж
	XK_Cyrillic_ve:                 '\u0432', // в
	XK_Cyrillic_softsign:           '\u044c', // ь
	XK_Cyrillic_yeru:               '\u044b', // ы
	XK_Cyrillic_ze:                 '\u0437', // з
	XK_Cyrillic_sha:                '\u0448', // ш
	XK_Cyrillic_e:                  '\u044d', // э
	XK_Cyrillic_shcha:              '\u0449', // щ
	XK_Cyrillic_che:                '\u0447', // ч
	XK_Cyrillic_hardsign:           '\u044a', // ъ
	XK_Cyrillic_YU:                 '\u042e', // Ю
	XK_Cyrillic_A:                  '\u0410', // А
	XK_Cyrillic_BE:                 '\u0411', // Б
	XK_Cyrillic_TSE:                '\u0426', // Ц
	XK_Cyrillic_DE:                 '\u0414', // Д
	XK_Cyrillic_IE:                 '\u0415', // Е
	XK_Cyrillic_EF:                 '\u0424', // Ф
	XK_Cyrillic_GHE:                '\u0413', // Г
	XK_Cyrillic_HA:                 '\u0425', // Х
	XK_Cyrillic_I:                  '\u0418', // И
	XK_Cyrillic_SHORTI:             '\u0419', // Й
	XK_Cyrillic_KA:                 '\u041a', // К
	XK_Cyrillic_EL:                 '\u041b', // Л
	XK_Cyrillic_EM:                 '\u041c', // М
	XK_Cyrillic_EN:                 '\u041d', // Н
	XK_Cyrillic_O:                  '\u041e', // О
	XK_Cyrillic_PE:                 '\u041f', // П
	XK_Cyrillic_YA:                 '\u042f', // Я
	XK_Cyrillic_ER:                 '\u0420', // Р
	XK_Cyrillic_ES:                 '\u0421', // С
	XK_Cyrillic_TE:                 '\u0422', // Т
	XK_Cyrillic_U:                  '\u0423', // У
	XK_Cyrillic_ZHE:                '\u0416', // Ж
	XK_Cyrillic_VE:                 '\u0412', // В
	XK_Cyrillic_SOFTSIGN:           '\u042c', // Ь
	XK_Cyrillic_YERU:               '\u042b', // Ы
	XK_Cyrillic_ZE:                 '\u0417', // З
	XK_Cyrillic_SHA:                '\u0428', // Ш
	XK_Cyrillic_E:                  '\u042d', // Э
	XK_Cyrillic_SHCHA:              '\u0429', // Щ
	XK_Cyrillic_CHE:                '\u0427', // Ч
	XK_Cyrillic_HARDSIGN:           '\u042a', // Ъ
	XK_Greek_ALPHAaccent:           '\u0386', // Ά
	XK_Greek_EPSILONaccent:         '\u0388', // Έ
	XK_Greek_ETAaccent:             '\u0389', // Ή
	XK_Greek_IOTAaccent:            '\u038a', // Ί
	XK_Greek_IOTAdieresis:          '\u03aa', // Ϊ
	XK_Greek_OMICRONaccent:         '\u038c', // Ό
	XK_Greek_UPSILONaccent:         '\u038e', // Ύ
	XK_Greek_UPSILONdieresis:       '\u03ab', // Ϋ
	XK_Greek_OMEGAaccent:           '\u038f', // Ώ
	XK_Greek_accentdieresis:        '\u0385', // ΅
	XK_Greek_horizbar:              '\u2015', // ―
	XK_Greek_alphaaccent:           '\u03ac', // ά
	XK_Greek_epsilonaccent:         '\u03ad', // έ
	XK_Greek_etaaccent:             '\u03ae', // ή
	XK_Greek_iotaaccent:            '\u03af', // ί
	XK_Greek_iotadieresis:          '\u03ca', // ϊ
	XK_Greek_iotaaccentdieresis:    '\u0390', // ΐ
	XK_Greek_omicronaccent:         '\u03cc', // ό
	XK_Greek_upsilonaccent:         '\u03cd', // ύ
	XK_Greek_upsilondieresis:       '\u03cb', // ϋ
	XK_Greek_upsilonaccentdieresis: '\u03b0', // ΰ
	XK_Greek_omegaaccent:           '\u03ce', // ώ
	XK_Greek_ALPHA:                 '\u0391', // Α
	XK_Greek_BETA:                  '\u0392', // Β
	XK_Greek_GAMMA:                 '\u0393', // Γ
	XK_Greek_DELTA:                 '\u0394', // Δ
	XK_Greek_EPSILON:               '\u0395', // Ε
	XK_Greek_ZETA:                  '\u0396', // Ζ
	XK_Greek_ETA:                   '\u0397', // Η
	XK_Greek_THETA:                 '\u0398', // Θ
	XK_Greek_IOTA:                  '\u0399', // Ι
	XK_Greek_KAPPA:                 '\u039a', // Κ
	XK_Greek_LAMDA:                 '\u039b', // Λ
	// XK_Greek_LAMBDA == XK_Greek_LAMDA # U+039B GREEK CAPITAL LETTER LAMDA
	XK_Greek_MU:      '\u039c', // Μ
	XK_Greek_NU:      '\u039d', // Ν
	XK_Greek_XI:      '\u039e', // Ξ
	XK_Greek_OMICRON: '\u039f', // Ο
	XK_Greek_PI:      '\u03a0', // Π
	XK_Greek_RHO:     '\u03a1', // Ρ
	XK_Greek_SIGMA:   '\u03a3', // Σ
	XK_Greek_TAU:     '\u03a4', // Τ
	XK_Greek_UPSILON: '\u03a5', // Υ
	XK_Greek_PHI:     '\u03a6', // Φ
	XK_Greek_CHI:     '\u03a7', // Χ
	XK_Greek_PSI:     '\u03a8', // Ψ
	XK_Greek_OMEGA:   '\u03a9', // Ω
	XK_Greek_alpha:   '\u03b1', // α
	XK_Greek_beta:    '\u03b2', // β
	XK_Greek_gamma:   '\u03b3', // γ
	XK_Greek_delta:   '\u03b4', // δ
	XK_Greek_epsilon: '\u03b5', // ε
	XK_Greek_zeta:    '\u03b6', // ζ
	XK_Greek_eta:     '\u03b7', // η
	XK_Greek_theta:   '\u03b8', // θ
	XK_Greek_iota:    '\u03b9', // ι
	XK_Greek_kappa:   '\u03ba', // κ
	XK_Greek_lamda:   '\u03bb', // λ
	// XK_Greek_lambda == XK_Greek_lamda # U+03BB GREEK SMALL LETTER LAMDA
	XK_Greek_mu:              '\u03bc', // μ
	XK_Greek_nu:              '\u03bd', // ν
	XK_Greek_xi:              '\u03be', // ξ
	XK_Greek_omicron:         '\u03bf', // ο
	XK_Greek_pi:              '\u03c0', // π
	XK_Greek_rho:             '\u03c1', // ρ
	XK_Greek_sigma:           '\u03c3', // σ
	XK_Greek_finalsmallsigma: '\u03c2', // ς
	XK_Greek_tau:             '\u03c4', // τ
	XK_Greek_upsilon:         '\u03c5', // υ
	XK_Greek_phi:             '\u03c6', // φ
	XK_Greek_chi:             '\u03c7', // χ
	XK_Greek_psi:             '\u03c8', // ψ
	XK_Greek_omega:           '\u03c9', // ω
	XK_hebrew_doublelowline:  '\u2017', // ‗
	XK_hebrew_aleph:          '\u05d0', // א
	XK_hebrew_bet:            '\u05d1', // ב
	XK_hebrew_gimel:          '\u05d2', // ג
	XK_hebrew_dalet:          '\u05d3', // ד
	XK_hebrew_he:             '\u05d4', // ה
	XK_hebrew_waw:            '\u05d5', // ו
	XK_hebrew_zain:           '\u05d6', // ז
	XK_hebrew_chet:           '\u05d7', // ח
	XK_hebrew_tet:            '\u05d8', // ט
	XK_hebrew_yod:            '\u05d9', // י
	XK_hebrew_finalkaph:      '\u05da', // ך
	XK_hebrew_kaph:           '\u05db', // כ
	XK_hebrew_lamed:          '\u05dc', // ל
	XK_hebrew_finalmem:       '\u05dd', // ם
	XK_hebrew_mem:            '\u05de', // מ
	XK_hebrew_finalnun:       '\u05df', // ן
	XK_hebrew_nun:            '\u05e0', // נ
	XK_hebrew_samech:         '\u05e1', // ס
	XK_hebrew_ayin:           '\u05e2', // ע
	XK_hebrew_finalpe:        '\u05e3', // ף
	XK_hebrew_pe:             '\u05e4', // פ
	XK_hebrew_finalzade:      '\u05e5', // ץ
	XK_hebrew_zade:           '\u05e6', // צ
	XK_hebrew_qoph:           '\u05e7', // ק
	XK_hebrew_resh:           '\u05e8', // ר
	XK_hebrew_shin:           '\u05e9', // ש
	XK_hebrew_taw:            '\u05ea', // ת
	XK_Thai_kokai:            '\u0e01', // ก
	XK_Thai_khokhai:          '\u0e02', // ข
	XK_Thai_khokhuat:         '\u0e03', // ฃ
	XK_Thai_khokhwai:         '\u0e04', // ค
	XK_Thai_khokhon:          '\u0e05', // ฅ
	XK_Thai_khorakhang:       '\u0e06', // ฆ
	XK_Thai_ngongu:           '\u0e07', // ง
	XK_Thai_chochan:          '\u0e08', // จ
	XK_Thai_choching:         '\u0e09', // ฉ
	XK_Thai_chochang:         '\u0e0a', // ช
	XK_Thai_soso:             '\u0e0b', // ซ
	XK_Thai_chochoe:          '\u0e0c', // ฌ
	XK_Thai_yoying:           '\u0e0d', // ญ
	XK_Thai_dochada:          '\u0e0e', // ฎ
	XK_Thai_topatak:          '\u0e0f', // ฏ
	XK_Thai_thothan:          '\u0e10', // ฐ
	XK_Thai_thonangmontho:    '\u0e11', // ฑ
	XK_Thai_thophuthao:       '\u0e12', // ฒ
	XK_Thai_nonen:            '\u0e13', // ณ
	XK_Thai_dodek:            '\u0e14', // ด
	XK_Thai_totao:            '\u0e15', // ต
	XK_Thai_thothung:         '\u0e16', // ถ
	XK_Thai_thothahan:        '\u0e17', // ท
	XK_Thai_thothong:         '\u0e18', // ธ
	XK_Thai_nonu:             '\u0e19', // น
	XK_Thai_bobaimai:         '\u0e1a', // บ
	XK_Thai_popla:            '\u0e1b', // ป
	XK_Thai_phophung:         '\u0e1c', // ผ
	XK_Thai_fofa:             '\u0e1d', // ฝ
	XK_Thai_phophan:          '\u0e1e', // พ
	XK_Thai_fofan:            '\u0e1f', // ฟ
	XK_Thai_phosamphao:       '\u0e20', // ภ
	XK_Thai_moma:             '\u0e21', // ม
	XK_Thai_yoyak:            '\u0e22', // ย
	XK_Thai_rorua:            '\u0e23', // ร
	XK_Thai_ru:               '\u0e24', // ฤ
	XK_Thai_loling:           '\u0e25', // ล
	XK_Thai_lu:               '\u0e26', // ฦ
	XK_Thai_wowaen:           '\u0e27', // ว
	XK_Thai_sosala:           '\u0e28', // ศ
	XK_Thai_sorusi:           '\u0e29', // ษ
	XK_Thai_sosua:            '\u0e2a', // ส
	XK_Thai_hohip:            '\u0e2b', // ห
	XK_Thai_lochula:          '\u0e2c', // ฬ
	XK_Thai_oang:             '\u0e2d', // อ
	XK_Thai_honokhuk:         '\u0e2e', // ฮ
	XK_Thai_paiyannoi:        '\u0e2f', // ฯ
	XK_Thai_saraa:            '\u0e30', // ะ
	XK_Thai_maihanakat:       '\u0e31', // ั
	XK_Thai_saraaa:           '\u0e32', // า
	XK_Thai_saraam:           '\u0e33', // ำ
	XK_Thai_sarai:            '\u0e34', // ิ
	XK_Thai_saraii:           '\u0e35', // ี
	XK_Thai_saraue:           '\u0e36', // ึ
	XK_Thai_sarauee:          '\u0e37', // ื
	XK_Thai_sarau:            '\u0e38', // ุ
	XK_Thai_sarauu:           '\u0e39', // ู
	XK_Thai_phinthu:          '\u0e3a', // ฺ
	XK_Thai_baht:             '\u0e3f', // ฿
	XK_Thai_sarae:            '\u0e40', // เ
	XK_Thai_saraae:           '\u0e41', // แ
	XK_Thai_sarao:            '\u0e42', // โ
	XK_Thai_saraaimaimuan:    '\u0e43', // ใ
	XK_Thai_saraaimaimalai:   '\u0e44', // ไ
	XK_Thai_lakkhangyao:      '\u0e45', // ๅ
	XK_Thai_maiyamok:         '\u0e46', // ๆ
	XK_Thai_maitaikhu:        '\u0e47', // ็
	XK_Thai_maiek:            '\u0e48', // ่
	XK_Thai_maitho:           '\u0e49', // ้
	XK_Thai_maitri:           '\u0e4a', // ๊
	XK_Thai_maichattawa:      '\u0e4b', // ๋
	XK_Thai_thanthakhat:      '\u0e4c', // ์
	XK_Thai_nikhahit:         '\u0e4d', // ํ
	XK_Thai_leksun:           '\u0e50', // ๐
	XK_Thai_leknung:          '\u0e51', // ๑
	XK_Thai_leksong:          '\u0e52', // ๒
	XK_Thai_leksam:           '\u0e53', // ๓
	XK_Thai_leksi:            '\u0e54', // ๔
	XK_Thai_lekha:            '\u0e55', // ๕
	XK_Thai_lekhok:           '\u0e56', // ๖
	XK_Thai_lekchet:          '\u0e57', // ๗
	XK_Thai_lekpaet:          '\u0e58', // ๘
	XK_Thai_lekkao:           '\u0e59', // ๙
	XK_Korean_Won:            '\u20a9', // ₩
	XK_Armenian_ligature_ew:  '\u0587', // և
	XK_Armenian_full_stop:    '\u0589', // ։
	// XK_Armenian_verjaket == XK_Armenian_full_stop # U+0589 ARMENIAN FULL STOP
	XK_Armenian_separation_mark: '\u055d', // ՝
	// XK_Armenian_but == XK_Armenian_separation_mark # U+055D ARMENIAN COMMA
	XK_Armenian_hyphen: '\u058a', // ֊
	// XK_Armenian_yentamna == XK_Armenian_hyphen # U+058A ARMENIAN HYPHEN
	XK_Armenian_exclam: '\u055c', // ՜
	// XK_Armenian_amanak == XK_Armenian_exclam # U+055C ARMENIAN EXCLAMATION MARK
	XK_Armenian_accent: '\u055b', // ՛
	// XK_Armenian_shesht == XK_Armenian_accent # U+055B ARMENIAN EMPHASIS MARK
	XK_Armenian_question: '\u055e', // ՞
	// XK_Armenian_paruyk == XK_Armenian_question # U+055E ARMENIAN QUESTION MARK
	XK_Armenian_AYB:        '\u0531', // Ա
	XK_Armenian_ayb:        '\u0561', // ա
	XK_Armenian_BEN:        '\u0532', // Բ
	XK_Armenian_ben:        '\u0562', // բ
	XK_Armenian_GIM:        '\u0533', // Գ
	XK_Armenian_gim:        '\u0563', // գ
	XK_Armenian_DA:         '\u0534', // Դ
	XK_Armenian_da:         '\u0564', // դ
	XK_Armenian_YECH:       '\u0535', // Ե
	XK_Armenian_yech:       '\u0565', // ե
	XK_Armenian_ZA:         '\u0536', // Զ
	XK_Armenian_za:         '\u0566', // զ
	XK_Armenian_E:          '\u0537', // Է
	XK_Armenian_e:          '\u0567', // է
	XK_Armenian_AT:         '\u0538', // Ը
	XK_Armenian_at:         '\u0568', // ը
	XK_Armenian_TO:         '\u0539', // Թ
	XK_Armenian_to:         '\u0569', // թ
	XK_Armenian_ZHE:        '\u053a', // Ժ
	XK_Armenian_zhe:        '\u056a', // ժ
	XK_Armenian_INI:        '\u053b', // Ի
	XK_Armenian_ini:        '\u056b', // ի
	XK_Armenian_LYUN:       '\u053c', // Լ
	XK_Armenian_lyun:       '\u056c', // լ
	XK_Armenian_KHE:        '\u053d', // Խ
	XK_Armenian_khe:        '\u056d', // խ
	XK_Armenian_TSA:        '\u053e', // Ծ
	XK_Armenian_tsa:        '\u056e', // ծ
	XK_Armenian_KEN:        '\u053f', // Կ
	XK_Armenian_ken:        '\u056f', // կ
	XK_Armenian_HO:         '\u0540', // Հ
	XK_Armenian_ho:         '\u0570', // հ
	XK_Armenian_DZA:        '\u0541', // Ձ
	XK_Armenian_dza:        '\u0571', // ձ
	XK_Armenian_GHAT:       '\u0542', // Ղ
	XK_Armenian_ghat:       '\u0572', // ղ
	XK_Armenian_TCHE:       '\u0543', // Ճ
	XK_Armenian_tche:       '\u0573', // ճ
	XK_Armenian_MEN:        '\u0544', // Մ
	XK_Armenian_men:        '\u0574', // մ
	XK_Armenian_HI:         '\u0545', // Յ
	XK_Armenian_hi:         '\u0575', // յ
	XK_Armenian_NU:         '\u0546', // Ն
	XK_Armenian_nu:         '\u0576', // ն
	XK_Armenian_SHA:        '\u0547', // Շ
	XK_Armenian_sha:        '\u0577', // շ
	XK_Armenian_VO:         '\u0548', // Ո
	XK_Armenian_vo:         '\u0578', // ո
	XK_Armenian_CHA:        '\u0549', // Չ
	XK_Armenian_cha:        '\u0579', // չ
	XK_Armenian_PE:         '\u054a', // Պ
	XK_Armenian_pe:         '\u057a', // պ
	XK_Armenian_JE:         '\u054b', // Ջ
	XK_Armenian_je:         '\u057b', // ջ
	XK_Armenian_RA:         '\u054c', // Ռ
	XK_Armenian_ra:         '\u057c', // ռ
	XK_Armenian_SE:         '\u054d', // Ս
	XK_Armenian_se:         '\u057d', // ս
	XK_Armenian_VEV:        '\u054e', // Վ
	XK_Armenian_vev:        '\u057e', // վ
	XK_Armenian_TYUN:       '\u054f', // Տ
	XK_Armenian_tyun:       '\u057f', // տ
	XK_Armenian_RE:         '\u0550', // Ր
	XK_Armenian_re:         '\u0580', // ր
	XK_Armenian_TSO:        '\u0551', // Ց
	XK_Armenian_tso:        '\u0581', // ց
	XK_Armenian_VYUN:       '\u0552', // Ւ
	XK_Armenian_vyun:       '\u0582', // ւ
	XK_Armenian_PYUR:       '\u0553', // Փ
	XK_Armenian_pyur:       '\u0583', // փ
	XK_Armenian_KE:         '\u0554', // Ք
	XK_Armenian_ke:         '\u0584', // ք
	XK_Armenian_O:          '\u0555', // Օ
	XK_Armenian_o:          '\u0585', // օ
	XK_Armenian_FE:         '\u0556', // Ֆ
	XK_Armenian_fe:         '\u0586', // ֆ
	XK_Armenian_apostrophe: '\u055a', // ՚
	XK_Georgian_an:         '\u10d0', // ა
	XK_Georgian_ban:        '\u10d1', // ბ
	XK_Georgian_gan:        '\u10d2', // გ
	XK_Georgian_don:        '\u10d3', // დ
	XK_Georgian_en:         '\u10d4', // ე
	XK_Georgian_vin:        '\u10d5', // ვ
	XK_Georgian_zen:        '\u10d6', // ზ
	XK_Georgian_tan:        '\u10d7', // თ
	XK_Georgian_in:         '\u10d8', // ი
	XK_Georgian_kan:        '\u10d9', // კ
	XK_Georgian_las:        '\u10da', // ლ
	XK_Georgian_man:        '\u10db', // მ
	XK_Georgian_nar:        '\u10dc', // ნ
	XK_Georgian_on:         '\u10dd', // ო
	XK_Georgian_par:        '\u10de', // პ
	XK_Georgian_zhar:       '\u10df', // ჟ
	XK_Georgian_rae:        '\u10e0', // რ
	XK_Georgian_san:        '\u10e1', // ს
	XK_Georgian_tar:        '\u10e2', // ტ
	XK_Georgian_un:         '\u10e3', // უ
	XK_Georgian_phar:       '\u10e4', // ფ
	XK_Georgian_khar:       '\u10e5', // ქ
	XK_Georgian_ghan:       '\u10e6', // ღ
	XK_Georgian_qar:        '\u10e7', // ყ
	XK_Georgian_shin:       '\u10e8', // შ
	XK_Georgian_chin:       '\u10e9', // ჩ
	XK_Georgian_can:        '\u10ea', // ც
	XK_Georgian_jil:        '\u10eb', // ძ
	XK_Georgian_cil:        '\u10ec', // წ
	XK_Georgian_char:       '\u10ed', // ჭ
	XK_Georgian_xan:        '\u10ee', // ხ
	XK_Georgian_jhan:       '\u10ef', // ჯ
	XK_Georgian_hae:        '\u10f0', // ჰ
	XK_Georgian_he:         '\u10f1', // ჱ
	XK_Georgian_hie:        '\u10f2', // ჲ
	XK_Georgian_we:         '\u10f3', // ჳ
	XK_Georgian_har:        '\u10f4', // ჴ
	XK_Georgian_hoe:        '\u10f5', // ჵ
	XK_Georgian_fi:         '\u10f6', // ჶ
	XK_Xabovedot:           '\u1e8a', // Ẋ
	XK_Ibreve:              '\u012c', // Ĭ
	XK_Zstroke:             '\u01b5', // Ƶ
	XK_Gcaron:              '\u01e6', // Ǧ
	XK_Ocaron:              '\u01d2', // ǒ
	XK_Obarred:             '\u019f', // Ɵ
	XK_xabovedot:           '\u1e8b', // ẋ
	XK_ibreve:              '\u012d', // ĭ
	XK_zstroke:             '\u01b6', // ƶ
	XK_gcaron:              '\u01e7', // ǧ
	XK_ocaron:              '\u01d2', // ǒ
	XK_obarred:             '\u0275', // ɵ
	XK_SCHWA:               '\u018f', // Ə
	XK_schwa:               '\u0259', // ə
	XK_EZH:                 '\u01b7', // Ʒ
	XK_ezh:                 '\u0292', // ʒ
	XK_Lbelowdot:           '\u1e36', // Ḷ
	XK_lbelowdot:           '\u1e37', // ḷ
	XK_Abelowdot:           '\u1ea0', // Ạ
	XK_abelowdot:           '\u1ea1', // ạ
	XK_Ahook:               '\u1ea2', // Ả
	XK_ahook:               '\u1ea3', // ả
	XK_Acircumflexacute:    '\u1ea4', // Ấ
	XK_acircumflexacute:    '\u1ea5', // ấ
	XK_Acircumflexgrave:    '\u1ea6', // Ầ
	XK_acircumflexgrave:    '\u1ea7', // ầ
	XK_Acircumflexhook:     '\u1ea8', // Ẩ
	XK_acircumflexhook:     '\u1ea9', // ẩ
	XK_Acircumflextilde:    '\u1eaa', // Ẫ
	XK_acircumflextilde:    '\u1eab', // ẫ
	XK_Acircumflexbelowdot: '\u1eac', // Ậ
	XK_acircumflexbelowdot: '\u1ead', // ậ
	XK_Abreveacute:         '\u1eae', // Ắ
	XK_abreveacute:         '\u1eaf', // ắ
	XK_Abrevegrave:         '\u1eb0', // Ằ
	XK_abrevegrave:         '\u1eb1', // ằ
	XK_Abrevehook:          '\u1eb2', // Ẳ
	XK_abrevehook:          '\u1eb3', // ẳ
	XK_Abrevetilde:         '\u1eb4', // Ẵ
	XK_abrevetilde:         '\u1eb5', // ẵ
	XK_Abrevebelowdot:      '\u1eb6', // Ặ
	XK_abrevebelowdot:      '\u1eb7', // ặ
	XK_Ebelowdot:           '\u1eb8', // Ẹ
	XK_ebelowdot:           '\u1eb9', // ẹ
	XK_Ehook:               '\u1eba', // Ẻ
	XK_ehook:               '\u1ebb', // ẻ
	XK_Etilde:              '\u1ebc', // Ẽ
	XK_etilde:              '\u1ebd', // ẽ
	XK_Ecircumflexacute:    '\u1ebe', // Ế
	XK_ecircumflexacute:    '\u1ebf', // ế
	XK_Ecircumflexgrave:    '\u1ec0', // Ề
	XK_ecircumflexgrave:    '\u1ec1', // ề
	XK_Ecircumflexhook:     '\u1ec2', // Ể
	XK_ecircumflexhook:     '\u1ec3', // ể
	XK_Ecircumflextilde:    '\u1ec4', // Ễ
	XK_ecircumflextilde:    '\u1ec5', // ễ
	XK_Ecircumflexbelowdot: '\u1ec6', // Ệ
	XK_ecircumflexbelowdot: '\u1ec7', // ệ
	XK_Ihook:               '\u1ec8', // Ỉ
	XK_ihook:               '\u1ec9', // ỉ
	XK_Ibelowdot:           '\u1eca', // Ị
	XK_ibelowdot:           '\u1ecb', // ị
	XK_Obelowdot:           '\u1ecc', // Ọ
	XK_obelowdot:           '\u1ecd', // ọ
	XK_Ohook:               '\u1ece', // Ỏ
	XK_ohook:               '\u1ecf', // ỏ
	XK_Ocircumflexacute:    '\u1ed0', // Ố
	XK_ocircumflexacute:    '\u1ed1', // ố
	XK_Ocircumflexgrave:    '\u1ed2', // Ồ
	XK_ocircumflexgrave:    '\u1ed3', // ồ
	XK_Ocircumflexhook:     '\u1ed4', // Ổ
	XK_ocircumflexhook:     '\u1ed5', // ổ
	XK_Ocircumflextilde:    '\u1ed6', // Ỗ
	XK_ocircumflextilde:    '\u1ed7', // ỗ
	XK_Ocircumflexbelowdot: '\u1ed8', // Ộ
	XK_ocircumflexbelowdot: '\u1ed9', // ộ
	XK_Ohornacute:          '\u1eda', // Ớ
	XK_ohornacute:          '\u1edb', // ớ
	XK_Ohorngrave:          '\u1edc', // Ờ
	XK_ohorngrave:          '\u1edd', // ờ
	XK_Ohornhook:           '\u1ede', // Ở
	XK_ohornhook:           '\u1edf', // ở
	XK_Ohorntilde:          '\u1ee0', // Ỡ
	XK_ohorntilde:          '\u1ee1', // ỡ
	XK_Ohornbelowdot:       '\u1ee2', // Ợ
	XK_ohornbelowdot:       '\u1ee3', // ợ
	XK_Ubelowdot:           '\u1ee4', // Ụ
	XK_ubelowdot:           '\u1ee5', // ụ
	XK_Uhook:               '\u1ee6', // Ủ
	XK_uhook:               '\u1ee7', // ủ
	XK_Uhornacute:          '\u1ee8', // Ứ
	XK_uhornacute:          '\u1ee9', // ứ
	XK_Uhorngrave:          '\u1eea', // Ừ
	XK_uhorngrave:          '\u1eeb', // ừ
	XK_Uhornhook:           '\u1eec', // Ử
	XK_uhornhook:           '\u1eed', // ử
	XK_Uhorntilde:          '\u1eee', // Ữ
	XK_uhorntilde:          '\u1eef', // ữ
	XK_Uhornbelowdot:       '\u1ef0', // Ự
	XK_uhornbelowdot:       '\u1ef1', // ự
	XK_Ybelowdot:           '\u1ef4', // Ỵ
	XK_ybelowdot:           '\u1ef5', // ỵ
	XK_Yhook:               '\u1ef6', // Ỷ
	XK_yhook:               '\u1ef7', // ỷ
	XK_Ytilde:              '\u1ef8', // Ỹ
	XK_ytilde:              '\u1ef9', // ỹ
	XK_Ohorn:               '\u01a0', // Ơ
	XK_ohorn:               '\u01a1', // ơ
	XK_Uhorn:               '\u01af', // Ư
	XK_uhorn:               '\u01b0', // ư
	XK_EcuSign:             '\u20a0', // ₠
	XK_ColonSign:           '\u20a1', // ₡
	XK_CruzeiroSign:        '\u20a2', // ₢
	XK_FFrancSign:          '\u20a3', // ₣
	XK_LiraSign:            '\u20a4', // ₤
	XK_MillSign:            '\u20a5', // ₥
	XK_NairaSign:           '\u20a6', // ₦
	XK_PesetaSign:          '\u20a7', // ₧
	XK_RupeeSign:           '\u20a8', // ₨
	XK_WonSign:             '\u20a9', // ₩
	XK_NewSheqelSign:       '\u20aa', // ₪
	XK_DongSign:            '\u20ab', // ₫
	XK_EuroSign:            '\u20ac', // €
	XK_zerosuperior:        '\u2070', // ⁰
	XK_foursuperior:        '\u2074', // ⁴
	XK_fivesuperior:        '\u2075', // ⁵
	XK_sixsuperior:         '\u2076', // ⁶
	XK_sevensuperior:       '\u2077', // ⁷
	XK_eightsuperior:       '\u2078', // ⁸
	XK_ninesuperior:        '\u2079', // ⁹
	XK_zerosubscript:       '\u2080', // ₀
	XK_onesubscript:        '\u2081', // ₁
	XK_twosubscript:        '\u2082', // ₂
	XK_threesubscript:      '\u2083', // ₃
	XK_foursubscript:       '\u2084', // ₄
	XK_fivesubscript:       '\u2085', // ₅
	XK_sixsubscript:        '\u2086', // ₆
	XK_sevensubscript:      '\u2087', // ₇
	XK_eightsubscript:      '\u2088', // ₈
	XK_ninesubscript:       '\u2089', // ₉
	XK_partdifferential:    '\u2202', // ∂
	XK_emptyset:            '\u2205', // ∅
	XK_elementof:           '\u2208', // ∈
	XK_notelementof:        '\u2209', // ∉
	XK_because:             '\u2235', // ∵
	XK_approxeq:            '\u2245', // ≅
	XK_notapproxeq:         '\u2247', // ≇
	XK_notidentical:        '\u2262', // ≢
	XK_stricteq:            '\u2263', // ≣
	XK_braille_blank:       '\u2800', // ⠀
	XK_braille_dots_1:      '\u2801', // ⠁
	XK_braille_dots_2:      '\u2802', // ⠂
	XK_braille_dots_12:     '\u2803', // ⠃
	XK_braille_dots_3:      '\u2804', // ⠄
	XK_braille_dots_13:     '\u2805', // ⠅
	XK_braille_dots_23:     '\u2806', // ⠆
	XK_braille_dots_123:    '\u2807', // ⠇
	XK_braille_dots_4:      '\u2808', // ⠈
	XK_braille_dots_14:     '\u2809', // ⠉
	XK_braille_dots_5:      '\u2810', // ⠐
	XK_braille_dots_15:     '\u2811', // ⠑
	XK_braille_dots_25:     '\u2812', // ⠒
	XK_braille_dots_125:    '\u2813', // ⠓
	XK_braille_dots_35:     '\u2814', // ⠔
	XK_braille_dots_135:    '\u2815', // ⠕
	XK_braille_dots_235:    '\u2816', // ⠖
	XK_braille_dots_1235:   '\u2817', // ⠗
	XK_braille_dots_45:     '\u2818', // ⠘
	XK_braille_dots_145:    '\u2819', // ⠙
	XK_braille_dots_6:      '\u2820', // ⠠
	XK_braille_dots_16:     '\u2821', // ⠡
	XK_braille_dots_26:     '\u2822', // ⠢
	XK_braille_dots_126:    '\u2823', // ⠣
	XK_braille_dots_36:     '\u2824', // ⠤
	XK_braille_dots_136:    '\u2825', // ⠥
	XK_braille_dots_236:    '\u2826', // ⠦
	XK_braille_dots_1236:   '\u2827', // ⠧
	XK_braille_dots_46:     '\u2828', // ⠨
	XK_braille_dots_146:    '\u2829', // ⠩
	XK_braille_dots_56:     '\u2830', // ⠰
	XK_braille_dots_156:    '\u2831', // ⠱
	XK_braille_dots_256:    '\u2832', // ⠲
	XK_braille_dots_1256:   '\u2833', // ⠳
	XK_braille_dots_356:    '\u2834', // ⠴
	XK_braille_dots_1356:   '\u2835', // ⠵
	XK_braille_dots_2356:   '\u2836', // ⠶
	XK_braille_dots_12356:  '\u2837', // ⠷
	XK_braille_dots_456:    '\u2838', // ⠸
	XK_braille_dots_1456:   '\u2839', // ⠹
	XK_braille_dots_7:      '\u2840', // ⡀
	XK_braille_dots_17:     '\u2841', // ⡁
	XK_braille_dots_27:     '\u2842', // ⡂
	XK_braille_dots_127:    '\u2843', // ⡃
	XK_braille_dots_37:     '\u2844', // ⡄
	XK_braille_dots_137:    '\u2845', // ⡅
	XK_braille_dots_237:    '\u2846', // ⡆
	XK_braille_dots_1237:   '\u2847', // ⡇
	XK_braille_dots_47:     '\u2848', // ⡈
	XK_braille_dots_147:    '\u2849', // ⡉
	XK_braille_dots_57:     '\u2850', // ⡐
	XK_braille_dots_157:    '\u2851', // ⡑
	XK_braille_dots_257:    '\u2852', // ⡒
	XK_braille_dots_1257:   '\u2853', // ⡓
	XK_braille_dots_357:    '\u2854', // ⡔
	XK_braille_dots_1357:   '\u2855', // ⡕
	XK_braille_dots_2357:   '\u2856', // ⡖
	XK_braille_dots_12357:  '\u2857', // ⡗
	XK_braille_dots_457:    '\u2858', // ⡘
	XK_braille_dots_1457:   '\u2859', // ⡙
	XK_braille_dots_67:     '\u2860', // ⡠
	XK_braille_dots_167:    '\u2861', // ⡡
	XK_braille_dots_267:    '\u2862', // ⡢
	XK_braille_dots_1267:   '\u2863', // ⡣
	XK_braille_dots_367:    '\u2864', // ⡤
	XK_braille_dots_1367:   '\u2865', // ⡥
	XK_braille_dots_2367:   '\u2866', // ⡦
	XK_braille_dots_12367:  '\u2867', // ⡧
	XK_braille_dots_467:    '\u2868', // ⡨
	XK_braille_dots_1467:   '\u2869', // ⡩
	XK_braille_dots_567:    '\u2870', // ⡰
	XK_braille_dots_1567:   '\u2871', // ⡱
	XK_braille_dots_2567:   '\u2872', // ⡲
	XK_braille_dots_12567:  '\u2873', // ⡳
	XK_braille_dots_3567:   '\u2874', // ⡴
	XK_braille_dots_13567:  '\u2875', // ⡵
	XK_braille_dots_23567:  '\u2876', // ⡶
	XK_braille_dots_123567: '\u2877', // ⡷
	XK_braille_dots_4567:   '\u2878', // ⡸
	XK_braille_dots_14567:  '\u2879', // ⡹
	XK_braille_dots_8:      '\u2880', // ⢀
	XK_braille_dots_18:     '\u2881', // ⢁
	XK_braille_dots_28:     '\u2882', // ⢂
	XK_braille_dots_128:    '\u2883', // ⢃
	XK_braille_dots_38:     '\u2884', // ⢄
	XK_braille_dots_138:    '\u2885', // ⢅
	XK_braille_dots_238:    '\u2886', // ⢆
	XK_braille_dots_1238:   '\u2887', // ⢇
	XK_braille_dots_48:     '\u2888', // ⢈
	XK_braille_dots_148:    '\u2889', // ⢉
	XK_braille_dots_58:     '\u2890', // ⢐
	XK_braille_dots_158:    '\u2891', // ⢑
	XK_braille_dots_258:    '\u2892', // ⢒
	XK_braille_dots_1258:   '\u2893', // ⢓
	XK_braille_dots_358:    '\u2894', // ⢔
	XK_braille_dots_1358:   '\u2895', // ⢕
	XK_braille_dots_2358:   '\u2896', // ⢖
	XK_braille_dots_12358:  '\u2897', // ⢗
	XK_braille_dots_458:    '\u2898', // ⢘
	XK_braille_dots_1458:   '\u2899', // ⢙
	XK_Sinh_ng:             '\u0d82', // ං
	XK_Sinh_h2:             '\u0d83', // ඃ
	XK_Sinh_a:              '\u0d85', // අ
	XK_Sinh_aa:             '\u0d86', // ආ
	XK_Sinh_ae:             '\u0d87', // ඇ
	XK_Sinh_aee:            '\u0d88', // ඈ
	XK_Sinh_i:              '\u0d89', // ඉ
	XK_Sinh_ii:             '\u0d8a', // ඊ
	XK_Sinh_u:              '\u0d8b', // උ
	XK_Sinh_uu:             '\u0d8c', // ඌ
	XK_Sinh_ri:             '\u0d8d', // ඍ
	XK_Sinh_rii:            '\u0d8e', // ඎ
	XK_Sinh_lu:             '\u0d8f', // ඏ
	XK_Sinh_luu:            '\u0d90', // ඐ
	XK_Sinh_e:              '\u0d91', // එ
	XK_Sinh_ee:             '\u0d92', // ඒ
	XK_Sinh_ai:             '\u0d93', // ඓ
	XK_Sinh_o:              '\u0d94', // ඔ
	XK_Sinh_oo:             '\u0d95', // ඕ
	XK_Sinh_au:             '\u0d96', // ඖ
	XK_Sinh_ka:             '\u0d9a', // ක
	XK_Sinh_kha:            '\u0d9b', // ඛ
	XK_Sinh_ga:             '\u0d9c', // ග
	XK_Sinh_gha:            '\u0d9d', // ඝ
	XK_Sinh_ng2:            '\u0d9e', // ඞ
	XK_Sinh_nga:            '\u0d9f', // ඟ
	XK_Sinh_ca:             '\u0da0', // ච
	XK_Sinh_cha:            '\u0da1', // ඡ
	XK_Sinh_ja:             '\u0da2', // ජ
	XK_Sinh_jha:            '\u0da3', // ඣ
	XK_Sinh_nya:            '\u0da4', // ඤ
	XK_Sinh_jnya:           '\u0da5', // ඥ
	XK_Sinh_nja:            '\u0da6', // ඦ
	XK_Sinh_tta:            '\u0da7', // ට
	XK_Sinh_ttha:           '\u0da8', // ඨ
	XK_Sinh_dda:            '\u0da9', // ඩ
	XK_Sinh_ddha:           '\u0daa', // ඪ
	XK_Sinh_nna:            '\u0dab', // ණ
	XK_Sinh_ndda:           '\u0dac', // ඬ
	XK_Sinh_tha:            '\u0dad', // ත
	XK_Sinh_thha:           '\u0dae', // ථ
	XK_Sinh_dha:            '\u0daf', // ද
	XK_Sinh_dhha:           '\u0db0', // ධ
	XK_Sinh_na:             '\u0db1', // න
	XK_Sinh_ndha:           '\u0db3', // ඳ
	XK_Sinh_pa:             '\u0db4', // ප
	XK_Sinh_pha:            '\u0db5', // ඵ
	XK_Sinh_ba:             '\u0db6', // බ
	XK_Sinh_bha:            '\u0db7', // භ
	XK_Sinh_ma:             '\u0db8', // ම
	XK_Sinh_mba:            '\u0db9', // ඹ
	XK_Sinh_ya:             '\u0dba', // ය
	XK_Sinh_ra:             '\u0dbb', // ර
	XK_Sinh_la:             '\u0dbd', // ල
	XK_Sinh_va:             '\u0dc0', // ව
	XK_Sinh_sha:            '\u0dc1', // ශ
	XK_Sinh_ssha:           '\u0dc2', // ෂ
	XK_Sinh_sa:             '\u0dc3', // ස
	XK_Sinh_ha:             '\u0dc4', // හ
	XK_Sinh_lla:            '\u0dc5', // ළ
	XK_Sinh_fa:             '\u0dc6', // ෆ
	XK_Sinh_al:             '\u0dca', // ්
	XK_Sinh_aa2:            '\u0dcf', // ා
	XK_Sinh_ae2:            '\u0dd0', // ැ
	XK_Sinh_aee2:           '\u0dd1', // ෑ
	XK_Sinh_i2:             '\u0dd2', // ි
	XK_Sinh_ii2:            '\u0dd3', // ී
	XK_Sinh_u2:             '\u0dd4', // ු
	XK_Sinh_uu2:            '\u0dd6', // ූ
	XK_Sinh_ru2:            '\u0dd8', // ෘ
	XK_Sinh_e2:             '\u0dd9', // ෙ
	XK_Sinh_ee2:            '\u0dda', // ේ
	XK_Sinh_ai2:            '\u0ddb', // ෛ
	XK_Sinh_o2:             '\u0ddc', // ො
	XK_Sinh_oo2:            '\u0ddd', // ෝ
	XK_Sinh_au2:            '\u0dde', // ෞ
	XK_Sinh_lu2:            '\u0ddf', // ෟ
	XK_Sinh_ruu2:           '\u0df2', // ෲ
	XK_Sinh_luu2:           '\u0df3', // ෳ
	XK_Sinh_kunddaliya:     '\u0df4', // ෴
}
