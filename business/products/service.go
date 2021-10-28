package products

import (
	"altaStore/business"
	"altaStore/util/validator"
	"fmt"
)

type service struct {
	repository Repository
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

func InitProductsService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (service *service) FindAllProducts(limit, offset int) (*[]Products, error) {
	list_products, err := service.repository.GetAllProducts(limit, offset)
	if err != nil {
		return nil, err
	}
	return list_products, nil
}

func (service *service) DetailProducts(id_products int) (*Products, error) {
	fmt.Println("***")
	fmt.Println(id_products)
	products, err := service.repository.GetDetailProducts(id_products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (service *service) InsertProducts(id_admin int, products_spec ProductsSpec, createdBy string) error {
	err := validator.GetValidator().Struct(&products_spec)
	if err != nil {
		return err
	}
	products := NewProducts(id_admin, products_spec)
	err = service.repository.CreateProducts(products, createdBy)
	if err != nil {
		return err
	}
	return nil
}

func (service *service) ModifyProducts(id_admin, id_products int, products_updatable ProductsUpdatable, modifiedBy string) error {
	err := validator.GetValidator().Struct(&products_updatable)
	if err != nil {
		return err
	}
	products, err := service.repository.GetDetailProducts(id_products)
	if err != nil {
		return err
	}

	// check if admin have authorized to edit product
	if products.AdminID != id_admin {
		return business.ErrUnauthorized
	}

	new_products := products.ModifyProducts(id_admin, products_updatable)

	err = service.repository.UpdateProducts(id_products, new_products, modifiedBy)
	if err != nil {
		return err
	}
	return nil
}
