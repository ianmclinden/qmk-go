// qmk-go - go client library for VIA-enabled QMK keyboards
// Copyright (c) 2022 Ian McLinden. All rights reserved
//
// This file is released under GNU LGPL 2.1 on Linux,
// and under the 3-clause BSD license on all other platforms

package rgblight

import (
	"testing"
)

var hsvTests = []struct {
	Hue            Hue
	HueByte        byte
	Saturation     Saturation
	SaturationByte byte
	Brightness     Brightness
	BrightnessByte byte
}{
	{Hue(0), 0, Saturation(0), 0, Brightness(0), 0},
	{Hue(181), 128, Saturation(50), 128, Brightness(50), 128},
	{Hue(360), 255, Saturation(100), 255, Brightness(100), 255},
}

func TestHSVFromByte(t *testing.T) {
	for i, test := range hsvTests {
		hue := HueFromByte(test.HueByte)
		if test.Hue != hue {
			t.Errorf("[%d] wanted rgblight hue %v, got %v", i, test.Hue, hue)
		}
		saturation := SaturationFromByte(test.SaturationByte)
		if test.Saturation != saturation {
			t.Errorf("[%d] wanted rgblight saturation %v, got %v", i, test.Saturation, saturation)

		}
		brightness := BrightnessFromByte(test.BrightnessByte)
		if test.Brightness != brightness {
			t.Errorf("[%d] wanted rgblight brightness %v, got %v", i, test.Brightness, brightness)
		}
	}
}

func TestByteFromHSV(t *testing.T) {
	for i, test := range hsvTests {
		hue := test.Hue.ToByte()
		if test.HueByte != hue {
			t.Errorf("[%d] wanted rgblight hue byte %v, got %v", i, test.HueByte, hue)
		}
		saturation := test.Saturation.ToByte()
		if test.SaturationByte != saturation {
			t.Errorf("[%d] wanted rgblight saturation byte %v, got %v", i, test.SaturationByte, saturation)

		}
		brightness := test.Brightness.ToByte()
		if test.BrightnessByte != brightness {
			t.Errorf("[%d] wanted rgblight brightness byte %v, got %v", i, test.BrightnessByte, brightness)
		}
	}
}

