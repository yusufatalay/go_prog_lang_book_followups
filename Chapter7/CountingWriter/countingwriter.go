package main

import (
	"fmt"
	"io"
	"os"
)

type ByteCounter struct {
	len int64
	w   io.Writer
}

func (b *ByteCounter) Write(p []byte) (n int, err error) {
	n, err = b.w.Write(p)
	b.len += int64(n)

	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	b := ByteCounter{0, w}
	return &b, &b.len
}

func main() {
	// also counts the new line character
	l, w := CountingWriter(os.Stdout)
	fmt.Fprintf(l, "%d\n", 420)
	fmt.Println(*w)

}
