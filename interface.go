// qmk-go - go client library for VIA-enabled QMK keyboards
// Copyright (c) 2022 Ian McLinden. All rights reserved
//
// This file is released under GNU LGPL 2.1 on Linux,
// and under the 3-clause BSD license on all other platforms

package qmk

import (
	"github.com/ianmclinden/qmk-go/backlight"
	"github.com/ianmclinden/qmk-go/keycode"
	"github.com/ianmclinden/qmk-go/rgblight"
)

// Client to bind and configure QMK Keyboard
type Client interface {
	Keyboard() Keyboard

	// id_get_protocol_version
	GetProtocolVersion() (uint16, error)
	// id_get_keyboard_value -> id_uptime (ms)
	GetUptime() (uint32, error)

	// id_get_keyboard_value -> id_layout_options
	GetLayoutOptions() (uint32, error)
	// id_set_keyboard_value -> id_layout_options
	SetLayoutOptions(uint32) error

	// id_get_keyboard_value -> id_switch_matrix_state
	GetSwitchMatrixState() ([]byte, error)

	// id_get_keyboard_value -> default (raw_hid_receive_kb)
	GetRawKeyboardValue(uint8) ([]byte, error)
	// id_set_keyboard_value -> default (raw_hid_receive_kb)
	SetRawKeyboardValue(uint8, []byte) error

	// id_dynamic_keymap_get_keycode
	GetDynamicKeymapKeycode(uint8, uint8, uint8) (keycode.Keycode, error)
	// id_dynamic_keymap_set_keycode
	SetDynamicKeymapKeycode(uint8, uint8, uint8, keycode.Keycode) error
	// id_dynamic_keymap_reset
	ResetDynamicKeymap() error

	// id_lighting_get_value -> id_qmk_backlight_brightness
	GetBacklightBrightness() (backlight.Brightness, error)
	// id_lighting_set_value -> id_qmk_backlight_brightness
	SetBacklightBrightness(backlight.Brightness) error
	// id_lighting_get_value -> id_qmk_backlight_effect
	GetBacklightEffect() (backlight.Effect, error)
	// id_lighting_set_value -> id_qmk_backlight_effect
	SetBacklightEffect(backlight.Effect) error

	// Get RGB light brightness
	GetRgblightBrightness() (rgblight.Brightness, error)
	// Set RGB light brightness
	SetRgblightBrightness(rgblight.Brightness) error
	// Get RGB light color
	GetRgblightColor() (rgblight.Color, error)
	// SetRGB Light color (optionally set brightness from color)
	SetRgblightColor(rgblight.Color, bool) error
	// Get RGB light effect
	GetRgblightEffect() (rgblight.Effect, error)
	// Set RGB light effect
	SetRgblightEffect(rgblight.Effect) error
	// Get RGB light effect speed
	GetRgblightEffectSpeed() (rgblight.Speed, error)
	// Set RGB light effect speed
	SetRgblightEffectSpeed(rgblight.Speed) error

	// Save rgblight and backlight to EEPROM
	SaveLighting() error

	// Get the number of supported macros
	GetDynamicKeymapMacroCount() (uint8, error)
	// Get macro by VIA index (preferred)
	GetDynamicKeymapMacro(uint8) ([]byte, error)
	// Set macro by VIA index (preferred)
	SetDynamicKeymapMacro(uint8, []byte) error
	// Reset the dynamic keymap
	ResetDynamicKeymapMacro() error

	// Get the number of supported keymap layers
	GetDynamicKeymapLayerCount() (uint8, error)
	// Get the number of bytes allocated for dynamic macros
	GetDynamicKeymapMacroBufferSize() (uint16, error)
	// Get dynamic macros by buffer (id_dynamic_keymap_macro_get_buffer)
	GetDynamicKeymapMacroBuffer(uint16, uint8) ([]byte, error)
	// Set dynamic macros by buffer (id_dynamic_keymap_macro_set_buffer)
	SetDynamicKeymapMacroBuffer(uint16, uint8, []byte) error

	// Get multiple keycodes by buffer (id_dynamic_keymap_get_buffer)
	GetDynamicKeymapBuffer(uint16, uint8) ([]byte, error)
	// Set multiple keycodes by buffer (id_dynamic_keymap_set_buffer)
	SetDynamicKeymapBuffer(uint16, uint8, []byte) error

	// Reset EEPROM
	ResetEeprom() error
}
