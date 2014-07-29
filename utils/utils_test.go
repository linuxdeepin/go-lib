package utils

import (
	"fmt"
	"io/ioutil"
	. "launchpad.net/gocheck"
	"os"
	"testing"
)

type UtilsTest struct{}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func init() {
	Suite(&UtilsTest{})
}

func (*UtilsTest) TestIsURI(c *C) {
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
		c.Check(IsURI(d.value), Equals, d.result)
	}
}

func (*UtilsTest) TestGetURIScheme(c *C) {
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
		c.Check(GetURIScheme(d.value), Equals, d.result)
	}
}

func (*UtilsTest) TestGetURIContent(c *C) {
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
		c.Check(GetURIContent(d.value), Equals, d.result)
	}
}

func (*UtilsTest) TestEncodeURI(c *C) {
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
		c.Check(EncodeURI(d.value, d.scheme), Equals, d.result)
	}
}

func (*UtilsTest) TestDecodeURI(c *C) {
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
		c.Check(DecodeURI(d.value), Equals, d.result)
	}
}

func (*UtilsTest) TestPathToURI(c *C) {
	var data = []struct {
		value, scheme, result string
	}{
		{"", SCHEME_FILE, ""},
		// {"", SCHEME_FILE, "file://"},
		{"/usr/lib/share/test", SCHEME_FILE, "file:///usr/lib/share/test"},
		{"/usr/lib/share/test", SCHEME_FTP, "ftp:///usr/lib/share/test"},
		{"/usr/lib/share/test", SCHEME_HTTP, "http:///usr/lib/share/test"},
		{"/usr/lib/share/test", SCHEME_HTTPS, "https:///usr/lib/share/test"},
		{"/usr/lib/share/test", SCHEME_SMB, "smb:///usr/lib/share/test"},
		// TODO
		{"/usr/lib/share/%E4%B8%AD%E6%96%87%E8%B7%AF%E5%BE%84/1%202%203", SCHEME_FILE, "file:///usr/lib/share/%E4%B8%AD%E6%96%87%E8%B7%AF%E5%BE%84/1%202%203"},
		// {"/usr/lib/share/中文路径/1 2 3", SCHEME_FILE, "file:///usr/lib/share/%E4%B8%AD%E6%96%87%E8%B7%AF%E5%BE%84/1%202%203"},
		// {"/home/fsh/Wallpapers/中文 name with %.jpg", SCHEME_FILE, "file:///home/fsh/Wallpapers/%E4%B8%AD%E6%96%87%20name%20with%20%25.jpg"},
		// {"file:///home/fsh/Wallpapers/%E4%B8%AD%E6%96%87%20name%20with%20%25.jpg", SCHEME_FILE, "file:///home/fsh/Wallpapers/%E4%B8%AD%E6%96%87%20name%20with%20%25.jpg"},
		// {"file:///usr/lib/share/test", SCHEME_FILE, "file:///usr/lib/share/test"},
	}
	for _, d := range data {
		c.Check(PathToURI(d.value, d.scheme), Equals, d.result)
	}
}

func (*UtilsTest) TestURIToPath(c *C) {
	var data = []struct {
		value, result string
	}{
		{"", ""},
		{"file:///usr/lib/share/test", "/usr/lib/share/test"},
		{"ftp:///usr/lib/share/test", "/usr/lib/share/test"},
		{"http:///usr/lib/share/test", "/usr/lib/share/test"},
		{"https:///usr/lib/share/test", "/usr/lib/share/test"},
		{"smb:///usr/lib/share/test", "/usr/lib/share/test"},
		// TODO
		{"file:///usr/lib/share/%E4%B8%AD%E6%96%87%E8%B7%AF%E5%BE%84/1%202%203", "/usr/lib/share/%E4%B8%AD%E6%96%87%E8%B7%AF%E5%BE%84/1%202%203"},
		// {"file:///usr/lib/share/%E4%B8%AD%E6%96%87%E8%B7%AF%E5%BE%84/1%202%203", "/usr/lib/share/中文路径/1 2 3"},
		// {"file:///home/fsh/Wallpapers/%E4%B8%AD%E6%96%87%20name%20with%20%25.jpg", "/home/fsh/Wallpapers/中文 name with %.jpg"},
		{"/usr/lib/share/test", "/usr/lib/share/test"},
	}
	for _, d := range data {
		c.Check(URIToPath(d.value), Equals, d.result)
	}
}

func (*UtilsTest) TestIsFileExist(c *C) {
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
		c.Check(IsFileExist(d.path), Equals, false)
		ioutil.WriteFile(d.path, nil, 0644)
		c.Check(IsFileExist(d.path), Equals, true)
		c.Check(IsFileExist(d.uri), Equals, true)
		os.Remove(d.path)
	}
}

func (*UtilsTest) TestIsDir(c *C) {
	testDir := "/tmp/deepin_test_dir"
	os.RemoveAll(testDir)
	c.Check(IsDir(testDir), Equals, false)
	os.MkdirAll(testDir, 0644)
	c.Check(IsDir(testDir), Equals, true)
	os.RemoveAll(testDir)
}

func (*UtilsTest) TestIsSymlink(c *C) {
	testFile := "/tmp/deepin_test_file"
	testSymlink := "/tmp/deepin_test_symlink"
	os.Remove(testSymlink)
	c.Check(IsSymlink(testSymlink), Equals, false)
	ioutil.WriteFile(testFile, nil, 0644)
	os.Symlink(testFile, testSymlink)
	c.Check(IsSymlink(testSymlink), Equals, true)
	os.Remove(testSymlink)
	os.Remove(testFile)
}

func (*UtilsTest) TestIsEnvExists(c *C) {
	testEnvName := "test_is_env_exists"
	testEnvValue := "test_env_value"
	c.Check(false, Equals, IsEnvExists(testEnvName))
	os.Setenv(testEnvName, testEnvValue)
	c.Check(true, Equals, IsEnvExists(testEnvName))
}

func (*UtilsTest) TestUnsetEnv(c *C) {
	testEnvName := "test_unset_env"
	testEnvValue := "test_env_value"
	c.Check(false, Equals, IsEnvExists(testEnvName))
	os.Setenv(testEnvName, testEnvValue)
	c.Check(true, Equals, IsEnvExists(testEnvName))
	envCount := len(os.Environ())
	UnsetEnv(testEnvName)
	c.Check(false, Equals, IsEnvExists(testEnvName))
	c.Check(len(os.Environ()), Equals, envCount-1)
}

func (*UtilsTest) TestGenUuid(c *C) {
	for i := 0; i < 5; i++ {
		fmt.Println("GenUuid:", GenUuid())
	}
}

func (*UtilsTest) TestRandString(c *C) {
	for i := 0; i < 5; i++ {
		fmt.Println("RandString:", RandString(10))
	}
}

func (*UtilsTest) TestIsInterfaceNil(c *C) {
	c.Check(IsInterfaceNil(1), Equals, false)
	c.Check(IsInterfaceNil(true), Equals, false)
	c.Check(IsInterfaceNil(nil), Equals, true)
	var a []int = nil
	c.Check(IsInterfaceNil(a), Equals, true)
}
