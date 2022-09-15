// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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
