// qmk-go - go client library for VIA-enabled QMK keyboards
// Copyright (c) 2022 Ian McLinden. All rights reserved
//
// This file is released under GNU LGPL 2.1 on Linux,
// and under the 3-clause BSD license on all other platforms

package keycode

import (
	"errors"
	"strings"
)

var (
	ErrorUnknownKeycode = errors.New("unknown keycode")
)

type Keycode uint16

func KeycodeFromBytes(msb byte, lsb byte) Keycode {
	return Keycode(uint16(msb)<<8 | uint16(lsb))
}

func (k Keycode) ToBytes() []byte {
	return []byte{byte((k >> 8) & 0xFF), byte(k & 0xFF)}
}

const (
	KC_UNKNOWN Keycode = 0x00
)

const (
	KC_TRANSPARENT Keycode = 0x01
	KC_TRNS                = KC_TRANSPARENT
)

// Punctuation
const (
	KC_ENT  Keycode = KC_ENTER
	KC_ESC          = KC_ESCAPE
	KC_BSPC         = KC_BACKSPACE
	KC_SPC          = KC_SPACE
	KC_MINS         = KC_MINUS
	KC_EQL          = KC_EQUAL
	KC_LBRC         = KC_LEFT_BRACKET
	KC_RBRC         = KC_RIGHT_BRACKET
	KC_BSLS         = KC_BACKSLASH
	KC_NUHS         = KC_NONUS_HASH
	KC_SCLN         = KC_SEMICOLON
	KC_QUOT         = KC_QUOTE
	KC_GRV          = KC_GRAVE
	KC_COMM         = KC_COMMA
	KC_SLSH         = KC_SLASH
	KC_NUBS         = KC_NONUS_BACKSLASH
)

// Lock Keys
const (
	KC_CAPS Keycode = KC_CAPS_LOCK
	KC_SCRL         = KC_SCROLL_LOCK
	KC_NUM          = KC_NUM_LOCK
	KC_LCAP         = KC_LOCKING_CAPS_LOCK
	KC_LNUM         = KC_LOCKING_NUM_LOCK
	KC_LSCR         = KC_LOCKING_SCROLL_LOCK
)

// Commands
const (
	KC_PSCR Keycode = KC_PRINT_SCREEN
	KC_PAUS         = KC_PAUSE
	KC_BRK          = KC_PAUSE
	KC_INS          = KC_INSERT
	KC_PGUP         = KC_PAGE_UP
	KC_DEL          = KC_DELETE
	KC_PGDN         = KC_PAGE_DOWN
	KC_RGHT         = KC_RIGHT
	KC_APP          = KC_APPLICATION
	KC_EXEC         = KC_EXECUTE
	KC_SLCT         = KC_SELECT
	KC_AGIN         = KC_AGAIN
	KC_PSTE         = KC_PASTE
	KC_ERAS         = KC_ALTERNATE_ERASE
	KC_SYRQ         = KC_SYSTEM_REQUEST
	KC_CNCL         = KC_CANCEL
	KC_CLR          = KC_CLEAR
	KC_PRIR         = KC_PRIOR
	KC_RETN         = KC_RETURN
	KC_SEPR         = KC_SEPARATOR
	KC_CLAG         = KC_CLEAR_AGAIN
	KC_CRSL         = KC_CRSEL
	KC_EXSL         = KC_EXSEL
)

// Keypad
const (
	KC_PSLS Keycode = KC_KP_SLASH
	KC_PAST         = KC_KP_ASTERISK
	KC_PMNS         = KC_KP_MINUS
	KC_PPLS         = KC_KP_PLUS
	KC_PENT         = KC_KP_ENTER
	KC_P1           = KC_KP_1
	KC_P2           = KC_KP_2
	KC_P3           = KC_KP_3
	KC_P4           = KC_KP_4
	KC_P5           = KC_KP_5
	KC_P6           = KC_KP_6
	KC_P7           = KC_KP_7
	KC_P8           = KC_KP_8
	KC_P9           = KC_KP_9
	KC_P0           = KC_KP_0
	KC_PDOT         = KC_KP_DOT
	KC_PEQL         = KC_KP_EQUAL
	KC_PCMM         = KC_KP_COMMA
)

// Language Specific
const (
	KC_INT1 Keycode = KC_INTERNATIONAL_1
	KC_INT2         = KC_INTERNATIONAL_2
	KC_INT3         = KC_INTERNATIONAL_3
	KC_INT4         = KC_INTERNATIONAL_4
	KC_INT5         = KC_INTERNATIONAL_5
	KC_INT6         = KC_INTERNATIONAL_6
	KC_INT7         = KC_INTERNATIONAL_7
	KC_INT8         = KC_INTERNATIONAL_8
	KC_INT9         = KC_INTERNATIONAL_9
	KC_LNG1         = KC_LANGUAGE_1
	KC_LNG2         = KC_LANGUAGE_2
	KC_LNG3         = KC_LANGUAGE_3
	KC_LNG4         = KC_LANGUAGE_4
	KC_LNG5         = KC_LANGUAGE_5
	KC_LNG6         = KC_LANGUAGE_6
	KC_LNG7         = KC_LANGUAGE_7
	KC_LNG8         = KC_LANGUAGE_8
	KC_LNG9         = KC_LANGUAGE_9
)

// Modifiers
const (
	KC_LCTL Keycode = KC_LEFT_CTRL
	KC_LSFT         = KC_LEFT_SHIFT
	KC_LALT         = KC_LEFT_ALT
	KC_LOPT         = KC_LEFT_ALT
	KC_LGUI         = KC_LEFT_GUI
	KC_LCMD         = KC_LEFT_GUI
	KC_LWIN         = KC_LEFT_GUI
	KC_RCTL         = KC_RIGHT_CTRL
	KC_RSFT         = KC_RIGHT_SHIFT
	KC_RALT         = KC_RIGHT_ALT
	KC_ALGR         = KC_RIGHT_ALT
	KC_ROPT         = KC_RIGHT_ALT
	KC_RGUI         = KC_RIGHT_GUI
	KC_RCMD         = KC_RIGHT_GUI
	KC_RWIN         = KC_RIGHT_GUI
)

// Generic Desktop Page (0x01)
const (
	KC_PWR  Keycode = KC_SYSTEM_POWER
	KC_SLEP         = KC_SYSTEM_SLEEP
	KC_WAKE         = KC_SYSTEM_WAKE
)

// Consumer Page (0x0C)
const (
	KC_MUTE Keycode = KC_AUDIO_MUTE
	KC_VOLU         = KC_AUDIO_VOL_UP
	KC_VOLD         = KC_AUDIO_VOL_DOWN
	KC_MNXT         = KC_MEDIA_NEXT_TRACK
	KC_MPRV         = KC_MEDIA_PREV_TRACK
	KC_MSTP         = KC_MEDIA_STOP
	KC_MPLY         = KC_MEDIA_PLAY_PAUSE
	KC_MSEL         = KC_MEDIA_SELECT
	KC_EJCT         = KC_MEDIA_EJECT
	KC_CALC         = KC_CALCULATOR
	KC_MYCM         = KC_MY_COMPUTER
	KC_WSCH         = KC_WWW_SEARCH
	KC_WHOM         = KC_WWW_HOME
	KC_WBAK         = KC_WWW_BACK
	KC_WFWD         = KC_WWW_FORWARD
	KC_WSTP         = KC_WWW_STOP
	KC_WREF         = KC_WWW_REFRESH
	KC_WFAV         = KC_WWW_FAVORITES
	KC_MFFD         = KC_MEDIA_FAST_FORWARD
	KC_MRWD         = KC_MEDIA_REWIND
	KC_BRIU         = KC_BRIGHTNESS_UP
	KC_BRID         = KC_BRIGHTNESS_DOWN
)

// System Specific
const (
	KC_BRMU Keycode = KC_PAUSE
	KC_BRMD         = KC_SCROLL_LOCK
)

