package main

import (
	"fmt"
	"net/http"
)

func registHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	login := r.FormValue("login")
	email := r.FormValue("email")
	//password := r.FormValue("password")
	if login == "" {
		http.Error(w, "Incorrect input", http.StatusBadRequest)
		return
	}
	city := r.FormValue("city")
	fmt.Fprint(w, "Регистрация пройдена успешно!")
	fmt.Fprintf(w, "Логин: %v", login)
	fmt.Fprintf(w, "Email: %v", email)
	fmt.Fprintf(w, "Город: %v", city)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "registrator.html")
	})
	http.HandleFunc("/register", registHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
