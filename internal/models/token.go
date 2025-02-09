package models

// Token represents a JWT token in our system
type Token struct {
	Value  string `json:"value"`
	UserId string `json:"user_id"`
}
