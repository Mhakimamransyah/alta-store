package productsimages

import (
	"altaStore/business"
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
	for _, img := range files {
		err := service.repository.CreateImages(products_image, img, createdBy)
		if err != nil {
			continue
		}
	}

	return nil
}

func (service *service) RemoveProductsImages(id_admins, id_products, id_products_images int, deletedById int) error {
	products_images, err := service.repository.GetProductsImagesById(id_products_images)
	if err != nil {
		return business.ErrNotFound
	}

	// Check admin authority
	if id_admins != deletedById {
		return business.ErrUnauthorized
	}

	err = service.repository.DeleteProductsImages(products_images, deletedById)
	if err != nil {
		return business.ErrInternalServerError
	}

	return nil
}
