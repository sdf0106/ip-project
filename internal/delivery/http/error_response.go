package delivery

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func newErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	logrus.Error(message)
	http.Error(w, message, statusCode)
}
