package main

import (
	"machine"
	"time"
)

const (
	max_led_num = 8
)

func main() {

	delay := func(t int64) { // implicit time delay function
		time.Sleep(time.Duration(1000000 * t))
	}

	var leds = [max_led_num]machine.Pin{
		machine.D12,
		machine.D11,
		machine.D10,
		machine.D9,
		machine.D6,
		machine.D5,
		machine.D21,
		machine.D22,
	} // array of pins

	button := machine.D4
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	for i := 0; i < len(leds); i++ { // setup leds
		leds[i].Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	for {

		if button.Get() {
			// blink one by one forward
			for i := 0; i < len(leds); i++ {
				leds[i].High()
				delay(75)
				leds[i].Low()
			}

			// blink one by one backward
			for i := len(leds) - 1; i >= 0; i-- {
				leds[i].High()
				delay(75)
				leds[i].Low()
			}
		}
	}
}
