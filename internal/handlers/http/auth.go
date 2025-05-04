package http

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3filter"

	"github.com/kudrmax/perfectPetProject/internal/handlers/http/api"
	"github.com/kudrmax/perfectPetProject/internal/services/auth"
	"github.com/kudrmax/perfectPetProject/internal/services/jwt_provider"
)

func extractBearerToken(header string) string {
	const prefix = "Bearer "
	if strings.HasPrefix(header, prefix) {
		return strings.TrimPrefix(header, prefix)
	}
	return ""
}

func MyyAuthFunc(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	jwtTokenDuration := time.Minute * 15
	jwtSecret := "super-secret"
	jwtProviderService := jwt_provider.NewService(jwtSecret, jwtTokenDuration)

	if input.SecuritySchemeName == "BearerAuth" {
		tokenStr := extractBearerToken(
			input.RequestValidationInput.Request.Header.Get("Authorization"),
		)

		userId, err := jwtProviderService.ParseToken(tokenStr)
		if err != nil {
			return fmt.Errorf("unauthorized: %w", err)
		}

		// Прокинем userID в context для последующего использования
		ctxWithUser := context.WithValue(ctx, "userId", userId)
		input.RequestValidationInput.Request.WithContext(ctxWithUser)
		return nil
	}
	return fmt.Errorf("unsupported security scheme: %s", input.SecuritySchemeName)
}

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
