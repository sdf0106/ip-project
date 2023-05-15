package service

import (
	"github.com/sdf0106/ip-project/internal/domain"
	"github.com/sdf0106/ip-project/internal/repository"
)

type AgentService struct {
	repo repository.Agent
}

func NewAgentService(repo repository.Agent) *AgentService {
	return &AgentService{
		repo: repo,
	}
}

func (s *AgentService) GetAllAgents() ([]domain.Agent, error) {
	return s.repo.GetAllAgents()
}
func (s *AgentService) GetAgentById(id int) (domain.Agent, error) {
	return s.repo.GetAgentById(id)
}
func (s *AgentService) GetAgentHouses(id int) ([]domain.House, error) {
	return s.repo.GetAgentHouses(id)
}
