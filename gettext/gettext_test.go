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

package gettext

import (
	"os"
	"os/exec"
	"testing"

	C "gopkg.in/check.v1"
)

type gettext struct{} //nolint:golint,unused

func Test(t *testing.T) { C.TestingT(t) }

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

func (*gettext) TestTr(c *C.C) {
	_ = os.Setenv("LC_ALL", "en_US.UTF-8")
	_ = os.Setenv("LANGUAGE", "ar")

	InitI18n()

	Bindtextdomain("test", "testdata/locale")
	Textdomain("test")
	c.Check(Tr("Back"), C.Equals, "الخلف")
}

func (*gettext) TestDGettext(c *C.C) {
	_ = os.Setenv("LC_ALL", "en_US.UTF-8")
	_ = os.Setenv("LANGUAGE", "zh_CN")
	InitI18n()
	Bindtextdomain("test", "testdata/locale")
	c.Check(DGettext("test", "Back"), C.Equals, "返回")
}

func (*gettext) TestFailed(c *C.C) {
	_ = os.Setenv("LC_ALL", "en_US.UTF-8")
	InitI18n()
	Bindtextdomain("test", "testdata/locale")
	c.Check(DGettext("test", "notfound"), C.Equals, "notfound")
	c.Check(DGettext("test", "未找到"), C.Equals, "未找到")
}

func (*gettext) TestNTr(c *C.C) {
	_ = os.Setenv("LC_ALL", "en_US.UTF-8")
	Bindtextdomain("test", "testdata/plural/locale")
	Textdomain("test")

	_ = os.Setenv("LANGUAGE", "es")
	InitI18n()

	c.Check(NTr("%d apple", "%d apples", 1), C.Equals, "%d manzana")
	c.Check(NTr("%d apple", "%d apples", 2), C.Equals, "%d manzanas")

	_ = os.Setenv("LANGUAGE", "zh_CN")
	InitI18n()

	c.Check(NTr("%d apple", "%d apples", 0), C.Equals, "%d苹果")
	c.Check(NTr("%d apple", "%d apples", 1), C.Equals, "%d苹果")
	c.Check(NTr("%d apple", "%d apples", 2), C.Equals, "%d苹果")
}

func (*gettext) TestDNGettext(c *C.C) {
	_ = os.Setenv("LC_ALL", "en_US.UTF-8")
	Bindtextdomain("test", "testdata/plural/locale")

	_ = os.Setenv("LANGUAGE", "es")
	InitI18n()
	c.Check(DNGettext("test", "%d person", "%d persons", 1), C.Equals, "%d persona")
	c.Check(DNGettext("test", "%d person", "%d persons", 2), C.Equals, "%d personas")

	_ = os.Setenv("LANGUAGE", "zh_CN")
	InitI18n()
	c.Check(DNGettext("test", "%d person", "%d persons", 0), C.Equals, "%d人")
	c.Check(DNGettext("test", "%d person", "%d persons", 1), C.Equals, "%d人")
	c.Check(DNGettext("test", "%d person", "%d persons", 2), C.Equals, "%d人")
}

func (*gettext) TestQueryLang(c *C.C) {
	_ = os.Setenv("LC_ALL", "zh_CN.UTF-8")
	_ = os.Setenv("LC_MESSAGE", "zh_TW.")
	_ = os.Setenv("LANGUAGE", "en_US.12")
	_ = os.Setenv("LANG", "it")

	c.Check(QueryLang(), C.Equals, "en_US")

	_ = os.Setenv("LANGUAGE", "")
	c.Check(QueryLang(), C.Equals, "zh_CN")

	_ = os.Setenv("LC_ALL", "")
	c.Check(QueryLang(), C.Equals, "zh_TW")
}
