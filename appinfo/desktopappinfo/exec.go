package desktopappinfo

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"pkg.deepin.io/lib/utils"
	"strings"
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
			reader.UnreadByte()
			break
		}
	}
}

func splitExec(exec string) ([]string, error) {
	var buf bytes.Buffer
	var outlist []string
	reader := strings.NewReader(exec)
	var in bool
	for {
		ch, err := reader.ReadByte()
		if err != nil {
			// err is EOF
			if in {
				return nil, ErrQuotingNotClosed
			}
			outlist = append(outlist, buf.String())
			buf.Reset()
			break
		}

		switch ch {
		case ' ':
			if in {
				buf.WriteByte(ch)
			} else {
				eatAllSpace(reader)
				outlist = append(outlist, buf.String())
				buf.Reset()
			}
		case '"':
			in = !in
			if !in {
				ch0, err0 := reader.ReadByte()
				if err0 != nil {
					continue
				}
				if ch0 != ' ' {
					return nil, ErrNoSpaceAfterQuoting
				}
				reader.UnreadByte()
			}

		case '\\':
			if in {
				ch1, err1 := reader.ReadByte()
				if err1 != nil {
					// err1 is EOF
					return nil, ErrEscapeCharAtEnd
				}

				if shouldEscapeChar(ch1) {
					// \#
					buf.WriteByte(ch1)
				} else {
					// \b
					return nil, ErrInvalidEscapeSequence{ch1}
				}

			} else {
				return nil, ErrReservedCharNotQuoted{ch}
			}

		default:
			if isReservedChar(ch) {
				if in {
					if shouldEscapeChar(ch) {
						return nil, ErrCharNotEscaped{ch}
					}
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

func pathToURI(filepath string) string {
	u := url.URL{Path: filepath}
	u.Scheme = "file"
	return u.String()
}

func (ai *DesktopAppInfo) expandFieldCode(cmdline, files []string) ([]string, error) {
	return expandFieldCode(cmdline, files, ai.GetName(), ai.GetIcon(), ai.GetFileName())
}

var ErrBadFieldCode = errors.New("bad field code")

func expandFieldCode(cmdline, files []string, translatedName, icon, desktopFile string) ([]string, error) {
	var ret []string
	for _, arg := range cmdline {
		if len(arg) == 2 && arg[0] == '%' {
			switch arg[1] {
			case 'f':
				// a single file name
				if len(files) > 0 {
					ret = append(ret, files[0])
				}
			case 'F':
				// a list of files
				ret = append(ret, files...)
			case 'u':
				// a single URL
				if len(files) > 0 {
					ret = append(ret, pathToURI(files[0]))
				}

			case 'U':
				// a list of URLs
				for _, file := range files {
					ret = append(ret, pathToURI(file))
				}

			case 'i':
				// icon
				if icon != "" {
					ret = append(ret, "--icon", icon)
				}

			case 'c':
				//  translated name
				ret = append(ret, translatedName)

			case 'k':
				// location of desktop file or URI
				ret = append(ret, desktopFile)

			case 'd', 'D', 'n', 'N', 'v', 'm':
				// Deprecated

			case '%':
				ret = append(ret, "%")

			default:
				return nil, ErrBadFieldCode
			}
		} else {
			ret = append(ret, arg)
		}
	}
	return ret, nil
}
