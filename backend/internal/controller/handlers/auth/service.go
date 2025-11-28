package auth

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/goawwer/devclash/pkg/s3"
)

func saveLogoAtServer(ctx context.Context, f multipart.File, h *multipart.FileHeader, orgName string) (string, error) {
	defer f.Close()

	filename := fmt.Sprintf("%s-%d%s", orgName, time.Now().Unix(), filepath.Ext(h.Filename))

	return s3.Upload(&s3.S3UploadFileParameters{
		Ctx:         ctx,
		Prefix:      "logos",
		Filename:    filename,
		Reader:      f,
		Size:        h.Size,
		ContentType: h.Header.Get("Content-Type"),
	})
}
