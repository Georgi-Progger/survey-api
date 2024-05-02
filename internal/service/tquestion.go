package service

import (
	model "github.com/Georgi-Progger/survey-api/internal/model"
	"github.com/Georgi-Progger/survey-api/internal/repository"
)

type TQuestionService struct {
	repo repository.TQuestion
}

func NewTQuestionService(repo repository.TQuestion) *TQuestionService {
	return &TQuestionService{repo: repo}
}

func (s *TQuestionService) GetAll() ([]model.TestQuestion, error) {
	return s.repo.GetAll()
}

func (s *TQuestionService) InsertAnswers(userId int, answers []model.UserTestAnswer) error {
	return s.repo.InsertAnswers(userId, answers)
}

func (s *TQuestionService) GetUserAnswers(userId int) ([]model.UserTestAnswer, error) {
	return s.repo.GetUserAnswers(userId)
}