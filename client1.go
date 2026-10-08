package main

import (
	"bufio"
	"os"

	//"flag"
	"fmt"
	"net"
)

func read(conn net.Conn) {
	//reads incoming data from the TCP connection
	reader := bufio.NewReader(conn)

	//reads data until it reaches a newline and stores it in message, then print
	msg, _ := reader.ReadString('\n')

	fmt.Printf(msg)
}

func main() {
	stdin := bufio.NewReader(os.Stdin)

	//create a command-line option called -msg
	//value is the default value if the flag isn't specified
	//and usage is the description/help text
	//flag.String returns a pointer to a string, so need to use *msgP to access the string
	//msgP := flag.String("msg", "Default message", "The message you want to send")

	//use flag.Parse() to actually process the data that the client has entered
	//flag.Parse()

	//create a TCP connection to the server at address 127.0.0.1 (localhost), port 8030
	conn, _ := net.Dial("tcp", "127.0.0.1:8030")

	for {
		//conn, _ := net.Dial("tcp", "127.0.0.1:8030")
		fmt.Printf("Enter text...")
		msg, _ := stdin.ReadString('\n')

		//sends the content of msg to the server
		fmt.Fprintln(conn, msg)
		read(conn)
	}
}
