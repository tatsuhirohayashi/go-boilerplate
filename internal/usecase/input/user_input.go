package input

import "errors"

type GetUserByEmailInput struct {
	Email string `json:"email" validate:"required,email"`
}

func (i *GetUserByEmailInput) Validate() error {
	if i.Email == "" {
		return errors.New("email is required")
	}
	return nil
}