package main

import (
	"fmt"
	"time"
)

// Имитация сервиса погоды
func fetchWeather(ch chan string) {
	// Имитируем небольшую задержку сети
	time.Sleep(100 * time.Millisecond)
	ch <- "Погода: +25°C, Солнечно"
}

// Имитация сервиса валют
func fetchCurrency(ch chan string) {
	time.Sleep(50 * time.Millisecond)
	ch <- "Курс: 1 USD = 90 RUB"
}

// Имитация сервиса новостей
func fetchNews(ch chan string) {
	time.Sleep(150 * time.Millisecond)
	ch <- "Новости: Go 1.26 официально выпущен!"
}

func main() {
	ch1, ch2, ch3 := make(chan string), make(chan string), make(chan string)
	go fetchCurrency(ch1)
	go fetchNews(ch2)
	go fetchWeather(ch3)

	for i := 1; i <= 3; i++ {
		select {
		case ch1_val := <-ch1:
			fmt.Println(ch1_val)
		case ch2_val := <-ch2:
			fmt.Println(ch2_val)
		case ch3_val := <-ch3:
			fmt.Println(ch3_val)
		}

	}
}
