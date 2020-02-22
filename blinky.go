package main

import (
	"machine"
	"time"
)

func main() {
	led5 := machine.D5
	led5.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led6 := machine.D6
	led6.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led6.High()
		led5.Low()
		time.Sleep(time.Millisecond * 200)

		led5.High()
		led6.Low()
		time.Sleep(time.Millisecond * 150)
	}
}
