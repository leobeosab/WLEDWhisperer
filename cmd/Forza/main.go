package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Forza Telemetry Server")

	port := ":8080"
	protocol := "udp"

	udpAddr, err := net.ResolveUDPAddr(protocol, port)
	if err != nil {
		panic(err)
	}

	log.Printf("Server starting at localhost%s, with protocol: %s", port, protocol)

	udpConn, err := net.ListenUDP(protocol, udpAddr)
	if err != nil {
		panic(err)
	}

	for {
		var buff [2048]byte
		n, err := udpConn.Read(buff[0:])
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(hex.EncodeToString(buff[0:n]))
	}
}
