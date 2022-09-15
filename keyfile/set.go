// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package keyfile

import (
	"bytes"
	"strconv"
)

func stringEscape(in string) string {
	var buf bytes.Buffer
	for _, r := range in {
		switch r {
		case '\n':
			buf.WriteString(`\n`)
		case '\t':
			buf.WriteString(`\t`)
		case '\r':
			buf.WriteString(`\r`)
		case '\\':
			buf.WriteString(`\\`)
		default:
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

func (f *KeyFile) SetString(section, key, value string) {
	f.SetValue(section, key, stringEscape(value))
}

// TODO SetLocaleString

func (f *KeyFile) SetBool(section, key string, value bool) {
	f.SetValue(section, key, strconv.FormatBool(value))
}

func (f *KeyFile) SetInt(section, key string, value int) {
	f.SetValue(section, key, strconv.Itoa(value))
}

func (f *KeyFile) SetInt64(section, key string, value int64) {
	f.SetValue(section, key, strconv.FormatInt(value, 10))
}

func (f *KeyFile) SetUint64(section, key string, value uint64) {
	f.SetValue(section, key, strconv.FormatUint(value, 10))
}

func (f *KeyFile) SetFloat64(section, key string, value float64) {
	f.SetValue(section, key, strconv.FormatFloat(value, 'g', -1, 64))
}

func (f *KeyFile) SetStringList(section, key string, values []string) {
	var buf bytes.Buffer
	for _, val := range values {
		for _, r := range val {
			switch r {
			case ' ':
				buf.WriteString(`\s`)
			case '\n':
				buf.WriteString(`\n`)
			case '\t':
				buf.WriteString(`\t`)
			case '\r':
				buf.WriteString(`\r`)
			case '\\':
				buf.WriteString(`\\`)
			case rune(f.ListSeparator):
				buf.WriteByte('\\')
				buf.WriteRune(r)
			default:
				buf.WriteRune(r)
			}
		}
		buf.WriteByte(f.ListSeparator)
	}
	f.SetValue(section, key, buf.String())
}

func (f *KeyFile) SetBoolList(section, key string, values []bool) {
	var buf bytes.Buffer
	for _, val := range values {
		buf.WriteString(strconv.FormatBool(val))
		buf.WriteByte(f.ListSeparator)
	}
	f.SetValue(section, key, buf.String())
}

func (f *KeyFile) SetIntList(section, key string, values []int) {
	var buf bytes.Buffer
	for _, val := range values {
		buf.WriteString(strconv.Itoa(val))
		buf.WriteByte(f.ListSeparator)
	}
	f.SetValue(section, key, buf.String())
}

func (f *KeyFile) SetFloat64List(section, key string, values []float64) {
	var buf bytes.Buffer
	for _, val := range values {
		buf.WriteString(strconv.FormatFloat(val, 'g', -1, 64))
		buf.WriteByte(f.ListSeparator)
	}
	f.SetValue(section, key, buf.String())
}
