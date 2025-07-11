package templates

import (
	"embed"
	"html/template"
	"log"

	"go.uber.org/fx"
)

//go:embed **/*.html
var FS embed.FS

func NewTemplate(lc fx.Lifecycle) *template.Template {
	tmpl, err := template.ParseFS(FS, "**/*.html")
	if err != nil {
		log.Fatal(err)
	}
	return tmpl
}
