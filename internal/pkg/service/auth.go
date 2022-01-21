package service

import (
	"github.com/ziyadovea/todo-app/internal/pkg/repository"
	"github.com/ziyadovea/todo-app/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (as *AuthService) CreateUser(user *models.User) (int, error) {
	hash, err := generatePasswordHash(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hash
	return as.repo.CreateUser(user)
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}
