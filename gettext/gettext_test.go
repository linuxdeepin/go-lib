/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package gettext

import (
	C "launchpad.net/gocheck"
	"os"
	"os/exec"
	"testing"
)

type gettext struct{}

func Test(t *testing.T) { C.TestingT(t) }

func init() {
	// use ./build_test_locale_data to update locale def if need
	os.Setenv("LOCPATH", "testdata/locale_def/")
	os.Setenv("LC_ALL", "en_US.UTF-8")

	cmd := exec.Command("/usr/bin/locale", "-a")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	C.Suite(&gettext{})
}

func (*gettext) TestTr(c *C.C) {
	os.Setenv("LC_ALL", "en_US.UTF-8")
	os.Setenv("LANGUAGE", "ar")

	InitI18n()

	Bindtextdomain("test", "testdata/locale")
	Textdomain("test")
	c.Check(Tr("Back"), C.Equals, "الخلف")
}

func (*gettext) TestDGettext(c *C.C) {
	os.Setenv("LC_ALL", "en_US.UTF-8")
	os.Setenv("LANGUAGE", "zh_CN")
	InitI18n()
	Bindtextdomain("test", "testdata/locale")
	c.Check(DGettext("test", "Back"), C.Equals, "返回")
}

func (*gettext) TestFailed(c *C.C) {
	os.Setenv("LC_ALL", "en_US.UTF-8")
	InitI18n()
	Bindtextdomain("test", "testdata/locale")
	c.Check(DGettext("test", "notfound"), C.Equals, "notfound")
	c.Check(DGettext("test", "未找到"), C.Equals, "未找到")
}

func (*gettext) TestNTr(c *C.C) {
	os.Setenv("LC_ALL", "en_US.UTF-8")
	Bindtextdomain("test", "testdata/plural/locale")
	Textdomain("test")

	os.Setenv("LANGUAGE", "es")
	InitI18n()

	c.Check(NTr("%d apple", "%d apples", 1), C.Equals, "%d manzana")
	c.Check(NTr("%d apple", "%d apples", 2), C.Equals, "%d manzanas")

	os.Setenv("LANGUAGE", "zh_CN")
	InitI18n()

	c.Check(NTr("%d apple", "%d apples", 0), C.Equals, "%d苹果")
	c.Check(NTr("%d apple", "%d apples", 1), C.Equals, "%d苹果")
	c.Check(NTr("%d apple", "%d apples", 2), C.Equals, "%d苹果")
}

func (*gettext) TestDNGettext(c *C.C) {
	os.Setenv("LC_ALL", "en_US.UTF-8")
	Bindtextdomain("test", "testdata/plural/locale")

	os.Setenv("LANGUAGE", "es")
	InitI18n()
	c.Check(DNGettext("test", "%d person", "%d persons", 1), C.Equals, "%d persona")
	c.Check(DNGettext("test", "%d person", "%d persons", 2), C.Equals, "%d personas")

	os.Setenv("LANGUAGE", "zh_CN")
	InitI18n()
	c.Check(DNGettext("test", "%d person", "%d persons", 0), C.Equals, "%d人")
	c.Check(DNGettext("test", "%d person", "%d persons", 1), C.Equals, "%d人")
	c.Check(DNGettext("test", "%d person", "%d persons", 2), C.Equals, "%d人")
}

func (*gettext) TestQueryLang(c *C.C) {
	os.Setenv("LC_ALL", "zh_CN.UTF-8")
	os.Setenv("LC_MESSAGE", "zh_TW.")
	os.Setenv("LANGUAGE", "en_US.12")
	os.Setenv("LANG", "it")

	c.Check(QueryLang(), C.Equals, "en_US")

	os.Setenv("LANGUAGE", "")
	c.Check(QueryLang(), C.Equals, "zh_CN")

	os.Setenv("LC_ALL", "")
	c.Check(QueryLang(), C.Equals, "zh_TW")
}
