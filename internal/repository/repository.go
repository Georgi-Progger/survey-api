package repository

import (
	"context"
	"database/sql"

	model "github.com/Georgi-Progger/survey-api/internal/model"
)

type Candidate interface {
	Create(ctx context.Context, candidate model.Candidate) error
}

type Role interface {
	GetByName(ctx context.Context, name string) (*model.Role, error)
}

type User interface {
	Save(ctx context.Context, user model.User) (int, error)
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
	Role
	User
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Candidate: NewCandidateImpl(db),
		Interview: NewInterviewImpl(db),
		Video:     NewVideoImpl(db),
		Role:      NewRoleRepository(db),
		User:      NewUserRepository(db),
	}
}
