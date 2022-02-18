// qmk-go - go client library for VIA-enabled QMK keyboards
// Copyright (c) 2022 Ian McLinden. All rights reserved
//
// This file is released under GNU LGPL 2.1 on Linux,
// and under the 3-clause BSD license on all other platforms

package qmk

import (
	"bytes"
	"errors"
	"fmt"
	"math"

	"github.com/ianmclinden/qmk-go/backlight"
	"github.com/ianmclinden/qmk-go/keycode"
	"github.com/ianmclinden/qmk-go/rgblight"
	"github.com/karalabe/hid"
)

type client struct {
	device *hid.Device
}

var (
	ErrorNoMatchingDevice = errors.New("no matching devices found")
	ErrorVersionMismatch  = errors.New("keyboard does not match this VIA version")
	ErrorBadMessageSize   = errors.New("incorrect QMK Message size")
	ErrorReadWrite        = errors.New("could not read/write to QMK device")
	ErrorUnknownCommand   = errors.New("unknown VIA command")
	ErrorBadBufferSize    = fmt.Errorf("incorrect buffer size (<=%d)", MaxDynamicKeymapBufferSize)
	ErrorMacroNotInBytes  = errors.New("macro was not found in bytes")
	ErrorMacroCopy        = errors.New("could not copy macro into buffer")
)

var (
	dynamicKeymapMacroCache []byte
)

func NewClient(vid uint16, pid uint16, serial string) (Client, error) {

	keyboards, err := ListKeyboards()
	if err != nil {
		return nil, err
	}
	serials := []hid.DeviceInfo{}

	for i := range keyboards {
		if serial == "" || keyboards[i].Serial == serial {
			serials = append(serials, keyboards[i])
		}
	}

	if len(serials) == 0 {
		return nil, ErrorNoMatchingDevice
	}

	di := serials[0]

	var device *hid.Device
	for i := 0; i < 20; i++ {
		device, err = di.Open()
		if err != nil {
			continue
		}
		c := &client{device}
		version, err := c.GetProtocolVersion()
		if err != nil {
			return nil, err
		}
		if version != ViaProtocolVersion {
			return nil, ErrorVersionMismatch
		}
		return c, nil
	}
	return nil, err
}

func (c *client) sendMessage(message []byte, retries int) error {
	if len(message) != HidMessageSize {
		return ErrorBadMessageSize
	}
	for i := 0; i < retries; i++ {
		wrote, err := c.device.Write(message)
		if err != nil || wrote != HidMessageSize {
			continue
		}
		read, err := c.device.Read(message)
		if err != nil || read != HidMessageSize {
			continue
		}
		if message[0] == UnhandledId {
			return ErrorUnknownCommand
		}
		return nil
	}
	return ErrorReadWrite
}

func (c *client) Keyboard() Keyboard {
	return Keyboard(c.device.DeviceInfo)
}

func (c *client) GetProtocolVersion() (uint16, error) {
	buffer := [HidMessageSize]byte{GetProtocolVersionId}
	err := c.sendMessage(buffer[:], 20)
	if err != nil {
		return 0, err
	}

	return uint16(buffer[1])<<8 | uint16(buffer[2]), nil
}

func (c *client) GetUptime() (uint32, error) {
	buffer := [HidMessageSize]byte{GetKeyboardValueId, UptimeId}
	err := c.sendMessage(buffer[:], 20)
	if err != nil {
		return 0, err
	}

	uptime := uint32(buffer[2])<<24 | uint32(buffer[3])<<16 | uint32(buffer[4])<<8 | uint32(buffer[5])
	return uptime, nil
}

func (c *client) GetLayoutOptions() (uint32, error) {
	buffer := [HidMessageSize]byte{GetKeyboardValueId, LayoutOptionsId}
	err := c.sendMessage(buffer[:], 20)
	if err != nil {
		return 0, err
	}

	options := uint32(buffer[2])<<24 | uint32(buffer[3])<<16 | uint32(buffer[4])<<8 | uint32(buffer[5])
	return options, nil
}

func (c *client) GetSwitchMatrixState() ([]byte, error) {
	buffer := [HidMessageSize]byte{GetKeyboardValueId, SwitchMatrixStateId}
	err := c.sendMessage(buffer[:], 20)
	if err != nil {
		return nil, err
	}

	return buffer[2:], nil
}

