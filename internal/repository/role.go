package repository

import (
	"context"
	"database/sql"

	model "github.com/Georgi-Progger/survey-api/internal/model"
)

type RoleImpl struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) *RoleImpl {
	return &RoleImpl{db: db}
}

func (r *RoleImpl) GetByName(ctx context.Context, name string) (*model.Role, error) {
	query := "SELECT * FROM roles WHERE name=$1;"

	rows, err := r.db.QueryContext(ctx, query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	role := model.Role{}
	for rows.Next() {
		err := rows.Scan(&role.Id, &role.Name)
		if err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &role, nil
}
