package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sdf0106/ip-project/internal/domain"
)

type ClientCartPostgres struct {
	db *pgxpool.Pool
}

func NewClientCartPostgres(db *pgxpool.Pool) *ClientCartPostgres {
	return &ClientCartPostgres{
		db: db,
	}
}

func (r *ClientCartPostgres) GetCart(clientId int) ([]domain.House, error) {
	var houses []domain.House

	query := fmt.Sprintf("SELECT ht.id, ht.address, ht.owner_id, ht.agent_id, ht.build_date, ht.price FROM %s cct INNER JOIN %s ht ON cct.house_id = ht.id WHERE cct.client_id = $1 ORDER BY ht.id", clientCartTable, housesTable)
	rows, err := r.db.Query(context.Background(), query, clientId)

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

func (r *ClientCartPostgres) AddToCart(clientId int, houseId int) error {
	query := fmt.Sprintf("INSERT INTO %s (client_id, house_id) VALUES ($1, $2)", clientCartTable)
	_, err := r.db.Exec(context.Background(), query, clientId, houseId)

	return err
}

func (r *ClientCartPostgres) RemoveFromCart(clientId int, houseId int) error {
	query := fmt.Sprintf("DELETE FROM %s cct WHERE cct.client_id = $1 AND cct.house_id = $2", clientCartTable)
	_, err := r.db.Exec(context.Background(), query, clientId, houseId)

	return err
}