// id_get_keyboard_value -> default (raw_hid_receive_kb)
func (c *client) GetRawKeyboardValue(id uint8) ([]byte, error) {
	buffer := [HidMessageSize]byte{GetKeyboardValueId, byte(id)}
	err := c.sendMessage(buffer[:], 20)
	if err != nil {
		return nil, err
	}

	return buffer[2:], nil
}

func (c *client) SetLayoutOptions(options uint32) error {
	buffer := [HidMessageSize]byte{
		GetKeyboardValueId,
		LayoutOptionsId,
		byte((options >> 24) & 0xFF),
		byte((options >> 16) & 0xFF),
		byte((options >> 8) & 0xFF),
		byte((options) & 0xFF),
	}
	return c.sendMessage(buffer[:], 20)
}

func (c *client) SetRawKeyboardValue(id uint8, data []byte) error {
	if len(data) > HidMessageSize-2 {
		return fmt.Errorf("data slice was too long (<=%d)", HidMessageSize-2)
	}
	buffer := append([]byte{SetKeyboardValueId, byte(id)}, data...)
	return c.sendMessage(buffer[:], 20)
}

func (c *client) GetDynamicKeymapKeycode(layer uint8, row uint8, column uint8) (keycode.Keycode, error) {
	buffer := [HidMessageSize]byte{
		DynamicKeymapGetKeycodeId,
		byte(layer),
		byte(row),
		byte(column),
	}
	err := c.sendMessage(buffer[:], 20)
	if err != nil {
		return 0, err
	}

	return keycode.KeycodeFromBytes(buffer[4], buffer[5]), nil
}

func (c *client) SetDynamicKeymapKeycode(layer uint8, row uint8, column uint8, keycode keycode.Keycode) error {
	buffer := [HidMessageSize]byte{
		DynamicKeymapSetKeycodeId,
		byte(layer),
		byte(row),
		byte(column),
		keycode.ToBytes()[0],
		keycode.ToBytes()[1],
	}
	return c.sendMessage(buffer[:], 20)
}

func (c *client) ResetDynamicKeymap() error {
	buffer := [HidMessageSize]byte{DynamicKeymapResetId}
	return c.sendMessage(buffer[:], 20)
}

func (c *client) GetBacklightBrightness() (backlight.Brightness, error) {
	buffer := [HidMessageSize]byte{LightingGetValueId, BacklightBrightnessId}
	err := c.sendMessage(buffer[:], 20)
	if err != nil {
		return 0, err
	}

	return backlight.BrightnessFromByte(buffer[2]), nil
}

func (c *client) GetBacklightEffect() (backlight.Effect, error) {
	buffer := [HidMessageSize]byte{LightingGetValueId, BacklightEffectId}
	err := c.sendMessage(buffer[:], 20)
	if err != nil {
		return 0, err
	}

	return backlight.EffectFromByte(buffer[2]), nil
}

func (c *client) GetRgblightBrightness() (rgblight.Brightness, error) {
	buffer := [HidMessageSize]byte{LightingGetValueId, RgblightBrightnessId}
	err := c.sendMessage(buffer[:], 20)
	if err != nil {
		return 0, err
	}

	return rgblight.BrightnessFromByte(buffer[2]), nil
}

func (c *client) GetRgblightEffect() (rgblight.Effect, error) {
	buffer := [HidMessageSize]byte{LightingGetValueId, RgblightEffectId}
	err := c.sendMessage(buffer[:], 20)
	if err != nil {
		return rgblight.EffectUnknown, err
	}

	return rgblight.EffectFromByte(buffer[2]), nil
}

func (c *client) GetRgblightEffectSpeed() (rgblight.Speed, error) {
	buffer := [HidMessageSize]byte{LightingGetValueId, RgblightEffectSpeedId}
	err := c.sendMessage(buffer[:], 20)
	if err != nil {
		return 0, err
	}

	return rgblight.SpeedFromByte(buffer[2]), nil
}

func (c *client) GetRgblightColor() (rgblight.Color, error) {
	buffer := [HidMessageSize]byte{LightingGetValueId, RgblightColorId}
	err := c.sendMessage(buffer[:], 20)
	if err != nil {
		return rgblight.ColorOff, err
	}

	var (
		hue = rgblight.HueFromByte(buffer[2])
		sat = rgblight.SaturationFromByte(buffer[3])
	)

	val, err := c.GetRgblightBrightness()
	if err != nil {
		return rgblight.ColorOff, err
	}

	return rgblight.Color{Hue: hue, Saturation: sat, Brightness: val}, nil
}

