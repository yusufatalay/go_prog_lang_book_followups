// Netcat1 is a read-only TCP client.
// modified this program to listen the specified port
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	portFlag := flag.String("port", "8000", "port number for the connection")
	flag.Parse()
	connectionString := fmt.Sprintf("localhost:%s", *portFlag)

	conn, err := net.Dial("tcp", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
