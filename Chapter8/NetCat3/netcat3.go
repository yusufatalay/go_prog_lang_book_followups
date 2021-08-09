package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:8000")
	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, tcpConn) // ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(tcpConn, os.Stdin)
	// Only close the Write part of the tcp connection to print all the current
	// values before terminating the connection.
	tcpConn.CloseWrite()
	<-done // wait for background goroutine to finish
}
