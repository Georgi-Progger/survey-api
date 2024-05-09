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

func (r *UserRepository) GetAllWithRole(roleId int) ([]model.UserWithInfo, error) {
	query := `
	SELECT u.id, u.role_id, u.phonenumber, c.first_name, c.last_name, c.middle_name, c.date_of_birth, c.city, c.education, c.reason_dismissal, c.email, c.year_work_experience, c.resume_path, c.creation_date
	FROM users u
	JOIN candidates c ON c.user_id = u.id
	WHERE role_id = $1;
`
	rows, err := r.db.Query(query, roleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.UserWithInfo
	for rows.Next() {
		var user model.UserWithInfo

		err = rows.Scan(&user.Id, &user.RoleId, &user.Phonenumber, &user.FirstName, &user.LastName, &user.MiddleName, &user.DateOfBirth, &user.City, &user.Education, &user.ReasonDismissal, &user.Email, &user.YearWorkExperience, &user.ResumePath, &user.CreationDate)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) Update(user model.User) error {
	query := "UPDATE public.users SET role_id=$1, phonenumber=$2, email=$3, password=$4 WHERE id=$5;"
	_, err := r.db.Exec(query, user.RoleId, user.Phonenumber, user.Email, user.Password, user.Id)
	return err
}
