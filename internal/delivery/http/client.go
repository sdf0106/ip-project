package delivery

import (
	"encoding/json"
	"github.com/sdf0106/ip-project/internal/domain"
	"github.com/sdf0106/ip-project/internal/dto"
	"github.com/sdf0106/ip-project/pkg/utils"
	"net/http"
)

func (h *Handler) getCart(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.RetrieveId(r)

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	houses, err := h.services.GetCart(userId)

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response, err := json.Marshal(map[string][]domain.House{
		"cart": houses,
	})

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (h *Handler) addToCart(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.RetrieveId(r)

	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input dto.HouseInCartInput

	input, err = utils.RequestBodyParser(r, input)

	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.AddToCart(userId, input.HouseId); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) removeFromCart(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.RetrieveId(r)

	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input dto.HouseInCartInput

	input, err = utils.RequestBodyParser(r, input)

	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.RemoveFromCart(userId, input.HouseId); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
