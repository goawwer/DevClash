package s3

import (
	"context"
	"io"

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
		Secure: true,
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

func Upload(ctx context.Context, prefix string, filename string, r io.Reader, size int64, contentType string) (string, error) {
	key := prefix + "/" + filename

	_, err := s3.client.PutObject(ctx, s3.bucketName, key, r, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", err
	}

	return key, nil
}
