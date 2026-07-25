package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	httpRequest := "GET / HTTP/1.1\n" +
		"Host: example.com\n\n"
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	if _, err = conn.Write([]byte(httpRequest)); err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(os.Stdout, conn)
}
