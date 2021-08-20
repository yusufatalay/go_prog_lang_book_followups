package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("rocket goes to the space")
}

func main() {
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	fmt.Println("Commencing countdown.")
	select {
	case <-time.After(10 * time.Second):
		// Do nothing
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}
