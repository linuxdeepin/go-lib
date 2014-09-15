package utils

import (
	C "launchpad.net/gocheck"
)

type testConfig struct {
	core Config
	Data string
}

func newTestConfig() (c *testConfig) {
	c = &testConfig{}
	c.core.SetConfigName("test")
	return
}
func (c *testConfig) save() {
	c.core.Save(c)
}
func (c *testConfig) load() {
	c.core.Load(c)
}

func (*testWrapper) TestConfig(c *C.C) {
	conf := newTestConfig()
	conf.core.RemoveConfigFile()
	conf.Data = "data"
	conf.save()
	conf.Data = ""
	conf.load()
	c.Check(conf.Data, C.Equals, "data")
}
