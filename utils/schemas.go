/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package utils

import (
	"fmt"
	"gir/gio-2.0"
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
