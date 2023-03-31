// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package iso

import (
	"os"
	"testing"

	. "github.com/linuxdeepin/go-lib/gettext"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

const envLanguage = "LANGUAGE"

func TestGetCountryDatabase(t *testing.T) {
	database, err := GetCountryDatabase()
	require.NoError(t, err)
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
	assert.Error(t, err)
	_, err = GetCountryCodeForLanguage("en.US_UTF-8")
	assert.Error(t, err)
	_, err = GetCountryCodeForLanguage("en_.US.UTF-8")
	assert.Error(t, err)
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
