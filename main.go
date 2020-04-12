package main

import (
	"fmt"
	"log"
	"net"
)

// Turns led 3 & 6 purple
func CreatePacket(index int, red int, green int, blue int) []byte {
	return []byte{1, 1, 3, 255, 0, 255, 6, 255, 0, 255}
}

func main() {
	data := CreatePacket(3, 0, 200, 20)
	port := ":21324"
	address := "192.168.1.19"

	fmt.Println(data)

	broadcast_addr, err := net.ResolveUDPAddr("udp", address+port)
	if err != nil {
		log.Println(err)
	}

	local_addr, err := net.ResolveUDPAddr("udp", "192.168.1.15"+port)
	if err != nil {
		fmt.Println(err)
	}

	connection, err := net.DialUDP("udp", local_addr, broadcast_addr)
	if err != nil {
		fmt.Println(err)
	}
	defer connection.Close()

	connection.Write(data)

}
