package capsule

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Controller interface {
	CreateController(w http.ResponseWriter, r *http.Request)
	OpenController(w http.ResponseWriter, r *http.Request)
	AddMessageController(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	capsuleService Service
}

func NewController(capsuleService Service) *controller {
	return &controller{
		capsuleService: capsuleService,
	}
}

func (c *controller) CreateController(w http.ResponseWriter, r *http.Request) {
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

	id, err := c.capsuleService.CreateCapsule(name, description, dateToOpen)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/capsules/%s", id))
	w.WriteHeader(http.StatusCreated)
}

func (c *controller) OpenController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := r.PathValue("id")

	if err := c.capsuleService.OpenCapsule(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *controller) AddMessageController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := r.PathValue("id")

	var mapData map[string]any

	if err := json.NewDecoder(r.Body).Decode(&mapData); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	message := mapData["message"].(string)

	capsule, err := c.capsuleService.AddMessageToCapsule(id, message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/capsules/%s", capsule.Id))
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(map[string]any{
		"messages": capsule.Messages,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
}
