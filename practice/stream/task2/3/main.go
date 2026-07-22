package main

import (
	"fmt"
	"io"
	"os"
)

type byteCounter struct {
	counter int
}

func (b *byteCounter) Write(bs []byte) (int, error) {
	if len(bs) == 0 {
		return 0, nil
	}
	b.counter += len(bs)
	return len(bs), nil
}

func main() {
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	writer := &byteCounter{
		counter: 0}
	n, _ := io.Copy(writer, file)
	fmt.Println(writer.counter, n)
}
