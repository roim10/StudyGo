package main

import (
	"fmt"
	"sync"
)

func generate(out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- 12
}

func modify(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	num := <-in
	num *= 5
	out <- num
}
func printResult(in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(<-in)
}
func main() {
	var wg sync.WaitGroup
	ch1, ch2 := make(chan int), make(chan int)
	wg.Add(3)
	go generate(ch1, &wg)
	go modify(ch1, ch2, &wg)
	go printResult(ch2, &wg)
	wg.Wait()
}
