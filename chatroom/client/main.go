package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	startClient("127.0.0.1:8081")
}

func startClient(address string) {

	connectionClient, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}

	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')

		fmt.Fprint(connectionClient, text+"\n")
	}
}
