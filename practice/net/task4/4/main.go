package main

import (
	"fmt"
	"net/http"
)

func fetchStatus(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

func main() {
	n, err := fetchStatus("https://github.com")
	fmt.Printf("код статуса: %v и Ошибка: %v\n", n, err)
}
