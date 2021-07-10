package main

import "fmt"

func main() {

	panicreturn("")
}

func panicreturn(input string) {
	defer func() (err error) {
		if p := recover(); p != nil {
			fmt.Printf("%v", p)
			err = fmt.Errorf("%v", p)
		}
		return
	}()
	if input == "" {
		panic("empty string\n")
	}

}