// Mouse Keys
const (
	KC_MS_U Keycode = KC_MS_UP
	KC_MS_D         = KC_MS_DOWN
	KC_MS_L         = KC_MS_LEFT
	KC_MS_R         = KC_MS_RIGHT
	KC_BTN1         = KC_MS_BTN1
	KC_BTN2         = KC_MS_BTN2
	KC_BTN3         = KC_MS_BTN3
	KC_BTN4         = KC_MS_BTN4
	KC_BTN5         = KC_MS_BTN5
	KC_BTN6         = KC_MS_BTN6
	KC_BTN7         = KC_MS_BTN7
	KC_BTN8         = KC_MS_BTN8
	KC_WH_U         = KC_MS_WH_UP
	KC_WH_D         = KC_MS_WH_DOWN
	KC_WH_L         = KC_MS_WH_LEFT
	KC_WH_R         = KC_MS_WH_RIGHT
	KC_ACL0         = KC_MS_ACCEL0
	KC_ACL1         = KC_MS_ACCEL1
	KC_ACL2         = KC_MS_ACCEL2
)

// Media and Function keys
const (
	// Generic Desktop Page (0x01)
	KC_SYSTEM_POWER Keycode = iota + 0xA5
	KC_SYSTEM_SLEEP
	KC_SYSTEM_WAKE

	// Consumer Page (0x0C)
	KC_AUDIO_MUTE
	KC_AUDIO_VOL_UP
	KC_AUDIO_VOL_DOWN
	KC_MEDIA_NEXT_TRACK
	KC_MEDIA_PREV_TRACK
	KC_MEDIA_STOP
	KC_MEDIA_PLAY_PAUSE
	KC_MEDIA_SELECT
	KC_MEDIA_EJECT // 0xB0
	KC_MAIL
	KC_CALCULATOR
	KC_MY_COMPUTER
	KC_WWW_SEARCH
	KC_WWW_HOME
	KC_WWW_BACK
	KC_WWW_FORWARD
	KC_WWW_STOP
	KC_WWW_REFRESH
	KC_WWW_FAVORITES
	KC_MEDIA_FAST_FORWARD
	KC_MEDIA_REWIND
	KC_BRIGHTNESS_UP
	KC_BRIGHTNESS_DOWN
)

// Keyboard/Keypad Page (0x07)
const (
	KC_NO Keycode = iota
	KC_ROLL_OVER
	KC_POST_FAIL
	KC_UNDEFINED
	KC_A
	KC_B
	KC_C
	KC_D
	KC_E
	KC_F
	KC_G
	KC_H
	KC_I
	KC_J
	KC_K
	KC_L
	KC_M // 0x10
	KC_N
	KC_O
	KC_P
	KC_Q
	KC_R
	KC_S
	KC_T
	KC_U
	KC_V
	KC_W
	KC_X
	KC_Y
	KC_Z
	KC_1
	KC_2
	KC_3 // 0x20
	KC_4
	KC_5
	KC_6
	KC_7
	KC_8
	KC_9
	KC_0
	KC_ENTER
	KC_ESCAPE
	KC_BACKSPACE
	KC_TAB
	KC_SPACE
	KC_MINUS
	KC_EQUAL
	KC_LEFT_BRACKET
	KC_RIGHT_BRACKET // 0x30
	KC_BACKSLASH
	KC_NONUS_HASH
	KC_SEMICOLON
	KC_QUOTE
	KC_GRAVE
	KC_COMMA
	KC_DOT
	KC_SLASH
	KC_CAPS_LOCK
	KC_F1
	KC_F2
	KC_F3
	KC_F4
	KC_F5
	KC_F6
	KC_F7 // 0x40
	KC_F8
	KC_F9
	KC_F10
	KC_F11
	KC_F12
	KC_PRINT_SCREEN
	KC_SCROLL_LOCK
	KC_PAUSE
	KC_INSERT
	KC_HOME
	KC_PAGE_UP
	KC_DELETE
	KC_END
	KC_PAGE_DOWN
	KC_RIGHT
	KC_LEFT // 0x50
	KC_DOWN
	KC_UP
	KC_NUM_LOCK
	KC_KP_SLASH
	KC_KP_ASTERISK
	KC_KP_MINUS
	KC_KP_PLUS
	KC_KP_ENTER
	KC_KP_1
	KC_KP_2
	KC_KP_3
	KC_KP_4
	KC_KP_5
	KC_KP_6
	KC_KP_7
	KC_KP_8 // 0x60
	KC_KP_9
	KC_KP_0
	KC_KP_DOT
	KC_NONUS_BACKSLASH
	KC_APPLICATION
	KC_KB_POWER
	KC_KP_EQUAL
	KC_F13
	KC_F14
	KC_F15
	KC_F16
	KC_F17
	KC_F18
	KC_F19
	KC_F20
	KC_F21 // 0x70
	KC_F22
	KC_F23
	KC_F24
	KC_EXECUTE
	KC_HELP
	KC_MENU
	KC_SELECT
	KC_STOP
	KC_AGAIN
	KC_UNDO
	KC_CUT
	KC_COPY
	KC_PASTE
	KC_FIND
	KC_KB_MUTE
	KC_KB_VOLUME_UP // 0x80
	KC_KB_VOLUME_DOWN
	KC_LOCKING_CAPS_LOCK
	KC_LOCKING_NUM_LOCK
	KC_LOCKING_SCROLL_LOCK
	KC_KP_COMMA
	KC_KP_EQUAL_AS400
	KC_INTERNATIONAL_1
	KC_INTERNATIONAL_2
	KC_INTERNATIONAL_3
	KC_INTERNATIONAL_4
	KC_INTERNATIONAL_5
	KC_INTERNATIONAL_6
	KC_INTERNATIONAL_7
	KC_INTERNATIONAL_8
	KC_INTERNATIONAL_9
	KC_LANGUAGE_1 // 0x90
	KC_LANGUAGE_2
	KC_LANGUAGE_3
	KC_LANGUAGE_4
	KC_LANGUAGE_5
	KC_LANGUAGE_6
	KC_LANGUAGE_7
	KC_LANGUAGE_8
	KC_LANGUAGE_9
	KC_ALTERNATE_ERASE
	KC_SYSTEM_REQUEST
	KC_CANCEL
	KC_CLEAR
	KC_PRIOR
	KC_RETURN
	KC_SEPARATOR
	KC_OUT // 0xA0
	KC_OPER
	KC_CLEAR_AGAIN
	KC_CRSEL
	KC_EXSEL
)

// Modifiers
const (
	KC_LEFT_CTRL Keycode = iota + 0xE0
	KC_LEFT_SHIFT
	KC_LEFT_ALT
	KC_LEFT_GUI
	KC_RIGHT_CTRL
	KC_RIGHT_SHIFT
	KC_RIGHT_ALT
	KC_RIGHT_GUI

	// **********************************************
	// * 0xF0-0xFF are unallocated in the HID spec. *
	// * QMK uses these for Mouse Keys - see below. *
	// **********************************************
)

// Fn keys
const (
	KC_FN0 Keycode = iota + 0xC0
	KC_FN1
	KC_FN2
	KC_FN3
	KC_FN4
	KC_FN5
	KC_FN6
	KC_FN7
	KC_FN8
	KC_FN9
	KC_FN10
	KC_FN11
	KC_FN12
	KC_FN13
	KC_FN14
	KC_FN15
	KC_FN16 // 0xD0
	KC_FN17
	KC_FN18
	KC_FN19
	KC_FN20
	KC_FN21
	KC_FN22
	KC_FN23
	KC_FN24
	KC_FN25
	KC_FN26
	KC_FN27
	KC_FN28
	KC_FN29
	KC_FN30
	KC_FN31
)

// Mouse Buttons
const (
	KC_MS_UP Keycode = iota + 0xF0
	KC_MS_DOWN
	KC_MS_LEFT
	KC_MS_RIGHT
	KC_MS_BTN1
	KC_MS_BTN2
	KC_MS_BTN3
	KC_MS_BTN4
	KC_MS_BTN5
	KC_MS_BTN6 = KC_MS_BTN5
	KC_MS_BTN7 = KC_MS_BTN5
	KC_MS_BTN8 = KC_MS_BTN5
)

// Mouse Wheel
const (
	KC_MS_WH_UP Keycode = iota + 0xF9
	KC_MS_WH_DOWN
	KC_MS_WH_LEFT
	KC_MS_WH_RIGHT

	// Acceleration
	KC_MS_ACCEL0
	KC_MS_ACCEL1
	KC_MS_ACCEL2 // 0xFF
)

