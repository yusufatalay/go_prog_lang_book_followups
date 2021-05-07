// Boiling prints the boiling point of water
package main

import "fmt"

const boilingC = 100.0

func main() {
	var c = boilingC
	var f = (c * 9 / 5) + 32
	fmt.Printf("boiling point of water = %gC or %gF\n", c, f)
}
