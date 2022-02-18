// qmk-go - go client library for VIA-enabled QMK keyboards
// Copyright (c) 2022 Ian McLinden. All rights reserved
//
// This file is released under GNU LGPL 2.1 on Linux,
// and under the 3-clause BSD license on all other platforms

package qmk

// VIA Protocol
// This is changed only when the command IDs change,
// so VIA Configurator can detect compatible firmware.
const (
	ViaProtocolVersion = 0x0009
)

// HID Usage Page
const (
	HidMessageSize = 32
	HidUsagePage   = 0xFF60
	HidUsage       = 0x61
)

// VIA Command IDs
const (
	GetProtocolVersionId              = 0x01
	GetKeyboardValueId                = 0x02
	SetKeyboardValueId                = 0x03
	DynamicKeymapGetKeycodeId         = 0x04
	DynamicKeymapSetKeycodeId         = 0x05
	DynamicKeymapResetId              = 0x06
	LightingSetValueId                = 0x07
	LightingGetValueId                = 0x08
	LightingSaveId                    = 0x09
	EepromResetId                     = 0x0A
	BootloaderJumpId                  = 0x0B
	DynamicKeymapMacroGetCountId      = 0x0C
	DynamicKeymapMacroGetBufferSizeId = 0x0D
	DynamicKeymapMacroGetBufferId     = 0x0E
	DynamicKeymapMacroSetBufferId     = 0x0F
	DynamicKeymapMacroResetId         = 0x10
	DynamicKeymapGetLayerCountId      = 0x11
	DynamicKeymapGetBufferId          = 0x12
	DynamicKeymapSetBufferId          = 0x13
	UnhandledId                       = 0xFF
)

// VIA Keyboard Value IDs
const (
	UptimeId            = 0x01
	LayoutOptionsId     = 0x02
	SwitchMatrixStateId = 0x03
)

// VIA Lighting Value IDs
const (
	BacklightBrightnessId = 0x09
	BacklightEffectId     = 0x0A

	RgblightBrightnessId  = 0x80
	RgblightEffectId      = 0x81
	RgblightEffectSpeedId = 0x82
	RgblightColorId       = 0x83
)

// Dynamic Keymap
const (
	MaxDynamicKeymapBufferSize = 28
)
