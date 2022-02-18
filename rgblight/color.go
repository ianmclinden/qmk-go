// qmk-go - go client library for VIA-enabled QMK keyboards
// Copyright (c) 2022 Ian McLinden. All rights reserved
//
// This file is released under GNU LGPL 2.1 on Linux,
// and under the 3-clause BSD license on all other platforms

package rgblight

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

var (
	ErrorUnknownColorFormat = errors.New("unknown or invalid color format")
)

type (
	Hue        uint16 // 0-360°
	Saturation uint8  // 0-100%
	Brightness uint8  //0-100%
)

func HueFromByte(value byte) Hue {
	return Hue(int(math.Round((float64(value) * 360.0) / 255.0)))
}

func (h Hue) ToByte() byte {
	return byte(math.Round((float64(h) * 255.0) / 360.0))
}

func SaturationFromByte(value byte) Saturation {
	return Saturation(int(math.Round((float64(value) * 100.0) / 255.0)))
}

func (s Saturation) ToByte() byte {
	return byte(math.Round((float64(s) * 255.0) / 100.0))
}

func BrightnessFromByte(value byte) Brightness {
	return Brightness(int(math.Round((float64(value) * 100.0) / 255.0)))
}

func (b Brightness) ToByte() byte {
	return byte(math.Round((float64(b) * 255.0) / 100.0))
}

type hsvColor struct {
	Hue        Hue        // 0-360°
	Saturation Saturation // 0-100%
	Brightness Brightness // 0-100%
}

func (c hsvColor) toString() string {
	return fmt.Sprintf("hsv(%d,%d,%d)", c.Hue, c.Saturation, c.Brightness)
}

func (c hsvColor) toRGB() rgbColor {
	var (
		h       = float64(c.Hue)
		s       = float64(c.Saturation) / 100.0
		v       = float64(c.Brightness) / 100.0
		d       = s * v
		x       = d * (1 - math.Abs(math.Mod(h/60.0, 2)-1))
		m       = v - d
		r, g, b float64
	)

	if h >= 0 && h < 60 {
		r, g, b = d, x, 0
	} else if h >= 60 && h < 120 {
		r, g, b = x, d, 0
	} else if h >= 120 && h < 180 {
		r, g, b = 0, d, x
	} else if h >= 180 && h < 240 {
		r, g, b = 0, x, d
	} else if h >= 240 && h < 300 {
		r, g, b = x, 0, d
	} else {
		r, g, b = d, 0, x
	}
	var (
		rv = uint8(math.Round((r + m) * 255.0))
		gv = uint8(math.Round((g + m) * 255.0))
		bv = uint8(math.Round((b + m) * 255.0))
	)
	return rgbColor{Red(rv), Green(gv), Blue(bv)}
}

func (c hsvColor) toColor() Color {
	return Color(c)
}

type (
	Red   byte // 0-255
	Green byte // 0-255
	Blue  byte // 0-255
)

type rgbColor struct {
	Red   Red   // 0-255
	Green Green // 0-255
	Blue  Blue  // 0-255
}

func (c rgbColor) toString() string {
	return fmt.Sprintf("rgb(%d,%d,%d)", c.Red, c.Green, c.Blue)
}

func (c rgbColor) toHexString() string {
	return fmt.Sprintf("#%02x%02x%02x", c.Red, c.Green, c.Blue)
}

func (c rgbColor) toColor() Color {
	return c.toHSV().toColor()
}

func (c rgbColor) toHSV() hsvColor {
	var (
		r           = float64(c.Red)
		g           = float64(c.Green)
		b           = float64(c.Blue)
		cmax        = math.Max(r, math.Max(g, b))
		cmin        = math.Min(r, math.Min(g, b))
		diff        = cmax - cmin
		h    uint16 = 0
		s    uint8  = uint8(math.Round(diff * 100.0 / cmax))
		l    uint8  = uint8(math.Round((cmax * 100.0) / 255.0))
	)

	if cmax == 0 {
		s = 0
	} else if cmax == r {
		h = uint16(math.Round(60*((g-b)/diff)+360)) % 360
	} else if cmax == g {
		h = uint16(math.Round(60*((b-r)/diff)+120)) % 360
	} else if cmax == b {
		h = uint16(math.Round(60*((r-g)/diff)+240)) % 360
	}

	return hsvColor{Hue(h), Saturation(s), Brightness(l)}
}

type Color hsvColor

// QMK Default Colors as HSV (19)
var (
	ColorBlack       = Color{0, 0, 0}
	ColorWhite       = Color{0, 0, 100}
	ColorRed         = Color{0, 100, 100}
	ColorCoral       = Color{16, 69, 100}
	ColorGold        = Color{42, 100, 85}
	ColorOrange      = Color{40, 100, 100}
	ColorGoldenrod   = Color{43, 85, 85}
	ColorYellow      = Color{61, 100, 100}
	ColorChartreuse  = Color{90, 100, 100}
	ColorGreen       = Color{120, 100, 100}
	ColorSpringGreen = Color{150, 100, 100}
	ColorTurquoise   = Color{174, 35, 44}
	ColorTeal        = Color{181, 100, 50}
	ColorCyan        = Color{181, 100, 100}
	ColorAzure       = Color{186, 40, 100}
	ColorBlue        = Color{240, 100, 100}
	ColorPurple      = Color{270, 100, 100}
	ColorMagenta     = Color{301, 100, 100}
	ColorPink        = Color{330, 50, 100}
)

