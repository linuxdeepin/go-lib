/*
 * Copyright (C) 2014 ~ 2017 Deepin Technology Co., Ltd.
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

package utils

import (
	C "gopkg.in/check.v1"
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
