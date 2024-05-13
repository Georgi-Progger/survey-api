package service

import (
	"context"

	"github.com/Georgi-Progger/survey-api/internal/repository"
)

type VideoService struct {
	repo repository.Video
}

func NewVideoService(repo repository.Video) *VideoService {
	return &VideoService{repo: repo}
}

func (s *VideoService) Save(ctx context.Context, vquestionId, userId int, filePath string) error {
	return s.repo.Save(ctx, vquestionId, userId, filePath)
}
