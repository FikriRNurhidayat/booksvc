package image_presentation

import (
	image_model "github.com/fikrirnurhidayat/booksvc/internal/image/domain/model"
)

type ImageResponse struct {
	Image Image `json:"image"`
}

type ListImagesResponse struct {
	Images Images `json:"images"`
}

func NewImageResponse(image image_model.Image) ImageResponse {
	return ImageResponse{
		Image: Image{
			ID:   image.ID.String(),
			Name: image.Name.String(),
			Alt:  image.Alt.String(),
			URL:  image.URL.String(),
		},
	}
}

func NewListImagesResponse(images image_model.Images) ListImagesResponse {
	return ListImagesResponse{
		Images: NewImages(images),
	}
}
