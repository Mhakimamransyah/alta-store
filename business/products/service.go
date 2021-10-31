package products

import (
	"altaStore/business"
	productsimages "altaStore/business/products_images"
	"altaStore/util/validator"
)

type service struct {
	products_repository        Repository
	products_images_repository productsimages.Repository
}

type FilterProducts struct {
	CategoriesId int
	Query        string
	Sort         string
	SortPrice    string
	Price_max    int
	Price_min    int
	Page         int
	Per_page     int
	Status       string
}

type ProductsSpec struct {
	Title         string  `form:"title" validate:"required,max=100"`
	Price         int     `form:"price" validate:"required"`
	Description   string  `form:"description" validate:"max=200"`
	Weight        float64 `form:"weight"`
	Id_categories int     `form:"id_categories" validate:"required"`
	Stock         int     `form:"stock"`
}

type ProductsUpdatable struct {
	Title        string  `form:"title" validate:"required,max=100"`
	Price        int     `form:"price" validate:"required"`
	Description  string  `form:"description" validate:"max=200"`
	Stock        int     `form:"stock"`
	Status       string  `form:"status"`
	Weight       float64 `form:"weight"`
	CategoriesID int     `form:"id_categories" validate:"required"`
}

func InitProductsService(products_repo Repository, products_img_repo productsimages.Repository) *service {
	return &service{
		products_repository:        products_repo,
		products_images_repository: products_img_repo,
	}
}

func (service *service) FindAllProducts(filter FilterProducts) (*[]Products, error) {
	list_products, err := service.products_repository.GetAllProducts(filter)
	if err != nil {
		return nil, err
	}

	var list_products_with_products_images []Products
	for _, data := range *list_products {
		list_products_images, error := service.products_images_repository.GetListProductsImagesByIdProducts(data.ID)
		if error != nil {
			continue
		}
		data.Products_images = *list_products_images
		list_products_with_products_images = append(list_products_with_products_images, data)
	}

	return &list_products_with_products_images, nil
}

func (service *service) DetailProducts(id_products int) (*Products, error) {
	products, err := service.products_repository.GetDetailProducts(id_products)
	if err != nil {
		return nil, err
	}
	products_with_images, err := service.products_images_repository.GetListProductsImagesByIdProducts(products.ID)
	products.Products_images = *products_with_images
	return products, nil
}

func (service *service) InsertProducts(id_admin int, products_spec ProductsSpec, createdBy int) (*Products, error) {
	err := validator.GetValidator().Struct(&products_spec)
	if err != nil {
		return nil, err
	}
	products := NewProducts(id_admin, products_spec)
	products, err = service.products_repository.CreateProducts(products, createdBy)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (service *service) ModifyProducts(id_admin, id_products int, products_updatable ProductsUpdatable, modifiedBy int) error {
	err := validator.GetValidator().Struct(&products_updatable)
	if err != nil {
		return err
	}
	products, err := service.products_repository.GetDetailProducts(id_products)
	if err != nil {
		return err
	}

	// check if admin have authorized to edit product
	if products.AdminID != id_admin {
		return business.ErrUnauthorized
	}

	new_products := products.ModifyProducts(id_admin, products_updatable)

	err = service.products_repository.UpdateProducts(id_products, new_products, modifiedBy)
	if err != nil {
		return err
	}
	return nil
}

func (service *service) RemoveProducts(id_products int, deletedById int) error {
	products, err := service.products_repository.GetDetailProducts(id_products)
	if err != nil {
		return err
	}
	if deletedById != products.AdminID {
		return business.ErrUnauthorized
	}
	err = service.products_repository.DeleteProducts(products)
	if err != nil {
		return err
	}
	return nil
}
