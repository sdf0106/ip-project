package delivery

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const authHeader = "Authorization"

func (h *Handler) authMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := getTokenFromHeader(r)

		if err != nil {
			logrus.Error(err.Error())

			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		userId, err := h.services.ParseToken(token)

		if err != nil {
			errorLog := fmt.Sprintf("JWT error: %s", err.Error())

			newErrorResponse(w, http.StatusUnauthorized, errorLog)
		}

		ctx := context.WithValue(r.Context(), "id", userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func getTokenFromHeader(r *http.Request) (string, error) {
	header := r.Header.Get(authHeader)

	if header == "" {
		return "", errors.New("empty Authorization Header")
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		return "", errors.New("empty authorization header")
	}

	return headerParts[1], nil
}
