package main

import (
	"fmt"
	"net"
)

func main() {
	// listen on default redis TCP port
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

	//start infinite loop, to read bytes from connection
	for {
		resp := NewResp(connection)
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		_ = value

		writer := NewWriter(connection)
		writer.Write(Value{typ: "string", str: "OK"})
	}
}
