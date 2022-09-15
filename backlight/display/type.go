// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package display

type ControllerType uint

const (
	ControllerTypeUnknown ControllerType = iota
	ControllerTypeFirmware
	ControllerTypePlatform
	ControllerTypeRaw
)

func (t ControllerType) String() string {
	switch t {
	case ControllerTypeFirmware:
		return "firmware"
	case ControllerTypePlatform:
		return "platform"
	case ControllerTypeRaw:
		return "raw"
	}
	return "unknown"
}

func ControllerTypeFromString(str string) ControllerType {
	switch str {
	case "firmware":
		return ControllerTypeFirmware
	case "platform":
		return ControllerTypePlatform
	case "raw":
		return ControllerTypeRaw
	}
	return ControllerTypeUnknown
}
