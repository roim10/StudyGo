package main

import "fmt"

type Engine struct {
	Horsepower  int
	Reliability int
}

type Car struct {
	Brand      string
	TopSpeed   int
	EngineInfo Engine
}

type Driver struct {
	Name       string
	Experience int
}

func main() {
	car := Car{Brand: "ferrari", TopSpeed: 212, EngineInfo: Engine{Horsepower: 699, Reliability: 78}}
	driver := Driver{Name: "Leclerc", Experience: 10}
	fmt.Println(car)
	fmt.Println(driver)
	UpgradeEngine(&car.EngineInfo, 15)
	fmt.Println(car)
	car1 := Car{Brand: "mercedes", TopSpeed: 231, EngineInfo: Engine{Horsepower: 731, Reliability: 87}}
	CompareCars(car1, car)
}

func UpgradeEngine(e *Engine, boost int) {
	e.Horsepower += boost
	e.Reliability -= 5
}

func AssignDriver(c *Car, d Driver) {
	fmt.Printf("Пилот %v сел за руль %v", d.Name, c.Brand)
}
func CompareCars(c1 Car, c2 Car) {
	if c1.TopSpeed > c2.TopSpeed {
		fmt.Println(c1.Brand)
	} else if c1.TopSpeed < c2.TopSpeed {
		fmt.Println(c2.Brand)
	} else {
		fmt.Println("они равны")
	}
}
