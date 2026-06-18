package main

import (
	"fmt"
)

func main() {
	var lib libiry = libiry{"koe", "poe"}
	lib.stad()
	tom := person{age: 18, name: "tom", vegent: false}
	tom.stud()
	tom.studa("carrot")
}

type libiry []string

func (l libiry) stad() {
	for _, val := range l {
		fmt.Println(val)
	}
}

type person struct {
	name   string
	age    int
	vegent bool
}

func (p person) stud() {
	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.vegent)
}
func (p person) studa(meal string) {
	fmt.Println(p.name, "eat", meal)
}
