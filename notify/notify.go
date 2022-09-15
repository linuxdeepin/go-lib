// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package notify

import (
	"github.com/linuxdeepin/go-lib/notify/dbusnotify"
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
