package gettext

import (
	C "launchpad.net/gocheck"
	"os"
	"testing"
)

type gettext struct{}

func Test(t *testing.T) { C.TestingT(t) }

func init() {
	Bindtextdomain("test", "testdata/locale")
	C.Suite(&gettext{})
}

func (*gettext) TestTr(c *C.C) {
	os.Setenv("LANGUAGE", "ar")
	InitI18n()

	Textdomain("test")
	c.Check(Tr("Back"), C.Equals, "الخلف")
}

func (*gettext) TestBGettext(c *C.C) {
	os.Setenv("LANGUAGE", "zh_CN")
	InitI18n()
	c.Check(DGettext("test", "Back"), C.Equals, "返回")
}

func (*gettext) TestFailed(c *C.C) {
	c.Check(DGettext("test", "notfound"), C.Equals, "notfound")
	c.Check(DGettext("test", "未找到"), C.Equals, "未找到")
}
