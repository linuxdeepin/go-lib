// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package procfs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFile(t *testing.T) {
	p := Process(1)
	assert.Equal(t, p.getFile("cwd"), "/proc/1/cwd")
}

func TestExist(t *testing.T) {
	p := Process(os.Getpid())
	assert.True(t, p.Exist())
}

func TestCmdline(t *testing.T) {
	p := Process(os.Getpid())
	cmdline, err := p.Cmdline()
	require.NoError(t, err)
	t.Log("cmdline:", cmdline)
	assert.True(t, len(cmdline) > 0)
}

func TestCwd(t *testing.T) {
	p := Process(os.Getpid())
	cwd, err := p.Cwd()
	require.NoError(t, err)
	t.Log("cwd:", cwd)

	osWd, err1 := os.Getwd()
	require.Nil(t, err1)
	assert.Equal(t, cwd, osWd)
}

func TestExe(t *testing.T) {
	p := Process(os.Getpid())
	exe, err := p.Exe()
	require.NoError(t, err)
	t.Log("exe:", exe)
	assert.True(t, len(exe) > 0)
}

func TestEnvVars(t *testing.T) {
	vars := EnvVars{
		"PWD=/a/b/c",
	}
	pwd, ok := vars.Lookup("PWD")
	assert.Equal(t, pwd, "/a/b/c")
	assert.True(t, ok)

	abc, ok := vars.Lookup("abc")
	assert.Equal(t, abc, "")
	assert.False(t, ok)

	pwd = vars.Get("PWD")
	assert.Equal(t, pwd, "/a/b/c")

	abc = vars.Get("abc")
	assert.Equal(t, abc, "")
}

func TestEnvion(t *testing.T) {
	p := Process(os.Getpid())
	environ, err := p.Environ()
	require.NoError(t, err)
	assert.True(t, len(environ) > 0)
	for _, aVar := range environ {
		t.Log(string(aVar))
	}

	path, ok := environ.Lookup("PATH")
	assert.True(t, ok)
	assert.True(t, path != "")

	home, ok := environ.Lookup("HOME")
	assert.True(t, ok)
	assert.True(t, home != "")

	xxx, ok := environ.Lookup("XXXXXXXXXXXXXXX")
	assert.False(t, ok)
	assert.Equal(t, xxx, "")
}

func TestStatus(t *testing.T) {
	p := Process(os.Getpid())
	status, err := p.Status()
	require.NoError(t, err)
	assert.NotEmpty(t, status)

	// test lookup
	val, err := status.lookup("XXX")
	assert.Empty(t, val)
	assert.Equal(t, err, StatusFieldNotFoundErr{"XXX"})
	assert.Equal(t, err.Error(), "field XXX is not found in proc status file")

	// test Uids
	uids, err := status.Uids()
	require.NoError(t, err)
	t.Log("uids:", uids)
	assert.Equal(t, uids[0], uint(os.Getuid()))
	assert.Equal(t, uids[1], uint(os.Geteuid()))

	// test PPid
	ppid, err := status.PPid()
	require.NoError(t, err)
	t.Log("ppid:", ppid)
	assert.True(t, ppid > 0)
}
