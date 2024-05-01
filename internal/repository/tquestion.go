package repository

import (
	"database/sql"

	"github.com/Georgi-Progger/survey-api/internal/model"
)

type TQuestionRepository struct {
	db *sql.DB
}

func NewTQuestionRepository(db *sql.DB) *TQuestionRepository {
	return &TQuestionRepository{db: db}
}

func (r *TQuestionRepository) GetAll() ([]model.TestQuestion, error) {
	query := `SELECT test_question.id, question, test_answer.id, answer  FROM test_question
	LEFT JOIN test_answer ON test_question_id = test_question.id;`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []model.TestQuestion
	prevId := -1
	for rows.Next() {
		var tq model.TestQuestion
		var ta model.TestQuestionAnswer
		err = rows.Scan(&tq.Id, &tq.Question, &ta.Id, &ta.Answer)
		if err != nil {
			return nil, err
		}

		if tq.Id != prevId {
			prevId = tq.Id
			res = append(res, tq)
		}
		res[len(res)-1].Answers = append(res[len(res)-1].Answers, ta)
	}
	return res, nil
}

func (r *TQuestionRepository) InsertAnswers(userId int, answers []model.UserTestAnswer) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	query := "INSERT INTO public.test_user_answer VALUES ($1, $2, $3);"
	for _, ans := range answers {
		_, err = tx.Exec(query, ans.TestQuestionId, userId, ans.TestAnswerId)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (r *TQuestionRepository) GetUserAnswers(userId int) ([]model.UserTestAnswer, error) {
	// query := `SELECT a.test_question_id, a.test_answer_id FROM public.test_user_answer a
	// WHERE a.user_id = 1` 
	return nil, nil // TODO закончить
}