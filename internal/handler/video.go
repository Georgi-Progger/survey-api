package handler

import (
	"log"
	"net/http"

	. "github.com/Georgi-Progger/survey-api/pkg/s3storage"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
)

func (h *Handler) UploadFile(c echo.Context) error {
	sess, err := CreateSession()
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

	err = h.services.Video.Save(c.Request().Context(), fileName)
	if err != nil {
		log.Println("Failed to save video in DB:", err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "File uploaded successfully"})
}
