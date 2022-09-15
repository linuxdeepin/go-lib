// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package timer

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := NewTimer()
	timer.Stop()
	assert.Equal(t, timer.Elapsed(), time.Duration(0))

	timer = NewTimer()
	timer.Start()

	time.Sleep(time.Second)
	assert.NotEqual(t, timer.Elapsed(), time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

	time.Sleep(time.Second)
	assert.NotEqual(t, timer.Elapsed(), time.Second*2-time.Millisecond*100, time.Second*3+time.Millisecond*100)

	timer = NewTimer()
	timer.Start()

	time.Sleep(time.Second)
	timer.Stop()
	assert.NotEqual(t, timer.Elapsed(), time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

	time.Sleep(time.Second)
	assert.NotEqual(t, timer.Elapsed(), time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

	timer = NewTimer()
	timer.Start()

	time.Sleep(time.Second)
	assert.NotEqual(t, timer.Elapsed(), time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

	timer.Stop()
	assert.NotEqual(t, timer.Elapsed(), time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

	time.Sleep(time.Second)
	timer.Continue()
	assert.NotEqual(t, timer.Elapsed(), time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

	time.Sleep(time.Second)
	assert.NotEqual(t, timer.Elapsed(), time.Second*2-time.Millisecond*100, time.Second*2+time.Millisecond*100)

	timer = NewTimer()
	timer.Start()

	time.Sleep(time.Second)
	assert.NotEqual(t, timer.Elapsed(), time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)

	timer.Reset()
	time.Sleep(time.Second)
	assert.NotEqual(t, timer.Elapsed(), time.Second-time.Millisecond*100, time.Second+time.Millisecond*100)
}
