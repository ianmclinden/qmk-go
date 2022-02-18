// qmk-go - go client library for VIA-enabled QMK keyboards
// Copyright (c) 2022 Ian McLinden. All rights reserved
//
// This file is released under GNU LGPL 2.1 on Linux,
// and under the 3-clause BSD license on all other platforms

package qmk

import (
	"errors"
	"sort"

	"github.com/karalabe/hid"
)

var (
	ErrorHIDUnsupported   = errors.New("HID API is not supported")
	ErrorNoKeyboardsFound = errors.New("no via-enabled keyboards found")
)

type Keyboard = hid.DeviceInfo

type byProduct []Keyboard

func (a byProduct) Len() int           { return len(a) }
func (a byProduct) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byProduct) Less(i, j int) bool { return a[i].Product < a[j].Product }

func ListKeyboards() ([]Keyboard, error) {
	var (
		devices   = hid.Enumerate(0, 0)
		keyboards = []Keyboard{}
	)

	for i := range devices {
		if devices[i].UsagePage == HidUsagePage && devices[i].Usage == HidUsage {
			keyboards = append(keyboards, devices[i])
		}
	}

	if len(keyboards) == 0 {
		return nil, ErrorNoKeyboardsFound
	}

	sort.Sort(byProduct(keyboards))
	return keyboards, nil
}
