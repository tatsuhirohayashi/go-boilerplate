package usecase

import (
	"context"
	"go-boilerplate/internal/infrastructure/persistence/dto"
	apperrors "go-boilerplate/internal/pkg/errors"
	"go-boilerplate/internal/repository"
	"go-boilerplate/internal/usecase/input"
	"go-boilerplate/internal/usecase/output"
)

type TodoUseCase interface {
	ListTodo(ctx context.Context, input *input.ListTodoInput) (*output.TodoListOutput, error)
	GetTodo(ctx context.Context, input *input.GetTodoInput) (*output.TodoOutput, error)
	CreateTodo(ctx context.Context, input *input.CreateTodoInput) (*output.TodoOutput, error)
	UpdateTodo(ctx context.Context, input *input.UpdateTodoInput) (*output.TodoOutput, error)
	DeleteTodo(ctx context.Context, input *input.DeleteTodoInput) error
}

type todoUseCase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUseCase(todoRepo repository.TodoRepository) TodoUseCase {
	return &todoUseCase{todoRepo: todoRepo}
}

func (u *todoUseCase) ListTodo(ctx context.Context, input *input.ListTodoInput) (*output.TodoListOutput, error) {
	todos, err := u.todoRepo.FindAll(ctx, &dto.FindAllInput{
		UserID: input.UserID,
	})
	if err != nil {
		return nil, err
	}

	return output.NewTodoListOutput(todos), nil
}

func (u *todoUseCase) GetTodo(ctx context.Context, input *input.GetTodoInput) (*output.TodoOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, apperrors.NewValidationError("invalid input parameters", err)
	}
	inputDTO := &dto.FindByIDInput{
		ID:     input.ID,
		UserID: input.UserID,
	}
	todo, err := u.todoRepo.FindByID(ctx, inputDTO)
	if err != nil {
		return nil, err
	}

	return output.NewTodoOutput(todo), nil
}

func (u *todoUseCase) CreateTodo(ctx context.Context, input *input.CreateTodoInput) (*output.TodoOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, apperrors.NewValidationError("invalid input parameters", err)
	}
	inputDTO := &dto.CreateTodoInput{
		UserID:  input.UserID,
		Title:   input.Title,
		Content: input.Content,
	}
	todo, err := u.todoRepo.Create(ctx, inputDTO)
	if err != nil {
		return nil, err
	}

	return output.NewTodoOutput(todo), nil
}

func (u *todoUseCase) UpdateTodo(ctx context.Context, input *input.UpdateTodoInput) (*output.TodoOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, apperrors.NewValidationError("invalid input parameters", err)
	}
	inputFindDTO := &dto.FindByIDInput{
		ID:     input.ID,
		UserID: input.UserID,
	}
	existing, err := u.todoRepo.FindByID(ctx, inputFindDTO)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		return nil, apperrors.NewNotFoundError("todo not found", nil)
	}

	inputUpdateDTO := &dto.UpdateTodoInput{
		ID:      input.ID,
		UserID:  input.UserID,
		Title:   input.Title,
		Content: input.Content,
	}

	updated, err := u.todoRepo.Update(ctx, inputUpdateDTO)
	if err != nil {
		return nil, err
	}

	return output.NewTodoOutput(updated), nil
}

func (u *todoUseCase) DeleteTodo(ctx context.Context, input *input.DeleteTodoInput) error {
	if err := input.Validate(); err != nil {
		return apperrors.NewValidationError("invalid input parameters", err)
	}
	inputFindDTO := &dto.FindByIDInput{
		ID:     input.ID,
		UserID: input.UserID,
	}
	existing, err := u.todoRepo.FindByID(ctx, inputFindDTO)
	if err != nil {
		return err
	}
	if existing == nil {
		return apperrors.NewNotFoundError("todo not found", nil)
	}
	inputDeleteDTO := &dto.DeleteTodoInput{
		ID: input.ID,
	}
	return u.todoRepo.Delete(ctx, inputDeleteDTO)
}