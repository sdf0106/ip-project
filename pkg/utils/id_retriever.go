package utils

import (
	"errors"
	"net/http"
)

func RetrieveId(r *http.Request) (int, error) {
	userId, ok := r.Context().Value("id").(int)

	if !ok {
		return 0, errors.New("user id not found")
	}

	return userId, nil
}
