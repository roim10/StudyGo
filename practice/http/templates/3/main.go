package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Product struct {
	Title string
	Price int
}
type GoodsDat struct {
	Products []Product
	Total    int
}

func templateHandle(w http.ResponseWriter, r *http.Request) {
	data := []Product{
		{"Ноутбук", 55000},
		{"Мышь", 1200},
		{"Клавиатура", 2500},
	}
	result := 0
	for _, p := range data {
		result += p.Price
	}
	dat := GoodsDat{
		Products: data,
		Total:    result,
	}
	tmpl := template.Must(template.ParseFiles("goods.html"))
	tmpl.Execute(w, dat)
}

func main() {
	http.HandleFunc("/goods", templateHandle)
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
