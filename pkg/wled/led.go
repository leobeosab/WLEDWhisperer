package wled

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

	return data
}

// Refactor later
func DimLEDs(r byte, g byte, b byte, brightness float32) []byte {
	if brightness > 1.0 {
		brightness = 1.0
	}

	r = byte(float32(r) * brightness)
	g = byte(float32(g) * brightness)
	b = byte(float32(b) * brightness)

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

func SetPercentageLEDs(p float32, l int, r byte, g byte, b byte) []LED {
	// LED Percentage
	lP := p * float32(l)
	// Fully Lit LED Number
	fL := int(lP)
	// Brightness of last LED
	lB := lP - float32(fL)

	leds := make([]LED, l)

	for i := 0; i < l; i++ {
		brightness := float32(1.0)

		if i == fL {
			brightness = lB
		}
		if i > fL {
			brightness = 0
		}

		// This could be made prettier
		colors := DimLEDs(r, g, b, brightness)
		leds = append(leds, LED{
			Index: byte(i),
			Red:   colors[0],
			Green: colors[1],
			Blue:  colors[2],
		})
	}

	return leds
}
