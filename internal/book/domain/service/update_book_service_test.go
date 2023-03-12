package book_service_test

import (
	"context"
	"errors"
	"testing"

	book_model "github.com/fikrirnurhidayat/booksvc/internal/book/domain/model"
	book_service "github.com/fikrirnurhidayat/booksvc/internal/book/domain/service"
	image_model "github.com/fikrirnurhidayat/booksvc/internal/image/domain/model"
	mock_book_repository "github.com/fikrirnurhidayat/booksvc/internal/mocks/book/domain/repository"
	mock_image_repository "github.com/fikrirnurhidayat/booksvc/internal/mocks/image/domain/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type MockUpdateBookService struct {
	bookRepository  *mock_book_repository.BookRepository
	imageRepository *mock_image_repository.ImageRepository
}

func TestUpdateBookService(t *testing.T) {
	type input struct {
		ctx  context.Context
		book book_model.Book
	}

	type output struct {
		book book_model.Book
		err  error
	}

	type testcase struct {
		name   string
		in  *input
		out *output
		setup  func(*MockUpdateBookService, *input, *output)
	}

	// NOTE: Move this somewhere else, so it can be shared with other test suite
	fixtures := book_model.Books{
		{
			ID:    uuid.New(),
			Title: "The Gay Science",
			ISBN:  "9780394719856",
			Cover: image_model.Image{
				ID:   uuid.New(),
				Name: "The Gay Science Cover",
				Alt:  "The Gay Science Cover",
				URL:  "https://picsum.photos/200/300",
			},
			Thumbnail: image_model.Image{
				ID:   uuid.New(),
				Name: "The Gay Science Thumbnail",
				Alt:  "The Gay Science Thumbnail",
				URL:  "https://picsum.photos/200/300",
			},
			DisplayImages: []image_model.Image{
				{
					ID:   uuid.New(),
					Name: "The Gay Science Display Image 1",
					Alt:  "The Gay Science Display Image 1",
					URL:  "https://picsum.photos/200/300",
				},
				{
					ID:   uuid.New(),
					Name: "The Gay Science Display Image 2",
					Alt:  "The Gay Science Display Image 2",
					URL:  "https://picsum.photos/200/300",
				},
				{
					ID:   uuid.New(),
					Name: "The Gay Science Display Image 3",
					Alt:  "The Gay Science Display Image 3",
					URL:  "https://picsum.photos/200/300",
				},
			},
		},
	}

	for _, tt := range []testcase{
		{
			name: "cover doesn't exist",
			in: &input{
				ctx:  context.Background(),
				book: fixtures[0],
			},
			out: &output{
				book: book_model.Book{},
				err:  errors.New("image not found"),
			},
			setup: func(m *MockUpdateBookService, i *input, o *output) {
				m.imageRepository.On("GetImage", i.ctx, i.book.Cover.ID).Return(image_model.Image{}, o.err)
			},
		},
		{
			name: "thumbnail doesn't exist",
			in: &input{
				ctx:  context.Background(),
				book: fixtures[0],
			},
			out: &output{
				book: book_model.Book{},
				err:  errors.New("image not found"),
			},
			setup: func(m *MockUpdateBookService, i *input, o *output) {
				m.imageRepository.On("GetImage", i.ctx, i.book.Cover.ID).Return(o.book.Cover, nil)
				m.imageRepository.On("GetImage", i.ctx, i.book.Thumbnail.ID).Return(image_model.Image{}, o.err)
			},
		},
		{
			name: "display image doesn't exist",
			in: &input{
				ctx:  context.Background(),
				book: fixtures[0],
			},
			out: &output{
				book: book_model.Book{},
				err:  errors.New("image not found"),
			},
			setup: func(m *MockUpdateBookService, i *input, o *output) {
				m.imageRepository.On("GetImage", i.ctx, i.book.Cover.ID).Return(i.book.Cover, nil)
				m.imageRepository.On("GetImage", i.ctx, i.book.Thumbnail.ID).Return(i.book.Thumbnail, nil)

				for _, di := range(i.book.DisplayImages) {
					m.imageRepository.On("GetImage", i.ctx, di.ID).Return(image_model.Image{}, o.err)
				}
			},
		},
		{
			name: "failed to save book",
			in: &input{
				ctx:  context.Background(),
				book: fixtures[0],
			},
			out: &output{
				book: book_model.Book{},
				err:  errors.New("failed to save book"),
			},
			setup: func(m *MockUpdateBookService, i *input, o *output) {
				m.imageRepository.On("GetImage", i.ctx, i.book.Cover.ID).Return(i.book.Cover, nil)
				m.imageRepository.On("GetImage", i.ctx, i.book.Thumbnail.ID).Return(i.book.Thumbnail, nil)

				for _, di := range(i.book.DisplayImages) {
					m.imageRepository.On("GetImage", i.ctx, di.ID).Return(di, nil)
				}

				m.bookRepository.On("SaveBook", i.ctx, i.book).Return(book_model.Book{}, o.err)
			},
		},
		{
			name: "ok",
			in: &input{
				ctx:  context.Background(),
				book: fixtures[0],
			},
			out: &output{
				book: fixtures[0],
				err:  nil,
			},
			setup: func(m *MockUpdateBookService, i *input, o *output) {
				m.imageRepository.On("GetImage", i.ctx, i.book.Cover.ID).Return(i.book.Cover, nil)
				m.imageRepository.On("GetImage", i.ctx, i.book.Thumbnail.ID).Return(i.book.Thumbnail, nil)

				for _, di := range(i.book.DisplayImages) {
					m.imageRepository.On("GetImage", i.ctx, di.ID).Return(di, nil)
				}

				m.bookRepository.On("SaveBook", i.ctx, i.book).Return(o.book, nil)
			},
		},
	} {
		t.Run(tt.name, func (t *testing.T) {
			m := &MockUpdateBookService{
				bookRepository:  &mock_book_repository.BookRepository{},
				imageRepository: &mock_image_repository.ImageRepository{},
			}

			if (tt.setup != nil) {
				tt.setup(m, tt.in, tt.out)
			}

			updateBookService := book_service.NewUpdateBookService(m.bookRepository, m.imageRepository) 
			book, err := updateBookService.Call(tt.in.ctx, tt.in.book)

			assert.Equal(t, tt.out.book, book)
			assert.Equal(t, tt.out.err, err)
		})
	}
}
