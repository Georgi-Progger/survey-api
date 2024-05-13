package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

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
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": "Failed to open file: " + err.Error() })
	}

	svc := s3.New(sess)

	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Failed to open file:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": "Failed to open file: " + err.Error()})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": "Failed to read file: " + err.Error()})
	}
	defer src.Close()
	ext := strings.Split(file.Filename, ".")
	fileName := uuid.New().String() + "." + ext[len(ext) - 1]
	size := file.Size
	log.Println(ext[len(ext) - 1])

	buffer := make([]byte, size)
	_, err = src.Read(buffer)
	if err != nil {
		log.Println("Failed to read file:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": "Failed to read file: " + err.Error()})
	}

	contentType := http.DetectContentType(buffer)

	rawData := c.FormValue("data")

	var candidate model.Candidate
	if err := json.Unmarshal([]byte(rawData), &candidate); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON provided: " + err.Error()})
	}
	candidate.ResumePath = fileName
	candidate.UserId = jwt.GetUserIdFromContext(c)
	if err := h.services.Candidate.Create(c.Request().Context(), candidate); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to insert candidate: " + err.Error()})
	}

	err = UploadToS3(svc, buffer, fileName, contentType)
	if err != nil {
		log.Println("Failed to upload file to S3:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": "Failed to upload file to S3: " + err.Error()})
	}
	log.Println(fileName)
	return c.JSON(http.StatusCreated, map[string]string{"message": "Candidate successfully created: " + err.Error()})
}
