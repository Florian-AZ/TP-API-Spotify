package fonction

import (
	"ApiSpotify/structure"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func GetAccessToken() *string {

	// URL de L'API pour obtenir le token
	urlApi := "https://accounts.spotify.com/api/token"

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", "4ba3c025599744ef9bc882e88bf3b5cc")
	data.Set("client_secret", "6ad26881c3bb4a01a83da348a3714030")

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodPost, urlApi, strings.NewReader(data.Encode()))
	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errReq.Error())
	}

	// Ajout d'une métadonnée dans le header, User_Agent permet d'identifier l'application, système ....
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errResp.Error())
		return (nil)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData structure.SpotifyToken

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	return (&decodeData.AccessToken)
}

func GetAlbums(token, artistID string) structure.AlbumsResponse {
	urlApi := fmt.Sprintf("https://api.spotify.com/v1/artists/%s/albums", artistID)
	// Client HTTP
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errReq.Error())
	}

	// Ajout d'une métadonnée dans le header, User_Agent permet d'identifier l'application, système ....
	req.Header.Add("Authorization", "Bearer "+token)

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errResp.Error())
		return (structure.AlbumsResponse{})
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData structure.AlbumsResponse

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	return (decodeData)
}

func GetTracks(token, trackID string) structure.TracksResponse {
	urlApi := fmt.Sprintf("https://api.spotify.com/v1/tracks/" + trackID)
	// Client HTTP
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}
	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errReq.Error())
	}
	// Ajout d'une métadonnée dans le header, User_Agent permet d'identifier l'application, système ....
	req.Header.Add("Authorization", "Bearer "+token)
	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errResp.Error())
		return (structure.TracksResponse{})
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errBody.Error())
	}
	// Déclaration de la variable qui va contenir les données
	var decodeData structure.TracksResponse
	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)
	// Affichage des données
	return (decodeData)
}
