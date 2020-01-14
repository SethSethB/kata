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

var gradle bool

// javaCmd represents the java command
var javaCmd = &cobra.Command{
	Use:   "java",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		kataName := createKataName(args)

		targetDir := path.Join("./", kataName)

		if gradle == true {
			fmt.Println("Creating gradle project")
		} else {
			fmt.Println("Creating maven project...")
			createMaven(kataName)
		}

		if git == true {
			initGit(targetDir, []string{
				".classpath",
				".project",
				".settings/",
				"target/",
				".idea/",
			})
		}
		finalMessage := fmt.Sprintf("Complete! \nRun the command \"cd %s && mvn test\" to run test suite", kataName)
		defer fmt.Println(finalMessage)
	},
}

func createMaven(n string) {

	targetDir := path.Join("./", n)
	className := convertLowerCamelCaseToUpper(n)

	classDir := path.Join(targetDir, "/src/main/java/com/kata")
	testDir := path.Join(targetDir, "/src/test/java/com/kata")

	os.MkdirAll(classDir, os.ModePerm)
	os.MkdirAll(testDir, os.ModePerm)

	createKataFile(className, className, classDir, "/java/mainClass.java")
	createKataFile(className+"Test", className, testDir, "/java/testClass.java")
	createKataFile("pom", "", targetDir, "/java/pom.xml")
}

func init() {
	RootCmd.AddCommand(javaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// javaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	javaCmd.Flags().BoolVarP(&gradle, "gradle", "", false, "Initialises as gradle rather than maven")
}
