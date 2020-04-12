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

	leds := []wled.LED{wled.LED{Index: 5, Red: 255, Green: 255, Blue: 255}}

	s.Connection.Write(wled.CreatePacket(5, leds))
}
