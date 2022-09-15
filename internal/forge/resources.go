package forge

import (
	"embed"
	"html/template"
	"io"
	"io/fs"
)

// Resources is foobar
type Resources struct {
	EmbedFS embed.FS
	LocalFS fs.FS
}

// MustOpenDirectory is foobar
func (resources *Resources) MustOpenDirectory(dir string) fs.FS {
	// { // Dev Env
	// 	actualFilesOnDisk := os.DirFS("resources/" + dir)
	// 	_, err := actualFilesOnDisk.Open(".")
	// 	if err == nil {
	// 		return actualFilesOnDisk
	// 	}
	// }

	dest, err := fs.Sub(resources.EmbedFS, dir)
	if err != nil {
		panic(err)
	}

	return dest
}

// MustOpenFile is foobar
func (resources *Resources) MustOpenFile(fileName string) fs.File {
	// { // Dev Env
	// 	actualFilesOnDisk := os.DirFS("resources/" + dir)
	// 	_, err := actualFilesOnDisk.Open(".")
	// 	if err == nil {
	// 		return actualFilesOnDisk
	// 	}
	// }

	dest, err := resources.EmbedFS.Open(fileName)
	if err != nil {
		panic(err)
	}

	return dest
}

// MustOpenFileContents is foobar
func (resources *Resources) MustOpenFileContents(fileName string) string {
	file := resources.MustOpenFile(fileName)
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return string(fileBytes)
}

// MustParseHTMLTemplate is foobar
func (resources *Resources) MustParseHTMLTemplate(fileName string) *template.Template {
	file := resources.MustOpenFile(fileName)
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	t, err := template.New("theme").Parse(string(fileBytes))
	if err != nil {
		panic(err)
	}

	return t
}
