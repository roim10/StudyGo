package main

import "fmt"

type Vehicle interface {
	move()
}

type car struct{}
type plane struct{}

func (c car) move() {
	fmt.Println("car edet")
}
func (p plane) move() {
	fmt.Println("plane letit")
}

func main() {
	var ferrari Vehicle
	fmt.Println(ferrari)
	var car Vehicle = car{}
	var plane Vehicle = plane{}
	plane.move()
	car.move()
	myPhone := SMS{phoneNumber: "+79991112233"}
	myMail := Email{adress: "tom@example.com"}
	not := []nin{myMail, myPhone}
	for _, v := range not {
		Modif(v, "lol")
	}
}
func print(value any) {
	fmt.Println(value)
}

type nin interface {
	Send(msg string)
}
type SMS struct {
	phoneNumber string
}

func (s SMS) Send(msg string) {
	fmt.Printf("sms %s: %s\n", s.phoneNumber, msg)
}

type Email struct {
	adress string
}

func (e Email) Send(msg string) {
	fmt.Printf("email %s: %s\n", e.adress, msg)
}
func Modif(n nin, text string) {
	n.Send(text)
}
