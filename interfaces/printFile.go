package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	printFile()
}

func printFile() {
	// read filename from CLI
	fmt.Printf("os.Args: %s\n", os.Args)
	filename := os.Args[1]
	fmt.Printf("Reading file: %s\n", filename)

	rf := readFile{filename}

	// final piping bytes from File to Stdout
	io.Copy(os.Stdout, rf)
}

type readFile struct {
	filename string
}

func (rf readFile) Read(bs []byte) (int, error) {
	file, err := os.Open(rf.filename)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}

	size, err := file.Read(bs)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Size of file: %d\n", size)

	return len(bs), nil
}
