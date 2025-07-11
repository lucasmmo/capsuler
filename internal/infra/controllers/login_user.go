package controllers

import (
	"capsuler/internal/domain/user/services"
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"go.uber.org/fx"
)

type LoginRequestJson struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUser struct {
	Service  *services.LoginUser
	Template *template.Template
}

func NewLoginController(lc fx.Lifecycle, service *services.LoginUser, template *template.Template) *LoginUser {
	return &LoginUser{
		Service:  service,
		Template: template,
	}
}

func (c *LoginUser) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.Template.ExecuteTemplate(w, "login.html", nil)
		return
	}
	contentType := r.Header.Get("Content-Type")

	if contentType == "application/json" {
		defer r.Body.Close()
		var reqBody LoginRequestJson
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			json.NewEncoder(w).Encode(map[string]any{"message": err.Error()})
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		token, err := c.Service.Login(reqBody.Email, reqBody.Password)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]any{"message": err.Error()})
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(map[string]any{"token": token})
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	token, err := c.Service.Login(email, password)
	if err != nil {
		c.Template.ExecuteTemplate(w, "login.html", map[string]any{"message": err.Error()})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   string(token),
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 1),
	})
	http.Redirect(w, r, "/capsules/dashboard", http.StatusSeeOther)
	return
}
