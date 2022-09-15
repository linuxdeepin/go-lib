// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package common

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

type Controller struct {
	Path          string
	Name          string
	MaxBrightness int
}

var errMaxBrightnessIsZero = errors.New("max brightness of controller is zero")

func NewController(path string) (*Controller, error) {
	c := &Controller{
		Path: path,
		Name: filepath.Base(path),
	}

	var err error
	c.MaxBrightness, err = c.GetPropInt("max_brightness")
	if err != nil {
		return nil, err
	}
	if c.MaxBrightness == 0 {
		return nil, errMaxBrightnessIsZero
	}

	return c, nil
}

func (c *Controller) GetBrightness() (int, error) {
	brightness, err := c.GetPropInt("brightness")
	if err != nil {
		return 0, err
	}
	return brightness, nil
}

func (c *Controller) GetPropBytes(name string) ([]byte, error) {
	filename := filepath.Join(c.Path, name)
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func (c *Controller) GetPropString(name string) (string, error) {
	bytes, err := c.GetPropBytes(name)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (c *Controller) GetPropInt(name string) (int, error) {
	str, err := c.GetPropString(name)
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		return 0, err
	}
	return num, nil
}

func ListControllerPaths(dir string) ([]string, error) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	paths := make([]string, len(fileInfos))

	for i, fileInfo := range fileInfos {
		paths[i] = filepath.Join(dir, fileInfo.Name())
	}
	return paths, nil
}
