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

	"github.com/spf13/cobra"
)

// bookmarkCmd represents the bookmark command
var bookmarkCmd = &cobra.Command{
	Use:   "bookmark",
	Short: "Bookmark website links",
	Long:  `Stores and Saves website links`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bookmark called")
	},
}

func init() {
	rootCmd.AddCommand(bookmarkCmd)
}
