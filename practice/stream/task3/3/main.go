package main

import (
	"fmt"
	"os"
)

type survey struct {
	name       string
	age        int
	favoritePl string
}

func (s *survey) survey() {
	fmt.Print("Ваше имя: ")
	_, err := fmt.Scanln(&s.name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("Ваш возраст: ")
	_, err = fmt.Scanln(&s.age)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("Ваш любимый язык программирования: ")
	_, err = fmt.Scanln(&s.favoritePl)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	var s survey
	s.survey()
	file, err := os.OpenFile("profile.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	fmt.Fprintln(file, s.name, s.age, s.favoritePl)
}
