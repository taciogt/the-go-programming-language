package xkcd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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

	for i := 1; err == nil; i++ {
		if i == http.StatusNotFound {
			continue // comic #404 do not exists
		}

		comicInfo, err = GetComicInfo(i)
		if err != nil {
			if errors.Is(err, ErrNotFound{}) {
				log.Printf("last comic found (#%d)\n", i-1)
			} else {
				if _, err := fmt.Fprintf(os.Stdout, "error=%s", err); err != nil {
					panic(err)
				}
			}
			close(ch)
			return
		}

		ch <- comicInfo
	}
	close(ch)
}

const indexFileName = "comics-index.json"

func BuildIndex() error {
	comics := make([]ComicInfo, 0)
	for c := range ListAllComics() {
		comics = append(comics, c)

		if c.Num%100 == 1 {
			log.Printf("indexing comic #%d...\n", c.Num)
		}
	}

	comicsSerialized, err := json.Marshal(comics)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(indexFileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	if _, err = io.Copy(f, bytes.NewReader(comicsSerialized)); err != nil {
		return err
	}

	return nil
}

func LoadIndex() ([]ComicInfo, error) {
	data, err := os.ReadFile(indexFileName)
	if err != nil {
		return nil, err
	}
	comics := make([]ComicInfo, 0)
	if err := json.Unmarshal(data, &comics); err != nil {
		return nil, err
	}

	return comics, nil
}

func SearchTerm(comics []ComicInfo, searchTerm string) ([]ComicInfo, error) {
	result := make([]ComicInfo, 0)
	searchTerm = strings.ToLower(searchTerm)

	for i, c := range comics {
		if strings.Contains(strings.ToLower(c.Title), searchTerm) {
			result = append(result, comics[i])
		}
	}
	return result, nil
}
