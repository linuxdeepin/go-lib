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
	"gir/gio-2.0"
	"os"
	"pkg.deepin.io/lib/log"
	"pkg.deepin.io/lib/utils"
)

const (
	envAutoProxy  = "auto_proxy"
	envHttpProxy  = "http_proxy"
	envHttpsProxy = "https_proxy"
	envFtpProxy   = "ftp_proxy"
	envSocksProxy = "SOCKS_SERVER"

	gsettingsIdProxy = "com.deepin.dde.proxy"
	gkeyProxyMethod  = "proxy-method"

	proxyMethodNone   = "none"
	proxyMethodManual = "manual"
	proxyMethodAuto   = "auto"

	gkeyAutoProxy = "auto-proxy"

	gkeyHttpProxy  = "http-proxy"
	gkeyHttpsProxy = "https-proxy"
	gkeyFtpProxy   = "ftp-proxy"
	gkeySocksProxy = "socks-proxy"
)

var (
	proxySettings = gio.NewSettings(gsettingsIdProxy)
	logger        = log.NewLogger("com.deepin.dlib.proxy")
)

// SetupProxy setup system proxy, need followed with glib.StartLoop().
func SetupProxy() {
	updateProxyEnvs()
	listenProxyGsettings()
}

func listenProxyGsettings() {
	proxySettings.Connect("changed", func(s *gio.Settings, key string) {
		logger.Debug("proxy keys in gsettings changed", key, proxySettings.GetString(key))
		updateProxyEnvs()
	})
}

func showEnvs() {
	showEnv(envAutoProxy)
	showEnv(envHttpProxy)
	showEnv(envHttpsProxy)
	showEnv(envFtpProxy)
	showEnv(envSocksProxy)
}

func showEnv(envName string) {
	if utils.IsEnvExists(envName) {
		logger.Debug(envName, os.Getenv(envName))
	} else {
		logger.Debug(envName, "<not exists>")
	}
}

// BUGGY. DO NOT UNCOMMENT UNTIL IT'S FIXED.
func updateProxyEnvs() {
//	utils.UnsetEnv(envAutoProxy)
//	utils.UnsetEnv(envHttpProxy)
//	utils.UnsetEnv(envHttpsProxy)
//	utils.UnsetEnv(envFtpProxy)
//	utils.UnsetEnv(envSocksProxy)
//	proxyMethod := proxySettings.GetString(gkeyProxyMethod)
//	switch proxyMethod {
//	case proxyMethodNone:
//	case proxyMethodAuto:
//		autoProxy := proxySettings.GetString(gkeyAutoProxy)
//		if len(autoProxy) > 0 {
//			os.Setenv(envAutoProxy, autoProxy)
//		}
//	case proxyMethodManual:
//		httpProxy := proxySettings.GetString(gkeyHttpProxy)
//		if len(httpProxy) > 0 {
//			os.Setenv(envHttpProxy, httpProxy)
//		}
//
//		httpsProxy := proxySettings.GetString(gkeyHttpsProxy)
//		if len(httpsProxy) > 0 {
//			os.Setenv(envHttpsProxy, httpsProxy)
//		}
//
//		ftpProxy := proxySettings.GetString(gkeyFtpProxy)
//		if len(ftpProxy) > 0 {
//			os.Setenv(envFtpProxy, ftpProxy)
//		}
//
//		socksProxy := proxySettings.GetString(gkeySocksProxy)
//		if len(socksProxy) > 0 {
//			os.Setenv(envSocksProxy, socksProxy)
//		}
//	}
//	showEnvs()
}
