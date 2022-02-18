// qmk-go - go client library for VIA-enabled QMK keyboards
// Copyright (c) 2022 Ian McLinden. All rights reserved
//
// This file is released under GNU LGPL 2.1 on Linux,
// and under the 3-clause BSD license on all other platforms

package rgblight

import (
	"testing"
)

var speedTests = []struct {
	Speed Speed
	Byte  byte
}{
	{Speed(0), 0},
	{Speed(50), 128},
	{Speed(100), 255},
}

func TestSpeedFromByte(t *testing.T) {
	for i, test := range speedTests {
		speed := SpeedFromByte(test.Byte)
		if test.Speed != speed {
			t.Errorf("[%d] wanted rgblight speed %v, got %v", i, test.Speed, speed)
		}
	}
}

func TestByteFromSpeed(t *testing.T) {
	for i, test := range speedTests {
		b := test.Speed.ToByte()
		if test.Byte != b {
			t.Errorf("[%d] wanted rgblight speed byte %v, got %v", i, test.Byte, b)
		}
	}
}

var effectTests = []struct {
	Input  string
	Name   string
	Effect Effect
}{
	/*  0*/ {"AllOff", "All Off", EffectAllOff},
	/*  1*/ {"alloff", "All Off", EffectAllOff},
	/*  2*/ {"All Off ", "All Off", EffectAllOff},
	/*  3*/ {"all off ", "All Off", EffectAllOff},
	/*  4*/ {"SolidColor", "Solid Color", EffectSolidColor},
	/*  5*/ {"solidcolor", "Solid Color", EffectSolidColor},
	/*  6*/ {"Solid Color ", "Solid Color", EffectSolidColor},
	/*  7*/ {"solid color ", "Solid Color", EffectSolidColor},
	/*  8*/ {"Breathing1", "Breathing 1", EffectBreathing1},
	/*  9*/ {"breathing1", "Breathing 1", EffectBreathing1},
	/* 10*/ {"Breathing 1", "Breathing 1", EffectBreathing1},
	/* 11*/ {"breathing 1", "Breathing 1", EffectBreathing1},
	/* 12*/ {"Breathing2", "Breathing 2", EffectBreathing2},
	/* 13*/ {"breathing2", "Breathing 2", EffectBreathing2},
	/* 14*/ {"Breathing 2", "Breathing 2", EffectBreathing2},
	/* 15*/ {"breathing 2", "Breathing 2", EffectBreathing2},
	/* 16*/ {"Breathing3", "Breathing 3", EffectBreathing3},
	/* 17*/ {"breathing3", "Breathing 3", EffectBreathing3},
	/* 18*/ {"Breathing 3", "Breathing 3", EffectBreathing3},
	/* 19*/ {"breathing 3", "Breathing 3", EffectBreathing3},
	/* 20*/ {"Breathing4", "Breathing 4", EffectBreathing4},
	/* 21*/ {"breathing4", "Breathing 4", EffectBreathing4},
	/* 22*/ {"Breathing 4", "Breathing 4", EffectBreathing4},
	/* 23*/ {"breathing 4", "Breathing 4", EffectBreathing4},
	/* 24*/ {"RainbowMood1", "Rainbow Mood 1", EffectRainbowMood1},
	/* 25*/ {"rainbowMood1", "Rainbow Mood 1", EffectRainbowMood1},
	/* 26*/ {"RainbowMood 1", "Rainbow Mood 1", EffectRainbowMood1},
	/* 27*/ {"rainbowMood 1", "Rainbow Mood 1", EffectRainbowMood1},
	/* 28*/ {"RainbowMood2", "Rainbow Mood 2", EffectRainbowMood2},
	/* 29*/ {"rainbowMood2", "Rainbow Mood 2", EffectRainbowMood2},
	/* 30*/ {"RainbowMood 2", "Rainbow Mood 2", EffectRainbowMood2},
	/* 31*/ {"rainbowMood 2", "Rainbow Mood 2", EffectRainbowMood2},
	/* 32*/ {"RainbowMood3", "Rainbow Mood 3", EffectRainbowMood3},
	/* 33*/ {"rainbowMood3", "Rainbow Mood 3", EffectRainbowMood3},
	/* 34*/ {"RainbowMood 3", "Rainbow Mood 3", EffectRainbowMood3},
	/* 35*/ {"rainbowMood 3", "Rainbow Mood 3", EffectRainbowMood3},
	/* 36*/ {"RainbowSwirl1", "Rainbow Swirl 1", EffectRainbowSwirl1},
	/* 37*/ {"rainbowSwirl1", "Rainbow Swirl 1", EffectRainbowSwirl1},
	/* 38*/ {"RainbowSwirl 1", "Rainbow Swirl 1", EffectRainbowSwirl1},
	/* 39*/ {"rainbowSwirl 1", "Rainbow Swirl 1", EffectRainbowSwirl1},
	/* 40*/ {"RainbowSwirl2", "Rainbow Swirl 2", EffectRainbowSwirl2},
	/* 41*/ {"rainbowSwirl2", "Rainbow Swirl 2", EffectRainbowSwirl2},
	/* 42*/ {"RainbowSwirl 2", "Rainbow Swirl 2", EffectRainbowSwirl2},
	/* 43*/ {"rainbowSwirl 2", "Rainbow Swirl 2", EffectRainbowSwirl2},
	/* 44*/ {"RainbowSwirl3", "Rainbow Swirl 3", EffectRainbowSwirl3},
	/* 45*/ {"rainbowSwirl3", "Rainbow Swirl 3", EffectRainbowSwirl3},
	/* 46*/ {"RainbowSwirl 3", "Rainbow Swirl 3", EffectRainbowSwirl3},
	/* 47*/ {"rainbowSwirl 3", "Rainbow Swirl 3", EffectRainbowSwirl3},
	/* 48*/ {"RainbowSwirl4", "Rainbow Swirl 4", EffectRainbowSwirl4},
	/* 49*/ {"rainbowSwirl4", "Rainbow Swirl 4", EffectRainbowSwirl4},
	/* 50*/ {"RainbowSwirl 4", "Rainbow Swirl 4", EffectRainbowSwirl4},
	/* 51*/ {"rainbowSwirl 4", "Rainbow Swirl 4", EffectRainbowSwirl4},
	/* 52*/ {"RainbowSwirl5", "Rainbow Swirl 5", EffectRainbowSwirl5},
	/* 53*/ {"rainbowSwirl5", "Rainbow Swirl 5", EffectRainbowSwirl5},
	/* 54*/ {"RainbowSwirl 5", "Rainbow Swirl 5", EffectRainbowSwirl5},
	/* 55*/ {"rainbowSwirl 5", "Rainbow Swirl 5", EffectRainbowSwirl5},
	/* 56*/ {"RainbowSwirl6", "Rainbow Swirl 6", EffectRainbowSwirl6},
	/* 57*/ {"rainbowSwirl6", "Rainbow Swirl 6", EffectRainbowSwirl6},
	/* 58*/ {"RainbowSwirl 6", "Rainbow Swirl 6", EffectRainbowSwirl6},
	/* 59*/ {"rainbowSwirl 6", "Rainbow Swirl 6", EffectRainbowSwirl6},
	/* 60*/ {"Snake1", "Snake 1", EffectSnake1},
	/* 61*/ {"snake1", "Snake 1", EffectSnake1},
	/* 62*/ {"Snake 1", "Snake 1", EffectSnake1},
	/* 63*/ {"snake 1", "Snake 1", EffectSnake1},
	/* 64*/ {"Snake2", "Snake 2", EffectSnake2},
	/* 65*/ {"snake2", "Snake 2", EffectSnake2},
	/* 66*/ {"Snake 2", "Snake 2", EffectSnake2},
	/* 67*/ {"snake 2", "Snake 2", EffectSnake2},
	/* 68*/ {"Snake3", "Snake 3", EffectSnake3},
	/* 69*/ {"snake3", "Snake 3", EffectSnake3},
	/* 70*/ {"Snake 3", "Snake 3", EffectSnake3},
	/* 71*/ {"snake 3", "Snake 3", EffectSnake3},
	/* 72*/ {"Snake4", "Snake 4", EffectSnake4},
	/* 73*/ {"snake4", "Snake 4", EffectSnake4},
	/* 74*/ {"Snake 4", "Snake 4", EffectSnake4},
	/* 75*/ {"snake 4", "Snake 4", EffectSnake4},
	/* 76*/ {"Snake5", "Snake 5", EffectSnake5},
	/* 77*/ {"snake5", "Snake 5", EffectSnake5},
	/* 78*/ {"Snake 5", "Snake 5", EffectSnake5},
	/* 79*/ {"snake 5", "Snake 5", EffectSnake5},
	/* 80*/ {"Snake6", "Snake 6", EffectSnake6},
	/* 81*/ {"snake6", "Snake 6", EffectSnake6},
	/* 82*/ {"Snake 6", "Snake 6", EffectSnake6},
	/* 83*/ {"snake 6", "Snake 6", EffectSnake6},
	/* 84*/ {"Knight1", "Knight 1", EffectKnight1},
	/* 85*/ {"knight1", "Knight 1", EffectKnight1},
	/* 86*/ {"Knight 1", "Knight 1", EffectKnight1},
	/* 87*/ {"knight 1", "Knight 1", EffectKnight1},
	/* 88*/ {"Knight2", "Knight 2", EffectKnight2},
	/* 89*/ {"knight2", "Knight 2", EffectKnight2},
	/* 90*/ {"Knight 2", "Knight 2", EffectKnight2},
	/* 91*/ {"knight 2", "Knight 2", EffectKnight2},
	/* 92*/ {"Knight3", "Knight 3", EffectKnight3},
	/* 93*/ {"knight3", "Knight 3", EffectKnight3},
	/* 94*/ {"Knight 3", "Knight 3", EffectKnight3},
	/* 95*/ {"knight 3", "Knight 3", EffectKnight3},
	/* 96*/ {"Christmas", "Christmas", EffectChristmas},
	/* 97*/ {"christmas", "Christmas", EffectChristmas},
	/* 98*/ {"Christmas ", "Christmas", EffectChristmas},
	/* 99*/ {"gradient1", "Gradient 1", EffectGradient1},
	/*100*/ {"Gradient 1", "Gradient 1", EffectGradient1},
	/*101*/ {"gradient 1", "Gradient 1", EffectGradient1},
	/*102*/ {"Gradient2", "Gradient 2", EffectGradient2},
	/*103*/ {"gradient2", "Gradient 2", EffectGradient2},
	/*104*/ {"Gradient 2", "Gradient 2", EffectGradient2},
	/*105*/ {"gradient 2", "Gradient 2", EffectGradient2},
	/*106*/ {"Gradient3", "Gradient 3", EffectGradient3},
	/*107*/ {"gradient3", "Gradient 3", EffectGradient3},
	/*108*/ {"Gradient 3", "Gradient 3", EffectGradient3},
	/*109*/ {"gradient 3", "Gradient 3", EffectGradient3},
	/*110*/ {"Gradient4", "Gradient 4", EffectGradient4},
	/*111*/ {"gradient4", "Gradient 4", EffectGradient4},
	/*112*/ {"Gradient 4", "Gradient 4", EffectGradient4},
	/*113*/ {"gradient 4", "Gradient 4", EffectGradient4},
	/*114*/ {"Gradient5", "Gradient 5", EffectGradient5},
	/*115*/ {"gradient5", "Gradient 5", EffectGradient5},
	/*116*/ {"Gradient 5", "Gradient 5", EffectGradient5},
	/*117*/ {"gradient 5", "Gradient 5", EffectGradient5},
	/*118*/ {"Gradient6", "Gradient 6", EffectGradient6},
	/*119*/ {"gradient6", "Gradient 6", EffectGradient6},
	/*120*/ {"Gradient 6", "Gradient 6", EffectGradient6},
	/*121*/ {"gradient 6", "Gradient 6", EffectGradient6},
	/*122*/ {"Gradient7", "Gradient 7", EffectGradient7},
	/*123*/ {"gradient7", "Gradient 7", EffectGradient7},
	/*124*/ {"Gradient 7", "Gradient 7", EffectGradient7},
	/*125*/ {"gradient 7", "Gradient 7", EffectGradient7},
	/*126*/ {"Gradient8", "Gradient 8", EffectGradient8},
	/*127*/ {"gradient8", "Gradient 8", EffectGradient8},
	/*128*/ {"Gradient 8", "Gradient 8", EffectGradient8},
	/*129*/ {"gradient 8", "Gradient 8", EffectGradient8},
	/*130*/ {"Gradient9", "Gradient 9", EffectGradient9},
	/*131*/ {"gradient9", "Gradient 9", EffectGradient9},
	/*132*/ {"Gradient 9", "Gradient 9", EffectGradient9},
	/*133*/ {"gradient 9", "Gradient 9", EffectGradient9},
	/*134*/ {"Gradient10", "Gradient 10", EffectGradient10},
	/*135*/ {"gradient10", "Gradient 10", EffectGradient10},
	/*136*/ {"Gradient 10", "Gradient 10", EffectGradient10},
	/*137*/ {"gradient 10", "Gradient 10", EffectGradient10},
	/*138*/ {"RGBTest", "RGB Test", EffectRGBTest},
	/*139*/ {"rgbtest", "RGB Test", EffectRGBTest},
	/*140*/ {"RGB Test ", "RGB Test", EffectRGBTest},
	/*141*/ {"rgb test ", "RGB Test", EffectRGBTest},
	/*142*/ {"Alternating", "Alternating", EffectAlternating},
	/*143*/ {"alternating", "Alternating", EffectAlternating},
	/*144*/ {"Unknown", "Unknown", EffectUnknown},
	/*145*/ {"unknown", "Unknown", EffectUnknown},
	/*146*/ {"the macarena", "Unknown", EffectUnknown},
	/*147*/ {"", "Unknown", EffectUnknown},
}

func TestEffectFromString(t *testing.T) {
	for i, test := range effectTests {
		effect := EffectFromString(test.Name)
		if test.Effect != effect {
			t.Errorf("[%d] wanted rgblight effect %v, got %v", i, test.Effect, effect)
		}
	}
}

func TestEffectToName(t *testing.T) {
	for i, test := range effectTests {
		name := test.Effect.Name()
		if test.Name != name {
			t.Errorf("[%d] wanted rgblight effect name %v, got %v", i, test.Name, name)
		}
	}
}

func TestAllEffects(t *testing.T) {
	effects := AllEffects()
	if len(effects) != int(EffectUnknown) {
		t.Errorf("wanted %d rgblight effects, got %d", int(EffectUnknown), len(effects))
	}
	for i, effect := range effects {
		for _, test := range effectTests {
			if test.Effect == effect {
				return // we have at least some test coverage
			}
		}
		t.Errorf("rgblight effect [%d] %v has no test coverage", i, effect)
	}
}
