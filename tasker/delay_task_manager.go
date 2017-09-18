/*
 * Copyright (C) 2017 ~ 2017 Deepin Technology Co., Ltd.
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

package tasker

import (
	"fmt"
	"sync"
	"time"
)

type DelayTaskManager struct {
	taskMap map[string]*DelayTask
	locker  sync.Mutex
}

func NewDelayTaskManager() *DelayTaskManager {
	return &DelayTaskManager{
		taskMap: make(map[string]*DelayTask),
	}
}

func (m *DelayTaskManager) AddTask(name string, duration time.Duration, cb interface{}) error {
	m.locker.Lock()
	defer m.locker.Unlock()
	if _, ok := m.taskMap[name]; ok {
		return fmt.Errorf("Task '%s' has exists", name)
	}

	task, err := NewDelayTask(duration, cb)
	if err != nil {
		return err
	}
	m.taskMap[name] = task
	return nil
}

func (m *DelayTaskManager) GetTask(name string) (*DelayTask, error) {
	m.locker.Lock()
	defer m.locker.Unlock()
	task, ok := m.taskMap[name]
	if !ok {
		return nil, fmt.Errorf("No task '%s' exists", name)
	}
	return task, nil
}

func (m *DelayTaskManager) DeleteTask(name string) {
	m.locker.Lock()
	defer m.locker.Unlock()
	task, ok := m.taskMap[name]
	if !ok {
		return
	}
	task.Stop()
	task = nil
	delete(m.taskMap, name)
}

func (m *DelayTaskManager) Destroy() {
	m.locker.Lock()
	defer m.locker.Unlock()
	for _, task := range m.taskMap {
		task.Stop()
		task = nil
	}
	m.taskMap = nil
}
