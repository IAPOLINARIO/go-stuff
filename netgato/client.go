package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	listen = flag.Bool("l", false, "Listen")
	host   = flag.String("h", "localhost", "Host")
	// Default 0 port means to choose random one
	port = flag.Int("p", 0, "Port")
)

func main() {
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Hostname and port required")
		return
	}

	serverHost := flag.Arg(0)
	serverPort := flag.Arg(1)

	startClient(fmt.Sprintf("%s:%s", serverHost, serverPort))
}

func startClient(addr string) {
	fmt.Printf("Connecting to: %v...\n", addr)
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		fmt.Printf("Can't connect to server: %s\n", err)
		return
	}

	for {
		fmt.Print("> ")

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		conn.Write([]byte(fmt.Sprintf("%s\r\n", text)))
		b := make([]byte, 512)

		conn.Read(b)

		fmt.Println(string(b))
	}

}
