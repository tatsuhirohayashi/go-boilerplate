package dto

import (
	"go-boilerplate/internal/domain"
	"time"

	"github.com/google/uuid"
)

type FindAllInput struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

type FindByIDInput struct {
	ID     uuid.UUID `json:"id" validate:"required"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

type CreateTodoInput struct {
	UserID  uuid.UUID `json:"user_id" validate:"required"`
	Title   string    `json:"title" validate:"required,min=1,max=100"`
	Content *string   `json:"content" validate:"omitempty,max=1000"`
}

type UpdateTodoInput struct {
	ID      uuid.UUID `json:"id" validate:"required"`
	UserID  uuid.UUID `json:"user_id" validate:"required"`
	Title   string    `json:"title" validate:"required,min=1,max=100"`
	Content *string   `json:"content" validate:"omitempty,max=1000"`
}

type DeleteTodoInput struct {
	ID     uuid.UUID `json:"id" validate:"required"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

type TodoOutput struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	Title     string    `json:"title"`
	Content   *string   `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TodoListOutput struct {
	Todos []TodoOutput `json:"todos"`
	Total int64        `json:"total"`
}

func ConvertTodoOutput(todo *domain.Todo) *TodoOutput {
	return &TodoOutput{
		ID:        todo.ID,
		UserID:    todo.UserID,
		Title:     todo.Title,
		Content:   todo.Content,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

func ConvertTodoListOutput(todos []*domain.Todo, total int64) *TodoListOutput {
	outputs := make([]TodoOutput, len(todos))
	for i, todo := range todos {
		outputs[i] = *ConvertTodoOutput(todo)
	}
	return &TodoListOutput{
		Todos: outputs,
		Total: total,
	}
}