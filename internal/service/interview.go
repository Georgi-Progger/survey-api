package service

import (
	"context"

	model "github.com/Georgi-Progger/survey-api/internal/model"
	"github.com/Georgi-Progger/survey-api/internal/repository"
)

type InterviewService struct {
	repo repository.Interview
}

func NewInterviewService(repo repository.Interview) *InterviewService {
	return &InterviewService{repo: repo}
}

func (s *InterviewService) GetInterviewQuestion(ctx context.Context, nameInterview string) (*model.Interview, error) {
	return s.repo.GetInterviewQuestion(ctx, nameInterview)
}
