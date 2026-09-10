package main

import (
	"devconnect/internal/handler"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	mux.HandleFunc("POST /api/v1/auth/register", handler.Register)

	fmt.Println("DevConnect server running on :8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println(err)
	}
}
