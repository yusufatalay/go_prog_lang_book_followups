// Cf conferts it numeric argument to Celcius , Fahrenheit or Kelvin
package main

import (
	"fmt"
	"os"
	"strconv"
	"tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64) // convert strign to  float64

		if err != nil {
			fmt.Fprintf(os.Stderr, "cd : %v\n", err)
			os.Exit(1)
		}

		k := tempconv.Kelvin(t)
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		fmt.Printf("%s = %s, %s = %s, %s = %s \n", f, tempconv.FToC(f), c, tempconv.CToF(c), k, tempconv.KToC(k))
	}
}
