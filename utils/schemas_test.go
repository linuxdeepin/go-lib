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
	"testing"

	"github.com/stretchr/testify/assert"
)

type schemaTest struct {
	schema string
	exist  bool
}

func TestGSchemaIsExist(t *testing.T) {
	schemas := []string{
		"com.deepin.wacom",
		"com.deepin.touchpad",
	}

	list1 := []schemaTest{
		schemaTest{"com.deepin.wacom", true},
		schemaTest{"com.deepin.touchpad", true},
	}

	list2 := []schemaTest{
		schemaTest{"org.123.123", false},
		schemaTest{"org/11/11", false},
		schemaTest{"sdsdsvfdsfs", false},
		schemaTest{"/dsfd/assasd", false},
		schemaTest{".sds.sadsd.", false},
		schemaTest{"-sds-sds-ss", false},
		schemaTest{"(jjjj)", false},
		schemaTest{"$fgg$", false},
	}

	for _, l := range list1 {
		assert.Equal(t, isSchemaInList(l.schema, schemas), l.exist)
	}

	for _, l := range list2 {
		assert.Equal(t, IsGSchemaExist(l.schema), l.exist)
		assert.Equal(t, isSchemaInList(l.schema, schemas), l.exist)
	}
}
