package service

import (
	"golang.org/x/crypto/bcrypt"

	"devconnect/internal/model"
	"devconnect/internal/repository"
)

type AuthService struct {
	userRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) Register(request model.RegisterRequest) (model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		// a go password is a string and to convert the string into byte what the bycrypt expects thats why we have created []byte(request.Password)
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: string(hashedPassword),
	}

	id, err := s.userRepository.CreateUser(user)

	if err != nil {
		return model.User{}, err
	}

	user.ID = id

	return user, nil
}
