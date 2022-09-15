// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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
