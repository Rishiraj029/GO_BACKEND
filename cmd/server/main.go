package main

import (
	"devconnect/internal/config"
	"devconnect/internal/database"
	"devconnect/internal/handler"
	"devconnect/internal/repository"
	"devconnect/internal/service"
	"fmt"
	"net/http"
)

func main() {

	cfg, err := config.Load()

	if err != nil {
		fmt.Println("Failed to load Configuration:", err)
		return
	}

	db, err := database.Connect(cfg)

	if err != nil {
		fmt.Println("Database connection failed:", err)
		return
	}

	defer db.Close()

	userRepository := repository.NewUserRepository(db)

	authService := service.NewAuthService(userRepository)

	authHandler := handler.NewAuthHandler(authService)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	mux.HandleFunc(
		"POST /api/v1/auth/register", authHandler.Register,
	)

	fmt.Println("DevConnect server running on :8080")

	error := http.ListenAndServe(":8080", mux)

	if error != nil {
		fmt.Println(error)
	}
}
