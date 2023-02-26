package book_presentation

import book_model "github.com/fikrirnurhidayat/booksvc/internal/book/domain/model"

type (
	BookResponse struct {
		Book Book `json:"book"`
	}

	ListBooksResponse struct {
		Books Books `json:"books"`
	}
)

func NewBookResponse(book book_model.Book) BookResponse {
	return BookResponse{
		Book: NewBook(book),
	}
}

func NewListBooksResponse(books book_model.Books) ListBooksResponse {
	return ListBooksResponse{
		Books: NewBooks(books),
	}
}
