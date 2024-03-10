package repository

import (
	"context"
	"database/sql"

	model "github.com/Georgi-Progger/survey-api/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(ctx context.Context, user model.User) (int, error) {
	query := "INSERT INTO users(role_id, phonenumber, email, password) VALUES($1, $2, $3, $4) RETURNING id;"
	id := 0
	rows := r.db.QueryRowContext(ctx, query, user.RoleId, user.Phonenumber, user.Email, user.Password)
	if err := rows.Scan(&id); err != nil {
		return 0, err
	}
	if err := rows.Err(); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepository) GetUserByPhonenumber(phonenumber string) (model.User, error) {
	query := "SELECT * FROM users WHERE phonenumber=$1;"
	rows := r.db.QueryRow(query, phonenumber)
	user := model.User{}
	err := rows.Scan(&user.Id, &user.RoleId, &user.Phonenumber, &user.Email, &user.Password)
	return user, err
}