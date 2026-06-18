package main

import (
	"fmt"
)

type AudioPlayer interface {
	PlaySound(trackName string)
}
type LightController interface {
	ToggleLight(on bool)
}
type SmartSpeaker struct {
	Name         string
	Volume       int
	IsLightOn    bool
	CurrentTrack string
}

func (s *SmartSpeaker) PlaySound(trackName string) {
	s.CurrentTrack = trackName
	fmt.Printf("Воспроизведение трека: %s \n", s.CurrentTrack)
}
func (s *SmartSpeaker) ToggleLight(on bool) {
	s.IsLightOn = on
	fmt.Printf("Подсветка переключена:%t \n", s.IsLightOn)
}
func (s *SmartSpeaker) SetVolume(level int) {
	s.Volume = level
	fmt.Printf("Уровень громкости изменен до: %v \n", s.Volume)
}

func StartParty(player AudioPlayer, track string) {
	player.PlaySound(track)
}
func ControlAmbiance(light LightController, status bool) {
	light.ToggleLight(status)
}

type ConfigReader interface {
	ReadConfig() string
}
type ConfigWriter interface {
	WriteConfig(data string)
}
type ConfigManager interface {
	ConfigReader
	ConfigWriter
}
type BackupDrive struct {
	payload string
}

func (b *BackupDrive) ReadConfig() string {
	return b.payload
}

func (b *BackupDrive) WriteConfig(data string) {
	b.payload = data
	fmt.Println(b.payload)
}

type AdvancedConfigManager interface {
	ConfigReader
	WriteConfig(data string)
}

func SyncData(manager AdvancedConfigManager, newData string) {
	manager.ReadConfig()
	manager.WriteConfig(newData)
}
func main() {
	var SS SmartSpeaker
	SS.SetVolume(75)
	StartParty(&SS, "Oxxxymiron - не от мира сего")
	ControlAmbiance(&SS, false)
	SS.PlaySound("Oxxxymiron - не от мира сего")
	SS.ToggleLight(true)
	bd := &BackupDrive{
		payload: "123",
	}
	fmt.Println(bd)
	SyncData(bd, "lol")
}
