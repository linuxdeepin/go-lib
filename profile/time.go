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

package profile

import (
	"time"
)

// Timer recoreds time cost.
type Timer struct {
	startTime       time.Time
	moduleStartTime time.Time
	endTime         time.Time
}

// NewTimer creates a new timer and start it.
func NewTimer() *Timer {
	t := &Timer{}
	t.startTime = time.Now()
	t.endTime = t.startTime
	return t
}

// Elapsed returns the time duration from the time of start or last elapsed.
func (t *Timer) Elapsed() time.Duration {
	endTime := time.Now()
	sub := endTime.Sub(t.moduleStartTime)
	t.moduleStartTime = endTime
	return sub
}

// TotalCost will stop the timer and return the total cost time.
func (t *Timer) TotalCost() time.Duration {
	if t.startTime == t.endTime {
		t.endTime = time.Now()
	}
	return t.endTime.Sub(t.startTime)
}
