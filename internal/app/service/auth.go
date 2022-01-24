package service

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/ziyadovea/todo-app/internal/app/repository"
	"github.com/ziyadovea/todo-app/models"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

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

func (as *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := as.repo.GetUser(username, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: user.ID,
	})
	return token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
}

func (as *AuthService) ParseToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("incorrect type of token claims")
	}

	return claims.UserID, nil
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}
