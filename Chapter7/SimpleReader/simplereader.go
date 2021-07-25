// exercise 7.4
package simplereader

import "io"

type readstring struct {
	s string
}

func (r *readstring) Read(b []byte) (n int, err error) {
	n, err = r.Read(b[:])
	if len(r.s) == 0 {
		err = io.EOF
	}
	n = len(r.s)
	return
}

func NewReader(s string) io.Reader {
	return &readstring{s}
}
