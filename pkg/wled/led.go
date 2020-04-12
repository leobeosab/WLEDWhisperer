package wled

import "log"

type LED struct {
	Index byte
	Red   byte
	Green byte
	Blue  byte
}

// Create packet according to WARLS format
func CreatePacket(timeout byte, leds []LED) []byte {
	data := make([]byte, 2)

	// Protocol
	data[0] = 1
	// timeout in seconds before wled device returns to normal mode
	data[1] = timeout

	// LED bytes are index,r,g,b.. reapeat until end
	for _, l := range leds {
		data = append(data, l.Index, l.Red, l.Green, l.Blue)
	}

	log.Println(data)

	return data
}

func SetStripLEDs(l int, r byte, g byte, b byte) []LED {
	leds := make([]LED, l)
	for i := 0; i < l; i++ {
		leds[i] = LED{
			Index: byte(i),
			Red:   r,
			Green: g,
			Blue:  b,
		}
	}

	return leds
}
