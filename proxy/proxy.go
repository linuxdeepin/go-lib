/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package proxy

import (
	"fmt"
	"gir/gio-2.0"
	"os"
	"pkg.deepin.io/lib/log"
	"pkg.deepin.io/lib/utils"
	"strconv"
	"strings"
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

	gkeyProxyAuto        = "autoconfig-url"
	gkeyProxyIgnoreHosts = "ignore-hosts"
	gkeyProxyHost        = "host"
	gkeyProxyPort        = "port"

	gchildProxyHttp  = "http"
	gchildProxyHttps = "https"
	gchildProxyFtp   = "ftp"
	gchildProxySocks = "socks"
)

var (
	logger                  = log.NewLogger("go-lib/proxy")
	proxySettings           = gio.NewSettings(gsettingsIdProxy)
	proxyChildSettingsHttp  *gio.Settings
	proxyChildSettingsHttps *gio.Settings
	proxyChildSettingsFtp   *gio.Settings
	proxyChildSettingsSocks *gio.Settings
)

// SetupProxy setup system proxy, need followed with glib.StartLoop().
func SetupProxy() {
	proxyChildSettingsHttp = proxySettings.GetChild(gchildProxyHttp)
	proxyChildSettingsHttps = proxySettings.GetChild(gchildProxyHttps)
	proxyChildSettingsFtp = proxySettings.GetChild(gchildProxyFtp)
	proxyChildSettingsSocks = proxySettings.GetChild(gchildProxySocks)
	updateProxyEnvs()
	listenProxyGsettings()
}

func listenProxyGsettings() {
	proxySettings.Connect("changed", func(s *gio.Settings, key string) {
		updateProxyEnvs()
	})
	proxyChildSettingsHttp.Connect("changed", func(s *gio.Settings, key string) {
		updateProxyEnvs()
	})
	proxyChildSettingsHttps.Connect("changed", func(s *gio.Settings, key string) {
		updateProxyEnvs()
	})
	proxyChildSettingsFtp.Connect("changed", func(s *gio.Settings, key string) {
		updateProxyEnvs()
	})
	proxyChildSettingsSocks.Connect("changed", func(s *gio.Settings, key string) {
		updateProxyEnvs()
	})
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

	utils.UnsetEnv(envHttpProxy)
	utils.UnsetEnv(envHttpsProxy)
	utils.UnsetEnv(envFtpProxy)
	utils.UnsetEnv(envSocksProxy)
	utils.UnsetEnv(envAutoProxy)
	utils.UnsetEnv(envAllProxy)
	utils.UnsetEnv(envNoProxy)
	proxyMode := proxySettings.GetString(gkeyProxyMode)
	switch proxyMode {
	case proxyModeNone:
	case proxyModeAuto:
		doSetEnv(envAutoProxy, proxySettings.GetString(gkeyProxyAuto))
	case proxyModeManual:
		doSetEnv(envHttpProxy, getProxyValue(proxyTypeHttp))
		doSetEnv(envHttpsProxy, getProxyValue(proxyTypeHttps))
		doSetEnv(envFtpProxy, getProxyValue(proxyTypeFtp))
		doSetEnv(envSocksProxy, getProxyValue(proxyTypeSocks))

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
}

func getProxyValue(proxyType string) (proxyValue string) {
	childSettings, err := getProxyChildSettings(proxyType)
	if err != nil {
		return
	}
	host := childSettings.GetString(gkeyProxyHost)
	if len(host) == 0 {
		return
	}
	port := strconv.Itoa(int(childSettings.GetInt(gkeyProxyPort)))
	proxyValue = fmt.Sprintf("%s://%s:%s", proxyType, host, port)
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
