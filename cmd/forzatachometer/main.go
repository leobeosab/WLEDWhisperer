package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math"
	"net"
	"time"

	"github.com/leobeosab/wledwhisperer/pkg/wled"
)

func main() {
	fmt.Println("Forza Telemetry Server")

	var hostIp string
	var wledIp string
	var ledCount int

	// This is passed as a flag because I don't know what interface to grab the IP from
	flag.StringVar(&hostIp, "host-ip", "", "IP address of this machine")
	flag.StringVar(&wledIp, "wled-ip", "", "IP address of the WLED strip")
	flag.IntVar(&ledCount, "led-coount", 0, "Number of leds in the light strip")

	flag.Parse()

	if hostIp == "" || wledIp == "" || ledCount <= 0 {
		panic("You need to pass --led-count, --host-ip, and --wled-ip")
	}

	// Temporary
	s := &wled.Settings{
		Address:     wledIp,
		FromAddress: hostIp,
		Port:        ":21324",
		LedCount:    ledCount,
	}
	wled.CreateConnection(s)
	defer s.Connection.Close()

	port := hostIp + ":8080"
	protocol := "udp"

	udpAddr, err := net.ResolveUDPAddr(protocol, port)
	if err != nil {
		panic(err)
	}

	log.Printf("Server starting at %s, with protocol: %s", port, protocol)

	udpConn, err := net.ListenUDP(protocol, udpAddr)
	if err != nil {
		panic(err)
	}

	oldTime := time.Now()

	for {
		var buff [2048]byte
		_, err := udpConn.Read(buff[0:])
		if err != nil {
			log.Println(err)
			continue
		}

		// Byte 1 is on or off
		//p := buff[0]

		EngineMaxRPM := ByteTo32BitFloat(buff[8:12])
		//IdleRPM := ByteTo32BitFloat(buff[12:16])
		CurrentRPM := ByteTo32BitFloat(buff[16:20])

		//fmt.Printf("Paused: %v Max Engine RPM: %f Idle RPM: %f Current RPM: %f \n", p == 0, EngineMaxRPM, IdleRPM, CurrentRPM)
		//fmt.Println(hex.EncodeToString(buff[0:24]))

		prog := CurrentRPM / EngineMaxRPM
		prog += 0.1

		// Run 10 times a second
		if time.Since(oldTime) > time.Duration(20*time.Millisecond) {
			data := wled.SetPercentageLEDs(prog, s.LedCount, byte(prog*255), byte((1.0-prog)*255), 0)
			packet := wled.CreatePacketWithBrightness(255, data, 1.0)
			s.Connection.Write(packet)
			fmt.Println(prog)

			oldTime = time.Now()
		}
	}
}

// ByteTo32BitFloat takes a slice of 4 bytes and turns into an int32
func ByteTo32BitFloat(data []byte) float32 {
	bits := binary.LittleEndian.Uint32(data)
	return math.Float32frombits(bits)
}
