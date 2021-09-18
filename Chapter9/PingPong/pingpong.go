package main

import (
	"fmt"
	"time"
)

var left = make(chan struct{})
var right = make(chan struct{})
var timeout = make(chan struct{})
var connection_count int

func main() {

	// start the game from the left

	tmr := time.NewTimer(1 * time.Second)

	go func() {
		for {
			select {
			case <-timeout:
				fmt.Println(connection_count)
				break
			case b := <-left:
				connection_count++
				right <- b
			}
		}
	}()
	go func() {
		for {
			select {
			case <-timeout:
				fmt.Println(connection_count)
				break
			case b := <-right:
				connection_count++
				left <- b
			}
		}
	}()

	left <- struct{}{}
	<-tmr.C
	close(timeout)
	fmt.Println(connection_count)

}
