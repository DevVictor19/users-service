package usecase

import (
	"context"
	"fmt"
	"usersservice/internal/users"
	"usersservice/internal/users/dtos"
	"usersservice/internal/users/models"
	"usersservice/pkg/env"
	"usersservice/pkg/httpErrors"
	"usersservice/pkg/logger"
	"usersservice/pkg/utils"
)

type usersUseCase struct {
	cfg       *env.Config
	usersRepo users.Repository
	logger    logger.Logger
}

func NewUsersUseCase(cfg *env.Config, usersRepo users.Repository, logger logger.Logger) users.UseCase {
	return &usersUseCase{
		cfg:       cfg,
		usersRepo: usersRepo,
		logger:    logger,
	}
}

func (u *usersUseCase) SignUp(ctx context.Context, input *dtos.SignUpInputDTO) (*dtos.UserDTO, error) {
	if err := utils.ValidateStruct(ctx, input); err != nil {
		return nil, httpErrors.NewBadRequestError(fmt.Errorf("usersUseCase.SignUp: failed to validate input: %w", err))
	}

	user := models.NewUser("test_id", input.FirstName, input.LastName, input.Email)

	user, err := u.usersRepo.Create(ctx, user)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(fmt.Errorf("usersUseCase.SignUp: failed to create user: %w", err))
	}

	output := &dtos.UserDTO{
		Identifier: user.Identifier,
		ExternalID: user.ExternalID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		UserName:   user.UserName,
		Email:      user.Email,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}

	return output, nil
}
