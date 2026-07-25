package main

import (
	"fmt"
	"net"
)

func main() {
	var str string
	fmt.Println("напишите любую строку")
	fmt.Scanln(&str)
	conn, err := net.Dial("tcp", "tcpbin.com:4242")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	if n1, err := conn.Write([]byte(str + "\n")); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(n1)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	fmt.Printf("n=%d, err=%v, buf=%q\n", n, err, buf[:n])
	if err != nil {
		fmt.Println(err)
		return
	}
}
