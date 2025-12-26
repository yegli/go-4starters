package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	address := "127.0.0.1:8081"
	ln, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			return
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	go readFromConn(conn)
	writeToConn(conn)
}

func readFromConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Print(msg)
	}
}

func writeToConn(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		conn.Write([]byte("> " + text))
	}
}
