package storage

import (
	"context"
	"io"
)

// folder storage menjadi tempat semua implementasi penyimpanan file.

type Uploader interface {
	Upload(ctx context.Context, file io.Reader, fileName string) (UploadResult, error)
	Delete(ctx context.Context, fileID string) error
}

type UploadResult struct {
	URL    string
	FileID string
}
