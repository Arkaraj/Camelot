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
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var provider string
var listOfProviders string = `
- Google (default)
- Amazon
- bing
- codepen
- crunchyroll
- dribbble
- duckduckgo
- facebook
- flipkart
- giphy
- github
- hackernews
- imdb
- instagram
- linkedin
- medium
- netflix
- npm
- reddit
- soundcloud
- spotify
- stackoverflow
- steam
- pinterest
- quora
- twitchtv
- twitter
- uwatch
- wikipedia
- yahoo
- youtube
- 9anime`

var rawUrl string = ""

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Searches the given query in google(by default), providers from terminal",
	Long:  `Searches and opens new tab in Chrome by the entered keyword/query From Terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.Join(args[0:], "+")

		incog, _ := cmd.Flags().GetBool("incognito")

		var code []string
		switch runtime.GOOS {
		case "darwin":
			if incog {
				code = []string{"open", "-na", "Google Chrome", "--args", "--incognito"}
			} else {
				code = []string{"open"}
			}
		case "windows":
			if incog {
				code = []string{"cmd", "/c", "start", "chrome", "--incognito"}
			} else {
				code = []string{"cmd", "/c", "start"}
			}
		default:
			if incog {
				code = []string{"xdg-open"}
			} else {
				code = []string{"xdg-open"}
			}
		}
		// default
		queryString := "https://www.google.com/search?q="

		path, err := cmd.Flags().GetBool("output_path")

		if err != nil {
			os.Exit(1)
		}

		listFlag, _ := cmd.Flags().GetBool("list")

		if listFlag {
			fmt.Println(listOfProviders)
			return
		}

		if len(query) == 0 && len(rawUrl) == 0 {
			fmt.Println("Please Enter a Search Term.")
			return
		}

		prvd, _ := cmd.Flags().GetString("provider")
		if len(prvd) > 0 {
			switch provider {
			case "google":
				queryString = "https://www.google.com/search?q="
				break
			case "youtube":
				queryString = "https://www.youtube.com/results?search_query="
				break
			case "amazon":
				queryString = "https://www.amazon.com/s/?keywords="
				break
			case "github":
				queryString = "https://github.com/search?utf8=✓&q="
				break
			case "bing":
				queryString = "https://www.bing.com/search?q="
				break
			case "codepen":
				queryString = "http://codepen.io/search/pens?q="
				break
			case "crunchyroll":
				queryString = "https://www.crunchyroll.com/search?q="
				break
			case "dribbble":
				queryString = "https://dribbble.com/search?q="
				break
			case "duckduckgo":
				queryString = "https://duckduckgo.com/?q="
				break
			case "facebook":
				queryString = "https://www.facebook.com/search/top/?q="
				break
			case "flipkart":
				queryString = "http://www.flipkart.com/search?q="
				break
			case "giphy":
				queryString = "https://giphy.com/search/"
				break
			case "hackernews":
				queryString = "https://hn.algolia.com/?q="
				break
			case "imdb":
				queryString = "https://www.imdb.com/find?q="
				break
			case "instagram":
				queryString = "https://www.instagram.com/explore/tags/"
				break
			case "medium":
				queryString = "https://medium.com/search?q="
				break
			case "linkedin":
				queryString = "https://www.linkedin.com/search/results/all/?keywords="
				break
			case "netflix":
				queryString = "http://www.netflix.com/search/"
				break
			case "npm":
				queryString = "https://www.npmjs.com/search?q="
				break
			case "reddit":
				queryString = "https://www.reddit.com/search?q="
				break
			case "soundcloud":
				queryString = "https://soundcloud.com/search?q="
				break
			case "spotify":
				queryString = "https://play.spotify.com/search/"
				break
			case "stackoverflow":
				queryString = "https://stackoverflow.com/search?q="
				break
			case "steam":
				queryString = "https://store.steampowered.com/search/?term="
				break
			case "pinterest":
				queryString = "https://www.pinterest.com/search/pins/?q="
				break
			case "quora":
				queryString = "https://www.quora.com/search?q="
				break
			case "twitchtv":
				queryString = "https://www.twitch.tv/search?query="
				break
			case "twitter":
				queryString = "https://twitter.com/search?q="
				break
			case "uwatch":
				queryString = "https://www.uwatchfree.as/?s="
				break
			case "wikipedia":
				queryString = "https://en.wikipedia.org/wiki/Special:Search?search="
				break
			case "yahoo":
				queryString = "https://search.yahoo.com/search?p="
				break
			case "9anime":
				queryString = "https://9anime.to/search?keyword="
				break
			default:
				// os error stating no such providers
				fmt.Println("[Error] No Provider " + prvd + " found")
				return
				// queryString = "https://www.google.com/search?q="
			}
		}

		if path {
			fmt.Println("Link: " + queryString + query)
			return
		}

		rawUrl, _ = cmd.Flags().GetString("url")

		if len(rawUrl) > 0 {
			queryString = rawUrl
			query = ""
		}

		exe := exec.Command(code[0], append(code[1:], queryString+query)...)
		exe.Start()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.PersistentFlags().StringVarP(&provider, "provider", "p", "", "Pass in provider name (default google)")

	searchCmd.PersistentFlags().StringVarP(&rawUrl, "url", "u", "", "Pass in the web url")

	searchCmd.PersistentFlags().BoolP("output_path", "o", false, "View the query path")

	// -l will list the providers
	searchCmd.PersistentFlags().BoolP("list", "l", false, "List all the Providers")

	// opens chrome in incognito mode
	searchCmd.PersistentFlags().BoolP("incognito", "i", false, "Opens tab in chrome's incognito mode, only works for chrome")

}
