// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package proxy

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/godbus/dbus/v5"

	"github.com/linuxdeepin/go-gir/gio-2.0"
	"github.com/linuxdeepin/go-lib/gsettings"
	"github.com/linuxdeepin/go-lib/log"
	"github.com/linuxdeepin/go-lib/utils"
)

// Synchronize proxy gsettings to environment variables.
//
// Examples of proxy environment variables:
// http_proxy="http://user:pass@127.0.0.1:8080/"
// https_proxy="https://127.0.0.1:8080/"
// ftp_proxy="ftp://127.0.0.1:8080/"
// all_proxy="http://127.0.0.1:8080/"
// SOCKS_SERVER=socks5://127.0.0.1:8000/
// no_proxy="localhost,127.0.0.0/8,::1"

const (
	// general proxy environment variables, works for wget/curl/aria2c
	envAutoProxy  = "auto_proxy"
	envHttpProxy  = "http_proxy"
	envHttpsProxy = "https_proxy"
	envFtpProxy   = "ftp_proxy"
	envAllProxy   = "all_proxy"
	envNoProxy    = "no_proxy"

	// special proxy environment variable for chrome
	envSocksProxy = "SOCKS_SERVER"

	gsettingsIdProxy = "com.deepin.wrap.gnome.system.proxy"

	proxyTypeHttp  = "http"
	proxyTypeHttps = "https"
	proxyTypeFtp   = "ftp"
	proxyTypeSocks = "socks5"

	gkeyProxyMode   = "mode"
	proxyModeNone   = "none"
	proxyModeManual = "manual"
	proxyModeAuto   = "auto"

	gkeyProxyAuto                   = "autoconfig-url"
	gkeyProxyIgnoreHosts            = "ignore-hosts"
	gkeyProxyHost                   = "host"
	gkeyProxyPort                   = "port"
	gkeyProxyUseAuthentication      = "use-authentication"
	gkeyProxyAuthenticationUser     = "authentication-user"
	gkeyProxyAuthenticationPassword = "authentication-password"

	gchildProxyHttp  = "http"
	gchildProxyHttps = "https"
	gchildProxyFtp   = "ftp"
	gchildProxySocks = "socks"
)

var (
	logger                  = log.NewLogger("go-lib/proxy")
	proxySettings           *gio.Settings
	proxyChildSettingsHttp  *gio.Settings
	proxyChildSettingsHttps *gio.Settings
	proxyChildSettingsFtp   *gio.Settings
	proxyChildSettingsSocks *gio.Settings
)

// SetupProxy setup system proxy, need followed with glib.StartLoop().
func SetupProxy() {
	proxySettings = gio.NewSettings(gsettingsIdProxy)
	proxyChildSettingsHttp = proxySettings.GetChild(gchildProxyHttp)
	proxyChildSettingsHttps = proxySettings.GetChild(gchildProxyHttps)
	proxyChildSettingsFtp = proxySettings.GetChild(gchildProxyFtp)
	proxyChildSettingsSocks = proxySettings.GetChild(gchildProxySocks)
	updateProxyEnvs()
	listenProxyGsettings()
}

func listenProxyGsettings() {
	changedHandler := func(key string) {
		updateProxyEnvs()
	}

	const systemProxy = "system.proxy"
	gsettings.ConnectChanged(systemProxy, "*", changedHandler)
	gsettings.ConnectChanged(systemProxy+"."+gchildProxyHttp, "*", changedHandler)
	gsettings.ConnectChanged(systemProxy+"."+gchildProxyHttps, "*", changedHandler)
	gsettings.ConnectChanged(systemProxy+"."+gchildProxyFtp, "*", changedHandler)
	gsettings.ConnectChanged(systemProxy+"."+gchildProxySocks, "*", changedHandler)
}

func showEnvs() {
	showEnv(envHttpProxy)
	showEnv(envHttpsProxy)
	showEnv(envFtpProxy)
	showEnv(envSocksProxy)
	showEnv(envAllProxy)
	showEnv(envAutoProxy)
	showEnv(envNoProxy)
}

func showEnv(envName string) {
	if utils.IsEnvExists(envName) {
		logger.Debug(envName, os.Getenv(envName))
	} else {
		logger.Debug(envName, "<not exists>")
	}
}

