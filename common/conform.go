package common

import "github.com/leebenson/conform"

// Conform will ensure strings are trimmed and sanitized.
func Conform(s interface{}) error {
	return conform.Strings(s)
}
