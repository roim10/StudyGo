package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func isPortOpen(host string, port int) bool {
	conn, err := net.DialTimeout("tcp", host+":"+strconv.Itoa(port), 8*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func main() {
	n := isPortOpen("google.com", 80)
	fmt.Println(n)
}
