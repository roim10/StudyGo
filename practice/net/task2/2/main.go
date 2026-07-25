package main

import (
	"bufio"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:9090", 8*time.Second)
	if err != nil {
		return
	}
	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		str := scanner.Text()
		_, err = conn.Write([]byte(str + "\n"))
		if err != nil {
			return
		}
		buf := make([]byte, 1024)
		_, err = conn.Read(buf)
		if err != nil {
			return
		}
		if str == "exit" {
			break
		}
	}

}
