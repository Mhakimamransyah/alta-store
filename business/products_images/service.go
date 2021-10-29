package productsimages

import (
	"mime/multipart"
)

type service struct {
	repository Repository
}

func InitProductsImagesService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (service *service) InsertNewImages(products_image *ProductImages, files []*multipart.FileHeader, createdBy int) error {
	err := service.repository.CreateImages(products_image, files, createdBy)
	if err != nil {
		return err
	}
	return nil
}
