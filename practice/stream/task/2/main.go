package main

import "fmt"

type vowelWriter struct{}

func (v vowelWriter) Write(bs []byte) (int, error) {
	if len(bs) == 0 {
		return 0, nil
	}
	for i := range bs {
		switch bs[i] {
		case 'i', 'e', 'a', 'o', 'u':
			fmt.Print(string(bs[i]))
		}
	}
	return len(bs), nil
}

func main() {
	byte0 := []byte("When forty winters shall besiege thy brow, And dig deep trenches in thy beauty field, Thy youth proud livery so gazed on now")
	writer := vowelWriter{}
	writer.Write(byte0)
}
