package iso

import (
	C "launchpad.net/gocheck"
	"os"
	. "pkg.deepin.io/lib/gettext"
)

const envLanguage = "LANGUAGE"

func (*testWrapper) TestGetCountryDatabase(c *C.C) {
	database, err := GetCountryDatabase()
	c.Check(database, C.NotNil)
	c.Check(err, C.Equals, nil)
}

func (*testWrapper) TestGetLocaleCountryInfo(c *C.C) {
	oldLanguage := os.Getenv(envLanguage)
	defer os.Setenv(envLanguage, oldLanguage)

	testData := []struct {
		language, code, name string
	}{
		{"zh_CN.UTF-8", "CN", "中国"},
		{"en_US.UTF-8", "US", "United States"},
	}
	for _, d := range testData {
		InitI18n()
		os.Setenv(envLanguage, d.language)
		code, _ := GetLocaleCountryCode()
		c.Check(code, C.Equals, d.code)
		name, _ := GetLocaleCountryName()
		c.Check(name, C.Equals, d.name)
	}
}

func (*testWrapper) TestGetCountryCodeForLanguage(c *C.C) {
	testData := []struct {
		language, code string
	}{
		{"de_DE.ISO-8859-1", "DE"},
		{"en_US.UTF-8", "US"},
		{"zh_CN.UTF-8", "CN"},
		{"zh_CN", "CN"},
	}
	for _, d := range testData {
		code, _ := GetCountryCodeForLanguage(d.language)
		c.Check(code, C.Equals, d.code)
	}

	// check invalid format
	_, err := GetCountryCodeForLanguage("")
	c.Check(err, C.NotNil)
	_, err = GetCountryCodeForLanguage("en.US_UTF-8")
	c.Check(err, C.NotNil)
	_, err = GetCountryCodeForLanguage("en_.US.UTF-8")
	c.Check(err, C.NotNil)
}

func (*testWrapper) TestGetCountryInfoForCode(c *C.C) {
	oldLanguage := os.Getenv(envLanguage)
	defer os.Setenv(envLanguage, oldLanguage)

	testData := []struct {
		language, code, name string
	}{
		{"zh_CN.UTF-8", "CN", "中国"},
		{"zh_CN.UTF-8", "US", "美国"},
		{"en_US.UTF-8", "CN", "China"},
		{"en_US.UTF-8", "US", "United States"},
		{"en_US.UTF-8", "Cn", "China"},
		{"en_US.UTF-8", "cn", "China"},
	}
	for _, d := range testData {
		InitI18n()
		os.Setenv(envLanguage, d.language)
		name, _ := GetCountryNameForCode(d.code)
		c.Check(name, C.Equals, d.name)
	}
}
