package persistence_gorm

import (
	"context"
	"go-boilerplate/internal/domain"
	"go-boilerplate/internal/infrastructure/persistence/dto"
	"go-boilerplate/internal/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByEmail(ctx context.Context, input *dto.FindUserByEmailInput) (*dto.UserOutput, error) {
	var user domain.User
	if err := r.db.First(&user, "email = ?", input.Email).Error; err != nil {
		return nil, HandleDBError(err, "user")
	}
	return dto.ConvertUserOutput(&user), nil
}

func (r *userRepository) Create(ctx context.Context, input *dto.CreateUserInput) (*dto.UserOutput, error) {
	user := domain.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	if err := r.db.Create(&user).Error; err != nil {
		return nil, HandleDBError(err, "user")
	}
	return dto.ConvertUserOutput(&user), nil
}