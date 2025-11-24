package scanner

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type FileData struct {
	Name       string
	Path       string
	CreatedAt  time.Time
	ModifiedAt time.Time
	Signature  fileSignature
}

type fileSignature struct {
	Extension string
	Signature []byte
	Offset    int
}

var knownSignatures = []fileSignature{
	{Extension: "jpg", Signature: []byte{0xFF, 0xD8, 0xFF, 0xE0}},
	{Extension: "jpg", Signature: []byte{0xFF, 0xD8, 0xFF, 0xDB}},
	{Extension: "jpg", Signature: []byte{0xFF, 0xD8, 0xFF, 0xEE}},
	{Extension: "jpg", Signature: []byte{0xFF, 0xD8, 0xFF, 0xE1}},
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
		sig := fileSignature{
			Extension: filepath.Ext(full_path),
		}
		item := FileData{
			Name:       element.Name(),
			Path:       full_path,
			ModifiedAt: info.ModTime(),
			Signature:  sig,
		}
		files = append(files, item)
	}
	return files
}

func magicType(filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file %v\n", err)
	}
	defer f.Close()
	magicBytes := make([]byte, 4)
	_, err = io.ReadFull(f, magicBytes)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file %v\n", err)
	}
	return "temp"
}
