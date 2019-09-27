package common

import "github.com/pkg/errors"

// Check will conform and validate the input structs.
func Check(s interface{}) error {
	// Conform to ensure we're validated.
	if err := Conform(s); err != nil {
		return errors.Wrap(err, "could not conform")
	}

	// Validate that is correct.
	if err := Validate(s); err != nil {
		return errors.Wrap(err, "could not validate")
	}

	return nil
}
