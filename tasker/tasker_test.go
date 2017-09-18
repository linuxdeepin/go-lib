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
	"testing"
	"time"
)

func handleTask(id int, name string) {
	fmt.Printf("%v: %s\n", id, name)
}

func TestDelayTask(t *testing.T) {
	task, err := NewDelayTask(time.Second*1, nil)
	if err == nil {
		panic("Failed: should error, because if invalid callback function")
	}

	task, err = NewDelayTask(time.Second*1, handleTask)
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
