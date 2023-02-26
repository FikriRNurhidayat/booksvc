package book_presentation

import (
	book_model "github.com/fikrirnurhidayat/booksvc/internal/book/domain/model"
	"github.com/fikrirnurhidayat/booksvc/internal/image/presentation"
)

type (
	Book struct {
		ID            string                    `json:"id"`
		Title         string                    `json:"title"`
		ISBN          string                    `json:"isbn"`
		Cover         image_presentation.Image  `json:"cover"`
		Thumbnail     image_presentation.Image  `json:"thumbnail"`
		DisplayImages image_presentation.Images `json:"display_images"`
	}

	Books []Book
)

func NewBook(book book_model.Book) Book {
	return Book{
		ID:            book.ID.String(),
		Title:         book.Title.String(),
		ISBN:          book.ISBN.String(),
		Cover:         image_presentation.NewImage(book.Cover),
		Thumbnail:     image_presentation.NewImage(book.Thumbnail),
		DisplayImages: image_presentation.NewImages(book.DisplayImages),
	}
}

func NewBooks(books book_model.Books) Books {
	result := Books{}

	for _, book := range books {
		result = append(result, NewBook(book))
	}

	return result
}
