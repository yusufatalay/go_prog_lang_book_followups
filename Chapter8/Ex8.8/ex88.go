// Also contains exercise 8.8
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	afk := time.NewTimer(10 * time.Second)
	text := make(chan string)
	var wg sync.WaitGroup
	go func() {
		for input.Scan() {
			text <- input.Text()
		}
		close(text)
	}()
	
	for {
		select{
	case t, ok := <-text:
		if ok{
			wg.Add(1)
			timeout.Reset(10*time.Second) // reset the time if the user shouts
			go func(){
				defer wg.Done()
				echo(c,t, 1*time.Second)
			}()
		}else{
			wg.Wait()
			c.Close()
			return
		}
	case <- afk.C:
		afk.Stop()
		c.Close()
		fmt.Println("Connection timed out by afking 10 seconds")
		return 

	}

	c.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // connection aborted
			continue
		}
		go handleConn(conn)
	}
}
