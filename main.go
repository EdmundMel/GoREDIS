package main

import (
	"fmt"
	"net"
	"strings"
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

		if value.typ != "array" {
			fmt.Println("Invalid request, expected array length > 0")
			continue
		}

		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		writer := NewWriter(connection)

		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(Value{typ: "string", str: ""})
			continue
		}

		result := handler(args)
		writer.Write(result)
	}
}
