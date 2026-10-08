package main

import (
	"bufio"
	"fmt"
	"net"
)

//bufio = buffered input and output, instead of reading data one bit at a time, a buffer will
//hold some data so it can be read more efficiently

//takes in a value of type Conn defined in the net package
func handleConnection(conn net.Conn) {
	//reads incoming data from the TCP connection
	reader := bufio.NewReader(conn)

	for {
		msg, _ := reader.ReadString('\n')

		fmt.Println(msg)
		fmt.Fprintln(conn, "OK")
	}
	//reads data until it reaches a newline and stores it in message, then print
}

func main() {
	//create a tcp listener on port 8030, waiting for a client, the _ represents an error variable
	ln, _ := net.Listen("tcp", ":8030")

	for {
		//wait until a client connects, then create a connection, program waits here until connection
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}
}
