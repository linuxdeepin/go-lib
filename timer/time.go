package timer

// #include <time.h>
import "C"

import (
	"fmt"
)

// TimeSpec holds times.
type TimeSpec struct {
	seconds     int64
	nanoSeconds int64
}

func (ts TimeSpec) String() string {
	return fmt.Sprintf("%d.%d", ts.Seconds, ts.NanoSeconds)
}

// Seconds returns the seconds.
func (ts TimeSpec) Seconds() int64 {
	return ts.seconds
}

// Milliseconds returns the milliseconds
func (ts TimeSpec) Milliseconds() int64 {
	return ts.seconds*1000 + ts.nanoSeconds/1000000
}

// MicroSeconds returns milliseconds
func (ts TimeSpec) MicroSeconds() int64 {
	return ts.seconds*1000000 + ts.nanoSeconds/1000
}

// NanoSeconds returns nanoseconds
func (ts TimeSpec) NanoSeconds() int64 {
	return ts.nanoSeconds
}

// GetMonotonicTime returns the monotonic time, used when the time should be effected by time adjust, like timer.
// according to the sources of golang, the time.Now() is the CLOCK_REALTIME
// which will be affected by discontinuous jumps in the system time.
func GetMonotonicTime() TimeSpec {
	var ts C.struct_timespec

	// see http://man7.org/linux/man-pages/man2/clock_gettime.2.html for a list of CLOCK_ constants
	C.clock_gettime(C.CLOCK_MONOTONIC_RAW, &ts)

	// convert to something easier to manipulate
	return TimeSpec{int64(ts.tv_sec), int64(ts.tv_nsec)}
}
