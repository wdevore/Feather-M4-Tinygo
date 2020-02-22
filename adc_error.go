package main

import (
	"fmt"
	"time"
)

func main() {
	// machine.InitADC()

	// sensor := machine.ADC{machine.A2}
	// sensor.Configure()

	i := 0.0
	for {
		// val := sensor.Get()
		// println("val: ", val)
		// flt := i
		s := fmt.Sprintf("%f", i) // "i" causes it to fail.
		println("flt: ", s)
		i += 1.0
		time.Sleep(time.Millisecond * 100)
	}
}
