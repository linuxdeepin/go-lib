/*
 * Copyright (C) 2017 ~ 2018 Deepin Technology Co., Ltd.
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

package locale

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestExplodeLocale(t *testing.T) {
	Convey("ExplodeLocale", t, func(c C) {
		cs := ExplodeLocale("zh_CN.UTF-8@hubei")
		c.So(cs, ShouldResemble, &Components{
			Language:  "zh",
			Territory: "CN",
			Codeset:   "UTF-8",
			Modifier:  "hubei",
			Mask:      ComponentTerritory | ComponentCodeset | ComponentModifier,
		})

		cs = ExplodeLocale("zh_CN.UTF-8")
		c.So(cs, ShouldResemble, &Components{
			Language:  "zh",
			Territory: "CN",
			Codeset:   "UTF-8",
			Mask:      ComponentTerritory | ComponentCodeset,
		})

		cs = ExplodeLocale("zh_CN")
		c.So(cs, ShouldResemble, &Components{
			Language:  "zh",
			Territory: "CN",
			Mask:      ComponentTerritory,
		})

		cs = ExplodeLocale("zh")
		c.So(cs, ShouldResemble, &Components{
			Language: "zh",
		})

		cs = ExplodeLocale("")
		c.So(cs, ShouldResemble, &Components{})

		cs = ExplodeLocale("_.@")
		c.So(cs, ShouldResemble, &Components{
			Mask: ComponentTerritory | ComponentCodeset | ComponentModifier,
		})
	})
}

func TestGetLocaleVariants(t *testing.T) {
	Convey("GetLocaleVariants", t, func(c C) {
		variants := GetLocaleVariants("zh_CN")
		c.So(variants, ShouldResemble, []string{"zh_CN", "zh"})

		variants = GetLocaleVariants("zh_CN.UTF-8")
		c.So(variants, ShouldResemble, []string{"zh_CN.UTF-8", "zh_CN", "zh.UTF-8", "zh"})

		variants = GetLocaleVariants("zh_CN.UTF-8@hubei")
		c.So(variants, ShouldResemble, []string{"zh_CN.UTF-8@hubei", "zh_CN@hubei", "zh.UTF-8@hubei", "zh@hubei", "zh_CN.UTF-8", "zh_CN", "zh.UTF-8", "zh"})
	})
}

func Test_readAliases(t *testing.T) {
	Convey("readAliases", t, func(c C) {
		aliases := readAliases("testdata/locale.alias")
		c.So(aliases, ShouldResemble, map[string]string{"bokmal": "nb_NO.ISO-8859-1", "catalan": "ca_ES.ISO-8859-1", "croatian": "hr_HR.ISO-8859-2"})

		aliases = readAliases("testdata/x")
		c.So(aliases, ShouldBeNil)
		c.So(aliases["nil"], ShouldEqual, "")
	})
}

func Test_unaliasLang(t *testing.T) {
	Convey("unaliasLang", t, func(c C) {
		aliasFile = "testdata/locale.alias"
		c.So(unaliasLang("zh_CN"), ShouldEqual, "zh_CN")
		c.So(unaliasLang(""), ShouldEqual, "")
		c.So(unaliasLang("bokmal"), ShouldEqual, "nb_NO.ISO-8859-1")
	})
}

func TestGetLanguageNames(t *testing.T) {
	Convey("GetLanguageNames", t, func(c C) {
		os.Setenv("LANGUAGE", "zh_CN")
		c.So(GetLanguageNames(), ShouldResemble, []string{"zh_CN", "zh", "C"})

		os.Setenv("LANGUAGE", "en_US")
		c.So(GetLanguageNames(), ShouldResemble, []string{"en_US", "en", "C"})
	})
}
