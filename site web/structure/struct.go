package structure

type SpotifyToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type Album struct {
	Name        string `json:"name"`
	ReleaseDate string `json:"release_date"`
	TotalTracks int    `json:"total_tracks"`
}

type AlbumsResponse struct {
	Items []Album `json:"items"`
}

// Artiste pour track
type Artist struct {
	Name string `json:"name"`
}

// Album pour track
type TrackAlbum struct {
	Name        string `json:"name"`
	ReleaseDate string `json:"release_date"`
	Images      []struct {
		URL string `json:"url"`
	} `json:"images"`
}
