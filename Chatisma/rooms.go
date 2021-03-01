package Chatisma

import (
	"fmt"
	"io"
	"sync"
)

type room struct {
	name           string
	MessageChannel chan string
	clients        map[chan<- string]struct{}
	Quit           chan struct{}
	*sync.RWMutex
}

func CreateRoom(name string) *room {
	r := &room{
		name:           name,
		MessageChannel: make(chan string),
		RWMutex:        new(sync.RWMutex),
		clients:        make(map[chan<- string]struct{}),
		Quit:           make(chan struct{}),
	}

	r.Run()

	return r
}

func (r *room) Run() {
	fmt.Println("Starting chat room", r.name)

	go func() {
		for message := range r.MessageChannel {
			r.broadcastMessage(message)
		}
	}()
}

func (r *room) broadcastMessage(message string) {
	r.RLock()
	defer r.RUnlock()

	fmt.Println("Received message: ", message)

	for wc, _ := range r.clients {
		go func(wc chan<- string) {
			wc <- message
		}(wc)
	}
}

func (r *room) AddClient(client io.ReadWriteCloser) {
	r.Lock()
	wc, done := StartClient(r.MessageChannel, c, r.Quit)
	r.clients[wc] = struct{}{}
	r.Unlock()
}

func (r *room) ClientCount() int {
	return len(r.clients)
}

func (r *room) RemoveClient(wc chan<- string) {
	fmt.Println("Removing client ")
	r.Lock()
	close(wc)
	delete(r.clients, wc)
	r.Unlock()

	select {
	case <-r.Quit:
		if len(r.clients) == 0 {
			close(r.MessageChannel)
		}
	default:
	}
}
