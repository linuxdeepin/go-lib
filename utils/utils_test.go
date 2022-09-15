// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func TestIsURI(t *testing.T) {
	var data = []struct {
		value  string
		result bool
	}{
		{"", false},
		{":", false},
		{"://", false},
		{"file:/", false},
		{"file://", true},
		{"file:///", true},
		{"file:///usr/share", true},
		{"unknown:///usr/share", true},
	}
	for _, d := range data {
		assert.Equal(t, IsURI(d.value), d.result)
	}
}

func TestGetURIScheme(t *testing.T) {
	var data = []struct {
		value, result string
	}{
		{"", ""},
		{":", ""},
		{"://", ""},
		{"file:/", ""},
		{"file://", "file"},
		{"file:///", "file"},
		{"file:///usr/share", "file"},
		{"unknown:///usr/share", "unknown"},
	}
	for _, d := range data {
		assert.Equal(t, GetURIScheme(d.value), d.result)
	}
}

func TestGetURIContent(t *testing.T) {
	var data = []struct {
		value, result string
	}{
		{"", ""},
		{":", ""},
		{"://", ""},
		{"file:/", ""},
		{"file://", ""},
		{"file:///", "/"},
		{"file:///usr/share", "/usr/share"},
		{"unknown:///usr/share", "/usr/share"},
	}
	for _, d := range data {
		assert.Equal(t, GetURIContent(d.value), d.result)
	}
}

func TestEncodeURI(t *testing.T) {
	var data = []struct {
		value, scheme, result string
	}{
		{"", SCHEME_FILE, "file://"},
		{"/usr/lib/share/test", SCHEME_FILE, "file:///usr/lib/share/test"},
		{"/usr/lib/share/test", SCHEME_FTP, "ftp:///usr/lib/share/test"},
		{"/usr/lib/share/test", SCHEME_HTTP, "http:///usr/lib/share/test"},
		{"/usr/lib/share/test", SCHEME_HTTPS, "https:///usr/lib/share/test"},
		{"/usr/lib/share/test", SCHEME_SMB, "smb:///usr/lib/share/test"},
		{"/usr/lib/share/中文路径/1 2 3", SCHEME_FILE, "file:///usr/lib/share/%E4%B8%AD%E6%96%87%E8%B7%AF%E5%BE%84/1%202%203"},
		{"file:///usr/lib/share/test", SCHEME_FILE, "file:///usr/lib/share/test"},
		{"file:///usr/lib/share/test", SCHEME_FTP, "ftp:///usr/lib/share/test"},
		{"file:///usr/lib/share/%E4%B8%AD%E6%96%87%E8%B7%AF%E5%BE%84/1%202%203", SCHEME_FILE, "file:///usr/lib/share/%E4%B8%AD%E6%96%87%E8%B7%AF%E5%BE%84/1%202%203"},
		{"file:///usr/lib/share/中文路径/1 2 3", SCHEME_FILE, "file:///usr/lib/share/%E4%B8%AD%E6%96%87%E8%B7%AF%E5%BE%84/1%202%203"},
		{"/usr/lib/share/中文路径/1 2 3", SCHEME_FILE, "file:///usr/lib/share/%E4%B8%AD%E6%96%87%E8%B7%AF%E5%BE%84/1%202%203"},
		{"/home/fsh/Wallpapers/中文 name with %.jpg", SCHEME_FILE, "file:///home/fsh/Wallpapers/%E4%B8%AD%E6%96%87%20name%20with%20%25.jpg"},
		{"file:///home/fsh/Wallpapers/%E4%B8%AD%E6%96%87%20name%20with%20%25.jpg", SCHEME_FILE, "file:///home/fsh/Wallpapers/%E4%B8%AD%E6%96%87%20name%20with%20%25.jpg"},
		{"file:///usr/lib/share/test", SCHEME_FILE, "file:///usr/lib/share/test"},
	}
	for _, d := range data {
		assert.Equal(t, EncodeURI(d.value, d.scheme), d.result)
	}
}

