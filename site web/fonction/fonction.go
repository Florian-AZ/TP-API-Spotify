package fonction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func GetAccessToken(clientID, clientSecret string) (string, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var token SpotifyToken
	if err := json.Unmarshal(body, &token); err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func main() {
	// Remplace par tes identifiants Spotify
	clientID := "<4ba3c025599744ef9bc882e88bf3b5cc>"
	clientSecret := "<6ad26881c3bb4a01a83da348a3714030>"

	token, err := GetAccessToken(clientID, clientSecret)
	if err != nil {
		fmt.Println("Erreur pour récupérer le token :", err)
		return
	}
	fmt.Println("Access Token :", token)

	// récupérer les albums de Damso
	artistID := "0eKTWbC8N7IV9lEdhT4sq2" // ID Spotify de Damso
	urlAPI := fmt.Sprintf("https://api.spotify.com/v1/artists/%s/albums", artistID)

	// Client HTTP
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, _ := http.NewRequest("GET", urlAPI, nil)

	req.Header.Add("Authorization", "Bearer "+token)

	req.Header.Add("User-Agent", "Ynov Campus Cours")

	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Oups une erreur est survenue:", err)
		return
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var albums AlbumsResponse
	if err := json.Unmarshal(body, &albums); err != nil {
		fmt.Println("Erreur décodage JSON :", err)
		return
	}

	// Affichage des albums
	for i, album := range albums.Items {
		fmt.Printf("%d. %s | Date de sortie : %s | Nombre de titres : %d\n", i+1, album.Name, album.ReleaseDate, album.TotalTracks)
	}
}