var colorTests = []struct {
	Input    string
	Name     string
	HEX      string
	RGB      string
	HSV      string
	RGBColor rgbColor
	HSVColor hsvColor
	Color    Color
}{
	// Pre-defined colors
	/* 0*/ {"Black", "Black", "#000000", "rgb(0,0,0)", "hsv(0,0,0)", rgbColor{0, 0, 0}, hsvColor{0, 0, 0}, ColorBlack},
	/* 1*/ {"#000000", "Black", "#000000", "rgb(0,0,0)", "hsv(0,0,0)", rgbColor{0, 0, 0}, hsvColor{0, 0, 0}, ColorBlack},
	/* 2*/ {"rgb(0,0,0)", "Black", "#000000", "rgb(0,0,0)", "hsv(0,0,0)", rgbColor{0, 0, 0}, hsvColor{0, 0, 0}, ColorBlack},
	/* 3*/ {"hsv(0,0,0)", "Black", "#000000", "rgb(0,0,0)", "hsv(0,0,0)", rgbColor{0, 0, 0}, hsvColor{0, 0, 0}, ColorBlack},
	/* 4*/ {"White", "White", "#ffffff", "rgb(255,255,255)", "hsv(0,0,100)", rgbColor{255, 255, 255}, hsvColor{0, 0, 100}, ColorWhite},
	/* 5*/ {"#ffffff", "White", "#ffffff", "rgb(255,255,255)", "hsv(0,0,100)", rgbColor{255, 255, 255}, hsvColor{0, 0, 100}, ColorWhite},
	/* 6*/ {"rgb(255,255,255)", "White", "#ffffff", "rgb(255,255,255)", "hsv(0,0,100)", rgbColor{255, 255, 255}, hsvColor{0, 0, 100}, ColorWhite},
	/* 7*/ {"hsv(0,0,100)", "White", "#ffffff", "rgb(255,255,255)", "hsv(0,0,100)", rgbColor{255, 255, 255}, hsvColor{0, 0, 100}, ColorWhite},
	/* 8*/ {"Red", "Red", "#ff0000", "rgb(255,0,0)", "hsv(0,100,100)", rgbColor{255, 0, 0}, hsvColor{0, 100, 100}, ColorRed},
	/* 9*/ {"#ff0000", "Red", "#ff0000", "rgb(255,0,0)", "hsv(0,100,100)", rgbColor{255, 0, 0}, hsvColor{0, 100, 100}, ColorRed},
	/*10*/ {"rgb(255,0,0)", "Red", "#ff0000", "rgb(255,0,0)", "hsv(0,100,100)", rgbColor{255, 0, 0}, hsvColor{0, 100, 100}, ColorRed},
	/*11*/ {"hsv(0,100,100)", "Red", "#ff0000", "rgb(255,0,0)", "hsv(0,100,100)", rgbColor{255, 0, 0}, hsvColor{0, 100, 100}, ColorRed},
	/*12*/ {"Coral", "Coral", "#ff7e4f", "rgb(255,126,79)", "hsv(16,69,100)", rgbColor{255, 126, 79}, hsvColor{16, 69, 100}, ColorCoral},
	/*13*/ {"#ff7e4f", "Coral", "#ff7e4f", "rgb(255,126,79)", "hsv(16,69,100)", rgbColor{255, 126, 79}, hsvColor{16, 69, 100}, ColorCoral},
	/*14*/ {"rgb(255,126,79)", "Coral", "#ff7e4f", "rgb(255,126,79)", "hsv(16,69,100)", rgbColor{255, 126, 79}, hsvColor{16, 69, 100}, ColorCoral},
	/*15*/ {"hsv(16,69,100)", "Coral", "#ff7e4f", "rgb(255,126,79)", "hsv(16,69,100)", rgbColor{255, 126, 79}, hsvColor{16, 69, 100}, ColorCoral},
	/*16*/ {"Orange", "Orange", "#ffaa00", "rgb(255,170,0)", "hsv(40,100,100)", rgbColor{255, 170, 0}, hsvColor{40, 100, 100}, ColorOrange},
	/*17*/ {"#ffaa00", "Orange", "#ffaa00", "rgb(255,170,0)", "hsv(40,100,100)", rgbColor{255, 170, 0}, hsvColor{40, 100, 100}, ColorOrange},
	/*18*/ {"rgb(255,170,0)", "Orange", "#ffaa00", "rgb(255,170,0)", "hsv(40,100,100)", rgbColor{255, 170, 0}, hsvColor{40, 100, 100}, ColorOrange},
	/*19*/ {"hsv(40,100,100)", "Orange", "#ffaa00", "rgb(255,170,0)", "hsv(40,100,100)", rgbColor{255, 170, 0}, hsvColor{40, 100, 100}, ColorOrange},
	/*20*/ {"Gold", "Gold", "#d99800", "rgb(217,152,0)", "hsv(42,100,85)", rgbColor{217, 152, 0}, hsvColor{42, 100, 85}, ColorGold},
	/*21*/ {"#d99800", "Gold", "#d99800", "rgb(217,152,0)", "hsv(42,100,85)", rgbColor{217, 152, 0}, hsvColor{42, 100, 85}, ColorGold},
	/*22*/ {"rgb(217,152,0)", "Gold", "#d99800", "rgb(217,152,0)", "hsv(42,100,85)", rgbColor{217, 152, 0}, hsvColor{42, 100, 85}, ColorGold},
	/*23*/ {"hsv(42,100,85)", "Gold", "#d99800", "rgb(217,152,0)", "hsv(42,100,85)", rgbColor{217, 152, 0}, hsvColor{42, 100, 85}, ColorGold},
	/*24*/ {"Goldenrod", "Goldenrod", "#d9a521", "rgb(217,165,33)", "hsv(43,85,85)", rgbColor{217, 165, 33}, hsvColor{43, 85, 85}, ColorGoldenrod},
	/*25*/ {"#d9a521", "Goldenrod", "#d9a521", "rgb(217,165,33)", "hsv(43,85,85)", rgbColor{217, 165, 33}, hsvColor{43, 85, 85}, ColorGoldenrod},
	/*26*/ {"rgb(217,165,33)", "Goldenrod", "#d9a521", "rgb(217,165,33)", "hsv(43,85,85)", rgbColor{217, 165, 33}, hsvColor{43, 85, 85}, ColorGoldenrod},
	/*27*/ {"hsv(43,85,85)", "Goldenrod", "#d9a521", "rgb(217,165,33)", "hsv(43,85,85)", rgbColor{217, 165, 33}, hsvColor{43, 85, 85}, ColorGoldenrod},
	/*28*/ {"Yellow", "Yellow", "#fbff00", "rgb(251,255,0)", "hsv(61,100,100)", rgbColor{251, 255, 0}, hsvColor{61, 100, 100}, ColorYellow},
	/*29*/ {"#fbff00", "Yellow", "#fbff00", "rgb(251,255,0)", "hsv(61,100,100)", rgbColor{251, 255, 0}, hsvColor{61, 100, 100}, ColorYellow},
	/*30*/ {"rgb(251,255,0)", "Yellow", "#fbff00", "rgb(251,255,0)", "hsv(61,100,100)", rgbColor{251, 255, 0}, hsvColor{61, 100, 100}, ColorYellow},
	/*31*/ {"hsv(61,100,100)", "Yellow", "#fbff00", "rgb(251,255,0)", "hsv(61,100,100)", rgbColor{251, 255, 0}, hsvColor{61, 100, 100}, ColorYellow},
	/*32*/ {"Chartreuse", "Chartreuse", "#80ff00", "rgb(128,255,0)", "hsv(90,100,100)", rgbColor{128, 255, 0}, hsvColor{90, 100, 100}, ColorChartreuse},
	/*33*/ {"#80ff00", "Chartreuse", "#80ff00", "rgb(128,255,0)", "hsv(90,100,100)", rgbColor{128, 255, 0}, hsvColor{90, 100, 100}, ColorChartreuse},
	/*34*/ {"rgb(128,255,0)", "Chartreuse", "#80ff00", "rgb(128,255,0)", "hsv(90,100,100)", rgbColor{128, 255, 0}, hsvColor{90, 100, 100}, ColorChartreuse},
	/*35*/ {"hsv(90,100,100)", "Chartreuse", "#80ff00", "rgb(128,255,0)", "hsv(90,100,100)", rgbColor{128, 255, 0}, hsvColor{90, 100, 100}, ColorChartreuse},
	/*36*/ {"Green", "Green", "#00ff00", "rgb(0,255,0)", "hsv(120,100,100)", rgbColor{0, 255, 0}, hsvColor{120, 100, 100}, ColorGreen},
	/*37*/ {"#00ff00", "Green", "#00ff00", "rgb(0,255,0)", "hsv(120,100,100)", rgbColor{0, 255, 0}, hsvColor{120, 100, 100}, ColorGreen},
	/*38*/ {"rgb(0,255,0)", "Green", "#00ff00", "rgb(0,255,0)", "hsv(120,100,100)", rgbColor{0, 255, 0}, hsvColor{120, 100, 100}, ColorGreen},
	/*39*/ {"hsv(120,100,100)", "Green", "#00ff00", "rgb(0,255,0)", "hsv(120,100,100)", rgbColor{0, 255, 0}, hsvColor{120, 100, 100}, ColorGreen},
	/*40*/ {"SpringGreen", "SpringGreen", "#00ff80", "rgb(0,255,128)", "hsv(150,100,100)", rgbColor{0, 255, 128}, hsvColor{150, 100, 100}, ColorSpringGreen},
	/*41*/ {"#00ff80", "SpringGreen", "#00ff80", "rgb(0,255,128)", "hsv(150,100,100)", rgbColor{0, 255, 128}, hsvColor{150, 100, 100}, ColorSpringGreen},
	/*42*/ {"rgb(0,255,128)", "SpringGreen", "#00ff80", "rgb(0,255,128)", "hsv(150,100,100)", rgbColor{0, 255, 128}, hsvColor{150, 100, 100}, ColorSpringGreen},
	/*43*/ {"hsv(150,100,100)", "SpringGreen", "#00ff80", "rgb(0,255,128)", "hsv(150,100,100)", rgbColor{0, 255, 128}, hsvColor{150, 100, 100}, ColorSpringGreen},
	/*44*/ {"Turquoise", "Turquoise", "#49706c", "rgb(73,112,108)", "hsv(174,35,44)", rgbColor{73, 112, 108}, hsvColor{174, 35, 44}, ColorTurquoise},
	/*45*/ {"#49706c", "Turquoise", "#49706c", "rgb(73,112,108)", "hsv(174,35,44)", rgbColor{73, 112, 108}, hsvColor{174, 35, 44}, ColorTurquoise},
	/*46*/ {"rgb(73,112,108)", "Turquoise", "#49706c", "rgb(73,112,108)", "hsv(174,35,44)", rgbColor{73, 112, 108}, hsvColor{174, 35, 44}, ColorTurquoise},
	/*47*/ {"hsv(174,35,44)", "Turquoise", "#49706c", "rgb(73,112,108)", "hsv(174,35,44)", rgbColor{73, 112, 108}, hsvColor{174, 35, 44}, ColorTurquoise},
	/*48*/ {"Teal", "Teal", "#007d80", "rgb(0,125,128)", "hsv(181,100,50)", rgbColor{0, 125, 128}, hsvColor{181, 100, 50}, ColorTeal},
	/*49*/ {"#007d80", "Teal", "#007d80", "rgb(0,125,128)", "hsv(181,100,50)", rgbColor{0, 125, 128}, hsvColor{181, 100, 50}, ColorTeal},
	/*50*/ {"rgb(0,125,128)", "Teal", "#007d80", "rgb(0,125,128)", "hsv(181,100,50)", rgbColor{0, 125, 128}, hsvColor{181, 100, 50}, ColorTeal},
	/*51*/ {"hsv(181,100,50)", "Teal", "#007d80", "rgb(0,125,128)", "hsv(181,100,50)", rgbColor{0, 125, 128}, hsvColor{181, 100, 50}, ColorTeal},
	/*52*/ {"Cyan", "Cyan", "#00fbff", "rgb(0,251,255)", "hsv(181,100,100)", rgbColor{0, 251, 255}, hsvColor{181, 100, 100}, ColorCyan},
	/*53*/ {"#00fbff", "Cyan", "#00fbff", "rgb(0,251,255)", "hsv(181,100,100)", rgbColor{0, 251, 255}, hsvColor{181, 100, 100}, ColorCyan},
	/*54*/ {"rgb(0,251,255)", "Cyan", "#00fbff", "rgb(0,251,255)", "hsv(181,100,100)", rgbColor{0, 251, 255}, hsvColor{181, 100, 100}, ColorCyan},
	/*55*/ {"hsv(181,100,100)", "Cyan", "#00fbff", "rgb(0,251,255)", "hsv(181,100,100)", rgbColor{0, 251, 255}, hsvColor{181, 100, 100}, ColorCyan},
	/*56*/ {"Azure", "Azure", "#99f5ff", "rgb(153,245,255)", "hsv(186,40,100)", rgbColor{153, 245, 255}, hsvColor{186, 40, 100}, ColorAzure},
	/*57*/ {"#99f5ff", "Azure", "#99f5ff", "rgb(153,245,255)", "hsv(186,40,100)", rgbColor{153, 245, 255}, hsvColor{186, 40, 100}, ColorAzure},
	/*58*/ {"rgb(153,245,255)", "Azure", "#99f5ff", "rgb(153,245,255)", "hsv(186,40,100)", rgbColor{153, 245, 255}, hsvColor{186, 40, 100}, ColorAzure},
	/*59*/ {"hsv(186,40,100)", "Azure", "#99f5ff", "rgb(153,245,255)", "hsv(186,40,100)", rgbColor{153, 245, 255}, hsvColor{186, 40, 100}, ColorAzure},
	/*60*/ {"Blue", "Blue", "#0000ff", "rgb(0,0,255)", "hsv(240,100,100)", rgbColor{0, 0, 255}, hsvColor{240, 100, 100}, ColorBlue},
	/*61*/ {"#0000ff", "Blue", "#0000ff", "rgb(0,0,255)", "hsv(240,100,100)", rgbColor{0, 0, 255}, hsvColor{240, 100, 100}, ColorBlue},
	/*62*/ {"rgb(0,0,255)", "Blue", "#0000ff", "rgb(0,0,255)", "hsv(240,100,100)", rgbColor{0, 0, 255}, hsvColor{240, 100, 100}, ColorBlue},
	/*63*/ {"hsv(240,100,100)", "Blue", "#0000ff", "rgb(0,0,255)", "hsv(240,100,100)", rgbColor{0, 0, 255}, hsvColor{240, 100, 100}, ColorBlue},
	/*64*/ {"Purple", "Purple", "#8000ff", "rgb(128,0,255)", "hsv(270,100,100)", rgbColor{128, 0, 255}, hsvColor{270, 100, 100}, ColorPurple},
	/*65*/ {"#8000ff", "Purple", "#8000ff", "rgb(128,0,255)", "hsv(270,100,100)", rgbColor{128, 0, 255}, hsvColor{270, 100, 100}, ColorPurple},
	/*66*/ {"rgb(128,0,255)", "Purple", "#8000ff", "rgb(128,0,255)", "hsv(270,100,100)", rgbColor{128, 0, 255}, hsvColor{270, 100, 100}, ColorPurple},
	/*67*/ {"hsv(270,100,100)", "Purple", "#8000ff", "rgb(128,0,255)", "hsv(270,100,100)", rgbColor{128, 0, 255}, hsvColor{270, 100, 100}, ColorPurple},
	/*68*/ {"Magenta", "Magenta", "#ff00fb", "rgb(255,0,251)", "hsv(301,100,100)", rgbColor{255, 0, 251}, hsvColor{301, 100, 100}, ColorMagenta},
	/*69*/ {"#ff00fb", "Magenta", "#ff00fb", "rgb(255,0,251)", "hsv(301,100,100)", rgbColor{255, 0, 251}, hsvColor{301, 100, 100}, ColorMagenta},
	/*70*/ {"rgb(255,0,251)", "Magenta", "#ff00fb", "rgb(255,0,251)", "hsv(301,100,100)", rgbColor{255, 0, 251}, hsvColor{301, 100, 100}, ColorMagenta},
	/*71*/ {"hsv(301,100,100)", "Magenta", "#ff00fb", "rgb(255,0,251)", "hsv(301,100,100)", rgbColor{255, 0, 251}, hsvColor{301, 100, 100}, ColorMagenta},
	/*72*/ {"Pink", "Pink", "#ff80bf", "rgb(255,128,191)", "hsv(330,50,100)", rgbColor{255, 128, 191}, hsvColor{330, 50, 100}, ColorPink},
	/*73*/ {"#ff80bf", "Pink", "#ff80bf", "rgb(255,128,191)", "hsv(330,50,100)", rgbColor{255, 128, 191}, hsvColor{330, 50, 100}, ColorPink},
	/*74*/ {"rgb(255,128,191)", "Pink", "#ff80bf", "rgb(255,128,191)", "hsv(330,50,100)", rgbColor{255, 128, 191}, hsvColor{330, 50, 100}, ColorPink},
	/*75*/ {"hsv(330,50,100)", "Pink", "#ff80bf", "rgb(255,128,191)", "hsv(330,50,100)", rgbColor{255, 128, 191}, hsvColor{330, 50, 100}, ColorPink},
	// Custom colors - no test from name
	/*76*/ {"#012345", "Unknown", "#012345", "rgb(1,35,69)", "hsv(210,99,27)", rgbColor{1, 35, 69}, hsvColor{210, 99, 27}, Color{210, 99, 27}},
	/*77*/ {"rgb(1,35,69)", "Unknown", "#012345", "rgb(1,35,69)", "hsv(210,99,27)", rgbColor{1, 35, 69}, hsvColor{210, 99, 27}, Color{210, 99, 27}},
	/*78*/ {"hsv(210,99,27)", "Unknown", "#012345", "rgb(1,35,69)", "hsv(210,99,27)", rgbColor{1, 35, 69}, hsvColor{210, 99, 27}, Color{210, 99, 27}}}

