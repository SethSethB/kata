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
	"path"
	"strings"

	"github.com/spf13/cobra"
)

var jest bool

var javascriptCmd = &cobra.Command{
	Use:   "javascript",
	Short: "Creates the boilerplate files for a javascript kata",
	Long: `The javascript command does the following:
	
	Creates a new directory
	Creates class & test files based on the name provided
	If no args are provided it will prompt for a name.
	Names are converted to camelcase automatically. E.g. "some EXAMPLE name" will be named someExampleName
	Inititates npm and installs mocha/chai as dev dependencies by default
	If run with the jest flag will initialise with jest rather than mocha/chai
	If run with the git flag it will create a new repository and commit the inital files`,
	Run: func(cmd *cobra.Command, args []string) {

		name := createKataName(args)

		os.Mkdir(name, os.ModePerm)
		targetDir := path.Join("./", name)

		executeCmdInKataDir(command{
			dir:    targetDir,
			name:   "npm",
			args:   []string{"init", "-y"},
			msg:    "Initialising npm...",
			errMsg: "Error inialising npm",
		})

		fmt.Println("Writing kata files...")
		mainContents := createContents(name, "/javascript/mainFunction.js")
		createKataFile(mainContents, name+".js", targetDir)

		var testScript string
		var dependencies []string
		var testContents []byte

		if jest {
			testScript = "jest --watchAll"
			dependencies = []string{"jest"}
			testContents = createContents(name, "/javascript/testSuiteJest.js")
		} else {
			testScript = "mocha *.spec.js"
			dependencies = []string{"mocha", "chai"}
			testContents = createContents(name, "/javascript/testSuite.js")
		}

		createKataFile(testContents, name+".spec.js", targetDir)
		installDependencies(targetDir, dependencies)
		configureTestScript(targetDir, testScript)

		if git {
			initGit(targetDir, []string{"node_modules"})
		}

		finalMessage := fmt.Sprintf("\nComplete! \nRun the command \"cd %s && npm test\" to run test suite", name)
		fmt.Println(finalMessage)
	},
}

func installDependencies(targetDir string, deps []string) {

	args := []string{"install", "-D"}
	args = append(args, deps...)

	executeCmdInKataDir(command{
		dir:    targetDir,
		name:   "npm",
		args:   args,
		msg:    "Installing Dependencies...",
		errMsg: "Error inialising npm",
	})
}

func configureTestScript(targetDir, testScript string) {
	bs, err := ioutil.ReadFile(path.Join(targetDir, "/package.json"))
	if err != nil {
		fmt.Println("error reading package.json", err)
	}
	contents := strings.ReplaceAll(string(bs), "echo \\\"Error: no test specified\\\" && exit 1", testScript)

	err = ioutil.WriteFile(path.Join(targetDir, "/package.json"), []byte(contents), os.ModePerm)
	if err != nil {
		fmt.Println("error updating package.json", err)
	}
}

func init() {
	RootCmd.AddCommand(javascriptCmd)
	javascriptCmd.Flags().BoolVarP(&jest, "jest", "j", false, "creates kata with jest as test runner")

}
