//go:build windows

package scanner

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func obtainCreationDate(file string) (time.Time, error) {
	info, err := os.Stat(file)
	if err != nil {
		fmt.Printf("Failed to obtain %v creation date\n", file)
		return time.Time{}, err
	}
	sysData := info.Sys().(*syscall.Win32FileAttributeData)
	creation := time.Unix(0, sysData.CreationTime.Nanoseconds())
	return creation, nil
}

func obtainPerms(info os.FileInfo) bool {
	// TEMPORARY CONTINUE, I NEED TO CHECK
	// HOW DOES GO HANDLE FILE PERMS
	return true
}
