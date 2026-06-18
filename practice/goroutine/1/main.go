package main

import (
	"fmt"
	"time"
)

func processRequest(id int, duration time.Duration) {
	fmt.Printf("Запрос %v начал обрабатываться...\n", id)
	time.Sleep(duration)
	fmt.Printf("Запрос %v успешно обработан за %v мс!\n", id, duration)
}

func main() {
	for i := range 5 {
		go func(id int, duration time.Duration) {
			fmt.Printf("Запрос %v начал обрабатываться...\n", id)
			time.Sleep(duration)
			fmt.Printf("Запрос %v успешно обработан за %v мс!\n", id, duration)
		}(i, time.Duration(i)*50*time.Millisecond)
		//	go processRequest(i, time.Duration(i)*50*time.Millisecond)
	}
	time.Sleep(1 * time.Second)
}
