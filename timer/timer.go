/*
 * Copyright (C) 2014 ~ 2017 Deepin Technology Co., Ltd.
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

package timer

import (
	"time"
)

type _TimerState uint16

const (
	_TimerStateNotStarted _TimerState = iota
	_TimerStateStarted
	_TimerStateStopped
)

// Timer is used for reporting rate.
type Timer struct {
	timer        time.Time
	duration     time.Duration
	lastDuration time.Duration
	state        _TimerState
}

// NewTimer creates a new timer.
func NewTimer() *Timer {
	timer := &Timer{
		timer:    time.Now(),
		duration: 0,
		state:    _TimerStateNotStarted,
	}
	return timer
}

// Start the timer.
func (timer *Timer) Start() {
	timer.state = _TimerStateStarted
	timer.timer = time.Now()
}

// Stop timer.
func (timer *Timer) Stop() {
	if timer.state&_TimerStateStarted != 1 {
		return
	}
	timer.state = _TimerStateStopped
	timer.duration += time.Since(timer.timer)
	timer.lastDuration = timer.duration
}

// Continue a stopped timer.
func (timer *Timer) Continue() {
	timer.Start()
}

// Reset a timer.
func (timer *Timer) Reset() {
	timer.duration = time.Duration(0)
	timer.Start()
}

// Elapsed returns the duration from start to end if timer is stopped.
// Elapsed returns the duration from start to now if timer is not stopped.
func (timer *Timer) Elapsed() time.Duration {
	if timer.state == _TimerStateStarted {
		return timer.lastDuration + time.Since(timer.timer)
	}

	return timer.duration
}
