package main

import (
	"fmt"
)

func main() {
	fmt.Println(ToFahrenheit(12))
	pointers()
	a := 12
	Double(a)
	fmt.Println(a)
	DoublePtr(&a)
	var id1 UserId = 1
	defer fmt.Println(id1)
	UpdateId(&id1)
	fmt.Println(id1)
	UpdateId(&id1)
	fmt.Println(id1)
	UpdateId(&id1)
	fmt.Println(id1)
	UpdateId(&id1)
	fmt.Println(id1)
	LastTrack()
}

type Celsius float64

type Fahrenheit float64

func ToFahrenheit(c Celsius) Fahrenheit {
	var FinC Fahrenheit
	fmt.Scan(&c)
	FinC = Fahrenheit(c*9/5 + 32)
	return FinC
}

func pointers() {
	var x int = 10
	var p *int = &x
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(*p)
	fmt.Println(&p)
	*p = 42
	fmt.Println(x)
}
func Double(n int) {
	n *= 2
	fmt.Println(n)
}

func DoublePtr(n *int) {
	*n *= 2
	fmt.Println(*n)
}

type UserId int

func UpdateId(id *UserId) {
	*id += 1
}
func LastTrack() {
	for i := range 10 {
		defer fmt.Println(i)
		fmt.Println("закончен")
	}
}
