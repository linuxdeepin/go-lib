// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package appinfo

type AppInfo interface {
	GetId() string
	GetName() string
	GetIcon() string
	GetExecutable() string
	GetFileName() string
	GetCommandline() string
	Launch(files []string, launchContext *AppLaunchContext) error
	GetStartupWMClass() string
}
