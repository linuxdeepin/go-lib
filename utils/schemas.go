// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

import (
	"fmt"

	"github.com/linuxdeepin/go-gir/gio-2.0"
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
