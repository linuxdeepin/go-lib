// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package userdir

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"github.com/linuxdeepin/go-lib/xdg/basedir"
	"sync"
	"unicode"
)

const (
	Desktop     = "DESKTOP"
	Download    = "DOWNLOAD"
	Templates   = "TEMPLATES"
	PublicShare = "PUBLICSHARE"
	Documents   = "DOCUMENTS"
	Music       = "MUSIC"
	Pictures    = "PICTURES"
	Videos      = "VIDEOS"
)

func parseUserDirsConfig(file string) (map[string]string, error) {
	result := make(map[string]string)
	homeDir := basedir.GetUserHomeDir()

	fh, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	reader := bufio.NewReader(fh)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		// remove newline at end
		line = bytes.TrimRightFunc(line, unicode.IsSpace)

		// skip comments
		if !bytes.HasPrefix(line, []byte("XDG_")) {
			continue
		}

		parts := bytes.SplitN(line, []byte{'='}, 2)
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		if !bytes.HasSuffix(key, []byte("_DIR")) {
			continue
		}
		// key match regexp /XDG_.*_DIR/
		value, err := parseValue(parts[1], homeDir)
		if err != nil {
			continue
		}

		result[string(key)] = value
	}

	return result, nil
}

var errBadValueFormat = errors.New("bad value format")

// Foramt of value is "$HOME/yyy" , where yyy is a shell-escaped
// homedir-relative path, or "/yyy", where /yyy is an absolute path.
// No other format is supported.
func parseValue(val []byte, homeDir string) (string, error) {
	if len(val) < 3 {
		// min "/"
		return "", errBadValueFormat
	}

	// The first character and the last character must be equal to a
	if val[0] != '"' || val[len(val)-1] != '"' {
		return "", errBadValueFormat
	}
	val = val[1 : len(val)-1]
	var isRelative bool
	if bytes.HasPrefix(val, []byte("$HOME")) {
		isRelative = true
		val = val[5:]
	}
	if len(val) == 0 || val[0] != '/' {
		return "", errBadValueFormat
	}
	value := string(val)
	if isRelative {
		value = filepath.Join(homeDir, value)
	}
	return filepath.Clean(value), nil
}

var userDirsCache map[string]string
var mutex sync.Mutex

func Get(dir string) string {
	mutex.Lock()
	defer mutex.Unlock()

	if userDirsCache == nil {
		cfg, err := parseUserDirsConfig(getUserDirsConfigFile())
		if err != nil {
			return basedir.GetUserHomeDir()
		}
		userDirsCache = cfg
	}

	if dir, ok := userDirsCache["XDG_"+dir+"_DIR"]; ok {
		return dir
	} else {
		return basedir.GetUserHomeDir()
	}
}

func ReloadCache() error {
	mutex.Lock()
	defer mutex.Unlock()

	cfg, err := parseUserDirsConfig(getUserDirsConfigFile())
	if err != nil {
		return err
	}
	userDirsCache = cfg
	return nil
}

func getUserDirsConfigFile() string {
	// default ~/.config/user-dirs.dirs
	return filepath.Join(basedir.GetUserConfigDir(), "user-dirs.dirs")
}
