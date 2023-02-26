package book_repository

import (
	"context"

	"github.com/fikrirnurhidayat/booksvc/internal/book/domain/model"
	"github.com/google/uuid"
)

type BookRepository interface {
	SearchBooks(ctx context.Context) (book_model.Books, error)
	SaveBook(ctx context.Context, book book_model.Book) (book_model.Book, error)
	GetBook(ctx context.Context, id uuid.UUID) (book_model.Book, error)
	DeleteBook(ctx context.Context, id uuid.UUID) error
}
