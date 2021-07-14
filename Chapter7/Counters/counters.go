package main

import (
	"bufio"
	"fmt"
	"strings"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(s []byte) (int, error) {
	counter := 0

	scanner := bufio.NewScanner(strings.NewReader(string(s[:])))

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		counter++
	}

	*c += WordCounter(counter)
	return counter, nil
}

type LineCounter int

func (c *LineCounter) Write(s []byte) (int, error) {
	counter := 0

	scanner := bufio.NewScanner(strings.NewReader(string(s[:])))

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		counter++
	}
	*c += LineCounter(counter)
	return counter, nil

}

func main() {
	var b ByteCounter
	b.Write([]byte("hello"))
	fmt.Println(b)

	b = 0 // reset the counter
	var bname = "Desmond"
	fmt.Fprintf(&b, "hello, %s", bname)

	fmt.Println(b)

	var w WordCounter
	w.Write([]byte("     hello mah name jeff    "))
	fmt.Println(w)

	w = 0
	var wname = "Desmond"
	fmt.Fprintf(&w, "yellow, %s", wname)
	fmt.Println(w)

	var l LineCounter
	l.Write([]byte("     hello mah \nname jeff    "))
	fmt.Println(l)

	l = 0
	var lname = "Desmond"
	fmt.Fprintf(&l, "yel\nlow\n, %s", lname)
	fmt.Println(l)

}
