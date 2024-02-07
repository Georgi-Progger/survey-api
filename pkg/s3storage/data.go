package s3storage

import "io"

type config struct {
	accessKeyID     string
	secretAccessKey string
	region          string
	bucketName      string
}

func NewConfig(accessKeyId string, secretAccessKey string, region string, bucketName string) *config {
	return &config{
		accessKeyID:     accessKeyId,
		secretAccessKey: secretAccessKey,
		region:          region,
		bucketName:      bucketName,
	}
}

type fileInfo struct {
	fileFullName string
	body         io.ReadSeeker
}

func NewFile(fileName string, body io.ReadSeeker) *fileInfo {
	return &fileInfo{
		fileFullName: fileName,
		body:         body}
}
