package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type MobileAlertStater interface {
	Alert() string
}

// MobileAlert implements an alert depending on its state.
type MobileAlert struct {
	state MobileAlertStater
}

// Alert returns a alert string
func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

// SetState changes state
func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

// NewMobileAlert is the MobileAlert constructor.
func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
}

// MobileAlertVibration implements vibration alert
type MobileAlertVibration struct {
}

// Alert returns a alert string
func (a *MobileAlertVibration) Alert() string {
	return "Vrrr... Brrr... Vrrr..."
}

// MobileAlertSong implements beep alert
type MobileAlertSong struct {
}

// Alert returns a alert string
func (a *MobileAlertSong) Alert() string {
	return "Белые розы, Белые розы. Беззащитны шипы..."
}

func main() {

	mobile := NewMobileAlert()

	result := mobile.Alert()
	result += mobile.Alert()

	mobile.SetState(&MobileAlertSong{})

	result += mobile.Alert()
	fmt.Println(result)
}
