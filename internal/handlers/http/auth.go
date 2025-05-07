package http

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/openapi3filter"

	"github.com/kudrmax/perfectPetProject/internal/handlers/http/api"
	"github.com/kudrmax/perfectPetProject/internal/services/auth"
)

func (h *Handler) RegisterUser(ctx context.Context, request api.RegisterUserRequestObject) (api.RegisterUserResponseObject, error) {
	token, err := h.authService.Register(request.Body.Name, request.Body.Username, request.Body.Password)

	if err != nil {
		switch err {
		case auth.UserAlreadyExistsErr:
			return api.RegisterUser409Response{}, nil
		default:
			return api.RegisterUser500JSONResponse{}, nil
		}
	}

	return api.RegisterUser201JSONResponse{
		AuthResponseJSONResponse: api.AuthResponseJSONResponse{
			AccessToken: token,
		},
	}, nil
}

func (h *Handler) LoginUser(ctx context.Context, request api.LoginUserRequestObject) (api.LoginUserResponseObject, error) {
	token, err := h.authService.Login(request.Body.Username, request.Body.Password)

	if err != nil {
		switch err {
		case auth.UserNotFoundErr:
			return api.LoginUser404JSONResponse{}, nil
		case auth.WrongPasswordErr:
			return api.LoginUser401Response{}, nil
		default:
			return api.LoginUser500JSONResponse{}, nil
		}
	}

	return api.LoginUser200JSONResponse{
		AuthResponseJSONResponse: api.AuthResponseJSONResponse{
			AccessToken: token,
		},
	}, nil
}

func (h *Handler) LogoutUser(ctx context.Context, request api.LogoutUserRequestObject) (api.LogoutUserResponseObject, error) {
	// TODO удалять userId из контекста
	// TODO делать токен невалидным
	userId := ctx.Value("userId")
	_ = userId
	return api.LogoutUser204Response{}, nil
}

func AuthMiddleware2(authService *auth.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "missing or invalid Authorization header", http.StatusUnauthorized)
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			userID, err := authService.ValidateTokenAndGetUserId(tokenStr)
			if err != nil {
				http.Error(w, "invalid or expired token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "userId", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (h *Handler) AuthMiddleware(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	if input.SecuritySchemeName != "BearerAuth" {
		return fmt.Errorf("unsupported security scheme: %s", input.SecuritySchemeName)
	}

	tokenStr := h.extractBearerToken(
		input.RequestValidationInput.Request.Header.Get("Authorization"),
	)

	userId, err := h.authService.ValidateTokenAndGetUserId(tokenStr)
	if err != nil {
		return fmt.Errorf("unauthorized: %w", err)
	}

	ctxWithUser := context.WithValue(ctx, "userId", userId)
	input.RequestValidationInput.Request = input.RequestValidationInput.Request.WithContext(ctxWithUser)
	ctx = ctxWithUser

	return nil
}

func (h *Handler) extractBearerToken(header string) string {
	const prefix = "Bearer "
	if strings.HasPrefix(header, prefix) {
		return strings.TrimPrefix(header, prefix)
	}
	return ""
}
