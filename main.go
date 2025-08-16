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
	var currentLine string
	for {
		txt := make([]byte, 8)
		_, err = msgFile.Read(txt)
		if err == io.EOF {
			break
		}
		ok, k := checkAndFindIfNewLineInText(txt)
		if !ok {
			currentLine = currentLine + string(txt)
		} else {
			currentLine = currentLine + string(txt[:k])
			fmt.Printf("read: %s\n", currentLine)
			currentLine = string(txt[k+1:])
		}
	}
}

func checkAndFindIfNewLineInText(txt []byte) (bool, int) {
	var i int
	for ; i < len(txt); i++ {
		if string(txt[i]) == fmt.Sprintf("\n") {
			return true, i
		}
	}
	return false, -1
}
