package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Klithik/bowerbird/internal/manipulator"
	"github.com/Klithik/bowerbird/internal/scanner"
)

func main() {
	var source_dir, end_dir string
	var category, ignore_hidden, year, month, datePrio bool

	const categoryDoc string = "The elements should be sorted based on file type (extension)"
	flag.BoolVar(&category, "type", true, categoryDoc)

	const source_dirDoc string = "The directory to be sorted, by default it uses the current directory"
	flag.StringVar(&source_dir, "source", ".", source_dirDoc)

	const end_dirDoc string = "The directory where the folders to be created should be placed, by default it uses the current directory"
	flag.StringVar(&end_dir, "end", ".", end_dirDoc)

	const ignoreHiddenDoc string = "Ignores hidden files (files starting with a '.' character)"
	flag.BoolVar(&ignore_hidden, "ignore_hidden", true, ignoreHiddenDoc)

	const yearDoc string = "Sort elements by last modification year"
	flag.BoolVar(&year, "year", false, yearDoc)

	const monthDoc string = "Sort elements by last modification month. If this flag is enabled, year is also used"
	flag.BoolVar(&month, "month", false, monthDoc)

	const datePrioDoc string = "Define if the top directory created by bowerbird is file type or date, so if true the directory of a file would be for example, 2025>October>Image, and if false Image>2025>October"

	if month == true {
		year = true
	}

	flag.Parse()

	if dir_verification(source_dir) == false {
		os.Exit(1)
	}
	if dir_verification(end_dir) == false {
		os.Exit(1)
	}

	files := scanner.Scan(source_dir, ignore_hidden, true)
	manipulator.MoveFiles(files, end_dir, year, month, category, datePrio)
}

func dir_verification(input string) bool {
	info, err := os.Stat(input)
	if os.IsNotExist(err) {
		fmt.Println("Path doesnt exist", input)
		return false
	} else if err != nil {
		fmt.Println("Directory could not be verified", input)
		return false
	} else if info.IsDir() == false {
		fmt.Println("Path provided is not a directory")
		return false
	}
	return true
}
