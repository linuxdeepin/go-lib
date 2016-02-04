/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package dbus

import "fmt"

func ExampleConn_Emit() {
	conn, err := SessionBus()
	if err != nil {
		panic(err)
	}

	conn.Emit("/foo/bar", "foo.bar.Baz", uint32(0xDAEDBEEF))
}

func ExampleObject_Call() {
	var list []string

	conn, err := SessionBus()
	if err != nil {
		panic(err)
	}

	err = conn.BusObject().Call("org.freedesktop.DBus.ListNames", 0).Store(&list)
	if err != nil {
		panic(err)
	}
	for _, v := range list {
		fmt.Println(v)
	}
}

func ExampleObject_Go() {
	conn, err := SessionBus()
	if err != nil {
		panic(err)
	}

	ch := make(chan *Call, 10)
	conn.BusObject().Go("org.freedesktop.DBus.ListActivatableNames", 0, ch)
	select {
	case call := <-ch:
		if call.Err != nil {
			panic(err)
		}
		list := call.Body[0].([]string)
		for _, v := range list {
			fmt.Println(v)
		}
		// put some other cases here
	}
}
