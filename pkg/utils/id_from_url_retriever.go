package utils

import (
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

func RetrieveIdFromUrl(r *http.Request, key string) (int, error) {
	idFromRequest := chi.URLParam(r, key)

	return strconv.Atoi(idFromRequest)
}
