package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ApiData struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func Refreshtoken() {

	// URL de L'API
	urlApi := "https://accounts.spotify.com/api/token"

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodPost, urlApi, nil)
	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errReq.Error())
	}

	// Ajout d'une métadonnée dans le header, User_Agent permet d'identifier l'application, système ....
	req.Header.Add("Content-Type", "application/x-www-form-urlencodedpair_29b80f29b47b4ff286b64740d5f48bbf")

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errResp.Error())
		return
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
	var decodeData ApiData

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	// Affichage des données
	fmt.Println(decodeData)
}

// Home handler simple
func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./template/index.html")
}

func Damso(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./template/damso.html")
}

func Laylow(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./template/laylow.html")
}
