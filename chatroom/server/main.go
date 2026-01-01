package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

var (
	conns []net.Conn
	mu    sync.Mutex // required to prevent concurrent slice access
)

func main() {
	address := "127.0.0.1:8081"
	ln, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	go writeToConn()

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		mu.Lock()
		conns = append(conns, conn)
		mu.Unlock()

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	readFromConn(conn)
	mu.Lock()
	for i, c := range conns {
		if c == conn {
			conns = append(conns[:i], conns[i+1:]...)
			break
		}
	}
	mu.Unlock()
	conn.Close()
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

func writeToConn() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		mu.Lock()
		for _, c := range conns {
			c.Write([]byte("> " + text))
		}
		mu.Unlock()
	}
}