func TestColorFromString(t *testing.T) {
	for i, test := range colorTests {
		color, err := ColorFromString(test.Input)
		if err != nil {
			t.Errorf("[%d] %s", i, err.Error())
		}
		if test.Color != color {
			t.Errorf("[%d] wanted rgblight color %v, got %v", i, test.Color, color)
		}
	}

}

func TestColorToName(t *testing.T) {
	for i, test := range colorTests {
		name := test.Color.Name()
		if test.Name != name {
			t.Errorf("[%d] wanted rgblight color name %v, got %v", i, test.Name, name)
		}
	}
}

func TestColorToRGB(t *testing.T) {
	for i, test := range colorTests {
		rgb := test.Color.ToStringRGB()
		if test.RGB != rgb {
			t.Errorf("[%d] wanted rgblight color rgb string %v, got %v", i, test.RGB, rgb)
		}
	}
}

func TestColorToHSV(t *testing.T) {
	for i, test := range colorTests {
		hsv := test.Color.ToStringHSV()
		if test.HSV != hsv {
			t.Errorf("[%d] wanted rgblight color hsv string %v, got %v", i, test.HSV, hsv)
		}
	}
}

func TestColorToHEX(t *testing.T) {
	for i, test := range colorTests {
		hex := test.Color.ToStringHEX()
		if test.HEX != hex {
			t.Errorf("[%d] wanted rgblight color hex string %v, got %v", i, test.HEX, hex)
		}
	}
}

func TestRGBToHSV(t *testing.T) {
	for i, test := range colorTests {
		hsv := test.RGBColor.toHSV()
		if test.HSVColor != hsv {
			t.Errorf("[%d] wanted rgblight hsvColor %v, got %v", i, test.HSVColor, hsv)
		}
	}
}

func TestHSVToRGB(t *testing.T) {
	for i, test := range colorTests {
		rgb := test.HSVColor.toRGB()
		if test.RGBColor != rgb {
			t.Errorf("[%d] wanted rgblight rgbColor %v, got %v", i, test.RGBColor, rgb)
		}
	}
}

func TestAllColors(t *testing.T) {
	colors := AllColors()
	if len(colors) != 19 {
		t.Errorf("wanted %d rgblight colors, got %d", 19, len(colors))
	}
	for i, color := range colors {
		for _, test := range colorTests {
			if test.Color == color {
				return // we have at least some test coverage
			}
		}
		t.Errorf("rgblight color [%d] %v has no test coverage", i, color)
	}
}