// Legacy / Deprecated
const (
	KC_BSPACE         Keycode = KC_BACKSPACE
	KC_LBRACKET               = KC_LEFT_BRACKET
	KC_RBRACKET               = KC_RIGHT_BRACKET
	KC_BSLASH                 = KC_BACKSLASH
	KC_SCOLON                 = KC_SEMICOLON
	KC_CAPSLOCK               = KC_CAPS_LOCK
	KC_PSCREEN                = KC_PRINT_SCREEN
	KC_SCROLLLOCK             = KC_SCROLL_LOCK
	KC_PGDOWN                 = KC_PAGE_DOWN
	KC_NUMLOCK                = KC_NUM_LOCK
	KC_NONUS_BSLASH           = KC_NONUS_BACKSLASH
	KC_POWER                  = KC_KB_POWER
	KC__MUTE                  = KC_KB_MUTE
	KC__VOLUP                 = KC_KB_VOLUME_UP
	KC__VOLDOWN               = KC_KB_VOLUME_DOWN
	KC_LOCKING_CAPS           = KC_LOCKING_CAPS_LOCK
	KC_LOCKING_NUM            = KC_LOCKING_NUM_LOCK
	KC_LOCKING_SCROLL         = KC_LOCKING_SCROLL_LOCK
	KC_LANG1                  = KC_LANGUAGE_1
	KC_LANG2                  = KC_LANGUAGE_2
	KC_LANG3                  = KC_LANGUAGE_3
	KC_LANG4                  = KC_LANGUAGE_4
	KC_LANG5                  = KC_LANGUAGE_5
	KC_LANG6                  = KC_LANGUAGE_6
	KC_LANG7                  = KC_LANGUAGE_7
	KC_LANG8                  = KC_LANGUAGE_8
	KC_LANG9                  = KC_LANGUAGE_9
	KC_ALT_ERASE              = KC_ALTERNATE_ERASE
	KC_SYSREQ                 = KC_SYSTEM_REQUEST

	KC_LCTRL  = KC_LEFT_CTRL
	KC_LSHIFT = KC_LEFT_SHIFT
	KC_RCTRL  = KC_RIGHT_CTRL
	KC_RSHIFT = KC_RIGHT_SHIFT

	KC_ZKHK = KC_GRAVE
	KC_RO   = KC_INTERNATIONAL_1
	KC_KANA = KC_INTERNATIONAL_2
	KC_JYEN = KC_INTERNATIONAL_3
	KC_HENK = KC_INTERNATIONAL_4
	KC_MHEN = KC_INTERNATIONAL_5
	KC_HAEN = KC_LANGUAGE_1
	KC_HANJ = KC_LANGUAGE_2

	KC_CLCK = KC_CAPS_LOCK
	KC_SLCK = KC_SCROLL_LOCK
	KC_NLCK = KC_NUM_LOCK
)

// VIA keycodes
const (
	FN_MO13 Keycode = iota + 0x5F10
	FN_MO23
	MACRO00
	MACRO01
	MACRO02
	MACRO03
	MACRO04
	MACRO05
	MACRO06
	MACRO07
	MACRO08
	MACRO09
	MACRO10
	MACRO11
	MACRO12
	MACRO13
	MACRO14
	MACRO15
)

// User keycodes
const (
	USER00 Keycode = iota + 0x5F80
	USER01
	USER02
	USER03
	USER04
	USER05
	USER06
	USER07
	USER08
	USER09
	USER10
	USER11
	USER12
	USER13
	USER14
	USER15
)

