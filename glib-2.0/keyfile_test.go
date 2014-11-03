package glib

import (
	C "launchpad.net/gocheck"
	"os"
	"strings"
	"testing"
)

type glib struct{}

func Test(t *testing.T) { C.TestingT(t) }

func init() {
	C.Suite(&glib{})
}

var desktop_file = `
[Desktop Entry]
Categories=GNOME;GTK;Utility;System;TerminalEmulator;
Comment=Use the command line
Comment[am]=የ ትእዛዝ መስመር ይጠቀሙ
Comment[ar]=إستعمال سطر الأوامر
Comment[zh_CN]=使用命令行
Keywords=shell;prompt;command;commandline;
Keywords[an]=enterprite ;indicador;comando;linia de comandos;
Keywords[ar]=طرفية;صدفة;سطر;أوامر;
Keywords[as]=শ্বেল;প্ৰমপ্ট;কমান্ড;কমান্ডশাৰী;
Keywords[zh_CN]=shell;prompt;command;commandline;命令;提示符;命令行;
Name=Deepin Terminal
Name[es_AR]=Terminal Deepin
Name[ja]=ターミナル
StartupNotify=true
TryExec=deepin-terminal
Type=Application

[NewQuake Shortcut Group]
Exec=deepin-terminal --quake-mode
Name=New Quake Window

[NewWindow Shortcut Group]
Exec=deepin-terminal
Name=New Window
Name[vi]=Cửa sổ Mới
Name[zh_CN]=新建窗口
Name[zh_TW]=開啟新視窗
`

func checkDesktopFile(f *KeyFile, c *C.C) {
	{
		r, err := f.GetBoolean("Desktop Entry", "StartupNotify")
		c.Check(err, C.Equals, nil)
		c.Check(r, C.Equals, true)
	}
	{
		r, err := f.GetString("Desktop Entry", "Type")
		c.Check(err, C.Equals, nil)
		c.Check(r, C.Equals, "Application")
	}
	{
		l, r, err := f.GetStringList("Desktop Entry", "Keywords")
		c.Check(err, C.Equals, nil)
		c.Check(l, C.Equals, uint64(4))
		c.Check(strings.Join(r, ";"), C.Equals, "shell;prompt;command;commandline")
	}
	{
		r, err := f.GetLocaleString("NewWindow Shortcut Group", "Name", "zh_CN")
		c.Check(err, C.Equals, nil)
		c.Check(r, C.Equals, "新建窗口")

		os.Setenv("LANGUAGE", "zh_TW")
		r, err = f.GetLocaleString("NewWindow Shortcut Group", "Name", "\x00")
		c.Check(err, C.Equals, nil)
		c.Check(r, C.Equals, "開啟新視窗")
	}
	{
		l, r, err := f.GetLocaleStringList("Desktop Entry", "Keywords", "zh_CN")
		c.Check(err, C.Equals, nil)
		c.Check(l, C.Equals, uint64(7))
		c.Check(strings.Join(r, ";"), C.Equals, "shell;prompt;command;commandline;命令;提示符;命令行")
	}
	{
		_, err := f.GetDouble("Can'tFind", "Name")
		c.Check(err, C.ErrorMatches, "Key file does not have group 'Can'tFind'")
	}
}

func (*glib) TestKeyFileFromFile(c *C.C) {
	f := NewKeyFile()
	ok, err := f.LoadFromFile("testdata/deepin-terminal.desktop", KeyFileFlagsKeepTranslations)
	if !ok || err != nil {
		c.Fatal(ok, err)
	}
	checkDesktopFile(f, c)
}

func (*glib) TestUserDirs(c *C.C) {
	dirs := []UserDirectory{
		UserDirectoryDirectoryDesktop,
		UserDirectoryDirectoryDocuments,
		UserDirectoryDirectoryDownload,
		UserDirectoryDirectoryMusic,
		UserDirectoryDirectoryPictures,
		UserDirectoryDirectoryPublicShare,
		UserDirectoryDirectoryTemplates,
		UserDirectoryDirectoryVideos,
	}

	for _, d := range dirs {
		GetUserSpecialDir(d)
	}
}