var (
	ColorOff = ColorBlack
)

func (c Color) toHSV() hsvColor {
	return hsvColor(c)
}

func (c Color) toRGB() rgbColor {
	return hsvColor(c).toRGB()
}

func (c Color) ToStringHSV() string {
	return c.toHSV().toString()
}

func (c Color) ToStringRGB() string {
	return c.toRGB().toString()
}

func (c Color) ToStringHEX() string {
	return c.toRGB().toHexString()
}

func AllColors() []Color {
	return []Color{
		ColorBlack, ColorWhite, ColorRed, ColorCoral, ColorOrange, ColorGold, ColorGoldenrod, ColorYellow, ColorChartreuse, ColorGreen, ColorSpringGreen, ColorTurquoise, ColorTeal, ColorCyan, ColorAzure, ColorBlue, ColorPurple, ColorMagenta, ColorPink,
	}
}

func (c Color) Name() string {
	switch c {
	case ColorBlack:
		return "Black"
	case ColorWhite:
		return "White"
	case ColorRed:
		return "Red"
	case ColorCoral:
		return "Coral"
	case ColorOrange:
		return "Orange"
	case ColorGold:
		return "Gold"
	case ColorGoldenrod:
		return "Goldenrod"
	case ColorYellow:
		return "Yellow"
	case ColorChartreuse:
		return "Chartreuse"
	case ColorGreen:
		return "Green"
	case ColorSpringGreen:
		return "SpringGreen"
	case ColorTurquoise:
		return "Turquoise"
	case ColorTeal:
		return "Teal"
	case ColorCyan:
		return "Cyan"
	case ColorAzure:
		return "Azure"
	case ColorBlue:
		return "Blue"
	case ColorPurple:
		return "Purple"
	case ColorMagenta:
		return "Magenta"
	case ColorPink:
		return "Pink"
	default:
		return "Unknown"
	}
}

func ColorFromString(value string) (Color, error) {
	s := strings.ToLower(value)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "color", "", -1)

	if strings.HasPrefix(s, "hsv") {
		return parseHSV(s)
	}
	if strings.HasPrefix(s, "rgb") {
		return parseRGB(s)
	}
	if strings.HasPrefix(s, "#") {
		return parseHex(s)
	}

	switch s {
	case "black":
		return ColorBlack, nil
	case "white":
		return ColorWhite, nil
	case "red":
		return ColorRed, nil
	case "coral":
		return ColorCoral, nil
	case "orange":
		return ColorOrange, nil
	case "gold":
		return ColorGold, nil
	case "goldenrod":
		return ColorGoldenrod, nil
	case "yellow":
		return ColorYellow, nil
	case "chartreuse":
		return ColorChartreuse, nil
	case "green":
		return ColorGreen, nil
	case "springgreen":
		return ColorSpringGreen, nil
	case "turquoise":
		return ColorTurquoise, nil
	case "teal":
		return ColorTeal, nil
	case "cyan":
		return ColorCyan, nil
	case "azure":
		return ColorAzure, nil
	case "blue":
		return ColorBlue, nil
	case "purple":
		return ColorPurple, nil
	case "magenta":
		return ColorMagenta, nil
	case "pink":
		return ColorPink, nil
	default:
		return ColorOff, ErrorUnknownColorFormat
	}
}

func parseHSV(value string) (Color, error) {
	str := strings.Replace(value, "hsv(", "", -1)
	str = strings.Replace(str, ")", "", -1)

	var (
		h    uint16
		s, v uint8
	)
	n, err := fmt.Sscanf(str, "%d,%d,%d", &h, &s, &v)
	if err != nil || n != 3 {
		return ColorOff, err
	}
	if h > 360 || s > 100 || v > 100 {
		return ColorOff, ErrorUnknownColorFormat
	}

	hsv := hsvColor{Hue(h), Saturation(s), Brightness(v)}
	return hsv.toColor(), nil
}

func parseRGB(value string) (Color, error) {
	s := strings.Replace(value, "rgb(", "", -1)
	s = strings.Replace(s, ")", "", -1)

	var r, g, b uint8
	n, err := fmt.Sscanf(s, "%d,%d,%d", &r, &g, &b)
	if err != nil || n != 3 {
		return ColorOff, err
	}

	rgb := rgbColor{Red(r), Green(g), Blue(b)}
	return rgb.toColor(), nil
}

func parseHex(value string) (Color, error) {
	s := strings.Replace(value, "#", "", -1)
	if len(s) != 6 {
		return ColorOff, ErrorUnknownColorFormat
	}

	var r, g, b uint8
	n, err := fmt.Sscanf(s, "%2x%2x%2x", &r, &g, &b)
	if err != nil || n != 3 {
		return ColorOff, err
	}

	rgb := rgbColor{Red(r), Green(g), Blue(b)}
	return rgb.toColor(), nil
}
