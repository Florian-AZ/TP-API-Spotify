package structure

type SpotifyToken struct {
	AccessToken string `json:"access_token"`
}

type Album struct {
	Name          string `json:"name"`
	ReleaseDate   string `json:"release_date"`
	TotalTracks   int    `json:"total_tracks"`
	External_urls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
}

type AlbumsResponse struct {
	Items []Album `json:"items"`
}

// Artiste pour track
type Artist struct {
	Name string `json:"name"`
}

type TrackAlbum struct {
	Name        string `json:"name"`
	ReleaseDate string `json:"release_date"`
	Images      []struct {
		URL string `json:"url"`
	} `json:"images"`
}

type TracksResponse struct {
	Name    string `json:"name"` // Nom du track maladresse
	Artists []struct {
		Name string `json:"name"` // nom de l'artiste
	} `json:"artists"`
	Album struct {
		Images []struct {
			URL string `json:"url"` // url de l'image de l'album
		} `json:"images"`
		Name         string `json:"name"`         // nom de l'album
		Release_Date string `json:"release_date"` // date de sortie de l'album
	} `json:"album"`
	External_urls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
}

type DamsoData struct {
	Name        string
	ReleaseDate string
	TotalTracks int
	Items       []Album
}

type LaylowData struct {
	Name          string
	Artist        string
	Album         string
	Release_Date  string
	External_urls string
}
