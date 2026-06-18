package main

import "fmt"

type Device interface {
	TurnOn()
}

type TV struct {
	Volume  int
	Channel int
}
type AudioSystem struct {
	Volume    int
	BassBoost bool
}
type Projector struct {
	Brightness int
}

func (tv TV) TurnOn() {
	fmt.Printf("Телевизор включен, текущий канал %v\n", tv.Channel)
}
func (au AudioSystem) TurnOn() {
	fmt.Printf("Аудиосистема включена, громкость %v\n", au.Volume)
}
func (pr Projector) TurnOn() {
	fmt.Printf("Проектор включен, текущая яркость %v\n", pr.Brightness)
}
func (tv *TV) ChangeChannel(newChannel int) {
	tv.Channel = newChannel
	fmt.Printf("канал переключен на %v\n", tv.Channel)
}
func (au AudioSystem) ToggleBass() {
	au.BassBoost = !au.BassBoost
	fmt.Printf("бассбуст %v\n", au.BassBoost)
}
func IncreaseVolumeTV(tv int) {
	tv += 10
}
func IncreaseVolumeAU(au int) {
	au += 10
}

func BoostVolume(d Device) {
	if val, i := d.(*TV); i {
		IncreaseVolumeTV(val.Volume)
	} else if val, i := d.(*AudioSystem); i {
		IncreaseVolumeAU(val.Volume)
	} else if _, i := d.(*Projector); i {
		fmt.Println("У проектора нет звуковой системы!")
	} else {
		fmt.Println("не получилось")
	}
}
func ProcessDevices(devices []any) {
	for _, dev := range devices {
		switch v := dev.(type) {
		case *TV:
			v.TurnOn()
			v.ChangeChannel(5)

		case *AudioSystem:
			v.TurnOn()
			v.ToggleBass()

		case *Projector:
			v.TurnOn()

		case nil:
			fmt.Println("тут ничего нет")

		default:
			fmt.Println("Неизвестный тип гаджета")
		}
	}
}
func main() {
	var COm []any = []any{
		&TV{Volume: 20, Channel: 1},
		&AudioSystem{Volume: 15, BassBoost: false},
		&Projector{Brightness: 1500},
		nil,
		"Pharaon ебанный псих курит сканк за четверых",
	}
	ProcessDevices(COm)
}
