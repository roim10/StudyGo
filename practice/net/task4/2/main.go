package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		//Timeout: 3* time.Second,
		Timeout: 1 * time.Nanosecond,
	}
	resp, err := client.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
}
