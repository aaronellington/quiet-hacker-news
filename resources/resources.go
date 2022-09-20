package resources

import (
	"embed"
	"html/template"
	"io/fs"
	"os"

	"github.com/aaronellington/quiet-hacker-news/internal/forge"
)

// Public is foobar
var Public fs.FS

// Index template
var Index *template.Template

//go:embed *
var everything embed.FS

func init() {
	resources := forge.Resources{
		FileSystems: []fs.FS{
			os.DirFS("resources"),
			everything,
		},
	}

	Public = resources.MustOpenDirectory("public")
	Index = resources.MustParseHTMLTemplate("index.go.html")
}
