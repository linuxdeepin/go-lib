// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package basedir

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserHomeDir(t *testing.T) {
	os.Setenv("HOME", "/home/test")
	dir := GetUserHomeDir()
	assert.Equal(t, dir, "/home/test")
}

func TestGetUserDataDir(t *testing.T) {
	os.Setenv("HOME", "/home/test")
	os.Setenv("XDG_DATA_HOME", "")
	dir := GetUserDataDir()
	assert.Equal(t, dir, "/home/test/.local/share")

	os.Setenv("XDG_DATA_HOME", "a invalid path")
	dir = GetUserDataDir()
	assert.Equal(t, dir, "/home/test/.local/share")

	os.Setenv("XDG_DATA_HOME", "/home/test/xdg")
	dir = GetUserDataDir()
	assert.Equal(t, dir, "/home/test/xdg")
}

func TestFilterNotAbs(t *testing.T) {
	result := filterNotAbs([]string{
		"a/is/invald", "b/is invalid", "c is invald", "/d/is/ok", "/e/is/ok/"})
	assert.Equal(t, result, []string{"/d/is/ok", "/e/is/ok"})
}

func TestGetSystemDataDirs(t *testing.T) {
	os.Setenv("XDG_DATA_DIRS", "/a:/b:/c")
	dirs := GetSystemDataDirs()
	assert.Equal(t, dirs, []string{"/a", "/b", "/c"})

	os.Setenv("XDG_DATA_DIRS", "/a:/b/:/c/")
	dirs = GetSystemDataDirs()
	assert.Equal(t, dirs, []string{"/a", "/b", "/c"})

	os.Setenv("XDG_DATA_DIRS", "/a:/b/:c is invald")
	dirs = GetSystemDataDirs()
	assert.Equal(t, dirs, []string{"/a", "/b"})

	os.Setenv("XDG_DATA_DIRS", "a/is/invald:b/is invalid :c is invald")
	dirs = GetSystemDataDirs()
	assert.Equal(t, dirs, []string{"/usr/local/share", "/usr/share"})

	os.Setenv("XDG_DATA_DIRS", "")
	dirs = GetSystemDataDirs()
	assert.Equal(t, dirs, []string{"/usr/local/share", "/usr/share"})
}

func TestGetUserConfigDir(t *testing.T) {
	os.Setenv("XDG_CONFIG_HOME", "")
	os.Setenv("HOME", "/home/test")
	dir := GetUserConfigDir()
	assert.Equal(t, dir, "/home/test/.config")

	os.Setenv("XDG_CONFIG_HOME", "/home/test/myconfig")
	dir = GetUserConfigDir()
	assert.Equal(t, dir, "/home/test/myconfig")
}

func TestGetSystemConfigDirs(t *testing.T) {
	os.Setenv("XDG_CONFIG_DIRS", "/a:/b:/c")
	dirs := GetSystemConfigDirs()
	assert.Equal(t, dirs, []string{"/a", "/b", "/c"})

	os.Setenv("XDG_CONFIG_DIRS", "")
	dirs = GetSystemConfigDirs()
	assert.Equal(t, dirs, []string{"/etc/xdg"})
}

func TestGetUserCacheDir(t *testing.T) {
	os.Setenv("XDG_CACHE_HOME", "/cache/user/a")
	dir := GetUserCacheDir()
	assert.Equal(t, dir, "/cache/user/a")

	os.Setenv("XDG_CACHE_HOME", "")
	os.Setenv("HOME", "/home/test")
	dir = GetUserCacheDir()
	assert.Equal(t, dir, "/home/test/.cache")
}

func TestGetUserRuntimeDir(t *testing.T) {
	os.Setenv("XDG_RUNTIME_DIR", "/runtime/user/test")
	dir, err := GetUserRuntimeDir(true)
	require.NoError(t, err)
	assert.Equal(t, dir, "/runtime/user/test")

	os.Setenv("XDG_RUNTIME_DIR", "")
	dir, err = GetUserRuntimeDir(true)
	assert.Error(t, err)
	assert.Equal(t, dir, "")

	os.Setenv("XDG_RUNTIME_DIR", "a invalid path")
	dir, err = GetUserRuntimeDir(true)
	assert.Error(t, err)
	assert.Equal(t, dir, "")

	os.Setenv("XDG_RUNTIME_DIR", "")
	dir, err = GetUserRuntimeDir(false)
	require.NoError(t, err)
	assert.Equal(t, dir, fmt.Sprintf("/tmp/goxdg-runtime-dir-fallback-%d", os.Getuid()))
}
