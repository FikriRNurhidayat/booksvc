package book_infrastructure_repository

import (
	"context"

	"github.com/fikrirnurhidayat/booksvc/internal/book/domain/model"
	"github.com/fikrirnurhidayat/booksvc/internal/book/domain/repository"
	"github.com/fikrirnurhidayat/booksvc/internal/common/domain/model"
	"github.com/fikrirnurhidayat/booksvc/internal/config"
	"github.com/fikrirnurhidayat/booksvc/internal/image/infrastructure/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MongoBookCollection = "books"

type MongoBookRepository struct {
	db     *mongo.Collection
	logger echo.Logger
}

type BookDocument struct {
	ID            string                                         `bson:"_id"`
	Title         string                                         `bson:"title"`
	ISBN          string                                         `bson:"isbn"`
	Cover         image_infrastructure_repository.ImageDocument  `bson:"cover"`
	Thumbnail     image_infrastructure_repository.ImageDocument  `bson:"thumbnail"`
	DisplayImages image_infrastructure_repository.ImageDocuments `bson:"display_images"`
}

func NewBookDocument(book book_model.Book) *BookDocument {
	return &BookDocument{
		ID:            book.ID.String(),
		Title:         book.Title.String(),
		ISBN:          book.ISBN.String(),
		Cover:         image_infrastructure_repository.NewImageDocument(book.Cover),
		Thumbnail:     image_infrastructure_repository.NewImageDocument(book.Thumbnail),
		DisplayImages: image_infrastructure_repository.NewImageDocuments(book.DisplayImages),
	}
}

func (doc *BookDocument) ToEntity() book_model.Book {
	return book_model.Book{
		ID:            uuid.MustParse(doc.ID),
		Title:         common_model.ShortText(doc.Title),
		ISBN:          book_model.ISBN(doc.ISBN),
		Cover:         doc.Cover.ToEntity(),
		Thumbnail:     doc.Thumbnail.ToEntity(),
		DisplayImages: doc.DisplayImages.ToEntities(),
	}
}

type BookDocuments []BookDocument

func (docs BookDocuments) ToEntites() book_model.Books {
	books := book_model.Books{}

	for _, doc := range docs {
		books = append(books, doc.ToEntity())
	}

	return books
}

func (r *MongoBookRepository) DeleteBook(ctx context.Context, id uuid.UUID) error {
	opts := options.Delete().SetHint(bson.D{{"_id", 1}})

	if _, err := r.db.DeleteMany(ctx, bson.D{{"_id", id.String()}}, opts); err != nil {
		r.logger.Errorf("failed to delete document: %s", err.Error())
		return err
	}

	return nil
}

func (r *MongoBookRepository) GetBook(ctx context.Context, id uuid.UUID) (book_model.Book, error) {
	var doc BookDocument
	res := r.db.FindOne(ctx, bson.D{{"_id", id.String()}})
	err := res.Decode(&doc)
	if err != nil {
		r.logger.Errorf("failed to parse document: %s", err.Error())
		return book_model.Book{}, err
	}

	return doc.ToEntity(), nil
}

func (r *MongoBookRepository) SaveBook(ctx context.Context, book book_model.Book) (book_model.Book, error) {
	opts := options.Update().SetUpsert(true)
	result, err := r.db.UpdateByID(ctx, book.ID.String(), bson.D{{"$set", NewBookDocument(book)}}, opts)
	if err != nil {
		r.logger.Errorf("failed to execute updateByID: %s", err.Error())
		return book_model.Book{}, err
	}

	r.logger.Debugf("success to execute updateByID: %s", result.UpsertedID)

	return book, nil

	// result, err := r.db.InsertOne(ctx, NewBookDocument(book))

	// if err != nil {
	// 	r.logger.Errorf("failed to execute insertOne: %s", err.Error())
	// 	return book_model.Book{}, err
	// }

	// r.logger.Debugf("success to execute insertOne: %s", result.InsertedID)

	// return book, nil
}

func (r *MongoBookRepository) SearchBooks(ctx context.Context) (book_model.Books, error) {
	var docs BookDocuments

	result, err := r.db.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if err := result.All(ctx, &docs); err != nil {
		return nil, err
	}

	return docs.ToEntites(), nil
}

func NewMongoBookRepository(client *mongo.Client, logger echo.Logger) book_repository.BookRepository {
	collection := client.Database(config.GetMongoDatabaseName()).Collection(MongoBookCollection)

	return &MongoBookRepository{
		db:     collection,
		logger: logger,
	}
}
