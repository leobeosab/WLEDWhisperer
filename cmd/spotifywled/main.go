package main

import (
	"log"
	"time"

	"github.com/leobeosab/wledwhisperer/pkg/spotify"
	"github.com/leobeosab/wledwhisperer/pkg/wled"
	sapi "github.com/zmb3/spotify"
)

var (
	s *wled.Settings
)

func main() {
	// Temporary
	s = &wled.Settings{
		Address:     "192.168.1.19",
		FromAddress: "192.168.1.15",
		Port:        ":21324",
		LedCount:    14,
	}

	wled.CreateConnection(s)
	spotify.SetupClient(ProgressLoop)
}

// Loop working with LEDS
// Passed into client Setup
// Ran when auth succeeds
func ProgressLoop(client *sapi.Client) {
	for {
		time.Sleep(200 * time.Millisecond)
		p, err := client.PlayerCurrentlyPlaying()
		if err != nil {
			log.Println(err)
			continue
		}

		prog := float32(p.Progress) / float32(p.Item.Duration)

		// progess (.xx) * number of leds should light up the correct percentage of leds :)
		data := wled.SetPercentageLEDs(prog, s.LedCount, 125, 255, 125)
		s.Connection.Write(wled.CreatePacketWithBrightness(5, data, 1.0))
	}
}
