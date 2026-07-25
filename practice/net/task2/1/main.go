package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		return
	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		buf := make([]byte, 1024)

		go func() {
			defer conn.Close()
			for {
				n, err := conn.Read(buf)
				if err != nil {
					return
				}
				str := string(buf[:n])
				str = strings.ToUpper(str)
				fmt.Printf("n=%d, err=%v, buf=%q\n", n, err, str)
				_, err = conn.Write([]byte(str))
				if err != nil {
					return
				}
			}
		}()

	}
}
