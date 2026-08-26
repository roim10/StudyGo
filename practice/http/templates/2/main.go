package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type ExamResult struct {
	Score  int
	Passed bool
}

func templateHandle(w http.ResponseWriter, r *http.Request) {
	scoreStr := r.URL.Query().Get("score")
	score, err := strconv.Atoi(scoreStr)
	if err != nil {
		http.Error(w, "internal input", http.StatusBadRequest)
		return
	}
	if score < 0 || score > 100 {
		http.Error(w, "Internal input", http.StatusBadRequest)
		return
	}
	var pass bool
	if score >= 60 {
		pass = true
	} else {
		pass = false
	}
	result := ExamResult{
		Score:  score,
		Passed: pass,
	}
	tmpl := template.Must(template.ParseFiles("result.html"))
	tmpl.Execute(w, result)
}

func main() {
	http.HandleFunc("/exam", templateHandle)
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