func updateProxyEnvs() {
	logger.Debug("update proxy environment variables...")

	os.Unsetenv(envHttpProxy)
	os.Unsetenv(envHttpsProxy)
	os.Unsetenv(envFtpProxy)
	os.Unsetenv(envSocksProxy)
	os.Unsetenv(envAutoProxy)
	os.Unsetenv(envAllProxy)
	os.Unsetenv(envNoProxy)
	systemdAndDbusUnSetEnv(envHttpProxy)
	systemdAndDbusUnSetEnv(envHttpsProxy)
	systemdAndDbusUnSetEnv(envFtpProxy)
	systemdAndDbusUnSetEnv(envSocksProxy)
	systemdAndDbusUnSetEnv(envAutoProxy)
	systemdAndDbusUnSetEnv(envAllProxy)
	systemdAndDbusUnSetEnv(envNoProxy)
	proxyMode := proxySettings.GetString(gkeyProxyMode)
	switch proxyMode {
	case proxyModeNone:
	case proxyModeAuto:
		doSetEnv(envAutoProxy, proxySettings.GetString(gkeyProxyAuto))
	case proxyModeManual:
		doSetEnv(envHttpProxy, getProxyValue(proxyTypeHttp, proxyTypeHttp))
		doSetEnv(envHttpsProxy, getProxyValue(proxyTypeHttps, proxyTypeHttp))
		doSetEnv(envFtpProxy, getProxyValue(proxyTypeFtp, proxyTypeHttp))
		doSetEnv(envAllProxy, getProxyValue(proxyTypeSocks, proxyTypeSocks))
		doSetEnv(envSocksProxy, getProxyValue(proxyTypeSocks, proxyTypeSocks))

		arrayIgnoreHosts := proxySettings.GetStrv(gkeyProxyIgnoreHosts)
		ignoreHosts := strings.Join(arrayIgnoreHosts, ",")
		doSetEnv(envNoProxy, ignoreHosts)

		// fallback socks proxy value to http to be compatible with Qt>=4.6
		if utils.IsEnvExists(envSocksProxy) && !utils.IsEnvExists(envHttpProxy) {
			doSetEnv(envHttpProxy, os.Getenv(envSocksProxy))
		}
	}
	showEnvs()
}

func doSetEnv(env, value string) {
	if len(value) > 0 {
		os.Setenv(env, value)
	}
	systemdAndDbusSetEnv(env, value)
}

func systemdAndDbusSetEnv(env, value string) {
	if len(value) > 0 {
		bus, err := dbus.SessionBus()
		if err != nil {
			logger.Warning(err)
			return
		}

		systemdObj := bus.Object("org.freedesktop.systemd1", "/org/freedesktop/systemd1")
		err = systemdObj.Call("org.freedesktop.systemd1.Manager.SetEnvironment", 0, []string{fmt.Sprintf("%s=%s", env, value)}).Err
		if err != nil {
			logger.Warning(err)
		}
		dbusObj := bus.Object("org.freedesktop.DBus", "/org/freedesktop/DBus")
		envMap := make(map[string]string)
		envMap[env] = value
		err = dbusObj.Call("org.freedesktop.DBus.UpdateActivationEnvironment", 0, envMap).Err
		if err != nil {
			logger.Warning(err)
		}
		logger.Debug("update dbus and systemd env")
	}
}

func systemdAndDbusUnSetEnv(env string) {
	bus, err := dbus.SessionBus()
	if err != nil {
		logger.Warning(err)
		return
	}

	dbusObj := bus.Object("org.freedesktop.DBus", "/org/freedesktop/DBus")
	envMap := make(map[string]string)
	envMap[env] = ""
	err = dbusObj.Call("org.freedesktop.DBus.UpdateActivationEnvironment", 0, envMap).Err
	if err != nil {
		logger.Warning(err)
	}

	// dbus will register with systemd when cleaning environment variables, so systemd needs to be cleaned up at the end
	systemdObj := bus.Object("org.freedesktop.systemd1", "/org/freedesktop/systemd1")
	err = systemdObj.Call("org.freedesktop.systemd1.Manager.UnsetEnvironment", 0, []string{env}).Err
	if err != nil {
		logger.Warning(err)
	}

	logger.Debug("unset dbus and systemd env")
}

func getProxyValue(proxyType string, protocol string) (proxyValue string) {
	childSettings, err := getProxyChildSettings(proxyType)
	if err != nil {
		return
	}
	host := childSettings.GetString(gkeyProxyHost)
	if len(host) == 0 {
		return
	}
	port := strconv.Itoa(int(childSettings.GetInt(gkeyProxyPort)))

	useAuthentication := false
	if childSettings.GetSchema().HasKey(gkeyProxyUseAuthentication) {
		useAuthentication = childSettings.GetBoolean(gkeyProxyUseAuthentication)
	}
	if useAuthentication {
		user := childSettings.GetString(gkeyProxyAuthenticationUser)
		password := childSettings.GetString(gkeyProxyAuthenticationPassword)
		proxyValue = fmt.Sprintf("%s://%s:%s@%s:%s", protocol, user, password, host, port)
	} else {
		proxyValue = fmt.Sprintf("%s://%s:%s", protocol, host, port)
	}
	return
}

func getProxyChildSettings(proxyType string) (childSettings *gio.Settings, err error) {
	switch proxyType {
	case proxyTypeHttp:
		childSettings = proxyChildSettingsHttp
	case proxyTypeHttps:
		childSettings = proxyChildSettingsHttps
	case proxyTypeFtp:
		childSettings = proxyChildSettingsFtp
	case proxyTypeSocks:
		childSettings = proxyChildSettingsSocks
	default:
		err = fmt.Errorf("not a valid proxy type: %s", proxyType)
		logger.Error(err)
	}
	return
}
