package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.github.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("http code: %v\n", resp.StatusCode)
	n, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("количество байт в теле ответа: %v", n)
}
