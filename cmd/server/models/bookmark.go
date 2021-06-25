package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
)

// Bookmark
type Bookmark struct {
	ID   string `json:"id"`
	Link string `json:"link"`
}

func GetBookmarks() (links []Bookmark) {

	path, err := filepath.Abs("./cmd/server/models/bookmark.json")
	if err != nil {
		log.Fatal(err)
	}

	// path := filepath.Join(dir, "bookmark.json")
	fileBytes, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &links)

	if err != nil {
		panic(err)
	}

	return links

}

func CreateBookmark(links []Bookmark) {

	linkBytes, err := json.Marshal(links)

	if err != nil {
		panic(err)
	}

	path, err := filepath.Abs("./cmd/server/models/bookmark.json")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(path, linkBytes, 0644)

	if err != nil {
		panic(err)
	}
}
