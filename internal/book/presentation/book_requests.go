package book_presentation

import (
	book_model "github.com/fikrirnurhidayat/booksvc/internal/book/domain/model"
	common_model "github.com/fikrirnurhidayat/booksvc/internal/common/domain/model"
	image_model "github.com/fikrirnurhidayat/booksvc/internal/image/domain/model"
	"github.com/google/uuid"
)

type (
	BookParams struct {
		Title           string   `json:"title"`
		ISBN            string   `json:"isbn"`
		CoverID         string   `json:"cover_id"`
		ThumbnailID     string   `json:"thumbnail_id"`
		DisplayImageIDs []string `json:"display_image_ids"`
	}

	BookRequest struct {
		Book BookParams `json:"book"`
	}
)

func (request BookRequest) ToEntity() book_model.Book {
	book := book_model.Book{}
	book.Title = common_model.ShortText(request.Book.Title)
	book.ISBN = book_model.ISBN(request.Book.ISBN)
	book.Cover = image_model.Image{
		ID: uuid.MustParse(request.Book.CoverID),
	}
	book.Thumbnail = image_model.Image{
		ID: uuid.MustParse(request.Book.ThumbnailID),
	}

	book.DisplayImages = image_model.Images{}

	for _, dii := range request.Book.DisplayImageIDs {
		book.DisplayImages = append(book.DisplayImages, image_model.Image{
			ID: uuid.MustParse(dii),
		})
	}

	return book
}
