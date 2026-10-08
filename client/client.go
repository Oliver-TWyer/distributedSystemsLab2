package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	//create a reader that reads from the connection
	reader := bufio.NewReader(conn)

	for {
		msg, _ := reader.ReadString('\n')
		fmt.Printf(msg)
	}
}

func write(conn net.Conn, terminate chan bool) {
	//TODO Continually get input from the user and send messages to the server.

	reader := bufio.NewReader(os.Stdin)
	for {
		//read keyboard messages
		fmt.Println("Send a message->")
		msg, _ := reader.ReadString('\n')

		if msg == "exit" {
			fmt.Fprintf(conn, "endchat")
			terminate <- true
		}
		fmt.Fprintf(conn, msg)
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	//We create a flag that is called ip, and we pass the default value of localhost:8030 unless otherwise
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	//TODO Try to connect to the server
	//we need to get the string where addrPtr is pointing too, and dial the connection
	conn, _ := net.Dial("tcp", *addrPtr)

	terminate := make(chan bool)

	//TODO Start asynchronously reading and displaying messages
	go read(conn)
	//TODO Start getting and sending user messages
	go write(conn, terminate)

	select {
	case <-terminate:
		fmt.Println("Exiting the chat...")
	}
}
