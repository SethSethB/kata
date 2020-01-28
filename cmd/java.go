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

var javaCmd = &cobra.Command{
	Use:   "java",
	Short: "Creates the boilerplate files for a java kata",
	Long: `The java command does the following:
	
	Creates a new directory
	Creates a function file & test file based on the name provided
	If no args are provided it will prompt for a function name.
	Names are converted to camelcase automatically. E.g. "some EXAMPLE name" will be named someExampleName
	If run with the git flag it will create a new repository and commit the inital files`,
	Run: func(cmd *cobra.Command, args []string) {
		kataName := createKataName(args)
		targetDir := path.Join("./", kataName)

		createMaven(kataName)

		if git {
			initGit(targetDir, []string{
				".classpath",
				".project",
				".settings/",
				"target/",
				".idea/",
			})
		}
		finalMessage := fmt.Sprintf("\nComplete! \nRun the command \"cd %s && mvn test\" to run test suite", kataName)
		fmt.Println(finalMessage)
	},
}

func createMaven(n string) {
	fmt.Println("Creating maven project...")

	targetDir := path.Join("./", n)
	classDir := path.Join(targetDir, "/src/main/java/com/kata")
	testDir := path.Join(targetDir, "/src/test/java/com/kata")

	os.MkdirAll(classDir, os.ModePerm)
	os.MkdirAll(testDir, os.ModePerm)

	mainContents := createContents(n, "/java/mainClass.java")
	testContents := createContents(n, "/java/testClass.java")
	pomContents := createContents("pom.xml", "/java/pom.xml")

	className := convertLowerCamelCaseToUpper(n)
	createKataFile(mainContents, className+".java", classDir)
	createKataFile(testContents, className+"Test.java", testDir)
	createKataFile(pomContents, "pom.xml", targetDir)
}

func init() {
	RootCmd.AddCommand(javaCmd)
}
