package book_application

import (
	"net/http"

	book_repository "github.com/fikrirnurhidayat/booksvc/internal/book/domain/repository"
	book_service "github.com/fikrirnurhidayat/booksvc/internal/book/domain/service"
	book_presentation "github.com/fikrirnurhidayat/booksvc/internal/book/presentation"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type BookController interface {
	CreateBook(echo.Context) error
	UpdateBook(echo.Context) error
	GetBook(echo.Context) error
	DeleteBook(echo.Context) error
	ListBooks(echo.Context) error
}

type BookControllerImpl struct {
	bookRepository    book_repository.BookRepository
	insertBookService book_service.InsertBookService
	updateBookService book_service.UpdateBookService
}

func (c *BookControllerImpl) UpdateBook(ctx echo.Context) (err error) {
	var request book_presentation.BookRequest

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return err
	}

	if err := ctx.Bind(&request); err != nil {
		return err
	}

	book := request.ToEntity()
	book.ID = id

	if book, err = c.updateBookService.Call(ctx.Request().Context(), book); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, book_presentation.NewBookResponse(book))
}

func (c *BookControllerImpl) DeleteBook(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return err
	}

	if err := c.bookRepository.DeleteBook(ctx.Request().Context(), id); err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *BookControllerImpl) CreateBook(ctx echo.Context) (err error) {
	var request book_presentation.BookRequest

	if err = ctx.Bind(&request); err != nil {
		ctx.Logger().Error(err)
		return err
	}

	book := request.ToEntity()

	if book, err = c.insertBookService.Call(ctx.Request().Context(), book); err != nil {
		ctx.Logger().Error(err)
		return err
	}

	response := book_presentation.NewBookResponse(book)

	return ctx.JSON(http.StatusCreated, response)
}

func (c *BookControllerImpl) GetBook(ctx echo.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return err
	}

	book, err := c.bookRepository.GetBook(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	response := book_presentation.NewBookResponse(book)

	return ctx.JSON(http.StatusOK, response)
}

func (c *BookControllerImpl) ListBooks(ctx echo.Context) error {
	books, err := c.bookRepository.SearchBooks(ctx.Request().Context())
	if err != nil {
		return err
	}

	for _, book := range books {
		ctx.Logger().Error(book.ID)
	}

	return ctx.JSON(http.StatusOK, book_presentation.NewListBooksResponse(books))
}

func NewBookController(bookRepository book_repository.BookRepository, insertBookService book_service.InsertBookService, updateBookService book_service.UpdateBookService) BookController {
	return &BookControllerImpl{
		bookRepository:    bookRepository,
		insertBookService: insertBookService,
		updateBookService: updateBookService,
	}
}
