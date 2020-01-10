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
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"

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

		var name string

		if len(args) == 0 {
			name = promptName()
		} else {
			name = args[0]
		}

		os.Mkdir(name, os.ModePerm)
		targetDir := path.Join("./", name)

		setupNode(targetDir)

		fmt.Println("Writing kata files...")
		createFile(name, name, "mainFunction.js")
		createFile(name, name+".spec", "testSuite.js")

		if git == true {
			initGit(targetDir)
		}

		finalMessage := fmt.Sprintf("Complete! \nRun the command \"cd %s && npm test\" to run test suite", name)
		fmt.Println(finalMessage)
	},
}

func promptName() string {
	fmt.Print("Enter kata name: ")

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n := s.Text()

	if len(n) == 0 {
		return promptName()
	}

	return n
}

func setupNode(targetDir string) {
	initCmd := exec.Command("npm", "init", "-y")
	initCmd.Dir = targetDir

	fmt.Println("Initialising npm...")
	err := initCmd.Run()
	if err != nil {
		fmt.Println("Error inialising npm: ", err)
	}

	fmt.Println("Installing node dependencies...")
	installCmd := exec.Command("npm", "install", "-D", "mocha", "chai")
	installCmd.Dir = targetDir
	err = installCmd.Run()
	if err != nil {
		fmt.Println("Error installing node modules: ", err)
	}

	bs, err := ioutil.ReadFile(path.Join(targetDir, "/package.json"))
	if err != nil {
		fmt.Println("error reading package.json", err)
	}
	contents := strings.ReplaceAll(string(bs), "echo \\\"Error: no test specified\\\" && exit 1", "mocha *.spec.js")
	err = ioutil.WriteFile(path.Join(targetDir, "/package.json"), []byte(contents), os.ModePerm)
	if err != nil {
		fmt.Println("error updating package.json", err)
	}
}

func createFile(kataName string, fileName string, fileTemplate string) {

	bs := createContents(kataName, fileTemplate)

	file, err := os.Create(path.Join("./", kataName, fileName+".js"))

	if err != nil {
		fmt.Println("Error creating kata file", kataName, err)
	}

	file.Write(bs)
}

func createContents(name string, template string) []byte {
	gopath := os.Getenv("GOPATH")
	bs, err := ioutil.ReadFile(path.Join(gopath, "/src/github.com/sethsethb/kata-gen/templates/", template))

	if err != nil {
		fmt.Println("error reading template: ", err)
	}

	contents := strings.ReplaceAll(string(bs), "KATANAME", name)
	return []byte(contents)
}

func initGit(targetDir string) {
	initCmd := exec.Command("git", "init")
	initCmd.Dir = targetDir

	fmt.Println("Initialising git with initial commit...")
	err := initCmd.Run()
	if err != nil {
		fmt.Println("error initalising git: ", err)
	}

	file, _ := os.Create(path.Join("./", targetDir, ".gitignore"))

	bs := []byte("node_modules/")
	file.Write(bs)

	addCmd := exec.Command("git", "add", ".")
	addCmd.Dir = targetDir
	addCmd.Run()

	commitCmd := exec.Command("git", "commit", "-m", "Initial commit")
	commitCmd.Dir = targetDir
	commitCmd.Run()

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
