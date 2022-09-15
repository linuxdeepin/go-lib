// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5Sum(t *testing.T) {
	testStr := "hello world"
	if ret, ok := SumStrMd5(testStr); !ok {
		t.Errorf("SumStrMd5 '%s' Faild", testStr)
		return
	} else {
		assert.Equal(t, ret, "5eb63bbbe01eeed093cb22bb8f5acdc3")
	}

	testFile := "testdata/testfile"
	if ret, ok := SumFileMd5(testFile); !ok {
		t.Errorf("SumFileMd5 '%s' Failed", testFile)
		return
	} else {
		assert.Equal(t, ret, "0a75266cc21da8c88a940b00d4d535b7")
	}
}
