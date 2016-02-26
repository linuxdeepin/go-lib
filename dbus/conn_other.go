/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

// +build !darwin

package dbus

import (
	"bytes"
	"errors"
	"os/exec"
)

func sessionBusPlatform() (*Conn, error) {
	cmd := exec.Command("/usr/bin/dbus-launch")
	b, err := cmd.CombinedOutput()

	if err != nil {
		return nil, err
	}

	i := bytes.IndexByte(b, '=')
	j := bytes.IndexByte(b, '\n')

	if i == -1 || j == -1 {
		return nil, errors.New("dbus: couldn't determine address of session bus")
	}

	return Dial(string(b[i+1 : j]))
}
