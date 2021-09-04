package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client struct {
	msgchn chan<- string // an outgoing message channel
	id     string        // client's id
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.msgchn <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			cli.msgchn <- "Active users:"
			for c := range clients {
				cli.msgchn <- c.id
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.msgchn)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages

	// write to current client
	go clientWriter(conn, ch)

	// initialize a client struct and set the id&msg fields

	// get the connected clients ID
	who := conn.RemoteAddr().String()
	cl := client{ch, who}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cl

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// NOTE ignoring potential errors from input.Err()

	leaving <- cl
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
