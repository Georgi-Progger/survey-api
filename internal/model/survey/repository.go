package survey

import (
	"context"
	"database/sql"
	"log"
)

type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *repo {
	return &repo{
		db: db,
	}
}

func (r *repo) NewCandidates(ctx context.Context, candidate Candidate) error {
	query := `
      INSERT INTO candidates (first_name, last_name, 
        middle_name, date_of_birth, city, education, reason_dismissal,
        email, phone, year_work_experience, employee_entered_info)
      VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
  `

	_, err := r.db.ExecContext(ctx, query, candidate.FirstName, candidate.LastName,
		candidate.MiddleName, candidate.BirthDate,
		candidate.City, candidate.Education, candidate.ReasonDismissal,
		candidate.Email, candidate.PhoneNumper, candidate.YearWorkExperience,
		candidate.EmployeeInfo)
	if err != nil {
		log.Panic(err)
	}
	return nil
}

func (r *repo) GetInterviewQuestion(ctx context.Context, nameInterview string) (*Interview, error) {
	query := `
		SELECT i.id, i.interview_name, iq.id, iq.text_answer
		FROM interviews i
		JOIN interview_questions iq ON i.id = iq.interview_id
		WHERE i.interview_name = $1;
	`
	var interview Interview
	var question InterviewQuestion

	rows, err := r.db.QueryContext(ctx, query, nameInterview)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&interview.Id, &interview.InterviewName, &question.Id, &question.TextAnswer)
		if err != nil {
			return nil, err
		}
		interview.Questions = append(interview.Questions, question)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &interview, nil
}

// func (r *repo) SaveVideo(ctx context.Context, filePath string) {
// 	query = `
// 			INSERT
// 	`
// }
