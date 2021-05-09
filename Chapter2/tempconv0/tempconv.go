// Package tempconv perfonms Celcius and Fahrenheit temperature computations
package tempconv

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


func CToF(c Celsius) Fahrenheit {return Fahrenheit(c*9/5+32)}		// Fahrenheit() is a type convertion not a function
func FToC(f Fahrenheit) Celsius {return Celsius((f-32)*5/9)}
func KToC(k Kelvin) Celsius {return Celsius(k+273.15)}
func CToK(c Celsius) Kelvin {return Kelcin(c-273.15)}
func FToK(f Fahrenheit) Kelvin {return Kelvin(CToK(FToC(f)))}
func KToF(k kelvin) Fahrenheit {return Fahrenheit(CToF(FToK(k)))}

