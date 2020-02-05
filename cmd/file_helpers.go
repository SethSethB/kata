package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func createKataFile(contents []byte, fileName, directory string) error {
	file, err := appFs.Create(path.Join("./", directory, fileName))

	if err != nil {
		fmt.Println("Error creating kata file", fileName, err)
		return err

	} else {
		_, err = file.Write(contents)
		file.Close()
		return err
	}

}

func createContents(name string, t string) []byte {
	gopath := os.Getenv("GOPATH")

	f, _ := appFs.Open(path.Join(gopath, "/src/github.com/sethsethb/kata/templates/", t))

	templateContent, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Println("error reading template:", err)
	}

	return replacePlaceholders(templateContent, name)
}

func replacePlaceholders(bs []byte, n string) []byte {
	contents := strings.ReplaceAll(string(bs), "kataName", n)
	contents = strings.ReplaceAll(contents, "KataName", convertLowerCamelCaseToUpper(n))
	contents = strings.ReplaceAll(contents, "kataname", strings.ToLower(n))
	return []byte(contents)
}
