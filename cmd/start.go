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
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bhatipradeep/go-kv-shell/gokvshell"
	"github.com/spf13/cobra"
)

func attendCalls() {
	reader := bufio.NewReader(os.Stdin)
	transactionStack := gokvshell.NewTransactionStack(gokvshell.TRANSACTION_STACK_LIMIT)
	fmt.Println("START")
	var passErr error
	for {
		passErr = nil
		text, _ := reader.ReadString('\n')
		input := strings.Fields(text)
		switch input[0] {
		case "BEGIN":
			transaction := gokvshell.NewTransaction()
			passErr = (*transactionStack).PushTransaction(transaction)
			if passErr == nil {
				fmt.Println("--> Entered new transaction")
			}

		case "GET":
			if len(input) != 2 {
				passErr = gokvshell.InvalidGetArgumentsError{}

			} else {
				val, err := transactionStack.Get(input[1])
				if val != "" {
					fmt.Println(val)
				}
				passErr = err
			}

		case "SET":
			if len(input) < 3 {
				passErr = gokvshell.InvalidSetArgumentsError{}

			} else {
				transactionStack.Set(input[1], strings.Join(input[2:], ","))
			}

		case "DELETE":
			transactionStack.Delete(input[1])

		case "COMMIT":
			passErr = transactionStack.Commit()

		case "ROLLBACK":
			passErr = transactionStack.Rollback()

		case "END":
			fmt.Println("--> Ending out of top transaction")
			passErr = transactionStack.PopTransaction()

		case "EXIT":
			fmt.Println("--> Exiting go-kv-shell")
		default:
			fmt.Println("--> Enter valid operation")
		}

		if passErr != nil {
			fmt.Println("--> Error: " + passErr.Error())
		}
	}
}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start key value strore shell",
	Long:  "This command will start the transactional key value store shell",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("The key value store is initializing. You can start typing input once `START` prints.")
		attendCalls()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
