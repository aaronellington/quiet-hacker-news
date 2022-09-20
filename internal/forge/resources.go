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
	{ // Prefer local
		localDir, err := fs.Sub(resources.LocalFS, dir)
		if err == nil {
			return localDir
		}
	}

	embedDir, err := fs.Sub(resources.EmbedFS, dir)
	if err != nil {
		panic(err)
	}

	return embedDir
}

// MustOpenFile is foobar
func (resources *Resources) MustOpenFile(fileName string) fs.File {
	{ // Prefer local
		localFile, err := resources.LocalFS.Open(fileName)
		if err == nil {
			return localFile
		}
	}

	embedFile, err := resources.EmbedFS.Open(fileName)
	if err != nil {
		panic(err)
	}

	return embedFile
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
