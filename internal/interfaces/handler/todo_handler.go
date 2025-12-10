package handler

import (
	"encoding/json"
	"go-boilerplate/internal/pkg/constants"
	apperrors "go-boilerplate/internal/pkg/errors"
	"go-boilerplate/internal/usecase"
	"go-boilerplate/internal/usecase/input"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type TodoHandler interface {
	RegisterTodoHandlers(r *mux.Router)
	ListTodo(w http.ResponseWriter, r *http.Request)
	GetTodo(w http.ResponseWriter, r *http.Request)
	CreateTodo(w http.ResponseWriter, r *http.Request)
	UpdateTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}
type todoHandler struct {
	BaseHandler
	todoUseCase usecase.TodoUseCase
	userUseCase usecase.UserUseCase
}

func NewTodoHandler(todoUseCase usecase.TodoUseCase, userUseCase usecase.UserUseCase) TodoHandler {
	return &todoHandler{todoUseCase: todoUseCase, userUseCase: userUseCase}
}

func (h *todoHandler) RegisterTodoHandlers(r *mux.Router) {
	todoRouter := r.PathPrefix(constants.TodosPath).Subrouter()
	todoRouter.Use(h.authMiddleware)

	todoRouter.HandleFunc("", h.ListTodo).Methods(http.MethodGet, http.MethodOptions)
	todoRouter.HandleFunc("/{id}", h.GetTodo).Methods(http.MethodGet, http.MethodOptions)
	todoRouter.HandleFunc("", h.CreateTodo).Methods(http.MethodPost, http.MethodOptions)
	todoRouter.HandleFunc("/{id}", h.UpdateTodo).Methods(http.MethodPut, http.MethodOptions)
	todoRouter.HandleFunc("/{id}", h.DeleteTodo).Methods(http.MethodDelete, http.MethodOptions)
}

func (h *todoHandler) ListTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := h.getUserEmail(r)
	user, err := h.userUseCase.GetUserByEmail(ctx, &input.GetUserByEmailInput{Email: email})
	if err != nil {
		h.respondError(w, err)
		return
	}

	output, err := h.todoUseCase.ListTodo(ctx, &input.ListTodoInput{UserID: user.ID})
	if err != nil {
		h.respondError(w, err)
		return
	}

	h.respondJSON(w, http.StatusOK, output)
}

func (h *todoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	email := h.getUserEmail(r)
	user, err := h.userUseCase.GetUserByEmail(ctx, &input.GetUserByEmailInput{Email: email})
	if err != nil {
		h.respondError(w, err)
		return
	}

	todoID, err := uuid.Parse(vars["id"])
	if err != nil {
		h.respondError(w, apperrors.NewValidationError("invalid todo id", err))
		return
	}

	input := &input.GetTodoInput{
		ID:     todoID,
		UserID: user.ID,
	}

	if err := input.Validate(); err != nil {
		h.respondError(w, apperrors.NewValidationError("validation failed", err))
		return
	}

	output, err := h.todoUseCase.GetTodo(ctx, input)
	if err != nil {
		h.respondError(w, err)
		return
	}

	h.respondJSON(w, http.StatusOK, output)
}

func (h *todoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	email := h.getUserEmail(r)
	user, err := h.userUseCase.GetUserByEmail(ctx, &input.GetUserByEmailInput{Email: email})
	if err != nil {
		h.respondError(w, err)
		return
	}

	var input input.CreateTodoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.respondError(w, apperrors.NewValidationError("invalid request body", err))
		return
	}
	input.UserID = user.ID

	if err := input.Validate(); err != nil {
		h.respondError(w, apperrors.NewValidationError("validation failed", err))
		return
	}

	output, err := h.todoUseCase.CreateTodo(ctx, &input)
	if err != nil {
		h.respondError(w, err)
		return
	}

	h.respondJSON(w, http.StatusCreated, output)
}

func (h *todoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	email := h.getUserEmail(r)
	user, err := h.userUseCase.GetUserByEmail(ctx, &input.GetUserByEmailInput{Email: email})
	if err != nil {
		h.respondError(w, err)
		return
	}

	todoID, err := uuid.Parse(vars["id"])
	if err != nil {
		h.respondError(w, apperrors.NewValidationError("invalid todo id", err))
		return
	}

	var input input.UpdateTodoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.respondError(w, apperrors.NewValidationError("invalid request body", err))
		return
	}
	input.ID = todoID
	input.UserID = user.ID

	if err := input.Validate(); err != nil {
		h.respondError(w, apperrors.NewValidationError("validation failed", err))
		return
	}

	output, err := h.todoUseCase.UpdateTodo(ctx, &input)
	if err != nil {
		h.respondError(w, err)
		return
	}

	h.respondJSON(w, http.StatusOK, output)
}

func (h *todoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	email := h.getUserEmail(r)
	user, err := h.userUseCase.GetUserByEmail(ctx, &input.GetUserByEmailInput{Email: email})
	if err != nil {
		h.respondError(w, err)
		return
	}

	todoID, err := uuid.Parse(vars["id"])
	if err != nil {
		h.respondError(w, apperrors.NewValidationError("invalid todo id", err))
		return
	}

	input := &input.DeleteTodoInput{
		ID:     todoID,
		UserID: user.ID,
	}

	if err := input.Validate(); err != nil {
		h.respondError(w, apperrors.NewValidationError("validation failed", err))
		return
	}

	if err := h.todoUseCase.DeleteTodo(ctx, input); err != nil {
		h.respondError(w, err)
		return
	}

	h.respondJSON(w, http.StatusNoContent, nil)
}