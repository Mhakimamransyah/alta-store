package productsimages

import "mime/multipart"

type Service interface {
	InsertNewImages(products_image ProductImages, files []*multipart.FileHeader, createdBy string) error
}

type Repository interface {
	CreateImages(products_images *ProductImages, files []*multipart.FileHeader, createdBy string) error
	GetListProductsImagesByIdProducts(id_products int) (*[]ProductImages, error)
}
