package image_infrastructure_service

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

type ImageService interface {
	Upload(ctx context.Context, file *multipart.FileHeader) (url string, err error)
	Download(ctx context.Context, url string) (downloadUrl string, err error)
}

type LocalImageService struct {
	Location  string
	PrefixURL string
}

func (s *LocalImageService) Upload(ctx context.Context, file *multipart.FileHeader) (url string, err error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(path.Join(s.Location, file.Filename))
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return strings.Join([]string{s.PrefixURL, file.Filename}, "/"), nil
}

func (s *LocalImageService) Download(ctx context.Context, url string) (string, error) {
	// TODO: Check whether image exist on local storage
	fileName := strings.ReplaceAll(url, s.PrefixURL+"/", "")

	return path.Join(s.Location, fileName), nil
}

func NewLocalImageUploadService(location string, prefixURL string) ImageService {
	return &LocalImageService{
		Location:  location,
		PrefixURL: prefixURL,
	}
}
