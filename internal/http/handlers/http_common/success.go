package http_common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteResponse[T any](w http.ResponseWriter, status int, v T) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("error encode response json: %w", err)
	}
	return nil
}

func GetRequestBody[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("error decode body json: %w", err)
	}
	return v, nil
}
