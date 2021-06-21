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
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var yt string;

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
		if(yt == "yes") {
			queryString = "https://www.youtube.com/results?search_query="
		}
		exe := exec.Command(code[0], append(code[1:], queryString+query)...)
		exe.Start()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	searchCmd.PersistentFlags().StringVarP(&yt, "youtube", "y", "", "Pass in keyword to search in youtube")
}
