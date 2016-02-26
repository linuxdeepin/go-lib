/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

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
