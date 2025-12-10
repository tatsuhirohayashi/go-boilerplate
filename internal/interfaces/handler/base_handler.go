package handler

import (
	"context"
	"encoding/json"
	"errors"
	apperrors "go-boilerplate/internal/pkg/errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type BaseHandler struct{}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type contextKey string

const userContextKey contextKey = "user"

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (h *BaseHandler) respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		h.respondError(w, apperrors.NewInternalError("failed to marshal response", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

// エラーレスポンスを返す共通メソッド
func (h *BaseHandler) respondError(w http.ResponseWriter, err error) {
	var status int
	var response ErrorResponse

	// アプリケーションのエラー型の場合
	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		switch appErr.Type {
		case apperrors.NotFound:
			status = http.StatusNotFound
		case apperrors.ValidationError:
			status = http.StatusBadRequest
		case apperrors.PermissionDenied:
			status = http.StatusForbidden
		case apperrors.Unauthorized:
			status = http.StatusUnauthorized
		case apperrors.AlreadyExists:
			status = http.StatusConflict
		case apperrors.BusinessRuleError:
			status = http.StatusUnprocessableEntity
		default:
			status = http.StatusInternalServerError
		}

		response = ErrorResponse{
			Code:    string(appErr.Type),
			Message: appErr.Message,
		}
	} else {
		// 未知のエラーの場合
		status = http.StatusInternalServerError
		response = ErrorResponse{
			Code:    string(apperrors.InternalError),
			Message: "internal server error",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func (h *BaseHandler) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			h.respondError(w, apperrors.NewUnauthorizedError("authorization header is required", nil))
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			h.respondError(w, apperrors.NewUnauthorizedError("invalid authorization header format", nil))
			return
		}
		claims := &Claims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			h.respondError(w, apperrors.NewUnauthorizedError("invalid token", err))
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, claims.Email)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (h *BaseHandler) getUserEmail(r *http.Request) string {
	email, ok := r.Context().Value(userContextKey).(string)
	if !ok {
		return ""
	}
	return email
}