// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package appinfo

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/linuxdeepin/go-x11-client"
)

var (
	pid      int
	prog     string
	hostname string
)

func init() {
	pid = os.Getpid()
	prog = filepath.Base(os.Args[0])
	hostname, _ = os.Hostname()
}

type AppLaunchContext struct {
	sync.Mutex
	conn        *x.Conn
	count       uint
	timestamp   uint32
	cmdPrefixes []string
	cmdSuffixes []string
	env         []string
}

func NewAppLaunchContext(conn *x.Conn) *AppLaunchContext {
	return &AppLaunchContext{
		conn: conn,
	}
}

func (ctx *AppLaunchContext) SetEnv(env []string) {
	ctx.env = env
}

func (ctx *AppLaunchContext) SetTimestamp(timestamp uint32) {
	ctx.timestamp = timestamp
}

func (ctx *AppLaunchContext) GetTimestamp() uint32 {
	return ctx.timestamp
}

func (ctx *AppLaunchContext) SetCmdPrefixes(v []string) {
	ctx.cmdPrefixes = v
}

func (ctx *AppLaunchContext) GetCmdPrefixes() []string {
	return ctx.cmdPrefixes
}

func (ctx *AppLaunchContext) SetCmdSuffixes(v []string) {
	ctx.cmdSuffixes = v
}

func (ctx *AppLaunchContext) GetEnv() []string {
	return ctx.env
}

func (ctx *AppLaunchContext) GetCmdSuffixes() []string {
	return ctx.cmdSuffixes
}

func (ctx *AppLaunchContext) GetStartupNotifyId(appInfo AppInfo, files []string) (string, error) {
	execBase := filepath.Base(appInfo.GetExecutable())
	snId := fmt.Sprintf("%s-%d-%s-%s-%d_TIME%d", prog, pid, hostname, execBase, ctx.count, ctx.timestamp)
	ctx.count++

	screenStr := strconv.Itoa(ctx.conn.ScreenNumber)
	// send new msg
	msg := &StartupNotifyMessage{
		Type: "new",
		KeyValues: map[string]string{
			"ID":     snId,
			"SCREEN": screenStr,
			"NAME":   appInfo.GetName(),
		},
	}
	startupWMClass := appInfo.GetStartupWMClass()
	if startupWMClass != "" {
		msg.KeyValues["WMCLASS"] = startupWMClass
	}

	err := msg.Broadcast(ctx.conn)
	if err != nil {
		return "", err
	}
	return snId, nil
}

func (ctx *AppLaunchContext) LaunchFailed(startupNotifyId string) error {
	// send remove msg
	msg := &StartupNotifyMessage{
		Type: "remove",
		KeyValues: map[string]string{
			"ID": startupNotifyId,
		},
	}
	return msg.Broadcast(ctx.conn)
}
