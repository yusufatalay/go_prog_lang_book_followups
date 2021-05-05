// echo2 prints its command-line arguments
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	s := ""
	// edited for execise 1.2
	for i, arg := range os.Args[1:] {
		s = arg

		fmt.Printf("Args[%d]: %s", i+1, s)
		fmt.Println()
	}

	duration := time.Since(start)
	fmt.Println("I took ", duration)
}
