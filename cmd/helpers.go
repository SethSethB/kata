package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

//InitGit Initialises git, creates .gitignore and adds initial commit
func InitGit(targetDir string, ignores []string) {
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
