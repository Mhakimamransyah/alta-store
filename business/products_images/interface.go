package productsimages

import "mime/multipart"

type Service interface {
	InsertNewImages(products_image *ProductImages, files []*multipart.FileHeader, createdById int) error
}

type Repository interface {
	CreateImages(products_images *ProductImages, files []*multipart.FileHeader, createdById int) error
	GetListProductsImagesByIdProducts(id_products int) (*[]ProductImages, error)
	DeleteProductsImages(products_images *ProductImages, deletedById int) error
	GetProductsImagesById(id_products_image int) (*ProductImages, error)
}
