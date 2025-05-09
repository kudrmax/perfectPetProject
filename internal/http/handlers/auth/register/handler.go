package register

import (
	"encoding/json"
	"net/http"

	"github.com/kudrmax/perfectPetProject/internal/services/auth"
)

type (
	authService interface {
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

type Request struct {
	Name     string `json:"name"`
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
	json.NewEncoder(w).Encode(Response{AccessToken: token})
}
