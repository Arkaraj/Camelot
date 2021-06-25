/*
Copyright © 2021 Arkaraj arkaraj2017@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"Camelot/cmd/server/models"
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var link string
var id string

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomStringId(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

// bookmarkCmd represents the bookmark command
var bookmarkCmd = &cobra.Command{
	Use:   "bookmark",
	Short: "Bookmark website links",
	Long:  `Stores and Saves website links`,
	Run: func(cmd *cobra.Command, args []string) {

		listFlag, _ := cmd.Flags().GetBool("list")

		if listFlag {

			links := models.GetBookmarks()

			fmt.Printf("Id \t bookmarked Link \n")

			for _, l := range links {
				fmt.Printf("%v \t %v \n", l.ID, l.Link)
			}

			return
		}

		url, _ := cmd.Flags().GetString("add")

		id, _ := cmd.Flags().GetString("rmv")

		if len(url) > 0 {

			list := models.Bookmark{
				ID:   RandomStringId(6),
				Link: url,
			}

			links := models.GetBookmarks()

			links = append(links, list)

			models.CreateBookmark(links)
			fmt.Println("Added To Bookmarks")
			return
		}

		if len(id) > 0 {
			list := []models.Bookmark{}

			links := models.GetBookmarks()

			for i := range links {
				if links[i].ID != id {
					list = append(list, links[i])
				}
			}
			models.CreateBookmark(list)
			fmt.Println("Removed Bookmark")
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(bookmarkCmd)

	bookmarkCmd.PersistentFlags().StringVarP(&link, "add", "a", "", "Add link to bookmarks")

	// -l will list the Bookmarks
	bookmarkCmd.PersistentFlags().BoolP("list", "l", false, "List all the Bookmarked websites")

	bookmarkCmd.PersistentFlags().StringVarP(&id, "rmv", "r", "", "Remove link from bookmarks by entering the Bookmark Id")
}
