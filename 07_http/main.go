package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Read function doesn't resize the byte slice
	/*
		resp.Body.Read(bs)
		fmt.Println(string(bs))
		bs := make([]byte, 99999)
	*/
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Wrote this many bytes", len(bs))
	return len(bs), nil
}
