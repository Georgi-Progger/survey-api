package repository

import (
	"context"
	"database/sql"
	"log"
)

type VideoImpl struct {
	db *sql.DB
}

func NewVideoImpl(db *sql.DB) *VideoImpl {
	return &VideoImpl{db: db}
}

func (r *VideoImpl) Save(ctx context.Context, vquestionId, userId int, filePath string) error {
	query := `
	INSERT INTO question_answer(
		video_question_id, user_id, video_path)
		VALUES ($1, $2, $3);
`
	_, err := r.db.ExecContext(ctx, query, vquestionId, userId, filePath)
	if err != nil {
		log.Println("Failed to save video:", err)
		return err
	}

	return nil
}
