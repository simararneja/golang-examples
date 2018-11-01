package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		io.Copy(os.Stdout, f)
	}

}
