// Package entropy provides reusable sources of random data.
package entropy

import (
	"crypto/rand"
	"errors"
)

// Bytes returns length cryptographically secure random bytes.
func Bytes(length int) ([]byte, error) {
	if length < 0 {
		return nil, errors.New("entropy: length must not be negative")
	}

	data := make([]byte, length)
	if _, err := rand.Read(data); err != nil {
		return nil, err
	}
	return data, nil
}
