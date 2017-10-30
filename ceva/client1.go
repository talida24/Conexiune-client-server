package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
)

func main() {

	// connect to this socket
	conn1, _ := net.Dial("tcp", "127.0.0.1:8082")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Alegeti un loc(1-10): ")
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn1, text + "\n")

		message, _ := bufio.NewReader(conn1).ReadString('\n')
		fmt.Print("Mesajul primit de la server: "+message)
	}
}
