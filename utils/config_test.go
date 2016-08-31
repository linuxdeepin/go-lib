/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

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
func (c *testConfig) save() error {
	return c.core.Save(c)
}
func (c *testConfig) load() {
	c.core.Load(c)
}

func (*testWrapper) TestConfig(c *C.C) {
	conf := newTestConfig()
	conf.core.RemoveConfigFile()
	conf.Data = "data"
	err := conf.save()
	if err != nil {
		c.Skip("Save config failed:" + err.Error())
		return
	}
	conf.Data = ""
	conf.load()
	c.Check(conf.Data, C.Equals, "data")
}
