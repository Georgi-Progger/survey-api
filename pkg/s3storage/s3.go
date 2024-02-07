package s3storage

import (
	"bytes"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	accessKeyID     = "ACCESS_KEY_ID"
	secretAccessKey = "SECRET_ACCESS_KEY"
	region          = "REGION_S3"
	bucketName      = "S3_BUCKET"
	fileName        = "example.txt"
)

func СreateSession() (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})
}

func ReadFile(file *os.File) ([]byte, string) {
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	contentType := http.DetectContentType(buffer)

	return buffer, contentType
}

func UploadToS3(svc *s3.S3, buffer []byte, contentType string) error {
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
