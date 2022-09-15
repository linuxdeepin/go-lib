// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"sync"
)

var (
	DefaultHomeConfigPrefix   = os.Getenv("HOME") + "/.config/deepin/"
	DefaultSystemConfigPrefix = "/var/cache/deepin/"
	DefaultConfigExt          = ".json"
)

type Config struct {
	configFile string
	lock       sync.Mutex
}

func (c *Config) Lock() {
	c.lock.Lock()
}

func (c *Config) Unlock() {
	c.lock.Unlock()
}

func (c *Config) SetConfigFile(file string) {
	c.configFile = file
}

func (c *Config) GetConfigFile() string {
	return c.configFile
}

func (c *Config) SetConfigName(name string) {
	c.SetConfigFile(DefaultHomeConfigPrefix + name + DefaultConfigExt)
}

func (c *Config) SetSystemConfigName(name string) {
	c.SetConfigFile(DefaultSystemConfigPrefix + name + DefaultConfigExt)
}

func (c *Config) IsConfigFileExists() bool {
	return IsFileExist(c.configFile)
}

func (c *Config) RemoveConfigFile() error {
	return os.Remove(c.configFile)
}

func (c *Config) Load(v interface{}) (err error) {
	if IsFileExist(c.configFile) {
		var fileContent []byte
		fileContent, err = ioutil.ReadFile(c.configFile)
		if err != nil {
			return
		}
		err = json.Unmarshal(fileContent, v)
	} else {
		err = c.Save(v)
	}
	return
}

func (c *Config) Save(v interface{}) (err error) {
	c.Lock()
	defer c.Unlock()
	_ = EnsureDirExist(path.Dir(c.configFile))
	var fileContent []byte
	fileContent, err = c.GetFileContentToSave(v)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(c.configFile, fileContent, 0644)
	return
}

func (c *Config) GetFileContentToSave(v interface{}) (fileContent []byte, err error) {
	fileContent, err = json.Marshal(v)
	return
}
