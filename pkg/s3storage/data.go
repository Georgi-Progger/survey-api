package s3storage

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