func (c *client) SetBacklightBrightness(brightness backlight.Brightness) error {
	buffer := [HidMessageSize]byte{
		LightingSetValueId,
		BacklightBrightnessId,
		brightness.ToByte(),
	}
	return c.sendMessage(buffer[:], 20)
}

func (c *client) SetBacklightEffect(effect backlight.Effect) error {
	buffer := [HidMessageSize]byte{
		LightingSetValueId,
		BacklightEffectId,
		effect.ToByte(),
	}
	return c.sendMessage(buffer[:], 20)
}

func (c *client) SetRgblightBrightness(brightness rgblight.Brightness) error {
	buffer := [HidMessageSize]byte{
		LightingSetValueId,
		RgblightBrightnessId,
		brightness.ToByte(),
	}
	return c.sendMessage(buffer[:], 20)
}

func (c *client) SetRgblightEffect(effect rgblight.Effect) error {
	buffer := [HidMessageSize]byte{
		LightingSetValueId,
		RgblightEffectId,
		effect.ToByte(),
	}

	// Send twice - if previous color mode is 0/Off then the first send will
	// enable solid color mode, not the desired mode
	err := c.sendMessage(buffer[:], 20)
	if err != nil {
		return err
	}

	return c.sendMessage(buffer[:], 20)
}

func (c *client) SetRgblightEffectSpeed(speed rgblight.Speed) error {
	buffer := [HidMessageSize]byte{
		LightingSetValueId,
		RgblightEffectSpeedId,
		speed.ToByte(),
	}
	return c.sendMessage(buffer[:], 20)
}

func (c *client) SetRgblightColor(color rgblight.Color, setBrightness bool) error {

	buffer := [HidMessageSize]byte{
		LightingSetValueId,
		RgblightColorId,
		color.Hue.ToByte(),
		color.Saturation.ToByte(),
	}
	err := c.sendMessage(buffer[:], 20)
	if err != nil {
		return err
	}
	if setBrightness {
		return c.SetRgblightBrightness(color.Brightness)
	}
	return err
}

func (c *client) SaveLighting() error {
	buffer := [HidMessageSize]byte{LightingSaveId}
	return c.sendMessage(buffer[:], 20)
}

func (c *client) ResetEeprom() error {
	buffer := [HidMessageSize]byte{EepromResetId}
	return c.sendMessage(buffer[:], 20)
}

func (c *client) GetDynamicKeymapMacroCount() (uint8, error) {
	buffer := [HidMessageSize]byte{DynamicKeymapMacroGetCountId}
	err := c.sendMessage(buffer[:], 20)
	if err != nil {
		return 0, err
	}

	return uint8(buffer[1]), nil
}

func (c *client) GetDynamicKeymapMacroBufferSize() (uint16, error) {
	buffer := [HidMessageSize]byte{DynamicKeymapMacroGetBufferSizeId}
	err := c.sendMessage(buffer[:], 20)
	if err != nil {
		return 0, err
	}

	return uint16(buffer[1])<<8 | uint16(buffer[2]), nil
}

func (c *client) GetDynamicKeymapMacroBuffer(offset uint16, size uint8) ([]byte, error) {
	buffer := [HidMessageSize]byte{
		DynamicKeymapMacroGetBufferId,
		byte((offset >> 8) & 0xFF),
		byte(offset & 0xFF),
		byte(size),
	}
	err := c.sendMessage(buffer[:], 20)

	if err != nil {
		return nil, err
	}

	return buffer[4:], nil
}

func (c *client) SetDynamicKeymapMacroBuffer(offset uint16, size uint8, value []byte) error {
	if size <= 0 || size > MaxDynamicKeymapBufferSize {
		return ErrorBadBufferSize
	}
	if uint8(len(value)) < size {
		size = uint8(len(value))
	}
	buffer := [HidMessageSize]byte{
		DynamicKeymapMacroSetBufferId,
		byte((offset >> 8) & 0xFF),
		byte(offset & 0xFF),
		byte(size),
	}
	for i := 0; i < int(size); i++ {
		buffer[i+4] = byte(value[i])
	}
	return c.sendMessage(buffer[:], 20)
}

func (c *client) ResetDynamicKeymapMacro() error {
	buffer := [HidMessageSize]byte{DynamicKeymapMacroResetId}
	return c.sendMessage(buffer[:], 20)
}

