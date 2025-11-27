package manipulator

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/Klithik/bowerbird/internal/scanner"
)

func MoveFilesbyType(files []scanner.FileData, target string) [][2]string {
	var output [][2]string
	for _, element := range files {
		targetPath := filepath.Join(target, element.Category)
		_, err := os.Stat(targetPath)
		if os.IsNotExist(err) {
			info, _ := os.Stat(target)
			os.Mkdir(targetPath, info.Mode().Perm())
		}
		targetPath = filepath.Join(targetPath, element.Name)
		err = os.Rename(element.Path, targetPath)
		var mov_success bool = true
		if err != nil {
			mov_success = false
		}
		result := [2]string{element.Name, strconv.FormatBool(mov_success)}
		output = append(output, result)
	}
	return output
}
