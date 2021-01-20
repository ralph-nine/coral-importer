package coral

import (
	"sync"
	"time"
)

var createdAtCounter struct {
	enabled bool
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
	if createdAtCounter.enabled {
		// We increment the milliseconds value for each unique time.
		createdAtCounter.mutex.Lock()

		var counter time.Duration

		createdAtCounter.counter += 1 * time.Millisecond
		counter = createdAtCounter.counter

		createdAtCounter.mutex.Unlock()

		return Time{
			Time: time.Now().Add(counter),
		}
	}

	return Time{
		Time: time.Now(),
	}
}
