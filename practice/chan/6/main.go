package main

import (
	"fmt"
	"math/rand"
)

func generateCPUMetrics(count int, ch chan int) {
	defer close(ch)
	for i := 1; i <= count; i++ {
		num := rand.Intn(91) + 10
		ch <- num
	}
}

func main() {
	ch := make(chan int)
	var n float64
	var num1 float64

	go generateCPUMetrics(10, ch)
	for num := range ch {
		fmt.Println(num)
		num1 += float64(num)
		n += 1
		if num >= 80 {
			fmt.Printf("[WARNING] Зафиксирован пик нагрузки: %v%% \n", num)
		}
	}

	fmt.Println(num1 / n)
}
