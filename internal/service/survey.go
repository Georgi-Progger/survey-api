package service

import (
	"log"
	"net/http"

	model "github.com/Georgi-Progger/survey-api/internal/model"
	survey "github.com/Georgi-Progger/survey-api/internal/repository"
	. "github.com/Georgi-Progger/survey-api/pkg/s3storage"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
)

func (s *Service) InsertCandidate(c echo.Context) error {
	var candidate model.Candidate
	if err := c.Bind(&candidate); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	surveyRepo := survey.NewRepository(s.Db)
	if err := surveyRepo.NewCandidates(c.Request().Context(), candidate); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert candidate"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Candidate successfully created"})
}

func (s *Service) SelectInterview(c echo.Context) error {
	name := c.FormValue("nameInterview")

	surveyRepo := survey.NewRepository(s.Db)
	res, err := surveyRepo.GetInterviewQuestion(c.Request().Context(), name)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func (s *Service) UploadFile(c echo.Context) error {
	sess, err := CreateSession()
	surveyRepo := survey.NewRepository(s.Db)
	if err != nil {
		log.Fatal("Failed to create session:", err)
	}

	svc := s3.New(sess)

	file, err := c.FormFile("file")
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	fileName := file.Filename
	size := file.Size

	buffer := make([]byte, size)
	_, err = src.Read(buffer)
	if err != nil {
		log.Fatal("Failed to read file:", err)
	}

	contentType := http.DetectContentType(buffer)

	err = UploadToS3(svc, buffer, fileName, contentType)
	if err != nil {
		log.Fatal("Failed to upload file to S3:", err)
	}

	err = surveyRepo.SaveVideo(c.Request().Context(), fileName)
	if err != nil {
		log.Println("Failed to save video in DB:", err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "File uploaded successfully"})
}
