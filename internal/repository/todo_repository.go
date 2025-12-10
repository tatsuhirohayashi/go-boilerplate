package repository

import (
	"context"
	"go-boilerplate/internal/infrastructure/persistence/dto"
)

type TodoRepository interface {
	FindAll(ctx context.Context, input *dto.FindAllInput) (*dto.TodoListOutput, error)
	FindByID(ctx context.Context, input *dto.FindByIDInput) (*dto.TodoOutput, error)
	Create(ctx context.Context, input *dto.CreateTodoInput) (*dto.TodoOutput, error)
	Update(ctx context.Context, input *dto.UpdateTodoInput) (*dto.TodoOutput, error)
	Delete(ctx context.Context, input *dto.DeleteTodoInput) error
}