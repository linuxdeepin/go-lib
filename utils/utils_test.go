package utils

import (
	"io/ioutil"
	C "launchpad.net/gocheck"
	"os"
	"regexp"
	"testing"
)

type testWrapper struct{}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { C.TestingT(t) }

func init() {
	C.Suite(&testWrapper{})
}

func (*testWrapper) TestIsURI(c *C.C) {
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
		c.Check(IsURI(d.value), C.Equals, d.result)
	}
}

func (*testWrapper) TestGetURIScheme(c *C.C) {
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
		c.Check(GetURIScheme(d.value), C.Equals, d.result)
	}
}

func (*testWrapper) TestGetURIContent(c *C.C) {
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
		c.Check(GetURIContent(d.value), C.Equals, d.result)
	}
}

func (*testWrapper) TestEncodeURI(c *C.C) {
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
		c.Check(EncodeURI(d.value, d.scheme), C.Equals, d.result)
	}
}

func (*testWrapper) TestDecodeURI(c *C.C) {
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
		c.Check(DecodeURI(d.value), C.Equals, d.result)
	}
}

func (*testWrapper) TestPathToURI(c *C.C) {
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
		c.Check(PathToURI(d.value, d.scheme), C.Equals, d.result)
	}
}

func (*testWrapper) TestURIToPath(c *C.C) {
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
		c.Check(URIToPath(d.value), C.Equals, d.result)
	}
}

func (*testWrapper) TestIsFileExist(c *C.C) {
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
		c.Check(IsFileExist(d.path), C.Equals, false)
		ioutil.WriteFile(d.path, nil, 0644)
		c.Check(IsFileExist(d.path), C.Equals, true)
		c.Check(IsFileExist(d.uri), C.Equals, true)
		os.Remove(d.path)
	}
}

func (*testWrapper) TestIsDir(c *C.C) {
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
		c.Check(IsDir(info.dir), C.Equals, info.ok)
	}
}

func (*testWrapper) TestIsSymlink(c *C.C) {
	testFile := "/tmp/deepin_test_file"
	testSymlink := "/tmp/deepin_test_symlink"
	os.Remove(testSymlink)
	c.Check(IsSymlink(testSymlink), C.Equals, false)
	ioutil.WriteFile(testFile, nil, 0644)
	os.Symlink(testFile, testSymlink)
	c.Check(IsSymlink(testSymlink), C.Equals, true)
	os.Remove(testSymlink)
	os.Remove(testFile)
}

func (*testWrapper) TestIsEnvExists(c *C.C) {
	testEnvName := "test_is_env_exists"
	testEnvValue := "test_env_value"
	c.Check(false, C.Equals, IsEnvExists(testEnvName))
	os.Setenv(testEnvName, testEnvValue)
	c.Check(true, C.Equals, IsEnvExists(testEnvName))
}

func (*testWrapper) TestUnsetEnv(c *C.C) {
	testEnvName := "test_unset_env"
	testEnvValue := "test_env_value"
	c.Check(false, C.Equals, IsEnvExists(testEnvName))
	os.Setenv(testEnvName, testEnvValue)
	c.Check(true, C.Equals, IsEnvExists(testEnvName))
	envCount := len(os.Environ())
	UnsetEnv(testEnvName)
	c.Check(false, C.Equals, IsEnvExists(testEnvName))
	c.Check(len(os.Environ()), C.Equals, envCount-1)
}

func (*testWrapper) TestGenUuid(c *C.C) {
	validUuid := regexp.MustCompile(`^[a-z0-9]{8}(-[a-z0-9]{4}){3}-[a-z0-9]{12}$`)
	var lastUuid string
	for i := 0; i < 5; i++ {
		currentUuid := GenUuid()
		c.Check(currentUuid != lastUuid, C.Equals, true)
		c.Check(validUuid.MatchString(currentUuid), C.Equals, true)
		lastUuid = currentUuid
	}
}

func (*testWrapper) TestRandString(c *C.C) {
	validStr := regexp.MustCompile(`^[0-za-f]{10}$`)
	var lastStr string
	for i := 0; i < 5; i++ {
		currentStr := RandString(10)
		c.Check(currentStr != lastStr, C.Equals, true)
		c.Check(validStr.MatchString(currentStr), C.Equals, true)
		lastStr = currentStr
	}
}

func (*testWrapper) TestIsInterfaceNil(c *C.C) {
	c.Check(IsInterfaceNil(1), C.Equals, false)
	c.Check(IsInterfaceNil(true), C.Equals, false)
	c.Check(IsInterfaceNil(nil), C.Equals, true)
	var a []int = nil
	c.Check(IsInterfaceNil(a), C.Equals, true)
}
