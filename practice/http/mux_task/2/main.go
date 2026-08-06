package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	category := vars["category"]
	response := fmt.Sprintf("Order category=%v, id=%v", category, id)
	fmt.Fprintln(w, response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/orders/{category}/{id:[0-9]+}", ProductHandler)
	err := http.ListenAndServe(":8181", router)
	if err != nil {
		fmt.Println(err)
		return
	}
}
