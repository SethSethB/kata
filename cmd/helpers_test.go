package cmd

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"testing"

	. "github.com/franela/goblin"
)

func TestHelpersSuite(t *testing.T) {
	g := Goblin(t)

	g.Describe("ConvertToCamelCase", func() {
		g.It("Handles simple case", func() {
			g.Assert("simple").Equal(convertToCamelCase("simple"))
		})

		g.It("Converts words separated by space", func() {
			g.Assert("lessSimple").Equal(convertToCamelCase("less Simple"))
		})

		g.It("Converts case mixtures", func() {
			g.Assert("caseMixtureName").Equal(convertToCamelCase("CasE mIXture namE"))
		})

		g.It("Handles multiple spaces", func() {
			g.Assert("multipleSpaces").Equal(convertToCamelCase("MultiPle    spaceS"))
		})
	})

	g.Describe("convertLowerCamelCaseToUpper", func() {
		g.It("Updates first letter to upperCase", func() {
			g.Assert("UpperCamelCase").Equal(convertLowerCamelCaseToUpper("upperCamelCase"))
		})
	})

	g.Describe("initGit", func() {

		var kataPath string

		g.BeforeEach(func() {
			tempDir := path.Join(os.TempDir(), "testKataDir")

			os.RemoveAll(tempDir)
			kataPath = path.Join(tempDir, "testKata")
			os.MkdirAll(kataPath, 0777)
		})

		g.It("initialises git - command 'git status' will not throw an error", func() {
			initGit(kataPath, []string{})

			initCmd := exec.Command("git", "status")
			initCmd.Dir = kataPath
			msg, _ := initCmd.StderrPipe()
			initCmd.Start()
			bs, _ := ioutil.ReadAll(msg)

			g.Assert(len(bs)).Equal(0)
		})

		g.It("commits all existing files", func() {

			f, _ := os.Create(path.Join(kataPath, "testfile.txt"))
			f.Close()

			initGit(kataPath, []string{})

			initCmd := exec.Command("git", "status")
			initCmd.Dir = kataPath
			msg, _ := initCmd.StdoutPipe()
			initCmd.Start()
			bs, _ := ioutil.ReadAll(msg)

			g.Assert(string(bs)).Equal("On branch master\nnothing to commit, working tree clean\n")
		})

		g.It("creates a .gitignore file with files added", func() {

			initGit(kataPath, []string{"testfile", "node_modules"})

			bs, err := ioutil.ReadFile(path.Join(kataPath, ".gitignore"))

			g.Assert(err == nil).IsTrue()
			g.Assert(string(bs)).Equal("testfile\nnode_modules")

		})
	})

	g.Describe("replacePlaceholders", func() {
		g.It("does not change content with no placeholders", func() {
			name := "funnyFacesKata"
			content := []byte("nothing to replace")
			g.Assert(string(content)).Equal(string(replacePlaceholders(content, name)))
		})

		g.It("updates kataName places holders with name", func() {
			name := "funnyFacesKata"
			content := []byte("this kataName and /nkataName will change")
			expected := []byte("this " + name + " and /n" + name + " will change")
			g.Assert(string(expected)).Equal(string(replacePlaceholders(content, name)))
		})

		g.It("updates KataName places holders with Name", func() {
			name := "funnyFacesKata"
			content := []byte("this kataName and /nKataName will change")
			expected := []byte("this " + name + " and /nFunnyFacesKata will change")
			g.Assert(string(expected)).Equal(string(replacePlaceholders(content, name)))
		})

		g.It("updates kataname places holders with Name", func() {
			name := "funnyFacesKata"
			content := []byte("this kataName and /nKataName will change. And also kataname")
			expected := []byte("this " + name + " and /nFunnyFacesKata will change. And also funnyfaceskata")
			g.Assert(string(expected)).Equal(string(replacePlaceholders(content, name)))
		})

	})

}
