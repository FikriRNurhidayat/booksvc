package image_application

import (
	"net/http"

	common_model "github.com/fikrirnurhidayat/booksvc/internal/common/domain/model"
	image_model "github.com/fikrirnurhidayat/booksvc/internal/image/domain/model"
	image_repository "github.com/fikrirnurhidayat/booksvc/internal/image/domain/repository"
	image_infrastructure_service "github.com/fikrirnurhidayat/booksvc/internal/image/infrastructure/service"
	"github.com/fikrirnurhidayat/booksvc/internal/image/presentation"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ImageController interface {
	UploadImage(ctx echo.Context) error
	DownloadImage(ctx echo.Context) error
}

type ImageControllerImpl struct {
	imageRepository image_repository.ImageRepository
	imageService    image_infrastructure_service.ImageService
}

func (c *ImageControllerImpl) DownloadImage(ctx echo.Context) error {
	id := ctx.Param("id")
	image, err := c.imageRepository.GetImage(ctx.Request().Context(), uuid.MustParse(id))
	if err != nil {
		ctx.Logger().Error(err.Error())
		return err
	}

	path, err := c.imageService.Download(ctx.Request().Context(), image.URL.String())
	if err != nil {
		ctx.Logger().Error(err.Error())
		return err
	}

	ctx.Logger().Debug(path)

	return ctx.File(path)
}

func (c *ImageControllerImpl) UploadImage(ctx echo.Context) error {
	name := ctx.FormValue("image[name]")
	alt := ctx.FormValue("image[alt]")

	file, err := ctx.FormFile("image[file]")
	if err != nil {
		return err
	}

	url, err := c.imageService.Upload(ctx.Request().Context(), file)
	if err != nil {
		return err
	}

	image := image_model.NewImage()
	image.Name = common_model.ShortText(name)
	image.Alt = common_model.ShortText(alt)
	image.URL = common_model.URL(url)

	if err := c.imageRepository.CreateImage(ctx.Request().Context(), image); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, image_presentation.NewImageResponse(image))
}

func NewImageController(imageRepository image_repository.ImageRepository, imageService image_infrastructure_service.ImageService) ImageController {
	return &ImageControllerImpl{
		imageRepository: imageRepository,
		imageService:    imageService,
	}
}
