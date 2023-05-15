package delivery

import (
	"encoding/json"
	"github.com/sdf0106/ip-project/internal/domain"
	"github.com/sdf0106/ip-project/pkg/utils"
	"net/http"
)

func (h *Handler) getAllHouses(w http.ResponseWriter, r *http.Request) {
	houses, err := h.services.GetAllHouses()

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

func (h *Handler) getHouseById(w http.ResponseWriter, r *http.Request) {
	id, err := utils.RetrieveIdFromUrl(r, "id")

	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	house, err := h.services.GetHouseById(id)

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
