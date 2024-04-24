package handler

import (
	"log"
	"net/http"

	model "github.com/Georgi-Progger/survey-api/internal/model"
	"github.com/Georgi-Progger/survey-api/pkg/jwt"
	"github.com/labstack/echo/v4"

	. "github.com/Georgi-Progger/survey-api/pkg/s3storage"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

func (h *Handler) InsertCandidate(c echo.Context) error {
	sess, err := CreateSession()
	if err != nil {
		log.Println("Failed to create session:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": "Failed to open file"})
	}

	svc := s3.New(sess)

	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Failed to open file:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": "Failed to open file"})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": "Failed to read file"})
	}
	defer src.Close()
	fileName := uuid.New().String()
	size := file.Size

	buffer := make([]byte, size)
	_, err = src.Read(buffer)
	if err != nil {
		log.Println("Failed to read file:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": "Failed to read file"})
	}

	contentType := http.DetectContentType(buffer)

	err = UploadToS3(svc, buffer, fileName, contentType)
	if err != nil {
		log.Println("Failed to upload file to S3:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": "Failed to upload file to S3"})
	}

	var candidate model.Candidate
	if err := c.Bind(&candidate); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}
	candidate.ResumePath = fileName
	candidate.UserId = jwt.GetUserIdFromContext(c)
	if err := h.services.Candidate.Create(c.Request().Context(), candidate); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert candidate"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Candidate successfully created"})
}
