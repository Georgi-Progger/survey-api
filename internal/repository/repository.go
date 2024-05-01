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
	SetRole(userId, roleId int) error
}

type User interface {
	Save(ctx context.Context, user model.User) (int, error)
	GetUserByPhonenumber(phonenumber string) (model.User, error)
	GetAllWithRole(roleId int) ([]model.UserWithName, error)
	Update(user model.User) error
}

type Interview interface {
	GetInterviewQuestion(ctx context.Context, nameInterview string) (*model.Interview, error)
}

type Video interface {
	Save(ctx context.Context, vquestionId, userId int, filePath string) error
}

type VQuestion interface {
	GetAll() ([]model.VQuestion, error)
}

type TQuestion interface {
	GetAll() ([]model.TestQuestion, error)
	InsertAnswers(userId int, answers []model.UserTestAnswer) error
	GetUserAnswers(userId int) ([]model.UserTestAnswer, error)
}

type Repository struct {
	Candidate
	Interview
	Video
	Role
	User
	VQuestion
	TQuestion
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Candidate: NewCandidateImpl(db),
		Interview: NewInterviewImpl(db),
		Video:     NewVideoImpl(db),
		Role:      NewRoleRepository(db),
		User:      NewUserRepository(db),
		VQuestion: NewVQuestionRepository(db),
		TQuestion: NewTQuestionRepository(db),
	}
}
