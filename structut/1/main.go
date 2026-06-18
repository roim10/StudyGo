package main

import (
	"fmt"
)

type person struct {
	name   string
	age    int
	female bool
}
type class struct {
	string
	int
}

func main() {
	// var pe person = person{name: "Tom", age: 12, female: false}
	// var per person = person{"Lola", 23, true}
	// fmt.Println(pe)
	// fmt.Println(per)
	// fmt.Println(pe.female)
	// fmt.Println(per.female)
	// team := struct {
	// 	name string
	// 	age  int
	// }{
	// 	name: "Temo",
	// 	age:  12,
	// }
	// // fmt.Println(team)
	// a9 := class{"nasz", 12}
	// fmt.Println(a9)
	// var p_a9 *class = &a9
	// var a10 *class = &class{"til", 12}
	// fmt.Println(a10.int)
	// fmt.Println(p_a9.string)
	// a10.int = 13
	// fmt.Println(a10.int)
	// Belov := new(gouverment)
	// fmt.Println(*Belov)
	// Belov.posol = false
	// Belov.ministr = true
	// fmt.Println(*Belov)
	// bob := new(struct{name, company string; age int})
	// tom := person{name: "tom", age: 12, female: false}
	// tomas := tom
	// tomas.age = 23
	// fmt.Println(tom.age)
	// fmt.Println(tomas.age)
	// tim := person{name: "tim", age: 12, female: false}
	// sum(&tim)
	// fmt.Println(tim)
	tom := person{"Tom", 41, false}
	bob := person{"Bob", 41, false}
	tomas := person{"Tom", 41, false}
	fmt.Println(tom == bob)
	fmt.Println(tom == tomas)
}

func sum(p *person) {
	p.age += 1
	fmt.Println(p.age)
}

type gouverment struct {
	posol   bool
	ministr bool
}
