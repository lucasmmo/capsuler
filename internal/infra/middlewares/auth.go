package middlewares

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType == "application/json" {
			tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer")
			if err := Validate(tokenString); err != nil {
				json.NewEncoder(w).Encode(map[string]any{"message": err.Error()})
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			next.ServeHTTP(w, r)
			return
		}

		token, err := r.Cookie("token")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		tokenString := strings.TrimPrefix(token.Value, "Bearer")
		if err := Validate(tokenString); err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func Validate(tokenString string) error {
	if tokenString == "" {
		return errors.New("Invalid token value")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return errors.New("Unauthorized: Invalid token")
	}
	return nil
}
