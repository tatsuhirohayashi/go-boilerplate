package repository

import (
	"context"
	"go-boilerplate/internal/infrastructure/persistence/dto"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, input *dto.FindUserByEmailInput) (*dto.UserOutput, error)
	Create(ctx context.Context, input *dto.CreateUserInput) (*dto.UserOutput, error)
}