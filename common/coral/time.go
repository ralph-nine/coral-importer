package coral

import (
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"
)

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"$date": t.Time,
	})
}

func (t *Time) UnmarshalJSON(data []byte) error {
	input := map[string]time.Time{}
	if err := json.Unmarshal(data, &input); err != nil {
		logrus.WithError(err).Warn("could not parse time")
		return nil
	}

	t.Time = input["$date"]

	return nil
}
