// qmk-go - go client library for VIA-enabled QMK keyboards
// Copyright (c) 2022 Ian McLinden. All rights reserved
//
// This file is released under GNU LGPL 2.1 on Linux,
// and under the 3-clause BSD license on all other platforms

package backlight

import "testing"

var brightnessTests = []struct {
	Brightness Brightness
	Byte       byte
}{
	{Brightness(0), 0},
	{Brightness(50), 128},
	{Brightness(100), 255},
}

func TestBrightnessFromByte(t *testing.T) {
	for i, test := range brightnessTests {
		brightness := BrightnessFromByte(test.Byte)
		if test.Brightness != brightness {
			t.Errorf("[%d] wanted backlight brightness %v, got %v", i, test.Brightness, brightness)
		}
	}
}

func TestByteFromBrightness(t *testing.T) {
	for i, test := range brightnessTests {
		b := test.Brightness.ToByte()
		if test.Byte != b {
			t.Errorf("[%d] wanted backlight brightness byte %v, got %v", i, test.Byte, b)
		}
	}
}
