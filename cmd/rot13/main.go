package main

import (
	"io"
	"os"
	"strings"
)

type Rot13Reader struct {
	r io.Reader
}

func (rot Rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b)

	for j := range b {
		if b[j] >= 'A' && b[j] <= 'z' {
			if (b[j] > 'M' && b[j] <= 'Z') || (b[j] > 'm') {
				b[j] -= 13
			} else {
				b[j] += 13
			}
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := Rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
