package models

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            *int64    `db:"id"`
	Identifier    string    `db:"identifier"`
	ExternalID    string    `db:"external_id"`
	FirstName     string    `db:"first_name"`
	LastName      string    `db:"last_name"`
	UserName      string    `db:"username"`
	Email         string    `db:"email"`
	EmailVerified bool      `db:"email_verified"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

func NewUser(externalID, firstName, lastName, email string) *User {
	now := time.Now()
	uuid, err := uuid.NewV7()
	if err != nil {
		log.Fatalf("User.NewUser: failed to generate uuid: %v", err)
	}
	userName := fmt.Sprintf("%s.%s", firstName, lastName)

	return &User{
		Identifier: uuid.String(),
		ExternalID: externalID,
		FirstName:  firstName,
		LastName:   lastName,
		UserName:   userName,
		Email:      email,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}
