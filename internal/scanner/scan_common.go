package scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileData struct {
	Name       string
	Path       string
	CreatedAt  time.Time
	ModifiedAt time.Time
	Signature  fileSignature
	Category   string
}

type fileSignature struct {
	Extension string
	Signature []byte
	Offset    int
}

const (
	CategoryImage      = "Image"
	CategoryVideo      = "Video"
	CategoryAudio      = "Audio"
	CategoryDocument   = "Document"
	CategoryExecutable = "Executable"
	CategoryArchive    = "Archive"
	CategoryCode       = "Code"
	CategoryUnknown    = "Unknown"
)

var knownSignatures = []fileSignature{
	{Extension: "jpg", Signature: []byte{0xFF, 0xD8, 0xFF, 0xE0}},
	{Extension: "jpg", Signature: []byte{0xFF, 0xD8, 0xFF, 0xDB}},
	{Extension: "jpg", Signature: []byte{0xFF, 0xD8, 0xFF, 0xEE}},
	{Extension: "jpg", Signature: []byte{0xFF, 0xD8, 0xFF, 0xE1}},
}

func Scan(directoryPath string, hidden bool, ignore_dir bool, look_creation bool) []FileData {
	var files []FileData
	read, err := os.ReadDir(directoryPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(read) == 0 {
		fmt.Println("Empty directory")
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
		var creationDate time.Time = time.Time{}
		if look_creation == true {
			creationDate, err = obtainCreationDate(full_path)
			if err != nil {
				creationDate = time.Time{}
			}
		}
		item := FileData{
			Name:       element.Name(),
			Path:       full_path,
			ModifiedAt: info.ModTime(),
			Signature:  sig,
			CreatedAt:  creationDate,
			Category:   obtainCategory(element.Name()),
		}
		files = append(files, item)
	}
	return files
}

func obtainCategory(file string) string {
	ext := strings.ToLower(filepath.Ext(file))
	//CATEGORY MAP
	categoryMap := map[string]string{
		//IMAGE
		".jpg": CategoryImage, ".jpeg": CategoryImage, ".png": CategoryImage, ".gif": CategoryImage,
		".svg": CategoryImage, ".webp": CategoryImage, ".bmp": CategoryImage,
		//VIDEO
		".mp4": CategoryVideo, ".avi": CategoryVideo, ".mkv": CategoryVideo, ".mov": CategoryVideo,
		".wmv": CategoryVideo, ".flv": CategoryVideo,
		//AUDIO
		".mp3": CategoryAudio, ".wav": CategoryAudio, ".flac": CategoryAudio, ".aac": CategoryAudio,
		".ogg": CategoryAudio,
		//DOCUMENTS
		".pdf": CategoryDocument, ".doc": CategoryDocument, ".docx": CategoryDocument, ".txt": CategoryDocument,
		".rtf": CategoryDocument, ".pptx": CategoryDocument, ".odt": CategoryDocument, ".xlsx": CategoryDocument,
		//EXECUTABLES
		".exe": CategoryExecutable, ".msi": CategoryExecutable, ".bat": CategoryExecutable,
		".sh": CategoryExecutable, ".app": CategoryExecutable,
		//ARCHIVES
		".tar": CategoryArchive, ".zip": CategoryArchive, ".rar": CategoryArchive, ".7z": CategoryArchive,
		".gz": CategoryArchive,
		//CODE
		".go": CategoryCode, ".py": CategoryCode, ".js": CategoryCode, ".java": CategoryCode,
		".cpp": CategoryCode, ".c": CategoryCode, ".ts": CategoryCode, ".rb": CategoryCode,
	}
	if category, exists := categoryMap[ext]; exists {
		return category
	}
	return "Unknown"
}
