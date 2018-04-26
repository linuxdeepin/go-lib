package pulse

//#include "dde-pulse.h"
import "C"
import "runtime"
import "fmt"
import "unsafe"

// ONLY C file and below function can use C.pa_threaded_mainloop_lock/unlock

// It Must not be send any fn to pendingCallback in the callback,
// otherwise DEADLOCK.
// This means at least don't run any unknown code which passed from user.
var pendingCallback = make(chan func())

func startHandleCallbacks() {
	for fnWithOutPendingCode := range pendingCallback {
		if fnWithOutPendingCode != nil {
			fnWithOutPendingCode()
		}
	}
}

func init() {
	// TODO: encapsulate the logic to Context and adding Start/Stop logic.
	go startHandleCallbacks()
}

func freeContext(ctx *Context) {
	runtime.LockOSThread()
	C.pa_threaded_mainloop_lock(ctx.loop)
	// The pa_context_unref must be protected.
	// This operation will cancel all pending operations which will touch  mainloop object.
	C.pa_context_unref(ctx.ctx)
	C.pa_threaded_mainloop_unlock(ctx.loop)
	// There no need to call pa_threaded_mainloop_stop.
	C.pa_threaded_mainloop_free(ctx.loop)
	runtime.UnlockOSThread()

	ctx.loop = nil
	ctx.ctx = nil
}

// safeDo invoke an function with lock
func (c *Context) safeDo(fn func()) {
	runtime.LockOSThread()
	C.pa_threaded_mainloop_lock(c.loop)
	// NOTE: fn() can't hold any lock except the c.loop,
	//  Otherwise there will be a deadlock situation.
	// See also mainloop_callbacks.go
	fn()
	C.pa_threaded_mainloop_unlock(c.loop)
	runtime.UnlockOSThread()
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
			go cb(event_type, idx)
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
			go cb()
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
			go cb(v)
		}
	}
}

//export go_receive_some_info
func go_receive_some_info(cookie int64, infoType int, info unsafe.Pointer, status int) {
	paInfo := NewPaInfo(info, infoType)

	ck := fetchCookie(cookie)
	if ck == nil {
		fmt.Println("Warning: recieve_some_info with nil cookie", cookie, infoType, info, status)
		return
	}

	switch {
	case status == 1:
		ck.EndOfList()
	case status == 0:
		// 1. the Feed will be blocked until the ck creator received.
		// 2. the creator of ck may be invoked under PendingCallback
		// so this must not be moved to PendingCallback.
		ck.Feed(paInfo)
	case status < 0:
		ck.Failed()
	}
}
