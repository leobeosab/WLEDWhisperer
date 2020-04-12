package main

import "github.com/leobeosab/wledwhisperer/pkg/wled"

// Turns led 3 & 6 purple
func CreatePacket(index int, red int, green int, blue int) []byte {
	return []byte{1, 1, 3, 255, 0, 255, 6, 255, 0, 255, 10, 255, 0, 0}
}

func main() {
	data := CreatePacket(3, 0, 200, 20)
	s := &wled.Settings{
		Address:     "192.168.1.19",
		FromAddress: "192.168.1.15",
		Port:        ":21324",
		LedCount:    14,
	}

	wled.CreateConnection(s)

	s.Connection.Write(data)
}
