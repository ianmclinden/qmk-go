// qmk-go - go client library for VIA-enabled QMK keyboards
// Copyright (c) 2022 Ian McLinden. All rights reserved
//
// This file is released under GNU LGPL 2.1 on Linux,
// and under the 3-clause BSD license on all other platforms

package backlight

import (
	"testing"
)

var effectTests = []struct {
	Byte   byte
	Input  string
	Name   string
	Effect Effect
}{
	/* 0*/ {0, "Off", "Breathing Off", EffectBreathingOff},
	/* 1*/ {0, "off", "Breathing Off", EffectBreathingOff},
	/* 2*/ {0, "Breathing Off ", "Breathing Off", EffectBreathingOff},
	/* 3*/ {0, "breathing Off ", "Breathing Off", EffectBreathingOff},
	/* 4*/ {0, "BreathingOff ", "Breathing Off", EffectBreathingOff},
	/* 5*/ {0, "breathingoff ", "Breathing Off", EffectBreathingOff},
	/* 6*/ {1, "On", "Breathing On", EffectBreathingOn},
	/* 7*/ {1, "on", "Breathing On", EffectBreathingOn},
	/* 8*/ {1, "Breathing ", "Breathing On", EffectBreathingOn},
	/* 9*/ {1, "breathing ", "Breathing On", EffectBreathingOn},
	/*10*/ {1, "Breathing On ", "Breathing On", EffectBreathingOn},
	/*11*/ {1, "breathing on ", "Breathing On", EffectBreathingOn},
	/*12*/ {1, "BreathingOn ", "Breathing On", EffectBreathingOn},
	/*13*/ {1, "breathingon ", "Breathing On", EffectBreathingOn},
	/*14*/ {2, "Unknown", "Unknown", EffectUnknown},
	/*15*/ {2, "unknown", "Unknown", EffectUnknown},
	/*16*/ {2, "the macarena", "Unknown", EffectUnknown},
	/*17*/ {2, "", "Unknown", EffectUnknown},
}

func TestEffectFromString(t *testing.T) {
	for i, test := range effectTests {
		effect := EffectFromString(test.Name)
		if test.Effect != effect {
			t.Errorf("[%d] wanted backlight effect %v, got %v", i, test.Effect, effect)
		}
	}
}

func TestEffectFromByte(t *testing.T) {
	for i, test := range effectTests {
		effect := EffectFromByte(test.Byte)
		if test.Effect != effect {
			t.Errorf("[%d] wanted backlight effect %v, got %v", i, test.Effect, effect)
		}
	}
}

func TestEffectToName(t *testing.T) {
	for i, test := range effectTests {
		name := test.Effect.Name()
		if test.Name != name {
			t.Errorf("[%d] wanted backlight effect name %v, got %v", i, test.Name, name)
		}
	}
}

func TestEffectToByte(t *testing.T) {
	for i, test := range effectTests {
		b := test.Effect.ToByte()
		if test.Byte != b {
			t.Errorf("[%d] wanted backlight effect byte %v, got %v", i, test.Byte, b)
		}
	}
}

func TestAllEffects(t *testing.T) {
	effects := AllEffects()
	if len(effects) != int(EffectUnknown) {
		t.Errorf("wanted %d backlight effects, got %d", int(EffectUnknown), len(effects))
	}
	for i, effect := range effects {
		for _, test := range effectTests {
			if test.Effect == effect {
				return // we have at least some test coverage
			}
		}
		t.Errorf("backlight effect [%d] %v has no test coverage", i, effect)
	}
}
