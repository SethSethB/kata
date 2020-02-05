package cmd

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"

	. "github.com/franela/goblin"
	"github.com/spf13/afero"
)

func TestFileHelpersSuite(t *testing.T) {

	g := Goblin(t)

	g.Describe("File Helper functions", func() {

		g.BeforeEach(func() {
			appFs = afero.NewMemMapFs()
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

		g.Describe("createKataFile", func() {
			g.It("creates new file with content", func() {
				c := []byte("test content")
				n := "testname"
				d := "testDirectory"
				createKataFile(c, n, d)

				expectedFile := path.Join(d, n)
				f, _ := appFs.Open(expectedFile)

				result, err := ioutil.ReadAll(f)
				g.Assert(err == nil).IsTrue()
				g.Assert(result).Equal(c)

			})

			g.It("Returns filesystem error", func() {

				appFs = &mockFs{}

				c := []byte("test content")
				n := "testname"
				d := "testDirectory"
				err := createKataFile(c, n, d)

				g.Assert(err == nil).IsFalse()
				g.Assert(err.Error()).Equal(errMockFs.Error())

			})

		})

		g.Describe("createContents", func() {

			g.It("getsContents from file and replaces them", func() {
				gopath := os.Getenv("GOPATH")
				n := "testtemplate"
				d := path.Join(gopath, "/src/github.com/sethsethb/kata/templates/")
				appFs.MkdirAll(d, os.ModePerm)
				f, _ := appFs.Create(path.Join(d, n))
				c := "test content with kataName"
				f.Write([]byte(c))

				result := createContents("findTheGoose", "testtemplate")
				g.Assert(string(result)).Equal("test content with findTheGoose")

			})
		})
	})
}

type mockFs struct{}
type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	Stat() (os.FileInfo, error)
}

var errMockFs = errors.New("File System Error")

func (mockFs) Create(name string) (afero.File, error)                      { return nil, errMockFs }
func (mockFs) Mkdir(name string, perm os.FileMode) error                   { return errMockFs }
func (mockFs) MkdirAll(path string, perm os.FileMode) error                { return errMockFs }
func (mockFs) Open(name string) (afero.File, error)                        { return nil, errMockFs }
func (mockFs) Remove(name string) error                                    { return errMockFs }
func (mockFs) RemoveAll(path string) error                                 { return errMockFs }
func (mockFs) Rename(oldname, newname string) error                        { return errMockFs }
func (mockFs) Stat(name string) (os.FileInfo, error)                       { return nil, errMockFs }
func (mockFs) Name() string                                                { return "" }
func (mockFs) Chmod(name string, mode os.FileMode) error                   { return errMockFs }
func (mockFs) Chtimes(name string, atime time.Time, mtime time.Time) error { return errMockFs }
func (mockFs) OpenFile(name string, flag int, perm os.FileMode) (afero.File, error) {
	return nil, errMockFs
}
