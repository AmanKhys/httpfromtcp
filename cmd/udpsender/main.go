package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	udpEndPoint, err := net.ResolveUDPAddr("", "localhost:42069")
	if err != nil {
		log.Fatalln("udp endpoint not resolved:", err)
	}

	udpConn, err := net.DialUDP("udp", nil, udpEndPoint)
	if err != nil {
		log.Fatalln("udp connection not establised:", err)
	}
	defer udpConn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		str, err := reader.ReadString(byte('\n'))
		if err == io.EOF {
			fmt.Println("successfully read the entire line")
		} else if err != nil {
			log.Println("hehe:", err)
		}
		_, err = udpConn.Write([]byte(str))
		if err != nil {
			log.Println("huhu:", err)
		}
	}

}
