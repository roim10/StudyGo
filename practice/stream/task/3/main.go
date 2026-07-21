package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("numbers.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		os.Exit(1)
	}
	_, err = file.Write([]byte("12345"))
	if err != nil {
		os.Exit(1)
	}
	file.Close()
	file, err = os.Open("numbers.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	buf := make([]byte, 5)
	var n int
	n, err = file.Read(buf)
	if err != nil && err != io.EOF {
		os.Exit(1)
	}
	fmt.Println(string(buf[:n]))

}
