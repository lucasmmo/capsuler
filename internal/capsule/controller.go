package capsule

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type controller struct {
	creator      *Creator
	opener       *Opener
	messageAdder *MessageAdder
}

func NewController(creator *Creator, opener *Opener, messageAdder *MessageAdder) *controller {
	return &controller{
		creator:      creator,
		opener:       opener,
		messageAdder: messageAdder,
	}
}

func (c *controller) CreateCapsule(w http.ResponseWriter, r *http.Request) {
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

	id, err := c.creator.Create(name, description, dateToOpen)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/capsules/%s", id))
	w.WriteHeader(http.StatusCreated)
}

func (c *controller) OpenCapsule(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := r.PathValue("id")

	err := c.opener.Open(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *controller) AddMessageToCapsule(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := r.PathValue("id")

	var mapData map[string]any

	if err := json.NewDecoder(r.Body).Decode(&mapData); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	message := mapData["message"].(string)

	messages, err := c.messageAdder.Add(id, message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(map[string]any{
		"messages": messages,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
