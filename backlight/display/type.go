/*
 * Copyright (C) 2017 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

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
