/*
 * Copyright (C) 2017 ~ 2018 Deepin Technology Co., Ltd.
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

package notify

import (
	"pkg.deepin.io/lib/notify/dbusnotify"
)

var (
	defaultAppName string
	notifier       *dbusnotify.Notifier
	isInitted      bool
)

// This must be called before anny other functions.
func Init(appName string) bool {
	if appName == "" {
		return false
	}

	if isInitted {
		return true
	}

	defaultAppName = appName

	var err error
	notifier, err = dbusnotify.NewNotifier(dbusDest, dbusPath)
	if err != nil {
		panic(err)
	}
	isInitted = true
	return true
}

// Gets whether or not libnotify is initialized.
func IsInitted() bool {
	return isInitted
}

// This should be called when the program no longer needs libnotify for
// the rest of its lifecycle, typically just before exitting.
func Destroy() {
	dbusnotify.DestroyNotifier(notifier)
	notifier = nil
	isInitted = false
}

func GetAppName() string {
	return defaultAppName
}

func SetAppName(name string) {
	defaultAppName = name
}

func GetServerCaps() ([]string, error) {
	return notifier.GetCapabilities()
}

type ServerInfo struct {
	Name, Vendor, Version, SpecVersion string
}

//name string, vendor string, version string, spec_version string
func GetServerInfo() (*ServerInfo, error) {
	name, vendor, version, specVersion, err := notifier.GetServerInformation()
	if err != nil {
		return nil, err
	}
	return &ServerInfo{
		Name:        name,
		Vendor:      vendor,
		Version:     version,
		SpecVersion: specVersion,
	}, nil
}
