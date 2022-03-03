package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	s := deferChangeNamedReturnValue()
	fmt.Println(s)
	// IMPORTANT to defer before panic!!!
	defer func() { fmt.Println("Defer in main") }()

	err := writeToTempFile("Some text")
	if err != nil {
		log.Panicf(err.Error())
	}
	fmt.Printf("Write to file ok.")

}

func writeToTempFile(text string) error {
	defer func() { fmt.Println("Defer in writeToTempFile") }()

	file, err := os.Open("temp.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	n, err := file.WriteString("Some text")
	if err != nil {
		return err
	}
	fmt.Printf("Num of bytes written: %d", n)
	// file.Close()
	return nil
}

func deferChangeNamedReturnValue() (size int) {
	defer func() { size = 0 }()
	size = 10
	return
}
