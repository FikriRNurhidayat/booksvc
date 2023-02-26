package image_repository

import (
	"context"

	"github.com/fikrirnurhidayat/booksvc/internal/image/domain/model"
	"github.com/google/uuid"
)

type ImageRepository interface {
	CreateImage(ctx context.Context, image image_model.Image) error
	GetImage(ctx context.Context, id uuid.UUID) (image_model.Image, error)
	SearchImages(ctx context.Context) (image_model.Images, error)
	DeleteImage(ctx context.Context, id uuid.UUID) error
}
