package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calcHandle(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	op := r.URL.Query().Get("op")
	num1, err := strconv.Atoi(a)
	if err != nil {
		http.Error(w, "invalid number: "+a, http.StatusBadRequest)
		return
	}
	num2, err := strconv.Atoi(b)
	if err != nil {
		http.Error(w, "invalid number: "+b, http.StatusBadRequest)
		return
	}
	if num2 == 0 && op == "/" {
		http.Error(w, "division by zero", http.StatusBadRequest)
		return
	}
	result := 0
	switch op {
	case "+":
		result = num1 + num2
		fmt.Fprintf(w, "result: %v", result)
	case "-":
		result = num1 - num2
		fmt.Fprintf(w, "result: %v", result)
	case "/":
		result = num1 / num2
		fmt.Fprintf(w, "result: %v", result)
	case "*":
		result = num1 * num2
		fmt.Fprintf(w, "result: %v", result)
	default:
		http.Error(w, "unknown operation", http.StatusBadRequest)
	}
}
func main() {
	http.HandleFunc("/calc", calcHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
