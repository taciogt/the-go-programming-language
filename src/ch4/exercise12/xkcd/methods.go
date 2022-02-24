package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ComicInfo struct {
	Num        int    `json:"num"`
	Month      int    `json:"month,string"`
	Year       int    `json:"year,string"`
	Link       string `json:"link"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Alt        string `json:"alt"`
	Image      string `json:"img"`
	Title      string `json:"title"`
	Day        int    `json:"day,string"`
	Transcript string `json:"transcript"`
}

func GetComicInfo(comicNumber int) (comicInfo ComicInfo, err error) {
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", comicNumber)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		err = NewErrNotFound(comicNumber)
		return
	} else if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpect response status code: %s", resp.Status)
	}

	if err = json.NewDecoder(resp.Body).Decode(&comicInfo); err != nil {
		return
	}

	return
}
