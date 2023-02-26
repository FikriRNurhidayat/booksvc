package image_model

import (
	common_model "github.com/fikrirnurhidayat/booksvc/internal/common/domain/model"
	"github.com/google/uuid"
)

type Image struct {
	ID   uuid.UUID
	Name common_model.ShortText
	Alt  common_model.ShortText
	URL  common_model.URL
}

type Images []Image

func NewImage() Image {
	return Image{
		ID:   uuid.New(),
		Name: "",
		Alt:  "",
		URL:  "",
	}
}
