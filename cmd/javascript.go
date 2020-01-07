// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// javascriptCmd represents the javascript command
var javascriptCmd = &cobra.Command{
	Use:   "javascript",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		createFile(name + ".js")
	},
}

func createFile(name string) {

	bs := createContents(name)

	fmt.Println(string(bs))
	file, err := os.Create(name)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error creating kata file", name)
	}

	file.Write(bs)
}

func createContents(name string) []byte {
	bs, err := ioutil.ReadFile("/home/sbell5/go/src/kata/templates/mainFunction.js")

	if err != nil {
		fmt.Println("error reading template", err)
	}
	return bs
}

func init() {
	RootCmd.AddCommand(javascriptCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// javascriptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// javascriptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
