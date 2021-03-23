package resources

import (
	"embed"
	"html/template"
	"io/fs"
)

// Index template
var Index *template.Template

// Public filesystem for philote
var Public fs.FS

//go:embed *
var resources embed.FS

//go:embed index.go.html
var themeContent string

func init() {
	var err error

	Public, err = fs.Sub(resources, "public")
	if err != nil {
		panic(err)
	}

	Index, err = template.New("theme").Parse(themeContent)
	if err != nil {
		panic(err)
	}
}
