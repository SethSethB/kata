/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"path"

	"github.com/spf13/cobra"
)

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Creates the boilerplate files for a go kata",
	Long: `The go command does the following:
	
	Creates a new directory
	Creates a function & test files based on the name provided
	If no args are provided it will prompt for a  name.
	Names are converted to camelcase automatically. E.g. "some EXAMPLE name" will be named someExampleName
	If run with the git flag it will create a new repository and commit the inital files
	`,
	Run: func(cmd *cobra.Command, args []string) {
		name := createKataName(args)
		os.Mkdir(name, os.ModePerm)
		targetDir := path.Join("./", name)

		fmt.Println("Writing kata files...")
		mainContents := createContents(name, "/go/mainFunction.go")
		testContents := createContents(name, "/go/testSuiteGoblin.go")
		createKataFile(mainContents, name+".go", targetDir)
		createKataFile(testContents, name+"_test.go", targetDir)

		executeCmdInKataDir(command{
			dir:    targetDir,
			name:   "go",
			args:   []string{"mod", "init", "kata/" + name},
			msg:    "Initialising go mod...",
			errMsg: "Error initialising go module",
		})

		if git {
			initGit(targetDir, []string{"node_modules"})
		}

		finalMessage := fmt.Sprintf("Complete! \nRun the command \"cd %s && go test\" to run test suite", name)
		fmt.Println(finalMessage)
	},
}

func init() {
	RootCmd.AddCommand(goCmd)
}
