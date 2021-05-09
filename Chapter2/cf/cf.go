// Cf conferts it numeric argument to Celcius , Fahrenheit or Kelvin
package main

import (
	"fmt"
	"os"
	"strconv"
	"tempconv0"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64) // convert strign to  float64

		if err != nil {
			fmt.Fprintf(os.Stderr, "cd : %v\n", err)
			os.Exit(1)
		}
	}

	k := tempconv0.Kelvin(t)
	f := tempconv0.Fahrenheit(t)
	c := tempconv0.Celcius(t)

	fmt.Printf("%s = %s, %s = %s, %s = %s \n", f, tempconv0.FtoC(f), c, tempconv0.CToF(c), k, tempconv0.CToK(k))
}
