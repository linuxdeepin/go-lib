// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKeyFile(t *testing.T) {
	file := "testdata/test_keyfile.ini"
	group := "Test"
	key := "id"
	value := "1234"

	assert.Equal(t, WriteKeyToKeyFile(file, group, key, value), true)
	ifc, ok := ReadKeyFromKeyFile(file, group, key, "")
	assert.Equal(t, ok, true)

	assert.Equal(t, value, ifc.(string))

	os.Remove(file)
}
