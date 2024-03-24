package handler

import (
	"log"
	"net/http"
	"strconv"

	. "github.com/Georgi-Progger/survey-api/pkg/s3storage"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) UploadFile(c echo.Context) error {
	vquestId, err := strconv.Atoi(c.FormValue("vquestId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"err": "Not all form values are set"})
	}
	userId, err := strconv.Atoi(c.FormValue("userId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"err": "Not all form values are set"})
	}
	sess, err := CreateSession()
	if err != nil {
		log.Fatal("Failed to create session:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": "Failed to open file"})
	}

	svc := s3.New(sess)

	file, err := c.FormFile("file")
	if err != nil {
		log.Fatal("Failed to open file:", err)
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
		log.Fatal("Failed to read file:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": "Failed to read file"})
	}

	contentType := http.DetectContentType(buffer)

	err = UploadToS3(svc, buffer, fileName, contentType)
	if err != nil {
		log.Fatal("Failed to upload file to S3:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": "Failed to upload file to S3"})
	}

	err = h.services.Video.Save(c.Request().Context(), vquestId, userId, fileName)
	if err != nil {
		log.Println("Failed to save video in DB:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": "Failed to save video in DB"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "File uploaded successfully"})
}
