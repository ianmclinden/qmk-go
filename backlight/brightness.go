// qmk-go - go client library for VIA-enabled QMK keyboards
// Copyright (c) 2022 Ian McLinden. All rights reserved
//
// This file is released under GNU LGPL 2.1 on Linux,
// and under the 3-clause BSD license on all other platforms

package backlight

import (
	"math"
)

type Brightness uint8

func BrightnessFromByte(value byte) Brightness {
	return Brightness(int(math.Round((float64(value) * 100.0) / 255.0)))
}

func (b Brightness) ToByte() byte {
	return byte(math.Round((float64(b) * 255.0) / 100.0))
}
