package book

import (
	book_application "github.com/fikrirnurhidayat/booksvc/internal/book/application"
	book_repository "github.com/fikrirnurhidayat/booksvc/internal/book/domain/repository"
	book_service "github.com/fikrirnurhidayat/booksvc/internal/book/domain/service"
	book_infrastructure_repository "github.com/fikrirnurhidayat/booksvc/internal/book/infrastructure/repository"
	image_repository "github.com/fikrirnurhidayat/booksvc/internal/image/domain/repository"
	image_infrastructure_repository "github.com/fikrirnurhidayat/booksvc/internal/image/infrastructure/repository"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookModule struct {
	bookController    book_application.BookController
	bookRepository    book_repository.BookRepository
	imageRepository   image_repository.ImageRepository
	insertBookService book_service.InsertBookService
	updateBookService book_service.UpdateBookService
	logger            echo.Logger
}

func (m *BookModule) ApplyRoute(e *echo.Echo) {
	e.POST("/v1/books", m.bookController.CreateBook)
	e.GET("/v1/books", m.bookController.ListBooks)
	e.GET("/v1/books/:id", m.bookController.GetBook)
	e.DELETE("/v1/books/:id", m.bookController.DeleteBook)
	e.PUT("/v1/books/:id", m.bookController.UpdateBook)
}

func NewBookModule(client *mongo.Client, logger echo.Logger) *BookModule {
	module := &BookModule{}

	module.logger = logger
	module.bookRepository = book_infrastructure_repository.NewMongoBookRepository(client, module.logger)
	module.imageRepository = image_infrastructure_repository.NewMongoImageRepository(client, module.logger)
	module.insertBookService = book_service.NewInsertBookService(module.bookRepository, module.imageRepository)
	module.updateBookService = book_service.NewUpdateBookService(module.bookRepository, module.imageRepository)
	module.bookController = book_application.NewBookController(module.bookRepository, module.insertBookService, module.updateBookService)

	return module
}
