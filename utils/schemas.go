/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
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

package utils

import (
	"fmt"

	"pkg.deepin.io/gir/gio-2.0"
)

func CheckAndNewGSettings(schema string) (*gio.Settings, error) {
	if !IsGSchemaExist(schema) {
		return nil, fmt.Errorf("Not found this schema: %s", schema)
	}

	return gio.NewSettings(schema), nil
}

func IsGSchemaExist(schema string) bool {
	if isSchemaInList(schema, gio.SettingsListSchemas()) ||
		isSchemaInList(schema, gio.SettingsListRelocatableSchemas()) {
		return true
	}

	return false
}

func isSchemaInList(schema string, list []string) bool {
	for _, s := range list {
		if schema == s {
			return true
		}
	}

	return false
}
