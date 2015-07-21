package operations

import (
	"container/list"
	"pkg.deepin.io/lib/gio-2.0"
	"pkg.deepin.io/lib/timer"
	"sync"
)

// TODO: make a real Enumerator.
type Enumerator struct {
	ch     chan interface{}
	closed bool
}

func NewEnumerator(ch chan interface{}) *Enumerator {
	return &Enumerator{
		ch: ch,
	}
}

func (e *Enumerator) Next() <-chan interface{} {
	return e.ch
}

func (e *Enumerator) IsClosed() bool {
	return e.closed
}

func (e *Enumerator) Close() {
	if !e.closed {
		close(e.ch)
		e.closed = true
	}
}

// ReactorElement holds signal handler and a id for that.
type ReactorElement struct {
	fn interface{}
	id int64
}

func newReactorElement(id int64, fn interface{}) *ReactorElement {
	return &ReactorElement{
		fn: fn,
		id: id,
	}
}

// SignalReactor is a reactor of one signal.
type SignalReactor struct {
	signalName  string
	elements    *list.List // there is no priority, priority queue is not necessary.
	lock        sync.Mutex
	cancellable *gio.Cancellable
}

// NewSignalReactor creates a new SignalReactor.
func NewSignalReactor(signalName string, cancellable *gio.Cancellable) *SignalReactor {
	return &SignalReactor{
		signalName:  signalName,
		elements:    list.New(),
		cancellable: cancellable,
	}
}

func (l *SignalReactor) newDetacher(id int64) func() {
	return func() {
		l.lock.Lock()
		defer l.lock.Unlock()
		iter := l.elements.Front()
		var e *list.Element
		for iter != nil {
			if iter.Value.(ReactorElement).id == id {
				e = iter
				break
			}
		}
		l.elements.Remove(e)
	}
}

// Add adds a new handler to signal.
func (l *SignalReactor) Add(fn interface{}) func() {
	l.lock.Lock()
	defer l.lock.Unlock()

	id := timer.GetMonotonicTime().NanoSeconds()
	l.elements.PushBack(newReactorElement(id, fn))
	return l.newDetacher(id)
}

// Enumerator return a channel of handler.
func (l *SignalReactor) Enumerator() *Enumerator {
	e := NewEnumerator(make(chan interface{}))
	go func() {
		iter := l.elements.Front()
		for iter != nil {
			cancellable := l.cancellable
			if cancellable != nil && cancellable.IsCancelled() {
				break
			}
			if e.IsClosed() {
				break
			}
			listener := iter.Value.(*ReactorElement)
			e.ch <- listener.fn
			iter = iter.Next()
		}
		e.Close()
	}()

	return e
}
