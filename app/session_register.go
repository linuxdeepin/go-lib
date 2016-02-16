/**
 * Copyright (C) 2015 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package app

import (
	"dbus/com/deepin/sessionmanager"
	"fmt"
	"os"
	"pkg.deepin.io/lib/utils"
)

// DDESessionRegister will register to session manager if program is started from startdde.
func DDESessionRegister() {
	cookie := os.ExpandEnv("$DDE_SESSION_PROCESS_COOKIE_ID")
	utils.UnsetEnv("DDE_SESSION_PROCESS_COOKIE_ID")
	if cookie == "" {
		fmt.Println("get DDE_SESSION_PROCESS_COOKIE_ID failed")
		return
	}
	go func() {
		manager, err := sessionmanager.NewSessionManager("com.deepin.SessionManager", "/com/deepin/SessionManager")
		if err != nil {
			fmt.Println("register failed:", err)
			return
		}
		manager.Register(cookie)
	}()
}
