package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func contextHandle(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	defer cancel()
	result := make(chan string, 1)
	go func() {
		time.Sleep(5 * time.Second)
		result <- "данные получены"
	}()
	work(ctx, w, result)
}

func work(ctx context.Context, w http.ResponseWriter, data chan string) {
	select {
	case <-ctx.Done():
		http.Error(w, "операция отменена по таймауту", http.StatusRequestTimeout)
		return
	case result := <-data:
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, result)
	}
}

func main() {
	http.HandleFunc("/search", contextHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
