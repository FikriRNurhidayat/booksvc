package image

import (
	"github.com/fikrirnurhidayat/booksvc/internal/config"
	image_application "github.com/fikrirnurhidayat/booksvc/internal/image/application"
	image_repository "github.com/fikrirnurhidayat/booksvc/internal/image/domain/repository"
	image_infrastructure_repository "github.com/fikrirnurhidayat/booksvc/internal/image/infrastructure/repository"
	image_infrastructure_service "github.com/fikrirnurhidayat/booksvc/internal/image/infrastructure/service"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type ImageModule struct {
	logger             echo.Logger
	imageUploadService image_infrastructure_service.ImageService
	imageRepository    image_repository.ImageRepository
	imageController    image_application.ImageController
}

func (m *ImageModule) ApplyRoute(e *echo.Echo) {
	e.POST("/v1/images", m.imageController.UploadImage)
	e.GET("/v1/images/:id", m.imageController.DownloadImage)
}

func NewImageModule(client *mongo.Client, logger echo.Logger) *ImageModule {
	module := &ImageModule{}

	module.logger = logger
	module.imageUploadService = image_infrastructure_service.NewLocalImageUploadService(config.GetLocalStorageDirectory(), config.GetPublicURL())
	module.imageRepository = image_infrastructure_repository.NewMongoImageRepository(client, module.logger)
	module.imageController = image_application.NewImageController(module.imageRepository, module.imageUploadService)

	return module
}
