// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package log

import (
	"fmt"
	"strings"
)

// same with fmt.Sprintln() but trim the additional end line
func fmtSprint(v ...interface{}) (s string) {
	s = fmt.Sprintln(v...)
	s = strings.TrimSuffix(s, "\n")
	return
}

func isStringInArray(s string, arr []string) bool {
	for _, t := range arr {
		if t == s {
			return true
		}
	}
	return false
}
