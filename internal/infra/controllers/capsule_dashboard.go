package controllers

import (
	"html/template"
	"net/http"

	"go.uber.org/fx"
)

type CapsuleDashboard struct {
	Template *template.Template
}

func NewCapsuleDashboardController(lc fx.Lifecycle, template *template.Template) *CapsuleDashboard {
	return &CapsuleDashboard{
		Template: template,
	}
}

func (c *CapsuleDashboard) Dashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.Template.ExecuteTemplate(w, "capsule_dashboard.html", nil)
		return
	}
}
