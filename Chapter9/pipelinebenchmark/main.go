// contains exercise 9.4
package main

import (
	"fmt"
	"time"
)

const N = 999999

func main() {
	message := "hi"
	var mid [N]chan string

	for i := 0; i < N; i++ {
		mid[i] = make(chan string)
	}

	head, tail := mid[0], mid[N-1]

	for i := 0; i < N-1; i++ {
		go func(i int) {
			mid[i+1] <- <-mid[i]
		}(i)
	}

	s := time.Now()
	head <- message
	<-tail
	fmt.Printf("%d goroutines, %fs\n", N, time.Since(s).Seconds())
}
