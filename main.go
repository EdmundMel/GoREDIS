package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := l.Accept()
	if err != nil {
		fmt.Println(err)
	}

	defer connection.Close()

	for {
		buf := make([]byte, 1024)

		_, err = connection.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}

		// ignore request
		connection.Write([]byte("+OK\r\n"))
	}
}
