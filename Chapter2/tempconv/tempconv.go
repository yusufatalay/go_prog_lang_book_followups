// Package tempconv performs Celcius and Fahrenheit temperature computations

import "fmt"

type Celsius float64			// even though C and F has same underlying type their types are different
type Fahrenheit float64
type Kelvin float64
cosnt (
		AbsoluteZeroC = -273.15
		FreezingC	  = 0
		BoilingC	  = 100
		)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }


