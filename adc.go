package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	machine.InitADC()

	sensor := machine.ADC{machine.A2}
	sensor.Configure()

	// Note: The pot is a 10K
	for {
		val := sensor.Get()
		vals := (val - 2000) / 500
		// println("ADC: ", val)
		l := ""
		for i := uint16(0); i < vals; i++ {
			l += "."
		}
		x := fmt.Sprintf("%s %d (%d)", l, val, vals)
		println(x)
		time.Sleep(time.Millisecond * 100)
	}
}
