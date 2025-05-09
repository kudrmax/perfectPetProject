package login

import (
	"encoding/json"
	"net/http"

	"github.com/kudrmax/perfectPetProject/internal/services/auth"
)

type (
	authService interface {
		Login(username, password string) (accessToken string, err error)
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

type Request struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Response struct {
	AccessToken string `json:"accessToken"`
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	var body Request
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
	json.NewEncoder(w).Encode(Response{AccessToken: token})
}
