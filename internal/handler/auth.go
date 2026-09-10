package handler

import (
	"encoding/json"
	"net/http"

	"devconnect/internal/model"
	"devconnect/internal/validator"
)

func Register(w http.ResponseWriter, r *http.Request) {

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

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(request)

}
