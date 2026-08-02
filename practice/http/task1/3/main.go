package main

import (
	"fmt"
	"net/http"
)

func staticHandle(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/page/index":
		http.ServeFile(w, r, "static/index.html")
	case "/page/about":
		http.ServeFile(w, r, "static/about.html")
	default:
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/page/", staticHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
