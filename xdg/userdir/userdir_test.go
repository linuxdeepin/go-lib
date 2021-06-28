/*
 * Copyright (C) 2016 ~ 2018 Deepin Technology Co., Ltd.
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
	require.Nil(t, err)
	assert.Equal(t, value, "/home/test/Desktop")

	value, err = parseValue([]byte(`"/home/test/DesktopA"`), homeDir)
	require.Nil(t, err)
	assert.Equal(t, value, "/home/test/DesktopA")

	value, err = parseValue([]byte(`"$HOME/"`), homeDir)
	require.Nil(t, err)
	assert.Equal(t, value, "/home/test")

	value, err = parseValue([]byte(`"/"`), homeDir)
	require.Nil(t, err)
	assert.Equal(t, value, "/")

	value, err = parseValue([]byte(""), homeDir)
	assert.NotNil(t, err)
	assert.Equal(t, value, "")

	value, err = parseValue([]byte("$HOME"), homeDir)
	assert.NotNil(t, err)
	assert.Equal(t, value, "")

	value, err = parseValue([]byte(`"not abs"`), homeDir)
	assert.NotNil(t, err)
	assert.Equal(t, value, "")
}

func TestParseUserDirsConfig(t *testing.T) {
	os.Setenv("HOME", "/home/test")
	cfg, err := parseUserDirsConfig("./testdata/user-dirs.dirs")
	require.Nil(t, err)
	assert.Equal(t, cfg, map[string]string{"XDG_DESKTOP_DIR": "/home/test/桌面", "XDG_DOCUMENTS_DIR": "/home/test/文档", "XDG_DOWNLOAD_DIR": "/home/test/下载", "XDG_MUSIC_DIR": "/home/test/音乐", "XDG_PICTURES_DIR": "/home/test/图片", "XDG_PUBLICSHARE_DIR": "/home/test/.Public", "XDG_TEMPLATES_DIR": "/home/test/.Templates", "XDG_VIDEOS_DIR": "/home/test/视频"})
}

func TestGet(t *testing.T) {
	os.Setenv("HOME", "/home/test")
	testDataDir, err := filepath.Abs("./testdata")
	require.Nil(t, err)

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
	require.Nil(t, err)

	os.Setenv("XDG_CONFIG_HOME", testDataDir)
	assert.Equal(t, Get(Desktop), "/home/test/桌面")

	testDataDir2, err := filepath.Abs("./testdata2")
	require.Nil(t, err)
	os.Setenv("XDG_CONFIG_HOME", testDataDir2)
	err = ReloadCache()
	require.Nil(t, err)
	assert.Equal(t, Get(Desktop), "/home/test/MyDesktop")
}
