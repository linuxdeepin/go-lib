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
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExplodeLocale(t *testing.T) {
	cs := ExplodeLocale("zh_CN.UTF-8@hubei")
	assert.Equal(t, cs, &Components{
		Language:  "zh",
		Territory: "CN",
		Codeset:   "UTF-8",
		Modifier:  "hubei",
		Mask:      ComponentTerritory | ComponentCodeset | ComponentModifier,
	})

	cs = ExplodeLocale("zh_CN.UTF-8")
	assert.Equal(t, cs, &Components{
		Language:  "zh",
		Territory: "CN",
		Codeset:   "UTF-8",
		Mask:      ComponentTerritory | ComponentCodeset,
	})

	cs = ExplodeLocale("zh_CN")
	assert.Equal(t, cs, &Components{
		Language:  "zh",
		Territory: "CN",
		Mask:      ComponentTerritory,
	})

	cs = ExplodeLocale("zh")
	assert.Equal(t, cs, &Components{
		Language: "zh",
	})

	cs = ExplodeLocale("")
	assert.Equal(t, cs, &Components{})

	cs = ExplodeLocale("_.@")
	assert.Equal(t, cs, &Components{
		Mask: ComponentTerritory | ComponentCodeset | ComponentModifier,
	})
}

func TestGetLocaleVariants(t *testing.T) {
	variants := GetLocaleVariants("zh_CN")
	assert.Equal(t, variants, []string{"zh_CN", "zh"})

	variants = GetLocaleVariants("zh_CN.UTF-8")
	assert.Equal(t, variants, []string{"zh_CN.UTF-8", "zh_CN", "zh.UTF-8", "zh"})

	variants = GetLocaleVariants("zh_CN.UTF-8@hubei")
	assert.Equal(t, variants, []string{"zh_CN.UTF-8@hubei", "zh_CN@hubei", "zh.UTF-8@hubei", "zh@hubei", "zh_CN.UTF-8", "zh_CN", "zh.UTF-8", "zh"})
}

func Test_readAliases(t *testing.T) {
	aliases := readAliases("testdata/locale.alias")
	assert.Equal(t, aliases, map[string]string{"bokmal": "nb_NO.ISO-8859-1", "catalan": "ca_ES.ISO-8859-1", "croatian": "hr_HR.ISO-8859-2"})

	aliases = readAliases("testdata/x")
	require.Nil(t, aliases)
	assert.Equal(t, aliases["nil"], "")
}

func Test_unaliasLang(t *testing.T) {
	aliasFile = "testdata/locale.alias"
	assert.Equal(t, unaliasLang("zh_CN"), "zh_CN")
	assert.Equal(t, unaliasLang(""), "")
	assert.Equal(t, unaliasLang("bokmal"), "nb_NO.ISO-8859-1")
}

func TestGetLanguageNames(t *testing.T) {
	os.Setenv("LANGUAGE", "zh_CN")
	assert.Equal(t, GetLanguageNames(), []string{"zh_CN", "zh", "C"})

	os.Setenv("LANGUAGE", "en_US")
	assert.Equal(t, GetLanguageNames(), []string{"en_US", "en", "C"})
}
