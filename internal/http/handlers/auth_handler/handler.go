package auth_handler

import (
	"net/http"

	"github.com/kudrmax/perfectPetProject/internal/http/handlers/http_common"
	"github.com/kudrmax/perfectPetProject/internal/http/http_model"
	"github.com/kudrmax/perfectPetProject/internal/services/auth"
)

type (
	authService interface {
		Login(username, password string) (accessToken string, err error)
		Register(name, username, password string) (accessToken string, err error)
	}
)

type Handler struct {
	authService authService
}

func New(
	authService authService,
) *Handler {
	return &Handler{
		authService: authService,
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	body, err := http_common.GetRequestBody[http_model.RegisterRequest](r)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.authService.Register(body.Name, body.Username, body.Password)

	if err != nil {
		switch err {
		case auth.UserAlreadyExistsErr:
			http.Error(w, "user already exists", http.StatusConflict)
			return
		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
			// TODO log
			return
		}
	}

	err = http_common.WriteResponse(w, http.StatusOK, http_model.AuthResponse{AccessToken: token})
	if err != nil {
		// TODO log
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	body, err := http_common.GetRequestBody[http_model.LoginRequest](r)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(body.Username, body.Password)

	if err != nil {
		switch err {
		case auth.UserNotFoundErr:
			http.Error(w, "user not found", http.StatusNotFound)
			return
		case auth.WrongPasswordErr:
			http.Error(w, "wrong password", http.StatusUnauthorized)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			// TODO log
			return
		}
	}

	err = http_common.WriteResponse(w, http.StatusOK, http_model.AuthResponse{AccessToken: token})
	if err != nil {
		// TODO log
	}
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO implement logout
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	// TODO implement refresh
}
