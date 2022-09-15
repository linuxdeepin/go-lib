// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package desktopappinfo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// 测试简单的情况
func TestFailDetectorSimple(t *testing.T) {
	ch := make(chan struct{}, 2)
	fd := newFailDetector(ch)
	_, err := fd.Write([]byte("Failed to invoke: Booster: abcdef"))
	require.NoError(t, err)
	assert.True(t, fd.done)

	fd = newFailDetector(ch)
	_, err = fd.Write([]byte("deepin-turbo-invoker: error: xxxx not a file\n"))
	require.NoError(t, err)
	assert.True(t, fd.done)
}

// 测试多次写入，\n 后有其他字符的情况
func TestFailDetectorWrite(t *testing.T) {
	ch := make(chan struct{}, 2)
	fd := newFailDetector(ch)
	_, err := fd.Write([]byte("line1\n"))
	require.NoError(t, err)
	assert.False(t, fd.done)
	assert.Len(t, fd.buf.Bytes(), 0)

	_, err = fd.Write([]byte("line2\nabc"))
	require.NoError(t, err)
	assert.False(t, fd.done)
	assert.Equal(t, []byte("abc"), fd.buf.Bytes())

	_, err = fd.Write(nil)
	require.NoError(t, err)
	assert.False(t, fd.done)
	assert.Equal(t, []byte("abc"), fd.buf.Bytes())

	_, err = fd.Write([]byte("abc-abc\nFailed to invoke: Booster: abcdef\nline 3"))
	require.NoError(t, err)
	assert.True(t, fd.done)
	assert.Zero(t, fd.buf.Bytes())
}
