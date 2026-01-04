package dtos

import "time"

type UserDTO struct {
	Identifier    string    `json:"identifier"`
	ExternalID    string    `json:"external_id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	UserName      string    `json:"username"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"email_verified"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
