package coral

import "encoding/json"

type HTML string

func (html HTML) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(html))
}
