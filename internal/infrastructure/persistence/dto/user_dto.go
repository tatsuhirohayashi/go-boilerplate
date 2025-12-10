package dto

import (
	"go-boilerplate/internal/domain"
	"time"

	"github.com/google/uuid"
)


type FindUserByEmailInput struct {
	Email string `json:"email" validate:"required,email"`
}

type CreateUserInput struct {
	Name     string `json:"name" validate:"required,min=1,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type UserOutput struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ConvertUserOutput(user *domain.User) *UserOutput {
	return &UserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}