package proxy

import (
	"dlib/glib-2.0"
	. "launchpad.net/gocheck"
	"testing"
)

type ProxyTest struct{}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func init() {
	Suite(&ProxyTest{})
}

func (*ProxyTest) TestProxy(c *C) {
	SetupProxy()
	glib.StartLoop()
}
