package service

import "database/sql"

type Service struct {
	Db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		Db: db,
	}
}
