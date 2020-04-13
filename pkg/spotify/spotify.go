package spotify

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zmb3/spotify"
)

const redirectURI = "http://localhost:8888/callback"

var (
	auth  = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadCurrentlyPlaying, spotify.ScopeUserReadPlaybackState)
	ch    = make(chan *spotify.Client)
	state = ""
)

func SetupClient(loop func(*spotify.Client)) {
	var client *spotify.Client

	http.HandleFunc("/callback", CompleteAuth)

	go func() {
		url := auth.AuthURL(state)
		fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

		client = <-ch

		user, err := client.CurrentUser()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("You are logged in as: ", user.ID)
		loop(client)
	}()

	http.ListenAndServe(":8888", nil)
}

func CompleteAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := auth.NewClient(tok)
	w.Header().Set("Content-Type", "text/html")
	ch <- &client
}
