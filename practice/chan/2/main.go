package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type LogMessage struct {
	ID      int
	Message string
}

func main() {
	packet := []LogMessage{}
	wg.Add(2)
	LMChan := make(chan LogMessage, 5)
	go func() {
		defer wg.Done()
		for i := 0; i <= 7; i++ {
			LM := LogMessage{
				ID: i,
			}
			LMChan <- LM
		}
		close(LMChan)
	}()
	go func() {
		defer wg.Done()
		for val := range LMChan {
			packet = append(packet, val)
			if len(packet) == 3 {
				fmt.Println("БД: Записан пакет из 3 логов!")
				packet = nil
			}
		}
		if len(packet) > 0 {
			fmt.Printf("Отправлен пакет в БД:%v\n", packet)
		}
	}()
	wg.Wait()
}
