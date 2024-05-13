package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Georgi-Progger/survey-api/pkg/jwt"
	. "github.com/Georgi-Progger/survey-api/pkg/s3storage"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) UploadFile(c echo.Context) error {
	vquestId, err := strconv.Atoi(c.FormValue("vquestId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"err": "Not all form values are set: " + err.Error()})
	}
	userId := jwt.GetUserIdFromContext(c)
	sess, err := CreateSession()
	if err != nil {
		log.Println("Failed to create session:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": "Failed to open file: " + err.Error()})
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

	err = UploadToS3(svc, buffer, fileName, contentType)
	if err != nil {
		log.Println("Failed to upload file to S3:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": "Failed to upload file to S3: " + err.Error()})
	}

	err = h.services.Video.Save(c.Request().Context(), vquestId, userId, fileName)
	if err != nil {
		log.Println("Failed to save video in DB:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"err": "Failed to save video in DB: " + err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "File uploaded successfully: " + err.Error()})
}
