package storage

import (
	"awesomeProject/internal/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserStorage struct {
	users map[string]models.User
}

func NewUserStorage() *UserStorage {
	return &UserStorage{users: make(map[string]models.User)}
}

// CreateUser now uses ID
func (s *UserStorage) CreateUser(user models.User) error {
	if _, exists := s.users[user.Email]; exists {
		return errors.New("user already exists")
	}
	s.users[user.Email] = user
	return nil
}

// AuthenticateUser now checks against hashed passwords
func (s *UserStorage) AuthenticateUser(email, password string) (models.User, error) {
	user, exists := s.users[email]
	if !exists {
		return models.User{}, errors.New("user not found")
	}
	// Compare the provided password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, errors.New("invalid password")
	}
	return user, nil
}
