package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(p []byte) (n int, err error) {
	buffer := make([]byte, 1024)
	buf_len, err := rot.r.Read(buffer)
	for i := 0; i < buf_len; i++ {
		switch val := buffer[i]; {
		case val >= 'A' && val <= 'M',
			val >= 'a' && val <= 'm':
			p[i] = val + 13
		case val >= 'N' && val <= 'Z',
			val >= 'n' && val <= 'z':
			p[i] = val - 13
		default:
			p[i] = val
		}
	}
	return buf_len, err
}

func main() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
