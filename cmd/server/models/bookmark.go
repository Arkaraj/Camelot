package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Bookmark
type Bookmark struct {
	ID   int    `json:"id"`
	Link string `json:"link"`
}

func GetBookmarks(path string) (links []Bookmark) {

	// path := filepath.Join(dir, "bookmark.json")
	fileBytes, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Command only works in Camelot directory")
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &links)

	if err != nil {
		panic(err)
	}

	return links

}

func CreateBookmark(links []Bookmark, path string) {

	linkBytes, err := json.Marshal(links)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path, linkBytes, 0644)

	if err != nil {
		panic(err)
	}
}
