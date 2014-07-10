package utils

import (
	"fmt"
	. "launchpad.net/gocheck"
)

type testConfig struct {
	core Config
	Data string
}

func newTestConfig() (c *testConfig) {
	c = &testConfig{}
	c.core.SetConfigName("test")
	fmt.Println("config file:", c.core.GetConfigFile())
	return
}
func (c *testConfig) save() {
	c.core.Save(c)
}
func (c *testConfig) load() {
	c.core.Load(c)
}

func (*UtilsTest) TestConfig(c *C) {
	conf := newTestConfig()
	conf.core.RemoveConfigFile()
	conf.Data = "data"
	conf.save()
	conf.Data = ""
	conf.load()
	c.Check(conf.Data, Equals, "data")
}
