package s3

import (
	"context"
	"io"
	"time"
)

type S3UploadFileParameters struct {
	Ctx         context.Context
	Prefix      string
	Filename    string
	Reader      io.Reader
	Size        int64
	ContentType string
}

type S3GetFileParameters struct {
	Ctx      context.Context
	FileName string
	Expires  time.Duration
}
