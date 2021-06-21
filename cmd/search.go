/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

var provider string;

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Google Searches the given query from terminal",
	Long: `Searches and opens new tab in Chrome by the entered keyword/query From Terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		query := strings.Join(args[0:],"+")
		var code []string
		switch runtime.GOOS {
		case "darwin":
			code = []string{"open"}
		case "windows":
			code = []string{"cmd", "/c", "start"}
		default:
			code = []string{"xdg-open"}
		}
		queryString := "https://www.google.com/search?q="

		path, err := cmd.Flags().GetBool("output_path")

		if err != nil {
			os.Exit(1)
		} 

		switch provider {
		case "youtube":
			queryString = "https://www.youtube.com/results?search_query="
			break
		case "amazon":
			queryString = "https://www.amazon.com/s/?keywords="
			break
		case "github":
			queryString = "https://github.com/search?utf8=✓&q="
			break
		case "wikipedia":
			queryString = "https://en.wikipedia.org/wiki/Special:Search?search="
			break
		default:
			fmt.Println("Opening in Google...")
			queryString = "https://www.google.com/search?q="
		}

		if(path) {
			fmt.Println(queryString+query)
			return
		}

		exe := exec.Command(code[0], append(code[1:], queryString+query)...)
		exe.Start()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// searchCmd.PersistentFlags().BoolP("youtube", "y", false, "Search keyword in Youtube")

	searchCmd.PersistentFlags().StringVarP(&provider, "provider", "p", "", "Pass in provider name (default google)")

	searchCmd.PersistentFlags().BoolP("output_path", "o", false, "View the query path")

}
