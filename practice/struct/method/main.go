package main

import "fmt"

func main() {
	ferrari := racecar{model: "ferrari", Fuel: 98, TyreWear: 7, EngineTemp: 12, IsPitLimited: true}
	ferrari.GetStatus()
	ferrari.DriveLap()
	ferrari.PitStop()
	ferrari.TogglePitLimiter()
	ferrari.GetStatus()
	mercedes := racecar{model: "mercedes", Fuel: 100, TyreWear: 5, EngineTemp: 12, IsPitLimited: false}
	mercedes.GetStatus()
	CompareTelemetry(ferrari, mercedes)
	mercedes.GetStatus()
	ferrari.GetStatus()
}

type racecar struct {
	model        string
	Fuel         float32
	TyreWear     float32
	EngineTemp   float64
	IsPitLimited bool
}

func (f *racecar) DriveLap() {
	f.Fuel -= 3
	f.TyreWear += 5
	f.EngineTemp += 2
}
func (f *racecar) PitStop() {
	f.TyreWear = 0
	f.Fuel = 100
}
func (f racecar) GetStatus() {
	fmt.Println(f.model)
	fmt.Println(f.Fuel)
	fmt.Println(f.EngineTemp)
	fmt.Println(f.IsPitLimited)
	fmt.Println(f.TyreWear)
}
func (f *racecar) TogglePitLimiter() {
	f.IsPitLimited = !f.IsPitLimited
}
func CompareTelemetry(car1, car2 racecar) {
	car1.breaking()
	car2.breaking()
}
func (r *racecar) breaking() {
	r.EngineTemp = 101
}
