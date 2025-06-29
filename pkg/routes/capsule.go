package routes

import (
	"capsuler/pkg/services"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func CreateCapsule(capsuleRepository services.CapsuleRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var mapData map[string]any

		if err := json.NewDecoder(r.Body).Decode(&mapData); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		name := mapData["name"].(string)
		description := mapData["description"].(string)
		dateToOpenInt := mapData["date_to_open"].(float64)

		dateToOpen := time.Unix(int64(dateToOpenInt), 0)

		id, err := services.CreateCapsule(name, description, dateToOpen, capsuleRepository)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}

		w.Header().Set("Location", fmt.Sprintf("/capsules/%s", id))
		w.WriteHeader(http.StatusCreated)
	}
}

func OpenCapsule(capsuleRepository services.CapsuleRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		id := r.PathValue("id")

		if err := services.OpenCapsule(id, capsuleRepository); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func AddMessageToCapsule(capsuleRepository services.CapsuleRepository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		id := r.PathValue("id")

		var mapData map[string]any

		if err := json.NewDecoder(r.Body).Decode(&mapData); err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}

		message := mapData["message"].(string)

		capsule, err := services.AddMessageToCapsule(id, message, capsuleRepository)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}

		w.Header().Set("Location", fmt.Sprintf("/capsules/%s", capsule.GetId()))
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(map[string]any{
			"messages": capsule.GetMessages(),
		}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err.Error())
			return
		}
	}
}
