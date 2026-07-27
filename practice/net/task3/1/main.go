package main

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		input := make([]byte, (1024 * 4))
		conn.SetReadDeadline(time.Now().Add(time.Second * 5))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			var netErr net.Error
			if errors.As(err, &netErr) && netErr.Timeout() {
				fmt.Println("Таймаут")
				conn.Write([]byte("timeout"))
				continue
			}
			conn.Write([]byte("Ошибка: соеденение потеряно\n"))
			break

		}
		source := string(input[:n])
		partSource := strings.Fields(source)
		if len(partSource) != 3 {
			conn.Write([]byte("Напишите математическое выражение по шаблону: 1 число знак 2 число"))
			continue
		}
		num1, err := strconv.Atoi(partSource[0])
		if err != nil {
			conn.Write([]byte("Ошибка: первое значение не является числом\n"))
			continue
		}
		num2, err := strconv.Atoi(partSource[2])
		if err != nil {
			conn.Write([]byte("Ошибка: второе значение не является числом\n"))
			continue
		}
		op := partSource[1]
		result := 0
		if num2 == 0 && op == "/" {
			conn.Write([]byte("Ошибка: делить на ноль нельзя"))
			continue
		}
		switch op {
		case "+":
			result = num1 + num2
		case "/":
			result = num1 / num2
		case "*":
			result = num1 * num2
		case "-":
			result = num1 - num2
		default:
			fmt.Println("unknown operator")
			conn.Write([]byte("Неизвестный оператор. Пожалйста выберите оператора из приведенного списка"))
			continue
		}
		resultStr := strconv.Itoa(result)
		conn.Write([]byte(resultStr))
	}
}

func main() {
	listener, err := net.Listen("tcp", ":4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("server listening")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn)
	}
}
