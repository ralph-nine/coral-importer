package coral

import (
	"sync"
	"time"
)

var createdAtCounter struct {
	enabled bool
	lastMs  time.Duration
	counter time.Duration
	mutex   sync.Mutex
}

// EnableMonotonicCursorTime will ensure that all emitted cursor times emitted
// are unique to the ms.
func EnableMonotonicCursorTime() {
	createdAtCounter.enabled = true
}

// NewCursorTime is used to return time values that are used by pagination
// fields.
func NewCursorTime() Time {
	now := time.Now()

	if createdAtCounter.enabled {
		unixMs := time.Duration(now.UnixNano()) * time.Nanosecond / time.Millisecond

		// We increment the milliseconds value for each unique time.
		createdAtCounter.mutex.Lock()

		var counter time.Duration

		// If the last timestamp we saw (which includes the counter) is greater than
		// or equal to the current timestamp, then we should add one to both. If the
		// current timestamp is less than the last one, we should reset the counter.
		if createdAtCounter.lastMs >= unixMs {
			createdAtCounter.counter += 1 * time.Millisecond
			createdAtCounter.lastMs += 1 * time.Millisecond
			counter = createdAtCounter.counter
		} else {
			createdAtCounter.lastMs = unixMs
			createdAtCounter.counter = 0
			counter = 0
		}

		createdAtCounter.mutex.Unlock()

		return Time{
			Time: now.Add(counter),
		}
	}

	return Time{
		Time: now,
	}
}
