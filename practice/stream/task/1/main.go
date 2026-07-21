package main

import (
	"io"
	"os"
	"strings"
)

type upperReader string

func (up upperReader) Read(p []byte) (int, error) {
	count := 0
	upper := strings.ToUpper(string(up))
	for i := range upper {
		p[i] = upper[i]
		count++
	}
	return count, io.EOF
}

func main() {
	file, err := os.Create("result.txt")
	if err != nil {
		os.Exit(1)
	}
	hw := upperReader("Hello world")
	buf := make([]byte, len(hw))
	n, err := hw.Read(buf)
	if err != nil && err != io.EOF {
		os.Exit(1)
	}
	_, err = file.Write(buf[:n])
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
}
