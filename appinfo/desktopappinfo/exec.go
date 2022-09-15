// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package desktopappinfo

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/linuxdeepin/go-lib/utils"
)

func getDefaultTerminal() (exec string, execArg string) {
	gs, err := utils.CheckAndNewGSettings("com.deepin.desktop.default-applications.terminal")
	if err == nil {
		exec = gs.GetString("exec")
		execArg = gs.GetString("exec-arg")
	} else {
		exec = "xterm"
		execArg = "-e"
	}
	return
}

// The characters that must be escaped
var escapeChars = []byte{'`', '$', '\\', '"'}

// Reserved characters
var reservedChars = []byte{'`', '$', '\\', '"', '\t', '\n', '\'', '>', '<', '~', '|', '&', ';', '*', '?', '#', '(', ')'}

func isReservedChar(ch byte) bool {
	for _, v := range reservedChars {
		if v == ch {
			return true
		}
	}
	return false
}

func shouldEscapeChar(ch byte) bool {
	for _, v := range escapeChars {
		if v == ch {
			return true
		}
	}
	return false
}

var ErrQuotingNotClosed = errors.New("quoting is not be closed")
var ErrEscapeCharAtEnd = errors.New("escape char \\ at end of string")
var ErrNoSpaceAfterQuoting = errors.New("no space character found after a quoting")

type ErrReservedCharNotQuoted struct {
	Char byte
}

func (err ErrReservedCharNotQuoted) Error() string {
	return fmt.Sprintf("reserved character %q is not be quoted", err.Char)
}

type ErrCharNotEscaped struct {
	Char byte
}

func (err ErrCharNotEscaped) Error() string {
	return fmt.Sprintf("character %q is not be escaped", err.Char)
}

type ErrInvalidEscapeSequence struct {
	Char byte
}

func (err ErrInvalidEscapeSequence) Error() string {
	return fmt.Sprintf("invalid escape sequence %q", "\\"+string(err.Char))
}

func eatAllSpace(reader *strings.Reader) {
	for {
		ch, err := reader.ReadByte()
		if err != nil {
			// EOF
			return
		}
		if ch != ' ' {
			_ = reader.UnreadByte()
			break
		}
	}
}

func splitExec(exec string) ([]string, error) {
	var buf bytes.Buffer
	var outlist []string
	reader := strings.NewReader(exec)
	var quoteChar byte
	for {
		ch, err := reader.ReadByte()
		if err != nil {
			// err is EOF
			if quoteChar != 0 {
				return nil, ErrQuotingNotClosed
			}
			outlist = append(outlist, buf.String())
			break
		}

		switch ch {
		case ' ':
			if quoteChar != 0 {
				buf.WriteByte(ch)
			} else {
				eatAllSpace(reader)
				outlist = append(outlist, buf.String())
				buf.Reset()
			}
		case '"', '\'':
			if quoteChar == 0 {
				quoteChar = ch
			} else if quoteChar == ch {
				quoteChar = 0
				ch0, err0 := reader.ReadByte()
				if err0 != nil {
					continue
				}
				if ch0 != ' ' {
					return nil, ErrNoSpaceAfterQuoting
				}
				_ = reader.UnreadByte()
			} else {
				buf.WriteByte(ch)
			}

		case '\\':
			if quoteChar != 0 {
				ch1, err1 := reader.ReadByte()
				if err1 != nil {
					// err1 is EOF
					return nil, ErrEscapeCharAtEnd
				}

				if quoteChar == ch1 {
					buf.WriteByte(ch1)
					continue
				}

				if shouldEscapeChar(ch1) {
					// \#
					buf.WriteByte(ch1)
				} else {
					// \b
					return nil, ErrInvalidEscapeSequence{ch1}
				}

			} else {
				ch1, err1 := reader.ReadByte()
				if err1 != nil {
					// \ at end
					return nil, ErrEscapeCharAtEnd
				}

				switch ch1 {
				case ' ':
					// \ + space => space
					buf.WriteByte(' ')
				case '\\':
					// \ + \ => \ + \
					buf.WriteString(`\\`)
				default:
					// \ + w => w
					_ = reader.UnreadByte()
				}
			}

		default:
			if isReservedChar(ch) {
				if quoteChar != 0 {
					buf.WriteByte(ch)
				} else {
					return nil, ErrReservedCharNotQuoted{ch}
				}
			} else {
				buf.WriteByte(ch)
			}
		}
	}

	return outlist, nil
}

func toURL(path string) string {
	var u *url.URL
	if strings.HasPrefix(path, "/") {
		u = &url.URL{
			Path:   path,
			Scheme: "file",
		}
	} else {
		var err error
		u, err = url.Parse(path)
		if err != nil {
			return ""
		}
	}
	return u.String()
}

func toLocalPath(in string) string {
	if strings.HasPrefix(in, "file://") {
		u, err := url.Parse(in)
		if err == nil {
			return u.Path
		}
	}
	return in
}

func (ai *DesktopAppInfo) expandFieldCode(cmdline, files []string) ([]string, error) {
	return expandFieldCode(cmdline, files, ai.GetName(), ai.GetIcon(), ai.GetFileName())
}

var ErrBadFieldCode = errors.New("bad field code")

func expandFieldCode(cmdline, files []string, translatedName, icon, desktopFile string) ([]string, error) {
	// element of files can be local path (starts with /) or uri (starts with file:///)
	var ret []string
	var buf bytes.Buffer
	submitBuf := func() {
		ret = append(ret, buf.String())
		buf.Reset()
	}

	for _, arg := range cmdline {
		argR := strings.NewReader(arg)
		for {
			c, err := argR.ReadByte()
			if err != nil {
				break
			}
			if c == '%' {
				fieldCode, err := argR.ReadByte()
				if err != nil {
					break
				}
				switch fieldCode {
				case 'f':
					// a single filepath
					if len(files) > 0 {
						buf.WriteString(toLocalPath(files[0]))
					}
				case 'F':
					// a list of filepaths
					for _, file := range files {
						buf.WriteString(toLocalPath(file))
						submitBuf()
					}
				case 'u':
					// a single filepath or URL
					// file URI 转为 filepath 传入，避免应用不兼容 file URI 的情况
					if len(files) > 0 {
						buf.WriteString(toLocalPath(files[0]))
					}
				case 'U':
					// a list of filepaths or URLs
					// file URI 转为 filepath 传入，避免应用不兼容 file URI 的情况
					for _, file := range files {
						buf.WriteString(toLocalPath(file))
						submitBuf()
					}
				case 'i':
					// icon
					if icon != "" {
						buf.WriteString("--icon")
						submitBuf()
						buf.WriteString(icon)
					}
				case 'c':
					//  translated name
					buf.WriteString(translatedName)
				case 'k':
					// location of desktop file or URI
					buf.WriteString(desktopFile)
				case '%':
					buf.WriteByte('%')
				case 'd', 'D', 'n', 'N', 'v', 'm':
					// Deprecated
				default:
					return nil, ErrBadFieldCode
				}
			} else {
				buf.WriteByte(c)
			}
		}
		if len(buf.Bytes()) > 0 {
			submitBuf()
		}
	}
	return ret, nil
}
