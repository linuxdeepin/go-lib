// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package shell

import (
	"bytes"
	"strings"
)

const specialChars = "`~!#$&*()|\\;'\"<>? "

func isSpecialChar(c byte) bool {
	return strings.IndexByte(specialChars, c) >= 0
}

// Encode returns a sh string literal representing s
func Encode(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isSpecialChar(c) {
			buf.WriteByte('\\')
			buf.WriteByte(c)
		} else if c == '\t' {
			buf.WriteString(`'\t'`)
		} else if c == '\r' {
			buf.WriteString(`'\r'`)
		} else if c == '\n' {
			buf.WriteString(`'\n'`)
		} else {
			buf.WriteByte(c)
		}
	}
	return buf.String()
}
