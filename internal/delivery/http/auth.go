package delivery

import (
	"encoding/json"
	"github.com/sdf0106/ip-project/internal/dto"
	"github.com/sdf0106/ip-project/pkg/utils"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var input dto.SignUpInput

	input, err := utils.RequestBodyParser(r, input)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.services.CreateUser(input.InputToEntity())
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response, err := json.Marshal(map[string]int{
		"id": userId,
	})

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input dto.SignInInput

	input, err := utils.RequestBodyParser(r, input)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response, err := json.Marshal(map[string]string{
		"token": token,
	})

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}
