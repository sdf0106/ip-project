package delivery

import (
	"encoding/json"
	"errors"
	"github.com/sdf0106/ip-project/internal/domain"
	"github.com/sdf0106/ip-project/internal/dto"
	"github.com/sdf0106/ip-project/pkg/utils"
	"net/http"
)

func (h *Handler) getMyHouses(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("id").(int)

	if !ok {
		newErrorResponse(w, http.StatusBadRequest, errors.New("id of the user can't be gotten").Error())
		return
	}

	houses, err := h.services.GetMyHouses(userId)

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response, err := json.Marshal(map[string][]domain.House{
		"houses": houses,
	})

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (h *Handler) createHouse(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.RetrieveId(r)

	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input dto.CreateHouseInput

	input, err = utils.RequestBodyParser(r, input)

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	input.OwnerId = userId

	houseId, err := h.services.CreateHouse(userId, input.InputToEntity())

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response, err := json.Marshal(map[string]int{
		"id": houseId,
	})

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *Handler) deleteHouse(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.RetrieveId(r)

	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	houseId, err := utils.RetrieveIdFromUrl(r, "id")

	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	h.services.DeleteHouse(userId, houseId)

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) updateHouse(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.RetrieveId(r)

	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input dto.UpdateHouseInput

	input, err = utils.RequestBodyParser(r, input)

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	house, err := h.services.UpdateHouse(userId, input.InputToEntity())

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response, err := json.Marshal(house)

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *Handler) hireAgent(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.RetrieveId(r)

	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input dto.HireAgentInput

	input, err = utils.RequestBodyParser(r, input)

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.HireAgent(userId, input.HouseId, input.AgentId)

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response, err := json.Marshal(map[string]int{
		"id": id,
	})

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
