package validator

import (
	"strings"

	"devconnect/internal/model"
)

func ValidateRegisterRequest(request model.RegisterRequest) string {
	if strings.TrimSpace(request.Name) == "" {
		return "name is required"
	}

	if strings.TrimSpace(request.Email) == "" {
		return "email is required"
	}

	if strings.TrimSpace(request.Password) == "" {
		return "password is required"
	}

	if len(request.Password) < 8 {
		return "password must be at least 8 characters"
	}

	return ""
}

func ValidateLoginRequest(request model.LoginRequest) string {
	if strings.TrimSpace(request.Email) == "" {
		return "email is required"
	}

	if strings.TrimSpace(request.Password) == "" {
		return "password is required"
	}

	return ""
}
