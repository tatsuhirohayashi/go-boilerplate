package input

import "errors"

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

func (i *LoginInput) Validate() error {
	if i.Email == "" {
		return errors.New("email is required")
	}
	if i.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

type RegisterUserInput struct {
	Name     string `json:"name" validate:"required,min=1,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

func (i *RegisterUserInput) Validate() error {
	if i.Name == "" {
		return errors.New("name is required")
	}
	if i.Email == "" {
		return errors.New("email is required")
	}
	if i.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

type CheckAuthenticationInput struct {
	Email    string `json:"email" validate:"required,email"`
}

func (i *CheckAuthenticationInput) Validate() error {
	if i.Email == "" {
		return errors.New("email is required")
	}
	return nil
}