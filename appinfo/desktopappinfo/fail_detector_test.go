package desktopappinfo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 测试简单的情况
func TestFailDetectorSimple(t *testing.T) {
	ch := make(chan struct{})
	fd := newFailDetector(ch)
	_, err := fd.Write([]byte("Failed to invoke: Booster: abcdef"))
	assert.Nil(t, err)
	assert.True(t, fd.failed)

	fd = newFailDetector(ch)
	_, err = fd.Write([]byte("deepin-turbo-invoker: error: xxxx not a file\n"))
	assert.Nil(t, err)
	assert.True(t, fd.failed)
}

// 测试多次写入，\n 后有其他字符的情况
func TestFailDetectorWrite(t *testing.T) {
	ch := make(chan struct{})
	fd := newFailDetector(ch)
	_, err := fd.Write([]byte("line1\n"))
	assert.Nil(t, err)
	assert.False(t, fd.failed)
	assert.Len(t, fd.buf.Bytes(), 0)

	_, err = fd.Write([]byte("line2\nabc"))
	assert.Nil(t, err)
	assert.False(t, fd.failed)
	assert.Equal(t, []byte("abc"), fd.buf.Bytes())

	_, err = fd.Write(nil)
	assert.Nil(t, err)
	assert.False(t, fd.failed)
	assert.Equal(t, []byte("abc"), fd.buf.Bytes())

	_, err = fd.Write([]byte("abc-abc\nFailed to invoke: Booster: abcdef\nline 3"))
	assert.Nil(t, err)
	assert.True(t, fd.failed)
	assert.Zero(t, fd.buf.Bytes())
}
