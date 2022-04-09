package omdbapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Client struct {
	ApiKey string
}

const baseUrl = "https://www.omdbapi.com/"

func NewClient(apiKey string) Client {
	return Client{ApiKey: apiKey}
}

func (c Client) SearchMovie(searchTerm string) ([]Movie, error) {
	u := fmt.Sprintf("%s?apikey=%s&s=%s", baseUrl, c.ApiKey, url.QueryEscape(searchTerm))
	fmt.Printf("url: %s\n", u)

	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	var moviesResponse SearchMovieResponse
	if err = json.NewDecoder(resp.Body).Decode(&moviesResponse); err != nil {
		return nil, err
	}

	return moviesResponse.Movies, nil
}

func (c Client) DownloadPoster(movie Movie) (string, error) {
	resp, err := http.Get(movie.Poster)
	if err != nil {
		return "", err
	}
	defer func() { resp.Body.Close() }()

	filename := fmt.Sprintf("%s.jpg", strings.ToLower(movie.Title))
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", err
	}

	return filename, nil
}
