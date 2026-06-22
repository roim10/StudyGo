package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID          int
	Amount      float64
	IsProcessed bool
}

func GenerateOrders(count int) chan int {
	ch := make(chan int)
	go func() {
		for i := range count {
			ch <- i
		}
		close(ch)
	}()
	return ch
}
func ProcessOrders(inChan chan int, outChan chan Order) {
	for id := range inChan {
		order := Order{
			ID:          id,
			Amount:      float64(id) * 10.5,
			IsProcessed: true,
		}
		outChan <- order
	}
	close(outChan)
}

func main() {
	outChan := make(chan Order, 3)
	inChan := GenerateOrders(3)
	go ProcessOrders(inChan, outChan)
	time.Sleep(10 * time.Millisecond)
	i := cap(outChan)
	fmt.Println(i)
	i = len(outChan)
	fmt.Println(i)
	for order := range outChan {
		fmt.Printf("Получен обработанный заказ: ID=%d, Amount=%.1f, Processed=%t\n",
			order.ID, order.Amount, order.IsProcessed)
	}
}
