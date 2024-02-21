package repository

import (
	"context"
	"database/sql"

	model "github.com/Georgi-Progger/survey-api/internal/model"
)

type InterviewImpl struct {
	db *sql.DB
}

func NewInterviewImpl(db *sql.DB) *InterviewImpl {
	return &InterviewImpl{db: db}
}

func (r *InterviewImpl) GetInterviewQuestion(ctx context.Context, nameInterview string) (*model.Interview, error) {
	query := `
	SELECT i.id, i.interview_name, iq.id, iq.text_answer
	FROM interviews i
	JOIN interview_questions iq ON i.id = iq.interview_id
	WHERE i.interview_name = $1;
`
	var interview model.Interview
	var question model.InterviewQuestion

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
