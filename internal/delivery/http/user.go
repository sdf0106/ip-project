package delivery

import (
	"encoding/json"
	"github.com/sdf0106/ip-project/internal/domain"
	"github.com/sdf0106/ip-project/internal/dto"
	"github.com/sdf0106/ip-project/pkg/utils"
	"net/http"
)

func (h *Handler) chooseRole(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.RetrieveId(r)

	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input dto.UserRoleInput

	input, err = utils.RequestBodyParser(r, input)

	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.ChooseRole(userId, input.Role); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) updateRole(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.RetrieveId(r)

	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input dto.ChangeRoleInput

	input, err = utils.RequestBodyParser(r, input)

	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.ChangeRole(userId, input.PrevRole, input.Role); err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) updateUserInfo(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.RetrieveId(r)

	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input dto.UpdateUserInput

	input, err = utils.RequestBodyParser(r, input)

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.services.UpdateUserInfo(userId, input.InputToEntity())

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response, err := json.Marshal(map[string]domain.User{
		"user": user,
	})

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
