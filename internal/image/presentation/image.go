package image_presentation

import image_model "github.com/fikrirnurhidayat/booksvc/internal/image/domain/model"

type (
	Image struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Alt  string `json:"alt"`
		URL  string `json:"url"`
	}

	Images []Image
)

func NewImage(image image_model.Image) Image {
	return Image{
		ID:   image.ID.String(),
		Name: image.Name.String(),
		Alt:  image.Alt.String(),
		URL:  image.URL.String(),
	}
}

func NewImages(images image_model.Images) Images {
	results := Images{}

	for _, image := range images {
		results = append(results, NewImage(image))
	}

	return results
}
