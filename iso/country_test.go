/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
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

package iso

import (
	"github.com/stretchr/testify/require"
	"os"
	. "pkg.deepin.io/lib/gettext"
	"testing"

	"github.com/stretchr/testify/assert"
)

const envLanguage = "LANGUAGE"

func TestGetCountryDatabase(t *testing.T) {
	database, err := GetCountryDatabase()
	require.Nil(t, err)
	assert.NotNil(t, database)
}

func TestGetLocaleCountryInfo(t *testing.T) {
	cur := os.Getenv("LC_ALL")
	if cur == "C" || cur == "POSIX" {
		t.Skip("Unsupported locale")
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
		assert.Equal(t, code, d.code)
		name, _ := GetLocaleCountryName()
		assert.Equal(t, name, d.name)
	}
}

func TestGetCountryCodeForLanguage(t *testing.T) {
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
		assert.Equal(t, code, d.code)
	}

	// check invalid format
	_, err := GetCountryCodeForLanguage("")
	assert.NotNil(t, err)
	_, err = GetCountryCodeForLanguage("en.US_UTF-8")
	assert.NotNil(t, err)
	_, err = GetCountryCodeForLanguage("en_.US.UTF-8")
	assert.NotNil(t, err)
}

func TestGetCountryInfoForCode(t *testing.T) {
	cur := os.Getenv("LC_ALL")
	if cur == "C" || cur == "POSIX" {
		t.Skip("Unsupported locale")
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
		assert.Equal(t, name, d.name)
	}
}
