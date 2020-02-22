package main

import (
	"machine"
	"time"
)

const (
	max_led_num = 8
)

func main() {
	println("Starting app...")
	machine.InitADC()

	sensor := machine.ADC{machine.A2}
	sensor.Configure()

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

	for i := 0; i < len(leds); i++ { // setup leds
		leds[i].Configure(machine.PinConfig{Mode: machine.PinOutput})
	}

	// val := sensor.Get()
	// vals := (val - 2000) / 500
	// // vals 0 -> 125
	// // 15.6 bands
	// mod := int(float64(vals) / 17)

	// // prevMod := mod
	// // Note: The pot is a 10K
	for {
		// 2650 -> 59000 = 56350
		val := uint16(sensor.Get())
		println("val: ", val)
		flt := float64(val)
		println("flt: ", flt)

		// vals := (val - 2650) / 500
		// vals 0 -> 125
		// 15.6 bands
		// md := int(val / 17.0)
		md := 2

		// 	// if prevMod != mod {
		for i := 0; i < len(leds); i++ {
			leds[i].Low()
		}

		leds[md].High()
		// prevMod = mod
		// }

		// 	// println("ADC: ", val)
		// l := "-"
		// for i := uint16(0); i < uint16(vals); i++ {
		// 	l += "."
		// }
		// x := fmt.Sprintf("%s %d (%d) m:%d", l, val, vals, md)
		// x := fmt.Sprintf("%s %d (%d)", l, val, vals)
		// println(x)
		time.Sleep(time.Millisecond * 100)
	}
}
