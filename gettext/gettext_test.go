// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gettext

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	// use ./build_test_locale_data to update locale def if need
	_ = os.Setenv("LOCPATH", "testdata/locale_def/")
	_ = os.Setenv("LC_ALL", "en_US.UTF-8")

	cmd := exec.Command("/usr/bin/locale", "-a")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()

	//C.Suite(&gettext{})
}

func Test_Tr(t *testing.T) {
	_ = os.Setenv("LC_ALL", "en_US.UTF-8")
	_ = os.Setenv("LANGUAGE", "ar")

	InitI18n()

	Bindtextdomain("test", "testdata/locale")
	Textdomain("test")
	assert.Equal(t, Tr("Back"), "الخلف")
}

func Test_DGettext(t *testing.T) {
	_ = os.Setenv("LC_ALL", "en_US.UTF-8")
	_ = os.Setenv("LANGUAGE", "zh_CN")
	InitI18n()
	Bindtextdomain("test", "testdata/locale")
	assert.Equal(t, DGettext("test", "Back"), "返回")
}

func Test_Failed(t *testing.T) {
	_ = os.Setenv("LC_ALL", "en_US.UTF-8")
	InitI18n()
	Bindtextdomain("test", "testdata/locale")
	assert.Equal(t, DGettext("test", "notfound"), "notfound")
	assert.Equal(t, DGettext("test", "未找到"), "未找到")
}

func Test_NTr(t *testing.T) {
	_ = os.Setenv("LC_ALL", "en_US.UTF-8")
	Bindtextdomain("test", "testdata/plural/locale")
	Textdomain("test")

	_ = os.Setenv("LANGUAGE", "es")
	InitI18n()

	assert.Equal(t, NTr("%d apple", "%d apples", 1), "%d manzana")
	assert.Equal(t, NTr("%d apple", "%d apples", 2), "%d manzanas")

	_ = os.Setenv("LANGUAGE", "zh_CN")
	InitI18n()

	assert.Equal(t, NTr("%d apple", "%d apples", 0), "%d苹果")
	assert.Equal(t, NTr("%d apple", "%d apples", 1), "%d苹果")
	assert.Equal(t, NTr("%d apple", "%d apples", 2), "%d苹果")
}

func Test_DNGettext(t *testing.T) {
	_ = os.Setenv("LC_ALL", "en_US.UTF-8")
	Bindtextdomain("test", "testdata/plural/locale")

	_ = os.Setenv("LANGUAGE", "es")
	InitI18n()
	assert.Equal(t, DNGettext("test", "%d person", "%d persons", 1), "%d persona")
	assert.Equal(t, DNGettext("test", "%d person", "%d persons", 2), "%d personas")

	_ = os.Setenv("LANGUAGE", "zh_CN")
	InitI18n()
	assert.Equal(t, DNGettext("test", "%d person", "%d persons", 0), "%d人")
	assert.Equal(t, DNGettext("test", "%d person", "%d persons", 1), "%d人")
	assert.Equal(t, DNGettext("test", "%d person", "%d persons", 2), "%d人")
}

func Test_QueryLang(t *testing.T) {
	_ = os.Setenv("LC_ALL", "zh_CN.UTF-8")
	_ = os.Setenv("LC_MESSAGE", "zh_TW.")
	_ = os.Setenv("LANGUAGE", "en_US.12")
	_ = os.Setenv("LANG", "it")

	assert.Equal(t, QueryLang(), "en_US")

	_ = os.Setenv("LANGUAGE", "")
	assert.Equal(t, QueryLang(), "zh_CN")

	_ = os.Setenv("LC_ALL", "")
	assert.Equal(t, QueryLang(), "zh_TW")
}
