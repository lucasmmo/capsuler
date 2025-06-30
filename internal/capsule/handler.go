package capsule

import "net/http"

func NewHandler(capsuleController *controller) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /", capsuleController.CreateController)
	mux.HandleFunc("POST /{id}/open", capsuleController.OpenController)
	mux.HandleFunc("POST /{id}/messages", capsuleController.AddMessageController)

	return mux
}
