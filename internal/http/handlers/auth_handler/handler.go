package auth_handler

import (
	"encoding/json"
	"net/http"

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

func NewHandler(
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

	var body http_model.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := h.authService.Register(body.Name, body.Username, body.Password)

	if err != nil {
		switch err {
		case auth.UserAlreadyExistsErr:
			http.Error(w, err.Error(), http.StatusConflict)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(http_model.AuthResponse{AccessToken: token})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	var body http_model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
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
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(http_model.AuthResponse{AccessToken: token})
}
