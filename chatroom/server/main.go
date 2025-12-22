package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	startServer("127.0.0.1:8081")
}

func startServer(address string) {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	connectionServer, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(connectionServer).ReadString('\n')
		connectionServer.Write([]byte(message + "\n"))

		fmt.Print(message)
	}
}
