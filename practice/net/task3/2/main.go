package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:4545", 8*time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Напишите пожалуйста математическое выражение, обязательно с пробелами между числаи и знаком действия. Доступные знаки действия: +, -, *, /")
		scanner.Scan()
		str := scanner.Text()
		if n, err := conn.Write([]byte(str)); n == 0 || err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("Ответ: ")
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
		fmt.Println()
	}
}
