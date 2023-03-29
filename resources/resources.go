package resources

import (
	"embed"
	"html/template"
	"io/fs"
	"os"

	"github.com/kyberbits/forge/forge"
)

//go:embed *
var everything embed.FS

type Resources struct {
	Public fs.FS
	Index  *template.Template
}

func NewResources() Resources {
	resources := forge.NewResources([]fs.FS{
		os.DirFS("resources"),
		everything,
	})

	return Resources{
		Public: resources.MustOpenDirectory("public"),
		Index:  resources.MustParseHTMLTemplate("index.go.html"),
	}
}
