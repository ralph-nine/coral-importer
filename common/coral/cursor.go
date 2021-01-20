package coral

import (
	"sync"
	"time"
)

var createdAtCounter struct {
	enabled bool
	last    int64
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
		unix := now.Unix()

		// We increment the milliseconds value for each unique time.
		createdAtCounter.mutex.Lock()

		var counter time.Duration

		// If the unix timestamp is the same as the last one, then we need to
		// increment the counter. If it's not, then we need to reset the counter.
		if createdAtCounter.last == unix {
			createdAtCounter.counter += 1 * time.Millisecond
			counter = createdAtCounter.counter
		} else {
			createdAtCounter.last = unix
			createdAtCounter.counter = 0
			counter = 0
		}

		createdAtCounter.mutex.Unlock()

		return Time{
			Time: time.Now().Add(counter),
		}
	}

	return Time{
		Time: time.Now(),
	}
}
