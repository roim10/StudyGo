package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func searchHandle(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	pagestr := r.FormValue("page")
	page := 1
	if pagestr != "" {
		p, err := strconv.Atoi(pagestr)
		if err != nil {
			http.Error(w, "Internal input", http.StatusBadRequest)
			return
		}
		page = p
	}
	limitstr := r.FormValue("limit")
	limit := 10
	if limitstr != "" {
		lim, err := strconv.Atoi(limitstr)
		if err != nil {
			http.Error(w, "Incorrect input", http.StatusBadRequest)
			return
		}
		limit = lim
	}
	fmt.Fprintf(w, "Поиск по запросу: \"%s\"\n", query)
	fmt.Fprintf(w, "Страница: %v, результатов на странице: %v\n", page, limit)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "search.html")
	})
	http.HandleFunc("/search", searchHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
