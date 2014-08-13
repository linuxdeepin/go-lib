/**
 * Copyright (c) 2013 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 Xu FaSheng
 *
 * Author:      Xu FaSheng <fasheng.xu@gmail.com>
 * Maintainer:  Xu FaSheng <fasheng.xu@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
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
