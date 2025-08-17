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
	messagesChan := getLinesChannel(msgFile)
	for msg := range messagesChan {
		fmt.Printf("read: %s\n", msg)
	}

}

func getLinesChannel(f io.ReadCloser) <-chan string {
	var currentLine string
	var stream = make(chan string)
	go func() {
		for {
			txt := make([]byte, 8)
			_, err := f.Read(txt)
			if err == io.EOF {
				close(stream)
				break
			}
			ok, k := checkAndFindIfNewLineInText(txt)
			if !ok {
				currentLine = currentLine + string(txt)
			} else {
				currentLine = currentLine + string(txt[:k])
				stream <- currentLine
				currentLine = string(txt[k+1:])
			}
		}
	}()
	return stream
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
