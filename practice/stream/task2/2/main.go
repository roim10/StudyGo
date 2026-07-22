package main

import (
	"io"
	"os"
)

type vowelReader string

func (v vowelReader) Read(p []byte) (int, error) {
	count := 0
	for _, ch := range v {
		switch ch {
		case 'i', 'e', 'a', 'o', 'u':
			p[count] = byte(ch)
			count++
		}
	}
	return count, io.EOF
}

func main() {
	hw := vowelReader("Hello world")
	io.Copy(os.Stdout, hw)
}
