package storage

import (
	"GO-AUTH-JWT/exception"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/imagekit-developer/imagekit-go/v2"
	"github.com/imagekit-developer/imagekit-go/v2/option"
)

type ImagekitUploader struct {
	client *imagekit.Client
}

// Initial imagekit client
func InitImageKit() *imagekit.Client {
	client := imagekit.NewClient( // ? imagekit.NewClient(opts ...option.RequestOption)(r imagekit.Client) => untuk membuat client imagekit,
		option.WithPrivateKey(os.Getenv("IMAGEKIT_PRIVATE_KEY")), // ? option.WithPrivateKey(privateKey string) => untuk mengatur private key
	)
	return &client
}

func NewImagekitUploader(client *imagekit.Client) Uploader {
	return &ImagekitUploader{client: client}
}

func (uploader *ImagekitUploader) Upload(ctx context.Context, file io.Reader, fileName string) (UploadResult, error) {
	// upload image to imagekit
	responseUpload, err := uploader.client.Files.Upload(ctx, imagekit.FileUploadParams{
		File:     file,
		FileName: fileName,
	})
	if err != nil {
		panic(exception.NewCustomInternalServerError("Gagal upload gambar Imagekit"))
	}

	return UploadResult{URL: responseUpload.URL, FileIDImageKit: responseUpload.FileID}, nil
}

func (uploader *ImagekitUploader) Delete(ctx context.Context, fileID string) error {
	// delete image from imagekit
	err := uploader.client.Files.Delete(ctx, fileID)
	if err != nil {
		fmt.Println(err)
		panic(exception.NewCustomInternalServerError("Gagal menghapus gambar Imagekit"))
	}
	fmt.Println("Gambar berhasil dihapus")
	return nil
}
