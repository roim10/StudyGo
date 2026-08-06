package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func IdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("user id=%v", id)
	fmt.Fprintln(w, response)
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome")
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users/{id:[0-9]+}", IdHandler)
	router.HandleFunc("/", WelcomeHandler)
	err := http.ListenAndServe(":8181", router)
	if err != nil {
		fmt.Println(err)
		return
	}
}
