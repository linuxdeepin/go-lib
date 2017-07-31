/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package iso

import (
	C "gopkg.in/check.v1"
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
	cur := os.Getenv("LC_ALL")
	if cur == "C" || cur == "POSIX" {
		c.Skip("Unsupported locale")
		return
	}
	oldLanguage := os.Getenv(envLanguage)
	defer os.Setenv(envLanguage, oldLanguage)

	testData := []struct {
		language, code, name string
	}{
		{"zh_CN.UTF-8", "CN", "中国"},
		{"en_US.UTF-8", "US", "United States"},
	}
	for _, d := range testData {
		os.Setenv("LC_MESSAGES", "en_US.UTF-8")
		os.Setenv("LC_CTYPE", "en_US.UTF-8")
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
	cur := os.Getenv("LC_ALL")
	if cur == "C" || cur == "POSIX" {
		c.Skip("Unsupported locale")
		return
	}
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
		os.Setenv("LC_MESSAGES", "en_US.UTF-8")
		os.Setenv("LC_CTYPE", "en_US.UTF-8")
		InitI18n()
		os.Setenv(envLanguage, d.language)
		name, _ := GetCountryNameForCode(d.code)
		c.Check(name, C.Equals, d.name)
	}
}
