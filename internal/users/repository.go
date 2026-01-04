package users

import (
	"context"
	"usersservice/internal/users/models"
)

type Repository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
}
