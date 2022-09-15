// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package pinyin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_HansToPinyin(t *testing.T) {
	pinyinMap := make(map[int]string)
	array := HansToPinyin("统信软件")
	for i, a := range array {
		pinyinMap[i] = a
	}
	assert.Len(t, pinyinMap, 2)
	assert.Equal(t, pinyinMap[0], "tongxinruanjian")
	assert.Equal(t, pinyinMap[1], "tongshenruanjian")
}