func (k Keycode) Name() string {
	switch k {

	case KC_TRANSPARENT:
		return "KC_TRANSPARENT"

	case KC_SYSTEM_POWER:
		return "KC_SYSTEM_POWER"
	case KC_SYSTEM_SLEEP:
		return "KC_SYSTEM_SLEEP"
	case KC_SYSTEM_WAKE:
		return "KC_SYSTEM_WAKE"

	case KC_AUDIO_MUTE:
		return "KC_AUDIO_MUTE"
	case KC_AUDIO_VOL_UP:
		return "KC_AUDIO_VOL_UP"
	case KC_AUDIO_VOL_DOWN:
		return "KC_AUDIO_VOL_DOWN"
	case KC_MEDIA_NEXT_TRACK:
		return "KC_MEDIA_NEXT_TRACK"
	case KC_MEDIA_PREV_TRACK:
		return "KC_MEDIA_PREV_TRACK"
	case KC_MEDIA_STOP:
		return "KC_MEDIA_STOP"
	case KC_MEDIA_PLAY_PAUSE:
		return "KC_MEDIA_PLAY_PAUSE"
	case KC_MEDIA_SELECT:
		return "KC_MEDIA_SELECT"
	case KC_MEDIA_EJECT:
		return "KC_MEDIA_EJECT"
	case KC_MAIL:
		return "KC_MAIL"
	case KC_CALCULATOR:
		return "KC_CALCULATOR"
	case KC_MY_COMPUTER:
		return "KC_MY_COMPUTER"
	case KC_WWW_SEARCH:
		return "KC_WWW_SEARCH"
	case KC_WWW_HOME:
		return "KC_WWW_HOME"
	case KC_WWW_BACK:
		return "KC_WWW_BACK"
	case KC_WWW_FORWARD:
		return "KC_WWW_FORWARD"
	case KC_WWW_STOP:
		return "KC_WWW_STOP"
	case KC_WWW_REFRESH:
		return "KC_WWW_REFRESH"
	case KC_WWW_FAVORITES:
		return "KC_WWW_FAVORITES"
	case KC_MEDIA_FAST_FORWARD:
		return "KC_MEDIA_FAST_FORWARD"
	case KC_MEDIA_REWIND:
		return "KC_MEDIA_REWIND"
	case KC_BRIGHTNESS_UP:
		return "KC_BRIGHTNESS_UP"
	case KC_BRIGHTNESS_DOWN:
		return "KC_BRIGHTNESS_DOWN"

	case KC_NO:
		return "KC_NO"
	case KC_POST_FAIL:
		return "KC_POST_FAIL"
	case KC_UNDEFINED:
		return "KC_UNDEFINED"
	case KC_A:
		return "KC_A"
	case KC_B:
		return "KC_B"
	case KC_C:
		return "KC_C"
	case KC_D:
		return "KC_D"
	case KC_E:
		return "KC_E"
	case KC_F:
		return "KC_F"
	case KC_G:
		return "KC_G"
	case KC_H:
		return "KC_H"
	case KC_I:
		return "KC_I"
	case KC_J:
		return "KC_J"
	case KC_K:
		return "KC_K"
	case KC_L:
		return "KC_L"
	case KC_M:
		return "KC_M"
	case KC_N:
		return "KC_N"
	case KC_O:
		return "KC_O"
	case KC_P:
		return "KC_P"
	case KC_Q:
		return "KC_Q"
	case KC_R:
		return "KC_R"
	case KC_S:
		return "KC_S"
	case KC_T:
		return "KC_T"
	case KC_U:
		return "KC_U"
	case KC_V:
		return "KC_V"
	case KC_W:
		return "KC_W"
	case KC_X:
		return "KC_X"
	case KC_Y:
		return "KC_Y"
	case KC_Z:
		return "KC_Z"
	case KC_1:
		return "KC_1"
	case KC_2:
		return "KC_2"
	case KC_3:
		return "KC_3"
	case KC_4:
		return "KC_4"
	case KC_5:
		return "KC_5"
	case KC_6:
		return "KC_6"
	case KC_7:
		return "KC_7"
	case KC_8:
		return "KC_8"
	case KC_9:
		return "KC_9"
	case KC_0:
		return "KC_0"
	case KC_ENTER:
		return "KC_ENTER"
	case KC_ESCAPE:
		return "KC_ESCAPE"
	case KC_BACKSPACE:
		return "KC_BACKSPACE"
	case KC_TAB:
		return "KC_TAB"
	case KC_SPACE:
		return "KC_SPACE"
	case KC_MINUS:
		return "KC_MINUS"
	case KC_EQUAL:
		return "KC_EQUAL"
	case KC_LEFT_BRACKET:
		return "KC_LEFT_BRACKET"
	case KC_RIGHT_BRACKET:
		return "KC_RIGHT_BRACKET"
	case KC_BACKSLASH:
		return "KC_BACKSLASH"
	case KC_NONUS_HASH:
		return "KC_NONUS_HASH"
	case KC_SEMICOLON:
		return "KC_SEMICOLON"
	case KC_QUOTE:
		return "KC_QUOTE"
	case KC_GRAVE:
		return "KC_GRAVE"
	case KC_COMMA:
		return "KC_COMMA"
	case KC_DOT:
		return "KC_DOT"
	case KC_SLASH:
		return "KC_SLASH"
	case KC_CAPS_LOCK:
		return "KC_CAPS_LOCK"
	case KC_F1:
		return "KC_F1"
	case KC_F2:
		return "KC_F2"
	case KC_F3:
		return "KC_F3"
	case KC_F4:
		return "KC_F4"
	case KC_F5:
		return "KC_F5"
	case KC_F6:
		return "KC_F6"
	case KC_F7:
		return "KC_F7"
	case KC_F8:
		return "KC_F8"
	case KC_F9:
		return "KC_F9"
	case KC_F10:
		return "KC_F10"
	case KC_F11:
		return "KC_F11"
	case KC_F12:
		return "KC_F12"
	case KC_PRINT_SCREEN:
		return "KC_PRINT_SCREEN"
	case KC_SCROLL_LOCK:
		return "KC_SCROLL_LOCK"
	case KC_PAUSE:
		return "KC_PAUSE"
	case KC_INSERT:
		return "KC_INSERT"
	case KC_HOME:
		return "KC_HOME"
	case KC_PAGE_UP:
		return "KC_PAGE_UP"
	case KC_DELETE:
		return "KC_DELETE"
	case KC_END:
		return "KC_END"
	case KC_PAGE_DOWN:
		return "KC_PAGE_DOWN"
	case KC_RIGHT:
		return "KC_RIGHT"
	case KC_LEFT:
		return "KC_LEFT"
	case KC_DOWN:
		return "KC_DOWN"
	case KC_UP:
		return "KC_UP"
	case KC_NUM_LOCK:
		return "KC_NUM_LOCK"
	case KC_KP_SLASH:
		return "KC_KP_SLASH"
	case KC_KP_ASTERISK:
		return "KC_KP_ASTERISK"
	case KC_KP_MINUS:
		return "KC_KP_MINUS"
	case KC_KP_PLUS:
		return "KC_KP_PLUS"
	case KC_KP_ENTER:
		return "KC_KP_ENTER"
	case KC_KP_1:
		return "KC_KP_1"
	case KC_KP_2:
		return "KC_KP_2"
	case KC_KP_3:
		return "KC_KP_3"
	case KC_KP_4:
		return "KC_KP_4"
	case KC_KP_5:
		return "KC_KP_5"
	case KC_KP_6:
		return "KC_KP_6"
	case KC_KP_7:
		return "KC_KP_7"
	case KC_KP_8:
		return "KC_KP_8"
	case KC_KP_9:
		return "KC_KP_9"
	case KC_KP_0:
		return "KC_KP_0"
	case KC_KP_DOT:
		return "KC_KP_DOT"
	case KC_NONUS_BACKSLASH:
		return "KC_NONUS_BACKSLASH"
	case KC_APPLICATION:
		return "KC_APPLICATION"
	case KC_KB_POWER:
		return "KC_KB_POWER"
	case KC_KP_EQUAL:
		return "KC_KP_EQUAL"
	case KC_F13:
		return "KC_F13"
	case KC_F14:
		return "KC_F14"
	case KC_F15:
		return "KC_F15"
	case KC_F16:
		return "KC_F16"
	case KC_F17:
		return "KC_F17"
	case KC_F18:
		return "KC_F18"
	case KC_F19:
		return "KC_F19"
	case KC_F20:
		return "KC_F20"
	case KC_F21:
		return "KC_F21"
	case KC_F22:
		return "KC_F22"
	case KC_F23:
		return "KC_F23"
	case KC_F24:
		return "KC_F24"
	case KC_EXECUTE:
		return "KC_EXECUTE"
	case KC_HELP:
		return "KC_HELP"
	case KC_MENU:
		return "KC_MENU"
	case KC_SELECT:
		return "KC_SELECT"
	case KC_STOP:
		return "KC_STOP"
	case KC_AGAIN:
		return "KC_AGAIN"
	case KC_UNDO:
		return "KC_UNDO"
	case KC_CUT:
		return "KC_CUT"
	case KC_COPY:
		return "KC_COPY"
	case KC_PASTE:
		return "KC_PASTE"
	case KC_FIND:
		return "KC_FIND"
	case KC_KB_MUTE:
		return "KC_KB_MUTE"
	case KC_KB_VOLUME_UP:
		return "KC_KB_VOLUME_UP"
	case KC_KB_VOLUME_DOWN:
		return "KC_KB_VOLUME_DOWN"
	case KC_LOCKING_CAPS_LOCK:
		return "KC_LOCKING_CAPS_LOCK"
	case KC_LOCKING_NUM_LOCK:
		return "KC_LOCKING_NUM_LOCK"
	case KC_LOCKING_SCROLL_LOCK:
		return "KC_LOCKING_SCROLL_LOCK"
	case KC_KP_COMMA:
		return "KC_KP_COMMA"
	case KC_KP_EQUAL_AS400:
		return "KC_KP_EQUAL_AS400"
	case KC_INTERNATIONAL_1:
		return "KC_INTERNATIONAL_1"
	case KC_INTERNATIONAL_2:
		return "KC_INTERNATIONAL_2"
	case KC_INTERNATIONAL_3:
		return "KC_INTERNATIONAL_3"
	case KC_INTERNATIONAL_4:
		return "KC_INTERNATIONAL_4"
	case KC_INTERNATIONAL_5:
		return "KC_INTERNATIONAL_5"
	case KC_INTERNATIONAL_6:
		return "KC_INTERNATIONAL_6"
	case KC_INTERNATIONAL_7:
		return "KC_INTERNATIONAL_7"
	case KC_INTERNATIONAL_8:
		return "KC_INTERNATIONAL_8"
	case KC_INTERNATIONAL_9:
		return "KC_INTERNATIONAL_9"
	case KC_LANGUAGE_1:
		return "KC_LANGUAGE_1"
	case KC_LANGUAGE_2:
		return "KC_LANGUAGE_2"
	case KC_LANGUAGE_3:
		return "KC_LANGUAGE_3"
	case KC_LANGUAGE_4:
		return "KC_LANGUAGE_4"
	case KC_LANGUAGE_5:
		return "KC_LANGUAGE_5"
	case KC_LANGUAGE_6:
		return "KC_LANGUAGE_6"
	case KC_LANGUAGE_7:
		return "KC_LANGUAGE_7"
	case KC_LANGUAGE_8:
		return "KC_LANGUAGE_8"
	case KC_LANGUAGE_9:
		return "KC_LANGUAGE_9"
	case KC_ALTERNATE_ERASE:
		return "KC_ALTERNATE_ERASE"
	case KC_SYSTEM_REQUEST:
		return "KC_SYSTEM_REQUEST"
	case KC_CANCEL:
		return "KC_CANCEL"
	case KC_CLEAR:
		return "KC_CLEAR"
	case KC_PRIOR:
		return "KC_PRIOR"
	case KC_RETURN:
		return "KC_RETURN"
	case KC_SEPARATOR:
		return "KC_SEPARATOR"
	case KC_OUT:
		return "KC_OUT"
	case KC_OPER:
		return "KC_OPER"
	case KC_CLEAR_AGAIN:
		return "KC_CLEAR_AGAIN"
	case KC_CRSEL:
		return "KC_CRSEL"
	case KC_EXSEL:
		return "KC_EXSEL"

	case KC_LEFT_CTRL:
		return "KC_LEFT_CTRL"
	case KC_LEFT_SHIFT:
		return "KC_LEFT_SHIFT"
	case KC_LEFT_ALT:
		return "KC_LEFT_ALT"
	case KC_LEFT_GUI:
		return "KC_LEFT_GUI"
	case KC_RIGHT_CTRL:
		return "KC_RIGHT_CTRL"
	case KC_RIGHT_SHIFT:
		return "KC_RIGHT_SHIFT"
	case KC_RIGHT_ALT:
		return "KC_RIGHT_ALT"
	case KC_RIGHT_GUI:
		return "KC_RIGHT_GUI"

	case KC_FN0:
		return "KC_FN0"
	case KC_FN1:
		return "KC_FN1"
	case KC_FN2:
		return "KC_FN2"
	case KC_FN3:
		return "KC_FN3"
	case KC_FN4:
		return "KC_FN4"
	case KC_FN5:
		return "KC_FN5"
	case KC_FN6:
		return "KC_FN6"
	case KC_FN7:
		return "KC_FN7"
	case KC_FN8:
		return "KC_FN8"
	case KC_FN9:
		return "KC_FN9"
	case KC_FN10:
		return "KC_FN10"
	case KC_FN11:
		return "KC_FN11"
	case KC_FN12:
		return "KC_FN12"
	case KC_FN13:
		return "KC_FN13"
	case KC_FN14:
		return "KC_FN14"
	case KC_FN15:
		return "KC_FN15"
	case KC_FN16:
		return "KC_FN16"
	case KC_FN17:
		return "KC_FN17"
	case KC_FN18:
		return "KC_FN18"
	case KC_FN19:
		return "KC_FN19"
	case KC_FN20:
		return "KC_FN20"
	case KC_FN21:
		return "KC_FN21"
	case KC_FN22:
		return "KC_FN22"
	case KC_FN23:
		return "KC_FN23"
	case KC_FN24:
		return "KC_FN24"
	case KC_FN25:
		return "KC_FN25"
	case KC_FN26:
		return "KC_FN26"
	case KC_FN27:
		return "KC_FN27"
	case KC_FN28:
		return "KC_FN28"
	case KC_FN29:
		return "KC_FN29"
	case KC_FN30:
		return "KC_FN30"
	case KC_FN31:
		return "KC_FN31"

	case KC_MS_UP:
		return "KC_MS_UP"
	case KC_MS_DOWN:
		return "KC_MS_DOWN"
	case KC_MS_LEFT:
		return "KC_MS_LEFT"
	case KC_MS_RIGHT:
		return "KC_MS_RIGHT"
	case KC_MS_BTN1:
		return "KC_MS_BTN1"
	case KC_MS_BTN2:
		return "KC_MS_BTN2"
	case KC_MS_BTN3:
		return "KC_MS_BTN3"
	case KC_MS_BTN4:
		return "KC_MS_BTN4"
	case KC_MS_BTN5:
		return "KC_MS_BTN5"

	case KC_MS_WH_UP:
		return "KC_MS_WH_UP"
	case KC_MS_WH_DOWN:
		return "KC_MS_WH_DOWN"
	case KC_MS_WH_LEFT:
		return "KC_MS_WH_LEFT"
	case KC_MS_WH_RIGHT:
		return "KC_MS_WH_RIGHT"

	case KC_MS_ACCEL0:
		return "KC_MS_ACCEL0"
	case KC_MS_ACCEL1:
		return "KC_MS_ACCEL1"
	case KC_MS_ACCEL2:
		return "KC_MS_ACCEL2"

	case FN_MO13:
		return "FN_MO13"
	case FN_MO23:
		return "FN_MO23"
	case MACRO00:
		return "MACRO00"
	case MACRO01:
		return "MACRO01"
	case MACRO02:
		return "MACRO02"
	case MACRO03:
		return "MACRO03"
	case MACRO04:
		return "MACRO04"
	case MACRO05:
		return "MACRO05"
	case MACRO06:
		return "MACRO06"
	case MACRO07:
		return "MACRO07"
	case MACRO08:
		return "MACRO08"
	case MACRO09:
		return "MACRO09"
	case MACRO10:
		return "MACRO10"
	case MACRO11:
		return "MACRO11"
	case MACRO12:
		return "MACRO12"
	case MACRO13:
		return "MACRO13"
	case MACRO14:
		return "MACRO14"
	case MACRO15:
		return "MACRO15"

	case USER00:
		return "USER00"
	case USER01:
		return "USER01"
	case USER02:
		return "USER02"
	case USER03:
		return "USER03"
	case USER04:
		return "USER04"
	case USER05:
		return "USER05"
	case USER06:
		return "USER06"
	case USER07:
		return "USER07"
	case USER08:
		return "USER08"
	case USER09:
		return "USER09"
	case USER10:
		return "USER10"
	case USER11:
		return "USER11"
	case USER12:
		return "USER12"
	case USER13:
		return "USER13"
	case USER14:
		return "USER14"
	case USER15:
		return "USER15"

	default:
		return "UNKNOWN"
	}
}

