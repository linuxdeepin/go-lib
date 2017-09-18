/*
 * Copyright (C) 2017 ~ 2017 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package keyfile

import (
	"bytes"
	"fmt"
	libLocale "pkg.deepin.io/lib/locale"
	"strconv"
	"strings"
	"unicode/utf8"
)

func parseValueAsBool(value string) (bool, error) {
	switch value {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, InvalidValueError{value}
	}
}

func (f *KeyFile) GetInt(section, key string) (int, error) {
	value, err := f.GetValue(section, key)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(value)
}

func (f *KeyFile) GetInt64(section, key string) (int64, error) {
	value, err := f.GetValue(section, key)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(value, 10, 64)
}

func (f *KeyFile) GetUint64(section, key string) (uint64, error) {
	value, err := f.GetValue(section, key)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(value, 10, 64)
}

func (f *KeyFile) GetFloat64(section, key string) (float64, error) {
	value, err := f.GetValue(section, key)
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(value, 64)
}

type ParseValueAsStringError struct {
	Value  string
	Reason string
}

func (err ParseValueAsStringError) Error() string {
	return fmt.Sprintf("value %q %s", err.Value, err.Reason)
}

// support escape characters:
// \s space
// \n newline
// \t tab
// \r carriage return
// \\ backslash
func parseValueAsString(value string, wantArray bool, listSeparator byte) (string, []string, error) {
	var buf bytes.Buffer
	var outlist []string
	reader := strings.NewReader(value)
	for {
		ch, err := reader.ReadByte()
		if err != nil {
			break
		}

		if ch == '\\' {
			ch, err := reader.ReadByte()
			if err != nil {
				return "", nil, ParseValueAsStringError{value, "contains escape character at end of line"}
			}

			switch ch {
			case 's':
				buf.WriteByte(' ')
			case 'n':
				buf.WriteByte('\n')
			case 't':
				buf.WriteByte('\t')
			case 'r':
				buf.WriteByte('\r')
			case '\\':
				buf.WriteByte('\\')
			default:
				if wantArray && ch == listSeparator {
					buf.WriteByte(';')
				} else {
					return "", nil, ParseValueAsStringError{value, "contains invalid escape sequence \\" + string(ch)}
				}
			}
		} else if wantArray && ch == listSeparator {
			outlist = append(outlist, buf.String())
			buf.Reset()
		} else {
			buf.WriteByte(ch)
		}
	}

	if wantArray {
		if buf.Len() != 0 {
			outlist = append(outlist, buf.String())
		}
		return "", outlist, nil
	} else {
		return buf.String(), nil, nil
	}
}

func (f *KeyFile) GetString(section, key string) (string, error) {
	value, err := f.GetValue(section, key)
	if err != nil {
		return "", err
	}

	if !utf8.ValidString(value) {
		return "", ValueInvalidUTF8Error{section, key}
	}

	str, _, err := parseValueAsString(value, false, f.ListSeparator)
	return str, err
}

func (f *KeyFile) GetLocaleString(section, key, locale string) (string, error) {
	var languages []string
	if locale == "" {
		languages = libLocale.GetLanguageNames()
	} else {
		languages = libLocale.GetLocaleVariants(locale)
	}

	for _, lang := range languages {
		translated, err := f.GetString(section, fmt.Sprintf("%s[%s]", key, lang))
		if err == nil {
			return translated, nil
		}
	}

	// NOTE: not support key Gettext-Domain
	// fallback to default key
	return f.GetString(section, key)
}

func (f *KeyFile) GetStringList(section, key string) ([]string, error) {
	value, err := f.GetValue(section, key)
	if err != nil {
		return nil, err
	}

	if !utf8.ValidString(value) {
		return nil, ValueInvalidUTF8Error{section, key}
	}

	_, list, err := parseValueAsString(value, true, f.ListSeparator)
	return list, err
}

func (f *KeyFile) GetLocaleStringList(section, key, locale string) ([]string, error) {
	var languages []string
	if locale == "" {
		languages = libLocale.GetLanguageNames()
	} else {
		languages = libLocale.GetLocaleVariants(locale)
	}

	for _, lang := range languages {
		translated, err := f.GetStringList(section, fmt.Sprintf("%s[%s]", key, lang))
		if err == nil {
			return translated, nil
		}
	}

	// fallback to default key
	return f.GetStringList(section, key)
}

type ValueInvalidUTF8Error struct {
	Section string
	Key     string
}

func (err ValueInvalidUTF8Error) Error() string {
	return fmt.Sprintf("value of key %q in section %q is not a valid utf8 string", err.Key, err.Section)
}

func (f *KeyFile) GetBoolList(section, key string) ([]bool, error) {
	value, err := f.GetValue(section, key)
	if err != nil {
		return nil, err
	}

	_, list, err := parseValueAsString(value, true, f.ListSeparator)
	if err != nil {
		return nil, err
	}
	ret := make([]bool, len(list))
	for i, str := range list {
		ret[i], err = parseValueAsBool(str)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func (f *KeyFile) GetIntList(section, key string) ([]int, error) {
	value, err := f.GetValue(section, key)
	if err != nil {
		return nil, err
	}

	_, list, err := parseValueAsString(value, true, f.ListSeparator)
	if err != nil {
		return nil, err
	}
	ret := make([]int, len(list))
	for i, str := range list {
		ret[i], err = strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func (f *KeyFile) GetFloat64List(section, key string) ([]float64, error) {
	value, err := f.GetValue(section, key)
	if err != nil {
		return nil, err
	}

	_, list, err := parseValueAsString(value, true, f.ListSeparator)
	if err != nil {
		return nil, err
	}
	ret := make([]float64, len(list))
	for i, str := range list {
		ret[i], err = strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}
