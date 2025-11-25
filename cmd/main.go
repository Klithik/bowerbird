package main

import (
	"flag"
	"fmt"

	"github.com/Klithik/bowerbird/internal/manipulator"
	"github.com/Klithik/bowerbird/internal/scanner"
)

func main() {
	var strat, source_dir, end_dir string
	var ignore_hidden bool

	const stratDoc string = "What should the folder be sorted by: the possible values are:\n-year: in order to place all valid files inside a year of creation folder\n-type: in order to place all valid files inside a folder based on its filetype (it tries to determine the file type based on magic bytes, if it fails, it uses the file extension)"
	flag.StringVar(&strat, "strategy", "year", stratDoc)

	const source_dirDoc string = "The directory to be sorted, by default it uses the current directory"
	flag.StringVar(&source_dir, "source", ".", source_dirDoc)

	const end_dirDoc string = "The directory where the folders to be created should be placed, by default it uses the current directory"
	flag.StringVar(&end_dir, "end", ".", end_dirDoc)

	const ignoreHiddenDoc string = "Ignores hidden files (files starting with a '.' character)"
	flag.BoolVar(&ignore_hidden, "ignore_hidden", true, ignoreHiddenDoc)

	flag.Parse()

	files := scanner.Scan(source_dir, ignore_hidden, true, false, false, false)
	manipulator.MoveFiles(files, end_dir)
}
