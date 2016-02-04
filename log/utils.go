/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

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
