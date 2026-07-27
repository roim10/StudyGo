package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}
	req, err := http.NewRequest(
		"GET", "https://httpbin.org/headers", nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "go-practice-app")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
