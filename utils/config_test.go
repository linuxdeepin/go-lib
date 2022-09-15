// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	_ = c.core.Load(c)
}

func TestConfig(t *testing.T) {
	conf := newTestConfig()
	_ = conf.core.RemoveConfigFile()
	conf.Data = "data"
	err := conf.save()
	if err != nil {
		t.Skip("Save config failed:" + err.Error())
		return
	}
	conf.Data = ""
	conf.load()
	assert.Equal(t, conf.Data, "data")
}