func KeycodeFromString(value string) (Keycode, error) {
	value = strings.Replace(value, "{", "", -1)
	value = strings.Replace(value, "}", "", -1)
	value = strings.Replace(value, "KC_", "", -1)
	switch strings.ToUpper(value) {
	case "TRANSPARENT":
		return KC_TRANSPARENT, nil
	case "TRNS":
		return KC_TRANSPARENT, nil

	case "ENT":
		return KC_ENTER, nil
	case "ESC":
		return KC_ESCAPE, nil
	case "BSPC":
		return KC_BACKSPACE, nil
	case "SPC":
		return KC_SPACE, nil
	case "MINS":
		return KC_MINUS, nil
	case "EQL":
		return KC_EQUAL, nil
	case "LBRC":
		return KC_LEFT_BRACKET, nil
	case "RBRC":
		return KC_RIGHT_BRACKET, nil
	case "BSLS":
		return KC_BACKSLASH, nil
	case "NUHS":
		return KC_NONUS_HASH, nil
	case "SCLN":
		return KC_SEMICOLON, nil
	case "QUOT":
		return KC_QUOTE, nil
	case "GRV":
		return KC_GRAVE, nil
	case "COMM":
		return KC_COMMA, nil
	case "SLSH":
		return KC_SLASH, nil
	case "NUBS":
		return KC_NONUS_BACKSLASH, nil

	case "CAPS":
		return KC_CAPS_LOCK, nil
	case "SCRL":
		return KC_SCROLL_LOCK, nil
	case "NUM":
		return KC_NUM_LOCK, nil
	case "LCAP":
		return KC_LOCKING_CAPS_LOCK, nil
	case "LNUM":
		return KC_LOCKING_NUM_LOCK, nil
	case "LSCR":
		return KC_LOCKING_SCROLL_LOCK, nil

	case "PSCR":
		return KC_PRINT_SCREEN, nil
	case "PAUS":
		return KC_PAUSE, nil
	case "BRK":
		return KC_PAUSE, nil
	case "INS":
		return KC_INSERT, nil
	case "PGUP":
		return KC_PAGE_UP, nil
	case "DEL":
		return KC_DELETE, nil
	case "PGDN":
		return KC_PAGE_DOWN, nil
	case "RGHT":
		return KC_RIGHT, nil
	case "APP":
		return KC_APPLICATION, nil
	case "EXEC":
		return KC_EXECUTE, nil
	case "SLCT":
		return KC_SELECT, nil
	case "AGIN":
		return KC_AGAIN, nil
	case "PSTE":
		return KC_PASTE, nil
	case "ERAS":
		return KC_ALTERNATE_ERASE, nil
	case "SYRQ":
		return KC_SYSTEM_REQUEST, nil
	case "CNCL":
		return KC_CANCEL, nil
	case "CLR":
		return KC_CLEAR, nil
	case "PRIR":
		return KC_PRIOR, nil
	case "RETN":
		return KC_RETURN, nil
	case "SEPR":
		return KC_SEPARATOR, nil
	case "CLAG":
		return KC_CLEAR_AGAIN, nil
	case "CRSL":
		return KC_CRSEL, nil
	case "EXSL":
		return KC_EXSEL, nil

	case "PSLS":
		return KC_KP_SLASH, nil
	case "PAST":
		return KC_KP_ASTERISK, nil
	case "PMNS":
		return KC_KP_MINUS, nil
	case "PPLS":
		return KC_KP_PLUS, nil
	case "PENT":
		return KC_KP_ENTER, nil
	case "P1":
		return KC_KP_1, nil
	case "P2":
		return KC_KP_2, nil
	case "P3":
		return KC_KP_3, nil
	case "P4":
		return KC_KP_4, nil
	case "P5":
		return KC_KP_5, nil
	case "P6":
		return KC_KP_6, nil
	case "P7":
		return KC_KP_7, nil
	case "P8":
		return KC_KP_8, nil
	case "P9":
		return KC_KP_9, nil
	case "P0":
		return KC_KP_0, nil
	case "PDOT":
		return KC_KP_DOT, nil
	case "PEQL":
		return KC_KP_EQUAL, nil
	case "PCMM":
		return KC_KP_COMMA, nil

	case "INT1":
		return KC_INTERNATIONAL_1, nil
	case "INT2":
		return KC_INTERNATIONAL_2, nil
	case "INT3":
		return KC_INTERNATIONAL_3, nil
	case "INT4":
		return KC_INTERNATIONAL_4, nil
	case "INT5":
		return KC_INTERNATIONAL_5, nil
	case "INT6":
		return KC_INTERNATIONAL_6, nil
	case "INT7":
		return KC_INTERNATIONAL_7, nil
	case "INT8":
		return KC_INTERNATIONAL_8, nil
	case "INT9":
		return KC_INTERNATIONAL_9, nil
	case "LNG1":
		return KC_LANGUAGE_1, nil
	case "LNG2":
		return KC_LANGUAGE_2, nil
	case "LNG3":
		return KC_LANGUAGE_3, nil
	case "LNG4":
		return KC_LANGUAGE_4, nil
	case "LNG5":
		return KC_LANGUAGE_5, nil
	case "LNG6":
		return KC_LANGUAGE_6, nil
	case "LNG7":
		return KC_LANGUAGE_7, nil
	case "LNG8":
		return KC_LANGUAGE_8, nil
	case "LNG9":
		return KC_LANGUAGE_9, nil

	case "LCTL":
		return KC_LEFT_CTRL, nil
	case "LSFT":
		return KC_LEFT_SHIFT, nil
	case "LALT":
		return KC_LEFT_ALT, nil
	case "LOPT":
		return KC_LEFT_ALT, nil
	case "LGUI":
		return KC_LEFT_GUI, nil
	case "LCMD":
		return KC_LEFT_GUI, nil
	case "LWIN":
		return KC_LEFT_GUI, nil
	case "RCTL":
		return KC_RIGHT_CTRL, nil
	case "RSFT":
		return KC_RIGHT_SHIFT, nil
	case "RALT":
		return KC_RIGHT_ALT, nil
	case "ALGR":
		return KC_RIGHT_ALT, nil
	case "ROPT":
		return KC_RIGHT_ALT, nil
	case "RGUI":
		return KC_RIGHT_GUI, nil
	case "RCMD":
		return KC_RIGHT_GUI, nil
	case "RWIN":
		return KC_RIGHT_GUI, nil

	case "PWR":
		return KC_SYSTEM_POWER, nil
	case "SLEP":
		return KC_SYSTEM_SLEEP, nil
	case "WAKE":
		return KC_SYSTEM_WAKE, nil

	case "MUTE":
		return KC_AUDIO_MUTE, nil
	case "VOLU":
		return KC_AUDIO_VOL_UP, nil
	case "VOLD":
		return KC_AUDIO_VOL_DOWN, nil
	case "MNXT":
		return KC_MEDIA_NEXT_TRACK, nil
	case "MPRV":
		return KC_MEDIA_PREV_TRACK, nil
	case "MSTP":
		return KC_MEDIA_STOP, nil
	case "MPLY":
		return KC_MEDIA_PLAY_PAUSE, nil
	case "MSEL":
		return KC_MEDIA_SELECT, nil
	case "EJCT":
		return KC_MEDIA_EJECT, nil
	case "CALC":
		return KC_CALCULATOR, nil
	case "MYCM":
		return KC_MY_COMPUTER, nil
	case "WSCH":
		return KC_WWW_SEARCH, nil
	case "WHOM":
		return KC_WWW_HOME, nil
	case "WBAK":
		return KC_WWW_BACK, nil
	case "WFWD":
		return KC_WWW_FORWARD, nil
	case "WSTP":
		return KC_WWW_STOP, nil
	case "WREF":
		return KC_WWW_REFRESH, nil
	case "WFAV":
		return KC_WWW_FAVORITES, nil
	case "MFFD":
		return KC_MEDIA_FAST_FORWARD, nil
	case "MRWD":
		return KC_MEDIA_REWIND, nil
	case "BRIU":
		return KC_BRIGHTNESS_UP, nil
	case "BRID":
		return KC_BRIGHTNESS_DOWN, nil

	case "BRMU":
		return KC_PAUSE, nil
	case "BRMD":
		return KC_SCROLL_LOCK, nil

	case "MS_U":
		return KC_MS_UP, nil
	case "MS_D":
		return KC_MS_DOWN, nil
	case "MS_L":
		return KC_MS_LEFT, nil
	case "MS_R":
		return KC_MS_RIGHT, nil
	case "BTN1":
		return KC_MS_BTN1, nil
	case "BTN2":
		return KC_MS_BTN2, nil
	case "BTN3":
		return KC_MS_BTN3, nil
	case "BTN4":
		return KC_MS_BTN4, nil
	case "BTN5":
		return KC_MS_BTN5, nil
	case "BTN6":
		return KC_MS_BTN6, nil
	case "BTN7":
		return KC_MS_BTN7, nil
	case "BTN8":
		return KC_MS_BTN8, nil
	case "WH_U":
		return KC_MS_WH_UP, nil
	case "WH_D":
		return KC_MS_WH_DOWN, nil
	case "WH_L":
		return KC_MS_WH_LEFT, nil
	case "WH_R":
		return KC_MS_WH_RIGHT, nil
	case "ACL0":
		return KC_MS_ACCEL0, nil
	case "ACL1":
		return KC_MS_ACCEL1, nil
	case "ACL2":
		return KC_MS_ACCEL2, nil

	case "SYSTEM_POWER":
		return KC_SYSTEM_POWER, nil
	case "SYSTEM_SLEEP":
		return KC_SYSTEM_SLEEP, nil
	case "SYSTEM_WAKE":
		return KC_SYSTEM_WAKE, nil

	case "AUDIO_MUTE":
		return KC_AUDIO_MUTE, nil
	case "AUDIO_VOL_UP":
		return KC_AUDIO_VOL_UP, nil
	case "AUDIO_VOL_DOWN":
		return KC_AUDIO_VOL_DOWN, nil
	case "MEDIA_NEXT_TRACK":
		return KC_MEDIA_NEXT_TRACK, nil
	case "MEDIA_PREV_TRACK":
		return KC_MEDIA_PREV_TRACK, nil
	case "MEDIA_STOP":
		return KC_MEDIA_STOP, nil
	case "MEDIA_PLAY_PAUSE":
		return KC_MEDIA_PLAY_PAUSE, nil
	case "MEDIA_SELECT":
		return KC_MEDIA_SELECT, nil
	case "MEDIA_EJECT":
		return KC_MEDIA_EJECT, nil
	case "MAIL":
		return KC_MAIL, nil
	case "CALCULATOR":
		return KC_CALCULATOR, nil
	case "MY_COMPUTER":
		return KC_MY_COMPUTER, nil
	case "WWW_SEARCH":
		return KC_WWW_SEARCH, nil
	case "WWW_HOME":
		return KC_WWW_HOME, nil
	case "WWW_BACK":
		return KC_WWW_BACK, nil
	case "WWW_FORWARD":
		return KC_WWW_FORWARD, nil
	case "WWW_STOP":
		return KC_WWW_STOP, nil
	case "WWW_REFRESH":
		return KC_WWW_REFRESH, nil
	case "WWW_FAVORITES":
		return KC_WWW_FAVORITES, nil
	case "MEDIA_FAST_FORWARD":
		return KC_MEDIA_FAST_FORWARD, nil
	case "MEDIA_REWIND":
		return KC_MEDIA_REWIND, nil
	case "BRIGHTNESS_UP":
		return KC_BRIGHTNESS_UP, nil
	case "BRIGHTNESS_DOWN":
		return KC_BRIGHTNESS_DOWN, nil

	case "NO":
		return KC_NO, nil
	case "ROLL_OVER":
		return KC_ROLL_OVER, nil
	case "POST_FAIL":
		return KC_POST_FAIL, nil
	case "UNDEFINED":
		return KC_UNDEFINED, nil
	case "A":
		return KC_A, nil
	case "B":
		return KC_B, nil
	case "C":
		return KC_C, nil
	case "D":
		return KC_D, nil
	case "E":
		return KC_E, nil
	case "F":
		return KC_F, nil
	case "G":
		return KC_G, nil
	case "H":
		return KC_H, nil
	case "I":
		return KC_I, nil
	case "J":
		return KC_J, nil
	case "K":
		return KC_K, nil
	case "L":
		return KC_L, nil
	case "M":
		return KC_M, nil
	case "N":
		return KC_N, nil
	case "O":
		return KC_O, nil
	case "P":
		return KC_P, nil
	case "Q":
		return KC_Q, nil
	case "R":
		return KC_R, nil
	case "S":
		return KC_S, nil
	case "T":
		return KC_T, nil
	case "U":
		return KC_U, nil
	case "V":
		return KC_V, nil
	case "W":
		return KC_W, nil
	case "X":
		return KC_X, nil
	case "Y":
		return KC_Y, nil
	case "Z":
		return KC_Z, nil
	case "1":
		return KC_1, nil
	case "2":
		return KC_2, nil
	case "3":
		return KC_3, nil
	case "4":
		return KC_4, nil
	case "5":
		return KC_5, nil
	case "6":
		return KC_6, nil
	case "7":
		return KC_7, nil
	case "8":
		return KC_8, nil
	case "9":
		return KC_9, nil
	case "0":
		return KC_0, nil
	case "ENTER":
		return KC_ENTER, nil
	case "ESCAPE":
		return KC_ESCAPE, nil
	case "BACKSPACE":
		return KC_BACKSPACE, nil
	case "TAB":
		return KC_TAB, nil
	case "SPACE":
		return KC_SPACE, nil
	case "MINUS":
		return KC_MINUS, nil
	case "EQUAL":
		return KC_EQUAL, nil
	case "LEFT_BRACKET":
		return KC_LEFT_BRACKET, nil
	case "RIGHT_BRACKET":
		return KC_RIGHT_BRACKET, nil
	case "BACKSLASH":
		return KC_BACKSLASH, nil
	case "NONUS_HASH":
		return KC_NONUS_HASH, nil
	case "SEMICOLON":
		return KC_SEMICOLON, nil
	case "QUOTE":
		return KC_QUOTE, nil
	case "GRAVE":
		return KC_GRAVE, nil
	case "COMMA":
		return KC_COMMA, nil
	case "DOT":
		return KC_DOT, nil
	case "SLASH":
		return KC_SLASH, nil
	case "CAPS_LOCK":
		return KC_CAPS_LOCK, nil
	case "F1":
		return KC_F1, nil
	case "F2":
		return KC_F2, nil
	case "F3":
		return KC_F3, nil
	case "F4":
		return KC_F4, nil
	case "F5":
		return KC_F5, nil
	case "F6":
		return KC_F6, nil
	case "F7":
		return KC_F7, nil
	case "F8":
		return KC_F8, nil
	case "F9":
		return KC_F9, nil
	case "F10":
		return KC_F10, nil
	case "F11":
		return KC_F11, nil
	case "F12":
		return KC_F12, nil
	case "PRINT_SCREEN":
		return KC_PRINT_SCREEN, nil
	case "SCROLL_LOCK":
		return KC_SCROLL_LOCK, nil
	case "PAUSE":
		return KC_PAUSE, nil
	case "INSERT":
		return KC_INSERT, nil
	case "HOME":
		return KC_HOME, nil
	case "PAGE_UP":
		return KC_PAGE_UP, nil
	case "DELETE":
		return KC_DELETE, nil
	case "END":
		return KC_END, nil
	case "PAGE_DOWN":
		return KC_PAGE_DOWN, nil
	case "RIGHT":
		return KC_RIGHT, nil
	case "LEFT":
		return KC_LEFT, nil
	case "DOWN":
		return KC_DOWN, nil
	case "UP":
		return KC_UP, nil
	case "NUM_LOCK":
		return KC_NUM_LOCK, nil
	case "KP_SLASH":
		return KC_KP_SLASH, nil
	case "KP_ASTERISK":
		return KC_KP_ASTERISK, nil
	case "KP_MINUS":
		return KC_KP_MINUS, nil
	case "KP_PLUS":
		return KC_KP_PLUS, nil
	case "KP_ENTER":
		return KC_KP_ENTER, nil
	case "KP_1":
		return KC_KP_1, nil
	case "KP_2":
		return KC_KP_2, nil
	case "KP_3":
		return KC_KP_3, nil
	case "KP_4":
		return KC_KP_4, nil
	case "KP_5":
		return KC_KP_5, nil
	case "KP_6":
		return KC_KP_6, nil
	case "KP_7":
		return KC_KP_7, nil
	case "KP_8":
		return KC_KP_8, nil
	case "KP_9":
		return KC_KP_9, nil
	case "KP_0":
		return KC_KP_0, nil
	case "KP_DOT":
		return KC_KP_DOT, nil
	case "NONUS_BACKSLASH":
		return KC_NONUS_BACKSLASH, nil
	case "APPLICATION":
		return KC_APPLICATION, nil
	case "KB_POWER":
		return KC_KB_POWER, nil
	case "KP_EQUAL":
		return KC_KP_EQUAL, nil
	case "F13":
		return KC_F13, nil
	case "F14":
		return KC_F14, nil
	case "F15":
		return KC_F15, nil
	case "F16":
		return KC_F16, nil
	case "F17":
		return KC_F17, nil
	case "F18":
		return KC_F18, nil
	case "F19":
		return KC_F19, nil
	case "F20":
		return KC_F20, nil
	case "F21":
		return KC_F21, nil
	case "F22":
		return KC_F22, nil
	case "F23":
		return KC_F23, nil
	case "F24":
		return KC_F24, nil
	case "EXECUTE":
		return KC_EXECUTE, nil
	case "HELP":
		return KC_HELP, nil
	case "MENU":
		return KC_MENU, nil
	case "SELECT":
		return KC_SELECT, nil
	case "STOP":
		return KC_STOP, nil
	case "AGAIN":
		return KC_AGAIN, nil
	case "UNDO":
		return KC_UNDO, nil
	case "CUT":
		return KC_CUT, nil
	case "COPY":
		return KC_COPY, nil
	case "PASTE":
		return KC_PASTE, nil
	case "FIND":
		return KC_FIND, nil
	case "KB_MUTE":
		return KC_KB_MUTE, nil
	case "KB_VOLUME_UP":
		return KC_KB_VOLUME_UP, nil
	case "KB_VOLUME_DOWN":
		return KC_KB_VOLUME_DOWN, nil
	case "LOCKING_CAPS_LOCK":
		return KC_LOCKING_CAPS_LOCK, nil
	case "LOCKING_NUM_LOCK":
		return KC_LOCKING_NUM_LOCK, nil
	case "LOCKING_SCROLL_LOCK":
		return KC_LOCKING_SCROLL_LOCK, nil
	case "KP_COMMA":
		return KC_KP_COMMA, nil
	case "KP_EQUAL_AS400":
		return KC_KP_EQUAL_AS400, nil
	case "INTERNATIONAL_1":
		return KC_INTERNATIONAL_1, nil
	case "INTERNATIONAL_2":
		return KC_INTERNATIONAL_2, nil
	case "INTERNATIONAL_3":
		return KC_INTERNATIONAL_3, nil
	case "INTERNATIONAL_4":
		return KC_INTERNATIONAL_4, nil
	case "INTERNATIONAL_5":
		return KC_INTERNATIONAL_5, nil
	case "INTERNATIONAL_6":
		return KC_INTERNATIONAL_6, nil
	case "INTERNATIONAL_7":
		return KC_INTERNATIONAL_7, nil
	case "INTERNATIONAL_8":
		return KC_INTERNATIONAL_8, nil
	case "INTERNATIONAL_9":
		return KC_INTERNATIONAL_9, nil
	case "LANGUAGE_1":
		return KC_LANGUAGE_1, nil
	case "LANGUAGE_2":
		return KC_LANGUAGE_2, nil
	case "LANGUAGE_3":
		return KC_LANGUAGE_3, nil
	case "LANGUAGE_4":
		return KC_LANGUAGE_4, nil
	case "LANGUAGE_5":
		return KC_LANGUAGE_5, nil
	case "LANGUAGE_6":
		return KC_LANGUAGE_6, nil
	case "LANGUAGE_7":
		return KC_LANGUAGE_7, nil
	case "LANGUAGE_8":
		return KC_LANGUAGE_8, nil
	case "LANGUAGE_9":
		return KC_LANGUAGE_9, nil
	case "ALTERNATE_ERASE":
		return KC_ALTERNATE_ERASE, nil
	case "SYSTEM_REQUEST":
		return KC_SYSTEM_REQUEST, nil
	case "CANCEL":
		return KC_CANCEL, nil
	case "CLEAR":
		return KC_CLEAR, nil
	case "PRIOR":
		return KC_PRIOR, nil
	case "RETURN":
		return KC_RETURN, nil
	case "SEPARATOR":
		return KC_SEPARATOR, nil
	case "OUT":
		return KC_OUT, nil
	case "OPER":
		return KC_OPER, nil
	case "CLEAR_AGAIN":
		return KC_CLEAR_AGAIN, nil
	case "CRSEL":
		return KC_CRSEL, nil
	case "EXSEL":
		return KC_EXSEL, nil

	case "LEFT_CTRL":
		return KC_LEFT_CTRL, nil
	case "LEFT_SHIFT":
		return KC_LEFT_SHIFT, nil
	case "LEFT_ALT":
		return KC_LEFT_ALT, nil
	case "LEFT_GUI":
		return KC_LEFT_GUI, nil
	case "RIGHT_CTRL":
		return KC_RIGHT_CTRL, nil
	case "RIGHT_SHIFT":
		return KC_RIGHT_SHIFT, nil
	case "RIGHT_ALT":
		return KC_RIGHT_ALT, nil
	case "RIGHT_GUI":
		return KC_RIGHT_GUI, nil

	case "FN0":
		return KC_FN0, nil
	case "FN1":
		return KC_FN1, nil
	case "FN2":
		return KC_FN2, nil
	case "FN3":
		return KC_FN3, nil
	case "FN4":
		return KC_FN4, nil
	case "FN5":
		return KC_FN5, nil
	case "FN6":
		return KC_FN6, nil
	case "FN7":
		return KC_FN7, nil
	case "FN8":
		return KC_FN8, nil
	case "FN9":
		return KC_FN9, nil
	case "FN10":
		return KC_FN10, nil
	case "FN11":
		return KC_FN11, nil
	case "FN12":
		return KC_FN12, nil
	case "FN13":
		return KC_FN13, nil
	case "FN14":
		return KC_FN14, nil
	case "FN15":
		return KC_FN15, nil
	case "FN16":
		return KC_FN16, nil
	case "FN17":
		return KC_FN17, nil
	case "FN18":
		return KC_FN18, nil
	case "FN19":
		return KC_FN19, nil
	case "FN20":
		return KC_FN20, nil
	case "FN21":
		return KC_FN21, nil
	case "FN22":
		return KC_FN22, nil
	case "FN23":
		return KC_FN23, nil
	case "FN24":
		return KC_FN24, nil
	case "FN25":
		return KC_FN25, nil
	case "FN26":
		return KC_FN26, nil
	case "FN27":
		return KC_FN27, nil
	case "FN28":
		return KC_FN28, nil
	case "FN29":
		return KC_FN29, nil
	case "FN30":
		return KC_FN30, nil
	case "FN31":
		return KC_FN31, nil

	case "MS_UP":
		return KC_MS_UP, nil
	case "MS_DOWN":
		return KC_MS_DOWN, nil
	case "MS_LEFT":
		return KC_MS_LEFT, nil
	case "MS_RIGHT":
		return KC_MS_RIGHT, nil
	case "MS_BTN1":
		return KC_MS_BTN1, nil
	case "MS_BTN2":
		return KC_MS_BTN2, nil
	case "MS_BTN3":
		return KC_MS_BTN3, nil
	case "MS_BTN4":
		return KC_MS_BTN4, nil
	case "MS_BTN5":
		return KC_MS_BTN5, nil
	case "MS_BTN6":
		return KC_MS_BTN6, nil
	case "MS_BTN7":
		return KC_MS_BTN7, nil
	case "MS_BTN8":
		return KC_MS_BTN8, nil

	case "MS_WH_UP":
		return KC_MS_WH_UP, nil
	case "MS_WH_DOWN":
		return KC_MS_WH_DOWN, nil
	case "MS_WH_LEFT":
		return KC_MS_WH_LEFT, nil
	case "MS_WH_RIGHT":
		return KC_MS_WH_RIGHT, nil

	case "MS_ACCEL0":
		return KC_MS_ACCEL0, nil
	case "MS_ACCEL1":
		return KC_MS_ACCEL1, nil
	case "MS_ACCEL2":
		return KC_MS_ACCEL2, nil

	case "BSPACE":
		return KC_BACKSPACE, nil
	case "LBRACKET":
		return KC_LEFT_BRACKET, nil
	case "RBRACKET":
		return KC_RIGHT_BRACKET, nil
	case "BSLASH":
		return KC_BACKSLASH, nil
	case "SCOLON":
		return KC_SEMICOLON, nil
	case "CAPSLOCK":
		return KC_CAPS_LOCK, nil
	case "PSCREEN":
		return KC_PRINT_SCREEN, nil
	case "SCROLLLOCK":
		return KC_SCROLL_LOCK, nil
	case "PGDOWN":
		return KC_PAGE_DOWN, nil
	case "NUMLOCK":
		return KC_NUM_LOCK, nil
	case "NONUS_BSLASH":
		return KC_NONUS_BACKSLASH, nil
	case "POWER":
		return KC_KB_POWER, nil
	case "_MUTE":
		return KC_KB_MUTE, nil
	case "_VOLUP":
		return KC_KB_VOLUME_UP, nil
	case "_VOLDOWN":
		return KC_KB_VOLUME_DOWN, nil
	case "LOCKING_CAPS":
		return KC_LOCKING_CAPS_LOCK, nil
	case "LOCKING_NUM":
		return KC_LOCKING_NUM_LOCK, nil
	case "LOCKING_SCROLL":
		return KC_LOCKING_SCROLL_LOCK, nil
	case "LANG1":
		return KC_LANGUAGE_1, nil
	case "LANG2":
		return KC_LANGUAGE_2, nil
	case "LANG3":
		return KC_LANGUAGE_3, nil
	case "LANG4":
		return KC_LANGUAGE_4, nil
	case "LANG5":
		return KC_LANGUAGE_5, nil
	case "LANG6":
		return KC_LANGUAGE_6, nil
	case "LANG7":
		return KC_LANGUAGE_7, nil
	case "LANG8":
		return KC_LANGUAGE_8, nil
	case "LANG9":
		return KC_LANGUAGE_9, nil
	case "ALT_ERASE":
		return KC_ALTERNATE_ERASE, nil
	case "SYSREQ":
		return KC_SYSTEM_REQUEST, nil

	case "LSHIFT":
		return KC_LEFT_SHIFT, nil
	case "RCTRL":
		return KC_RIGHT_CTRL, nil
	case "RSHIFT":
		return KC_RIGHT_SHIFT, nil

	case "RO":
		return KC_INTERNATIONAL_1, nil
	case "KANA":
		return KC_INTERNATIONAL_2, nil
	case "JYEN":
		return KC_INTERNATIONAL_3, nil
	case "HENK":
		return KC_INTERNATIONAL_4, nil
	case "MHEN":
		return KC_INTERNATIONAL_5, nil
	case "HAEN":
		return KC_LANGUAGE_1, nil
	case "HANJ":
		return KC_LANGUAGE_2, nil

	case "SLCK":
		return KC_SCROLL_LOCK, nil
	case "NLCK":
		return KC_NUM_LOCK, nil

	case "FN_MO13":
		return FN_MO13, nil
	case "FN_MO23":
		return FN_MO23, nil
	case "MACRO00":
		return MACRO00, nil
	case "MACRO01":
		return MACRO01, nil
	case "MACRO02":
		return MACRO02, nil
	case "MACRO03":
		return MACRO03, nil
	case "MACRO04":
		return MACRO04, nil
	case "MACRO05":
		return MACRO05, nil
	case "MACRO06":
		return MACRO06, nil
	case "MACRO07":
		return MACRO07, nil
	case "MACRO08":
		return MACRO08, nil
	case "MACRO09":
		return MACRO09, nil
	case "MACRO10":
		return MACRO10, nil
	case "MACRO11":
		return MACRO11, nil
	case "MACRO12":
		return MACRO12, nil
	case "MACRO13":
		return MACRO13, nil
	case "MACRO14":
		return MACRO14, nil
	case "MACRO15":
		return MACRO15, nil

	case "USER00":
		return USER00, nil
	case "USER01":
		return USER01, nil
	case "USER02":
		return USER02, nil
	case "USER03":
		return USER03, nil
	case "USER04":
		return USER04, nil
	case "USER05":
		return USER05, nil
	case "USER06":
		return USER06, nil
	case "USER07":
		return USER07, nil
	case "USER08":
		return USER08, nil
	case "USER09":
		return USER09, nil
	case "USER10":
		return USER10, nil
	case "USER11":
		return USER11, nil
	case "USER12":
		return USER12, nil
	case "USER13":
		return USER13, nil
	case "USER14":
		return USER14, nil
	case "USER15":
		return USER15, nil

	default:
		return KC_NO, ErrorUnknownKeycode
	}
}
