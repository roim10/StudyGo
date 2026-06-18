package main

import "fmt"

type Sender interface {
	Send(message string) bool
	GetChannelName() string
}
type EmailSender struct {
	EmailAddress string
}
type SMSSeneder struct {
	PhoneNumber string
}
type PushSender struct {
	DeviceID string
}

func (ES EmailSender) Send(msg string) bool {
	ConnectionLoss := false
	if !ConnectionLoss {
		fmt.Printf("email %s: %s\n", ES.EmailAddress, msg)
	} else {
		fmt.Println("no")
	}
	return ConnectionLoss
}
func (ES EmailSender) GetChannelName() string {
	return "email"
}
func (SMSS SMSSeneder) Send(msg string) bool {
	ConnectionLoss := false
	if !ConnectionLoss {
		fmt.Printf("SMS %s: %s\n", SMSS.PhoneNumber, msg)
	} else {
		fmt.Println("no")
	}
	return ConnectionLoss
}
func (SMSS SMSSeneder) GetChannelName() string {
	return "SMS"
}
func (PS PushSender) Send(msg string) bool {
	ConnectionLoss := false
	if !ConnectionLoss {
		fmt.Printf("Push %s: %s\n", PS.DeviceID, msg)
	} else {
		fmt.Println("no")
	}
	return ConnectionLoss
}
func (PS PushSender) GetChannelName() string {
	return "Push"
}
func Notify(s Sender, text string) {
	s.GetChannelName()
	s.Send(text)
}
func LogNotification(data any) {
	fmt.Println(data)
}

type code struct {
	CodeName string
}

func main() {
	var spisok = []Sender{
		EmailSender{EmailAddress: "user@example.com"},
		SMSSeneder{PhoneNumber: "+79991112233"},
		PushSender{DeviceID: "iPhone15"},
	}
	for _, s := range spisok {
		Notify(s, "Hello!")
	}
	code_1 := code{"code 200"}
	LogNotification(code_1)
	email := EmailSender{"user@example.com"}
	LogNotification(email)
}