func (c *client) GetDynamicKeymapLayerCount() (uint8, error) {
	buffer := [HidMessageSize]byte{DynamicKeymapGetLayerCountId}
	err := c.sendMessage(buffer[:], 20)
	if err != nil {
		return 0, err
	}

	return uint8(buffer[1]), nil
}

func (c *client) GetDynamicKeymapBuffer(offset uint16, size uint8) ([]byte, error) {
	buffer := [HidMessageSize]byte{
		DynamicKeymapGetBufferId,
		byte((offset >> 8) & 0xFF),
		byte(offset & 0xFF),
		byte(size),
	}
	err := c.sendMessage(buffer[:], 20)

	if err != nil {
		return nil, err
	}

	return buffer[4:], nil
}

func (c *client) SetDynamicKeymapBuffer(offset uint16, size uint8, value []byte) error {
	if size <= 0 || size > MaxDynamicKeymapBufferSize {
		return ErrorBadBufferSize
	}
	if uint8(len(value)) < size {
		size = uint8(len(value))
	}
	buffer := [HidMessageSize]byte{
		DynamicKeymapSetBufferId,
		byte((offset >> 8) & 0xFF),
		byte(offset & 0xFF),
		byte(size),
	}
	for i := 0; i < int(size); i++ {
		buffer[i+4] = byte(value[i])
	}
	return c.sendMessage(buffer[:], 20)
}

func (c *client) getMacroBuffer() ([]byte, error) {
	if dynamicKeymapMacroCache != nil {
		return dynamicKeymapMacroCache, nil
	}

	size, err := c.GetDynamicKeymapMacroBufferSize()
	if err != nil {
		return nil, err
	}
	buffer := []byte{}

	// Read `reads` whole max sized blocks (overrun reads return 0x00)
	reads := int(math.Ceil(float64(size) / float64(MaxDynamicKeymapBufferSize)))
	for i := 0; i < reads; i++ {
		chunk, err := c.GetDynamicKeymapMacroBuffer(uint16(i*MaxDynamicKeymapBufferSize), MaxDynamicKeymapBufferSize)
		if err != nil {
			return nil, err
		}

		buffer = append(buffer, chunk...)
	}
	// Trim the buffer to max size
	dynamicKeymapMacroCache = buffer[:size]
	return dynamicKeymapMacroCache, nil
}

func (c *client) setMacroBuffer(buffer []byte) error {
	size, err := c.GetDynamicKeymapMacroBufferSize()
	if err != nil {
		return err
	}

	writes := int(math.Ceil(float64(size) / float64(MaxDynamicKeymapBufferSize)))
	bytes := writes * MaxDynamicKeymapBufferSize

	// Expand the buffer to the max size
	if cap(buffer) < bytes {
		buffer = append(buffer, make([]byte, bytes-cap(buffer))...)
	}

	// Clear the buffer cache
	dynamicKeymapMacroCache = nil

	// Write `writes` whole max sized blocks (overrun writes are noop)
	for i := 0; i < writes; i++ {
		start := i * MaxDynamicKeymapBufferSize
		end := start + MaxDynamicKeymapBufferSize
		err := c.SetDynamicKeymapMacroBuffer(uint16(start), MaxDynamicKeymapBufferSize, buffer[start:end])
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *client) GetDynamicKeymapMacro(number uint8) ([]byte, error) {
	count, err := c.GetDynamicKeymapMacroCount()
	if err != nil {
		return nil, err
	}
	if number > count {
		return nil, fmt.Errorf("empty or invalid macro number (0-%d)", count)
	}

	buffer, err := c.getMacroBuffer()
	if err != nil {
		return nil, err
	}

	macros := bytes.SplitAfterN(buffer, []byte("\x00"), -1)
	if int(number) >= len(macros) {
		return nil, ErrorMacroNotInBytes
	}
	macro := macros[number]
	return macro[:len(macro)-1], nil
}

func (c *client) SetDynamicKeymapMacro(number uint8, macro []byte) error {
	count, err := c.GetDynamicKeymapMacroCount()
	if err != nil {
		return err
	}
	if number > count {
		return fmt.Errorf("empty or invalid macro number (0-%d)", count)
	}

	buffer, err := c.getMacroBuffer()
	if err != nil {
		return err
	}

	macros := bytes.SplitAfterN(buffer, []byte("\x00"), -1)
	if int(number) >= len(macros) {
		return ErrorMacroNotInBytes
	}

	macros[number] = append(macro, byte(0))
	buffer = bytes.Join(macros, nil)

	return c.setMacroBuffer(buffer)
}
