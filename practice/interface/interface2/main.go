package main

import (
	"fmt"
	"time"
)

type SessionReader interface {
	GetToken() string
	IsValid() bool
}
type SessionWriter interface {
	Extend(duration time.Duration)
	Block()
}
type UserSession struct {
	Token     string
	ExpiresAt time.Time
	IsActive  bool
}

func (US UserSession) GetToken() string {
	return US.Token
}
func (US UserSession) IsValid() bool {
	if US.IsActive == true && time.Now().Before(US.ExpiresAt) {
		return true
	} else {
		return false
	}
}
func (US *UserSession) Extend(duration time.Duration) {
	US.ExpiresAt = time.Now().Add(duration)
}
func (US *UserSession) Block() {
	US.IsActive = false
}
func CheckAccess(SR SessionReader) {
	if SR.IsValid() {
		fmt.Printf("Доступ РАЗРЕШЕН для токена: %s\n", SR.GetToken())
	} else {
		fmt.Println("Доступ ЗАПРЕЩЕН: сессия недействительна")
	}
}
func RefreshSession(SW SessionWriter) {
	SW.Extend(10 * time.Minute)
	fmt.Println("Сессия успешно продлена на 10 минут.")
}
func (US UserSession) BuggyBlock() {
	US.IsActive = false
}
func main() {
	var MySession UserSession
	MySession.Token = "test-token"
	MySession.IsActive = true
	CheckAccess(MySession)
	CheckAccess(&MySession)
	//RefreshSession(MySession) ОШИБКА: Обычная структура не удовлетворяет SessionWriter, так как методы реализованы для *UserSession
	RefreshSession(&MySession)
	CheckAccess(&MySession)
	MySession.BuggyBlock()
	fmt.Println("Статус после BuggyBlock:", MySession.IsActive) //потому нодо поинтер ставить. и у меня есть функция блок
}
