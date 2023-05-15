package delivery

import (
	"encoding/json"
	"github.com/sdf0106/ip-project/internal/domain"
	"github.com/sdf0106/ip-project/pkg/utils"
	"net/http"
)

func (h *Handler) getAllAgents(w http.ResponseWriter, r *http.Request) {
	agents, err := h.services.GetAllAgents()

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response, err := json.Marshal(map[string][]domain.Agent{
		"agents": agents,
	})

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (h *Handler) getAgentById(w http.ResponseWriter, r *http.Request) {
	id, err := utils.RetrieveIdFromUrl(r, "id")

	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	agent, err := h.services.GetAgentById(id)

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	response, err := json.Marshal(agent)

	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (h *Handler) getAgentHouses(w http.ResponseWriter, r *http.Request) {
	id, err := utils.RetrieveIdFromUrl(r, "id")

	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	houses, err := h.services.GetAgentHouses(id)

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
