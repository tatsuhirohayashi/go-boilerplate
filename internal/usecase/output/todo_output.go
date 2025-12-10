package output

import (
	"go-boilerplate/internal/infrastructure/persistence/dto"
	"time"

	"github.com/google/uuid"
)

type TodoOutput struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   *string   `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TodoListOutput struct {
	Todos []TodoOutput `json:"todos"`
	Total int64        `json:"total"`
}

func NewTodoOutput(todo *dto.TodoOutput) *TodoOutput {
	return &TodoOutput{
		ID:        todo.ID,
		Title:     todo.Title,
		Content:   todo.Content,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

func NewTodoListOutput(todos *dto.TodoListOutput) *TodoListOutput {
	outputs := make([]TodoOutput, len(todos.Todos))
	for i, todo := range todos.Todos {
		outputs[i] = *NewTodoOutput(&todo)
	}
	return &TodoListOutput{
		Todos: outputs,
		Total: todos.Total,
	}
}