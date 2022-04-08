package xkcd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	defer func() { _ = resp.Body.Close() }()

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

func ListAllComics() chan ComicInfo {
	comics := make(chan ComicInfo)
	go listComics(comics)
	return comics
}

func listComics(ch chan ComicInfo) {
	var comicInfo ComicInfo
	var err error

	//for i := 1; err == nil; i++ {
	for i := 1; err == nil && i < 100; i++ {
		comicInfo, err = GetComicInfo(i)
		if err != nil {
			//fmt.Printf("error=%s", err)
			if _, err := fmt.Fprintf(os.Stdout, "error=%s", err); err != nil {
				panic(err)
			}
			close(ch)
			return
		}
		ch <- comicInfo
	}
	close(ch)
}

func BuildIndex() error {
	comics := make([]ComicInfo, 0)
	for c := range ListAllComics() {
		comics = append(comics, c)
	}

	comicsSerialized, err := json.Marshal(comics)
	if err != nil {
		return err
	}
	f, err := os.OpenFile("comics-index.json", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	if _, err = io.Copy(f, bytes.NewReader(comicsSerialized)); err != nil {
		return err
	}

	return nil
}
