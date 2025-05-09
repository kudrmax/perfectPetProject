package auth_middleware

import (
	"context"
	"net/http"
	"strings"
)

type (
	authService interface {
		ValidateTokenAndGetUserId(token string) (userId int, err error)
	}
)

type Handler struct {
	authService authService
}

func NewHandler(
	authService authService,
) *Handler {
	return &Handler{
		authService: authService,
	}
}

func (h *Handler) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := h.extractBearerToken(
			r.Header.Get("Authorization"),
		)

		userId, err := h.authService.ValidateTokenAndGetUserId(tokenStr)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "userId", userId)

		next(w, r.WithContext(ctx))
	}
}

func (h *Handler) extractBearerToken(header string) string {
	const prefix = "Bearer "
	if strings.HasPrefix(header, prefix) {
		return strings.TrimPrefix(header, prefix)
	}
	return ""
}
