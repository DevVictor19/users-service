package users

import (
	"context"
	"usersservice/internal/users/dtos"
)

type UseCase interface {
	SignUp(ctx context.Context, input *dtos.SignUpInputDTO) (*dtos.UserDTO, error)
}
