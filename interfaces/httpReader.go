package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	httpReader()
}

func httpReader() {
	resp, err := http.Get("http://golang.org")
	if err != nil {
		fmt.Printf("ERR GET URL: %v", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// to get body io.ReaderCloser
	io.Copy(lw, resp.Body)
}

type logWriter struct {
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Writting this many bytes %d\n", len(bs))
	return len(bs), nil
}
