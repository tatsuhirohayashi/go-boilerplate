package input

import (
	"errors"

	"github.com/google/uuid"
)

type ListTodoInput struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

func (i *ListTodoInput) Validate() error {
	if i.UserID == uuid.Nil {
		return errors.New("user_id is required")
	}
	return nil
}

type GetTodoInput struct {
	ID     uuid.UUID `json:"id" validate:"required"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

func (i *GetTodoInput) Validate() error {
	if i.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if i.UserID == uuid.Nil {
		return errors.New("user_id is required")
	}
	return nil
}


type CreateTodoInput struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
	Title     string  `json:"title" validate:"required,min=1,max=100"`
	Content   *string `json:"content" validate:"omitempty,max=1000"`
}

type UpdateTodoInput struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
	Title     string    `json:"title" validate:"required,min=1,max=100"`
	Content   *string   `json:"content" validate:"omitempty,max=1000"`
}

type DeleteTodoInput struct {
	ID     uuid.UUID `json:"id" validate:"required"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

func (i *CreateTodoInput) Validate() error {
	if i.UserID == uuid.Nil {
		return errors.New("user_id is required")
	}
	if i.Title == "" {
		return errors.New("title is required")
	}
	if len(i.Title) > 100 {
		return errors.New("title must be less than 100 characters")
	}
	if i.Content != nil && len(*i.Content) > 1000 {
		return errors.New("content must be less than 1000 characters")
	}
	return nil
}

func (i *UpdateTodoInput) Validate() error {
	if i.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if i.UserID == uuid.Nil {
		return errors.New("user_id is required")
	}
	if i.Title == "" {
		return errors.New("title is required")
	}
	if len(i.Title) > 100 {
		return errors.New("title must be less than 100 characters")
	}
	if i.Content != nil && len(*i.Content) > 1000 {
		return errors.New("content must be less than 1000 characters")
	}
	return nil
}


func (i *DeleteTodoInput) Validate() error {
	if i.ID == uuid.Nil {
		return errors.New("id is required")
	}
	if i.UserID == uuid.Nil {
		return errors.New("user_id is required")
	}
	return nil
}