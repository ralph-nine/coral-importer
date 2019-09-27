package coral

import (
	"encoding/json"
	"time"
)

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"$date": t.Time,
	})
}
