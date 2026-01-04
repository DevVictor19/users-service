package repositories

import (
	"context"
	"fmt"
	"usersservice/internal/users"
	"usersservice/internal/users/models"

	"github.com/jmoiron/sqlx"
)

type usersRepo struct {
	db *sqlx.DB
}

func NewMySQLUsersRepository(db *sqlx.DB) users.Repository {
	return &usersRepo{
		db: db,
	}
}

func (r *usersRepo) Create(ctx context.Context, user *models.User) (*models.User, error) {
	if _, err := r.db.ExecContext(ctx, userInsertQR,
		user.Identifier,
		user.ExternalID,
		user.FirstName,
		user.LastName,
		user.UserName,
		user.Email,
		user.EmailVerified,
		user.CreatedAt,
		user.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("usersRepo.Create: failed to create user: %w", err)
	}

	var createdUser models.User
	if err := r.db.GetContext(ctx, &createdUser, userSelectByIdentifierQR, user.Identifier); err != nil {
		return nil, fmt.Errorf("usersRepo.Create: failed to get created user: %w", err)
	}

	return &createdUser, nil
}
