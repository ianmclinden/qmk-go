// qmk-go - go client library for VIA-enabled QMK keyboards
// Copyright (c) 2022 Ian McLinden. All rights reserved
//
// This file is released under GNU LGPL 2.1 on Linux,
// and under the 3-clause BSD license on all other platforms

package rgblight

import (
	"math"
	"strings"
)

type Speed uint8 // 0-100%

func SpeedFromByte(value byte) Speed {
	return Speed(uint8(math.Round((float64(value) * 100.0) / 255.0)))
}

func (h Speed) ToByte() byte {
	return byte(math.Round((float64(h) * 255.0) / 100.0))
}

type Effect uint8

const (
	EffectAllOff Effect = iota
	EffectSolidColor
	EffectBreathing1
	EffectBreathing2
	EffectBreathing3
	EffectBreathing4
	EffectRainbowMood1
	EffectRainbowMood2
	EffectRainbowMood3
	EffectRainbowSwirl1
	EffectRainbowSwirl2
	EffectRainbowSwirl3
	EffectRainbowSwirl4
	EffectRainbowSwirl5
	EffectRainbowSwirl6
	EffectSnake1
	EffectSnake2
	EffectSnake3
	EffectSnake4
	EffectSnake5
	EffectSnake6
	EffectKnight1
	EffectKnight2
	EffectKnight3
	EffectChristmas
	EffectGradient1
	EffectGradient2
	EffectGradient3
	EffectGradient4
	EffectGradient5
	EffectGradient6
	EffectGradient7
	EffectGradient8
	EffectGradient9
	EffectGradient10
	EffectRGBTest
	EffectAlternating
	EffectUnknown
)

func EffectFromByte(value byte) Effect {
	return Effect(value)
}

func (e Effect) ToByte() byte {
	return byte(e)
}

func AllEffects() []Effect {
	es := make([]Effect, EffectUnknown)
	for i := 0; i < int(EffectUnknown); i++ {
		es[i] = Effect(i)
	}
	return es
}

func (e Effect) Name() string {
	switch e {
	case EffectAllOff:
		return "All Off"
	case EffectSolidColor:
		return "Solid Color"
	case EffectBreathing1:
		return "Breathing 1"
	case EffectBreathing2:
		return "Breathing 2"
	case EffectBreathing3:
		return "Breathing 3"
	case EffectBreathing4:
		return "Breathing 4"
	case EffectRainbowMood1:
		return "Rainbow Mood 1"
	case EffectRainbowMood2:
		return "Rainbow Mood 2"
	case EffectRainbowMood3:
		return "Rainbow Mood 3"
	case EffectRainbowSwirl1:
		return "Rainbow Swirl 1"
	case EffectRainbowSwirl2:
		return "Rainbow Swirl 2"
	case EffectRainbowSwirl3:
		return "Rainbow Swirl 3"
	case EffectRainbowSwirl4:
		return "Rainbow Swirl 4"
	case EffectRainbowSwirl5:
		return "Rainbow Swirl 5"
	case EffectRainbowSwirl6:
		return "Rainbow Swirl 6"
	case EffectSnake1:
		return "Snake 1"
	case EffectSnake2:
		return "Snake 2"
	case EffectSnake3:
		return "Snake 3"
	case EffectSnake4:
		return "Snake 4"
	case EffectSnake5:
		return "Snake 5"
	case EffectSnake6:
		return "Snake 6"
	case EffectKnight1:
		return "Knight 1"
	case EffectKnight2:
		return "Knight 2"
	case EffectKnight3:
		return "Knight 3"
	case EffectChristmas:
		return "Christmas"
	case EffectGradient1:
		return "Gradient 1"
	case EffectGradient2:
		return "Gradient 2"
	case EffectGradient3:
		return "Gradient 3"
	case EffectGradient4:
		return "Gradient 4"
	case EffectGradient5:
		return "Gradient 5"
	case EffectGradient6:
		return "Gradient 6"
	case EffectGradient7:
		return "Gradient 7"
	case EffectGradient8:
		return "Gradient 8"
	case EffectGradient9:
		return "Gradient 9"
	case EffectGradient10:
		return "Gradient 10"
	case EffectRGBTest:
		return "RGB Test"
	case EffectAlternating:
		return "Alternating"
	default:
		return "Unknown"
	}
}

func EffectFromString(value string) Effect {
	s := strings.ToLower(value)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "effect", "", -1)
	switch s {
	case "off", "alloff":
		return EffectAllOff
	case "solid", "static", "color", "solidcolor", "staticcolor":
		return EffectSolidColor
	case "breathing", "breathing1":
		return EffectBreathing1
	case "breathing2":
		return EffectBreathing2
	case "breathing3":
		return EffectBreathing3
	case "breathing4":
		return EffectBreathing4
	case "mood", "mood1", "rainbowmood", "rainbowmood1":
		return EffectRainbowMood1
	case "mood2", "rainbowmood2":
		return EffectRainbowMood2
	case "mood3", "rainbowmood3":
		return EffectRainbowMood3
	case "swirl", "swirl1", "rainbowswirl", "rainbowswirl1":
		return EffectRainbowSwirl1
	case "swirl2", "rainbowswirl2":
		return EffectRainbowSwirl2
	case "swirl3", "rainbowswirl3":
		return EffectRainbowSwirl3
	case "swirl4", "rainbowswirl4":
		return EffectRainbowSwirl4
	case "swirl5", "rainbowswirl5":
		return EffectRainbowSwirl5
	case "swirl6", "rainbowswirl6":
		return EffectRainbowSwirl6
	case "snake", "snake1":
		return EffectSnake1
	case "snake2":
		return EffectSnake2
	case "snake3":
		return EffectSnake3
	case "snake4":
		return EffectSnake4
	case "snake5":
		return EffectSnake5
	case "snake6":
		return EffectSnake6
	case "knight", "knight1":
		return EffectKnight1
	case "knight2":
		return EffectKnight2
	case "knight3":
		return EffectKnight3
	case "christmas":
		return EffectChristmas
	case "gradient", "gradient1":
		return EffectGradient1
	case "gradient2":
		return EffectGradient2
	case "gradient3":
		return EffectGradient3
	case "gradient4":
		return EffectGradient4
	case "gradient5":
		return EffectGradient5
	case "gradient6":
		return EffectGradient6
	case "gradient7":
		return EffectGradient7
	case "gradient8":
		return EffectGradient8
	case "gradient9":
		return EffectGradient9
	case "gradient10":
		return EffectGradient10
	case "test", "rgbtest":
		return EffectRGBTest
	case "alternating":
		return EffectAlternating
	}
	return EffectUnknown
}
