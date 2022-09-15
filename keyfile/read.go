// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package keyfile

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func (f *KeyFile) LoadFromReader(reader io.Reader) error {
	var comments string
	var section string
	// Parse line by line
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineLength := len(line)

		switch {
		case lineLength == 0 || line[0] == '#': // Comment
			// Append comments
			if len(comments) == 0 {
				comments = line
			} else {
				comments += (LineBreak + line)
			}
			continue

		case line[0] == '[' && line[lineLength-1] == ']': // New section
			section = strings.TrimSpace(line[1 : lineLength-1])
			if len(section) == 0 {
				return BlankSectionNameError{}
			}
			if len(comments) > 0 {
				f.SetSectionComments(section, comments)
				comments = ""
			}
			continue

		default:
			idx := strings.IndexRune(line, '=')
			if idx == -1 {
				return ParseError{line}
			}
			if section == "" {
				return EntryNotInSectionError{line}
			}
			key := strings.TrimRightFunc(line[:idx], unicode.IsSpace)
			if key == "" {
				return errEmptyKey
			}
			if f.keyReg != nil {
				if !f.keyReg.MatchString(key) {
					return InvalidKeyError{key}
				}
			}
			value := strings.TrimLeftFunc(line[idx+1:], unicode.IsSpace)
			f.SetValue(section, key, value)
			if len(comments) > 0 {
				f.SetKeyComments(section, key, comments)
				comments = ""
			}
		}
	}

	return nil
}

func (f *KeyFile) LoadFromData(data []byte) error {
	return f.LoadFromReader(bytes.NewBuffer(data))
}

func (f *KeyFile) LoadFromFile(filename string) error {
	fh, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fh.Close()
	return f.LoadFromReader(bufio.NewReader(fh))
}

type BlankSectionNameError struct{}

func (err BlankSectionNameError) Error() string {
	return "empty section name not allowed"
}

type EntryNotInSectionError struct {
	Line string
}

func (err EntryNotInSectionError) Error() string {
	return fmt.Sprintf("entry %q not in any section", err.Line)
}

type InvalidKeyError struct {
	Key string
}

func (err InvalidKeyError) Error() string {
	return fmt.Sprintf("invalid key name %q", err.Key)
}

var errEmptyKey = errors.New("key is empty")

type ParseError struct {
	Line string
}

func (err ParseError) Error() string {
	return fmt.Sprintf("could not parse line: %q", err.Line)
}
