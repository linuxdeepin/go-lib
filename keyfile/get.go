// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package keyfile

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	libLocale "github.com/linuxdeepin/go-lib/locale"
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
				break
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
					buf.WriteByte('\\')
					buf.WriteByte(ch)
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
