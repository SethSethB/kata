package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

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
