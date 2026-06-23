package main

import "fmt"

var data []int

func fillData(n int, ready chan struct{}) {
	for i := range n {
		data = append(data, i)
	}
	close(ready)
}
func processElement(index int, done chan bool) {
	data[index] *= 10
	fmt.Printf("Результат: %v \n", data[index])
	done <- true
}

func main() {
	readyCh := make(chan struct{})
	go fillData(5, readyCh)
	<-readyCh
	count := len(data)
	fmt.Printf("длина массива: %v \n", count)
	DoneCh := make(chan bool)
	for i := range count {
		go processElement(i, DoneCh)
	}
	for range count {
		<-DoneCh
	}
	fmt.Println("Все горутины завершили работу. Конец.")
}
