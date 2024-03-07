package service

import (
	"context"

	model "github.com/Georgi-Progger/survey-api/internal/model"
	"github.com/Georgi-Progger/survey-api/internal/repository"
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

type User interface {
	Save(ctx context.Context, user model.User) (int, error)
}

type Sender interface {
	Send(destination, message string) error
}

type Role interface {
	GetByName(ctx context.Context, name string) (*model.Role, error)
}

type Service struct {
	Candidate
	Interview
	Video
	Sender
	Role
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Candidate: NewCandidateService(repo.Candidate),
		Interview: NewInterviewService(repo.Interview),
		Video:     NewVideoService(repo.Video),
		Sender:    NewSmsSenderService(),
		Role:      NewRoleService(repo.Role),
		User:      NewUserService(repo.User),
	}
}
