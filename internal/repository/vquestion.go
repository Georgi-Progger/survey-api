package repository

import (
	"database/sql"

	"github.com/Georgi-Progger/survey-api/internal/model"
)

type VQuestionRepository struct {
	db *sql.DB
}

func NewVQuestionRepository(db *sql.DB) *VQuestionRepository {
	return &VQuestionRepository{db: db}
}

func (r *VQuestionRepository) GetAll() ([]model.VQuestion, error) {
	query := "SELECT * FROM video_question;"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var vqsts []model.VQuestion
	for rows.Next() {
		vqst := model.VQuestion{}
		err := rows.Scan(&vqst.Id, &vqst.QuestionText)
		if err != nil {
			return nil, err
		}
		vqsts = append(vqsts, vqst)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return vqsts, nil
}

func (r *VQuestionRepository) GetAllByUserIdWithQuestions(userId int) ([]model.VQuestionAndAnswer, error) {
	query := `
	SELECT video_path, question FROM public.question_answer q
	JOIN video_question AS v ON v.id = q.video_question_id
	WHERE  q.user_id = $1
	`

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]model.VQuestionAndAnswer, 0)
	for rows.Next() {
		var row model.VQuestionAndAnswer
		err = rows.Scan(&row.AnswerPath, &row.Question)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}
