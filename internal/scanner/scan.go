package scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type FileData struct {
	Name       string
	Path       string
	CreatedAt  time.Time
	ModifiedAt time.Time
	FileType   string
}

func Scan(directoryPath string, hidden bool, ignore_files []string, ignore_dir bool, get_creation bool, get_modified bool, get_type bool) []FileData {
	var files []FileData
	read, err := os.ReadDir(directoryPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, element := range read {
		if element.IsDir() == true && ignore_dir == true {
			continue
		}
		if element.Name()[0] == '.' && hidden == true {
			continue
		}
		full_path := filepath.Join(directoryPath, element.Name())
		info, _ := os.Stat(full_path)
		item := FileData{
			Name:       element.Name(),
			Path:       full_path,
			ModifiedAt: info.ModTime(),
		}
		files = append(files, item)
	}
	return files
}
