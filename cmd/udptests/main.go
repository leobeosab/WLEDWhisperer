package main

import "github.com/leobeosab/wledwhisperer/pkg/wled"

func main() {
	s := &wled.Settings{
		Address:     "192.168.1.19",
		FromAddress: "192.168.1.15",
		Port:        ":21324",
		LedCount:    14,
	}

	wled.CreateConnection(s)

	leds := wled.SetStripLEDs(s.LedCount, 255, 255, 255)

	s.Connection.Write(wled.CreatePacket(5, leds))
}
