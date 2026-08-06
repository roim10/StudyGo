package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("id=%v", id)
	fmt.Fprintln(w, response)
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index Page")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/products/{id:[0-9]+}", ProductHandler)
	router.HandleFunc("/articles/{id:[0-9]+}", ProductHandler)
	router.HandleFunc("/", PageHandler)
	err := http.ListenAndServe(":8181", router)
	if err != nil {
		fmt.Println(err)
		return
	}
}
