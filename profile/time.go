/**
 * Copyright (C) 2015 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

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
