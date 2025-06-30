package main

import (
	"fmt"
	"net/http"
)

func main() {
	deps := InitDependencies()

	mux := http.NewServeMux()
	mux.Handle("/capsules", deps.capsuleHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}