func TestDecodeURI(t *testing.T) {
	var data = []struct {
		value, result string
	}{
		{"", ""},
		{"file:///usr/lib/share/test", "/usr/lib/share/test"},
		{"ftp:///usr/lib/share/test", "/usr/lib/share/test"},
		{"http:///usr/lib/share/test", "/usr/lib/share/test"},
		{"https:///usr/lib/share/test", "/usr/lib/share/test"},
		{"smb:///usr/lib/share/test", "/usr/lib/share/test"},
		{"file:///usr/lib/share/%E4%B8%AD%E6%96%87%E8%B7%AF%E5%BE%84/1%202%203", "/usr/lib/share/中文路径/1 2 3"},
		{"file:///usr/lib/share/中文路径/1 2 3", "/usr/lib/share/中文路径/1 2 3"},
		{"file:///home/fsh/Wallpapers/%E4%B8%AD%E6%96%87%20name%20with%20%25.jpg", "/home/fsh/Wallpapers/中文 name with %.jpg"},
		{"/usr/lib/share/test", "/usr/lib/share/test"},
	}
	for _, d := range data {
		assert.Equal(t, DecodeURI(d.value), d.result)
	}
}

func TestPathToURI(t *testing.T) {
	var data = []struct {
		value, scheme, result string
	}{
		{"", SCHEME_FILE, ""},
		{"/usr/lib/share/test", SCHEME_FILE, "file:///usr/lib/share/test"},
		{"/usr/lib/share/test", SCHEME_FTP, "ftp:///usr/lib/share/test"},
		{"/usr/lib/share/test", SCHEME_HTTP, "http:///usr/lib/share/test"},
		{"/usr/lib/share/test", SCHEME_HTTPS, "https:///usr/lib/share/test"},
		{"/usr/lib/share/test", SCHEME_SMB, "smb:///usr/lib/share/test"},
		{"/usr/lib/share/%E4%B8%AD%E6%96%87%E8%B7%AF%E5%BE%84/1%202%203", SCHEME_FILE, "file:///usr/lib/share/%E4%B8%AD%E6%96%87%E8%B7%AF%E5%BE%84/1%202%203"},
	}
	for _, d := range data {
		assert.Equal(t, PathToURI(d.value, d.scheme), d.result)
	}
}

func TestURIToPath(t *testing.T) {
	var data = []struct {
		value, result string
	}{
		{"", ""},
		{"file:///usr/lib/share/test", "/usr/lib/share/test"},
		{"ftp:///usr/lib/share/test", "/usr/lib/share/test"},
		{"http:///usr/lib/share/test", "/usr/lib/share/test"},
		{"https:///usr/lib/share/test", "/usr/lib/share/test"},
		{"smb:///usr/lib/share/test", "/usr/lib/share/test"},
		{"file:///usr/lib/share/%E4%B8%AD%E6%96%87%E8%B7%AF%E5%BE%84/1%202%203", "/usr/lib/share/%E4%B8%AD%E6%96%87%E8%B7%AF%E5%BE%84/1%202%203"},
		{"/usr/lib/share/test", "/usr/lib/share/test"},
	}
	for _, d := range data {
		assert.Equal(t, URIToPath(d.value), d.result)
	}
}

func TestIsFileExist(t *testing.T) {
	var data = []struct {
		path, uri string
	}{
		{"/tmp/deepin_test_file", "file:///tmp/deepin_test_file"},
		{"/tmp/deepin_test_中文 name with %.jpg", "file:///tmp/deepin_test_%E4%B8%AD%E6%96%87%20name%20with%20%25.jpg"},
		{"/tmp/deepin_test_file 中文路径", "file:///tmp/deepin_test_file%20%E4%B8%AD%E6%96%87%E8%B7%AF%E5%BE%84"},
		{"/tmp/deepin_test_file 中文路径", "file:///tmp/deepin_test_file 中文路径"},
	}
	for _, d := range data {
		os.Remove(d.path)
		assert.Equal(t, IsFileExist(d.path), false)
		_ = ioutil.WriteFile(d.path, nil, 0644)
		assert.Equal(t, IsFileExist(d.path), true)
		assert.Equal(t, IsFileExist(d.uri), true)
		os.Remove(d.path)
	}
}

