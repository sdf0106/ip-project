package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sdf0106/ip-project/internal/domain"
)

type HousePostgres struct {
	db *pgxpool.Pool
}

func NewHousePostgres(db *pgxpool.Pool) *HousePostgres {
	return &HousePostgres{db: db}
}

func (r *HousePostgres) GetAllHouses() ([]domain.House, error) {
	var houses []domain.House

	query := fmt.Sprintf("SELECT ht.id, ht.address, ht.owner_id, ht.agent_id, ht.build_date, ht.price FROM %s ht ORDER BY ht.id", housesTable)
	rows, err := r.db.Query(context.Background(), query)

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
func (r *HousePostgres) GetHouseById(houseId int) (domain.House, error) {
	var house domain.House

	query := fmt.Sprintf("SELECT ht.id, ht.address, ht.owner_id, ht.agent_id, ht.build_date, ht.price FROM %s ht WHERE ht.id=$1", housesTable)
	row := r.db.QueryRow(context.Background(), query, houseId)

	if err := row.Scan(&house.Id, &house.Address, &house.OwnerId, &house.AgentId, &house.BuildDate, &house.Price); err != nil {
		return house, err
	}

	return house, nil
}
