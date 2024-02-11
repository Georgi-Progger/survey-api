package s3storage

import (
	"bytes"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

var (
	accessKeyID     string
	secretAccessKey string
	region          string
	bucketName      string
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	accessKeyID = os.Getenv("ACCESS_KEY_ID")
	secretAccessKey = os.Getenv("SECRET_ACCESS_KEY")
	region = os.Getenv("REGION_S3")
	bucketName = os.Getenv("S3_BUCKET")
}

func CreateSession() (*session.Session, error) {
	Init()
	s3Endpoint := "https://s3.timeweb.com"

	return session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Endpoint:    aws.String(s3Endpoint),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})
}

func ReadFile(file *os.File) ([]byte, string, string, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, "", "", err
	}
	size := fileInfo.Size()
	buffer := make([]byte, size)
	_, err = file.Read(buffer)
	if err != nil {
		return nil, "", "", err
	}

	contentType := http.DetectContentType(buffer)

	return buffer, fileInfo.Name(), contentType, nil
}

func UploadToS3(svc *s3.S3, buffer []byte, fileName, contentType string) error {
	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(bucketName),
		Key:           aws.String(fileName),
		ACL:           aws.String("public-read"),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(int64(len(buffer))),
		ContentType:   aws.String(contentType),
	})
	return err
}
