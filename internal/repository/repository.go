package repository

import (
	"context"
	"database/sql"

	model "github.com/Georgi-Progger/survey-api/internal/model"
)

type Candidate interface {
	Create(ctx context.Context, candidate model.Candidate) error
}

type Interview interface {
	GetInterviewQuestion(ctx context.Context, nameInterview string) (*model.Interview, error)
}

type Video interface {
	Save(ctx context.Context, filePath string) error
}

type Repository struct {
	Candidate
	Interview
	Video
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Candidate: NewCandidateImpl(db),
		Interview: NewInterviewImpl(db),
		Video:     NewVideoImpl(db),
	}
}

// func (r *Repository) NewCandidates(ctx context.Context, candidate model.Candidate) error {
// 	query := `
//       INSERT INTO candidates (first_name, last_name,
//         middle_name, date_of_birth, city, education, reason_dismissal,
//         email, phone, year_work_experience, employee_entered_info)
//       VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
//   `

// 	_, err := r.db.ExecContext(ctx, query, candidate.FirstName, candidate.LastName,
// 		candidate.MiddleName, candidate.BirthDate,
// 		candidate.City, candidate.Education, candidate.ReasonDismissal,
// 		candidate.Email, candidate.PhoneNumper, candidate.YearWorkExperience,
// 		candidate.EmployeeInfo)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	return nil
// }

// func (r *Repository) GetInterviewQuestion(ctx context.Context, nameInterview string) (*model.Interview, error) {
// 	query := `
// 		SELECT i.id, i.interview_name, iq.id, iq.text_answer
// 		FROM interviews i
// 		JOIN interview_questions iq ON i.id = iq.interview_id
// 		WHERE i.interview_name = $1;
// 	`
// 	var interview model.Interview
// 	var question model.InterviewQuestion

// 	rows, err := r.db.QueryContext(ctx, query, nameInterview)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		err := rows.Scan(&interview.Id, &interview.InterviewName, &question.Id, &question.TextAnswer)
// 		if err != nil {
// 			return nil, err
// 		}
// 		interview.Questions = append(interview.Questions, question)
// 	}

// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return &interview, nil
// }

// func (r *Repository) SaveVideo(ctx context.Context, filePath string) error {
// 	query := `
// 		INSERT INTO videos (file_path) VALUES ($1)
// 	`
// 	_, err := r.db.ExecContext(ctx, query, filePath)
// 	if err != nil {
// 		log.Println("Failed to save video:", err)
// 		return err
// 	}

// 	return nil
// }
