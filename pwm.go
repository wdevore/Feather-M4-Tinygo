package main

import (
	"machine"
	"time"
)

func main() {

	machine.InitPWM() // initialize PWM
	red := machine.PWM{machine.D13}

	red.Configure() // start the pin's PWM

	duty, delta := 0, 1024

	for {

		red.Set(uint16(duty)) // set PWM duty cycle
		duty += delta

		if duty < 0 || duty > 65535 {
			delta *= -1
			duty += delta
		}

		time.Sleep(time.Millisecond * 5)
	}
}
