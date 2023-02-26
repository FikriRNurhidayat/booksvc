package book_model

import (
	common_model "github.com/fikrirnurhidayat/booksvc/internal/common/domain/model"
	image_model "github.com/fikrirnurhidayat/booksvc/internal/image/domain/model"
	"github.com/google/uuid"
)

type Book struct {
	ID            uuid.UUID
	Title         common_model.ShortText
	ISBN          ISBN
	Cover         image_model.Image
	Thumbnail     image_model.Image
	DisplayImages image_model.Images
}

type Books []Book

type BookConstructor func(Book)

func NewBook() Book {
	return Book{
		ID:            uuid.New(),
		Title:         "",
		ISBN:          "",
		Cover:         image_model.Image{},
		Thumbnail:     image_model.Image{},
		DisplayImages: image_model.Images{},
	}
}
