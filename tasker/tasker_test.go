// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package tasker

import (
	"fmt"
	"testing"
	"time"
)

func handleTask(id int, name string) {
	fmt.Printf("%v: %s\n", id, name)
}

func TestDelayTask(t *testing.T) {
	_, err := NewDelayTask(time.Second*1, nil)
	if err == nil {
		panic("Failed: should error, because if invalid callback function")
	}

	task, err := NewDelayTask(time.Second*1, handleTask)
	if err != nil {
		panic("Failed: should no error")
	}

	err = task.Start(1, "First")
	if err != nil {
		panic("Failed: should no error")
	}
	time.Sleep(time.Second * 2)
	err = task.Start(2, "Second")
	if err != nil {
		panic("Failed: should no error")
	}
	time.Sleep(time.Millisecond * 500)
	err = task.Start(3, "Third")
	if err != nil {
		panic("Failed: should no error")
	}
	time.Sleep(time.Millisecond * 1200)
}
