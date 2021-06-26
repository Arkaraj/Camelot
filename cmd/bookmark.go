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
	"path/filepath"

	"github.com/spf13/cobra"
)

var link string
var dp string
var id int

// bookmarkCmd represents the bookmark command
var bookmarkCmd = &cobra.Command{
	Use:   "bookmark",
	Short: "Bookmark website links",
	Long:  `Stores and Saves website links`,
	Run: func(cmd *cobra.Command, args []string) {

		dirpath,_ := filepath.Abs("./cmd/server/models/bookmark.json")

		dp, _ = cmd.Flags().GetString("directorypath")

		if (len(dp) > 0) {
			dirpath = dp+"/cmd/server/models/bookmark.json"
		}

		// path,_ := filepath.Abs(dirpath)

		listFlag, _ := cmd.Flags().GetBool("list")

		if listFlag {
			links := models.GetBookmarks(dirpath)

			fmt.Printf("Id \t bookmarked Link \n")

			for _, l := range links {
				fmt.Printf("%v \t %v \n", l.ID, l.Link)
			}

			return
		}

		url, _ := cmd.Flags().GetString("add")

		id, _ := cmd.Flags().GetInt("rmv")

		if len(url) > 0 {
			// path,_ := filepath.Abs(".")
			links := models.GetBookmarks(dirpath)

			list := models.Bookmark{
				ID:   len(links) + 1,
				Link: url,
			}

			links = append(links, list)

			models.CreateBookmark(links, dirpath)
			fmt.Println("Added To Bookmarks")
			return
		}

		if id != 0 {
			list := []models.Bookmark{}

			// path,_ := filepath.Abs("./bookmark.json")
			links := models.GetBookmarks(dirpath)

			for i := range links {
				if links[i].ID != id {
					list = append(list, links[i])
				}
			}
			models.CreateBookmark(list, dirpath)
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

	bookmarkCmd.PersistentFlags().StringVarP(&dp, "directorypath", "d", "", "Enter the directory path where Camelot is")

	bookmarkCmd.PersistentFlags().IntVarP(&id, "rmv", "r", 0, "Remove link from bookmarks by entering the Bookmark Id")

}
