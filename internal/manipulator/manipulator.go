package manipulator

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/Klithik/bowerbird/internal/scanner"
)

func MoveFiles(files []scanner.FileData, target string, year bool, month bool, category bool, datePrio bool) [][2]string {
	var output [][2]string
	for _, element := range files {
		targetPath := target
		if datePrio == true {
			if month == true {
				fileMonth := element.ModifiedAt.Month().String()
				fileYear := strconv.Itoa(element.ModifiedAt.Year())
				targetPath = filepath.Join(target, fileYear)
				targetPath = filepath.Join(target, fileMonth)
			} else if year == true {
				fileYear := strconv.Itoa(element.ModifiedAt.Year())
				targetPath = filepath.Join(target, fileYear)
			}
			if category == true {
				targetPath = filepath.Join(target, element.Category)
			}
		} else {
			if category == true {
				targetPath = filepath.Join(target, element.Category)
			}
			if month == true {
				fileMonth := element.ModifiedAt.Month().String()
				fileYear := strconv.Itoa(element.ModifiedAt.Year())
				targetPath = filepath.Join(target, fileYear)
				targetPath = filepath.Join(target, fileMonth)
			} else if year == true {
				fileYear := strconv.Itoa(element.ModifiedAt.Year())
				targetPath = filepath.Join(target, fileYear)
			}
		}
		_, err := os.Stat(targetPath)
		if os.IsNotExist(err) {
			info, _ := os.Stat(target)
			os.MkdirAll(targetPath, info.Mode().Perm())
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