func TestIsDir(t *testing.T) {
	var infos = []struct {
		dir string
		ok  bool
	}{
		{
			dir: "testdata",
			ok:  true,
		},
		// test symlink dir
		//{
		//dir: "testdata/utils",
		//ok:  true,
		//},
		{
			dir: "testdata/test1",
			ok:  false,
		},
	}

	for _, info := range infos {
		assert.Equal(t, IsDir(info.dir), info.ok)
	}
}

func TestIsSymlink(t *testing.T) {
	testFile := "/tmp/deepin_test_file"
	testSymlink := "/tmp/deepin_test_symlink"
	os.Remove(testSymlink)
	assert.Equal(t, IsSymlink(testSymlink), false)
	_ = ioutil.WriteFile(testFile, nil, 0644)
	_ = os.Symlink(testFile, testSymlink)
	assert.Equal(t, IsSymlink(testSymlink), true)
	os.Remove(testSymlink)
	os.Remove(testFile)
}

func TestIsEnvExists(t *testing.T) {
	testEnvName := "test_is_env_exists"
	testEnvValue := "test_env_value"
	assert.Equal(t, false, IsEnvExists(testEnvName))
	os.Setenv(testEnvName, testEnvValue)
	assert.Equal(t, true, IsEnvExists(testEnvName))
}

func TestUnsetEnv(t *testing.T) {
	testEnvName := "test_unset_env"
	testEnvValue := "test_env_value"
	assert.Equal(t, false, IsEnvExists(testEnvName))
	os.Setenv(testEnvName, testEnvValue)
	assert.Equal(t, true, IsEnvExists(testEnvName))
	envCount := len(os.Environ())
	_ = UnsetEnv(testEnvName)
	assert.Equal(t, false, IsEnvExists(testEnvName))
	assert.Equal(t, len(os.Environ()), envCount-1)
}

func TestGenUuid(t *testing.T) {
	validUuid := regexp.MustCompile(`^[a-z0-9]{8}(-[a-z0-9]{4}){3}-[a-z0-9]{12}$`)
	var lastUuid string
	for i := 0; i < 5; i++ {
		currentUuid := GenUuid()
		assert.Equal(t, currentUuid != lastUuid, true)
		assert.Equal(t, validUuid.MatchString(currentUuid), true)
		lastUuid = currentUuid
	}
}

func TestRandString(t *testing.T) {
	validStr := regexp.MustCompile(`^[0-za-f]{10}$`)
	var lastStr string
	for i := 0; i < 5; i++ {
		currentStr := RandString(10)
		assert.Equal(t, currentStr != lastStr, true)
		assert.Equal(t, validStr.MatchString(currentStr), true)
		lastStr = currentStr
	}
}

func TestIsInterfaceNil(t *testing.T) {
	assert.Equal(t, IsInterfaceNil(nil), true)
	assert.Equal(t, IsInterfaceNil(1), false)
	assert.Equal(t, IsInterfaceNil(true), false)
	assert.Equal(t, IsInterfaceNil(""), false)
	assert.Equal(t, IsInterfaceNil("str"), false)

	var a1 []int = nil
	assert.Equal(t, IsInterfaceNil(a1), true)
	var a2 []int = make([]int, 0)
	assert.Equal(t, IsInterfaceNil(a2), false)

	var ifcWrapper interface{} = "str"
	assert.Equal(t, IsInterfaceNil(ifcWrapper), false)
}
