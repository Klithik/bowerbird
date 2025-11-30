//go:build windows

package scanner

import (
	"time"
)

func obtainCreationDate(file string) (time.Time, error) {
	return time.Time{}, nil
}
