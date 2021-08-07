// Contains Exercise 8.1
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	portFlag := flag.String("port", "8000", "port number for the connection")
	flag.Parse()
	connectionString := fmt.Sprintf("localhost:%s", *portFlag)
	listener, err := net.Listen("tcp", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g. connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g. client disconnected

		}
		time.Sleep(1 * time.Second)
	}
}
