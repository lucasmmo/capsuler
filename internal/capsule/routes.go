package capsule

import "net/http"

func NewRouter(capsuleController *controller) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", capsuleController.CreateCapsule)
	mux.HandleFunc("POST /{id}/open", capsuleController.OpenCapsule)
	mux.HandleFunc("POST /{id}/messages", capsuleController.AddMessageToCapsule)

	return mux
}
