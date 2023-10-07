package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var StatusOK = "HTTP/1.1 200 OK\r\n\r\nContent-Type: text/plain\r\n\r\nContent-Length: %s\r\n\r\n%s"

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	req := make([]byte, 1024)

	_, err = conn.Read(req)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		os.Exit(1)
	}

	// split the request by spaces
	requestParts := strings.Split(string(req), " ")

	fmt.Println(requestParts)
	// the second part of the request is the path
	path := requestParts[1]

	if path == "/" {
		_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
		if err != nil {
			fmt.Println("Error writing response: ", err.Error())
		}
	} else if path == "/echo" {
		_, err = conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
		if err != nil {
			fmt.Println("Error writing response: ", err.Error())
		}
	} else {
		_, err = conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
		if err != nil {
			fmt.Println("Error writing response: ", err.Error())
		}
	}
}
