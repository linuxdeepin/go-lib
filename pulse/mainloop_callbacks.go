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

type safeDoCtx struct {
	fn   func()
	loop *C.pa_threaded_mainloop
}

var pendingSafeDo = make(chan safeDoCtx)

func startHandleCallbacks() {
	for fnWithOutPendingCode := range pendingCallback {
		if fnWithOutPendingCode != nil {
			fnWithOutPendingCode()
		}
	}
}

func startHandleSafeDo() {
	// Move all safeDo fn to here to avoid creating too many OS Thread.
	// Because GOMAXPROC only apply to go-runtime.
	for c := range pendingSafeDo {
		if c.fn != nil {
			runtime.LockOSThread()
			C.pa_threaded_mainloop_lock(c.loop)
			c.fn()
			C.pa_threaded_mainloop_unlock(c.loop)
			runtime.UnlockOSThread()
		}
	}
}

func init() {
	// TODO: encapsulate the logic to Context and adding Start/Stop logic.
	go startHandleCallbacks()
	go startHandleSafeDo()
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
	pendingSafeDo <- safeDoCtx{fn, c.loop}
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
		ctx := GetContext()
		ctx.mu.Lock()
		cbs, ok := ctx.cbs[facility]
		ctx.mu.Unlock()
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
		ctx := GetContext()
		ctx.mu.Lock()
		cbs, ok := ctx.stateCbs[state]
		ctx.mu.Unlock()
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
func go_receive_some_info(cookie int64, infoType int, info unsafe.Pointer, eol int) {
	ck := fetchCookie(cookie)
	if ck == nil {
		if eol == 0 {
			// Sometimes pulseaudio will send eol=1 to pa_context_get_*_by_index/name.
			fmt.Println("Warning: recieve_some_info with nil cookie", cookie, infoType, info)
		}
		return
	}

	switch {
	case eol == 1:
		ck.EndOfList()
	case eol == 0:
		// 1. the Feed will be blocked until the ck creator received.
		// 2. the creator of ck may be invoked under PendingCallback
		// so this must not be moved to PendingCallback.
		paInfo := NewPaInfo(info, infoType)
		ck.Feed(paInfo)
	case eol < 0:
		ck.Failed()
	}
}
