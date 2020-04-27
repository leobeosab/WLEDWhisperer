# WLEDWhisperer Readme

Go communication with WLED plus various examples and integrations, just for fun

## Dependencies | Required Parts

- [zmb2/spotify](https://github.com/zmb3/spotify)
- A compatible [Aircoookie/WLED](https://github.com/Aircoookie/WLED) device, esp32, esp8266 etc
- A individually addressable RGB led strip

## Demo

### Forza
[Youtube Video](https://youtu.be/UF4cU1qSNPA)
![WLEDWhisperer%20Readme/51DB74C0-9F97-4141-A57D-91AE59E54EE4.gif](Readme/51DB74C0-9F97-4141-A57D-91AE59E54EE4.gif)

### Spotify

![WLEDWhisperer%20Readme/1BB453F6-0574-4AE8-8A8A-1C1B8B8F847F.gif](Readme/1BB453F6-0574-4AE8-8A8A-1C1B8B8F847F.gif)

## Running Locally

`go get [github.com/leobeosab/WLEDWhisperer](https://github.com/leobeosab/WLEDWhisperer)`

### WLED Strip setup

This will change in the future but for now you have to go into the main.go file of the program you want to use ex `cmd/udptest/main.go` and change the settings struct to your specific computer's and LED strip's ip and ports as well as the number of LEDs on the strip

```go
s := &wled.Settings{
        Address:     "192.168.1.19", // IP of WLED device
        FromAddress: "192.168.1.15", // IP of local machine on same network as WLED
        Port:        ":21324", // Port of WLED device default is 21324
        ledLedCount:    14, // Number of LEDs on strip connectied to WLED device
}
```

### Spotify Progress Bar Specific Setup

Setup a Spotify App with your Callback url set to [http://localhost:8888/callback](http://localhost:8888/callback) more info [here](https://github.com/zmb3/spotify)

Make sure you have your app's credentials as environment variables ex:

```bash
export SPOTIFY_ID=[yourid]
export SPOTIFY_SECRET=[yoursecret]
```

If everything is working correctly you should be able to just do the following: 

`go run cmd/spotifywled/main.go` or `go install ./... && spotifywled`

and go to the address printed in the console + authorize your app and if music is playing you should see your led strip lit up relative to how far into a song yo are

### Forza Specific Setup

Open Forza, go to hud and gameplay, turn data output on and set ip address and port to the ip of the device running WLED whisperer and the port you set (default is 8080)