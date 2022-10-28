package forge

import (
	"html/template"
	"io"
	"io/fs"
)

type Resources struct {
	FileSystems []fs.FS
}

func (resources *Resources) MustOpenDirectory(dir string) fs.FS {
	for i, fileSystem := range resources.FileSystems {
		_, openTestErr := fileSystem.Open(dir)
		if openTestErr != nil {
			continue
		}

		directory, err := fs.Sub(fileSystem, dir)
		if err == nil {
			return directory
		}

		if i == (len(resources.FileSystems) - 1) {
			panic(err)
		}
	}

	panic("no fileSystems")
}

func (resources *Resources) MustOpenFile(fileName string) fs.File {
	for i, fileSystem := range resources.FileSystems {
		file, err := fileSystem.Open(fileName)
		if err == nil {
			return file
		}

		// If the last filesystem, panic
		if i == (len(resources.FileSystems) - 1) {
			panic(err)
		}
	}

	panic("no fileSystems")
}

func (resources *Resources) MustOpenFileContents(fileName string) string {
	file := resources.MustOpenFile(fileName)
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return string(fileBytes)
}

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
