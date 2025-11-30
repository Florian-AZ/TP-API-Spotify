package controller

import (
	"ApiSpotify/fonction"
	"ApiSpotify/structure"
	"fmt"
	"html/template"
	"net/http"
)

var Token *string

// Page d'accueil
func Home(w http.ResponseWriter, r *http.Request) {
	Token = fonction.GetAccessToken()
	fmt.Println("Access Token:", *Token)
}

// Page albums Damso
func AlbumDamso(w http.ResponseWriter, r *http.Request) {
	Token = fonction.GetAccessToken()
	albums := fonction.GetAlbums(*Token, "4tZwfgrHOc3mvqYlEYSvVi")
	if r.Method == http.MethodPost {
		for _, i := range A.albumsItems {
			data := structure.data{
				Name:        albums[0].Artists[0].Name,
				ReleaseDate: albums[0].ReleaseDate,
				TotalTracks: len(albums),
				Items:       []albums,
			}
			AHTML.Data = append(AHTML.Data, data)
		}
	data := PageData{
		Name:        albums[0].Artists[0].Name,
		ReleaseDate: albums[0].ReleaseDate,
		TotalTracks: len(albums),
		Items:       albums,
	}
	tmpl := template.Must(template.ParseFiles("./template/damso.html"))
	tmpl.Execute(w, data) // albums est une slice d'Album
}

// Page track Maladresse Laylow
func TrackMaladresse(w http.ResponseWriter, r *http.Request) {
	Token = fonction.GetAccessToken()
	track := fonction.GetTracks(*Token, "67Pf31pl0PfjBfUmvYNDCL")
	if r.Method == http.MethodPost {
		data := structure.LaylowData{
			Name:         track.Name,
			Artist:       track.Artists[0].Name,
			Album:        track.Album.Name,
			Release_Date: track.Album.Release_Date,
		}
		tmpl := template.Must(template.ParseFiles("./template/laylow.html"))
		tmpl.Execute(w, data)
	}
}
