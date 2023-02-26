package image_infrastructure_repository

import (
	"context"

	common_model "github.com/fikrirnurhidayat/booksvc/internal/common/domain/model"
	"github.com/fikrirnurhidayat/booksvc/internal/config"
	image_model "github.com/fikrirnurhidayat/booksvc/internal/image/domain/model"
	image_repository "github.com/fikrirnurhidayat/booksvc/internal/image/domain/repository"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const MongoImageCollection = "images"

type MongoImageRepository struct {
	db     *mongo.Collection
	logger echo.Logger
}

type ImageDocument struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
	Alt  string `bson:"alt"`
	URL  string `bson:"url"`
}

func NewImageDocument(image image_model.Image) ImageDocument {
	return ImageDocument{
		ID:   image.ID.String(),
		Name: image.Name.String(),
		Alt:  image.Alt.String(),
		URL:  image.URL.String(),
	}
}

func (doc ImageDocument) ToEntity() image_model.Image {
	return image_model.Image{
		ID:   uuid.MustParse(doc.ID),
		Name: common_model.ShortText(doc.Name),
		Alt:  common_model.ShortText(doc.Alt),
		URL:  common_model.URL(doc.URL),
	}
}

type ImageDocuments []ImageDocument

func (docs ImageDocuments) ToEntities() image_model.Images {
	var images image_model.Images

	for _, imageDocument := range docs {
		images = append(images, imageDocument.ToEntity())
	}

	return images
}

func NewImageDocuments(images image_model.Images) ImageDocuments {
	var imageDocuments ImageDocuments

	for _, image := range images {
		imageDocuments = append(imageDocuments, ImageDocument{
			ID:   image.ID.String(),
			Name: image.Name.String(),
			Alt:  image.Alt.String(),
			URL:  image.URL.String(),
		})
	}

	return imageDocuments
}

func (*MongoImageRepository) DeleteImage(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

func (r *MongoImageRepository) GetImage(ctx context.Context, id uuid.UUID) (image_model.Image, error) {
	var doc ImageDocument
	result := r.db.FindOne(ctx, bson.D{{"_id", id.String()}})

	if result.Err() != nil {
		return image_model.Image{}, result.Err()
	}

	if err := result.Decode(&doc); err != nil {
		return image_model.Image{}, err
	}

	return doc.ToEntity(), nil
}

func (r *MongoImageRepository) CreateImage(ctx context.Context, image image_model.Image) error {
	result, err := r.db.InsertOne(ctx, NewImageDocument(image))
	if err != nil {
		r.logger.Errorf("failed to insertOne: %s", err.Error())
	}

	r.logger.Debugf("success to insertOne: %s", result.InsertedID)

	return nil
}

func (*MongoImageRepository) SearchImages(ctx context.Context) (image_model.Images, error) {
	panic("unimplemented")
}

func NewMongoImageRepository(client *mongo.Client, logger echo.Logger) image_repository.ImageRepository {
	collection := client.Database(config.GetMongoDatabaseName()).Collection(MongoImageCollection)

	return &MongoImageRepository{
		db:     collection,
		logger: logger,
	}
}
