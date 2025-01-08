package cli

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrNoInputFile  = errors.New("no input file provided")
	ErrFileNotExist = errors.New("file does not exist")
	ErrNotAFile     = errors.New("not a file")
	ErrFileOpen     = errors.New("error opening file")
)

const (
	Version = "0.0.3"
	Author  = "elliot40404<avishek40404@gmail.com>"
	Name    = "modo"
	Desc    = "Modo is a simple cli app that lets you interact any text file that contains checkboxes in the markdown format."
	Example = "modo <file>"
)

func help() {
	fmt.Printf(`%s %s

%s

Usage:
  %s %s

Options:
  -h, --help    Show this help message
  -v, --version Show version information

Examples:
  %s
`, Name, Version, Desc, Name, Example, Example)
}

func ParseArgs() (*os.File, error) {
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}
	arg := os.Args[1]
	if arg == "-h" || arg == "--help" {
		help()
		os.Exit(1)
	}
	if arg == "-v" || arg == "--version" {
		fmt.Printf("%s %s\n", Name, Version)
		os.Exit(1)
	}
	fi, err := os.Stat(arg)
	if err != nil {
		return nil, ErrFileNotExist
	}
	if fi.IsDir() {
		return nil, ErrNotAFile
	}
	f, err := os.OpenFile(arg, os.O_RDWR, 0o644)
	if err != nil {
		return nil, ErrFileOpen
	}
	return f, nil
}
