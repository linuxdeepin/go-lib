package locale

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestExplodeLocale(t *testing.T) {
	Convey("ExplodeLocale", t, func() {
		cs := ExplodeLocale("zh_CN.UTF-8@hubei")
		So(cs, ShouldResemble, &Components{
			Language:  "zh",
			Territory: "CN",
			Codeset:   "UTF-8",
			Modifier:  "hubei",
			Mask:      ComponentTerritory | ComponentCodeset | ComponentModifier,
		})

		cs = ExplodeLocale("zh_CN.UTF-8")
		So(cs, ShouldResemble, &Components{
			Language:  "zh",
			Territory: "CN",
			Codeset:   "UTF-8",
			Mask:      ComponentTerritory | ComponentCodeset,
		})

		cs = ExplodeLocale("zh_CN")
		So(cs, ShouldResemble, &Components{
			Language:  "zh",
			Territory: "CN",
			Mask:      ComponentTerritory,
		})

		cs = ExplodeLocale("zh")
		So(cs, ShouldResemble, &Components{
			Language: "zh",
		})

		cs = ExplodeLocale("")
		So(cs, ShouldResemble, &Components{})

		cs = ExplodeLocale("_.@")
		So(cs, ShouldResemble, &Components{
			Mask: ComponentTerritory | ComponentCodeset | ComponentModifier,
		})
	})
}

func TestGetLocaleVariants(t *testing.T) {
	Convey("GetLocaleVariants", t, func() {
		variants := GetLocaleVariants("zh_CN")
		So(variants, ShouldResemble, []string{"zh_CN", "zh"})

		variants = GetLocaleVariants("zh_CN.UTF-8")
		So(variants, ShouldResemble, []string{"zh_CN.UTF-8", "zh_CN", "zh.UTF-8", "zh"})

		variants = GetLocaleVariants("zh_CN.UTF-8@hubei")
		So(variants, ShouldResemble, []string{"zh_CN.UTF-8@hubei", "zh_CN@hubei", "zh.UTF-8@hubei", "zh@hubei", "zh_CN.UTF-8", "zh_CN", "zh.UTF-8", "zh"})
	})
}

func Test_readAliases(t *testing.T) {
	Convey("readAliases", t, func() {
		aliases := readAliases("testdata/locale.alias")
		So(aliases, ShouldResemble, map[string]string{"bokmal": "nb_NO.ISO-8859-1", "catalan": "ca_ES.ISO-8859-1", "croatian": "hr_HR.ISO-8859-2"})

		aliases = readAliases("testdata/x")
		So(aliases, ShouldBeNil)
		So(aliases["nil"], ShouldEqual, "")
	})
}

func Test_unaliasLang(t *testing.T) {
	Convey("unaliasLang", t, func() {
		aliasFile = "testdata/locale.alias"
		So(unaliasLang("zh_CN"), ShouldEqual, "zh_CN")
		So(unaliasLang(""), ShouldEqual, "")
		So(unaliasLang("bokmal"), ShouldEqual, "nb_NO.ISO-8859-1")
	})
}

func TestGetLanguageNames(t *testing.T) {
	Convey("GetLanguageNames", t, func() {
		os.Setenv("LANGUAGE", "zh_CN")
		So(GetLanguageNames(), ShouldResemble, []string{"zh_CN", "zh", "C"})

		os.Setenv("LANGUAGE", "en_US")
		So(GetLanguageNames(), ShouldResemble, []string{"en_US", "en", "C"})
	})
}
