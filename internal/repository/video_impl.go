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

func (r *VideoImpl) Save(ctx context.Context, filePath string) error {
	query := `
	INSERT INTO videos (file_path) VALUES ($1)
`
	_, err := r.db.ExecContext(ctx, query, filePath)
	if err != nil {
		log.Println("Failed to save video:", err)
		return err
	}

	return nil
}
