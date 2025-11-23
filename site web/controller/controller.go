package controller

import (
	//"ApiSpotify/fonction"
	//"ApiSpotify/structure"
	"fmt"
	"html/template"
	"net/http"
)

var clientID = "<4ba3c025599744ef9bc882e88bf3b5cc>"
var clientSecret = "<6ad26881c3bb4a01a83da348a3714030>"

// Page d'accueil
func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./template/index.html"))

	data := map[string]string{
		"DamsoURL":  "/album/damso",
		"LaylowURL": "/track/laylow",
	}
	tmpl.Execute(w, data)
}

// Page albums Damso
func AlbumDamso(w http.ResponseWriter, r *http.Request) {
	token, err := GetAccessToken(clientID, clientSecret)
	if err != nil {
		fmt.Fprintln(w, "Erreur token :", err)
		return
	}

	albums, err := GetAlbums(token, "0eKTWbC8N7IV9lEdhT4sq2")
	if err != nil {
		fmt.Fprintln(w, "Erreur albums :", err)
		return
	}

	tmpl := template.Must(template.ParseFiles("./template/damso.html"))
	tmpl.Execute(w, albums) // albums est une slice d'Album
}

// Page track Maladresse Laylow
func TrackMaladresse(w http.ResponseWriter, r *http.Request) {
	token, err := GetAccessToken(clientID, clientSecret)
	if err != nil {
		fmt.Fprintln(w, "Erreur token :", err)
		return
	}

	track, err := GetTrackInfo(token, "<ID_TRACK_MALADRESSE>")
	if err != nil {
		fmt.Fprintln(w, "Erreur track :", err)
		return
	}

	tmpl := template.Must(template.ParseFiles("./template/laylow.html"))
	tmpl.Execute(w, track)
}
