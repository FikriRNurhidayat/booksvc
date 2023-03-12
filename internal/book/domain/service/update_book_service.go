package book_service

import (
	"context"

	book_model "github.com/fikrirnurhidayat/booksvc/internal/book/domain/model"
	book_repository "github.com/fikrirnurhidayat/booksvc/internal/book/domain/repository"
	image_repository "github.com/fikrirnurhidayat/booksvc/internal/image/domain/repository"
)

type UpdateBookService interface {
	Call(ctx context.Context, book book_model.Book) (book_model.Book, error)
}

type UpdateBookServiceImpl struct {
	bookRepository  book_repository.BookRepository
	imageRepository image_repository.ImageRepository
}

func (s *UpdateBookServiceImpl) Call(ctx context.Context, book book_model.Book) (book_model.Book, error) {
	var err error

	// Book's cover MUST exist
	if book.Cover, err = s.imageRepository.GetImage(ctx, book.Cover.ID); err != nil {
		return book_model.Book{}, err
	}

	// Book's thumbnail MUST exist
	if book.Thumbnail, err = s.imageRepository.GetImage(ctx, book.Thumbnail.ID); err != nil {
		return book_model.Book{}, err
	}

	// Given book display images MUST exist
	for i, displayImage := range book.DisplayImages {
		if book.DisplayImages[i], err = s.imageRepository.GetImage(ctx, displayImage.ID); err != nil {
			return book_model.Book{}, err
		}
	}

	if book, err = s.bookRepository.SaveBook(ctx, book); err != nil {
		return book_model.Book{}, err
	}

	return book, nil
}

func NewUpdateBookService(bookRepository book_repository.BookRepository, imageRepository image_repository.ImageRepository) UpdateBookService {
	return &UpdateBookServiceImpl{
		bookRepository:  bookRepository,
		imageRepository: imageRepository,
	}
}
