package service

import (
	"github.com/Georgi-Progger/survey-api/internal/model"
	"github.com/Georgi-Progger/survey-api/internal/repository"
)

type VQuestionService struct {
	repo repository.VQuestion
}

func NewVQuestionService(repo repository.VQuestion) *VQuestionService {
	return &VQuestionService{repo: repo}
}

func (s *VQuestionService) GetAll() ([]model.VQuestion, error) {
	return s.repo.GetAll()
}

func (s *VQuestionService) 	GetAllByUserIdWithQuestions(userId int) ([]model.VQuestionAndAnswer, error) {
	return s.repo.GetAllByUserIdWithQuestions(userId)
}