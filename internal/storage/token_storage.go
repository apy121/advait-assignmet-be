package storage

import (
	"awesomeProject/internal/models"
	"sync"
)

var TokenStore = NewTokenStorage()

type TokenStorage struct {
	tokens map[string]models.Token
	mu     sync.Mutex
}

func NewTokenStorage() *TokenStorage {
	return &TokenStorage{tokens: make(map[string]models.Token)}
}

// RevokeToken removes the token from the storage
func (s *TokenStorage) RevokeToken(token string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.tokens, token)
	return nil
}

// StoreToken adds a new token to the storage
func (s *TokenStorage) StoreToken(token models.Token) {
	s.mu.Lock()
	s.tokens[token.Value] = token
	s.mu.Unlock()
}

// CheckToken verifies if a token exists and hasn't been revoked
func (s *TokenStorage) CheckToken(token string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.tokens[token]
	return exists
}
