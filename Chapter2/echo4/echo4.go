// Echo4 prints its command-line arguments
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing new line") // flag.Bool creates new bool typre flag
var sep = flag.String("s", " ", "seperator")

func main() {
	flag.Parse()                               // flag.Parse must be called before using flags
	fmt.Print(strings.Join(flag.Args(), *sep)) // in order to use this flags in code we have to use them as pointers

	if !*n {
		fmt.Println()
	}
}
