package pulse

import "C"
import "fmt"
import "unsafe"

var pendingCallback = make(chan func(), 10)

func startHandleCallbacks() {
	for fn := range pendingCallback {
		if fn != nil {
			fn()
		}
	}
}

func init() {
	// TODO: encapsulate the logic to Context and adding Start/Stop logic.
	go startHandleCallbacks()
}

// ALL of below functions are invoked from pulse's mainloop thread.
// So
//   1. Don't hold go-runtime lock in current thread.
//   2. Don't hold pa_threaded_mainloop in current thread.
//
// NOTE:
// In pa_threaded_mainloop, _prepare_ and _dispatch_ is running
// under lock. Only phrase _poll_ is release the lock. So callback
// must not try to hold lock.

//export go_handle_changed
func go_handle_changed(facility int, event_type int, idx uint32) {
	pendingCallback <- func() {
		cbs, ok := GetContext().cbs[facility]
		if !ok {
			fmt.Println("unknow event", facility, event_type, idx)
			return
		}

		for _, cb := range cbs {
			cb(event_type, idx)
		}
	}
}

//export go_handle_state_changed
func go_handle_state_changed(state int) {
	pendingCallback <- func() {
		cbs, ok := GetContext().stateCbs[state]
		if !ok {
			fmt.Println("Unregiste state:", state)
		}

		for _, cb := range cbs {
			cb()
		}
	}
}

//export go_update_volume_meter
func go_update_volume_meter(source_index uint32, sink_index uint32, v float64) {
	pendingCallback <- func() {
		sourceMeterLock.RLock()
		cb, ok := sourceMeterCBs[source_index]
		sourceMeterLock.RUnlock()

		if ok && cb != nil {
			cb(v)
		}
	}
}

//export go_receive_some_info
func go_receive_some_info(cookie int64, infoType int, info unsafe.Pointer, status int) {
	paInfo := NewPaInfo(info, infoType)

	pendingCallback <- func() {
		c := fetchCookie(cookie)
		if c == nil {
			fmt.Println("Warning: recieve_some_info with nil cookie", cookie, infoType, info, status)
			return
		}

		switch {
		case status == 1:
			c.EndOfList()
		case status == 0:
			c.Feed(paInfo)
		case status < 0:
			c.Failed()
		}
	}
}
