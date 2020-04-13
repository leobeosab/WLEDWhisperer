package wled

import "log"

type LED struct {
	Index byte
	Red   byte
	Green byte
	Blue  byte
}

func CreatePacket(timeout byte, leds []LED) []byte {
	return CreatePacketWithBrightness(timeout, leds, 1)
}

// Create packet according to WARLS format
func CreatePacketWithBrightness(timeout byte, leds []LED, b float32) []byte {
	if b > 1.0 {
		b = 1.0
	}

	data := make([]byte, 2)

	// Protocol
	data[0] = 1
	// timeout in seconds before wled device returns to normal mode
	data[1] = timeout

	// LED bytes are index,r,g,b.. reapeat until end
	for _, l := range leds {
		data = append(data, l.Index)
		data = append(data, DimLEDs(l.Red, l.Green, l.Blue, b)...)
	}

	log.Println(data)

	return data
}

func DimLEDs(r byte, g byte, b byte, bright float32) []byte {
	r = byte(float32(r) * bright)
	g = byte(float32(g) * bright)
	b = byte(float32(b) * bright)

	return []byte{r, g, b}
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
