package proxy

import (
	"dlib/glib-2.0"
	. "launchpad.net/gocheck"
	"os/exec"
	"testing"
	"time"
)

type ProxyTest struct{}

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

func init() {
	Suite(&ProxyTest{})
}

func (*ProxyTest) TestProxy(c *C) {
	SetupProxy()
	go func() {
		// execute default terminal every 10s
		for {
			time.Sleep(10 * time.Second)
			exec.Command("/usr/bin/default-terminal").Run()
		}
	}()
	glib.StartLoop()
}
