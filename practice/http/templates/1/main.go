package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ProfileData struct {
	Name string
	Age  string
	City string
}

func templeHandle(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")
	city := r.URL.Query().Get("city")
	if name == "" {
		name = "не указано"
	}
	if age == "" {
		age = "не указано"
	}
	if city == "" {
		city = "не указано"
	}
	data := ProfileData{
		Name: name,
		Age:  age,
		City: city,
	}
	tmpl, err := template.New("data").Parse("<h1>Здравствуйте, {{ .Name }}!</h1><ul><li>Возраст: {{ .Age }}</li><li>Город: {{ .City }}</li></ul>")
	if err != nil {
		http.Error(w, "Ошибка при создании шаблона", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/profile", templeHandle)
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
