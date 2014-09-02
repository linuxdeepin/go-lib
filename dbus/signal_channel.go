package dbus

type signalChannel struct {
	input, output chan *Signal
	caches        []*Signal
}

func newSignalChannel() *signalChannel {
	ch := &signalChannel{
		input:  make(chan *Signal),
		output: make(chan *Signal),
		caches: make([]*Signal, 0, 30),
	}
	go ch.run()
	return ch
}

func (ch *signalChannel) In() chan<- *Signal {
	return ch.input
}

func (ch *signalChannel) Out() <-chan *Signal {
	return ch.output
}

func (ch *signalChannel) Close() {
	close(ch.input)
}

//add/remove could be improved by allocate vector capacatiy
func (ch *signalChannel) add(s *Signal) {
	ch.caches = append(ch.caches, s)
}
func (ch *signalChannel) remove() {
	ch.caches = ch.caches[1:]
}
func (ch *signalChannel) peek() *Signal {
	return ch.caches[0]
}

func (ch *signalChannel) shutdown() {
	for len(ch.caches) > 0 {
		ch.output <- ch.peek()
		ch.remove()
	}
	close(ch.output)
}

func (ch *signalChannel) run() {
	for {
		if len(ch.caches) == 0 {
			elem, open := <-ch.input
			if open {
				ch.add(elem)
			} else {
				ch.shutdown()
				return
			}
		} else {
			select {
			case elem, open := <-ch.input:
				if open {
					ch.add(elem)
				} else {
					ch.shutdown()
					return
				}
			case ch.output <- ch.peek():
				ch.remove()
			}
		}
	}
}

func (conn *Conn) sendSignals(signal *Signal) {
	conn.signalsLck.Lock()
	for _, ch := range conn.signals {
		ch.In() <- signal
	}
	conn.signalsLck.Unlock()
}

// Signal registers the given channel to be passed all received signal messages.
// The caller has to make sure that ch is sufficiently buffered; if a message
// arrives when a write to c is not possible, it is discarded.
//
// Multiple of these channels can be registered at the same time. Passing a
// channel that already is registered will remove it from the list of the
// registered channels.
//
// These channels are "overwritten" by Eavesdrop; i.e., if there currently is a
// channel for eavesdropped messages, this channel receives all signals, and
// none of the channels passed to Signal will receive any signals.
func (conn *Conn) Signal() <-chan *Signal {
	conn.signalsLck.Lock()
	ch := newSignalChannel()
	conn.signals[ch.Out()] = ch
	conn.signalsLck.Unlock()
	return ch.Out()
}

func (conn *Conn) DetachSignal(removingChan <-chan *Signal) {
	conn.signalsLck.Lock()
	if v, ok := conn.signals[removingChan]; ok {
		delete(conn.signals, removingChan)
		v.Close()
	}
	conn.signalsLck.Unlock()
}
