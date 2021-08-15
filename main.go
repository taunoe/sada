// Project: https://github.com/taunoe/sada
// Started: 14.08.2021 Tauno erik
package main

import (
	"machine"
	"time"
	// https://github.com/tinygo-org/drivers/blob/release/examples/shiftregister/
	//"tinygo.org/x/drivers/shiftregister"
)

var (
	latch = machine.D12 // D12 Pin latch connected to ST_CP of 74HC595 (12)
	clock = machine.D11 // D11 Pin clock connected to SH_CP of 74HC595 (11)
	out   = machine.D10 // D10 Pin data connected to DS of 74HC595 (14)
)

func shift_out(mask uint8) {
	for i := 0; i < int(8); i++ {
		clock.Low()
		out.Set(mask&1 != 0)
		mask = mask >> 1
		clock.High()
	}
}

func main() {
	latch.Configure(machine.PinConfig{Mode: machine.PinOutput})
	clock.Configure(machine.PinConfig{Mode: machine.PinOutput})
	out.Configure(machine.PinConfig{Mode: machine.PinOutput})
	latch.High()

	// Loop
	for {
		// Column (-) and row (+) pins:
		// c1, c2, c3, c4, c5, c6, c7, c8
		// c9,c10, R1, R2, R3, R4, R5, R6
		// R7, R8, R9,R10, nc, nc, nc, nc
		var matrix = []uint8{
			0b10000000,
			0b00100000,
			0b00010000,
		}

		latch.Low()
		shift_out(matrix[2])
		shift_out(matrix[1])
		shift_out(matrix[0])
		latch.High()

		delay(500)

		// Some fun with masks
		//for _, pattern := range patterns {
		//	d.WriteMask(pattern)
		//	delay(100)
		//}
		//delay(500)
		//d.WriteMask(0x00)
	}

	/*
		led := machine.D13

		led.Configure(machine.PinConfig{Mode: machine.PinOutput})

		// Loop
		for {
			led.Low()
			delay(600)

			led.High()
			delay(600)

			led.Low()
			delay(300)

			led.High()
			delay(300)
		}
	*/
}

// Delay function
func delay(delay uint16) {
	time.Sleep(time.Duration(1000000 * uint32(delay)))
}
