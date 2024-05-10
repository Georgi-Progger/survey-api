package service

import (
	"context"

	model "github.com/Georgi-Progger/survey-api/internal/model"
	"github.com/Georgi-Progger/survey-api/internal/repository"
)

type CandidateService struct {
	repo repository.Candidate
}

func NewCandidateService(repo repository.Candidate) *CandidateService {
	return &CandidateService{repo: repo}
}

func (s *CandidateService) Create(ctx context.Context, candidate model.Candidate) error {
	return s.repo.Create(ctx, candidate)
}

func (s *CandidateService) GetByUserId(id int) (model.Candidate, error) {
	return s.repo.GetByUserId(id)
}
