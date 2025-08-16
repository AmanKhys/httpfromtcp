package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	msgFile, err := os.Open("messages.txt")
	if err != nil {
		panic(err)
	}
	for {
		txt := make([]byte, 8)
		_, err = msgFile.Read(txt)
		if err == io.EOF {
			break
		}
		fmt.Printf("read: %s\n", txt)
	}
}
