// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package arch

import (
	"runtime"
)

type ArchFamilyType int

const (
	Unknown ArchFamilyType = iota
	AMD64
	Sunway
)

func Get() ArchFamilyType {
	switch runtime.GOARCH {
	case "sw_64","sw64":
		return Sunway
	case "amd64":
		return AMD64
	default:
		return Unknown
	}
}
