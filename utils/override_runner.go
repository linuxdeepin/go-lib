/**
 * Copyright (c) 2014 Deepin, Inc.
 *               2014 Xu FaSheng
 *
 * Author:      Xu FaSheng <fasheng.xu@gmail.com>
 * Maintainer:  Xu FaSheng <fasheng.xu@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package utils

import (
	"sync"
)

// OverrideRunner could run a series of tasks, and if possible, the
// later tasks will cancel the previous tasks.
type OverrideRunner struct {
	wg             sync.WaitGroup
	done           chan bool
	channTaskGroup chan []func()
	lastTaskDone   chan bool
}

func NewOverrideRunner() (r *OverrideRunner) {
	r = &OverrideRunner{}
	r.done = make(chan bool)
	r.channTaskGroup = make(chan []func())
	r.lastTaskDone = make(chan bool)
	return
}

// AddTaskGroup add a series of task as group.
func (r *OverrideRunner) AddTaskGroup(taskGroup ...func()) {
	go func() {
		r.channTaskGroup <- taskGroup
	}()
}

// Loop start the loop to run tasks.
func (r *OverrideRunner) Loop() {
	for {
		select {
		case taskGroup := <-r.channTaskGroup:
			close(r.lastTaskDone)            // cancel last task group
			r.wg.Wait()                      // wait for last task finished
			r.lastTaskDone = make(chan bool) // reset channel
			r.run(r.lastTaskDone, taskGroup...)
		case <-r.done:
			return
		}
	}
}

func (r *OverrideRunner) run(done chan bool, taskGroup ...func()) {
	r.wg.Add(1)
	runTask := func(cb func()) (continueNextTask bool) {
		select {
		case <-done:
			continueNextTask = false
		default:
			cb()
			continueNextTask = true
		}
		return
	}
	go func() {
		var continueNextTask bool
		for _, task := range taskGroup {
			continueNextTask = runTask(task)
			if !continueNextTask {
				// this task group was override
				break
			}
		}
		r.wg.Done()
	}()
}

// Wait wait for the end of all tasks.
func (r *OverrideRunner) Wait() {
	r.wg.Wait()
}
