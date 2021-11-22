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

	"github.com/spf13/cobra"
)

func attendCalls() {
	// To-Do : Add code to respond to kv store calls.
}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start key value strore shell",
	Long:  "This command will start the transactional key value store shell",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("The key value store is initializing. You can start typing input once `START` prints.")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
