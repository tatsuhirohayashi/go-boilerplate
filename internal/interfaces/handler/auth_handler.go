package handler

import (
	"encoding/json"
	"go-boilerplate/internal/pkg/constants"
	apperrors "go-boilerplate/internal/pkg/errors"
	"go-boilerplate/internal/usecase"
	"go-boilerplate/internal/usecase/input"
	"net/http"

	"github.com/gorilla/mux"
)

type AuthHandler interface {
	RegisterAuthHandlers(r *mux.Router)
	Login(w http.ResponseWriter, r *http.Request)
	Signup(w http.ResponseWriter, r *http.Request)
	CheckAuthentication(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	BaseHandler
	authUseCase usecase.AuthUseCase
}

func NewAuthHandler(authUseCase usecase.AuthUseCase) AuthHandler {
	return &authHandler{authUseCase: authUseCase}
}

func (h *authHandler) RegisterAuthHandlers(r *mux.Router) {
	authRouter := r.PathPrefix(constants.AuthPath).Subrouter()
	isAuthCheckRouter := r.PathPrefix(constants.AuthPath).Subrouter()
	isAuthCheckRouter.Use(h.authMiddleware)

	authRouter.HandleFunc("/login", h.Login).Methods(http.MethodPost, http.MethodOptions)
	authRouter.HandleFunc("/signup", h.Signup).Methods(http.MethodPost, http.MethodOptions)
	isAuthCheckRouter.HandleFunc("/authentication", h.CheckAuthentication).Methods(http.MethodPost, http.MethodOptions)
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := &input.LoginInput{}
	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		h.respondError(w, apperrors.NewValidationError("invalid request body", err))
		return
	}

	output, err := h.authUseCase.Login(ctx, input)
	if err != nil {
		h.respondError(w, err)
		return
	}

	h.respondJSON(w, http.StatusOK, output)
}

func (h *authHandler) Signup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := &input.RegisterUserInput{}
	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		h.respondError(w, apperrors.NewValidationError("invalid request body", err))
		return
	}

	output, err := h.authUseCase.RegisterUser(ctx, input)
	if err != nil {
		h.respondError(w, err)
		return
	}

	h.respondJSON(w, http.StatusCreated, output)
}

func (h *authHandler) CheckAuthentication(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := h.getUserEmail(r)

	output, err := h.authUseCase.CheckAuthentication(ctx, &input.CheckAuthenticationInput{Email: email})
	if err != nil {
		h.respondError(w, err)
		return
	}

	h.respondJSON(w, http.StatusOK, output)
}
