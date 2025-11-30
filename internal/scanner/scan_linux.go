//go:build linux || darwin

package scanner

import (
	"fmt"
	"time"

	"golang.org/x/sys/unix"
)

func obtainCreationDate(file string) (time.Time, error) {
	var stat unix.Statx_t
	err := unix.Statx(unix.AT_FDCWD, file, 0, unix.STATX_BTIME, &stat)
	if err != nil {
		return time.Time{}, err
	}
	if stat.Mask&unix.STATX_BTIME == 0 {
		return time.Time{}, fmt.Errorf("birth time not available")
	}
	return time.Unix(stat.Btime.Sec, int64(stat.Btime.Nsec)), nil
}
