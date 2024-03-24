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
