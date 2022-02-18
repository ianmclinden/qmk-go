// qmk-go - go client library for VIA-enabled QMK keyboards
// Copyright (c) 2022 Ian McLinden. All rights reserved
//
// This file is released under GNU LGPL 2.1 on Linux,
// and under the 3-clause BSD license on all other platforms

package backlight

import "strings"

type Effect uint8

const (
	EffectBreathingOff Effect = iota
	EffectBreathingOn
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
	case EffectBreathingOff:
		return "Breathing Off"
	case EffectBreathingOn:
		return "Breathing On"
	default:
		return "Unknown"
	}
}

func EffectFromString(value string) Effect {
	s := strings.ToLower(value)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "effect", "", -1)
	switch s {
	case "off", "breathingoff":
		return EffectBreathingOff
	case "on", "breathing", "breathingon":
		return EffectBreathingOn
	default:
		return EffectUnknown
	}
}
