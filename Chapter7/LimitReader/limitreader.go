// exercise 7.5
package main

import (
	"fmt"
	"io"
	"strings"
)

type limitreader struct {
	r io.Reader
	l int // limit of this reader
}

func (l *limitreader) Read(b []byte) (n int, err error) {
	n, err = l.r.Read(b[:l.l])
	if n >= l.l {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &limitreader{r: r, l: n}
}

func main() {
	s := strings.NewReader("hellooo mi boiiiii")
	toprint := LimitReader(s, 4)

	buf := new(strings.Builder)

	_, err := io.Copy(buf, toprint)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf.String())
}
