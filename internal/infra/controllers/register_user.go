package controllers

import (
	"capsuler/internal/domain/user/services"
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"go.uber.org/fx"
)

type RegisterUser struct {
	RegisterService *services.RegisterUser
	LoginService    *services.LoginUser
	Template        *template.Template
}

type RegisterRequestJson struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewRegisterController(lc fx.Lifecycle, registerService *services.RegisterUser, loginService *services.LoginUser, template *template.Template) *RegisterUser {
	return &RegisterUser{
		RegisterService: registerService,
		LoginService:    loginService,
		Template:        template,
	}
}

func (c *RegisterUser) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		c.Template.ExecuteTemplate(w, "register.html", nil)
		return
	}

	contentType := r.Header.Get("Content-Type")

	if contentType == "application/json" {
		defer r.Body.Close()
		var reqBody RegisterRequestJson
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			json.NewEncoder(w).Encode(map[string]any{"message": err.Error()})
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		if err := c.RegisterService.Register(reqBody.Username, reqBody.Email, reqBody.Password); err != nil {
			json.NewEncoder(w).Encode(map[string]any{"message": err.Error()})
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(map[string]any{"message": "user created"})
		w.WriteHeader(http.StatusCreated)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if err := c.RegisterService.Register(username, email, password); err != nil {
		c.Template.ExecuteTemplate(w, "register.html", map[string]any{"message": err.Error()})
		return
	}

	token, err := c.LoginService.Login(email, password)
	if err != nil {
		c.Template.ExecuteTemplate(w, "register.html", map[string]any{"message": err.Error()})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   string(token),
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 1),
	})
	http.Redirect(w, r, "/capsules/dashboard", http.StatusSeeOther)
}
