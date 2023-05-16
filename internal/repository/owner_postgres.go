package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sdf0106/ip-project/internal/domain"
)

type OwnerPostgres struct {
	db *pgxpool.Pool
}

func NewOwnerPostgres(db *pgxpool.Pool) *OwnerPostgres {
	return &OwnerPostgres{
		db: db,
	}
}

func (r *OwnerPostgres) GetMyHouses(userId int) ([]domain.House, error) {
	var houses []domain.House

	query := fmt.Sprintf("SELECT ht.id, ht.address, ht.owner_id, ht.agent_id, ht.build_date, ht.price FROM %s ot INNER JOIN %s ht ON at.id = ht.owner_id WHERE at.id = $1 ORDER BY at.id", ownersTable, housesTable)
	rows, err := r.db.Query(context.Background(), query, userId)

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

func (r *OwnerPostgres) CreateHouse(userId int, house domain.House) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (address, owner_id, build_date, price) VALUES ($1, $2, $3, $4) RETURNING ID", housesTable)
	row := r.db.QueryRow(context.Background(), query, house.Address, house.OwnerId, house.BuildDate, house.Price)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *OwnerPostgres) DeleteHouse(ownerId int, houseId int) error {
	query := fmt.Sprintf("DELETE FROM %s ht USING %s ot WHERE ot.id=$1 AND ot.user_id=ht.owner_id AND ht.house_id=$2", housesTable, ownersTable)
	_, err := r.db.Exec(context.Background(), query, ownerId, houseId)

	return err
}

func (r *OwnerPostgres) UpdateHouse(ownerId int, house domain.House) (domain.House, error) {
	//query := fmt.Sprintf("UPDATE %s ht SET ht.agent_id FROM %s ot WHERE ot.userId = $1 AND ot.")

	return house, nil
}

func (r *OwnerPostgres) HireAgent(ownerId int, houseId int, agentId int) (int, error) {

	query := fmt.Sprintf("UPDATE %s AS ht SET ht.agent_id=$1 FROM %s AS ot WHERE ot.user_id=$2 AND ot.id=ht.owner_id AND ht.id=$3", housesTable, ownersTable)
	_, err := r.db.Exec(context.Background(), query, agentId, ownerId, houseId)

	if err != nil {
		return 0, err
	}

	return houseId, nil
}
