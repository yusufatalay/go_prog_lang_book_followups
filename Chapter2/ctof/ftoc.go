// ftoc prints two Celcius-toFahrenheit conversions
package main

import "fmt"

func cToF(c float64) float64 {
	return (c * 9 / 5) + 32
}

func main() {
	var boilingC, freezingC = 100.0, 0.0

	fmt.Printf("Boiling point of water is %g\n", cToF(boilingC))
	fmt.Printf("Freezin gpoint of water is %g\n", cToF(freezingC))
}
