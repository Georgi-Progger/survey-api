package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	model "github.com/Georgi-Progger/survey-api/internal/model"
)

type CandidateImpl struct {
	db *sql.DB
}

func NewCandidateImpl(db *sql.DB) *CandidateImpl {
	return &CandidateImpl{db: db}
}

func (r *CandidateImpl) Create(ctx context.Context, candidate model.Candidate) error {
	query := `
	INSERT INTO candidates (first_name, last_name, 
	  middle_name, date_of_birth, city, education, reason_dismissal,
	  email, year_work_experience, employee_entered_info, user_id, resume_path)
	VALUES ($1, $2, $3, $4::date, $5, $6, $7, $8, $9, $10, $11, $12)
`

	layout := "02.01.2006"
	date, err := time.Parse(layout, candidate.BirthDate)
	if err != nil {
		fmt.Println("Ошибка при преобразовании строки во временной формат")
		return fmt.Errorf("не верный формат даты")
	}

	dateForPostgres := date.Format("2006-01-02")

	_, err = r.db.ExecContext(ctx, query, candidate.FirstName, candidate.LastName,
		candidate.MiddleName, dateForPostgres,
		candidate.City, candidate.Education, candidate.ReasonDismissal,
		candidate.Email, candidate.YearWorkExperience,
		candidate.EmployeeInfo, candidate.UserId, candidate.ResumePath)
	return err
}

func (r *CandidateImpl) GetByUserId(id int) (model.Candidate, error) {
	query := "SELECT * FROM candidates WHERE user_id=$1;"
	rows := r.db.QueryRow(query, id)
	cnd := model.Candidate{}
	err := rows.Scan(&cnd.Id, &cnd.FirstName, &cnd.LastName,
		&cnd.MiddleName, &cnd.BirthDate,
		&cnd.City, &cnd.Education, &cnd.ReasonDismissal,
		&cnd.Email, &cnd.YearWorkExperience,
		&cnd.EmployeeInfo, &cnd.UserId, &cnd.ResumePath, &cnd.CreationDate)
	return cnd, err
}
