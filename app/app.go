/*
 * Copyright (C) 2015 ~ 2017 Deepin Technology Co., Ltd.
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

// app provide a convenient application structure with commandline and profile features.
package app

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"pkg.deepin.io/lib/log"
	"pkg.deepin.io/lib/profile"
	"strings"
)

func toLogLevel(name string) (log.Priority, error) {
	name = strings.ToLower(name)
	logLevel := log.LevelInfo
	var err error
	switch name {
	case "":
	case "error":
		logLevel = log.LevelError
	case "warn":
		logLevel = log.LevelWarning
	case "info":
		logLevel = log.LevelInfo
	case "debug":
		logLevel = log.LevelDebug
	case "no":
		logLevel = log.LevelDisable
	default:
		err = fmt.Errorf("%s is not support", name)
	}

	return logLevel, err
}

// App is app structure with commandline and profile features.
type App struct {
	cmd      *kingpin.Application
	verbose  *bool
	logLevel *string
	memprof  *string
	cpuprof  *string
	profile  *profile.Config
}

// ParseCommandLine will parse command line which should not contains the executable name,
// and then return the sub-command. Exit if parse failed.
func (app *App) ParseCommandLine(args []string) string {
	subcmd := kingpin.MustParse(app.cmd.Parse(args))

	app.profile.CPUProfile = app.CpuProf()
	app.profile.MemProfile = app.MemProf()

	return subcmd
}

// Flag extends a new global flag. see kingpin for details.
func (app *App) Flag(longName string, desc string) *kingpin.FlagClause {
	return app.cmd.Flag(longName, desc)
}

// Command extends an sub-command. see kingpin for details.
func (app *App) Command(name string, desc string) *kingpin.CmdClause {
	return app.cmd.Command(name, desc)
}

// StartProfile starts all possible profiles.
func (app *App) StartProfile() error {
	return app.profile.Start()
}

// StopProfile stop profile, this should be called when shutdown hook is disabled,
// and this method does nothing when shutdown hook is enabled.
func (app *App) StopProfile() {
	app.profile.Stop()
}

// EnableNoHookShutdown enables or disable shutdown hook.
// If shutdown hook is disabled, StopProfile is needed to be called.
func (app *App) EnableNoShutdownHook(noShutdownHook bool) {
	app.profile.NoShutdownHook = noShutdownHook
}

// LogLevel returns the log level.
func (app *App) LogLevel() log.Priority {
	if *app.logLevel == "" && *app.verbose {
		return log.LevelDebug
	}
	lv, _ := toLogLevel(*app.logLevel)
	return lv
}

func (app *App) IsLogLevelNone() bool {
	return *app.verbose == false && *app.logLevel == ""
}

// MemProf returns memory profile's path.
func (app *App) MemProf() string {
	return *app.memprof
}

// CpuProf returns cpu profile's path
func (app *App) CpuProf() string {
	return *app.cpuprof
}

// New creates a new application according to name, description and version.
// There are some default command line flag:
// 	verbose(v for short): show much more message, shorthand for --loglevel debug which will be ignored if loglevel is specificed.
// 	loglevel(l for short): set log level, possible value is error/warn/info/debug/no, info is default.
// 	memprof: the file to save memory profile.
// 	cpuprof: the file to save cpu profile.
func New(name string, desc string, version string) *App {
	cmd := kingpin.New(name, desc)
	cmd.Version(version)

	app := &App{
		cmd:      cmd,
		profile:  &profile.Config{},
		verbose:  cmd.Flag("verbose", "Show much more message, shorthand for --loglevel debug which will be ignored if loglevel is specificed.").Short('v').Bool(),
		logLevel: cmd.Flag("loglevel", "Set log level, possible value is error/warn/info/debug/no, info is default.").Short('l').String(),
		memprof:  cmd.Flag("memprof", "Write memory profile to specific file").String(),
		cpuprof:  cmd.Flag("cpuprof", "Write cpu profile to specific file").String(),
	}
	return app
}
