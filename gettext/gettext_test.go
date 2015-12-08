package gettext

import (
	C "launchpad.net/gocheck"
	"os"
	"testing"
)

type gettext struct{}

func Test(t *testing.T) { C.TestingT(t) }

func init() {
	C.Suite(&gettext{})
}

func (*gettext) TestTr(c *C.C) {
	InitI18n()
	Bindtextdomain("test", "testdata/locale")
	os.Setenv("LANGUAGE", "ar")

	c.Check(Tr("Back"), C.Equals, "الخلف")
}

func (*gettext) TestDGettext(c *C.C) {
	InitI18n()
	Bindtextdomain("test", "testdata/locale")
	os.Setenv("LANGUAGE", "zh_CN")
	c.Check(DGettext("test", "Back"), C.Equals, "返回")
}

func (*gettext) TestFailed(c *C.C) {
	InitI18n()
	Bindtextdomain("test", "testdata/locale")
	c.Check(DGettext("test", "notfound"), C.Equals, "notfound")
	c.Check(DGettext("test", "未找到"), C.Equals, "未找到")
}

func (*gettext) TestNTr(c *C.C) {

	Bindtextdomain("test", "testdata/plural/locale")
	Textdomain("test")

	InitI18n()
	os.Setenv("LANGUAGE", "es")
	c.Check(NTr("%d apple", "%d apples", 1), C.Equals, "%d manzana")
	c.Check(NTr("%d apple", "%d apples", 2), C.Equals, "%d manzanas")

	InitI18n()
	os.Setenv("LANGUAGE", "zh_CN")
	c.Check(NTr("%d apple", "%d apples", 0), C.Equals, "%d苹果")
	c.Check(NTr("%d apple", "%d apples", 1), C.Equals, "%d苹果")
	c.Check(NTr("%d apple", "%d apples", 2), C.Equals, "%d苹果")
}

func (*gettext) TestDNGettext(c *C.C) {
	Bindtextdomain("test", "testdata/plural/locale")

	InitI18n()
	os.Setenv("LANGUAGE", "es")
	c.Check(DNGettext("test", "%d person", "%d persons", 1), C.Equals, "%d persona")
	c.Check(DNGettext("test", "%d person", "%d persons", 2), C.Equals, "%d personas")
	InitI18n()
	os.Setenv("LANGUAGE", "zh_CN")
	c.Check(DNGettext("test", "%d person", "%d persons", 0), C.Equals, "%d人")
	c.Check(DNGettext("test", "%d person", "%d persons", 1), C.Equals, "%d人")
	c.Check(DNGettext("test", "%d person", "%d persons", 2), C.Equals, "%d人")
}
