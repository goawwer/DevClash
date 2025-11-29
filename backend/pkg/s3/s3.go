package s3

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

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

func Delete(p *S3RemoveFileParameters) error {
	return s3.client.RemoveObject(p.Ctx, s3.bucketName, p.Filename, minio.RemoveObjectOptions{})
}

func StorePictureAtS3(ctx context.Context, f multipart.File, h *multipart.FileHeader, entityName string, prefix string) (string, error) {
	defer f.Close()

	filename := fmt.Sprintf("%s-%d%s", entityName, time.Now().Unix(), filepath.Ext(h.Filename))

	return Upload(&S3UploadFileParameters{
		Ctx:         ctx,
		Prefix:      prefix,
		Filename:    filename,
		Reader:      f,
		Size:        h.Size,
		ContentType: h.Header.Get("Content-Type"),
	})
}
