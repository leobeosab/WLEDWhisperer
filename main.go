package main

import (
	"fmt"
	"log"
	"net"
)

func CreateWARLSByte(byteN int, n int) byte {
	return byte(byteN + n*4)
}

func CreatePacket(index int, red int, green int, blue int) []byte {
	return []byte{1, 1, CreateWARLSByte(2, index), CreateWARLSByte(3, red), CreateWARLSByte(4, green), CreateWARLSByte(5, blue)}
}

func GetLocalIP() string {
	var localIP string
	addr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("GetLocalIP in communication failed")
		return "localhost"
	}
	for _, val := range addr {
		if ip, ok := val.(*net.IPNet); ok && !ip.IP.IsLoopback() {
			if ip.IP.To4() != nil {
				localIP = ip.IP.String()
			}
		}
	}
	return localIP
}

func main() {
	data := CreatePacket(3, 0, 20, 20)
	port := ":21324"
	address := "192.168.1.19"

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
