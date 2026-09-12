package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"devconnect/internal/apperror"
	"devconnect/internal/model"
	"devconnect/internal/service"
	"devconnect/internal/validator"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var request model.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	validationError := validator.ValidateRegisterRequest(request)

	if validationError != "" {
		http.Error(w, validationError, http.StatusBadRequest)
		return
	}

	user, err := h.authService.Register(request)

	if err != nil {
		if errors.Is(err, apperror.ErrEmailAlreadyExists) {
			http.Error(
				w,
				"email already exists",
				http.StatusConflict,
			)
			return
		}

		http.Error(
			w,
			"Failed to register user",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}
