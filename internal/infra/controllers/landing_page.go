package controllers

import (
	"html/template"
	"net/http"

	"go.uber.org/fx"
)

type LandingPage struct {
	Template *template.Template
}

func NewLandingPageController(
	lc fx.Lifecycle,
	template *template.Template,
) *LandingPage {
	return &LandingPage{
		Template: template,
	}
}

func (c *LandingPage) LandingPage(w http.ResponseWriter, r *http.Request) {
	c.Template.ExecuteTemplate(w, "index.html", nil)
}
