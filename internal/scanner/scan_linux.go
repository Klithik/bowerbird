//go:build linux

package scanner

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"time"
)

// Obtains the creation date of a unix file, if the birth time is supported but not defined in a particular file, it returns the last modification date.
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

// Checks the current user's permissions regarding a particular path.
// If the user has write and read access, returns true
func obtainPerms(info os.FileInfo) bool {
	perms := info.Mode().Perm()
	if perms&0b110000000 == 0b110000000 {
		return true
	}
	return false
}
