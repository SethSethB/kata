package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func createKataFile(fileName, funcName, directory, template string) {

	extension := filepath.Ext(template)

	bs := createContents(funcName, template)

	file, err := os.Create(path.Join("./", directory, fileName+extension))

	defer file.Close()

	if err != nil {
		fmt.Println("Error creating kata file", fileName, err)
	}

	file.Write(bs)
}

func createContents(n string, t string) []byte {
	gopath := os.Getenv("GOPATH")

	bs, err := ioutil.ReadFile(path.Join(gopath, "/src/github.com/sethsethb/kata-gen/templates/", t))

	if err != nil {
		fmt.Println("error reading template:", err)
	}

	contents := strings.ReplaceAll(string(bs), "KATANAME", n)
	return []byte(contents)
}

func initGit(targetDir string, ignores []string) {
	initCmd := exec.Command("git", "init")
	initCmd.Dir = targetDir

	fmt.Println("Initialising git with initial commit...")
	err := initCmd.Run()
	if err != nil {
		fmt.Println("error initalising git: ", err)
	}

	gitIgnore, _ := os.Create(path.Join("./", targetDir, ".gitignore"))
	defer gitIgnore.Close()

	for i, ignore := range ignores {

		if i != 0 {
			ignore = "\n" + ignore
		}

		bs := []byte(ignore)
		gitIgnore.Write(bs)
	}

	addCmd := exec.Command("git", "add", ".")
	addCmd.Dir = targetDir
	addCmd.Run()

	commitCmd := exec.Command("git", "commit", "-m", "Initial commit")
	commitCmd.Dir = targetDir
	commitCmd.Run()

}

func convertToCamelCase(s string) string {
	isUpper := false

	var r string
	letters := strings.Split(s, "")

	for _, l := range letters {

		if l == " " {
			isUpper = true
			continue
		}

		if isUpper {
			r += strings.ToUpper(l)
			isUpper = false
		} else {
			r += strings.ToLower(l)
		}
	}
	return r
}

func convertToUpperCamelCase(s string) string {

	c := convertToCamelCase(s)
	bs := []byte(c)

	firstChar := string(bs[0])

	return strings.Replace(c, firstChar, strings.ToUpper(firstChar), 1)
}
