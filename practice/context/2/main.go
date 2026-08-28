package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func reportHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	for i := range 5 {
		select {
		case <-ctx.Done():
			log.Printf("клиент отменил запрос, шаг %v", i)
			return
		case <-time.After(1 * time.Second):
		}
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "отчёт готов")

}

func main() {
	http.HandleFunc("/report", reportHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
