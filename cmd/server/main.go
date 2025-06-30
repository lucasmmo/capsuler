package main

import (
	"fmt"
	"net/http"
)

func main() {
	deps := InitDependencies()

	fs := http.FileServer(http.Dir("web/app/dist"))

	mux := http.NewServeMux()
	mux.Handle("/", fs)
	mux.Handle("/capsules", deps.capsuleRoutes)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}
