package s3

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Config struct {
	Endpoint         string `env:"S3_ENDPOINT"`
	Bucket           string `env:"S3_BUCKET_NAME"`
	AccessKeyID      string `env:"S3_ACCESS_KEY_ID"`
	PrivateAccessKey string `env:"S3_PRIVATE_ACCESS_KEY"`
	Region           string `env:"S3_REGION"`
}

type S3Storage struct {
	client     *minio.Client
	bucketName string
}

var s3 *S3Storage

func Init(cfg *Config) error {
	c, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.PrivateAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return err
	}

	s3 = &S3Storage{
		client:     c,
		bucketName: cfg.Bucket,
	}

	return nil
}

func Upload(p *S3UploadFileParameters) (string, error) {
	key := p.Prefix + "/" + p.Filename

	_, err := s3.client.PutObject(p.Ctx, s3.bucketName, key, p.Reader, p.Size, minio.PutObjectOptions{
		ContentType: p.ContentType,
	})

	if err != nil {
		return "", err
	}

	return key, nil
}

func PresignKey(p *S3GetFileParameters) (string, error) {
	presignedURL, err := s3.client.PresignedGetObject(p.Ctx, s3.bucketName, p.FileName, p.Expires, nil)
	return presignedURL.String(), err
}
