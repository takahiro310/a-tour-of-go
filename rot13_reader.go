package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (n int, e error) {

	const ROT13 = 13
	n, e = r.r.Read(b)

	if e != nil {
		return
	}

	for i := 0; i < len(b); i++ {
		if 'A' <= b[i] && b[i] <= 'Z' {
			b[i] = (b[i]-'A'+ROT13)%26 + 'A'
		}
		if 'a' <= b[i] && b[i] <= 'z' {
			b[i] = (b[i]-'a'+ROT13)%26 + 'a'
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
