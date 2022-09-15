// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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
