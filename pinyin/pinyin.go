// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package pinyin

import (
	"strconv"
	"strings"
	"unicode"
)

func HansToPinyin(hans string) []string {
	return getPinyinFromKey(hans)
}

func getPinyinFromKey(key string) []string {
	rets := []string{}
	for _, c := range key {
		if unicode.Is(unicode.Scripts["Han"], c) {
			array := getPinyinByHan(int64(c))
			if len(rets) == 0 {
				rets = array
				continue
			}
			rets = rangeArray(rets, array)
		} else {
			array := []string{string(c)}
			if len(rets) == 0 {
				rets = array
			} else {
				rets = rangeArray(rets, array)
			}
		}
	}

	return rets
}

func getPinyinByHan(han int64) []string {
	code := strconv.FormatInt(han, 16)
	value := PinyinDataMap[strings.ToUpper(code)]
	array := strings.Split(value, ";")
	return array
}

func rangeArray(a1, a2 []string) []string {
	rets := []string{}
	for _, v := range a1 {
		for _, r := range a2 {
			rets = append(rets, v+r)
		}
	}

	return rets
}
