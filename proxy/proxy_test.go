package proxy

import (
	"dlib/gio-2.0"
	"dlib/glib-2.0"
	"dlib/utils"
	"fmt"
	. "launchpad.net/gocheck"
	"os"
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
	showEnv := func(envName string) {
		if utils.IsEnvExists(envName) {
			fmt.Println(envName, os.Getenv(envName))
		} else {
			fmt.Println(envName, "not exists")
		}
	}
	proxySettings.Connect("changed", func(s *gio.Settings, key string) {
		fmt.Println("proxy gsettings changed", key, proxySettings.GetString(key))
		go func() {
			time.Sleep(500 * time.Millisecond)
			fmt.Println()
			showEnv(envAutoProxy)
			showEnv(envHttpProxy)
			showEnv(envHttpsProxy)
			showEnv(envFtpProxy)
			showEnv(envSocksProxy)
			showEnv(envSocksVersion)
		}()
	})
	glib.StartLoop()
}
