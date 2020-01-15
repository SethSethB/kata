package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

func createKataFile(contents []byte, fileName, directory string) {

	file, err := os.Create(path.Join("./", directory, fileName))

	defer file.Close()

	if err != nil {
		fmt.Println("Error creating kata file", fileName, err)
	}

	file.Write(contents)
}

func createContents(n string, t string) []byte {
	gopath := os.Getenv("GOPATH")

	bs, err := ioutil.ReadFile(path.Join(gopath, "/src/github.com/sethsethb/kata-gen/templates/", t))

	if err != nil {
		fmt.Println("error reading template:", err)
	}

	contents := strings.ReplaceAll(string(bs), "kataName", n)
	contents = strings.ReplaceAll(contents, "KataName", convertLowerCamelCaseToUpper(n))
	contents = strings.ReplaceAll(contents, "kataname", strings.ToLower(n))

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

func convertLowerCamelCaseToUpper(s string) string {

	bs := []byte(s)

	firstChar := string(bs[0])

	return strings.Replace(s, firstChar, strings.ToUpper(firstChar), 1)
}

func createKataName(args []string) string {
	if len(args) == 0 {
		return promptName()
	}
	return args[0]
}

func promptName() string {
	fmt.Print("Enter kata name: ")

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n := s.Text()

	if len(n) == 0 {
		return promptName()
	}

	return convertToCamelCase(n)
}
