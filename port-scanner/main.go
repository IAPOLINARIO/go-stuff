package main

import (
	"fmt"
	"net"
	"sort"
)

const (
	baseAddress = "scanme.nmap.org"
)

func main() {
	//An int value of 100, is provided to make() here
	//This allows the channel to be buffered, which means you can send it an item without waiting for
	//a receiver to read the item. Buffered channels are ideal for maintaining and tracking work for multiple producers and consumers.
	// We’ve capped the channel at 100, meaning it can hold 100 items before the sender will block.
	ports := make(chan int, 100)

	results := make(chan int)
	var openports []int

	//a for loop to start the desired number of workers. In this case, 100
	for i := 0; i < cap(ports); i++ {
		fmt.Printf("Creating worker: %d\n", i)
		go worker(ports, results)
	}

	//need to send to the workers in a separate goroutine
	//because the result-gathering loop needs to start before more than 100 items of work can continue.
	go func() {
		for i := 0; i < 1024; i++ {
			ports <- i
		}
	}()

	//The result-gathering loop receives on the results channel 1024 times.
	//If the port doesn’t equal 0, it’s appended to the slice. After closing the channels
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	//closing the channels
	close(ports)
	close(results)

	//Sort the open ports
	sort.Ints(openports)

	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}

func worker(ports, results chan int) {
	//use range to continuously receive from the ports channel, looping until the channel is closed.
	for p := range ports {
		address := fmt.Sprintf("%v:%d", baseAddress, p)
		fmt.Printf("port:%d\n", p)

		conn, err := net.Dial("tcp", address)

		if err != nil {
			results <- 0
			continue
		}

		conn.Close()
		results <- p
	}
}
