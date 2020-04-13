package wled

import (
	"log"
	"net"
)

type Settings struct {
	Address     string
	FromAddress string
	Port        string
	LedCount    int
	Connection  *net.UDPConn
}

func CreateConnection(settings *Settings) {
	broadcast_addr, err := net.ResolveUDPAddr("udp", settings.Address+settings.Port)
	if err != nil {
		log.Printf("Error resolving target address with %s:%s", settings.Address, settings.Port)
		log.Println(err)
	}

	local_addr, err := net.ResolveUDPAddr("udp", settings.FromAddress+settings.Port)
	if err != nil {
		log.Printf("Error resolving target address with %s:%s", settings.FromAddress, settings.Port)
		log.Println(err)
	}

	conn, err := net.DialUDP("udp", local_addr, broadcast_addr)
	if err != nil {
		log.Println("Error dialing udp")
		log.Println(err)
	}

	settings.Connection = conn
}
