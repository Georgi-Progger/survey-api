package repository

import (
	"context"
	"database/sql"

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
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
`

	_, err := r.db.ExecContext(ctx, query, candidate.FirstName, candidate.LastName,
		candidate.MiddleName, candidate.BirthDate,
		candidate.City, candidate.Education, candidate.ReasonDismissal,
		candidate.Email, candidate.YearWorkExperience,
		candidate.EmployeeInfo, candidate.UserId, candidate.ResumePath)
	return err
}
