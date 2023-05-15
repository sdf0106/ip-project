package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sdf0106/ip-project/internal/domain"
)

type AgentPostgres struct {
	db *pgxpool.Pool
}

func NewAgentPostgres(db *pgxpool.Pool) *AgentPostgres {
	return &AgentPostgres{
		db: db,
	}
}

func (r *AgentPostgres) GetAllAgents() ([]domain.Agent, error) {
	var agents []domain.Agent

	query := fmt.Sprintf("SELECT at.id, ut.name, ut.email, at.phone FROM %s ut INNER JOIN %s at ON ut.id = at.user_id ORDER BY at.id", usersTable, agentsTable)
	rows, err := r.db.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var agent domain.Agent
		if err = rows.Scan(&agent.Id, &agent.Name, &agent.Email, &agent.Phone); err != nil {
			return nil, err
		}
		agents = append(agents, agent)
	}

	return agents, nil
}

func (r *AgentPostgres) GetAgentById(id int) (domain.Agent, error) {
	var agent domain.Agent

	query := fmt.Sprintf("SELECT at.id, ut.name, ut.email, at.phone FROM %s ut INNER JOIN %s at ON ut.id = at.user_id WHERE ut.id = $1 ORDER BY at.id", usersTable, agentsTable)
	row, err := r.db.Query(context.Background(), query, id)

	if err != nil {
		return agent, err
	}

	if err = row.Scan(&agent.Id, &agent.Name, &agent.Email, &agent.Phone); err != nil {
		return agent, err
	}

	return agent, nil
}

func (r *AgentPostgres) GetAgentHouses(id int) ([]domain.House, error) {
	var houses []domain.House

	query := fmt.Sprintf("SELECT ht.id, ht.address, ht.owner_id, ht.agent_id, ht.build_date, ht.price FROM %s at INNER JOIN %s ht ON at.id = ht.agent_id WHERE at.id = $1 ORDER BY at.id", agentsTable, housesTable)
	rows, err := r.db.Query(context.Background(), query, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var house domain.House
		if err = rows.Scan(&house.Id, &house.AgentId, &house.OwnerId, &house.Address, &house.BuildDate, &house.Price); err != nil {
			return nil, err
		}
		houses = append(houses, house)
	}

	return houses, nil
}
