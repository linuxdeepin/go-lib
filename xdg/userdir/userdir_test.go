// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package userdir

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseValue(t *testing.T) {
	homeDir := "/home/test"
	value, err := parseValue([]byte(`"$HOME/Desktop/"`), homeDir)
	require.NoError(t, err)
	assert.Equal(t, value, "/home/test/Desktop")

	value, err = parseValue([]byte(`"/home/test/DesktopA"`), homeDir)
	require.NoError(t, err)
	assert.Equal(t, value, "/home/test/DesktopA")

	value, err = parseValue([]byte(`"$HOME/"`), homeDir)
	require.NoError(t, err)
	assert.Equal(t, value, "/home/test")

	value, err = parseValue([]byte(`"/"`), homeDir)
	require.NoError(t, err)
	assert.Equal(t, value, "/")

	value, err = parseValue([]byte(""), homeDir)
	assert.Error(t, err)
	assert.Equal(t, value, "")

	value, err = parseValue([]byte("$HOME"), homeDir)
	assert.Error(t, err)
	assert.Equal(t, value, "")

	value, err = parseValue([]byte(`"not abs"`), homeDir)
	assert.Error(t, err)
	assert.Equal(t, value, "")
}

func TestParseUserDirsConfig(t *testing.T) {
	os.Setenv("HOME", "/home/test")
	cfg, err := parseUserDirsConfig("./testdata/user-dirs.dirs")
	require.NoError(t, err)
	assert.Equal(t, cfg, map[string]string{"XDG_DESKTOP_DIR": "/home/test/桌面", "XDG_DOCUMENTS_DIR": "/home/test/文档", "XDG_DOWNLOAD_DIR": "/home/test/下载", "XDG_MUSIC_DIR": "/home/test/音乐", "XDG_PICTURES_DIR": "/home/test/图片", "XDG_PUBLICSHARE_DIR": "/home/test/.Public", "XDG_TEMPLATES_DIR": "/home/test/.Templates", "XDG_VIDEOS_DIR": "/home/test/视频"})
}

func TestGet(t *testing.T) {
	os.Setenv("HOME", "/home/test")
	testDataDir, err := filepath.Abs("./testdata")
	require.NoError(t, err)

	os.Setenv("XDG_CONFIG_HOME", testDataDir)

	assert.Equal(t, Get(Desktop), "/home/test/桌面")
	assert.Equal(t, Get(Download), "/home/test/下载")
	assert.Equal(t, Get(Templates), "/home/test/.Templates")
	assert.Equal(t, Get(PublicShare), "/home/test/.Public")
	assert.Equal(t, Get(Documents), "/home/test/文档")
	assert.Equal(t, Get(Music), "/home/test/音乐")
	assert.Equal(t, Get(Pictures), "/home/test/图片")
	assert.Equal(t, Get(Videos), "/home/test/视频")
	assert.Equal(t, Get("XXXX"), "/home/test")
}

func TestReloadCache(t *testing.T) {
	os.Setenv("HOME", "/home/test")
	testDataDir, err := filepath.Abs("./testdata")
	require.NoError(t, err)

	os.Setenv("XDG_CONFIG_HOME", testDataDir)
	assert.Equal(t, Get(Desktop), "/home/test/桌面")

	testDataDir2, err := filepath.Abs("./testdata2")
	require.NoError(t, err)
	os.Setenv("XDG_CONFIG_HOME", testDataDir2)
	err = ReloadCache()
	require.NoError(t, err)
	assert.Equal(t, Get(Desktop), "/home/test/MyDesktop")
}
